package main

import (
	"fault/ast"
	"fault/execute"
	"fault/listener"
	"fault/llvm"
	"fault/preprocess"
	"fault/reachability"
	"fault/smt"
	smtvar "fault/smt/variables"
	"fault/swaps"
	"fault/types"
	"fault/util"
	"fault/visualize"
	"flag"
	"fmt"
	"log"
	"os"
	gopath "path"
	"strings"

	_ "github.com/olekukonko/tablewriter"
)

func parse(data string, path string, file string, filetype string, reach bool, visu bool) (*ast.Spec, *listener.FaultListener, *types.Checker, string) {
	//Confirm that the filetype and file declaration match
	if !validate_filetype(data, filetype) {
		log.Fatalf("malformatted file: declaration does not match filetype.")
	}
	flags := make(map[string]bool)
	flags["specType"] = (filetype == "fspec")
	flags["testing"] = false
	flags["skipRun"] = false
	lstnr := listener.Execute(data, path, flags)

	pre := preprocess.Execute(lstnr)

	ty := types.Execute(pre.Processed, pre)

	sw := swaps.NewPrecompiler(ty)
	tree = sw.Swap(tree)

	var visual string
	if visu {
		vis := visualize.NewVisual(ty.Checked)
		vis.Build()
		visual = vis.Render()
	}

	if reach {
		r := reachability.NewTracer()
		r.Scan(ty.Checked)
	}
	return ty.Checked, lstnr, ty, visual
}

func validate_filetype(data string, filetype string) bool {
	if filetype == "fspec" && data[0:4] == "spec" {
		return true
	}
	if filetype == "fsystem" && data[0:6] == "system" {
		return true
	}
	return false
}

func smt2(ir string, runs int16, uncertains map[string][]float64, unknowns []string, asserts []*ast.AssertionStatement, assumes []*ast.AssertionStatement) *smt.Generator {
	generator := smt.NewGenerator()
	generator.LoadMeta(runs, uncertains, unknowns, asserts, assumes)
	generator.Run(ir)
	return generator
}

func probability(smt string, uncertains map[string][]float64, unknowns []string, results map[string][]*smtvar.VarChange) (*execute.ModelChecker, map[string]execute.Scenario) {
	ex := execute.NewModelChecker()
	ex.LoadModel(smt, uncertains, unknowns, results)
	ok, err := ex.Check()
	if err != nil {
		log.Fatalf("model checker has failed: %s", err)
	}
	if !ok {
		fmt.Println("Fault could not find a failure case.")
		return ex, nil
	}
	scenario, err := ex.Solve()
	if err != nil {
		log.Fatalf("error found fetching solution from solver: %s", err)
	}
	data := ex.Filter(scenario)
	return ex, data
}

func run(filepath string, mode string, input string, reach bool) {
	filetype := util.DetectMode(filepath)
	if filetype == "" {
		log.Fatal("file provided is not a .fspec or .fsystem file")
	}

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
		tree, lstnr, ty, visual := parse(d, path, filepath, filetype, reach, mode == "visualize")
		if lstnr == nil {
			log.Fatal("Fault parser returned nil")
		}

		if mode == "ast" {
			fmt.Println(lstnr.AST)
			return
		}

		compiler := llvm.Execute(tree, ty.SpecStructs, lstnr.Uncertains, lstnr.Unknowns, false)
		uncertains = compiler.Uncertains
		unknowns = compiler.Unknowns

		if mode == "ir" {
			fmt.Println(compiler.GetIR())
			return
		}

		if !compiler.IsValid && visual != "" {
			fmt.Println(visual)
			fmt.Printf("\n\n")
			return
		}

		if !compiler.IsValid {
			fmt.Println("Fault found nothing to run. Missing run block or start block.")
			return
		}

		generator := smt.Execute(compiler)
		if mode == "smt" {
			fmt.Println(generator.SMT())
			return
		}

		mc, data := probability(generator.SMT(), uncertains, unknowns, generator.Results)
		if mode == "visualize" {
			fmt.Println(visual)
			fmt.Printf("\n\n")
			mc.Mermaid()
			return
		}

		if data != nil {
			mc.LoadMeta(generator.GetForks())
			fmt.Println("~~~~~~~~~~\n  Fault found the following scenario\n~~~~~~~~~~")
			mc.Format(data)
		}
	case "ll":
		generator := smt2(d, 0, uncertains, unknowns, nil, nil)
		if mode == "smt" {
			fmt.Println(generator.SMT())
			return
		}

		mc, data := probability(generator.SMT(), uncertains, unknowns, generator.Results)
		if mode == "visualize" {
			mc.Mermaid()
			return
		}
		if data != nil {
			mc.LoadMeta(generator.GetForks())
			fmt.Println("~~~~~~~~~~\n  Fault found the following scenario\n~~~~~~~~~~")
			mc.Format(data)
		}
	case "smt2":
		mc, data := probability(d, uncertains, unknowns, make(map[string][]*smtvar.VarChange))

		if mode == "visualize" {
			mc.Mermaid()
			return
		}
		if data != nil {
			fmt.Println("~~~~~~~~~~\n  Fault found the following scenario\n~~~~~~~~~~")
			mc.Format(data)
		}
	}
}

func main() {
	var mode string
	var input string
	var filepath string
	var reach bool
	modeCommand := flag.String("m", "check", "stop compiler at certain milestones: ast, ir, smt, or check")
	inputCommand := flag.String("i", "fspec", "format of the input file (default: fspec)")
	fpCommand := flag.String("f", "", "path to file to compile")
	reachCommand := flag.String("c", "false", "make sure the transitions to all defined states are specified in the model")

	flag.Parse()

	if *fpCommand == "" {
		flag.PrintDefaults()
		fmt.Println("must provide path of file to compile")
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
		case "visualize":
		default:
			fmt.Printf("%s is not a valid mode\n", mode)
			os.Exit(1)
		}
	}

	//Check if solver is set
	if (mode == "check" || mode == "visualize") &&
		os.Getenv("SOLVERCMD") == "" || os.Getenv("SOLVERARG") == "" {
		fmt.Printf("\n no solver configured, defaulting to SMT output without model checking. Please set SOLVERCMD and SOLVERARG variables.\n\n")
		mode = "smt"
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
			fmt.Printf("%s is not a valid input format\n", input)
			os.Exit(1)
		}
	}

	if *reachCommand == "" {
		reach = false
	} else {
		r := strings.ToLower(*reachCommand)
		switch r {
		case "true":
			reach = true
		case "false":
			reach = false
		case "t":
			reach = true
		case "f":
			reach = false
		default:
			fmt.Printf("%s is not a valid option for completeness please use true or false\n", r)
			os.Exit(1)
		}
	}

	run(filepath, mode, input, reach)
}
