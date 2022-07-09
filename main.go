package main

import (
	"fault/ast"
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

func parse(data string, path string, file string) (*listener.FaultListener, *types.Checker) {
	// Setup the input
	is := antlr.NewInputStream(data)

	// Create the Lexer
	lexer := parser.NewFaultLexer(is)
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(&listener.FaultErrorListener{Filename: file})
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser
	p := parser.NewFaultParser(stream)
	p.RemoveErrorListeners()
	p.AddErrorListener(&listener.FaultErrorListener{Filename: file})

	// Finally parse the expression
	lstnr := listener.NewListener(false, false)
	lstnr.Path = path
	antlr.ParseTreeWalkerDefault.Walk(lstnr, p.Spec())

	// Infer Types and Build Symbol Table
	ty := &types.Checker{}
	err := ty.Check(lstnr.AST)
	if err != nil {
		log.Fatal(err)
	}
	return lstnr, ty
}

func ll(lstnr *listener.FaultListener, ty *types.Checker) *llvm.Compiler {
	compiler := llvm.NewCompiler()
	compiler.LoadMeta(ty.SpecStructs, lstnr.Uncertains, lstnr.Unknowns)
	err := compiler.Compile(lstnr.AST)
	if err != nil {
		log.Fatal(err)
	}
	return compiler
}

func smt2(ir string, uncertains map[string][]float64, unknowns []string, asserts []*ast.AssertionStatement, assumes []*ast.AssumptionStatement) *smt.Generator {
	generator := smt.NewGenerator()
	generator.LoadMeta(uncertains, unknowns, asserts, assumes)
	generator.Run(ir)
	return generator
}

func probability(smt string, uncertains map[string][]float64, unknowns []string) (*execute.ModelChecker, map[string]execute.Scenario) {
	ex := execute.NewModelChecker("z3")
	ex.LoadModel(smt, uncertains, unknowns)
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
	uncertains := make(map[string][]float64)
	unknowns := []string{}

	data, err := os.ReadFile(filepath)
	if err != nil {
		log.Fatal(err)
	}
	d := string(data)
	path := gopath.Dir(filepath)

	switch input {
	case "fspec":
		lstnr, ty := parse(d, path, filepath)
		if lstnr == nil {
			log.Fatal("Fault parser returned nil")
		}
		uncertains = lstnr.Uncertains
		unknowns = lstnr.Unknowns

		if mode == "ast" {
			fmt.Println(lstnr.AST)
			return
		}

		compiler := ll(lstnr, ty)
		uncertains = compiler.Uncertains
		unknowns = compiler.Unknowns

		if mode == "ir" {
			fmt.Println(compiler.GetIR())
			return
		}

		generator := smt2(compiler.GetIR(), uncertains, unknowns, compiler.Asserts, compiler.Assumes)
		if mode == "smt" {
			fmt.Println(generator.SMT())
			return
		}

		mc, data := probability(generator.SMT(), uncertains, unknowns)
		mc.LoadMeta(generator.Branches, generator.BranchTrail)
		fmt.Println("~~~~~~~~~~\n  Fault found the following scenario\n~~~~~~~~~~")
		mc.Format(data)
	case "ll":
		generator := smt2(d, uncertains, unknowns, nil, nil)
		if mode == "smt" {
			fmt.Println(generator.SMT())
			return
		}

		mc, data := probability(generator.SMT(), uncertains, unknowns)
		mc.LoadMeta(generator.Branches, generator.BranchTrail)
		fmt.Println("~~~~~~~~~~\n  Fault found the following scenario\n~~~~~~~~~~")
		mc.Format(data)
	case "smt2":
		mc, data := probability(d, uncertains, unknowns)

		fmt.Println("~~~~~~~~~~\n  Fault found the following scenario\n~~~~~~~~~~")
		mc.Format(data)
	}
}

// func main() {
// 	p := tea.NewProgram(bubbles.New())
// 	if err := p.Start(); err != nil {
// 		fmt.Printf("Alas, there's been an error: %v", err)
// 		os.Exit(1)
// 	}
// }

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
