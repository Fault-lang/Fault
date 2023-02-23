package smt

import (
	"fault/smt/rules"
	"testing"

	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/enum"
	irtypes "github.com/llir/llvm/ir/types"
)

func TestCreateCondRule(t *testing.T) {
	g := NewGenerator()
	cond := &rules.Infix{
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
	inst := g.createCondRule(cond)

	if inst.(*rules.Infix).Y.(*rules.Wrap).Value != "y" {
		t.Fatalf("createCondRule returns the wrong value. got=%s", inst.(*rules.Infix).String())
	}

	cond2 := &rules.Infix{
		X: &rules.Wrap{
			Value: "x",
			All:   true,
		},
		Y: &rules.Wrap{
			Value: "y",
			All:   true,
		},
		Op: "true",
	}
	inst2 := g.createCondRule(cond2)

	if inst2.(*rules.Infix).Y.(*rules.Wrap).Value != "True" {
		t.Fatalf("createCondRule returns the wrong value. got=%s", inst2.(*rules.Infix).String())
	}
}

func TestParseTerms(t *testing.T) {
	g := NewGenerator()

	b := ir.NewBlock("test")
	alloc := b.NewAlloca(irtypes.Double)
	alloc.SetName("test_this_var")
	val := constant.NewFloat(irtypes.Double, 2)
	store := b.NewStore(val, alloc)
	g.variables.SSA["test_this_var"] = 0

	alloc2 := b.NewAlloca(irtypes.Double)
	alloc2.SetName("test_this_var2")
	val2 := constant.NewFloat(irtypes.Double, 3)
	store2 := b.NewStore(val2, alloc2)
	g.variables.SSA["test_this_var2"] = 0

	terms := ir.NewBlock("test-true")
	terms.NewFCmp(enum.FPredOGT, store.Src, store2.Src)
	g.parseTerms([]*ir.Block{terms})
	if len(g.variables.Ref) != 1 {
		t.Fatal("parse terms failed to save a rule.")
	}
	if g.variables.Ref["@__run-%0"].(*rules.Infix).X.(*rules.Wrap).Value != "2.0" {
		t.Fatalf("parse terms produced the wrong x value. got=%s", g.variables.Ref["%0"].(*rules.Infix).X.(*rules.Wrap).Value)
	}

	if g.variables.Ref["@__run-%0"].(*rules.Infix).Y.(*rules.Wrap).Value != "3.0" {
		t.Fatalf("parse terms produced the wrong y value. got=%s", g.variables.Ref["%0"].(*rules.Infix).Y.(*rules.Wrap).Value)
	}
}

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

func TestParallelPermutations(t *testing.T) {
	g := NewGenerator()
	test1 := []string{"foo", "bar"}
	results1 := g.parallelPermutations(test1)

	test2 := []string{"foo", "bar", "fizz", "buzz"}
	results2 := g.parallelPermutations(test2)

	test3 := []string{"foo", "bar", "fizz", "buzz", "foosh"}
	results3 := g.parallelPermutations(test3)

	if len(results1) != 2 {
		t.Fatalf("wrong number of permutations on set 1got=%d", len(results1))
	}

	if len(results2) != 24 {
		t.Fatalf("wrong number of permutations on set 2 got=%d", len(results2))
	}

	if len(results3) != 120 {
		t.Fatalf("wrong number of permutations on set 3 got=%d", len(results3))
	}
}

func TestConvertIdent(t *testing.T) {
	g := NewGenerator()
	b := ir.NewBlock("test")
	alloc := b.NewAlloca(irtypes.I32)
	alloc.SetName("test_this_var")
	val := constant.NewInt(irtypes.I32, 0)
	store := b.NewStore(val, alloc)
	g.variables.Loads["@__run-%1"] = store.Dst
	g.variables.SSA["test_this_var"] = 0

	if g.variables.ConvertIdent("@__run", "%test_this_var") != "test_this_var_0" {
		t.Fatalf("ConvertIdent returned the wrong value. got=%s", g.variables.ConvertIdent("@__run", "%test_this_var"))
	}

	if g.variables.ConvertIdent("@__run", "%1") != "test_this_var_0" {
		t.Fatalf("ConvertIdent returned the wrong value. got=%s", g.variables.ConvertIdent("@__run", "%1"))
	}

}

func TestNewConstants(t *testing.T) {
	g := NewGenerator()
	globals := []*ir.Global{
		ir.NewGlobalDef("test1", constant.NewFloat(irtypes.Double, 10)),
		ir.NewGlobalDef("test2", constant.NewFloat(irtypes.Double, 20)),
		ir.NewGlobalDef("test3", constant.NewFloat(irtypes.Double, 30)),
	}
	g.variables.SSA["test1"] = 0
	g.variables.SSA["test2"] = 2
	g.variables.SSA["test3"] = 5

	results := g.newConstants(globals)

	if len(results) != 3 {
		t.Fatalf("newConstants returned an incorrect number of results. got=%d", len(results))
	}

	if results[0] != "(assert (= test1_1 10.0))" {
		t.Fatalf("newConstants returned an incorrect value at index 0. got=%s", results[0])
	}

	if results[1] != "(assert (= test2_3 20.0))" {
		t.Fatalf("newConstants returned an incorrect value at index 1. got=%s", results[1])
	}

	if results[2] != "(assert (= test3_6 30.0))" {
		t.Fatalf("newConstants returned an incorrect value at index 2. got=%s", results[2])
	}
}
