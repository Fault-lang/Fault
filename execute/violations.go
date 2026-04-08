package execute

import (
	"fault/ast"
	"fmt"
	"strconv"
)

// EvaluateViolations checks each assertion against the model stored in
// ResultValues and sets a.Violated accordingly.
//
// Assertions stored in the AST have already been negated by the LLVM compiler
// (e.g., source `X > 0` becomes stored `X <= 0`). So if the stored condition
// holds for a given round, the original assertion was violated in that round.
func (mc *ModelChecker) EvaluateViolations(asserts []*ast.AssertionStatement) {
	for _, a := range asserts {
		a.Violated = assertionViolated(a, mc.ResultValues)
	}
}

func assertionViolated(a *ast.AssertionStatement, values map[string]string) bool {
	if a.TemporalFilter != "" {
		// nft/nmt filters require combinatorial counting; not yet supported.
		return false
	}

	rounds := roundsForConstraint(a.Constraint, values)
	if len(rounds) == 0 {
		return false
	}

	violatedCount := 0
	for _, round := range rounds {
		if holds, ok := evalConstraint(a.Constraint, round, values); ok && holds {
			violatedCount++
		}
	}

	switch a.Temporal {
	case "eventually":
		// Original `eventually X` was negated to `always not X` (and across all
		// rounds). It is violated when the stored (negated) condition holds in
		// every round — meaning X was never true.
		return violatedCount == len(rounds)
	default:
		// Original `always X` or simple `X` was negated. Violated when the
		// stored condition holds in at least one round.
		return violatedCount > 0
	}
}

// evalConstraint evaluates the top-level InvariantClause at a given round.
func evalConstraint(c *ast.InvariantClause, round int16, values map[string]string) (bool, bool) {
	switch c.Operator {
	case "||":
		lv, lok := evalBoolExpr(c.Left, round, values)
		rv, rok := evalBoolExpr(c.Right, round, values)
		if !lok || !rok {
			return false, false
		}
		return lv || rv, true
	case "&&":
		lv, lok := evalBoolExpr(c.Left, round, values)
		rv, rok := evalBoolExpr(c.Right, round, values)
		if !lok || !rok {
			return false, false
		}
		return lv && rv, true
	default:
		lv, lok := evalNumericExpr(c.Left, round, values)
		rv, rok := evalNumericExpr(c.Right, round, values)
		if !lok || !rok {
			return false, false
		}
		return applyCompare(c.Operator, lv, rv), true
	}
}

// evalBoolExpr evaluates an expression that is expected to produce a boolean.
func evalBoolExpr(expr ast.Expression, round int16, values map[string]string) (bool, bool) {
	switch e := expr.(type) {
	case *ast.InfixExpression:
		switch e.Operator {
		case "||":
			lv, lok := evalBoolExpr(e.Left, round, values)
			rv, rok := evalBoolExpr(e.Right, round, values)
			if !lok || !rok {
				return false, false
			}
			return lv || rv, true
		case "&&":
			lv, lok := evalBoolExpr(e.Left, round, values)
			rv, rok := evalBoolExpr(e.Right, round, values)
			if !lok || !rok {
				return false, false
			}
			return lv && rv, true
		default:
			lv, lok := evalNumericExpr(e.Left, round, values)
			rv, rok := evalNumericExpr(e.Right, round, values)
			if !lok || !rok {
				return false, false
			}
			return applyCompare(e.Operator, lv, rv), true
		}
	case *ast.PrefixExpression:
		if e.Operator == "!" {
			rv, ok := evalBoolExpr(e.Right, round, values)
			if !ok {
				return false, false
			}
			return !rv, true
		}
	case *ast.AssertVar:
		for _, inst := range e.Instances {
			key := fmt.Sprintf("%s_%d", inst, round)
			if raw, ok := values[key]; ok {
				switch raw {
				case "true":
					return true, true
				case "false":
					return false, true
				}
			}
		}
	case *ast.Boolean:
		return e.Value, true
	}
	return false, false
}

// evalNumericExpr evaluates an expression to a float64.
func evalNumericExpr(expr ast.Expression, round int16, values map[string]string) (float64, bool) {
	switch e := expr.(type) {
	case *ast.AssertVar:
		for _, inst := range e.Instances {
			key := fmt.Sprintf("%s_%d", inst, round)
			if raw, ok := values[key]; ok {
				return parseModelFloat(raw)
			}
		}
		return 0, false

	case *ast.FloatLiteral:
		return e.Value, true

	case *ast.IntegerLiteral:
		return float64(e.Value), true

	case *ast.Boolean:
		if e.Value {
			return 1, true
		}
		return 0, true

	case *ast.InfixExpression:
		lv, lok := evalNumericExpr(e.Left, round, values)
		rv, rok := evalNumericExpr(e.Right, round, values)
		if !lok || !rok {
			return 0, false
		}
		switch e.Operator {
		case "+":
			return lv + rv, true
		case "-":
			return lv - rv, true
		case "*":
			return lv * rv, true
		case "/":
			if rv == 0 {
				return 0, false
			}
			return lv / rv, true
		}
		// Comparison operators inside a sub-expression return 0 or 1.
		if applyCompare(e.Operator, lv, rv) {
			return 1, true
		}
		return 0, true

	case *ast.PrefixExpression:
		rv, ok := evalNumericExpr(e.Right, round, values)
		if !ok {
			return 0, false
		}
		switch e.Operator {
		case "-":
			return -rv, true
		case "!":
			if rv == 0 {
				return 1, true
			}
			return 0, true
		}
	}
	return 0, false
}

// roundsForConstraint returns all round indices present in the model for the
// variables referenced by the constraint.
func roundsForConstraint(c *ast.InvariantClause, values map[string]string) []int16 {
	vars := collectAssertVars(c.Left)
	vars = append(vars, collectAssertVars(c.Right)...)

	seen := make(map[int16]bool)
	var rounds []int16
	for _, av := range vars {
		for _, inst := range av.Instances {
			for key := range values {
				roundStr, base := splitIdent(key)
				if base != inst {
					continue
				}
				r, err := strconv.ParseInt(roundStr, 10, 16)
				if err != nil {
					continue
				}
				k := int16(r)
				if !seen[k] {
					seen[k] = true
					rounds = append(rounds, k)
				}
			}
		}
	}
	return rounds
}

// collectAssertVars recursively collects all AssertVar nodes in an expression.
func collectAssertVars(expr ast.Expression) []*ast.AssertVar {
	switch e := expr.(type) {
	case *ast.AssertVar:
		return []*ast.AssertVar{e}
	case *ast.InfixExpression:
		return append(collectAssertVars(e.Left), collectAssertVars(e.Right)...)
	case *ast.PrefixExpression:
		return collectAssertVars(e.Right)
	case *ast.IndexExpression:
		return collectAssertVars(e.Left)
	}
	return nil
}

func parseModelFloat(s string) (float64, bool) {
	switch s {
	case "true":
		return 1, true
	case "false":
		return 0, true
	}
	v, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0, false
	}
	return v, true
}

func applyCompare(op string, l, r float64) bool {
	switch op {
	case ">":
		return l > r
	case "<":
		return l < r
	case ">=":
		return l >= r
	case "<=":
		return l <= r
	case "==":
		return l == r
	case "!=":
		return l != r
	}
	return false
}
