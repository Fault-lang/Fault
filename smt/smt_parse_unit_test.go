package smt

import (
	"testing"

	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/enum"
	irtypes "github.com/llir/llvm/ir/types"
)

func TestTagRules(t *testing.T) {
	g := NewGenerator()
	rules := []rule{&wrap{value: "x"}, &infix{
		x: &wrap{
			value: "x",
			all:   true,
		},
		y: &wrap{
			value: "y",
			all:   true,
		},
		op: ">",
	}, &ite{
		cond: &wrap{
			value: "x",
			all:   true,
		},
		t: []rule{&wrap{
			value: "x",
			all:   true,
		}},
		f: []rule{&wrap{
			value: "y",
			all:   true,
		}},
	}, &vwrap{
		value: constant.NewInt(irtypes.I32, 0),
	}}
	r := g.tagRules(rules, "foo", "bar")
	if r[0].(*wrap).tag.branch != "foo" || r[0].(*wrap).tag.block != "bar" {
		t.Fatalf("tag not set correctly for rule %s. got=%s", rules[0], r[0].(*wrap).tag.String())
	}

	if r[1].(*infix).tag.branch != "foo" || r[1].(*infix).tag.block != "bar" {
		t.Fatalf("tag not set correctly for rule %s. got=%s", rules[1], r[1].(*infix).tag.String())
	}

	if r[2].(*ite).tag.branch != "foo" || r[2].(*ite).tag.block != "bar" {
		t.Fatalf("tag not set correctly for rule %s. got=%s", rules[2], r[2].(*ite).tag.String())
	}

	if r[3].(*vwrap).tag.branch != "foo" || r[3].(*vwrap).tag.block != "bar" {
		t.Fatalf("tag not set correctly for rule %s. got=%s", rules[3], r[3].(*vwrap).tag.String())
	}
}

func TestParseCond(t *testing.T) {
	g := NewGenerator()
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
	inst := g.parseCond(cond)

	if inst.(*infix).y.(*wrap).value != "y" {
		t.Fatalf("parseCond returns the wrong value. got=%s", inst.(*infix).String())
	}

	cond2 := &infix{
		x: &wrap{
			value: "x",
			all:   true,
		},
		y: &wrap{
			value: "y",
			all:   true,
		},
		op: "true",
	}
	inst2 := g.parseCond(cond2)

	if inst2.(*infix).y.(*wrap).value != "True" {
		t.Fatalf("parseCond returns the wrong value. got=%s", inst2.(*infix).String())
	}
}

func TestParseTerms(t *testing.T) {
	g := NewGenerator()

	b := ir.NewBlock("test")
	alloc := b.NewAlloca(irtypes.Double)
	alloc.SetName("test_this_var")
	val := constant.NewFloat(irtypes.Double, 2)
	store := b.NewStore(val, alloc)
	g.variables.ssa["test_this_var"] = 0

	alloc2 := b.NewAlloca(irtypes.Double)
	alloc2.SetName("test_this_var2")
	val2 := constant.NewFloat(irtypes.Double, 3)
	store2 := b.NewStore(val2, alloc2)
	g.variables.ssa["test_this_var2"] = 0

	terms := ir.NewBlock("test")
	terms.NewFCmp(enum.FPredOGT, store.Src, store2.Src)
	g.parseTerms([]*ir.Block{terms})
	if len(g.variables.ref) != 1 {
		t.Fatal("parse terms failed to save a rule.")
	}
	if g.variables.ref["%0"].(*infix).x.(*wrap).value != "2.0" {
		t.Fatalf("parse terms produced the wrong x value. got=%s", g.variables.ref["%0"].(*infix).x.(*wrap).value)
	}

	if g.variables.ref["%0"].(*infix).y.(*wrap).value != "3.0" {
		t.Fatalf("parse terms produced the wrong y value. got=%s", g.variables.ref["%0"].(*infix).y.(*wrap).value)
	}
}
