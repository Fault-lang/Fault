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
)

type ProgressUpdate struct {
	Phase   ProgressPhase
	Status  string  // "Parsing AST...", "Type checking complete", etc.
	Percent float64 // 0.0 to 1.0
	Done    bool
	Error   error
}

type CompilationConfig struct {
	Filepath string
	Mode     string // ast, ir, smt, check
	Input    string // fspec, ll, smt2
	Output   string // log, smt, static, legacy, visualize
	Reach    bool
}

type CompilationOutput struct {
	ResultLog *scenario.Logger
	SMT       string
	AST       *ast.Spec
	IR        string
	Error     error
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
	lstnr := listener.Execute(data, path, flags)

	r.sendProgress(PhaseParsing, "Parsing complete", 0.14, true)

	r.sendProgress(PhasePreprocessing, "Preprocessing...", 0.14, false)
	pre := preprocess.Execute(lstnr)
	r.sendProgress(PhasePreprocessing, "Preprocessing complete", 0.28, true)

	r.sendProgress(PhaseTypeChecking, "Type checking...", 0.28, false)
	ty := types.Execute(pre.Processed, pre)
	r.sendProgress(PhaseTypeChecking, "Type checking complete", 0.42, true)

	sw := swaps.NewPrecompiler(ty)
	tree := sw.Swap(ty.Checked)

	if reach {
		reacher := reachability.NewTracer()
		reacher.Scan(ty.Checked)
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
	g := generator.Execute(compiler)
	return g
}

func (r *Runner) plainSolve(smt string) (string, error) {
	ex := execute.NewModelChecker()
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
	ex := execute.NewModelChecker()
	ex.LoadModel(smt, uncertains, unknowns)
	ok, err := ex.Check()
	if err != nil {
		return nil, fmt.Errorf("model checker has failed: %s", err)
	}
	if !ok {
		return nil, fmt.Errorf("Fault could not find a failure case")
	}
	err = ex.Solve()
	if err != nil {
		return nil, fmt.Errorf("error found fetching solution from solver: %s", err)
	}
	return ex, nil
}

func (r *Runner) Run() *CompilationOutput {
	output := &CompilationOutput{}

	filetype := util.DetectMode(r.config.Filepath)
	if filetype == "" {
		err := fmt.Errorf("file provided is not a .fspec or .fsystem file")
		r.sendError(PhaseParsing, err)
		output.Error = err
		return output
	}

	filepath := util.Filepath(r.config.Filepath)
	uncertains := make(map[string][]float64)
	unknowns := []string{}

	data, err := os.ReadFile(filepath)
	if err != nil {
		r.sendError(PhaseParsing, err)
		output.Error = err
		return output
	}
	d := string(data)
	path := gopath.Dir(filepath)

	switch r.config.Input {
	case "fspec":
		tree, lstnr, ty, alias, err := r.parse(d, path, filepath, filetype, r.config.Reach)
		if err != nil {
			r.sendError(PhaseParsing, err)
			output.Error = err
			return output
		}
		if lstnr == nil {
			err := fmt.Errorf("Fault parser returned nil")
			r.sendError(PhaseParsing, err)
			output.Error = err
			return output
		}

		if r.config.Mode == "ast" {
			output.AST = lstnr.AST
			return output
		}

		r.sendProgress(PhaseLLVM, "Generating LLVM IR...", 0.42, false)
		compiler := llvm.Execute(tree, ty.SpecStructs, lstnr.Uncertains, lstnr.Unknowns, alias, false)
		uncertains = compiler.RawInputs.Uncertains
		unknowns = compiler.RawInputs.Unknowns
		r.sendProgress(PhaseLLVM, "LLVM IR generated", 0.56, true)

		if r.config.Mode == "ir" {
			output.IR = compiler.GetIR()
			return output
		}

		if !compiler.IsValid {
			err := fmt.Errorf("Fault found nothing to run. Missing run block or start block")
			r.sendError(PhaseLLVM, err)
			output.Error = err
			return output
		}

		r.sendProgress(PhaseSMT, "Generating SMT constraints...", 0.56, false)
		g := generator.Execute(compiler)
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
				return output
			}
			r.sendProgress(PhaseModelChecking, "Model checking complete", 0.85, true)
			output.SMT = scenario
			return output
		}

		r.sendProgress(PhaseModelChecking, "Running model checker...", 0.70, false)
		r.sendProgress(PhaseModelChecking, "Checking satisfiability...", 0.75, false)
		mc, err := r.probability(g.SMT(), uncertains, unknowns)
		if err != nil {
			r.sendError(PhaseModelChecking, err)
			output.Error = err
			return output
		}
		r.sendProgress(PhaseModelChecking, "Model checking complete", 0.85, true)

		r.sendProgress(PhaseResults, "Processing results...", 0.85, false)
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
