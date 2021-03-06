package fuzzer

import (
	"fault/listener"
	"fault/parser"
	"fault/types"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func evaluate(spec string) (code int, err error) {
	defer func() {
		if out := recover(); out != nil {
			code = 1
			err = out.(error)
		}
	}()
	is := antlr.NewInputStream(spec)

	// Create the Lexer
	lexer := parser.NewFaultLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser
	p := parser.NewFaultParser(stream)

	// Finally parse the expression
	listener := &listener.FaultListener{}
	antlr.ParseTreeWalkerDefault.Walk(listener, p.Spec())

	// Infer Types and Build Symbol Table
	ty := &types.Checker{}
	err = ty.Check(listener.AST)
	if err != nil {
		code = 2
	}
	return
}
