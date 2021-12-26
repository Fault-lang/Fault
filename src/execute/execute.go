package execute

import (
	"bytes"
	"errors"
	"fault/execute/parser"
	"fault/execute/solvers"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"gonum.org/v1/gonum/stat/distuv"
)

// Takes SMTLib2 and runs z3. If Uncertain types are present
// execute will calculate the odds of z3 suggested state actually
// occurring and rerun the model.

type ModelChecker struct {
	SMT         string
	Uncertains  map[string][]float64
	mode        string
	solver      map[string]*solvers.Solver
	spath       string
	lpath       []string
	branches    map[string][]string
	branchTrail map[string]map[string][]string
}

func NewModelChecker(mode string) *ModelChecker {
	abs, err := filepath.Abs("./")
	if err != nil {
		panic(err)
	}

	p := strings.Split(abs, string(os.PathSeparator))

	mc := &ModelChecker{
		mode:        mode,
		spath:       abs,
		lpath:       p,
		branches:    make(map[string][]string),
		branchTrail: make(map[string]map[string][]string),
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

func (mc *ModelChecker) LoadMeta(branches map[string][]string, trail map[string]map[string][]string) {
	// Load metadata that helps the results display nicely
	mc.branches = branches
	mc.branchTrail = trail
}

func (mc *ModelChecker) run(command string, actions []string) (string, error) {
	var path []string
	if mc.lpath[len(mc.lpath)-1] != "execute" {
		path = append([]string{}, mc.spath, "execute", mc.solver[command].Command)
	} else {
		path = append([]string{}, mc.spath, mc.solver[command].Command)
	}
	bin := filepath.Join(path...)

	cmd := exec.Command(bin,
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
