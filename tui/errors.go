package tui

import (
	"fault/runner"
	"fmt"
	"os"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

type ErrorCategory int

const (
	ErrorFile ErrorCategory = iota
	ErrorParsing
	ErrorTypeCheck
	ErrorLLVM
	ErrorSMT
	ErrorSolver
	ErrorInternal
)

// EnhancedError wraps an error with additional context
type EnhancedError struct {
	Category    ErrorCategory
	Phase       runner.ProgressPhase
	Message     string
	Detail      string
	Suggestion  string
	OriginalErr error
}

func (e *EnhancedError) Error() string {
	return e.Message
}

// CategorizeError analyzes an error and provides categorization and suggestions
func CategorizeError(err error, phase runner.ProgressPhase) *EnhancedError {
	if err == nil {
		return nil
	}

	errMsg := err.Error()
	enhanced := &EnhancedError{
		OriginalErr: err,
		Phase:       phase,
		Message:     errMsg,
	}

	// Categorize based on error message content first, then fall back to phase.
	// More specific patterns are listed before broader ones.
	switch {

	// ── File / setup errors ──────────────────────────────────────────────────

	case strings.Contains(errMsg, "no such file") || strings.Contains(errMsg, "cannot find"):
		enhanced.Category = ErrorFile
		enhanced.Suggestion = "Check that the file path is correct and the file exists."

	case strings.Contains(errMsg, "permission denied"):
		enhanced.Category = ErrorFile
		enhanced.Suggestion = "Check file permissions — you may need read access to this file."

	case strings.Contains(errMsg, "malformatted file") || strings.Contains(errMsg, "declaration does not match"):
		enhanced.Category = ErrorFile
		enhanced.Suggestion = "The file type does not match its declaration.\n" +
			"  .fspec files must start with 'spec <name>;'\n" +
			"  .fsystem files must start with 'system <name>;'"

	case strings.Contains(errMsg, "not a .fspec or .fsystem file"):
		enhanced.Category = ErrorFile
		enhanced.Suggestion = "Fault only accepts .fspec and .fsystem files.\n" +
			"Rename your file or choose a different one."

	// ── Solver / environment errors ──────────────────────────────────────────

	case strings.Contains(errMsg, "missing SOLVERCMD"):
		enhanced.Category = ErrorSolver
		enhanced.Detail = "The SOLVERCMD environment variable is not set."
		enhanced.Suggestion = "Set SOLVERCMD to your solver binary, e.g.:\n  export SOLVERCMD=z3"

	case strings.Contains(errMsg, "missing SOLVERARG"):
		enhanced.Category = ErrorSolver
		enhanced.Detail = "The SOLVERARG environment variable is not set."
		enhanced.Suggestion = "Set SOLVERARG to the solver's stdin flag, e.g.:\n  export SOLVERARG=-in"

	case strings.Contains(errMsg, "unexpected empty response from solver"):
		enhanced.Category = ErrorSolver
		enhanced.Detail = "The solver ran but produced no output."
		enhanced.Suggestion = "Check that SOLVERCMD points to a working solver binary and that\n" +
			"the solver version is compatible."

	case strings.Contains(errMsg, "no model found"):
		enhanced.Category = ErrorSolver
		enhanced.Detail = "The solver reported the constraints are unsatisfiable."
		enhanced.Suggestion = "Fault could not find a failure case. Your assertions may be too strict,\n" +
			"or the model may be correct. Try relaxing assert conditions."

	case strings.Contains(errMsg, "model checker has failed"):
		enhanced.Category = ErrorSolver
		enhanced.Detail = "The SMT solver returned an unexpected error."
		enhanced.Suggestion = "Check your SOLVERCMD and SOLVERARG settings, then retry."

	// ── Parsing errors (specific patterns first) ─────────────────────────────

	case strings.Contains(errMsg, "Variable names must be only letters or numbers"):
		enhanced.Category = ErrorParsing
		enhanced.Detail = "A variable or property name contains an invalid character."
		enhanced.Suggestion = "Variable names may only contain letters (a–z, A–Z) and digits (0–9).\n" +
			"Remove underscores, hyphens, dots, or other special characters."

	case strings.Contains(errMsg, "A function cannot be empty") ||
		strings.Contains(errMsg, "A state function cannot be empty"):
		enhanced.Category = ErrorParsing
		enhanced.Detail = "An empty function body was found."
		enhanced.Suggestion = "Every func{} block must contain at least one statement.\n" +
			"Add a body, e.g.: func{ stay(); }"

	case strings.Contains(errMsg, "Too few statements"):
		enhanced.Category = ErrorParsing
		enhanced.Detail = "The specification has too few top-level statements."
		enhanced.Suggestion = "A valid spec needs at minimum a declaration and one component definition.\n" +
			"Check that the file is complete and not truncated."

	case strings.Contains(errMsg, "stack underflow"):
		enhanced.Category = ErrorInternal
		enhanced.Detail = "The parser encountered an unexpected structure."
		enhanced.Suggestion = "This usually means the file has a syntax error that confused the parser.\n" +
			"Look for mismatched braces, missing semicolons, or incomplete expressions."

	case phase == runner.PhaseParsing:
		enhanced.Category = ErrorParsing
		enhanced.Detail = "The parser could not understand the file syntax."
		enhanced.Suggestion = "Check for syntax errors: mismatched braces, missing semicolons,\n" +
			"or unrecognised keywords."

	// ── Type-checking errors ─────────────────────────────────────────────────

	case phase == runner.PhaseTypeChecking:
		enhanced.Category = ErrorTypeCheck
		enhanced.Detail = "Type checking found an inconsistency in declarations or types."
		enhanced.Suggestion = "Check variable declarations and that all referenced types are defined.\n" +
			"Make sure stock/flow/component names are spelled consistently."

	// ── LLVM / compilation errors ────────────────────────────────────────────

	case strings.Contains(errMsg, "Missing run block") || strings.Contains(errMsg, "Missing start block") ||
		strings.Contains(errMsg, "missing run block") || strings.Contains(errMsg, "missing start block"):
		enhanced.Category = ErrorLLVM
		enhanced.Detail = "No run or start block was found in the specification."
		enhanced.Suggestion = "Add a run block to your .fspec:\n  for <n> run { ... }\n" +
			"Or a start block to your .fsystem:\n  start { <component>: <state>, };"

	case strings.Contains(errMsg, "Internal compiler stacktrace"):
		enhanced.Category = ErrorInternal
		enhanced.Detail = errMsg
		enhanced.Message = "An internal compiler error occurred."
		enhanced.Suggestion = "This is likely a bug in Fault. Please report it at\n" +
			"https://github.com/fault-lang/fault/issues with the full error message."

	case phase == runner.PhaseLLVM:
		enhanced.Category = ErrorLLVM
		enhanced.Detail = "LLVM IR generation encountered an issue."
		enhanced.Suggestion = "Ensure your specification has a valid run or start block\n" +
			"and that all referenced components are defined."

	// ── SMT generation errors ────────────────────────────────────────────────

	case phase == runner.PhaseSMT:
		enhanced.Category = ErrorSMT
		enhanced.Detail = "SMT constraint generation failed."
		enhanced.Suggestion = "Check your logic expressions, assert conditions, and\n" +
			"that all variables used in asserts are declared."

	// ── Reachability errors ──────────────────────────────────────────────────

	case strings.Contains(errMsg, "unreachable"):
		enhanced.Category = ErrorLLVM
		enhanced.Detail = "Reachability analysis found states that can never be reached."
		enhanced.Suggestion = "Check that every state listed in the error is referenced by an\n" +
			"advance() call or a start block entry. Remove or connect unreachable states."

	// ── Model-checking phase fallback ────────────────────────────────────────

	case phase == runner.PhaseModelChecking:
		enhanced.Category = ErrorSolver
		enhanced.Detail = "The model checker encountered a problem."
		enhanced.Suggestion = "Verify the solver is correctly configured (SOLVERCMD / SOLVERARG)\n" +
			"and that the generated SMT is well-formed."

	default:
		enhanced.Category = ErrorInternal
		enhanced.Suggestion = "This is an unexpected error. Try again, or report it at\n" +
			"https://github.com/fault-lang/fault/issues"
	}

	return enhanced
}

// ValidateSetupConfig validates the configuration before starting compilation
func ValidateSetupConfig(config runner.CompilationConfig) error {
	// Check if file exists
	if _, err := os.Stat(config.Filepath); os.IsNotExist(err) {
		return fmt.Errorf("file not found: %s", config.Filepath)
	}

	// Check file permissions
	file, err := os.Open(config.Filepath)
	if err != nil {
		return fmt.Errorf("cannot read file: %v", err)
	}
	file.Close()

	// Check solver configuration if mode is "model"
	if config.Mode == "model" {
		if os.Getenv("SOLVERCMD") == "" || os.Getenv("SOLVERARG") == "" {
			return fmt.Errorf("solver not configured: set SOLVERCMD and SOLVERARG environment variables")
		}
	}

	return nil
}

// RenderPhaseProgress renders the phase progression showing where the error occurred
func RenderPhaseProgress(failedPhase runner.ProgressPhase) string {
	phases := []struct {
		phase runner.ProgressPhase
		name  string
	}{
		{runner.PhaseParsing, "Parsing"},
		{runner.PhasePreprocessing, "Preprocessing"},
		{runner.PhaseTypeChecking, "Type Checking"},
		{runner.PhaseLLVM, "LLVM IR Generation"},
		{runner.PhaseSMT, "SMT Generation"},
		{runner.PhaseModelChecking, "Model Checking"},
		{runner.PhaseResults, "Results"},
	}

	var b strings.Builder
	b.WriteString(SubtitleStyle.Render("Pipeline Progress:"))
	b.WriteString("\n\n")

	for _, p := range phases {
		var status, symbol string
		var style lipgloss.Style

		if p.phase < failedPhase {
			symbol = "✓"
			status = "complete"
			style = PhaseDoneStyle
		} else if p.phase == failedPhase {
			symbol = "✗"
			status = "failed"
			style = ErrorStyle
		} else {
			symbol = "⋯"
			status = "not started"
			style = PhasePendingStyle
		}

		line := fmt.Sprintf("%s %s (%s)", symbol, p.name, status)
		b.WriteString(style.Render(line))
		b.WriteString("\n")
	}

	return b.String()
}

// GetPhaseName returns a human-readable name for a phase
func GetPhaseName(phase runner.ProgressPhase) string {
	switch phase {
	case runner.PhaseParsing:
		return "Parsing"
	case runner.PhasePreprocessing:
		return "Preprocessing"
	case runner.PhaseTypeChecking:
		return "Type Checking"
	case runner.PhaseLLVM:
		return "LLVM IR Generation"
	case runner.PhaseSMT:
		return "SMT Generation"
	case runner.PhaseModelChecking:
		return "Model Checking"
	case runner.PhaseResults:
		return "Results Processing"
	default:
		return "Unknown Phase"
	}
}
