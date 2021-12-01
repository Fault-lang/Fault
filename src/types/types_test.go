package types

import (
	"fault/ast"
	"fault/listener"
	"fault/parser"
	"testing"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func TestAddOK(t *testing.T) {
	test := `spec test1;
			const x = 2+2;
	`
	checker, err := prepTest(test)

	consts := checker.Constants["test1"]

	if err != nil {
		t.Fatalf("Type checking failed on valid expression. got=%s", err)
	}

	if consts["x"].(*ast.InfixExpression).InferredType.Type != "INT" {
		t.Fatalf("Constant x does not have an int type. got=%T", consts["x"])
	}

}

func TestTypeError(t *testing.T) {
	test := `spec test1;
			const x = 2+"2";
	`
	_, err := prepTest(test)
	if err == nil {
		t.Fatalf("Type checking failed to catch int string mismatch.")
	}
}

func TestComplex(t *testing.T) {
	test := `spec test1;
			const x = (2.1*8)+2.3/(5-2);
	`
	checker, err := prepTest(test)

	if err != nil {
		t.Fatalf("Type checking failed on valid expression. got=%s", err)
	}

	consts := checker.Constants["test1"]

	if consts["x"].(*ast.InfixExpression).InferredType.Type != "FLOAT" {
		t.Fatalf("Constant x does not have an float type. got=%T", consts["x"])
	}

	if consts["x"].(*ast.InfixExpression).InferredType.Scope != 10 {
		t.Fatalf("Constant x has the wrong scope. got=%d", consts["x"].(*ast.InfixExpression).InferredType.Scope)
	}

}

func TestScopes(t *testing.T) {
	test := `spec test1;
			const x = 2.2;
			const y = 2.0200;
			const z = uncertain(10, 5.2);
			const a = .005;
			const b = 103.40000;
	`
	checker, err := prepTest(test)

	if err != nil {
		t.Fatalf("Type checking failed on valid expression. got=%s", err)
	}

	consts := checker.Constants["test1"]

	if consts["x"].(*ast.FloatLiteral).InferredType.Scope != 10 {
		t.Fatalf("Constant x has the wrong scope. got=%d", consts["x"].(*ast.FloatLiteral).InferredType.Scope)
	}

	if consts["y"].(*ast.FloatLiteral).InferredType.Scope != 100 {
		t.Fatalf("Constant y has the wrong scope. got=%d", consts["y"].(*ast.FloatLiteral).InferredType.Scope)
	}

	if consts["z"].(*ast.Uncertain).InferredType.Scope != 0 {
		t.Fatalf("Constant z has the wrong scope. got=%d", consts["z"].(*ast.Uncertain).InferredType.Scope)
	}

	if consts["z"].(*ast.Uncertain).InferredType.Parameters[0].Scope != 1 {
		t.Fatalf("Constant z mean has the wrong scope. got=%d", consts["z"].(*ast.Uncertain).InferredType.Parameters[0].Scope)
	}

	if consts["z"].(*ast.Uncertain).InferredType.Parameters[1].Scope != 10 {
		t.Fatalf("Constant z sigma has the wrong scope. got=%d", consts["z"].(*ast.Uncertain).InferredType.Parameters[1].Scope)
	}

	if consts["a"].(*ast.FloatLiteral).InferredType.Scope != 1000 {
		t.Fatalf("Constant a has the wrong scope. got=%d", consts["a"].(*ast.FloatLiteral).InferredType.Scope)
	}

	if consts["b"].(*ast.FloatLiteral).InferredType.Scope != 10 {
		t.Fatalf("Constant b has the wrong scope. got=%d", consts["b"].(*ast.FloatLiteral).InferredType.Scope)
	}

}

func TestTypesInStruct(t *testing.T) {
	test := `spec test1;
			const a = 2.3;
			
			def foo = stock{
				foosh: 3,
				bar: "hello!",
				fizz: a,
			};

			def zoo = flow{
				con: new foo,
				rate: func{
					con.foosh + 2;
				},
			};
	`
	checker, err := prepTest(test)

	if err != nil {
		t.Fatalf("Type checking failed on valid expression. got=%s", err)
	}

	str := checker.SpecStructs["test1"]

	fooStock, ok := str["foo"]
	if !ok {
		t.Fatal("stock foo not stored in symbol table correctly.", str["foo"])
	}

	if fooStock["foosh"].(*ast.IntegerLiteral).InferredType.Type != "INT" {
		t.Fatalf("stock property not typed correctly. got=%s", fooStock["foosh"].(*ast.IntegerLiteral).InferredType.Type)
	}

	if fooStock["bar"].(*ast.StringLiteral).InferredType.Type != "STRING" {
		t.Fatalf("stock property not typed correctly. got=%s", fooStock["bar"].(*ast.StringLiteral).InferredType.Type)
	}

	if fooStock["fizz"].(*ast.Identifier).InferredType.Type != "FLOAT" {
		t.Fatalf("stock property not typed correctly. got=%s", fooStock["fizz"].(*ast.Identifier).InferredType.Type)
	}

	zooFlow, ok := str["zoo"]
	if !ok {
		t.Fatal("flow zoo not stored in symbol table correctly.")
	}

	if zooFlow["con"].(*ast.Instance).InferredType.Type != "STOCK" {
		t.Fatalf("flow property not typed correctly. got=%s", zooFlow["con"].(*ast.Instance).InferredType.Type)
	}

	if zooFlow["rate"].(*ast.BlockStatement).InferredType.Type != "INT" {
		t.Fatalf("flow property not typed correctly. got=%s", zooFlow["rate"].(*ast.IntegerLiteral).InferredType.Type)
	}
}

func TestInvalidAssert(t *testing.T) {
	test := `spec test1;
			const a = 2.3;

			assert a + 5;
	`
	_, err := prepTest(test)
	//sym := checker.SymbolTypes

	actual := "assert statement not testing a Boolean expression. got=INT"

	if err.Error() != actual {
		t.Fatalf("Type checking failed to catch invalid expression. got=%s", err)
	}

}

func TestValidAssert(t *testing.T) {
	test := `spec test1;
			const a = 2.3;
			
			assert a > 5;
	`
	_, err := prepTest(test)

	if err != nil {
		t.Fatalf("Type checking failed on a valid expression. got=%s", err)
	}

}

func TestPrefix(t *testing.T) {
	test := `spec test1;
			const a = !2.3;
			const b = -2.3;
	`
	checker, err := prepTest(test)

	if err != nil {
		t.Fatalf("Type checking failed on a valid expression. got=%s", err)
	}

	consts := checker.Constants["test1"]

	if consts["a"].(*ast.PrefixExpression).InferredType.Type != "BOOL" {
		t.Fatalf("Constant a does not have an boolean type. got=%s", consts["a"].(*ast.Boolean).InferredType.Type)
	}

	if consts["b"].(*ast.FloatLiteral).InferredType.Type != "FLOAT" {
		t.Fatalf("Constant a does not have an float type. got=%s", consts["b"].(*ast.FloatLiteral).InferredType.Type)
	}

}

func TestNatural(t *testing.T) {
	test := `spec test1;
			const a = natural(2);
	`
	checker, err := prepTest(test)

	if err != nil {
		t.Fatalf("Type checking failed on a valid expression. got=%s", err)
	}

	consts := checker.Constants["test1"]

	if consts["a"].(*ast.Natural).InferredType.Type != "NATURAL" {
		t.Fatalf("Constant a does not have an natural type. got=%s", consts["a"].(*ast.Natural).InferredType.Type)
	}

}

func TestBoolean(t *testing.T) {
	test := `spec test1;
			const a = true;
	`
	checker, err := prepTest(test)

	if err != nil {
		t.Fatalf("Type checking failed on a valid expression. got=%s", err)
	}

	consts := checker.Constants["test1"]

	if consts["a"].(*ast.Boolean).InferredType.Type != "BOOL" {
		t.Fatalf("Constant a does not have an Boolean type. got=%s", consts["a"].(*ast.Boolean).InferredType.Type)
	}

}

func TestString(t *testing.T) {
	test := `spec test1;
			const a = "Hello!";
	`
	checker, err := prepTest(test)

	if err != nil {
		t.Fatalf("Type checking failed on a valid expression. got=%s", err)
	}

	consts := checker.Constants["test1"]

	if consts["a"].(*ast.StringLiteral).InferredType.Type != "STRING" {
		t.Fatalf("Constant a does not have a string type. got=%s", consts["a"].(*ast.StringLiteral).InferredType.Type)
	}

}

// Infix, Prefix, ... what other types of expressions?
// Type check init matches expression type. init cannot be an uncertain. Uncertains are immutable... can only be declared as constants?
// check float + float returns a the larger scope
// "ignore x=5" <-- syntax to remove scenarios from the model checker?

func prepTest(test string) (*Checker, error) {
	is := antlr.NewInputStream(test)
	lexer := parser.NewFaultLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	p := parser.NewFaultParser(stream)
	l := &listener.FaultListener{}
	antlr.ParseTreeWalkerDefault.Walk(l, p.Spec())
	ty := &Checker{}
	err := ty.Check(l.AST)
	return ty, err
}
