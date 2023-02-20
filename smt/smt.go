package smt

import (
	"bytes"
	"fault/ast"
	"fmt"
	"strconv"
	"strings"

	"github.com/llir/llvm/asm"
	"github.com/llir/llvm/ir"
)

type Generator struct {
	currentFunction string
	currentBlock    string
	branchId        int

	// Raw input
	Uncertains map[string][]float64
	Unknowns   []string
	functions  map[string]*ir.Func
	rawAsserts []*ast.AssertionStatement
	rawAssumes []*ast.AssertionStatement
	rawRules   [][]rule

	// Generated SMT
	inits     []string
	constants []string
	rules     []string
	asserts   []string

	variables      *variables
	blocks         map[string][]rule
	localCallstack []string

	forks            []Fork
	storedChoice     map[string]*stateChange
	inPhiState       *PhiState //Flag, are we in a conditional or parallel?
	parallelGrouping string
	parallelRunStart bool      //Flag, make sure all branches with parallel runs begin from the same point
	returnVoid       *PhiState //Flag, escape parseFunc before moving to next block

	Rounds     int
	RoundVars  [][][]string
	RVarLookup map[string][][]int
	Results    map[string][]*VarChange
}

func NewGenerator() *Generator {
	return &Generator{
		variables:       NewVariables(),
		functions:       make(map[string]*ir.Func),
		blocks:          make(map[string][]rule),
		storedChoice:    make(map[string]*stateChange),
		currentFunction: "@__run",
		Uncertains:      make(map[string][]float64),
		inPhiState:      NewPhiState(),
		returnVoid:      NewPhiState(),
		Results:         make(map[string][]*VarChange),
		RVarLookup:      make(map[string][][]int),
	}
}

func (g *Generator) LoadMeta(runs int16, uncertains map[string][]float64, unknowns []string, asserts []*ast.AssertionStatement, assumes []*ast.AssertionStatement) {
	if runs == 0 {
		g.Rounds = 1 //even if runs are zero we need to generate asserts for initialization
	} else {
		g.Rounds = int(runs)
	}

	g.Uncertains = uncertains
	g.Unknowns = unknowns
	g.rawAsserts = asserts
	g.rawAssumes = assumes
}

func (g *Generator) Run(llopt string) {
	m, err := asm.ParseString("", llopt) //"/" because ParseString has a path variable
	if err != nil {
		panic(err)
	}
	g.newCallgraph(m)

}

func (g *Generator) newRound() {
	g.RoundVars = append(g.RoundVars, [][]string{})
}

func (g *Generator) initVarRound(base string, num int) {
	g.RoundVars = [][][]string{{{base, fmt.Sprint(num)}}}
}

func (g *Generator) currentRound() int {
	return len(g.RoundVars) - 1
}

func (g *Generator) addVarToRound(base string, num int) {
	if g.currentRound() == -1 {
		g.initVarRound(base, num)
		g.addVarToRoundLookup(base, num, 0, len(g.RoundVars[g.currentRound()])-1)
		return
	}

	g.RoundVars[g.currentRound()] = append(g.RoundVars[g.currentRound()], []string{base, fmt.Sprint(num)})
	g.addVarToRoundLookup(base, num, g.currentRound(), len(g.RoundVars[g.currentRound()])-1)
}

func (g *Generator) addVarToRoundLookup(base string, num int, idx int, idx2 int) {
	g.RVarLookup[base] = append(g.RVarLookup[base], []int{num, idx, idx2})
}

func (g *Generator) lookupVarRounds(base string, num string) [][]int {
	if num == "" {
		return g.RVarLookup[base]
	}

	state, err := strconv.Atoi(num)
	if err != nil {
		panic(err)
	}

	return g.lookupVarSpecificState(base, state)
}

func (g *Generator) lookupVarSpecificState(base string, state int) [][]int {
	for _, b := range g.RVarLookup[base] {
		if b[0] == state {
			return [][]int{b}
		}
	}
	panic(fmt.Errorf("state %d of variable %s is missing", state, base))
}

func (g *Generator) varRounds(base string, num string) map[int][]string {
	ir := make(map[int][]string)
	states := g.lookupVarRounds(base, num)
	for _, s := range states {
		ir[s[1]] = append(ir[s[1]], fmt.Sprintf("%s_%d", base, s[0]))
	}
	return ir
}

func (g *Generator) GetForks() []Fork {
	return g.forks
}

func (g *Generator) newConstants(globals []*ir.Global) []string {
	// Constants cannot be changed and therefore don't increment
	// in SSA. So instead of return a *rule we can skip directly
	// to a set of strings
	r := []string{}
	for _, gl := range globals {
		id := g.variables.formatIdent(gl.GlobalIdent.Ident())
		r = append(r, g.constantRule(id, gl.Init))
	}
	return r
}

func (g *Generator) newAssumes(asserts []*ast.AssertionStatement) {
	for _, v := range asserts {
		a := g.parseAssert(v)
		rule := g.writeAssert("", a)
		g.asserts = append(g.asserts, rule)
	}
}

func (g *Generator) newAsserts(asserts []*ast.AssertionStatement) {
	var arule []string
	for _, v := range asserts {
		a := g.parseAssert(v)
		arule = append(arule, a)
	}

	if len(arule) == 0 {
		return
	}

	if len(arule) > 1 {
		g.asserts = append(g.asserts, g.writeAssert("or", strings.Join(arule, "")))
	} else {
		g.asserts = append(g.asserts, g.writeAssert("", arule[0]))
	}
}

func (g *Generator) sortFuncs(funcs []*ir.Func) {
	//Iterate through all the function blocks and store them by
	// function call name.
	for _, f := range funcs {
		// Get function name.
		fname := f.Ident()

		if fname != "@__run" {
			g.functions[f.Ident()] = f
			continue
		}
	}
}

func (g *Generator) newCallgraph(m *ir.Module) {
	g.constants = g.newConstants(m.Globals)
	g.sortFuncs(m.Funcs)

	run := g.parseRunBlock(m.Funcs)
	g.rawRules = append(g.rawRules, run)

	g.rules = append(g.rules, g.generateRules()...)

	g.newAsserts(g.rawAsserts)
	g.newAssumes(g.rawAssumes)

}

func (g *Generator) generateFromCallstack(callstack []string) []rule {
	if len(callstack) == 0 {
		return []rule{}
	}

	if len(callstack) > 1 {
		//Generate parallel runs

		perm := g.parallelPermutations(callstack)
		return g.runParallel(perm)
	} else {
		fname := callstack[0]
		v := g.functions[fname]
		return g.parseFunction(v)
	}
}

////////////////////////
// Generating SMT
///////////////////////

func (g *Generator) SMT() string {
	var out bytes.Buffer

	out.WriteString(strings.Join(g.inits, "\n"))
	out.WriteString(strings.Join(g.constants, "\n"))
	out.WriteString(strings.Join(g.rules, "\n"))
	out.WriteString(strings.Join(g.asserts, "\n"))

	return out.String()
}

////////////////////////
// Temporal Logic
///////////////////////

func (g *Generator) applyTemporalLogic(temp string, ir []string, temporalFilter string, on string, off string) string {
	switch temp {
	case "eventually":
		if len(ir) > 1 {
			or := fmt.Sprintf("(%s %s)", on, strings.Join(ir, " "))
			return or
		}
		return ir[0]
	case "always":
		if len(ir) > 1 {
			or := fmt.Sprintf("(%s %s)", off, strings.Join(ir, " "))
			return or
		}
		return ir[0]
	case "eventually-always":
		if len(ir) > 1 {
			or := g.eventuallyAlways(ir)
			return or
		}
		return ir[0]
	default:
		if len(ir) > 1 {
			var op string
			switch temporalFilter {
			case "nft":
				op = "or"
			case "nmt":
				op = "or"
			default:
				op = off
			}
			or := fmt.Sprintf("(%s %s)", op, strings.Join(ir, " "))
			return or
		}
		return ir[0]
	}
}

func (g *Generator) eventuallyAlways(ir []string) string {
	var progression []string
	for i := range ir {
		s := fmt.Sprintf("(and %s)", strings.Join(ir[i:], " "))
		progression = append(progression, s)
	}
	return fmt.Sprintf("(or %s)", strings.Join(progression, " "))
}
