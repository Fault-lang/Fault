package smt

import (
	"testing"

	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	irtypes "github.com/llir/llvm/ir/types"
)

func TestConvertIdent(t *testing.T) {
	g := NewGenerator()
	b := ir.NewBlock("test")
	alloc := b.NewAlloca(irtypes.I32)
	alloc.SetName("test_this_var")
	val := constant.NewInt(irtypes.I32, 0)
	store := b.NewStore(val, alloc)
	g.variables.loads["@__run-%1"] = store.Dst
	g.variables.ssa["test_this_var"] = 0

	if g.variables.convertIdent("@__run", "%test_this_var") != "test_this_var_0" {
		t.Fatalf("convertIdent returned the wrong value. got=%s", g.variables.convertIdent("@__run", "%test_this_var"))
	}

	if g.variables.convertIdent("@__run", "%1") != "test_this_var_0" {
		t.Fatalf("convertIdent returned the wrong value. got=%s", g.variables.convertIdent("@__run", "%1"))
	}

}

func TestNewConstants(t *testing.T) {
	g := NewGenerator()
	globals := []*ir.Global{
		ir.NewGlobalDef("test1", constant.NewFloat(irtypes.Double, 10)),
		ir.NewGlobalDef("test2", constant.NewFloat(irtypes.Double, 20)),
		ir.NewGlobalDef("test3", constant.NewFloat(irtypes.Double, 30)),
	}
	g.variables.ssa["test1"] = 0
	g.variables.ssa["test2"] = 2
	g.variables.ssa["test3"] = 5

	results := g.newConstants(globals)

	if len(results) != 3 {
		t.Fatalf("newConstants returned an incorrect number of results. got=%d", len(results))
	}

	if results[0] != "(assert (= test1_1 double 10.0))" {
		t.Fatalf("newConstants returned an incorrect value at index 0. got=%s", results[0])
	}

	if results[1] != "(assert (= test2_3 double 20.0))" {
		t.Fatalf("newConstants returned an incorrect value at index 1. got=%s", results[1])
	}

	if results[2] != "(assert (= test3_6 double 30.0))" {
		t.Fatalf("newConstants returned an incorrect value at index 2. got=%s", results[2])
	}
}

func TestAssrtType(t *testing.T) {
	a := &assrt{
		variable: &wrap{
			value: "x",
			all:   true,
		},
		conjunction: "&&",
		assertion: &wrap{
			value: "y",
			all:   true,
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
	i := &infix{
		x: &wrap{
			value: "x",
			all:   true,
		},
		y: &wrap{
			value: "y",
			all:   true,
		},
		op: ">",
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
	cond := &infix{
		x: &wrap{
			value: "x",
			all:   true,
		},
		y: &wrap{
			value: "y",
			all:   true,
		},
		op: ">",
	}

	i := &ite{
		cond: cond,
		t: []rule{&wrap{
			value: "x",
			all:   true,
		}},
		f: []rule{&wrap{
			value: "y",
			all:   true,
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
	i := &invariant{
		left: &wrap{
			value: "x",
			all:   true,
		},
		operator: "&&",
		right: &wrap{
			value: "y",
			all:   true,
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
	w := &wrap{
		value: "x",
		all:   true,
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
	w := &vwrap{
		value: val,
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
	wg := &wrapGroup{
		wraps: []*wrap{{
			value: "x",
			all:   true,
		}, {
			value: "y",
			all:   true,
		}, {
			value: "z",
			all:   true,
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
