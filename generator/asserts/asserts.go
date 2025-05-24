package asserts

import (
	"fault/ast"
	"fault/generator/rules"
	"fault/util"
	"fmt"
	"strconv"
	"strings"
)

type Time struct {
	Filter string
	Type   string
	N      int
}

type Constraint struct {
	Raw      *ast.InvariantClause
	Left     *rules.VarSets
	Right    *rules.VarSets
	Op       string
	On       string
	Off      string
	Temporal *Time
	Then     bool
	Assume   bool
	Rounds   int
	Registry map[string][][]string
}

func NewConstraint(a *ast.AssertionStatement, rounds int, registry map[string][][]string) *Constraint {
	var operator string
	stateRange := a.Constraint.Operator == "then"
	if stateRange && (a.TemporalFilter != "" || a.Temporal != "") {
		panic("cannot mix temporal logic with when/then assertions")
	}

	operator = smtlibOperators(a.Constraint.Operator)

	if stateRange {
		operator = "and"
	}

	var on, off string
	if a.Assume {
		on = "or"
		off = "and"
	} else {
		on = "and"
		off = "or"
	}

	return &Constraint{
		Raw:  a.Constraint,
		Then: stateRange,
		Op:   operator,
		Temporal: &Time{
			Filter: a.TemporalFilter,
			Type:   a.Temporal,
			N:      a.TemporalN,
		},
		On:       on,
		Off:      off,
		Assume:   a.Assume,
		Rounds:   rounds,
		Registry: registry,
	}
}

func (c *Constraint) RegistryConstant(cons string) map[string][]string {
	subset := make(map[string][]string)
	for k := range c.Registry {
		subset[k] = []string{cons}
	}
	return subset
}

func (c *Constraint) FilterRegistry(v string, constant bool, all bool) map[string][]string {
	// Filter the registry to only include rounds+scopes where the
	// variable is relevant

	subset := make(map[string][]string)
	constant_value := ""
	for k, vars := range c.Registry {
		for _, var_ssa := range vars {
			full_ssa := strings.Join(var_ssa, "_")

			if constant && (var_ssa[0] == v) && (constant_value == "") {
				constant_value = full_ssa
			}

			if !all && full_ssa == v {
				//If not all values of variable, return only the round+scope with the exact SSA value
				subset[k] = []string{full_ssa}
			}

			if !constant && var_ssa[0] == v {
				//Otherwise return every round+scope with the base name of the variable or any Instances
				if _, ok := subset[k]; !ok {
					subset[k] = []string{full_ssa}
					continue
				}
				subset[k] = append(subset[k], full_ssa)
			}

		}

		if constant {
			//If var is constant, populate every round+scope with the correct SSA value
			if _, ok := subset[k]; !ok {
				subset[k] = []string{constant_value}
				continue
			}
			subset[k] = append(subset[k], constant_value)
		}
	}
	return subset
}

func (c *Constraint) FilterRegistryByIndex(v string, idx string) map[string][]string {
	subset := make(map[string][]string)
	for k, vars := range c.Registry {
		for _, var_ssa := range vars {
			if var_ssa[0] == v && var_ssa[1] == idx {
				full_ssa := strings.Join(var_ssa, "_")
				subset[k] = []string{full_ssa}
			}
		}
	}
	return subset
}

func (c *Constraint) Parse() []string {
	c.Left = c.parseNode(c.Raw.Left)
	c.Right = c.parseNode(c.Raw.Right)

	// Then
	// If left is true, right is true
	// If left is false, right doesn't matter

	// Temporals
	// If left and right are instances of the same variable
	l := c.applyTemporal()

	// Expand based on temporal filter

	// If assume, the different asserts need to be joined by
	// an "and" instead of and "or"

	// if c.Then {
	// 	sg := c.mergeInvariantInfix(left, right, "or")
	// 	return c.joinStates(sg, c.Op)
	// }

	//If left and right are asserts on the same variable
	// dset := util.DiffStrSets(c.Left.Bases, c.Right.Bases)
	// if dset.Len() == 0 && (c.Temporal.Type != "" || c.Temporal.Filter != "") {
	// 	sg := c.merge(left, right)
	// 	ir, chain := c.flatten(sg)
	// 	assertChain := c.NewAssertChain(ir, chain, "")
	// 	return c.applyTemporal(assertChain)
	// }

	// if c.Temporal.Type != "" || c.Temporal.Filter != "" {
	// 	ir := c.expand(left, right, c.Op, c.Temporal.Filter, c.Temporal.N)
	// 	return c.applyTemporal(c.Temporal.Type, ir, c.Temporal.Filter, c.On, c.Off)
	// }
	// if c.Assume {
	// 	operator := "and"

	// 	// if operator == op { // (and (and ) (and )) is redundant
	// 	// 	sg := rules.NewPossibleVars()
	// 	// 	sg.Wraps = append(left.Wraps, right.Wraps...)
	// 	// 	return g.joinStates(sg, operator)
	// 	// }

	// 	sg := c.merge(left, right)
	// 	return c.join(sg, operator)
	// }

	// operator := "or"
	// // if operator == c.Op {
	// // 	sg := rules.NewPossibleVars()
	// // 	sg.Wraps = append(left.Wraps, right.Wraps...)
	// // 	return g.joinStates(sg, operator)
	// // }
	// sg := c.merge(left, right)
	// return c.join(sg, operator)
	var smt []string
	smt = append(smt, fmt.Sprintf("(assert %s)", l))
	return smt
}

func (c *Constraint) parseNode(exp ast.Expression) *rules.VarSets {
	switch e := exp.(type) {
	case *ast.InfixExpression:
		operator := smtlibOperators(e.Operator)
		left := c.parseNode(e.Left)
		right := c.parseNode(e.Right)

		return c.merge(left, right, operator)

	case *ast.AssertVar:
		var registery_subset = make(map[string][]string)
		for _, v := range e.Instances {
			var subset2 map[string][]string
			_, a, cons := captureState(v)
			subset2 = c.FilterRegistry(v, cons, a)
			registery_subset = util.MergeStringSliceMaps(registery_subset, subset2) //If the variable is not in the registry, add it
		}
		return rules.NewVarSets(registery_subset)

	case *ast.IntegerLiteral:
		reg := c.RegistryConstant(fmt.Sprintf("%d", e.Value))
		return rules.NewVarSets(reg)
	case *ast.FloatLiteral:
		reg := c.RegistryConstant(fmt.Sprintf("%v", e.Value))
		return rules.NewVarSets(reg)
	case *ast.Boolean:
		reg := c.RegistryConstant(fmt.Sprintf("%v", e.Value))
		return rules.NewVarSets(reg)
	case *ast.StringLiteral:
		reg := c.RegistryConstant(e.Value)
		return rules.NewVarSets(reg)
	case *ast.PrefixExpression:
		var operator string
		right := c.parseNode(e.Right)
		if e.Operator == "!" { //Not valid in SMTLib
			operator = "not"
		} else {
			operator = smtlibOperators(e.Operator)
		}

		prefix := make(map[string][]string)
		for k, v := range right.Vars {
			prefix[k] = c.Prefix(v, operator)
		}
		return rules.NewVarSets(prefix)

	case *ast.Nil:
	case *ast.IndexExpression:
		subset := c.FilterRegistryByIndex(e.Left.String(), e.Index.String())
		return rules.NewVarSets(subset)
	default:
		pos := e.Position()
		panic(fmt.Sprintf("illegal node %T in assert or assume line: %d, col: %d", e, pos[0], pos[1]))
	}
	return nil
}

func (c *Constraint) Prefix(x []string, op string) []string {
	var product []string
	for _, a := range x {
		product = append(product, fmt.Sprintf("(%s %s)", op, a))
	}
	return product
}

func captureState(id string) (string, bool, bool) {
	//Returns base name, where it is a constant (c), and
	// whether to apply assert to all SSA versions of
	// the variable (a)
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
	if err != nil { //Last part is not a number so constant or all
		return "", a, c
	} else { //Last part is a number so rule only applies to ONE specific instance of the variable
		return raw[len(raw)-1], false, false
	}

}

func (c *Constraint) merge(left *rules.VarSets, right *rules.VarSets, operator string) *rules.VarSets {
	merged := rules.NewVarSets(make(map[string][]string))

	if len(left.Vars) > 0 {
		for k, l := range left.Vars {
			for k2, l2 := range right.Vars {
				if k == k2 { //For now
					combos := util.PairCombinations(l, l2)
					merged.Vars[k] = c.Package(combos, operator)
				}
			}
		}
	}
	return merged
}

func (c *Constraint) Package(x [][]string, op string) []string {
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

// func (c *Constraint) flatten(sg *rules.PossibleVars) ([]string, []int) {
// 	var asserts []string
// 	var chains []int
// 	for _, w := range sg.Wraps {
// 		for i := 0; i <= c.Rounds; i++ {
// 			if s, ok := w.States[i]; ok {
// 				asserts = append(asserts, s.Values...)
// 				chains = append(chains, s.Chain...)
// 			}
// 		}
// 	}
// 	return asserts, chains
// }

func (c *Constraint) applyTemporal() string {
	// full := ir.Values
	// first := ir.Values[0]
	// clause := strings.Join(full, " ")

	switch c.Temporal.Type {
	case "eventually": // At least one state is true
		m := c.merge(c.Left, c.Right, c.Op)
		return fmt.Sprintf("(%s %s)", c.On, strings.Join(m.List(), " "))
	case "always": // Every state is true
		m := c.merge(c.Left, c.Right, c.Op)
		return fmt.Sprintf("(%s %s)", c.Off, strings.Join(m.List(), " "))
	case "eventually-always": // Once the statement is true, it stays true
		m := c.merge(c.Left, c.Right, c.Op)
		or := c.eventuallyAlways(m.List())
		return or
		// if len(full) > 1 {
		// 	or := c.eventuallyAlways(ir)
		// 	return or
		// }
		// return first
	default:
		var op string
		switch c.Temporal.Filter {
		case "nft": // True no fewer than X times
			op = "or"
		case "nmt": // True no more than X times
			op = "or"
		default:
			op = c.Off
		}
		clause := c.merge(c.Left, c.Right, c.Op)

		if len(clause.List()) == 1 {
			return clause.List()[0]
		}

		or := fmt.Sprintf("(%s %s)", op, strings.Join(clause.List(), " "))
		return or
	}
}

// func (c *Constraint) expand() *rules.AssertChain {
// 	var x [][]string
// 	//list1, chain1 := g.flattenStates(left)
// 	//list2, chain2 := g.flattenStates(right)
// 	list1, _ := g.flattenStates(left)
// 	list2, _ := g.flattenStates(right)
// 	c := util.Cartesian(list1, list2)
// 	//chains := g.matchChainToCombo(chain1, chain2, c)

// 	switch temporalFilter {
// 	// For logic like "no more than X times" "no fewer than X times"
// 	// We need to flip some of the operators and build out more
// 	// states before packaging the asserts
// 	case "nmt":
// 		// (and (or on on on) (and off off))
// 		combos := util.Combinations(c, temporalN) //generate all combinations of possible on states
// 		pairs := impliesOnOffPairs(combos, c)     //for each combination prepare a list of states that must logically be off
// 		for _, p := range pairs {
// 			var o []string
// 			var f []string
// 			var chainOn []int
// 			var chainOff []int
// 			for _, on := range p[0] {
// 				// Write the clauses
// 				i := g.Log.NewAssert(on[0], on[1], op)
// 				chainOn = append(chainOn, i)
// 				clause := fmt.Sprintf("(%s %s %s)", op, on[0], on[1])
// 				o = append(o, clause)
// 				g.Log.AddChain(clause, g.NewMultiVAssertChain(on, []int{}, op))
// 			}
// 			// For nmt any of the potential on states can be on
// 			var onStr string
// 			if len(o) == 1 {
// 				g.Log.AddChain(o[0], g.NewAssertChain(o, chainOn, op))
// 				onStr = o[0]
// 			} else {
// 				clause := strings.Join(o, " ")
// 				g.Log.NewMultiClauseAssert(o, "or")
// 				g.Log.AddChain(clause, g.NewMultiVAssertChain(o, chainOn, "or"))
// 				onStr = fmt.Sprintf("(%s %s)", "or", clause)
// 			}

// 			offOp := util.OP_NEGATE[op]
// 			for _, off := range p[1] {
// 				if op == "=" {
// 					clause := fmt.Sprintf("(%s (%s %s %s))", "not", op, off[0], off[1])
// 					g.Log.AddChain(clause, g.NewMultiVAssertChain(off, []int{}, "!="))
// 					i := g.Log.NewAssert(off[0], off[1], "!=")
// 					chainOff = append(chainOff, i)
// 					f = append(f, clause)
// 				} else {
// 					clause := fmt.Sprintf("(%s %s %s)", offOp, off[0], off[1])
// 					g.Log.AddChain(clause, g.NewMultiVAssertChain(off, []int{}, offOp))
// 					i := g.Log.NewAssert(off[0], off[1], offOp)
// 					chainOff = append(chainOff, i)
// 					f = append(f, clause)
// 				}
// 			}
// 			// But these states must be off
// 			var offStr string
// 			if len(f) == 1 {
// 				g.Log.AddChain(f[0], g.NewAssertChain(f, chainOff, ""))
// 				offStr = f[0]
// 			} else {
// 				clause := strings.Join(f, " ")
// 				g.Log.AddChain(clause, g.NewMultiVAssertChain(f, chainOff, "and"))
// 				g.Log.NewMultiClauseAssert(f, "and")
// 				offStr = fmt.Sprintf("(%s %s)", "and", clause)
// 			}
// 			x = append(x, []string{onStr, offStr})
// 		}
// 		return g.packageStateGraph(x, "and", []int{}, [][]int{})
// 	case "nft":
// 		// (or (and on on on))
// 		combos := util.Combinations(c, temporalN)
// 		pairs := impliesOnOffPairs(combos, c)
// 		for _, p := range pairs {
// 			var o []string
// 			for _, on := range p[0] {
// 				o = append(o, fmt.Sprintf("(%s %s %s)", op, on[0], on[1]))
// 			}
// 			// For nft all on states in this possibility MUST be on
// 			var onStr string
// 			if len(o) == 1 {
// 				onStr = o[0]
// 			} else {
// 				g.Log.NewMultiClauseAssert(o, "and")
// 				onStr = fmt.Sprintf("(%s %s)", "and", strings.Join(o, " "))
// 			}
// 			x = append(x, []string{onStr})
// 		}
// 		return g.packageStateGraph(x, "or", []int{}, [][]int{})
// 	default:
// 		return g.packageStateGraph(c, op, []int{}, [][]int{})
// 	}
// }

func (c *Constraint) eventuallyAlways(values []string) string {
	var clause string
	var progression []string
	for i := 1; i <= len(values); i++ {
		clause = strings.Join(values[len(values)-i:], " ")
		s := fmt.Sprintf("(and %s)", clause)
		progression = append(progression, s)
	}

	parentClause := strings.Join(progression, " ")
	return fmt.Sprintf("(or %s)", parentClause)
}

// func (c *Constraint) mergeByRound(left_base *rules.VarSets, right *rules.VarSets, operator string) *rules.VarSets {
// 	ret := &rules.VarSets{}

// 	st := make(map[int]*rules.AssertChain)
// 	if left.Constant && right.Constant {
// 		combos := util.PairCombinations(left.States[0].Values, right.States[0].Values)
// 		st[0] = g.packageStateGraph(combos, operator, left.States[0].Chain, [][]int{})
// 		ret.States = st
// 		return ret
// 	}

// 	if left.Base == right.Base &&
// 		(left.Base != "" && right.Base != "") {
// 		ret.Base = left.Base

// 		var long map[int]*rules.AssertChain
// 		var short map[int]*rules.AssertChain
// 		var leftLead bool

// 		if len(left.States) >= len(right.States) {
// 			long = left.States
// 			short = right.States
// 			leftLead = true
// 		} else {
// 			long = right.States
// 			short = left.States
// 			leftLead = false
// 		}

// 		//Pair based on same state
// 		for i := 0; i <= g.Rounds; i++ {
// 			var pairs [][]string
// 			var slast *rules.AssertChain

// 			if _, ok := long[i]; !ok {
// 				long[i] = &rules.AssertChain{}
// 			}

// 			var chains []int
// 			for idx, s := range long[i].Values {
// 				if sstates, ok := short[i]; ok {
// 					slast = sstates
// 					if len(sstates.Values) > idx {
// 						p := g.mergePairs(s, sstates.Values[idx], leftLead)
// 						pairs = append(pairs, p)
// 						if len(long[i].Chain) > 0 {
// 							chains = append(chains, long[i].Chain[idx])
// 						}
// 						if len(short[i].Chain) > 0 {
// 							chains = append(chains, short[i].Chain[idx])
// 						}
// 						continue
// 					}

// 					p := g.mergePairs(s, sstates.Values[len(sstates.Values)-1], leftLead)
// 					if len(long[i].Chain) > 0 {
// 						chains = append(chains, long[i].Chain[idx])
// 					}
// 					if len(sstates.Chain) > 0 {
// 						chains = append(chains, sstates.Chain[len(sstates.Chain)-1])
// 					}
// 					pairs = append(pairs, p)
// 					continue
// 				}
// 				p := g.mergePairs(s, slast.Values[len(slast.Values)-1], leftLead)
// 				chains = append(chains, []int{long[i].Chain[idx], slast.Chain[len(slast.Chain)-1]}...)
// 				pairs = append(pairs, p)

// 			}
// 			st[i] = g.packageStateGraph(pairs, operator, chains, [][]int{})
// 		}
// 		ret.States = st
// 		return ret
// 	}

// 	if left.Constant {
// 		ret.Base = right.Base
// 		st := g.balance(right, left, operator)
// 		ret.States = st
// 		return ret
// 	}

// 	if right.Constant {
// 		ret.Base = left.Base
// 		st := g.balance(left, right, operator)
// 		ret.States = st
// 		return ret
// 	}

// 	if left.Terminal && right.Terminal {
// 		var chains []int
// 		combos := g.termCombos(left.Base, right.Base)
// 		for i, c := range combos {
// 			st[i] = g.packageStateGraph(c, operator, chains, [][]int{})
// 		}
// 		ret.States = st
// 		return ret
// 	}

// 	var llast, rlast []string
// 	for i := 0; i <= g.Rounds; i++ {
// 		var l, r *rules.AssertChain
// 		var okleft, okright bool
// 		if l, okleft = left.States[i]; !okleft {
// 			if llast == nil {
// 				if invalidBase(left.Base) {
// 					panic("assert left variable base name is invalid")
// 				}
// 				l.Values = []string{fmt.Sprintf("%s_%s", left.Base, "0")}
// 			} else {
// 				l.Values = llast
// 			}
// 		}

// 		if r, okright = right.States[i]; !okright {
// 			if rlast == nil {
// 				if invalidBase(right.Base) {
// 					panic("assert left variable base name is invalid")
// 				}
// 				r.Values = []string{fmt.Sprintf("%s_%s", right.Base, "0")}
// 			} else {
// 				r.Values = rlast
// 			}
// 		}

// 		combos := util.PairCombinations(l.Values, r.Values)
// 		chains := g.matchChainToCombo(left.GetChains(), right.GetChains(), combos)
// 		st[i] = g.packageStateGraph(combos, operator, []int{}, chains)
// 		llast = l.Values[len(l.Values)-1:]
// 		rlast = r.Values[len(r.Values)-1:]
// 	}
// 	ret.States = st
// 	return ret
// }

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
