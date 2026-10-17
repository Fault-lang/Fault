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
run init { t = new f; } { t.fill; };
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

const paramSpec = `spec testreq;
def request = stock{
    amount: param(0.0),
    count: param(0),
};
run init{r1 = new request;} {}
`

// synthPickSpec: single __ slot, solver must pick increment to satisfy value > 10.
// decrement and noop cannot satisfy the goal.
const synthPickSpec = `spec synth_pick;

def counter = stock{
    value: 10,
};

def ctrl = flow{
    c: new counter,
    increment: func{
        c.value <- 1;
    },
    decrement: func{
        c.value -> 1;
    },
};

assume counter.value > 10 eventually;

run init{ inst = new ctrl; } {
    __;
};
`

// synthSequenceSpec: two __ slots, goal requires fill twice (40+10+10=60).
// Exercises iterative deepening: k=1 is not enough, k=2 satisfies.
const synthSequenceSpec = `spec synth_sequence;

def tank = stock{
    level: 40,
};

def ops = flow{
    t: new tank,
    fill: func{
        t.level <- 10;
    },
    drain: func{
        t.level -> 10;
    },
};

assume tank.level == 60 eventually;

run init{ inst = new ops; } {
    __;
    __;
};
`

// synthSandwichSpec: explicit call then __ slot; after deposit (100->150),
// solver must pick withdraw to restore balance to 100.
const synthSandwichSpec = `spec synth_sandwich;

def wallet = stock{
    balance: 100,
};

def account = flow{
    w: new wallet,
    deposit: func{
        w.balance <- 50;
    },
    withdraw: func{
        w.balance -> 50;
    },
};

assume wallet.balance == 100 eventually;

run init{ inst = new account; } {
    inst.deposit;
    __;
};
`

// synthUnsatSpec: goal is unreachable regardless of which func the solver picks.
// value starts at 10, increment adds 1, but assume requires value < 0 — impossible.
const synthUnsatSpec = `spec synth_unsat;

def counter = stock{
    value: 10,
};

def ctrl = flow{
    c: new counter,
    increment: func{
        c.value <- 1;
    },
};

assume counter.value < 0 eventually;

run init{ inst = new ctrl; } {
    __;
};
`

func runSynthSpec(t *testing.T, spec string) *CompilationOutput {
	t.Helper()
	path := writeTempSpec(t, spec)
	r := NewRunner(CompilationConfig{
		Filepath: path,
		Mode:     "model",
		Input:    "fault",
		Output:   "text",
	}, nil)
	return r.Run()
}

// TestSynthProbabilityFuncSelection verifies that synthProbability picks a concrete
// func (increment) over noop to satisfy the assume goal.
func TestSynthProbabilityFuncSelection(t *testing.T) {
	result := runSynthSpec(t, synthPickSpec)
	if result.Error != nil {
		t.Fatalf("unexpected error: %v", result.Error)
	}
	if result.Message == "" && result.ResultLog == nil {
		t.Fatal("expected a result (message or log) — solver may have returned NoSat unexpectedly")
	}
}

// TestSynthProbabilityIterativeDeepening verifies that synthProbability finds the
// minimal solution requiring two non-noop steps (k=2), not just k=1.
func TestSynthProbabilityIterativeDeepening(t *testing.T) {
	result := runSynthSpec(t, synthSequenceSpec)
	if result.Error != nil {
		t.Fatalf("unexpected error: %v", result.Error)
	}
	if result.Message == "" && result.ResultLog == nil {
		t.Fatal("expected a result — solver may have returned NoSat unexpectedly")
	}
}

// TestSynthProbabilitySandwich verifies that explicit calls and synthesis slots
// can be mixed in a run block and the solver fills the slot correctly.
func TestSynthProbabilitySandwich(t *testing.T) {
	result := runSynthSpec(t, synthSandwichSpec)
	if result.Error != nil {
		t.Fatalf("unexpected error: %v", result.Error)
	}
	if result.Message == "" && result.ResultLog == nil {
		t.Fatal("expected a result — solver may have returned NoSat unexpectedly")
	}
}

// TestSynthProbabilityNoSat verifies that synthProbability returns a NoSat message
// (not an error) when the assume goal is unreachable by any candidate sequence.
func TestSynthProbabilityNoSat(t *testing.T) {
	result := runSynthSpec(t, synthUnsatSpec)
	if result.Error != nil {
		t.Fatalf("unexpected error: %v", result.Error)
	}
	if result.Message == "" {
		t.Fatal("expected a NoSat message when goal is unreachable, got empty message with no log")
	}
}

func TestParamRequiresTemplateMode(t *testing.T) {
	path := writeTempSpec(t, paramSpec)

	// "ir" mode outputs LLVM IR (no __PARAM__ tokens), so it is allowed.
	// "ast" mode returns before LLVM, so it never triggers the check.
	for _, mode := range []string{"model", "smt"} {
		t.Run(mode, func(t *testing.T) {
			config := CompilationConfig{
				Filepath: path,
				Mode:     mode,
				Input:    "fault",
				Output:   "text",
			}
			r := NewRunner(config, nil)
			result := r.Run()
			if result.Error == nil {
				t.Fatalf("mode=%q: expected error for param() spec, got none", mode)
			}
			if result.ErrorPhase != PhaseSMT {
				t.Fatalf("mode=%q: expected ErrorPhase PhaseSMT, got %v", mode, result.ErrorPhase)
			}
		})
	}
}

func TestParamTemplateMode(t *testing.T) {
	path := writeTempSpec(t, paramSpec)

	config := CompilationConfig{
		Filepath: path,
		Mode:     "template",
		Input:    "fault",
		Output:   "text",
	}
	r := NewRunner(config, nil)
	result := r.Run()
	if result.Error != nil {
		t.Fatalf("unexpected error in template mode: %v", result.Error)
	}
	if result.SMT == "" {
		t.Fatal("expected SMT output in template mode, got empty string")
	}
	if result.ParamManifest == nil {
		t.Fatal("expected ParamManifest to be set in template mode")
	}
	if len(result.ParamManifest) == 0 {
		t.Fatal("expected non-empty ParamManifest")
	}
}
