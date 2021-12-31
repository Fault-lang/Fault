package main

import (
	"fault/execute"
	"fault/listener"
	"fault/llvm"
	"fault/parser"
	"fault/smt"
	"fault/types"
	"fault/util"
	"flag"
	"fmt"
	"log"
	"os"
	gopath "path"
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	_ "github.com/olekukonko/tablewriter"
)

func parse(data string, path string) (*listener.FaultListener, *types.Checker) {
	// Setup the input
	is := antlr.NewInputStream(data)

	// Create the Lexer
	lexer := parser.NewFaultLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser
	p := parser.NewFaultParser(stream)

	// Finally parse the expression
	listener := &listener.FaultListener{}
	listener.Path = path
	antlr.ParseTreeWalkerDefault.Walk(listener, p.Spec())

	// Infer Types and Build Symbol Table
	ty := &types.Checker{}
	err := ty.Check(listener.AST)
	if err != nil {
		log.Fatal(err)
	}
	return listener, ty
}

func ll(listener *listener.FaultListener, ty *types.Checker) *llvm.Compiler {
	compiler := llvm.NewCompiler(ty.SpecStructs)
	err := compiler.Compile(listener.AST)
	if err != nil {
		log.Fatal(err)
	}
	return compiler
}

func smt2(ir string) *smt.Generator {
	generator := smt.NewGenerator()
	generator.Run(ir)
	return generator
}

func probability(smt string, uncertains map[string][]float64) (*execute.ModelChecker, map[string]execute.Scenario) {
	ex := execute.NewModelChecker("z3")
	ex.LoadModel(smt, uncertains)
	ok, err := ex.Check()
	if err != nil {
		log.Fatal(err)
	}
	if !ok {
		fmt.Print("Fault could not find a failure case.")
		os.Exit(0)
	}
	scenario, err := ex.Solve()
	if err != nil {
		log.Fatal(err)
	}
	data := ex.Filter(scenario)
	return ex, data
}

func run(filepath string, mode string, input string) {
	filepath = util.Filepath(filepath)

	data, err := os.ReadFile(filepath)
	if err != nil {
		log.Fatal(err)
	}
	d := string(data)
	path := gopath.Dir(filepath)

	switch input {
	case "fspec":
		listener, ty := parse(d, path)
		if listener == nil {
			log.Fatal("Fault parser returned nil")
		}

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

		if mode == "smt" {
			fmt.Println(generator.SMT())
			return
		}

		mc, data := probability(generator.SMT(), make(map[string][]float64))
		mc.LoadMeta(generator.Branches, generator.BranchTrail)
		fmt.Println("~~~~~~~~~~\n  Fault found the following scenario\n~~~~~~~~~~")
		mc.Format(data)
	case "ll":
		generator := smt2(d)
		if mode == "smt" {
			fmt.Println(generator.SMT())
			return
		}

		mc, data := probability(generator.SMT(), make(map[string][]float64))
		mc.LoadMeta(generator.Branches, generator.BranchTrail)
		fmt.Println("~~~~~~~~~~\n  Fault found the following scenario\n~~~~~~~~~~")
		mc.Format(data)
	case "smt2":
		mc, data := probability(d, make(map[string][]float64))

		fmt.Println("~~~~~~~~~~\n  Fault found the following scenario\n~~~~~~~~~~")
		mc.Format(data)
	}
}

func main() {
	var mode string
	var input string
	var filepath string
	modeCommand := flag.String("mode", "check", "stop compiler at certain milestones: ast, ir, smt, or check")
	inputCommand := flag.String("input", "fspec", "format of the input file (default: fspec)")
	fpCommand := flag.String("filepath", "", "path to file to compile")
	//helpCommand := flag.Bool("help", false, "path to file to compile")

	flag.Parse()

	if *fpCommand == "" {
		flag.PrintDefaults()
		fmt.Printf("must provide path of file to compile")
		os.Exit(1)
	}
	filepath = *fpCommand

	if *modeCommand == "" {
		mode = "check"
	} else {
		mode = strings.ToLower(*modeCommand)
		switch mode {
		case "ast":
		case "ir":
		case "smt":
		case "check":
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
