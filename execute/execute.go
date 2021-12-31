package execute

import (
	"bytes"
	"errors"
	"fault/execute/parser"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"gonum.org/v1/gonum/stat/distuv"
)

// Takes SMTLib2 and runs z3. If Uncertain types are present
// execute will calculate the odds of z3 suggested state actually
// occurring and rerun the model.

type Solver struct {
	Command   string
	Arguments []string
}

type ModelChecker struct {
	SMT         string
	Uncertains  map[string][]float64
	solver      map[string]*Solver
	branches    map[string][]string
	branchTrail map[string]map[string][]string
}

func NewModelChecker(mode string) *ModelChecker {

	mc := &ModelChecker{
		solver:      GenerateSolver(),
		branches:    make(map[string][]string),
		branchTrail: make(map[string]map[string][]string),
	}
	return mc
}

func GenerateSolver() map[string]*Solver {
	command, _ := os.LookupEnv("SOLVERCMD")
	if command == "" {
		panic("No solver is loaded, missing SOLVERCMD")
	}

	args, _ := os.LookupEnv("SOLVERARG")
	if args == "" {
		panic("No solver is loaded, missing SOLVERARG")
	}

	s := make(map[string]*Solver)
	s["basic_run"] = &Solver{
		Command:   command,
		Arguments: []string{args},
		/*Command: "z3",
		Arguments: []string{"-in"}*/
	}
	return s
}

func (mc *ModelChecker) LoadModel(smt string, uncertains map[string][]float64) {
	mc.SMT = smt
	mc.Uncertains = uncertains
}

func (mc *ModelChecker) LoadMeta(branches map[string][]string, trail map[string]map[string][]string) {
	// Load metadata that helps the results display nicely
	mc.branches = branches
	mc.branchTrail = trail
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
	/*model := fmt.Sprint(mc.SMT, strings.Join(actions, "\n"))
	ret, err := mc.solver[command].Run(model)
	if err != nil {
		panic(err)
	}
	return strings.TrimSpace(ret), nil*/
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
