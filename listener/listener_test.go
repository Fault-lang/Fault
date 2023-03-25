package listener

import (
	"fault/ast"
	"testing"
)

func TestSpecDecl(t *testing.T) {
	test := `spec test1;`
	flags := make(map[string]bool)
	flags["specType"] = true
	_, spec := prepTest(test, flags)

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
	flags := make(map[string]bool)
	flags["specType"] = true
	_, spec := prepTest(test, flags)

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
	flags := make(map[string]bool)
	flags["specType"] = true
	_, spec := prepTest(test, flags)

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
	flags := make(map[string]bool)
	flags["specType"] = true
	_, spec := prepTest(test, flags)

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

func TestStockDecl(t *testing.T) {
	test := `spec test1;
			 def foo = stock{
				value: 100,
				test: buzz,
				call: test2.lol,
			 };
			`
	flags := make(map[string]bool)
	flags["specType"] = true
	_, spec := prepTest(test, flags)

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
	flags := make(map[string]bool)
	flags["specType"] = true
	_, spec := prepTest(test, flags)

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
	flags := make(map[string]bool)
	flags["specType"] = true
	_, spec := prepTest(test, flags)

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
	flags := make(map[string]bool)
	flags["specType"] = true
	_, spec := prepTest(test, flags)
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
				st: new foo,
			 };
			`
	flags := make(map[string]bool)
	flags["specType"] = true
	_, spec := prepTest(test, flags)
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
	flags := make(map[string]bool)
	flags["specType"] = true
	_, spec := prepTest(test, flags)
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
	flags := make(map[string]bool)
	flags["specType"] = true
	_, spec := prepTest(test, flags)
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
	flags := make(map[string]bool)
	flags["specType"] = true
	_, spec := prepTest(test, flags)
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
	flags := make(map[string]bool)
	flags["specType"] = true
	_, spec := prepTest(test, flags)
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
	flags := make(map[string]bool)
	flags["specType"] = true
	_, spec := prepTest(test, flags)
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
	flags := make(map[string]bool)
	flags["specType"] = true
	_, spec := prepTest(test, flags)
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
	flags := make(map[string]bool)
	flags["specType"] = true
	_, spec := prepTest(test, flags)
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
		in, ok := ife.Condition.(*ast.InfixExpression)
		if !ok {
			t.Fatalf("If Condition operand not wrapped. got=%T", ife.Condition)
		}

		if _, ok = in.Left.(*ast.Identifier); !ok {
			t.Fatalf("If Condition does not contain an Identifier. got=%T", in.Left)
		}
		if _, ok = in.Right.(*ast.Boolean); !ok {
			t.Fatalf("If Condition does not contain a boolean. got=%T", in.Right)
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
	flags := make(map[string]bool)
	flags["specType"] = true
	_, spec := prepTest(test, flags)
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
		in, ok := ife.Condition.(*ast.InfixExpression)
		if !ok {
			t.Fatalf("If Condition does not contain an Infix. got=%T", ife.Condition)
		}
		_, ok = in.Left.(*ast.Identifier)
		if !ok {
			t.Fatalf("If Condition does not contain an Identifier. got=%T", in.Left)
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
	flags := make(map[string]bool)
	flags["specType"] = true
	_, spec := prepTest(test, flags)
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
	flags := make(map[string]bool)
	flags["specType"] = true
	_, spec := prepTest(test, flags)
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
	flags := make(map[string]bool)
	flags["specType"] = false
	_, spec := prepTest(test, flags)
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
	flags := make(map[string]bool)
	flags["specType"] = false
	_, spec := prepTest(test, flags)
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
	flags := make(map[string]bool)
	flags["specType"] = false
	_, spec := prepTest(test, flags)
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
	flags := make(map[string]bool)
	flags["specType"] = true
	_, spec := prepTest(test, flags)
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
	flags := make(map[string]bool)
	flags["specType"] = true
	_, spec := prepTest(test, flags)
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

func TestRunIfBlock(t *testing.T) {
	test := `spec test1;
			 for 5 run{
				d = new foo;
				if true {
					d.fn;
				}else if false {
					d.fn2;
				}else{
					d.fn3;
				}
			 };
			`
	flags := make(map[string]bool)
	flags["specType"] = true
	_, spec := prepTest(test, flags)
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

	ifblock, ok := forSt.Body.Statements[1].(*ast.ExpressionStatement).Expression.(*ast.IfExpression)
	if !ok {
		t.Fatalf("forSt.Body.Statements[1] is not an IfExpression. got=%T", forSt.Body.Statements[1])
	}

	expr, ok := ifblock.Consequence.Statements[0].(*ast.ParallelFunctions)
	if !ok {
		t.Fatalf("if consequence is not packaged as a ParallelFunctions. got=%T", ifblock.Consequence.Statements[0])
	}

	id, ok := expr.Expressions[0].(*ast.ParameterCall)
	if !ok {
		t.Fatalf("expr.Expression is not an function call. got=%T", expr.Expressions[0])
	}

	if id.Value[0] != "d" && id.Value[0] != "fn" {
		t.Fatalf("Identifier is not d.fn. got=%s", id.Value)
	}

	expr1, ok := ifblock.Elif.Consequence.Statements[0].(*ast.ParallelFunctions)
	if !ok {
		t.Fatalf("if consequence is not packaged as a ParallelFunctions. got=%T", ifblock.Elif.Consequence.Statements[0])
	}

	id1, ok := expr1.Expressions[0].(*ast.ParameterCall)
	if !ok {
		t.Fatalf("expr.Expression is not an function call. got=%T", expr1.Expressions[0])
	}

	if id1.Value[0] != "d" && id1.Value[0] != "fn2" {
		t.Fatalf("Identifier is not d.fn2. got=%s", id1.Value)
	}

	expr2, ok := ifblock.Elif.Alternative.Statements[0].(*ast.ParallelFunctions)
	if !ok {
		t.Fatalf("if consequence is not packaged as a ParallelFunctions. got=%T", ifblock.Elif.Alternative.Statements[0])
	}

	id2, ok := expr2.Expressions[0].(*ast.ParameterCall)
	if !ok {
		t.Fatalf("expr.Expression is not an function call. got=%T", expr2.Expressions[0])
	}

	if id2.Value[0] != "d" && id2.Value[0] != "fn3" {
		t.Fatalf("Identifier is not d.fn3. got=%s", id2.Value)
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
	flags["specType"] = true

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
	flags := make(map[string]bool)
	flags["specType"] = true
	_, spec := prepTest(test, flags)
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
	flags := make(map[string]bool)
	flags["specType"] = true
	_, spec := prepTest(test, flags)
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
	flags := make(map[string]bool)
	flags["specType"] = true
	_, spec := prepTest(test, flags)
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
	flags := make(map[string]bool)
	flags["specType"] = true
	_, spec := prepTest(test, flags)
	assert, ok := spec.Statements[1].(*ast.AssertionStatement)
	if !ok {
		t.Fatalf("spec.Statements[1] is not an AssertionStatement. got=%T", spec.Statements[1])
	}

	if assert.Constraint.Left.(*ast.Identifier).Value != "x" {
		t.Fatalf("assert variable is not correct. got=%s, want=x", assert.Constraint.Left.(*ast.Identifier).Value)
	}

	if assert.Constraint.Operator != ">" {
		t.Fatalf("assert comparison is not correct. got=%s, want=>", assert.Constraint.Operator)
	}

	if assert.Constraint.Right.String() != "y" {
		t.Fatalf("assert comparison is not correct. got=%s, want=y", assert.Constraint.Right.(*ast.Identifier).Value)
	}

}

func TestAssertionCompound(t *testing.T) {
	test := `spec test1;
			 assert x > y && x > 1;
			`
	flags := make(map[string]bool)
	flags["specType"] = true
	_, spec := prepTest(test, flags)
	assert, ok := spec.Statements[1].(*ast.AssertionStatement)
	if !ok {
		t.Fatalf("spec.Statements[1] is not an AssertionStatement. got=%T", spec.Statements[1])
	}

	if assert.Constraint.Operator != "&&" {
		t.Fatalf("assert comparison is not correct. got=%s, want=&&", assert.Constraint.Operator)
	}

}

func TestAssertionCompound2(t *testing.T) {
	test := `spec test1;
			 assert x > y || x > 1;
			`
	flags := make(map[string]bool)
	flags["specType"] = true
	_, spec := prepTest(test, flags)
	assert, ok := spec.Statements[1].(*ast.AssertionStatement)
	if !ok {
		t.Fatalf("spec.Statements[1] is not an AssertionStatement. got=%T", spec.Statements[1])
	}

	if assert.Constraint.Operator != "||" {
		t.Fatalf("assert comparison is not correct. got=%s, want=||", assert.Constraint.Operator)
	}

}

func TestAssumption(t *testing.T) {
	test := `spec test1;
			 assume x == 5;
			`
	flags := make(map[string]bool)
	flags["specType"] = true
	_, spec := prepTest(test, flags)
	assert, ok := spec.Statements[1].(*ast.AssertionStatement)
	if !ok {
		t.Fatalf("spec.Statements[1] is not an AssumptionStatement. got=%T", spec.Statements[1])
	}

	if !assert.Assume {
		t.Fatal("Assumption not parsed as assumption.")
	}

	if assert.Constraint.Left.(*ast.Identifier).Value != "x" {
		t.Fatalf("assumption variable is not correct. got=%s, want=x", assert.Constraint.Left.(*ast.Identifier).Value)
	}

	if assert.Constraint.Operator != "==" {
		t.Fatalf("assumption comparison is not correct. got=%s, want=>", assert.Constraint.Operator)
	}

	if assert.Constraint.Right.String() != "5" {
		t.Fatalf("assumption comparison is not correct. got=%d, want=5", assert.Constraint.Right.(*ast.IntegerLiteral).Value)
	}

}

func TestAssumptionCompound(t *testing.T) {
	test := `spec test1;
			 assume x == 5 || y > 1;
			`
	flags := make(map[string]bool)
	flags["specType"] = true
	_, spec := prepTest(test, flags)
	assert, ok := spec.Statements[1].(*ast.AssertionStatement)
	if !ok {
		t.Fatalf("spec.Statements[1] is not an AssumptionStatement. got=%T", spec.Statements[1])
	}

	if assert.Constraint.Operator != "||" {
		t.Fatalf("assumption comparison is not correct. got=%s, want=&&", assert.Constraint.Operator)
	}

}

func TestAssumptionCompound2(t *testing.T) {
	test := `spec test1;
			 assume x == 5 && y > 1;
			`
	flags := make(map[string]bool)
	flags["specType"] = true
	_, spec := prepTest(test, flags)
	assert, ok := spec.Statements[1].(*ast.AssertionStatement)
	if !ok {
		t.Fatalf("spec.Statements[1] is not an AssumptionStatement. got=%T", spec.Statements[1])
	}

	if assert.Constraint.Operator != "&&" {
		t.Fatalf("assumption comparison is not correct. got=%s, want=&&", assert.Constraint.Operator)
	}

}

func TestTemporal(t *testing.T) {
	test := `spec test1;
			 assert x > y eventually;
			`
	flags := make(map[string]bool)
	flags["specType"] = true
	_, spec := prepTest(test, flags)
	assert, ok := spec.Statements[1].(*ast.AssertionStatement)
	if !ok {
		t.Fatalf("spec.Statements[1] is not an AssertionStatement. got=%T", spec.Statements[1])
	}

	if assert.Constraint.Left.(*ast.Identifier).Value != "x" {
		t.Fatalf("assert variable is not correct. got=%s, want=x", assert.Constraint.Left.(*ast.Identifier).Value)
	}

	if assert.Constraint.Operator != ">" {
		t.Fatalf("assert comparison is not correct. got=%s, want=>", assert.Constraint.Operator)
	}

	if assert.Constraint.Right.String() != "y" {
		t.Fatalf("assert comparison is not correct. got=%s, want=y", assert.Constraint.Right.(*ast.Identifier).Value)
	}

	if assert.Temporal != "eventually" {
		t.Fatalf("assert comparison is not correct. got=%s, want=eventually", assert.Temporal)
	}

}

func TestTemporalFilter(t *testing.T) {
	test := `spec test1;
			 assert x > y nmt 3;
			`
	flags := make(map[string]bool)
	flags["specType"] = true
	_, spec := prepTest(test, flags)
	assert, ok := spec.Statements[1].(*ast.AssertionStatement)
	if !ok {
		t.Fatalf("spec.Statements[1] is not an AssertionStatement. got=%T", spec.Statements[1])
	}

	if assert.Constraint.Left.(*ast.Identifier).Value != "x" {
		t.Fatalf("assert variable is not correct. got=%s, want=x", assert.Constraint.Left.(*ast.Identifier).Value)
	}

	if assert.Constraint.Operator != ">" {
		t.Fatalf("assert comparison is not correct. got=%s, want=>", assert.Constraint.Operator)
	}

	if assert.Constraint.Right.String() != "y" {
		t.Fatalf("assert comparison is not correct. got=%s, want=y", assert.Constraint.Right.(*ast.Identifier).Value)
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
	flags := make(map[string]bool)
	flags["specType"] = true
	_, spec := prepTest(test, flags)
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
	flags := make(map[string]bool)
	flags["specType"] = true
	_, spec := prepTest(test, flags)
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
	flags := make(map[string]bool)
	flags["specType"] = true
	_, spec := prepTest(test, flags)
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
			 for 1 run {
				b[1][2];
			 }
			`
	flags := make(map[string]bool)
	flags["specType"] = true
	_, spec := prepTest(test, flags)
	con, ok := spec.Statements[1].(*ast.ForStatement)
	if !ok {
		t.Fatalf("spec.Statements[1] is not a ForStatement. got=%T", spec.Statements[1])
	}

	idx1, ok := con.Body.Statements[0].(*ast.ExpressionStatement).Expression.(*ast.IndexExpression)
	if !ok {
		t.Fatalf("Constant is not an IndexExpression. got=%T", con.Body.Statements[0].(*ast.ExpressionStatement).Expression)
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
			for 1 run {
				b[a[2]];
			}
			`
	flags := make(map[string]bool)
	flags["specType"] = true
	_, spec := prepTest(test, flags)
	con, ok := spec.Statements[1].(*ast.ForStatement)
	if !ok {
		t.Fatalf("spec.Statements[1] is not a ForStatement. got=%T", spec.Statements[1])
	}

	idx1, ok := con.Body.Statements[0].(*ast.ExpressionStatement).Expression.(*ast.IndexExpression)
	if !ok {
		t.Fatalf("Constant is not an IndexExpression. got=%T", con.Body.Statements[0].(*ast.ExpressionStatement).Expression)
	}

	if idx1.Left.(*ast.Identifier).Value != "b" {
		t.Fatalf("IndexExpression Left is not b. got=%s", idx1.Left.(*ast.Identifier).Value)
	}
}

func TestNegInt(t *testing.T) {
	test := `spec test1;
			 const a = -13;
			`
	flags := make(map[string]bool)
	flags["specType"] = true
	_, spec := prepTest(test, flags)
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
	flags := make(map[string]bool)
	flags["specType"] = true
	_, spec := prepTest(test, flags)
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
	flags := make(map[string]bool)
	flags["specType"] = true
	_, spec := prepTest(test, flags)
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
	flags := make(map[string]bool)
	flags["specType"] = true
	l, spec := prepTest(test, flags)
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

func TestInstanceOrder(t *testing.T) {
	test := `spec test1;

			 def f = stock{
				test: new foo.bar,
				test1: 12,
				test2: -3,
			 };

			
			for 1 run {
				car = new f;
			}
			`
	flags := make(map[string]bool)
	flags["specType"] = true
	_, spec := prepTest(test, flags)

	fst, ok4 := spec.Statements[2].(*ast.ForStatement)
	if !ok4 {
		t.Fatalf("spec.Statements[2] is not a ForStatement. got=%T", spec.Statements[2])
	}

	ins, ok5 := fst.Body.Statements[0].(*ast.ExpressionStatement).Expression.(*ast.Instance)
	if !ok5 {
		t.Fatalf("ForStatement Statement is not an instance. got=%T", fst.Body.Statements[0].(*ast.ExpressionStatement).Expression)
	}

	order := ins.Order
	if len(order) != 3 || order[0] != "test" || order[1] != "test1" || order[2] != "test2" {
		t.Fatalf("instance order not correct. got=%s", order)
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
	flags := make(map[string]bool)
	flags["specType"] = true
	l, spec := prepTest(test, flags)
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

			component c = states{
				initial: func{
					advance(this.next);
				},
				close: func{
					advance(this.initial);
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
	flags := make(map[string]bool)
	flags["specType"] = false
	_, sys := prepTest(test, flags)

	decl, ok := sys.Statements[0].(*ast.SysDeclStatement)
	if !ok {
		t.Fatalf("sys.Statements[0] is not an SysDeclStatement. got=%T", sys.Statements[1])
	}

	if decl.Name.Value != "test1" {
		t.Fatalf("system is not named correctly got=%s", decl.Name)
	}

	_, ok2 := sys.Statements[1].(*ast.ImportStatement)
	if !ok2 {
		t.Fatalf("sys.Statements[1] is not an ImportStatement. got=%T", sys.Statements[1])
	}

	component, ok3 := sys.Statements[2].(*ast.DefStatement).Value.(*ast.ComponentLiteral)
	if !ok3 {
		t.Fatalf("sys.Statements[3] is not a ComponentLiteral. got=%T", sys.Statements[2])
	}

	if len(component.Pairs) != 3 {
		t.Fatalf("wrong number of component pairs. got=%d", len(component.Pairs))
	}

	for k, v := range component.Pairs {
		if f, ok := v.(*ast.FunctionLiteral); ok {
			if exp, ok2 := f.Body.Statements[0].(*ast.ExpressionStatement); ok2 {
				if ifblock, ok3 := exp.Expression.(*ast.IfExpression); !ok3 {
					if _, ok4 := exp.Expression.(*ast.BuiltIn); !ok4 {
						t.Fatalf("state %s in component not wrapped with conditional got expression=%s", k, exp.Expression)
					}
				} else {
					if b, ok4 := ifblock.Consequence.Statements[0].(*ast.ExpressionStatement).Expression.(*ast.BuiltIn); !ok4 {
						t.Fatal("missing built in function")
					} else {
						if k.Value == "initial" && (b.Function != "advance" || b.Parameters["toState"].(*ast.ParameterCall).String() != "this.next") {
							t.Fatalf("builtin in state %s formatted incorrectly", k.Value)
						}
						if k.Value == "close" && (b.Function != "advance" || b.Parameters["toState"].(*ast.ParameterCall).String() != "this.initial") {
							t.Fatalf("builtin in state %s formatted incorrectly", k.Value)
						}
						if k.Value == "next" && b.Function != "stay" {
							t.Fatalf("builtin in state %s formatted incorrectly", k.Value)
						}
					}
				}
			} else {
				t.Fatalf("state %s in component not wrapped with conditional got=%s", k, f.Body.Statements[0])
			}
		}
	}
}

func TestSysGlobal(t *testing.T) {
	test := `system test1;

			import "foo.fspec";

			global t = new foo.test;
			`
	flags := make(map[string]bool)
	flags["specType"] = false
	_, sys := prepTest(test, flags)

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
				idle: func{
					stay();
				},
				active: func{
					stay();
				},
			};

			component test2 = states{
				idle: func{
					stay();
				},
				active: func{
					stay();
				},
			};

			start {
				test:idle,
				test2:active,
			};
			`
	flags := make(map[string]bool)
	flags["specType"] = false
	_, sys := prepTest(test, flags)

	starts, ok := sys.Statements[3].(*ast.StartStatement)
	if !ok {
		t.Fatalf("sys.Statements[3] is not a StartStatement. got=%T", sys.Statements[3])
	}

	if len(starts.Pairs) != 2 {
		t.Fatalf("start block has the wrong number of expressions. got=%d", len(starts.Pairs))
	}

}

func prepTest(test string, flags map[string]bool) (*FaultListener, *ast.Spec) {
	var testRun bool
	if flags["skipRun"] {
		testRun = true
	}

	var specType bool
	if flags["specType"] {
		specType = true
	}

	listener := Execute(test, "", specType, testRun)
	return listener, listener.AST
}
