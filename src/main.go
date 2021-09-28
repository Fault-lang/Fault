package main

import (
	"fault/listener"
	"fault/llvm"
	"fault/parser"
	"fault/smt"
	"fault/types"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func parse(data string) (*listener.FaultListener, *types.Checker) {
	// Setup the input
	is := antlr.NewInputStream(data)

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
		fmt.Println(err)
	}
	return listener, ty
}

func ll(listener *listener.FaultListener, ty *types.Checker) *llvm.Compiler {
	compiler := llvm.NewCompiler(ty.SpecStructs)
	err := compiler.Compile(listener.AST)
	if err != nil {
		fmt.Println(err)
	}
	return compiler
}

func smt2(ir string) *smt.Generator {
	generator := smt.NewGenerator()
	generator.Run(ir)
	return generator
}

func run(filepath string, mode string, input string) {
	data, err := os.ReadFile(filepath)
	d := string(data)
	if err != nil {
		panic(err)
	}

	switch input {
	case "fspec":
		listener, ty := parse(d)
		if mode == "ast" {
			fmt.Println(listener.AST)
			return
		}

		compiler := ll(listener, ty)

		if mode == "ir" {
			fmt.Println(compiler.GetIR())
			return
		}

		generator := smt2(compiler.GetIR())
		fmt.Println(generator.SMT())
	case "ll":
		generator := smt2(d)
		fmt.Println(generator.SMT())

	case "smt2":
	}

	// if mode == "SMT" {
	// 	fmt.Println(generator.SMT())
	// 	return
	// }
}

func main() {
	var mode string
	var input string
	var filepath string
	modeCommand := flag.String("mode", "smt", "stop compiler at certain milestones: ast, ir, or smt")
	inputCommand := flag.String("input", "fspec", "format of the input file (default: fspec)")
	fpCommand := flag.String("filepath", "", "path to file to compile")
	//helpCommand := flag.Bool("help", false, "path to file to compile")

	flag.Parse()

	if *fpCommand == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}
	filepath = *fpCommand

	if *modeCommand == "" {
		mode = "smt"
	} else {
		mode = strings.ToLower(*modeCommand)
		switch mode {
		case "ast":
		case "ir":
		case "smt":
		default:
			fmt.Printf("%s is not a valid mode", mode)
			os.Exit(1)
		}
	}

	if *inputCommand == "" {
		input = "fspec"
	} else {
		input = strings.ToLower(*inputCommand)
		switch input {
		case "fspec":
		case "ll":
		case "smt2":
		default:
			fmt.Printf("%s is not a valid input format", input)
			os.Exit(1)
		}
	}

	run(filepath, mode, input)
}
