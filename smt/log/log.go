package log

import (
	"fault/smt/rules"
	"fault/util"
	"fmt"
	"strconv"
	"strings"
)

type ResultLog struct {
	Events        []*Event
	Lookup        map[string]int
	Changes       map[string]bool
	Asserts       []*Assert
	AssertClauses map[string]bool
	AssertChains  map[string]*rules.AssertChain
}

type Event struct {
	Round       int
	Type        string
	Scope       string
	Variable    string
	Previous    string
	Current     string
	Probability string
	Dead        bool //Filters out events not in solution
}

type Clause interface {
	Type() string
	GetFloat() float64
	GetInt() int64
	GetBool() bool
	GetString() string
	String() string
}

type FlClause struct {
	Clause
	Value float64
}

func (flc *FlClause) Type() string {
	return "FLOAT"
}
func (flc *FlClause) GetFloat() float64 {
	return flc.Value
}
func (flc *FlClause) GetInt() int64 {
	return int64(flc.Value)
}
func (flc *FlClause) GetBool() bool {
	return false
}
func (flc *FlClause) GetString() string {
	return ""
}
func (flc *FlClause) String() string {
	return fmt.Sprintf("%f", flc.Value)
}

type IntClause struct {
	Clause
	Value int64
}

func (intc *IntClause) Type() string {
	return "INT"
}
func (intc *IntClause) GetFloat() float64 {
	return float64(intc.Value)
}
func (intc *IntClause) GetInt() int64 {
	return intc.Value
}
func (intc *IntClause) GetBool() bool {
	return false
}
func (intc *IntClause) GetString() string {
	return ""
}
func (intc *IntClause) String() string {
	return fmt.Sprintf("%d", intc.Value)
}

type BoolClause struct {
	Clause
	Value bool
}

func (boolc *BoolClause) Type() string {
	return "BOOL"
}
func (boolc *BoolClause) GetFloat() float64 {
	return 0.0
}
func (boolc *BoolClause) GetInt() int64 {
	return 0
}
func (boolc *BoolClause) GetBool() bool {
	return boolc.Value
}
func (boolc *BoolClause) GetString() string {
	return ""
}
func (boolc *BoolClause) String() string {
	return fmt.Sprintf("%v", boolc.Value)
}

type StringClause struct {
	Clause
	Value string
}

func (strc *StringClause) Type() string {
	return "STRING"
}
func (strc *StringClause) GetFloat() float64 {
	return 0.0
}
func (strc *StringClause) GetInt() int64 {
	return 0
}
func (strc *StringClause) GetBool() bool {
	return false
}
func (strc *StringClause) GetString() string {
	return strc.Value
}
func (strc *StringClause) String() string {
	return strc.Value
}

type Assert struct {
	Left  Clause
	Right Clause
	Op    string
}

func (a *Assert) String() string {
	return fmt.Sprintf("(%s %s %s)", a.Op, a.Left.String(), a.Right.String())
}

func (rl *ResultLog) NewAssert(l string, r string, op string) int {
	left := rl.NewClause(l)
	right := rl.NewClause(r)
	rl.Asserts = append(rl.Asserts, &Assert{Left: left, Right: right, Op: op})
	return len(rl.Asserts) - 1
}

func (rl *ResultLog) NewClause(x string) Clause {
	if x == "true" || x == "false" { // 0 and 1 as well as partials like "t" are not valid anyway
		b, err := strconv.ParseBool(x)
		if err == nil {
			return &BoolClause{Value: b}
		}
	}

	f, err := strconv.ParseFloat(x, 64)
	if err == nil && strings.Contains(x, ".") {
		return &FlClause{Value: f}
	}

	i, err := strconv.ParseInt(x, 10, 64)
	if err == nil {
		return &IntClause{Value: i}
	}

	return &StringClause{Value: x}

}

func (rl *ResultLog) StoreEval(a *Assert, res bool) {
	key := a.String()
	rl.AssertClauses[key] = res
}

func (e *Event) String() string {
	return fmt.Sprintf("%d,%s,%s,%s,%s,%s,%s\n", e.Round, e.Type, e.Scope, e.Variable, e.Previous, e.Current, e.Probability)
}

func (e *Event) Kill() {
	e.Dead = true

}

func NewLog() *ResultLog {
	return &ResultLog{
		Lookup:        make(map[string]int),
		Changes:       make(map[string]bool),
		AssertClauses: make(map[string]bool),
		AssertChains:  make(map[string]*rules.AssertChain),
	}
}

func NewInit(round int, scope string, variable string) *Event {
	return &Event{
		Round:    round,
		Type:     "INIT",
		Scope:    scope,
		Variable: variable,
	}
}

func NewChange(round int, scope string, variable string) *Event {
	return &Event{
		Round:    round,
		Type:     "CHANGE",
		Scope:    scope,
		Variable: variable,
	}
}

func NewStateVar(round int, scope string, variable string) *Event {
	return &Event{
		Round:    round,
		Type:     "STATEVAR",
		Scope:    scope,
		Variable: variable,
	}
}

func NewTransition(round int, previous string, current string) *Event {
	return &Event{
		Round:    round,
		Type:     "TRANSITION",
		Variable: "__state",
		Previous: previous,
		Current:  current,
	}
}

func NewTrigger(round int, scope string, variable string) *Event {
	return &Event{
		Round:    round,
		Type:     "TRIGGER",
		Scope:    scope,
		Variable: variable,
	}
}

func (rl *ResultLog) Index(name string) int {
	if i, ok := rl.Lookup[name]; ok {
		return i
	}
	return -1
}

func (rl *ResultLog) FilterStateTransitions() {
	for idx, l := range rl.Events {
		if idx > 1 && l.Type == "TRANSITION" {
			previous := l.Previous
			current := l.Current
			if rl.deadTransition(previous, idx-1) &&
				rl.deadTransition(current, idx-2) {
				//Kill everything! :D
				rl.Events[idx].Kill()
				rl.Events[idx-1].Kill()
				rl.Events[idx-2].Kill()

			}

		}
	}
}

func (rl *ResultLog) FilterStateVars() {
	for idx, l := range rl.Events {
		if l.Type == "STATEVAR" {
			rl.Events[idx].Kill()
		}
	}
}

func (rl *ResultLog) FilterOut(deadVars []string) {
	for _, dvar := range deadVars {
		if idx, ok := rl.Lookup[dvar]; ok {
			rl.Events[idx].Kill()

			// If this branch is really the result of
			// a function call, remove that too.
			if idx != 0 && rl.Events[idx-1].Type == "TRIGGER" {
				rl.Events[idx-1].Kill()
			}
		}
	}

	rl.FilterStateTransitions()
	rl.FilterStateVars()
}

func (rl *ResultLog) deadTransition(stateVar string, idx int) bool {
	base, _ := util.GetVarBase(rl.Events[idx].Variable)
	return base == stateVar && rl.Events[idx].Dead
}

func (rl *ResultLog) String() string {
	var str = "Round,Type,Scope,Variable,Previous,Current,Probability\n"
	for _, l := range rl.Events {
		if !l.Dead {
			str = fmt.Sprintf("%s%s", str, l.String())
		}
	}
	return str
}

func (rl *ResultLog) Add(e *Event) {
	if !rl.Changes[e.Variable] {
		rl.Events = append(rl.Events, e)
		if e.Variable != "" {
			rl.Lookup[e.Variable] = len(rl.Events) - 1
		}

		if e.Variable != "" && e.Type == "CHANGE" {
			rl.Changes[e.Variable] = true
		}
	}
}

func (rl *ResultLog) UpdateCurrent(idx int, val string) {
	rl.Events[idx].Current = val
}

func (rl *ResultLog) UpdatePrevious(idx int, val string) {
	rl.Events[idx].Previous = val
}

func (rl *ResultLog) UpdateProbability(idx int, p float64) {
	rl.Events[idx].Probability = fmt.Sprintf("%f", p)
}
