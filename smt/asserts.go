package smt

import (
	"fault/ast"
	"fault/util"
	"fmt"
	"strconv"
	"strings"
)

func (g *Generator) parseAssert(assert ast.Node) ([]*assrt, []*assrt, string) {
	switch e := assert.(type) {
	case *ast.AssertionStatement:
		a1 := g.generateAsserts(e.Constraints.Variable, e.Constraints.Comparison, e.Constraints)
		a2 := g.generateAsserts(e.Constraints.Expression, e.Constraints.Comparison, e.Constraints)

		if e.Constraints.Conjuction != "" {
			return a1, a2, e.Constraints.Conjuction
		} else {
			a1, a2 := removeDuplicates(a1, a2)
			return append(a1, a2...), nil, ""
		}
	case *ast.AssumptionStatement:
		a1 := g.generateAsserts(e.Constraints.Variable, e.Constraints.Comparison, e.Constraints)
		a2 := g.generateAsserts(e.Constraints.Expression, e.Constraints.Comparison, e.Constraints)
		if e.Constraints.Conjuction != "" {
			return a1, a2, e.Constraints.Conjuction
		} else {
			return append(a1, a2...), nil, ""
		}
	default:
		pos := e.Position()
		panic(fmt.Sprintf("not a valid assert or assumption line: %d, col: %d", pos[0], pos[1]))
	}
}

func (g *Generator) generateAsserts(exp ast.Expression, comp string, constr ast.Expression) []*assrt {
	var ident []string
	var assrt []*assrt
	switch v := exp.(type) {
	case *ast.InfixExpression:
		ident = g.findIdent(v)
		for _, id := range ident {
			assrt = append(assrt, g.packageAssert(id, v.Operator, v))
		}
	case *ast.Identifier:
		ident = g.findIdent(v)
		for _, id := range ident {
			assrt = append(assrt, g.packageAssert(id, comp, constr))
		}
	case *ast.ParameterCall:
		ident = g.findIdent(v)
		for _, id := range ident {
			assrt = append(assrt, g.packageAssert(id, comp, constr))
		}
	case *ast.AssertVar:
		for _, v := range v.Instances {
			assrt = append(assrt, g.packageAssert(v, comp, constr))
		}
	case *ast.IntegerLiteral:
		return assrt
	case *ast.FloatLiteral:
		return assrt
	case *ast.StringLiteral:
		return assrt
	case *ast.Boolean:
		return assrt
	default:
		panic(fmt.Sprintf("poorly formatted assert statement got=%T", exp))
	}
	return assrt
}

func (g *Generator) parseInvariant(ex ast.Expression) rule {

	switch e := ex.(type) {
	case *ast.Invariant:
		left := g.parseInvariant(e.Variable)
		right := g.parseInvariant(e.Expression)
		var conj string
		if e.Conjuction != "" {
			conj = e.Conjuction
		} else {
			conj = e.Comparison
		}
		return &invariant{
			left:        left,
			conjunction: conj,
			right:       right,
		}
	case *ast.InfixExpression:
		left := g.parseInvariant(e.Left)
		right := g.parseInvariant(e.Right)
		return &invariant{
			left:        left,
			conjunction: e.Operator,
			right:       right,
		}

	case *ast.AssertVar:
		if len(e.Instances) == 1 {
			s, a, c := captureState(e.Instances[0])
			return &wrap{value: e.Instances[0],
				state:    s,
				all:      a,
				constant: c,
			}
		}
		var wg *wrapGroup
		for _, v := range e.Instances {
			s, a, c := captureState(v)
			wg.wraps = append(wg.wraps, &wrap{value: v,
				state:    s,
				all:      a,
				constant: c,
			})
		}
		return wg
	// case *ast.Identifier:
	// 	s, a, c := captureState(e.Value)
	// 	return &wrap{value: e.Value,
	// 		state:    s,
	// 		all:      a,
	// 		constant: c,
	// 	}
	case *ast.IntegerLiteral:
		return &wrap{value: fmt.Sprint(e.Value),
			state:    "",
			all:      false,
			constant: true,
		}
	case *ast.FloatLiteral:
		return &wrap{value: fmt.Sprint(e.Value),
			state:    "",
			all:      false,
			constant: true,
		}
	case *ast.Boolean:
		return &wrap{value: fmt.Sprint(e.Value),
			state:    "",
			all:      false,
			constant: true,
		}
	case *ast.StringLiteral:
		return &wrap{value: e.Value,
			state:    "",
			all:      false,
			constant: true,
		}
	//case *ast.Natural:
	//case *ast.Uncertain:
	case *ast.PrefixExpression:
	case *ast.Nil:
	case *ast.IndexExpression:
	default:
		pos := e.Position()
		panic(fmt.Sprintf("illegal node %T in assert or assume line: %d, col: %d", e, pos[0], pos[1]))
	}
	return nil
}

func (g *Generator) packageAssert(ident string, comp string, expr ast.Expression) *assrt {
	s, a, c := captureState(ident)
	w := &wrap{value: ident,
		state:    s,
		all:      a,
		constant: c,
	}
	return &assrt{
		variable:    w,
		conjunction: comp,
		assertion:   g.parseInvariant(expr)}
}

func (g *Generator) findIdent(n ast.Node) []string {
	switch v := n.(type) {
	case *ast.InfixExpression:
		return g.findIdent(v.Left)
	case *ast.PrefixExpression:
		return g.findIdent(v.Right)
	case *ast.Identifier:
		return []string{v.Value}
	case *ast.ParameterCall:
		return []string{strings.Join(v.Value, "_")}
	case *ast.AssertVar:
		return v.Instances
	default:
		pos := n.Position()
		panic(fmt.Sprintf("improperly formatted assert or assume line: %d, col: %d", pos[0], pos[1]))
	}
}

func (g *Generator) generateAssertRules(ru rule) []string {
	// assert x == true;
	// negated: x != true;
	// (assert (or (= x0 false) (= x1 false) (= x2 false)))

	// assert x > 5 && x < 2;
	// negated: x <= 5 || x >= 2
	// (assert (or (<= x0 5) (<= x1 5) (<= x2 5) (>= x0 2) (>= x1 2) (>= x2 2)))

	// assert x > 5 || x < 2;
	// negated: x <= 5 && x >= 2
	// (assert (and (or (<= x0 5) (<= x1 5) (<= x2 5))
	//				(or (>= x0 2) (>= x1 2) (>= x2 2))))

	//assert x > y
	// negated: x <= y
	// What states overlap in the same round? :/
	// Actually does this matter?
	// (assert (or (<= x0 y0) (<= x1 y1) (<= x2 y2)))
	var i *invariant
	switch r := ru.(type) {
	case *assrt:
		return g.generateAssertRules(r.assertion)
	case *invariant:
		i = r
	case *wrap:
		return g.wrapPerm(r)
	case *wrapGroup:
		var wg []string
		for _, v := range r.wraps {
			wg = append(wg, g.wrapPerm(v)...)
		}
		return wg
	default:
		panic(fmt.Sprintf("Improperly formatted assertion %s", r.String()))
	}

	var left, right []string

	switch l := i.left.(type) {
	case *invariant:
		left = g.generateAssertRules(l)
	case *wrap:
		left = g.wrapPerm(l)
	}

	switch r := i.right.(type) {
	case *invariant:
		right = g.generateAssertRules(r)
	case *wrap:
		right = g.wrapPerm(r)
	}

	return cartesianAsserts(left, right, i.conjunction)
}

func (g *Generator) generateCompound(a1 []*assrt, a2 []*assrt, op string) []string {
	var left, right []string
	for _, l := range a1 {
		left = append(left, g.generateAssertRules(l)...)
	}

	for _, r := range a2 {
		right = append(right, g.generateAssertRules(r)...)
	}

	switch op {
	case "&&":
		lor := fmt.Sprintf("(or %s)", strings.Join(left, " "))
		ror := fmt.Sprintf("(or %s)", strings.Join(right, " "))
		return []string{
			g.writeAssert("", fmt.Sprintf("(and %s %s)", lor, ror))}
	case "||":
		combo := append(left, right...)
		return []string{g.writeAssert("", fmt.Sprintf("(or %s)", strings.Join(combo, " ")))}
	default:
		panic(fmt.Sprintf("improper conjunction for assert got=%s", op))
	}
}

func (g *Generator) wrapPerm(w *wrap) []string {
	if w.constant {
		return []string{w.value}
	} else if w.state != "" {
		state := fmt.Sprint(w.value, "_", w.state)
		return []string{state}
	}
	if w.all {
		var states []string
		end := g.ssa[w.value]
		for i := 0; i < int(end+1); i++ {
			states = append(states, fmt.Sprint(w.value, "_", i))
		}
		return states
	}
	panic(fmt.Sprintf("Inproperly formatted metadata for value %s in assert", w.value))
}

func cartesianAsserts(list1 []string, list2 []string, op string) []string {
	var product []string
	for _, a := range util.Cartesian(list1, list2) {
		s := fmt.Sprintf("(%s %s %s)", op, a[0], a[1])
		product = append(product, s)
	}
	return product
}

func captureState(id string) (string, bool, bool) {
	var a, c bool
	raw := strings.Split(id, "_")
	if len(raw) > 2 { //Not a constant
		c = false
		a = true
	} else {
		c = true
		a = false
	}

	_, err := strconv.Atoi(raw[len(raw)-1])
	if err != nil {
		return "", a, c
	} else {
		return raw[len(raw)-1], false, false
	}

}

func removeDuplicates(a1 []*assrt, a2 []*assrt) ([]*assrt, []*assrt) {
	for _, x := range a1 {
		for k, y := range a2 {
			if x.assertion.String() == y.assertion.String() {
				a2 = append(a2[0:k], a2[k+1:]...)
			}
		}
	}
	return a1, a2
}
