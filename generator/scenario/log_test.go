package scenario

import (
	"strings"
	"testing"
)

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
