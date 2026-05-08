package asserts

import (
	"fault/ast"
	"fmt"
	"sort"
	"strings"
	"testing"
)

// registry builds a minimal SSA registry for a single variable.
// e.g. registry("spec_s_value", 3) gives round-0_@__run → [["spec_s_value","0"],["spec_s_value","1"],...]
// The key must match the "round-N_@__run" pattern that VarSets.GetByRunRound expects.
func registry(varName string, ssaCount int) map[string][][]string {
	vars := make([][]string, ssaCount+1)
	for i := 0; i <= ssaCount; i++ {
		vars[i] = []string{varName, fmt.Sprintf("%d", i)}
	}
	return map[string][][]string{"round-0_@__run": vars}
}

// assertStmt builds an AssertionStatement for a simple (left op right) constraint.
func assertStmt(left []string, op string, right float64, assume bool, temporal, filter string) *ast.AssertionStatement {
	return &ast.AssertionStatement{
		Token: ast.Token{},
		Constraint: &ast.InvariantClause{
			Left:     &ast.AssertVar{Instances: left},
			Right:    &ast.FloatLiteral{Value: right},
			Operator: op,
		},
		Assume:         assume,
		Temporal:       temporal,
		TemporalFilter: filter,
	}
}

// sortedSMT normalises an SMT OR/AND clause for deterministic comparison,
// since StringSet iteration order is random.
func sortedSMT(s string) string {
	s = strings.TrimSpace(s)
	// Strip outer wrapper (assert ...) and outer paren group
	inner := s
	if strings.HasPrefix(inner, "(assert ") {
		inner = inner[len("(assert "):]
		inner = inner[:len(inner)-1]
	}
	// Get the operator and terms
	inner = strings.TrimSpace(inner)
	if len(inner) < 2 || inner[0] != '(' {
		return s
	}
	inner = inner[1 : len(inner)-1]
	parts := strings.SplitN(inner, " ", 2)
	if len(parts) != 2 {
		return s
	}
	op := parts[0]
	rest := parts[1]
	// Split terms (simple split on ") (" boundary)
	terms := splitTerms(rest)
	sort.Strings(terms)
	return fmt.Sprintf("(assert (%s %s))", op, strings.Join(terms, " "))
}

func splitTerms(s string) []string {
	var terms []string
	depth := 0
	start := 0
	for i, c := range s {
		switch c {
		case '(':
			depth++
		case ')':
			depth--
			if depth == 0 {
				terms = append(terms, strings.TrimSpace(s[start:i+1]))
				start = i + 1
			}
		case ' ':
			if depth == 0 && strings.TrimSpace(s[start:i]) != "" {
				terms = append(terms, strings.TrimSpace(s[start:i]))
				start = i + 1
			}
		}
	}
	if t := strings.TrimSpace(s[start:]); t != "" {
		terms = append(terms, t)
	}
	return terms
}

// ---- smtlibOperators ----

func TestSmtlibOperators(t *testing.T) {
	cases := []struct {
		input string
		want  string
	}{
		{"==", "="},
		{"!=", "not"},
		{"||", "or"},
		{"&&", "and"},
		{"<", "<"},
		{">", ">"},
		{"<=", "<="},
		{">=", ">="},
	}
	for _, tc := range cases {
		got := smtlibOperators(tc.input)
		if got != tc.want {
			t.Errorf("smtlibOperators(%q) = %q, want %q", tc.input, got, tc.want)
		}
	}
}

// ---- NewConstraint polarity ----

func TestNewConstraint_AssertPolarity(t *testing.T) {
	stmt := assertStmt([]string{"spec_s_v"}, "==", 10, false, "", "")
	c := NewConstraint(stmt, 1, registry("spec_s_v", 2), map[string][]map[string]string{})
	if c.On != "and" {
		t.Errorf("assert: On = %q, want %q", c.On, "and")
	}
	if c.Off != "or" {
		t.Errorf("assert: Off = %q, want %q", c.Off, "or")
	}
	if c.Assume {
		t.Error("assert: Assume should be false")
	}
}

func TestNewConstraint_AssumePolarity(t *testing.T) {
	stmt := assertStmt([]string{"spec_s_v"}, "==", 10, true, "", "")
	c := NewConstraint(stmt, 1, registry("spec_s_v", 2), map[string][]map[string]string{})
	if c.On != "or" {
		t.Errorf("assume: On = %q, want %q", c.On, "or")
	}
	if c.Off != "and" {
		t.Errorf("assume: Off = %q, want %q", c.Off, "and")
	}
	if !c.Assume {
		t.Error("assume: Assume should be true")
	}
}

func TestNewConstraint_OperatorMapped(t *testing.T) {
	stmt := assertStmt([]string{"spec_s_v"}, "==", 10, false, "", "")
	c := NewConstraint(stmt, 1, registry("spec_s_v", 2), map[string][]map[string]string{})
	if c.Op != "=" {
		t.Errorf("Op = %q, want %q", c.Op, "=")
	}
}

// ---- Package ----

func TestPackage_EqualityOperator(t *testing.T) {
	c := &Constraint{}
	result := c.Package([][]string{{"x_0", "10"}, {"x_1", "10"}}, "=")
	vals := result.Values()
	for _, v := range vals {
		if !strings.HasPrefix(v, "(= ") {
			t.Errorf("Package with = produced %q, want (= ...)", v)
		}
	}
}

func TestPackage_NotEqualOperator(t *testing.T) {
	c := &Constraint{}
	result := c.Package([][]string{{"x_0", "10"}}, "not")
	vals := result.Values()
	if len(vals) != 1 {
		t.Fatalf("expected 1 value, got %d", len(vals))
	}
	if vals[0] != "(not (= x_0 10))" {
		t.Errorf("got %q, want %q", vals[0], "(not (= x_0 10))")
	}
}

func TestPackage_NotEqualWithFalseLHS(t *testing.T) {
	// When left side is "false", not wraps the right side
	c := &Constraint{}
	result := c.Package([][]string{{"false", "x_0"}}, "not")
	vals := result.Values()
	if len(vals) != 1 {
		t.Fatalf("expected 1 value, got %d", len(vals))
	}
	if vals[0] != "(not x_0)" {
		t.Errorf("got %q, want %q", vals[0], "(not x_0)")
	}
}

func TestPackage_ComparisonOperator(t *testing.T) {
	c := &Constraint{}
	result := c.Package([][]string{{"x_0", "5"}}, ">")
	vals := result.Values()
	if len(vals) != 1 || vals[0] != "(> x_0 5)" {
		t.Errorf("got %v, want [(> x_0 5)]", vals)
	}
}

// ---- captureState ----

func TestCaptureState_AllVersions(t *testing.T) {
	// Three-part name with no trailing number → apply to all SSA versions
	_, all, constant := captureState("spec_s_value")
	if !all {
		t.Error("expected all=true for multi-part name without numeric suffix")
	}
	if constant {
		t.Error("expected constant=false for non-constant")
	}
}

func TestCaptureState_Constant(t *testing.T) {
	// Two-part name → constant
	_, all, constant := captureState("myconst")
	if all {
		t.Error("expected all=false for constant")
	}
	if !constant {
		t.Error("expected constant=true for two-part name")
	}
}

func TestCaptureState_SpecificSSA(t *testing.T) {
	// Trailing numeric suffix → specific SSA version only
	idx, all, constant := captureState("spec_s_value_3")
	if all || constant {
		t.Errorf("expected all=false, constant=false for specific SSA; got all=%v constant=%v", all, constant)
	}
	if idx != "3" {
		t.Errorf("expected idx=3, got %q", idx)
	}
}

// ---- HasActiveInstance / IsRelevant ----

func TestHasActiveInstance_Empty(t *testing.T) {
	if HasActiveInstance([]string{}, map[string]string{"x": "Real"}) {
		t.Error("empty instances should return false")
	}
}

func TestHasActiveInstance_MultipleInstances(t *testing.T) {
	// More than one instance → always active regardless of vars
	if !HasActiveInstance([]string{"a", "b"}, map[string]string{}) {
		t.Error("multiple instances should always return true")
	}
}

func TestHasActiveInstance_SinglePresent(t *testing.T) {
	if !HasActiveInstance([]string{"spec_s_v"}, map[string]string{"spec_s_v": "Real"}) {
		t.Error("single instance present in vars should return true")
	}
}

func TestHasActiveInstance_SingleAbsent(t *testing.T) {
	if HasActiveInstance([]string{"spec_s_v"}, map[string]string{"other": "Real"}) {
		t.Error("single instance absent from vars should return false")
	}
}

// ---- eventuallyAlways ----

func TestEventuallyAlways_Single(t *testing.T) {
	c := &Constraint{}
	result := c.eventuallyAlways([]string{"(= x_0 5)"})
	if result != "(or (= x_0 5))" {
		t.Errorf("got %q, want %q", result, "(or (= x_0 5))")
	}
}

func TestEventuallyAlways_Multiple(t *testing.T) {
	c := &Constraint{}
	result := c.eventuallyAlways([]string{"(= x_0 5)", "(= x_1 5)", "(= x_2 5)"})
	// Should be: (or (= x_2 5) (and (= x_1 5) (= x_2 5)) (and (= x_0 5) (= x_1 5) (= x_2 5)))
	if !strings.HasPrefix(result, "(or ") {
		t.Errorf("expected or clause, got %q", result)
	}
	// The last term (all three) must be present
	if !strings.Contains(result, "(and (= x_0 5) (= x_1 5) (= x_2 5))") {
		t.Errorf("missing full conjunction in eventually-always: %q", result)
	}
}

// ---- NoFew / NoMore ----

func TestNoFew_N1(t *testing.T) {
	c := &Constraint{}
	terms := []string{"(> x_0 5)", "(> x_1 5)", "(> x_2 5)"}
	result := c.NoFew(terms, 1)
	// n=1 → return all terms unchanged
	if len(result) != 3 {
		t.Errorf("NoFew n=1: got %d results, want 3", len(result))
	}
}

func TestNoFew_N2(t *testing.T) {
	c := &Constraint{}
	terms := []string{"(> x_0 5)", "(> x_1 5)", "(> x_2 5)"}
	result := c.NoFew(terms, 2)
	// C(3,2)=3 pairs, each wrapped in (and ...)
	if len(result) != 3 {
		t.Errorf("NoFew n=2 of 3: got %d results, want 3", len(result))
	}
	for _, r := range result {
		if !strings.HasPrefix(r, "(and ") {
			t.Errorf("NoFew n=2: expected (and ...), got %q", r)
		}
	}
}

func TestNoMore_N1(t *testing.T) {
	c := &Constraint{}
	terms := []string{"(= x_0 5)", "(= x_1 5)", "(= x_2 5)"}
	result := c.NoMore(terms, 1)
	// n=1 → return terms unchanged
	if len(result) != 3 {
		t.Errorf("NoMore n=1: got %d results, want 3", len(result))
	}
}

func TestNoMore_N2(t *testing.T) {
	c := &Constraint{}
	terms := []string{"(= x_0 5)", "(= x_1 5)", "(= x_2 5)"}
	result := c.NoMore(terms, 2)
	// C(3,2)=3 combinations, each (and (or chosen) (and nots))
	if len(result) != 3 {
		t.Errorf("NoMore n=2 of 3: got %d results, want 3", len(result))
	}
}

// ---- Parse(): full constraint → SMT string ----

func TestParse_AssumeEventually_Equality(t *testing.T) {
	// assume spec_s_value == 60 eventually
	// → (assert (or (= spec_s_value_0 60) (= spec_s_value_1 60) (= spec_s_value_2 60)))
	stmt := assertStmt([]string{"spec_s_value"}, "==", 60, true, "eventually", "")
	reg := registry("spec_s_value", 2)
	c := NewConstraint(stmt, 1, reg, map[string][]map[string]string{})
	results := c.Parse()
	if len(results) != 1 {
		t.Fatalf("expected 1 assertion, got %d", len(results))
	}
	got := results[0]
	if !strings.HasPrefix(got, "(assert (or ") {
		t.Errorf("assume eventually should produce (assert (or ...)), got: %s", got)
	}
	for i := 0; i <= 2; i++ {
		term := fmt.Sprintf("(= spec_s_value_%d 60)", i)
		if !strings.Contains(got, term) {
			t.Errorf("missing term %q in: %s", term, got)
		}
	}
}

func TestParse_AssumeAlways_Equality(t *testing.T) {
	// assume spec_s_value == 60 always
	// → (assert (and (= spec_s_value_0 60) (= spec_s_value_1 60) ...))
	stmt := assertStmt([]string{"spec_s_value"}, "==", 60, true, "always", "")
	reg := registry("spec_s_value", 2)
	c := NewConstraint(stmt, 1, reg, map[string][]map[string]string{})
	results := c.Parse()
	if len(results) != 1 {
		t.Fatalf("expected 1 assertion, got %d", len(results))
	}
	got := results[0]
	if !strings.HasPrefix(got, "(assert (and ") {
		t.Errorf("assume always should produce (assert (and ...)), got: %s", got)
	}
}

func TestParse_AssumeEventually_GreaterThan(t *testing.T) {
	// assume counter.value > 10 eventually (synthesis goal)
	// → (assert (or (> spec_c_value_0 10) (> spec_c_value_1 10)))
	stmt := assertStmt([]string{"spec_c_value"}, ">", 10, true, "eventually", "")
	reg := registry("spec_c_value", 1)
	c := NewConstraint(stmt, 1, reg, map[string][]map[string]string{})
	results := c.Parse()
	if len(results) != 1 {
		t.Fatalf("expected 1 assertion, got %d", len(results))
	}
	got := results[0]
	if !strings.Contains(got, "(> spec_c_value_0 10)") {
		t.Errorf("missing (> spec_c_value_0 10) in: %s", got)
	}
	if !strings.Contains(got, "(> spec_c_value_1 10)") {
		t.Errorf("missing (> spec_c_value_1 10) in: %s", got)
	}
}

func TestParse_AssertNoTemporal_NegatedEquality(t *testing.T) {
	// assert (after compiler negation: != operator)
	// The compiler mutates == to != before calling the generator.
	// → (assert (or (not (= spec_s_value_0 60)) (not (= spec_s_value_1 60)) ...))
	stmt := assertStmt([]string{"spec_s_value"}, "!=", 60, false, "", "")
	reg := registry("spec_s_value", 2)
	c := NewConstraint(stmt, 1, reg, map[string][]map[string]string{})
	results := c.Parse()
	if len(results) != 1 {
		t.Fatalf("expected 1 assertion, got %d", len(results))
	}
	got := results[0]
	if !strings.HasPrefix(got, "(assert (or ") {
		t.Errorf("assert != should produce (assert (or ...)), got: %s", got)
	}
	if !strings.Contains(got, "(not (= ") {
		t.Errorf("assert != should contain (not (= ...)), got: %s", got)
	}
}

func TestParse_AssumeNoTemporal_Equality(t *testing.T) {
	// assume spec_s_value == 60 (no temporal)
	// → (assert (and (= spec_s_value_0 60) (= spec_s_value_1 60) ...))
	// Off="and" for assume, no temporal → uses Off
	stmt := assertStmt([]string{"spec_s_value"}, "==", 60, true, "", "")
	reg := registry("spec_s_value", 2)
	c := NewConstraint(stmt, 1, reg, map[string][]map[string]string{})
	results := c.Parse()
	if len(results) != 1 {
		t.Fatalf("expected 1 assertion, got %d", len(results))
	}
	got := results[0]
	if !strings.HasPrefix(got, "(assert (and ") {
		t.Errorf("assume no temporal should produce (assert (and ...)), got: %s", got)
	}
}

func TestParse_AssumeEventuallyAlways(t *testing.T) {
	// assume spec_s_value == 5 eventually-always
	// → (assert (or v2 (and v1 v2) (and v0 v1 v2)))
	stmt := assertStmt([]string{"spec_s_value"}, "==", 5, true, "eventually-always", "")
	reg := registry("spec_s_value", 2)
	c := NewConstraint(stmt, 1, reg, map[string][]map[string]string{})
	results := c.Parse()
	if len(results) != 1 {
		t.Fatalf("expected 1 assertion, got %d", len(results))
	}
	got := results[0]
	if !strings.HasPrefix(got, "(assert (or ") {
		t.Errorf("eventually-always should produce (assert (or ...)), got: %s", got)
	}
	// Must contain the "all versions true" conjunction
	if !strings.Contains(got, "(and ") {
		t.Errorf("eventually-always should contain conjunctions, got: %s", got)
	}
}

func TestParse_MultipleInstances(t *testing.T) {
	// When both template and instance variables exist (e.g. tank_level and inst_t_level)
	// both should appear in the assertion
	reg := map[string][][]string{
		"round-0_@__run": {
			{"spec_tank_level", "0"},
			{"spec_tank_level", "1"},
			{"spec_inst_t_level", "0"},
			{"spec_inst_t_level", "1"},
		},
	}
	stmt := &ast.AssertionStatement{
		Constraint: &ast.InvariantClause{
			Left:     &ast.AssertVar{Instances: []string{"spec_tank_level", "spec_inst_t_level"}},
			Right:    &ast.FloatLiteral{Value: 60},
			Operator: "==",
		},
		Assume:   true,
		Temporal: "eventually",
	}
	c := NewConstraint(stmt, 1, reg, map[string][]map[string]string{})
	results := c.Parse()
	if len(results) != 1 {
		t.Fatalf("expected 1 assertion, got %d", len(results))
	}
	got := results[0]
	if !strings.Contains(got, "spec_tank_level") {
		t.Errorf("missing spec_tank_level in: %s", got)
	}
	if !strings.Contains(got, "spec_inst_t_level") {
		t.Errorf("missing spec_inst_t_level in: %s", got)
	}
}
