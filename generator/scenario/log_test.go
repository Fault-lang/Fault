package scenario

import (
	"fault/ast"
	"strings"
	"testing"
)

// ---- helpers ----

func makeAssertVar(instance string) *ast.AssertVar {
	return &ast.AssertVar{Instances: []string{instance}}
}

func makeViolatedAssertion(varName string) *ast.AssertionStatement {
	return &ast.AssertionStatement{
		Constraint: &ast.InvariantClause{
			Left:     makeAssertVar(varName),
			Operator: "!=",
			Right:    &ast.FloatLiteral{Value: 1},
		},
		Violated: true,
	}
}

func TestNewLogger(t *testing.T) {
	logger := NewLogger()
	if logger == nil {
		t.Errorf("NewLogger() = %v, want non-nil logger", logger)
	}
	if logger.Events == nil {
		t.Errorf("NewLogger().Events = %v, want non-nil Events", logger.Events)
	}
	if logger.BranchIndexes == nil {
		t.Errorf("NewLogger().BranchIndexes = %v, want non-nil BranchIndexes", logger.BranchIndexes)
	}
	if logger.BranchVars == nil {
		t.Errorf("NewLogger().BranchVars = %v, want non-nil BranchVars", logger.BranchVars)
	}
	if logger.ForksCaps == nil {
		t.Errorf("NewLogger().ForksCaps = %v, want non-nil ForksCaps", logger.ForksCaps)
	}
	if logger.Results == nil {
		t.Errorf("NewLogger().Results = %v, want non-nil Results", logger.Results)
	}
}

func TestLogger_EnterFunction(t *testing.T) {
	logger := NewLogger()
	logger.EnterFunction("foo", 1)
	if len(logger.Events) != 1 {
		t.Errorf("Logger.EnterFunction() = %v, want %v", len(logger.Events), 1)
	}
	if logger.Events[0].(*FunctionCall).FunctionName != "foo" {
		t.Errorf("Logger.EnterFunction().FunctionName = %v, want %v", logger.Events[0].(*FunctionCall).FunctionName, "foo")
	}
	if logger.Events[0].(*FunctionCall).Round != "1" {
		t.Errorf("Logger.EnterFunction().Round = %v, want %v", logger.Events[0].(*FunctionCall).Round, 1)
	}
	if logger.Events[0].(*FunctionCall).Type != "Entry" {
		t.Errorf("Logger.EnterFunction().Type = %v, want %v", logger.Events[0].(*FunctionCall).Type, "Entry")
	}
}

func TestLogger_ExitFunction(t *testing.T) {
	logger := NewLogger()
	logger.ExitFunction("foo", 1)
	if len(logger.Events) != 1 {
		t.Errorf("Logger.ExitFunction() = %v, want %v", len(logger.Events), 1)
	}
	if logger.Events[0].(*FunctionCall).FunctionName != "foo" {
		t.Errorf("Logger.ExitFunction().FunctionName = %v, want %v", logger.Events[0].(*FunctionCall).FunctionName, "foo")
	}
	if logger.Events[0].(*FunctionCall).Round != "1" {
		t.Errorf("Logger.ExitFunction().Round = %v, want %v", logger.Events[0].(*FunctionCall).Round, 1)
	}
	if logger.Events[0].(*FunctionCall).Type != "Exit" {
		t.Errorf("Logger.ExitFunction().Type = %v, want %v", logger.Events[0].(*FunctionCall).Type, "Exit")
	}
}

func TestLogger_UpdateVariable(t *testing.T) {
	logger := NewLogger()
	logger.UpdateVariable("foo", false)
	if len(logger.Events) != 1 {
		t.Errorf("Logger.UpdateVariable() = %v, want %v", len(logger.Events), 1)
	}
	if logger.Events[0].(*VariableUpdate).Variable != "foo" {
		t.Errorf("Logger.UpdateVariable().Variable = %v, want %v", logger.Events[0].(*VariableUpdate).Variable, "foo")
	}
}

func TestLogger_AddPhiOption(t *testing.T) {
	logger := NewLogger()
	logger.AddPhiOption("foo", "bar")
	if len(logger.ForksCaps["foo"]) != 1 {
		t.Errorf("Logger.AddPhiOption() = %v, want %v", len(logger.ForksCaps["foo"]), 1)
	}
	if logger.ForksCaps["foo"][0] != "bar" {
		t.Errorf("Logger.AddPhiOption() = %v, want %v", logger.ForksCaps["foo"][0], "bar")
	}
}

func TestFunctionCall_MarkDead(t *testing.T) {
	logger := NewLogger()
	b1 := logger.NewBranchSelector("b", int(1), []string{}, []string{"a_1"})
	b2 := logger.NewBranchSelector("b", int(2), []string{}, []string{"a_2"})
	logger.BranchSelectors = []*BranchSelector{b1, b2}
	logger.EnterFunction("test1", 1)
	logger.UpdateVariable("a_1", false)
	logger.ExitFunction("test1", 1)
	logger.EnterFunction("test2", 1)
	logger.UpdateVariable("a_2", false)
	logger.ExitFunction("test2", 1)
	logger.AddPhiOption("a_3", "a_1")
	logger.AddPhiOption("a_3", "a_2")

	logger.Results["a_1"] = "1"
	logger.Results["a_2"] = "5"
	logger.Results["a_3"] = "1"
	logger.Results["b_1"] = "true"
	logger.Results["b_2"] = "false"

	logger.Trace()
	logger.Kill()

	if logger.Events[1].IsDead() {
		t.Errorf("FunctionCall.IsDead() = %v, want %v", logger.Events[1].IsDead(), false)
	}
	if !logger.Events[4].IsDead() {
		t.Errorf("FunctionCall.IsDead() = %v, want %v", logger.Events[4].IsDead(), true)
	}
}

// ---- isSynthSlotName ----

func TestIsSynthSlotName_Valid(t *testing.T) {
	cases := []string{"synth_0", "synth_1", "synth_42"}
	for _, c := range cases {
		if !isSynthSlotName(c) {
			t.Errorf("isSynthSlotName(%q) = false, want true", c)
		}
	}
}

func TestIsSynthSlotName_Invalid(t *testing.T) {
	cases := []string{
		"synth_",       // empty suffix
		"synth_abc",    // non-numeric suffix
		"synth_1a",     // mixed
		"synth_1_fill", // has extra segment
		"other_1",      // wrong prefix
		"",             // empty
		"synth",        // no underscore
	}
	for _, c := range cases {
		if isSynthSlotName(c) {
			t.Errorf("isSynthSlotName(%q) = true, want false", c)
		}
	}
}

// ---- synthChoice ----

func TestSynthChoice_FoundTrue(t *testing.T) {
	l := NewLogger()
	l.Results["synth_1_fill_1"] = "true"
	l.Results["synth_1_drain_1"] = "false"
	got := l.synthChoice("synth_1")
	if got != "fill" {
		t.Errorf("synthChoice = %q, want %q", got, "fill")
	}
}

func TestSynthChoice_NoneTrue(t *testing.T) {
	l := NewLogger()
	l.Results["synth_1_fill_1"] = "false"
	l.Results["synth_1_drain_1"] = "false"
	got := l.synthChoice("synth_1")
	if got != "" {
		t.Errorf("synthChoice with all false = %q, want %q", got, "")
	}
}

func TestSynthChoice_NoResults(t *testing.T) {
	l := NewLogger()
	got := l.synthChoice("synth_1")
	if got != "" {
		t.Errorf("synthChoice with empty results = %q, want %q", got, "")
	}
}

func TestSynthChoice_WrongSlot(t *testing.T) {
	l := NewLogger()
	l.Results["synth_2_fill_2"] = "true"
	got := l.synthChoice("synth_1") // asking about slot 1, not 2
	if got != "" {
		t.Errorf("synthChoice wrong slot = %q, want %q", got, "")
	}
}

// ---- IsInternalVariable ----

func TestIsInternalVariable_BlockSelectors(t *testing.T) {
	l := NewLogger()
	// block*true_N and block*false_N are internal
	if !l.IsInternalVariable("blockABCtrue_0") {
		t.Error("blockABCtrue_0 should be internal")
	}
	if !l.IsInternalVariable("blockXYZfalse_1") {
		t.Error("blockXYZfalse_1 should be internal")
	}
}

func TestIsInternalVariable_SynthSelectors(t *testing.T) {
	l := NewLogger()
	if !l.IsInternalVariable("synth_1_fill_1") {
		t.Error("synth_1_fill_1 should be internal")
	}
}

func TestIsInternalVariable_UserVars(t *testing.T) {
	l := NewLogger()
	cases := []string{"spec_c_value_0", "spec_tank_level_2", "myvar_3"}
	for _, c := range cases {
		if l.IsInternalVariable(c) {
			t.Errorf("IsInternalVariable(%q) = true, want false", c)
		}
	}
}

// ---- Kill: synthesis slot protection ----

func TestKill_SynthSlotEntryStaysAlive(t *testing.T) {
	// When the drain candidate's selector is false, drain's variable update is dead.
	// But the outer synth_1 Entry/Exit must remain alive (isSynthSlotName guards it).
	l := NewLogger()

	// Events:
	// [0] Enter synth_1
	// [1]   Enter synth_1_drain  (candidate)
	// [2]     VariableUpdate spec_t_level_1
	// [3]   Exit  synth_1_drain
	// [4] Exit  synth_1
	l.EnterFunction("synth_1", 1)
	l.EnterFunction("synth_1_drain", 1)
	l.UpdateVariable("spec_t_level_1", false)
	l.ExitFunction("synth_1_drain", 1)
	l.ExitFunction("synth_1", 1)

	// synth_1_drain selector is false → drain's vars are dead
	sel := l.NewBranchSelector("synth_1_drain", 1, []string{"(= phi spec_t_level_1)"}, []string{"spec_t_level_1"})
	l.AddBranchSelector(sel)
	l.Results["synth_1_drain_1"] = "false"

	l.Trace()
	l.Kill()

	// The variable update inside drain should be dead
	if !l.Events[2].IsDead() {
		t.Error("spec_t_level_1 update should be dead (in false branch)")
	}

	// The outer synth_1 Entry must NOT be dead
	if l.Events[0].IsDead() {
		t.Error("synth_1 Entry must stay alive (synthesis slot)")
	}
	// The outer synth_1 Exit must NOT be dead
	if l.Events[4].IsDead() {
		t.Error("synth_1 Exit must stay alive (synthesis slot)")
	}
}

func TestKill_DeadBranchKillsFunction(t *testing.T) {
	// A normal function whose selector is false should have Entry/Exit marked dead.
	l := NewLogger()

	// [0] Enter foo
	// [1]   VariableUpdate var_x_0
	// [2] Exit foo
	l.EnterFunction("foo", 0)
	l.UpdateVariable("var_x_0", false)
	l.ExitFunction("foo", 0)

	sel := l.NewBranchSelector("blockABC", 0, []string{"(= phi var_x_0)"}, []string{"var_x_0"})
	l.AddBranchSelector(sel)
	l.Results["blockABC_0"] = "false"

	l.Trace()
	l.Kill()

	if !l.Events[1].IsDead() {
		t.Error("var_x_0 update should be dead")
	}
	if !l.Events[0].IsDead() {
		t.Error("foo Entry should be dead (all vars dead)")
	}
	if !l.Events[2].IsDead() {
		t.Error("foo Exit should be dead (all vars dead)")
	}
}

func TestKill_LiveBranchKeepsFunction(t *testing.T) {
	// A function whose selector is true: nothing should be marked dead.
	l := NewLogger()

	l.EnterFunction("bar", 0)
	l.UpdateVariable("var_y_0", false)
	l.ExitFunction("bar", 0)

	sel := l.NewBranchSelector("blockDEF", 0, []string{"(= phi var_y_0)"}, []string{"var_y_0"})
	l.AddBranchSelector(sel)
	l.Results["blockDEF_0"] = "true"

	l.Trace()
	l.Kill()

	for i, e := range l.Events {
		if e.IsDead() {
			t.Errorf("event[%d] should be alive (true branch), got dead", i)
		}
	}
}

func TestKill_NoBranchSelectors(t *testing.T) {
	// Kill is a no-op when there are no dead selectors.
	l := NewLogger()
	l.EnterFunction("baz", 0)
	l.UpdateVariable("var_z_0", false)
	l.ExitFunction("baz", 0)

	l.Trace()
	l.Kill() // no BranchSelectors, should not panic

	for i, e := range l.Events {
		if e.IsDead() {
			t.Errorf("event[%d] should be alive (no selectors), got dead", i)
		}
	}
}

// ---- synthChoice: multiple true candidates ----

func TestSynthChoice_MultipleTrueCandidates(t *testing.T) {
	// When more than one candidate selector is true the function must return
	// one of the true candidates without panicking.
	l := NewLogger()
	l.Results["synth_1_fill_1"] = "true"
	l.Results["synth_1_drain_1"] = "true"
	got := l.synthChoice("synth_1")
	if got != "fill" && got != "drain" {
		t.Errorf("synthChoice with multiple true candidates = %q, want fill or drain", got)
	}
}

// ---- Kill: empty Results map ----

func TestKill_EmptyResults(t *testing.T) {
	// Kill must not panic when Results is empty.
	l := NewLogger()
	l.EnterFunction("foo", 0)
	l.UpdateVariable("var_x_0", false)
	l.ExitFunction("foo", 0)

	sel := l.NewBranchSelector("blockABC", 0, []string{"(= phi var_x_0)"}, []string{"var_x_0"})
	l.AddBranchSelector(sel)
	// Results is empty — selector value is missing, treated as not-true.

	l.Trace()
	l.Kill() // must not panic
}

// ---- IsNegated / _neg suffix rendering in String() ----

func TestIsNegated_Positive(t *testing.T) {
	l := NewLogger()
	base, negated := l.IsNegated("spec_x_neg")
	if !negated {
		t.Error("expected negated=true for 'spec_x_neg'")
	}
	if base != "spec_x" {
		t.Errorf("expected base %q, got %q", "spec_x", base)
	}
}

func TestIsNegated_Negative(t *testing.T) {
	l := NewLogger()
	base, negated := l.IsNegated("spec_x")
	if negated {
		t.Error("expected negated=false for 'spec_x'")
	}
	if base != "spec_x" {
		t.Errorf("expected base unchanged %q, got %q", "spec_x", base)
	}
}

func TestString_NegatedVariableInOutput(t *testing.T) {
	l := NewLogger()

	l.EnterFunction("@__run", 1)
	l.EnterFunction("foo", 1)
	// Variable name encodes negation via _neg suffix.
	l.UpdateVariable("spec_x_neg_1", false)
	l.ExitFunction("foo", 1)
	l.ExitFunction("@__run", 1)

	l.Results["spec_x_neg_1"] = "true"
	l.StringRules["spec_x"] = "x"
	l.IsStringRule["spec_x"] = true

	l.Trace()
	out := l.String()
	if !strings.Contains(out, "not x") {
		t.Errorf("String() should contain 'not x' for negated variable, got:\n%s", out)
	}
}

// ---- Trace: nested function scopes ----

func TestTrace_NestedFunctionScopes(t *testing.T) {
	// Ensure Trace() completes without panic for two levels of nesting.
	l := NewLogger()
	l.EnterFunction("@__run", 2)
	l.EnterFunction("outer", 2)
	l.EnterFunction("inner", 2)
	l.UpdateVariable("spec_y_1", false)
	l.ExitFunction("inner", 2)
	l.ExitFunction("outer", 2)
	l.ExitFunction("@__run", 2)

	l.Results["spec_y_1"] = "42"
	l.Trace() // must not panic
}

// ---- String(): synthesis choice appears in output ----

func TestString_SynthChoiceInOutput(t *testing.T) {
	l := NewLogger()

	l.EnterFunction("@__run", 1)
	l.EnterFunction("synth_1", 1)
	l.EnterFunction("synth_1_fill", 1)
	l.UpdateVariable("spec_t_level_1", false)
	l.ExitFunction("synth_1_fill", 1)
	l.ExitFunction("synth_1", 1)
	l.ExitFunction("@__run", 1)

	l.Results["synth_1_fill_1"] = "true"
	l.Results["synth_1_drain_1"] = "false"
	l.Results["spec_t_level_1"] = "60"

	l.Trace()

	out := l.String()
	if !strings.Contains(out, "Fault chose fill") {
		t.Errorf("String() should contain 'Fault chose fill', got:\n%s", out)
	}
}

// ============================================================
// Phase 2: failing tests for spec-type-aware rendering split
// ============================================================

// ---- Boolean logic: true-initialized string rule appears (issue #80) ----

// TestString_BooleanLogic_TrueRuleVisibleInTrace is the exact scenario from
// issue #80: a string-rule variable whose _0 result is "true" is absent from
// the step trace because the pre-seed loop sets currentState[base]="true",
// then the display guard suppresses output when hasOldValue && oldValue==newValue.
// A "false"-initialized rule appears correctly because it is never pre-seeded.
// After the split the boolean logic path must not pre-seed at all.
func TestString_BooleanLogic_TrueRuleVisibleInTrace(t *testing.T) {
	l := NewLogger()

	// Temporal spec with a string rule: the step function updates the rule variable.
	l.EnterFunction("@__run", 1)
	l.EnterFunction("spec_fl_fn", 1)
	l.UpdateVariable("test_str1_1", false) // step fires, still true
	l.ExitFunction("spec_fl_fn", 1)
	l.ExitFunction("@__run", 1)

	// _0 is the initial value (true), _1 is after the step (still true).
	l.Results["test_str1_0"] = "true"
	l.Results["test_str1_1"] = "true"
	l.StringRules["test_str1"] = "is a fish"
	l.IsStringRule["test_str1"] = true

	l.Trace()
	out := l.String()

	// The step trace must show "is a fish is TRUE" inside the step function body.
	// The Initialize model section also shows it (via latestResult), so we must
	// verify it appears AFTER "Run function" — i.e. in the trace, not just the header.
	// Currently the pre-seed makes hasOldValue=true with oldValue="true", so
	// the first-appearance branch never fires and the line is silently dropped from
	// the trace even though Initialize model shows it correctly.
	runIdx := strings.Index(out, "Run function")
	ruleIdx := strings.LastIndex(out, "is a fish is TRUE")
	if runIdx == -1 {
		t.Fatalf("issue #80: step function not found in output:\n%s", out)
	}
	if ruleIdx == -1 || ruleIdx < runIdx {
		t.Errorf("issue #80: true-initialized string rule must appear in step trace (after 'Run function'), got:\n%s", out)
	}
}

// TestString_BooleanLogic_FalseRuleVisibleInTrace verifies that a false-initialized
// string rule appears in the trace (this already works today — it is not pre-seeded).
func TestString_BooleanLogic_FalseRuleVisibleInTrace(t *testing.T) {
	l := NewLogger()

	l.EnterFunction("@__run", 1)
	l.EnterFunction("spec_fl_fn", 1)
	l.UpdateVariable("test_str1_1", false)
	l.ExitFunction("spec_fl_fn", 1)
	l.ExitFunction("@__run", 1)

	l.Results["test_str1_0"] = "false"
	l.Results["test_str1_1"] = "false"
	l.StringRules["test_str1"] = "is a fish"
	l.IsStringRule["test_str1"] = true

	l.Trace()
	out := l.String()
	if !strings.Contains(out, "is a fish is FALSE") {
		t.Errorf("false-initialized string rule must appear in step trace, got:\n%s", out)
	}
}

// TestString_BooleanLogic_NoStartModel verifies that boolean logic specs
// (empty run block) never emit a "Start model" section.
func TestString_BooleanLogic_NoStartModel(t *testing.T) {
	l := NewLogger()

	l.EnterFunction("@__run", 1)
	l.UpdateVariable("test_str1_0", false)
	l.ExitFunction("@__run", 1)

	l.Results["test_str1_0"] = "true"
	l.StringRules["test_str1"] = "is a fish"
	l.IsStringRule["test_str1"] = true

	l.Trace()
	out := l.String()
	if strings.Contains(out, "Start model") {
		t.Errorf("Boolean logic: 'Start model' must not appear in output, got:\n%s", out)
	}
}

// ---- Temporal: true-initialized boolean stock still transitions true → false ----

// TestString_Temporal_TrueToFalseTransition verifies that the pre-seed loop
// is preserved for temporal specs: a stock initialized to true that later
// becomes false should render as "true → false", not "Set variable X to false".
func TestString_Temporal_TrueToFalseTransition(t *testing.T) {
	l := NewLogger()

	// Temporal: has a step function.
	l.EnterFunction("@__run", 1)
	l.EnterFunction("spec_fl_fn", 1)
	l.UpdateVariable("spec_st_value_1", false)
	l.ExitFunction("spec_fl_fn", 1)
	l.ExitFunction("@__run", 1)

	// _0 is true (initial), _1 is false (after the step).
	l.Results["spec_st_value_0"] = "true"
	l.Results["spec_st_value_1"] = "false"

	l.Trace()
	out := l.String()
	if !strings.Contains(out, "true → false") {
		t.Errorf("Temporal: true→false transition must appear, got:\n%s", out)
	}
}

// ---- Temporal: __state hoisting ----

// TestString_Temporal_StateHoisting verifies that for a function whose display
// name ends in "__state", the first "Set variable X to value true" line is
// hoisted out of the function body and appears before the section divider.
func TestString_Temporal_StateHoisting(t *testing.T) {
	l := NewLogger()

	l.EnterFunction("@__run", 1)
	l.EnterFunction("spec_comp_active__state", 1)
	l.UpdateVariable("spec_comp_active_1", false)
	l.ExitFunction("spec_comp_active__state", 1)
	l.ExitFunction("@__run", 1)

	l.Results["spec_comp_active_1"] = "true"

	l.Trace()
	out := l.String()

	startModel := "Start model"
	stateSet := "Set variable spec_comp_active to value true"

	startIdx := strings.Index(out, startModel)
	stateIdx := strings.Index(out, stateSet)

	if stateIdx == -1 {
		t.Fatalf("Temporal: hoisted state line not found in output:\n%s", out)
	}
	if startIdx == -1 {
		t.Fatalf("Temporal: 'Start model' section not found in output:\n%s", out)
	}
	// The hoisted state line must appear before "Start model", not inside the step function body.
	if stateIdx > startIdx {
		t.Errorf("Temporal: hoisted state line must appear before 'Start model', got:\n%s", out)
	}
}

// ---- Simulation mode: no asserts → show everything ----

// TestString_SimulationMode_ShowsAllVars verifies that when l.Asserts is empty
// (simulation run, no assertions), all variable updates are shown unfiltered.
func TestString_SimulationMode_ShowsAllVars(t *testing.T) {
	l := NewLogger()

	l.EnterFunction("@__run", 1)
	l.EnterFunction("spec_fl_fill", 1)
	l.UpdateVariable("spec_st_level_1", false)
	l.ExitFunction("spec_fl_fill", 1)
	l.ExitFunction("@__run", 1)

	l.Results["spec_st_level_1"] = "50"
	// No l.Asserts set — simulation mode.

	l.Trace()
	out := l.String()
	if !strings.Contains(out, "spec_st_level") {
		t.Errorf("Simulation mode: variable must appear when no asserts set, got:\n%s", out)
	}
}
