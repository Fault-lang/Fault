package tui

import "fault/runner"

// SetupCompleteMsg is sent when the user completes the setup view
type SetupCompleteMsg struct {
	Config runner.CompilationConfig
}

// ProgressUpdateMsg wraps a single progress update from the runner.
// progressCh and resultCh are threaded through so the Update handler can
// chain the next waitForProgress read without storing channels on the model.
type ProgressUpdateMsg struct {
	Update     runner.ProgressUpdate
	progressCh <-chan runner.ProgressUpdate
	resultCh   <-chan *runner.CompilationOutput
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

// LargeSMTWarningMsg is sent when the SMT formula exceeds the size threshold.
// The TUI transitions to a confirmation dialog; the runner goroutine is blocked
// waiting for a response on ConfirmCh.
type LargeSMTWarningMsg struct {
	SMTLines   int
	ConfirmCh  chan bool
	ProgressCh <-chan runner.ProgressUpdate
	ResultCh   <-chan *runner.CompilationOutput
}
