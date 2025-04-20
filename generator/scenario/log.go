package scenario

import (
	"fmt"
	"strconv"
	"strings"

	"golang.org/x/exp/slices"
	"gonum.org/v1/gonum/stat/distuv"
)

type Logger struct {
	Events        []Event
	Uncertains    map[string][]float64
	BranchIndexes map[string][]int    // function_name : [event_index_1, event_index_2]
	FuncIndexes   map[string][]int    // function_name : [entry_index, exit_index]
	BranchVars    map[string][]string // function_name : [var_name_1, var_name_2]
	Forks         map[string][]string // var_name_phi : [var_name_1, var_name_2]
	Results       map[string]string   // var_name : solution value
	StringRules   map[string]string   // var_name : string rule
	IsStringRule  map[string]bool
	IsCompound    map[string]bool // Filter display of compound rules
}

func NewLogger() *Logger {
	return &Logger{
		Events:        []Event{},
		Uncertains:    make(map[string][]float64),
		Forks:         make(map[string][]string),
		Results:       make(map[string]string),
		BranchIndexes: make(map[string][]int),
		BranchVars:    make(map[string][]string),
		FuncIndexes:   make(map[string][]int),
		StringRules:   make(map[string]string),
		IsStringRule:  make(map[string]bool),
	}
}

func (l *Logger) EnterFunction(fname string, round string) {
	l.Events = append(l.Events, &FunctionCall{
		FunctionName: fname,
		Round:        round,
		Type:         "Entry",
	})
}

func (l *Logger) ExitFunction(fname string, round string) {
	l.Events = append(l.Events, &FunctionCall{
		FunctionName: fname,
		Round:        round,
		Type:         "Exit",
	})
}

func (l *Logger) UpdateVariable(variable string) {
	l.Events = append(l.Events, &VariableUpdate{
		Variable: variable,
	})
}

func (l *Logger) UpdateSolvable(variable string) {
	l.Events = append(l.Events, &Solvable{
		Variable: variable,
	})
}

func (l *Logger) AddPhiOption(phi string, end string) {
	if _, ok := l.Forks[phi]; !ok {
		l.Forks[phi] = []string{end}
		return
	}
	l.Forks[phi] = append(l.Forks[phi], end)
}

type Event interface {
	MarkDead()
	IsDead() bool
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

func (l *Logger) Kill() {
	var dead []string
	for phi, options := range l.Forks {
		phi_value := l.Results[phi]
		for i, o := range options {
			if phi_value == l.Results[o] {
				dead = append(dead, options[0:i]...)
				if i+1 < len(options) {
					dead = append(dead, options[i+1:]...)
				}
				break
			}
		}
	}

	for fname, vars := range l.BranchVars {
		for _, v := range vars {
			if slices.Contains(dead, v) {
				// Kill Variable Updates
				for _, i := range l.BranchIndexes[fname] {
					l.Events[i].MarkDead()
				}
				// Kill Function Calls themselves
				for _, i := range l.FuncIndexes[fname] {
					l.Events[i].MarkDead()
				}

				break
			}
		}
	}
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

func (l *Logger) Print() {
	fmt.Print("\n===================================\n")
	fmt.Printf("Fault found the following scenario\n")
	identLevel := ""
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
			v := getBase(event.Variable)
			s, negated := l.IsNegated(v)
			if l.IsStringRule[s] == true {
				s = l.StringRules[s]
				if negated {
					fmt.Printf("%s not %s is %s\n", identLevel, s, l.Results[event.Variable])
				} else {
					fmt.Printf("%s %s is %s\n", identLevel, s, l.Results[event.Variable])
				}
			} else {
				fmt.Printf("%sUpdate variable %s to value %s\n", identLevel, getBase(event.Variable), l.Results[event.Variable])
			}

		case *Solvable:
			if event.Type == "Uncertain" {
				fmt.Printf("%sResolving variable %s to value %s (%s) \n", identLevel, getBase(event.Variable), l.Results[event.Variable], event.Probability)
			} else {
				fmt.Printf("%sResolving variable %s to value %s\n", identLevel, getBase(event.Variable), l.Results[event.Variable])
			}
		}
	}
	fmt.Print("\n")
}
