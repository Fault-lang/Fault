package log

import "fmt"

type ResultLog struct {
	Events []*Event
	Lookup map[string]int
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
		Lookup: make(map[string]int),
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

func (rl *ResultLog) FilterOut(deadVars []string) {
	for _, dvar := range deadVars {
		idx := rl.Lookup[dvar]
		rl.Events[idx].Kill()
	}
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
	rl.Events = append(rl.Events, e)
	if e.Variable != "" {
		rl.Lookup[e.Variable] = len(rl.Events) - 1
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
