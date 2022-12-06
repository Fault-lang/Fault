package preprocess

import (
	"fault/ast"
	"fault/listener"
	"fault/parser"
	"testing"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

func TestConstants(t *testing.T) {
	test := `spec test1;
	const x = 2;
	const y = 3.1;
	const z = unknown(a);`

	process := prepTest(test)

	consts := process.Specs["test1"]

	all := []string{"x", "y", "z"}
	for _, v := range consts.Order {
		for i, a := range all {
			if a == v[0] {
				all = append(all[0:i], all[i+1:]...)
			}
		}
	}

	if len(all) != 0 {
		t.Fatalf("constant %s missing from order", all)
	}

	x, _ := consts.FetchConstant("x")

	if _, ok := x.(*ast.IntegerLiteral); !ok {
		t.Fatalf("Constant x does not have the right type. got=%T", x)
	}

	if x.(*ast.IntegerLiteral).Value != 2 {
		t.Fatalf("Constant x does not have the right value. got=%d", x)
	}

	y, _ := consts.FetchConstant("y")
	if _, ok1 := y.(*ast.FloatLiteral); !ok1 {
		t.Fatalf("right y node does not have the right type. got=%T", y)
	}

	z, _ := consts.FetchConstant("z")
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

	foo, _ := variables.FetchStock("foo")
	if len(foo) != 4 {
		t.Fatalf("stock foo returns the wrong number of properties got=%d want=4", len(foo))
	}

	fizz := foo["fizz"].(*ast.Identifier).RawId()
	if len(fizz) != 2 || fizz[0] != "test1" || fizz[1] != "a" {
		t.Fatalf("identifier not converted to correct context got=%s", fizz)
	}

	zoo, _ := variables.FetchFlow("zoo")
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
	foo, _ := variables.FetchComponent("foo")
	if foo == nil {
		t.Fatal("component named foo not found")
	}

	if len(foo) != 3 {
		t.Fatalf("component foo returns the wrong number of states got=%d want=3", len(foo))
	}

	ifblock := foo["initial"].(*ast.FunctionLiteral).Body.Statements[0].(*ast.ExpressionStatement).Expression.(*ast.IfExpression)
	ifcond := ifblock.Condition.(*ast.InfixExpression)
	left := ifcond.Left.(*ast.InfixExpression)
	state := left.Left.(*ast.This).RawId()
	if len(state) != 3 || state[0] != "test" || state[1] != "foo" || state[2] != "initial" {
		t.Fatalf("state conditional wrap incorrect got=%s", state)
	}

	right := ifcond.Right.(*ast.InfixExpression)
	this := right.Left.(*ast.This).RawId()
	if len(this) != 3 || this[0] != "test" || this[1] != "foo" || this[2] != "x" {
		t.Fatalf("this special word not converted to correct context got=%s", this)
	}

	trueblock := ifblock.Consequence.Statements[0].(*ast.ExpressionStatement).Expression.(*ast.BuiltIn)
	if trueblock.IdString() != "test_foo_initial_stay" {
		t.Fatalf("built in stay not named correctly got=%s", trueblock.IdString())
	}

	elseblock := ifblock.Elif.Consequence.Statements[0].(*ast.ExpressionStatement).Expression.(*ast.BuiltIn)
	if elseblock.IdString() != "test_foo_initial_advance" {
		t.Fatalf("built in advance not named correctly got=%s", elseblock.IdString())
	}

	if elseblock.Parameters["toState"].(ast.Nameable).IdString() != "test_foo_alarm" {
		t.Fatalf("built in advance has the wrong input got=%s", elseblock.Parameters["toState"].(ast.Nameable).IdString())
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

	bar, _ := variables.FetchStock("str2_bar")
	if bar == nil {
		t.Fatal("stock names str2_bar not found")
	}

	if len(bar) != 1 {
		t.Fatalf("stock str2_bar returns the wrong number of properties got=%d want=4", len(bar))
	}

	zoo, _ := variables.FetchStock("fl_buzz_foosh_bar")
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

	fl, _ := variables.FetchFlow("f")
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
			if v.(*ast.StructInstance).ComplexScope != "" {
				t.Fatalf("struct complex scope incorrect got=%s", v.(*ast.StructInstance).ComplexScope)
			}

			valId := v.(*ast.StructInstance).RawId()
			if valId[0] != keyId[0] || valId[1] != keyId[1] || keyId[2] != valId[2] {
				t.Fatalf("field name is not correct value=%s", valId)
			}
			props := v.(*ast.StructInstance).Properties
			propId := props["foosh"].ProcessedName
			if propId[0] != "test1" || propId[1] != "fl" || propId[2] != "buzz" || propId[3] != "foosh" {
				t.Fatalf("field name is not correct value=%s", valId)
			}

			if props["foosh"].Value.(*ast.StructInstance).ComplexScope != "fl_buzz" {
				t.Fatalf("struct complex scope incorrect got=%s", props["foosh"].Value.(*ast.StructInstance).ComplexScope)
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

func TestAsserts(t *testing.T) {
	test := `spec test1;
	const a = 2;
	def foo = stock{
		bar: 1,
		x: 0,
		y: 5,
		z: 3,
	};

	assert a >= 10;
	assume foo.bar != 3;
	assert foo.x > foo.y[1] && foo.y[2] < foo.z; 
	`

	process := prepTest(test)
	tree := process.Processed
	spec := tree.(*ast.Spec).Statements

	str1 := spec[3].(*ast.AssertionStatement).Constraints.Left.(ast.Nameable).RawId()
	if len(str1) != 2 || str1[0] != "test1" || str1[1] != "a" {
		t.Fatalf("assumption var name 1 not correct got=%s", str1)
	}

	str1a := spec[4].(*ast.AssumptionStatement).Constraints.Left.(ast.Nameable).RawId()
	if len(str1a) != 3 || str1a[0] != "test1" || str1a[1] != "foo" || str1a[2] != "bar" {
		t.Fatalf("assumption var name 2 not correct got=%s", str1a)
	}

	str3p1 := spec[5].(*ast.AssertionStatement).Constraints.Left.(*ast.InfixExpression)
	str3p2 := spec[5].(*ast.AssertionStatement).Constraints.Right.(*ast.InfixExpression)

	str3 := str3p1.Left.(ast.Nameable).RawId()
	if len(str3) != 3 || str3[0] != "test1" || str3[1] != "foo" || str3[2] != "x" {
		t.Fatalf("assumption var name 3 not correct got=%s", str3)
	}

	str3a := str3p1.Right.(ast.Nameable).RawId()
	if len(str3a) != 3 || str3a[0] != "test1" || str3a[1] != "foo" || str3a[2] != "y" {
		t.Fatalf("assumption var name 4 not correct got=%s", str3a)
	}

	str4 := str3p2.Left.(ast.Nameable).RawId()
	if len(str4) != 3 || str4[0] != "test1" || str4[1] != "foo" || str4[2] != "y" {
		t.Fatalf("assumption var name 5 not correct got=%s", str4)
	}

	str4a := str3p2.Right.(ast.Nameable).RawId()
	if len(str4a) != 3 || str4a[0] != "test1" || str4[1] != "foo" || str4a[2] != "z" {
		t.Fatalf("assumption var name 6 not correct got=%s", str4a)
	}
}

func TestCollapseIf(t *testing.T) {
	test := `spec test1;
	const a = 2;
	for 1 run {
			if a == 2{
				if a != 0{
					3;
				}
			}
	};
	`
	process := prepTest(test)
	tree := process.Processed
	spec := tree.(*ast.Spec).Statements

	if1, ok := spec[2].(*ast.ForStatement).Body.Statements[0].(*ast.ExpressionStatement).Expression.(*ast.IfExpression)
	if !ok {
		t.Fatalf("statement not an IfExpression got=%T", spec[2].(*ast.ForStatement).Body.Statements[0].(*ast.ExpressionStatement).Expression)
	}

	cond, _ := if1.Condition.(*ast.InfixExpression)
	cond1, ok := cond.Right.(*ast.InfixExpression)
	cond2, ok2 := cond.Left.(*ast.InfixExpression)

	if cond.Operator != "&&" || !ok || !ok2 {
		t.Fatalf("multicond not collapsed got=%s", cond)
	}
	if cond1.Right.(*ast.IntegerLiteral).Value != 0 {
		t.Fatalf("collapsed multicond wrong right value got=%s", cond1)
	}

	if cond2.Right.(*ast.IntegerLiteral).Value != 2 {
		t.Fatalf("collapsed multicond wrong left value got=%s", cond2)
	}
}

func TestCollapseIfElse(t *testing.T) {
	test := `spec test1;
	const a = 2;
	for 1 run {
			if true {
				3;
			}else if a != 0{
				if a == 2{
					3;
				}
			}
	};
	`
	process := prepTest(test)
	tree := process.Processed
	spec := tree.(*ast.Spec).Statements

	if1, ok := spec[2].(*ast.ForStatement).Body.Statements[0].(*ast.ExpressionStatement).Expression.(*ast.IfExpression)
	if !ok {
		t.Fatalf("statement not an IfExpression got=%T", spec[2].(*ast.ForStatement).Body.Statements[0].(*ast.ExpressionStatement).Expression)
	}

	cond, _ := if1.Elif.Condition.(*ast.InfixExpression)
	cond1, ok := cond.Right.(*ast.InfixExpression)
	cond2, ok2 := cond.Left.(*ast.InfixExpression)

	if cond.Operator != "&&" || !ok || !ok2 {
		t.Fatalf("multicond not collapsed got=%s", cond)
	}
	if cond1.Right.(*ast.IntegerLiteral).Value != 2 {
		t.Fatalf("collapsed multicond wrong right value got=%s", cond1)
	}

	if cond2.Right.(*ast.IntegerLiteral).Value != 0 {
		t.Fatalf("collapsed multicond wrong left value got=%s", cond2)
	}
}

func TestCollapseElse(t *testing.T) {
	test := `spec test1;
	const a = 2;
	for 1 run {
			if true {
				3;
			}else{
				if a == 2{
					3;
				}
			}
	};
	`
	process := prepTest(test)
	tree := process.Processed
	spec := tree.(*ast.Spec).Statements

	if1, ok := spec[2].(*ast.ForStatement).Body.Statements[0].(*ast.ExpressionStatement).Expression.(*ast.IfExpression)
	if !ok {
		t.Fatalf("statement not an IfExpression got=%T", spec[2].(*ast.ForStatement).Body.Statements[0].(*ast.ExpressionStatement).Expression)
	}

	cond, _ := if1.Elif.Condition.(*ast.InfixExpression)

	if cond.Operator != "==" {
		t.Fatalf("multicond not collapsed got=%s", cond)
	}
	if cond.Right.(*ast.IntegerLiteral).Value != 2 {
		t.Fatalf("collapsed multicond wrong right value got=%s", cond)
	}
}

func TestCondCollapse(t *testing.T) {
	test := `spec test1;
	const a = 2;
	for 1 run {
			if a == 2{
				if a != 0{
					3;
				}else if a < 1 {
					if a >= 2 {
					true;
					}
				}
			}else if a !=5 {
				true;
			}else{
				if a > 4 {
					false;
				}
			}
	};
`
	process := prepTest(test)
	tree := process.Processed
	spec := tree.(*ast.Spec).Statements

	if1, ok := spec[2].(*ast.ForStatement).Body.Statements[0].(*ast.ExpressionStatement).Expression.(*ast.IfExpression)
	if !ok {
		t.Fatalf("statement not an IfExpression got=%T", spec[2].(*ast.ForStatement).Body.Statements[0].(*ast.ExpressionStatement).Expression)
	}

	cond, _ := if1.Condition.(*ast.InfixExpression)
	subcond1, ok := cond.Right.(*ast.InfixExpression)
	subcond2, ok2 := cond.Left.(*ast.InfixExpression)

	if cond.Operator != "&&" || !ok || !ok2 {
		t.Fatalf("multicond not collapsed got=%s", cond)
	}
	if subcond1.Operator != "!=" && subcond1.Right.(*ast.IntegerLiteral).Value != 0 {
		t.Fatalf("collapsed multicond wrong right value got=%s", subcond1)
	}

	if subcond2.Operator != "==" && subcond2.Right.(*ast.IntegerLiteral).Value != 2 {
		t.Fatalf("collapsed multicond wrong left value got=%s", subcond2)
	}

	cond2 := if1.Elif.Condition.(*ast.InfixExpression)

	if cond2.Operator != "!=" {
		t.Fatalf("multicond else if not collapsed got=%s", cond2)
	}
	if cond2.Right.(*ast.IntegerLiteral).Value != 5 {
		t.Fatalf("collapsed multicond wrong right value got=%s", cond2.Right)
	}

	cond3 := if1.Elif.Elif.Condition.(*ast.InfixExpression)
	subcond3, ok3 := cond3.Left.(*ast.InfixExpression)
	tempcond4, _ := cond3.Right.(*ast.InfixExpression)
	subcond4, ok4 := tempcond4.Left.(*ast.InfixExpression)
	subcond5, ok5 := tempcond4.Right.(*ast.InfixExpression)

	if cond3.Operator != "&&" || !ok3 || !ok4 || !ok5 {
		t.Fatalf("multicond else if not collapsed got=%s", cond3)
	}
	if subcond3.Operator != "==" && subcond3.Right.(*ast.IntegerLiteral).Value != 2 {
		t.Fatalf("collapsed multicond wrong right value got=%s", subcond3)
	}

	if subcond4.Operator != "<" && subcond4.Right.(*ast.IntegerLiteral).Value != 1 {
		t.Fatalf("collapsed multicond wrong left value got=%s", subcond4)
	}

	if subcond5.Operator != ">=" && subcond5.Right.(*ast.IntegerLiteral).Value != 2 {
		t.Fatalf("collapsed multicond wrong right value got=%s", subcond5)
	}

	cond4 := if1.Elif.Elif.Elif.Condition.(*ast.InfixExpression)

	if cond4.Operator != ">" {
		t.Fatalf("multicond else if not collapsed got=%s", cond4)
	}
	if cond4.Right.(*ast.IntegerLiteral).Value != 4 {
		t.Fatalf("collapsed multicond wrong right value got=%s", cond4.Right)
	}

	// if a == 2 && a != 0{
	// 		3;
	// }elif a !=5 {
	// 	true;
	// 	}elif a == 2 && a < 1 && a >= 2{
	// 		true;
	// }elif a > 4 {
	// 		false;
	// };

}

func TestInstanceFlatten(t *testing.T) {
	p := NewProcesser()
	p.trail = p.trail.PushSpec("test")
	p.Specs["test"] = NewSpecRecord()
	p.Specs["test"].SpecName = "test"
	p.initialPass = false

	stockdata := make(map[string]ast.Node)
	stockdata["zoo"] = &ast.IntegerLiteral{Value: 2}

	p.Specs["test"].AddStock("foo", stockdata)
	p.Specs["test"].Index("STOCK", "foo")
	p.structTypes["test"] = map[string]string{"foo": "STOCK"}

	test := &ast.Instance{Value: &ast.Identifier{Spec: "test", Value: "foo"}, Name: "bar", Order: []string{"zoo"}}

	node, err := p.walk(test)
	if err != nil {
		t.Fatalf("test errored: %s", err.Error())
	}

	n, ok := node.(*ast.StructInstance)
	if !ok {
		t.Fatalf("instance not converted to StructInstance got=%T", node)
	}

	if n.Parent[0] != "test" || n.Parent[1] != "foo" {
		t.Fatalf("StructInstance has the incorrect parent information got=%s", n.Parent)
	}

	if n.Properties["zoo"] == nil {
		t.Fatalf("StructInstance missing it's property")
	}

	i, ok := n.Properties["zoo"].Value.(*ast.IntegerLiteral)
	if !ok {
		t.Fatalf("StructInstanct property is the wrong type got=%T", n.Properties["zoo"].Value)
	}

	if i.Value != 2 {
		t.Fatalf("StructInstanct property is the wrong value got=%d", i.Value)
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
	pre := NewProcesser()
	pre.StructsPropertyOrder = l.StructsPropertyOrder
	pre.Run(l.AST)
	return pre
}

func prepSysTest(test string) *Processor {
	path := ""
	is := antlr.NewInputStream(test)
	lexer := parser.NewFaultLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	p := parser.NewFaultParser(stream)
	l := listener.NewListener(path, true, false)
	antlr.ParseTreeWalkerDefault.Walk(l, p.SysSpec())
	pre := NewProcesser()
	pre.StructsPropertyOrder = l.StructsPropertyOrder
	pre.Run(l.AST)
	return pre
}
