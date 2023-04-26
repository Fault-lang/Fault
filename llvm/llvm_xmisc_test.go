package llvm

// "xmisc" is a little hack to keep the order of test files
// from screwing up the block numbering on LLVM IR (and thus
// causing a bunch of tests to fail.

import (
	"fault/ast"
	"testing"
)

func TestNegateTemporal(t *testing.T) {
	op1, n1 := negateTemporal("nft", 2)
	if op1 != "nmt" || n1 != 1 {
		t.Fatal("negateTemporal incorrect for nft")
	}
	op2, n2 := negateTemporal("nmt", 2)

	if op2 != "nft" || n2 != 3 {
		t.Fatal("negateTemporal incorrect for nmt")
	}
}

func TestValidOperator(t *testing.T) {
	c := NewCompiler()
	boolTy := &ast.Type{Type: "BOOL"}
	floatTy := &ast.Type{Type: "FLOAT"}
	test := &ast.InfixExpression{
		Left:     &ast.Identifier{InferredType: boolTy},
		Right:    &ast.Identifier{InferredType: boolTy},
		Operator: "&&"}
	test1 := &ast.InfixExpression{
		Left:     &ast.Identifier{InferredType: floatTy},
		Right:    &ast.Identifier{InferredType: floatTy},
		Operator: "&&"}

	if !c.validOperator(test, true) {
		t.Fatal("operator is valid but validOperator returned false")
	}

	if c.validOperator(test, false) {
		t.Fatal("operator is invalid but validOperator returned true")
	}

	if !c.validOperator(test1, true) {
		t.Fatal("operator is valid but validOperator returned false")
	}

	if !c.validOperator(test1, false) {
		t.Fatal("operator is valid but validOperator returned false")
	}
}
