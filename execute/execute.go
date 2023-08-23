package execute

import (
	"bytes"
	"errors"
	"fault/execute/parser"
	"fault/smt/forks"
	resultlog "fault/smt/log"
	"fault/smt/rules"
	"fault/smt/variables"
	"fault/util"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
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
	SMT          string
	Uncertains   map[string][]float64
	Unknowns     []string
	Results      map[string][]*variables.VarChange
	ResultValues map[string]string
	Log          *resultlog.ResultLog
	solver       map[string]*Solver
	Forks        *forks.Fork
}

func NewModelChecker() *ModelChecker {

	mc := &ModelChecker{
		solver:       GenerateSolver(),
		ResultValues: make(map[string]string),
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

func (mc *ModelChecker) LoadModel(smt string, uncertains map[string][]float64, unknowns []string, results map[string][]*variables.VarChange, log *resultlog.ResultLog) {
	mc.SMT = smt
	mc.Uncertains = uncertains
	mc.Unknowns = unknowns
	mc.Results = results
	mc.Log = log
}

func (mc *ModelChecker) LoadMeta(frks *forks.Fork) {
	// Load metadata that helps the results display nicely
	mc.Forks = frks
}

func (mc *ModelChecker) run(command string, actions []string) (string, error) {
	cmd := exec.Command(mc.solver[command].Command,
		mc.solver[command].Arguments...)
	cmd.Stdin = strings.NewReader(fmt.Sprint(mc.SMT, strings.Join(actions, "\n")))

	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	err := cmd.Run()

	if err != nil {
		return "", fmt.Errorf(fmt.Sprint(err) + ": " + stderr.String())
	}
	return strings.TrimSpace(out.String()), err
}

func (mc *ModelChecker) Check() (bool, error) {
	results, err := mc.run("basic_run", []string{"(check-sat)"})
	if err != nil {
		return false, err
	}

	if util.FromEnd(results, 5) == "unsat" {
		return false, nil
	} else if util.FromEnd(results, 3) == "sat" {
		return true, nil
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
	results = cleanExtraOutputs(results)

	is := antlr.NewInputStream(results)
	lexer := parser.NewSMTLIBv2Lexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	p := parser.NewSMTLIBv2Parser(stream)
	l := NewSMTListener()
	antlr.ParseTreeWalkerDefault.Walk(l, p.Start())

	mc.ResultValues = l.Values

	return l.Results, err
}

func (mc *ModelChecker) Filter(results map[string]Scenario) map[string]Scenario {
	for k, uncertain := range mc.Uncertains {
		if results[k] != nil {
			dist := distuv.Normal{
				Mu:    uncertain[0],
				Sigma: uncertain[1],
			}

			results[k] = mc.stateAssessment(dist, results[k])
		}
	}
	return results
}

func (mc *ModelChecker) Eval(a *resultlog.Assert) bool {
	switch a.Op {
	case "=":
		return mc.EvalAmbiguous(a) // Could be either bool == bool or float == float
	case "not":
		return mc.EvalAmbiguous(a)
	case ">":
		left := a.Left.GetFloat()
		right := a.Right.GetFloat()
		res := left > right
		mc.Log.StoreEval(a, res)
		return res
	case ">=":
		left := a.Left.GetFloat()
		right := a.Right.GetFloat()
		res := left >= right
		mc.Log.StoreEval(a, res)
		return res
	case "<":
		left := a.Left.GetFloat()
		right := a.Right.GetFloat()
		res := left < right
		mc.Log.StoreEval(a, res)
		return res
	case "<=":
		left := a.Left.GetFloat()
		right := a.Right.GetFloat()
		res := left <= right
		mc.Log.StoreEval(a, res)
		return res
	case "and":
		left, err := mc.EvalClause(a.Left)
		if err != nil {
			panic(err)
		}
		right, err := mc.EvalClause(a.Right)
		if err != nil {
			panic(err)
		}
		res := left && right
		mc.Log.StoreEval(a, res)
		return res
	case "or":
		left, err := mc.EvalClause(a.Left)
		if err != nil {
			panic(err)
		}
		right, err := mc.EvalClause(a.Right)
		if err != nil {
			panic(err)
		}
		res := left || right
		mc.Log.StoreEval(a, res)
		return res
	default:
		panic(fmt.Sprintf("no option for operator %s", a.Op))
	}
}

func (mc *ModelChecker) EvalAmbiguous(a *resultlog.Assert) bool {
	if a.Left.Type() != a.Right.Type() {
		panic(fmt.Sprintf("improperly formatted assertion clause %s got type left %s and type right %s", a.String(), a.Left.Type(), a.Right.Type()))
	}

	var res bool
	switch a.Left.Type() {
	case "FLOAT":
		if a.Op == "=" {
			res = a.Left.GetFloat() == a.Right.GetFloat()
		}

		if a.Op == "not" {
			res = a.Left.GetFloat() != a.Right.GetFloat()
		}
	case "BOOL":
		if a.Op == "=" {
			res = a.Left.GetBool() == a.Right.GetBool()
		}

		if a.Op == "not" {
			res = a.Left.GetBool() != a.Right.GetBool()
		}
	case "STRING":
		if a.Op == "=" {
			left := mc.ResultValues[a.Left.GetString()]
			right := mc.ResultValues[a.Right.GetString()]
			res = left == right
		}

		if a.Op == "not" {
			left := mc.ResultValues[a.Left.GetString()]
			right := mc.ResultValues[a.Right.GetString()]
			res = left != right
		}
	}
	mc.Log.StoreEval(a, res)
	return res
}

func (mc *ModelChecker) EvalClause(c resultlog.Clause) (bool, error) {
	switch c.Type() {
	case "BOOL":
		return c.GetBool(), nil
	case "STRING":
		if c.String() == "" { // Happens when with clauses like (and x y z)
			return false, nil // where Left clause will be x y z and Right clause will be ""
		}

		if cl, ok := mc.Log.AssertClauses[c.GetString()]; ok {
			return cl, nil
		}
		if ch, ok2 := mc.Log.AssertChains[c.GetString()]; ok2 {
			chain := make(map[string]*rules.AssertChain)
			chain[c.GetString()] = ch
			mc.CheckAsserts(chain)
			return mc.Log.AssertClauses[c.GetString()], nil
		}

		return false, fmt.Errorf("assertion clause %s not found", c.GetString())
	default:
		return false, fmt.Errorf("illegal assertion clause %s typed %s", c.GetString(), c.Type())
	}
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

func cleanExtraOutputs(results string) string {
	for results[0:1] != "(" {
		newline := strings.Index(results, "\n")
		results = results[newline+1:]
	}
	return results
}
