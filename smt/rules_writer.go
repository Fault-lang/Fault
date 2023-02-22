package smt

import (
	"fault/smt/rules"
	"fmt"
	"strings"
)

func (g *Generator) writeAssertlessRule(op string, x string, y string) string {
	if y != "" {
		return fmt.Sprintf("(%s %s %s)", op, x, y)
	} else {
		return fmt.Sprintf("(%s %s)", op, x)
	}
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

func (g *Generator) writeBranchRule(r *rules.Infix) string {
	y := g.unpackRule(r.Y)
	x := g.unpackRule(r.X)

	return fmt.Sprintf("(%s %s %s)", r.Op, x, y)
}

func (g *Generator) writeRule(ru rules.Rule) string {
	switch r := ru.(type) {
	case *rules.Infix:
		y := g.unpackRule(r.Y)
		x := g.unpackRule(r.X)

		if y == "0x3DA3CA8CB153A753" { //An uncertain or unknown value
			g.declareVar(x, r.Ty)
			return ""
		}

		if r.Op == "or" {
			stmt := fmt.Sprintf("%s%s", x, y)
			return g.writeAssert("or", stmt)
		}

		if r.Op != "" && r.Op != "=" {
			return g.writeAssertlessRule(r.Op, x, y)
		}

		return g.writeInitRule(x, r.Ty, y)
	case *rules.Ite:
		cond := g.writeCond(r.Cond.(*rules.Infix))
		var tRule, fRule string
		var tEnds, fEnds []string
		for _, t := range r.T {
			tEnds = append(tEnds, g.writeBranchRule(t.(*rules.Infix)))
		}

		for _, f := range r.F {
			fEnds = append(fEnds, g.writeBranchRule(f.(*rules.Infix)))
		}

		if len(tEnds) > 1 {
			stmt := strings.Join(tEnds, " ")
			tRule = g.writeAssertlessRule("and", stmt, "")
		} else if len(tEnds) == 1 {
			tRule = tEnds[0]
		}

		if len(fEnds) > 1 {
			stmt := strings.Join(fEnds, " ")
			fRule = g.writeAssertlessRule("and", stmt, "")
		} else if len(fEnds) == 1 {
			fRule = fEnds[0]
		}

		br := g.writeBranch("ite", cond, tRule, fRule)
		return g.writeAssert("", br)
	case *rules.Wrap:
		return r.Value
	case *rules.Phi:
		g.declareVar(r.EndState, g.variables.lookupType(r.BaseVar, nil))
		ends := g.formatEnds(r.BaseVar, r.Nums, r.EndState)
		return g.writeAssert("or", ends)
	case *rules.Ands:
		var ands string
		for _, x := range r.X {
			var s string
			switch x := x.(type) {
			case *rules.Infix:
				s = g.writeBranchRule(x)
			default:
				s = g.writeRule(x)
			}
			ands = fmt.Sprintf("%s%s", ands, s)
		}
		return g.writeAssertlessRule("and", ands, "")
	case *rules.Choices:
		var ands string
		var s string
		for _, x := range r.X {
			s = g.writeRule(x)
			ands = fmt.Sprintf("%s%s", ands, s)
		}
		if r.Op == "or" {
			return g.writeAssert(r.Op, ands)
		}
		return g.writeAssert("", ands)
	default:
		panic(fmt.Sprintf("%T is not a valid rule type", r))
	}
}

func (g *Generator) writeCond(r *rules.Infix) string {
	y := g.unpackCondRule(r.Y)
	x := g.unpackCondRule(r.X)

	return g.writeAssertlessRule(r.Op, x, y)
}

func (g *Generator) unpackCondRule(x rules.Rule) string {
	switch r := x.(type) {
	case *rules.Wrap:
		return r.Value
	case *rules.Infix:
		x := g.unpackCondRule(r.X)
		y := g.unpackCondRule(r.Y)
		return g.writeAssertlessRule(r.Op, x, y)
	default:
		panic(fmt.Sprintf("%T is not a valid rule type", r))
	}
}

func (g *Generator) unpackRule(x rules.Rule) string {
	switch r := x.(type) {
	case *rules.Wrap:
		return r.Value
	case *rules.Infix:
		return g.writeRule(r)
	case *rules.Ands:
		return g.writeRule(r)
	case *rules.Choices:
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

func (g *Generator) generateRules() []string {
	var rules []string
	for _, v := range g.rawRules {
		for _, ru := range v {
			rules = append(rules, g.writeRule(ru))
		}
	}
	return rules
}
