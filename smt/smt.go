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

func (g *Generator) addVarToRound(base string, num int) {
	idx := len(g.RoundVars) - 1
	if idx == -1 {
		idx = 0
		g.RoundVars = [][][]string{{{base, fmt.Sprint(num)}}}
	} else {
		g.RoundVars[idx] = append(g.RoundVars[idx], []string{base, fmt.Sprint(num)})
	}

	idx2 := len(g.RoundVars[idx]) - 1
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

	for _, b := range g.RVarLookup[base] {
		if b[0] == state {
			return [][]int{b}
		}
	}
	panic(fmt.Errorf("state %s of variable %s is missing", num, base))
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

	if len(g.rawAsserts) > 0 {
		var asserts []string
		for _, v := range g.rawAsserts {
			a := g.parseAssert(v)
			// if op != "&&" && op != "||" {
			// 	var aa []string
			// 	for _, assrt := range a1 {
			// 		ir := g.generateAssertRules(assrt, assrt.temporalFilter, assrt.temporalN)
			// 		aa = append(aa, g.applyTemporalLogic(v.Temporal, ir, assrt.temporalFilter, "and", "or"))
			// 	}
			// 	if len(a1) > 1 {
			// 		asserts = append(asserts, g.writeAssert("and", strings.Join(aa, " ")))
			// 	} else {
			// 		asserts = append(asserts, aa...)
			// 	}
			// } else {
			// 	asserts = append(asserts, g.generateCompound(a1, a2, op))
			// }
			asserts = append(asserts, a)
		}
		if len(asserts) > 1 {
			g.asserts = append(g.asserts, g.writeAssert("or", strings.Join(asserts, "")))
		} else {
			g.asserts = append(g.asserts, g.writeAssert("", asserts[0]))
		}
	}

	for _, v := range g.rawAssumes {
		a := g.parseAssert(v)
		// a1, a2, op := g.parseAssert(v)
		// if op != "&&" && op != "||" {
		// 	for _, assrt := range a1 {
		// 		ir := g.generateAssertRules(assrt, assrt.temporalFilter, assrt.temporalN)
		// 		assume := g.writeAssert("", g.applyTemporalLogic(v.Temporal, ir, assrt.temporalFilter, "or", "and"))
		// 		g.asserts = append(g.asserts, assume)
		// 	}
		// } else {
		// 	assume := g.writeAssert("", g.generateCompound(a1, a2, op))
		// 	g.asserts = append(g.asserts, assume)
		// }
		assume := g.writeAssert("", a)
		g.asserts = append(g.asserts, assume)
	}
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
