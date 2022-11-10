package listener

import (
	"fault/ast"
	"fault/parser"
	"testing"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func TestSpecDecl(t *testing.T) {
	test := `spec test1;`
	_, spec := prepTest(test, nil)

	if spec == nil {
		t.Fatalf("prepTest() returned nil")
	}
	if len(spec.Statements) != 1 {
		t.Fatalf("spec.Statements does not contain 1 statement. got=%d", len(spec.Statements))
	}
	if spec.Statements[0].TokenLiteral() != "SPEC_DECL" {
		t.Fatalf("spec.Statement[0] is not SPEC_DECL. got=%s", spec.Statements[0].TokenLiteral())
	}
	if spec.Statements[0].(*ast.SpecDeclStatement).Name.Value != "test1" {
		t.Fatalf("Spec name is not test1. got=%s", spec.Statements[0].(*ast.SpecDeclStatement).Name.Value)
	}
}

func TestConstDecl(t *testing.T) {
	test := `spec test1;
			 const x = 5;
			`
	_, spec := prepTest(test, nil)

	if spec == nil {
		t.Fatalf("prepTest() returned nil")
	}
	if len(spec.Statements) != 2 {
		t.Fatalf("spec.Statements does not contain 2 statements. got=%d", len(spec.Statements))
	}
	if spec.Statements[1].TokenLiteral() != "CONST_DECL" {
		t.Fatalf("spec.Statement[1] is not CONST_DECL. got=%s", spec.Statements[1].TokenLiteral())
	}
	if spec.Statements[1].(*ast.ConstantStatement).Name.Value != "x" {
		t.Fatalf("Constant identifier is not x. got=%s", spec.Statements[1].(*ast.ConstantStatement).Name.Value)
	}
}

func TestBoolean(t *testing.T) {
	test := `spec test1;
			 const x = false;
			`
	_, spec := prepTest(test, nil)

	if spec == nil {
		t.Fatalf("prepTest() returned nil")
	}
	if len(spec.Statements) != 2 {
		t.Fatalf("spec.Statements does not contain 2 statements. got=%d", len(spec.Statements))
	}
	if spec.Statements[1].TokenLiteral() != "CONST_DECL" {
		t.Fatalf("spec.Statement[1] is not CONST_DECL. got=%s", spec.Statements[1].TokenLiteral())
	}
	if _, ok := spec.Statements[1].(*ast.ConstantStatement).Value.(*ast.Boolean); !ok {
		t.Fatalf("Constant is not a Boolean. got=%T", spec.Statements[1].(*ast.ConstantStatement).Value)
	}
}

func TestConstMultiDecl(t *testing.T) {
	test := `spec test1;
			 const x,y = 5;
			`
	_, spec := prepTest(test, nil)

	if spec == nil {
		t.Fatalf("prepTest() returned nil")
	}
	if len(spec.Statements) != 3 {
		t.Fatalf("spec.Statements does not contain 3 statements. got=%d", len(spec.Statements))
	}
	if spec.Statements[1].TokenLiteral() != "CONST_DECL" {
		t.Fatalf("spec.Statement[1] is not CONST_DECL. got=%s", spec.Statements[1].TokenLiteral())
	}
	if spec.Statements[1].(*ast.ConstantStatement).Name.Value != "x" {
		t.Fatalf("Constant identifier is not x. got=%s", spec.Statements[1].(*ast.ConstantStatement).Name.Value)
	}

	if spec.Statements[2].TokenLiteral() != "CONST_DECL" {
		t.Fatalf("spec.Statement[2] is not CONST_DECL. got=%s", spec.Statements[2].TokenLiteral())
	}
	if spec.Statements[2].(*ast.ConstantStatement).Name.Value != "y" {
		t.Fatalf("Constant identifier is not y. got=%s", spec.Statements[2].(*ast.ConstantStatement).Name.Value)
	}
}

func TestConstMultiWExpressDecl(t *testing.T) {
	test := `spec test1;
			 const x = 1;
	         const y = 2 * (x + 1);;
			`
	_, spec := prepTest(test, nil)

	if spec == nil {
		t.Fatalf("prepTest() returned nil")
	}
	if len(spec.Statements) != 3 {
		t.Fatalf("spec.Statements does not contain 3 statements. got=%d", len(spec.Statements))
	}
	if spec.Statements[1].TokenLiteral() != "CONST_DECL" {
		t.Fatalf("spec.Statement[1] is not CONST_DECL. got=%s", spec.Statements[1].TokenLiteral())
	}
	if spec.Statements[1].(*ast.ConstantStatement).Name.Value != "x" {
		t.Fatalf("Constant identifier is not x. got=%s", spec.Statements[1].(*ast.ConstantStatement).Name.Value)
	}

	if spec.Statements[2].TokenLiteral() != "CONST_DECL" {
		t.Fatalf("spec.Statement[2] is not CONST_DECL. got=%s", spec.Statements[2].TokenLiteral())
	}
	if spec.Statements[2].(*ast.ConstantStatement).Name.Value != "y" {
		t.Fatalf("Constant identifier is not y. got=%s", spec.Statements[2].(*ast.ConstantStatement).Name.Value)
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
				test: buzz,
				call: test2.lol,
			 };
			`
	_, spec := prepTest(test, nil)

	if spec == nil {
		t.Fatalf("prepTest() returned nil")
	}
	if len(spec.Statements) != 2 {
		t.Fatalf("spec.Statements does not contain 2 statements. got=%d", len(spec.Statements))
	}
	if spec.Statements[1].TokenLiteral() != "STOCK" {
		t.Fatalf("spec.Statement[1] is not STOCK. got='%s'", spec.Statements[1].TokenLiteral())
	}
	if spec.Statements[1].(*ast.DefStatement).Name.Value != "foo" {
		t.Fatalf("Stock identifier is not foo. got=%s", spec.Statements[1].(*ast.DefStatement).Name.Value)
	}

	stock := spec.Statements[1].(*ast.DefStatement).Value.(*ast.StockLiteral).Pairs

	for k, v := range stock {
		if k.Value == "value" {
			_, ok := v.(*ast.IntegerLiteral)
			if !ok {
				t.Fatalf("Property is not an integer. got=%T", v)
			}
		} else if k.Value == "test" {
			buzz, ok := v.(*ast.Identifier)
			if !ok {
				t.Fatalf("Property is not an indentifier. got=%T", v)
			}
			if buzz.Value != "buzz" {
				t.Fatalf("Property is incorrect. got=%s want=buzz", buzz.Value)
			}
		} else if k.Value == "call" {
			call, ok := v.(*ast.ParameterCall)
			if !ok {
				t.Fatalf("Property is not an call. got=%T", v)
			}

			if call.Value[0] != "test2" {
				t.Fatalf("Property is incorrect. got=%s want=test2", call.Value[0])
			}

		}
	}
	if len(stock) != 3 {
		t.Fatalf("a key is missing from stock")
	}

}

func TestStockDeclFloat(t *testing.T) {
	test := `spec test1;
			 def foo = stock{
				value: 10.0,
			 };
			`
	_, spec := prepTest(test, nil)

	if spec == nil {
		t.Fatalf("prepTest() returned nil")
	}
	if len(spec.Statements) != 2 {
		t.Fatalf("spec.Statements does not contain 2 statements. got=%d", len(spec.Statements))
	}
	if spec.Statements[1].TokenLiteral() != "STOCK" {
		t.Fatalf("spec.Statement[1] is not STOCK. got='%s'", spec.Statements[1].TokenLiteral())
	}
	if spec.Statements[1].(*ast.DefStatement).Name.Value != "foo" {
		t.Fatalf("Stock identifier is not foo. got=%s", spec.Statements[1].(*ast.DefStatement).Name.Value)
	}

	stock := spec.Statements[1].(*ast.DefStatement).Value.(*ast.StockLiteral).Pairs
	for _, v := range stock {
		_, ok := v.(*ast.FloatLiteral)
		if !ok {
			t.Fatalf("Property is not a float. got=%T", v)
		}
	}

	if len(stock) != 1 {
		t.Fatalf("key 'value' is missing from stock")
	}
}

func TestFlowDecl(t *testing.T) {
	test := `spec test1;
			 def foo = flow{
				bar: "here's a string",
			 };
			`
	_, spec := prepTest(test, nil)

	if spec == nil {
		t.Fatalf("prepTest() returned nil")
	}
	if len(spec.Statements) != 2 {
		t.Fatalf("spec.Statements does not contain 2 statements. got=%d", len(spec.Statements))
	}
	if spec.Statements[1].TokenLiteral() != "FLOW" {
		t.Fatalf("spec.Statement[1] is not FLOW. got='%s'", spec.Statements[1].TokenLiteral())
	}
	if spec.Statements[1].(*ast.DefStatement).Name.Value != "foo" {
		t.Fatalf("Flow identifier is not foo. got=%s", spec.Statements[1].(*ast.DefStatement).Name.Value)
	}

	flow := spec.Statements[1].(*ast.DefStatement).Value.(*ast.FlowLiteral).Pairs
	for _, v := range flow {
		_, ok := v.(*ast.StringLiteral)
		if !ok {
			t.Fatalf("Property is not a string. got=%T", v)
		}
	}

	if len(flow) != 1 {
		t.Fatalf("key 'bar' is missing from flow")
	}
}

func TestStockConnection(t *testing.T) {
	test := `spec test1;
			 def foo = flow{
				bar: new fizz,
			 };
			`
	_, spec := prepTest(test, nil)
	flow := spec.Statements[1].(*ast.DefStatement).Value.(*ast.FlowLiteral).Pairs
	for _, v := range flow {
		f, ok := v.(*ast.Instance)
		if !ok {
			t.Fatalf("Property is not an instance. got=%T", v)
		}
		if f.Value.Value != "fizz" {
			t.Fatalf("wrong element in call expression. got=%s", f.Value.Value)
		}

		if f.Name != "bar" {
			t.Fatalf("wrong name in call expression. got=%s", f.Name)
		}

		if f.Value.Spec != "test1" {
			t.Fatalf("wrong spec for call expression. got=%s", f.Value.Spec)
		}
	}
}

func TestStructOrder(t *testing.T) {
	test := `spec test1;
			 def foo = stock{
				bar: 2.0,
				bash: 5,
				barg: true,
			 };

			 def zoo = flow{
				st: new foo
			 }
			`
	_, spec := prepTest(test, nil)
	stock := spec.Statements[1].(*ast.DefStatement).Value.(*ast.StockLiteral)
	if len(stock.Order) != 3 {
		t.Fatalf("Struct has incorrect number of properties in order. got=%d", len(stock.Order))
	}

	if stock.Order[0] != "bar" || stock.Order[1] != "bash" || stock.Order[2] != "barg" {
		t.Fatalf("Struct order is wrong. got=%s", stock.Order)
	}

	flow := spec.Statements[2].(*ast.DefStatement).Value.(*ast.FlowLiteral).Pairs
	for _, v := range flow {
		f, ok := v.(*ast.Instance)
		if !ok {
			t.Fatalf("Property is not an instance. got=%T", v)
		}
		if len(f.Order) != 3 {
			t.Fatalf("Instance has incorrect number of properties in order. got=%d", len(f.Order))
		}

		if f.Order[0] != "bar" || f.Order[1] != "bash" || f.Order[2] != "barg" {
			t.Fatalf("Instance order is wrong. got=%s", f.Order)
		}
	}
}

func TestStockImport(t *testing.T) {
	test := `spec test1;
			 def foo = flow{
				bar: new test2.fizz,
			 };
			`
	_, spec := prepTest(test, nil)
	flow := spec.Statements[1].(*ast.DefStatement).Value.(*ast.FlowLiteral).Pairs
	for _, v := range flow {
		f, ok := v.(*ast.Instance)
		if !ok {
			t.Fatalf("Property is not an instance. got=%T", v)
		}
		if f.Value.Value != "fizz" {
			t.Fatalf("wrong element in call expression. got=%s", f.Value.Value)
		}

		if f.Value.Spec != "test2" {
			t.Fatalf("wrong spec for call expression. got=%s", f.Value.Spec)
		}
	}
}

func TestFunctionBlock(t *testing.T) {
	test := `spec test1;
			 def foo = flow{
				bar: func{1+2;},
			 };
			`
	_, spec := prepTest(test, nil)
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

func TestThis(t *testing.T) {
	test := `spec test1;
			 def foo = flow{
				bar: func{1+this;},
			 };
			`
	_, spec := prepTest(test, nil)
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
		infix, ok := s.Expression.(*ast.InfixExpression)
		if !ok {
			t.Fatalf("Function body missing InfixExpression. got=%T", s.Expression)
		}
		this, ok := infix.Right.(*ast.This)
		if !ok {
			t.Fatalf("right infix is not This. got=%T", infix.Right)
		}
		if len(this.Value) != 2 {
			t.Fatalf("this has the wrong scope. got=%s", this.Value)
		}

		if this.Value[0] != "foo" {
			t.Fatalf("this has the wrong scope. got=%s", this.Value)
		}

		if this.Value[1] != "bar" {
			t.Fatalf("this has the wrong scope. got=%s", this.Value)
		}
	}
}

func TestClock(t *testing.T) {
	test := `spec test1;
			 def foo = flow{
				bar: func{1+clock;},
			 };
			`
	_, spec := prepTest(test, nil)
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
		infix, ok := s.Expression.(*ast.InfixExpression)
		if !ok {
			t.Fatalf("Function body missing InfixExpression. got=%T", s.Expression)
		}
		clock, ok := infix.Right.(*ast.Clock)
		if !ok {
			t.Fatalf("right infix is not This. got=%T", infix.Right)
		}
		if clock.Value != "foo.bar" {
			t.Fatalf("this has the wrong scope. got=%s", clock.Value)
		}

	}
}

func TestClockRun(t *testing.T) {
	test := `spec test1;
			 def foo = flow{
				bar: func{1+this;},
			 };
			 for 1 run {
				 clock;
			 }
			`
	_, spec := prepTest(test, nil)
	run := spec.Statements[2].(*ast.ForStatement).Body.Statements
	clock, ok := run[0].(*ast.ExpressionStatement).Expression.(*ast.Clock)
	if !ok {
		t.Fatal("clock missing from run block")
	}

	if clock.Value != "" {
		t.Fatal("scope failed to reset")
	}
}

func TestPrefix(t *testing.T) {
	test := `spec test1;
			 def foo = flow{
				bar: func{!true;
						-a;
						},
			 };
			`
	_, spec := prepTest(test, nil)
	flow := spec.Statements[1].(*ast.DefStatement).Value.(*ast.FlowLiteral).Pairs
	for _, v := range flow {
		f, ok := v.(*ast.FunctionLiteral)
		if !ok {
			t.Fatalf("Property is not a function. got=%T", v)
		}
		if len(f.Body.Statements) != 2 {
			t.Fatalf("function BlockStatement does not contain 2 statements. got=%d", len(f.Body.Statements))
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

		s2, ok := f.Body.Statements[1].(*ast.ExpressionStatement)
		if !ok {
			t.Fatalf("Function body missing ExpressionStatement. got=%T", f.Body.Statements[1])
		}
		pre2, ok := s2.Expression.(*ast.PrefixExpression)
		if !ok {
			t.Fatalf("Function body missing PrefixExpression. got=%T", s2.Expression)
		}
		_, ok = pre2.Right.(*ast.Identifier)
		if !ok {
			t.Fatalf("Prefix does not contain an Identifier. got=%T", pre2.Right)
		}
	}
}
func TestSimpleConditional(t *testing.T) {
	test := `spec test1;
			 def foo = flow{
				bar: func{!true;
					if(x){
						2+3;
					}
				},
			 };
			`
	_, spec := prepTest(test, nil)
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
		if len(ife.Consequence.Statements) == 0 {
			t.Fatalf("If Condition does not contain an consequence clause. got=%s", ife.Consequence)
		}
		if ife.Alternative != nil {
			t.Fatalf("If Condition contain an alternative clause when it shouldn't. got=%s", ife.Alternative)
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
	_, spec := prepTest(test, nil)
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
		if len(ife.Consequence.Statements) == 0 {
			t.Fatalf("If Condition does not contain an consequence clause. got=%s", ife.Consequence)
		}
		if ife.Alternative == nil {
			t.Fatalf("If Condition does not contain an alternative clause. got=%s", ife.String())
		}
	}
}

func TestElseIf(t *testing.T) {
	test := `spec test1;
			 def foo = flow{
				bar: func{!true;
					if(x){
						2+3;
					}else if(y){
						1+1;
					}
				},
			 };
			`
	_, spec := prepTest(test, nil)
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
		if len(ife.Consequence.Statements) == 0 {
			t.Fatalf("If Condition does not contain an consequence clause. got=%s", ife.Consequence)
		}
		if ife.Elif == nil {
			t.Fatalf("If Condition does not contain an else if clause. got=%s", ife.String())
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
	_, spec := prepTest(test, nil)
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
	test := `system test1;
			 import "hello";
			`
	_, spec := prepSysTest(test, nil)
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
	if imp.Name.Value != "hello" {
		t.Fatalf("Import name is not hello. got=%s", imp.Name.Value)
	}

	if imp.Path.String() != `"hello"` {
		t.Fatalf("Import path is not correct. got=%s", imp.Path.String())
	}
}

func TestImportWIdent(t *testing.T) {
	test := `system test1;
			 import helloWorld "../../hello";
			`
	_, spec := prepSysTest(test, nil)
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
	if imp.Name.Value != "helloWorld" {
		t.Fatalf("Import name is not helloWorld. got=%s", imp.Name.Value)
	}

	if imp.Path.String() != `"../../hello"` {
		t.Fatalf("Import path is not correct. got=%s", imp.Path.String())
	}
}

func TestMultiImport(t *testing.T) {
	test := `system test1;
			 import("hello"
			         x "world");
			`
	_, spec := prepSysTest(test, nil)
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
	if imp.Name.Value != "hello" {
		t.Fatalf("Import name is not hello. got=%s", imp.Name.Value)
	}

	if imp.Path.String() != `"hello"` {
		t.Fatalf("Import path is not correct. got=%s", imp.Path.String())
	}

	imp2, ok := spec.Statements[2].(*ast.ImportStatement)
	if !ok {
		t.Fatalf("spec.Statement[2] is not an import statement. got=%T", spec.Statements[2])
	}
	if imp2.Name.Value != "x" {
		t.Fatalf("Import name is not x. got=%s", imp2.Name.Value)
	}

	if imp2.Path.String() != `"world"` {
		t.Fatalf("Import path is not correct. got=%s", imp2.Path.String())
	}
}

func TestForStatement(t *testing.T) {
	test := `spec test1;
			 for 5 run{};
			`
	_, spec := prepTest(test, nil)
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
	_, spec := prepTest(test, nil)
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

	inst, ok := forSt.Body.Statements[0].(*ast.ExpressionStatement).Expression.(*ast.Instance)
	if !ok {
		t.Fatalf("forSt.Body.Statements[1] is not an ParallelFunctions. got=%T", forSt.Body.Statements[0].(*ast.ExpressionStatement).Expression)
	}
	if inst.Value.Spec != "test1" {
		t.Fatalf("instance has the wrong name. got=%s", inst.Value.Spec)
	}
	if inst.Value.Value != "foo" {
		t.Fatalf("instance has the wrong value. got=%s", inst.Value.Value)
	}
	if inst.Name != "d" {
		t.Fatalf("instance has the wrong name. got=%s", inst.Name)
	}

	expr, ok := forSt.Body.Statements[1].(*ast.ParallelFunctions)
	if !ok {
		t.Fatalf("forSt.Body.Statements[1] is not an ParallelFunctions. got=%T", forSt.Body.Statements[1])
	}

	id, ok := expr.Expressions[0].(*ast.ParameterCall)
	if !ok {
		t.Fatalf("expr.Expression is not an function call. got=%T", expr.Expressions[0])
	}

	if id.Value[0] != "d" && id.Value[0] != "fn" {
		t.Fatalf("Identifier is not d.fn. got=%s", id.Value)
	}

}

func TestSkipRun(t *testing.T) {
	test := `spec test1;
			 const a = 5;
			 for 5 run{
				d = new foo;
				d.fn;
			 };
			 `
	flags := make(map[string]bool)
	flags["skipRun"] = true

	_, spec := prepTest(test, flags)
	if spec == nil {
		t.Fatalf("prepTest() returned nil")
	}
	if len(spec.Statements) != 2 {
		t.Fatalf("spec.Statements does not contain 2 statements. got=%d", len(spec.Statements))
	}
	_, ok := spec.Statements[1].(*ast.ConstantStatement)
	if !ok {
		t.Fatalf("spec.Statements[1] is not a ConstantStatement. got=%T", spec.Statements[1])
	}

}

func TestRunInit(t *testing.T) {
	test := `spec test1;
			 for 5 run{
				d = new test2.foo;
			 };
			`
	_, spec := prepTest(test, nil)
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

	inst, ok := forSt.Body.Statements[0].(*ast.ExpressionStatement).Expression.(*ast.Instance)
	if !ok {
		t.Fatalf("forSt.Body.Statements[1] is not an ParallelFunctions. got=%T", forSt.Body.Statements[0].(*ast.ExpressionStatement).Expression)
	}
	if inst.Value.Spec != "test2" {
		t.Fatalf("instance has the wrong spec name. got=%s", inst.Value.Spec)
	}
	if inst.Value.Value != "foo" {
		t.Fatalf("instance has the wrong value. got=%s", inst.Value.Value)
	}
	if inst.Name != "d" {
		t.Fatalf("instance has the wrong name. got=%s", inst.Name)
	}

}

func TestIncr(t *testing.T) {
	test := `spec test1;
			 for 5 run{
				i++;
			 };
			`
	_, spec := prepTest(test, nil)
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

	if infix.Token.Type != "PLUS" {
		t.Fatalf("infix token is not PLUS. got=%s", infix.Token.Type)
	}

	if infix.Right.String() != "1" {
		t.Fatalf("infix right side is not 1. got=%s", infix.Right.String())
	}

	if infix.Left.(*ast.Identifier).Value != "i" {
		t.Fatalf("infix left side is not i. got=%s", infix.Left.(*ast.Identifier).Value)
	}

	if infix.Operator != "+" {
		t.Fatalf("infix operator is not +. got=%s", infix.Operator)
	}

}

func TestDecr(t *testing.T) {
	test := `spec test1;
			 for 5 run{
				i--;
			 };
			`
	_, spec := prepTest(test, nil)
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

	if infix.Token.Type != "MINUS" {
		t.Fatalf("infix token is not MINUS. got=%s", infix.Token.Type)
	}

	if infix.Right.String() != "1" {
		t.Fatalf("infix right side is not 1. got=%s", infix.Right.String())
	}

	if infix.Left.(*ast.Identifier).Value != "i" {
		t.Fatalf("infix left side is not i. got=%s", infix.Left.(*ast.Identifier).Value)
	}

	if infix.Operator != "-" {
		t.Fatalf("infix operator is not -. got=%s", infix.Operator)
	}

}

func TestAssertion(t *testing.T) {
	test := `spec test1;
			 assert x > y;
			`
	_, spec := prepTest(test, nil)
	assert, ok := spec.Statements[1].(*ast.AssertionStatement)
	if !ok {
		t.Fatalf("spec.Statements[1] is not an AssertionStatement. got=%T", spec.Statements[1])
	}

	if assert.Constraints.Left.(*ast.Identifier).Value != "x" {
		t.Fatalf("assert variable is not correct. got=%s, want=x", assert.Constraints.Left.(*ast.Identifier).Value)
	}

	if assert.Constraints.Operator != ">" {
		t.Fatalf("assert comparison is not correct. got=%s, want=>", assert.Constraints.Operator)
	}

	if assert.Constraints.Right.String() != "y" {
		t.Fatalf("assert comparison is not correct. got=%s, want=y", assert.Constraints.Right.(*ast.Identifier).Value)
	}

}

func TestAssertionCompound(t *testing.T) {
	test := `spec test1;
			 assert x > y && x > 1;
			`
	_, spec := prepTest(test, nil)
	assert, ok := spec.Statements[1].(*ast.AssertionStatement)
	if !ok {
		t.Fatalf("spec.Statements[1] is not an AssertionStatement. got=%T", spec.Statements[1])
	}

	if assert.Constraints.Operator != "&&" {
		t.Fatalf("assert comparison is not correct. got=%s, want=&&", assert.Constraints.Operator)
	}

}

func TestAssertionCompound2(t *testing.T) {
	test := `spec test1;
			 assert x > y || x > 1;
			`
	_, spec := prepTest(test, nil)
	assert, ok := spec.Statements[1].(*ast.AssertionStatement)
	if !ok {
		t.Fatalf("spec.Statements[1] is not an AssertionStatement. got=%T", spec.Statements[1])
	}

	if assert.Constraints.Operator != "||" {
		t.Fatalf("assert comparison is not correct. got=%s, want=||", assert.Constraints.Operator)
	}

}

func TestAssumption(t *testing.T) {
	test := `spec test1;
			 assume x == 5;
			`
	_, spec := prepTest(test, nil)
	assert, ok := spec.Statements[1].(*ast.AssumptionStatement)
	if !ok {
		t.Fatalf("spec.Statements[1] is not an AssumptionStatement. got=%T", spec.Statements[1])
	}

	if assert.Constraints.Left.(*ast.Identifier).Value != "x" {
		t.Fatalf("assumption variable is not correct. got=%s, want=x", assert.Constraints.Left.(*ast.Identifier).Value)
	}

	if assert.Constraints.Operator != "==" {
		t.Fatalf("assumption comparison is not correct. got=%s, want=>", assert.Constraints.Operator)
	}

	if assert.Constraints.Right.String() != "5" {
		t.Fatalf("assumption comparison is not correct. got=%d, want=5", assert.Constraints.Right.(*ast.IntegerLiteral).Value)
	}

}

func TestAssumptionCompound(t *testing.T) {
	test := `spec test1;
			 assume x == 5 || y > 1;
			`
	_, spec := prepTest(test, nil)
	assert, ok := spec.Statements[1].(*ast.AssumptionStatement)
	if !ok {
		t.Fatalf("spec.Statements[1] is not an AssumptionStatement. got=%T", spec.Statements[1])
	}

	if assert.Constraints.Operator != "||" {
		t.Fatalf("assumption comparison is not correct. got=%s, want=&&", assert.Constraints.Operator)
	}

}

func TestAssumptionCompound2(t *testing.T) {
	test := `spec test1;
			 assume x == 5 && y > 1;
			`
	_, spec := prepTest(test, nil)
	assert, ok := spec.Statements[1].(*ast.AssumptionStatement)
	if !ok {
		t.Fatalf("spec.Statements[1] is not an AssumptionStatement. got=%T", spec.Statements[1])
	}

	if assert.Constraints.Operator != "&&" {
		t.Fatalf("assumption comparison is not correct. got=%s, want=&&", assert.Constraints.Operator)
	}

}

func TestTemporal(t *testing.T) {
	test := `spec test1;
			 assert x > y eventually;
			`
	_, spec := prepTest(test, nil)
	assert, ok := spec.Statements[1].(*ast.AssertionStatement)
	if !ok {
		t.Fatalf("spec.Statements[1] is not an AssertionStatement. got=%T", spec.Statements[1])
	}

	if assert.Constraints.Left.(*ast.Identifier).Value != "x" {
		t.Fatalf("assert variable is not correct. got=%s, want=x", assert.Constraints.Left.(*ast.Identifier).Value)
	}

	if assert.Constraints.Operator != ">" {
		t.Fatalf("assert comparison is not correct. got=%s, want=>", assert.Constraints.Operator)
	}

	if assert.Constraints.Right.String() != "y" {
		t.Fatalf("assert comparison is not correct. got=%s, want=y", assert.Constraints.Right.(*ast.Identifier).Value)
	}

	if assert.Temporal != "eventually" {
		t.Fatalf("assert comparison is not correct. got=%s, want=eventually", assert.Temporal)
	}

}

func TestTemporalFilter(t *testing.T) {
	test := `spec test1;
			 assert x > y nmt 3;
			`
	_, spec := prepTest(test, nil)
	assert, ok := spec.Statements[1].(*ast.AssertionStatement)
	if !ok {
		t.Fatalf("spec.Statements[1] is not an AssertionStatement. got=%T", spec.Statements[1])
	}

	if assert.Constraints.Left.(*ast.Identifier).Value != "x" {
		t.Fatalf("assert variable is not correct. got=%s, want=x", assert.Constraints.Left.(*ast.Identifier).Value)
	}

	if assert.Constraints.Operator != ">" {
		t.Fatalf("assert comparison is not correct. got=%s, want=>", assert.Constraints.Operator)
	}

	if assert.Constraints.Right.String() != "y" {
		t.Fatalf("assert comparison is not correct. got=%s, want=y", assert.Constraints.Right.(*ast.Identifier).Value)
	}

	if assert.TemporalFilter != "nmt" {
		t.Fatalf("assert comparison is not correct. got=%s, want=nmt", assert.TemporalFilter)
	}

	if assert.TemporalN != 3 {
		t.Fatalf("assert comparison is not correct. got=%d, want=3", assert.TemporalN)
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
	_, spec := prepTest(test, nil)
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

		_, ok = assign.Right.(*ast.InfixExpression)
		if !ok {
			t.Fatalf("Right value is not an infix. got=%T", assign.Right)
		}
		if assign.Operator != "<-" {
			t.Fatalf("Operator is not <-. got=%s", assign.Operator)
		}
	}
}

func TestMiscAssign(t *testing.T) {
	test := `spec test1;
			 def foo = flow{
				bar: func{
					test.fuzz = 10;
				},
			 };
			`
	_, spec := prepTest(test, nil)
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
		if assign.Left.String() != "test.fuzz" {
			t.Fatalf("Left value is not test.fuzz. got=%s", assign.Left.String())
		}

		_, ok = assign.Right.(*ast.IntegerLiteral)
		if !ok {
			t.Fatalf("Right value is not an integer. got=%T", assign.Right)
		}
		if assign.Operator != "=" {
			t.Fatalf("Operator is not =. got=%s", assign.Operator)
		}
	}
}

func TestNil(t *testing.T) {
	test := `spec test1;
			 const a = nil;
			`
	_, spec := prepTest(test, nil)
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
	_, spec := prepTest(test, nil)
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

	if idx2.Left.(*ast.Identifier).Value != "b" {
		t.Fatalf("IndexExpression Left is not b. got=%s", idx2.Left.(*ast.Identifier).Value)
	}
}

func TestAccessHistory2(t *testing.T) {
	test := `spec test1;
			 const a = b[a[2]];
			`
	_, spec := prepTest(test, nil)
	con, ok := spec.Statements[1].(*ast.ConstantStatement)
	if !ok {
		t.Fatalf("spec.Statements[1] is not a ConstantStatement. got=%T", spec.Statements[1])
	}

	idx1, ok := con.Value.(*ast.IndexExpression)
	if !ok {
		t.Fatalf("Constant is not an IndexExpression. got=%T", con.Value)
	}

	if idx1.Left.(*ast.Identifier).Value != "b" {
		t.Fatalf("IndexExpression Left is not b. got=%s", idx1.Left.(*ast.Identifier).Value)
	}
}

func TestNegInt(t *testing.T) {
	test := `spec test1;
			 const a = -13;
			`
	_, spec := prepTest(test, nil)
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
	_, spec := prepTest(test, nil)
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
	_, spec := prepTest(test, nil)
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

func TestDeclaredType(t *testing.T) {
	test := `spec test1;
			 const a = natural(1);
			 const b = uncertain(10, 2.3);
			`
	l, spec := prepTest(test, nil)
	con, ok := spec.Statements[1].(*ast.ConstantStatement)
	if !ok {
		t.Fatalf("spec.Statements[1] is not a ConstantStatement. got=%T", spec.Statements[1])
	}

	nat, ok := con.Value.(*ast.Natural)
	if !ok {
		t.Fatalf("Constant is not a Natural. got=%T", con.Value)
	}

	if nat.Value != 1 {
		t.Fatalf("Natural is not 1. got=%d", nat.Value)
	}

	con1, ok := spec.Statements[2].(*ast.ConstantStatement)
	if !ok {
		t.Fatalf("spec.Statements[2] is not a ConstantStatement. got=%T", spec.Statements[2])
	}

	uncer, ok := con1.Value.(*ast.Uncertain)
	if !ok {
		t.Fatalf("Constant is not an Uncertain. got=%T", con1.Value)
	}

	if uncer.Mean != 10 {
		t.Fatalf("Uncertain mean is not 10. got=%f", uncer.Mean)
	}

	if uncer.Sigma != 2.3 {
		t.Fatalf("Uncertain sigma is not 2.3. got=%f", uncer.Sigma)
	}

	if l.Uncertains["b"] != nil {
		t.Fatal("Uncertain b not indexed")
	}
}

func TestUnknown(t *testing.T) {
	test := `spec test1;
			 const a = unknown();

			 def f = flow{
				foo,
				buzz: unknown(),
				bizz: func{
					x = unknown();
					z = x + unknown(y);
				},
			 };

			 const b;
			`
	l, spec := prepTest(test, nil)
	con, ok := spec.Statements[1].(*ast.ConstantStatement)
	if !ok {
		t.Fatalf("spec.Statements[1] is not a ConstantStatement. got=%T", spec.Statements[1])
	}

	unkn, ok := con.Value.(*ast.Unknown)
	if !ok {
		t.Fatalf("Constant is not an Unknown. got=%T", con.Value)
	}

	if unkn.Name == nil {
		t.Fatal("Unknown identifier Name is nil.")
	}

	if unkn.Name.Value != "a" {
		t.Fatalf("Unknown identifier is not correctly set. got=%s", unkn.Name.Value)
	}

	if unkn.Name.Spec != "test1" {
		t.Fatalf("Unknown spec is not correctly set. got=%s", unkn.Name.Spec)
	}

	df, ok := spec.Statements[2].(*ast.DefStatement)
	if !ok {
		t.Fatalf("spec.Statements[2] is not a DefStatement. got=%T", spec.Statements[2])
	}

	fl, ok := df.Value.(*ast.FlowLiteral)
	if !ok {
		t.Fatalf("Def block not a flow. got=%T", df.Value)
	}

	var pass int
	for ident, v := range fl.Pairs {

		if ident.Value == "foo" {
			val, ok := v.(*ast.Unknown)
			if !ok {
				t.Fatalf("Improper value for property pair %s. got=%T", ident.Value, v)
			}

			if val.Name == nil {
				t.Fatal("Unknown foo identifier Name is nil.")
			}

			if val.Name.Value != ident.Value {
				t.Fatalf("Unknown has wrong identifier. got=%s", val.Name.Value)
			}

			if val.Name.Spec != "test1" {
				t.Fatalf("Unknown has wrong spec. got=%s", val.Name.Spec)
			}
			pass++
		} else if ident.Value == "buzz" {
			val, ok := v.(*ast.Unknown)
			if !ok {
				t.Fatalf("Improper value for property pair %s. got=%T", ident.Value, v)
			}

			if val.Name == nil {
				t.Fatal("Unknown buzz identifier Name is nil.")
			}

			if val.Name.Value != ident.Value {
				t.Fatalf("Unknown has wrong identifier. got=%s", val.Name.Value)
			}

			if val.Name.Spec != "test1" {
				t.Fatalf("Unknown has wrong spec. got=%s", val.Name.Spec)
			}
			pass++
		} else if ident.Value == "bizz" {
			val, ok := v.(*ast.FunctionLiteral)
			if !ok {
				t.Fatalf("Improper value for property pair %s. got=%T", ident.Value, v)
			}

			infix1, ok := val.Body.Statements[0].(*ast.ExpressionStatement).Expression.(*ast.InfixExpression)
			if !ok {
				t.Fatalf("Incorrect function statement want=*ast.InfixExpression got=%T", val.Body.Statements[0].(*ast.ExpressionStatement).Expression)
			}

			id1, ok := infix1.Left.(*ast.Identifier)
			if !ok {
				t.Fatalf("Incorrect function statement want=*ast.Identifier got=%T", infix1.Left)
			}

			un1, ok := infix1.Right.(*ast.Unknown)
			if !ok {
				t.Fatalf("Incorrect function statement want=*ast.Unknown got=%T", infix1.Right)
			}

			if un1.Name == nil {
				t.Fatal("Unknown x identifier Name is nil.")
			}

			if un1.Name.Value != id1.Value {
				t.Fatalf("Unknown has wrong identifier. got=%s", un1.Name.Value)
			}

			if un1.Name.Spec != "test1" {
				t.Fatalf("Unknown has wrong spec. got=%s", un1.Name.Spec)
			}

			infix2, ok := val.Body.Statements[1].(*ast.ExpressionStatement).Expression.(*ast.InfixExpression)
			if !ok {
				t.Fatalf("Incorrect function statement want=*ast.InfixExpression got=%T", val.Body.Statements[1].(*ast.ExpressionStatement).Expression)
			}

			_, ok = infix2.Left.(*ast.Identifier)
			if !ok {
				t.Fatalf("Incorrect function statement want=*ast.Identifier got=%T", infix2.Left)
			}

			infix3, ok := infix2.Right.(*ast.InfixExpression)
			if !ok {
				t.Fatalf("Incorrect function statement want=*ast.InfixExpression got=%T", val.Body.Statements[1].(*ast.ExpressionStatement).Expression)
			}

			un2, ok := infix3.Right.(*ast.Unknown)
			if !ok {
				t.Fatalf("Incorrect function statement want=*ast.Unknown got=%T", infix3.Right)
			}

			if un2.Name == nil {
				t.Fatal("Unknown y identifier Name is nil.")
			}

			if un2.Name.Value != "y" {
				t.Fatalf("Unknown has wrong identifier. got=%s", un1.Name.Value)
			}

			if un2.Name.Spec != "test1" {
				t.Fatalf("Unknown has wrong spec. got=%s", un1.Name.Spec)
			}

			pass++
		} else {
			t.Fatalf("Flow has extra property. got=%s", ident.Value)
		}
	}
	if pass != 3 {
		t.Fatalf("Flow is missing property. want=3 got=%d", pass)
	}

	con2, ok := spec.Statements[3].(*ast.ConstantStatement)
	if !ok {
		t.Fatalf("spec.Statements[3] is not a ConstantStatement. got=%T", spec.Statements[3])
	}

	unk2, ok := con2.Value.(*ast.Unknown)
	if !ok {
		t.Fatalf("Constant is not an Unknown. got=%T", con2.Value)
	}

	if unk2.Name.Value != "b" {
		t.Fatalf("Unknown identifier is not correctly set. got=%s", unk2.Name.Value)
	}

	if unk2.Name.Spec != "test1" {
		t.Fatalf("Unknown spec is not correctly set. got=%s", unk2.Name.Spec)
	}

	if len(l.Unknowns) != 5 {
		t.Fatalf("missing an unknown want=5 got=%d", len(l.Unknowns))
	}
}

func TestSysSpec(t *testing.T) {
	test := `system test1;

			import "foo.fspec";

			component f = states{
				test: new foo.bar,
				initial: func{
					advance(this.next);
				},
				next: func{
					stay();
				},
			 };

			
			for 1 run {
				car = new f;
				bot = new foo.bar;
				car.test = bot;
			}
			`
	_, sys := prepSysTest(test, nil)

	_, ok := sys.Statements[0].(*ast.SysDeclStatement)
	if !ok {
		t.Fatalf("sys.Statements[0] is not an SysDeclStatement. got=%T", sys.Statements[1])
	}

	_, ok2 := sys.Statements[1].(*ast.ImportStatement)
	if !ok2 {
		t.Fatalf("sys.Statements[1] is not an ImportStatement. got=%T", sys.Statements[1])
	}

	component, ok3 := sys.Statements[2].(*ast.DefStatement).Value.(*ast.ComponentLiteral)
	if !ok3 {
		t.Fatalf("sys.Statements[2] is not a ComponentLiteral. got=%T", sys.Statements[2])
	}

	if len(component.Pairs) != 3 {
		t.Fatalf("wrong number of component pairs. got=%d", len(component.Pairs))
	}

	_, ok4 := sys.Statements[3].(*ast.ForStatement)
	if !ok4 {
		t.Fatalf("sys.Statements[3] is not a ForStatement. got=%T", sys.Statements[3])
	}

}

func TestSysGlobal(t *testing.T) {
	test := `system test1;

			import "foo.fspec";

			global t = new foo.test;
			`
	_, sys := prepSysTest(test, nil)

	global, ok := sys.Statements[2].(*ast.DefStatement).Value.(*ast.Instance)
	if !ok {
		t.Fatalf("sys.Statements[2] is not a ComponentLiteral. got=%T", sys.Statements[2])
	}

	if global.Value.Value != "test" {
		t.Fatalf("wrong value for global instance. got=%s", global.Value.Value)
	}

	if global.Value.Spec != "foo" {
		t.Fatalf("wrong spec value for global instance. got=%s", global.Value.Spec)
	}

	if global.Name != "t" {
		t.Fatalf("wrong name for global instance. got=%s", global.Name)
	}

}

func TestSysStart(t *testing.T) {
	test := `system test1;

			component test = states{
				idle: func{},
				active: func{},
			};

			component test2 = states{
				idle: func{},
				active: func{},
			};

			start {
				test:idle,
				test2:active,
			};
			`
	_, sys := prepSysTest(test, nil)

	starts, ok := sys.Statements[3].(*ast.StartStatement)
	if !ok {
		t.Fatalf("sys.Statements[3] is not a StartStatement. got=%T", sys.Statements[3])
	}

	if len(starts.Pairs) != 2 {
		t.Fatalf("start block has the wrong number of expressions. got=%d", len(starts.Pairs))
	}

}

func prepTest(test string, flags map[string]bool) (*FaultListener, *ast.Spec) {
	path := ""
	is := antlr.NewInputStream(test)
	lexer := parser.NewFaultLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	p := parser.NewFaultParser(stream)
	var listener *FaultListener
	if flags != nil && flags["skipRun"] {
		listener = NewListener(path, true, true)
	} else {
		listener = NewListener(path, true, false)
	}
	antlr.ParseTreeWalkerDefault.Walk(listener, p.Spec())
	return listener, listener.AST
}

func prepSysTest(test string, flags map[string]bool) (*FaultListener, *ast.Spec) {
	path := ""
	is := antlr.NewInputStream(test)
	lexer := parser.NewFaultLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	p := parser.NewFaultParser(stream)
	var listener *FaultListener
	if flags != nil && flags["skipRun"] {
		listener = NewListener(path, true, true)
	} else {
		listener = NewListener(path, true, false)
	}
	antlr.ParseTreeWalkerDefault.Walk(listener, p.SysSpec())
	return listener, listener.AST
}
