package log

import (
	"fault/util"
	"fmt"
)

type ResultLog struct {
	Events  []*Event
	Lookup  map[string]int
	Changes map[string]bool
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

func (e *Event) String() string {
	return fmt.Sprintf("%d,%s,%s,%s,%s,%s,%s\n", e.Round, e.Type, e.Scope, e.Variable, e.Previous, e.Current, e.Probability)
}

func (e *Event) Kill() {
	e.Dead = true

}

func NewLog() *ResultLog {
	return &ResultLog{
		Lookup:  make(map[string]int),
		Changes: make(map[string]bool),
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
		// Kill state variables
		if l.Type == "STATEVAR" {
			rl.Events[idx].Kill()
		}
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
