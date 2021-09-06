package smt

import (
	"fault/listener"
	"fault/llvm"
	"fault/parser"
	"fault/types"
	"testing"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func TestBathTub(t *testing.T) {
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
	generator := NewGenerator(compiler.GetIR())
	return generator.SMT(), nil
}
