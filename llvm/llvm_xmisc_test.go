package llvm

// "xmisc" is a little hack to keep the order of test files
// from screwing up the block numbering on LLVM IR (and thus
// causing a bunch of tests to fail.

import (
	"fault/ast"
	"fault/preprocess"
	"strings"
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

func TestProcessSpecInvalidRootNode(t *testing.T) {
	c := NewCompiler()
	err := c.Compile(&ast.ConstantStatement{})
	if err == nil {
		t.Fatal("expected error for invalid root node, got nil")
	}
	if !strings.Contains(err.Error(), "spec file improperly formatted. Root node is") {
		t.Fatalf("unexpected error message. got=%q", err.Error())
	}
}

func TestProcessSpecMissingSpecDecl(t *testing.T) {
	c := NewCompiler()
	c.isTesting = true
	spec := &ast.Spec{
		Statements: []ast.Statement{
			&ast.ConstantStatement{},
		},
	}
	err := c.Compile(spec)
	if err == nil {
		t.Fatal("expected error for missing spec declaration, got nil")
	}
	if !strings.Contains(err.Error(), "spec file improperly formatted. Missing spec declaration, got") {
		t.Fatalf("unexpected error message. got=%q", err.Error())
	}
}

func TestProcessSpecFetchInstanceStrMapError(t *testing.T) {
	c := NewCompiler()
	c.isTesting = true
	sr := preprocess.NewSpecRecord()
	sr.SpecName = "specname"
	sr.Order = [][]string{{"STOCK", "instancename"}}
	c.specStructs["specname"] = sr

	spec := &ast.Spec{
		Statements: []ast.Statement{
			&ast.SysDeclStatement{
				Name: &ast.Identifier{Value: "sysname", ProcessedName: []string{"sysname"}},
			},
			&ast.DefStatement{
				Name: &ast.Identifier{
					Value:         "specname_instancename",
					ProcessedName: []string{"specname", "instancename"},
				},
				Value: &ast.StructInstance{
					Parent: []string{"sysname", "parentname"},
				},
			},
		},
	}
	err := c.Compile(spec)
	if err == nil {
		t.Fatal("expected error when FetchInstanceStrMap fails, got nil")
	}
	if !strings.Contains(err.Error(), "no stock found with name instancename") {
		t.Fatalf("unexpected error message. got=%q", err.Error())
	}
}

func TestProcessSpecFetchComponentError(t *testing.T) {
	c := NewCompiler()
	c.isTesting = true
	sr := preprocess.NewSpecRecord()
	sr.SpecName = "specname"
	c.specStructs["specname"] = sr

	spec := &ast.Spec{
		Statements: []ast.Statement{
			&ast.SysDeclStatement{
				Name: &ast.Identifier{Value: "sysname", ProcessedName: []string{"sysname"}},
			},
			&ast.DefStatement{
				Name: &ast.Identifier{
					Value:         "specname_componentname",
					ProcessedName: []string{"specname", "componentname"},
				},
				Value: &ast.ComponentLiteral{
					ProcessedName: []string{"specname", "componentname"},
				},
			},
		},
	}
	err := c.Compile(spec)
	if err == nil {
		t.Fatal("expected error when FetchComponent fails, got nil")
	}
	if !strings.Contains(err.Error(), "no component found with name componentname") {
		t.Fatalf("unexpected error message. got=%q", err.Error())
	}
}
