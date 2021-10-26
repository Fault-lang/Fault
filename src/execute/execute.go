package execute

import (
	"bytes"
	"errors"
	"fault/execute/parser"
	"fault/execute/solvers"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/olekukonko/tablewriter"
	"gonum.org/v1/gonum/stat/distuv"
)

// Takes SMTLib2 and runs z3. If Uncertain types are present
// execute will calculate the odds of z3 suggested state actually
// occurring and rerun the model.

type ModelChecker struct {
	SMT        string
	Uncertains map[string][]float64
	mode       string
	solver     map[string]*solvers.Solver
}

func NewModelChecker(mode string) *ModelChecker {
	mc := &ModelChecker{
		mode: mode,
	}
	switch mc.mode { // Possible support for different smt solvers
	case "z3":
		mc.solver = solvers.Z3()
	}
	return mc
}

func (mc *ModelChecker) LoadModel(smt string, uncertains map[string][]float64) {
	mc.SMT = smt
	mc.Uncertains = uncertains
}

func (mc *ModelChecker) run(command string, actions []string) (string, error) {
	cmd := exec.Command(mc.solver[command].Command,
		mc.solver[command].Arguments...)

	cmd.Stdin = strings.NewReader(fmt.Sprint(mc.SMT, strings.Join(actions, "\n")))

	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()

	if err != nil {
		return "", err
	}
	return strings.TrimSpace(out.String()), err
}

func (mc *ModelChecker) Check() (bool, error) {
	results, err := mc.run("basic_run", []string{"(check-sat)"})
	if err != nil {
		return false, err
	}

	if results == "sat" {
		return true, nil
	} else if results == "unsat" {
		return false, nil
	} else {
		return false, errors.New(results)
	}
}

func (mc *ModelChecker) Solve() (map[string]Scenario, error) {
	results, err := mc.run("basic_run", []string{"(check-sat)", "(get-model)"})
	if err != nil {
		return nil, err
	}
	// Remove extra output (ie "sat")
	if results[0:1] != "(" {
		newline := strings.Index(results, "\n")
		results = results[newline:]
	}

	is := antlr.NewInputStream(results)
	lexer := parser.NewSMTLIBv2Lexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	p := parser.NewSMTLIBv2Parser(stream)
	l := NewSMTListener()
	antlr.ParseTreeWalkerDefault.Walk(l, p.Start())

	return l.Results, err
}

func (mc *ModelChecker) Filter(results map[string]Scenario) map[string]Scenario {
	likelihood := make(map[string]Scenario)
	if len(mc.Uncertains) != 0 {
		for k, uncertain := range mc.Uncertains {
			if results[k] != nil {
				dist := distuv.Normal{
					Mu:    uncertain[0],
					Sigma: uncertain[1],
				}

				likelihood[k] = mc.stateAssessment(dist, results[k])
			}
		}
		return likelihood
	}
	return results
}

func (mc *ModelChecker) stateAssessment(dist distuv.Normal, states Scenario) Scenario {
	var weighted Scenario
	switch s := states.(type) {
	case *FloatTrace:
		weighted = NewFloatTrace()
		weighted.(*FloatTrace).results = s.results
		for i, state := range s.results {
			weighted.(*FloatTrace).AddWeight(i, dist.Prob(state))
		}
	case *IntTrace:
		weighted = NewIntTrace()
		weighted.(*IntTrace).results = s.results
		for i, state := range s.results {
			weighted.(*IntTrace).AddWeight(i, dist.Prob(float64(state)))
		}
	case *BoolTrace:
		//Requires Gaussian distr, TODO LATER
		/*weighted = NewBoolTrace()
		weighted.(*BoolTrace).results = s.results
		for i, state := range s.results {
			weighted.(*BoolTrace).AddWeight(i, dist.Prob(state))
		}*/
	}
	return weighted
}

func (mc *ModelChecker) Format(results map[string]Scenario) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Variable", "State (Weight)"})
	var row []string
	for k, v := range results {
		row = append(row, k)
		switch s := v.(type) {
		case *FloatTrace:
			var r []string
			weights := s.GetWeights()
			for i, n := range s.Get() {
				if int(i) == len(r) {
					r = append(r, hWToString(n, i, weights))
				} else if int(i) > len(r) {
					for j := len(r) - 1; j < int(i)-1; j++ {
						r = append(r, "[branch]")
					}
					r = append(r, hWToString(n, i, weights))

				} else if int(i) < len(r) {
					r2 := r[0:i]
					r2 = append(r2, hWToString(n, i, weights))
					r = append(r2, r[i+1:]...)
				}
			}
			row = append(row, strings.Join(r, " "))
		case *IntTrace:
			var r []string
			weights := s.GetWeights()
			for i, n := range s.Get() {
				if int(i) == len(r) {
					r = append(r, hWToString(n, i, weights))
				} else if int(i) > len(r) {
					for j := len(r) - 1; j < int(i); j++ {
						r = append(r, "[branch]")
					}
					r = append(r, hWToString(n, i, weights))

				} else if int(i) < len(r) {
					r[i] = hWToString(n, i, weights)
				}
			}
			row = append(row, strings.Join(r, " "))
		case *BoolTrace:
			var r []string
			weights := s.GetWeights()
			for i, n := range s.Get() {
				if int(i) == len(r) {
					r = append(r, hWToString(n, i, weights))
				} else if int(i) > len(r) {
					for j := len(r) - 1; j < int(i); j++ {
						r = append(r, "[branch]")
					}
					r = append(r, hWToString(n, i, weights))

				} else if int(i) < len(r) {
					r[i] = hWToString(n, i, weights)
				}
			}
			row = append(row, strings.Join(r, " "))
		}
		table.Append(row)
	}
	table.Render()
	return
}

func hWToString(n interface{}, i int64, weights map[int64]float64) string {
	switch h := n.(type) {
	case float64:
		if val, ok := weights[i]; ok {
			return fmt.Sprintf("-> %f (%f)", h, val)
		} else {
			return fmt.Sprintf("-> %f", h)
		}
	case int64:
		if val, ok := weights[i]; ok {
			return fmt.Sprintf("-> %d (%f)", h, val)
		} else {
			return fmt.Sprintf("-> %d", h)
		}
	case bool:
		if val, ok := weights[i]; ok {
			return fmt.Sprintf("-> %v (%f)", h, val)
		} else {
			return fmt.Sprintf("-> %v", h)
		}
	default:
		panic(fmt.Sprintf("type %T not allowed", n))
	}
}
