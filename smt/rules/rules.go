package rules

import (
	"bytes"
	"fault/util"
	"fmt"

	"github.com/llir/llvm/ir/value"
)

type Rule interface {
	ruleNode()
	String() string
	Assertless() string
}

type Ands struct {
	Rule
	X   []Rule
	tag *branch
}

func (a *Ands) ruleNode() {}
func (a *Ands) String() string {
	var out bytes.Buffer
	for _, r := range a.X {
		out.WriteString(r.String())
	}
	return out.String()
}
func (a *Ands) Assertless() string {
	var ands string
	for _, asrt := range a.X {
		ands = fmt.Sprintf("%s %s", ands, asrt.Assertless())
	}
	return fmt.Sprintf("(and %s)", ands)
}
func (a *Ands) Tag(k1 string, k2 string) {
	a.tag = &branch{
		branch: k1,
		block:  k2,
	}
}

type States struct {
	Rule
	Terminal bool
	Base     string
	States   map[int][]string
	Constant bool
	tag      *branch
}

func (s *States) ruleNode() {}
func (s *States) String() string {
	return s.Base
}
func (s *States) Assertless() string {
	return ""
}
func (s *States) Tag(k1 string, k2 string) {
	s.tag = &branch{
		branch: k1,
		block:  k2,
	}
}

type Assrt struct {
	Rule
	Variable       *Wrap
	Conjunction    string
	Assertion      Rule
	tag            *branch
	TemporalFilter string
	TemporalN      int
}

func (a *Assrt) ruleNode() {}
func (a *Assrt) String() string {
	return a.Variable.String() + a.Conjunction + a.Assertion.String()
}
func (a *Assrt) Assertless() string {
	return ""
}
func (a *Assrt) Tag(k1 string, k2 string) {
	a.tag = &branch{
		branch: k1,
		block:  k2,
	}
}

type Choices struct {
	Rule
	X   []*Ands
	Op  string
	tag *branch
}

func (c *Choices) ruleNode() {}
func (c *Choices) String() string {
	var out bytes.Buffer
	for i, ru := range c.X {
		out.WriteString(fmt.Sprintf("branch-%d: ", i))
		for _, r := range ru.X {
			out.WriteString(r.String())
		}
	}
	return out.String()
}
func (c *Choices) Assertless() string {
	return ""
}
func (c *Choices) Tag(k1 string, k2 string) {
	c.tag = &branch{
		branch: k1,
		block:  k2,
	}
}

type Infix struct {
	Rule
	X   Rule
	Y   Rule
	Ty  string
	Op  string
	tag *branch
}

func (i *Infix) ruleNode() {}
func (i *Infix) String() string {
	return fmt.Sprintf("%s %s %s", i.X.String(), i.Op, i.Y.String())
}
func (i *Infix) Assertless() string {
	return fmt.Sprintf("(%s %s %s)", i.Op, i.X.Assertless(), i.Y.Assertless())
}
func (i *Infix) Tag(k1 string, k2 string) {
	i.tag = &branch{
		branch: k1,
		block:  k2,
	}
}

type Prefix struct {
	Rule
	X   Rule
	Ty  string
	Op  string
	tag *branch
}

func (pr *Prefix) ruleNode() {}
func (pr *Prefix) String() string {
	return fmt.Sprintf("%s %s", pr.Op, pr.X.String())
}
func (pr *Prefix) Assertless() string {
	return fmt.Sprintf("(%s %s)", pr.Op, pr.X)
}
func (pr *Prefix) Tag(k1 string, k2 string) {
	pr.tag = &branch{
		branch: k1,
		block:  k2,
	}
}

type Ite struct {
	Rule
	Cond Rule
	T    []Rule
	F    []Rule
	tag  *branch
}

func (it *Ite) ruleNode() {}
func (it *Ite) String() string {
	return fmt.Sprintf("if %s then %s else %s", it.Cond.String(), it.T, it.F)
}
func (it *Ite) Assertless() string {
	var t, f string
	for _, tr := range it.T {
		t = fmt.Sprintf("%s %s", t, tr.Assertless())
	}

	for _, fa := range it.F {
		t = fmt.Sprintf("%s %s", f, fa.Assertless())
	}
	return fmt.Sprintf("(ite (%s) (%s) (%s))", it.Cond.Assertless(), t, f)
}

func (ite *Ite) Tag(k1 string, k2 string) {
	ite.tag = &branch{
		branch: k1,
		block:  k2,
	}
}

type Invariant struct {
	Rule
	Left     Rule
	Operator string
	Right    Rule
	tag      *branch
}

func (i *Invariant) ruleNode() {}
func (i *Invariant) String() string {
	if i.Left == nil { //Prefixes like !a
		return fmt.Sprint(i.Operator, i.Right.String())
	}
	return fmt.Sprint(i.Left.String(), i.Operator, i.Right.String())
}
func (i *Invariant) Assertless() string {
	return fmt.Sprintf("(%s %s %s)", i.Operator, i.Left.Assertless(), i.Right.Assertless())
}
func (i *Invariant) Tag(k1 string, k2 string) {
	i.tag = &branch{
		branch: k1,
		block:  k2,
	}
}

type Phi struct {
	Rule
	BaseVar  string
	Nums     []int16
	EndState string
	tag      *branch
}

func (p *Phi) ruleNode() {}
func (p *Phi) String() string {
	var out bytes.Buffer
	for _, n := range p.Nums {
		r := fmt.Sprintf("%s = %s_%d || ", p.EndState, p.BaseVar, n)
		out.WriteString(r)
	}
	return out.String()
}
func (p *Phi) Assertless() string {
	return ""
}
func (p *Phi) Tag(k1 string, k2 string) {
	p.tag = &branch{
		branch: k1,
		block:  k2,
	}
}

type StateChange struct {
	Rule
	Ands  []value.Value
	Ors   []value.Value
	Rules Rule
	tag   *branch
}

func (sc *StateChange) ruleNode() {}
func (sc *StateChange) String() string {
	var out bytes.Buffer
	for _, n := range sc.Ands {
		r := fmt.Sprintf("and %s ", n)
		out.WriteString(r)
	}
	for _, n := range sc.Ors {
		r := fmt.Sprintf("or %s ", n)
		out.WriteString(r)
	}
	return out.String()
}
func (sc *StateChange) Assertless() string {
	return ""
}
func (sc *StateChange) Tag(k1 string, k2 string) {
	sc.tag = &branch{
		branch: k1,
		block:  k2,
	}
}
func (sc *StateChange) Empty() bool {
	if len(sc.Ands) > 0 {
		return false
	}
	if len(sc.Ors) > 0 {
		return false
	}
	if sc.Rules != nil {
		return false
	}
	return true
}

type Wrap struct { //wrapper for constant values to be used in infix as rules
	Rule
	Value    string
	State    string //invariant only for one state
	All      bool   // invariant for all states
	Constant bool   // this is a constant
	tag      *branch
}

func (w *Wrap) ruleNode() {}
func (w *Wrap) String() string {
	return w.Value
}
func (w *Wrap) Assertless() string {
	return w.String()
}
func (w *Wrap) Tag(k1 string, k2 string) {
	w.tag = &branch{
		branch: k1,
		block:  k2,
	}
}

type StateGroup struct {
	Rule
	Bases *util.StringSet
	Wraps []*States
	tag   *branch
}

func NewStateGroup() *StateGroup {
	sg := &StateGroup{}
	sg.Bases = util.NewStrSet()
	return sg
}
func (sg *StateGroup) ruleNode() {}
func (sg *StateGroup) AddWrap(w *States) {
	sg.Wraps = append(sg.Wraps, w)
}
func (sg *StateGroup) String() string {
	var out bytes.Buffer
	for _, v := range sg.Wraps {
		out.WriteString(v.Base)
	}
	return out.String()
}
func (sg *StateGroup) Assertless() string {
	return ""
}
func (sg *StateGroup) Tag(k1 string, k2 string) {
	sg.tag = &branch{
		branch: k1,
		block:  k2,
	}
}

type WrapGroup struct {
	Rule
	Wraps []*Wrap
	tag   *branch
}

func (wg *WrapGroup) ruleNode() {}
func (wg *WrapGroup) String() string {
	var out bytes.Buffer
	for _, v := range wg.Wraps {
		out.WriteString(v.Value)
	}
	return out.String()
}
func (wg *WrapGroup) Assertless() string {
	return ""
}
func (wg *WrapGroup) Tag(k1 string, k2 string) {
	wg.tag = &branch{
		branch: k1,
		block:  k2,
	}
}

type Vwrap struct {
	Rule
	Value value.Value
	tag   *branch
}

func (vw *Vwrap) ruleNode() {}
func (vw *Vwrap) String() string {
	return vw.Value.String()
}
func (vw *Vwrap) Assertless() string {
	return ""
}
func (vw *Vwrap) Tag(k1 string, k2 string) {
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

func TagRules(ru []Rule, branch string, block string) []Rule {
	var tagged []Rule
	for i := 0; i < len(ru); i++ {
		tagged = append(tagged, TagRule(ru[i], branch, block))
	}
	return tagged
}

func TagRule(ru Rule, branch string, block string) Rule {
	switch r := ru.(type) {
	case *Infix:
		r.X = TagRule(r.X, branch, block)
		r.Y = TagRule(r.Y, branch, block)
		r.Tag(branch, block)
		return r
	case *Ite:
		r.Cond = TagRule(r.Cond, branch, block)
		r.T = TagRules(r.T, branch, block)
		r.F = TagRules(r.F, branch, block)
		r.Tag(branch, block)
		return r
	case *Wrap:
		r.Tag(branch, block)
		return r
	case *Vwrap:
		r.Tag(branch, block)
		return r
	case *Phi:
		r.Tag(branch, block)
		return r
	case *Ands:
		r.X = TagRules(r.X, branch, block)
		r.Tag(branch, block)
		return r
	case *Choices:
		var tagged []*Ands
		for _, v := range r.X {
			r2 := TagRule(v, branch, block)
			tagged = append(tagged, r2.(*Ands))
		}
		r.X = tagged
		r.Tag(branch, block)
		return r
	default:
		panic(fmt.Sprintf("%T is not a valid rule type", ru))
	}
}
