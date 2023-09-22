package execute

import (
	"bytes"
	"errors"
	"fault/execute/parser"
	"fault/smt/forks"
	"fault/smt/log"
	resultlog "fault/smt/log"
	"fault/smt/variables"
	"fault/util"
	"fmt"
	"os"
	"os/exec"
	"strconv"
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

func (mc *ModelChecker) PlainSolve() (string, error) {
	return mc.run("basic_run", []string{"(check-sat)", "(get-model)"})
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
	key := a.String()
	a.Left = mc.ConvertVars(a.Left)
	a.Right = mc.ConvertVars(a.Right)

	switch a.Op {
	case "=":
		res := mc.EvalAmbiguous(a) // Could be either bool == bool or float == float
		mc.Log.StoreEval(key, res)
		return res
	case "not":
		res := mc.EvalAmbiguous(a)
		mc.Log.StoreEval(key, res)
		return res
	case ">":
		left := mc.ConvertClause(a.Left)
		right := mc.ConvertClause(a.Right)

		res := left > right
		mc.Log.StoreEval(key, res)
		return res
	case ">=":
		left := mc.ConvertClause(a.Left)
		right := mc.ConvertClause(a.Right)

		res := left >= right
		mc.Log.StoreEval(key, res)
		return res
	case "<":
		left := mc.ConvertClause(a.Left)
		right := mc.ConvertClause(a.Right)

		res := left < right
		mc.Log.StoreEval(key, res)
		return res
	case "<=":
		left := mc.ConvertClause(a.Left)
		right := mc.ConvertClause(a.Right)

		res := left <= right
		mc.Log.StoreEval(key, res)
		return res
	case "and":
		if a.Left.Type() == "MULTI" {
			//Make sure the subclauses are handled first
			if ch, ok2 := mc.Log.AssertChains[a.Left.String()]; ok2 {
				if len(ch.Chain) > 0 {
					mc.CheckChain(ch)
				} else {
					i := mc.LookupClause(a.Left.String())
					if i < 0 {
						panic(fmt.Errorf("cannot find clause for %s", a.Left.String()))
					}
					mc.Eval(mc.Log.Asserts[i])
				}

				// Now handle the main clause
				res := mc.EvalAmbiguous(a)
				mc.Log.StoreEval(key, res)
				return res
			} else {
				panic(fmt.Errorf("cannot find clause for %s", a.Left.String()))
			}
		}

		// Or just try the normal way :)
		left, err := mc.EvalClause(a.Left)
		if err != nil {
			panic(err)
		}
		right, err := mc.EvalClause(a.Right)
		if err != nil {
			panic(err)
		}
		res := left && right
		mc.Log.StoreEval(key, res)
		return res

	case "or":
		if a.Left.Type() == "MULTI" {
			//Make sure the subclauses are handled first
			if ch, ok2 := mc.Log.AssertChains[a.Left.String()]; ok2 {
				if len(ch.Chain) > 0 {
					mc.CheckChain(ch)
				} else {
					i := mc.LookupClause(a.Left.String())
					if i < 0 {
						panic(fmt.Errorf("cannot find clause for %s", a.Left.String()))
					}
					mc.Eval(mc.Log.Asserts[i])
				}
			}
			res := mc.EvalMultiClause(a.Left.(*resultlog.MultiClause), a.Op)
			mc.Log.StoreEval(key, res)
			return res
		}

		left, err := mc.EvalClause(a.Left)
		if err != nil {
			panic(err)
		}
		right, err := mc.EvalClause(a.Right)
		if err != nil {
			panic(err)
		}
		res := left || right
		mc.Log.StoreEval(key, res)
		return res
	default:
		panic(fmt.Sprintf("no option for operator %s", a.Op))
	}
}

func (mc *ModelChecker) EvalMultiClause(m *resultlog.MultiClause, op string) bool {
	for _, v := range m.Value {
		var res bool
		var err error
		if ret, ok2 := mc.ResultValues[v]; ok2 {
			res, err = strconv.ParseBool(ret)
			if err != nil {
				panic(err)
			}
		} else {
			i := mc.LookupClause(v)
			if i < 0 {
				panic(fmt.Errorf("cannot find clause for %s", v))
			}
			res = mc.Eval(mc.Log.Asserts[i])
		}

		if res && op == "or" {
			return true
		}

		if res && op == "and" {
			return false
		}

	}

	if op == "or" {
		return false
	}

	return false
}

func (mc *ModelChecker) mixedClauseTypes(ltype string, rtype string) bool {
	if ltype == "STRING" && rtype == "BOOL" || ltype == "BOOL" && rtype == "STRING" {
		return true
	}
	return false
}

func (mc *ModelChecker) EvalMixedClauses(a *resultlog.Assert) bool {
	var left string
	var right string

	switch l := a.Left.(type) {
	case *resultlog.BoolClause:
		left = a.Left.String()
	case *resultlog.StringClause:
		left = mc.EvalStringClause(l)
	}

	switch r := a.Right.(type) {
	case *resultlog.BoolClause:
		right = a.Right.String()
	case *resultlog.StringClause:
		right = mc.EvalStringClause(r)
	}

	if a.Op == "not" {
		return left != right
	} else {
		return left == right
	}

}

func (mc *ModelChecker) EvalStringClause(c *resultlog.StringClause) string {
	var ret string
	var ok bool
	key := c.GetString()
	if ret, ok = mc.ResultValues[key]; !ok {
		if retres, ok := mc.Log.AssertClauses[key]; ok {
			ret = fmt.Sprintf("%v", retres)
		} else if retClause, ok2 := mc.Log.AssertChains[key]; ok2 {
			mc.CheckChain(retClause)
			retres := mc.Log.AssertClauses[key]
			ret = fmt.Sprintf("%v", retres)
		} else {
			panic(fmt.Sprintf("Cannot find clause %s", key))
		}
	}
	return ret
}

func (mc *ModelChecker) EvalAmbiguous(a *resultlog.Assert) bool {
	if mc.mixedClauseTypes(a.Left.Type(), a.Right.Type()) {
		return mc.EvalMixedClauses(a)
	}

	if a.Left.Type() == "INT" && a.Right.Type() == "FLOAT" {
		l := &resultlog.FlClause{}
		l.Value = mc.ConvertClause(a.Left)
		a.Left = l
	}

	if a.Left.Type() == "FLOAT" && a.Right.Type() == "INT" {
		r := &resultlog.FlClause{}
		r.Value = mc.ConvertClause(a.Right)
		a.Right = r
	}

	if a.Left.Type() != a.Right.Type() && a.Left.Type() != "MULTI" {
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
		if a.Op == "=" || a.Op == "not" {
			left := mc.EvalStringClause(a.Left.(*resultlog.StringClause))
			right := mc.EvalStringClause(a.Right.(*resultlog.StringClause))

			if a.Op == "not" {
				res = left != right
			} else {
				res = left == right
			}
		}
	case "MULTI":
		if a.Op == "and" { // ANDs every clause must be true
			var ok bool
			for _, v := range a.Left.(*log.MultiClause).Value {
				if res, ok = mc.Log.AssertClauses[v]; ok {
					if !res {
						break
					}
				} else {
					c := mc.Log.AssertChains[v]
					c.Chain = []int{mc.LookupClause(v)}
					mc.CheckChain(c)
					if res, ok = mc.Log.AssertClauses[v]; ok {
						if !res {
							break
						}
					} else {
						panic(fmt.Sprintf("missing clause %s", v))
					}
				}
			}
			res = true
		} else { //ORs only one need be true
			for _, v := range a.Left.(*log.MultiClause).Value {
				if r, ok := mc.Log.AssertClauses[v]; ok {
					if r {
						res = true
						break
					}
				} else {
					c := mc.Log.AssertChains[v]
					c.Chain = []int{mc.LookupClause(v)}
					mc.CheckChain(c)
					if res, ok = mc.Log.AssertClauses[v]; ok {
						if !res {
							break
						}
					} else {
						panic(fmt.Sprintf("missing clause %s", v))
					}
				}
			}
			res = false
		}
	}
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

		val, ok := mc.ResultValues[c.GetString()]
		if ok {
			ret, err := strconv.ParseBool(val)
			if err != nil {
				return false, fmt.Errorf("assertion on a non boolean value %s", c.GetString())
			}
			return ret, nil
		}

		ret, ok := mc.Log.AssertClauses[c.GetString()]
		if !ok {
			return false, fmt.Errorf("assertion clause %s not found", c.GetString())
		}
		return ret, nil

	default:
		return false, fmt.Errorf("illegal assertion clause %s typed %s", c.GetString(), c.Type())
	}
}

func (mc *ModelChecker) ConvertVars(a resultlog.Clause) resultlog.Clause {
	// Convert "STRING" clauses that are really idents into proper types and values
	switch a.Type() {
	case "INT":
		return a
	case "FLOAT":
		return a
	case "BOOL":
		return a
	case "STRING":
		if temp, ok := mc.ResultValues[a.String()]; ok {
			return mc.Log.NewClause(temp)
		}
	}
	return a
}

func (mc *ModelChecker) ConvertClause(a resultlog.Clause) float64 {
	var val float64
	var err error
	switch a.Type() {
	case "INT":
		val = float64(a.GetInt())
	case "FLOAT":
		val = a.GetFloat()
	case "STRING":
		temp := mc.ResultValues[a.String()]
		val, err = strconv.ParseFloat(temp, 64)
		if err != nil {
			panic(err)
		}
	}
	return val
}

func (mc *ModelChecker) LookupClause(clause string) int {
	for i, a := range mc.Log.Asserts {
		if a.String() == clause {
			return i
		}
	}
	return -1
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
