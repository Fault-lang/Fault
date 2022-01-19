package types

import (
	"fault/ast"
	"fault/listener"
	"fault/parser"
	"testing"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

//// TEST STOCKFLOW FUNCTIONS FIRST ////
func TestSFAdd(t *testing.T) {
	s := StockFlow{}
	s.Add("test", "bar", &ast.Nil{})

	if _, ok := s["test"]; !ok {
		t.Fatal("struct test not added to StockFlow")
	}

	if _, ok := s["test"]["bar"]; !ok {
		t.Fatal("parameter bar not added to StockFlow")
	}

}

func TestSFBulk(t *testing.T) {
	s := StockFlow{}
	pairs := make(map[string]ast.Node)
	pairs["bar1"] = &ast.Nil{}
	pairs["bar2"] = &ast.Nil{}
	s.Bulk("test", pairs)

	if _, ok := s["test"]; !ok {
		t.Fatal("struct test not added to StockFlow")
	}

	if _, ok := s["test"]["bar1"]; !ok {
		t.Fatal("parameter bar1 not added to StockFlow")
	}

	if _, ok := s["test"]["bar2"]; !ok {
		t.Fatal("parameter bar2 not added to StockFlow")
	}

}

func TestSFGet(t *testing.T) {
	s := StockFlow{}
	s.Add("test", "bar", &ast.Nil{})
	n := s.Get("test", "bar")

	if n == nil {
		t.Fatal("struct test not added to StockFlow")
	}

	if _, ok := n.(*ast.Nil); !ok {
		t.Fatal("StockFlow did not return the correct node")
	}
}

func TestSFGetStruct(t *testing.T) {
	s := StockFlow{}
	s.Add("test", "bar", &ast.Nil{})
	n := s.GetStruct("test")

	if n == nil {
		t.Fatal("struct test not added to StockFlow")
	}

	if _, ok := n["bar"]; !ok {
		t.Fatal("StockFlow did not return a valid struct")
	}

}

func TestImportTrail(t *testing.T) {
	it := importTrail{}
	it = it.PushSpec("test")
	it = it.PushSpec("this")
	it = it.PushSpec("trail")

	if len(it) != 3 {
		t.Fatal("specs not added to trail correctly")
	}

	i, it2 := it.PopSpec()
	if i != "trail" {
		t.Fatalf("trail entry incorrect. got=%s, want=trail", i)
	}

	if len(it2) != 2 {
		t.Fatal("specs not popped off trail correctly")
	}
}

func TestAddOK(t *testing.T) {
	test := `spec test1;
			const x = 2+2;
			const y = 2+3.1;
	`
	checker, err := prepTest(test)

	consts := checker.Constants["test1"]

	if err != nil {
		t.Fatalf("Type checking failed on valid expression. got=%s", err)
	}

	if consts["x"].(*ast.InfixExpression).InferredType.Type != "INT" {
		t.Fatalf("Constant x does not have an int type. got=%T", consts["x"])
	}

	if consts["x"].(*ast.InfixExpression).Right.(*ast.IntegerLiteral).InferredType.Type != "INT" {
		t.Fatalf("right node does not have an int type. got=%T", consts["x"].(*ast.InfixExpression).Right)
	}

	if consts["x"].(*ast.InfixExpression).Left.(*ast.IntegerLiteral).InferredType.Type != "INT" {
		t.Fatalf("left node does not have an int type. got=%T", consts["x"].(*ast.InfixExpression).Left)
	}

	if consts["y"].(*ast.InfixExpression).InferredType.Type != "FLOAT" {
		t.Fatalf("Constant y does not have a float type. got=%T", consts["y"])
	}

	if consts["y"].(*ast.InfixExpression).Right.(*ast.FloatLiteral).InferredType.Type != "FLOAT" {
		t.Fatalf("right y node does not have an int type. got=%T", consts["y"].(*ast.InfixExpression).Right)
	}

	if consts["y"].(*ast.InfixExpression).Left.(*ast.IntegerLiteral).InferredType.Type != "INT" {
		t.Fatalf("left y node does not have an int type. got=%T", consts["y"].(*ast.InfixExpression).Left)
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

func TestStructTypeError(t *testing.T) {
	test := `spec test1;
			def foo = stock{
				bar: 5,
			};

			def fizz = stock{
				buzz: new foo,
				bash: func{
					buzz.bar <- 2;
				},
			};
	`
	_, err := prepTest(test)
	//sym := checker.SymbolTypes

	actual := "stock is the store of values please use a flow for test1.fizz"

	if err.Error() != actual {
		t.Fatalf("Type checking failed to catch invalid expression. got=%s", err)
	}

}

func TestInstanceError(t *testing.T) {
	test := `spec test1;
			def foo = stock{
				bar: 5,
			};

			def fizz = flow{
				buzz: new foo,
				bash: func{
					buzz <- 2;
				},
			};
	`
	_, err := prepTest(test)
	//sym := checker.SymbolTypes

	actual := "struct buzz missing property, line:7, col:10"

	if err.Error() != actual {
		t.Fatalf("Type checking failed to catch invalid expression. got=%s", err)
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
				fizz2: -a,
			};

			def zoo = flow{
				con: new foo,
				rate: func{
					con.foosh + 2;
				},
				rate2: func{
					2 - a;
					a - 2;
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

	if fooStock["fizz2"].(*ast.PrefixExpression).InferredType.Type != "FLOAT" {
		t.Fatalf("stock property not typed correctly. got=%s", fooStock["fizz2"].(*ast.PrefixExpression).InferredType.Type)
	}

	zooFlow, ok := str["zoo"]
	if !ok {
		t.Fatal("flow zoo not stored in symbol table correctly.")
	}

	if zooFlow["con"].(*ast.Instance).InferredType.Type != "STOCK" {
		t.Fatalf("flow property not typed correctly. got=%s", zooFlow["con"].(*ast.Instance).InferredType.Type)
	}

	if zooFlow["rate"].(*ast.BlockStatement).InferredType.Type != "INT" {
		t.Fatalf("flow property not typed correctly. got=%s", zooFlow["rate"].(*ast.BlockStatement).InferredType.Type)
	}

	infix, ok := zooFlow["rate"].(*ast.BlockStatement).Statements[0].(*ast.ExpressionStatement).Expression.(*ast.InfixExpression)
	if !ok {
		t.Fatalf("expecting a infix expression. got=%T", zooFlow["rate"].(*ast.BlockStatement).Statements[0].(*ast.ExpressionStatement).Expression)
	}
	if infix.Right.(*ast.IntegerLiteral).InferredType.Type != "INT" {
		t.Fatalf("flow property not typed correctly. got=%s", infix.Right.(*ast.IntegerLiteral).InferredType.Type)
	}

	if infix.Left.(*ast.ParameterCall).InferredType.Type != "INT" {
		t.Fatalf("flow property not typed correctly. got=%s", infix.Left.(*ast.ParameterCall).InferredType.Type)
	}

	if zooFlow["rate2"].(*ast.BlockStatement).InferredType.Type != "FLOAT" {
		t.Fatalf("flow property not typed correctly. got=%s", zooFlow["rate2"].(*ast.BlockStatement).InferredType.Type)
	}

	infix2, ok := zooFlow["rate2"].(*ast.BlockStatement).Statements[1].(*ast.ExpressionStatement).Expression.(*ast.InfixExpression)
	if !ok {
		t.Fatalf("expecting a infix expression. got=%T", zooFlow["rate2"].(*ast.BlockStatement).Statements[1].(*ast.ExpressionStatement).Expression)
	}
	if infix2.Right.(*ast.IntegerLiteral).InferredType.Type != "INT" {
		t.Fatalf("flow property not typed correctly. got=%s", infix2.Right.(*ast.IntegerLiteral).InferredType.Type)
	}

	if infix2.Left.(*ast.Identifier).InferredType.Type != "FLOAT" {
		t.Fatalf("flow property not typed correctly. got=%s", infix2.Left.(*ast.Identifier).InferredType.Type)
	}

	infix3, ok := zooFlow["rate2"].(*ast.BlockStatement).Statements[0].(*ast.ExpressionStatement).Expression.(*ast.InfixExpression)
	if !ok {
		t.Fatalf("expecting a infix expression. got=%T", zooFlow["rate2"].(*ast.BlockStatement).Statements[0].(*ast.ExpressionStatement).Expression)
	}
	if infix3.Right.(*ast.Identifier).InferredType.Type != "FLOAT" {
		t.Fatalf("flow property not typed correctly. got=%s", infix3.Right.(*ast.Identifier).InferredType.Type)
	}

	if infix3.Left.(*ast.IntegerLiteral).InferredType.Type != "INT" {
		t.Fatalf("flow property not typed correctly. got=%s", infix3.Left.(*ast.IntegerLiteral).InferredType.Type)
	}
}

func TestNils(t *testing.T) {
	test := `spec test1;
			const x = nil + 3;
			const y = 4 + nil;
			const z = nil + nil;`
	checker, err := prepTest(test)

	if err != nil {
		t.Fatalf("Type checking failed on valid expression. got=%s", err)
	}

	consts := checker.Constants["test1"]

	if consts["x"].(*ast.InfixExpression).InferredType.Type != "INT" {
		t.Fatalf("Constant x does not have an int type. got=%s", consts["x"].(*ast.InfixExpression).InferredType.Type)
	}

	if consts["x"].(*ast.InfixExpression).Right.(*ast.IntegerLiteral).InferredType.Type != "INT" {
		t.Fatalf("x right node does not have an int type. got=%s", consts["x"].(*ast.InfixExpression).Right.(*ast.IntegerLiteral).InferredType.Type)
	}

	if consts["x"].(*ast.InfixExpression).Left.(*ast.Nil).InferredType.Type != "NIL" {
		t.Fatalf("x left node does not have an nil type. got=%s", consts["x"].(*ast.InfixExpression).Left.(*ast.Nil).InferredType.Type)
	}

	if consts["y"].(*ast.InfixExpression).InferredType.Type != "INT" {
		t.Fatalf("Constant y does not have an int type. got=%s", consts["y"].(*ast.InfixExpression).InferredType.Type)
	}

	if consts["y"].(*ast.InfixExpression).Right.(*ast.Nil).InferredType.Type != "NIL" {
		t.Fatalf("y right node does not have an nil type. got=%s", consts["y"].(*ast.InfixExpression).Right.(*ast.Nil).InferredType.Type)
	}

	if consts["y"].(*ast.InfixExpression).Left.(*ast.IntegerLiteral).InferredType.Type != "INT" {
		t.Fatalf("y left node does not have an int type. got=%s", consts["y"].(*ast.InfixExpression).Left.(*ast.IntegerLiteral).InferredType.Type)
	}

	if consts["z"].(*ast.InfixExpression).InferredType.Type != "NIL" {
		t.Fatalf("Constant z does not have a nil type. got=%s", consts["z"].(*ast.InfixExpression).InferredType.Type)
	}

	if consts["z"].(*ast.InfixExpression).Right.(*ast.Nil).InferredType.Type != "NIL" {
		t.Fatalf("z right node does not have a nil type. got=%s", consts["z"].(*ast.InfixExpression).Right.(*ast.Nil).InferredType.Type)
	}

	if consts["z"].(*ast.InfixExpression).Left.(*ast.Nil).InferredType.Type != "NIL" {
		t.Fatalf("z left node does not have a nil type. got=%s", consts["z"].(*ast.InfixExpression).Left.(*ast.Nil).InferredType.Type)
	}
}

func TestInConditionals(t *testing.T) {
	test := `spec test1;
			def foo = stock{
				foosh: 3,
				bar: "hello!",
			};

			def zoo = flow{
				con: new foo,
				rate: func{
					if con.foosh == 3 {
						2;
					}else if con.foosh == 5{
						false;
					}else{
						nil;
					}
				},
				rate2: func{
					if con.foosh == 3 {
						2+2;
					}else if con.foosh == 5{
						7*8;
					}else{
						5.333 / 3 * 2;
					}
				},
			};
	`
	checker, err := prepTest(test)

	if err != nil {
		t.Fatalf("Type checking failed on valid expression. got=%s", err)
	}

	str := checker.SpecStructs["test1"]

	zooFlow, ok := str["zoo"]
	if !ok {
		t.Fatal("flow zoo not stored in symbol table correctly.")
	}

	if zooFlow["rate"].(*ast.BlockStatement).InferredType.Type != "INT" {
		t.Fatalf("flow property not typed correctly. got=%s", zooFlow["rate"].(*ast.BlockStatement).InferredType.Type)
	}

	if zooFlow["rate"].(*ast.BlockStatement).Statements[0].(*ast.ExpressionStatement).InferredType.Type != "INT" {
		t.Fatalf("flow property not typed correctly. got=%s", zooFlow["rate"].(*ast.BlockStatement).Statements[0].(*ast.ExpressionStatement).InferredType.Type)
	}

	ife, ok := zooFlow["rate"].(*ast.BlockStatement).Statements[0].(*ast.ExpressionStatement).Expression.(*ast.IfExpression)
	if !ok {
		t.Fatalf("expecting a If expression. got=%T", zooFlow["rate"].(*ast.BlockStatement).Statements[0].(*ast.ExpressionStatement).Expression)
	}

	if ife.InferredType.Type != "INT" {
		t.Fatalf("if expression not typed correctly. got=%s", ife.InferredType.Type)
	}

	if ife.Condition.(*ast.InfixExpression).InferredType.Type != "BOOL" {
		t.Fatalf("if condition not typed correctly. got=%s", ife.Condition.(*ast.InfixExpression).InferredType.Type)
	}

	if ife.Consequence.InferredType.Type != "INT" {
		t.Fatalf("if consequence block not typed correctly. got=%s", ife.Consequence.InferredType.Type)
	}

	if ife.Elif.InferredType.Type != "BOOL" {
		t.Fatalf("if else if block not typed correctly. got=%s", ife.Elif.InferredType.Type)
	}

	if ife.Elif.Condition.(*ast.InfixExpression).InferredType.Type != "BOOL" {
		t.Fatalf("if else if condition not typed correctly. got=%s", ife.Elif.Condition.(*ast.InfixExpression).InferredType.Type)
	}

	if ife.Elif.Consequence.InferredType.Type != "BOOL" {
		t.Fatalf("if else if consequence block not typed correctly. got=%s", ife.Elif.Consequence.InferredType.Type)
	}

	if ife.Elif.Alternative.InferredType.Type != "NIL" {
		t.Fatalf("if alternative block not typed correctly. got=%s", ife.Elif.Alternative.InferredType.Type)
	}

	ife2, ok := zooFlow["rate2"].(*ast.BlockStatement).Statements[0].(*ast.ExpressionStatement).Expression.(*ast.IfExpression)
	if !ok {
		t.Fatalf("expecting a If expression. got=%T", zooFlow["rate2"].(*ast.BlockStatement).Statements[0].(*ast.ExpressionStatement).Expression)
	}

	if ife2.InferredType.Type != "INT" {
		t.Fatalf("if expression not typed correctly. got=%s", ife2.InferredType.Type)
	}

	if ife2.Condition.(*ast.InfixExpression).InferredType.Type != "BOOL" {
		t.Fatalf("if condition not typed correctly. got=%s", ife2.Condition.(*ast.InfixExpression).InferredType.Type)
	}

	if ife2.Consequence.InferredType.Type != "INT" {
		t.Fatalf("if consequence block not typed correctly. got=%s", ife2.Consequence.InferredType.Type)
	}

	if ife2.Elif.InferredType.Type != "INT" {
		t.Fatalf("if else if block not typed correctly. got=%s", ife2.Elif.InferredType.Type)
	}

	if ife2.Elif.Condition.(*ast.InfixExpression).InferredType.Type != "BOOL" {
		t.Fatalf("if else if condition not typed correctly. got=%s", ife2.Elif.Condition.(*ast.InfixExpression).InferredType.Type)
	}

	if ife2.Elif.Consequence.InferredType.Type != "INT" {
		t.Fatalf("if else if consequence block not typed correctly. got=%s", ife2.Elif.Consequence.InferredType.Type)
	}

	if ife2.Elif.Alternative.InferredType.Type != "FLOAT" {
		t.Fatalf("if alternative block not typed correctly. got=%s", ife2.Elif.Alternative.InferredType.Type)
	}

}

func TestComplexStruct(t *testing.T) {
	test := `spec test1;
			def str = stock{
				foo: 3,
			};

			def str2 = stock{
				bar: new str,
			};

			def str3 = flow{
				buzz: new str2,
				fizz: func{
					buzz.bar.foo <- 5;
				},
				bash: new str,
				foosh: func{
					bash.foo <- 5;
				},
			};
	`
	checker, err := prepTest(test)

	if err != nil {
		t.Fatalf("Type checking failed on valid expression. got=%s", err)
	}

	str := checker.SpecStructs["test1"]

	_, ok := str["str2"]["___base"].(*ast.StockLiteral)
	if !ok {
		t.Fatalf("struct not a stock has wrong type. got=%T", str["str2"]["___base"])
	}

	inst, ok := str["str2"]["bar"].(*ast.Instance)
	if !ok {
		t.Fatalf("property is not an instance. got=%T", str["str2"]["bar"])
	}

	if inst.InferredType.Type != "STOCK" {
		t.Fatalf("instance has wrong type. got=%s", inst.InferredType.Type)
	}

	if !inst.Complex {
		t.Fatalf("instance should be complex")
	}

	fl, ok := str["str3"]["___base"].(*ast.FlowLiteral)
	if !ok {
		t.Fatalf("struct not a flow has wrong type. got=%T", str["str3"]["___base"])
	}

	if fl.InferredType.Type != "FLOW" {
		t.Fatalf("flow has wrong type. got=%s", fl.InferredType.Type)
	}

	inst2, ok := str["str3"]["buzz"].(*ast.Instance)
	if !ok {
		t.Fatalf("property is not an instance. got=%T", str["str3"]["buzz"])
	}

	if inst2.InferredType.Type != "STOCK" {
		t.Fatalf("instance has wrong type. got=%s", inst2.InferredType.Type)
	}

	if inst2.Complex {
		t.Fatalf("instance not should be complex")
	}

	inst3, ok := str["str3"]["bash"].(*ast.Instance)
	if !ok {
		t.Fatalf("property is not an instance. got=%T", str["str3"]["bash"])
	}

	if inst3.InferredType.Type != "STOCK" {
		t.Fatalf("instance has wrong type. got=%s", inst3.InferredType.Type)
	}

	if inst3.Complex {
		t.Fatalf("instance not should be complex")
	}
}

func TestReallyComplexStruct(t *testing.T) {
	test := `spec test1;
			def str = stock{
				foo: 3,
			};

			def str2 = stock{
				bar: new str,
			};

			def str3 = stock{
				foosh: new str2,
			};

			def fl = flow{
				buzz: new str3,
				fizz: func{
					buzz.foosh.bar.foo <- 5;
				},
			};
	`
	checker, err := prepTest(test)

	if err != nil {
		t.Fatalf("Type checking failed on valid expression. got=%s", err)
	}

	str := checker.SpecStructs["test1"]

	inst, ok := str["str3"]["foosh"].(*ast.Instance)
	if !ok {
		t.Fatalf("property is not an instance. got=%T", str["str3"]["foosh"])
	}

	if inst.InferredType.Type != "STOCK" {
		t.Fatalf("instance has wrong type. got=%s", inst.InferredType.Type)
	}

	if !inst.Complex {
		t.Fatalf("instance should be complex")
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

func TestInvalidAssert2(t *testing.T) {
	test := `spec test1;
			const a = 2.3;

			assert 5 + a;
	`
	_, err := prepTest(test)
	//sym := checker.SymbolTypes

	actual := "assert statement not testing a Boolean expression. got=FLOAT"

	if err.Error() != actual {
		t.Fatalf("Type checking failed to catch invalid expression. got=%s", err)
	}

}

func TestInvalidAssert3(t *testing.T) {
	test := `spec test1;
			const a = 2.3;

			assert true + a;
	`
	_, err := prepTest(test)
	//sym := checker.SymbolTypes

	actual := "assert statement not testing a Boolean expression. got=FLOAT"

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

func TestInvalidInfix(t *testing.T) {
	test := `spec test1;
			const a = 2 + "world";
	`
	_, err := prepTest(test)
	//sym := checker.SymbolTypes

	actual := "type mismatch: got=INT,STRING"

	if err.Error() != actual {
		t.Fatalf("Type checking failed to catch invalid expression. got=%s", err)
	}

}

func TestInvalidInfix2(t *testing.T) {
	test := `spec test1;
			const a = "hello" + 4;
	`
	_, err := prepTest(test)
	//sym := checker.SymbolTypes

	actual := "type mismatch: got=STRING,INT"

	if err.Error() != actual {
		t.Fatalf("Type checking failed to catch invalid expression. got=%s", err)
	}

}

func TestValidCompoundAssert(t *testing.T) {
	test := `spec test1;
			const a = 2.3;
			
			assert a > 5 && b == 4 || c != "hello!";
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

	float, ok := consts["a"].(*ast.PrefixExpression).Right.(*ast.FloatLiteral)
	if !ok {
		t.Fatalf("prefix base is not a float. got=%T", consts["a"].(*ast.PrefixExpression).Right)
	}

	if float.InferredType.Type != "FLOAT" {
		t.Fatalf("Prefix base does not have a float type. got=%s", float.InferredType.Type)
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
