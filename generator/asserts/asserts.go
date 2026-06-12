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
	VarTypes map[string]string // SMT sort for each base variable name
}

func NewConstraint(a *ast.AssertionStatement, rounds int, registry map[string][][]string, whens map[string][]map[string]string, varTypes map[string]string) (*Constraint, error) {
	var operator string
	stateRange := a.Constraint.Operator == "then"
	if stateRange && (a.TemporalFilter != "" || a.Temporal != "") {
		return nil, fmt.Errorf("cannot mix temporal logic with when/then assertions (%s)", a.GetToken().Location())
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
		VarTypes: varTypes,
	}, nil
}

func IsRelevant(v map[string]string, c *ast.InvariantClause) bool {
	// Check c.Left. Is it a *ast.AssertVar? If not, skip.
	if leftAssertVar, ok := c.Left.(*ast.AssertVar); ok {
		left := HasActiveInstance(leftAssertVar.Instances, v)
		if !left {
			return false
		}
	}

	//Repeat for c.Right
	if rightAssertVar, ok := c.Right.(*ast.AssertVar); ok {
		return HasActiveInstance(rightAssertVar.Instances, v)
	}

	return true
}

func HasActiveInstance(insts []string, vars map[string]string) bool {
	// Check the length of the Instances property.
	// If len == 0 return false
	// If len > 1 return true
	// If len == 1, check to see if the string is a key in v.
	// If not, then the variable is not active in this spec
	// return false. Otherwise return true

	if len(insts) == 0 {
		return false
	}
	if len(insts) > 1 {
		return true
	}
	if _, exists := vars[insts[0]]; !exists {
		return false
	}
	return true
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

	// First pass: find constant_value across all scopes before populating,
	// so map iteration order doesn't cause empty strings for earlier keys.
	constant_value := ""
	if constant {
		for _, vars := range c.Registry {
			for _, var_ssa := range vars {
				if var_ssa[0] == v {
					constant_value = strings.Join(var_ssa, "_")
					break
				}
			}
			if constant_value != "" {
				break
			}
		}
	}

	for k, vars := range c.Registry {
		if _, ok := subset[k]; !ok {
			subset[k] = util.NewStrSet()
		}

		for _, var_ssa := range vars {
			full_ssa := strings.Join(var_ssa, "_")

			if !all && full_ssa == v {
				//If not all values of variable, return only the round+scope with the exact SSA value
				subset[k].Add(full_ssa)
			}

			if !constant && var_ssa[0] == v {
				//Otherwise return every round+scope with the base name of the variable or any Instances
				subset[k].Add(full_ssa)
			}

		}

		if constant && constant_value != "" {
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
	if l != "" {
		smt = append(smt, fmt.Sprintf("(assert %s)", l))
	}
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
		assertVar, ok := e.Left.(*ast.AssertVar)
		if !ok {
			subset := c.FilterRegistryByIndex(e.Left.String(), e.Index.String())
			return rules.NewVarSets(subset)
		}
		merged := make(map[string]*util.StringSet)
		for _, inst := range assertVar.Instances {
			subset := c.FilterRegistryByIndex(inst, e.Index.String())
			merged = util.MergeStringSets(merged, subset)
		}
		return rules.NewVarSets(merged)
	case *ast.Param:
		if len(e.ProcessedName) > 0 {
			placeholder := fmt.Sprintf("__PARAM_%s__", strings.Join(e.ProcessedName, "_"))
			reg := c.RegistryConstant(placeholder)
			return rules.NewVarSets(reg)
		}
		panic(fmt.Sprintf("param() in assume/assert must be directly compared to a variable: %s", e.GetToken().Location()))
	default:
		panic(fmt.Sprintf("illegal node %T in assert or assume %s", e, e.GetToken().Location()))
	}
	return nil
}

func (c *Constraint) applyWhen() string {
	whens := c.Whens
	if len(whens) == 0 {
		// No flow steps ran — generate a static round-0 constraint for each real instance.
		whens = c.buildStaticWhens()
	}

	var ru []string
	var op string
	for _, w := range whens {
		l := c.parseWhenThen(c.Raw.Left, w)
		r := c.parseWhenThen(c.Raw.Right, w)
		if l == "" || r == "" {
			continue
		}
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
	if len(ru) == 0 {
		return ""
	}
	return fmt.Sprintf("(%s %s)", op, strings.Join(ru, " "))
}

// collectAssertVarNodes returns all *ast.AssertVar nodes in an expression tree.
func collectAssertVarNodes(node ast.Expression) []*ast.AssertVar {
	switch e := node.(type) {
	case *ast.AssertVar:
		return []*ast.AssertVar{e}
	case *ast.InfixExpression:
		return append(collectAssertVarNodes(e.Left), collectAssertVarNodes(e.Right)...)
	case *ast.PrefixExpression:
		return collectAssertVarNodes(e.Right)
	default:
		return nil
	}
}

// hasRound0Entry reports whether base appears in the round-0 registry.
func (c *Constraint) hasRound0Entry(base string) bool {
	for key, vars := range c.Registry {
		var round int
		if _, err := fmt.Sscanf(key, "round-%d_", &round); err != nil || round != 0 {
			continue
		}
		for _, v := range vars {
			if v[0] == base {
				return true
			}
		}
	}
	return false
}

// buildStaticWhens generates synthetic when/then maps for specs with no active rounds
// (c.Whens is empty because no flow functions ran). All AssertVar nodes in the
// constraint share the same BFS-expanded Instances slice, so Instances[i] on the
// left corresponds to Instances[i] on the right. For each positional index where
// the instance has a round-0 registry entry, one map is produced covering every
// AssertVar at that index.
func (c *Constraint) buildStaticWhens() []map[string]string {
	allNodes := append(collectAssertVarNodes(c.Raw.Left), collectAssertVarNodes(c.Raw.Right)...)
	if len(allNodes) == 0 {
		return nil
	}

	nInstances := len(allNodes[0].Instances)
	var result []map[string]string
	for i := 0; i < nInstances; i++ {
		w := make(map[string]string)
		hasReal := false
		for _, av := range allNodes {
			if i >= len(av.Instances) {
				continue
			}
			base := av.Instances[i]
			if c.hasRound0Entry(base) {
				w[base] = fmt.Sprintf("%s_0", base)
				hasReal = true
			}
		}
		if hasReal {
			result = append(result, w)
		}
	}
	return result
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
		if e.Operator == "choose" {
			return right
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
		// Prefer instances that have a real declaration in the registry (concrete
		// instantiations) over template-level names that appear first in the BFS
		// but have no corresponding LLVM variable.
		for _, inst := range e.Instances {
			if v, ok := w[inst]; ok && v != "" && c.hasRound0Entry(inst) {
				return v
			}
		}
		// Fallback: return the first match without registry check (covers constants
		// and other cases not tracked in the round-0 registry).
		for _, inst := range e.Instances {
			if v, ok := w[inst]; ok && v != "" {
				return v
			}
		}
		return ""
	case *ast.Param:
		if len(e.ProcessedName) > 0 {
			return fmt.Sprintf("__PARAM_%s__", strings.Join(e.ProcessedName, "_"))
		}
		panic(fmt.Sprintf("param() in assume/assert must be directly compared to a variable: %s", e.GetToken().Location()))
	default:
		panic(fmt.Sprintf("illegal node %T in assert or assume %s", e, e.GetToken().Location()))
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

// mergeSameIndex is like merge but only pairs left and right values that share
// the same SSA index suffix. This prevents cross-index pairings where a helper
// variable at one index is combined with a real state variable at another index,
// which would create spurious violation conditions for "always" assertions.
//
// Like merge(), it only pairs values from matching registry keys (k == k2).
// If no indexed values are found in the right side (e.g., right is a constant),
// it falls back to the standard cross-product merge.
func (c *Constraint) mergeSameIndex(left *rules.VarSets, right *rules.VarSets, operator string) *rules.VarSets {
	merged := rules.NewVarSets(make(map[string]*util.StringSet))

	if len(left.Vars) == 0 {
		return merged
	}

	anyRightIndexed := false
	for _, vs := range right.Vars {
		for _, v := range vs.Values() {
			if ssaIndexOf(v) != "" {
				anyRightIndexed = true
				break
			}
		}
		if anyRightIndexed {
			break
		}
	}

	// If right has no SSA-indexed variables (e.g., it's a constant literal),
	// fall back to the full cross-product merge — no index mismatch possible.
	if !anyRightIndexed {
		return c.merge(left, right, operator)
	}

	// For each registry key that appears in both left and right, pair values
	// that share the same SSA numeric index.
	for k, lVs := range left.Vars {
		rVs, ok := right.Vars[k]
		if !ok {
			continue
		}

		leftByIdx := make(map[string][]string)
		for _, v := range lVs.Values() {
			if idx := ssaIndexOf(v); idx != "" {
				leftByIdx[idx] = append(leftByIdx[idx], v)
			}
		}

		rightByIdx := make(map[string][]string)
		for _, v := range rVs.Values() {
			if idx := ssaIndexOf(v); idx != "" {
				rightByIdx[idx] = append(rightByIdx[idx], v)
			}
		}

		var combos [][]string
		for idx, lefts := range leftByIdx {
			rights, ok := rightByIdx[idx]
			if !ok {
				continue
			}
			for _, l := range lefts {
				for _, r := range rights {
					combos = append(combos, []string{l, r})
				}
			}
		}

		if len(combos) > 0 {
			merged.Vars[k] = c.Package(combos, operator)
		}
	}

	return merged
}

// ssaIndexOf extracts the trailing numeric SSA index from an SMT expression.
// Handles wrapped expressions like "(not waterpumpmonitor_Alarm_Silent_2)" → "2".
func ssaIndexOf(expr string) string {
	s := strings.TrimSpace(expr)
	// Unwrap single-argument SMT operators: (not X), (- X), etc.
	for strings.HasPrefix(s, "(") && strings.HasSuffix(s, ")") {
		inner := strings.TrimSpace(s[1 : len(s)-1])
		fields := strings.Fields(inner)
		if len(fields) == 2 {
			s = fields[1]
		} else {
			break
		}
	}
	i := strings.LastIndex(s, "_")
	if i < 0 {
		return ""
	}
	suffix := s[i+1:]
	if _, err := strconv.Atoi(suffix); err != nil {
		return ""
	}
	return suffix
}

// varSMTType returns the SMT sort for an SSA-versioned variable name,
// stripping the trailing numeric suffix to look up the base name in VarTypes.
func (c *Constraint) varSMTType(ssaName string) string {
	if c.VarTypes == nil {
		return ""
	}
	if ty, ok := c.VarTypes[ssaName]; ok {
		return ty
	}
	// Strip trailing _N SSA suffix
	parts := strings.Split(ssaName, "_")
	if len(parts) > 1 {
		if _, err := strconv.Atoi(parts[len(parts)-1]); err == nil {
			base := strings.Join(parts[:len(parts)-1], "_")
			return c.VarTypes[base]
		}
	}
	return ""
}

func (c *Constraint) Package(x [][]string, op string) *util.StringSet {
	product := util.NewStrSet()
	for _, a := range x {
		if len(a) == 1 {
			product.Add(a[0])
		} else {
			var s string
			if op == "not" && a[0] == "false" {
				// pair is [false, varName]: negation of (var == false) → var is not false
				if c.varSMTType(a[1]) == "Real" {
					s = fmt.Sprintf("(not (= %s 0.0))", a[1])
				} else {
					s = fmt.Sprintf("(%s %s)", op, a[1])
				}
			} else if op == "not" && a[1] == "false" {
				// pair is [varName, false]: negation of (var == false) → var is not false
				if c.varSMTType(a[0]) == "Real" {
					s = fmt.Sprintf("(not (= %s 0.0))", a[0])
				} else {
					s = fmt.Sprintf("(%s %s)", op, a[0])
				}
			} else if op == "not" {
				// negation of (var == something): replace bool literals with numerics for Real vars
				lhs, rhs := a[0], a[1]
				if c.varSMTType(lhs) == "Real" {
					if rhs == "true" {
						rhs = "1.0"
					} else if rhs == "false" {
						rhs = "0.0"
					}
				} else if c.varSMTType(rhs) == "Real" {
					if lhs == "true" {
						lhs = "1.0"
					} else if lhs == "false" {
						lhs = "0.0"
					}
				}
				s = fmt.Sprintf("(%s (= %s %s))", op, lhs, rhs)
			} else {
				// plain comparison: replace bool literals with numerics for Real vars
				lhs, rhs := a[0], a[1]
				if c.varSMTType(lhs) == "Real" {
					if rhs == "true" {
						rhs = "1.0"
					} else if rhs == "false" {
						rhs = "0.0"
					}
				} else if c.varSMTType(rhs) == "Real" {
					if lhs == "true" {
						lhs = "1.0"
					} else if lhs == "false" {
						lhs = "0.0"
					}
				}
				s = fmt.Sprintf("(%s %s %s)", op, lhs, rhs)
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
		if len(m.List()) == 0 {
			return ""
		}
		return fmt.Sprintf("(%s %s)", c.On, strings.Join(m.List(), " "))
	case "always": // Every state is true
		m := c.mergeSameIndex(c.Left, c.Right, c.Op)
		if len(m.List()) == 0 {
			return ""
		}
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
		inverse, err := util.SliceOfIndex(len(merged))
		if err != nil {
			panic(err) // len() is always >= 0; this cannot happen
		}

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
