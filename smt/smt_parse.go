package smt

import (
	"fault/util"
	"fmt"
	"strings"

	"github.com/llir/llvm/ir"
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
	g.currentFunction = f.Ident()

	for _, block := range f.Blocks {
		r := g.parseBlock(block)
		rules = append(rules, r...)
	}

	return rules
}

func (g *Generator) parseBlock(block *ir.Block) []rule {
	var rules []rule
	g.currentBlock = block.Ident()
	if g.skipBlocks[g.currentBlock] != 0 {
		return rules
	}

	// For each non-branching instruction of the basic block.
	r := g.parseInstruct(block)
	rules = append(rules, r...)

	var r1 []rule
	switch term := block.Term.(type) {
	case *ir.TermCondBr:
		r1 = g.parseTermCon(term)
	default:
		stack := util.Copy(g.localCallstack)
		g.localCallstack = []string{}
		r1 = g.generateFromCallstack(stack)
	}
	rules = append(rules, r1...)
	return rules
}

func (g *Generator) parseTermCon(term *ir.TermCondBr) []rule {

	r1 := g.parseTerms(term.Succs())
	g.inPhiState = true
	id := term.Cond.Ident()
	if g.variables.isTemp(id) {
		if v, ok := g.variables.ref[id]; ok {
			r1 = g.findParseIte(r1, v)
			//r1 = append(r1, r2...)
		}
	} else if g.variables.isBolean(id) ||
		g.variables.isNumeric(id) {
		r1 = g.findParseIte(r1, &wrap{value: id})
		//r1 = append(r1, r2...)
	}
	g.inPhiState = false

	stack := util.Copy(g.localCallstack)
	g.localCallstack = []string{}
	rtemp := g.generateFromCallstack(stack)

	r1 = append(r1, rtemp...)
	return r1
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
			rules = append(rules, g.storeRule(inst)...)
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
				g.variables.ref[id] = r
				return rules
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
				g.variables.ref[id] = r
				return rules
			}

			rules = append(rules, r)
		case *ir.InstCall:
			callee := inst.Callee.Ident()
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
		case *ir.InstXor:
			r := g.xorRule(inst)
			g.tempRule(inst, r)
		default:
			panic(fmt.Sprintf("unrecognized instruction: %T", inst))

		}
	}
	return rules
}

func (g *Generator) parseTerms(terms []*ir.Block) []rule {
	var rules []rule
	//Conditionals are considered terminals
	if len(terms) > 1 { //more than one terminal == branch
		var t, f, a []rule
		var tvars, fvars map[string]string
		g.branchId = g.branchId + 1
		branch := fmt.Sprint("branch_", g.branchId)
		for _, term := range terms {
			bname := strings.Split(term.Ident(), "-")
			switch bname[len(bname)-1] {
			case "true":
				g.inPhiState = true
				branchBlock := "true"
				g.skipBlocks[term.Ident()] = 1
				t = g.parseInstruct(term)
				t = g.tagRules(t, branch, branchBlock)
				rules = append(rules, t...)
				g.inPhiState = false
			case "false":
				g.inPhiState = true
				branchBlock := "false"
				g.skipBlocks[term.Ident()] = 1
				f = g.parseInstruct(term)
				f = g.tagRules(f, branch, branchBlock)
				rules = append(rules, f...)
				g.inPhiState = false
			case "after":
				//g.skipBlocks[term.Ident()] = 1
				a = g.parseInstruct(term)
				//rules = append(rules, a...)
			default:
				panic(fmt.Sprintf("unrecognized terminal branch: %s", term.Ident()))
			}
		}
		if t != nil || f != nil {
			g.newFork()
			g.buildForkChoice(t, "true")
			g.buildForkChoice(f, "false")

			tEnds, phis := g.capCond("true", make(map[string]string))
			fEnds, _ := g.capCond("false", phis)

			// Keep variable names in sync across branches
			tSync, fSync := g.capCondSyncRules()
			tEnds = append(tEnds, tSync...)
			fEnds = append(fEnds, fSync...)

			rules = append(rules, &ite{cond: nil, t: tEnds, tvars: tvars, f: fEnds, fvars: fvars})
		}
		rules = append(rules, a...) //Because it's AFTER
	}
	if len(terms) == 1 { // Jump to that block
		var r []rule
		g.skipBlocks[terms[0].Ident()] = 1
		r = g.parseInstruct(terms[0])
		rules = append(rules, r...)
	}
	return rules
}

func (g *Generator) parseRule(id string, val string, ty string, op string) rule {
	wid := &wrap{value: id}
	wval := &wrap{value: val}
	return &infix{x: wid, ty: ty, y: wval, op: op}
}

func (g *Generator) parseInfix(id string, x string, y string, op string) rule {
	x = g.convertInfixVar(x)
	y = g.convertInfixVar(y)

	g.variables.ref[id] = g.parseRule(x, y, "", op)
	return g.variables.ref[id]
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

func (g *Generator) findParseIte(r []rule, w rule) []rule {
	for k, v := range r {
		if ite, ok := v.(*ite); ok {
			switch w.(type) {
			case *wrap:
				ite.cond = w
			default:
				ite.cond = g.parseCond(w)
			}
			r[k] = ite
		}
	}
	return r
}
