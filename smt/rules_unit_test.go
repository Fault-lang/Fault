package smt

import (
	"testing"

	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	irtypes "github.com/llir/llvm/ir/types"
)

func TestFetchIdent(t *testing.T) {
	g := NewGenerator()
	b := ir.NewBlock("test")
	alloc := b.NewAlloca(irtypes.I32)
	alloc.SetName("test_this_var")
	val := constant.NewInt(irtypes.I32, 0)
	store := b.NewStore(val, alloc)
	g.variables.loads["%1"] = store.Dst
	g.variables.ssa["test_this_var"] = 0
	g.variables.ref["%2"] = &infix{
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

	test1 := g.fetchIdent("%1", &wrap{
		value: "x",
		all:   true,
	})

	test2 := g.fetchIdent("%2", &wrap{
		value: "x",
		all:   true,
	})

	if test1.(*wrap).value != "%test_this_var_0" {
		t.Fatalf("fetchIdent returned the wrong result. got=%s", test1.String())
	}

	if test2.(*infix).x.String() != "x" || test2.(*infix).y.String() != "y" {
		t.Fatalf("fetchIdent returned the wrong result. got=%s", test2.String())
	}

	test3 := g.tempToIdent(g.variables.ref["%2"])

	if test3.(*infix).x.String() != "x" || test3.(*infix).y.String() != "y" {
		t.Fatalf("tempToIdent returned the wrong result. got=%s", test3.String())
	}
}

func TestStoreRule(t *testing.T) {
	g := NewGenerator()
	b := ir.NewBlock("test")
	alloc := b.NewAlloca(irtypes.I32)
	alloc.SetName("test_this_var")
	val := constant.NewInt(irtypes.I32, 0)
	store := b.NewStore(val, alloc)

	alloc2 := b.NewAlloca(irtypes.I32)
	alloc2.SetName("test_this_var2")
	store2 := b.NewStore(val, alloc2)

	g.variables.loads["%1"] = store.Dst
	g.variables.ssa["test_this_var"] = 0
	g.variables.ref["%2"] = &infix{
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

	test1 := g.storeRule(store, []rule{})
	if len(test1) != 1 {
		t.Fatalf("storeRule did not store new rule. got=%d", len(test1))
	}
	test2 := g.storeRule(store2, []rule{})
	if len(test2) != 1 {
		t.Fatalf("storeRule did not store new rule. got=%d", len(test2))
	}
	if test2[0].(*infix).x.String() != "test_this_var2_0" {
		t.Fatalf("storeRule malformed rule got=%s", test2[0].(*infix).x.String())
	}

	if test2[0].(*infix).y.String() != "0" {
		t.Fatalf("storeRule malformed rule got=%s", test2[0].(*infix).y.String())
	}
}
