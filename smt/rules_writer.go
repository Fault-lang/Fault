package smt

import (
	"fmt"
	"strings"
)

func (g *Generator) writeInfix(x string, y string, op string) string {
	return fmt.Sprintf("(%s %s %s)", op, x, y)
}

func (g *Generator) writeBranch(ty string, cond string, t string, f string) string {
	return fmt.Sprintf("(%s %s %s %s)", ty, cond, t, f)
}

func (g *Generator) declareVar(id string, t string) {
	def := fmt.Sprintf("(declare-fun %s () %s)", id, t)
	g.inits = append(g.inits, def)
}
func (g *Generator) writeAssert(op string, stmt string) string {
	if op == "" {
		return fmt.Sprintf("(assert %s)", stmt)
	}
	return fmt.Sprintf("(assert (%s %s))", op, stmt)
}

func (g *Generator) writeBranchRule(r *infix) string {
	y := g.unpackRule(r.y)
	x := g.unpackRule(r.x)

	return fmt.Sprintf("(%s %s %s)", r.op, x, y)
}

func (g *Generator) writeRule(ru rule) string {
	switch r := ru.(type) {
	case *infix:
		y := g.unpackRule(r.y)
		x := g.unpackRule(r.x)

		if y == "0x3DA3CA8CB153A753" { //An uncertain or unknown value
			g.declareVar(x, r.ty)
			return ""
		}

		if r.op != "" && r.op != "=" {
			return g.writeInfix(x, y, r.op)
		}

		return g.writeInitRule(x, r.ty, y)
	case *ite:
		cond := g.writeCond(r.cond.(*infix))
		var tRule, fRule string
		var tEnds, fEnds []string
		for _, t := range r.t {
			tEnds = append(tEnds, g.writeBranchRule(t.(*infix)))
		}

		for _, f := range r.f {
			fEnds = append(fEnds, g.writeBranchRule(f.(*infix)))
		}

		if len(tEnds) > 1 {
			tRule = fmt.Sprintf("(and %s)", strings.Join(tEnds, " "))
		} else if len(tEnds) == 1 {
			tRule = tEnds[0]
		}

		if len(fEnds) > 1 {
			fRule = fmt.Sprintf("(and %s)", strings.Join(fEnds, " "))
		} else if len(fEnds) == 1 {
			fRule = fEnds[0]
		}

		br := g.writeBranch("ite", cond, tRule, fRule)
		return g.writeAssert("", br)
	case *wrap:
		return r.value
	default:
		panic(fmt.Sprintf("%T is not a valid rule type", r))
	}
}

func (g *Generator) writeCond(r *infix) string {
	y := g.unpackRule(r.y)
	x := g.unpackRule(r.x)

	return g.writeInfix(x, y, r.op)
}

func (g *Generator) unpackRule(x rule) string {
	switch r := x.(type) {
	case *wrap:
		return r.value
	case *infix:
		return g.writeRule(r)
	default:
		panic(fmt.Sprintf("%T is not a valid rule type", r))
	}
}

func (g *Generator) writeInitRule(id string, t string, val string) string {
	// Initialize: x = Int("x")
	g.declareVar(id, t)
	// Set rule: s.add(x == 2)
	return fmt.Sprintf("(assert (= %s %s))", id, val)
}

func (g *Generator) generateRules(raw []rule) []string {
	var rules []string
	for _, v := range raw {
		rules = append(rules, g.writeRule(v))
	}
	return rules
}
