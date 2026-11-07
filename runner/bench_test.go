package runner

import (
	"os"
	"path/filepath"
	"testing"
)

// writeTempSpecB is writeTempSpec adapted for benchmarks.
func writeTempSpecB(b *testing.B, content string) string {
	b.Helper()
	dir := b.TempDir()
	path := filepath.Join(dir, "test.fspec")
	if err := os.WriteFile(path, []byte(content), 0600); err != nil {
		b.Fatalf("could not write temp spec: %v", err)
	}
	return path
}

// copyTestdata copies a file from generator/testdata into a temp dir so the
// runner can resolve relative imports correctly.
func copyTestdata(b *testing.B, relPath string) string {
	b.Helper()
	src := filepath.Join("..", "generator", "testdata", relPath)
	data, err := os.ReadFile(src)
	if err != nil {
		b.Fatalf("could not read testdata %s: %v", relPath, err)
	}
	dir := b.TempDir()
	dst := filepath.Join(dir, filepath.Base(relPath))
	if err := os.WriteFile(dst, data, 0600); err != nil {
		b.Fatalf("could not write testdata copy: %v", err)
	}
	return dst
}

// runTo runs the compiler pipeline up to the given mode and reports any error.
func runTo(b *testing.B, path, mode string) {
	b.Helper()
	r := NewRunner(CompilationConfig{
		Filepath: path,
		Mode:     mode,
		Input:    "fault",
		Output:   "text",
	}, nil)
	result := r.Run()
	if result.Error != nil {
		b.Fatalf("mode=%s: %v", mode, result.Error)
	}
}

// ----------------------------------------------------------------------------
// Fixtures
// ----------------------------------------------------------------------------

// minimal: one stock, one flow, one assert. Smallest meaningful spec.
const minimalSpec = `spec minimal;
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

// branching: adds a conditional inside the flow function.
const branchingSpec = `spec branching;
def s = stock{
	value: 30,
};
def f = flow{
	vault: new s,
	fn: func{
		if vault.value > 4 {
			vault.value <- vault.value - 2;
		}
	},
};
run init { l = new f; } {
	l.fn;
	l.fn;
	l.fn;
};
`

// parallel: two flows operating concurrently via | in the run block.
const parallelSpec = `spec parallel;
def tub = stock{
	level: 5,
};
def faucet = flow{
	water: new tub,
	in: func{
		water.level <- 10;
	},
};
def drain = flow{
	water: new tub,
	out: func{
		water.level -> 20;
	},
};
run init { drawn = new faucet; pipe = new drain; } {
	drawn.in | pipe.out;
	drawn.in | pipe.out;
	drawn.in | pipe.out;
	drawn.in | pipe.out;
};
`

// unknowns: free variables + assume/assert, exercises QF_NRA path.
const unknownsSpec = `spec unknowns;
def s = stock{
	a,
	b: 2,
	c: 0,
};
def f = flow{
	data: new s,
	fn: func{
		data.c <- data.a + data.b;
	},
};
assume s.a > 5;
assert s.a <= 6;
run init { loop = new f; } {
	loop.fn;
	loop.fn;
	loop.fn;
};
`

// ----------------------------------------------------------------------------
// Per-stage benchmarks: measure each pipeline stage in isolation.
// Each benchmark stops at the named checkpoint so cost is cumulative up to
// that stage, not just the marginal cost of the stage itself.
// ----------------------------------------------------------------------------

func BenchmarkStage_AST_Minimal(b *testing.B) {
	for b.Loop() {
		path := writeTempSpecB(b, minimalSpec)
		runTo(b, path, "ast")
	}
}

func BenchmarkStage_IR_Minimal(b *testing.B) {
	for b.Loop() {
		path := writeTempSpecB(b, minimalSpec)
		runTo(b, path, "ir")
	}
}

func BenchmarkStage_SMT_Minimal(b *testing.B) {
	for b.Loop() {
		path := writeTempSpecB(b, minimalSpec)
		runTo(b, path, "smt")
	}
}

func BenchmarkStage_AST_Branching(b *testing.B) {
	for b.Loop() {
		path := writeTempSpecB(b, branchingSpec)
		runTo(b, path, "ast")
	}
}

func BenchmarkStage_IR_Branching(b *testing.B) {
	for b.Loop() {
		path := writeTempSpecB(b, branchingSpec)
		runTo(b, path, "ir")
	}
}

func BenchmarkStage_SMT_Branching(b *testing.B) {
	for b.Loop() {
		path := writeTempSpecB(b, branchingSpec)
		runTo(b, path, "smt")
	}
}

func BenchmarkStage_AST_Parallel(b *testing.B) {
	for b.Loop() {
		path := writeTempSpecB(b, parallelSpec)
		runTo(b, path, "ast")
	}
}

func BenchmarkStage_IR_Parallel(b *testing.B) {
	for b.Loop() {
		path := writeTempSpecB(b, parallelSpec)
		runTo(b, path, "ir")
	}
}

func BenchmarkStage_SMT_Parallel(b *testing.B) {
	for b.Loop() {
		path := writeTempSpecB(b, parallelSpec)
		runTo(b, path, "smt")
	}
}

func BenchmarkStage_AST_Unknowns(b *testing.B) {
	for b.Loop() {
		path := writeTempSpecB(b, unknownsSpec)
		runTo(b, path, "ast")
	}
}

func BenchmarkStage_IR_Unknowns(b *testing.B) {
	for b.Loop() {
		path := writeTempSpecB(b, unknownsSpec)
		runTo(b, path, "ir")
	}
}

func BenchmarkStage_SMT_Unknowns(b *testing.B) {
	for b.Loop() {
		path := writeTempSpecB(b, unknownsSpec)
		runTo(b, path, "smt")
	}
}
