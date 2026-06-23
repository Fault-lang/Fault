package main

import (
	"fault/listener"
	"fault/preprocess"
	"fault/types"
	"fault/util"
	"fmt"
	"os"
	gopath "path/filepath"

	"github.com/spf13/cobra"
)

func newLintCmd() *cobra.Command {
	var warnOnly bool

	cmd := &cobra.Command{
		Use:   "lint",
		Short: "Report type errors and semantic issues without compiling",
		Long: `Lint runs the parser and type checker against a .fspec or .fsystem file,
collecting all recoverable errors rather than stopping at the first one.

Exits with code 1 if any issues are found (use --warn-only to suppress).

Example:
  fault lint -f=model.fspec
  fault lint -f=model.fspec --warn-only`,
		RunE: func(cmd *cobra.Command, args []string) error {
			fp, _ := cmd.Flags().GetString("file")
			if fp == "" {
				return fmt.Errorf("required flag \"file\" not set")
			}
			return runLint(fp, warnOnly)
		},
	}

	cmd.Flags().StringP("file", "f", "", "path to .fspec or .fsystem file")
	cmd.Flags().BoolVar(&warnOnly, "warn-only", false, "report issues but always exit 0")
	return cmd
}

func runLint(filepath string, warnOnly bool) error {
	filetype := util.DetectMode(filepath)
	if filetype != "fspec" && filetype != "fsystem" {
		return fmt.Errorf("%s is not a .fspec or .fsystem file", filepath)
	}

	resolved := util.Filepath(filepath)
	data, err := os.ReadFile(resolved)
	if err != nil {
		return fmt.Errorf("could not read file: %w", err)
	}
	path := gopath.Dir(resolved)

	flags := map[string]bool{
		"specType": filetype == "fspec",
		"testing":  false,
		"skipRun":  false,
	}

	// Listener (parse) — fatal if this fails
	lstnr, err := listener.Execute(string(data), path, flags)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: parse error: %v\n", filepath, err)
		if warnOnly {
			return nil
		}
		os.Exit(1)
	}

	// Preprocess in lint mode — collects recoverable errors
	pre, err := preprocess.ExecuteLint(lstnr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: fatal preprocess error: %v\n", filepath, err)
		if warnOnly {
			return nil
		}
		os.Exit(1)
	}

	// Type check in lint mode — collects all recoverable errors
	ty, err := types.ExecuteLint(pre.Processed, pre)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: fatal type error: %v\n", filepath, err)
		if warnOnly {
			return nil
		}
		os.Exit(1)
	}

	errs := append(pre.Errors(), ty.Errors()...)
	if len(errs) == 0 {
		fmt.Printf("%s: no issues found\n", filepath)
		return nil
	}

	for _, e := range errs {
		fmt.Fprintf(os.Stderr, "%s: %v\n", filepath, e)
	}

	if !warnOnly {
		os.Exit(1)
	}
	return nil
}
