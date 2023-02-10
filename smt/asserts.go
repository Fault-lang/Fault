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

func (g *Generator) parseAssert(a *ast.AssertionStatement) string {
	stateRange := a.Constraint.Operator == "then"
	if stateRange && (a.TemporalFilter != "" || a.Temporal != "") {
		panic("cannot mix temporal logic with when/then assertions")
	}

	left := g.parseInvariantNode(a.Constraint.Left, stateRange)
	right := g.parseInvariantNode(a.Constraint.Right, stateRange)

	op := smtlibOperators(a.Constraint.Operator)

	if stateRange {
		operator := "and"
		sg := g.mergeInvariantInfix(left, right, "or")
		return g.joinStates(sg, operator)
	}

	var on, off string
	if a.Assume {
		on = "or"
		off = "and"
	} else {
		on = "and"
		off = "or"
	}

	//If left and right are asserts on the same variable
	dset := util.DiffStrSets(left.bases, right.bases)
	if dset.Len() == 0 && (a.Temporal != "" || a.TemporalFilter != "") {
		sg := g.mergeInvariantInfix(left, right, smtlibOperators(a.Constraint.Operator))
		ir := g.flattenStates(sg)
		return g.applyTemporalLogic(a.Temporal, ir, a.TemporalFilter, on, off)
	}

	if a.Temporal != "" || a.TemporalFilter != "" {
		ir := expandAssertStateGraph(g.flattenStates(left), g.flattenStates(right), smtlibOperators(a.Constraint.Operator), a.TemporalFilter, a.TemporalN)
		return g.applyTemporalLogic(a.Temporal, ir, a.TemporalFilter, on, off)
	}
	if a.Assume {
		operator := "and"

		if operator == op { // (and (and ) (and )) is redundant
			sg := NewStateGroup()
			sg.wraps = append(left.wraps, right.wraps...)
			return g.joinStates(sg, operator)
		}

		sg := g.mergeInvariantInfix(left, right, op)
		return g.joinStates(sg, operator)
	}

	operator := "or"
	if operator == op {
		sg := NewStateGroup()
		sg.wraps = append(left.wraps, right.wraps...)
		return g.joinStates(sg, operator)
	}
	sg := g.mergeInvariantInfix(left, right, op)
	return g.joinStates(sg, operator)
}

func (g *Generator) parseInvariantNode(exp ast.Expression, stateRange bool) *stateGroup {
	switch e := exp.(type) {
	case *ast.InfixExpression:
		operator := smtlibOperators(e.Operator)
		left := g.parseInvariantNode(e.Left, stateRange)
		right := g.parseInvariantNode(e.Right, stateRange)

		return g.mergeInvariantInfix(left, right, operator)

	case *ast.AssertVar:
		var wg = NewStateGroup()
		for _, v := range e.Instances {
			wg.bases.Add(v)
			st, _, c := captureState(v)
			vr := g.varRounds(v, st)
			wg.AddWrap(&states{base: v,
				terminal: true,
				states:   vr,
				constant: c,
			})
		}
		return wg
	case *ast.IntegerLiteral:
		s := make(map[int][]string)
		s[0] = []string{fmt.Sprint(e.Value)}
		sg := NewStateGroup()
		sg.AddWrap(&states{
			base:     "__int",
			states:   s,
			constant: true})
		return sg
	case *ast.FloatLiteral:
		s := make(map[int][]string)
		s[0] = []string{fmt.Sprint(e.Value)}
		sg := NewStateGroup()
		sg.AddWrap(&states{
			base:     "__float",
			states:   s,
			constant: true,
		})
		return sg
	case *ast.Boolean:
		s := make(map[int][]string)
		s[0] = []string{fmt.Sprint(e.Value)}
		sg := NewStateGroup()
		sg.AddWrap(&states{
			base:     "__bool",
			states:   s,
			constant: true,
		})
		return sg
	case *ast.StringLiteral:
		s := make(map[int][]string)
		s[0] = []string{fmt.Sprint(e.Value)}
		sg := NewStateGroup()
		sg.AddWrap(&states{
			base:     "__string",
			states:   s,
			constant: true,
		})
		return sg
	case *ast.PrefixExpression:
		var operator string
		right := g.parseInvariant(e.Right)
		if e.Operator == "!" { //Not valid in SMTLib
			operator = "not"
		} else {
			operator = smtlibOperators(e.Operator)
		}

		if r, ok := right.(*states); ok {
			return g.mergeInvariantPrefix([]*states{r}, operator)
		} else {
			return g.mergeInvariantPrefix(right.(*stateGroup).wraps, operator)
		}

	case *ast.Nil:
	case *ast.IndexExpression:
		var wg = NewStateGroup()
		for _, v := range e.Left.(*ast.AssertVar).Instances {
			wg.bases.Add(v)
			vr := g.varRounds(v, e.Index.String())
			wg.AddWrap(&states{base: v,
				states:   vr,
				constant: true,
			})
		}
		return wg
	default:
		pos := e.Position()
		panic(fmt.Sprintf("illegal node %T in assert or assume line: %d, col: %d", e, pos[0], pos[1]))
	}
	return nil
}

func (g *Generator) mergeInvariantPrefix(right []*states, operator string) *stateGroup {
	sg := NewStateGroup()
	for _, r := range right {
		states := make(map[int][]string)
		for i := 0; i <= g.Rounds; i++ {
			if s, ok := r.states[i]; ok {
				states[i] = append(states[i], fmt.Sprintf("(%s %s)", operator, s))
			}
		}
		r.states = states
		sg.AddWrap(r)
	}
	return sg
}

func (g *Generator) mergeInvariantInfix(left *stateGroup, right *stateGroup, operator string) *stateGroup {
	sg := NewStateGroup()
	for _, l := range left.wraps {
		for _, r := range right.wraps {
			state := g.mergeByRound(l, r, operator)
			sg.AddWrap(state)
		}
	}

	return sg

}

func (g *Generator) mergePairs(l string, r string, leftLead bool) []string {
	if leftLead {
		return []string{l, r}
	}
	return []string{r, l}
}

func (g *Generator) mergeByRound(left *states, right *states, operator string) *states {
	ret := &states{}

	st := make(map[int][]string)
	if left.constant && right.constant {
		combos := util.PairCombinations(left.states[0], right.states[0])
		st[0] = packageStateGraph(combos, operator)
		ret.states = st
		return ret
	}

	if left.base == right.base &&
		(left.base != "" && right.base != "") {
		ret.base = left.base

		var long map[int][]string
		var short map[int][]string
		var leftLead bool

		if len(left.states) >= len(right.states) {
			long = left.states
			short = right.states
			leftLead = true
		} else {
			long = right.states
			short = left.states
			leftLead = false
		}

		//Pair based on same state
		for i := 0; i <= g.Rounds; i++ {
			var pairs [][]string
			var slast []string
			for idx, s := range long[i] {
				if sstates, ok := short[i]; ok {
					slast = sstates
					if len(sstates) > idx {
						p := g.mergePairs(s, sstates[idx], leftLead)
						pairs = append(pairs, p)
						continue
					}

					p := g.mergePairs(s, sstates[len(sstates)-1], leftLead)
					pairs = append(pairs, p)
					continue
				}
				p := g.mergePairs(s, slast[len(slast)-1], leftLead)
				pairs = append(pairs, p)

			}
			st[i] = packageStateGraph(pairs, operator)
		}
		ret.states = st
		return ret
	}

	if left.constant {
		ret.base = right.base
		st := g.balance(right, left, operator)
		ret.states = st
		return ret
	}

	if right.constant {
		ret.base = left.base
		st := g.balance(left, right, operator)
		ret.states = st
		return ret
	}

	if left.terminal && right.terminal {
		combos := g.termCombos(left.base, right.base)
		for i, c := range combos {
			st[i] = packageStateGraph(c, operator)
		}
		ret.states = st
		return ret
	}

	var llast, rlast []string
	for i := 0; i <= g.Rounds; i++ {
		var l, r []string
		var okleft, okright bool
		if l, okleft = left.states[i]; !okleft {
			if llast == nil {
				if invalidBase(left.base) {
					panic("assert left variable base name is invalid")
				}
				l = []string{fmt.Sprintf("%s_%s", left.base, "0")}
			} else {
				l = llast
			}
		}

		if r, okright = right.states[i]; !okright {
			if rlast == nil {
				if invalidBase(right.base) {
					panic("assert left variable base name is invalid")
				}
				r = []string{fmt.Sprintf("%s_%s", right.base, "0")}
			} else {
				r = rlast
			}
		}

		combos := util.PairCombinations(l, r)
		st[i] = packageStateGraph(combos, operator)
		llast = l[len(l)-1:]
		rlast = r[len(r)-1:]
	}
	ret.states = st
	return ret
}

func (g *Generator) termCombos(lbase string, rbase string) map[int][][]string {
	var llast string
	var rlast string
	combos := make(map[int][][]string)

	for i, rounds := range g.RoundVars {
		var c [][]string
		for _, vr := range rounds {
			if vr[0] == lbase {
				llast = strings.Join(vr, "_")
			}

			if vr[0] == lbase && rlast != "" {
				c = append(c, []string{llast, rlast})
			}

			if vr[0] == rbase {
				rlast = strings.Join(vr, "_")
			}

			if vr[0] == rbase && llast != "" {
				c = append(c, []string{llast, rlast})
			}
		}
		combos[i] = c
	}
	return combos
}

func (g *Generator) balance(vr *states, con *states, operator string) map[int][]string {
	ret := make(map[int][]string)
	for i := 0; i <= g.Rounds; i++ {
		if v, ok := vr.states[i]; ok {
			combos := util.PairCombinations(v, con.states[0])
			ret[i] = packageStateGraph(combos, operator)
		}
	}
	return ret
}

func (g *Generator) flattenStates(sg *stateGroup) []string {
	var asserts []string
	for _, w := range sg.wraps {
		for i := 0; i <= g.Rounds; i++ {
			if s, ok := w.states[i]; ok {
				asserts = append(asserts, s...)
			}
		}
	}
	return asserts
}

func (g *Generator) joinStates(sg *stateGroup, operator string) string {
	asserts := g.flattenStates(sg)
	if len(asserts) == 1 {
		return asserts[0]
	}
	return g.writeAssertlessRule(operator, strings.Join(asserts, " "), "")
}

func invalidBase(base string) bool {
	switch base {
	case "__string":
		return true
	case "__bool":
		return true
	case "__float":
		return true
	case "__int":
		return true
	case "":
		return true
	}
	return false
}

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
			if op == "not" && a[0] == "false" {
				s = fmt.Sprintf("(%s %s)", op, a[1])

			} else if op == "not" && a[1] == "false" {
				s = fmt.Sprintf("(%s %s)", op, a[0])

			} else if op == "not" {
				s = fmt.Sprintf("(%s (= %s %s))", op, a[0], a[1])

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
		return "not"
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
