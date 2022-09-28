package llvm

import (
	"fault/ast"
	"fault/listener"
	"fault/parser"
	"fault/types"
	"fmt"
	"io"
	"os/exec"
	"strings"
	"testing"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
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
	expecting := `@test1_x = global double 2.0
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

func TestConstExpr(t *testing.T) {
	// This probably needs to change LLVM constants don't allow
	// for expressions and should changed ( a + 5, changes with the value of a)

	test := `spec test1;
			const a = 2+3;
			const b = 2-3.4;
			const c = 2.1 * 3;
			const d = 2/3;
			const e = 2 > 3;
			const f = 2 < 3;`
	expecting := `define void @__run() {
		block-1:
			%0 = fadd double 2.0, 3.0
			%test1_a = alloca double
			store double %0, double* %test1_a
			%1 = fsub double 2.0, 0x400B333333333333
			%test1_b = alloca double
			store double %1, double* %test1_b
			%2 = fmul double 0x4000CCCCCCCCCCCD, 3.0
			%test1_c = alloca double
			store double %2, double* %test1_c
			%3 = fdiv double 2.0, 3.0
			%test1_d = alloca double
			store double %3, double* %test1_d
			%4 = fcmp ogt double 2.0, 3.0
			%test1_e = alloca i1
			store i1 %4, i1* %test1_e
			%5 = fcmp olt double 2.0, 3.0
			%test1_f = alloca i1
			store i1 %5, i1* %test1_f
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
					a + buzz.a -> buzz.b;
				},
				fizz2: func{
					buzz.a - b -> buzz.b;
				},
				fizz3: func{
					buzz.b + b -> buzz.a;
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

	expecting := `@test1_a = global double 0x4002666666666666
	@test1_b = global double 2.0
	
	define void @__run() {
	block-2:
		%test1_test_buzz_a = alloca double
		store double 10.0, double* %test1_test_buzz_a
		%test1_test_buzz_b = alloca double
		store double 20.0, double* %test1_test_buzz_b
		call void @test1_test_fizz(double* %test1_test_buzz_a, double* %test1_test_buzz_b), !\344b9b452817d4d3ea103f1449105264c !DIBasicType(tag: DW_TAG_string_type)
		call void @test1_test_fizz2(double* %test1_test_buzz_a, double* %test1_test_buzz_b), !\344b9b452817d4d3ea103f1449105264c !DIBasicType(tag: DW_TAG_string_type)
		call void @test1_test_fizz3(double* %test1_test_buzz_a, double* %test1_test_buzz_b), !\3227b3938f885317dea9c644434cb82dd !DIBasicType(tag: DW_TAG_string_type)
		call void @test1_test_fizz(double* %test1_test_buzz_a, double* %test1_test_buzz_b), !\39a2b113b63e8232c2945f1018bf785f0 !DIBasicType(tag: DW_TAG_string_type)
		call void @test1_test_fizz2(double* %test1_test_buzz_a, double* %test1_test_buzz_b), !\39a2b113b63e8232c2945f1018bf785f0 !DIBasicType(tag: DW_TAG_string_type)
		call void @test1_test_fizz3(double* %test1_test_buzz_a, double* %test1_test_buzz_b), !\387c9dfef940c096cf145af18149d3600 !DIBasicType(tag: DW_TAG_string_type)
		call void @test1_test_fizz(double* %test1_test_buzz_a, double* %test1_test_buzz_b), !cb7fdc02d16d31723661579b54e31084 !DIBasicType(tag: DW_TAG_string_type)
		call void @test1_test_fizz2(double* %test1_test_buzz_a, double* %test1_test_buzz_b), !cb7fdc02d16d31723661579b54e31084 !DIBasicType(tag: DW_TAG_string_type)
		call void @test1_test_fizz3(double* %test1_test_buzz_a, double* %test1_test_buzz_b), !\370ad8bfb0509411d97738ac929bc3d01 !DIBasicType(tag: DW_TAG_string_type)
		call void @test1_test_fizz(double* %test1_test_buzz_a, double* %test1_test_buzz_b), !\37e3f85f9630519ec31a508b611b1d4bb !DIBasicType(tag: DW_TAG_string_type)
		call void @test1_test_fizz2(double* %test1_test_buzz_a, double* %test1_test_buzz_b), !\37e3f85f9630519ec31a508b611b1d4bb !DIBasicType(tag: DW_TAG_string_type)
		call void @test1_test_fizz3(double* %test1_test_buzz_a, double* %test1_test_buzz_b), !e57d2299cbd9024a113885402ef4e089 !DIBasicType(tag: DW_TAG_string_type)
		call void @test1_test_fizz(double* %test1_test_buzz_a, double* %test1_test_buzz_b), !\37fbb0459ad7da0f1a336cf5de1cf9068 !DIBasicType(tag: DW_TAG_string_type)
		call void @test1_test_fizz2(double* %test1_test_buzz_a, double* %test1_test_buzz_b), !\37fbb0459ad7da0f1a336cf5de1cf9068 !DIBasicType(tag: DW_TAG_string_type)
		call void @test1_test_fizz3(double* %test1_test_buzz_a, double* %test1_test_buzz_b), !\340b3487db7f69f408810fb4cb8b544eb !DIBasicType(tag: DW_TAG_string_type)
		ret void
	}
	
	define void @test1_test_fizz(double* %test1_test_buzz_a, double* %test1_test_buzz_b) {
	block-3:
		%0 = load double, double* %test1_test_buzz_b
		%1 = load double, double* @test1_a
		%2 = load double, double* %test1_test_buzz_a
		%3 = fadd double %1, %2
		%4 = fsub double %0, %3
		store double %4, double* %test1_test_buzz_b
		ret void
	}
	
	define void @test1_test_fizz2(double* %test1_test_buzz_a, double* %test1_test_buzz_b) {
	block-4:
		%0 = load double, double* %test1_test_buzz_b
		%1 = load double, double* %test1_test_buzz_a
		%2 = load double, double* @test1_b
		%3 = fsub double %1, %2
		%4 = fsub double %0, %3
		store double %4, double* %test1_test_buzz_b
		ret void
	}
	
	define void @test1_test_fizz3(double* %test1_test_buzz_a, double* %test1_test_buzz_b) {
	block-5:
		%0 = load double, double* %test1_test_buzz_a
		%1 = load double, double* %test1_test_buzz_b
		%2 = load double, double* @test1_b
		%3 = fadd double %1, %2
		%4 = fsub double %0, %3
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

	expecting := `@test1_a = global double 0x4002666666666666
	@test1_b = global double 2.0
	
	define void @__run() {
	block-6:
		%test1_test_buzz_a = alloca double
		store double 10.0, double* %test1_test_buzz_a
		%test1_test_buzz_b = alloca double
		store double 20.0, double* %test1_test_buzz_b
		call void @test1_test_fizz(double* %test1_test_buzz_a, double* %test1_test_buzz_b), !d44e0a3fc2944aa552d9118f291d3106 !DIBasicType(tag: DW_TAG_string_type)
		ret void
	}
	
	define void @test1_test_fizz(double* %test1_test_buzz_a, double* %test1_test_buzz_b) {
	block-7:
		%0 = load double, double* %test1_test_buzz_a
		%1 = fcmp ogt double %0, 2.0
		br i1 %1, label %block-9-true, label %block-10-false
	
	block-8-after:
		%2 = load double, double* %test1_test_buzz_b
		%3 = fsub double %2, 1.0
		store double %3, double* %test1_test_buzz_b
		ret void
	
	block-9-true:
		%4 = load double, double* %test1_test_buzz_a
		%5 = load double, double* @test1_b
		%6 = fsub double %4, %5
		store double %6, double* %test1_test_buzz_a
		br label %block-8-after
	
	block-10-false:
		store double 10.0, double* %test1_test_buzz_a
		br label %block-8-after
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

	expecting := `@test1_a = global double 0x3DA3CA8CB153A753
	@test1_b = global double 0x3DA3CA8CB153A753
	
	define void @__run() {
	block-11:
		%test1_t_u_x = alloca double
		store double 0x3DA3CA8CB153A753, double* %test1_t_u_x
		call void @test1_t_bar(double* %test1_t_u_x), !\34614d4e08724f278c8ce39e50955edbc !DIBasicType(tag: DW_TAG_string_type)
		call void @test1_t_bar(double* %test1_t_u_x), !ba735eefbef72f20ea6a264b981e9285 !DIBasicType(tag: DW_TAG_string_type)
		call void @test1_t_bar(double* %test1_t_u_x), !fa06e912698cf4825866672ced835870 !DIBasicType(tag: DW_TAG_string_type)
		call void @test1_t_bar(double* %test1_t_u_x), !f3a3858248090df83b9702c4852e0e28 !DIBasicType(tag: DW_TAG_string_type)
		call void @test1_t_bar(double* %test1_t_u_x), !ba20b5c159a59aeb04358b812e68f2d2 !DIBasicType(tag: DW_TAG_string_type)
		ret void
	}
	
	define void @test1_t_bar(double* %test1_t_u_x) {
	block-12:
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
	structs := make(map[string]types.StockFlow)
	c := NewCompiler()
	c.LoadMeta(structs, make(map[string][]float64), []string{})
	s := NewCompiledSpec("test")
	c.currentSpec = s
	c.currentSpecName = "test"
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

	p := ir.NewParam(strings.Join(id[1:], "_"), DoubleP)
	c.resetParaState([]*ir.Param{p})

	if s.vars.GetState(id) != 0 {
		t.Fatalf("var state is incorrect for %s. got=%d", id, s.GetSpecVarState(id))
	}
}

func TestListSpecs(t *testing.T) {
	structs := make(map[string]types.StockFlow)
	c := NewCompiler()
	c.LoadMeta(structs, make(map[string][]float64), []string{})
	s := NewCompiledSpec("test")
	c.currentSpec = s
	c.currentSpecName = "test"
	c.specs["test"] = s

	results := c.ListSpecs()
	if results[1] == "test " {
		t.Fatal("List of Specs failed to return spec test")
	}

}
func TestListSpecsVars(t *testing.T) {
	structs := make(map[string]types.StockFlow)
	c := NewCompiler()
	c.LoadMeta(structs, make(map[string][]float64), []string{})
	s := NewCompiledSpec("test")
	c.currentSpec = s
	c.currentSpecName = "test"
	c.specs["test"] = s

	id := []string{"test", "this", "func"}
	val1 := constant.NewInt(irtypes.I32, 0)
	s.DefineSpecVar(id, val1)

	results := c.ListSpecsAndVars()

	if results["test"] == nil {
		t.Fatal("List of Specs and Vars failed to return spec test")
	}
	if results["test"][0] != "test_this_func" {
		t.Fatalf("List of Specs and Vars doesn't include var test_this_func")
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

	if n4.(*ast.Boolean).Value != true {
		t.Fatalf("boolean has not been negated got=%s", n4.(*ast.Boolean).String())
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

func TestComponent(t *testing.T) {
	tests := &ast.DefStatement{
		Token: ast.Token{
			Type:     ast.TokenType("COMPONENT"),
			Literal:  "COMPONENT",
			Position: []int{0, 0, 0, 0},
		},
		Name: &ast.Identifier{
			Spec:  "test",
			Value: "foo",
		},
		Value: &ast.ComponentLiteral{
			Order: []string{"initial", "alert", "close"},
			Pairs: map[ast.Expression]ast.Expression{
				&ast.Identifier{Spec: "test", Value: "initial"}: &ast.StateLiteral{
					Body: &ast.BlockStatement{
						Statements: []ast.Statement{&ast.ExpressionStatement{Expression: &ast.IfExpression{
							Condition: &ast.InfixExpression{Left: &ast.Identifier{Spec: "test", Value: "x"},
								Operator: ">",
								Right:    &ast.IntegerLiteral{Value: 2}},
							Consequence: &ast.BlockStatement{
								Statements: []ast.Statement{&ast.ExpressionStatement{Expression: &ast.BuiltIn{
									Parameters: map[string]ast.Operand{"toState": &ast.ParameterCall{Value: []string{"this", "alert"}}},
									Function:   "advance",
								}}}},
							Alternative: &ast.BlockStatement{
								Statements: []ast.Statement{&ast.ExpressionStatement{Expression: &ast.BuiltIn{
									Parameters: map[string]ast.Operand{},
									Function:   "stay",
								}}}},
						}}},
					},
				},
				&ast.Identifier{Spec: "test", Value: "alert"}: &ast.StateLiteral{
					Body: &ast.BlockStatement{
						Statements: []ast.Statement{&ast.ExpressionStatement{Expression: &ast.IfExpression{
							Condition: &ast.InfixExpression{Left: &ast.Identifier{Spec: "test", Value: "y"},
								Operator: "==",
								Right:    &ast.IntegerLiteral{Value: 5}},
							Consequence: &ast.BlockStatement{
								Statements: []ast.Statement{&ast.ExpressionStatement{Expression: &ast.BuiltIn{
									Parameters: map[string]ast.Operand{"toState": &ast.ParameterCall{Value: []string{"this", "close"}}},
									Function:   "advance",
								}}}},
							Alternative: &ast.BlockStatement{
								Statements: []ast.Statement{&ast.ExpressionStatement{Expression: &ast.BuiltIn{
									Parameters: map[string]ast.Operand{},
									Function:   "stay",
								}}}},
						}}},
					},
				},
				&ast.Identifier{Spec: "test", Value: "close"}: &ast.StateLiteral{
					Body: &ast.BlockStatement{
						Statements: []ast.Statement{},
					},
				},
			},
		},
	}

	compiler := NewCompiler()
	compiler.currentSpec = NewCompiledSpec("test")
	compiler.currentSpecName = "test"
	compiler.specs[compiler.currentSpecName] = compiler.currentSpec
	id1 := []string{"test", "x"}
	val1 := constant.NewFloat(irtypes.Double, 7.0)
	compiler.currentSpec.DefineSpecVar(id1, val1)
	compiler.allocVariable(id1, val1, []int{0, 0, 0})
	compiler.currentSpec.DefineSpecType(id1, irtypes.Double)
	id2 := []string{"test", "y"}
	val2 := constant.NewFloat(irtypes.Double, 2.0)
	compiler.currentSpec.DefineSpecVar(id2, val2)
	compiler.allocVariable(id2, val2, []int{0, 0, 0})
	compiler.currentSpec.DefineSpecType(id2, irtypes.Double)

	compiler.compileStruct(tests)

	component, ok := compiler.Components["foo"]

	if !ok {
		t.Fatalf("components not found after compiling")
	}

	i, ok2 := component["initial"]

	if !ok2 {
		t.Fatalf("component foo missing state initial")
	}

	a, ok3 := component["alert"]

	if !ok3 {
		t.Fatalf("component foo missing state alert")
	}

	c, ok4 := component["close"]

	if !ok4 {
		t.Fatalf("component foo missing state close")
	}

	for _, k := range []string{a, i, c} {
		if !strings.Contains(compiler.GetIR(), k) {
			t.Fatalf("block %s missing from IR", k)
		}
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

	expecting := `define void @__run() {
		block-28:
			%test_foo_x = alloca double
			store double 8.0, double* %test_foo_x
			ret void
		}
		
		define void @test_foo_initial(double* %test_foo_x) {
		block-29:
			%0 = load double, double* %test_foo_x
			%1 = fcmp ogt double %0, 10.0
			br i1 %1, label %block-31-true, label %block-32-false
		
		block-30-after:
			ret void
		
		block-31-true:
			call void @stay()
			br label %block-30-after
		
		block-32-false:
			%2 = alloca [10 x i8]
			store [10 x i8] c"this.alarm", [10 x i8]* %2
			%3 = bitcast [10 x i8]* %2 to i8*
			call void @advance(i8* %3)
			br label %block-30-after
		}
		
		define void @stay() {
		block-33:
			ret void
		}
		
		define void @advance(i8* %toState) {
		block-34:
			ret void
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

// run block
// init values
// instances (target/source)
// clock
// index
// importing

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
	is := antlr.NewInputStream(test)
	lexer := parser.NewFaultLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	p := parser.NewFaultParser(stream)
	l := listener.NewListener(true, false)
	antlr.ParseTreeWalkerDefault.Walk(l, p.Spec())
	ty := &types.Checker{}
	err := ty.Check(l.AST)
	if err != nil {
		return "", err
	}
	compiler := NewCompiler()
	compiler.LoadMeta(ty.SpecStructs, l.Uncertains, l.Unknowns)
	err = compiler.Compile(l.AST)
	if err != nil {
		return "", err
	}
	//fmt.Println(compiler.GetIR())
	return compiler.GetIR(), err
}

func prepTestSys(test string) (string, error) {
	is := antlr.NewInputStream(test)
	lexer := parser.NewFaultLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	p := parser.NewFaultParser(stream)
	l := listener.NewListener(true, false)
	antlr.ParseTreeWalkerDefault.Walk(l, p.SysSpec())
	ty := &types.Checker{}
	err := ty.Check(l.AST)
	if err != nil {
		return "", err
	}
	compiler := NewCompiler()
	compiler.LoadMeta(ty.SpecStructs, l.Uncertains, l.Unknowns)
	err = compiler.Compile(l.AST)
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
