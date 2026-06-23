package listener

import (
	"fault/ast"
	"strings"
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

func TestStockExtends(t *testing.T) {
	test := `spec test1;
def generic = stock{
	id: "entity primary key",
	name: "entity name",
};
def person = stock{
	extends generic,
	occupation: "the person's job",
};
`
	flags := map[string]bool{"specType": true}
	_, spec := prepTest(test, flags)
	if spec == nil {
		t.Fatal("prepTest() returned nil")
	}

	personDef := spec.Statements[2].(*ast.DefStatement)
	if personDef.Name.Value != "person" {
		t.Fatalf("expected def name 'person', got %q", personDef.Name.Value)
	}

	person := personDef.Value.(*ast.StockLiteral)
	if person.Extends == nil {
		t.Fatal("person.Extends is nil, expected *ast.Identifier")
	}
	if person.Extends.Value != "generic" {
		t.Fatalf("person.Extends.Value = %q, want 'generic'", person.Extends.Value)
	}
	if len(person.Pairs) != 1 {
		t.Fatalf("person.Pairs should have 1 own property before type-check, got %d", len(person.Pairs))
	}
}

func TestStockExtendsWithExclude(t *testing.T) {
	test := `spec test1;
def generic = stock{
	id: "entity primary key",
	name: "entity name",
	age: "entity age",
};
def person = stock{
	extends generic,
	occupation: "the person's job",
	exclude age,
};
`
	flags := map[string]bool{"specType": true}
	_, spec := prepTest(test, flags)
	if spec == nil {
		t.Fatal("prepTest() returned nil")
	}

	person := spec.Statements[2].(*ast.DefStatement).Value.(*ast.StockLiteral)
	if person.Extends == nil {
		t.Fatal("person.Extends is nil")
	}
	if len(person.Excludes) != 1 || person.Excludes[0] != "age" {
		t.Fatalf("person.Excludes = %v, want [age]", person.Excludes)
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
				bar: func{1+now;},
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

func TestRunStatement(t *testing.T) {
	test := `spec test1;
			 run{};
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
	runSt, ok := spec.Statements[1].(*ast.RunStatement)
	if !ok {
		t.Fatalf("spec.Statements[1] is not a RunStatement. got=%T", spec.Statements[1])
	}
	if len(runSt.Steps) != 0 {
		t.Fatalf("RunStatement should have 0 steps. got=%d", len(runSt.Steps))
	}
}

func TestRunBlock(t *testing.T) {
	test := `spec test1;
			 run init{d = new foo;} {
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
	runSt, ok := spec.Statements[1].(*ast.RunStatement)
	if !ok {
		t.Fatalf("spec.Statements[1] is not a RunStatement. got=%T", spec.Statements[1])
	}

	inst, ok := runSt.Inits.Statements[0].(*ast.ExpressionStatement).Expression.(*ast.Instance)
	if !ok {
		t.Fatalf("runSt.Inits.Statements[0] is not an Instance. got=%T", runSt.Inits.Statements[0].(*ast.ExpressionStatement).Expression)
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

	if len(runSt.Steps) != 1 {
		t.Fatalf("runSt should have 1 step. got=%d", len(runSt.Steps))
	}
	step, ok := runSt.Steps[0].(*ast.CallStep)
	if !ok {
		t.Fatalf("runSt.Steps[0] is not a CallStep. got=%T", runSt.Steps[0])
	}
	if len(step.Calls) != 1 {
		t.Fatalf("CallStep does not have 1 call. got=%d", len(step.Calls))
	}
	id := step.Calls[0]
	if id.Value[0] != "d" && id.Value[0] != "fn" {
		t.Fatalf("Identifier is not d.fn. got=%s", id.Value)
	}
}


func TestSkipRun(t *testing.T) {
	test := `spec test1;
			 const a = 5;
			 run init{d = new foo;} {
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
			 run init{d = new test2.foo;} {};
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
	runSt, ok := spec.Statements[1].(*ast.RunStatement)
	if !ok {
		t.Fatalf("spec.Statements[1] is not a RunStatement. got=%T", spec.Statements[1])
	}

	inst, ok := runSt.Inits.Statements[0].(*ast.ExpressionStatement).Expression.(*ast.Instance)
	if !ok {
		t.Fatalf("runSt.Inits.Statements[0] is not an Instance. got=%T", runSt.Inits.Statements[0].(*ast.ExpressionStatement).Expression)
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

func TestRunStmtExplicitStep(t *testing.T) {
	test := `spec test1;
			 run init{d = new foo;} {
				d.fn;
			 };
			`
	flags := map[string]bool{"specType": true}
	_, spec := prepTest(test, flags)

	runSt, ok := spec.Statements[1].(*ast.RunStatement)
	if !ok {
		t.Fatalf("Statement is not a RunStatement. got=%T", spec.Statements[1])
	}
	if len(runSt.Steps) != 1 {
		t.Fatalf("RunStatement does not have 1 step. got=%d", len(runSt.Steps))
	}
	step, ok := runSt.Steps[0].(*ast.CallStep)
	if !ok {
		t.Fatalf("Step is not a CallStep. got=%T", runSt.Steps[0])
	}
	if len(step.Calls) != 1 {
		t.Fatalf("CallStep does not have 1 call. got=%d", len(step.Calls))
	}
	if step.Operator != "|" {
		t.Fatalf("CallStep operator should be | for single call from runStepExpr, got=%q", step.Operator)
	}
}

func TestRunStmtSolvableStep(t *testing.T) {
	test := `spec test1;
			 run init{d = new foo;} {
				__;
				d.fn;
				__;
			 };
			`
	flags := map[string]bool{"specType": true}
	_, spec := prepTest(test, flags)

	runSt, ok := spec.Statements[1].(*ast.RunStatement)
	if !ok {
		t.Fatalf("Statement is not a RunStatement. got=%T", spec.Statements[1])
	}
	if len(runSt.Steps) != 3 {
		t.Fatalf("RunStatement does not have 3 steps. got=%d", len(runSt.Steps))
	}
	if _, ok := runSt.Steps[0].(*ast.SolvableStep); !ok {
		t.Fatalf("Step[0] is not a SolvableStep. got=%T", runSt.Steps[0])
	}
	if _, ok := runSt.Steps[1].(*ast.CallStep); !ok {
		t.Fatalf("Step[1] is not a CallStep. got=%T", runSt.Steps[1])
	}
	if _, ok := runSt.Steps[2].(*ast.SolvableStep); !ok {
		t.Fatalf("Step[2] is not a SolvableStep. got=%T", runSt.Steps[2])
	}
}

func TestRunStmtChoiceStep(t *testing.T) {
	test := `spec test1;
			 run init{d = new foo;} {
				d.fn | d.fn2;
			 };
			`
	flags := map[string]bool{"specType": true}
	_, spec := prepTest(test, flags)

	runSt, ok := spec.Statements[1].(*ast.RunStatement)
	if !ok {
		t.Fatalf("Statement is not a RunStatement. got=%T", spec.Statements[1])
	}
	if len(runSt.Steps) != 1 {
		t.Fatalf("RunStatement does not have 1 step. got=%d", len(runSt.Steps))
	}
	step, ok := runSt.Steps[0].(*ast.CallStep)
	if !ok {
		t.Fatalf("Step is not a CallStep. got=%T", runSt.Steps[0])
	}
	if len(step.Calls) != 2 {
		t.Fatalf("CallStep does not have 2 calls. got=%d", len(step.Calls))
	}
	if step.Operator != "|" {
		t.Fatalf("CallStep operator should be |, got=%q", step.Operator)
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

			
			run init{car = new f;} {}
			`
	flags := make(map[string]bool)
	flags["specType"] = true
	_, spec := prepTest(test, flags)

	fst, ok4 := spec.Statements[2].(*ast.RunStatement)
	if !ok4 {
		t.Fatalf("spec.Statements[2] is not a RunStatement. got=%T", spec.Statements[2])
	}

	ins, ok5 := fst.Inits.Statements[0].(*ast.ExpressionStatement).Expression.(*ast.Instance)
	if !ok5 {
		t.Fatalf("RunStatement Statement is not an instance. got=%T", fst.Inits.Statements[0].(*ast.ExpressionStatement).Expression)
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
				buzz: unknown(0),
				baz: unknown(0.0),
				bizz: func{
					x = unknown(false);
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

	if unkn.TypeHint != "" {
		t.Fatalf("untyped unknown() should have empty TypeHint. got=%s", unkn.TypeHint)
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
			_, ok := v.(*ast.Unknown)
			if !ok {
				t.Fatalf("Improper value for property pair %s. got=%T", ident.Value, v)
			}
			pass++
		} else if ident.Value == "buzz" {
			val, ok := v.(*ast.Unknown)
			if !ok {
				t.Fatalf("Improper value for property pair %s. got=%T", ident.Value, v)
			}
			if val.TypeHint != "INT" {
				t.Fatalf("unknown(0) should have TypeHint INT. got=%s", val.TypeHint)
			}
			pass++
		} else if ident.Value == "baz" {
			val, ok := v.(*ast.Unknown)
			if !ok {
				t.Fatalf("Improper value for property pair %s. got=%T", ident.Value, v)
			}
			if val.TypeHint != "REAL" {
				t.Fatalf("unknown(0.0) should have TypeHint REAL. got=%s", val.TypeHint)
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

			un1, ok := infix1.Right.(*ast.Unknown)
			if !ok {
				t.Fatalf("Incorrect function statement want=*ast.Unknown got=%T", infix1.Right)
			}

			if un1.TypeHint != "BOOL" {
				t.Fatalf("unknown(false) should have TypeHint BOOL. got=%s", un1.TypeHint)
			}

			pass++
		} else {
			t.Fatalf("Flow has extra property. got=%s", ident.Value)
		}
	}
	if pass != 4 {
		t.Fatalf("Flow is missing property. want=4 got=%d", pass)
	}

	con2, ok := spec.Statements[3].(*ast.ConstantStatement)
	if !ok {
		t.Fatalf("spec.Statements[3] is not a ConstantStatement. got=%T", spec.Statements[3])
	}

	_, ok = con2.Value.(*ast.Unknown)
	if !ok {
		t.Fatalf("Constant is not an Unknown. got=%T", con2.Value)
	}

	if len(l.Unknowns) != 6 {
		t.Fatalf("missing an unknown want=6 got=%d", len(l.Unknowns))
	}

}

func TestParam(t *testing.T) {
	// param() uses a literal to hint the type, same convention as unknown().
	// param(0.0) → REAL, param(0) → INT, param(false) → BOOL.
	test := `spec test1;
			 def f = flow{
				amount: param(0.0),
				count: param(0),
				flagged: param(false),
			 };
			`
	flags := make(map[string]bool)
	flags["specType"] = true
	l, spec := prepTest(test, flags)

	df, ok := spec.Statements[1].(*ast.DefStatement)
	if !ok {
		t.Fatalf("spec.Statements[1] is not a DefStatement. got=%T", spec.Statements[1])
	}

	fl, ok := df.Value.(*ast.FlowLiteral)
	if !ok {
		t.Fatalf("Def block not a flow. got=%T", df.Value)
	}

	var pass int
	for ident, v := range fl.Pairs {
		switch ident.Value {
		case "amount":
			val, ok := v.(*ast.Param)
			if !ok {
				t.Fatalf("amount: expected *ast.Param, got=%T", v)
			}
			if val.TypeHint != "REAL" {
				t.Fatalf("param(0.0) should have TypeHint REAL. got=%s", val.TypeHint)
			}
			pass++
		case "count":
			val, ok := v.(*ast.Param)
			if !ok {
				t.Fatalf("count: expected *ast.Param, got=%T", v)
			}
			if val.TypeHint != "INT" {
				t.Fatalf("param(0) should have TypeHint INT. got=%s", val.TypeHint)
			}
			pass++
		case "flagged":
			val, ok := v.(*ast.Param)
			if !ok {
				t.Fatalf("flagged: expected *ast.Param, got=%T", v)
			}
			if val.TypeHint != "BOOL" {
				t.Fatalf("param(false) should have TypeHint BOOL. got=%s", val.TypeHint)
			}
			pass++
		default:
			t.Fatalf("Flow has unexpected property: %s", ident.Value)
		}
	}
	if pass != 3 {
		t.Fatalf("Expected 3 param properties, got=%d", pass)
	}

	if len(l.Params) != 3 {
		t.Fatalf("Expected 3 entries in l.Params, got=%d: %v", len(l.Params), l.Params)
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
			
			run {
				car = new f;
				bot = new foo.bar;
			}
			`
	flags := make(map[string]bool)
	flags["specType"] = false
	_, sys := prepTest(test, flags)

	decl, ok := sys.Statements[0].(*ast.SysDeclStatement)
	if !ok {
		t.Fatalf("sys.Statements[0] is not an SysDeclStatement. got=%T", sys.Statements[0])
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

func TestSysRunStmt(t *testing.T) {
	test := `system test1;

			import "foo.fspec";

			global fl = new foo.fl;

			component A = states{
				on: func{
					stay();
				},
				off: func{
					stay();
				},
			};

			component B = states{
				idle: func{
					stay();
				},
			};

			run {
				A.on | B.idle;
			}
			`
	flags := make(map[string]bool)
	flags["specType"] = false
	_, sys := prepTest(test, flags)

	if sys == nil {
		t.Fatal("prepTest() returned nil AST")
	}

	// RunStatement should be in the AST
	var runSt *ast.RunStatement
	for _, stmt := range sys.Statements {
		if r, ok := stmt.(*ast.RunStatement); ok {
			runSt = r
			break
		}
	}
	if runSt == nil {
		t.Fatal("no RunStatement found in AST")
	}

	// Steps should contain 1 CallStep
	if len(runSt.Steps) != 1 {
		t.Fatalf("RunStatement should have 1 step, got=%d", len(runSt.Steps))
	}

	cs, ok := runSt.Steps[0].(*ast.CallStep)
	if !ok {
		t.Fatalf("step 0 should be CallStep, got=%T", runSt.Steps[0])
	}

	if cs.Operator != "|" {
		t.Fatalf("CallStep operator should be |, got=%s", cs.Operator)
	}

	if len(cs.Calls) != 2 {
		t.Fatalf("CallStep should have 2 calls, got=%d", len(cs.Calls))
	}

	if cs.Calls[0].Value[len(cs.Calls[0].Value)-1] != "on" {
		t.Fatalf("first call should be A.on, got=%s", cs.Calls[0].Value)
	}
	if cs.Calls[1].Value[len(cs.Calls[1].Value)-1] != "idle" {
		t.Fatalf("second call should be B.idle, got=%s", cs.Calls[1].Value)
	}
}

func TestSwap(t *testing.T) {
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
			
			run init{
				bot = new foo.bar;
				car = new f;
				car.test = bot;
			} {}
			`
	flags := make(map[string]bool)
	flags["specType"] = false
	_, sys := prepTest(test, flags)

	runSt, ok := sys.Statements[3].(*ast.RunStatement)
	if !ok {
		t.Fatalf("sys.Statements[3] is not a RunStatement. got=%T", sys.Statements[3])
	}

	if len(runSt.Inits.Statements) != 2 {
		t.Fatalf("run block has the wrong number of statements. got=%d", len(runSt.Inits.Statements))
	}

	inst := runSt.Inits.Statements[1].(*ast.ExpressionStatement).Expression.(*ast.Instance)
	if inst.Swaps[0].TokenLiteral() != "SWAP" {
		t.Fatalf("swap incorrect in AST. got=%s", inst.Swaps[0].TokenLiteral())
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

			run {
				test.idle && test2.active;
			};
			`
	flags := make(map[string]bool)
	flags["specType"] = false
	_, sys := prepTest(test, flags)

	var runSt *ast.RunStatement
	for _, s := range sys.Statements {
		if r, ok := s.(*ast.RunStatement); ok {
			runSt = r
			break
		}
	}
	if runSt == nil {
		t.Fatal("no RunStatement found in AST")
	}

	if len(runSt.Steps) != 1 {
		t.Fatalf("expected 1 run step, got=%d", len(runSt.Steps))
	}

	cs, ok := runSt.Steps[0].(*ast.CallStep)
	if !ok {
		t.Fatalf("step 0 should be a CallStep, got=%T", runSt.Steps[0])
	}

	if len(cs.Calls) != 2 {
		t.Fatalf("CallStep should have 2 calls, got=%d", len(cs.Calls))
	}

}

func TestBoolCompound(t *testing.T) {
	test := `system test1;
	component test = states{
				idle: func{
					stay() || advance(this.active);
				},
				active: func{
					stay();
				},
			};
			`

	flags := make(map[string]bool)
	flags["specType"] = false
	_, sys := prepTest(test, flags)
	component, ok := sys.Statements[1].(*ast.DefStatement).Value.(*ast.ComponentLiteral)
	if !ok {
		t.Fatalf("sys.Statements[1] is not a ComponentLiteral. got=%T", sys.Statements[1])
	}

	var chooseFunc *ast.FunctionLiteral
	for k, v := range component.Pairs {
		if k.Value != "idle" && k.Value != "active" {
			t.Fatalf("unexpected state in component. got=%s", k.Value)
		}
		if fn, ok := v.(*ast.FunctionLiteral); ok {
			if k.Value == "idle" {
				chooseFunc = fn
			}
		} else {
			t.Fatalf("state %s is not a FunctionLiteral. got=%T", k.Value, v)
		}
	}

	expr, ok := chooseFunc.Body.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("Function body statement is not an ExpressionStatement. got=%T", chooseFunc.Body.Statements[0])
	}
	ifexpr, ok := expr.Expression.(*ast.IfExpression)
	if !ok {
		t.Fatalf("Expression is not a IfExpression. got=%T", expr.Expression)
	}
	_, ok = ifexpr.Consequence.Statements[0].(*ast.ExpressionStatement).Expression.(*ast.InfixExpression)
	if !ok {
		t.Fatalf("Expression is not a InfixExpression. got=%T", ifexpr.Consequence.Statements[0].(*ast.ExpressionStatement).Expression)
	}

}

func TestChoose(t *testing.T) {
	test := `system test1;
	component test = states{
				idle: func{
					choose stay() || advance(this.active);
				},
				active: func{
					stay();
				},
			};
			`

	flags := make(map[string]bool)
	flags["specType"] = false
	_, sys := prepTest(test, flags)
	component, ok := sys.Statements[1].(*ast.DefStatement).Value.(*ast.ComponentLiteral)
	if !ok {
		t.Fatalf("sys.Statements[1] is not a ComponentLiteral. got=%T", sys.Statements[1])
	}

	var chooseFunc *ast.FunctionLiteral
	for k, v := range component.Pairs {
		if k.Value != "idle" && k.Value != "active" {
			t.Fatalf("unexpected state in component. got=%s", k.Value)
		}
		if fn, ok := v.(*ast.FunctionLiteral); ok {
			if k.Value == "idle" {
				chooseFunc = fn
			}
		} else {
			t.Fatalf("state %s is not a FunctionLiteral. got=%T", k.Value, v)
		}
	}

	expr, ok := chooseFunc.Body.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("Function body statement is not an ExpressionStatement. got=%T", chooseFunc.Body.Statements[0])
	}
	ifexpr, ok := expr.Expression.(*ast.IfExpression)
	if !ok {
		t.Fatalf("Expression is not a IfExpression. got=%T", expr.Expression)
	}
	choose, ok := ifexpr.Consequence.Statements[0].(*ast.ExpressionStatement).Expression.(*ast.PrefixExpression)
	if !ok {
		t.Fatalf("Expression is not a PrefixExpression. got=%T", ifexpr.Consequence.Statements[0].(*ast.ExpressionStatement).Expression)
	}
	if choose.Operator != "choose" {
		t.Fatalf("Expression is not a choose expression. got=%s", choose.Operator)
	}

}

func TestLeave(t *testing.T) {
	test := `system test1;
	component test = states{
				idle: func{
					advance(this.active);
					leave();
				},
				active: func{
					stay() && leave(this.failure);
				},
				failure: func{
					stay();
				},
			};
			`

	flags := make(map[string]bool)
	flags["specType"] = false
	_, sys := prepTest(test, flags)
	component, ok := sys.Statements[1].(*ast.DefStatement).Value.(*ast.ComponentLiteral)
	if !ok {
		t.Fatalf("sys.Statements[1] is not a ComponentLiteral. got=%T", sys.Statements[1])
	}

	var chooseFunc, leaveFunc *ast.FunctionLiteral
	for k, v := range component.Pairs {
		if k.Value != "idle" && k.Value != "active" && k.Value != "failure" {
			t.Fatalf("unexpected state in component. got=%s", k.Value)
		}
		if fn, ok := v.(*ast.FunctionLiteral); ok {
			if k.Value == "idle" {
				chooseFunc = fn
			} else if k.Value == "active" {
				leaveFunc = fn
			}
		} else {
			t.Fatalf("state %s is not a FunctionLiteral. got=%T", k.Value, v)
		}
	}

	expr, ok := chooseFunc.Body.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("Function body statement is not an ExpressionStatement. got=%T", chooseFunc.Body.Statements[0])
	}
	ifexpr, ok := expr.Expression.(*ast.IfExpression)
	if !ok {
		t.Fatalf("Expression is not a IfExpression. got=%T", expr.Expression)
	}
	leave, ok := ifexpr.Consequence.Statements[1].(*ast.ExpressionStatement).Expression.(*ast.BuiltIn)
	if !ok {
		t.Fatalf("Expression is not a BuiltIn. got=%T", ifexpr.Consequence.Statements[1].(*ast.ExpressionStatement).Expression)
	}
	if leave.Function != "leave" {
		t.Fatalf("Expression is not a leave expression. got=%s", leave.Function)
	}

	expr2, ok := leaveFunc.Body.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("Function body statement is not an ExpressionStatement. got=%T", leaveFunc.Body.Statements[0])
	}
	ifexpr2, ok := expr2.Expression.(*ast.IfExpression)
	if !ok {
		t.Fatalf("Expression is not a IfExpression. got=%T", expr2.Expression)
	}
	infix2, ok := ifexpr2.Consequence.Statements[0].(*ast.ExpressionStatement).Expression.(*ast.InfixExpression)
	if !ok {
		t.Fatalf("Expression is not an InfixExpression. got=%T", ifexpr2.Consequence.Statements[0].(*ast.ExpressionStatement).Expression)
	}

	leave2, ok := infix2.Right.(*ast.BuiltIn)
	if !ok {
		t.Fatalf("Expression is not a BuiltIn. got=%T", infix2.Right)
	}

	if leave2.Function != "leave" {
		t.Fatalf("Expression is not a leave expression. got=%s", leave2.Function)
	}

	if leave2.Parameters["exitState"].(*ast.ParameterCall).String() != "this.failure" {
		t.Fatalf("Expression is not a leave expression. got=%s", leave2.Parameters["exitState"].(*ast.ParameterCall).String())
	}
}

// helpers for underscore validation tests

func assertUnderscoreError(t *testing.T, test string, specType bool) {
	t.Helper()
	flags := map[string]bool{
		"specType": specType,
		"testing":  true,
	}
	_, err := Execute(test, "", flags)
	if err == nil {
		t.Fatal("expected error for underscore in variable name, got nil")
	}
	if !strings.Contains(err.Error(), "must be only letters or numbers") {
		t.Fatalf("expected underscore error, got: %q", err.Error())
	}
}

func TestGlobalDeclUnderscoreError(t *testing.T) {
	assertUnderscoreError(t, `system test1;
global foo_bar = 1;`, false)
}

func TestStructDeclUnderscoreError(t *testing.T) {
	assertUnderscoreError(t, `spec test1;
def foo_bar = stock{
	value: 1,
};`, true)
}

func TestComponentDeclUnderscoreError(t *testing.T) {
	assertUnderscoreError(t, `system test1;
component foo_bar = states{
	idle: func{
		stay();
	},
};`, false)
}

func TestStringDeclUnderscoreError(t *testing.T) {
	assertUnderscoreError(t, `spec test1;
foo_bar = "hello";`, true)
}

func TestConstSpecUnderscoreError(t *testing.T) {
	assertUnderscoreError(t, `spec test1;
const foo_bar = 5;`, true)
}

func TestStateFuncUnderscoreError(t *testing.T) {
	assertUnderscoreError(t, `system test1;
component foo = states{
	foo_state: func{
		stay();
	},
};`, false)
}

func TestPropFuncUnderscoreError(t *testing.T) {
	assertUnderscoreError(t, `spec test1;
def foo = flow{
	foo_fn: func{},
};`, true)
}

func TestPropIntUnderscoreError(t *testing.T) {
	assertUnderscoreError(t, `spec test1;
def foo = stock{
	foo_val: 5,
};`, true)
}

func TestPropStringUnderscoreError(t *testing.T) {
	assertUnderscoreError(t, `spec test1;
def foo = stock{
	foo_val: "hello",
};`, true)
}

func TestPropBoolUnderscoreError(t *testing.T) {
	assertUnderscoreError(t, `spec test1;
def foo = stock{
	foo_val: true,
};`, true)
}

func TestPropVarUnderscoreError(t *testing.T) {
	assertUnderscoreError(t, `spec test1;
def foo = stock{
	value: 1,
	foo_val: value,
};`, true)
}

func TestPropSolvableUnderscoreError(t *testing.T) {
	assertUnderscoreError(t, `spec test1;
def foo = stock{
	foo_val: float(0.0, 1.0),
};`, true)
}

func TestRunInitUnderscoreError(t *testing.T) {
	assertUnderscoreError(t, `spec test1;
def foo = flow{
	x: 1,
};
run init { foo_inst = new foo; } {}`, true)
}

func TestValidVarName(t *testing.T) {
	cases := []struct {
		name  string
		input string
		valid bool
	}{
		{"letters only", "foo", true},
		{"mixed alphanumeric", "foo123", true},
		{"uppercase", "FooBar", true},
		{"underscore", "foo_bar", false},
		{"hyphen", "foo-bar", false},
		{"empty string", "", false},
		{"space", "foo bar", false},
		{"dot", "foo.bar", false},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := validVarName(tc.input); got != tc.valid {
				t.Errorf("validVarName(%q) = %v, want %v", tc.input, got, tc.valid)
			}
		})
	}
}

func TestEmptyFunctionLitError(t *testing.T) {
	// functionLit (used in flow/stock properties) must not have an empty block.
	// This mirrors the equivalent check in EnterStateBlock for stateLit.
	test := `spec test1;
def foo = flow{
	bar: func{},
};`
	flags := make(map[string]bool)
	flags["specType"] = true
	flags["testing"] = true

	_, err := Execute(test, "", flags)
	if err == nil {
		t.Fatal("expected error for empty function body, got nil")
	}

	expected := "A function cannot be empty"
	if !strings.Contains(err.Error(), expected) {
		t.Fatalf("expected error to contain %q, got %q", expected, err.Error())
	}
}

func TestEmptyStateBlockError(t *testing.T) {
	// StateBlock is used inside component = states{} in .fsystem files.
	// An empty func body (func{}) has fewer than 3 children (just { and }),
	// which should trigger the validation error.
	test := `system test1;

component x = states{
	foo: func{},
};`
	flags := make(map[string]bool)
	flags["specType"] = false // fsystem
	flags["testing"] = true

	_, err := Execute(test, "", flags)
	if err == nil {
		t.Fatal("expected error for empty state block, got nil")
	}

	expected := "A state function cannot be empty"
	if !strings.Contains(err.Error(), expected) {
		t.Fatalf("expected error to contain %q, got %q", expected, err.Error())
	}
}

func TestPopUnderflow(t *testing.T) {
	defer func() {
		r := recover()
		if r == nil {
			t.Fatal("expected panic on pop() from empty stack, got none")
		}
		msg, ok := r.(string)
		if !ok || !strings.Contains(msg, "stack underflow") {
			t.Fatalf("expected stack underflow message, got %v", r)
		}
	}()
	l := NewListener("", true, false)
	l.pop()
}

func TestPeekUnderflow(t *testing.T) {
	defer func() {
		r := recover()
		if r == nil {
			t.Fatal("expected panic on peek() from empty stack, got none")
		}
		msg, ok := r.(string)
		if !ok || !strings.Contains(msg, "stack underflow") {
			t.Fatalf("expected stack underflow message, got %v", r)
		}
	}()
	l := NewListener("", true, false)
	l.peek()
}

// --- validate() error paths ---

func TestValidateTooFewStatements(t *testing.T) {
	// A spec with only the declaration clause has 1 statement on the stack,
	// which is less than 2, so validate() should panic with "Too few statements".
	// We call Execute with testing=false to bypass the early return in validate().
	flags := map[string]bool{"specType": true, "testing": false}
	_, err := Execute(`spec test1;`, "", flags)
	if err == nil {
		t.Fatal("expected error for spec with only declaration, got nil")
	}
	if !strings.Contains(err.Error(), "Too few statements") {
		t.Fatalf("expected 'Too few statements' error, got %q", err.Error())
	}
}

func TestValidateNoModelPossible(t *testing.T) {
	// A spec with declaration + const but no assert/def/run block has 2 statements
	// on the stack, but neither qualifies as a model driver, so validate() should
	// panic with "No model possible".
	flags := map[string]bool{"specType": true, "testing": false}
	_, err := Execute(`spec test1;
const x = 5;`, "", flags)
	if err == nil {
		t.Fatal("expected error for spec with no assert/def/run, got nil")
	}
	if !strings.Contains(err.Error(), "No model possible") {
		t.Fatalf("expected 'No model possible' error, got %q", err.Error())
	}
}

// --- assert/assume inside func/unfunc bodies ---

// assert and assume inside func{} / unfunc{} bodies are syntax errors at the
// grammar level — the parser rejects them before the listener fires. These tests
// confirm the grammar enforcement (the listener guards are a defensive fallback
// for any future grammar relaxation).
func TestAssertInsideFunc(t *testing.T) {
	flags := map[string]bool{"specType": true, "testing": false}
	_, err := Execute(`spec test1;
def f = flow{
	change: func{ assert x > 0; },
};
assert x > 0;
`, "", flags)
	if err == nil {
		t.Fatal("expected error for assert inside func body, got nil")
	}
	// Grammar rejects it before the listener fires.
	if !strings.Contains(err.Error(), "Invalid spec syntax") {
		t.Fatalf("unexpected error: %q", err.Error())
	}
}

func TestAssumeInsideFunc(t *testing.T) {
	flags := map[string]bool{"specType": true, "testing": false}
	_, err := Execute(`spec test1;
def f = flow{
	change: func{ assume x > 0; },
};
assert x > 0;
`, "", flags)
	if err == nil {
		t.Fatal("expected error for assume inside func body, got nil")
	}
	if !strings.Contains(err.Error(), "Invalid spec syntax") {
		t.Fatalf("unexpected error: %q", err.Error())
	}
}

func TestAssertInsideUnfunc(t *testing.T) {
	flags := map[string]bool{"specType": false, "testing": false}
	_, err := Execute(`system test1;
component fetch = states{
	getByName: unfunc{
		requires fetch.ready,
		assert fetch.ready > 0,
	},
};
`, "", flags)
	if err == nil {
		t.Fatal("expected error for assert inside unfunc body, got nil")
	}
	if !strings.Contains(err.Error(), "Invalid spec syntax") {
		t.Fatalf("unexpected error: %q", err.Error())
	}
}

func TestAssumeInsideUnfunc(t *testing.T) {
	flags := map[string]bool{"specType": false, "testing": false}
	_, err := Execute(`system test1;
component fetch = states{
	getByName: unfunc{
		requires fetch.ready,
		assume fetch.ready > 0,
	},
};
`, "", flags)
	if err == nil {
		t.Fatal("expected error for assume inside unfunc body, got nil")
	}
	if !strings.Contains(err.Error(), "Invalid spec syntax") {
		t.Fatalf("unexpected error: %q", err.Error())
	}
}

// --- String declaration happy paths ---

func TestStringDeclHappyPath(t *testing.T) {
	test := `spec test1;
foo = "hello";
def bar = flow{
	baz: func{foo + 1;},
};
run{};
`
	flags := map[string]bool{"specType": true}
	_, spec := prepTest(test, flags)
	if spec == nil {
		t.Fatal("prepTest() returned nil")
	}

	// Statement[1] should be a DefStatement wrapping a StringLiteral.
	ds, ok := spec.Statements[1].(*ast.DefStatement)
	if !ok {
		t.Fatalf("spec.Statements[1] is not *ast.DefStatement, got %T", spec.Statements[1])
	}
	if ds.Name.Value != "foo" {
		t.Fatalf("string decl name: expected foo, got %s", ds.Name.Value)
	}
	if _, ok := ds.Value.(*ast.StringLiteral); !ok {
		t.Fatalf("string decl value is not *ast.StringLiteral, got %T", ds.Value)
	}
}

func TestStringDeclCompound(t *testing.T) {
	// A compound string (a || b) combines operand names with || and is
	// represented as an InfixExpression in the DefStatement value.
	test := `spec test1;
foo = a || b;
def bar = flow{
	baz: func{foo + 1;},
};
run{};
`
	flags := map[string]bool{"specType": true}
	_, spec := prepTest(test, flags)
	if spec == nil {
		t.Fatal("prepTest() returned nil")
	}

	ds, ok := spec.Statements[1].(*ast.DefStatement)
	if !ok {
		t.Fatalf("spec.Statements[1] is not *ast.DefStatement, got %T", spec.Statements[1])
	}
	if _, ok := ds.Value.(*ast.InfixExpression); !ok {
		t.Fatalf("compound string value is not *ast.InfixExpression, got %T", ds.Value)
	}
}

// --- Invariant definitions ---

func TestDefInvariant(t *testing.T) {
	// assert x = y; uses the defInvariant rule → InvariantClause with Operator "=="
	test := `spec test1;
def foo = stock{
	x: 5,
};
assert foo.x = 10;
`
	flags := map[string]bool{"specType": true}
	_, spec := prepTest(test, flags)
	if spec == nil {
		t.Fatal("prepTest() returned nil")
	}

	as, ok := spec.Statements[2].(*ast.AssertionStatement)
	if !ok {
		t.Fatalf("spec.Statements[2] is not *ast.AssertionStatement, got %T", spec.Statements[2])
	}
	if as.Constraint == nil {
		t.Fatal("AssertionStatement.Constraint is nil")
	}
	if as.Constraint.Operator != "==" {
		t.Fatalf("defInvariant operator: expected ==, got %q", as.Constraint.Operator)
	}
}

func TestStageInvariant(t *testing.T) {
	// assert when x > 0 then y > 0; uses the stageInvariant rule → InvariantClause with Operator "then"
	test := `spec test1;
def foo = stock{
	x: 5,
	y: 3,
};
assert when foo.x > 0 then foo.y > 0;
`
	flags := map[string]bool{"specType": true}
	_, spec := prepTest(test, flags)
	if spec == nil {
		t.Fatal("prepTest() returned nil")
	}

	as, ok := spec.Statements[2].(*ast.AssertionStatement)
	if !ok {
		t.Fatalf("spec.Statements[2] is not *ast.AssertionStatement, got %T", spec.Statements[2])
	}
	if as.Constraint == nil {
		t.Fatal("AssertionStatement.Constraint is nil")
	}
	if as.Constraint.Operator != "then" {
		t.Fatalf("stageInvariant operator: expected then, got %q", as.Constraint.Operator)
	}
}

// --- Multiple components in a system ---

func TestMultipleComponentsStateActivation(t *testing.T) {
	// A run block with a state activation step for two components should produce
	// a RunStatement with one CallStep containing two ParameterCalls.
	test := `system test1;

component foo = states{
	idle: func{
		stay();
	},
};

component bar = states{
	running: func{
		stay();
	},
};

run {
	foo.idle && bar.running;
};
`
	flags := map[string]bool{"specType": false}
	_, spec := prepTest(test, flags)
	if spec == nil {
		t.Fatal("prepTest() returned nil")
	}

	var runSt *ast.RunStatement
	for _, s := range spec.Statements {
		if r, ok := s.(*ast.RunStatement); ok {
			runSt = r
			break
		}
	}
	if runSt == nil {
		t.Fatal("no RunStatement found in spec")
	}
	if len(runSt.Steps) != 1 {
		t.Fatalf("expected 1 step, got %d", len(runSt.Steps))
	}
	cs, ok := runSt.Steps[0].(*ast.CallStep)
	if !ok {
		t.Fatalf("step 0 should be CallStep, got=%T", runSt.Steps[0])
	}
	if len(cs.Calls) != 2 {
		t.Fatalf("expected 2 calls in step, got %d", len(cs.Calls))
	}
}

// --- Nested conditionals ---

func TestNestedConditional(t *testing.T) {
	// An if inside an if inside a flow function body should parse without error
	// and produce nested IfExpression nodes.
	test := `spec test1;
def foo = flow{
	bar: func{
		if(x > 0){
			if(y > 0){
				1;
			}
		}
	},
};
`
	flags := map[string]bool{"specType": true}
	_, spec := prepTest(test, flags)
	if spec == nil {
		t.Fatal("prepTest() returned nil")
	}

	flow := spec.Statements[1].(*ast.DefStatement).Value.(*ast.FlowLiteral).Pairs
	if len(flow) == 0 {
		t.Fatal("no pairs in flow literal")
	}
	var fn *ast.FunctionLiteral
	for _, v := range flow {
		f, ok := v.(*ast.FunctionLiteral)
		if !ok {
			t.Fatalf("pair value is not a FunctionLiteral, got %T", v)
		}
		fn = f
		break
	}
	if len(fn.Body.Statements) == 0 {
		t.Fatal("function body is empty")
	}
	outerIf, ok := fn.Body.Statements[0].(*ast.ExpressionStatement).Expression.(*ast.IfExpression)
	if !ok {
		t.Fatalf("outer statement is not an IfExpression, got %T", fn.Body.Statements[0].(*ast.ExpressionStatement).Expression)
	}
	if len(outerIf.Consequence.Statements) == 0 {
		t.Fatal("outer if has empty consequence")
	}
	_, ok = outerIf.Consequence.Statements[0].(*ast.ExpressionStatement).Expression.(*ast.IfExpression)
	if !ok {
		t.Fatalf("inner statement is not an IfExpression, got %T", outerIf.Consequence.Statements[0].(*ast.ExpressionStatement).Expression)
	}
}

// --- Scope reset between struct definitions ---

func TestScopeResetBetweenDefs(t *testing.T) {
	// After parsing two def statements the listener scope should be empty,
	// confirming that the second definition was not accidentally scoped to the first.
	test := `spec test1;
def foo = stock{
	x: 5,
};
def bar = stock{
	y: 3,
};
assert foo.x > bar.y;
`
	flags := map[string]bool{"specType": true}
	l, spec := prepTest(test, flags)
	if spec == nil {
		t.Fatal("prepTest() returned nil")
	}
	if l.scope != "" {
		t.Fatalf("listener scope should be empty after parsing, got %q", l.scope)
	}
}

func prepTest(test string, flags map[string]bool) (*FaultListener, *ast.Spec) {
	flags["testing"] = true
	listener, _ := Execute(test, "", flags)
	return listener, listener.AST
}

// --- unfunc{} listener tests ---

// unfuncLiteral returns the single UnfuncLiteral from a component with one state.
func unfuncLiteralFromSpec(t *testing.T, spec *ast.Spec) *ast.UnfuncLiteral {
	t.Helper()
	comp := spec.Statements[1].(*ast.DefStatement).Value.(*ast.ComponentLiteral)
	for _, v := range comp.Pairs {
		uf, ok := v.(*ast.UnfuncLiteral)
		if !ok {
			t.Fatalf("component pair value is not UnfuncLiteral, got %T", v)
		}
		return uf
	}
	t.Fatal("no pairs in component")
	return nil
}

func TestUnfuncSimple(t *testing.T) {
	// A single unfunc state with one requires and one emits clause.
	test := `system test1;

component fetch = states{
	getByName: unfunc{
		requires generic.name,
		emits generic.id,
	},
};
`
	flags := map[string]bool{"specType": false}
	_, spec := prepTest(test, flags)
	if spec == nil {
		t.Fatal("prepTest() returned nil")
	}

	def, ok := spec.Statements[1].(*ast.DefStatement)
	if !ok {
		t.Fatalf("expected DefStatement, got %T", spec.Statements[1])
	}
	comp, ok := def.Value.(*ast.ComponentLiteral)
	if !ok {
		t.Fatalf("expected ComponentLiteral, got %T", def.Value)
	}
	var uf *ast.UnfuncLiteral
	for _, v := range comp.Pairs {
		u, ok := v.(*ast.UnfuncLiteral)
		if !ok {
			t.Fatalf("component pair value is not UnfuncLiteral, got %T", v)
		}
		uf = u
		break
	}

	req, ok := uf.Requires.(*ast.ParameterCall)
	if !ok {
		t.Fatalf("Requires is not a ParameterCall, got %T", uf.Requires)
	}
	if strings.Join(req.Value, ".") != "generic.name" {
		t.Errorf("Requires = %v, want generic.name", req.Value)
	}

	if len(uf.Emits) != 1 {
		t.Fatalf("expected 1 emit, got %d", len(uf.Emits))
	}
	emit, ok := uf.Emits[0].(*ast.ParameterCall)
	if !ok {
		t.Fatalf("Emits[0] is not a ParameterCall, got %T", uf.Emits[0])
	}
	if strings.Join(emit.Value, ".") != "generic.id" {
		t.Errorf("Emits = %v, want generic.id", emit.Value)
	}
}

func TestUnfuncMultipleRequires(t *testing.T) {
	// requires with && should produce an InfixExpression preserving the operator.
	test := `system test1;

component fetch = states{
	getWithJoin: unfunc{
		requires generic.id && generic.joinId,
		emits generic.id,
	},
};
`
	flags := map[string]bool{"specType": false}
	_, spec := prepTest(test, flags)
	if spec == nil {
		t.Fatal("prepTest() returned nil")
	}

	uf := unfuncLiteralFromSpec(t, spec)

	infix, ok := uf.Requires.(*ast.InfixExpression)
	if !ok {
		t.Fatalf("Requires is not an InfixExpression, got %T", uf.Requires)
	}
	if infix.Operator != "&&" {
		t.Errorf("operator = %q, want &&", infix.Operator)
	}
	if infix.Left.String() != "generic.joinId" && infix.Right.String() != "generic.joinId" {
		t.Error("neither side of Requires contains generic.joinId")
	}
}

func TestUnfuncMultipleEmits(t *testing.T) {
	// emits with multiple comma-separated items should produce a []Expression.
	test := `system test1;

component fetch = states{
	getDetails: unfunc{
		requires generic.id,
		emits generic.name, generic.joinId,
	},
};
`
	flags := map[string]bool{"specType": false}
	_, spec := prepTest(test, flags)
	if spec == nil {
		t.Fatal("prepTest() returned nil")
	}

	uf := unfuncLiteralFromSpec(t, spec)

	if len(uf.Emits) != 2 {
		t.Fatalf("expected 2 emits, got %d", len(uf.Emits))
	}
	names := make([]string, 2)
	for i, e := range uf.Emits {
		pc, ok := e.(*ast.ParameterCall)
		if !ok {
			t.Fatalf("Emits[%d] is not a ParameterCall, got %T", i, e)
		}
		names[i] = strings.Join(pc.Value, ".")
	}
	if names[0] != "generic.name" {
		t.Errorf("Emits[0] = %q, want generic.name", names[0])
	}
	if names[1] != "generic.joinId" {
		t.Errorf("Emits[1] = %q, want generic.joinId", names[1])
	}
}

func TestUnfuncEmitAssignment(t *testing.T) {
	// emits with explicit bool assignment: x = true, y = false
	test := `system test1;

component fetch = states{
	deactivate: unfunc{
		requires store.active,
		emits store.active = false, store.done = true,
	},
};
`
	flags := map[string]bool{"specType": false}
	_, spec := prepTest(test, flags)
	if spec == nil {
		t.Fatal("prepTest() returned nil")
	}

	uf := unfuncLiteralFromSpec(t, spec)

	if len(uf.Emits) != 2 {
		t.Fatalf("expected 2 emits, got %d", len(uf.Emits))
	}

	first, ok := uf.Emits[0].(*ast.InfixExpression)
	if !ok {
		t.Fatalf("Emits[0] is not an InfixExpression, got %T", uf.Emits[0])
	}
	if first.Operator != "=" {
		t.Errorf("Emits[0] operator = %q, want =", first.Operator)
	}
	firstRHS, ok := first.Right.(*ast.Boolean)
	if !ok {
		t.Fatalf("Emits[0] RHS is not a Boolean, got %T", first.Right)
	}
	if firstRHS.Value != false {
		t.Errorf("Emits[0] RHS = %v, want false", firstRHS.Value)
	}

	second, ok := uf.Emits[1].(*ast.InfixExpression)
	if !ok {
		t.Fatalf("Emits[1] is not an InfixExpression, got %T", uf.Emits[1])
	}
	secondRHS, ok := second.Right.(*ast.Boolean)
	if !ok {
		t.Fatalf("Emits[1] RHS is not a Boolean, got %T", second.Right)
	}
	if secondRHS.Value != true {
		t.Errorf("Emits[1] RHS = %v, want true", secondRHS.Value)
	}
}

func TestUnfuncRequiresOr(t *testing.T) {
	// requires with || should produce an InfixExpression with || operator.
	test := `system test1;

component fetch = states{
	getByNameOrId: unfunc{
		requires generic.name || generic.id,
		emits generic.id,
	},
};
`
	flags := map[string]bool{"specType": false}
	_, spec := prepTest(test, flags)
	if spec == nil {
		t.Fatal("prepTest() returned nil")
	}

	uf := unfuncLiteralFromSpec(t, spec)

	infix, ok := uf.Requires.(*ast.InfixExpression)
	if !ok {
		t.Fatalf("Requires is not an InfixExpression, got %T", uf.Requires)
	}
	if infix.Operator != "||" {
		t.Errorf("operator = %q, want ||", infix.Operator)
	}
}

func TestUnfuncRequiresNot(t *testing.T) {
	// requires with ! should produce a PrefixExpression wrapping the ParameterCall.
	test := `system test1;

component fetch = states{
	getIfNotDeleted: unfunc{
		requires !generic.deleted,
		emits generic.id,
	},
};
`
	flags := map[string]bool{"specType": false}
	_, spec := prepTest(test, flags)
	if spec == nil {
		t.Fatal("prepTest() returned nil")
	}

	uf := unfuncLiteralFromSpec(t, spec)

	prefix, ok := uf.Requires.(*ast.PrefixExpression)
	if !ok {
		t.Fatalf("Requires is not a PrefixExpression, got %T", uf.Requires)
	}
	if prefix.Operator != "!" {
		t.Errorf("operator = %q, want !", prefix.Operator)
	}
	inner, ok := prefix.Right.(*ast.ParameterCall)
	if !ok {
		t.Fatalf("inner expression is not a ParameterCall, got %T", prefix.Right)
	}
	if strings.Join(inner.Value, ".") != "generic.deleted" {
		t.Errorf("inner = %v, want generic.deleted", inner.Value)
	}
}

func TestUnfuncMixedWithFunc(t *testing.T) {
	// A component containing both a func{} state and an unfunc{} state.
	test := `system test1;

component fetch = states{
	idle: func{
		stay();
	},
	getByName: unfunc{
		requires generic.name,
		emits generic.id,
	},
};
`
	flags := map[string]bool{"specType": false}
	_, spec := prepTest(test, flags)
	if spec == nil {
		t.Fatal("prepTest() returned nil")
	}

	comp := spec.Statements[1].(*ast.DefStatement).Value.(*ast.ComponentLiteral)
	if len(comp.Pairs) != 2 {
		t.Fatalf("expected 2 component pairs, got %d", len(comp.Pairs))
	}

	var funcCount, unfuncCount int
	for _, v := range comp.Pairs {
		switch v.(type) {
		case *ast.FunctionLiteral:
			funcCount++
		case *ast.UnfuncLiteral:
			unfuncCount++
		}
	}
	if funcCount != 1 {
		t.Errorf("expected 1 FunctionLiteral, got %d", funcCount)
	}
	if unfuncCount != 1 {
		t.Errorf("expected 1 UnfuncLiteral, got %d", unfuncCount)
	}
}

func TestUnfuncInFlow(t *testing.T) {
	// An unfunc inside a flow{} struct.
	test := `spec test1;

def lookup = flow{
	getByName: unfunc{
		requires store.key,
		emits store.value,
	},
};
`
	flags := map[string]bool{"specType": true}
	_, spec := prepTest(test, flags)
	if spec == nil {
		t.Fatal("prepTest() returned nil")
	}

	var fl *ast.FlowLiteral
	for _, s := range spec.Statements {
		def, ok := s.(*ast.DefStatement)
		if !ok {
			continue
		}
		if f, ok := def.Value.(*ast.FlowLiteral); ok {
			fl = f
			break
		}
	}
	if fl == nil {
		t.Fatal("no FlowLiteral found in spec")
	}

	if len(fl.Pairs) != 1 {
		t.Fatalf("expected 1 property, got %d", len(fl.Pairs))
	}

	for _, v := range fl.Pairs {
		uf, ok := v.(*ast.UnfuncLiteral)
		if !ok {
			t.Fatalf("flow property is not UnfuncLiteral, got %T", v)
		}
		req, ok := uf.Requires.(*ast.ParameterCall)
		if !ok {
			t.Fatalf("Requires is not a ParameterCall, got %T", uf.Requires)
		}
		if strings.Join(req.Value, ".") != "store.key" {
			t.Errorf("Requires = %v, want store.key", req.Value)
		}
		if len(uf.Emits) != 1 {
			t.Fatalf("expected 1 emit, got %d", len(uf.Emits))
		}
		emit, ok := uf.Emits[0].(*ast.ParameterCall)
		if !ok {
			t.Fatalf("Emits[0] is not a ParameterCall, got %T", uf.Emits[0])
		}
		if strings.Join(emit.Value, ".") != "store.value" {
			t.Errorf("Emits = %v, want store.value", emit.Value)
		}
	}
}

func TestEmitNegation(t *testing.T) {
	// !x.field should produce an InfixExpression x.field = false
	test := `system test1;

component toggle = states{
	disable: unfunc{
		requires store.active,
		emits !store.active,
	},
};
`
	flags := map[string]bool{"specType": false}
	_, spec := prepTest(test, flags)
	if spec == nil {
		t.Fatal("prepTest() returned nil")
	}

	uf := unfuncLiteralFromSpec(t, spec)

	if len(uf.Emits) != 1 {
		t.Fatalf("expected 1 emit, got %d", len(uf.Emits))
	}
	infix, ok := uf.Emits[0].(*ast.InfixExpression)
	if !ok {
		t.Fatalf("emit is not an InfixExpression, got %T", uf.Emits[0])
	}
	if infix.Operator != "=" {
		t.Errorf("operator = %q, want =", infix.Operator)
	}
	b, ok := infix.Right.(*ast.Boolean)
	if !ok {
		t.Fatalf("RHS is not Boolean, got %T", infix.Right)
	}
	if b.Value != false {
		t.Errorf("RHS = %v, want false", b.Value)
	}
}

func TestEmitArithAssign(t *testing.T) {
	// emits x = x + 1 should produce an InfixExpression with arithmetic RHS
	test := `system test1;

component counter = states{
	increment: unfunc{
		requires store.count,
		emits store.count = store.count + 1,
	},
};
`
	flags := map[string]bool{"specType": false}
	_, spec := prepTest(test, flags)
	if spec == nil {
		t.Fatal("prepTest() returned nil")
	}

	uf := unfuncLiteralFromSpec(t, spec)

	if len(uf.Emits) != 1 {
		t.Fatalf("expected 1 emit, got %d", len(uf.Emits))
	}
	assign, ok := uf.Emits[0].(*ast.InfixExpression)
	if !ok {
		t.Fatalf("emit is not an InfixExpression, got %T", uf.Emits[0])
	}
	if assign.Operator != "=" {
		t.Errorf("operator = %q, want =", assign.Operator)
	}
	rhs, ok := assign.Right.(*ast.InfixExpression)
	if !ok {
		t.Fatalf("RHS is not an InfixExpression, got %T", assign.Right)
	}
	if rhs.Operator != "+" {
		t.Errorf("RHS operator = %q, want +", rhs.Operator)
	}
}

func TestEmitFlowAssign(t *testing.T) {
	// emits x <- 1 should desugar to InfixExpression{op:"<-", right:InfixExpression{op:"+", left:x, right:1}}
	// emits x -> 1 should desugar to InfixExpression{op:"<-", right:InfixExpression{op:"-", left:x, right:1}}
	tests := []struct {
		name     string
		src      string
		arithOp  string
	}{
		{"add", "emits store.count <- 1,", "+"},
		{"sub", "emits store.count -> 1,", "-"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			test := `system test1;

component counter = states{
	increment: unfunc{
		requires store.count,
		` + tt.src + `
	},
};
`
			flags := map[string]bool{"specType": false}
			_, spec := prepTest(test, flags)
			if spec == nil {
				t.Fatal("prepTest() returned nil")
			}
			uf := unfuncLiteralFromSpec(t, spec)
			if len(uf.Emits) != 1 {
				t.Fatalf("expected 1 emit, got %d", len(uf.Emits))
			}
			assign, ok := uf.Emits[0].(*ast.InfixExpression)
			if !ok {
				t.Fatalf("emit is not an InfixExpression, got %T", uf.Emits[0])
			}
			if assign.Operator != "<-" {
				t.Errorf("outer operator = %q, want <-", assign.Operator)
			}
			rhs, ok := assign.Right.(*ast.InfixExpression)
			if !ok {
				t.Fatalf("RHS is not an InfixExpression, got %T", assign.Right)
			}
			if rhs.Operator != tt.arithOp {
				t.Errorf("RHS operator = %q, want %q", rhs.Operator, tt.arithOp)
			}
		})
	}
}

func TestUnfuncNoRequires(t *testing.T) {
	// An unfunc with no requires clause should parse successfully with Requires == nil.
	test := `system test1;

component fetch = states{
	getAll: unfunc{
		emits generic.id,
	},
};
`
	flags := map[string]bool{"specType": false}
	_, spec := prepTest(test, flags)
	if spec == nil {
		t.Fatal("prepTest() returned nil")
	}

	uf := unfuncLiteralFromSpec(t, spec)

	if uf.Requires != nil {
		t.Errorf("Requires should be nil for unfunc with no requires clause, got %T", uf.Requires)
	}
	if len(uf.Emits) != 1 {
		t.Fatalf("expected 1 emit, got %d", len(uf.Emits))
	}
}

func TestUnfuncRequiresGrouped(t *testing.T) {
	// Parenthesized grouped expression: (a.x && b.y) || !c.z
	test := `system test1;

component fetch = states{
	getComplex: unfunc{
		requires (generic.id && generic.joinId) || !generic.deleted,
		emits generic.id,
	},
};
`
	flags := map[string]bool{"specType": false}
	_, spec := prepTest(test, flags)
	if spec == nil {
		t.Fatal("prepTest() returned nil")
	}

	uf := unfuncLiteralFromSpec(t, spec)

	outer, ok := uf.Requires.(*ast.InfixExpression)
	if !ok {
		t.Fatalf("Requires is not an InfixExpression, got %T", uf.Requires)
	}
	if outer.Operator != "||" {
		t.Errorf("outer operator = %q, want ||", outer.Operator)
	}
	_, ok = outer.Left.(*ast.InfixExpression)
	if !ok {
		t.Fatalf("LHS of || should be InfixExpression (&&), got %T", outer.Left)
	}
	_, ok = outer.Right.(*ast.PrefixExpression)
	if !ok {
		t.Fatalf("RHS of || should be PrefixExpression (!), got %T", outer.Right)
	}
}

func TestEmitLiteralAssign(t *testing.T) {
	// emits stock.x = 5 should produce InfixExpression{op:"=", right:IntegerLiteral}
	test := `system test1;

component counter = states{
	reset: unfunc{
		requires store.count,
		emits store.count = 5,
	},
};
`
	flags := map[string]bool{"specType": false}
	_, spec := prepTest(test, flags)
	if spec == nil {
		t.Fatal("prepTest() returned nil")
	}

	uf := unfuncLiteralFromSpec(t, spec)

	if len(uf.Emits) != 1 {
		t.Fatalf("expected 1 emit, got %d", len(uf.Emits))
	}
	assign, ok := uf.Emits[0].(*ast.InfixExpression)
	if !ok {
		t.Fatalf("emit is not an InfixExpression, got %T", uf.Emits[0])
	}
	if assign.Operator != "=" {
		t.Errorf("operator = %q, want =", assign.Operator)
	}
	_, ok = assign.Right.(*ast.IntegerLiteral)
	if !ok {
		t.Fatalf("RHS is not an IntegerLiteral, got %T", assign.Right)
	}
}

func TestEmitFieldToFieldAssign(t *testing.T) {
	// emits stock.x = stock.y should produce InfixExpression{op:"=", right:ParameterCall}
	test := `system test1;

component cache = states{
	copy: unfunc{
		requires store.src,
		emits store.dst = store.src,
	},
};
`
	flags := map[string]bool{"specType": false}
	_, spec := prepTest(test, flags)
	if spec == nil {
		t.Fatal("prepTest() returned nil")
	}

	uf := unfuncLiteralFromSpec(t, spec)

	if len(uf.Emits) != 1 {
		t.Fatalf("expected 1 emit, got %d", len(uf.Emits))
	}
	assign, ok := uf.Emits[0].(*ast.InfixExpression)
	if !ok {
		t.Fatalf("emit is not an InfixExpression, got %T", uf.Emits[0])
	}
	if assign.Operator != "=" {
		t.Errorf("operator = %q, want =", assign.Operator)
	}
	rhs, ok := assign.Right.(*ast.ParameterCall)
	if !ok {
		t.Fatalf("RHS is not a ParameterCall, got %T", assign.Right)
	}
	if strings.Join(rhs.Value, ".") != "store.src" {
		t.Errorf("RHS = %v, want store.src", rhs.Value)
	}
}

func TestEmitFieldArithmetic(t *testing.T) {
	// emits stock.x = stock.y + stock.z — two field operands
	test := `system test1;

component adder = states{
	add: unfunc{
		requires store.a,
		emits store.result = store.a + store.b,
	},
};
`
	flags := map[string]bool{"specType": false}
	_, spec := prepTest(test, flags)
	if spec == nil {
		t.Fatal("prepTest() returned nil")
	}

	uf := unfuncLiteralFromSpec(t, spec)

	assign, ok := uf.Emits[0].(*ast.InfixExpression)
	if !ok {
		t.Fatalf("emit is not an InfixExpression, got %T", uf.Emits[0])
	}
	rhs, ok := assign.Right.(*ast.InfixExpression)
	if !ok {
		t.Fatalf("RHS is not an InfixExpression, got %T", assign.Right)
	}
	if rhs.Operator != "+" {
		t.Errorf("RHS operator = %q, want +", rhs.Operator)
	}
	_, ok = rhs.Left.(*ast.ParameterCall)
	if !ok {
		t.Fatalf("RHS.Left is not a ParameterCall, got %T", rhs.Left)
	}
	_, ok = rhs.Right.(*ast.ParameterCall)
	if !ok {
		t.Fatalf("RHS.Right is not a ParameterCall, got %T", rhs.Right)
	}
}

func TestEmitFlowAssignByField(t *testing.T) {
	// emits x <- stock.a and emits x -> stock.a — increment/decrement by field value
	tests := []struct {
		name    string
		src     string
		arithOp string
	}{
		{"increment_by_field", "emits store.count <- store.delta,", "+"},
		{"decrement_by_field", "emits store.count -> store.delta,", "-"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			test := `system test1;

component counter = states{
	adjust: unfunc{
		requires store.count,
		` + tt.src + `
	},
};
`
			flags := map[string]bool{"specType": false}
			_, spec := prepTest(test, flags)
			if spec == nil {
				t.Fatal("prepTest() returned nil")
			}
			uf := unfuncLiteralFromSpec(t, spec)
			assign, ok := uf.Emits[0].(*ast.InfixExpression)
			if !ok {
				t.Fatalf("emit is not an InfixExpression, got %T", uf.Emits[0])
			}
			if assign.Operator != "<-" {
				t.Errorf("outer operator = %q, want <-", assign.Operator)
			}
			rhs, ok := assign.Right.(*ast.InfixExpression)
			if !ok {
				t.Fatalf("RHS is not an InfixExpression, got %T", assign.Right)
			}
			if rhs.Operator != tt.arithOp {
				t.Errorf("RHS operator = %q, want %q", rhs.Operator, tt.arithOp)
			}
			_, ok = rhs.Right.(*ast.ParameterCall)
			if !ok {
				t.Fatalf("RHS.Right is not a ParameterCall, got %T", rhs.Right)
			}
		})
	}
}

func TestEmitMultipleLiterals(t *testing.T) {
	// emits stock.x = 5, stock.y = 10 — multiple targets with literal values
	test := `system test1;

component setup = states{
	configure: unfunc{
		emits store.x = 5, store.y = 10,
	},
};
`
	flags := map[string]bool{"specType": false}
	_, spec := prepTest(test, flags)
	if spec == nil {
		t.Fatal("prepTest() returned nil")
	}

	uf := unfuncLiteralFromSpec(t, spec)

	if len(uf.Emits) != 2 {
		t.Fatalf("expected 2 emits, got %d", len(uf.Emits))
	}
	for i, e := range uf.Emits {
		assign, ok := e.(*ast.InfixExpression)
		if !ok {
			t.Fatalf("Emits[%d] is not an InfixExpression, got %T", i, e)
		}
		if assign.Operator != "=" {
			t.Errorf("Emits[%d] operator = %q, want =", i, assign.Operator)
		}
		if _, ok := assign.Right.(*ast.IntegerLiteral); !ok {
			t.Fatalf("Emits[%d] RHS is not an IntegerLiteral, got %T", i, assign.Right)
		}
	}
}

// --- invalid syntax tests ---

func assertUnfuncParseError(t *testing.T, src string) {
	t.Helper()
	flags := map[string]bool{"specType": false, "testing": true}
	_, err := Execute(src, "", flags)
	if err == nil {
		t.Fatal("expected parse error, got nil")
	}
}

func TestRequiresTrueParseFails(t *testing.T) {
	assertUnfuncParseError(t, `system test1;
component fetch = states{
	op: unfunc{
		requires true,
		emits generic.id,
	},
};`)
}

func TestRequiresFalseParseFails(t *testing.T) {
	assertUnfuncParseError(t, `system test1;
component fetch = states{
	op: unfunc{
		requires false,
		emits generic.id,
	},
};`)
}

func TestRequiresNumericParseFails(t *testing.T) {
	assertUnfuncParseError(t, `system test1;
component fetch = states{
	op: unfunc{
		requires 5,
		emits generic.id,
	},
};`)
}

func TestRequiresStringParseFails(t *testing.T) {
	assertUnfuncParseError(t, `system test1;
component fetch = states{
	op: unfunc{
		requires "hello",
		emits generic.id,
	},
};`)
}

func TestRequiresSolvableParseFails(t *testing.T) {
	assertUnfuncParseError(t, `system test1;
component fetch = states{
	op: unfunc{
		requires unknown(),
		emits generic.id,
	},
};`)
}

func TestEmitStringRHSParseFails(t *testing.T) {
	assertUnfuncParseError(t, `system test1;
component fetch = states{
	op: unfunc{
		requires generic.id,
		emits store.x = "hello",
	},
};`)
}

func TestEmitSolvableRHSParseFails(t *testing.T) {
	assertUnfuncParseError(t, `system test1;
component fetch = states{
	op: unfunc{
		requires generic.id,
		emits store.x = uncertain(0, 1),
	},
};`)
}

func TestEmitLogicalOpParseFails(t *testing.T) {
	assertUnfuncParseError(t, `system test1;
component fetch = states{
	op: unfunc{
		requires generic.id,
		emits store.x && store.y,
	},
};`)
}

func TestEmitComparisonRHSParseFails(t *testing.T) {
	assertUnfuncParseError(t, `system test1;
component fetch = states{
	op: unfunc{
		requires generic.id,
		emits store.x = store.y < 5,
	},
};`)
}

func TestEmitNegationLiteralParseFails(t *testing.T) {
	assertUnfuncParseError(t, `system test1;
component fetch = states{
	op: unfunc{
		requires generic.id,
		emits !5,
	},
};`)
}

func TestEmitBareTrueParseFails(t *testing.T) {
	assertUnfuncParseError(t, `system test1;
component fetch = states{
	op: unfunc{
		requires generic.id,
		emits true,
	},
};`)
}
