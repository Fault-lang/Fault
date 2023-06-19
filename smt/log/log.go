package log

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
	Probability float64
}

func NewLog() *ResultLog {
	return &ResultLog{
		Lookup: make(map[string]int),
	}
}

func (rl *ResultLog) NewChange(round int, scope string, variable string) *Event {
	return &Event{
		Round:    round,
		Type:     "CHANGE",
		Scope:    scope,
		Variable: variable,
	}
}

func (rl *ResultLog) NewTransition(round int, scope string) *Event {
	return &Event{
		Round: round,
		Type:  "TRANSITION",
		Scope: scope,
	}
}

func (rl *ResultLog) NewTrigger(round int, scope string, variable string) *Event {
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

func (rl *ResultLog) Add(e *Event) {
	rl.Events = append(rl.Events, e)
	if e.Variable != "" {
		rl.Lookup[e.Variable] = len(rl.Events)
	}
}

func (rl *ResultLog) UpdateCurrent(idx int, val string) {
	rl.Events[idx].Current = val
}

func (rl *ResultLog) UpdatePrevious(idx int, val string) {
	rl.Events[idx].Previous = val
}

func (rl *ResultLog) UpdateProbability(idx int, p float64) {
	rl.Events[idx].Probability = p
}
