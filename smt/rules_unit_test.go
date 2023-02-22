package smt

import (
	"fault/smt/rules"
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
	g.variables.Loads["@__run-%1"] = store.Dst
	g.variables.SSA["test_this_var"] = 0
	g.variables.Ref["@__run-%2"] = &rules.Infix{
		X: &rules.Wrap{
			Value: "x",
			All:   true,
		},
		Y: &rules.Wrap{
			Value: "y",
			All:   true,
		},
		Op: ">",
	}

	test1 := g.fetchIdent("%1", &rules.Wrap{
		Value: "x",
		All:   true,
	})

	test2 := g.fetchIdent("%2", &rules.Wrap{
		Value: "x",
		All:   true,
	})

	if test1.(*rules.Wrap).Value != "%test_this_var_0" {
		t.Fatalf("fetchIdent returned the wrong result. got=%s", test1.String())
	}

	if test2.(*rules.Infix).X.String() != "x" || test2.(*rules.Infix).Y.String() != "y" {
		t.Fatalf("fetchIdent returned the wrong result. got=%s", test2.String())
	}

	test3 := g.tempToIdent(g.variables.Ref["@__run-%2"])

	if test3.(*rules.Infix).X.String() != "x" || test3.(*rules.Infix).Y.String() != "y" {
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

	g.variables.Loads["@__run-%1"] = store.Dst
	g.variables.SSA["test_this_var"] = 0
	g.variables.Ref["@__run-%2"] = &rules.Infix{
		X: &rules.Wrap{
			Value: "x",
			All:   true,
		},
		Y: &rules.Wrap{
			Value: "y",
			All:   true,
		},
		Op: ">",
	}

	test1 := g.storeRule(store)
	if len(test1) != 1 {
		t.Fatalf("storeRule did not store new rule. got=%d", len(test1))
	}
	test2 := g.storeRule(store2)
	if len(test2) != 1 {
		t.Fatalf("storeRule did not store new rule. got=%d", len(test2))
	}
	if test2[0].(*rules.Infix).X.String() != "test_this_var2_0" {
		t.Fatalf("storeRule malformed rule got=%s", test2[0].(*rules.Infix).X.String())
	}

	if test2[0].(*rules.Infix).Y.String() != "0" {
		t.Fatalf("storeRule malformed rule got=%s", test2[0].(*rules.Infix).Y.String())
	}
}
