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
	generator.LoadMeta(compiler.Uncertains, compiler.Unknowns)
	generator.rawAsserts = compiler.Asserts
	generator.rawAssumes = compiler.Assumes
	//fmt.Println(compiler.GetIR())
	generator.Run(compiler.GetIR())
	return generator.SMT(), nil
}
