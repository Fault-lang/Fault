package rules

import (
	"bytes"
	"fault/generator/scenario"
	"fmt"
	"strings"

	"github.com/barkimedes/go-deepcopy"
	"github.com/llir/llvm/ir/value"
)

type SSA struct {
	variables map[string]int16
}

func (s *SSA) Get(k string) int16 {
	return s.variables[k]
}

func (s *SSA) Update(k string) int16 {
	if _, ok := s.variables[k]; !ok {
		s.variables[k] = 0
		return s.variables[k]
	}
	s.variables[k] = s.variables[k] + 1
	return s.variables[k]
}

func (s *SSA) Clone() *SSA {
	m, err := deepcopy.Anything(s.variables)
	if err != nil {
		panic(err)
	}

	return &SSA{
		variables: m.(map[string]int16),
	}
}

func (s *SSA) Iter() map[string]int16 {
	return s.variables
}

func NewSSA() *SSA {
	return &SSA{
		variables: make(map[string]int16),
	}
}

type Rule interface {
	ruleNode()
	LoadContext(int, map[string]bool, map[string][]int16, *scenario.Logger)
	String() string
	Assertless() string
	IsTagged() bool
	Choice() string
	Branch() string
	WriteRule(ssa *SSA) ([]*Init, string, *SSA)
}

type Basic struct {
	Rule
	X        Rule
	Y        Rule
	PhiLevel int
	HaveSeen map[string]bool
	OnEntry  map[string][]int16
	Log      *scenario.Logger
	tag      *branch
}

func (b *Basic) ruleNode() {}

func (b *Basic) LoadContext(PhiLevel int, HaveSeen map[string]bool, OnEntry map[string][]int16, Log *scenario.Logger) {
	b.PhiLevel = PhiLevel
	b.HaveSeen = HaveSeen
	b.OnEntry = OnEntry
	b.Log = Log
}

func (b *Basic) WriteRule(ssa *SSA) ([]*Init, string, *SSA) {
	b.X.LoadContext(b.PhiLevel, b.HaveSeen, b.OnEntry, b.Log)
	b.Y.LoadContext(b.PhiLevel, b.HaveSeen, b.OnEntry, b.Log)

	init1, x, ssa := b.X.WriteRule(ssa)
	init2, y, ssa := b.Y.WriteRule(ssa)
	init := append(init1, init2...)
	return init, fmt.Sprintf("(assert %s %s)", x, y), ssa
}

func (b *Basic) String() string {
	return fmt.Sprintf("basic %s %s", b.X, b.Y)
}

func (b *Basic) Assertless() string {
	return ""
}

func (b *Basic) IsTagged() bool {
	return b.tag != nil
}

func (b *Basic) Choice() string {
	return b.tag.block
}

func (b *Basic) Branch() string {
	return b.tag.branch
}

func (b *Basic) Tag(k1 string, k2 string) {
	b.tag = &branch{
		branch: k1,
		block:  k2,
	}
}

type Init struct {
	Rule
	Ident string //base variable name
	SSA   string //Specific instance of the variable
	//String so that we can tell the difference
	//between "" for constant and "0"
	Type  string
	Value string
	Log   *scenario.Logger
	tag   *branch
}

func (i *Init) ruleNode() {}

func (i *Init) LoadContext(PhiLevel int, HaveSeen map[string]bool, OnEntry map[string][]int16, Log *scenario.Logger) {
	i.Log = Log
}

func (i *Init) WriteRule(ssa *SSA) ([]*Init, string, *SSA) {
	return nil, fmt.Sprintf("(declare-fun %s_%s () %s)", i.Ident, i.SSA, i.Type), ssa
}

func (i *Init) Tuple() []string {
	return []string{i.Ident, i.SSA}
}

func (i *Init) FullVar() string {
	if i.SSA == "" {
		return i.Ident
	}
	return fmt.Sprintf("%s_%s", i.Ident, i.SSA)
}

func (i *Init) String() string {
	return fmt.Sprintf("init %s %s", i.Ident, i.Type)
}

func (i *Init) Assertless() string {
	return ""
}

func (i *Init) IsTagged() bool {
	return i.tag != nil
}

func (i *Init) Choice() string {
	return i.tag.block
}

func (i *Init) Branch() string {
	return i.tag.branch
}

func (i *Init) Tag(k1 string, k2 string) {
	i.tag = &branch{
		branch: k1,
		block:  k2,
	}
}

type Ands struct {
	Rule
	X        []Rule
	PhiLevel int
	HaveSeen map[string]bool
	OnEntry  map[string][]int16
	Log      *scenario.Logger
	tag      *branch
}

func (a *Ands) ruleNode() {}

func (a *Ands) LoadContext(PhiLevel int, HaveSeen map[string]bool, OnEntry map[string][]int16, Log *scenario.Logger) {
	a.PhiLevel = PhiLevel
	a.HaveSeen = HaveSeen
	a.OnEntry = OnEntry
	a.Log = Log
}

func (a *Ands) WriteRule(ssa *SSA) ([]*Init, string, *SSA) {
	var rules []string
	var ru string
	var init, i []*Init
	for _, r := range a.X {
		r.LoadContext(a.PhiLevel, a.HaveSeen, a.OnEntry, a.Log)
		init, ru, ssa = r.WriteRule(ssa)
		rules = append(rules, ru)
		i = append(i, init...)
	}
	return i, fmt.Sprintf("(and %s)", strings.Join(rules, " ")), ssa
}

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

func (a *Ands) IsTagged() bool {
	return a.tag != nil
}

func (a *Ands) Choice() string {
	return a.tag.block
}

func (a *Ands) Branch() string {
	return a.tag.branch
}

type AssertChain struct {
	Op     string
	Values []string
	Chain  []int
	Parent int
}

func (ac *AssertChain) String() string {
	if ac.Op == "" {
		return strings.Join(ac.Values, " ")
	}
	if ac.Op == "!=" {
		return fmt.Sprintf("(not (= %s)", strings.Join(ac.Values, " "))
	}
	return fmt.Sprintf("(%s %s)", ac.Op, strings.Join(ac.Values, " "))
}

// Used to generate SMT based on Assert/Assume logic.
// Lists all possible vars in a scope (if conditional,
// parallel branches, etc)
type PossibleVars struct {
	Rule
	Terminal bool
	Base     []string
	Vars     map[string]*AssertChain
	Constant bool
	tag      *branch
}

func NewPossibleVars() *PossibleVars {
	return &PossibleVars{
		Vars: make(map[string]*AssertChain),
	}
}

func (s *PossibleVars) ruleNode() {}

func (s *PossibleVars) String() string {
	return strings.Join(s.Base, " ")
}

func (s *PossibleVars) Assertless() string {
	return ""
}

func (s *PossibleVars) Tag(k1 string, k2 string) {
	s.tag = &branch{
		branch: k1,
		block:  k2,
	}
}

func (s *PossibleVars) IsTagged() bool {
	return s.tag != nil
}

func (s *PossibleVars) Choice() string {
	return s.tag.block
}

func (s *PossibleVars) Branch() string {
	return s.tag.branch
}

func (s *PossibleVars) Add(base string) {
	s.Base = append(s.Base, base)
}

func (s *PossibleVars) GetChains() []int {
	var ret []int
	for _, a := range s.Vars {
		ret = append(ret, a.Chain...)
	}
	return ret
}

func (s *PossibleVars) LoadContext(PhiLevel int, HaveSeen map[string]bool, OnEntry map[string][]int16, Log *scenario.Logger) {
	// Implement LoadContext if needed
}

func (s *PossibleVars) WriteRule(ssa *SSA) ([]*Init, string, *SSA) {
	// Implement WriteRule if needed
	return nil, "", ssa
}

// type Assrt struct {
// 	Rule
// 	Variable       *Wrap
// 	Conjunction    string
// 	Assertion      Rule
// 	tag            *branch
// 	TemporalFilter string
// 	TemporalN      int
// }

// func (a *Assrt) ruleNode() {}
// func (a *Assrt) String() string {
// 	return a.Variable.String() + a.Conjunction + a.Assertion.String()
// }
// func (a *Assrt) Assertless() string {
// 	return ""
// }
// func (a *Assrt) Tag(k1 string, k2 string) {
// 	a.tag = &branch{
// 		branch: k1,
// 		block:  k2,
// 	}
// }

// func (a *Assrt) IsTagged() bool {
// 	return a.tag != nil
// }

// func (a *Assrt) Choice() string {
// 	return a.tag.block
// }

// func (a *Assrt) Branch() string {
// 	return a.tag.branch
// }

type Parallels struct {
	Rule
	Permutations [][]string
	Calls        map[string][]Rule
	Round        int
	Log          *scenario.Logger
	tag          *branch
}

func NewParallels(permutations [][]string) *Parallels {
	return &Parallels{
		Permutations: permutations,
		Calls:        make(map[string][]Rule),
	}
}

func (p *Parallels) ruleNode() {}

func (p *Parallels) LoadContext(PhiLevel int, HaveSeen map[string]bool, OnEntry map[string][]int16, Log *scenario.Logger) {
}

func (p *Parallels) WriteRule(ssa *SSA) ([]*Init, string, *SSA) {
	return nil, "", ssa
}

func (p *Parallels) String() string {
	var out bytes.Buffer
	for k, v := range p.Calls {
		var rules []string
		for _, r := range v {
			rules = append(rules, r.String())
		}
		out.WriteString(fmt.Sprintf("%s: %s\n", k, strings.Join(rules, ", ")))
	}
	return out.String()
}

func (p *Parallels) Assertless() string {
	return ""
}

func (p *Parallels) Tag(k1 string, k2 string) {
	p.tag = &branch{
		branch: k1,
		block:  k2,
	}
}

func (p *Parallels) IsTagged() bool {
	return p.tag != nil
}

func (p *Parallels) Choice() string {
	return p.tag.block
}

func (p *Parallels) Branch() string {
	return p.tag.branch
}

// type Choices struct {
// 	Rule
// 	X   []*Ands
// 	Op  string
// 	tag *branch
// }

// func (c *Choices) ruleNode() {}
// func (c *Choices) String() string {
// 	var out bytes.Buffer
// 	for i, ru := range c.X {
// 		out.WriteString(fmt.Sprintf("branch-%d: ", i))
// 		for _, r := range ru.X {
// 			out.WriteString(r.String())
// 		}
// 	}
// 	return out.String()
// }
// func (c *Choices) Assertless() string {
// 	return ""
// }
// func (c *Choices) Tag(k1 string, k2 string) {
// 	c.tag = &branch{
// 		branch: k1,
// 		block:  k2,
// 	}
// }

// func (c *Choices) IsTagged() bool {
// 	return c.tag != nil
// }

// func (c *Choices) Choice() string {
// 	return c.tag.block
// }

// func (c *Choices) Branch() string {
// 	return c.tag.branch
// }

type Infix struct {
	Rule
	X   Rule
	Y   Rule
	Ty  string
	Op  string
	tag *branch
	Phi bool //Tag if this rule is a phi value capping a branch
	// If so we don't want to track it as a state change
	PhiLevel int
	HaveSeen map[string]bool
	OnEntry  map[string][]int16
	Log      *scenario.Logger
}

func (i *Infix) ruleNode() {}
func (i *Infix) LoadContext(PhiLevel int, HaveSeen map[string]bool, OnEntry map[string][]int16, Log *scenario.Logger) {
	i.PhiLevel = PhiLevel
	i.HaveSeen = HaveSeen
	i.OnEntry = OnEntry
	i.Log = Log
}
func (i *Infix) String() string {
	return fmt.Sprintf("%s %s %s", i.X.String(), i.Op, i.Y.String())
}

func (i *Infix) WriteRule(ssa *SSA) ([]*Init, string, *SSA) {
	i.Y.LoadContext(i.PhiLevel, i.HaveSeen, i.OnEntry, i.Log)
	i.X.LoadContext(i.PhiLevel, i.HaveSeen, i.OnEntry, i.Log)

	initY, y, ssa := i.Y.WriteRule(ssa) // Y first because nestled rules will assign ssa wrong (eg X' = X + N)
	initX, x, ssa := i.X.WriteRule(ssa)
	init := append(initX, initY...)

	if y == "0x3DA3CA8CB153A753" { //Unknown or uncertain type
		return init, "", ssa
	}

	if _, ok := i.X.(*Wrap); ok && i.Op == "=" {
		i.Log.UpdateVariable(x)
	}

	return init, fmt.Sprintf("(%s %s %s)", i.Op, x, y), ssa
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

func (i *Infix) IsTagged() bool {
	return i.tag != nil
}

func (i *Infix) Choice() string {
	return i.tag.block
}

func (i *Infix) Branch() string {
	return i.tag.branch
}

type Prefix struct {
	Rule
	X        Rule
	Ty       string
	Op       string
	PhiLevel int
	HaveSeen map[string]bool
	OnEntry  map[string][]int16
	Log      *scenario.Logger
	tag      *branch
}

func (pr *Prefix) ruleNode() {}
func (pr *Prefix) LoadContext(PhiLevel int, HaveSeen map[string]bool, OnEntry map[string][]int16, Log *scenario.Logger) {
	pr.PhiLevel = PhiLevel
	pr.HaveSeen = HaveSeen
	pr.OnEntry = OnEntry
	pr.Log = Log
}
func (pr *Prefix) String() string {
	return fmt.Sprintf("%s %s", pr.Op, pr.X.String())
}

func (pr *Prefix) WriteRule(ssa *SSA) ([]*Init, string, *SSA) {
	panic(fmt.Sprintf("WriteRule not implemented for %T", pr))
	//return "", ssa
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

func (pr *Prefix) IsTagged() bool {
	return pr.tag != nil
}

func (pr *Prefix) Choice() string {
	return pr.tag.block
}

func (pr *Prefix) Branch() string {
	return pr.tag.branch
}

type Ite struct {
	Rule
	Cond       Rule
	T          []Rule
	F          []Rule
	After      []Rule
	BlockNames map[string]string
	Log        *scenario.Logger
	tag        *branch
}

func (it *Ite) ruleNode() {}
func NewIte(cond Rule, t []Rule, f []Rule, a []Rule, block_names []string) *Ite {
	return &Ite{
		Cond:  cond,
		T:     t,
		F:     f,
		After: a,
		BlockNames: map[string]string{
			"true":  block_names[0],
			"false": block_names[1],
			"after": block_names[2],
		},
	}
}
func (it *Ite) LoadContext(PhiLevel int, HaveSeen map[string]bool, OnEntry map[string][]int16, Log *scenario.Logger) {
}
func (it *Ite) String() string {
	return fmt.Sprintf("if %s then %s else %s", it.Cond.String(), it.T, it.F)
}

func (it *Ite) WriteRule(ssa *SSA) ([]*Init, string, *SSA) {
	panic(fmt.Sprintf("WriteRule not implemented for %T", it))
	//return "", ssa
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

func (it *Ite) IsTagged() bool {
	return it.tag != nil
}

func (it *Ite) Choice() string {
	return it.tag.block
}

func (it *Ite) Branch() string {
	return it.tag.branch
}

func (ite *Ite) Tag(k1 string, k2 string) {
	ite.tag = &branch{
		branch: k1,
		block:  k2,
	}
}

// type Invariant struct {
// 	Rule
// 	Left     Rule
// 	Operator string
// 	Right    Rule
// 	tag      *branch
// }

// func (i *Invariant) ruleNode() {}
// func (i *Invariant) String() string {
// 	if i.Left == nil { //Prefixes like !a
// 		return fmt.Sprint(i.Operator, i.Right.String())
// 	}
// 	return fmt.Sprint(i.Left.String(), i.Operator, i.Right.String())
// }
// func (i *Invariant) Assertless() string {
// 	return fmt.Sprintf("(%s %s %s)", i.Operator, i.Left.Assertless(), i.Right.Assertless())
// }
// func (i *Invariant) Tag(k1 string, k2 string) {
// 	i.tag = &branch{
// 		branch: k1,
// 		block:  k2,
// 	}
// }

// func (i *Invariant) IsTagged() bool {
// 	return i.tag != nil
// }

// func (i *Invariant) Choice() string {
// 	return i.tag.block
// }

// func (i *Invariant) Branch() string {
// 	return i.tag.branch
// }

// type Phi struct {
// 	Rule
// 	BaseVar  string
// 	Nums     []int16
// 	EndState string
// 	tag      *branch
// }

// func (p *Phi) ruleNode() {}
// func (p *Phi) String() string {
// 	var out bytes.Buffer
// 	for _, n := range p.Nums {
// 		r := fmt.Sprintf("%s = %s_%d || ", p.EndState, p.BaseVar, n)
// 		out.WriteString(r)
// 	}
// 	return out.String()
// }
// func (p *Phi) Assertless() string {
// 	return ""
// }
// func (p *Phi) Tag(k1 string, k2 string) {
// 	p.tag = &branch{
// 		branch: k1,
// 		block:  k2,
// 	}
// }

// func (p *Phi) IsTagged() bool {
// 	return p.tag != nil
// }

// func (p *Phi) Choice() string {
// 	return p.tag.block
// }

// func (p *Phi) Branch() string {
// 	return p.tag.branch
// }

// type StateChange struct {
// 	Rule
// 	Ands  []value.Value
// 	Ors   []value.Value
// 	Rules Rule
// 	tag   *branch
// }

// func (sc *StateChange) ruleNode() {}
// func (sc *StateChange) String() string {
// 	var out bytes.Buffer
// 	for _, n := range sc.Ands {
// 		r := fmt.Sprintf("and %s ", n)
// 		out.WriteString(r)
// 	}
// 	for _, n := range sc.Ors {
// 		r := fmt.Sprintf("or %s ", n)
// 		out.WriteString(r)
// 	}
// 	return out.String()
// }
// func (sc *StateChange) Assertless() string {
// 	return ""
// }
// func (sc *StateChange) Tag(k1 string, k2 string) {
// 	sc.tag = &branch{
// 		branch: k1,
// 		block:  k2,
// 	}
// }

// func (sc *StateChange) IsTagged() bool {
// 	return sc.tag != nil
// }

// func (sc *StateChange) Choice() string {
// 	return sc.tag.block
// }

// func (sc *StateChange) Branch() string {
// 	return sc.tag.branch
// }
// func (sc *StateChange) Empty() bool {
// 	if len(sc.Ands) > 0 {
// 		return false
// 	}
// 	if len(sc.Ors) > 0 {
// 		return false
// 	}
// 	if sc.Rules != nil {
// 		return false
// 	}
// 	return true
// }

type Wrap struct { //wrapper for constant values to be used in infix as rules
	Rule
	Value    string
	Variable bool
	Type     string
	Init     bool              //Are we referencing a existing value or initializating a new one?
	Debugger map[string]string //For debugging, the location in the code where this rule was created
	PhiLevel int
	HaveSeen map[string]bool
	OnEntry  map[string][]int16
	Log      *scenario.Logger
	tag      *branch
}

func (w *Wrap) ruleNode() {}
func (w *Wrap) LoadContext(PhiLevel int, HaveSeen map[string]bool, OnEntry map[string][]int16, Log *scenario.Logger) {
	w.PhiLevel = PhiLevel
	w.HaveSeen = HaveSeen
	w.OnEntry = OnEntry
	w.Log = Log
}
func NewWrap(v string, t string, vr bool, file string, line int, init bool) *Wrap {
	return &Wrap{
		Value:    v,
		Variable: vr,
		Type:     t,
		Init:     init,
		Debugger: map[string]string{
			"file": file,
			"line": fmt.Sprintf("%d", line),
		},
	}
}

func (w *Wrap) String() string {
	return w.Value
}
func DefaultValue(t string) string {
	switch t {
	case "Int":
		return "0"
	case "Float":
		return "0.0"
	case "Real":
		return "0.0"
	case "Bool":
		return "false"
	default:
		panic(fmt.Sprintf("Type %s not supported", t))
	}
}
func (w *Wrap) WriteRule(ssa *SSA) ([]*Init, string, *SSA) {
	var rule string

	if w.Value == "0x3DA3CA8CB153A753" { //An uncertain or unknown value
		return nil, "0x3DA3CA8CB153A753", ssa
	}

	if w.Variable {
		if w.Init {
			rule = fmt.Sprintf("%s_%d", w.Value, ssa.Update(w.Value))
			default_value := DefaultValue(w.Type)
			i := &Init{
				Ident: w.Value,
				SSA:   fmt.Sprintf("%d", ssa.Get(w.Value)),
				Type:  w.Type,
				Value: default_value,
			}
			return []*Init{i}, rule, ssa
		}

		if w.HaveSeen[w.Value] ||
			len(w.OnEntry) == 0 {
			rule = fmt.Sprintf("%s_%d", w.Value, ssa.Get(w.Value))
			return nil, rule, ssa
		}

		rule = fmt.Sprintf("%s_%d", w.Value, w.OnEntry[w.Value][w.PhiLevel])
		w.HaveSeen[w.Value] = true
		return nil, rule, ssa
	}

	return nil, w.Value, ssa
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
func (w *Wrap) IsTagged() bool {
	return w.tag != nil
}

func (w *Wrap) Choice() string {
	return w.tag.block
}

func (w *Wrap) Branch() string {
	return w.tag.branch
}

type VarSets struct {
	Rule
	Vars map[string][]string // [round_0_scope_name] => {this_variable_0, this_variable_1}
	tag  *branch
}

func NewVarSets(vars map[string][]string) *VarSets {
	varset := &VarSets{
		Vars: vars,
	}
	return varset
}

func (sg *VarSets) ruleNode() {}

func (sg *VarSets) String() string {
	var out bytes.Buffer
	for _, v := range sg.Vars {
		out.WriteString(strings.Join(v, "\n"))
	}
	return out.String()
}

func (sg *VarSets) List() []string {
	var ret []string
	for _, v := range sg.Vars {
		ret = append(ret, v...)
	}
	return ret
}

func (sg *VarSets) Assertless() string {
	return ""
}

func (sg *VarSets) Tag(k1 string, k2 string) {
	sg.tag = &branch{
		branch: k1,
		block:  k2,
	}
}

func (sg *VarSets) IsTagged() bool {
	return sg.tag != nil
}

func (sg *VarSets) Choice() string {
	return sg.tag.block
}

func (sg *VarSets) Branch() string {
	return sg.tag.branch
}

func (sg *VarSets) LoadContext(PhiLevel int, HaveSeen map[string]bool, OnEntry map[string][]int16, Log *scenario.Logger) {
	// Implement LoadContext if needed
}

func (sg *VarSets) WriteRule(ssa *SSA) ([]*Init, string, *SSA) {
	// Implement WriteRule if needed
	return nil, "", ssa
}

type WrapGroup struct {
	Rule
	Wraps    []*Wrap
	PhiLevel int
	HaveSeen map[string]bool
	OnEntry  map[string][]int16
	Log      *scenario.Logger
	tag      *branch
}

func (wg *WrapGroup) ruleNode() {}

func (wg *WrapGroup) LoadContext(PhiLevel int, HaveSeen map[string]bool, OnEntry map[string][]int16, Log *scenario.Logger) {
	wg.PhiLevel = PhiLevel
	wg.HaveSeen = HaveSeen
	wg.OnEntry = OnEntry
	wg.Log = Log
}

func (wg *WrapGroup) String() string {
	var out bytes.Buffer
	for _, v := range wg.Wraps {
		out.WriteString(v.Value)
	}
	return out.String()
}

func (wg *WrapGroup) WriteRule(ssa *SSA) ([]*Init, string, *SSA) {
	panic(fmt.Sprintf("WriteRule not implemented for %T", wg))
	//return "", ssa
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

func (wg *WrapGroup) IsTagged() bool {
	return wg.tag != nil
}

func (wg *WrapGroup) Choice() string {
	return wg.tag.block
}

func (wg *WrapGroup) Branch() string {
	return wg.tag.branch
}

type Vwrap struct {
	Rule
	Value value.Value
	tag   *branch
}

func (vw *Vwrap) ruleNode() {}
func (vw *Vwrap) LoadContext(PhiLevel int, HaveSeen map[string]bool, OnEntry map[string][]int16, Log *scenario.Logger) {
}
func (vw *Vwrap) String() string {
	return vw.Value.String()
}

func (vw *Vwrap) WriteRule(ssa *SSA) ([]*Init, string, *SSA) {
	return nil, vw.Value.String(), ssa
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

func (vw *Vwrap) IsTagged() bool {
	return vw.tag != nil
}

func (vw *Vwrap) Choice() string {
	return vw.tag.block
}

func (vw *Vwrap) Branch() string {
	return vw.tag.branch
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
	if ru.IsTagged() {
		return ru //Don't retag something (nestled phis)
	}
	switch r := ru.(type) {
	case *Infix:
		r.X = TagRule(r.X, branch, block)
		r.Y = TagRule(r.Y, branch, block)
		r.Tag(branch, block)
		return r
	case *Prefix:
		r.X = TagRule(r.X, branch, block)
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
	// case *Phi:
	// 	r.Tag(branch, block)
	// 	return r
	case *Ands:
		r.X = TagRules(r.X, branch, block)
		r.Tag(branch, block)
		return r
	// case *Choices:
	// 	var tagged []*Ands
	// 	for _, v := range r.X {
	// 		r2 := TagRule(v, branch, block)
	// 		tagged = append(tagged, r2.(*Ands))
	// 	}
	// 	r.X = tagged
	// 	r.Tag(branch, block)
	// 	return r
	default:
		panic(fmt.Sprintf("%T is not a valid rule type", ru))
	}
}
