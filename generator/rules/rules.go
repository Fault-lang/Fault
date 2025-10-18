package rules

import (
	"bytes"
	"fault/generator/scenario"
	"fault/util"
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
	SetRound(int)
	GetRound() int
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
	Round    int
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

func (b *Basic) SetRound(r int) {
	b.Round = r
}

func (b *Basic) GetRound() int {
	return b.Round
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
	Ident    string //base variable name
	SSA      string
	Round    int //Round of the rule, used for SSA
	Global   bool
	Type     string
	Value    Rule
	Solvable bool //If this rule is solvable, meaning it can be used to solve the scenario
	Indexed  bool
	Log      *scenario.Logger
	tag      *branch
}

func (i *Init) ruleNode() {}

func (i *Init) LoadContext(PhiLevel int, HaveSeen map[string]bool, OnEntry map[string][]int16, Log *scenario.Logger) {
	i.Log = Log
}

func (i *Init) SetRound(r int) {
	i.Round = r
	if i.Value != nil {
		i.Value.SetRound(r)
	}
}

func (i *Init) GetRound() int {
	return i.Round
}

func (i *Init) WriteRule(ssa *SSA) ([]*Init, string, *SSA) {
	var id string
	var d string
	var val string
	var rule string

	id = fmt.Sprintf("%s_%s", i.Ident, i.SSA)

	if i.Global && !i.Log.IsCompound[i.Ident] { // Do not log intermediate states in compound string rules
		i.Log.UpdateVariable(id)
	}

	d = fmt.Sprintf("(declare-fun %s () %s)", id, i.Type)

	if i.Value != nil && i.Global && !i.Solvable {
		_, rule, ssa = i.Value.WriteRule(ssa)
		val = fmt.Sprintf("(assert (= %s %s))", id, rule)
		rule = fmt.Sprintf("%s\n%s\n", d, val)
	} else {
		rule = d
	}

	return nil, rule, ssa
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

type FuncCall struct { //A marker rule, generates no SMT but used by the Scenario logger to interpret results
	Rule
	FunctionName string
	Type         string //Enter or Exit
	Round        int
	tag          *branch
}

func NewFuncCall(name string, typ string, round int) *FuncCall {
	return &FuncCall{
		FunctionName: name,
		Type:         typ,
		Round:        round,
	}
}

func (e *FuncCall) ruleNode() {}
func (e *FuncCall) LoadContext(PhiLevel int, HaveSeen map[string]bool, OnEntry map[string][]int16, Log *scenario.Logger) {
	// Implement LoadContext if needed
}
func (e *FuncCall) SetRound(r int) {
	e.Round = r
}
func (e *FuncCall) GetRound() int {
	return e.Round
}
func (e *FuncCall) WriteRule(ssa *SSA) ([]*Init, string, *SSA) {
	// Implement WriteRule if needed
	return nil, "", ssa
}
func (e *FuncCall) String() string {
	return fmt.Sprintf("(enter %s)", e.FunctionName)
}
func (e *FuncCall) Assertless() string {
	return ""
}
func (e *FuncCall) IsTagged() bool {
	return e.tag != nil
}
func (e *FuncCall) Choice() string {
	return e.tag.block
}
func (e *FuncCall) Branch() string {
	return e.tag.branch
}
func (e *FuncCall) Tag(k1 string, k2 string) {
	e.tag = &branch{
		branch: k1,
		block:  k2,
	}
}

type Ands struct {
	Rule
	X        []Rule
	PhiLevel int
	Round    int
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
func (a *Ands) SetRound(r int) {
	a.Round = r
	for _, ru := range a.X {
		ru.SetRound(r)
	}
}
func (a *Ands) GetRound() int {
	return a.Round
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

type Ors struct {
	Rule
	X          [][]Rule
	PhiLevel   int
	Round      int
	HaveSeen   map[string]bool
	OnEntry    map[string][]int16
	BranchName string
	Log        *scenario.Logger
	tag        *branch
}

func (o *Ors) ruleNode() {}

func (o *Ors) LoadContext(PhiLevel int, HaveSeen map[string]bool, OnEntry map[string][]int16, Log *scenario.Logger) {
	o.PhiLevel = PhiLevel
	o.HaveSeen = HaveSeen
	o.OnEntry = OnEntry
	o.Log = Log
}

func (o *Ors) SetRound(r int) {
	o.Round = r
	for _, ru := range o.X {
		for _, ri := range ru {
			ri.SetRound(r)
		}
	}
}

func (o *Ors) GetRound() int {
	return o.Round
}

func (o *Ors) WriteRule(ssa *SSA) ([]*Init, string, *SSA) {
	var rules []string
	var ru string
	var init, i []*Init
	for _, r := range o.X {
		for _, ri := range r {
			ri.LoadContext(o.PhiLevel, o.HaveSeen, o.OnEntry, o.Log)
			init, ru, ssa = ri.WriteRule(ssa)
			rules = append(rules, ru)
			i = append(i, init...)
		}
	}
	return i, fmt.Sprintf("(or %s)", strings.Join(rules, " ")), ssa
}

func (o *Ors) String() string {
	return o.BranchName
}

func (o *Ors) Assertless() string {
	var ors string
	for _, asrt := range o.X {
		for _, a := range asrt {
			ors = fmt.Sprintf("%s %s", ors, a.Assertless())
		}
	}
	return fmt.Sprintf("(or %s)", ors)
}

func (o *Ors) Tag(k1 string, k2 string) {
	o.tag = &branch{
		branch: k1,
		block:  k2,
	}
}

func (o *Ors) IsTagged() bool {
	return o.tag != nil
}

func (o *Ors) Choice() string {
	return o.tag.block
}

func (o *Ors) Branch() string {
	return o.tag.branch
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
func (s *PossibleVars) SetRound(r int) {
	// Implement SetRound if needed
}
func (s *PossibleVars) GetRound() int {
	return 0 // No round concept for PossibleVars
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

func (p *Parallels) SetRound(r int) {
	p.Round = r
	for _, rules := range p.Calls {
		for _, rule := range rules {
			rule.SetRound(r)
		}
	}
}
func (p *Parallels) GetRound() int {
	return p.Round
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

type Infix struct {
	Rule
	X     Rule
	Y     Rule
	Ty    string
	Op    string
	Round int //Round of the rule, used for SSA
	tag   *branch
	Phi   bool //Tag if this rule is a phi value capping a branch
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
func (i *Infix) SetRound(r int) {
	i.Round = r
	i.X.SetRound(r)
	i.Y.SetRound(r)
}
func (i *Infix) GetRound() int {
	return i.Round
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
		i.Log.UpdateSolvable(x)
		return init, "", ssa
	}

	if _, ok := i.X.(*Wrap); ok && i.Op == "=" && !i.Log.IsCompound[x] {
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
	Round    int
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
func (pr *Prefix) SetRound(r int) {
	pr.Round = r
	pr.X.SetRound(r)
}
func (pr *Prefix) GetRound() int {
	return pr.Round
}
func (pr *Prefix) String() string {
	return fmt.Sprintf("(%s %s)", pr.Op, pr.X.String())
}

func (pr *Prefix) WriteRule(ssa *SSA) ([]*Init, string, *SSA) {
	init, x, ssa := pr.X.WriteRule(ssa)
	r := fmt.Sprintf("(%s %s)", pr.Op, x)
	return init, r, ssa
}

func (pr *Prefix) Assertless() string {
	return fmt.Sprintf("(%s %s)", pr.Op, pr.X.Assertless())
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
	Round      int
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
func (it *Ite) SetRound(r int) {
	it.Round = r
	for _, tr := range it.T {
		tr.SetRound(r)
	}
	for _, fa := range it.F {
		fa.SetRound(r)
	}
	for _, a := range it.After {
		a.SetRound(r)
	}
	it.Cond.SetRound(r)
}
func (it *Ite) GetRound() int {
	return it.Round
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

type Stay struct {
	Rule
	Round int
	Log   *scenario.Logger
}

func (s *Stay) ruleNode() {}
func (s *Stay) LoadContext(PhiLevel int, HaveSeen map[string]bool, OnEntry map[string][]int16, Log *scenario.Logger) {
	// Implement LoadContext if needed
}
func (s *Stay) SetRound(r int) {
	s.Round = r
}
func (s *Stay) GetRound() int {
	return s.Round
}
func (s *Stay) WriteRule(ssa *SSA) ([]*Init, string, *SSA) {
	// Implement WriteRule if needed
	return nil, "", ssa
}
func (s *Stay) String() string {
	return "stay"
}
func (s *Stay) Assertless() string {
	return "stay"
}
func (s *Stay) IsTagged() bool {
	return false
}
func (s *Stay) Choice() string {
	return ""
}
func (s *Stay) Branch() string {
	return ""
}

type Wrap struct { //wrapper for constant values to be used in infix as rules
	Rule
	Value    string
	Variable bool
	Type     string
	Init     bool //Are we referencing a existing value or initializating a new one?
	Indexed  bool
	Debugger map[string]string //For debugging, the location in the code where this rule was created
	Round    int
	PhiLevel int
	HaveSeen map[string]bool
	OnEntry  map[string][]int16
	Whens    map[string][]string //map of when asserts this variable is involved in
	Log      *scenario.Logger
	tag      *branch
}

func (w *Wrap) ruleNode() {}

func NewWrap(v string, t string, vr bool, file string, line int, init bool, indexed bool) *Wrap {
	return &Wrap{
		Value:    v,
		Variable: vr,
		Type:     t,
		Init:     init,
		Indexed:  indexed,
		Whens:    make(map[string][]string),
		Debugger: map[string]string{
			"file": file,
			"line": fmt.Sprintf("%d", line),
		},
	}
}

func (w *Wrap) LoadContext(PhiLevel int, HaveSeen map[string]bool, OnEntry map[string][]int16, Log *scenario.Logger) {
	w.PhiLevel = PhiLevel
	w.HaveSeen = HaveSeen
	w.OnEntry = OnEntry
	w.Log = Log
}

func (w *Wrap) SetWhensThens(whens map[string]map[string][]string) {
	if whens[w.Value] != nil {
		w.Whens = whens[w.Value]
	}
}

func (w *Wrap) SetRound(r int) {
	w.Round = r
}
func (w *Wrap) GetRound() int {
	return w.Round
}
func (w *Wrap) String() string {
	return w.Value
}

func (w *Wrap) Clone(phiLevel int) *Wrap {
	return &Wrap{
		Value:    w.Value,
		Variable: w.Variable,
		Type:     w.Type,
		Init:     false,
		Indexed:  false,
		PhiLevel: phiLevel,
		Debugger: make(map[string]string),
	}
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
		if w.Indexed {
			return nil, w.Value, ssa
		}

		if w.Init {
			rule = fmt.Sprintf("%s_%d", w.Value, ssa.Update(w.Value))
			//default_value := DefaultValue(w.Type)
			i := &Init{
				Ident: w.Value,
				SSA:   fmt.Sprintf("%d", ssa.Get(w.Value)),
				Type:  w.Type,
				//Value: &Wrap{Value: default_value},
				Value: nil,
			}
			i.SetRound(w.Round)
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
	Vars map[string]*util.StringSet // [round_0_scope_name] => {this_variable_0, this_variable_1}
	tag  *branch
}

func NewVarSets(vars map[string]*util.StringSet) *VarSets {
	varset := &VarSets{
		Vars: vars,
	}
	return varset
}

func (sg *VarSets) ruleNode() {}

func (sg *VarSets) String() string {
	var out bytes.Buffer
	for _, v := range sg.Vars {
		out.WriteString(strings.Join(v.Values(), "\n"))
	}
	return out.String()
}

func (sg *VarSets) GetByRunRound(round int) []string {
	var ret []string
	for k, v := range sg.Vars {
		key := fmt.Sprintf("round-%d_@__run", round)
		if k == key {
			return v.Values()
		}
	}
	return ret
}

func (sg *VarSets) List() []string {
	var ret []string
	// Get by run rounds to keep vars in order
	for i := 0; len(sg.Vars) > i; i++ {
		ret = append(ret, sg.GetByRunRound(i)...)
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
func (sg *VarSets) SetRound(r int) {
	// Implement SetRound if needed
}
func (sg *VarSets) GetRound() int {
	return 0 // No round concept for VarSets
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
	Round    int
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

func (wg *WrapGroup) SetRound(r int) {
	wg.Round = r
	for _, w := range wg.Wraps {
		w.SetRound(r)
	}
}
func (wg *WrapGroup) GetRound() int {
	return wg.Round
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
	Round int
	tag   *branch
}

func (vw *Vwrap) ruleNode() {}
func (vw *Vwrap) LoadContext(PhiLevel int, HaveSeen map[string]bool, OnEntry map[string][]int16, Log *scenario.Logger) {
}
func (vw *Vwrap) SetRound(r int) {
	vw.Round = r
}
func (vw *Vwrap) GetRound() int {
	return vw.Round
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
	case *Ands:
		r.X = TagRules(r.X, branch, block)
		r.Tag(branch, block)
		return r
	default:
		panic(fmt.Sprintf("%T is not a valid rule type", ru))
	}
}
