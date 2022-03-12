package smt

import (
	"fault/ast"
	"testing"
)

func TestGenerateAsserts(t *testing.T) {
	g := NewGenerator()
	comp := ">"
	constr := &ast.Boolean{}
	stmt := &ast.AssertionStatement{TemporalFilter: "", TemporalN: 0}

	exp1 := &ast.IntegerLiteral{}
	results1 := g.generateAsserts(exp1, comp, constr, stmt)
	if len(results1) != 0 {
		t.Fatalf("generateAsserts returned wrong number of values. got=%d", len(results1))
	}

	exp2 := &ast.FloatLiteral{}
	results2 := g.generateAsserts(exp2, comp, constr, stmt)
	if len(results2) != 0 {
		t.Fatalf("generateAsserts returned wrong number of values. got=%d", len(results2))
	}

	exp3 := &ast.StringLiteral{}
	results3 := g.generateAsserts(exp3, comp, constr, stmt)
	if len(results3) != 0 {
		t.Fatalf("generateAsserts returned wrong number of values. got=%d", len(results3))
	}

	exp4 := &ast.Boolean{}
	results4 := g.generateAsserts(exp4, comp, constr, stmt)
	if len(results4) != 0 {
		t.Fatalf("generateAsserts returned wrong number of values. got=%d", len(results4))
	}
}

func TestGenerateAssertRules(t *testing.T) {
	g := NewGenerator()
	ru := &wrap{value: "test", constant: true}
	stmt := &ast.AssertionStatement{TemporalFilter: "", TemporalN: 0}
	r := g.generateAssertRules(ru, stmt.TemporalFilter, stmt.TemporalN)
	if len(r) != 1 {
		t.Fatalf("assert generation failed")
	}
	if r[0] != "test" {
		t.Fatalf("assert generation failed. got=%s want=test", r[0])
	}

	left := &invariant{
		left:        &wrap{value: "x", constant: true},
		right:       &wrap{value: "y", constant: true},
		conjunction: "&&",
	}

	right := &invariant{
		left:        &wrap{value: "z", constant: true},
		right:       &wrap{value: "a", constant: true},
		conjunction: "||",
	}

	ru2 := &invariant{
		left:        left,
		right:       right,
		conjunction: "&&",
	}

	r2 := g.generateAssertRules(ru2, stmt.TemporalFilter, stmt.TemporalN)

	if len(r2) != 1 {
		t.Fatalf("assert generation failed")
	}

	if r2[0] != "(&& (&& x y) (|| z a))" {
		t.Fatalf("assert generation failed. got=%s want=(&& (&& x y) (|| z a))", r2[0])
	}
}

func TestCaptureState(t *testing.T) {
	test1 := "test_constant"
	name, a, c := captureState(test1)
	if name != "" || a || !c {
		t.Fatal("captureState failed on a constant")
	}

	test2 := "test_this_var"
	name2, a2, c2 := captureState(test2)
	if name2 != "" || !a2 || c2 {
		t.Fatalf("captureState failed on a general state variablegot=%s %v %v", name2, a2, c2)
	}

	test3 := "test_this_var_2"
	name3, a3, c3 := captureState(test3)
	if name3 != "2" || a3 || c3 {
		t.Fatalf("captureState failed on a specific state variable. got=%s %v %v", name3, a3, c3)
	}

}
