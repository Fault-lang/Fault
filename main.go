package main

import (
	"bufio"
	"fault/runner"
	"fault/tui"
	"fault/util"
	"flag"
	"fmt"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	_ "github.com/olekukonko/tablewriter"
)

func main() {
	var mode string
	var input string
	var output string
	var filepath string
	var reach bool
	modeCommand := flag.String("m", "model", "stop compiler at certain milestones: ast, ir, smt, or model")
	inputCommand := flag.String("i", "fault", "format of the input file (default: fault)")
	fpCommand := flag.String("f", "", "path to file to compile")
	reachCommand := flag.Bool("complete", false, "make sure the transitions to all defined states are specified in the model")
	outputCommand := flag.String("output", "text", "format of the output: text or smt")
	smtThresholdCommand := flag.Int("smt-threshold", 0, fmt.Sprintf("warn before sending SMT formulas larger than this many lines to the solver (default: %d)", runner.LargeSMTThreshold))

	flag.Parse()

	// HYBRID MODE DETECTION
	if *fpCommand == "" {
		// No file provided - launch interactive TUI mode
		runInteractiveMode()
		return
	}

	// Traditional CLI mode - file was provided
	filepath = *fpCommand

	if *modeCommand == "" {
		mode = "model"
	} else {
		mode = strings.ToLower(*modeCommand)
		switch mode {
		case "ast":
		case "ir":
		case "smt":
		case "model":
		default:
			fmt.Printf("%s is not a valid mode\n", mode)
			os.Exit(1)
		}
	}

	if *outputCommand == "" {
		output = "text"
	} else {
		output = strings.ToLower(*outputCommand)
		switch output {
		case "text":
		case "smt":
		default:
			fmt.Printf("%s is not a valid mode\n", output)
			os.Exit(1)
		}
	}

	// Check if solver is set
	if mode == "model" &&
		(os.Getenv("SOLVERCMD") == "" || os.Getenv("SOLVERARG") == "") {
		fmt.Printf("\nno solver configured, defaulting to SMT output without model checking. Please set SOLVERCMD and SOLVERARG variables.\n\n")
		mode = "smt"
	}

	if *inputCommand == "" {
		input = "fault"
	} else {
		input = strings.ToLower(*inputCommand)
		switch input {
		case "fault":
		case "ll":
		case "smt2":
		default:
			fmt.Printf("%s is not a valid input format\n", input)
			os.Exit(1)
		}
	}

	if *reachCommand {
		reach = true
	}

	runTraditionalMode(filepath, mode, input, output, reach, *smtThresholdCommand)
}

func runTraditionalMode(filepath, mode, input, output string, reach bool, smtThreshold int) {
	config := runner.CompilationConfig{
		Filepath:             filepath,
		Mode:                 mode,
		Input:                input,
		Output:               output,
		Reach:                reach,
		LargeSMTLineOverride: smtThreshold,
	}

	// Run without progress updates (nil channel)
	r := runner.NewRunner(config, nil)
	result := r.Run()

	if result.Error != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", result.Error)
		os.Exit(1)
	}

	// Large SMT: prompt before sending to solver.
	if result.LargeSMTLines > 0 && result.Pending != nil {
		fmt.Fprintf(os.Stderr, "\nWarning: the SMT formula is %d lines.\n", result.LargeSMTLines)
		fmt.Fprintf(os.Stderr, "Sending a formula this large to the solver may take a very long time.\n")
		fmt.Fprintf(os.Stderr, "Proceed with model checking? [y/N]: ")
		reader := bufio.NewReader(os.Stdin)
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(strings.ToLower(line))
		if line != "y" && line != "yes" {
			fmt.Fprintln(os.Stderr, "Aborted.")
			os.Exit(0)
		}
		result = r.Resume(result.Pending)
		if result.Error != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", result.Error)
			os.Exit(1)
		}
	}

	for _, w := range result.Warnings {
		fmt.Fprintln(os.Stderr, w)
	}

	if result.Message != "" {
		fmt.Println(result.Message)
	}

	// Print results based on mode
	switch mode {
	case "ast":
		if result.AST != nil {
			printAST(result.AST)
		}
	case "ir":
		fmt.Println(result.IR)
	case "smt":
		fmt.Println(result.SMT)
	case "model":
		if result.ResultLog != nil {
			result.ResultLog.Print()
			for _, a := range result.Asserts {
				fmt.Println(a.EvLogString(true))
			}
		} else {
			fmt.Println(result.SMT)
		}
	}
}

func printAST(spec interface{}) {
	fmt.Printf("%v\n", spec)
}

func runInteractiveMode() {
	p := tea.NewProgram(
		tui.NewModel(),
		tea.WithAltScreen(),       // Full-screen TUI
		tea.WithMouseCellMotion(), // Enable mouse support
	)

	if _, err := p.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error running TUI: %v\n", err)
		os.Exit(1)
	}
}

// Helper function to validate filetype (used by runner)
func validateFiletype(filepath string) bool {
	filetype := util.DetectMode(filepath)
	return filetype == "fspec" || filetype == "fsystem"
}
