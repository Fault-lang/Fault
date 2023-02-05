package smt

import (
	"fault/ast"
	"fault/llvm"
	"fault/util"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// func (g *Generator) parseAssert(assert ast.Node) ([]*assrt, []*assrt, string) {
// 	switch e := assert.(type) {
// 	case *ast.AssertionStatement:
// 		a1 := g.generateAsserts(e.Constraints.Left, e.Constraints.Operator, e.Constraints, e)
// 		a2 := g.generateAsserts(e.Constraints.Right, e.Constraints.Operator, e.Constraints, e)

// 		if e.Constraints.Operator == "&&" || e.Constraints.Operator == "||" {
// 			return a1, a2, e.Constraints.Operator
// 		} else {
// 			a2 = removeDuplicates(a1, a2)
// 			return append(a1, a2...), nil, ""
// 		}
// 	case *ast.AssumptionStatement:
// 		a1 := g.generateAsserts(e.Constraints.Left, e.Constraints.Operator, e.Constraints, e)
// 		a2 := g.generateAsserts(e.Constraints.Right, e.Constraints.Operator, e.Constraints, e)
// 		if e.Constraints.Operator == "&&" || e.Constraints.Operator == "||" {
// 			return a1, a2, e.Constraints.Operator
// 		} else {
// 			return append(a1, a2...), nil, ""
// 		}
// 	default:
// 		pos := e.Position()
// 		panic(fmt.Sprintf("not a valid assert or assumption line: %d, col: %d", pos[0], pos[1]))
// 	}
// }

func (g *Generator) generateAsserts(exp ast.Expression, comp string, constr ast.Expression, stmt *ast.AssertionStatement) []*assrt {
	var ident []string
	var assrt []*assrt
	switch v := exp.(type) {
	case *ast.InfixExpression:
		ident = g.findIdent(v)
		for _, id := range ident {
			assrt = append(assrt, g.packageAssert(id, comp, v, stmt))
		}
		return assrt
	case *ast.PrefixExpression:
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
		for _, inst := range v.Instances {
			assrt = append(assrt, g.packageAssert(inst, comp, constr, stmt))
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

		if e.Operator == "then" {
			return &invariant{
				left:     left,
				operator: "then",
				right:    right,
			}
		}

		i := &invariant{
			left:     left,
			operator: smtlibOperators(e.Operator),
			right:    right,
		}
		if e.Operator == "!=" { //Not valid in SMTLib
			return &invariant{operator: "not",
				right: i}
		}
		return i

	case *ast.InfixExpression:
		left := g.parseInvariant(e.Left)
		right := g.parseInvariant(e.Right)
		i := &invariant{
			left:     left,
			operator: smtlibOperators(e.Operator),
			right:    right,
		}
		if e.Operator == "!=" { //Not valid in SMTLib
			return &invariant{operator: "not",
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
	case *ast.PrefixExpression:
		right := g.parseInvariant(e.Right)
		i := &invariant{
			left:     nil,
			operator: smtlibOperators(e.Operator),
			right:    right,
		}
		if e.Operator == "!" { //Not valid in SMTLib
			return &invariant{operator: "not",
				right: i}
		}
		return i
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

type thenStates struct {
	roundClauses []string
	values       [][][]string
}

func (g *Generator) generateThenRules(inv *invariant) []string {
	when := g.whenInfixNode(inv.left)
	//then := g.thenInfixNode(inv.right)
	fmt.Println(when)
	//fmt.Println(then)
	return []string{}
}

// check rounds for any matching variable names

// Generate all permutations of variables in the assert,
// in the round, in between variable state change

func (g *Generator) whenInfixNode(ru rule) map[string]*thenStates {
	ret := make(map[string]*thenStates)
	switch r := ru.(type) {
	case *invariant:
		left := g.whenInfixNode(r.left)
		for k, v := range left {
			ret[k] = v
		}

		right := g.whenInfixNode(r.right)
		for k, v := range right {
			ret[k] = v
		}

		return ret
	case *wrap:
		roundClause, values := g.whenNode(r)
		ret[r.value] = &thenStates{
			roundClauses: roundClause,
			values:       values,
		}
		return ret
	}
	return ret
}

func (g *Generator) whenNode(when *wrap) ([]string, [][][]string) {
	var roundClauses []string
	var values [][][]string
	base := when.value
	rounds := g.lookupVarRounds(when.value, when.state)
	for _, r := range rounds {
		roundClauses = append(roundClauses, fmt.Sprintf("(= %s_%d %s)", base, r[0], "true"))
		values = append(values, g.RoundVars[r[1]][r[2]:])
	}
	return roundClauses, values
}

// func (g *Generator) thenNode(ru rule) {

// }

// func (g *Generator) generateThenRules(inv *invariant) []string {
// 	var rounds [][]int
// 	var roundClauses []string
// 	var values [][][]string
// 	var base string
// 	switch when := inv.left.(type) {
// 	case *invariant:
// 		if when.left == nil {
// 			base = when.right.(*wrap).value
// 			rounds = g.lookupVarRounds(base, when.right.(*wrap).state)
// 			for _, r := range rounds {
// 				roundClauses = append(roundClauses, fmt.Sprintf("(= %s_%d %s)", base, r[0], "false"))
// 				values = append(values, g.RoundVars[r[1]][r[2]:])
// 			}
// 		}
// 	case *wrap:
// 		base = when.value
// 		rounds = g.lookupVarRounds(when.value, when.state)
// 		for _, r := range rounds {
// 			roundClauses = append(roundClauses, fmt.Sprintf("(= %s_%d %s)", base, r[0], "true"))
// 			values = append(values, g.RoundVars[r[1]][r[2]:])
// 		}
// 	}

// 	var rules []string
// 	switch then := inv.right.(type) {
// 	case *invariant:
// 		if then.left == nil { //Prefix
// 			rules = g.constructThen(then.left.(*wrap), values, roundClauses, "false")
// 		}
// 	case *wrap:
// 		rules = g.constructThen(then, values, roundClauses, "true")
// 	}
// 	return rules
// }

// func (g *Generator) whenClauses(when *wrap) ([][]int, []string, [][][]string) {
// 	var roundClauses []string
// 	var values [][][]string
// 	base := when.value
// 	rounds := g.lookupVarRounds(when.value, when.state)
// 	for _, r := range rounds {
// 		roundClauses = append(roundClauses, fmt.Sprintf("(= %s_%d %s)", base, r[0], "true"))
// 		values = append(values, g.RoundVars[r[1]][r[2]:])
// 	}
// 	return rounds, roundClauses, values
// }

// func (g *Generator) thenInfix(ru rule) []string {
// 	switch when := ru.(type) {
// 	case *invariant:
// 		leftRounds, leftRC, leftValues := g.thenInfixNode(when.left)
// 		rightRounds, rightRC, rightValues := g.thenInfixNode(when.right)

// 	case *wrap:
// 		rounds, roundClauses, values := g.whenClauses(when)
// 	default:
// 		panic("unsupported rule")
// 	}
// }

// func (G *Generator) thenClauses(then *wrap, values [][][]string) []string {
// 	var or []string
// 	var rules []string
// 	for _, val := range values {
// 		for _, v := range val {
// 			if v[0] == then.value {
// 				vname := strings.Join(v, "_")
// 				or = append(or, fmt.Sprintf("(= %s %s)", vname, b))
// 			}
// 		}
// 		tclause := fmt.Sprintf("(or %s)", strings.Join(or, " "))

// 		rules = append(rules, tclause)

// 	}
// 	return rules
// }

// func (g *Generator) constructThen(then *wrap, values [][][]string, roundClauses []string, b string) []string {
// 	var or []string
// 	var rules []string
// 	for idx, val := range values {
// 		for _, v := range val {
// 			if v[0] == then.value {
// 				vname := strings.Join(v, "_")
// 				or = append(or, fmt.Sprintf("(= %s %s)", vname, b))
// 			}
// 		}
// 		wclause := roundClauses[idx]
// 		tclause := fmt.Sprintf("(or %s)", strings.Join(or, " "))

// 		rules = append(rules, fmt.Sprintf("(and %s %s)", wclause, tclause))

// 	}
// 	return rules
// }

func (g *Generator) packageAssert(ident string, comp string, expr ast.Expression, stmt *ast.AssertionStatement) *assrt {
	var temporalFilter string
	var temporalN int

	temporalFilter = stmt.TemporalFilter
	temporalN = stmt.TemporalN

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
		if r.operator == "then" {
			return g.generateThenRules(r)
		}
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
			ret = append(ret, fmt.Sprintf("(%s %s)", i.operator, r))
		}
		return ret
	}

	return expandAssertStateGraph(left, right, i.operator, t, tn)
}

func (g *Generator) generateCompound(a1 []*assrt, a2 []*assrt, op string) string {
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
		return fmt.Sprintf("(and %s %s)", lor, ror)
	case "||":
		combo := append(left, right...)
		return fmt.Sprintf("(or %s)", strings.Join(combo, " "))
	default:
		panic(fmt.Sprintf("improper conjunction for assert got=%s", op))
	}
}

func (g *Generator) filterOutTempStates(v string, i int16) bool {
	for _, opt := range g.forks {
		choices := opt[v]
		for _, c := range choices {
			if len(c.Values) == 1 {
				return false
			}

			c.Values = c.Values[1:] //First value is not temp
			n := len(c.Values)
			t := sort.Search(n, func(k int) bool { return c.Values[k] == i })
			if t == n { //If no match in Values slice Search returns n since it cannot be confused with an index
				return false
			}

			if c.Values[t] == i {
				return true
			}
		}
	}
	return false
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
		end := g.variables.ssa[w.value]
		for i := 0; i < int(end+1); i++ {
			if !g.filterOutTempStates(w.value, int16(i)) {
				states = append(states, fmt.Sprint(w.value, "_", i))
			}
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
			var s string
			if op == "=" && a[0] == "false" {
				s = fmt.Sprintf("(%s %s)", "not", a[1])
			} else if op == "=" && a[1] == "false" {
				s = fmt.Sprintf("(%s %s)", "not", a[0])
			} else {
				s = fmt.Sprintf("(%s %s %s)", op, a[0], a[1])
			}
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
	case "||":
		return "or"
	case "&&":
		return "and"
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
