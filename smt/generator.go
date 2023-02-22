package smt

import (
	"bytes"
	"fault/ast"
	"fault/smt/rules"
	"fault/smt/variables"
	"fault/util"
	"fmt"
	"strconv"
	"strings"

	"github.com/llir/llvm/asm"
	"github.com/llir/llvm/ir"
	irtypes "github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
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
	rawRules   [][]rules.Rule

	// Generated SMT
	inits     []string
	constants []string
	rules     []string
	asserts   []string

	variables      *variables.VarData
	blocks         map[string][]rules.Rule
	localCallstack []string

	forks            []Fork
	storedChoice     map[string]*rules.StateChange
	inPhiState       *PhiState //Flag, are we in a conditional or parallel?
	parallelGrouping string
	parallelRunStart bool      //Flag, make sure all branches with parallel runs begin from the same point
	returnVoid       *PhiState //Flag, escape parseFunc before moving to next block

	Rounds     int
	RoundVars  [][][]string
	RVarLookup map[string][][]int
	Results    map[string][]*variables.VarChange
}

func NewGenerator() *Generator {
	return &Generator{
		variables:       variables.NewVariables(),
		functions:       make(map[string]*ir.Func),
		blocks:          make(map[string][]rules.Rule),
		storedChoice:    make(map[string]*rules.StateChange),
		currentFunction: "@__run",
		Uncertains:      make(map[string][]float64),
		inPhiState:      NewPhiState(),
		returnVoid:      NewPhiState(),
		Results:         make(map[string][]*variables.VarChange),
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
		id := g.variables.FormatIdent(gl.GlobalIdent.Ident())
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

func (g *Generator) generateFromCallstack(callstack []string) []rules.Rule {
	if len(callstack) == 0 {
		return []rules.Rule{}
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
// Parsing LLVM IR
///////////////////////

func (g *Generator) parseRunBlock(fns []*ir.Func) []rules.Rule {
	var r []rules.Rule
	for _, f := range fns {
		fname := f.Ident()

		if fname != "@__run" {
			continue
		}

		r = g.parseFunction(f)
	}
	return r
}

func (g *Generator) parseFunction(f *ir.Func) []rules.Rule {
	var ru []rules.Rule

	if g.isBuiltIn(f.Ident()) {
		return ru
	}

	oldfunc := g.currentFunction
	g.currentFunction = f.Ident()

	for _, block := range f.Blocks {
		if !g.returnVoid.Check() {
			r := g.parseBlock(block)
			ru = append(ru, r...)
		}
	}

	g.returnVoid.Out()

	g.currentFunction = oldfunc
	return ru
}

func (g *Generator) parseBlock(block *ir.Block) []rules.Rule {
	var ru []rules.Rule
	oldBlock := g.currentBlock
	g.currentBlock = block.Ident()

	// For each non-branching instruction of the basic block.
	r := g.parseInstruct(block)
	ru = append(ru, r...)

	for k, v := range g.storedChoice {
		r0 := g.stateRules(k, v)
		ru = append(ru, r0)
	}

	//Make sure call stack is clear
	r1 := g.executeCallstack()
	ru = append(ru, r1...)

	var r2 []rules.Rule
	switch term := block.Term.(type) {
	case *ir.TermCondBr:
		r2 = g.parseTermCon(term)
	case *ir.TermRet:
		g.returnVoid.In()
	}
	ru = append(ru, r2...)

	g.currentBlock = oldBlock
	return ru
}

func (g *Generator) parseTermCon(term *ir.TermCondBr) []rules.Rule {
	var ru []rules.Rule
	var cond rules.Rule
	var phis map[string]int16

	g.inPhiState.In()
	id := term.Cond.Ident()
	if g.variables.IsTemp(id) {
		refname := fmt.Sprintf("%s-%s", g.currentFunction, id)
		if v, ok := g.variables.Ref[refname]; ok {
			cond = v
		}
	} else if g.variables.IsBoolean(id) ||
		g.variables.IsNumeric(id) {
		cond = &rules.Wrap{Value: id}
	}
	g.inPhiState.Out()

	g.variables.InitPhis()

	t, f, a := g.parseTerms(term.Succs())

	if !g.isBranchClosed(t, f) {
		var tEnds, fEnds []rules.Rule
		ru = append(ru, t...)
		ru = append(ru, f...)

		g.inPhiState.In() //We need to step back into a Phi state to make sure multiconditionals are handling correctly
		g.newFork()
		g.buildForkChoice(t, "true")
		g.buildForkChoice(f, "false")

		tEnds, phis = g.capCond("true", make(map[string]int16))
		fEnds, _ = g.capCond("false", phis)

		// Keep variable names in sync across branches
		syncs := g.capCondSyncRules([]string{"true", "false"})
		tEnds = append(tEnds, syncs["true"]...)
		fEnds = append(fEnds, syncs["false"]...)

		ru = append(ru, &rules.Ite{Cond: cond, T: tEnds, F: fEnds})
		g.inPhiState.Out()
	}

	g.variables.PopPhis()
	g.variables.AppendState(phis)

	if a != nil {
		after := g.parseAfterBlock(a)
		ru = append(ru, after...)
	}

	return ru
}

func (g *Generator) parseAfterBlock(term *ir.Block) []rules.Rule {
	a := g.parseBlock(term)
	a1 := g.executeCallstack()
	a = append(a, a1...)
	return a
}

func (g *Generator) parseInstruct(block *ir.Block) []rules.Rule {
	var ru []rules.Rule
	for _, inst := range block.Insts {
		// Type switch on instruction to find call instructions.
		switch inst := inst.(type) {
		case *ir.InstAlloca:
			//Do nothing
		case *ir.InstLoad:
			g.loadsRule(inst)
		case *ir.InstStore:
			vname := inst.Dst.Ident()
			if vname == "@__rounds" {
				//Clear the callstack first
				r := g.executeCallstack()
				ru = append(ru, r...)
				g.rawRules = append(g.rawRules, ru)
				ru = []rules.Rule{}

				//Initate new round
				g.newRound()
				continue
			}

			if vname == "@__parallelGroup" {
				continue
			}

			switch inst.Src.Type().(type) {
			case *irtypes.ArrayType:
				refname := fmt.Sprintf("%s-%s", g.currentFunction, inst.Dst.Ident())
				g.variables.Loads[refname] = inst.Src
			default:
				ru = append(ru, g.storeRule(inst)...)
			}
		case *ir.InstFAdd:
			var r rules.Rule
			r = g.createInfixRule(inst.Ident(),
				inst.X.Ident(), inst.Y.Ident(), "+")
			g.tempRule(inst, r)
		case *ir.InstFSub:
			var r rules.Rule
			r = g.createInfixRule(inst.Ident(),
				inst.X.Ident(), inst.Y.Ident(), "-")
			g.tempRule(inst, r)
		case *ir.InstFMul:
			var r rules.Rule
			r = g.createInfixRule(inst.Ident(),
				inst.X.Ident(), inst.Y.Ident(), "*")
			g.tempRule(inst, r)
		case *ir.InstFDiv:
			var r rules.Rule
			r = g.createInfixRule(inst.Ident(),
				inst.X.Ident(), inst.Y.Ident(), "/")
			g.tempRule(inst, r)
		case *ir.InstFRem:
			//Cannot be implemented because SMT solvers do poorly with modulo
		case *ir.InstFCmp:
			var r rules.Rule
			op, y := g.createCompareRule(inst.Pred.String())
			if op == "true" || op == "false" {
				r = g.createInfixRule(inst.Ident(),
					inst.X.Ident(), y.(*rules.Wrap).Value, op)
			} else {
				r = g.createInfixRule(inst.Ident(),
					inst.X.Ident(), inst.Y.Ident(), op)
			}

			// If LLVM is storing this is a temp var
			// Happens in conditionals
			id := inst.Ident()
			if g.variables.IsTemp(id) {
				refname := fmt.Sprintf("%s-%s", g.currentFunction, id)
				g.variables.Ref[refname] = r
				continue
			}

			ru = append(ru, r)
		case *ir.InstICmp:
			var r rules.Rule
			op, y := g.createCompareRule(inst.Pred.String())
			if op == "true" || op == "false" {
				r = g.createInfixRule(inst.Ident(),
					inst.X.Ident(), y.(*rules.Wrap).Value, op)
			} else {
				r = g.createInfixRule(inst.Ident(),
					inst.X.Ident(), inst.Y.Ident(), op)
			}

			id := inst.Ident()
			if g.variables.IsTemp(id) {
				refname := fmt.Sprintf("%s-%s", g.currentFunction, id)
				g.variables.Ref[refname] = r
				continue
			}

			ru = append(ru, r)
		case *ir.InstCall:
			callee := inst.Callee.Ident()
			if g.isBuiltIn(callee) {
				meta := inst.Metadata // Is this in a "b || b" construction?
				if len(meta) > 0 {
					id := inst.Ident()
					refname := fmt.Sprintf("%s-%s", g.currentFunction, id)
					inst.Metadata = nil // don't need this anymore
					g.variables.Loads[refname] = inst
				} else {
					r := g.parseBuiltIn(inst, false)
					ru = append(ru, r...)
				}
				continue
			}
			meta := inst.Metadata
			if g.isSameParallelGroup(meta) {
				g.localCallstack = append(g.localCallstack, callee)
			} else if g.singleParallelStep(callee) {
				r := g.executeCallstack()
				ru = append(ru, r...)

				r1 := g.generateFromCallstack([]string{callee})
				ru = append(ru, r1...)
			} else {
				r := g.executeCallstack()
				ru = append(ru, r...)

				g.localCallstack = append(g.localCallstack, callee)
			}
			g.updateParallelGroup(meta)
			g.returnVoid.Out()
		case *ir.InstXor:
			r := g.xorRule(inst)
			g.tempRule(inst, r)
		case *ir.InstAnd:
			if g.isStateChangeChain(inst) {
				sc := &rules.StateChange{
					Ors:  []value.Value{},
					Ands: []value.Value{},
				}
				andAd, _ := g.parseChoice(inst, sc)
				id := inst.Ident()
				refname := fmt.Sprintf("%s-%s", g.currentFunction, id)
				g.variables.Loads[refname] = inst
				g.storedChoice[refname] = andAd

			} else {
				r := g.andRule(inst)
				g.tempRule(inst, r)
			}
		case *ir.InstOr:
			if g.isStateChangeChain(inst) {
				sc := &rules.StateChange{
					Ors:  []value.Value{},
					Ands: []value.Value{},
				}
				orAd, _ := g.parseChoice(inst, sc)
				id := inst.Ident()
				refname := fmt.Sprintf("%s-%s", g.currentFunction, id)
				g.variables.Loads[refname] = inst
				g.storedChoice[refname] = orAd

			} else {
				r := g.orRule(inst)
				g.tempRule(inst, r)
			}
		case *ir.InstBitCast:
			//Do nothing
		default:
			panic(fmt.Sprintf("unrecognized instruction: %T", inst))

		}
	}
	return ru
}

func (g *Generator) parseTerms(terms []*ir.Block) ([]rules.Rule, []rules.Rule, *ir.Block) {
	var t, f []rules.Rule
	var a *ir.Block
	g.branchId = g.branchId + 1
	branch := fmt.Sprint("branch_", g.branchId)
	for _, term := range terms {
		bname := strings.Split(term.Ident(), "-")
		switch bname[len(bname)-1] {
		case "true":
			g.inPhiState.In()
			branchBlock := "true"
			t = g.parseBlock(term)

			t1 := g.executeCallstack()
			t = append(t, t1...)

			t = rules.TagRules(t, branch, branchBlock)
			g.inPhiState.Out()
		case "false":
			g.inPhiState.In()
			branchBlock := "false"
			f = g.parseBlock(term)

			g.localCallstack = []string{}
			f1 := g.executeCallstack()
			f = append(f, f1...)

			f = rules.TagRules(f, branch, branchBlock)
			g.inPhiState.Out()
		case "after":
			a = term
		default:
			panic(fmt.Sprintf("unrecognized terminal branch: %s", term.Ident()))
		}
	}

	return t, f, a
}

func (g *Generator) parseChoice(branch value.Value, sc *rules.StateChange) (*rules.StateChange, []value.Value) {
	var ret []value.Value
	switch branch := branch.(type) {
	case *ir.InstCall:
		return sc, append(ret, branch)
	case *ir.InstOr:
		refnamex := fmt.Sprintf("%s-%s", g.currentFunction, branch.X.Ident())
		vx := g.variables.Loads[refnamex]
		if g.peek(vx) != "infix" {
			sc, ret = g.parseChoice(vx, sc)
			sc.Ors = append(sc.Ors, ret...)
		} else {
			sc2 := g.storedChoice[refnamex]
			sc.Ands = append(sc.Ands, sc2.Ands...)
			sc.Ors = append(sc.Ors, sc2.Ors...)
		}
		delete(g.storedChoice, refnamex)

		refnamey := fmt.Sprintf("%s-%s", g.currentFunction, branch.Y.Ident())
		vy := g.variables.Loads[refnamey]
		if g.peek(vy) != "infix" {
			sc, ret = g.parseChoice(vy, sc)
			sc.Ors = append(sc.Ors, ret...)
		} else {
			sc2 := g.storedChoice[refnamey]
			sc.Ands = append(sc.Ands, sc2.Ands...)
			sc.Ors = append(sc.Ors, sc2.Ors...)
		}
		delete(g.storedChoice, refnamey)

		return sc, ret
	case *ir.InstAnd:
		refnamex := fmt.Sprintf("%s-%s", g.currentFunction, branch.X.Ident())
		vx := g.variables.Loads[refnamex]
		if g.peek(vx) != "infix" {
			sc, ret = g.parseChoice(vx, sc)
			sc.Ands = append(sc.Ands, ret...)
		} else {
			sc2 := g.storedChoice[refnamex]
			sc.Ands = append(sc.Ands, sc2.Ands...)
			sc.Ors = append(sc.Ors, sc2.Ors...)
		}
		delete(g.storedChoice, refnamex)

		refnamey := fmt.Sprintf("%s-%s", g.currentFunction, branch.Y.Ident())
		vy := g.variables.Loads[refnamey]
		if g.peek(vy) != "infix" {
			sc, ret = g.parseChoice(vy, sc)
			sc.Ands = append(sc.Ands, ret...)
		} else {
			sc2 := g.storedChoice[refnamey]
			sc.Ands = append(sc.Ands, sc2.Ands...)
			sc.Ors = append(sc.Ors, sc2.Ors...)
		}
		delete(g.storedChoice, refnamey)

		return sc, ret
	}
	return sc, ret
}

func (g *Generator) parseBuiltIn(call *ir.InstCall, complex bool) []rules.Rule {
	p := call.Args
	if len(p) == 0 {
		return []rules.Rule{}
	}

	bc, ok := p[0].(*ir.InstBitCast)
	if !ok {
		panic("improper argument to built in function")
	}

	id := bc.From.Ident()
	refname := fmt.Sprintf("%s-%s", g.currentFunction, id)
	state := g.variables.Loads[refname]
	newState := state.Ident()
	base := newState[2 : len(newState)-1] //Because this is a charArray LLVM adds c"..." formatting we need to remove
	n := g.variables.SSA[base]
	prev := fmt.Sprintf("%s_%d", base, n)
	if !g.inPhiState.Check() {
		g.variables.NewPhi(base, n+1)
	} else {
		g.variables.StoreLastState(base, n+1)
	}
	g.addVarToRound(base, int(n+1))
	newState = g.variables.AdvanceSSA(base)
	g.AddNewVarChange(base, newState, prev)

	if complex {
		g.declareVar(newState, "Bool")
	}
	r1 := g.createRule(newState, "true", "Bool", "=")

	if g.currentFunction[len(g.currentFunction)-7:] != "__state" {
		panic("calling advance from outside the state chart")
	}

	base2 := g.currentFunction[1 : len(g.currentFunction)-7]
	n2 := g.variables.SSA[base2]
	prev2 := fmt.Sprintf("%s_%d", base2, n2)
	if !g.inPhiState.Check() {
		g.variables.NewPhi(base2, n2+1)
	} else {
		g.variables.StoreLastState(base2, n2+1)
	}

	g.addVarToRound(base2, int(n2+1))
	currentState := g.variables.AdvanceSSA(base2)
	g.AddNewVarChange(base2, currentState, prev2)
	if complex {
		g.declareVar(currentState, "Bool")
	}
	r2 := g.createRule(currentState, "false", "Bool", "=")
	return []rules.Rule{r1, r2}
}

func (g *Generator) isBuiltIn(c string) bool {
	if c == "@advance" || c == "@stay" {
		return true
	}
	return false
}

func (g *Generator) isBranchClosed(t []rules.Rule, f []rules.Rule) bool {
	if len(t) == 0 && len(f) == 0 {
		return true
	}
	return false
}

func (g *Generator) createRule(id string, val string, ty string, op string) rules.Rule {
	wid := &rules.Wrap{Value: id}
	wval := &rules.Wrap{Value: val}
	return &rules.Infix{X: wid, Ty: ty, Y: wval, Op: op}
}

func (g *Generator) createInfixRule(id string, x string, y string, op string) rules.Rule {
	x = g.convertInfixVar(x)
	y = g.convertInfixVar(y)

	refname := fmt.Sprintf("%s-%s", g.currentFunction, id)
	g.variables.Ref[refname] = g.createRule(x, y, "", op)
	return g.variables.Ref[refname]
}

func (g *Generator) createCondRule(cond rules.Rule) rules.Rule {
	switch inst := cond.(type) {
	case *rules.Wrap:
		return inst
	case *rules.Infix:
		op, y := g.createCompareRule(inst.Op)
		inst.Op = op
		if op == "true" || op == "false" {
			inst.Y = y
		}
		return inst
	default:
		panic(fmt.Sprintf("Invalid conditional: %s", inst))
	}
}

func (g *Generator) createMultiCondRule(id string, x rules.Rule, y rules.Rule, op string) rules.Rule {
	refname := fmt.Sprintf("%s-%s", g.currentFunction, id)
	g.variables.Ref[refname] = &rules.Infix{X: x, Ty: "Bool", Y: y, Op: op}
	return g.variables.Ref[refname]
}

func (g *Generator) createCompareRule(op string) (string, rules.Rule) {
	var y *rules.Wrap
	op = g.compareRuleOp(op)
	switch op {
	case "false":
		y = &rules.Wrap{Value: "False"}
	case "true":
		y = &rules.Wrap{Value: "True"}
	}
	return op, y
}

func (g *Generator) compareRuleOp(op string) string {
	switch op {
	case "false":
		return "false"
	case "oeq":
		return "="
	case "eq":
		return "="
	case "oge":
		return ">="
	case "ogt":
		return ">"
	case "ole":
		return "<="
	case "olt":
		return "<"
	case "one":
		return "!="
	case "ne":
		return "!="
	case "true":
		return "true"
	case "ueq":
		return "="
	case "uge":
		return ">="
	case "ugt":
		return ">"
	case "ule":
		return "<="
	case "ult":
		return "<"
	case "une":
		return "!="
	default:
		return op
	}
}

func (g *Generator) executeCallstack() []rules.Rule {
	stack := util.Copy(g.localCallstack)
	g.localCallstack = []string{}
	r := g.generateFromCallstack(stack)
	return r
}

func (g *Generator) peek(inst value.Value) string {
	switch inst.(type) {
	case *ir.InstOr, *ir.InstAnd:
		return "infix"
	case *ir.InstCall:
		return "call"
	default:
		panic("unsupported instruction type")
	}
}

func (g *Generator) isStateChangeChain(inst ir.Instruction) bool {
	switch inst := inst.(type) {
	case *ir.InstAnd:
		if !g.variables.IsTemp(inst.X.Ident()) {
			return false
		}

		switch inst.X.(type) {
		case *ir.InstCall, *ir.InstAnd, *ir.InstOr:
		default:
			return false
		}

		if !g.variables.IsTemp(inst.Y.Ident()) {
			return false
		}

		switch inst.Y.(type) {
		case *ir.InstCall, *ir.InstAnd, *ir.InstOr:
		default:
			return false
		}

	case *ir.InstOr:
		if !g.variables.IsTemp(inst.X.Ident()) {
			return false
		}

		switch inst.X.(type) {
		case *ir.InstCall, *ir.InstAnd, *ir.InstOr:
		default:
			return false
		}

		if !g.variables.IsTemp(inst.Y.Ident()) {
			return false
		}

		switch inst.Y.(type) {
		case *ir.InstCall, *ir.InstAnd, *ir.InstOr:
		default:
			return false
		}

	default:
		return false
	}
	return true
}

////////////////////////
// Some functions specific to variable names in rules
////////////////////////

func (g *Generator) AddNewVarChange(base string, id string, parent string) {
	var v *variables.VarChange
	if id == parent {
		v = &variables.VarChange{Id: id, Parent: ""}
	} else {
		v = &variables.VarChange{Id: id, Parent: parent}
	}

	if len(g.Results[base]) == 0 {
		g.Results[base] = append(g.Results[base], v)
	} else {
		g.Results[base] = append(g.Results[base], v)
	}
}

func (g *Generator) VarChangePhi(base string, end string, nums []int16) {
	for _, n := range nums {
		start := fmt.Sprintf("%s_%d", base, n)
		g.AddNewVarChange(base, end, start)
	}
}

func (g *Generator) tempToIdent(ru rules.Rule) rules.Rule {
	switch r := ru.(type) {
	case *rules.Wrap:
		return g.fetchIdent(r.Value, r)
	case *rules.Infix:
		r.X = g.tempToIdent(r.X)
		r.Y = g.tempToIdent(r.Y)
		return r
	}
	return ru
}

func (g *Generator) fetchIdent(id string, r rules.Rule) rules.Rule {
	if g.variables.IsTemp(id) {
		refname := fmt.Sprintf("%s-%s", g.currentFunction, id)
		if v, ok := g.variables.Loads[refname]; ok {
			n := g.variables.SSA[id]
			if !g.inPhiState.Check() {
				g.variables.NewPhi(id, n+1)
			} else {
				g.variables.StoreLastState(id, n+1)
			}
			g.addVarToRound(id, int(n+1))
			id = g.variables.AdvanceSSA(v.Ident())
			wid := &rules.Wrap{Value: id}
			return wid
		} else if ref, ok := g.variables.Ref[refname]; ok {
			switch r := ref.(type) {
			case *rules.Infix:
				r.X = g.tempToIdent(r.X)
				r.Y = g.tempToIdent(r.Y)
				return r
			}
		} else {
			panic(fmt.Sprintf("smt generation error, value for %s not found", id))
		}
	}
	return r
}

func (g *Generator) convertInfixVar(x string) string {
	if g.variables.IsTemp(x) {
		refname := fmt.Sprintf("%s-%s", g.currentFunction, x)
		if v, ok := g.variables.Loads[refname]; ok {
			xid := v.Ident()
			xidNoPercent := g.variables.FormatIdent(xid)
			if g.parallelRunStart {
				n := g.variables.GetStartState(xidNoPercent)
				x = fmt.Sprintf("%s_%d", xidNoPercent, n)
				g.parallelRunStart = false
			} else {
				x = g.variables.GetSSA(xidNoPercent)
			}
		}
	}
	return x
}
func (g *Generator) isASolvable(id string) bool {
	id, _ = g.variables.GetVarBase(id)
	for _, v := range g.Unknowns {
		if v == id {
			return true
		}
	}
	for k := range g.Uncertains {
		if k == id {
			return true
		}
	}
	return false
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
