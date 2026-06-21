package main

import (
	"bufio"
	"encoding/json"
	"fault/generator"
	"fault/runner"
	"fault/tui"
	"fault/util"
	"flag"
	"fmt"
	"os"
	ospath "path/filepath"
	"strings"

	tea "charm.land/bubbletea/v2"
	_ "github.com/olekukonko/tablewriter"
)

var defaultFaultrc = `# Fault compiler configuration
# Values here are overridden by environment variables set in your shell.

# Path to your SMT solver binary (e.g. z3, cvc5)
# SOLVERCMD=z3

# Argument to make the solver read from stdin
# SOLVERARG=-in

# Base directory for resolving relative file paths (~ or ..).
# Only needed if you reference .fspec/.fsystem files using relative paths.
# FAULT_HOST=
`

func loadConfig() {
	home, err := os.UserHomeDir()
	if err != nil {
		return
	}
	configPath := ospath.Join(home, ".faultrc")
	f, err := os.Open(configPath)
	if os.IsNotExist(err) {
		fmt.Fprintf(os.Stderr, "Warning: no ~/.faultrc found. Creating a default one at %s\n", configPath)
		if werr := os.WriteFile(configPath, []byte(defaultFaultrc), 0644); werr != nil {
			fmt.Fprintf(os.Stderr, "Warning: could not create default config: %v\n", werr)
		}
		return
	}
	if err != nil {
		return
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		key, val, ok := strings.Cut(line, "=")
		if !ok {
			continue
		}
		// Only set if not already set — env always wins
		if os.Getenv(strings.TrimSpace(key)) == "" {
			os.Setenv(strings.TrimSpace(key), strings.TrimSpace(val))
		}
	}
}

func main() {
	loadConfig()
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
	smtTimeoutCommand := flag.Int("timeout", generator.DefaultSMTTimeout, "solver timeout in milliseconds via (set-option :timeout N); 0 = no limit")
	smtMemoryCommand := flag.Int("memory-max-size", generator.DefaultSMTMemoryMaxSize, "solver memory limit in MB via (set-option :memory_max_size N); 0 = no limit")

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
		case "template":
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

	runTraditionalMode(filepath, mode, input, output, reach, *smtThresholdCommand, *smtTimeoutCommand, *smtMemoryCommand)
}

func runTraditionalMode(filepath, mode, input, output string, reach bool, smtThreshold, smtTimeout, smtMemoryMaxSize int) {
	config := runner.CompilationConfig{
		Filepath:             filepath,
		Mode:                 mode,
		Input:                input,
		Output:               output,
		Reach:                reach,
		LargeSMTLineOverride: smtThreshold,
		SMTTimeout:           smtTimeout,
		SMTMemoryMaxSize:     smtMemoryMaxSize,
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
	case "template":
		// Write <filepath>.smt2.tmpl
		tmplPath := strings.TrimSuffix(filepath, ".fspec")
		tmplPath = strings.TrimSuffix(tmplPath, ".fsystem") + ".smt2.tmpl"
		if err := os.WriteFile(tmplPath, []byte(result.SMT), 0644); err != nil {
			fmt.Fprintf(os.Stderr, "Error writing template: %v\n", err)
			os.Exit(1)
		}
		fmt.Fprintf(os.Stderr, "Template written to %s\n", tmplPath)

		// Write <filepath>.params.json
		manifestPath := strings.TrimSuffix(tmplPath, ".smt2.tmpl") + ".params.json"
		manifestBytes, err := json.MarshalIndent(result.ParamManifest, "", "  ")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error marshaling param manifest: %v\n", err)
			os.Exit(1)
		}
		if err := os.WriteFile(manifestPath, manifestBytes, 0644); err != nil {
			fmt.Fprintf(os.Stderr, "Error writing param manifest: %v\n", err)
			os.Exit(1)
		}
		fmt.Fprintf(os.Stderr, "Param manifest written to %s\n", manifestPath)
	case "model":
		if result.ResultLog != nil {
			result.ResultLog.Print()
			sysPrefix := result.ResultLog.SystemName + "_"
			for _, a := range result.Asserts {
				s := a.EvLogString(true)
				if sysPrefix != "_" {
					s = strings.ReplaceAll(s, sysPrefix, "")
				}
				fmt.Println(s)
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
	p := tea.NewProgram(tui.NewModel())

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
