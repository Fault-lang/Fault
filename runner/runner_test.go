package runner

import (
	"os"
	"path/filepath"
	"testing"
)

// simpleSpec is a minimal fspec that produces a small but non-trivial SMT formula.
const simpleSpec = `spec test1;
def s = stock{
	x: 10.0,
};
def f = flow{
	tank: new s,
	fill: func{
		tank.x <- tank.x + 1.0;
	},
};
for 1 init { t = new f; } run { t.fill; };
assert t.tank.x < 20.0;
`

func writeTempSpec(t *testing.T, content string) string {
	t.Helper()
	dir := t.TempDir()
	path := filepath.Join(dir, "test.fspec")
	if err := os.WriteFile(path, []byte(content), 0600); err != nil {
		t.Fatalf("could not write temp spec: %v", err)
	}
	return path
}

// TestLargeSMTCLIPath verifies that when the SMT exceeds the threshold with no
// progress channel (CLI mode), Run returns early with LargeSMTLines and Pending
// populated, and that Resume completes model checking successfully.
func TestLargeSMTCLIPath(t *testing.T) {
	path := writeTempSpec(t, simpleSpec)

	config := CompilationConfig{
		Filepath:             path,
		Mode:                 "model",
		Input:                "fault",
		Output:               "text",
		LargeSMTLineOverride: 1, // trigger on any non-trivial SMT
	}

	r := NewRunner(config, nil)
	result := r.Run()

	if result.Error != nil {
		t.Fatalf("unexpected error: %v", result.Error)
	}
	if result.LargeSMTLines == 0 {
		t.Fatal("expected LargeSMTLines > 0, got 0 — threshold was not overridden")
	}
	if result.Pending == nil {
		t.Fatal("expected Pending to be set when LargeSMTLines > 0")
	}
	if result.Pending.SMT == "" {
		t.Fatal("expected Pending.SMT to be non-empty")
	}
	if result.Pending.ResultLog == nil {
		t.Fatal("expected Pending.ResultLog to be non-nil")
	}

	// Now resume — simulating the user saying "yes".
	resumed := r.Resume(result.Pending)
	if resumed.Error != nil {
		t.Fatalf("Resume returned error: %v", resumed.Error)
	}
	// The result should have either a message (no failure found) or a result log.
	if resumed.Message == "" && resumed.ResultLog == nil {
		t.Fatal("Resume returned neither a message nor a ResultLog")
	}
}

// TestLargeSMTTUIPath verifies that when a progress channel is present (TUI mode),
// Run sends a PhaseConfirmLargeSMT update and blocks; sending true unblocks it and
// compilation completes; sending false unblocks it and Run returns with LargeSMTLines set.
func TestLargeSMTTUIPath(t *testing.T) {
	t.Run("confirm proceeds", func(t *testing.T) {
		path := writeTempSpec(t, simpleSpec)

		progressCh := make(chan ProgressUpdate, 20)
		resultCh := make(chan *CompilationOutput, 1)

		config := CompilationConfig{
			Filepath:             path,
			Mode:                 "model",
			Input:                "fault",
			Output:               "text",
			LargeSMTLineOverride: 1,
		}

		go func() {
			r := NewRunner(config, progressCh)
			result := r.Run()
			resultCh <- result
			close(progressCh)
		}()

		// Drain progress updates until we see PhaseConfirmLargeSMT.
		var confirmCh chan bool
		for update := range progressCh {
			if update.Phase == PhaseConfirmLargeSMT {
				if update.SMTLines == 0 {
					t.Error("PhaseConfirmLargeSMT update has SMTLines == 0")
				}
				if update.ConfirmCh == nil {
					t.Error("PhaseConfirmLargeSMT update has nil ConfirmCh")
				}
				confirmCh = update.ConfirmCh
				break
			}
		}

		if confirmCh == nil {
			t.Fatal("never received PhaseConfirmLargeSMT update")
		}

		// Confirm: runner should continue and complete.
		confirmCh <- true

		// Drain remaining progress updates then read the result.
		for range progressCh {
		}
		result := <-resultCh
		if result.Error != nil {
			t.Fatalf("unexpected error after confirm: %v", result.Error)
		}
		if result.LargeSMTLines != 0 {
			t.Fatalf("expected LargeSMTLines == 0 after confirm, got %d", result.LargeSMTLines)
		}
		if result.Message == "" && result.ResultLog == nil {
			t.Fatal("expected a result after confirming")
		}
	})

	t.Run("abort returns early", func(t *testing.T) {
		path := writeTempSpec(t, simpleSpec)

		progressCh := make(chan ProgressUpdate, 20)
		resultCh := make(chan *CompilationOutput, 1)

		config := CompilationConfig{
			Filepath:             path,
			Mode:                 "model",
			Input:                "fault",
			Output:               "text",
			LargeSMTLineOverride: 1,
		}

		go func() {
			r := NewRunner(config, progressCh)
			result := r.Run()
			resultCh <- result
			close(progressCh)
		}()

		var confirmCh chan bool
		for update := range progressCh {
			if update.Phase == PhaseConfirmLargeSMT {
				confirmCh = update.ConfirmCh
				break
			}
		}

		if confirmCh == nil {
			t.Fatal("never received PhaseConfirmLargeSMT update")
		}

		// Abort: runner should return early.
		confirmCh <- false

		for range progressCh {
		}
		result := <-resultCh
		if result.Error != nil {
			t.Fatalf("unexpected error after abort: %v", result.Error)
		}
		if result.LargeSMTLines == 0 {
			t.Fatal("expected LargeSMTLines > 0 after abort")
		}
		if result.ResultLog != nil {
			t.Fatal("expected no ResultLog after abort")
		}
	})
}
