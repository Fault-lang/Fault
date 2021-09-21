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

// func TestBathTub(t *testing.T) {
// 	expecting, err := os.ReadFile("testdata/bathtub.smt2")
// 	if err != nil {
// 		panic("compiled spec bathtub is not valid")
// 	}

// 	data, err := os.ReadFile("testdata/bathtub.fspec")
// 	if err != nil {
// 		panic("spec bathtub is not valid")
// 	}

// 	smt, err := prepTest(string(data))

// 	if err != nil {
// 		t.Fatalf("compilation failed on valid spec. got=%s", err)
// 	}

// 	err = compareResults(smt, string(expecting))

// 	if err != nil {
// 		t.Fatalf(err.Error())
// 	}
// }

func TestBathTub2(t *testing.T) {
	expecting, err := os.ReadFile("testdata/bathtub.smt2")
	if err != nil {
		panic("compiled spec bathtub is not valid")
	}

	data, err := os.ReadFile("testdata/bathtub2.fspec")
	if err != nil {
		panic("spec bathtub is not valid")
	}

	smt, err := prepTest(string(data))

	if err != nil {
		t.Fatalf("compilation failed on valid spec. got=%s", err)
	}

	err = compareResults(smt, string(expecting))

	if err != nil {
		t.Fatalf(err.Error())
	}
}

func compareResults(smt string, expecting string) error {
	//if !strings.Contains(smt, "(declare-fun") {
	if !strings.Contains(smt, "failfailfail") {
		return fmt.Errorf("smt not valid. \ngot=%s", smt)
	}

	smt = stripAndEscape(smt)
	expecting = stripAndEscape(expecting)
	if len(smt) != len(expecting) {
		return fmt.Errorf("wrong instructions length.\nwant=%s\ngot=%s",
			expecting, smt)
	}

	if smt != expecting {
		return fmt.Errorf("SMT string does not match.\nwant=%q\ngot=%q",
			expecting, smt)
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
	compiler := llvm.NewCompiler(ty.SpecStructs)
	err = compiler.Compile(l.AST)
	if err != nil {
		return "", err
	}
	generator := NewGenerator()
	//fmt.Println(compiler.GetIR())
	generator.Run(compiler.GetIR())
	return generator.SMT(), nil
}
