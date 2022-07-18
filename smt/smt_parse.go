package smt

import (
	"fmt"
	"strings"

	"github.com/llir/llvm/ir"
)

func (g *Generator) parseFunction(f *ir.Func) []rule {
	var rules []rule
	g.currentFunction = f.Ident()
	for _, block := range f.Blocks {
		g.currentBlock = block.Ident()
		if g.skipBlocks[g.currentBlock] == 0 {
			// For each non-branching instruction of the basic block.
			r := g.parseInstruct(block)
			rules = append(rules, r...)
			r1 := g.parseTerms(block.Term.Succs())
			switch term := block.Term.(type) {
			case *ir.TermCondBr:
				g.inPhiState = true
				id := term.Cond.Ident()
				if g.variables.isTemp(id) {
					if v, ok := g.variables.ref[id]; ok {
						r1[len(r1)-1].(*ite).cond = g.parseCond(v)
					}
				}

				g.inPhiState = false
			}
			rules = append(rules, r1...)
		}
	}
	return rules
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
			rules = g.storeRule(inst, rules)
			g.blocks[g.currentBlock] = rules
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
		case *ir.InstCall:
			callee := g.callRule(inst)
			g.callstack[g.call] = append(g.callstack[g.call], callee)
		default:
			panic(fmt.Sprintf("unrecognized instruction: %s", inst))

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

			tEnds := g.capCond("true")
			fEnds := g.capCond("false")

			// Keep variable names in sync across branches
			// tSync, fSync := g.capCondSyncRules()
			// tEnds = append(tEnds, tSync...)
			// fEnds = append(fEnds, fSync...)

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
