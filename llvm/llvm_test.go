package llvm

import (
	"fault/ast"
	"fault/listener"
	"fault/preprocess"
	"fault/swaps"
	"fault/types"
	"fmt"
	"io"
	"os/exec"
	"strings"
	"testing"
	"unicode"

	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	irtypes "github.com/llir/llvm/ir/types"
)

// These tests need to be run with go test, when run separately the block numbers are different
// and the tests fail.

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
	@__parallelGroup = global [38xi8] c"\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00"
	@__choiceGroup = global [38xi8] c"\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00"
	@test1_x = global double 2.0
	@test1_y = global double 0x3FF3333333333333
	@test1_a = global i1 true
	@test1_b = global i1 false
	@test1_c = global i1 false
	
	define void @__run() {
	block-0:
		ret void
	}
	
`

	llvm, err := prepTest(test, true)

	if err != nil {
		t.Fatalf("compilation failed on valid spec. got=%s", err)
	}

	ir, err := validateIR(llvm)

	if err != nil {
		t.Fatalf("generated IR is not valid. got=%s", err)
	}

	err = compareResults(llvm, expecting, string(ir))

	if err != nil {
		t.Fatal(err.Error())
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

			for 5 init{test = new foo;} run{
				test.fizz | test.fizz2;
				test.fizz3;
			};
	`

	expecting := `@__rounds=globali160@__parallelGroup=global[38xi8]c"\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00"@__choiceGroup=global[38xi8]c"\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00"@test1_a=globaldouble0x4002666666666666@test1_b=globaldouble2.0definevoid@__run(){block-1:storei160,i16*@__rounds%test1_test_buzz_a=allocadoublestoredouble10.0,double*%test1_test_buzz_a%test1_test_buzz_b=allocadoublestoredouble20.0,double*%test1_test_buzz_bstorei161,i16*@__roundsstore[38xi8]c"44b9b452817d4d3ea103f1449105264c_start",[38xi8]*@__parallelGroupcallvoid@test1_test_fizz(double*%test1_test_buzz_a,double*%test1_test_buzz_b)callvoid@test1_test_fizz2(double*%test1_test_buzz_a,double*%test1_test_buzz_b)store[38xi8]c"44b9b452817d4d3ea103f1449105264c_close",[38xi8]*@__parallelGroupstore[38xi8]c"227b3938f885317dea9c644434cb82dd_start",[38xi8]*@__parallelGroupcallvoid@test1_test_fizz3(double*%test1_test_buzz_a,double*%test1_test_buzz_b)store[38xi8]c"227b3938f885317dea9c644434cb82dd_close",[38xi8]*@__parallelGroupstorei162,i16*@__roundsstore[38xi8]c"9a2b113b63e8232c2945f1018bf785f0_start",[38xi8]*@__parallelGroupcallvoid@test1_test_fizz(double*%test1_test_buzz_a,double*%test1_test_buzz_b)callvoid@test1_test_fizz2(double*%test1_test_buzz_a,double*%test1_test_buzz_b)store[38xi8]c"9a2b113b63e8232c2945f1018bf785f0_close",[38xi8]*@__parallelGroupstore[38xi8]c"87c9dfef940c096cf145af18149d3600_start",[38xi8]*@__parallelGroupcallvoid@test1_test_fizz3(double*%test1_test_buzz_a,double*%test1_test_buzz_b)store[38xi8]c"87c9dfef940c096cf145af18149d3600_close",[38xi8]*@__parallelGroupstorei163,i16*@__roundsstore[38xi8]c"cb7fdc02d16d31723661579b54e31084_start",[38xi8]*@__parallelGroupcallvoid@test1_test_fizz(double*%test1_test_buzz_a,double*%test1_test_buzz_b)callvoid@test1_test_fizz2(double*%test1_test_buzz_a,double*%test1_test_buzz_b)store[38xi8]c"cb7fdc02d16d31723661579b54e31084_close",[38xi8]*@__parallelGroupstore[38xi8]c"70ad8bfb0509411d97738ac929bc3d01_start",[38xi8]*@__parallelGroupcallvoid@test1_test_fizz3(double*%test1_test_buzz_a,double*%test1_test_buzz_b)store[38xi8]c"70ad8bfb0509411d97738ac929bc3d01_close",[38xi8]*@__parallelGroupstorei164,i16*@__roundsstore[38xi8]c"7e3f85f9630519ec31a508b611b1d4bb_start",[38xi8]*@__parallelGroupcallvoid@test1_test_fizz(double*%test1_test_buzz_a,double*%test1_test_buzz_b)callvoid@test1_test_fizz2(double*%test1_test_buzz_a,double*%test1_test_buzz_b)store[38xi8]c"7e3f85f9630519ec31a508b611b1d4bb_close",[38xi8]*@__parallelGroupstore[38xi8]c"e57d2299cbd9024a113885402ef4e089_start",[38xi8]*@__parallelGroupcallvoid@test1_test_fizz3(double*%test1_test_buzz_a,double*%test1_test_buzz_b)store[38xi8]c"e57d2299cbd9024a113885402ef4e089_close",[38xi8]*@__parallelGroupstorei165,i16*@__roundsstore[38xi8]c"7fbb0459ad7da0f1a336cf5de1cf9068_start",[38xi8]*@__parallelGroupcallvoid@test1_test_fizz(double*%test1_test_buzz_a,double*%test1_test_buzz_b)callvoid@test1_test_fizz2(double*%test1_test_buzz_a,double*%test1_test_buzz_b)store[38xi8]c"7fbb0459ad7da0f1a336cf5de1cf9068_close",[38xi8]*@__parallelGroupstore[38xi8]c"40b3487db7f69f408810fb4cb8b544eb_start",[38xi8]*@__parallelGroupcallvoid@test1_test_fizz3(double*%test1_test_buzz_a,double*%test1_test_buzz_b)store[38xi8]c"40b3487db7f69f408810fb4cb8b544eb_close",[38xi8]*@__parallelGroupretvoid}definevoid@test1_test_fizz(double*%test1_test_buzz_a,double*%test1_test_buzz_b){block-2:%0=loaddouble,double*%test1_test_buzz_b%1=loaddouble,double*@test1_a%2=loaddouble,double*%test1_test_buzz_a%3=fadddouble%1,%2%4=fadddouble%0,%3storedouble%4,double*%test1_test_buzz_bretvoid}definevoid@test1_test_fizz2(double*%test1_test_buzz_a,double*%test1_test_buzz_b){block-3:%0=loaddouble,double*%test1_test_buzz_b%1=loaddouble,double*%test1_test_buzz_a%2=loaddouble,double*@test1_b%3=fsubdouble%1,%2%4=fadddouble%0,%3storedouble%4,double*%test1_test_buzz_bretvoid}definevoid@test1_test_fizz3(double*%test1_test_buzz_a,double*%test1_test_buzz_b){block-4:%0=loaddouble,double*%test1_test_buzz_a%1=loaddouble,double*%test1_test_buzz_b%2=loaddouble,double*@test1_b%3=fadddouble%1,%2%4=fadddouble%0,%3storedouble%4,double*%test1_test_buzz_aretvoid}`
	//Should fadd have variable names or the values in those variables?

	llvm, err := prepTest(test, true)

	if err != nil {
		t.Fatalf("compilation failed on valid spec. got=%s", err)
	}
	ir, err := validateIR(llvm)

	if err != nil {
		t.Fatalf("generated IR is not valid. got=%s", err)
	}

	err = compareResults(llvm, expecting, string(ir))

	if err != nil {
		t.Fatal(err.Error())
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

			for 1 init{test = new foo;} run{
				test.fizz;
			};
	`

	expecting := `@__rounds=globali160@__parallelGroup=global[38xi8]c"\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00"@__choiceGroup=global[38xi8]c"\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00"@test1_a=globaldouble0x4002666666666666@test1_b=globaldouble2.0definevoid@__run(){block-5:storei160,i16*@__rounds%test1_test_buzz_a=allocadoublestoredouble10.0,double*%test1_test_buzz_a%test1_test_buzz_b=allocadoublestoredouble20.0,double*%test1_test_buzz_bstorei161,i16*@__roundsstore[38xi8]c"d44e0a3fc2944aa552d9118f291d3106_start",[38xi8]*@__parallelGroupcallvoid@test1_test_fizz(double*%test1_test_buzz_a,double*%test1_test_buzz_b)store[38xi8]c"d44e0a3fc2944aa552d9118f291d3106_close",[38xi8]*@__parallelGroupretvoid}definevoid@test1_test_fizz(double*%test1_test_buzz_a,double*%test1_test_buzz_b){block-6:%0=loaddouble,double*%test1_test_buzz_a%1=fcmpogtdouble%0,2.0bri1%1,label%block-8-true,label%block-9-falseblock-7-after:%2=loaddouble,double*%test1_test_buzz_b%3=fsubdouble%2,1.0storedouble%3,double*%test1_test_buzz_bretvoidblock-8-true:%4=loaddouble,double*%test1_test_buzz_a%5=loaddouble,double*@test1_b%6=fsubdouble%4,%5storedouble%6,double*%test1_test_buzz_abrlabel%block-7-afterblock-9-false:storedouble10.0,double*%test1_test_buzz_abrlabel%block-7-after}`

	llvm, err := prepTest(test, true)

	if err != nil {
		t.Fatalf("compilation failed on valid spec. got=%s", err)
	}

	ir, err := validateIR(llvm)

	if err != nil {
		t.Fatalf("generated IR is not valid. got=%s", err)
	}

	err = compareResults(llvm, expecting, string(ir))

	if err != nil {
		t.Fatal(err.Error())
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

	for 5 init{t = new test;} run {
		t.bar;
	};
	`

	expecting := `@__rounds=globali160@__parallelGroup=global[38xi8]c"\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00"@__choiceGroup=global[38xi8]c"\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00"@test1_a=globaldouble0x3DA3CA8CB153A753@test1_b=globaldouble0x3DA3CA8CB153A753definevoid@__run(){block-10:storei160,i16*@__rounds%test1_t_u_x=allocadoublestoredouble0x3DA3CA8CB153A753,double*%test1_t_u_xstorei161,i16*@__roundsstore[38xi8]c"4614d4e08724f278c8ce39e50955edbc_start",[38xi8]*@__parallelGroupcallvoid@test1_t_bar(double*%test1_t_u_x)store[38xi8]c"4614d4e08724f278c8ce39e50955edbc_close",[38xi8]*@__parallelGroupstorei162,i16*@__roundsstore[38xi8]c"ba735eefbef72f20ea6a264b981e9285_start",[38xi8]*@__parallelGroupcallvoid@test1_t_bar(double*%test1_t_u_x)store[38xi8]c"ba735eefbef72f20ea6a264b981e9285_close",[38xi8]*@__parallelGroupstorei163,i16*@__roundsstore[38xi8]c"fa06e912698cf4825866672ced835870_start",[38xi8]*@__parallelGroupcallvoid@test1_t_bar(double*%test1_t_u_x)store[38xi8]c"fa06e912698cf4825866672ced835870_close",[38xi8]*@__parallelGroupstorei164,i16*@__roundsstore[38xi8]c"f3a3858248090df83b9702c4852e0e28_start",[38xi8]*@__parallelGroupcallvoid@test1_t_bar(double*%test1_t_u_x)store[38xi8]c"f3a3858248090df83b9702c4852e0e28_close",[38xi8]*@__parallelGroupstorei165,i16*@__roundsstore[38xi8]c"ba20b5c159a59aeb04358b812e68f2d2_start",[38xi8]*@__parallelGroupcallvoid@test1_t_bar(double*%test1_t_u_x)store[38xi8]c"ba20b5c159a59aeb04358b812e68f2d2_close",[38xi8]*@__parallelGroupretvoid}definevoid@test1_t_bar(double*%test1_t_u_x){block-11:%0=loaddouble,double*%test1_t_u_x%1=loaddouble,double*@test1_a%2=loaddouble,double*@test1_b%3=fadddouble%1,%2%4=fadddouble%0,%3storedouble%4,double*%test1_t_u_xretvoid}
	`

	llvm, err := prepTest(test, true)

	if err != nil {
		t.Fatalf("compilation failed on valid spec. got=%s", err)
	}

	ir, err := validateIR(llvm)

	if err != nil {
		t.Fatalf("generated IR is not valid. got=%s", err)
	}

	err = compareResults(llvm, expecting, string(ir))

	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestParamReset(t *testing.T) {
	structs := make(map[string]*preprocess.SpecRecord)
	c := NewCompiler()
	c.LoadMeta(structs, make(map[string][]float64), []string{}, make(map[string]string), true)
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

	n := negate(test, false)

	if n.(*ast.InfixExpression).Operator != "!=" {
		t.Fatalf("operator has not been negated got=%s", n.(*ast.InfixExpression).Operator)
	}

	test2 := &ast.Boolean{Value: true}

	n2 := negate(test2, true)

	if n2.(*ast.Boolean).Value != false {
		t.Fatalf("boolean has not been negated got=%s", n2.(*ast.Boolean).String())
	}

	test3 := &ast.Boolean{Value: false}

	n3 := negate(test3, true)

	if n3.(*ast.Boolean).Value != true {
		t.Fatalf("boolean has not been negated got=%s", n3.(*ast.Boolean).String())
	}

	test4 := &ast.PrefixExpression{
		Operator: "!",
		Right:    &ast.Boolean{Value: false},
	}

	n4 := negate(test4, true)

	if n4.(*ast.Boolean).Value != false {
		t.Fatalf("boolean has not been negated got=%v", n4.(*ast.Boolean).Value)
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
	c.specStructs["test"].Index("COMPONENT", "this")
	if !c.isVarSet(test) {
		t.Fatal("isVarSet returned false on component, should return true")
	}

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
	c.specStructs["test"].Index("FLOW", "this")
	if !c.isVarSet(test) {
		t.Fatal("isVarSet returned false on flow, should return true")
	}

	if !c.isVarSet(test1) {
		t.Fatal("isStrVarSet returned false on a flow var, should return true")
	}

	c.specStructs["test"] = preprocess.NewSpecRecord()

	c.specStructs["test"].AddStock("this", val)
	c.specStructs["test"].Index("STOCK", "this")
	if !c.isVarSet(test) {
		t.Fatal("isVarSet returned false on stock, should return true")
	}

	if !c.isVarSet(test1) {
		t.Fatal("isStrVarSet returned false on a stock var, should return true")
	}

}

func TestUncertains(t *testing.T) {
	c := NewCompiler()
	c.specs["test"] = NewCompiledSpec("test")
	test := &ast.StructInstance{Spec: "test", Name: "foo", Parent: []string{"test", "zoo"}, Order: []string{"bar"}, ProcessedName: []string{"test", "foo"}, Properties: map[string]*ast.StructProperty{"bar": {Spec: "test", Name: "bar", ProcessedName: []string{"test", "foo", "bar"}, Value: &ast.Uncertain{Mean: 2.0, Sigma: .3, ProcessedName: []string{"test", "foo", "bar"}}}}}
	c.processStruct(test)

	if len(c.RawInputs.Uncertains["test_foo_bar"]) == 0 {
		t.Fatal("uncertain value not stored")
	}

	if c.RawInputs.Uncertains["test_foo_bar"][0] != 2.0 || c.RawInputs.Uncertains["test_foo_bar"][1] != .3 {
		t.Fatalf("uncertain stored value is incorrect, got=%f", c.RawInputs.Uncertains["test_foo_bar"])
	}

}

func TestUnknowns2(t *testing.T) {
	c := NewCompiler()
	c.specs["test"] = NewCompiledSpec("test")
	test := &ast.StructInstance{Spec: "test", Name: "foo", Parent: []string{"test", "zoo"}, Order: []string{"bar"}, ProcessedName: []string{"test", "foo"}, Properties: map[string]*ast.StructProperty{"bar": {Spec: "test", Name: "bar", ProcessedName: []string{"test", "foo", "bar"}, Value: &ast.Unknown{Name: &ast.Identifier{Spec: "test", Value: "bar"}, ProcessedName: []string{"test", "foo", "bar"}}}}}
	c.processStruct(test)

	if len(c.RawInputs.Unknowns) == 0 {
		t.Fatal("unknown value not stored")
	}

	if c.RawInputs.Unknowns[0] != "test_foo_bar" {
		t.Fatalf("unknowns stored value is incorrect, got=%s", c.RawInputs.Unknowns[0])
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

	expecting := `@__rounds=globali160@__parallelGroup=global[38xi8]c"\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00"@__choiceGroup=global[38xi8]c"\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00"definevoid@__run(){block-16:%test_foo_x=allocadoublestoredouble8.0,double*%test_foo_x%test_foo_initial=allocai1storei1false,i1*%test_foo_initial%test_foo_alarm=allocai1storei1false,i1*%test_foo_alarmstorei1true,i1*%test_foo_initialcallvoid@test_foo_initial__state(i1*%test_foo_alarm,i1*%test_foo_initial,double*%test_foo_x)callvoid@test_foo_alarm__state(i1*%test_foo_alarm,i1*%test_foo_initial,double*%test_foo_x)retvoid}definevoid@test_foo_initial__state(i1*%test_foo_alarm,i1*%test_foo_initial,double*%test_foo_x){block-17:%0=loadi1,i1*%test_foo_initial%1=icmpeqi1%0,truebri1%1,label%block-19-true,label%block-18-afterblock-18-after:%2=loadi1,i1*%test_foo_initial%3=icmpeqi1%2,true%4=loaddouble,double*%test_foo_x%5=fcmpogtdouble%4,10.0%6=andi1%3,%5bri1%6,label%block-22-true,label%block-21-afterblock-19-true:%7=alloca[14xi8]store[14xi8]c"test_foo_alarm",[14xi8]*%7%8=bitcast[14xi8]*%7toi8*%9=calli1@advance(i8*%8)brlabel%block-18-afterblock-21-after:retvoidblock-22-true:%10=calli1@stay()brlabel%block-21-after}definei1@advance(i8*%toState){block-20:reti1true}definei1@stay(){block-23:reti1true}definevoid@test_foo_alarm__state(i1*%test_foo_alarm,i1*%test_foo_initial,double*%test_foo_x){block-24:%0=loadi1,i1*%test_foo_alarm%1=icmpeqi1%0,truebri1%1,label%block-26-true,label%block-25-afterblock-25-after:retvoidblock-26-true:%2=alloca[14xi8]store[14xi8]c"test_foo_close",[14xi8]*%2%3=bitcast[14xi8]*%2toi8*%4=calli1@advance(i8*%3)brlabel%block-25-after}`

	llvm, err := prepTest(test, false)

	if err != nil {
		t.Fatalf("compilation failed on valid spec. got=%s", err)
	}

	ir, err := validateIR(llvm)

	if err != nil {
		t.Fatalf("generated IR is not valid. got=%s", err)
	}

	err = compareResults(llvm, expecting, string(ir))

	if err != nil {
		t.Fatal(err.Error())
	}

}

func TestChoose(t *testing.T) {
	test := `
	system test;

	component foo = states{
		initial: func{
			choose stay() || advance(this.alarm);
		},
		alarm: func{
			advance(this.close);
		},
		close: func{
			stay();
		},
	};

	start {
		foo: initial,
	};
	`

	expecting := ``

	llvm, err := prepTest(test, false)

	if err != nil {
		t.Fatalf("compilation failed on valid spec. got=%s", err)
	}

	fmt.Println(llvm)
	ir, err := validateIR(llvm)

	if err != nil {
		t.Fatalf("generated IR is not valid. got=%s", err)
	}

	err = compareResults(llvm, expecting, string(ir))

	if err != nil {
		t.Fatal(err.Error())
	}

}

func TestIndexExp(t *testing.T) {
	test := `spec test1;
			
			def foo = flow{
				buzz: new bar,
				fizz: func{
					buzz.a = buzz.a[1] - 2;  
				},
			};

			def bar = stock{
				a: 10,
			};

			for 1 init{test = new foo;} run{
				test.fizz;
			};
	`

	expecting := `@__rounds=globali160@__parallelGroup=global[38xi8]c"\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00"@__choiceGroup=global[38xi8]c"\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00"@test1_test_buzz_a_1=globaldouble0x3DA3CA8CB153A753definevoid@__run(){block-27:storei160,i16*@__rounds%test1_test_buzz_a=allocadoublestoredouble10.0,double*%test1_test_buzz_astorei161,i16*@__roundsstore[38xi8]c"7f977975ebdfc28b778ed4618a0af327_start",[38xi8]*@__parallelGroupcallvoid@test1_test_fizz(double*%test1_test_buzz_a)store[38xi8]c"7f977975ebdfc28b778ed4618a0af327_close",[38xi8]*@__parallelGroupretvoid}definevoid@test1_test_fizz(double*%test1_test_buzz_a){block-28:%0=loaddouble,double*@test1_test_buzz_a_1%1=fsubdouble%0,2.0storedouble%1,double*%test1_test_buzz_aretvoid}`

	llvm, err := prepTest(test, true)

	if err != nil {
		t.Fatalf("compilation failed on valid spec. got=%s", err)
	}

	ir, err := validateIR(llvm)

	if err != nil {
		t.Fatalf("generated IR is not valid. got=%s", err)
	}

	err = compareResults(llvm, expecting, string(ir))

	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestStringExp(t *testing.T) {
	test := `spec test;
		str1 = "is a fish";
		str2 = "tastes delicious with ginger";
		str3 = "native to North America";

		assume str1 && str3;
		assert str3;
	`

	expecting := `@__rounds=globali160@__parallelGroup=global[38xi8]c"\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00"@__choiceGroup=global[38xi8]c"\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00"@test_str1=globali1false@test_str2=globali1false@test_str3=globali1falsedefinevoid@__run(){block-29:retvoid}
	`

	llvm, err := prepTest(test, true)

	if err != nil {
		t.Fatalf("compilation failed on valid spec. got=%s", err)
	}

	ir, err := validateIR(llvm)

	if err != nil {
		t.Fatalf("generated IR is not valid. got=%s", err)
	}

	err = compareResults(llvm, expecting, string(ir))

	if err != nil {
		t.Fatal(err.Error())
	}
}

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

func prepTest(test string, specType bool) (string, error) {
	flags := make(map[string]bool)
	flags["specType"] = specType
	flags["testing"] = true
	flags["skipRun"] = false

	l := listener.Execute(test, "", flags)
	pre := preprocess.Execute(l)

	ty := types.Execute(pre.Processed, pre)
	sw := swaps.NewPrecompiler(ty)
	tree := sw.Swap(ty.Checked)
	compiler := NewCompiler()
	compiler.LoadMeta(ty.SpecStructs, l.Uncertains, l.Unknowns, sw.Alias, true)
	err := compiler.Compile(tree)

	if err != nil {
		return "", err
	}
	//fmt.Println(compiler.GetIR())
	return compiler.GetIR(), err
}

func validateIR(ir string) ([]byte, error) {
	//Run LLVM optimizer to check IR is valid
	cmd := exec.Command("opt", "-S", "--passes=mem2reg")
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
