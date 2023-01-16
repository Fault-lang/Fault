package smt

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/value"
)

type rule interface {
	ruleNode()
	String() string
}

type ands struct {
	rule
	x   []rule
	tag *branch
}

func (a *ands) ruleNode() {}
func (a *ands) String() string {
	var out bytes.Buffer
	for _, r := range a.x {
		out.WriteString(r.String())
	}
	return out.String()
}
func (a *ands) Tag(k1 string, k2 string) {
	a.tag = &branch{
		branch: k1,
		block:  k2,
	}
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

type choices struct {
	rule
	x   []*ands
	op  string
	tag *branch
}

func (c *choices) ruleNode() {}
func (c *choices) String() string {
	var out bytes.Buffer
	for i, ru := range c.x {
		out.WriteString(fmt.Sprintf("branch-%d: ", i))
		for _, r := range ru.x {
			out.WriteString(r.String())
		}
	}
	return out.String()
}
func (c *choices) Tag(k1 string, k2 string) {
	c.tag = &branch{
		branch: k1,
		block:  k2,
	}
}

type infix struct {
	rule
	x   rule
	y   rule
	ty  string
	op  string
	tag *branch
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
	cond rule
	t    []rule
	f    []rule
	tag  *branch
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

type phi struct {
	baseVar  string
	nums     []int16
	endState string
	tag      *branch
}

func (p *phi) ruleNode() {}
func (p *phi) String() string {
	var out bytes.Buffer
	for _, n := range p.nums {
		r := fmt.Sprintf("%s = %s_%d || ", p.endState, p.baseVar, n)
		out.WriteString(r)
	}
	return out.String()
}
func (p *phi) Tag(k1 string, k2 string) {
	p.tag = &branch{
		branch: k1,
		block:  k2,
	}
}

type stateChange struct {
	rule
	ands []value.Value
	ors  []value.Value
	tag  *branch
}

func (sc *stateChange) ruleNode() {}
func (sc *stateChange) String() string {
	var out bytes.Buffer
	for _, n := range sc.ands {
		r := fmt.Sprintf("and %s ", n)
		out.WriteString(r)
	}
	for _, n := range sc.ors {
		r := fmt.Sprintf("or %s ", n)
		out.WriteString(r)
	}
	return out.String()
}
func (sc *stateChange) Tag(k1 string, k2 string) {
	sc.tag = &branch{
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
	g.variables.loads[refname] = inst.Src
}

func (g *Generator) storeRule(inst *ir.InstStore) []rule {
	var rules []rule
	base := g.variables.formatIdent(inst.Dst.Ident())
	if g.variables.isTemp(inst.Src.Ident()) {
		srcId := inst.Src.Ident()
		refname := fmt.Sprintf("%s-%s", g.currentFunction, srcId)
		if val, ok := g.variables.loads[refname]; ok {
			ty := g.variables.lookupType(refname, val)
			n := g.variables.ssa[base]
			prev := fmt.Sprintf("%s_%d", base, n)
			if !g.inPhiState.Check() {
				g.variables.newPhi(base, n+1)
			} else {
				g.variables.storeLastState(base, n+1)
			}
			id := g.variables.advanceSSA(base)
			v := g.variables.formatValue(val)
			if !g.variables.isBolean(v) && !g.variables.isNumeric(v) {
				v = g.variables.formatIdent(v)
				v = fmt.Sprintf("%s_%d", v, n)
			}
			g.AddNewVarChange(base, id, prev)
			rules = append(rules, g.parseRule(id, v, ty, ""))
		} else if ref, ok := g.variables.ref[refname]; ok {
			switch r := ref.(type) {
			case *infix:
				r.x = g.tempToIdent(r.x)
				r.y = g.tempToIdent(r.y)
				n := g.variables.ssa[base]
				prev := fmt.Sprintf("%s_%d", base, n)
				if !g.inPhiState.Check() {
					g.variables.newPhi(base, n+1)
				} else {
					g.variables.storeLastState(base, n+1)
				}
				id := g.variables.advanceSSA(base)
				g.AddNewVarChange(base, id, prev)
				wid := &wrap{value: id}
				if g.variables.isBolean(r.y.String()) {
					rules = append(rules, &infix{x: wid, ty: "Bool", y: r, op: "="})
				} else if g.isASolvable(r.x.String()) {
					rules = append(rules, &infix{x: wid, ty: "Real", y: r, op: "="})
				} else {
					rules = append(rules, &infix{x: wid, ty: "Real", y: r})
				}
			default:
				n := g.variables.ssa[base]
				prev := fmt.Sprintf("%s_%d", base, n)
				if !g.inPhiState.Check() {
					g.variables.newPhi(base, n+1)
				} else {
					g.variables.storeLastState(base, n+1)
				}
				ty := g.variables.lookupType(base, nil)
				id := g.variables.advanceSSA(base)
				g.AddNewVarChange(base, id, prev)
				wid := &wrap{value: id}
				rules = append(rules, &infix{x: wid, ty: ty, y: r})
			}
		} else {
			panic(fmt.Sprintf("smt generation error, value for %s not found", base))
		}
	} else {
		ty := g.variables.lookupType(base, inst.Src)
		n := g.variables.ssa[base]
		prev := fmt.Sprintf("%s_%d", base, n)
		if !g.inPhiState.Check() {
			g.variables.newPhi(base, n+1)
		} else {
			g.variables.storeLastState(base, n+1)
		}
		id := g.variables.advanceSSA(base)
		g.AddNewVarChange(base, id, prev)
		rules = append(rules, g.parseRule(id, inst.Src.Ident(), ty, ""))
	}
	return rules
}

func (g *Generator) xorRule(inst *ir.InstXor) rule {
	id := inst.Ident()
	x := inst.X.Ident()
	xRule := g.variables.lookupCondPart(g.currentFunction, x)
	if xRule == nil {
		x = g.variables.convertIdent(g.currentFunction, x)
		xRule = &wrap{value: x}
	}
	return g.parseMultiCond(id, xRule, &wrap{}, "not")
}

func (g *Generator) andRule(inst *ir.InstAnd) rule {
	id := inst.Ident()
	x := inst.X.Ident()
	y := inst.Y.Ident()

	xRule := g.variables.lookupCondPart(g.currentFunction, x)
	if xRule == nil {
		x = g.variables.convertIdent(g.currentFunction, x)
		xRule = &wrap{value: x}
	}

	yRule := g.variables.lookupCondPart(g.currentFunction, y)
	if yRule == nil {
		y = g.variables.convertIdent(g.currentFunction, y)
		yRule = &wrap{value: y}
	}
	return g.parseMultiCond(id, xRule, yRule, "and")
}

func (g *Generator) orRule(inst *ir.InstOr) rule {
	x := inst.X.Ident()
	y := inst.Y.Ident()
	id := inst.Ident()
	xRule := g.variables.lookupCondPart(g.currentFunction, x)
	if xRule == nil {
		x = g.variables.convertIdent(g.currentFunction, x)
		xRule = &wrap{value: x}
	}

	yRule := g.variables.lookupCondPart(g.currentFunction, y)
	if yRule == nil {
		y = g.variables.convertIdent(g.currentFunction, y)
		yRule = &wrap{value: y}
	}
	return g.parseMultiCond(id, xRule, yRule, "or")
}

func (g *Generator) stateRules(key string, sc *stateChange) rule {
	if len(sc.ors) == 0 {
		and := g.andStateRule(key, sc.ands)
		a := &ands{
			x: and,
		}

		c := &choices{
			x:  []*ands{a},
			op: "and",
		}
		return c
	}

	and := g.andStateRule(key, sc.ands)
	ors := g.orStateRule(key, sc.ors)

	if len(sc.ands) != 0 {
		ors["joined_ands"] = and
	}

	x := g.syncStateRules(ors)

	r := &choices{
		x:  x,
		op: "or",
	}

	return r

}

func (g *Generator) orStateRule(choiceK string, choiceV []value.Value) map[string][]rule {
	g.inPhiState.In()

	and := make(map[string][]rule)
	for _, b := range choiceV {
		refname := fmt.Sprintf("%s-%s", g.currentFunction, b.Ident())
		and[refname] = g.parseBuiltIn(b.(*ir.InstCall), true)
	}
	delete(g.storedChoice, choiceK)

	g.inPhiState.Out()
	return and
}

func (g *Generator) andStateRule(andK string, andV []value.Value) []rule {
	g.inPhiState.In()

	var ands []rule
	for _, b := range andV {
		a := g.parseBuiltIn(b.(*ir.InstCall), true)
		ands = append(ands, a...)
	}
	delete(g.storedChoice, andK)

	g.inPhiState.Out()
	return ands
}

func (g *Generator) syncStateRules(branches map[string][]rule) []*ands {
	g.inPhiState.In()
	g.newFork()

	var e []rule
	var keys []string
	ends := make(map[string][]rule)
	phis := make(map[string]int16)
	var x []*ands

	for k, v := range branches {
		keys = append(keys, k)
		g.buildForkChoice(v, k)
		e, phis = g.capCond(k, phis)
		ends[k] = append(v, e...)
	}

	syncs := g.capCondSyncRules(keys)
	for k, v := range syncs {
		e2 := append(ends[k], v...)
		a := &ands{
			x: e2,
		}
		x = append(x, a)
	}
	g.inPhiState.Out()
	return x
}

func (g *Generator) tempRule(inst value.Value, r rule) {
	// If infix rule is stored in a temp variable
	id := inst.Ident()
	if g.variables.isTemp(id) {
		refname := fmt.Sprintf("%s-%s", g.currentFunction, id)
		g.variables.ref[refname] = r
	}
}
