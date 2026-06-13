package runner

import (
	"fault/ast"
	"fault/execute"
	"fault/generator"
	"fault/generator/scenario"
	"fault/listener"
	"fault/llvm"
	"fault/preprocess"
	"fault/reachability"
	"fault/swaps"
	"fault/types"
	"fault/util"
	"fmt"
	"os"
	gopath "path"
	"strings"
)

type ProgressPhase int

const (
	PhaseParsing ProgressPhase = iota
	PhasePreprocessing
	PhaseTypeChecking
	PhaseLLVM
	PhaseSMT
	PhaseModelChecking
	PhaseResults
	PhaseConfirmLargeSMT // paused: waiting for user to confirm large SMT
)

// LargeSMTThreshold is the number of SMT lines above which Fault pauses
// and asks the user whether to proceed with model checking.
const LargeSMTThreshold = 10_000

type ProgressUpdate struct {
	Phase     ProgressPhase
	Status    string  // "Parsing AST...", "Type checking complete", etc.
	Percent   float64 // 0.0 to 1.0
	Done      bool
	Error     error
	SMTLines  int       // non-zero when Phase == PhaseConfirmLargeSMT
	ConfirmCh chan bool  // send true to proceed, false to abort (Phase == PhaseConfirmLargeSMT)
}

type CompilationConfig struct {
	Filepath             string
	Mode                 string // ast, ir, smt, template, model
	Input                string // fault, ll, smt2
	Output               string // text, smt
	Reach                bool
	LargeSMTLineOverride int // if > 0, overrides the default LargeSMTThreshold constant
	SMTTimeout           int // (set-option :timeout N) in milliseconds; 0 = no limit
	SMTMemoryMaxSize     int // (set-option :memory_max_size N) in MB; 0 = no limit
}

// PendingModelCheck holds everything needed to resume model checking after
// the user confirms they want to proceed past a large-SMT warning.
type PendingModelCheck struct {
	SMT        string
	Uncertains map[string][]float64
	Unknowns   []string
	Asserts    []*ast.AssertionStatement
	HasSynth   bool // true when the run block contains synthesis slots (__)
	ResultLog  *scenario.Logger
}

type CompilationOutput struct {
	ResultLog     *scenario.Logger
	Asserts       []*ast.AssertionStatement
	Warnings      []string
	Message       string
	SMT           string
	ParamManifest map[string]string // non-nil when Mode == "template"
	AST           *ast.Spec
	IR            string
	Error         error
	ErrorPhase    ProgressPhase
	LargeSMTLines int                // non-zero: run paused before solving due to large SMT
	Pending       *PendingModelCheck // non-nil when LargeSMTLines > 0 (CLI resume path)
}

type Runner struct {
	config   CompilationConfig
	progress chan ProgressUpdate
}

func NewRunner(config CompilationConfig, progress chan ProgressUpdate) *Runner {
	return &Runner{
		config:   config,
		progress: progress,
	}
}

func (r *Runner) sendProgress(phase ProgressPhase, status string, percent float64, done bool) {
	if r.progress != nil {
		r.progress <- ProgressUpdate{
			Phase:   phase,
			Status:  status,
			Percent: percent,
			Done:    done,
			Error:   nil,
		}
	}
}

func (r *Runner) sendError(phase ProgressPhase, err error) {
	if r.progress != nil {
		r.progress <- ProgressUpdate{
			Phase:  phase,
			Status: err.Error(),
			Error:  err,
		}
	}
}

func (r *Runner) parse(data string, path string, file string, filetype string, reach bool) (*ast.Spec, *listener.FaultListener, *types.Checker, map[string]string, error) {
	// Confirm that the filetype and file declaration match
	if !r.validateFiletype(data, filetype) {
		return nil, nil, nil, nil, fmt.Errorf("malformatted file: declaration does not match filetype")
	}

	r.sendProgress(PhaseParsing, "Parsing AST...", 0.0, false)

	flags := make(map[string]bool)
	flags["specType"] = (filetype == "fspec")
	flags["testing"] = false
	flags["skipRun"] = false
	lstnr, err := listener.Execute(data, path, flags)
	if err != nil {
		return nil, nil, nil, nil, err
	}

	r.sendProgress(PhaseParsing, "Parsing complete", 0.14, true)

	r.sendProgress(PhasePreprocessing, "Preprocessing...", 0.14, false)
	pre, err := preprocess.Execute(lstnr)
	if err != nil {
		return nil, nil, nil, nil, err
	}
	r.sendProgress(PhasePreprocessing, "Preprocessing complete", 0.28, true)

	r.sendProgress(PhaseTypeChecking, "Type checking...", 0.28, false)
	ty, err := types.Execute(pre.Processed, pre)
	if err != nil {
		return nil, nil, nil, nil, err
	}
	r.sendProgress(PhaseTypeChecking, "Type checking complete", 0.42, true)

	sw := swaps.NewPrecompiler(ty)
	tree := sw.Swap(ty.Checked)

	if reach {
		reacher := reachability.NewTracer()
		if err := reacher.Scan(ty.Checked); err != nil {
			return nil, nil, nil, nil, err
		}
	}

	return tree, lstnr, ty, sw.Alias, nil
}

func (r *Runner) skipCommentsNl(data string) string {
	i := 0
	for i < len(data) {
		// Skip whitespace and newlines
		if data[i] == ' ' || data[i] == '\t' || data[i] == '\n' || data[i] == '\r' {
			i++
			continue
		}

		// Handle single-line comments starting with // or #
		if i < len(data)-1 && (data[i:i+2] == "//" || data[i] == '#') {
			// Skip to end of line
			for i < len(data) && data[i] != '\n' {
				i++
			}
			continue
		}

		// Handle multi-line comments starting with /*
		if i < len(data)-1 && data[i:i+2] == "/*" {
			i += 2 // Skip the /*
			// Skip until we find */
			for i < len(data)-1 {
				if data[i:i+2] == "*/" {
					i += 2 // Skip the */
					break
				}
				i++
			}
			continue
		}

		// If we reach here, we've found the first non-comment, non-whitespace character
		break
	}

	return data[i:]
}

func (r *Runner) validateFiletype(data string, filetype string) bool {
	data = r.skipCommentsNl(data)
	if filetype == "fspec" && len(data) >= 4 && data[0:4] == "spec" {
		return true
	}
	if filetype == "fsystem" && len(data) >= 6 && data[0:6] == "system" {
		return true
	}
	return false
}

func (r *Runner) smt2(ir string, compiler *llvm.Compiler) *generator.Generator {
	g := generator.Execute(compiler, generator.GeneratorOptions{
		Timeout:       r.config.SMTTimeout,
		MemoryMaxSize: r.config.SMTMemoryMaxSize,
	})
	return g
}

func (r *Runner) plainSolve(smt string) (string, error) {
	ex, err := execute.NewModelChecker()
	if err != nil {
		return "", err
	}
	ex.LoadModel(smt, nil, nil)
	ok, err := ex.Check()
	if err != nil {
		return "", fmt.Errorf("model checker has failed: %s", err)
	}
	if !ok {
		return "Fault could not find a failure case.", nil
	}
	scenario, err := ex.PlainSolve()
	if err != nil {
		return "", fmt.Errorf("error found fetching solution from solver: %s", err)
	}
	return scenario, nil
}

func (r *Runner) probability(smt string, uncertains map[string][]float64, unknowns []string) (*execute.ModelChecker, error) {
	ex, err := execute.NewModelChecker()
	if err != nil {
		return nil, err
	}
	ex.LoadModel(smt, uncertains, unknowns)
	ok, err := ex.Check()
	if err != nil {
		return nil, fmt.Errorf("model checker has failed: %s", err)
	}
	if !ok {
		ex.NoSat = true
		return ex, nil
	}
	err = ex.Solve()
	if err != nil {
		return nil, fmt.Errorf("error found fetching solution from solver: %s", err)
	}
	return ex, nil
}

// synthProbability solves a synthesis problem using iterative deepening to find
// the shortest valid operation sequence. It tries k=1,2,...,n non-noop steps and
// returns the first satisfying result, ensuring the minimal solution is returned.
func (r *Runner) synthProbability(smt string, uncertains map[string][]float64, unknowns []string) (*execute.ModelChecker, error) {
	noopSelectors := findNoopSelectors(smt)
	n := len(noopSelectors)
	if n == 0 {
		return r.probability(smt, uncertains, unknowns)
	}

	ex, err := execute.NewModelChecker()
	if err != nil {
		return nil, err
	}
	ex.LoadModel(smt, uncertains, unknowns)

	for k := 1; k <= n; k++ {
		extra := buildMinStepAsserts(noopSelectors, n-k)
		ok, err := ex.CheckWithAsserts(extra)
		if err != nil {
			return nil, fmt.Errorf("model checker has failed: %s", err)
		}
		if ok {
			err = ex.SolveWithAsserts(extra)
			if err != nil {
				return nil, fmt.Errorf("error found fetching solution from solver: %s", err)
			}
			return ex, nil
		}
	}

	ex.NoSat = true
	return ex, nil
}

// findNoopSelectors scans the SMT formula for noop candidate selector variable
// declarations. These are Bool variables named synth_N___noop___N that the
// synthesis unpacker emits for each step's noop candidate.
func findNoopSelectors(smt string) []string {
	var selectors []string
	for _, line := range strings.Split(smt, "\n") {
		trimmed := strings.TrimSpace(line)
		if strings.HasPrefix(trimmed, "(declare-fun") && strings.Contains(trimmed, "___noop___") {
			parts := strings.Fields(trimmed)
			if len(parts) >= 2 {
				selectors = append(selectors, parts[1])
			}
		}
	}
	return selectors
}

// buildMinStepAsserts returns an SMT assertion requiring at least minNoops of the
// given noop selectors to be true, which limits the solver to at most
// len(noopSelectors)-minNoops non-noop steps.
func buildMinStepAsserts(noopSelectors []string, minNoops int) []string {
	if minNoops <= 0 {
		return nil
	}
	var terms []string
	for _, sel := range noopSelectors {
		terms = append(terms, fmt.Sprintf("(ite %s 1 0)", sel))
	}
	var sum string
	if len(terms) == 1 {
		sum = terms[0]
	} else {
		sum = "(+ " + strings.Join(terms, " ") + ")"
	}
	return []string{fmt.Sprintf("(assert (>= %s %d))", sum, minNoops)}
}

func (r *Runner) Run() *CompilationOutput {
	output := &CompilationOutput{}

	filetype := util.DetectMode(r.config.Filepath)
	if filetype == "" {
		err := fmt.Errorf("file provided is not a .fspec or .fsystem file")
		r.sendError(PhaseParsing, err)
		output.Error = err
		output.ErrorPhase = PhaseParsing
		return output
	}

	filepath := util.Filepath(r.config.Filepath)
	uncertains := make(map[string][]float64)
	unknowns := []string{}

	data, err := os.ReadFile(filepath)
	if err != nil {
		r.sendError(PhaseParsing, err)
		output.Error = err
		output.ErrorPhase = PhaseParsing
		return output
	}
	d := string(data)
	path := gopath.Dir(filepath)

	switch r.config.Input {
	case "fault", "fspec":
		tree, lstnr, ty, alias, err := r.parse(d, path, filepath, filetype, r.config.Reach)
		if err != nil {
			r.sendError(PhaseParsing, err)
			output.Error = err
			output.ErrorPhase = PhaseParsing
			return output
		}
		if lstnr == nil {
			err := fmt.Errorf("Fault parser returned nil")
			r.sendError(PhaseParsing, err)
			output.Error = err
			output.ErrorPhase = PhaseParsing
			return output
		}

		if r.config.Mode == "ast" {
			output.AST = lstnr.AST
			return output
		}

		r.sendProgress(PhaseLLVM, "Generating LLVM IR...", 0.42, false)
		compiler, err := llvm.Execute(tree, ty.SpecStructs, lstnr.Uncertains, lstnr.Unknowns, lstnr.Wholes, lstnr.Params, alias, false)
		if err != nil {
			r.sendError(PhaseLLVM, err)
			output.Error = err
			output.ErrorPhase = PhaseLLVM
			return output
		}
		uncertains = compiler.RawInputs.Uncertains
		unknowns = compiler.RawInputs.Unknowns
		r.sendProgress(PhaseLLVM, "LLVM IR generated", 0.56, true)

		if r.config.Mode == "ir" {
			output.IR = compiler.GetIR()
			return output
		}

		if len(compiler.RawInputs.Params) > 0 && r.config.Mode != "template" && r.config.Mode != "ir" {
			err := fmt.Errorf("spec contains param() fields but mode is %q — param() fields produce __PARAM_...__ placeholder tokens that Z3 cannot parse; use --mode=template to generate a substitutable SMT template", r.config.Mode)
			r.sendError(PhaseSMT, err)
			output.Error = err
			output.ErrorPhase = PhaseSMT
			return output
		}

		r.sendProgress(PhaseSMT, "Generating SMT constraints...", 0.56, false)
		g := generator.Execute(compiler, generator.GeneratorOptions{
			Timeout:       r.config.SMTTimeout,
			MemoryMaxSize: r.config.SMTMemoryMaxSize,
		})
		r.sendProgress(PhaseSMT, "SMT generation complete", 0.70, true)

		if r.config.Mode == "smt" {
			output.SMT = g.SMT()
			return output
		}

		if r.config.Mode == "template" {
			output.SMT = g.SMT()
			output.ParamManifest = g.ParamManifest()
			return output
		}

		smt := g.SMT()
		smtLines := strings.Count(smt, "\n") + 1
		threshold := LargeSMTThreshold
		if r.config.LargeSMTLineOverride > 0 {
			threshold = r.config.LargeSMTLineOverride
		}
		if smtLines > threshold {
			if r.progress != nil {
				// TUI path: block until user confirms or cancels.
				confirmCh := make(chan bool, 1)
				r.progress <- ProgressUpdate{
					Phase:     PhaseConfirmLargeSMT,
					Status:    fmt.Sprintf("SMT formula is %d lines", smtLines),
					SMTLines:  smtLines,
					ConfirmCh: confirmCh,
				}
				if !<-confirmCh {
					output.LargeSMTLines = smtLines
					return output
				}
			} else {
				// CLI path: return early; caller prompts and calls Resume().
				output.LargeSMTLines = smtLines
				output.SMT = smt
				output.Pending = &PendingModelCheck{
					SMT:        smt,
					Uncertains: uncertains,
					Unknowns:   unknowns,
					Asserts:    compiler.RawInputs.Asserts,
					HasSynth:   hasSolvableSteps(tree),
					ResultLog:  g.ResultLog,
				}
				return output
			}
		}

		if r.config.Output == "smt" {
			r.sendProgress(PhaseModelChecking, "Running model checker...", 0.70, false)
			scenario, err := r.plainSolve(smt)
			if err != nil {
				r.sendError(PhaseModelChecking, err)
				output.Error = err
				output.ErrorPhase = PhaseModelChecking
				return output
			}
			r.sendProgress(PhaseModelChecking, "Model checking complete", 0.85, true)
			output.SMT = scenario
			return output
		}

		// Default output is "text" with full model checking
		r.sendProgress(PhaseModelChecking, "Running model checker...", 0.70, false)
		r.sendProgress(PhaseModelChecking, "Checking satisfiability...", 0.75, false)
		var mc *execute.ModelChecker
		if hasSolvableSteps(tree) {
			mc, err = r.synthProbability(smt, uncertains, unknowns)
		} else {
			mc, err = r.probability(smt, uncertains, unknowns)
		}
		if err != nil {
			r.sendError(PhaseModelChecking, err)
			output.Error = err
			output.ErrorPhase = PhaseModelChecking
			return output
		}
		r.sendProgress(PhaseModelChecking, "Model checking complete", 0.85, true)
		r.sendProgress(PhaseResults, "Processing results...", 0.85, false)
		if mc.NoSat {
			msg := noSatMessage(compiler.RawInputs.Asserts, hasSolvableSteps(tree))
			r.sendProgress(PhaseResults, msg, 1.0, true)
			output.Message = msg
			return output
		}
		mc.EvaluateViolations(compiler.RawInputs.Asserts)
		output.Asserts = compiler.RawInputs.Asserts
		g.ResultLog.SystemName = systemName(tree)
		g.ResultLog.Results = mc.ResultValues
		g.ResultLog.Trace()
		g.ResultLog.Validate()
		g.ResultLog.Kill()
		r.sendProgress(PhaseResults, "Results ready", 1.0, true)

		output.ResultLog = g.ResultLog

	case "ll":
		r.sendProgress(PhaseLLVM, "Loading LLVM IR...", 0.42, false)
		compiler := llvm.NewCompiler()
		compiler.RawInputs.Uncertains = uncertains
		compiler.RawInputs.Unknowns = unknowns
		r.sendProgress(PhaseLLVM, "LLVM IR loaded", 0.56, true)

		r.sendProgress(PhaseSMT, "Generating SMT constraints...", 0.56, false)
		g := r.smt2(d, compiler)
		r.sendProgress(PhaseSMT, "SMT generation complete", 0.70, true)

		if r.config.Mode == "smt" {
			output.SMT = g.SMT()
			return output
		}

		if r.config.Output == "smt" {
			r.sendProgress(PhaseModelChecking, "Running model checker...", 0.70, false)
			scenario, err := r.plainSolve(g.SMT())
			if err != nil {
				r.sendError(PhaseModelChecking, err)
				output.Error = err
				output.ErrorPhase = PhaseModelChecking
				return output
			}
			r.sendProgress(PhaseModelChecking, "Model checking complete", 0.85, true)
			output.SMT = scenario
			return output
		}

	case "smt2":
		r.sendProgress(PhaseSMT, "Loading SMT2 file...", 0.56, false)
		r.sendProgress(PhaseSMT, "SMT2 loaded", 0.70, true)

		if r.config.Output == "smt" {
			r.sendProgress(PhaseModelChecking, "Running model checker...", 0.70, false)
			scenario, err := r.plainSolve(d)
			if err != nil {
				r.sendError(PhaseModelChecking, err)
				output.Error = err
				return output
			}
			r.sendProgress(PhaseModelChecking, "Model checking complete", 0.85, true)
			output.SMT = scenario
			return output
		}
	}

	return output
}

// Resume completes model checking after the user confirms they want to proceed
// past a large-SMT warning. It is called by the CLI after prompting the user;
// the TUI uses the ConfirmCh approach instead (the runner goroutine unblocks).
func (r *Runner) Resume(pending *PendingModelCheck) *CompilationOutput {
	output := &CompilationOutput{}

	r.sendProgress(PhaseModelChecking, "Running model checker...", 0.70, false)
	r.sendProgress(PhaseModelChecking, "Checking satisfiability...", 0.75, false)
	var mc *execute.ModelChecker
	var err error
	if pending.HasSynth {
		mc, err = r.synthProbability(pending.SMT, pending.Uncertains, pending.Unknowns)
	} else {
		mc, err = r.probability(pending.SMT, pending.Uncertains, pending.Unknowns)
	}
	if err != nil {
		r.sendError(PhaseModelChecking, err)
		output.Error = err
		output.ErrorPhase = PhaseModelChecking
		return output
	}
	r.sendProgress(PhaseModelChecking, "Model checking complete", 0.85, true)
	r.sendProgress(PhaseResults, "Processing results...", 0.85, false)
	if mc.NoSat {
		msg := noSatMessage(pending.Asserts, pending.HasSynth)
		r.sendProgress(PhaseResults, msg, 1.0, true)
		output.Message = msg
		return output
	}
	mc.EvaluateViolations(pending.Asserts)
	output.Asserts = pending.Asserts
	pending.ResultLog.Results = mc.ResultValues
	pending.ResultLog.Trace()
	pending.ResultLog.Validate()
	pending.ResultLog.Kill()
	r.sendProgress(PhaseResults, "Results ready", 1.0, true)
	output.ResultLog = pending.ResultLog
	return output
}

// systemName extracts the spec or system name from the top-level AST declaration.
func systemName(tree *ast.Spec) string {
	for _, stmt := range tree.Statements {
		switch s := stmt.(type) {
		case *ast.SpecDeclStatement:
			return s.Name.Value
		case *ast.SysDeclStatement:
			return s.Name.Value
		}
	}
	return ""
}

// hasSolvableSteps reports whether the spec's run block contains any synthesis
// slots (__), indicating the user wants program synthesis rather than verification.
func hasSolvableSteps(tree *ast.Spec) bool {
	for _, stmt := range tree.Statements {
		rs, ok := stmt.(*ast.RunStatement)
		if !ok {
			continue
		}
		for _, step := range rs.Steps {
			if _, ok := step.(*ast.SolvableStep); ok {
				return true
			}
		}
	}
	return false
}

// noSatMessage returns the appropriate user-facing message when the solver
// returns unsat, based on whether the spec is in verification, synthesis, or
// simulation mode.
//
//   - Verification (has assert statements): unsat means no assertion is ever
//     violated — the model is correct.
//   - Synthesis (has __ slots, no asserts): unsat means no combination of
//     operations can satisfy all the assume constraints.
//   - Simulation (only assumes, no __ and no asserts): unsat means the assume
//     constraints are mutually contradictory.
func noSatMessage(asserts []*ast.AssertionStatement, hasSynth bool) string {
	if hasSynth {
		return "Fault could not find a valid operation sequence that satisfies all constraints."
	}
	if len(asserts) > 0 {
		return "Fault could not find a failure case. All good!"
	}
	return "Fault could not find a satisfying model. The assume constraints may be contradictory."
}
