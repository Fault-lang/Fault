package execute

import (
	"fault/ast"
	"testing"
)

// makeAssertVar builds an AssertVar with one instance name.
func makeAssertVar(instance string) *ast.AssertVar {
	return &ast.AssertVar{
		Instances: []string{instance},
	}
}

// makeFloat builds a FloatLiteral.
func makeFloat(v float64) *ast.FloatLiteral {
	return &ast.FloatLiteral{Value: v}
}

// makeAssertion builds a simple AssertionStatement with the given stored
// (already-negated) operator and temporal type.
func makeAssertion(left ast.Expression, op string, right ast.Expression, temporal string) *ast.AssertionStatement {
	return &ast.AssertionStatement{
		Constraint: &ast.InvariantClause{
			Left:     left,
			Operator: op,
			Right:    right,
		},
		Temporal: temporal,
	}
}

// TestViolationSimpleHolds: stored condition `X <= 0` holds → original `X > 0` violated.
func TestViolationSimpleHolds(t *testing.T) {
	a := makeAssertion(makeAssertVar("spec1_x"), "<=", makeFloat(0), "")
	values := map[string]string{
		"spec1_x_0": "-1.0",
		"spec1_x_1": "2.0",
	}
	mc := &ModelChecker{ResultValues: values}
	mc.EvaluateViolations([]*ast.AssertionStatement{a})
	if !a.Violated {
		t.Fatal("expected Violated=true when stored condition holds in round 0")
	}
}

// TestViolationSimpleDoesNotHold: stored condition `X <= 0` never holds → not violated.
func TestViolationSimpleDoesNotHold(t *testing.T) {
	a := makeAssertion(makeAssertVar("spec1_x"), "<=", makeFloat(0), "")
	values := map[string]string{
		"spec1_x_0": "1.0",
		"spec1_x_1": "2.0",
	}
	mc := &ModelChecker{ResultValues: values}
	mc.EvaluateViolations([]*ast.AssertionStatement{a})
	if a.Violated {
		t.Fatal("expected Violated=false when stored condition never holds")
	}
}

// TestViolationAlwaysAnyRound: temporal "always" → violated if stored condition holds in any round.
func TestViolationAlwaysAnyRound(t *testing.T) {
	a := makeAssertion(makeAssertVar("spec1_x"), "<=", makeFloat(0), "always")
	// Only round 2 violates the original assertion.
	values := map[string]string{
		"spec1_x_0": "5.0",
		"spec1_x_1": "3.0",
		"spec1_x_2": "-1.0",
	}
	mc := &ModelChecker{ResultValues: values}
	mc.EvaluateViolations([]*ast.AssertionStatement{a})
	if !a.Violated {
		t.Fatal("expected Violated=true when stored condition holds in at least one round")
	}
}

// TestViolationEventuallyAllRounds: temporal "eventually" → violated only if stored condition holds in ALL rounds.
func TestViolationEventuallyAllRounds(t *testing.T) {
	// Original: assert eventually X > 0. Stored: X <= 0. Violated if X <= 0 in ALL rounds.
	a := makeAssertion(makeAssertVar("spec1_x"), "<=", makeFloat(0), "eventually")

	// Not all rounds fail → not violated.
	values := map[string]string{
		"spec1_x_0": "-1.0", // stored cond holds
		"spec1_x_1": "3.0",  // stored cond doesn't hold
	}
	mc := &ModelChecker{ResultValues: values}
	mc.EvaluateViolations([]*ast.AssertionStatement{a})
	if a.Violated {
		t.Fatal("expected Violated=false when stored condition doesn't hold in all rounds")
	}

	// All rounds fail → violated.
	a.Violated = false
	values = map[string]string{
		"spec1_x_0": "-1.0",
		"spec1_x_1": "-2.0",
	}
	mc.ResultValues = values
	mc.EvaluateViolations([]*ast.AssertionStatement{a})
	if !a.Violated {
		t.Fatal("expected Violated=true when stored condition holds in all rounds")
	}
}

// TestViolationNoMatchingVariable: variable not in model → not violated.
func TestViolationNoMatchingVariable(t *testing.T) {
	a := makeAssertion(makeAssertVar("spec1_x"), "<=", makeFloat(0), "")
	values := map[string]string{
		"spec1_y_0": "5.0",
	}
	mc := &ModelChecker{ResultValues: values}
	mc.EvaluateViolations([]*ast.AssertionStatement{a})
	if a.Violated {
		t.Fatal("expected Violated=false when variable is not in model")
	}
}

// TestViolationEquality: stored condition `X != true` (original: assert X == true).
func TestViolationEquality(t *testing.T) {
	// Use float representation: true=1, false=0.
	a := makeAssertion(makeAssertVar("spec1_flag"), "!=", makeFloat(1), "")
	values := map[string]string{
		"spec1_flag_0": "false", // 0 != 1 → stored cond holds → violated
	}
	mc := &ModelChecker{ResultValues: values}
	mc.EvaluateViolations([]*ast.AssertionStatement{a})
	if !a.Violated {
		t.Fatal("expected Violated=true when flag is false and stored condition is !=1")
	}
}

// TestViolationMultipleAssertions: only the matching assertion is marked violated.
func TestViolationMultipleAssertions(t *testing.T) {
	a1 := makeAssertion(makeAssertVar("spec1_x"), "<=", makeFloat(0), "")  // violated
	a2 := makeAssertion(makeAssertVar("spec1_y"), "<=", makeFloat(0), "")  // not violated

	values := map[string]string{
		"spec1_x_0": "-1.0",
		"spec1_y_0": "5.0",
	}
	mc := &ModelChecker{ResultValues: values}
	mc.EvaluateViolations([]*ast.AssertionStatement{a1, a2})

	if !a1.Violated {
		t.Fatal("expected a1.Violated=true")
	}
	if a2.Violated {
		t.Fatal("expected a2.Violated=false")
	}
}

// TestViolationTemporalFilterSkipped: assertions with a TemporalFilter (nft/nmt)
// are not yet supported and must return false without panicking.
func TestViolationTemporalFilterSkipped(t *testing.T) {
	for _, filter := range []string{"nft", "nmt"} {
		a := &ast.AssertionStatement{
			Constraint: &ast.InvariantClause{
				Left:     makeAssertVar("spec1_x"),
				Operator: "<=",
				Right:    makeFloat(0),
			},
			TemporalFilter: filter,
		}
		values := map[string]string{
			"spec1_x_0": "-1.0", // would violate if evaluated
		}
		mc := &ModelChecker{ResultValues: values}
		mc.EvaluateViolations([]*ast.AssertionStatement{a})
		if a.Violated {
			t.Errorf("TemporalFilter=%q: expected Violated=false (unsupported filter), got true", filter)
		}
	}
}

// TestViolationZeroValueNotMissing: a variable that is zero (or false) is present
// in the model and should be evaluated, not treated as missing.
func TestViolationZeroValueNotMissing(t *testing.T) {
	// Stored condition: X <= 0. X=0 satisfies it → violated.
	a := makeAssertion(makeAssertVar("spec1_x"), "<=", makeFloat(0), "")
	values := map[string]string{
		"spec1_x_0": "0.0",
	}
	mc := &ModelChecker{ResultValues: values}
	mc.EvaluateViolations([]*ast.AssertionStatement{a})
	if !a.Violated {
		t.Fatal("expected Violated=true for zero-value variable satisfying stored condition")
	}
}

// TestViolationArithmetic: stored condition X + Y <= 0 with X=1, Y=-2 → holds → violated.
func TestViolationArithmetic(t *testing.T) {
	a := &ast.AssertionStatement{
		Constraint: &ast.InvariantClause{
			Left: &ast.InfixExpression{
				Left:     makeAssertVar("spec1_x"),
				Operator: "+",
				Right:    makeAssertVar("spec1_y"),
			},
			Operator: "<=",
			Right:    makeFloat(0),
		},
	}
	values := map[string]string{
		"spec1_x_0": "1.0",
		"spec1_y_0": "-2.0",
	}
	mc := &ModelChecker{ResultValues: values}
	mc.EvaluateViolations([]*ast.AssertionStatement{a})
	if !a.Violated {
		t.Fatal("expected Violated=true when X+Y=-1 satisfies stored condition <=0")
	}
}

// TestViolationPartialRoundsAlways: "always" with a variable present in only some
// rounds — held only in the round where it exists → violated.
func TestViolationPartialRoundsAlways(t *testing.T) {
	a := makeAssertion(makeAssertVar("spec1_x"), "<=", makeFloat(0), "always")
	// Only round 0 exists for x; condition holds.
	values := map[string]string{
		"spec1_x_0": "-1.0",
	}
	mc := &ModelChecker{ResultValues: values}
	mc.EvaluateViolations([]*ast.AssertionStatement{a})
	if !a.Violated {
		t.Fatal("expected Violated=true when condition holds in the only round present")
	}
}

// TestViolationDivisionByZeroSafe: stored condition X / Y <= 0 with Y=0 must not panic.
func TestViolationDivisionByZeroSafe(t *testing.T) {
	a := &ast.AssertionStatement{
		Constraint: &ast.InvariantClause{
			Left: &ast.InfixExpression{
				Left:     makeAssertVar("spec1_x"),
				Operator: "/",
				Right:    makeAssertVar("spec1_y"),
			},
			Operator: "<=",
			Right:    makeFloat(0),
		},
	}
	values := map[string]string{
		"spec1_x_0": "5.0",
		"spec1_y_0": "0.0",
	}
	mc := &ModelChecker{ResultValues: values}
	// Must not panic; result is indeterminate (ok=false) so not violated.
	mc.EvaluateViolations([]*ast.AssertionStatement{a})
	if a.Violated {
		t.Fatal("expected Violated=false when division by zero yields indeterminate result")
	}
}
