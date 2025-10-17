package asserts

import (
	"fault/ast"
	"fault/generator/rules"
	"fault/util"
	"fmt"
	"strconv"
	"strings"

	"gonum.org/v1/gonum/stat/combin"
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
	Whens    []map[string]string
}

func NewConstraint(a *ast.AssertionStatement, rounds int, registry map[string][][]string, whens map[string][]map[string]string) *Constraint {
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
		Whens:    whens[a.String()],
	}
}

func (c *Constraint) RegistryConstant(cons string) map[string]*util.StringSet {
	subset := make(map[string]*util.StringSet)
	for k := range c.Registry {
		if _, ok := subset[k]; !ok {
			subset[k] = util.NewStrSet()
		}

		subset[k].Add(cons)
	}
	return subset
}

func (c *Constraint) FilterRegistry(v string, constant bool, all bool) map[string]*util.StringSet {
	// Filter the registry to only include rounds+scopes where the
	// variable is relevant

	subset := make(map[string]*util.StringSet)
	constant_value := ""
	for k, vars := range c.Registry {
		if _, ok := subset[k]; !ok {
			subset[k] = util.NewStrSet()
		}

		for _, var_ssa := range vars {
			full_ssa := strings.Join(var_ssa, "_")

			if constant && (var_ssa[0] == v) && (constant_value == "") {
				constant_value = full_ssa
			}

			if !all && full_ssa == v {
				//If not all values of variable, return only the round+scope with the exact SSA value
				subset[k].Add(full_ssa)
			}

			if !constant && var_ssa[0] == v {
				//Otherwise return every round+scope with the base name of the variable or any Instances
				subset[k].Add(full_ssa)
			}

		}

		if constant {
			//If var is constant, populate every round+scope with the correct SSA value
			subset[k].Add(constant_value)
		}
	}
	return subset
}

func (c *Constraint) FilterRegistryByIndex(v string, idx string) map[string]*util.StringSet {
	subset := make(map[string]*util.StringSet)
	for k, vars := range c.Registry {
		if _, ok := subset[k]; !ok {
			subset[k] = util.NewStrSet()
		}

		for _, var_ssa := range vars {
			if var_ssa[0] == v && var_ssa[1] == idx {
				full_ssa := strings.Join(var_ssa, "_")
				subset[k].Add(full_ssa)
			}
		}
	}
	return subset
}

func (c *Constraint) Parse() []string {

	// Then
	// If left is true, right is true
	// If left is false, right doesn't matter

	// Temporals
	// If left and right are instances of the same variable
	var l string
	if c.Then {
		l = c.applyWhen()
	} else {
		c.Left = c.parseNode(c.Raw.Left)
		c.Right = c.parseNode(c.Raw.Right)
		l = c.applyTemporal()
	}

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
		var registery_subset = make(map[string]*util.StringSet)
		for _, v := range e.Instances {
			var subset2 map[string]*util.StringSet
			_, a, cons := captureState(v)
			subset2 = c.FilterRegistry(v, cons, a)
			registery_subset = util.MergeStringSets(registery_subset, subset2) //If the variable is not in the registry, add it
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

		prefix := make(map[string]*util.StringSet)
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

func (c *Constraint) applyWhen() string {
	var ru []string
	var op string
	for _, w := range c.Whens {
		l := c.parseWhenThen(c.Raw.Left, w)
		r := c.parseWhenThen(c.Raw.Right, w)
		var rule string
		if c.Assume {
			op = "and"
			rule = fmt.Sprintf("(=> %s %s)", l, r)
		} else {
			op = "or"
			rule = fmt.Sprintf("(and %s (not %s))", l, r)
		}

		ru = append(ru, rule)
	}
	return fmt.Sprintf("(%s %s)", op, strings.Join(ru, " "))
}

func (c *Constraint) parseWhenThen(node ast.Expression, w map[string]string) string {
	switch e := node.(type) {
	case *ast.InfixExpression:
		left := c.parseWhenThen(e.Left, w)
		right := c.parseWhenThen(e.Right, w)
		op := smtlibOperators(e.Operator)

		if op == "not" {
			return fmt.Sprintf("(distinct %s %s)", left, right)
		}
		return fmt.Sprintf("(%s %s %s)", op, left, right)

	case *ast.PrefixExpression:
		right := c.parseWhenThen(e.Right, w)
		if e.Operator == "!" { //Not valid in SMTLib
			return fmt.Sprintf("(not %s)", right)
		}
		return fmt.Sprintf("(%s %s)", smtlibOperators(e.Operator), right)

	case *ast.IntegerLiteral:
		return fmt.Sprintf("%d", e.Value)
	case *ast.FloatLiteral:
		return fmt.Sprintf("%v", e.Value)
	case *ast.Boolean:
		return fmt.Sprintf("%v", e.Value)
	case *ast.StringLiteral:
		return e.Value
	case *ast.AssertVar:
		return w[e.Instances[0]]
	default:
		pos := e.Position()
		panic(fmt.Sprintf("illegal node %T in assert or assume line: %d, col: %d", e, pos[0], pos[1]))
	}

}

func (c *Constraint) Prefix(x *util.StringSet, op string) *util.StringSet {
	product := util.NewStrSet()
	for _, a := range x.Values() {
		product.Add(fmt.Sprintf("(%s %s)", op, a))
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
	merged := rules.NewVarSets(make(map[string]*util.StringSet))

	if len(left.Vars) > 0 {
		for k, l := range left.Vars {
			for k2, l2 := range right.Vars {
				if k == k2 { //For now
					combos := util.PairCombinations(l.Values(), l2.Values())
					merged.Vars[k] = c.Package(combos, operator)
				}
			}
		}
	}
	return merged
}

func (c *Constraint) Package(x [][]string, op string) *util.StringSet {
	product := util.NewStrSet()
	for _, a := range x {
		if len(a) == 1 {
			product.Add(a[0])
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
			product.Add(s)
		}
	}

	return product
}

func (c *Constraint) applyTemporal() string {
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

	default:
		var op string
		clause := c.merge(c.Left, c.Right, c.Op)

		switch c.Temporal.Filter {
		case "nft": // True no fewer than X times
			op = "or"
			m := c.NoFew(clause.List(), c.Temporal.N)
			if m[0] == "" {
				return ""
			}
			if len(m) == 1 {
				return m[0]
			}
			return fmt.Sprintf("(%s %s)", op, strings.Join(m, " "))
		case "nmt": // True no more than X times
			op = "or"
			m := c.NoMore(clause.List(), c.Temporal.N)
			if m[0] == "" {
				return ""
			}
			if len(m) == 1 {
				return m[0]
			}
			return fmt.Sprintf("(%s %s)", op, strings.Join(m, " "))
		default:
			op = c.Off
		}

		if len(clause.List()) == 1 {
			return clause.List()[0]
		}

		or := fmt.Sprintf("(%s %s)", op, strings.Join(clause.List(), " "))
		return or
	}
}

func (c *Constraint) NoFew(merged []string, n int) []string {
	// No more than n times
	// No fewer than n times
	// If n is 1, then it is just a single assert
	if len(merged) < n {
		return []string{strings.Join(merged, " ")}
	}

	if n == 1 {
		return merged
	}

	if n == 0 {
		return []string{}
	}

	if merged[0][0:4] == "(and" {
		return merged
	}

	combos := combin.Combinations(len(merged), n)

	//Assembling combinations
	var ret []string
	for _, combo := range combos {
		var clause []string
		for _, i := range combo {
			clause = append(clause, merged[i])
		}
		if len(clause) == 1 {
			ret = append(ret, clause[0])
		} else if clause[0][0:4] == "(and" {
			ret = append(ret, clause...)
		} else {
			ret = append(ret, fmt.Sprintf("(and %s)", strings.Join(clause, " ")))
		}
	}
	return ret
}

func (c *Constraint) NoMore(merged []string, n int) []string {
	// No more than n times
	// No fewer than n times
	// If n is 1, then it is just a single assert
	if len(merged) < n {
		return []string{strings.Join(merged, " ")}
	}

	if n == 1 {
		return merged
	}

	if n == 0 {
		return []string{}
	}

	if merged[0][0:4] == "(and" {
		return merged
	}

	combos := combin.Combinations(len(merged), n)

	//Assembling combinations
	var ret []string
	for _, combo := range combos {
		// Create inverse of the combination
		inverse := util.SliceOfIndex(len(merged))

		var clause []string
		var nots []string
		for _, i := range combo {
			// Remove the index from the inverse
			inverse[i] = -1

			clause = append(clause, merged[i])
		}

		for _, i := range inverse {
			if i > -1 {
				nots = append(nots, fmt.Sprintf("(not %s)", merged[i]))
			}
		}

		if len(clause) == 1 {
			ret = append(ret, clause[0])
		} else if clause[0][0:4] == "(and" {
			ret = append(ret, clause...)
		} else {
			on := fmt.Sprintf("(or %s)", strings.Join(clause, " "))
			off := fmt.Sprintf("(and %s)", strings.Join(nots, " "))

			ret = append(ret, fmt.Sprintf("(and %s %s)", on, off))
		}
	}
	return ret
}

func (c *Constraint) eventuallyAlways(values []string) string {
	var clause string
	var progression []string
	for i := 1; i <= len(values); i++ {
		var s string
		if len(values[len(values)-i:]) > 1 { // Don't need and for single values
			clause = strings.Join(values[len(values)-i:], " ")
			s = fmt.Sprintf("(and %s)", clause)
		} else {
			s = values[len(values)-i]
		}
		progression = append(progression, s)
	}

	parentClause := strings.Join(progression, " ")
	return fmt.Sprintf("(or %s)", parentClause)
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
