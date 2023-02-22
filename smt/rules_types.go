package smt

import (
	"fault/smt/rules"
	"fmt"
	"strings"

	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/value"
)

////////////////////////////////////
// General rule store and load logic
///////////////////////////////////

func (g *Generator) constantRule(id string, c constant.Constant) string {
	switch val := c.(type) {
	case *constant.Float:
		ty := g.variables.LookupType(id, val)
		id = g.variables.AdvanceSSA(id)
		g.addVarToRound(id, int(g.variables.SSA[id]))
		if g.isASolvable(id) {
			g.declareVar(id, ty)
		} else {
			v := val.X.String()
			if strings.Contains(v, ".") {
				return g.writeInitRule(id, ty, v)
			}
			return g.writeInitRule(id, ty, v+".0")
		}
	}
	return ""
}

func (g *Generator) loadsRule(inst *ir.InstLoad) {
	id := inst.Ident()
	refname := fmt.Sprintf("%s-%s", g.currentFunction, id)
	g.variables.Loads[refname] = inst.Src
}

func (g *Generator) storeRule(inst *ir.InstStore) []rules.Rule {
	var ru []rules.Rule
	base := g.variables.FormatIdent(inst.Dst.Ident())
	if g.variables.IsTemp(inst.Src.Ident()) {
		srcId := inst.Src.Ident()
		refname := fmt.Sprintf("%s-%s", g.currentFunction, srcId)
		if val, ok := g.variables.Loads[refname]; ok {
			ty := g.variables.LookupType(refname, val)
			n := g.variables.SSA[base]
			prev := fmt.Sprintf("%s_%d", base, n)
			if !g.inPhiState.Check() {
				g.variables.NewPhi(base, n+1)
			} else {
				g.variables.StoreLastState(base, n+1)
			}
			id := g.variables.AdvanceSSA(base)
			g.addVarToRound(base, int(n+1))
			v := g.variables.FormatValue(val)
			if !g.variables.IsBoolean(v) && !g.variables.IsNumeric(v) {
				v = g.variables.FormatIdent(v)
				v = fmt.Sprintf("%s_%d", v, n)
			}
			g.AddNewVarChange(base, id, prev)
			ru = append(ru, g.createRule(id, v, ty, ""))
		} else if ref, ok := g.variables.Ref[refname]; ok {
			switch r := ref.(type) {
			case *rules.Infix:
				r.X = g.tempToIdent(r.X)
				r.Y = g.tempToIdent(r.Y)
				n := g.variables.SSA[base]
				prev := fmt.Sprintf("%s_%d", base, n)
				if !g.inPhiState.Check() {
					g.variables.NewPhi(base, n+1)
				} else {
					g.variables.StoreLastState(base, n+1)
				}
				id := g.variables.AdvanceSSA(base)
				g.addVarToRound(base, int(n+1))
				g.AddNewVarChange(base, id, prev)
				wid := &rules.Wrap{Value: id}
				if g.variables.IsBoolean(r.Y.String()) {
					ru = append(ru, &rules.Infix{X: wid, Ty: "Bool", Y: r, Op: "="})
				} else if g.isASolvable(r.X.String()) {
					ru = append(ru, &rules.Infix{X: wid, Ty: "Real", Y: r, Op: "="})
				} else {
					ru = append(ru, &rules.Infix{X: wid, Ty: "Real", Y: r})
				}
			default:
				n := g.variables.SSA[base]
				prev := fmt.Sprintf("%s_%d", base, n)
				if !g.inPhiState.Check() {
					g.variables.NewPhi(base, n+1)
				} else {
					g.variables.StoreLastState(base, n+1)
				}
				ty := g.variables.LookupType(base, nil)
				id := g.variables.AdvanceSSA(base)
				g.addVarToRound(base, int(n+1))
				g.AddNewVarChange(base, id, prev)
				wid := &rules.Wrap{Value: id}
				ru = append(ru, &rules.Infix{X: wid, Ty: ty, Y: r})
			}
		} else {
			panic(fmt.Sprintf("smt generation error, value for %s not found", base))
		}
	} else {
		ty := g.variables.LookupType(base, inst.Src)
		n := g.variables.SSA[base]
		prev := fmt.Sprintf("%s_%d", base, n)
		if !g.inPhiState.Check() {
			g.variables.NewPhi(base, n+1)
		} else {
			g.variables.StoreLastState(base, n+1)
		}
		id := g.variables.AdvanceSSA(base)
		g.addVarToRound(base, int(g.variables.SSA[base]))
		g.AddNewVarChange(base, id, prev)
		ru = append(ru, g.createRule(id, inst.Src.Ident(), ty, ""))
	}
	return ru
}

func (g *Generator) xorRule(inst *ir.InstXor) rules.Rule {
	id := inst.Ident()
	x := inst.X.Ident()
	xRule := g.variables.LookupCondPart(g.currentFunction, x)
	if xRule == nil {
		x = g.variables.ConvertIdent(g.currentFunction, x)
		xRule = &rules.Wrap{Value: x}
	}
	return g.createMultiCondRule(id, xRule, &rules.Wrap{}, "not")
}

func (g *Generator) andRule(inst *ir.InstAnd) rules.Rule {
	id := inst.Ident()
	x := inst.X.Ident()
	y := inst.Y.Ident()

	xRule := g.variables.LookupCondPart(g.currentFunction, x)
	if xRule == nil {
		x = g.variables.ConvertIdent(g.currentFunction, x)
		xRule = &rules.Wrap{Value: x}
	}

	yRule := g.variables.LookupCondPart(g.currentFunction, y)
	if yRule == nil {
		y = g.variables.ConvertIdent(g.currentFunction, y)
		yRule = &rules.Wrap{Value: y}
	}
	return g.createMultiCondRule(id, xRule, yRule, "and")
}

func (g *Generator) orRule(inst *ir.InstOr) rules.Rule {
	x := inst.X.Ident()
	y := inst.Y.Ident()
	id := inst.Ident()
	xRule := g.variables.LookupCondPart(g.currentFunction, x)
	if xRule == nil {
		x = g.variables.ConvertIdent(g.currentFunction, x)
		xRule = &rules.Wrap{Value: x}
	}

	yRule := g.variables.LookupCondPart(g.currentFunction, y)
	if yRule == nil {
		y = g.variables.ConvertIdent(g.currentFunction, y)
		yRule = &rules.Wrap{Value: y}
	}
	return g.createMultiCondRule(id, xRule, yRule, "or")
}

func (g *Generator) stateRules(key string, sc *rules.StateChange) rules.Rule {
	if len(sc.Ors) == 0 {
		and := g.andStateRule(key, sc.Ands)
		a := &rules.Ands{
			X: and,
		}

		c := &rules.Choices{
			X:  []*rules.Ands{a},
			Op: "and",
		}
		return c
	}

	and := g.andStateRule(key, sc.Ands)
	ors := g.orStateRule(key, sc.Ors)

	if len(sc.Ands) != 0 {
		ors["joined_ands"] = and
	}

	x := g.syncStateRules(ors)

	r := &rules.Choices{
		X:  x,
		Op: "or",
	}

	return r

}

func (g *Generator) orStateRule(choiceK string, choiceV []value.Value) map[string][]rules.Rule {
	g.inPhiState.In()

	and := make(map[string][]rules.Rule)
	for _, b := range choiceV {
		refname := fmt.Sprintf("%s-%s", g.currentFunction, b.Ident())
		and[refname] = g.parseBuiltIn(b.(*ir.InstCall), true)
	}
	delete(g.storedChoice, choiceK)

	g.inPhiState.Out()
	return and
}

func (g *Generator) andStateRule(andK string, andV []value.Value) []rules.Rule {
	g.inPhiState.In()

	var ands []rules.Rule
	for _, b := range andV {
		a := g.parseBuiltIn(b.(*ir.InstCall), true)
		ands = append(ands, a...)
	}
	delete(g.storedChoice, andK)

	g.inPhiState.Out()
	return ands
}

func (g *Generator) syncStateRules(branches map[string][]rules.Rule) []*rules.Ands {
	g.inPhiState.In()
	g.newFork()

	var e []rules.Rule
	var keys []string
	ends := make(map[string][]rules.Rule)
	phis := make(map[string]int16)
	var x []*rules.Ands

	for k, v := range branches {
		keys = append(keys, k)
		g.buildForkChoice(v, k)
		e, phis = g.capCond(k, phis)
		ends[k] = append(v, e...)
	}

	syncs := g.capCondSyncRules(keys)
	for k, v := range syncs {
		e2 := append(ends[k], v...)
		a := &rules.Ands{
			X: e2,
		}
		x = append(x, a)
	}
	g.inPhiState.Out()
	return x
}

func (g *Generator) tempRule(inst value.Value, r rules.Rule) {
	// If infix rule is stored in a temp variable
	id := inst.Ident()
	if g.variables.IsTemp(id) {
		refname := fmt.Sprintf("%s-%s", g.currentFunction, id)
		g.variables.Ref[refname] = r
	}
}
