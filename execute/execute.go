package execute

import (
	"bytes"
	"errors"
	"fault/execute/parser"
	"fault/generator/scenario"
	"fault/util"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	//"gonum.org/v1/gonum/stat/distuv"
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
	Results      map[string][]*VarChange
	ResultValues map[string]string
	Log          *scenario.Logger
	solver       map[string]*Solver
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

func (mc *ModelChecker) LoadModel(smt string, uncertains map[string][]float64, unknowns []string /*results map[string][]*variables.VarChange, log *resultscenario.Logger*/) {
	mc.SMT = smt
	mc.Uncertains = uncertains
	mc.Unknowns = unknowns
	//mc.Results = results
	//mc.Log = log
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

func (mc *ModelChecker) Solve() error {
	results, err := mc.run("basic_run", []string{"(check-sat)", "(get-model)"})
	if err != nil {
		return err
	}

	// Remove extra output (ie "sat")
	results = cleanExtraOutputs(results)

	is := antlr.NewInputStream(results)
	lexer := parser.NewSMTLIBv2Lexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	p := parser.NewSMTLIBv2Parser(stream)
	l := NewSMTListener()
	antlr.ParseTreeWalkerDefault.Walk(l, p.Start_())

	mc.ResultValues = l.Values

	return err
}

func (mc *ModelChecker) PlainSolve() (string, error) {
	return mc.run("basic_run", []string{"(check-sat)", "(get-model)"})
}

type VarChange struct {
	Id     string // SSA name of var
	Parent string // SSA name of proceeding var
}

func cleanExtraOutputs(results string) string {
	for results[0:1] != "(" {
		newline := strings.Index(results, "\n")
		results = results[newline+1:]
	}
	return results
}
