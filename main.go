package main

import (
	"bufio"
	"encoding/json"
	"fault/generator"
	"fault/runner"
	"fault/tui"
	"fault/util"
	"fmt"
	"os"
	"os/exec"
	ospath "path/filepath"
	"strings"

	tea "charm.land/bubbletea/v2"
	_ "github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
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

	var mode, input, output, filepath string
	var reach bool
	var smtThreshold, smtTimeout, smtMemory int

	rootCmd := &cobra.Command{
		Use:   "fault",
		Short: "Fault model checker",
		Long:  "Fault is a model checker for distributed systems. Run without a subcommand to compile a .fspec or .fsystem file, or omit -f to launch the interactive TUI.",
		RunE: func(cmd *cobra.Command, args []string) error {
			if filepath == "" {
				runInteractiveMode()
				return nil
			}
			return runTraditionalMode(filepath, mode, input, output, reach, smtThreshold, smtTimeout, smtMemory)
		},
	}

	rootCmd.Flags().StringVarP(&filepath, "file", "f", "", "path to file to compile")
	rootCmd.Flags().StringVarP(&mode, "mode", "m", "model", "stop compiler at certain milestones: ast, ir, smt, template, or model")
	rootCmd.Flags().StringVarP(&input, "input", "i", "fault", "format of the input file: fault, ll, or smt2")
	rootCmd.Flags().StringVar(&output, "output", "text", "format of the output: text or smt")
	rootCmd.Flags().BoolVar(&reach, "complete", false, "make sure transitions to all defined states are specified")
	rootCmd.Flags().IntVar(&smtThreshold, "smt-threshold", 0, fmt.Sprintf("warn before sending SMT formulas larger than this many lines to the solver (default: %d)", runner.LargeSMTThreshold))
	rootCmd.Flags().IntVar(&smtTimeout, "timeout", generator.DefaultSMTTimeout, "solver timeout in milliseconds via (set-option :timeout N); 0 = no limit")
	rootCmd.Flags().IntVar(&smtMemory, "memory-max-size", generator.DefaultSMTMemoryMaxSize, "solver memory limit in MB via (set-option :memory_max_size N); 0 = no limit")

	rootCmd.AddCommand(newConfigCmd())
	rootCmd.AddCommand(newUpdateCmd())

	rootCmd.SilenceErrors = true
	rootCmd.SilenceUsage = true

	if err := rootCmd.Execute(); err != nil {
		if tryPlugin(os.Args[1:]) {
			return
		}
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

// tryPlugin looks for a fault-<cmd> binary on PATH and executes it, passing
// all remaining args through. Returns true if the plugin was found and run.
func tryPlugin(args []string) bool {
	if len(args) == 0 {
		return false
	}
	pluginName := "fault-" + args[0]
	path, err := exec.LookPath(pluginName)
	if err != nil {
		return false
	}
	cmd := exec.Command(path, args[1:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		if exit, ok := err.(*exec.ExitError); ok {
			os.Exit(exit.ExitCode())
		}
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	return true
}

func newConfigCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "config",
		Short: "Get and set config values in ~/.faultrc",
		Long:  "Read or update values in ~/.faultrc. Flags set the corresponding config key; run with no flags to print current config.",
		RunE: func(cmd *cobra.Command, args []string) error {
			home, err := os.UserHomeDir()
			if err != nil {
				return fmt.Errorf("cannot determine home directory: %w", err)
			}
			configPath := ospath.Join(home, ".faultrc")

			changed := map[string]string{}
			cmd.Flags().Visit(func(f *pflag.Flag) {
				key := strings.ReplaceAll(strings.ToUpper(f.Name), "-", "_")
				changed[key] = f.Value.String()
			})

			if len(changed) == 0 {
				// Print current config
				data, err := os.ReadFile(configPath)
				if err != nil {
					return fmt.Errorf("could not read %s: %w", configPath, err)
				}
				fmt.Print(string(data))
				return nil
			}

			return updateConfigFile(configPath, changed)
		},
	}

	cmd.Flags().String("solverarg", "", "set SOLVERARG in ~/.faultrc")
	cmd.Flags().String("solvercmd", "", "set SOLVERCMD in ~/.faultrc")
	cmd.Flags().String("fault-host", "", "set FAULT_HOST in ~/.faultrc")

	return cmd
}

// updateConfigFile sets key=value pairs in the config file, updating existing
// uncommented or commented-out entries, and appending any that aren't found.
func updateConfigFile(configPath string, updates map[string]string) error {
	data, err := os.ReadFile(configPath)
	if err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("could not read %s: %w", configPath, err)
	}

	lines := strings.Split(string(data), "\n")
	applied := map[string]bool{}

	for i, line := range lines {
		trimmed := strings.TrimSpace(strings.TrimPrefix(strings.TrimSpace(line), "#"))
		key, _, ok := strings.Cut(trimmed, "=")
		if !ok {
			continue
		}
		key = strings.TrimSpace(key)
		if val, found := updates[key]; found {
			lines[i] = key + "=" + val
			applied[key] = true
		}
	}

	// Append any keys that weren't found in the file
	for key, val := range updates {
		if !applied[key] {
			lines = append(lines, key+"="+val)
		}
	}

	if err := os.WriteFile(configPath, []byte(strings.Join(lines, "\n")), 0644); err != nil {
		return fmt.Errorf("could not write %s: %w", configPath, err)
	}

	for key, val := range updates {
		fmt.Printf("Set %s=%s in %s\n", key, val, configPath)
	}
	return nil
}

func newUpdateCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "update",
		Short: "Check for a new version of fault",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("Update check not yet implemented.")
			return nil
		},
	}
}

func runTraditionalMode(filepath, mode, input, output string, reach bool, smtThreshold, smtTimeout, smtMemoryMaxSize int) error {
	mode = strings.ToLower(mode)
	switch mode {
	case "ast", "ir", "smt", "template", "model":
	default:
		return fmt.Errorf("%s is not a valid mode", mode)
	}

	output = strings.ToLower(output)
	switch output {
	case "text", "smt":
	default:
		return fmt.Errorf("%s is not a valid output format", output)
	}

	input = strings.ToLower(input)
	switch input {
	case "fault", "ll", "smt2":
	default:
		return fmt.Errorf("%s is not a valid input format", input)
	}

	if mode == "model" &&
		(os.Getenv("SOLVERCMD") == "" || os.Getenv("SOLVERARG") == "") {
		fmt.Fprintf(os.Stderr, "\nno solver configured, defaulting to SMT output without model checking. Please set SOLVERCMD and SOLVERARG in ~/.faultrc or your environment.\n\n")
		mode = "smt"
	}

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

	r := runner.NewRunner(config, nil)
	result := r.Run()

	if result.Error != nil {
		return result.Error
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
			return result.Error
		}
	}

	for _, w := range result.Warnings {
		fmt.Fprintln(os.Stderr, w)
	}

	if result.Message != "" {
		fmt.Println(result.Message)
	}

	switch mode {
	case "ast":
		if result.AST != nil {
			fmt.Printf("%v\n", result.AST)
		}
	case "ir":
		fmt.Println(result.IR)
	case "smt":
		fmt.Println(result.SMT)
	case "template":
		tmplPath := strings.TrimSuffix(filepath, ".fspec")
		tmplPath = strings.TrimSuffix(tmplPath, ".fsystem") + ".smt2.tmpl"
		if err := os.WriteFile(tmplPath, []byte(result.SMT), 0644); err != nil {
			return fmt.Errorf("error writing template: %w", err)
		}
		fmt.Fprintf(os.Stderr, "Template written to %s\n", tmplPath)

		manifestPath := strings.TrimSuffix(tmplPath, ".smt2.tmpl") + ".params.json"
		manifestBytes, err := json.MarshalIndent(result.ParamManifest, "", "  ")
		if err != nil {
			return fmt.Errorf("error marshaling param manifest: %w", err)
		}
		if err := os.WriteFile(manifestPath, manifestBytes, 0644); err != nil {
			return fmt.Errorf("error writing param manifest: %w", err)
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
	return nil
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
