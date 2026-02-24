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

	// Categorize based on error message and phase
	switch {
	case strings.Contains(errMsg, "no such file") || strings.Contains(errMsg, "cannot find"):
		enhanced.Category = ErrorFile
		enhanced.Suggestion = "Check that the file path is correct and the file exists"

	case strings.Contains(errMsg, "permission denied"):
		enhanced.Category = ErrorFile
		enhanced.Suggestion = "Check file permissions - you may need read access"

	case strings.Contains(errMsg, "malformatted file") || strings.Contains(errMsg, "declaration does not match"):
		enhanced.Category = ErrorFile
		enhanced.Suggestion = "Ensure the file starts with 'spec' or 'system' declaration"

	case strings.Contains(errMsg, "not a .fspec or .fsystem file"):
		enhanced.Category = ErrorFile
		enhanced.Suggestion = "Use a file with .fspec or .fsystem extension"

	case phase == runner.PhaseParsing:
		enhanced.Category = ErrorParsing
		enhanced.Detail = "The parser could not understand the file syntax"
		enhanced.Suggestion = "Check for syntax errors in your specification"

	case phase == runner.PhaseTypeChecking:
		enhanced.Category = ErrorTypeCheck
		enhanced.Detail = "Type checking found an issue with types or declarations"
		enhanced.Suggestion = "Check variable declarations and type compatibility"

	case phase == runner.PhaseLLVM:
		enhanced.Category = ErrorLLVM
		enhanced.Detail = "LLVM IR generation encountered an issue"
		enhanced.Suggestion = "Ensure your specification has a valid run or start block"

	case phase == runner.PhaseSMT:
		enhanced.Category = ErrorSMT
		enhanced.Detail = "SMT constraint generation failed"
		enhanced.Suggestion = "Check your logic expressions and constraints"

	case strings.Contains(errMsg, "model checker has failed") || strings.Contains(errMsg, "solver"):
		enhanced.Category = ErrorSolver
		enhanced.Detail = "The SMT solver encountered an error"
		if os.Getenv("SOLVERCMD") == "" || os.Getenv("SOLVERARG") == "" {
			enhanced.Suggestion = "Configure solver: set SOLVERCMD and SOLVERARG environment variables"
		} else {
			enhanced.Suggestion = "Check solver configuration and try again"
		}

	case phase == runner.PhaseModelChecking:
		enhanced.Category = ErrorSolver
		enhanced.Suggestion = "Verify solver is correctly configured and accessible"

	default:
		enhanced.Category = ErrorInternal
		enhanced.Suggestion = "Try again or report this issue if it persists"
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
