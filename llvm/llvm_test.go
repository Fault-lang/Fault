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
	"regexp"
	"strings"
	"testing"
	"unicode"

	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	irtypes "github.com/llir/llvm/ir/types"
)

var blockNumRe = regexp.MustCompile(`block-\d+`)
var groupHashRe = regexp.MustCompile(`[0-9a-f]{32}`)

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

			run init{test = new foo;} {
				test.fizz | test.fizz2;
				test.fizz3;
				test.fizz | test.fizz2;
				test.fizz3;
				test.fizz | test.fizz2;
				test.fizz3;
				test.fizz | test.fizz2;
				test.fizz3;
				test.fizz | test.fizz2;
				test.fizz3;
			};
	`

	expecting := `@__rounds=globali160@__parallelGroup=global[38xi8]c"\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00"@__choiceGroup=global[38xi8]c"\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00"@test1_a=globaldouble0x4002666666666666@test1_b=globaldouble2.0definevoid@__run(){block:storei160,i16*@__rounds%test1_test_buzz_a=allocadoublestoredouble10.0,double*%test1_test_buzz_a%test1_test_buzz_b=allocadoublestoredouble20.0,double*%test1_test_buzz_bstorei161,i16*@__roundsstore[38xi8]c"44b9b452817d4d3ea103f1449105264c_start",[38xi8]*@__parallelGroupcallvoid@test1_test_fizz(double*%test1_test_buzz_a,double*%test1_test_buzz_b)callvoid@test1_test_fizz2(double*%test1_test_buzz_a,double*%test1_test_buzz_b)store[38xi8]c"44b9b452817d4d3ea103f1449105264c_close",[38xi8]*@__parallelGroupstorei162,i16*@__roundscallvoid@test1_test_fizz3(double*%test1_test_buzz_a,double*%test1_test_buzz_b)storei163,i16*@__roundsstore[38xi8]c"9a2b113b63e8232c2945f1018bf785f0_start",[38xi8]*@__parallelGroupcallvoid@test1_test_fizz(double*%test1_test_buzz_a,double*%test1_test_buzz_b)callvoid@test1_test_fizz2(double*%test1_test_buzz_a,double*%test1_test_buzz_b)store[38xi8]c"9a2b113b63e8232c2945f1018bf785f0_close",[38xi8]*@__parallelGroupstorei164,i16*@__roundscallvoid@test1_test_fizz3(double*%test1_test_buzz_a,double*%test1_test_buzz_b)storei165,i16*@__roundsstore[38xi8]c"cb7fdc02d16d31723661579b54e31084_start",[38xi8]*@__parallelGroupcallvoid@test1_test_fizz(double*%test1_test_buzz_a,double*%test1_test_buzz_b)callvoid@test1_test_fizz2(double*%test1_test_buzz_a,double*%test1_test_buzz_b)store[38xi8]c"cb7fdc02d16d31723661579b54e31084_close",[38xi8]*@__parallelGroupstorei166,i16*@__roundscallvoid@test1_test_fizz3(double*%test1_test_buzz_a,double*%test1_test_buzz_b)storei167,i16*@__roundsstore[38xi8]c"7e3f85f9630519ec31a508b611b1d4bb_start",[38xi8]*@__parallelGroupcallvoid@test1_test_fizz(double*%test1_test_buzz_a,double*%test1_test_buzz_b)callvoid@test1_test_fizz2(double*%test1_test_buzz_a,double*%test1_test_buzz_b)store[38xi8]c"7e3f85f9630519ec31a508b611b1d4bb_close",[38xi8]*@__parallelGroupstorei168,i16*@__roundscallvoid@test1_test_fizz3(double*%test1_test_buzz_a,double*%test1_test_buzz_b)storei169,i16*@__roundsstore[38xi8]c"7fbb0459ad7da0f1a336cf5de1cf9068_start",[38xi8]*@__parallelGroupcallvoid@test1_test_fizz(double*%test1_test_buzz_a,double*%test1_test_buzz_b)callvoid@test1_test_fizz2(double*%test1_test_buzz_a,double*%test1_test_buzz_b)store[38xi8]c"7fbb0459ad7da0f1a336cf5de1cf9068_close",[38xi8]*@__parallelGroupstorei1610,i16*@__roundscallvoid@test1_test_fizz3(double*%test1_test_buzz_a,double*%test1_test_buzz_b)retvoid}definevoid@test1_test_fizz(double*%test1_test_buzz_a,double*%test1_test_buzz_b){block:%0=loaddouble,double*%test1_test_buzz_b%1=loaddouble,double*@test1_a%2=loaddouble,double*%test1_test_buzz_a%3=fadddouble%1,%2%4=fadddouble%0,%3storedouble%4,double*%test1_test_buzz_bretvoid}definevoid@test1_test_fizz2(double*%test1_test_buzz_a,double*%test1_test_buzz_b){block:%0=loaddouble,double*%test1_test_buzz_b%1=loaddouble,double*%test1_test_buzz_a%2=loaddouble,double*@test1_b%3=fsubdouble%1,%2%4=fadddouble%0,%3storedouble%4,double*%test1_test_buzz_bretvoid}definevoid@test1_test_fizz3(double*%test1_test_buzz_a,double*%test1_test_buzz_b){block:%0=loaddouble,double*%test1_test_buzz_a%1=loaddouble,double*%test1_test_buzz_b%2=loaddouble,double*@test1_b%3=fadddouble%1,%2%4=fadddouble%0,%3storedouble%4,double*%test1_test_buzz_aretvoid}`
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

func TestRunStmtExplicit(t *testing.T) {
	test := `spec test1;
			def foo = flow{
				buzz: new bar,
				fizz: func{
					buzz.a <- buzz.a + 1;
				},
			};

			def bar = stock{
				a: 10,
			};

			run init{test = new foo;} {
				test.fizz;
			};
	`

	llvm, err := prepTest(test, true)
	if err != nil {
		t.Fatalf("compilation failed on valid spec. got=%s", err)
	}

	ir, err := validateIR(llvm)
	if err != nil {
		t.Fatalf("generated IR is not valid. got=%s", err)
	}

	// __synthStep should NOT appear — no solvable steps
	if strings.Contains(string(ir), "__synthStep") {
		t.Fatalf("__synthStep marker should not appear in IR without solvable steps")
	}

	// The flow function should be defined and called
	if !strings.Contains(string(ir), "test1_test_fizz") {
		t.Fatalf("expected test1_test_fizz function in IR, got:\n%s", string(ir))
	}
}

func TestRunStmtSolvable(t *testing.T) {
	test := `spec test1;
			def foo = flow{
				buzz: new bar,
				fizz: func{
					buzz.a <- buzz.a + 1;
				},
			};

			def bar = stock{
				a: 10,
			};

			run init{test = new foo;} {
				__;
				test.fizz;
				__;
			};
	`

	llvm, err := prepTest(test, true)
	if err != nil {
		t.Fatalf("compilation failed on valid spec. got=%s", err)
	}

	ir, err := validateIR(llvm)
	if err != nil {
		t.Fatalf("generated IR is not valid. got=%s", err)
	}

	// __synthStep marker must appear
	if !strings.Contains(string(ir), "__synthStep") {
		t.Fatalf("expected __synthStep marker in IR for solvable steps, got:\n%s", string(ir))
	}

	// The explicit step should still compile
	if !strings.Contains(string(ir), "test1_test_fizz") {
		t.Fatalf("expected test1_test_fizz function in IR, got:\n%s", string(ir))
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

			run init{test = new foo;} {
				test.fizz;
			};
	`

	expecting := `@__rounds=globali160@__parallelGroup=global[38xi8]c"\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00"@__choiceGroup=global[38xi8]c"\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00"@test1_a=globaldouble0x4002666666666666@test1_b=globaldouble2.0definevoid@__run(){block-5:storei160,i16*@__rounds%test1_test_buzz_a=allocadoublestoredouble10.0,double*%test1_test_buzz_a%test1_test_buzz_b=allocadoublestoredouble20.0,double*%test1_test_buzz_bstorei161,i16*@__roundscallvoid@test1_test_fizz(double*%test1_test_buzz_a,double*%test1_test_buzz_b)retvoid}definevoid@test1_test_fizz(double*%test1_test_buzz_a,double*%test1_test_buzz_b){block-6:%0=loaddouble,double*%test1_test_buzz_a%1=fcmpogtdouble%0,2.0bri1%1,label%block-8-true,label%block-9-falseblock-7-after:%2=loaddouble,double*%test1_test_buzz_b%3=fsubdouble%2,1.0storedouble%3,double*%test1_test_buzz_bretvoidblock-8-true:%4=loaddouble,double*%test1_test_buzz_a%5=loaddouble,double*@test1_b%6=fsubdouble%4,%5storedouble%6,double*%test1_test_buzz_abrlabel%block-7-afterblock-9-false:storedouble10.0,double*%test1_test_buzz_abrlabel%block-7-after}`

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

	run init{t = new test;} {
		t.bar;
		t.bar;
		t.bar;
		t.bar;
		t.bar;
	};
	`

	expecting := `@__rounds=globali160@__parallelGroup=global[38xi8]c"\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00"@__choiceGroup=global[38xi8]c"\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00"@test1_a=globaldouble0x3DA3CA8CB153A753@test1_b=globaldouble0x3DA3CA8CB153A753definevoid@__run(){block-10:storei160,i16*@__rounds%test1_t_u_x=allocadoublestoredouble0x3DA3CA8CB153A753,double*%test1_t_u_xstorei161,i16*@__roundscallvoid@test1_t_bar(double*%test1_t_u_x)storei162,i16*@__roundscallvoid@test1_t_bar(double*%test1_t_u_x)storei163,i16*@__roundscallvoid@test1_t_bar(double*%test1_t_u_x)storei164,i16*@__roundscallvoid@test1_t_bar(double*%test1_t_u_x)storei165,i16*@__roundscallvoid@test1_t_bar(double*%test1_t_u_x)retvoid}definevoid@test1_t_bar(double*%test1_t_u_x){block-11:%0=loaddouble,double*%test1_t_u_x%1=loaddouble,double*@test1_a%2=loaddouble,double*@test1_b%3=fadddouble%1,%2%4=fadddouble%0,%3storedouble%4,double*%test1_t_u_xretvoid}
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
	c.LoadMeta(structs, make(map[string][]float64), []string{}, []string{}, []string{}, make(map[string]string), true)
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
	test := &ast.StructInstance{Spec: "test", Name: "foo", Parent: []string{"test", "zoo"}, Order: []string{"bar"}, ProcessedName: []string{"test", "foo"}, Properties: map[string]*ast.StructProperty{"bar": {Spec: "test", Name: "bar", ProcessedName: []string{"test", "foo", "bar"}, Value: &ast.Unknown{ProcessedName: []string{"test", "foo", "bar"}}}}}
	c.processStruct(test)

	if len(c.RawInputs.Unknowns) == 0 {
		t.Fatal("unknown value not stored")
	}

	if c.RawInputs.Unknowns[0] != "test_foo_bar" {
		t.Fatalf("unknowns stored value is incorrect, got=%s", c.RawInputs.Unknowns[0])
	}

}

func TestParamStored(t *testing.T) {
	c := NewCompiler()
	c.specs["test"] = NewCompiledSpec("test")
	test := &ast.StructInstance{
		Spec: "test", Name: "req", Parent: []string{"test", "zoo"},
		Order:         []string{"amount"},
		ProcessedName: []string{"test", "req"},
		Properties: map[string]*ast.StructProperty{
			"amount": {
				Spec: "test", Name: "amount",
				ProcessedName: []string{"test", "req", "amount"},
				Value:         &ast.Param{ProcessedName: []string{"test", "req", "amount"}, TypeHint: "REAL"},
			},
		},
	}
	c.processStruct(test)

	if len(c.RawInputs.Params) == 0 {
		t.Fatal("param value not stored in RawInputs.Params")
	}

	if c.RawInputs.Params[0] != "test_req_amount" {
		t.Fatalf("params stored value is incorrect, got=%s", c.RawInputs.Params[0])
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

	run {
		foo.initial;
	}
	`

	flags := make(map[string]bool)
	flags["specType"] = false
	flags["testing"] = true
	flags["skipRun"] = false

	l, _ := listener.Execute(test, "", flags)
	pre, err := preprocess.Execute(l)
	if err != nil {
		t.Fatalf("preprocessing failed: %s", err)
	}

	ty, err := types.Execute(pre.Processed, pre)
	if err != nil {
		t.Fatal(err)
	}
	sw := swaps.NewPrecompiler(ty)
	tree := sw.Swap(ty.Checked)
	compiler := NewCompiler()
	compiler.LoadMeta(ty.SpecStructs, l.Uncertains, l.Unknowns, l.Wholes, l.Params, sw.Alias, true)
	err = compiler.Compile(tree)
	if err != nil {
		t.Fatalf("compilation failed: %s", err)
	}

	ir := compiler.GetIR()

	// State activation should store true into initial state
	if !strings.Contains(ir, "store i1 true") {
		t.Fatal("IR should contain 'store i1 true' for state activation")
	}
	// Transition chain should be called
	if !strings.Contains(ir, "@test_foo_initial__state") {
		t.Fatal("IR should call @test_foo_initial__state")
	}
	if !strings.Contains(ir, "@test_foo_alarm__state") {
		t.Fatal("IR should call @test_foo_alarm__state")
	}

	_, err = validateIR(ir)
	if err != nil {
		t.Fatalf("generated IR is not valid: %s", err)
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

			run init{test = new foo;} {
				test.fizz;
			};
	`

	expecting := `@__rounds = global i16 0
@__parallelGroup = global [38 x i8] c"\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00"
@__choiceGroup = global [38 x i8] c"\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00"
@test1_test_buzz_a_1 = global double 0x3DA3CA8CB153A753

define void @__run() {
block-27:
	store i16 0, i16* @__rounds
	%test1_test_buzz_a = alloca double
	store double 10.0, double* %test1_test_buzz_a
	store i16 1, i16* @__rounds
	call void @test1_test_fizz(double* %test1_test_buzz_a)
	ret void
}

define void @test1_test_fizz(double* %test1_test_buzz_a) {
block-28:
	%0 = load double, double* @test1_test_buzz_a_1
	%1 = fsub double %0, 2.0
	store double %1, double* %test1_test_buzz_a
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

func TestStringExp(t *testing.T) {
	test := `spec test;
		str1 = "is a fish";
		str2 = "tastes delicious with ginger";
		str3 = "native to North America";

		assume str1 && str3;
		assert str3;
	`

	expecting := `@__rounds = global i16 0
@__parallelGroup = global [38 x i8] c"\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00"
@__choiceGroup = global [38 x i8] c"\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00"
@test_str1 = global i1 false
@test_str2 = global i1 false
@test_str3 = global i1 false

define void @__run() {
block-29:
	ret void
}`

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

	run {
		foo.initial;
	}
	`

	flags := make(map[string]bool)
	flags["specType"] = false
	flags["testing"] = true
	flags["skipRun"] = false

	l, _ := listener.Execute(test, "", flags)
	pre, err := preprocess.Execute(l)
	if err != nil {
		t.Fatalf("preprocessing failed: %s", err)
	}

	ty, err := types.Execute(pre.Processed, pre)
	if err != nil {
		t.Fatal(err)
	}
	sw := swaps.NewPrecompiler(ty)
	tree := sw.Swap(ty.Checked)
	compiler := NewCompiler()
	compiler.LoadMeta(ty.SpecStructs, l.Uncertains, l.Unknowns, l.Wholes, l.Params, sw.Alias, true)
	err = compiler.Compile(tree)
	if err != nil {
		t.Fatalf("compilation failed: %s", err)
	}

	ir := compiler.GetIR()

	if !strings.Contains(ir, "store i1 true") {
		t.Fatal("IR should contain 'store i1 true' for state activation")
	}
	if !strings.Contains(ir, "@test_foo_initial__state") {
		t.Fatal("IR should call @test_foo_initial__state")
	}
	if !strings.Contains(ir, "@test_foo_alarm__state") {
		t.Fatal("IR should call @test_foo_alarm__state")
	}
	if !strings.Contains(ir, "@test_foo_close__state") {
		t.Fatal("IR should call @test_foo_close__state")
	}
	if !strings.Contains(ir, "@__choiceGroup") {
		t.Fatal("IR should contain choice group markers")
	}

	_, err = validateIR(ir)
	if err != nil {
		t.Fatalf("generated IR is not valid: %s", err)
	}
}

func TestLeave(t *testing.T) {
	test := `
	system test;

	component foo = states{
		initial: func{
			advance(this.alarm) && leave();
		},
		alarm: func{
			stay();
		},
	};

	run {
		foo.initial;
	}
	`

	flags := make(map[string]bool)
	flags["specType"] = false
	flags["testing"] = true
	flags["skipRun"] = false

	l, _ := listener.Execute(test, "", flags)
	pre, err := preprocess.Execute(l)
	if err != nil {
		t.Fatalf("preprocessing failed: %s", err)
	}

	ty, err := types.Execute(pre.Processed, pre)
	if err != nil {
		t.Fatal(err)
	}
	sw := swaps.NewPrecompiler(ty)
	tree := sw.Swap(ty.Checked)
	compiler := NewCompiler()
	compiler.LoadMeta(ty.SpecStructs, l.Uncertains, l.Unknowns, l.Wholes, l.Params, sw.Alias, true)
	err = compiler.Compile(tree)
	if err != nil {
		t.Fatalf("compilation failed: %s", err)
	}

	ir := compiler.GetIR()

	if !strings.Contains(ir, "store i1 true") {
		t.Fatal("IR should contain 'store i1 true' for state activation")
	}
	if !strings.Contains(ir, "@test_foo_initial__state") {
		t.Fatal("IR should call @test_foo_initial__state")
	}
	if !strings.Contains(ir, "@test_foo_alarm__state") {
		t.Fatal("IR should call @test_foo_alarm__state")
	}
	if !strings.Contains(ir, "@leave") {
		t.Fatal("IR should contain @leave function")
	}

	_, err = validateIR(ir)
	if err != nil {
		t.Fatalf("generated IR is not valid: %s", err)
	}
}

func TestSysRunBlock(t *testing.T) {
	test := `
	system test;

	component a = states{
		on: func{
			stay();
		},
		off: func{
			stay();
		},
	};

	component b = states{
		idle: func{
			stay();
		},
	};

	run {
		a.on | b.idle;
	}
	`

	flags := make(map[string]bool)
	flags["specType"] = false
	flags["testing"] = true
	flags["skipRun"] = false

	l, _ := listener.Execute(test, "", flags)
	pre, err := preprocess.Execute(l)
	if err != nil {
		t.Fatalf("preprocessing failed: %s", err)
	}

	ty, err := types.Execute(pre.Processed, pre)
	if err != nil {
		t.Fatal(err)
	}
	sw := swaps.NewPrecompiler(ty)
	tree := sw.Swap(ty.Checked)
	compiler := NewCompiler()
	compiler.LoadMeta(ty.SpecStructs, l.Uncertains, l.Unknowns, l.Wholes, l.Params, sw.Alias, true)
	err = compiler.Compile(tree)
	if err != nil {
		t.Fatalf("compilation failed: %s", err)
	}

	ir := compiler.GetIR()

	// Both component state functions should appear in the IR
	if !strings.Contains(ir, "@test_a_on__state") {
		t.Fatal("IR should contain @test_a_on__state")
	}
	if !strings.Contains(ir, "@test_b_idle__state") {
		t.Fatal("IR should contain @test_b_idle__state")
	}
	_, err = validateIR(ir)
	if err != nil {
		t.Fatalf("generated IR is not valid: %s", err)
	}
}

func TestStateActivation(t *testing.T) {
	test := `
	system test;

	component a = states{
		on: func{
			stay();
		},
		off: func{
			stay();
		},
	};

	component b = states{
		idle: func{
			stay();
		},
	};

	run {
		a.on && b.idle;
	}
	`

	flags := make(map[string]bool)
	flags["specType"] = false
	flags["testing"] = true
	flags["skipRun"] = false

	l, _ := listener.Execute(test, "", flags)
	pre, err := preprocess.Execute(l)
	if err != nil {
		t.Fatalf("preprocessing failed: %s", err)
	}

	ty, err := types.Execute(pre.Processed, pre)
	if err != nil {
		t.Fatal(err)
	}
	sw := swaps.NewPrecompiler(ty)
	tree := sw.Swap(ty.Checked)
	compiler := NewCompiler()
	compiler.LoadMeta(ty.SpecStructs, l.Uncertains, l.Unknowns, l.Wholes, l.Params, sw.Alias, true)
	err = compiler.Compile(tree)
	if err != nil {
		t.Fatalf("compilation failed: %s", err)
	}

	ir := compiler.GetIR()

	// StartStates should be populated by the StateActivation step
	aPrefix := "test_a_"
	bPrefix := "test_b_"
	if _, ok := compiler.StartStates[aPrefix]; !ok {
		t.Fatalf("StartStates should contain prefix %q after StateActivation", aPrefix)
	}
	if _, ok := compiler.StartStates[bPrefix]; !ok {
		t.Fatalf("StartStates should contain prefix %q after StateActivation", bPrefix)
	}

	// The initial state store (true) should appear in the IR
	if !strings.Contains(ir, "store i1 true") {
		t.Fatal("IR should contain 'store i1 true' for state activation")
	}

	// Both component state functions should appear in the IR
	if !strings.Contains(ir, "@test_a_on__state") {
		t.Fatal("IR should contain @test_a_on__state")
	}
	if !strings.Contains(ir, "@test_b_idle__state") {
		t.Fatal("IR should contain @test_b_idle__state")
	}

	_, err = validateIR(ir)
	if err != nil {
		t.Fatalf("generated IR is not valid: %s", err)
	}
}

func TestUnfuncCompiles(t *testing.T) {
	test := `
	system test;

	component fetch = states{
		id: false,
		count: false,
		countVotes: unfunc{
			requires fetch.id,
			emits fetch.count,
		},
	};

	run {
		fetch.countVotes;
	}
	`

	flags := make(map[string]bool)
	flags["specType"] = false
	flags["testing"] = true
	flags["skipRun"] = false

	l, _ := listener.Execute(test, "", flags)
	pre, err := preprocess.Execute(l)
	if err != nil {
		t.Fatalf("preprocessing failed: %s", err)
	}

	ty, err := types.Execute(pre.Processed, pre)
	if err != nil {
		t.Fatal(err)
	}
	sw := swaps.NewPrecompiler(ty)
	tree := sw.Swap(ty.Checked)
	compiler := NewCompiler()
	compiler.LoadMeta(ty.SpecStructs, l.Uncertains, l.Unknowns, l.Wholes, l.Params, sw.Alias, true)
	err = compiler.Compile(tree)
	if err != nil {
		t.Fatalf("compilation failed: %s", err)
	}

	ir := compiler.GetIR()

	// Stub function should appear in LLVM IR
	if !strings.Contains(ir, "@test_fetch_countVotes__state") {
		t.Fatal("IR should contain @test_fetch_countVotes__state")
	}

	// ComponentOrder should include the unfunc state
	found := false
	for _, k := range compiler.ComponentOrder {
		if strings.Contains(k, "countVotes") {
			found = true
			break
		}
	}
	if !found {
		t.Fatalf("ComponentOrder should include countVotes, got %v", compiler.ComponentOrder)
	}

	// RawInputs.Unfuncs should contain the unfunc info
	if len(compiler.RawInputs.Unfuncs) == 0 {
		t.Fatal("RawInputs.Unfuncs should not be empty")
	}
	uf := compiler.RawInputs.Unfuncs[0]
	if uf.Requires == nil {
		t.Fatal("UnfuncInfo.Requires should not be nil")
	}
	if uf.Emits == nil {
		t.Fatal("UnfuncInfo.Emits should not be nil")
	}

	_, err = validateIR(ir)
	if err != nil {
		t.Fatalf("generated IR is not valid: %s", err)
	}
}

func TestUnfuncWithMultipleRequires(t *testing.T) {
	test := `
	system test;

	component ops = states{
		id: false,
		joinId: false,
		result: false,
		getWithJoin: unfunc{
			requires ops.id && ops.joinId,
			emits ops.result,
		},
	};

	run {
		ops.getWithJoin;
	}
	`

	flags := make(map[string]bool)
	flags["specType"] = false
	flags["testing"] = true
	flags["skipRun"] = false

	l, _ := listener.Execute(test, "", flags)
	pre, err := preprocess.Execute(l)
	if err != nil {
		t.Fatalf("preprocessing failed: %s", err)
	}

	ty, err := types.Execute(pre.Processed, pre)
	if err != nil {
		t.Fatal(err)
	}
	sw := swaps.NewPrecompiler(ty)
	tree := sw.Swap(ty.Checked)
	compiler := NewCompiler()
	compiler.LoadMeta(ty.SpecStructs, l.Uncertains, l.Unknowns, l.Wholes, l.Params, sw.Alias, true)
	err = compiler.Compile(tree)
	if err != nil {
		t.Fatalf("compilation failed: %s", err)
	}

	if len(compiler.RawInputs.Unfuncs) == 0 {
		t.Fatal("RawInputs.Unfuncs should not be empty")
	}

	uf := compiler.RawInputs.Unfuncs[0]
	if uf.Requires == nil {
		t.Fatal("Requires should not be nil for getWithJoin")
	}
	// The requires expression should be an InfixExpression (&&)
	if _, ok := uf.Requires.(*ast.InfixExpression); !ok {
		t.Fatalf("Requires should be an InfixExpression for &&, got %T", uf.Requires)
	}

	_, err = validateIR(compiler.GetIR())
	if err != nil {
		t.Fatalf("generated IR is not valid: %s", err)
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
	str = blockNumRe.ReplaceAllString(str, "block")
	str = groupHashRe.ReplaceAllString(str, "HASH")
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

func TestStockExtendsAssertPropagation(t *testing.T) {
	// An assert on parent.field should also constrain child.field when child extends parent.
	test := `spec test1;
def generic = stock{
	id: "primary key",
};
def person = stock{
	extends generic,
	name: "person name",
};
assert generic.id == true;
`
	compiler, err := prepTestCompiler(test, true)
	if err != nil {
		t.Fatalf("compilation failed: %s", err)
	}

	if len(compiler.RawInputs.Asserts) == 0 {
		t.Fatal("no asserts compiled")
	}

	assertVar, ok := compiler.RawInputs.Asserts[0].Constraint.Left.(*ast.AssertVar)
	if !ok {
		t.Fatalf("assert left is not AssertVar, got %T", compiler.RawInputs.Asserts[0].Constraint.Left)
	}

	// Should contain both generic_id and person_id
	found := make(map[string]bool)
	for _, inst := range assertVar.Instances {
		found[inst] = true
	}

	if !found["test1_generic_id"] {
		t.Errorf("assert did not include test1_generic_id, got %v", assertVar.Instances)
	}
	if !found["test1_person_id"] {
		t.Errorf("assert did not propagate to test1_person_id, got %v", assertVar.Instances)
	}
}

func prepTestCompiler(test string, specType bool) (*Compiler, error) {
	flags := make(map[string]bool)
	flags["specType"] = specType
	flags["testing"] = true
	flags["skipRun"] = false

	l, _ := listener.Execute(test, "", flags)
	pre, err := preprocess.Execute(l)
	if err != nil {
		return nil, err
	}

	ty, err := types.Execute(pre.Processed, pre)
	if err != nil {
		return nil, err
	}
	sw := swaps.NewPrecompiler(ty)
	tree := sw.Swap(ty.Checked)
	compiler := NewCompiler()
	compiler.LoadMeta(ty.SpecStructs, l.Uncertains, l.Unknowns, l.Wholes, l.Params, sw.Alias, true)
	err = compiler.Compile(tree)
	return compiler, err
}

func prepTest(test string, specType bool) (string, error) {
	flags := make(map[string]bool)
	flags["specType"] = specType
	flags["testing"] = true
	flags["skipRun"] = false

	l, _ := listener.Execute(test, "", flags)
	pre, err := preprocess.Execute(l)
	if err != nil {
		return "", err
	}

	ty, err := types.Execute(pre.Processed, pre)
	if err != nil {
		return "", err
	}
	sw := swaps.NewPrecompiler(ty)
	tree := sw.Swap(ty.Checked)
	compiler := NewCompiler()
	compiler.LoadMeta(ty.SpecStructs, l.Uncertains, l.Unknowns, l.Wholes, l.Params, sw.Alias, true)
	err = compiler.Compile(tree)

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
