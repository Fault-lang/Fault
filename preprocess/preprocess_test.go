package preprocess

import (
	"fault/ast"
	"fault/listener"
	"fault/parser"
	"testing"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func TestConstants(t *testing.T) {
	test := `spec test1;
	const x = 2;
	const y = 2+3.1;
	const z = unknown(a);`

	process := prepTest(test)

	consts := process.Specs["test1"]

	x := consts.FetchConstant("x")

	if _, ok := x.(*ast.IntegerLiteral); !ok {
		t.Fatalf("Constant x does not have the right type. got=%T", x)
	}

	if x.(*ast.IntegerLiteral).Value != 2 {
		t.Fatalf("Constant x does not have the right value. got=%d", x)
	}

	y := consts.FetchConstant("y")
	if _, ok := y.(*ast.InfixExpression); !ok {
		t.Fatalf("Constant y not the right type. got=%T", y)
	}

	if _, ok1 := y.(*ast.InfixExpression).Right.(*ast.FloatLiteral); !ok1 {
		t.Fatalf("right y node does not have the right type. got=%T", y.(*ast.InfixExpression).Right)
	}

	if _, ok2 := y.(*ast.InfixExpression).Left.(*ast.IntegerLiteral); !ok2 {
		t.Fatalf("left y node does not have the right type. got=%T", y.(*ast.InfixExpression).Left)
	}

	z := consts.FetchConstant("z")
	if _, ok3 := z.(*ast.Unknown); !ok3 {
		t.Fatalf("Constant z does not have the right type. got=%T", z)
	}
}

func TestStructDef(t *testing.T) {
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

	process := prepTest(test)

	variables := process.Specs["test1"]

	foo := variables.FetchStock("foo")
	if len(foo) != 4 {
		t.Fatalf("stock foo returns the wrong number of properties got=%d want=4", len(foo))
	}

	fizz := foo["fizz"].(*ast.Identifier).RawId()
	if len(fizz) != 2 || fizz[0] != "test1" || fizz[1] != "a" {
		t.Fatalf("identifier not converted to correct context got=%s", fizz)
	}

	zoo := variables.FetchFlow("zoo")
	if len(zoo) != 3 {
		t.Fatalf("flow zoo returns the wrong number of properties got=%d want=4", len(zoo))
	}

	infix := zoo["rate2"].(*ast.FunctionLiteral).Body.Statements[0].(*ast.ExpressionStatement).Expression.(*ast.InfixExpression)
	if infix.Right.(*ast.Identifier).IdString() != "test1_a" {
		t.Fatalf("variable a has the wrong scope, got=%s", infix.Right.(*ast.Identifier).IdString())
	}

}

func TestComponent(t *testing.T) {
	test := `system test;

	component foo = states{
		x: 8,
		initial: func{
			if this.x > 10{
				stay();
			}else{
				advance(this.alarm);
			}
		},
		alarm: func{
			advance(this.close);
		},
	};

	start { 
		foo: initial,
	};`

	process := prepSysTest(test)
	variables := process.Specs["test"]
	foo := variables.FetchComponent("foo")
	if foo == nil {
		t.Fatal("component named foo not found")
	}

	if len(foo) != 3 {
		t.Fatalf("component foo returns the wrong number of states got=%d want=3", len(foo))
	}

	ifcond := foo["initial"].(*ast.FunctionLiteral).Body.Statements[0].(*ast.ExpressionStatement).Expression.(*ast.IfExpression).Condition.(*ast.InfixExpression)
	this := ifcond.Left.(*ast.This).RawId()
	if len(this) != 3 || this[0] != "test" || this[1] != "foo" || this[2] != "x" {
		t.Fatalf("this special word not converted to correct context got=%s", this)
	}

}

func TestInstances(t *testing.T) {
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

	process := prepTest(test)

	variables := process.Specs["test1"]

	bar := variables.FetchStock("str2_bar")
	if bar == nil {
		t.Fatal("stock names str2_bar not found")
	}

	if len(bar) != 1 {
		t.Fatalf("stock str2_bar returns the wrong number of properties got=%d want=4", len(bar))
	}

	zoo := variables.FetchStock("fl_buzz_foosh_bar")
	if zoo == nil {
		t.Fatal("stock named fl_buzz_foosh_bar not found")
	}

	if len(zoo) != 1 {
		t.Fatalf("flow fl_buzz_foosh_bar returns the wrong number of properties got=%d want=4", len(zoo))
	}

	o := variables.FetchOrder()
	if len(o) != 10 {
		t.Fatalf("wrong number of instances got=%d want=10", len(o))
	}

	if o[0][0] != "STOCK" {
		t.Fatalf("instance has the wrong type in order got=%s want=STOCK", o[0][0])
	}

	if o[8][1] != "fl_buzz_foosh" {
		t.Fatalf("instance has the wrong name in order got=%s want=fl_buzz_foosh", o[8][1])
	}

}

func TestRunInstances(t *testing.T) {
	test := `spec test1;
	def str = stock{
		foo: 3,
	};

	def fl = flow{
		buzz: new str,
		fizz: func{
			buzz.foo <- 5;
		},
	};

	for 5 run {
		f =  new fl;
		f.fizz;
	}
	`

	process := prepTest(test)

	variables := process.Specs["test1"]

	fl := variables.FetchFlow("f")
	if fl == nil {
		t.Fatal("flow named f not found")
	}

	if len(fl) != 2 {
		t.Fatalf("flow f returns the wrong number of properties got=%d want=3", len(fl))
	}

	pc := process.Processed.(*ast.Spec).Statements[3].(*ast.ForStatement).Body.Statements[1].(*ast.ParallelFunctions).Expressions[0].(*ast.ParameterCall)
	if pc.IdString() != "test1_f_fizz" {
		t.Fatalf("flow not correctly named in runblock got=%s", pc.IdString())
	}

	pairs := process.Processed.(*ast.Spec).Statements[2].(*ast.DefStatement).Value.(*ast.FlowLiteral).Pairs
	var buzz *ast.StructInstance
	for k, v := range pairs {
		if k.Value == "buzz" {
			buzz = v.(*ast.StructInstance)
			break
		}
	}

	if buzz.Properties["foo"].IdString() != "test1_fl_buzz_foo" {
		t.Fatalf("flow property not correctly named in runblock got=%s", buzz.Properties["foo"].IdString())
	}

	if buzz.Properties["foo"].Value.(ast.Nameable).IdString() != "test1_fl_buzz_foo" {
		t.Fatalf("flow property value not correctly named in runblock got=%s", buzz.Properties["foo"].Value.(ast.Nameable).IdString())
	}

	run := process.Processed.(*ast.Spec).Statements[3].(*ast.ForStatement).Body.Statements[0].(*ast.ExpressionStatement).Expression.(*ast.StructInstance)
	for k, v := range run.Properties {
		if k == "fizz" {
			foo := v.Value.(*ast.FunctionLiteral).Body.Statements[0].(*ast.ExpressionStatement).Expression.(*ast.InfixExpression).Left.(*ast.ParameterCall).IdString()
			if foo != "test1_f_buzz_foo" {
				t.Fatalf("inner variable called from the runblock not named correctly got=%s", foo)
			}
		}

	}
}

func TestIds(t *testing.T) {
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

	process := prepTest(test)
	tree := process.Processed
	spec := tree.(*ast.Spec).Statements

	str1 := spec[1].(*ast.DefStatement).Name.RawId()
	if str1[0] != "test1" || str1[1] != "str" {
		t.Fatalf("struct1 name not correct got=%s", str1)
	}

	str2 := spec[2].(*ast.DefStatement).Name.RawId()
	if str2[0] != "test1" || str2[1] != "str2" {
		t.Fatalf("struct2 name not correct got=%s", str2)
	}

	str2f := spec[2].(*ast.DefStatement).Value.(*ast.StockLiteral).Pairs
	for k, v := range str2f {
		keyId := k.RawId()
		valId := v.(*ast.StructInstance).RawId()
		if valId[0] != keyId[0] || valId[1] != keyId[1] || keyId[2] != valId[2] {
			t.Fatalf("key id and val id do not match key=%s value=%s", keyId, valId)
		}
	}

	str3 := spec[3].(*ast.DefStatement).Name.RawId()
	if str3[0] != "test1" || str3[1] != "str3" {
		t.Fatalf("struct3 name not correct got=%s", str3)
	}

	str4 := spec[4].(*ast.DefStatement).Name.RawId()
	if str4[0] != "test1" || str4[1] != "fl" {
		t.Fatalf("struct4 name not correct got=%s", str4)
	}
	str4f := spec[4].(*ast.DefStatement).Value.(*ast.FlowLiteral).Pairs
	for k, v := range str4f {
		keyId := k.RawId()
		if keyId[2] == "buzz" {
			valId := v.(*ast.StructInstance).RawId()
			if valId[0] != keyId[0] || valId[1] != keyId[1] || keyId[2] != valId[2] {
				t.Fatalf("field name is not correct value=%s", valId)
			}
			props := v.(*ast.StructInstance).Properties
			propId := props["foosh"].ProcessedName
			if propId[0] != "test1" || propId[1] != "fl" || propId[2] != "buzz" || propId[3] != "foosh" {
				t.Fatalf("field name is not correct value=%s", valId)
			}
		}
	}

}

func TestUnknowns(t *testing.T) {
	test := `spec test1;
	const a;
	const b;

	def s = stock{
	   x: unknown(),
	};

	def test = flow{
		u: new s,
		bar: func{
		   u.x <- a + b;
		},
	};

	for 5 run {
		t = new test;
		t.bar;
	};
	`

	process := prepTest(test)
	tree := process.Processed
	spec := tree.(*ast.Spec).Statements

	str1 := spec[1].(*ast.ConstantStatement).Name.RawId()
	if str1[0] != "test1" || str1[1] != "a" {
		t.Fatalf("constant1 name not correct got=%s", str1)
	}

	str1a := spec[1].(*ast.ConstantStatement).Value.(*ast.Unknown).RawId()
	if str1a[0] != "test1" || str1a[1] != "a" {
		t.Fatalf("unknown1 name not correct got=%s", str1a)
	}

	str2 := spec[2].(*ast.ConstantStatement).Name.RawId()
	if str2[0] != "test1" || str2[1] != "b" {
		t.Fatalf("constant2 name not correct got=%s", str2)
	}

	str2a := spec[2].(*ast.ConstantStatement).Value.(*ast.Unknown).RawId()
	if str2a[0] != "test1" || str2a[1] != "b" {
		t.Fatalf("unknown1 name not correct got=%s", str2a)
	}

	str3 := spec[3].(*ast.DefStatement).Value.(*ast.StockLiteral).Pairs
	for k, v := range str3 {
		keyId := k.RawId()
		if keyId[2] == "x" {
			valId := v.(*ast.Unknown).RawId()
			if valId[0] != keyId[0] || valId[1] != keyId[1] || keyId[2] != valId[2] {
				t.Fatalf("field name is not correct value=%s", valId)
			}

		}
	}
}

func prepTest(test string) *Processor {
	path := ""
	is := antlr.NewInputStream(test)
	lexer := parser.NewFaultLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	p := parser.NewFaultParser(stream)
	l := listener.NewListener(path, true, false)
	antlr.ParseTreeWalkerDefault.Walk(l, p.Spec())
	pro := NewProcesser()
	pro.Run(l.AST)
	return pro
}

func prepSysTest(test string) *Processor {
	path := ""
	is := antlr.NewInputStream(test)
	lexer := parser.NewFaultLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	p := parser.NewFaultParser(stream)
	l := listener.NewListener(path, true, false)
	antlr.ParseTreeWalkerDefault.Walk(l, p.SysSpec())
	pro := NewProcesser()
	pro.Run(l.AST)
	return pro
}
