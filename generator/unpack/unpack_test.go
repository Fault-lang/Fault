package unpack

import (
	"fault/generator/rules"
	"fault/generator/scenario"
	"fmt"
	"strings"
	"testing"
)

// ---- strictOr ----

func TestStrictOr_Empty(t *testing.T) {
	result := strictOr([]string{})
	// No crash; result is vacuously an or of nothing
	_ = result
}

func TestStrictOr_Single(t *testing.T) {
	result := strictOr([]string{"A"})
	if result != "A" {
		t.Errorf("single element strictOr = %q, want %q", result, "A")
	}
}

func TestStrictOr_Two(t *testing.T) {
	result := strictOr([]string{"A", "B"})
	// Expect: (or (and A (not B)) (and B (not A)))
	if !strings.HasPrefix(result, "(or ") {
		t.Errorf("two-element strictOr should start with (or ...), got: %s", result)
	}
	if !strings.Contains(result, "(not A)") {
		t.Errorf("strictOr should negate A in B's branch, got: %s", result)
	}
	if !strings.Contains(result, "(not B)") {
		t.Errorf("strictOr should negate B in A's branch, got: %s", result)
	}
}

func TestStrictOr_Three(t *testing.T) {
	result := strictOr([]string{"A", "B", "C"})
	// Each branch uses (and X (not Y) (not Z))
	if !strings.HasPrefix(result, "(or ") {
		t.Errorf("three-element strictOr should start with (or ...), got: %s", result)
	}
	// Each element appears both bare and negated
	for _, name := range []string{"A", "B", "C"} {
		if !strings.Contains(result, name) {
			t.Errorf("strictOr result missing %q: %s", name, result)
		}
		if !strings.Contains(result, fmt.Sprintf("(not %s)", name)) {
			t.Errorf("strictOr result missing (not %s): %s", name, result)
		}
	}
}

// ---- InitsToList ----

func TestInitsToList_Empty(t *testing.T) {
	result := InitsToList(nil)
	if len(result) != 0 {
		t.Errorf("InitsToList(nil) = %v, want []", result)
	}
}

func TestInitsToList_Single(t *testing.T) {
	i := rules.NewInit("myvar", "Real", 3, nil, false, false)
	result := InitsToList([]*rules.Init{i})
	if len(result) != 1 || result[0] != "myvar_3" {
		t.Errorf("InitsToList single = %v, want [myvar_3]", result)
	}
}

func TestInitsToList_Multiple(t *testing.T) {
	inits := []*rules.Init{
		rules.NewInit("a", "Real", 0, nil, false, false),
		rules.NewInit("b", "Real", 1, nil, false, false),
		rules.NewInit("c", "Bool", 2, nil, false, false),
	}
	result := InitsToList(inits)
	if len(result) != 3 {
		t.Fatalf("InitsToList multiple: got %d results, want 3", len(result))
	}
	if result[0] != "a_0" || result[1] != "b_1" || result[2] != "c_2" {
		t.Errorf("InitsToList multiple = %v, want [a_0 b_1 c_2]", result)
	}
}

// ---- Register / registry key format ----

func TestRegister_Empty(t *testing.T) {
	u := NewUnpacker("@__run")
	u.Register(nil)
	if len(u.Registry) != 0 {
		t.Errorf("Register(nil) should not add entries, got %v", u.Registry)
	}
}

func TestRegister_BuildsCorrectKey(t *testing.T) {
	u := NewUnpacker("@__run")
	i := rules.NewInit("spec_s_value", "Real", 0, nil, false, false)
	i.SetRound(2)
	u.Register([]*rules.Init{i})

	key := "round-2_@__run"
	if _, ok := u.Registry[key]; !ok {
		t.Errorf("expected registry key %q, got keys: %v", key, keys(u.Registry))
	}
	if len(u.Registry[key]) != 1 {
		t.Errorf("expected 1 entry under key %q, got %d", key, len(u.Registry[key]))
	}
	if u.Registry[key][0][0] != "spec_s_value" {
		t.Errorf("registry entry[0] = %v, want [spec_s_value 0]", u.Registry[key][0])
	}
}

func TestRegister_AccumulatesAcrossCalls(t *testing.T) {
	u := NewUnpacker("@__run")
	i1 := rules.NewInit("a", "Real", 0, nil, false, false)
	i1.SetRound(0)
	i2 := rules.NewInit("b", "Real", 1, nil, false, false)
	i2.SetRound(0)
	u.Register([]*rules.Init{i1})
	u.Register([]*rules.Init{i2})

	key := "round-0_@__run"
	if len(u.Registry[key]) != 2 {
		t.Errorf("expected 2 entries under key %q, got %d", key, len(u.Registry[key]))
	}
}

// ---- InitVars ----

func TestInitVars_Empty(t *testing.T) {
	u := NewUnpacker("@__run")
	result := u.InitVars()
	if len(result) != 0 {
		t.Errorf("InitVars() on empty unpacker = %v, want []", result)
	}
}

func TestInitVars_DeclaresVariables(t *testing.T) {
	u := NewUnpacker("@__run")
	// Non-global init: WriteRule just emits (declare-fun ...) with no log access.
	u.Inits = append(u.Inits, rules.NewInit("myvar", "Real", 0, nil, false, false))
	result := u.InitVars()
	if len(result) != 1 {
		t.Fatalf("InitVars: expected 1 declaration, got %d: %v", len(result), result)
	}
	if !strings.Contains(result[0], "declare-fun") {
		t.Errorf("InitVars result should contain declare-fun, got: %s", result[0])
	}
	if !strings.Contains(result[0], "myvar_0") {
		t.Errorf("InitVars result should contain myvar_0, got: %s", result[0])
	}
}

func TestInitVars_DeduplicatesIdenticalVars(t *testing.T) {
	u := NewUnpacker("@__run")
	// Add the same variable twice (can happen when candidates share variables)
	init := rules.NewInit("myvar", "Real", 0, nil, false, false)
	u.Inits = append(u.Inits, init, init)
	result := u.InitVars()
	if len(result) != 1 {
		t.Errorf("InitVars should deduplicate: got %d declarations, want 1", len(result))
	}
}

// ---- GetPhis ----

func TestGetPhis_NoChange(t *testing.T) {
	u := NewUnpacker("@__run")
	start := rules.NewSSA()
	start.Update("x") // x=0
	end := start.Clone()
	phis := u.GetPhis(start, end)
	if len(phis) != 0 {
		t.Errorf("GetPhis with no changes should return empty map, got %v", phis)
	}
}

func TestGetPhis_OneChange(t *testing.T) {
	u := NewUnpacker("@__run")
	start := rules.NewSSA()
	start.Update("x") // x=0
	start.Update("y") // y=0
	end := start.Clone()
	end.Update("x") // x=1 (changed)
	// y unchanged

	phis := u.GetPhis(start, end)
	if len(phis) != 1 {
		t.Fatalf("expected 1 phi, got %d: %v", len(phis), phis)
	}
	if _, ok := phis["x"]; !ok {
		t.Errorf("expected phi for x, got: %v", phis)
	}
	if phis["x"][0] != 0 || phis["x"][1] != 1 {
		t.Errorf("phi for x = %v, want [0 1]", phis["x"])
	}
}

func TestGetPhis_MultipleChanges(t *testing.T) {
	u := NewUnpacker("@__run")
	start := rules.NewSSA()
	start.Update("a") // a=0
	start.Update("b") // b=0
	end := start.Clone()
	end.Update("a") // a=1
	end.Update("b") // b=1

	phis := u.GetPhis(start, end)
	if len(phis) != 2 {
		t.Errorf("expected 2 phis, got %d: %v", len(phis), phis)
	}
}

// ---- unpackSynthSlot ----

func TestUnpackSynthSlot_Empty(t *testing.T) {
	u := NewUnpacker("@__run")
	slot := &rules.SynthSlot{
		Round:      1,
		Candidates: map[string][]rules.Rule{},
	}
	inits, rule := u.unpackSynthSlot(slot)
	if inits != nil || rule != "" {
		t.Errorf("empty SynthSlot should return nil, \"\"; got %v, %q", inits, rule)
	}
}

// synthUnpacker creates an Unpacker with varName pre-populated in SSA and VarTypes,
// so that candidates modifying varName produce a detectable phi.
func synthUnpacker(varName string) *Unpacker {
	u := NewUnpacker("@__run")
	u.VarTypes[varName] = "Real"
	u.SSA.Update(varName) // varName now at SSA version 0
	return u
}

// stateChangeRule builds an Infix "= varName_new 10" rule that bumps varName's SSA.
func stateChangeRule(varName string) rules.Rule {
	lhs := rules.NewWrap(varName, "Real", true, "", 0, true, false) // Init=true → updates SSA
	rhs := rules.NewWrap("10", "Real", false, "", 0, false, false)
	return &rules.Infix{X: lhs, Y: rhs, Op: "="}
}

func TestUnpackSynthSlot_TwoCandidates_SelectorsCreated(t *testing.T) {
	// fill modifies level (state change); drain also modifies level (other direction).
	// Identity phis ensure both branches have non-empty caps.
	// Verifies: selector vars declared, XOR constraint emitted.
	varName := "spec_ops_t_level"
	u := synthUnpacker(varName)
	slot := &rules.SynthSlot{
		Round: 1,
		Candidates: map[string][]rules.Rule{
			"fill":  {stateChangeRule(varName)},
			"drain": {stateChangeRule(varName)},
		},
	}
	_, rule := u.unpackSynthSlot(slot)

	if !strings.Contains(rule, "synth_1_drain") {
		t.Errorf("expected selector synth_1_drain in output, got:\n%s", rule)
	}
	if !strings.Contains(rule, "synth_1_fill") {
		t.Errorf("expected selector synth_1_fill in output, got:\n%s", rule)
	}
	// The XOR (strictOr) constraint: one true, one false
	if !strings.Contains(rule, "(not synth_1_drain") && !strings.Contains(rule, "(not synth_1_fill") {
		t.Errorf("expected strictOr negations in output, got:\n%s", rule)
	}
}

func TestUnpackSynthSlot_SelectorsAreDeterministic(t *testing.T) {
	// Candidates are iterated in sorted order; output must be stable across runs.
	varName := "spec_ctrl_val"
	run := func() string {
		u := synthUnpacker(varName)
		slot := &rules.SynthSlot{
			Round: 0,
			Candidates: map[string][]rules.Rule{
				"z_func": {stateChangeRule(varName)},
				"a_func": {stateChangeRule(varName)},
				"m_func": {stateChangeRule(varName)},
			},
		}
		_, rule := u.unpackSynthSlot(slot)
		return rule
	}
	r1 := run()
	r2 := run()
	if r1 != r2 {
		t.Errorf("unpackSynthSlot output is non-deterministic:\nrun1: %s\nrun2: %s", r1, r2)
	}
}

func TestUnpackSynthSlot_SelectorNameFormat(t *testing.T) {
	// Selector name format: synth_{round}_{funcname}
	// The full declared name (used in assertions) is synth_{round}_{funcname}_{round}.
	varName := "spec_c_val"
	u := synthUnpacker(varName)
	slot := &rules.SynthSlot{
		Round: 3,
		Candidates: map[string][]rules.Rule{
			"increment": {stateChangeRule(varName)},
		},
	}
	_, rule := u.unpackSynthSlot(slot)

	if !strings.Contains(rule, "synth_3_increment") {
		t.Errorf("expected synth_3_increment in output, got:\n%s", rule)
	}
}

func TestUnpackSynthSlot_ImplicationAssertion(t *testing.T) {
	// The non-empty candidate branch must produce an implication: (assert (=> selector rule))
	varName := "spec_ctrl_c_value"
	u := synthUnpacker(varName)
	slot := &rules.SynthSlot{
		Round: 1,
		Candidates: map[string][]rules.Rule{
			"add": {stateChangeRule(varName)},
			"sub": {stateChangeRule(varName)},
		},
	}
	_, rule := u.unpackSynthSlot(slot)

	if !strings.Contains(rule, "(assert (=>") {
		t.Errorf("expected implication assertion (assert (=> ...)) in output, got:\n%s", rule)
	}
	if !strings.Contains(rule, "synth_1_add") {
		t.Errorf("missing synth_1_add selector in:\n%s", rule)
	}
	if !strings.Contains(rule, "synth_1_sub") {
		t.Errorf("missing synth_1_sub selector in:\n%s", rule)
	}
}

func TestUnpackSynthSlot_LogsEnterExit(t *testing.T) {
	// The synthesis slot should record Enter/Exit events in the scenario log.
	varName := "spec_ctrl_v"
	u := synthUnpacker(varName)
	slot := &rules.SynthSlot{
		Round: 2,
		Candidates: map[string][]rules.Rule{
			"pick": {stateChangeRule(varName)},
		},
	}
	u.unpackSynthSlot(slot)

	slotName := "synth_2"
	found := false
	for _, e := range u.Log.Events {
		if fc, ok := e.(*scenario.FunctionCall); ok && strings.Contains(fc.FunctionName, slotName) {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("expected Enter/Exit FunctionCall event for %q in log", slotName)
	}
}

// ---- helpers ----

func keys(m map[string][][]string) []string {
	var ks []string
	for k := range m {
		ks = append(ks, k)
	}
	return ks
}
