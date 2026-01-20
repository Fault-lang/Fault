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
	}
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

		// Also check for messages directly in this function
		if !hasLiveUpdates {
			for i := entryIdx + 1; i < exitIdx; i++ {
				if _, ok := l.Events[i].(*Message); ok && !l.Events[i].IsDead() {
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

	// deadSet := make(map[string]bool)

	// for phi, options := range l.ForksCaps {
	// 	phiVal, ok := l.Results[phi]
	// 	if !ok {
	// 		// No value for this phi in the model; skip or log.
	// 		continue
	// 	}

	// 	// Find which endstateVar matches the phi value.
	// 	chosenIdx := -1
	// 	for i, endstate := range options {
	// 		if l.Results[endstate] == phiVal {
	// 			chosenIdx = i
	// 			break
	// 		}
	// 	}
	// 	if chosenIdx == -1 {
	// 		// No matching option — either model is weird or encoding changed.
	// 		// You might want to log/return an error instead of silently skipping.
	// 		continue
	// 	}

	// 	// All other options are dead branches.
	// 	for i, endstate := range options {
	// 		if i == chosenIdx {
	// 			continue
	// 		}
	// 		// Kill all vars in that branch.
	// 		for _, v := range l.Forks[endstate] {
	// 			deadSet[v] = true
	// 		}
	// 	}
	// }

	// // Convert set to slice.
	// dead := make([]string, 0, len(deadSet))
	// for v := range deadSet {
	// 	dead = append(dead, v)
	// }

	// if len(dead) == 0 {
	// 	return
	// }

	// for i, event := range l.Events {
	// 	switch e := event.(type) {
	// 	case *VariableUpdate:
	// 		if slices.Contains(dead, e.Variable) {
	// 			l.Events[i].MarkDead()
	// 			for _, i := range l.FuncIndexes[e.Variable] {
	// 				l.Events[i].MarkDead()
	// 			}
	// 		}
	// 	}
	// }

	// for fname, vars := range l.BranchVars {
	// 	for _, v := range vars {
	// 		if slices.Contains(dead, v) {
	// 			// Kill Variable Updates
	// 			for _, i := range l.BranchIndexes[fname] {
	// 				l.Events[i].MarkDead()
	// 			}
	// 			// Kill Function Calls themselves
	// 			for _, i := range l.FuncIndexes[fname] {
	// 				l.Events[i].MarkDead()
	// 			}

	// 			break
	// 		}
	// 	}
	// }
}

// func (l *Logger) Kill() {
// 	var deadends, dead []string
// 	for phi, options := range l.ForksCaps {
// 		phi_value := l.Results[phi]
// 		for i, o := range options {
// 			if phi_value == l.Results[o] {
// 				deadends = append(deadends, options[0:i]...)
// 				if i+1 < len(options) {
// 					deadends = append(deadends, options[i+1:]...)
// 				}

// 				for _, d := range deadends {
// 					dead = append(dead, l.Forks[d]...)
// 				}
// 				deadends = []string{}
// 				break
// 			}
// 		}
// 	}

// 	if len(dead) == 0 {
// 		return
// 	}
// 	for fname, vars := range l.BranchVars {
// 		for _, v := range vars {
// 			if slices.Contains(dead, v) {
// 				// Kill Variable Updates
// 				for _, i := range l.BranchIndexes[fname] {
// 					l.Events[i].MarkDead()
// 				}
// 				// Kill Function Calls themselves
// 				for _, i := range l.FuncIndexes[fname] {
// 					l.Events[i].MarkDead()
// 				}

// 				break
// 			}
// 		}
// 	}
// }

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
	base := getBase(varName)

	if strings.HasPrefix(base, "block") && (strings.HasSuffix(base, "true") || strings.HasSuffix(base, "false")) {
		return true
	}

	if strings.Contains(base, "__state-%") {
		return true
	}

	return false
}

func (l *Logger) Print() {
	fmt.Print("\n===================================\n")
	fmt.Printf("Fault found the following scenario\n")
	identLevel := ""

	// Track current state to show transitions
	currentState := make(map[string]string)

	// Helper to get component prefix from state variable (e.g., "infusion_pump_idle" -> "infusion_pump")
	getComponentPrefix := func(varName string) string {
		parts := strings.Split(varName, "_")
		if len(parts) >= 3 {
			// For patterns like "infusion_pump_idle", return "infusion_pump"
			// Return all but the last part (which is the state name)
			return strings.Join(parts[:len(parts)-1], "_")
		}
		return ""
	}

	// Helper to find complementary state changes in the same component
	findComplementaryStates := func(varName string, newValue string) []string {
		if newValue != "true" {
			return nil
		}

		componentPrefix := getComponentPrefix(varName)
		if componentPrefix == "" {
			return nil
		}

		var complementary []string
		// Look through currentState to find other states in same component that were true
		for otherVar, otherValue := range currentState {
			if otherValue == "true" && otherVar != varName && strings.HasPrefix(otherVar, componentPrefix+"_") {
				complementary = append(complementary, otherVar)
			}
		}
		return complementary
	}

	// Initialize currentState with starting states by looking for state variables with low SSA indices
	// Only use _1 (the first state after initialization) and ensure only one state per component is set
	// Prefer common starting states like "idle", "silent", "empty", "off", etc.
	componentStates := make(map[string]string) // Track which components already have a state set (map to state name)
	preferredStates := []string{"idle", "silent", "empty", "off", "initial", "stopped"}

	// First pass: look for preferred starting states
	for varWithSSA, value := range l.Results {
		if !strings.HasSuffix(varWithSSA, "_1") || value != "true" || l.IsInternalVariable(varWithSSA) {
			continue
		}

		baseVar := getBase(varWithSSA)
		componentPrefix := getComponentPrefix(baseVar)
		if componentPrefix == "" {
			continue
		}

		// Check if this is a preferred starting state
		for _, preferred := range preferredStates {
			if strings.HasSuffix(baseVar, "_"+preferred) {
				if _, exists := componentStates[componentPrefix]; !exists {
					currentState[baseVar] = value
					componentStates[componentPrefix] = baseVar
				}
				break
			}
		}
	}

	// Second pass: fill in any components that don't have a state yet
	for varWithSSA, value := range l.Results {
		if !strings.HasSuffix(varWithSSA, "_1") || value != "true" || l.IsInternalVariable(varWithSSA) {
			continue
		}

		baseVar := getBase(varWithSSA)
		componentPrefix := getComponentPrefix(baseVar)
		if componentPrefix == "" {
			continue
		}

		if _, exists := componentStates[componentPrefix]; !exists {
			currentState[baseVar] = value
			componentStates[componentPrefix] = baseVar
		}
	}

	for _, e := range l.Events {
		if e.IsDead() {
			continue
		}
		switch event := e.(type) {
		case *FunctionCall:
			if event.FunctionName == "@__run" {
				fmt.Print("\n")
				fmt.Printf("%sStart model, run for %s rounds\n", identLevel, event.Round)
				fmt.Printf("-----------------------------------\n")
				identLevel += "   "
				continue
			}

			if event.Type == "Entry" {
				fmt.Printf("%sRun function %s (round %s)\n", identLevel, event.FunctionName, event.Round)
				identLevel += "   "
			}

			if event.Type == "Exit" {
				identLevel = identLevel[:len(identLevel)-3]
			}
		case *VariableUpdate:
			// Skip internal solver variables
			if l.IsInternalVariable(event.Variable) {
				continue
			}

			v := getBase(event.Variable)
			s, negated := l.IsNegated(v)
			newValue := l.Results[event.Variable]
			oldValue, hasOldValue := currentState[v]

			// Show complementary state transitions (e.g., when infusing becomes true, show idle becoming false)
			if newValue == "true" && newValue != oldValue {
				complementary := findComplementaryStates(v, newValue)
				for _, compVar := range complementary {
					fmt.Printf("%s%s: true → false\n", identLevel, compVar)
					currentState[compVar] = "false"
				}
			}

			// Update state
			currentState[v] = newValue

			if l.IsStringRule[s] == true {
				s = l.StringRules[s]
				if hasOldValue && oldValue != newValue {
					if negated {
						fmt.Printf("%s not %s: %s → %s\n", identLevel, s, oldValue, newValue)
					} else {
						fmt.Printf("%s %s: %s → %s\n", identLevel, s, oldValue, newValue)
					}
				} else {
					if negated {
						fmt.Printf("%s not %s is %s\n", identLevel, s, newValue)
					} else {
						fmt.Printf("%s %s is %s\n", identLevel, s, newValue)
					}
				}
			} else {
				if hasOldValue && oldValue != newValue {
					fmt.Printf("%s%s: %s → %s\n", identLevel, v, oldValue, newValue)
				} else {
					fmt.Printf("%sUpdate variable %s to value %s\n", identLevel, v, newValue)
				}
			}

		case *Solvable:
			// Skip internal solver variables
			if l.IsInternalVariable(event.Variable) {
				continue
			}

			v := getBase(event.Variable)
			newValue := l.Results[event.Variable]

			if event.Type == "Uncertain" {
				fmt.Printf("%sResolving variable %s to value %s (%s) \n", identLevel, v, newValue, event.Probability)
			} else {
				fmt.Printf("%sResolving variable %s to value %s\n", identLevel, v, newValue)
			}

			// Update state
			currentState[v] = newValue

		case *Message:
			fmt.Printf("%s%s\n", identLevel, event.Text)
		}
	}
	fmt.Print("\n")
}

func (l *Logger) PrintRaw() {
	fmt.Print("\n===================================\n")
	fmt.Printf("Fault found the following scenario\n")
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
