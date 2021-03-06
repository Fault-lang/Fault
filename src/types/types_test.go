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

	spec := checker.AST
	sym := checker.Symbols

	if err != nil {
		t.Fatalf("Type checking failed on valid expression. got=%s", err)
	}

	if sym["x"] != "INT" {
		t.Fatalf("Constant x does not have an int type. got=%s", sym["x"])
	}

	if spec == nil {
		t.Fatalf("prepTest() returned nil")
	}

}

// func TestAddConvert(t *testing.T) {
// 	test := `spec test1;
// 			const x = 2+2.3;
// 	`
// 	checker, err := prepTest(test)

// 	spec := checker.AST
// 	sym := checker.Symbols

// 	if err != nil {
// 		t.Fatalf("Type checking failed on valid expression. got=%s", err)
// 	}

// 	if sym["x"] != "FLOAT" {
// 		t.Fatalf("Constant x does not have an float type. got=%s", sym["x"])
// 	}

// 	if spec == nil {
// 		t.Fatalf("prepTest() returned nil")
// 	}

// 	_, ok := spec.Statements[1].(*ast.ConstantStatement).Value.(*ast.InfixExpression).Left.(*ast.FloatLiteral)
// 	if !ok {
// 		t.Fatalf("Left node not converted to float type. got=%T", spec.Statements[1].(*ast.ConstantStatement).Value.(*ast.InfixExpression).Left)
// 	}

//}

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

	spec := checker.AST
	sym := checker.Symbols

	if err != nil {
		t.Fatalf("Type checking failed on valid expression. got=%s", err)
	}

	if sym["x"] != "FLOAT" {
		t.Fatalf("Constant x does not have an float type. got=%s", sym["x"])
	}

	if spec == nil {
		t.Fatalf("prepTest() returned nil")
	}

	_, ok := spec.Statements[1].(*ast.ConstantStatement).Value.(*ast.InfixExpression).Left.(*ast.InfixExpression)
	if !ok {
		t.Fatalf("Complex structure not maintained. got=%T", spec.Statements[1].(*ast.ConstantStatement).Value.(*ast.InfixExpression).Left)
	}

}

// func TestConvertWithVariable(t *testing.T) {
// 	test := `spec test1;
// 			const a = 2.3;
// 			const x = 2+a;
// 	`
// 	checker, err := prepTest(test)
// 	sym := checker.Symbols

// 	if err != nil {
// 		t.Fatalf("Type checking failed on valid expression. got=%s", err)
// 	}

// 	if sym["a"] != "FLOAT" {
// 		t.Fatalf("Constant a does not have an float type. got=%s", sym["a"])
// 	}

// 	if sym["x"] != "FLOAT" {
// 		t.Fatalf("Constant x does not have an float type. got=%s", sym["x"])
// 	}
// }

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
	sym := checker.Symbols

	if err != nil {
		t.Fatalf("Type checking failed on valid expression. got=%s", err)
	}

	if sym["a"] != "FLOAT" {
		t.Fatalf("Constant a does not have an float type. got=%s", sym["a"])
	}

	fooStock, ok := sym["foo"].(map[string]string)
	if !ok {
		t.Fatalf("stock foo not stored in symbol table correctly. got=%T", sym["foo"])
	}

	if fooStock["foosh"] != "INT" {
		t.Fatalf("stock property not typed correctly. got=%s", fooStock["foosh"])
	}

	if fooStock["bar"] != "STRING" {
		t.Fatalf("stock property not typed correctly. got=%s", fooStock["bar"])
	}

	if fooStock["fizz"] != "FLOAT" {
		t.Fatalf("stock property not typed correctly. got=%s", fooStock["fizz"])
	}

	zooFlow, ok := sym["zoo"].(map[string]string)
	if !ok {
		t.Fatalf("flow zoo not stored in symbol table correctly. got=%T", sym["zoo"])
	}

	if zooFlow["con"] != "STOCK" {
		t.Fatalf("flow property not typed correctly. got=%s", zooFlow["con"])
	}

	if zooFlow["rate"] != "INT" {
		t.Fatalf("flow property not typed correctly. got=%s", zooFlow["rate"])
	}
}

func TestInvalidAssert(t *testing.T) {
	test := `spec test1;
			const a = 2.3;
			
			assert a + 5;
	`
	_, err := prepTest(test)
	//sym := checker.Symbols

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

// Infix, Prefix, ... what other types of expressions?

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
