package smt

import (
	"fault/ast"
	"fault/llvm"
	"fault/smt/rules"
	"fault/util"
	"fmt"
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
		ir, chain := g.flattenStates(sg)
		return g.applyTemporalLogic(a.Temporal, g.NewAssertChain(ir, chain, ""), a.TemporalFilter, on, off)
	}

	if a.Temporal != "" || a.TemporalFilter != "" {
		ir := g.expandAssertStateGraph(left, right, smtlibOperators(a.Constraint.Operator), a.TemporalFilter, a.TemporalN)
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
			// e.Instances not matching top level names
			// when imported
			vr := g.varRounds(v, st)
			wg.AddWrap(&rules.States{Base: v,
				Terminal: true,
				States:   vr,
				Constant: c,
			})
		}
		return wg
	case *ast.IntegerLiteral:
		s := make(map[int]*rules.AssertChain)
		s[0] = g.NewAssertChain([]string{fmt.Sprint(e.Value)}, []int{}, "")
		sg := rules.NewStateGroup()
		sg.AddWrap(&rules.States{
			Base:     "__int",
			States:   s,
			Constant: true})
		return sg
	case *ast.FloatLiteral:
		s := make(map[int]*rules.AssertChain)
		s[0] = g.NewAssertChain([]string{fmt.Sprint(e.Value)}, []int{}, "")
		sg := rules.NewStateGroup()
		sg.AddWrap(&rules.States{
			Base:     "__float",
			States:   s,
			Constant: true,
		})
		return sg
	case *ast.Boolean:
		s := make(map[int]*rules.AssertChain)
		s[0] = g.NewAssertChain([]string{fmt.Sprint(e.Value)}, []int{}, "")
		sg := rules.NewStateGroup()
		sg.AddWrap(&rules.States{
			Base:     "__bool",
			States:   s,
			Constant: true,
		})
		return sg
	case *ast.StringLiteral:
		s := make(map[int]*rules.AssertChain)
		s[0] = g.NewAssertChain([]string{fmt.Sprint(e.Value)}, []int{}, "")
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

		switch r := right.(type) {
		case *rules.Wrap:
			var numstr = r.State
			if r.Constant {
				numstr = "0"
			}
			if r.All {
				numstr = ""
			}
			vr := g.varRounds(r.Value, numstr)
			s := &rules.States{Base: r.Value,
				States:   vr,
				Constant: true,
			}
			return g.mergeInvariantPrefix([]*rules.States{s}, operator)
		case *rules.States:
			return g.mergeInvariantPrefix([]*rules.States{r}, operator)
		default:
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
		states := make(map[int]*rules.AssertChain)
		for i := 0; i <= g.Rounds; i++ {
			if st, ok := r.States[i]; ok {
				for _, s := range st.Values {
					states[i] = g.NewAssertChain([]string{}, st.Chain, "")
					states[i].Values = append(states[i].Values, fmt.Sprintf("(%s %s)", operator, s))
				}
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

	st := make(map[int]*rules.AssertChain)
	if left.Constant && right.Constant {
		combos := util.PairCombinations(left.States[0].Values, right.States[0].Values)
		st[0] = g.packageStateGraph(combos, operator, left.States[0].Chain, [][]int{})
		ret.States = st
		return ret
	}

	if left.Base == right.Base &&
		(left.Base != "" && right.Base != "") {
		ret.Base = left.Base

		var long map[int]*rules.AssertChain
		var short map[int]*rules.AssertChain
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
			var slast *rules.AssertChain

			if _, ok := long[i]; !ok {
				long[i] = &rules.AssertChain{}
			}

			var chains []int
			for idx, s := range long[i].Values {
				if sstates, ok := short[i]; ok {
					slast = sstates
					if len(sstates.Values) > idx {
						p := g.mergePairs(s, sstates.Values[idx], leftLead)
						pairs = append(pairs, p)
						if len(long[i].Chain) > 0 {
							chains = append(chains, long[i].Chain[idx])
						}
						if len(short[i].Chain) > 0 {
							chains = append(chains, short[i].Chain[idx])
						}
						continue
					}

					p := g.mergePairs(s, sstates.Values[len(sstates.Values)-1], leftLead)
					if len(long[i].Chain) > 0 {
						chains = append(chains, long[i].Chain[idx])
					}
					if len(sstates.Chain) > 0 {
						chains = append(chains, sstates.Chain[len(sstates.Chain)-1])
					}
					pairs = append(pairs, p)
					continue
				}
				p := g.mergePairs(s, slast.Values[len(slast.Values)-1], leftLead)
				chains = append(chains, []int{long[i].Chain[idx], slast.Chain[len(slast.Chain)-1]}...)
				pairs = append(pairs, p)

			}
			st[i] = g.packageStateGraph(pairs, operator, chains, [][]int{})
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
		var chains []int
		combos := g.termCombos(left.Base, right.Base)
		for i, c := range combos {
			st[i] = g.packageStateGraph(c, operator, chains, [][]int{})
		}
		ret.States = st
		return ret
	}

	var llast, rlast []string
	for i := 0; i <= g.Rounds; i++ {
		var l, r *rules.AssertChain
		var okleft, okright bool
		if l, okleft = left.States[i]; !okleft {
			if llast == nil {
				if invalidBase(left.Base) {
					panic("assert left variable base name is invalid")
				}
				l.Values = []string{fmt.Sprintf("%s_%s", left.Base, "0")}
			} else {
				l.Values = llast
			}
		}

		if r, okright = right.States[i]; !okright {
			if rlast == nil {
				if invalidBase(right.Base) {
					panic("assert left variable base name is invalid")
				}
				r.Values = []string{fmt.Sprintf("%s_%s", right.Base, "0")}
			} else {
				r.Values = rlast
			}
		}

		combos := util.PairCombinations(l.Values, r.Values)
		chains := g.matchChainToCombo(left.GetChains(), right.GetChains(), combos)
		st[i] = g.packageStateGraph(combos, operator, []int{}, chains)
		llast = l.Values[len(l.Values)-1:]
		rlast = r.Values[len(r.Values)-1:]
	}
	ret.States = st
	return ret
}

func (g *Generator) matchChainToCombo(left []int, right []int, combos [][]string) [][]int {
	var ret [][]int
	merge := append(left, right...)
	lookup := make(map[string]int)

	for _, c := range combos {
		var item []int
		if l1, ok := lookup[c[0]]; !ok {
			for _, m := range merge {
				if g.Log.Asserts[m].String() == c[0] {
					item = append(item, m)
					lookup[c[0]] = m
				}
			}
		} else {
			item = append(item, l1)
		}

		if l2, ok := lookup[c[1]]; !ok {
			for _, m := range merge {
				if g.Log.Asserts[m].String() == c[1] {
					item = append(item, m)
					lookup[c[1]] = m
				}
			}
		} else {
			item = append(item, l2)
		}
		ret = append(ret, item)
	}
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

func (g *Generator) balance(vr *rules.States, con *rules.States, operator string) map[int]*rules.AssertChain {
	ret := make(map[int]*rules.AssertChain)
	for i := 0; i <= g.Rounds; i++ {
		if v, ok := vr.States[i]; ok {
			combos := util.PairCombinations(v.Values, con.States[0].Values)
			ret[i] = g.packageStateGraph(combos, operator, v.Chain, [][]int{})
		}
	}
	return ret
}

func (g *Generator) flattenStates(sg *rules.StateGroup) ([]string, []int) {
	var asserts []string
	var chains []int
	for _, w := range sg.Wraps {
		for i := 0; i <= g.Rounds; i++ {
			if s, ok := w.States[i]; ok {
				asserts = append(asserts, s.Values...)
				chains = append(chains, s.Chain...)
			}
		}
	}
	return asserts, chains
}

func (g *Generator) joinStates(sg *rules.StateGroup, operator string) string {
	asserts, chains := g.flattenStates(sg)
	if len(asserts) == 1 {
		return asserts[0]
	}
	ret := g.writeAssertlessRule(operator, strings.Join(asserts, " "), "")
	g.Log.AssertChains[ret] = g.NewAssertChain(asserts, chains, operator)
	return ret
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

func (g *Generator) convertIndexExpr(idx *ast.IndexExpression) string {
	return strings.Join([]string{idx.Left.String(), idx.Index.String()}, "_")
}

func (g *Generator) expandAssertStateGraph(left *rules.StateGroup, right *rules.StateGroup, op string, temporalFilter string, temporalN int) *rules.AssertChain {
	var x [][]string
	list1, chain1 := g.flattenStates(left)
	list2, chain2 := g.flattenStates(right)
	c := util.Cartesian(list1, list2)
	chains := append(chain1, chain2...)
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
			var chainOn []int
			var chainOff []int
			for _, on := range p[0] {
				// Write the clauses
				i := g.Log.NewAssert(on[0], on[1], op)
				chainOn = append(chainOn, i)
				o = append(o, fmt.Sprintf("(%s %s %s)", op, on[0], on[1]))
			}
			// For nmt any of the potential on states can be on
			var onStr string
			if len(o) == 1 {
				g.Log.AssertChains[o[0]] = g.NewAssertChain(o, chainOn, op)
				onStr = o[0]
			} else {
				clause := strings.Join(o, " ")
				g.Log.AssertChains[clause] = g.NewAssertChain(o, chainOn, "or")
				onStr = fmt.Sprintf("(%s %s)", "or", clause)
			}

			offOp := llvm.OP_NEGATE[op]
			for _, off := range p[1] {
				if op == "=" {
					i := g.Log.NewAssert(off[0], off[1], "!=")
					chainOff = append(chainOff, i)
					f = append(f, fmt.Sprintf("(%s (%s %s %s))", "not", op, off[0], off[1]))
				} else {
					i := g.Log.NewAssert(off[0], off[1], offOp)
					chainOff = append(chainOff, i)
					f = append(f, fmt.Sprintf("(%s %s %s)", offOp, off[0], off[1]))
				}
			}
			// But these states must be off
			var offStr string
			if len(f) == 1 {
				g.Log.AssertChains[f[0]] = g.NewAssertChain(f, chainOff, "")
				offStr = f[0]
			} else {
				clause := strings.Join(f, " ")
				g.Log.AssertChains[clause] = g.NewAssertChain(f, chainOff, "and")
				offStr = fmt.Sprintf("(%s %s)", "and", clause)
			}
			x = append(x, []string{onStr, offStr})
		}
		return g.packageStateGraph(x, "and", chains, [][]int{})
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
		return g.packageStateGraph(x, "or", chains, [][]int{})
	default:
		return g.packageStateGraph(c, op, chains, [][]int{})
	}
}

func (g *Generator) packageStateGraph(x [][]string, op string, subchain []int, subchains [][]int) *rules.AssertChain {
	var product []string
	var chain []int
	for idx, a := range x {
		if len(subchains) > 0 {
			subchain = subchains[idx]
		}
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
			g.Log.AssertChains[s] = g.NewAssertChain(a, subchain, op)
			i := g.Log.NewAssert(a[0], a[1], op)
			chain = append(chain, i)
			product = append(product, s)
		}
	}
	return g.NewAssertChain(product, chain, "")
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
