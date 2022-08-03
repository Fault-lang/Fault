package smt

import (
	"bytes"
	"fmt"

	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/value"
)

type rule interface {
	ruleNode()
	String() string
}

type assrt struct {
	rule
	variable       *wrap
	conjunction    string
	assertion      rule
	tag            *branch
	temporalFilter string
	temporalN      int
}

func (a *assrt) ruleNode() {}
func (a *assrt) String() string {
	return a.variable.String() + a.conjunction + a.assertion.String()
}
func (a *assrt) Tag(k1 string, k2 string) {
	a.tag = &branch{
		branch: k1,
		block:  k2,
	}
}

type infix struct {
	rule
	x           rule
	y           rule
	ty          string
	op          string
	declareOnly bool //For solvables which need to be declared but have no starting value
	tag         *branch
}

func (i *infix) ruleNode() {}
func (i *infix) String() string {
	return fmt.Sprintf("%s %s %s", i.x.String(), i.op, i.y.String())
}
func (i *infix) Tag(k1 string, k2 string) {
	i.tag = &branch{
		branch: k1,
		block:  k2,
	}
}

type ite struct {
	rule
	cond  rule
	t     []rule
	tvars map[string]string
	f     []rule
	fvars map[string]string
	tag   *branch
}

func (it *ite) ruleNode() {}
func (it *ite) String() string {
	return fmt.Sprintf("if %s then %s else %s", it.cond.String(), it.t, it.f)
}
func (ite *ite) Tag(k1 string, k2 string) {
	ite.tag = &branch{
		branch: k1,
		block:  k2,
	}
}

type invariant struct {
	rule
	left     rule
	operator string
	right    rule
	tag      *branch
}

func (i *invariant) ruleNode() {}
func (i *invariant) String() string {
	return fmt.Sprint(i.left.String(), i.operator, i.right.String())
}
func (i *invariant) Tag(k1 string, k2 string) {
	i.tag = &branch{
		branch: k1,
		block:  k2,
	}
}

type wrap struct { //wrapper for constant values to be used in infix as rules
	rule
	value    string
	state    string //invariant only for one state
	all      bool   // invariant for all states
	constant bool   // this is a constant
	tag      *branch
}

func (w *wrap) ruleNode() {}
func (w *wrap) String() string {
	return w.value
}
func (w *wrap) Tag(k1 string, k2 string) {
	w.tag = &branch{
		branch: k1,
		block:  k2,
	}
}

type wrapGroup struct {
	rule
	wraps []*wrap
	tag   *branch
}

func (wg *wrapGroup) ruleNode() {}
func (wg *wrapGroup) String() string {
	var out bytes.Buffer
	for _, v := range wg.wraps {
		out.WriteString(v.value)
	}
	return out.String()
}
func (wg *wrapGroup) Tag(k1 string, k2 string) {
	wg.tag = &branch{
		branch: k1,
		block:  k2,
	}
}

type vwrap struct {
	rule
	value value.Value
	tag   *branch
}

func (vw *vwrap) ruleNode() {}
func (vw *vwrap) String() string {
	return vw.value.String()
}
func (vw *vwrap) Tag(k1 string, k2 string) {
	vw.tag = &branch{
		branch: k1,
		block:  k2,
	}
}

type branch struct {
	branch string
	block  string
}

func (b *branch) String() string {
	return b.branch + "." + b.block
}

////////////////////////////////////
// General rule store and load logic
///////////////////////////////////

func (g *Generator) constantRule(id string, c constant.Constant) string {
	switch val := c.(type) {
	case *constant.Float:
		ty := g.variables.lookupType(id, val)
		id = g.variables.advanceSSA(id)
		if g.isASolvable(id) {
			g.declareVar(id, ty)
		} else {
			return g.writeInitRule(id, ty, val.String())
		}
	}
	return ""
}

func (g *Generator) loadsRule(inst *ir.InstLoad) {
	id := inst.Ident()
	g.variables.loads[id] = inst.Src
}

func (g *Generator) storeRule(inst *ir.InstStore, rules []rule) []rule {
	id := g.variables.formatIdent(inst.Dst.Ident())
	if g.variables.isTemp(inst.Src.Ident()) {
		srcId := inst.Src.Ident()
		if val, ok := g.variables.loads[srcId]; ok {
			ty := g.variables.lookupType(srcId, val)
			n := g.variables.ssa[id]
			if !g.inPhiState {
				g.variables.storeLastState(id, n+1)
			}
			id = g.variables.advanceSSA(id)
			rules = append(rules, g.parseRule(id, g.variables.formatValue(val), ty, ""))
		} else if ref, ok := g.variables.ref[srcId]; ok {
			switch r := ref.(type) {
			case *infix:
				r.x = g.tempToIdent(r.x)
				r.y = g.tempToIdent(r.y)
				n := g.variables.ssa[id]
				if !g.inPhiState {
					g.variables.storeLastState(id, n+1)
				}
				id = g.variables.advanceSSA(id)
				//g.trackRounds(id, inst)
				wid := &wrap{value: id}
				if g.variables.isBolean(r.y.String()) {
					rules = append(rules, &infix{x: wid, ty: "Bool", y: r, declareOnly: true}) // Still need to declare the new state
					rules = append(rules, &infix{x: wid, ty: "Bool", y: r, op: "="})
				} else if g.isASolvable(r.x.String()) {
					rules = append(rules, &infix{x: wid, ty: "Real", y: r, declareOnly: true})
					rules = append(rules, &infix{x: wid, ty: "Real", y: r, op: "="})
				} else {
					rules = append(rules, &infix{x: wid, ty: "Real", y: r})
				}
			default:
				n := g.variables.ssa[id]
				if !g.inPhiState {
					g.variables.storeLastState(id, n+1)
				}
				ty := g.variables.lookupType(id, nil)
				id = g.variables.advanceSSA(id)
				wid := &wrap{value: id}
				rules = append(rules, &infix{x: wid, ty: ty, y: r})
			}
		} else {
			panic(fmt.Sprintf("smt generation error, value for %s not found", id))
		}
	} else {
		ty := g.variables.lookupType(id, inst.Src)
		n := g.variables.ssa[id]
		if !g.inPhiState {
			g.variables.storeLastState(id, n+1)
		}
		id = g.variables.advanceSSA(id)
		rules = append(rules, g.parseRule(id, inst.Src.Ident(), ty, ""))
	}
	return rules
}

func (g *Generator) callRule(inst *ir.InstCall) string {
	callee := inst.Callee.Ident()
	meta := inst.Metadata
	g.parallelMeta(g.parallelGrouping, meta)
	return callee
}

func (g *Generator) tempRule(inst value.Value, r rule) {
	// If infix rule is stored in a temp variable
	id := inst.Ident()
	if g.variables.isTemp(id) {
		g.variables.ref[id] = r
	}
}
