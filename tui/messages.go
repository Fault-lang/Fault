package tui

import "fault/runner"

// SetupCompleteMsg is sent when the user completes the setup view
type SetupCompleteMsg struct {
	Config runner.CompilationConfig
}

// ProgressUpdateMsg wraps progress updates from the runner
type ProgressUpdateMsg struct {
	Update runner.ProgressUpdate
}

// CompilationCompleteMsg is sent when compilation finishes successfully
type CompilationCompleteMsg struct {
	Output *runner.CompilationOutput
}

// CompilationErrorMsg is sent when compilation fails
type CompilationErrorMsg struct {
	Error error
	Phase runner.ProgressPhase
}

// ValidationErrorMsg is sent when validation fails before compilation
type ValidationErrorMsg struct {
	Error error
}

// RetryCompilationMsg is sent when user wants to retry after error
type RetryCompilationMsg struct{}

// BackToSetupMsg is sent when user wants to go back to setup from error
type BackToSetupMsg struct{}
