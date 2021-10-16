package types

import (
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

	sym := checker.SymbolTypes["test1"]

	if err != nil {
		t.Fatalf("Type checking failed on valid expression. got=%s", err)
	}

	if sym["x"].(*Type).Type != "INT" {
		t.Fatalf("Constant x does not have an int type. got=%s", sym["x"].(*Type).Type)
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

	sym := checker.SymbolTypes["test1"]

	if err != nil {
		t.Fatalf("Type checking failed on valid expression. got=%s", err)
	}

	if sym["x"].(*Type).Type != "FLOAT" {
		t.Fatalf("Constant x does not have an float type. got=%s", sym["x"].(*Type).Type)
	}

	if sym["x"].(*Type).Scope != 10 {
		t.Fatalf("Constant x has the wrong scope. got=%d", sym["x"].(*Type).Scope)
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

	sym := checker.SymbolTypes["test1"]

	if err != nil {
		t.Fatalf("Type checking failed on valid expression. got=%s", err)
	}

	if sym["x"].(*Type).Scope != 10 {
		t.Fatalf("Constant x has the wrong scope. got=%d", sym["x"].(*Type).Scope)
	}

	if sym["y"].(*Type).Scope != 100 {
		t.Fatalf("Constant y has the wrong scope. got=%d", sym["y"].(*Type).Scope)
	}

	if sym["z"].(*Type).Scope != 0 {
		t.Fatalf("Constant z has the wrong scope. got=%d", sym["z"].(*Type).Scope)
	}

	if sym["z"].(*Type).Parameters[0].Scope != 1 {
		t.Fatalf("Constant z mean has the wrong scope. got=%d", sym["z"].(*Type).Parameters[0].Scope)
	}

	if sym["z"].(*Type).Parameters[1].Scope != 10 {
		t.Fatalf("Constant z sigma has the wrong scope. got=%d", sym["z"].(*Type).Parameters[1].Scope)
	}

	if sym["a"].(*Type).Scope != 1000 {
		t.Fatalf("Constant a has the wrong scope. got=%d", sym["a"].(*Type).Scope)
	}

	if sym["b"].(*Type).Scope != 10 {
		t.Fatalf("Constant b has the wrong scope. got=%d", sym["b"].(*Type).Scope)
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
	sym := checker.SymbolTypes["test1"]

	if err != nil {
		t.Fatalf("Type checking failed on valid expression. got=%s", err)
	}

	if sym["a"].(*Type).Type != "FLOAT" {
		t.Fatalf("Constant a does not have an float type. got=%s", sym["a"].(*Type).Type)
	}

	fooStock, ok := sym["foo"].(map[string]*Type)
	if !ok {
		t.Fatalf("stock foo not stored in symbol table correctly. got=%T", sym["foo"])
	}

	if fooStock["foosh"].Type != "INT" {
		t.Fatalf("stock property not typed correctly. got=%s", fooStock["foosh"].Type)
	}

	if fooStock["bar"].Type != "STRING" {
		t.Fatalf("stock property not typed correctly. got=%s", fooStock["bar"].Type)
	}

	if fooStock["fizz"].Type != "FLOAT" {
		t.Fatalf("stock property not typed correctly. got=%s", fooStock["fizz"].Type)
	}

	zooFlow, ok := sym["zoo"].(map[string]*Type)
	if !ok {
		t.Fatalf("flow zoo not stored in symbol table correctly. got=%T", sym["zoo"])
	}

	if zooFlow["con"].Type != "STOCK" {
		t.Fatalf("flow property not typed correctly. got=%s", zooFlow["con"].Type)
	}

	if zooFlow["rate"].Type != "INT" {
		t.Fatalf("flow property not typed correctly. got=%s", zooFlow["rate"].Type)
	}
}

func TestInvalidAssert(t *testing.T) {
	test := `spec test1;
			const a = 2.3;

			assert a + 5;
	`
	_, err := prepTest(test)
	//sym := checker.SymbolTypes

	actual := "Assert statement not testing a Boolean expression. got=FLOAT"

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
	sym := checker.SymbolTypes["test1"]

	if err != nil {
		t.Fatalf("Type checking failed on a valid expression. got=%s", err)
	}

	if sym["a"].(*Type).Type != "BOOL" {
		t.Fatalf("Constant a does not have an boolean type. got=%s", sym["a"].(*Type).Type)
	}

	if sym["b"].(*Type).Type != "FLOAT" {
		t.Fatalf("Constant a does not have an float type. got=%s", sym["b"].(*Type).Type)
	}

}

func TestNatural(t *testing.T) {
	test := `spec test1;
			const a = natural(2);
	`
	checker, err := prepTest(test)
	sym := checker.SymbolTypes["test1"]

	if err != nil {
		t.Fatalf("Type checking failed on a valid expression. got=%s", err)
	}

	if sym["a"].(*Type).Type != "NATURAL" {
		t.Fatalf("Constant a does not have an natural type. got=%s", sym["a"].(*Type).Type)
	}

}

func TestBoolean(t *testing.T) {
	test := `spec test1;
			const a = true;
	`
	checker, err := prepTest(test)
	sym := checker.SymbolTypes["test1"]

	if err != nil {
		t.Fatalf("Type checking failed on a valid expression. got=%s", err)
	}

	if sym["a"].(*Type).Type != "BOOL" {
		t.Fatalf("Constant a does not have an Boolean type. got=%s", sym["a"].(*Type).Type)
	}

}

func TestString(t *testing.T) {
	test := `spec test1;
			const a = "Hello!";
	`
	checker, err := prepTest(test)
	sym := checker.SymbolTypes["test1"]

	if err != nil {
		t.Fatalf("Type checking failed on a valid expression. got=%s", err)
	}

	if sym["a"].(*Type).Type != "STRING" {
		t.Fatalf("Constant a does not have a string type. got=%s", sym["a"].(*Type).Type)
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
