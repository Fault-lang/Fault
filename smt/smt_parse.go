package smt

import (
	"fault/util"
	"fmt"
	"strings"

	"github.com/llir/llvm/ir"
	irtypes "github.com/llir/llvm/ir/types"
)

func (g *Generator) parseRunBlock(fns []*ir.Func) []rule {
	var r []rule
	for _, f := range fns {
		fname := f.Ident()

		if fname != "@__run" {
			continue
		}

		r = g.parseFunction(f)
	}
	return r
}

func (g *Generator) parseFunction(f *ir.Func) []rule {
	var rules []rule

	if g.isBuiltIn(f.Ident()) {
		return rules
	}

	oldfunc := g.currentFunction
	g.currentFunction = f.Ident()

	for _, block := range f.Blocks {
		if !g.returnVoid.Check() {
			r := g.parseBlock(block)
			rules = append(rules, r...)
		}
	}

	g.returnVoid.Out()

	g.currentFunction = oldfunc
	return rules
}

func (g *Generator) parseBlock(block *ir.Block) []rule {
	var rules []rule
	oldBlock := g.currentBlock
	g.currentBlock = block.Ident()

	// For each non-branching instruction of the basic block.
	r := g.parseInstruct(block)
	rules = append(rules, r...)

	var r1 []rule
	switch term := block.Term.(type) {
	case *ir.TermCondBr:
		r1 = g.parseTermCon(term)
	case *ir.TermRet:
		stack := util.Copy(g.localCallstack)
		g.localCallstack = []string{}
		r1 = g.generateFromCallstack(stack)
		g.returnVoid.In()
	default:
		stack := util.Copy(g.localCallstack)
		g.localCallstack = []string{}
		r1 = g.generateFromCallstack(stack)
	}
	rules = append(rules, r1...)

	g.currentBlock = oldBlock
	return rules
}

func (g *Generator) parseTermCon(term *ir.TermCondBr) []rule {
	var rules []rule
	var cond rule
	var phis map[string]int16

	g.inPhiState.In()
	id := term.Cond.Ident()
	if g.variables.isTemp(id) {
		refname := fmt.Sprintf("%s-%s", g.currentFunction, id)
		if v, ok := g.variables.ref[refname]; ok {
			cond = v
		}
	} else if g.variables.isBolean(id) ||
		g.variables.isNumeric(id) {
		cond = &wrap{value: id}
	}
	g.inPhiState.Out()

	g.variables.initPhis()

	t, f, a := g.parseTerms(term.Succs())

	if !g.isBranchClosed(t, f) {
		var tEnds, fEnds []rule
		rules = append(rules, t...)
		rules = append(rules, f...)

		g.inPhiState.In() //We need to step back into a Phi state to make sure multiconditionals are handling correctly
		g.newFork()
		g.buildForkChoice(t, "true")
		g.buildForkChoice(f, "false")

		tEnds, phis = g.capCond("true", make(map[string]int16))
		fEnds, _ = g.capCond("false", phis)

		// Keep variable names in sync across branches
		tSync, fSync := g.capCondSyncRules()
		tEnds = append(tEnds, tSync...)
		fEnds = append(fEnds, fSync...)

		rules = append(rules, &ite{cond: cond, t: tEnds, f: fEnds})
		g.inPhiState.Out()
	}

	g.variables.popPhis()
	g.variables.appendState(phis)

	if a != nil {
		after := g.parseAfterBlock(a)
		rules = append(rules, after...)
	}

	return rules
}

func (g *Generator) parseAfterBlock(term *ir.Block) []rule {
	a := g.parseBlock(term)
	stack := util.Copy(g.localCallstack)
	g.localCallstack = []string{}
	a1 := g.generateFromCallstack(stack)
	a = append(a, a1...)
	return a
}

func (g *Generator) parseInstruct(block *ir.Block) []rule {
	var rules []rule
	for _, inst := range block.Insts {
		// Type switch on instruction to find call instructions.
		switch inst := inst.(type) {
		case *ir.InstAlloca:
			//Do nothing
		case *ir.InstLoad:
			g.loadsRule(inst)
		case *ir.InstStore:
			switch inst.Src.Type().(type) {
			case *irtypes.ArrayType:
				refname := fmt.Sprintf("%s-%s", g.currentFunction, inst.Dst.Ident())
				g.variables.loads[refname] = inst.Src
			default:
				rules = append(rules, g.storeRule(inst)...)
			}
		case *ir.InstFAdd:
			var r rule
			r = g.parseInfix(inst.Ident(),
				inst.X.Ident(), inst.Y.Ident(), "+")
			g.tempRule(inst, r)
		case *ir.InstFSub:
			var r rule
			r = g.parseInfix(inst.Ident(),
				inst.X.Ident(), inst.Y.Ident(), "-")
			g.tempRule(inst, r)
		case *ir.InstFMul:
			var r rule
			r = g.parseInfix(inst.Ident(),
				inst.X.Ident(), inst.Y.Ident(), "*")
			g.tempRule(inst, r)
		case *ir.InstFDiv:
			var r rule
			r = g.parseInfix(inst.Ident(),
				inst.X.Ident(), inst.Y.Ident(), "/")
			g.tempRule(inst, r)
		case *ir.InstFRem:
			//Cannot be implemented because SMT solvers do poorly with modulo
		case *ir.InstFCmp:
			var r rule
			op, y := g.parseCompare(inst.Pred.String())
			if op == "true" || op == "false" {
				r = g.parseInfix(inst.Ident(),
					inst.X.Ident(), y.(*wrap).value, op)
			} else {
				r = g.parseInfix(inst.Ident(),
					inst.X.Ident(), inst.Y.Ident(), op)
			}

			// If LLVM is storing this is a temp var
			// Happens in conditionals
			id := inst.Ident()
			if g.variables.isTemp(id) {
				refname := fmt.Sprintf("%s-%s", g.currentFunction, id)
				g.variables.ref[refname] = r
				continue
			}

			rules = append(rules, r)
		case *ir.InstICmp:
			var r rule
			op, y := g.parseCompare(inst.Pred.String())
			if op == "true" || op == "false" {
				r = g.parseInfix(inst.Ident(),
					inst.X.Ident(), y.(*wrap).value, op)
			} else {
				r = g.parseInfix(inst.Ident(),
					inst.X.Ident(), inst.Y.Ident(), op)
			}

			id := inst.Ident()
			if g.variables.isTemp(id) {
				refname := fmt.Sprintf("%s-%s", g.currentFunction, id)
				g.variables.ref[refname] = r
				continue
			}

			rules = append(rules, r)
		case *ir.InstCall:
			callee := inst.Callee.Ident()
			if g.isBuiltIn(callee) {
				g.parseBuiltIn(inst)
			}
			meta := inst.Metadata
			if g.isSameParallelGroup(meta) {
				g.localCallstack = append(g.localCallstack, callee)
			} else if g.singleParallelStep(callee) {
				stack := util.Copy(g.localCallstack)
				g.localCallstack = []string{}
				r := g.generateFromCallstack(stack)
				rules = append(rules, r...)

				r1 := g.generateFromCallstack([]string{callee})
				rules = append(rules, r1...)
			} else {
				stack := util.Copy(g.localCallstack)
				g.localCallstack = []string{}
				r := g.generateFromCallstack(stack)
				rules = append(rules, r...)

				g.localCallstack = append(g.localCallstack, callee)
			}
			g.updateParallelGroup(meta)
			g.returnVoid.Out()
		case *ir.InstXor:
			r := g.xorRule(inst)
			g.tempRule(inst, r)
		case *ir.InstAnd:
			r := g.andRule(inst)
			g.tempRule(inst, r)
		case *ir.InstOr:
			r := g.orRule(inst)
			g.tempRule(inst, r)
		case *ir.InstBitCast:
			//Do nothing
		default:
			panic(fmt.Sprintf("unrecognized instruction: %T", inst))

		}
	}
	return rules
}

func (g *Generator) parseTerms(terms []*ir.Block) ([]rule, []rule, *ir.Block) {
	var t, f []rule
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

			stack := util.Copy(g.localCallstack)
			g.localCallstack = []string{}
			t1 := g.generateFromCallstack(stack)
			t = append(t, t1...)

			t = g.tagRules(t, branch, branchBlock)
			g.inPhiState.Out()
		case "false":
			g.inPhiState.In()
			branchBlock := "false"
			f = g.parseBlock(term)

			stack := util.Copy(g.localCallstack)
			g.localCallstack = []string{}
			f1 := g.generateFromCallstack(stack)
			f = append(f, f1...)

			f = g.tagRules(f, branch, branchBlock)
			g.inPhiState.Out()
		case "after":
			a = term
		default:
			panic(fmt.Sprintf("unrecognized terminal branch: %s", term.Ident()))
		}
	}

	return t, f, a
}

func (g *Generator) parseBuiltIn(call *ir.InstCall) rule {
	p := call.Args
	bc, ok := p[0].(*ir.InstBitCast)
	if !ok {
		panic("improper argument to built in function")
	}
	id := bc.From.Ident()
	refname := fmt.Sprintf("%s-%s", g.currentFunction, id)
	state := g.variables.loads[refname]
	return g.parseRule(state.Ident(), "true", "Bool", "=")
}

func (g *Generator) isBuiltIn(c string) bool {
	if c == "@advance" || c == "@stay" {
		return true
	}
	return false
}

func (g *Generator) isBranchClosed(t []rule, f []rule) bool {
	if len(t) == 0 && len(f) == 0 {
		return true
	}
	return false
}

func (g *Generator) parseRule(id string, val string, ty string, op string) rule {
	wid := &wrap{value: id}
	wval := &wrap{value: val}
	return &infix{x: wid, ty: ty, y: wval, op: op}
}

func (g *Generator) parseInfix(id string, x string, y string, op string) rule {
	x = g.convertInfixVar(x)
	y = g.convertInfixVar(y)

	refname := fmt.Sprintf("%s-%s", g.currentFunction, id)
	g.variables.ref[refname] = g.parseRule(x, y, "", op)
	return g.variables.ref[refname]
}

func (g *Generator) parseCond(cond rule) rule {
	switch inst := cond.(type) {
	case *wrap:
		return inst
	case *infix:
		op, y := g.parseCompare(inst.op)
		inst.op = op
		if op == "true" || op == "false" {
			inst.y = y
		}
		return inst
	default:
		panic(fmt.Sprintf("Invalid conditional: %s", inst))
	}
}

func (g *Generator) parseMultiCond(id string, x rule, y rule, op string) rule {
	refname := fmt.Sprintf("%s-%s", g.currentFunction, id)
	g.variables.ref[refname] = &infix{x: x, ty: "Bool", y: y, op: op}
	return g.variables.ref[refname]
}

func (g *Generator) parseCompare(op string) (string, rule) {
	var y *wrap
	op = g.parseCompareOp(op)
	switch op {
	case "false":
		y = &wrap{value: "False"}
	case "true":
		y = &wrap{value: "True"}
	}
	return op, y
}

func (g *Generator) parseCompareOp(op string) string {
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
