package rules

import (
	"testing"

	"github.com/llir/llvm/ir/constant"
	irtypes "github.com/llir/llvm/ir/types"
)

func TestTagRules(t *testing.T) {
	rules := []Rule{&Wrap{Value: "x"}, &Infix{
		X: &Wrap{
			Value: "x",
			All:   true,
		},
		Y: &Wrap{
			Value: "y",
			All:   true,
		},
		Op: ">",
	}, &Ite{
		Cond: &Wrap{
			Value: "x",
			All:   true,
		},
		T: []Rule{&Wrap{
			Value: "x",
			All:   true,
		}},
		F: []Rule{&Wrap{
			Value: "y",
			All:   true,
		}},
	}, &Vwrap{
		Value: constant.NewInt(irtypes.I32, 0),
	}}
	r := TagRules(rules, "foo", "bar")
	if r[0].(*Wrap).tag.branch != "foo" || r[0].(*Wrap).tag.block != "bar" {
		t.Fatalf("tag not set correctly for rule %s. got=%s", rules[0], r[0].(*Wrap).tag.String())
	}

	if r[1].(*Infix).tag.branch != "foo" || r[1].(*Infix).tag.block != "bar" {
		t.Fatalf("tag not set correctly for rule %s. got=%s", rules[1], r[1].(*Infix).tag.String())
	}

	if r[2].(*Ite).tag.branch != "foo" || r[2].(*Ite).tag.block != "bar" {
		t.Fatalf("tag not set correctly for rule %s. got=%s", rules[2], r[2].(*Ite).tag.String())
	}

	if r[3].(*Vwrap).tag.branch != "foo" || r[3].(*Vwrap).tag.block != "bar" {
		t.Fatalf("tag not set correctly for rule %s. got=%s", rules[3], r[3].(*Vwrap).tag.String())
	}
}

func TestAssrtType(t *testing.T) {
	a := &Assrt{
		Variable: &Wrap{
			Value: "x",
			All:   true,
		},
		Conjunction: "&&",
		Assertion: &Wrap{
			Value: "y",
			All:   true,
		},
	}

	a.Tag("test", "me")

	if a.tag.block != "me" || a.tag.branch != "test" {
		t.Fatalf("type tagged incorrectly. got=%s block %s branch", a.tag.block, a.tag.branch)
	}

	if a.String() != "x&&y" {
		t.Fatalf("String() failed got=%s", a.String())
	}
}

func TestInfixType(t *testing.T) {
	i := &Infix{
		X: &Wrap{
			Value: "x",
			All:   true,
		},
		Y: &Wrap{
			Value: "y",
			All:   true,
		},
		Op: ">",
	}

	i.Tag("test", "me")

	if i.tag.block != "me" || i.tag.branch != "test" {
		t.Fatalf("type tagged incorrectly. got=%s block %s branch", i.tag.block, i.tag.branch)
	}

	if i.String() != "x > y" {
		t.Fatalf("String() failed got=%s", i.String())
	}
}

func TestIfeType(t *testing.T) {
	cond := &Infix{
		X: &Wrap{
			Value: "x",
			All:   true,
		},
		Y: &Wrap{
			Value: "y",
			All:   true,
		},
		Op: ">",
	}

	i := &Ite{
		Cond: cond,
		T: []Rule{&Wrap{
			Value: "x",
			All:   true,
		}},
		F: []Rule{&Wrap{
			Value: "y",
			All:   true,
		}},
	}

	i.Tag("test", "me")

	if i.tag.block != "me" || i.tag.branch != "test" {
		t.Fatalf("type tagged incorrectly. got=%s block %s branch", i.tag.block, i.tag.branch)
	}

	if i.String() != "if x > y then [x] else [y]" {
		t.Fatalf("String() failed got=%s", i.String())
	}
}

func TestInvariantType(t *testing.T) {
	i := &Invariant{
		Left: &Wrap{
			Value: "x",
			All:   true,
		},
		Operator: "&&",
		Right: &Wrap{
			Value: "y",
			All:   true,
		},
	}

	i.Tag("test", "me")

	if i.tag.block != "me" || i.tag.branch != "test" {
		t.Fatalf("type tagged incorrectly. got=%s block %s branch", i.tag.block, i.tag.branch)
	}

	if i.String() != "x&&y" {
		t.Fatalf("String() failed got=%s", i.String())
	}
}

func TestWrapType(t *testing.T) {
	w := &Wrap{
		Value: "x",
		All:   true,
	}

	w.Tag("test", "me")

	if w.tag.block != "me" || w.tag.branch != "test" {
		t.Fatalf("assrt type tagged incorrectly. got=%s block %s branch", w.tag.block, w.tag.branch)
	}

	if w.String() != "x" {
		t.Fatalf("assrt String() failed got=%s", w.String())
	}
}

func TestVWrapType(t *testing.T) {
	val := constant.NewInt(irtypes.I32, 0)
	w := &Vwrap{
		Value: val,
	}

	w.Tag("test", "me")

	if w.tag.block != "me" || w.tag.branch != "test" {
		t.Fatalf("assrt type tagged incorrectly. got=%s block %s branch", w.tag.block, w.tag.branch)
	}

	if w.String() != "i32 0" {
		t.Fatalf("assrt String() failed got=%s", w.String())
	}
}

func TestWrapGroupType(t *testing.T) {
	wg := &WrapGroup{
		Wraps: []*Wrap{{
			Value: "x",
			All:   true,
		}, {
			Value: "y",
			All:   true,
		}, {
			Value: "z",
			All:   true,
		}}}

	wg.Tag("test", "me")

	if wg.tag.block != "me" || wg.tag.branch != "test" {
		t.Fatalf("assrt type tagged incorrectly. got=%s block %s branch", wg.tag.block, wg.tag.branch)
	}

	if wg.String() != "xyz" {
		t.Fatalf("assrt String() failed got=%s", wg.String())
	}
}

func TestBranchType(t *testing.T) {
	b := &branch{
		branch: "dummy",
		block:  "func",
	}

	if b.String() != "dummy.func" {
		t.Fatalf("assrt String() failed got=%s", b.String())
	}
}
