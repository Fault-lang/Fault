package llvm

import (
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

	expecting := `
	@test1_a = global double 0x4002666666666666
@test1_b = global double 2.0

define void @__run() {
block-2:
	%test1_test_buzz_a1 = alloca double
	store double 10.0, double* %test1_test_buzz_a1
	%test1_test_buzz_b1 = alloca double
	store double 20.0, double* %test1_test_buzz_b1
	%0 = load double, double* @test1_a
	%test1_test_buzz_a2 = alloca double
	store double 10.0, double* %test1_test_buzz_a2
	%1 = fadd double %0, 10.0
	%test1_test_buzz_b2 = alloca double
	store double %1, double* %test1_test_buzz_b2
	%test1_test_buzz_a3 = alloca double
	store double 10.0, double* %test1_test_buzz_a3
	%2 = load double, double* @test1_b
	%3 = fsub double 10.0, %2
	%test1_test_buzz_b3 = alloca double
	store double %3, double* %test1_test_buzz_b3
	%test1_test_buzz_b4 = alloca double
	store double 20.0, double* %test1_test_buzz_b4
	%4 = load double, double* @test1_b
	%5 = fadd double 20.0, %4
	%test1_test_buzz_a4 = alloca double
	store double %5, double* %test1_test_buzz_a4
	%6 = load double, double* @test1_a
	%test1_test_buzz_a5 = alloca double
	store double 10.0, double* %test1_test_buzz_a5
	%7 = fadd double %6, 10.0
	%test1_test_buzz_b5 = alloca double
	store double %7, double* %test1_test_buzz_b5
	%test1_test_buzz_a6 = alloca double
	store double 10.0, double* %test1_test_buzz_a6
	%8 = load double, double* @test1_b
	%9 = fsub double 10.0, %8
	%test1_test_buzz_b6 = alloca double
	store double %9, double* %test1_test_buzz_b6
	%test1_test_buzz_b7 = alloca double
	store double 20.0, double* %test1_test_buzz_b7
	%10 = load double, double* @test1_b
	%11 = fadd double 20.0, %10
	%test1_test_buzz_a7 = alloca double
	store double %11, double* %test1_test_buzz_a7
	%12 = load double, double* @test1_a
	%test1_test_buzz_a8 = alloca double
	store double 10.0, double* %test1_test_buzz_a8
	%13 = fadd double %12, 10.0
	%test1_test_buzz_b8 = alloca double
	store double %13, double* %test1_test_buzz_b8
	%test1_test_buzz_a9 = alloca double
	store double 10.0, double* %test1_test_buzz_a9
	%14 = load double, double* @test1_b
	%15 = fsub double 10.0, %14
	%test1_test_buzz_b9 = alloca double
	store double %15, double* %test1_test_buzz_b9
	%test1_test_buzz_b10 = alloca double
	store double 20.0, double* %test1_test_buzz_b10
	%16 = load double, double* @test1_b
	%17 = fadd double 20.0, %16
	%test1_test_buzz_a10 = alloca double
	store double %17, double* %test1_test_buzz_a10
	%18 = load double, double* @test1_a
	%test1_test_buzz_a11 = alloca double
	store double 10.0, double* %test1_test_buzz_a11
	%19 = fadd double %18, 10.0
	%test1_test_buzz_b11 = alloca double
	store double %19, double* %test1_test_buzz_b11
	%test1_test_buzz_a12 = alloca double
	store double 10.0, double* %test1_test_buzz_a12
	%20 = load double, double* @test1_b
	%21 = fsub double 10.0, %20
	%test1_test_buzz_b12 = alloca double
	store double %21, double* %test1_test_buzz_b12
	%test1_test_buzz_b13 = alloca double
	store double 20.0, double* %test1_test_buzz_b13
	%22 = load double, double* @test1_b
	%23 = fadd double 20.0, %22
	%test1_test_buzz_a13 = alloca double
	store double %23, double* %test1_test_buzz_a13
	%24 = load double, double* @test1_a
	%test1_test_buzz_a14 = alloca double
	store double 10.0, double* %test1_test_buzz_a14
	%25 = fadd double %24, 10.0
	%test1_test_buzz_b14 = alloca double
	store double %25, double* %test1_test_buzz_b14
	%test1_test_buzz_a15 = alloca double
	store double 10.0, double* %test1_test_buzz_a15
	%26 = load double, double* @test1_b
	%27 = fsub double 10.0, %26
	%test1_test_buzz_b15 = alloca double
	store double %27, double* %test1_test_buzz_b15
	%test1_test_buzz_b16 = alloca double
	store double 20.0, double* %test1_test_buzz_b16
	%28 = load double, double* @test1_b
	%29 = fadd double 20.0, %28
	%test1_test_buzz_a16 = alloca double
	store double %29, double* %test1_test_buzz_a16
	%30 = load double, double* @test1_a
	%test1_test_buzz_a17 = alloca double
	store double 10.0, double* %test1_test_buzz_a17
	%31 = fadd double %30, 10.0
	%test1_test_buzz_b17 = alloca double
	store double %31, double* %test1_test_buzz_b17
	%test1_test_buzz_a18 = alloca double
	store double 10.0, double* %test1_test_buzz_a18
	%32 = load double, double* @test1_b
	%33 = fsub double 10.0, %32
	%test1_test_buzz_b18 = alloca double
	store double %33, double* %test1_test_buzz_b18
	%test1_test_buzz_b19 = alloca double
	store double 20.0, double* %test1_test_buzz_b19
	%34 = load double, double* @test1_b
	%35 = fadd double 20.0, %34
	%test1_test_buzz_a19 = alloca double
	store double %35, double* %test1_test_buzz_a19
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

// run block
// init values
// conditionals
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
	l := &listener.FaultListener{}
	antlr.ParseTreeWalkerDefault.Walk(l, p.Spec())
	ty := &types.Checker{}
	err := ty.Check(l.AST)
	if err != nil {
		return "", err
	}
	compiler := NewCompiler(ty.SpecStructs)
	err = compiler.Compile(l.AST)
	if err != nil {
		return "", err
	}
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
