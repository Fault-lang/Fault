package smt

import (
	"fault/listener"
	"fault/llvm"
	"fault/parser"
	"fault/preprocess"
	"fault/types"
	"fault/util"
	"fmt"
	"os"
	gopath "path"
	"strconv"
	"strings"
	"testing"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

func TestTestData(t *testing.T) {
	specs := []string{
		//"testdata/bathtub.fspec",
		//"testdata/simple.fspec",
		//"testdata/bathtub2.fspec",
		//"testdata/booleans.fspec",
		"testdata/unknowns.fspec",
	}
	smt2s := []string{
		// "testdata/bathtub.smt2",
		// "testdata/simple.smt2",
		// "testdata/bathtub2.smt2",
		// "testdata/booleans.smt2",
		"testdata/unknowns.smt2",
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
		smt, err := prepTest(s, string(data))

		if err != nil {
			t.Fatalf("compilation failed on valid spec %s. got=%s", s, err)
		}

		err = compareResults(s, smt, string(expecting))

		if err != nil {
			t.Fatalf(err.Error())
		}
	}
}

func TestSys(t *testing.T) {
	specs := [][]string{
		{"testdata/statecharts/statechart.fsystem", "1"},
		{"testdata/statecharts/advanceor.fsystem", "0"},
		{"testdata/statecharts/multioradvance.fsystem", "0"},
		{"testdata/statecharts/advanceand.fsystem", "0"},
		{"testdata/statecharts/mixedcalls.fsystem", "1"},
		{"testdata/statecharts/triggerfunc.fsystem", "1"},
	}
	smt2s := []string{
		"testdata/statecharts/statechart.smt2",
		"testdata/statecharts/advanceor.smt2",
		"testdata/statecharts/multioradvance.smt2",
		"testdata/statecharts/advanceand.smt2",
		"testdata/statecharts/mixedcalls.smt2",
		"testdata/statecharts/triggerfunc.smt2",
	}
	for i, s := range specs {
		data, err := os.ReadFile(s[0])
		if err != nil {
			panic(fmt.Sprintf("spec %s is not valid", s[0]))
		}
		imports, _ := strconv.ParseBool(s[1])

		expecting, err := os.ReadFile(smt2s[i])
		if err != nil {
			panic(fmt.Sprintf("compiled spec %s is not valid", smt2s[i]))
		}
		smt, err := prepTestSys(s[0], string(data), imports)

		if err != nil {
			t.Fatalf("compilation failed on valid spec %s. got=%s", s[0], err)
		}

		err = compareResults(s[0], smt, string(expecting))

		if err != nil {
			t.Fatalf(err.Error())
		}
	}
}

func TestMultiCond(t *testing.T) {
	specs := []string{
		"testdata/conditionals/multicond.fspec",
		"testdata/conditionals/multicond2.fspec",
		"testdata/conditionals/multicond3.fspec",
		"testdata/conditionals/multicond4.fspec",
		"testdata/conditionals/multicond5.fspec",
		"testdata/conditionals/condwelse.fspec",
	}
	smt2s := []string{
		"testdata/conditionals/multicond.smt2",
		"testdata/conditionals/multicond2.smt2",
		"testdata/conditionals/multicond3.smt2",
		"testdata/conditionals/multicond4.smt2",
		"testdata/conditionals/multicond5.smt2",
		"testdata/conditionals/condwelse.smt2",
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
		smt, err := prepTest(s, string(data))

		if err != nil {
			t.Fatalf("compilation failed on valid spec %s. got=%s", s, err)
		}

		err = compareResults(s, smt, string(expecting))

		if err != nil {
			t.Fatalf(err.Error())
		}
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
		if !notStrictlyOrdered(expecting, smt) {
			return fmt.Errorf("SMT string does not match for spec %s.\nwant=%q\ngot=%q",
				s, expecting, smt)
		}
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

func prepTest(path string, test string) (string, error) {
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
	compiler := llvm.NewCompiler()
	compiler.LoadMeta(ty.SpecStructs, l.Uncertains, l.Unknowns)
	err = compiler.Compile(tree)
	if err != nil {
		return "", err
	}
	//fmt.Println(compiler.GetIR())
	generator := NewGenerator()
	generator.LoadMeta(compiler.RunRound, compiler.Uncertains, compiler.Unknowns, compiler.Asserts, compiler.Assumes)
	generator.Run(compiler.GetIR())
	//fmt.Println(generator.SMT())
	return generator.SMT(), nil
}

func prepTestSys(filepath string, test string, imports bool) (string, error) {
	filepath = util.Filepath(filepath)
	path := gopath.Dir(filepath)

	is := antlr.NewInputStream(test)
	lexer := parser.NewFaultLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	p := parser.NewFaultParser(stream)
	l := listener.NewListener(path, !imports, false) //imports being true means testing is false :)
	antlr.ParseTreeWalkerDefault.Walk(l, p.SysSpec())

	pre := preprocess.NewProcesser()
	pre.StructsPropertyOrder = l.StructsPropertyOrder
	tree := pre.Run(l.AST)

	ty := types.NewTypeChecker(pre.Specs)
	tree, err := ty.Check(tree)

	if err != nil {
		return "", err
	}
	compiler := llvm.NewCompiler()
	compiler.LoadMeta(ty.SpecStructs, l.Uncertains, l.Unknowns)
	err = compiler.Compile(tree)
	if err != nil {
		return "", err
	}
	//fmt.Println(compiler.GetIR())
	generator := NewGenerator()
	generator.LoadMeta(compiler.RunRound, compiler.Uncertains, compiler.Unknowns, compiler.Asserts, compiler.Assumes)
	generator.Run(compiler.GetIR())
	fmt.Println(generator.SMT())
	return generator.SMT(), nil
}

func notStrictlyOrdered(want string, got string) bool {
	// Fixing cases where lines of SMT end up in slightly
	// different orders. Only runs when shallow string
	// compare fails

	s := strings.Split(want, "")
	dedup := make(map[string]bool)
	var keys []string
	for _, v := range s {
		if _, ok := dedup[v]; !ok {
			dedup[v] = true
			keys = append(keys, v)
		}
	}

	for _, k := range keys {
		if strings.Count(want, k) != strings.Count(got, k) {
			return false
		}
	}
	return true
}
