package smt

import (
	"fault/ast"
	"fault/llvm"
	"fault/util"
	"fmt"
	"strconv"
	"strings"
)

func (g *Generator) parseAssert(assert ast.Node) ([]*assrt, []*assrt, string) {
	switch e := assert.(type) {
	case *ast.AssertionStatement:
		a1 := g.generateAsserts(e.Constraints.Left, e.Constraints.Operator, e.Constraints, e)
		a2 := g.generateAsserts(e.Constraints.Right, e.Constraints.Operator, e.Constraints, e)

		if e.Constraints.Operator != "&&" && e.Constraints.Operator != "||" {
			return a1, a2, e.Constraints.Operator
		} else {
			a2 = removeDuplicates(a1, a2)
			return append(a1, a2...), nil, ""
		}
	case *ast.AssumptionStatement:
		a1 := g.generateAsserts(e.Constraints.Left, e.Constraints.Operator, e.Constraints, e)
		a2 := g.generateAsserts(e.Constraints.Right, e.Constraints.Operator, e.Constraints, e)
		if e.Constraints.Operator != "&&" && e.Constraints.Operator != "||" {
			return a1, a2, e.Constraints.Operator
		} else {
			return append(a1, a2...), nil, ""
		}
	default:
		pos := e.Position()
		panic(fmt.Sprintf("not a valid assert or assumption line: %d, col: %d", pos[0], pos[1]))
	}
}

func (g *Generator) generateAsserts(exp ast.Expression, comp string, constr ast.Expression, stmt ast.Statement) []*assrt {
	var ident []string
	var assrt []*assrt
	switch v := exp.(type) {
	case *ast.InfixExpression:
		ident = g.findIdent(v)
		for _, id := range ident {
			assrt = append(assrt, g.packageAssert(id, comp, v, stmt))
		}
		return assrt
	case *ast.Identifier:
		ident = g.findIdent(v)
		for _, id := range ident {
			assrt = append(assrt, g.packageAssert(id, comp, constr, stmt))
		}
		return assrt
	case *ast.ParameterCall:
		ident = g.findIdent(v)
		for _, id := range ident {
			assrt = append(assrt, g.packageAssert(id, comp, constr, stmt))
		}
		return assrt
	case *ast.AssertVar:
		for _, v := range v.Instances {
			assrt = append(assrt, g.packageAssert(v, comp, constr, stmt))
		}
		return assrt
	case *ast.IndexExpression:
		ident = g.findIdent(v)
		for _, id := range ident {
			assrt = append(assrt, g.packageAssert(id, comp, constr, stmt))
		}
		return assrt
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
}

func (g *Generator) parseInvariant(ex ast.Expression) rule {
	switch e := ex.(type) {
	case *ast.InvariantClause:
		left := g.parseInvariant(e.Left)
		right := g.parseInvariant(e.Right)

		i := &invariant{
			left:        left,
			conjunction: smtlibOperators(e.Operator),
			right:       right,
		}
		if e.Operator == "!=" { //Not valid in SMTLib
			return &invariant{conjunction: "not",
				right: i}
		}
		return i

	case *ast.InfixExpression:
		left := g.parseInvariant(e.Left)
		right := g.parseInvariant(e.Right)
		i := &invariant{
			left:        left,
			conjunction: smtlibOperators(e.Operator),
			right:       right,
		}
		if e.Operator == "!=" { //Not valid in SMTLib
			return &invariant{conjunction: "not",
				right: i}
		}
		return i

	case *ast.AssertVar:
		if len(e.Instances) == 1 {
			s, a, c := captureState(e.Instances[0])
			return &wrap{value: e.Instances[0],
				state:    s,
				all:      a,
				constant: c,
			}
		}
		var wg = &wrapGroup{}
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
		return &wrap{value: g.convertIndexExpr(e),
			state:    "",
			all:      false,
			constant: true,
		}
	default:
		pos := e.Position()
		panic(fmt.Sprintf("illegal node %T in assert or assume line: %d, col: %d", e, pos[0], pos[1]))
	}
	return nil
}

func (g *Generator) packageAssert(ident string, comp string, expr ast.Expression, stmt ast.Statement) *assrt {
	var temporalFilter string
	var temporalN int

	switch st := stmt.(type) {
	case *ast.AssertionStatement:
		temporalFilter = st.TemporalFilter
		temporalN = st.TemporalN
	case *ast.AssumptionStatement:
		temporalFilter = st.TemporalFilter
		temporalN = st.TemporalN
	}

	s, a, c := captureState(ident)
	w := &wrap{value: ident,
		state:    s,
		all:      a,
		constant: c,
	}
	return &assrt{
		variable:       w,
		conjunction:    comp,
		assertion:      g.parseInvariant(expr),
		temporalFilter: temporalFilter,
		temporalN:      temporalN}
}

func (g *Generator) convertIndexExpr(idx *ast.IndexExpression) string {
	return strings.Join([]string{idx.Left.String(), idx.Index.String()}, "_")
}

func (g *Generator) findIdent(n ast.Node) []string {
	switch v := n.(type) {
	case *ast.InfixExpression:
		return g.findIdent(v.Left)
	case *ast.PrefixExpression:
		return g.findIdent(v.Right)
	case *ast.IndexExpression:
		s := g.convertIndexExpr(v)
		return []string{s}
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

func (g *Generator) generateAssertRules(ru rule, t string, tn int) []string {
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
		return g.generateAssertRules(r.assertion, t, tn)
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
		left = g.generateAssertRules(l, t, tn)
	case *wrap:
		left = g.wrapPerm(l)
	default:
		left = nil
	}

	switch r := i.right.(type) {
	case *invariant:
		right = g.generateAssertRules(r, t, tn)
	case *wrap:
		right = g.wrapPerm(r)
	default:
		right = nil
	}

	if left == nil { // Typically (not (some rule))
		var ret []string
		for _, r := range right {
			ret = append(ret, fmt.Sprintf("(%s %s)", i.conjunction, r))
		}
		return ret
	}

	return expandAssertStateGraph(left, right, i.conjunction, t, tn)
}

func (g *Generator) generateCompound(a1 []*assrt, a2 []*assrt, op string) []string {
	var left, right []string
	for _, l := range a1 {
		left = append(left, g.generateAssertRules(l, l.temporalFilter, l.temporalN)...)
	}

	for _, r := range a2 {
		right = append(right, g.generateAssertRules(r, r.temporalFilter, r.temporalN)...)
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

func expandAssertStateGraph(list1 []string, list2 []string, op string, temporalFilter string, temporalN int) []string {
	var x [][]string
	c := util.Cartesian(list1, list2)
	switch temporalFilter {
	// For logic like "no more than X times" "no fewer than X times"
	// We need to flip some of the operators and build out more
	// states before packaging the asserts
	case "nmt":
		// (and (or on on on) (and off off))
		combos := util.Combinations(c, temporalN) //generate all combinations of possible on states
		pairs := impliesOnOffPairs(combos, c)     //for each combination prepare a list of states that must logically be off
		for _, p := range pairs {
			var o []string
			var f []string
			for _, on := range p[0] {
				// Write the clauses
				o = append(o, fmt.Sprintf("(%s %s %s)", op, on[0], on[1]))
			}
			// For nmt any of the potential on states can be on
			var onStr string
			if len(o) == 1 {
				onStr = o[0]
			} else {
				onStr = fmt.Sprintf("(%s %s)", "or", strings.Join(o, " "))
			}

			offOp := llvm.OP_NEGATE[op]
			for _, off := range p[1] {
				if op == "=" {
					f = append(f, fmt.Sprintf("(%s (%s %s %s))", "not", op, off[0], off[1]))
				} else {
					f = append(f, fmt.Sprintf("(%s %s %s)", offOp, off[0], off[1]))
				}
			}
			// But these states must be off
			var offStr string
			if len(f) == 1 {
				offStr = f[0]
			} else {
				offStr = fmt.Sprintf("(%s %s)", "and", strings.Join(f, " "))
			}
			x = append(x, []string{onStr, offStr})
		}
		return packageStateGraph(x, "and")
	case "nft":
		// (or (and on on on))
		combos := util.Combinations(c, temporalN)
		pairs := impliesOnOffPairs(combos, c)
		for _, p := range pairs {
			var o []string
			//var f []string
			for _, on := range p[0] {
				o = append(o, fmt.Sprintf("(%s %s %s)", op, on[0], on[1]))
			}
			// For nft all on states in this possibility MUST be on
			var onStr string
			if len(o) == 1 {
				onStr = o[0]
			} else {
				onStr = fmt.Sprintf("(%s %s)", "and", strings.Join(o, " "))
			}

			// offOp := llvm.OP_NEGATE[op]
			// for _, off := range p[1] {
			// 	if op == "=" {
			// 		f = append(f, fmt.Sprintf("(%s (%s %s %s))", "not", op, off[0], off[1]))
			// 	} else {
			// 		f = append(f, fmt.Sprintf("(%s %s %s)", offOp, off[0], off[1]))
			// 	}
			// }
			// // The off states can be on or off
			// offStr := fmt.Sprintf("(%s %s)", "or", strings.Join(f, " "))

			x = append(x, []string{onStr})
		}
		return packageStateGraph(x, "or")
	default:
		return packageStateGraph(c, op)
	}
}

func packageStateGraph(x [][]string, op string) []string {
	var product []string
	for _, a := range x {
		if len(a) == 1 {
			product = append(product, a[0])
		} else {
			s := fmt.Sprintf("(%s %s %s)", op, a[0], a[1])
			product = append(product, s)
		}
	}
	return product
}

func smtlibOperators(op string) string {
	switch op {
	case "==":
		return "="
	case "!=": //Invalid in SMTLib
		return "="
	default:
		return op
	}
}

func impliesOnOffPairs(on [][][]string, c [][]string) [][][][]string {
	var oop [][][][]string
	for _, o := range on {
		off := util.NotInSet(o, c)
		p := [][][]string{o, off}
		oop = append(oop, p)
	}
	return oop
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

func removeDuplicates(a1 []*assrt, a2 []*assrt) []*assrt {
	for _, x := range a1 {
		found, index := valueInList(x, a2)
		if found {
			a2 = removeFromList(index, a2)
		}
	}
	return a2
}

func valueInList(x *assrt, l []*assrt) (bool, int) {
	for k, y := range l {
		if x.assertion.String() == y.assertion.String() {
			return true, k
		}
	}
	return false, 0
}

func removeFromList(k int, l []*assrt) []*assrt {
	if (k + 1) != len(l) {
		return append(l[0:k], l[k+1:]...)
	}
	return l[0:k]
}
