package main

import (
	"fault/listener"
	"fault/parser"
	"fault/types"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func main() {
	// Setup the input
	is := antlr.NewInputStream("spec test;")

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
	err := ty.Check(listener.AST)
	if err != nil {
		print(err)
	}
	print(listener.AST.String())

}
