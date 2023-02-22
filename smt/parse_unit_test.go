package smt

import (
	"fault/smt/rules"
	"testing"

	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/enum"
	irtypes "github.com/llir/llvm/ir/types"
)

func TestcreateCondRule(t *testing.T) {
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
