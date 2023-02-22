package smt

import (
	"fault/ast"
	"fault/llvm"
	"fault/smt/rules"
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
	dset := util.DiffStrSets(left.Bases, right.Bases)
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
			sg := rules.NewStateGroup()
			sg.Wraps = append(left.Wraps, right.Wraps...)
			return g.joinStates(sg, operator)
		}

		sg := g.mergeInvariantInfix(left, right, op)
		return g.joinStates(sg, operator)
	}

	operator := "or"
	if operator == op {
		sg := rules.NewStateGroup()
		sg.Wraps = append(left.Wraps, right.Wraps...)
		return g.joinStates(sg, operator)
	}
	sg := g.mergeInvariantInfix(left, right, op)
	return g.joinStates(sg, operator)
}

func (g *Generator) parseInvariantNode(exp ast.Expression, stateRange bool) *rules.StateGroup {
	switch e := exp.(type) {
	case *ast.InfixExpression:
		operator := smtlibOperators(e.Operator)
		left := g.parseInvariantNode(e.Left, stateRange)
		right := g.parseInvariantNode(e.Right, stateRange)

		return g.mergeInvariantInfix(left, right, operator)

	case *ast.AssertVar:
		var wg = rules.NewStateGroup()
		for _, v := range e.Instances {
			wg.Bases.Add(v)
			st, _, c := captureState(v)
			vr := g.varRounds(v, st)
			wg.AddWrap(&rules.States{Base: v,
				Terminal: true,
				States:   vr,
				Constant: c,
			})
		}
		return wg
	case *ast.IntegerLiteral:
		s := make(map[int][]string)
		s[0] = []string{fmt.Sprint(e.Value)}
		sg := rules.NewStateGroup()
		sg.AddWrap(&rules.States{
			Base:     "__int",
			States:   s,
			Constant: true})
		return sg
	case *ast.FloatLiteral:
		s := make(map[int][]string)
		s[0] = []string{fmt.Sprint(e.Value)}
		sg := rules.NewStateGroup()
		sg.AddWrap(&rules.States{
			Base:     "__float",
			States:   s,
			Constant: true,
		})
		return sg
	case *ast.Boolean:
		s := make(map[int][]string)
		s[0] = []string{fmt.Sprint(e.Value)}
		sg := rules.NewStateGroup()
		sg.AddWrap(&rules.States{
			Base:     "__bool",
			States:   s,
			Constant: true,
		})
		return sg
	case *ast.StringLiteral:
		s := make(map[int][]string)
		s[0] = []string{fmt.Sprint(e.Value)}
		sg := rules.NewStateGroup()
		sg.AddWrap(&rules.States{
			Base:     "__string",
			States:   s,
			Constant: true,
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

		if r, ok := right.(*rules.States); ok {
			return g.mergeInvariantPrefix([]*rules.States{r}, operator)
		} else {
			return g.mergeInvariantPrefix(right.(*rules.StateGroup).Wraps, operator)
		}

	case *ast.Nil:
	case *ast.IndexExpression:
		var wg = rules.NewStateGroup()
		for _, v := range e.Left.(*ast.AssertVar).Instances {
			wg.Bases.Add(v)
			vr := g.varRounds(v, e.Index.String())
			wg.AddWrap(&rules.States{Base: v,
				States:   vr,
				Constant: true,
			})
		}
		return wg
	default:
		pos := e.Position()
		panic(fmt.Sprintf("illegal node %T in assert or assume line: %d, col: %d", e, pos[0], pos[1]))
	}
	return nil
}

func (g *Generator) mergeInvariantPrefix(right []*rules.States, operator string) *rules.StateGroup {
	sg := rules.NewStateGroup()
	for _, r := range right {
		states := make(map[int][]string)
		for i := 0; i <= g.Rounds; i++ {
			if s, ok := r.States[i]; ok {
				states[i] = append(states[i], fmt.Sprintf("(%s %s)", operator, s))
			}
		}
		r.States = states
		sg.AddWrap(r)
	}
	return sg
}

func (g *Generator) mergeInvariantInfix(left *rules.StateGroup, right *rules.StateGroup, operator string) *rules.StateGroup {
	sg := rules.NewStateGroup()
	for _, l := range left.Wraps {
		for _, r := range right.Wraps {
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

func (g *Generator) mergeByRound(left *rules.States, right *rules.States, operator string) *rules.States {
	ret := &rules.States{}

	st := make(map[int][]string)
	if left.Constant && right.Constant {
		combos := util.PairCombinations(left.States[0], right.States[0])
		st[0] = packageStateGraph(combos, operator)
		ret.States = st
		return ret
	}

	if left.Base == right.Base &&
		(left.Base != "" && right.Base != "") {
		ret.Base = left.Base

		var long map[int][]string
		var short map[int][]string
		var leftLead bool

		if len(left.States) >= len(right.States) {
			long = left.States
			short = right.States
			leftLead = true
		} else {
			long = right.States
			short = left.States
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
		ret.States = st
		return ret
	}

	if left.Constant {
		ret.Base = right.Base
		st := g.balance(right, left, operator)
		ret.States = st
		return ret
	}

	if right.Constant {
		ret.Base = left.Base
		st := g.balance(left, right, operator)
		ret.States = st
		return ret
	}

	if left.Terminal && right.Terminal {
		combos := g.termCombos(left.Base, right.Base)
		for i, c := range combos {
			st[i] = packageStateGraph(c, operator)
		}
		ret.States = st
		return ret
	}

	var llast, rlast []string
	for i := 0; i <= g.Rounds; i++ {
		var l, r []string
		var okleft, okright bool
		if l, okleft = left.States[i]; !okleft {
			if llast == nil {
				if invalidBase(left.Base) {
					panic("assert left variable base name is invalid")
				}
				l = []string{fmt.Sprintf("%s_%s", left.Base, "0")}
			} else {
				l = llast
			}
		}

		if r, okright = right.States[i]; !okright {
			if rlast == nil {
				if invalidBase(right.Base) {
					panic("assert left variable base name is invalid")
				}
				r = []string{fmt.Sprintf("%s_%s", right.Base, "0")}
			} else {
				r = rlast
			}
		}

		combos := util.PairCombinations(l, r)
		st[i] = packageStateGraph(combos, operator)
		llast = l[len(l)-1:]
		rlast = r[len(r)-1:]
	}
	ret.States = st
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

func (g *Generator) balance(vr *rules.States, con *rules.States, operator string) map[int][]string {
	ret := make(map[int][]string)
	for i := 0; i <= g.Rounds; i++ {
		if v, ok := vr.States[i]; ok {
			combos := util.PairCombinations(v, con.States[0])
			ret[i] = packageStateGraph(combos, operator)
		}
	}
	return ret
}

func (g *Generator) flattenStates(sg *rules.StateGroup) []string {
	var asserts []string
	for _, w := range sg.Wraps {
		for i := 0; i <= g.Rounds; i++ {
			if s, ok := w.States[i]; ok {
				asserts = append(asserts, s...)
			}
		}
	}
	return asserts
}

func (g *Generator) joinStates(sg *rules.StateGroup, operator string) string {
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

func (g *Generator) generateAsserts(exp ast.Expression, comp string, constr ast.Expression, stmt *ast.AssertionStatement) []*rules.Assrt {
	var ident []string
	var assrt []*rules.Assrt
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

func (g *Generator) parseInvariant(ex ast.Expression) rules.Rule {
	switch e := ex.(type) {
	case *ast.InvariantClause:

		left := g.parseInvariant(e.Left)
		right := g.parseInvariant(e.Right)

		if e.Operator == "then" {
			return &rules.Invariant{
				Left:     left,
				Operator: "then",
				Right:    right,
			}
		}

		i := &rules.Invariant{
			Left:     left,
			Operator: smtlibOperators(e.Operator),
			Right:    right,
		}
		if e.Operator == "!=" { //Not valid in SMTLib
			return &rules.Invariant{Operator: "not",
				Right: i}
		}
		return i

	case *ast.InfixExpression:
		left := g.parseInvariant(e.Left)
		right := g.parseInvariant(e.Right)
		i := &rules.Invariant{
			Left:     left,
			Operator: smtlibOperators(e.Operator),
			Right:    right,
		}
		if e.Operator == "!=" { //Not valid in SMTLib
			return &rules.Invariant{Operator: "not",
				Right: i}
		}
		return i

	case *ast.AssertVar:
		if len(e.Instances) == 1 {
			s, a, c := captureState(e.Instances[0])
			return &rules.Wrap{Value: e.Instances[0],
				State:    s,
				All:      a,
				Constant: c,
			}
		}
		var wg = &rules.WrapGroup{}
		for _, v := range e.Instances {
			s, a, c := captureState(v)
			wg.Wraps = append(wg.Wraps, &rules.Wrap{Value: v,
				State:    s,
				All:      a,
				Constant: c,
			})
		}
		return wg
	case *ast.IntegerLiteral:
		return &rules.Wrap{Value: fmt.Sprint(e.Value),
			State:    "",
			All:      false,
			Constant: true,
		}
	case *ast.FloatLiteral:
		return &rules.Wrap{Value: fmt.Sprint(e.Value),
			State:    "",
			All:      false,
			Constant: true,
		}
	case *ast.Boolean:
		return &rules.Wrap{Value: fmt.Sprint(e.Value),
			State:    "",
			All:      false,
			Constant: true,
		}
	case *ast.StringLiteral:
		return &rules.Wrap{Value: e.Value,
			State:    "",
			All:      false,
			Constant: true,
		}
	case *ast.PrefixExpression:
		right := g.parseInvariant(e.Right)
		i := &rules.Invariant{
			Left:     nil,
			Operator: smtlibOperators(e.Operator),
			Right:    right,
		}
		if e.Operator == "!" { //Not valid in SMTLib
			return &rules.Invariant{Operator: "not",
				Right: i}
		}
		return i
	case *ast.Nil:
	case *ast.IndexExpression:
		return &rules.Wrap{Value: g.convertIndexExpr(e),
			State:    "",
			All:      false,
			Constant: true,
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

func (g *Generator) generateThenRules(inv *rules.Invariant) []string {
	when := g.whenInfixNode(inv.Left)
	//then := g.thenInfixNode(inv.Right)
	fmt.Println(when)
	//fmt.Println(then)
	return []string{}
}

// check rounds for any matching variable names

// Generate all permutations of variables in the assert,
// in the round, in between variable state change

func (g *Generator) whenInfixNode(ru rules.Rule) map[string]*thenStates {
	ret := make(map[string]*thenStates)
	switch r := ru.(type) {
	case *rules.Invariant:
		left := g.whenInfixNode(r.Left)
		for k, v := range left {
			ret[k] = v
		}

		right := g.whenInfixNode(r.Right)
		for k, v := range right {
			ret[k] = v
		}

		return ret
	case *rules.Wrap:
		roundClause, values := g.whenNode(r)
		ret[r.Value] = &thenStates{
			roundClauses: roundClause,
			values:       values,
		}
		return ret
	}
	return ret
}

func (g *Generator) whenNode(when *rules.Wrap) ([]string, [][][]string) {
	var roundClauses []string
	var values [][][]string
	base := when.Value
	rounds := g.lookupVarRounds(when.Value, when.State)
	for _, r := range rounds {
		roundClauses = append(roundClauses, fmt.Sprintf("(= %s_%d %s)", base, r[0], "true"))
		values = append(values, g.RoundVars[r[1]][r[2]:])
	}
	return roundClauses, values
}

func (g *Generator) packageAssert(ident string, comp string, expr ast.Expression, stmt *ast.AssertionStatement) *rules.Assrt {
	var temporalFilter string
	var temporalN int

	temporalFilter = stmt.TemporalFilter
	temporalN = stmt.TemporalN

	s, a, c := captureState(ident)
	w := &rules.Wrap{Value: ident,
		State:    s,
		All:      a,
		Constant: c,
	}
	return &rules.Assrt{
		Variable:       w,
		Conjunction:    comp,
		Assertion:      g.parseInvariant(expr),
		TemporalFilter: temporalFilter,
		TemporalN:      temporalN}
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

func (g *Generator) generateAssertRules(ru rules.Rule, t string, tn int) []string {
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
	var i *rules.Invariant
	switch r := ru.(type) {
	case *rules.Assrt:
		return g.generateAssertRules(r.Assertion, t, tn)
	case *rules.Invariant:
		if r.Operator == "then" {
			return g.generateThenRules(r)
		}
		i = r
	case *rules.Wrap:
		return g.wrapPerm(r)
	case *rules.WrapGroup:
		var wg []string
		for _, v := range r.Wraps {
			wg = append(wg, g.wrapPerm(v)...)
		}
		return wg
	default:
		panic(fmt.Sprintf("Improperly formatted assertion %s", r.String()))
	}

	var left, right []string
	switch l := i.Left.(type) {
	case *rules.Invariant:
		left = g.generateAssertRules(l, t, tn)
	case *rules.Wrap:
		left = g.wrapPerm(l)
	default:
		left = nil
	}

	switch r := i.Right.(type) {
	case *rules.Invariant:
		right = g.generateAssertRules(r, t, tn)
	case *rules.Wrap:
		right = g.wrapPerm(r)
	default:
		right = nil
	}

	if left == nil { // Typically (not (some rule))
		var ret []string
		for _, r := range right {
			ret = append(ret, fmt.Sprintf("(%s %s)", i.Operator, r))
		}
		return ret
	}

	return expandAssertStateGraph(left, right, i.Operator, t, tn)
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

func (g *Generator) wrapPerm(w *rules.Wrap) []string {
	if w.Constant {
		return []string{w.Value}
	} else if w.State != "" {
		state := fmt.Sprint(w.Value, "_", w.State)
		return []string{state}
	}
	if w.All {
		var states []string
		end := g.variables.ssa[w.Value]
		for i := 0; i < int(end+1); i++ {
			if !g.filterOutTempStates(w.Value, int16(i)) {
				states = append(states, fmt.Sprint(w.Value, "_", i))
			}
		}
		return states
	}
	panic(fmt.Sprintf("Inproperly formatted metadata for value %s in assert", w.Value))
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
