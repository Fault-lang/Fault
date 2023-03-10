package llvm

import (
	"fault/ast"
	"fault/listener"
	"fault/parser"
	"fault/preprocess"
	"fault/types"
	"fmt"
	"io"
	"os/exec"
	"strings"
	"testing"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	irtypes "github.com/llir/llvm/ir/types"
)

func TestSimpleConst(t *testing.T) {
	test := `spec test1;
			const x = 2;
			const y = 1.2;
			const a = true;
			const b = false;
			const c = "Hello World!";
	`
	expecting := `
	@__rounds = global i16 0
	@__parallelGroup = global [5 x i8] c"start"
	@test1_x = global double 2.0
	@test1_y = global double 0x3FF3333333333333
	@test1_a = global i1 true
	@test1_b = global i1 false
	@test1_c = global [14 x i8] c"\22Hello World!\22"
	
	define void @__run() {
	block-0:
		ret void
	}
	
`

	llvm, err := prepTest(test)

	if err != nil {
		t.Fatalf("compilation failed on valid spec. got=%s", err)
	}

	ir, err := validateIR(llvm)

	if err != nil {
		t.Fatalf("generated IR is not valid. got=%s", err)
	}

	err = compareResults(llvm, expecting, string(ir))

	if err != nil {
		t.Fatalf(err.Error())
	}
}

func TestRunBlock(t *testing.T) {
	test := `spec test1;
			const a = 2.3;
			const b = 2;

			def foo = flow{
				buzz: new bar,
				fizz: func{
					buzz.b <- a + buzz.a;
				},
				fizz2: func{
					buzz.b <- buzz.a - b;
				},
				fizz3: func{
					buzz.a <- buzz.b + b;
				},
			};

			def bar = stock{
				a: 10,
				b: 20,
			};

			for 5 run{
				test = new foo;
				test.fizz | test.fizz2;
				test.fizz3;
			};
	`

	expecting := `@__rounds = global i16 0
	@__parallelGroup = global [5 x i8] c"start"
	@test1_a = global double 0x4002666666666666
	@test1_b = global double 2.0
	
	define void @__run() {
	block-1:
		store i16 0, i16* @__rounds
		%test1_test_buzz_a = alloca double
		store double 10.0, double* %test1_test_buzz_a
		%test1_test_buzz_b = alloca double
		store double 20.0, double* %test1_test_buzz_b
		call void @test1_test_fizz(double* %test1_test_buzz_a, double* %test1_test_buzz_b), !\344b9b452817d4d3ea103f1449105264c !DIBasicType(tag: DW_TAG_string_type)
		call void @test1_test_fizz2(double* %test1_test_buzz_a, double* %test1_test_buzz_b), !\344b9b452817d4d3ea103f1449105264c !DIBasicType(tag: DW_TAG_string_type)
		call void @test1_test_fizz3(double* %test1_test_buzz_a, double* %test1_test_buzz_b), !\3227b3938f885317dea9c644434cb82dd !DIBasicType(tag: DW_TAG_string_type)
		store i16 1, i16* @__rounds
		call void @test1_test_fizz(double* %test1_test_buzz_a, double* %test1_test_buzz_b), !\39a2b113b63e8232c2945f1018bf785f0 !DIBasicType(tag: DW_TAG_string_type)
		call void @test1_test_fizz2(double* %test1_test_buzz_a, double* %test1_test_buzz_b), !\39a2b113b63e8232c2945f1018bf785f0 !DIBasicType(tag: DW_TAG_string_type)
		call void @test1_test_fizz3(double* %test1_test_buzz_a, double* %test1_test_buzz_b), !\387c9dfef940c096cf145af18149d3600 !DIBasicType(tag: DW_TAG_string_type)
		store i16 2, i16* @__rounds
		call void @test1_test_fizz(double* %test1_test_buzz_a, double* %test1_test_buzz_b), !cb7fdc02d16d31723661579b54e31084 !DIBasicType(tag: DW_TAG_string_type)
		call void @test1_test_fizz2(double* %test1_test_buzz_a, double* %test1_test_buzz_b), !cb7fdc02d16d31723661579b54e31084 !DIBasicType(tag: DW_TAG_string_type)
		call void @test1_test_fizz3(double* %test1_test_buzz_a, double* %test1_test_buzz_b), !\370ad8bfb0509411d97738ac929bc3d01 !DIBasicType(tag: DW_TAG_string_type)
		store i16 3, i16* @__rounds
		call void @test1_test_fizz(double* %test1_test_buzz_a, double* %test1_test_buzz_b), !\37e3f85f9630519ec31a508b611b1d4bb !DIBasicType(tag: DW_TAG_string_type)
		call void @test1_test_fizz2(double* %test1_test_buzz_a, double* %test1_test_buzz_b), !\37e3f85f9630519ec31a508b611b1d4bb !DIBasicType(tag: DW_TAG_string_type)
		call void @test1_test_fizz3(double* %test1_test_buzz_a, double* %test1_test_buzz_b), !e57d2299cbd9024a113885402ef4e089 !DIBasicType(tag: DW_TAG_string_type)
		store i16 4, i16* @__rounds
		call void @test1_test_fizz(double* %test1_test_buzz_a, double* %test1_test_buzz_b), !\37fbb0459ad7da0f1a336cf5de1cf9068 !DIBasicType(tag: DW_TAG_string_type)
		call void @test1_test_fizz2(double* %test1_test_buzz_a, double* %test1_test_buzz_b), !\37fbb0459ad7da0f1a336cf5de1cf9068 !DIBasicType(tag: DW_TAG_string_type)
		call void @test1_test_fizz3(double* %test1_test_buzz_a, double* %test1_test_buzz_b), !\340b3487db7f69f408810fb4cb8b544eb !DIBasicType(tag: DW_TAG_string_type)
		ret void
	}
	
	define void @test1_test_fizz(double* %test1_test_buzz_a, double* %test1_test_buzz_b) {
	block-2:
		%0 = load double, double* %test1_test_buzz_b
		%1 = load double, double* @test1_a
		%2 = load double, double* %test1_test_buzz_a
		%3 = fadd double %1, %2
		%4 = fadd double %0, %3
		store double %4, double* %test1_test_buzz_b
		ret void
	}
	
	define void @test1_test_fizz2(double* %test1_test_buzz_a, double* %test1_test_buzz_b) {
	block-3:
		%0 = load double, double* %test1_test_buzz_b
		%1 = load double, double* %test1_test_buzz_a
		%2 = load double, double* @test1_b
		%3 = fsub double %1, %2
		%4 = fadd double %0, %3
		store double %4, double* %test1_test_buzz_b
		ret void
	}
	
	define void @test1_test_fizz3(double* %test1_test_buzz_a, double* %test1_test_buzz_b) {
	block-4:
		%0 = load double, double* %test1_test_buzz_a
		%1 = load double, double* %test1_test_buzz_b
		%2 = load double, double* @test1_b
		%3 = fadd double %1, %2
		%4 = fadd double %0, %3
		store double %4, double* %test1_test_buzz_a
		ret void
	}		
`
	//Should fadd have variable names or the values in those variables?

	llvm, err := prepTest(test)

	if err != nil {
		t.Fatalf("compilation failed on valid spec. got=%s", err)
	}
	ir, err := validateIR(llvm)

	if err != nil {
		t.Fatalf("generated IR is not valid. got=%s", err)
	}

	err = compareResults(llvm, expecting, string(ir))

	if err != nil {
		t.Fatalf(err.Error())
	}
}

func TestIfCond(t *testing.T) {
	test := `spec test1;
			const a = 2.3;
			const b = 2;

			def foo = flow{
				buzz: new bar,
				fizz: func{
					if buzz.a > 2{
						buzz.a -> b;
					}else{
						buzz.a = 10;
					}
					buzz.b -> 1;
				},
			};

			def bar = stock{
				a: 10,
				b: 20,
			};

			for 1 run{
				test = new foo;
				test.fizz;
			};
	`

	expecting := `@__rounds = global i16 0
	@__parallelGroup = global [5 x i8] c"start"
	@test1_a = global double 0x4002666666666666
	@test1_b = global double 2.0
	
	define void @__run() {
	block-5:
		store i16 0, i16* @__rounds
		%test1_test_buzz_a = alloca double
		store double 10.0, double* %test1_test_buzz_a
		%test1_test_buzz_b = alloca double
		store double 20.0, double* %test1_test_buzz_b
		call void @test1_test_fizz(double* %test1_test_buzz_a, double* %test1_test_buzz_b), !d44e0a3fc2944aa552d9118f291d3106 !DIBasicType(tag: DW_TAG_string_type)
		ret void
	}
	
	define void @test1_test_fizz(double* %test1_test_buzz_a, double* %test1_test_buzz_b) {
	block-6:
		%0 = load double, double* %test1_test_buzz_a
		%1 = fcmp ogt double %0, 2.0
		br i1 %1, label %block-8-true, label %block-9-false
	
	block-7-after:
		%2 = load double, double* %test1_test_buzz_b
		%3 = fsub double %2, 1.0
		store double %3, double* %test1_test_buzz_b
		ret void
	
	block-8-true:
		%4 = load double, double* %test1_test_buzz_a
		%5 = load double, double* @test1_b
		%6 = fsub double %4, %5
		store double %6, double* %test1_test_buzz_a
		br label %block-7-after
	
	block-9-false:
		store double 10.0, double* %test1_test_buzz_a
		br label %block-7-after
	}		
	`

	llvm, err := prepTest(test)

	if err != nil {
		t.Fatalf("compilation failed on valid spec. got=%s", err)
	}

	ir, err := validateIR(llvm)

	if err != nil {
		t.Fatalf("generated IR is not valid. got=%s", err)
	}

	err = compareResults(llvm, expecting, string(ir))

	if err != nil {
		t.Fatalf(err.Error())
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

	expecting := `@__rounds = global i16 0
	@__parallelGroup = global [5 x i8] c"start"
	@test1_a = global double 0x3DA3CA8CB153A753
	@test1_b = global double 0x3DA3CA8CB153A753
	
	define void @__run() {
	block-10:
		store i16 0, i16* @__rounds
		%test1_t_u_x = alloca double
		store double 0x3DA3CA8CB153A753, double* %test1_t_u_x
		call void @test1_t_bar(double* %test1_t_u_x), !\34614d4e08724f278c8ce39e50955edbc !DIBasicType(tag: DW_TAG_string_type)
		store i16 1, i16* @__rounds
		call void @test1_t_bar(double* %test1_t_u_x), !ba735eefbef72f20ea6a264b981e9285 !DIBasicType(tag: DW_TAG_string_type)
		store i16 2, i16* @__rounds
		call void @test1_t_bar(double* %test1_t_u_x), !fa06e912698cf4825866672ced835870 !DIBasicType(tag: DW_TAG_string_type)
		store i16 3, i16* @__rounds
		call void @test1_t_bar(double* %test1_t_u_x), !f3a3858248090df83b9702c4852e0e28 !DIBasicType(tag: DW_TAG_string_type)
		store i16 4, i16* @__rounds
		call void @test1_t_bar(double* %test1_t_u_x), !ba20b5c159a59aeb04358b812e68f2d2 !DIBasicType(tag: DW_TAG_string_type)
		ret void
	}
	
	define void @test1_t_bar(double* %test1_t_u_x) {
	block-11:
		%0 = load double, double* %test1_t_u_x
		%1 = load double, double* @test1_a
		%2 = load double, double* @test1_b
		%3 = fadd double %1, %2
		%4 = fadd double %0, %3
		store double %4, double* %test1_t_u_x
		ret void
	}`

	llvm, err := prepTest(test)

	if err != nil {
		t.Fatalf("compilation failed on valid spec. got=%s", err)
	}

	ir, err := validateIR(llvm)

	if err != nil {
		t.Fatalf("generated IR is not valid. got=%s", err)
	}

	err = compareResults(llvm, expecting, string(ir))

	if err != nil {
		t.Fatalf(err.Error())
	}
}

func TestParamReset(t *testing.T) {
	structs := make(map[string]*preprocess.SpecRecord)
	c := NewCompiler()
	c.LoadMeta(structs, make(map[string][]float64), []string{}, true)
	s := NewCompiledSpec("test")
	c.currentSpec = "test"
	c.specs["test"] = s

	id := []string{"test", "this", "func"}
	val1 := constant.NewInt(irtypes.I32, 0)
	s.DefineSpecVar(id, val1)
	s.AddParam(id, val1)

	val2 := constant.NewInt(irtypes.I32, 5)
	s.DefineSpecVar(id, val2)

	if s.vars.GetState(id) != 1 {
		t.Fatalf("var state is incorrect for %s. got=%d", id, s.GetSpecVarState(id))
	}

	p := ir.NewParam(strings.Join(id, "_"), DoubleP)
	c.resetParaState([]*ir.Param{p})

	if s.vars.GetState(id) != 0 {
		t.Fatalf("var state is incorrect for %s. got=%d", id, s.GetSpecVarState(id))
	}
}

func TestNegate(t *testing.T) {
	test := &ast.InfixExpression{
		Left: &ast.Boolean{
			Value: true,
		},
		Right: &ast.Boolean{
			Value: false,
		},
		Operator: "==",
	}

	n := negate(test)

	if n.(*ast.InfixExpression).Operator != "!=" {
		t.Fatalf("operator has not been negated got=%s", n.(*ast.InfixExpression).Operator)
	}

	if n.(*ast.InfixExpression).Left.(*ast.Boolean).Value != false {
		t.Fatalf("left value of infix not negated got=%s.", n.(*ast.InfixExpression).Left)
	}

	if n.(*ast.InfixExpression).Right.(*ast.Boolean).Value != true {
		t.Fatalf("right value of infix not negated. got=%s", n.(*ast.InfixExpression).Right)
	}

	test2 := &ast.Boolean{Value: true}

	n2 := negate(test2)

	if n2.(*ast.Boolean).Value != false {
		t.Fatalf("boolean has not been negated got=%s", n2.(*ast.Boolean).String())
	}

	test3 := &ast.Boolean{Value: false}

	n3 := negate(test3)

	if n3.(*ast.Boolean).Value != true {
		t.Fatalf("boolean has not been negated got=%s", n3.(*ast.Boolean).String())
	}

	test4 := &ast.PrefixExpression{
		Operator: "!",
		Right:    &ast.Boolean{Value: false},
	}

	n4 := negate(test4)

	if n4.(*ast.Boolean).Value != false {
		t.Fatalf("boolean has not been negated got=%v", n4.(*ast.Boolean).Value)
	}
}

func TestEval(t *testing.T) {
	tests := []*ast.InfixExpression{{
		Left:  &ast.IntegerLiteral{Value: 2},
		Right: &ast.IntegerLiteral{Value: 2},
	},
		{
			Left:  &ast.FloatLiteral{Value: 2.5},
			Right: &ast.IntegerLiteral{Value: 2},
		},
		{
			Left:     &ast.IntegerLiteral{Value: 2},
			Operator: "+",
			Right:    &ast.FloatLiteral{Value: 2.5},
		}}

	operators := []string{"+", "-", "/", "*"}

	results := []ast.Node{
		&ast.IntegerLiteral{Value: 4},
		&ast.FloatLiteral{Value: 4.5},
		&ast.FloatLiteral{Value: 4.5},
		&ast.IntegerLiteral{Value: 0},
		&ast.FloatLiteral{Value: .5},
		&ast.FloatLiteral{Value: -.5},
		&ast.FloatLiteral{Value: 1},
		&ast.FloatLiteral{Value: 1.25},
		&ast.FloatLiteral{Value: .8},
		&ast.IntegerLiteral{Value: 4},
		&ast.FloatLiteral{Value: 5},
		&ast.FloatLiteral{Value: 5},
	}

	i := 0
	for _, o := range operators {
		for _, n := range tests {
			n.Operator = o
			test := evaluate(n)
			switch actual := test.(type) {
			case *ast.IntegerLiteral:
				expected, ok := results[i].(*ast.IntegerLiteral)
				if !ok {
					t.Fatalf("expected value a different type from actual expected=%s actual=%s", results[i], test)
				}
				if expected.Value != actual.Value {
					t.Fatalf("expected value a different from actual expected=%s actual=%s", expected, actual)
				}
			case *ast.FloatLiteral:
				expected, ok := results[i].(*ast.FloatLiteral)
				if !ok {
					t.Fatalf("expected value a different type from actual expected=%s actual=%s", results[i], test)
				}
				if expected.Value != actual.Value {
					t.Fatalf("expected value a different from actual expected=%s actual=%s", expected, actual)
				}
			}
			i++
		}
	}
}

func TestIsVarSet(t *testing.T) {
	c := NewCompiler()
	c.specStructs["test"] = preprocess.NewSpecRecord()

	test := []string{"test", "this"}
	test1 := []string{"test", "this", "func"}

	val := map[string]ast.Node{"func": &ast.Nil{}}

	if c.isVarSet(test) {
		t.Fatal("isVarSet returned true, should return false")
	}

	c.specStructs["test"].AddComponent("this", val)
	if !c.isVarSet(test) {
		t.Fatal("isVarSet returned false on component, should return true")
	}
	c.specStructs["test"].Index("COMPONENT", "this")

	if !c.isVarSet(test1) {
		t.Fatal("isStrVarSet returned false on a component var, should return true")
	}

	c.specStructs["test"] = preprocess.NewSpecRecord()

	c.specStructs["test"].AddConstant("this", val["func"])
	if !c.isVarSet(test) {
		t.Fatal("isVarSet returned false on constant, should return true")
	}

	c.specStructs["test"] = preprocess.NewSpecRecord()

	c.specStructs["test"].AddFlow("this", val)
	if !c.isVarSet(test) {
		t.Fatal("isVarSet returned false on flow, should return true")
	}
	c.specStructs["test"].Index("FLOW", "this")

	if !c.isVarSet(test1) {
		t.Fatal("isStrVarSet returned false on a flow var, should return true")
	}

	c.specStructs["test"] = preprocess.NewSpecRecord()

	c.specStructs["test"].AddStock("this", val)
	if !c.isVarSet(test) {
		t.Fatal("isVarSet returned false on stock, should return true")
	}

	c.specStructs["test"].Index("STOCK", "this")

	if !c.isVarSet(test1) {
		t.Fatal("isStrVarSet returned false on a stock var, should return true")
	}

}

func TestUncertains(t *testing.T) {
	c := NewCompiler()
	c.specs["test"] = NewCompiledSpec("test")
	test := &ast.StructInstance{Spec: "test", Name: "foo", Parent: []string{"test", "zoo"}, Order: []string{"bar"}, ProcessedName: []string{"test", "foo"}, Properties: map[string]*ast.StructProperty{"bar": {Spec: "test", Name: "bar", ProcessedName: []string{"test", "foo", "bar"}, Value: &ast.Uncertain{Mean: 2.0, Sigma: .3, ProcessedName: []string{"test", "foo", "bar"}}}}}
	c.processStruct(test)

	if len(c.Uncertains["test_foo_bar"]) == 0 {
		t.Fatal("uncertain value not stored")
	}

	if c.Uncertains["test_foo_bar"][0] != 2.0 || c.Uncertains["test_foo_bar"][1] != .3 {
		t.Fatalf("uncertain stored value is incorrect, got=%f", c.Uncertains["test_foo_bar"])
	}

}

func TestUnknowns2(t *testing.T) {
	c := NewCompiler()
	c.specs["test"] = NewCompiledSpec("test")
	test := &ast.StructInstance{Spec: "test", Name: "foo", Parent: []string{"test", "zoo"}, Order: []string{"bar"}, ProcessedName: []string{"test", "foo"}, Properties: map[string]*ast.StructProperty{"bar": {Spec: "test", Name: "bar", ProcessedName: []string{"test", "foo", "bar"}, Value: &ast.Unknown{Name: &ast.Identifier{Spec: "test", Value: "bar"}, ProcessedName: []string{"test", "foo", "bar"}}}}}
	c.processStruct(test)

	if len(c.Unknowns) == 0 {
		t.Fatal("unknown value not stored")
	}

	if c.Unknowns[0] != "test_foo_bar" {
		t.Fatalf("unknowns stored value is incorrect, got=%s", c.Unknowns[0])
	}

}

func TestComponentIR(t *testing.T) {
	test := `
	system test;

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
	};
	`

	expecting := `@__rounds = global i16 0
	@__parallelGroup = global [5 x i8] c"start"
	
	define void @__run() {
	block-16:
		%test_foo_x = alloca double
		store double 8.0, double* %test_foo_x
		%test_foo_initial = alloca i1
		store i1 false, i1* %test_foo_initial
		%test_foo_alarm = alloca i1
		store i1 false, i1* %test_foo_alarm
		store i1 true, i1* %test_foo_initial
		call void @test_foo_initial__state(i1* %test_foo_alarm, i1* %test_foo_initial, double* %test_foo_x)
		call void @test_foo_alarm__state(i1* %test_foo_alarm, i1* %test_foo_initial, double* %test_foo_x)
		ret void
	}
	
	define void @test_foo_initial__state(i1* %test_foo_alarm, i1* %test_foo_initial, double* %test_foo_x) {
	block-17:
		%0 = load i1, i1* %test_foo_initial
		%1 = icmp eq i1 %0, true
		br i1 %1, label %block-19-true, label %block-18-after
	
	block-18-after:
		%2 = load i1, i1* %test_foo_initial
		%3 = icmp eq i1 %2, true
		%4 = load double, double* %test_foo_x
		%5 = fcmp ogt double %4, 10.0
		%6 = and i1 %3, %5
		br i1 %6, label %block-22-true, label %block-21-after
	
	block-19-true:
		%7 = alloca [14 x i8]
		store [14 x i8] c"test_foo_alarm", [14 x i8]* %7
		%8 = bitcast [14 x i8]* %7 to i8*
		%9 = call i1 @advance(i8* %8)
		br label %block-18-after
	
	block-21-after:
		ret void
	
	block-22-true:
		%10 = call i1 @stay()
		br label %block-21-after
	}
	
	define i1 @advance(i8* %toState) {
	block-20:
		ret i1 true
	}
	
	define i1 @stay() {
	block-23:
		ret i1 true
	}
	
	define void @test_foo_alarm__state(i1* %test_foo_alarm, i1* %test_foo_initial, double* %test_foo_x) {
	block-24:
		%0 = load i1, i1* %test_foo_alarm
		%1 = icmp eq i1 %0, true
		br i1 %1, label %block-26-true, label %block-25-after
	
	block-25-after:
		ret void
	
	block-26-true:
		%2 = alloca [14 x i8]
		store [14 x i8] c"test_foo_close", [14 x i8]* %2
		%3 = bitcast [14 x i8]* %2 to i8*
		%4 = call i1 @advance(i8* %3)
		br label %block-25-after
	}`

	llvm, err := prepTestSys(test)

	if err != nil {
		t.Fatalf("compilation failed on valid spec. got=%s", err)
	}

	ir, err := validateIR(llvm)

	if err != nil {
		t.Fatalf("generated IR is not valid. got=%s", err)
	}

	err = compareResults(llvm, expecting, string(ir))

	if err != nil {
		t.Fatalf(err.Error())
	}

}

// init values
// clock
// index

func compareResults(llvm string, expecting string, ir string) error {
	if !strings.Contains(ir, "source_filename = \"<stdin>\"") {
		return fmt.Errorf("optimized ir not valid. \ngot=%s", ir)
	}

	llvm = stripAndEscape(llvm)
	expecting = stripAndEscape(expecting)
	if len(llvm) != len(expecting) {
		return fmt.Errorf("wrong instructions length.\nwant=%s\ngot=%s",
			expecting, llvm)
	}

	if llvm != expecting {
		return fmt.Errorf("LLVM IR String does not match.\nwant=%q\ngot=%q",
			expecting, llvm)
	}
	return nil
}

func stripAndEscape(str string) string {
	var output strings.Builder
	output.Grow(len(str))
	for _, ch := range str {
		if !unicode.IsSpace(ch) {
			if ch == '%' {
				output.WriteString("%%")
			} else {
				output.WriteRune(ch)
			}
		}
	}
	return output.String()
}

func prepTest(test string) (string, error) {
	path := ""
	is := antlr.NewInputStream(test)
	lexer := parser.NewFaultLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	p := parser.NewFaultParser(stream)
	l := listener.NewListener(path, true, false)
	antlr.ParseTreeWalkerDefault.Walk(l, p.Spec())
	pre := preprocess.NewProcesser()
	pre.StructsPropertyOrder = l.StructsPropertyOrder
	tree := pre.Run(l.AST)

	ty := types.NewTypeChecker(pre.Specs)
	tree, err := ty.Check(tree)

	if err != nil {
		return "", err
	}
	compiler := NewCompiler()
	compiler.LoadMeta(pre.Specs, l.Uncertains, l.Unknowns, true)
	err = compiler.Compile(tree)
	if err != nil {
		return "", err
	}
	//fmt.Println(compiler.GetIR())
	return compiler.GetIR(), err
}

func prepTestSys(test string) (string, error) {
	path := ""
	is := antlr.NewInputStream(test)
	lexer := parser.NewFaultLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	p := parser.NewFaultParser(stream)
	l := listener.NewListener(path, true, false)
	antlr.ParseTreeWalkerDefault.Walk(l, p.SysSpec())
	pre := preprocess.NewProcesser()
	pre.StructsPropertyOrder = l.StructsPropertyOrder
	tree := pre.Run(l.AST)

	ty := types.NewTypeChecker(pre.Specs)
	tree, err := ty.Check(tree)
	if err != nil {
		return "", err
	}
	compiler := NewCompiler()
	compiler.LoadMeta(pre.Specs, l.Uncertains, l.Unknowns, true)
	err = compiler.Compile(tree)
	if err != nil {
		return "", err
	}
	//fmt.Println(compiler.GetIR())
	return compiler.GetIR(), err
}

func validateIR(ir string) ([]byte, error) {
	//Run LLVM optimizer to check IR is valid
	cmd := exec.Command("opt", "-S", "-inline", "--mem2reg")
	stdin, err := cmd.StdinPipe()
	if err != nil {
		return nil, err
	}

	go func() {
		defer stdin.Close()
		io.WriteString(stdin, ir)
	}()

	out, err := cmd.CombinedOutput()
	if err != nil {
		return nil, err
	}

	return out, err
}
