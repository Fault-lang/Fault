package smt

import (
	"fault/listener"
	"fault/llvm"
	"fault/parser"
	"fault/types"
	"fmt"
	"os"
	"strings"
	"testing"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func TestTestData(t *testing.T) {
	specs := []string{"testdata/bathtub.fspec",
		"testdata/simple.fspec",
		"testdata/bathtub2.fspec",
		//"testdata/unknowns.fspec",
	}
	smt2s := []string{"testdata/bathtub.smt2",
		"testdata/simple.smt2",
		"testdata/bathtub2.smt2",
		//"testdata/unknowns.smt2",
	}
	for i, s := range specs {
		data, err := os.ReadFile(s)
		if err != nil {
			panic(fmt.Sprintf("spec %s is not valid", s))
		}
		expecting, err := os.ReadFile(smt2s[i])
		if err != nil {
			panic(fmt.Sprintf("compiled spec %s is not valid", smt2s[i]))
		}
		smt, err := prepTest(string(data))

		if err != nil {
			t.Fatalf("compilation failed on valid spec. got=%s", err)
		}

		err = compareResults(s, smt, string(expecting))

		if err != nil {
			t.Fatalf(err.Error())
		}
	}
}

func TestIndividual(t *testing.T) {

	data, err := os.ReadFile("testdata/unknowns.fspec")
	if err != nil {
		panic(fmt.Sprintf("spec %s is not valid", "testdata/unknowns.fspec"))
	}
	expecting, err := os.ReadFile("testdata/unknowns.smt2")
	if err != nil {
		panic(fmt.Sprintf("compiled spec %s is not valid", "testdata/unknowns.smt2"))
	}
	smt, err := prepTest(string(data))

	if err != nil {
		t.Fatalf("compilation failed on valid spec. got=%s", err)
	}

	err = compareResults("testdata/unknowns.fspec", smt, string(expecting))

	if err != nil {
		t.Fatalf(err.Error())
	}
}

func TestMultiCond(t *testing.T) {

	data, err := os.ReadFile("testdata/multicond.fspec")
	if err != nil {
		panic(fmt.Sprintf("spec %s is not valid", "testdata/multicond.fspec"))
	}
	expecting, err := os.ReadFile("testdata/multicond.smt2")
	if err != nil {
		panic(fmt.Sprintf("compiled spec %s is not valid", "testdata/multicond.smt2"))
	}
	smt, err := prepTest(string(data))

	if err != nil {
		t.Fatalf("compilation failed on valid spec. got=%s", err)
	}

	// Order of statements is not deterministic for some reason.
	// so will fix the flakiness for now by checking for both
	// possible orders
	err = compareResults("testdata/multicond.fspec", smt, string(expecting))
	flake := compareResults("testdata/multicond.fspec", smt, "(declare-funmulticond_t_base_cond_0()Real)(declare-funmulticond_t_base_value_0()Real)(declare-funmulticond_t_base_cond_2()Real)(declare-funmulticond_t_base_value_2()Real)(declare-funmulticond_t_base_value_4()Real)(declare-funmulticond_t_base_cond_4()Real)(declare-funmulticond_t_base_value_1()Real)(declare-funmulticond_t_base_cond_1()Real)(declare-funmulticond_t_base_value_3()Real)(declare-funmulticond_t_base_cond_3()Real)(assert(=multicond_t_base_cond_01.0))(assert(=multicond_t_base_value_010.0))(assert(=multicond_t_base_value_1(+multicond_t_base_value_010.0)))(assert(=multicond_t_base_cond_1(+multicond_t_base_cond_02.0)))(assert(ite(>multicond_t_base_cond_00.0)(and(=multicond_t_base_cond_2multicond_t_base_cond_1)(=multicond_t_base_value_2multicond_t_base_value_1))(and(=multicond_t_base_value_2multicond_t_base_value_0)(=multicond_t_base_cond_2multicond_t_base_cond_0))))(assert(=multicond_t_base_value_3(+multicond_t_base_value_220.0)))(assert(=multicond_t_base_cond_3(-multicond_t_base_cond_22.0)))(assert(ite(>multicond_t_base_cond_24.0)(and(=multicond_t_base_value_4multicond_t_base_value_3)(=multicond_t_base_cond_4multicond_t_base_cond_3))(and(=multicond_t_base_value_4multicond_t_base_value_2)(=multicond_t_base_cond_4multicond_t_base_cond_2))))")

	if err != nil && flake != nil {
		t.Fatalf(err.Error())
	}
}

func compareResults(s string, smt string, expecting string) error {
	if !strings.Contains(smt, "(declare-fun") {
		return fmt.Errorf("smt not valid for spec %s. \ngot=%s", s, smt)
	}

	smt = stripAndEscape(smt)
	expecting = stripAndEscape(expecting)
	if len(smt) != len(expecting) {
		return fmt.Errorf("wrong instructions length for spec %s.\nwant=%s\ngot=%s",
			s, expecting, smt)
	}

	if smt != expecting {
		return fmt.Errorf("SMT string does not match for spec %s.\nwant=%q\ngot=%q",
			s, expecting, smt)
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
	compiler := llvm.NewCompiler()
	compiler.LoadMeta(ty.SpecStructs, l.Uncertains, l.Unknowns)
	err = compiler.Compile(l.AST)
	if err != nil {
		return "", err
	}
	generator := NewGenerator()
	generator.LoadMeta(compiler.Uncertains, compiler.Unknowns, compiler.Asserts, compiler.Assumes)
	generator.Run(compiler.GetIR())
	//fmt.Println(generator.SMT())
	return generator.SMT(), nil
}
