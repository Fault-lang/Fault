package rules

import (
	"fault/generator/scenario"
	"testing"
)

func TestSSA_Get(t *testing.T) {
	ssa := NewSSA()
	ssa.variables["a"] = 1
	if got := ssa.Get("a"); got != 1 {
		t.Errorf("SSA.Get() = %v, want %v", got, 1)
	}
}

func TestSSA_Update(t *testing.T) {
	ssa := NewSSA()
	if got := ssa.Update("a"); got != 0 {
		t.Errorf("SSA.Update() = %v, want %v", got, 0)
	}
	if got := ssa.Update("a"); got != 1 {
		t.Errorf("SSA.Update() = %v, want %v", got, 1)
	}
}

func TestSSA_Clone(t *testing.T) {
	ssa := NewSSA()
	ssa.variables["a"] = 1
	clone := ssa.Clone()
	if got := clone.Get("a"); got != 1 {
		t.Errorf("SSA.Clone().Get() = %v, want %v", got, 1)
	}
	clone.Update("a")
	if got := ssa.Get("a"); got != 1 {
		t.Errorf("Original SSA.Get() after clone update = %v, want %v", got, 1)
	}
	if got := clone.Get("a"); got != 2 {
		t.Errorf("Clone SSA.Get() after update = %v, want %v", got, 2)
	}
}

func TestSSA_Iter(t *testing.T) {
	ssa := NewSSA()
	ssa.variables["a"] = 1
	iter := ssa.Iter()
	if got := iter["a"]; got != 1 {
		t.Errorf("SSA.Iter() = %v, want %v", got, 1)
	}
}

func TestNewSSA(t *testing.T) {
	ssa := NewSSA()
	if ssa.variables == nil {
		t.Errorf("NewSSA() = %v, want non-nil variables map", ssa.variables)
	}
}
func TestBasic_LoadContext(t *testing.T) {
	b := &Basic{}
	haveSeen := map[string]bool{"a": true}
	onEntry := map[string][]int16{"a": {1}}
	b.LoadContext(1, haveSeen, onEntry, scenario.NewLogger())
	if b.PhiLevel != 1 {
		t.Errorf("Basic.LoadContext() PhiLevel = %v, want %v", b.PhiLevel, 1)
	}
	if b.HaveSeen["a"] != true {
		t.Errorf("Basic.LoadContext() HaveSeen = %v, want %v", b.HaveSeen["a"], true)
	}
	if b.OnEntry["a"][0] != 1 {
		t.Errorf("Basic.LoadContext() OnEntry = %v, want %v", b.OnEntry["a"][0], 1)
	}
}

func TestBasic_WriteRule(t *testing.T) {
	x := &Wrap{Value: "x"}
	y := &Wrap{Value: "y"}
	b := &Basic{X: x, Y: y}
	ssa := NewSSA()
	init, rule, _ := b.WriteRule(ssa)
	if rule != "(assert x y)" {
		t.Errorf("Basic.WriteRule() = %v, want %v", rule, "(assert x y)")
	}
	if len(init) != 0 {
		t.Errorf("Basic.WriteRule() init length = %v, want %v", len(init), 0)
	}
}

func TestBasic_String(t *testing.T) {
	x := &Wrap{Value: "x"}
	y := &Wrap{Value: "y"}
	b := &Basic{X: x, Y: y}
	if got := b.String(); got != "basic x y" {
		t.Errorf("Basic.String() = %v, want %v", got, "basic x y")
	}
}

func TestBasic_Assertless(t *testing.T) {
	b := &Basic{}
	if got := b.Assertless(); got != "" {
		t.Errorf("Basic.Assertless() = %v, want %v", got, "")
	}
}

func TestBasic_IsTagged(t *testing.T) {
	b := &Basic{}
	if got := b.IsTagged(); got != false {
		t.Errorf("Basic.IsTagged() = %v, want %v", got, false)
	}
	b.Tag("branch1", "block1")
	if got := b.IsTagged(); got != true {
		t.Errorf("Basic.IsTagged() = %v, want %v", got, true)
	}
}

func TestBasic_Choice(t *testing.T) {
	b := &Basic{}
	b.Tag("branch1", "block1")
	if got := b.Choice(); got != "block1" {
		t.Errorf("Basic.Choice() = %v, want %v", got, "block1")
	}
}

func TestBasic_Branch(t *testing.T) {
	b := &Basic{}
	b.Tag("branch1", "block1")
	if got := b.Branch(); got != "branch1" {
		t.Errorf("Basic.Branch() = %v, want %v", got, "branch1")
	}
}

func TestBasic_Tag(t *testing.T) {
	b := &Basic{}
	b.Tag("branch1", "block1")
	if b.tag.branch != "branch1" {
		t.Errorf("Basic.Tag() branch = %v, want %v", b.tag.branch, "branch1")
	}
	if b.tag.block != "block1" {
		t.Errorf("Basic.Tag() block = %v, want %v", b.tag.block, "block1")
	}
}
func TestInfix_LoadContext(t *testing.T) {
	i := &Infix{}
	haveSeen := map[string]bool{"a": true}
	onEntry := map[string][]int16{"a": {1}}
	i.LoadContext(1, haveSeen, onEntry, scenario.NewLogger())
	if i.PhiLevel != 1 {
		t.Errorf("Infix.LoadContext() PhiLevel = %v, want %v", i.PhiLevel, 1)
	}
	if i.HaveSeen["a"] != true {
		t.Errorf("Infix.LoadContext() HaveSeen = %v, want %v", i.HaveSeen["a"], true)
	}
	if i.OnEntry["a"][0] != 1 {
		t.Errorf("Infix.LoadContext() OnEntry = %v, want %v", i.OnEntry["a"][0], 1)
	}
}

func TestInfix_WriteRule(t *testing.T) {
	x := &Wrap{Value: "x"}
	y := &Wrap{Value: "y"}
	i := &Infix{X: x, Y: y, Op: "+"}
	ssa := NewSSA()
	init, rule, _ := i.WriteRule(ssa)
	if rule != "(+ x y)" {
		t.Errorf("Infix.WriteRule() = %v, want %v", rule, "(+ x y)")
	}
	if len(init) != 0 {
		t.Errorf("Infix.WriteRule() init length = %v, want %v", len(init), 0)
	}
}

func TestInfix_String(t *testing.T) {
	x := &Wrap{Value: "x"}
	y := &Wrap{Value: "y"}
	i := &Infix{X: x, Y: y, Op: "+"}
	if got := i.String(); got != "x + y" {
		t.Errorf("Infix.String() = %v, want %v", got, "x + y")
	}
}

func TestInfix_Assertless(t *testing.T) {
	x := &Wrap{Value: "x"}
	y := &Wrap{Value: "y"}
	i := &Infix{X: x, Y: y, Op: "+"}
	if got := i.Assertless(); got != "(+ x y)" {
		t.Errorf("Infix.Assertless() = %v, want %v", got, "(+ x y)")
	}
}

func TestInfix_IsTagged(t *testing.T) {
	i := &Infix{}
	if got := i.IsTagged(); got != false {
		t.Errorf("Infix.IsTagged() = %v, want %v", got, false)
	}
	i.Tag("branch1", "block1")
	if got := i.IsTagged(); got != true {
		t.Errorf("Infix.IsTagged() = %v, want %v", got, true)
	}
}

func TestInfix_Choice(t *testing.T) {
	i := &Infix{}
	i.Tag("branch1", "block1")
	if got := i.Choice(); got != "block1" {
		t.Errorf("Infix.Choice() = %v, want %v", got, "block1")
	}
}

func TestInfix_Branch(t *testing.T) {
	i := &Infix{}
	i.Tag("branch1", "block1")
	if got := i.Branch(); got != "branch1" {
		t.Errorf("Infix.Branch() = %v, want %v", got, "branch1")
	}
}

func TestInfix_Tag(t *testing.T) {
	i := &Infix{}
	i.Tag("branch1", "block1")
	if i.tag.branch != "branch1" {
		t.Errorf("Infix.Tag() branch = %v, want %v", i.tag.branch, "branch1")
	}
	if i.tag.block != "block1" {
		t.Errorf("Infix.Tag() block = %v, want %v", i.tag.block, "block1")
	}
}

func TestWrap_LoadContext(t *testing.T) {
	w := &Wrap{}
	haveSeen := map[string]bool{"a": true}
	onEntry := map[string][]int16{"a": {1}}
	w.LoadContext(1, haveSeen, onEntry, scenario.NewLogger())
	if w.PhiLevel != 1 {
		t.Errorf("Wrap.LoadContext() PhiLevel = %v, want %v", w.PhiLevel, 1)
	}
	if w.HaveSeen["a"] != true {
		t.Errorf("Wrap.LoadContext() HaveSeen = %v, want %v", w.HaveSeen["a"], true)
	}
	if w.OnEntry["a"][0] != 1 {
		t.Errorf("Wrap.LoadContext() OnEntry = %v, want %v", w.OnEntry["a"][0], 1)
	}
}

func TestWrap_WriteRule(t *testing.T) {
	w := &Wrap{Value: "x", Variable: true, Type: "Int", Init: true}
	ssa := NewSSA()
	init, rule, _ := w.WriteRule(ssa)
	if rule != "x_0" {
		t.Errorf("Wrap.WriteRule() = %v, want %v", rule, "x_0")
	}
	if len(init) != 1 {
		t.Errorf("Wrap.WriteRule() init length = %v, want %v", len(init), 1)
	}
	if init[0].Ident != "x_0" {
		t.Errorf("Wrap.WriteRule() init Ident = %v, want %v", init[0].Ident, "x_0")
	}
	if init[0].Type != "Int" {
		t.Errorf("Wrap.WriteRule() init Type = %v, want %v", init[0].Type, "Int")
	}
	if init[0].Value != "0" {
		t.Errorf("Wrap.WriteRule() init Value = %v, want %v", init[0].Value, "0")
	}
}

func TestWrap_String(t *testing.T) {
	w := &Wrap{Value: "x"}
	if got := w.String(); got != "x" {
		t.Errorf("Wrap.String() = %v, want %v", got, "x")
	}
}

func TestWrap_Assertless(t *testing.T) {
	w := &Wrap{Value: "x"}
	if got := w.Assertless(); got != "x" {
		t.Errorf("Wrap.Assertless() = %v, want %v", got, "x")
	}
}

func TestWrap_IsTagged(t *testing.T) {
	w := &Wrap{}
	if got := w.IsTagged(); got != false {
		t.Errorf("Wrap.IsTagged() = %v, want %v", got, false)
	}
	w.Tag("branch1", "block1")
	if got := w.IsTagged(); got != true {
		t.Errorf("Wrap.IsTagged() = %v, want %v", got, true)
	}
}

func TestWrap_Choice(t *testing.T) {
	w := &Wrap{}
	w.Tag("branch1", "block1")
	if got := w.Choice(); got != "block1" {
		t.Errorf("Wrap.Choice() = %v, want %v", got, "block1")
	}
}

func TestWrap_Branch(t *testing.T) {
	w := &Wrap{}
	w.Tag("branch1", "block1")
	if got := w.Branch(); got != "branch1" {
		t.Errorf("Wrap.Branch() = %v, want %v", got, "branch1")
	}
}

func TestWrap_Tag(t *testing.T) {
	w := &Wrap{}
	w.Tag("branch1", "block1")
	if w.tag.branch != "branch1" {
		t.Errorf("Wrap.Tag() branch = %v, want %v", w.tag.branch, "branch1")
	}
	if w.tag.block != "block1" {
		t.Errorf("Wrap.Tag() block = %v, want %v", w.tag.block, "block1")
	}
}
