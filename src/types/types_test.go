package types

import (
	"fault/listener"
	"fault/parser"
	"testing"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func TestAddOK(t *testing.T) {
	test := `spec test1;
			const x = 2+2;
	`
	checker, err := prepTest(test)

	spec := checker.AST
	sym := checker.Symbols

	if err != nil {
		t.Fatalf("Type checking failed on valid expression. got=%s", err)
	}

	if sym["x"] != "INT" {
		t.Fatalf("Constant x does not have an int type. got=%s", sym["x"])
	}

	if spec == nil {
		t.Fatalf("prepTest() returned nil")
	}

}

func prepTest(test string) (*Checker, error) {
	is := antlr.NewInputStream(test)
	lexer := parser.NewFaultLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	p := parser.NewFaultParser(stream)
	l := &listener.FaultListener{}
	antlr.ParseTreeWalkerDefault.Walk(l, p.Spec())
	ty := &Checker{}
	err := ty.Check(l.AST)
	return ty, err
}
