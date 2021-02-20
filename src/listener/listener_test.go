package listener

import (
	"fault/ast"
	"fault/parser"
	"testing"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func TestSpecDecl(t *testing.T) {
	test := `spec test1;`
	spec := prepTest(test)

	if spec == nil {
		t.Fatalf("prepTest() returned nil")
	}
	if len(spec.Statements) != 1 {
		t.Fatalf("spec.Statements does not contain 1 statement. got=%d", len(spec.Statements))
	}
	if spec.Statements[0].TokenLiteral() != "SPEC_DECL" {
		t.Fatalf("spec.Statement[0] is not SPEC_DECL. got=%s", spec.Statements[0].TokenLiteral())
	}
	if spec.Statements[0].(*ast.SpecDeclStatement).Name.String() != "test1" {
		t.Fatalf("Spec name is not test1. got=%s", spec.Statements[0].(*ast.SpecDeclStatement).Name.String())
	}
}

func TestConstDecl(t *testing.T) {
	test := `spec test1;
			 const x = 5;
			`
	spec := prepTest(test)

	if spec == nil {
		t.Fatalf("prepTest() returned nil")
	}
	if len(spec.Statements) != 2 {
		t.Fatalf("spec.Statements does not contain 2 statements. got=%d", len(spec.Statements))
	}
	if spec.Statements[1].TokenLiteral() != "CONST_DECL" {
		t.Fatalf("spec.Statement[1] is not CONST_DECL. got=%s", spec.Statements[1].TokenLiteral())
	}
	if spec.Statements[1].(*ast.ConstantStatement).Name.String() != "x" {
		t.Fatalf("Constant identifier is not x. got=%s", spec.Statements[1].(*ast.ConstantStatement).Name.String())
	}
}

func TestConstMultiDecl(t *testing.T) {
	test := `spec test1;
			 const x,y = 5;
			`
	spec := prepTest(test)

	if spec == nil {
		t.Fatalf("prepTest() returned nil")
	}
	if len(spec.Statements) != 3 {
		t.Fatalf("spec.Statements does not contain 3 statements. got=%d", len(spec.Statements))
	}
	if spec.Statements[1].TokenLiteral() != "CONST_DECL" {
		t.Fatalf("spec.Statement[1] is not CONST_DECL. got=%s", spec.Statements[1].TokenLiteral())
	}
	if spec.Statements[1].(*ast.ConstantStatement).Name.String() != "x" {
		t.Fatalf("Constant identifier is not x. got=%s", spec.Statements[1].(*ast.ConstantStatement).Name.String())
	}

	if spec.Statements[2].TokenLiteral() != "CONST_DECL" {
		t.Fatalf("spec.Statement[2] is not CONST_DECL. got=%s", spec.Statements[2].TokenLiteral())
	}
	if spec.Statements[2].(*ast.ConstantStatement).Name.String() != "y" {
		t.Fatalf("Constant identifier is not y. got=%s", spec.Statements[2].(*ast.ConstantStatement).Name.String())
	}
}

func TestConstMultiWExpressDecl(t *testing.T) {
	test := `spec test1;
			 const x = 1;
	         const y = 2 * (x + 1);;
			`
	spec := prepTest(test)

	if spec == nil {
		t.Fatalf("prepTest() returned nil")
	}
	if len(spec.Statements) != 3 {
		t.Fatalf("spec.Statements does not contain 3 statements. got=%d", len(spec.Statements))
	}
	if spec.Statements[1].TokenLiteral() != "CONST_DECL" {
		t.Fatalf("spec.Statement[1] is not CONST_DECL. got=%s", spec.Statements[1].TokenLiteral())
	}
	if spec.Statements[1].(*ast.ConstantStatement).Name.String() != "x" {
		t.Fatalf("Constant identifier is not x. got=%s", spec.Statements[1].(*ast.ConstantStatement).Name.String())
	}

	if spec.Statements[2].TokenLiteral() != "CONST_DECL" {
		t.Fatalf("spec.Statement[2] is not CONST_DECL. got=%s", spec.Statements[2].TokenLiteral())
	}
	if spec.Statements[2].(*ast.ConstantStatement).Name.String() != "y" {
		t.Fatalf("Constant identifier is not y. got=%s", spec.Statements[2].(*ast.ConstantStatement).Name.String())
	}

	_, ok := spec.Statements[2].(*ast.ConstantStatement).Value.(*ast.InfixExpression)
	if !ok {
		t.Fatalf("Constant value is not an infix expression. got=%T", spec.Statements[2].(*ast.ConstantStatement).Value)
	}
}

func TestStockDecl(t *testing.T) {
	test := `spec test1;
			 def foo = stock{
				value: 100,
			 };
			`
	spec := prepTest(test)

	if spec == nil {
		t.Fatalf("prepTest() returned nil")
	}
	if len(spec.Statements) != 2 {
		t.Fatalf("spec.Statements does not contain 2 statements. got=%d", len(spec.Statements))
	}
	if spec.Statements[1].TokenLiteral() != "=" {
		t.Fatalf("spec.Statement[1] is not ASSIGN. got='%s'", spec.Statements[1].TokenLiteral())
	}
	if spec.Statements[1].(*ast.DefStatement).Name.String() != "foo" {
		t.Fatalf("Stock identifier is not foo. got=%s", spec.Statements[1].(*ast.DefStatement).Name.String())
	}

	stock := spec.Statements[1].(*ast.DefStatement).Value.(*ast.StockLiteral).Pairs
	for _, v := range stock {
		f, ok := v.(*ast.FunctionLiteral)
		if !ok {
			t.Fatalf("Stock property is not wrapped in a function. got=%T", v)
		}
		_, ok = f.Body.Statements[0].(*ast.ExpressionStatement).Expression.(*ast.IntegerLiteral)
		if !ok {
			t.Fatalf("Function does not return integer. got=%T", f.Body.Statements[0].(*ast.ExpressionStatement).Expression)
		}
	}
}

func TestFlowDecl(t *testing.T) {
	test := `spec test1;
			 def foo = flow{
				bar: "here's a string",
			 };
			`
	spec := prepTest(test)

	if spec == nil {
		t.Fatalf("prepTest() returned nil")
	}
	if len(spec.Statements) != 2 {
		t.Fatalf("spec.Statements does not contain 2 statements. got=%d", len(spec.Statements))
	}
	if spec.Statements[1].TokenLiteral() != "=" {
		t.Fatalf("spec.Statement[1] is not ASSIGN. got='%s'", spec.Statements[1].TokenLiteral())
	}
	if spec.Statements[1].(*ast.DefStatement).Name.String() != "foo" {
		t.Fatalf("Flow identifier is not foo. got=%s", spec.Statements[1].(*ast.DefStatement).Name.String())
	}

	flow := spec.Statements[1].(*ast.DefStatement).Value.(*ast.FlowLiteral).Pairs
	for _, v := range flow {
		f, ok := v.(*ast.FunctionLiteral)
		if !ok {
			t.Fatalf("Flow property is not wrapped in a function. got=%T", v)
		}
		_, ok = f.Body.Statements[0].(*ast.ExpressionStatement).Expression.(*ast.StringLiteral)
		if !ok {
			t.Fatalf("Function does not return string. got=%T", f.Body.Statements[0].(*ast.ExpressionStatement).Expression)
		}
	}
}

func TestStockConnection(t *testing.T) {
	test := `spec test1;
			 def foo = flow{
				bar: new fizz,
			 };
			`
	spec := prepTest(test)
	flow := spec.Statements[1].(*ast.DefStatement).Value.(*ast.FlowLiteral).Pairs
	for _, v := range flow {
		f, ok := v.(*ast.InstanceExpression)
		if !ok {
			t.Fatalf("Property is not an instance. got=%T", v)
		}
		i, ok := f.Stock.(*ast.Identifier)
		if !ok {
			t.Fatalf("Function parameter is not an identifier. got=%T", f.Stock)
		}
		if i.Value != "fizz" {
			t.Fatalf("wrong element in call expression. got=%s", i.Value)
		}
	}
}

func TestFunctionBlock(t *testing.T) {
	test := `spec test1;
			 def foo = flow{
				bar: func{1+2;},
			 };
			`
	spec := prepTest(test)
	flow := spec.Statements[1].(*ast.DefStatement).Value.(*ast.FlowLiteral).Pairs
	for _, v := range flow {
		f, ok := v.(*ast.FunctionLiteral)
		if !ok {
			t.Fatalf("Property is not a function. got=%T", v)
		}
		if len(f.Body.Statements) != 1 {
			t.Fatalf("function BlockStatement does not contain 1 statement. got=%d", len(f.Body.Statements))
		}
		s, ok := f.Body.Statements[0].(*ast.ExpressionStatement)
		if !ok {
			t.Fatalf("Function body missing ExpressionStatement. got=%T", f.Body.Statements[0])
		}
		_, ok = s.Expression.(*ast.InfixExpression)
		if !ok {
			t.Fatalf("Function body missing InfixExpression. got=%T", s.Expression)
		}
	}
}

func TestPrefix(t *testing.T) {
	test := `spec test1;
			 def foo = flow{
				bar: func{!true;},
			 };
			`
	spec := prepTest(test)
	flow := spec.Statements[1].(*ast.DefStatement).Value.(*ast.FlowLiteral).Pairs
	for _, v := range flow {
		f, ok := v.(*ast.FunctionLiteral)
		if !ok {
			t.Fatalf("Property is not a function. got=%T", v)
		}
		if len(f.Body.Statements) != 1 {
			t.Fatalf("function BlockStatement does not contain 1 statement. got=%d", len(f.Body.Statements))
		}
		s, ok := f.Body.Statements[0].(*ast.ExpressionStatement)
		if !ok {
			t.Fatalf("Function body missing ExpressionStatement. got=%T", f.Body.Statements[0])
		}
		pre, ok := s.Expression.(*ast.PrefixExpression)
		if !ok {
			t.Fatalf("Function body missing PrefixExpression. got=%T", s.Expression)
		}
		_, ok = pre.Right.(*ast.Boolean)
		if !ok {
			t.Fatalf("Prefix does not contain a Boolean. got=%T", pre.Right)
		}
	}
}

func TestConditional(t *testing.T) {
	test := `spec test1;
			 def foo = flow{
				bar: func{!true;
					if(x){
						2+3;
					}else{
						1+1;
					}
				},
			 };
			`
	spec := prepTest(test)
	flow := spec.Statements[1].(*ast.DefStatement).Value.(*ast.FlowLiteral).Pairs
	for _, v := range flow {
		f, ok := v.(*ast.FunctionLiteral)
		if !ok {
			t.Fatalf("Property is not a function. got=%T", v)
		}
		if len(f.Body.Statements) != 2 {
			t.Fatalf("function BlockStatement does not contain 2 statements. got=%d", len(f.Body.Statements))
		}
		s, ok := f.Body.Statements[1].(*ast.ExpressionStatement)
		if !ok {
			t.Fatalf("Function body missing ExpressionStatement. got=%T", f.Body.Statements[0])
		}
		ife, ok := s.Expression.(*ast.IfExpression)
		if !ok {
			t.Fatalf("Function body missing IfExpression. got=%T", s.Expression)
		}
		_, ok = ife.Condition.(*ast.Identifier)
		if !ok {
			t.Fatalf("If Condition does not contain an Identifier. got=%T", ife.Condition)
		}
	}
}

func TestInit(t *testing.T) {
	test := `spec test1;
			 def foo = flow{
				bar: func{
					init 5;
				},
			 };
			`
	spec := prepTest(test)
	flow := spec.Statements[1].(*ast.DefStatement).Value.(*ast.FlowLiteral).Pairs
	for _, v := range flow {
		f, ok := v.(*ast.FunctionLiteral)
		if !ok {
			t.Fatalf("Property is not a function. got=%T", v)
		}
		if len(f.Body.Statements) != 1 {
			t.Fatalf("function BlockStatement does not contain 1 statements. got=%d", len(f.Body.Statements))
		}
		s, ok := f.Body.Statements[0].(*ast.ExpressionStatement)
		if !ok {
			t.Fatalf("Function body missing ExpressionStatement. got=%T", f.Body.Statements[0])
		}
		init, ok := s.Expression.(*ast.InitExpression)
		if !ok {
			t.Fatalf("Function body missing InitExpression. got=%T", s.Expression)
		}
		if init.Expression.(*ast.IntegerLiteral).Value != 5 {
			t.Fatalf("Init value is not 5. got=%d", init.Expression.(*ast.IntegerLiteral).Value)
		}
	}
}

func TestImport(t *testing.T) {
	test := `spec test1;
			 import "hello";
			`
	spec := prepTest(test)
	if spec == nil {
		t.Fatalf("prepTest() returned nil")
	}
	if len(spec.Statements) != 2 {
		t.Fatalf("spec.Statements does not contain 2 statements. got=%d", len(spec.Statements))
	}

	imp, ok := spec.Statements[1].(*ast.ImportStatement)
	if !ok {
		t.Fatalf("spec.Statement[1] is not an import statement. got=%T", spec.Statements[1])
	}
	if imp.Name.String() != "hello" {
		t.Fatalf("Import name is not hello. got=%s", imp.Name.String())
	}

	if imp.Path.String() != `"hello"` {
		t.Fatalf("Import path is not correct. got=%s", imp.Path.String())
	}
}

func TestImportWIdent(t *testing.T) {
	test := `spec test1;
			 import helloWorld "../../hello";
			`
	spec := prepTest(test)
	if spec == nil {
		t.Fatalf("prepTest() returned nil")
	}
	if len(spec.Statements) != 2 {
		t.Fatalf("spec.Statements does not contain 2 statements. got=%d", len(spec.Statements))
	}

	imp, ok := spec.Statements[1].(*ast.ImportStatement)
	if !ok {
		t.Fatalf("spec.Statement[1] is not an import statement. got=%T", spec.Statements[1])
	}
	if imp.Name.String() != "helloWorld" {
		t.Fatalf("Import name is not helloWorld. got=%s", imp.Name.String())
	}

	if imp.Path.String() != `"../../hello"` {
		t.Fatalf("Import path is not correct. got=%s", imp.Path.String())
	}
}

func TestMultiImport(t *testing.T) {
	test := `spec test1;
			 import("hello"
			         x "world");
			`
	spec := prepTest(test)
	if spec == nil {
		t.Fatalf("prepTest() returned nil")
	}
	if len(spec.Statements) != 3 {
		t.Fatalf("spec.Statements does not contain 3 statements. got=%d", len(spec.Statements))
	}

	imp, ok := spec.Statements[1].(*ast.ImportStatement)
	if !ok {
		t.Fatalf("spec.Statement[1] is not an import statement. got=%T", spec.Statements[1])
	}
	if imp.Name.String() != "hello" {
		t.Fatalf("Import name is not hello. got=%s", imp.Name.String())
	}

	if imp.Path.String() != `"hello"` {
		t.Fatalf("Import path is not correct. got=%s", imp.Path.String())
	}

	imp2, ok := spec.Statements[2].(*ast.ImportStatement)
	if !ok {
		t.Fatalf("spec.Statement[2] is not an import statement. got=%T", spec.Statements[2])
	}
	if imp2.Name.String() != "x" {
		t.Fatalf("Import name is not x. got=%s", imp2.Name.String())
	}

	if imp2.Path.String() != `"world"` {
		t.Fatalf("Import path is not correct. got=%s", imp2.Path.String())
	}
}

func TestForStatement(t *testing.T) {
	test := `spec test1;
			 for 5 run{};
			`
	spec := prepTest(test)
	if spec == nil {
		t.Fatalf("prepTest() returned nil")
	}
	if len(spec.Statements) != 2 {
		t.Fatalf("spec.Statements does not contain 2 statements. got=%d", len(spec.Statements))
	}
	forSt, ok := spec.Statements[1].(*ast.ForStatement)
	if !ok {
		t.Fatalf("spec.Statements[1] is not a ForStatement. got=%T", spec.Statements[1])
	}

	if forSt.Rounds.Value != 5 {
		t.Fatalf("ForStatement does not have 5 rounds. got=%d", forSt.Rounds.Value)
	}
}

func TestRunBlock(t *testing.T) {
	test := `spec test1;
			 for 5 run{
				d = new foo;
				d.fn;
			 };
			`
	spec := prepTest(test)
	if spec == nil {
		t.Fatalf("prepTest() returned nil")
	}
	if len(spec.Statements) != 2 {
		t.Fatalf("spec.Statements does not contain 2 statements. got=%d", len(spec.Statements))
	}
	forSt, ok := spec.Statements[1].(*ast.ForStatement)
	if !ok {
		t.Fatalf("spec.Statements[1] is not a ForStatement. got=%T", spec.Statements[1])
	}

	expr, ok := forSt.Body.Statements[1].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("forSt.Body.Statements[1] is not an ExpressionStatement. got=%T", forSt.Body.Statements[1])
	}

	id, ok := expr.Expression.(*ast.Identifier)
	if !ok {
		t.Fatalf("expr.Expression is not an Identifier. got=%T", expr.Expression)
	}

	if id.Value != "d.fn" {
		t.Fatalf("Identifier is not d.fn. got=%s", id.Value)
	}

}

func TestIncr(t *testing.T) {
	test := `spec test1;
			 for 5 run{
				i++;
			 };
			`
	spec := prepTest(test)
	forSt, ok := spec.Statements[1].(*ast.ForStatement)
	if !ok {
		t.Fatalf("spec.Statements[1] is not a ForStatement. got=%T", spec.Statements[1])
	}

	expr, ok := forSt.Body.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("forSt.Body.Statements[0] is not an ExpressionStatement. got=%T", forSt.Body.Statements[1])
	}

	infix, ok := expr.Expression.(*ast.InfixExpression)
	if !ok {
		t.Fatalf("expr.Expression is not an InfixExpression. got=%T", expr.Expression)
	}

	if infix.Right.String() != "1" {
		t.Fatalf("infix right side is not 1. got=%s", infix.Right.String())
	}

	if infix.Left.String() != "i" {
		t.Fatalf("infix left side is not i. got=%s", infix.Left.String())
	}

	if infix.Operator != "+" {
		t.Fatalf("infix operator is not +. got=%s", infix.Operator)
	}

}

func TestAssertion(t *testing.T) {
	test := `spec test1;
			 assert x > y;
			`
	spec := prepTest(test)
	assert, ok := spec.Statements[1].(*ast.AssertionStatement)
	if !ok {
		t.Fatalf("spec.Statements[1] is not an AssertionStatement. got=%T", spec.Statements[1])
	}

	expr, ok := assert.Expression.(*ast.InfixExpression)
	if !ok {
		t.Fatalf("assert.Expression is not an InfixExpression. got=%T", assert.Expression)
	}

	if expr.Right.String() != "y" {
		t.Fatalf("Right side isn't equal to y. got=%s", expr.Right.String())
	}

	if expr.Left.String() != "x" {
		t.Fatalf("Left side isn't equal to x. got=%s", expr.Left.String())
	}

	if expr.Operator != ">" {
		t.Fatalf("Operator isn't equal to >. got=%s", expr.Operator)
	}

}

func TestFaultAssign(t *testing.T) {
	test := `spec test1;
			 def foo = flow{
				bar: func{
					fizz -> buzz;
				},
			 };
			`
	spec := prepTest(test)
	flow := spec.Statements[1].(*ast.DefStatement).Value.(*ast.FlowLiteral).Pairs
	for _, v := range flow {
		f, ok := v.(*ast.FunctionLiteral)
		if !ok {
			t.Fatalf("Property is not a function. got=%T", v)
		}
		if len(f.Body.Statements) != 1 {
			t.Fatalf("function BlockStatement does not contain 1 statements. got=%d", len(f.Body.Statements))
		}
		s, ok := f.Body.Statements[0].(*ast.ExpressionStatement)
		if !ok {
			t.Fatalf("Function body missing ExpressionStatement. got=%T", f.Body.Statements[0])
		}
		assign, ok := s.Expression.(*ast.InfixExpression)
		if !ok {
			t.Fatalf("Function body missing InfixExpression. got=%T", s.Expression)
		}
		if assign.Left.String() != "fizz" {
			t.Fatalf("Left value is not fizz. got=%s", assign.Left.String())
		}
		if assign.Right.String() != "buzz" {
			t.Fatalf("Right value is not buzz. got=%s", assign.Right.String())
		}
		if assign.Operator != "->" {
			t.Fatalf("Operator is not ->. got=%s", assign.Operator)
		}
	}
}

func TestNil(t *testing.T) {
	test := `spec test1;
			 const a = nil;
			`
	spec := prepTest(test)
	con, ok := spec.Statements[1].(*ast.ConstantStatement)
	if !ok {
		t.Fatalf("spec.Statements[1] is not a ConstantStatement. got=%T", spec.Statements[1])
	}

	_, ok = con.Value.(*ast.Nil)
	if !ok {
		t.Fatalf("Constant is not set to nil. got=%T", con.Value)
	}
}

func TestAccessHistory(t *testing.T) {
	test := `spec test1;
			 const a = b[1][2];
			`
	spec := prepTest(test)
	con, ok := spec.Statements[1].(*ast.ConstantStatement)
	if !ok {
		t.Fatalf("spec.Statements[1] is not a ConstantStatement. got=%T", spec.Statements[1])
	}

	idx1, ok := con.Value.(*ast.IndexExpression)
	if !ok {
		t.Fatalf("Constant is not an IndexExpression. got=%T", con.Value)
	}

	idx2, ok := idx1.Left.(*ast.IndexExpression)
	if !ok {
		t.Fatalf("IndexExpression Left is not b[1]. got=%s", idx1.Left.String())
	}

	if idx2.Left.String() != "b" {
		t.Fatalf("IndexExpression Left is not b. got=%s", idx2.Left.String())
	}
}

func TestAccessHistory2(t *testing.T) {
	test := `spec test1;
			 const a = b[a[2]];
			`
	spec := prepTest(test)
	con, ok := spec.Statements[1].(*ast.ConstantStatement)
	if !ok {
		t.Fatalf("spec.Statements[1] is not a ConstantStatement. got=%T", spec.Statements[1])
	}

	idx1, ok := con.Value.(*ast.IndexExpression)
	if !ok {
		t.Fatalf("Constant is not an IndexExpression. got=%T", con.Value)
	}

	if idx1.Left.String() != "b" {
		t.Fatalf("IndexExpression Left is not b. got=%s", idx1.Left.String())
	}
}

func TestNegInt(t *testing.T) {
	test := `spec test1;
			 const a = -13;
			`
	spec := prepTest(test)
	con, ok := spec.Statements[1].(*ast.ConstantStatement)
	if !ok {
		t.Fatalf("spec.Statements[1] is not a ConstantStatement. got=%T", spec.Statements[1])
	}

	int1, ok := con.Value.(*ast.IntegerLiteral)
	if !ok {
		t.Fatalf("Constant is not an IntegerLiteral. got=%T", con.Value)
	}

	if int1.Value != -13 {
		t.Fatalf("Integer is not -13. got=%d", int1.Value)
	}
}

func TestFloat(t *testing.T) {
	test := `spec test1;
			 const a = 1.2;
			`
	spec := prepTest(test)
	con, ok := spec.Statements[1].(*ast.ConstantStatement)
	if !ok {
		t.Fatalf("spec.Statements[1] is not a ConstantStatement. got=%T", spec.Statements[1])
	}

	float1, ok := con.Value.(*ast.FloatLiteral)
	if !ok {
		t.Fatalf("Constant is not an FloatLiteral. got=%T", con.Value)
	}

	if float1.Value != 1.2 {
		t.Fatalf("Float is not 1.2. got=%f", float1.Value)
	}
}

func TestNegFloat(t *testing.T) {
	test := `spec test1;
			 const a = -1.2;
			`
	spec := prepTest(test)
	con, ok := spec.Statements[1].(*ast.ConstantStatement)
	if !ok {
		t.Fatalf("spec.Statements[1] is not a ConstantStatement. got=%T", spec.Statements[1])
	}

	float1, ok := con.Value.(*ast.FloatLiteral)
	if !ok {
		t.Fatalf("Constant is not an FloatLiteral. got=%T", con.Value)
	}

	if float1.Value != -1.2 {
		t.Fatalf("Float is not -1.2. got=%f", float1.Value)
	}
}

/* THINGS TO TEST:
- check String() in ast does not return Token Literal
- Could DefStatement be Infix Expressions
- Check grammar for ?*+ and handle as list of branches
- Check Position() is declared for all
- Fault types (non-negative, non-zero)
- How do Constants works in Go? (Barak)
*/

func prepTest(test string) *ast.Spec {
	is := antlr.NewInputStream(test)
	lexer := parser.NewFaultLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	p := parser.NewFaultParser(stream)
	listener := &FaultListener{}
	antlr.ParseTreeWalkerDefault.Walk(listener, p.Spec())
	return listener.AST
}
