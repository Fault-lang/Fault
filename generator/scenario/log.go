package scenario

import (
	"fault/util"
	"fmt"
	"slices"
	"strconv"
	"strings"

	"gonum.org/v1/gonum/stat/distuv"
)

type Logger struct {
	Events          []Event
	Uncertains      map[string][]float64
	BranchIndexes   map[string][]int    // function_name : [event_index_1, event_index_2]
	FuncIndexes     map[string][]int    // function_name : [entry_index, exit_index]
	BranchVars      map[string][]string // function_name : [var_name_1, var_name_2]
	ForksCaps       map[string][]string // var_name_phi : [var_name_endstate1, var_name_endstate2]
	Forks           map[string][]string // branch_name : [var_name_1, var_name_2]
	ForkQueue       []*util.StringSet
	Results         map[string]string // var_name : solution value
	StringRules     map[string]string // var_name : string rule
	IsStringRule    map[string]bool
	IsCompound      map[string]bool // Filter display of compound rules
	IsPhi           map[string]bool
	BranchSelectors []*BranchSelector // rules to make the solution easier to parse
	// RoundPhis stores per-variable phi SSA history: index 0 = initial SSA (before any rounds),
	// index N = phi output SSA after round N. Used by HistoryWrap to resolve value[now-N].
	RoundPhis  map[string][]int16
	SystemName string // stripped from variable and function names in output
}

func NewLogger() *Logger {
	return &Logger{
		Events:        []Event{},
		Uncertains:    make(map[string][]float64),
		Forks:         make(map[string][]string),
		ForksCaps:     make(map[string][]string),
		Results:       make(map[string]string),
		BranchIndexes: make(map[string][]int),
		BranchVars:    make(map[string][]string),
		FuncIndexes:   make(map[string][]int),
		StringRules:   make(map[string]string),
		IsStringRule:  make(map[string]bool),
		IsCompound:    make(map[string]bool),
		IsPhi:         make(map[string]bool),
		RoundPhis:     make(map[string][]int16),
	}
}

func (l *Logger) AddRoundPhi(varName string, ssa int16) {
	l.RoundPhis[varName] = append(l.RoundPhis[varName], ssa)
}

func (l *Logger) EnterFunction(fname string, round int) {

	roundStr := fmt.Sprintf("%d", round)
	l.Events = append(l.Events, &FunctionCall{
		FunctionName: fname,
		Round:        roundStr,
		Type:         "Entry",
	})
}

func (l *Logger) ExitFunction(fname string, round int) {
	roundStr := fmt.Sprintf("%d", round)
	l.Events = append(l.Events, &FunctionCall{
		FunctionName: fname,
		Round:        roundStr,
		Type:         "Exit",
	})
}

func (l *Logger) UpdateVariable(variable string, omit bool) {
	l.Events = append(l.Events, &VariableUpdate{
		Variable: variable,
		Dead:     omit, //omit initialized values to make model output easier to read
	})
}

func (l *Logger) UpdateSolvable(variable string) {
	l.Events = append(l.Events, &Solvable{
		Variable: variable,
	})
}

func (l *Logger) QueueFork(inits []string) {
	if len(inits) == 0 {
		return
	}

	ut := util.NewStrSet()
	for _, v := range inits {
		ut.Add(v)
	}
	l.ForkQueue = append(l.ForkQueue, ut)
}

func (l *Logger) AddPhiOption(phi string, end string) {
	l.IsPhi[phi] = true

	if _, ok := l.ForksCaps[phi]; !ok {
		l.ForksCaps[phi] = []string{end}
	} else {
		l.ForksCaps[phi] = append(l.ForksCaps[phi], end)
	}
	for i, f := range l.ForkQueue {
		if f.In(end) {
			l.Forks[end] = append([]string{}, l.ForkQueue[i].Values()...)
			break
		}
	}
}

func (l *Logger) AddMessage(text string, round int) {
	roundStr := fmt.Sprintf("%d", round)
	l.Events = append(l.Events, &Message{
		Text:  text,
		Round: roundStr,
	})
}

type Event interface {
	MarkDead()
	IsDead() bool
}

type Message struct {
	// For example: hitting a stay() in the statechart
	Event
	Text  string
	Round string
	Dead  bool
}

func (m *Message) MarkDead() {
	m.Dead = true
}

func (m *Message) IsDead() bool {
	return m.Dead
}

type FunctionCall struct {
	Event
	FunctionName string
	Round        string
	Type         string //Entry or Exit
	Dead         bool
}

func (f *FunctionCall) MarkDead() {
	f.Dead = true
}

func (f *FunctionCall) IsDead() bool {
	return f.Dead
}

type VariableUpdate struct {
	Event
	Round    string
	Scope    string
	Variable string
	Dead     bool //Filters out events not in solution
}

func (v *VariableUpdate) MarkDead() {
	v.Dead = true
}

func (v *VariableUpdate) IsDead() bool {
	return v.Dead
}

type Solvable struct {
	Event
	Round        string
	Scope        string
	Variable     string
	Probability  string
	Type         string //Unknown or uncertain
	Distrubution string //Default to Normal
	Dead         bool   //Filters out events not in solution
}

func (s *Solvable) MarkDead() {
	s.Dead = true
}

func (s *Solvable) IsDead() bool {
	return s.Dead
}

func (s *Solvable) SetProbability(val string, mu float64, sigma float64) {
	if s.Type == "Unknown" {
		return
	}

	v, err := strconv.ParseFloat(val, 64)
	if err != nil {
		panic(fmt.Sprintf("failed to parse '%s' as float64: %v", val, err))
	}

	dist := distuv.Normal{
		Mu:    mu,
		Sigma: sigma,
	}
	s.Probability = fmt.Sprintf("%f", dist.Prob(v))
}

func (s *Solvable) GetProbability() string {
	return s.Probability
}

type Choice struct {
	Event
	Operator string
	Options  []string
	Dead     bool
}

func (c *Choice) MarkDead() {
	c.Dead = true
}

func (c *Choice) IsDead() bool {
	return c.Dead
}

func (l *Logger) Trace() {
	var functions []string
	n := 0

	for i, e := range l.Events {
		switch event := e.(type) {
		case *FunctionCall:
			if event.Type == "Entry" {
				fname := fmt.Sprintf("%s-%d", event.FunctionName, n)
				functions = append(functions, fname)
				l.BranchIndexes[fname] = []int{}
				l.BranchVars[fname] = []string{}
				l.FuncIndexes[fname] = []int{i}
				n++
			}
			if event.Type == "Exit" {
				l.FuncIndexes[functions[len(functions)-1]] = append(l.FuncIndexes[functions[len(functions)-1]], i)
				functions = functions[:len(functions)-1]
			}
		case *VariableUpdate:
			var scope string
			if len(functions) == 0 {
				scope = "__global"
			} else {
				scope = functions[len(functions)-1]
			}

			l.BranchIndexes[scope] = append(l.BranchIndexes[scope], i)
			l.BranchVars[scope] = append(l.BranchVars[scope], event.Variable)
		case *Solvable:
			scope := functions[len(functions)-1]
			l.BranchIndexes[scope] = append(l.BranchIndexes[scope], i)
			l.BranchVars[scope] = append(l.BranchVars[scope], event.Variable)

			if event.Type == "Uncertain" {
				// Set probability distribution
				if u, ok := l.Uncertains[event.Variable]; ok {
					event.SetProbability(event.Probability, u[0], u[1])
				} else {
					panic(fmt.Sprintf("Uncertain variable %s not found in Uncertains map", event.Variable))
				}
			}
		}
	}
}

func (l *Logger) IsLoggable(id string) bool {
	//Do not log intermediate states in compound phrases or phis
	return !l.IsCompound[id] && !l.IsPhi[id]
}

func (l *Logger) Validate() {
	//Vars are in the event log but not in a branch
	var missing []string
	var found bool
	for _, event := range l.Events {
		switch e := event.(type) {
		case *VariableUpdate:
			for _, b := range l.BranchSelectors {
				if slices.Contains(b.Vars, e.Variable) {
					found = true
					break
				}
			}
			if !found {
				missing = append(missing, e.Variable)
			}
			found = false
		default:
			break
		}
	}
}

func (l *Logger) Kill() {
	var dead []string
	for _, branch := range l.BranchSelectors {
		name := branch.Id()
		if l.Results[name] == "false" {
			// Kill all the variables in this branch
			dead = append(dead, branch.Vars...)
		}
	}

	if len(dead) == 0 {
		return
	}

	// First pass: mark dead variables
	for i, event := range l.Events {
		switch e := event.(type) {
		case *VariableUpdate:
			if slices.Contains(dead, e.Variable) {
				l.Events[i].MarkDead()
			}
		case *Solvable:
			if slices.Contains(dead, e.Variable) {
				l.Events[i].MarkDead()
			}
		}
	}

	// Second pass: mark functions dead if they have no live variable updates
	for fname, indices := range l.FuncIndexes {
		if len(indices) < 1 {
			continue
		}

		// Synthesis slot functions (synth_N) are always kept alive — Print() handles
		// their display via synthChoice, regardless of what their nested candidates do.
		funcBaseName := strings.TrimRight(fname, "0123456789")
		funcBaseName = strings.TrimSuffix(funcBaseName, "-")
		if isSynthSlotName(funcBaseName) {
			continue
		}

		entryIdx := indices[0]
		var exitIdx int
		if len(indices) >= 2 {
			exitIdx = indices[1]
		} else {
			// No exit event, use end of events list
			exitIdx = len(l.Events)
		}

		// Check if this function has any direct variable updates (not in nested functions)
		// Use BranchIndexes which tracks events directly in this function's scope
		hasLiveUpdates := false
		if varIndices, ok := l.BranchIndexes[fname]; ok {
			for _, idx := range varIndices {
				if l.Events[idx].IsDead() {
					continue
				}
				switch e := l.Events[idx].(type) {
				case *VariableUpdate:
					if !l.IsInternalVariable(e.Variable) {
						hasLiveUpdates = true
						break
					}
				case *Solvable:
					if !l.IsInternalVariable(e.Variable) {
						hasLiveUpdates = true
						break
					}
				}
			}
		}

		// Also check for messages or alive nested function calls directly in this function
		if !hasLiveUpdates {
			for i := entryIdx + 1; i < exitIdx; i++ {
				if l.Events[i].IsDead() {
					continue
				}
				if _, ok := l.Events[i].(*Message); ok {
					hasLiveUpdates = true
					break
				}
				if fc, ok := l.Events[i].(*FunctionCall); ok && fc.Type == "Entry" {
					hasLiveUpdates = true
					break
				}
			}
		}

		// If no live updates (only nested functions), mark the function entry (and exit if exists) as dead
		if !hasLiveUpdates {
			l.Events[entryIdx].MarkDead()
			if len(indices) >= 2 {
				l.Events[exitIdx].MarkDead()
			}
		}
	}
}

type BranchSelector struct {
	Name string
	SSA  int
	Cond []string
	Vars []string // all the vars in this branch
}

func (bs *BranchSelector) Id() string {
	return fmt.Sprintf("%s_%d", bs.Name, bs.SSA)
}

func (bs *BranchSelector) WriteRule() string {
	name := bs.Id()
	if len(bs.Cond) == 0 {
		panic(fmt.Sprintf("Branch Selector %s is empty", name))
	}

	if len(bs.Cond) == 1 {
		return fmt.Sprintf("(= %s %s)", name, bs.Cond[0])
	}
	return fmt.Sprintf("(= %s (and %s))", name, strings.Join(bs.Cond, "\n"))
}

func (bs *BranchSelector) String() string {
	return fmt.Sprintf("branch_selector %s_%d", bs.Name, bs.SSA)
}

func (l *Logger) NewBranchSelector(name string, ssa int, cond []string, inits []string) *BranchSelector {
	return &BranchSelector{
		Name: name,
		SSA:  ssa,
		Cond: cond,
		Vars: inits,
	}
}

func (l *Logger) AddBranchSelector(s *BranchSelector) {
	l.BranchSelectors = append(l.BranchSelectors, s)
}

// stripSysPrefix removes the system name prefix (e.g. "waterpumpmonitor_") from a
// variable or function name so output is readable without the model-level namespace.
func (l *Logger) stripSysPrefix(name string) string {
	if l.SystemName == "" {
		return name
	}
	prefix := l.SystemName + "_"
	if strings.HasPrefix(name, prefix) {
		return name[len(prefix):]
	}
	return name
}

// displayFuncName strips the LLVM basic block suffix (e.g. "-%8") from a function name.
func displayFuncName(fname string) string {
	if i := strings.LastIndex(fname, "-%"); i >= 0 {
		if suffix := fname[i+2:]; len(suffix) > 0 {
			allDigits := true
			for _, c := range suffix {
				if c < '0' || c > '9' {
					allDigits = false
					break
				}
			}
			if allDigits {
				return fname[:i]
			}
		}
	}
	return fname
}

// isSynthSlotName returns true if name is a synthesis slot identifier like "synth_1".
func isSynthSlotName(name string) bool {
	if !strings.HasPrefix(name, "synth_") {
		return false
	}
	suffix := name[len("synth_"):]
	for _, c := range suffix {
		if c < '0' || c > '9' {
			return false
		}
	}
	return len(suffix) > 0
}

func getBase(s string) string {
	// Remove the SSA number from the variable name
	parts := strings.Split(s, "_")
	return strings.Join(parts[:len(parts)-1], "_")
}

func (l *Logger) IsNegated(s string) (string, bool) {
	// Check if the string contains "not"
	if strings.Contains(s, "_neg") {
		parts := strings.Split(s, "_neg")
		if len(parts) > 1 {
			return parts[0], true
		}
		panic(fmt.Sprintf("malformed '%s' negated string", s))
	}
	return s, false
}

func (l *Logger) IsInternalVariable(varName string) bool {
	// Filter out internal solver variables:
	// - block selectors: block*true_*, block*false_*
	// - state selectors: *__state-%*
	// - synthesis selectors: synth_N_*
	base := getBase(varName)

	if strings.HasPrefix(base, "block") && (strings.HasSuffix(base, "true") || strings.HasSuffix(base, "false")) {
		return true
	}

	if strings.Contains(base, "__state-%") {
		return true
	}

	if strings.HasPrefix(base, "synth_") {
		return true
	}

	return false
}

// synthChoice finds which candidate function was chosen for a synthesis slot.
// slotName is "synth_N"; returns the function name if a true selector is found.
func (l *Logger) synthChoice(slotName string) string {
	prefix := slotName + "_"
	for varWithSSA, val := range l.Results {
		if val != "true" {
			continue
		}
		base := getBase(varWithSSA)
		if strings.HasPrefix(base, prefix) {
			return strings.TrimPrefix(base, prefix)
		}
	}
	return ""
}

func (l *Logger) Print() {
	fmt.Print(l.String())
}

func (l *Logger) PrintRaw() {
	identLevel := ""
	for _, e := range l.Events {
		switch event := e.(type) {
		case *FunctionCall:
			continue
		case *VariableUpdate:
			// Skip internal solver variables even in raw output
			if l.IsInternalVariable(event.Variable) {
				continue
			}
			fmt.Printf("%s = %s", event.Variable, l.Results[event.Variable])
			if event.Dead {
				fmt.Printf(" is dead")
			}
			fmt.Printf("\n")
		case *Solvable:
			// Skip internal solver variables even in raw output
			if l.IsInternalVariable(event.Variable) {
				continue
			}
			fmt.Printf("%s = %s", event.Variable, l.Results[event.Variable])
			if event.Dead {
				fmt.Printf(" is dead")
			}
			fmt.Printf("\n")
		case *Message:
			fmt.Printf("%s%s\n", identLevel, event.Text)
		}
	}
	fmt.Print("\n")
}

// String returns the formatted output as a string instead of printing.
//
// Each open function is buffered; on exit the buffer is only flushed to the
// parent if it contains at least one line. This means functions with no
// observable variable changes are silently omitted, regardless of model type.
func (l *Logger) String() string {
	type frame struct {
		displayName string
		buf         strings.Builder
	}

	var stack []*frame
	var root strings.Builder

	// initialStates collects "Set variable X to value true" lines for state
	// variables active at model start. Printed as a block before the "---" divider.
	var initialStates []string

	// write sends s to the innermost open frame, or to root when no frame is open.
	write := func(s string) {
		if len(stack) > 0 {
			stack[len(stack)-1].buf.WriteString(s)
		} else {
			root.WriteString(s)
		}
	}

	// indent returns whitespace appropriate for the current nesting depth.
	// len(stack)==0 means we are at the @__run body level (1 level of indent).
	indent := func() string {
		return strings.Repeat("   ", len(stack)+1)
	}

	// flush pops the topmost frame. If its buffer is non-empty the function
	// header and body are written to the new top (or root). Empty frames are
	// dropped — this is what hides functions with no observable events.
	flush := func() {
		if len(stack) == 0 {
			return
		}
		f := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		content := f.buf.String()
		if content == "" {
			return
		}
		ind := strings.Repeat("   ", len(stack)+1)

		// For state functions (Foo__state), if the first buffered line sets the
		// corresponding state variable to true, hoist it before the function header.
		// This makes clear that the state was activated before the function ran,
		// not that the function activated itself.
		if strings.HasSuffix(f.displayName, "__state") {
			stateName := strings.TrimSuffix(f.displayName, "__state")
			innerInd := ind + "   "
			marker := fmt.Sprintf("%sSet variable %s to value true\n", innerInd, stateName)
			if strings.HasPrefix(content, marker) {
				// Collect into the initial-state block rather than writing inline.
				initialStates = append(initialStates, fmt.Sprintf("%sSet variable %s to value true\n", ind, stateName))
				content = content[len(marker):]
			}
		}

		write(fmt.Sprintf("%sRun function %s\n", ind, f.displayName))
		if content != "" {
			write(content)
		}
	}

	// currentState: base variable name → latest known value.
	currentState := make(map[string]string)

	// Pre-scan the event log to seed currentState from run-block direct
	// assignments only. This avoids picking up SSA helper variables that
	// happen to share the same _1 index as real initial states.
	{
		inRun, depth := false, 0
		for _, e := range l.Events {
			switch ev := e.(type) {
			case *FunctionCall:
				if ev.FunctionName == "@__run" {
					if ev.Type == "Entry" {
						inRun = true
					} else {
						inRun = false
					}
				} else if inRun {
					if ev.Type == "Entry" {
						depth++
					} else {
						depth--
					}
				}
			case *VariableUpdate:
				if inRun && depth == 0 && !e.IsDead() && !l.IsInternalVariable(ev.Variable) {
					if val := l.Results[ev.Variable]; val == "true" {
						currentState[getBase(ev.Variable)] = val
					}
				}
			}
		}
	}

	// componentPrefix returns everything but the last underscore-segment of
	// varName, used to group state variables by component.
	componentPrefix := func(varName string) string {
		parts := strings.Split(varName, "_")
		if len(parts) >= 3 {
			return strings.Join(parts[:len(parts)-1], "_")
		}
		return ""
	}

	// complementaryStates returns base variable names in the same component
	// that are currently true (they will be set to false by the new true state).
	complementaryStates := func(varName string) []string {
		prefix := componentPrefix(varName)
		if prefix == "" {
			return nil
		}
		var out []string
		for other, val := range currentState {
			if val == "true" && other != varName && strings.HasPrefix(other, prefix+"_") {
				out = append(out, other)
			}
		}
		return out
	}

	for _, e := range l.Events {
		if e.IsDead() {
			continue
		}
		switch event := e.(type) {
		case *FunctionCall:
			displayName := displayFuncName(l.stripSysPrefix(event.FunctionName))

			if event.FunctionName == "@__run" {
				if event.Type == "Entry" {
					root.WriteString("\nInitialize model\n")
					root.WriteString("-----------------------------------\n")
					root.WriteString("\nStart model\n")
					root.WriteString("-----------------------------------\n")
				}
				// @__run exit: nothing to flush — content was written directly to root
				continue
			}

			if strings.HasPrefix(event.FunctionName, "synth_") {
				if event.Type == "Entry" {
					chosen := l.synthChoice(event.FunctionName)
					if chosen != "" {
						write(fmt.Sprintf("%sFault chose %s (step %s)\n", indent(), chosen, event.Round))
					} else {
						write(fmt.Sprintf("%sSynthesis step %s (unsatisfiable)\n", indent(), event.Round))
					}
					stack = append(stack, &frame{displayName: displayName})
				} else {
					flush()
				}
				continue
			}

			// Internal block scopes (-%N suffix) are transparent: skip them.
			if displayFuncName(event.FunctionName) != event.FunctionName {
				continue
			}

			if event.Type == "Entry" {
				stack = append(stack, &frame{displayName: displayName})
			} else {
				flush()
			}

		case *VariableUpdate:
			if l.IsInternalVariable(event.Variable) {
				continue
			}
			v := getBase(event.Variable)
			s, negated := l.IsNegated(v)
			newValue := l.Results[event.Variable]
			oldValue, hasOldValue := currentState[v]

			// When a state becomes true, emit implicit false transitions for
			// other states in the same component that were previously true.
			if newValue == "true" && newValue != oldValue {
				for _, comp := range complementaryStates(v) {
					write(fmt.Sprintf("%s%s: true → false\n", indent(), l.stripSysPrefix(comp)))
					currentState[comp] = "false"
				}
			}
			currentState[v] = newValue

			displayV := l.stripSysPrefix(v)
			if l.IsStringRule[s] {
				s = l.StringRules[s]
				if hasOldValue && oldValue != newValue {
					if negated {
						write(fmt.Sprintf("%s not %s: %s → %s\n", indent(), s, oldValue, newValue))
					} else {
						write(fmt.Sprintf("%s %s: %s → %s\n", indent(), s, oldValue, newValue))
					}
				} else if !hasOldValue {
					if negated {
						write(fmt.Sprintf("%s not %s is %s\n", indent(), s, newValue))
					} else {
						write(fmt.Sprintf("%s %s is %s\n", indent(), s, newValue))
					}
				}
			} else {
				if !hasOldValue && newValue != "false" {
					// Suppress first-time false: a variable being false for the
					// first time is not an observable event in the trace.
					write(fmt.Sprintf("%sSet variable %s to value %s\n", indent(), displayV, newValue))
				} else if hasOldValue && oldValue != newValue {
					write(fmt.Sprintf("%s%s: %s → %s\n", indent(), displayV, oldValue, newValue))
				}
			}

		case *Solvable:
			if l.IsInternalVariable(event.Variable) {
				continue
			}
			v := getBase(event.Variable)
			newValue := l.Results[event.Variable]
			if event.Type == "Uncertain" {
				write(fmt.Sprintf("%sResolving variable %s to value %s (%s)\n", indent(), l.stripSysPrefix(v), newValue, event.Probability))
			} else {
				write(fmt.Sprintf("%sResolving variable %s to value %s\n", indent(), l.stripSysPrefix(v), newValue))
			}
			currentState[v] = newValue

		case *Message:
			// Messages are encoding artifacts (e.g. "Stay in current state").
			// They are not meaningful model events and are suppressed.
		}
	}

	// Insert the collected initial states
	if len(initialStates) > 0 {
		const model_run = "\nStart model\n-----------------------------------\n"
		const model_init = "\nInitialize model\n-----------------------------------\n"
		s := root.String()
		if idx := strings.Index(s, model_init+model_run); idx >= 0 {
			insertAt := idx + len(model_init)
			root.Reset()
			root.WriteString(s[:insertAt])
			for _, line := range initialStates {
				root.WriteString(line)
			}
			root.WriteString(s[insertAt:])
		}
	}

	root.WriteString("\n")
	return root.String()
}
