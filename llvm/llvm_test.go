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
		call void @test1_test_fizz(double* %test1_test_buzz_a, double* %test1_test_buzz_b), !\383be9367ad466c3482a1d5aacdd11ef2 !DIBasicType(tag: DW_TAG_string_type)
		call void @test1_test_fizz2(double* %test1_test_buzz_a, double* %test1_test_buzz_b), !\383be9367ad466c3482a1d5aacdd11ef2 !DIBasicType(tag: DW_TAG_string_type)
		call void @test1_test_fizz3(double* %test1_test_buzz_a, double* %test1_test_buzz_b), !\38487cfbf67279c764a526213b15d255a !DIBasicType(tag: DW_TAG_string_type)
		call void @test1_test_fizz(double* %test1_test_buzz_a, double* %test1_test_buzz_b), !\36c8807a6f069bb7d51cc92b4c3dbc725 !DIBasicType(tag: DW_TAG_string_type)
		call void @test1_test_fizz2(double* %test1_test_buzz_a, double* %test1_test_buzz_b), !\36c8807a6f069bb7d51cc92b4c3dbc725 !DIBasicType(tag: DW_TAG_string_type)
		call void @test1_test_fizz3(double* %test1_test_buzz_a, double* %test1_test_buzz_b), !c9992c8be1f85a2932d63eba0e893a05 !DIBasicType(tag: DW_TAG_string_type)
		call void @test1_test_fizz(double* %test1_test_buzz_a, double* %test1_test_buzz_b), !d00e60656da4edd14ac5050f5cd9c890 !DIBasicType(tag: DW_TAG_string_type)
		call void @test1_test_fizz2(double* %test1_test_buzz_a, double* %test1_test_buzz_b), !d00e60656da4edd14ac5050f5cd9c890 !DIBasicType(tag: DW_TAG_string_type)
		call void @test1_test_fizz3(double* %test1_test_buzz_a, double* %test1_test_buzz_b), !\33240cd436a5c85a09944540236db4ad2 !DIBasicType(tag: DW_TAG_string_type)
		call void @test1_test_fizz(double* %test1_test_buzz_a, double* %test1_test_buzz_b), !\329996c047463487ecba33b1e25082ffe !DIBasicType(tag: DW_TAG_string_type)
		call void @test1_test_fizz2(double* %test1_test_buzz_a, double* %test1_test_buzz_b), !\329996c047463487ecba33b1e25082ffe !DIBasicType(tag: DW_TAG_string_type)
		call void @test1_test_fizz3(double* %test1_test_buzz_a, double* %test1_test_buzz_b), !\33ac50f9e8b06ceef11bd5da91ee43fcc !DIBasicType(tag: DW_TAG_string_type)
		call void @test1_test_fizz(double* %test1_test_buzz_a, double* %test1_test_buzz_b), !\3669091a688eb9e5a01453bb4aaf35abb !DIBasicType(tag: DW_TAG_string_type)
		call void @test1_test_fizz2(double* %test1_test_buzz_a, double* %test1_test_buzz_b), !\3669091a688eb9e5a01453bb4aaf35abb !DIBasicType(tag: DW_TAG_string_type)
		call void @test1_test_fizz3(double* %test1_test_buzz_a, double* %test1_test_buzz_b), !\31496d8c243d2229e0186bb2919665d37 !DIBasicType(tag: DW_TAG_string_type)
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
		call void @test1_test_fizz(double* %test1_test_buzz_a, double* %test1_test_buzz_b), !d13399dec07511570da1e17fb6f98374 !DIBasicType(tag: DW_TAG_string_type)
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
		call void @test1_t_bar(double* %test1_t_u_x), !ef486bcefdf32416542f0c0e7dafd2a7 !DIBasicType(tag: DW_TAG_string_type)
		call void @test1_t_bar(double* %test1_t_u_x), !\36b15db79f1be6a7c6e58e1703ad79489 !DIBasicType(tag: DW_TAG_string_type)
		call void @test1_t_bar(double* %test1_t_u_x), !e4b288b91e48b4956fd9ecd3cf0a7950 !DIBasicType(tag: DW_TAG_string_type)
		call void @test1_t_bar(double* %test1_t_u_x), !\38a55f57c625b7c5382fc6522c9ede93a !DIBasicType(tag: DW_TAG_string_type)
		call void @test1_t_bar(double* %test1_t_u_x), !\37242f7face5577b201458fee71ac6cd6 !DIBasicType(tag: DW_TAG_string_type)
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

func TestGetInstances(t *testing.T) {
	structs := make(map[string]types.StockFlow)
	c := NewCompiler()
	c.LoadMeta(structs, make(map[string][]float64), []string{})
	s := NewCompiledSpec("test")
	c.currentSpec = s
	c.currentSpecName = "test"
	c.specs["test"] = s
	c.instances["insta1"] = "fake1"
	c.instances["insta2"] = "fake2"
	c.instances["insta3"] = "fake1"

	infix := &ast.InfixExpression{
		Left: &ast.IndexExpression{Left: &ast.ParameterCall{Value: []string{"fake1", "prop1"}}},
		Right: &ast.InfixExpression{
			Left:  &ast.IndexExpression{Left: &ast.ParameterCall{Value: []string{"fake1", "prop3"}}},
			Right: &ast.PrefixExpression{Right: &ast.ParameterCall{Value: []string{"fake2", "prop2"}}}}}
	results := c.getInstances(infix)

	if len(results["fake1"]) != 2 {
		t.Fatalf("incorrect results returned. got=%d want=2", len(results["fake1"]))
	}
	if results["fake1"][0] != "insta1" && results["fake1"][0] != "insta3" {
		t.Fatalf("instance not correct. got=%s want=insta1", results["fake1"][0])
	}

	if results["fake1"][1] != "insta3" && results["fake1"][1] != "insta1" {
		t.Fatalf("instance not correct. got=%s want=insta3", results["fake1"][1])
	}

	if len(results["fake2"]) != 1 {
		t.Fatalf("incorrect results returned. got=%d want=1", len(results["fake2"]))
	}

	if results["fake2"][0] != "insta2" {
		t.Fatalf("instance not correct. got=%s want=insta2", results["fake2"][0])
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
