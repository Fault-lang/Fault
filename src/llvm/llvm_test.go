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

define void @test1_test_fizz(double* %test1_test_buzz_b, double* %test1_test_buzz_a) {
block-3:
	%0 = load double, double* %test1_test_buzz_b
	%1 = load double, double* @test1_a
	%2 = load double, double* %test1_test_buzz_a
	%3 = fadd double %1, 10.0
	%4 = fsub double 20.0, %3
	store double %4, double* %test1_test_buzz_b
	ret void
}

define void @test1_test_fizz2(double* %test1_test_buzz_a, double* %test1_test_buzz_b) {
block-4:
	%0 = load double, double* %test1_test_buzz_b
	%1 = load double, double* %test1_test_buzz_a
	%2 = load double, double* @test1_b
	%3 = fsub double 10.0, %2
	%4 = fsub double 20.0, %3
	store double %4, double* %test1_test_buzz_b
	ret void
}

define void @test1_test_fizz3(double* %test1_test_buzz_b, double* %test1_test_buzz_a) {
block-5:
	%0 = load double, double* %test1_test_buzz_a
	%1 = load double, double* %test1_test_buzz_b
	%2 = load double, double* @test1_b
	%3 = fadd double 20.0, %2
	%4 = fsub double 10.0, %3
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
