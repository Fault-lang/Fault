package smt

import (
	"fmt"
	"strings"

	"github.com/llir/llvm/ir"
)

func (g *Generator) parseFunction(f *ir.Func, startVars map[string]string) []rule {
	var rules []rule
	g.currentFunction = f.Ident()
	for _, block := range f.Blocks {
		g.currentBlock = block.Ident()
		if g.skipBlocks[g.currentBlock] == 0 {
			// For each non-branching instruction of the basic block.
			r, sv := g.parseInstruct(block, startVars)
			rules = append(rules, r...)
			startVars = sv
			r1, _ := g.parseTerms(block.Term.Succs(), startVars)
			switch term := block.Term.(type) {
			case *ir.TermCondBr:
				id := term.Cond.Ident()
				if g.isTemp(id) {
					if v, ok := g.ref[id]; ok {
						r1[len(r1)-1].(*ite).cond = g.parseCond(v, startVars)
					}
				}

			}
			rules = append(rules, r1...)
		}

	}
	return rules
}

func (g *Generator) parseInstruct(block *ir.Block, startVars map[string]string) ([]rule, map[string]string) {
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
			r, startVars = g.parseInfix(inst.Ident(),
				inst.X.Ident(), inst.Y.Ident(), "+", startVars)
			g.tempRule(inst, r)
		case *ir.InstFSub:
			var r rule
			r, startVars = g.parseInfix(inst.Ident(),
				inst.X.Ident(), inst.Y.Ident(), "-", startVars)
			g.tempRule(inst, r)
		case *ir.InstFMul:
			var r rule
			r, startVars = g.parseInfix(inst.Ident(),
				inst.X.Ident(), inst.Y.Ident(), "*", startVars)
			g.tempRule(inst, r)
		case *ir.InstFDiv:
			var r rule
			r, startVars = g.parseInfix(inst.Ident(),
				inst.X.Ident(), inst.Y.Ident(), "/", startVars)
			g.tempRule(inst, r)
		case *ir.InstFRem:
			//Cannot be implemented because SMT solvers do poorly with modulo
		case *ir.InstFCmp:
			var r rule
			op, y := g.parseCompare(inst.Pred.String())
			if op == "true" || op == "false" {
				r, startVars = g.parseInfix(inst.Ident(),
					inst.X.Ident(), y.(*wrap).value, op, startVars)
			} else {
				r, startVars = g.parseInfix(inst.Ident(),
					inst.X.Ident(), inst.Y.Ident(), op, startVars)
			}

			// If LLVM is storing this is a temp var
			// Happens in conditionals
			id := inst.Ident()
			if g.isTemp(id) {
				g.ref[id] = r
				return rules, startVars
			}
			rules = append(rules, r)
		case *ir.InstCall:
			callee := g.callRule(inst)
			g.callstack[g.call] = append(g.callstack[g.call], callee)
		default:
			fmt.Printf("%T", inst)
		}
	}
	return rules, startVars
}

func (g *Generator) parseTerms(terms []*ir.Block, startVars map[string]string) ([]rule, map[string]string) {
	var rules []rule
	var sv map[string]string
	//Conditionals are considered terminals
	if len(terms) > 1 { //more than one terminal == branch
		var t, f []rule
		var tvars, fvars map[string]string
		g.branchId = g.branchId + 1
		branch := fmt.Sprint("branch_", g.branchId)
		for _, term := range terms {
			bname := strings.Split(term.Ident(), "-")
			switch bname[len(bname)-1] {
			case "true":
				branchBlock := "true"
				g.skipBlocks[term.Ident()] = 1
				sv = make(map[string]string)
				t, tvars = g.parseInstruct(term, startVars)
				t = g.tagRules(t, branch, branchBlock)
				rules = append(rules, t...)
			case "false":
				branchBlock := "false"
				g.skipBlocks[term.Ident()] = 1
				sv = make(map[string]string)
				f, fvars = g.parseInstruct(term, startVars)
				f = g.tagRules(f, branch, branchBlock)
				rules = append(rules, f...)
			case "after":
				g.skipBlocks[term.Ident()] = 1
			default:
				panic(fmt.Sprintf("unrecognized terminal branch: %s", term.Ident()))
			}
		}
		rules = append(rules, &ite{cond: nil, t: t, tvars: tvars, f: f, fvars: fvars})
	}
	if len(terms) == 1 { // Jump to that block
		var r []rule
		sv = make(map[string]string)
		g.skipBlocks[terms[0].Ident()] = 1
		r, sv = g.parseInstruct(terms[0], startVars)
		rules = append(rules, r...)
	}
	g.last = nil
	return rules, sv
}

func (g *Generator) parseInfix(id string, x string, y string, op string, startVars map[string]string) (rule, map[string]string) {
	if g.isTemp(x) {
		if v, ok := g.loads[x]; ok {
			xid := v.Ident()
			xidNoPercent := g.formatIdent(xid)
			if id, ok := startVars[xidNoPercent]; ok {
				x = g.formatIdent(id)
				delete(startVars, xidNoPercent)
			} else {
				x = g.convertIdent(xid)
			}
		}
	}

	if g.isTemp(y) {
		if v, ok := g.loads[y]; ok {
			yid := v.Ident()
			yidNoPercent := g.formatIdent(yid)
			if id, ok := startVars[yidNoPercent]; ok {
				y = g.formatIdent(id)
				delete(startVars, yidNoPercent)
			} else {
				y = g.convertIdent(yid)
			}
		}
	}
	g.ref[id] = g.parseRule(x, y, "", op)
	g.last = g.ref[id]
	return g.ref[id], startVars
}

func (g *Generator) parseCond(cond rule, startVars map[string]string) rule {
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
		return "="
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
		return "="
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

func (g *Generator) parseRule(id string, val string, ty string, op string) rule {
	wid := &wrap{value: id}
	wval := &wrap{value: val}
	return &infix{x: wid, ty: ty, y: wval, op: op}
}

func (g *Generator) tagRules(rules []rule, branch string, block string) []rule {
	var tagged []rule
	for i := 0; i < len(rules); i++ {
		tagged = append(tagged, g.tagRule(rules[i], branch, block))
	}
	return tagged
}

func (g *Generator) tagRule(ru rule, branch string, block string) rule {
	switch r := ru.(type) {
	case *infix:
		r.x = g.tagRule(r.x, branch, block)
		r.y = g.tagRule(r.y, branch, block)
		r.Tag(branch, block)
		return r
	case *ite:
		r.cond = g.tagRule(r.cond, branch, block)
		r.t = g.tagRules(r.t, branch, block)
		r.f = g.tagRules(r.f, branch, block)
		r.Tag(branch, block)
		return r
	case *wrap:
		r.Tag(branch, block)
		return r
	case *vwrap:
		r.Tag(branch, block)
		return r
	default:
		panic(fmt.Sprintf("%T is not a valid rule type", ru))
	}
}
