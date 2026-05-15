package generator

import (
	"fault/ast"
	"fault/generator/asserts"
	"fault/generator/rules"
	"fault/generator/scenario"
	"fault/generator/unpack"
	"fault/generator/unroll"
	"fault/llvm"
	"fault/util"
	"fmt"
	"strings"

	"github.com/llir/llvm/asm"
	"github.com/llir/llvm/ir"
)

const (
	DefaultSMTTimeout       = 30_000 // milliseconds
	DefaultSMTMemoryMaxSize = 1_096  // MB
)

// GeneratorOptions controls optional SMT-LIB2 solver directives that are
// prepended to the generated formula.
type GeneratorOptions struct {
	Timeout       int // (set-option :timeout N) in milliseconds; 0 = omit
	MemoryMaxSize int // (set-option :memory_max_size N) in MB; 0 = omit
}

// Take LL IR and Generate SMTLib2
//
// Step 1: Translate LL IR into Rules by LLBlock
// and Function
//
// Step 2: Add Phi values necessary for the SMT model
// and flatten to a single set of rules

type Generator struct {
	constants   []rules.Rule
	functions   map[string]*ir.Func
	Env         *unroll.Env
	RawInputs   *llvm.RawInputs
	RunBlock    *unroll.LLFunc
	smt         []string
	ResultLog   *scenario.Logger
	StringRules map[string]string
	IsCompound  map[string]bool
}

func NewGenerator(ri *llvm.RawInputs, sr map[string]string, is map[string]bool, opts GeneratorOptions) *Generator {
	var preamble []string
	if opts.Timeout > 0 {
		preamble = append(preamble, fmt.Sprintf("(set-option :timeout %d)", opts.Timeout))
	}
	if opts.MemoryMaxSize > 0 {
		preamble = append(preamble, fmt.Sprintf("(set-option :memory_max_size %d)", opts.MemoryMaxSize))
	}
	preamble = append(preamble, "(set-logic QF_NRA)")
	return &Generator{
		functions:   make(map[string]*ir.Func),
		Env:         unroll.NewEnv(ri),
		smt:         preamble,
		RawInputs:   ri,
		StringRules: sr,
		IsCompound:  is,
	}
}

func Execute(compiler *llvm.Compiler, opts GeneratorOptions) *Generator {
	generator := NewGenerator(compiler.RawInputs, compiler.StringRules, compiler.IsCompound, opts)
	generator.Run(compiler.GetOptimizedIR())
	return generator
}

func (g *Generator) AppendSMT(new_smt []string) {
	g.smt = append(g.smt, new_smt...)
}

func (g *Generator) Run(llopt string) {
	m, err := asm.ParseString("", llopt) //"" because ParseString has a path variable
	if err != nil {
		panic(err)
	}
	g.newCallgraph(m)

}

func (g *Generator) newCallgraph(m *ir.Module) {
	g.Env.MutableVars = unroll.FindMutableVars(m.Funcs)
	g.Env.UsedVars = unroll.FindUsedVars(m.Funcs)
	g.Env.WriteSets = unroll.FindWriteSets(m.Funcs)
	g.Env.StringRules = g.StringRules
	g.constants = unroll.NewConstants(g.Env, m.Globals, g.RawInputs)
	g.sortFuncs(m.Funcs)

	g.Env.WhensThens = unroll.WhenThen(g.Env.RawInputs.Asserts)

	g.RunBlock = unroll.NewLLFunc(g.Env, g.functions, g.functions["__run"])
	g.RunBlock.Unroll()

	p := unpack.NewUnpacker(g.RunBlock.Ident)
	p.LoadStringRules(g.StringRules, g.IsCompound)
	p.Log.Uncertains = g.RawInputs.Uncertains

	p.VarTypes = g.Env.VarTypes
	smt := p.Unpack(g.constants, g.RunBlock)
	g.AppendSMT(p.InitVars())
	g.AppendSMT(smt)

	g.ResultLog = p.Log

	assertSMT := g.ProcessAsserts(g.RawInputs.Asserts, g.Env.CurrentRound, p.Registry, p.Whens)
	g.AppendSMT(assertSMT)
	assumeSMT := g.ProcessAsserts(g.RawInputs.Assumes, g.Env.CurrentRound, p.Registry, p.Whens)
	g.AppendSMT(assumeSMT)

	unfuncSMT := g.ProcessUnfuncs(g.RawInputs.Unfuncs, g.Env.CurrentRound)
	g.AppendSMT(unfuncSMT)
}

func (g *Generator) ProcessAsserts(assertList []*ast.AssertionStatement, rounds int, registry map[string][][]string, whens map[string][]map[string]string) []string {
	var rules []string

	for _, as := range assertList {
		if !asserts.IsRelevant(g.Env.VarTypes, as.Constraint){ //If the assert is on a variable that is not used, drop the assert
			continue;
		}
		c := asserts.NewConstraint(as, rounds, registry, whens)
		rules = append(rules, c.Parse()...)
	}
	return rules
}

func (g *Generator) sortFuncs(funcs []*ir.Func) {
	//Iterate through all the function blocks and store them by
	// function call name.
	for _, f := range funcs {
		// Get function name.
		g.functions[util.FormatIdent(f.Ident())] = f
		continue
	}
}

// ProcessUnfuncs generates SMT constraints for unfunc states.
//
// Each field referenced in any unfunc's requires or emits gets an auto-generated
// shadow Bool variable named <field>_available_N. The numeric field itself is
// left free for the solver; only the shadow is constrained here.
//
// For each unfunc, it emits:
//   - shadow Bool declarations for every unique field across all unfuncs
//   - a declare-fun for the uninterpreted function
//   - per-step activation variables
//   - activation guards (active => requires_available_N)
//   - write effects on the shadow (active => field_available_N+1 = true)
//   - frame conditions on the shadow (not active => field_available_N+1 = field_available_N)
func (g *Generator) ProcessUnfuncs(unfuncs []*llvm.UnfuncInfo, rounds int) []string {
	if len(unfuncs) == 0 {
		return nil
	}
	// Bound the trace to at least the number of unfunc states so the solver
	// has enough steps to find a plan even when the run block has no explicit
	// rounds.
	steps := rounds
	if steps < len(unfuncs) {
		steps = len(unfuncs)
	}

	var smt []string

	// Declare _available shadow Bool variables for every unique field across all
	// unfuncs. Steps 0..steps inclusive because write effects reference n+1 where
	// n goes up to steps-1.
	for _, fieldBase := range collectAllUnfuncFields(unfuncs) {
		for n := 0; n <= steps; n++ {
			smt = append(smt, fmt.Sprintf("(declare-fun %s_available_%d () Bool)", fieldBase, n))
		}
	}

	for _, uf := range unfuncs {
		stateId := util.FormatIdent(uf.StateKey)
		smt = append(smt, fmt.Sprintf("(declare-fun %s (Bool) Bool)", stateId))

		for n := 0; n < steps; n++ {
			activeVar := fmt.Sprintf("%s_%d_active", stateId, n)
			smt = append(smt, fmt.Sprintf("(declare-fun %s () Bool)", activeVar))

			// Activation guard: if active, required fields must be available
			if uf.Requires != nil {
				reqSMT := unfuncExprToSMTAvail(uf.Requires, n)
				smt = append(smt, fmt.Sprintf("(assert (=> %s %s))", activeVar, reqSMT))
			}

			// Write effects and frame conditions on _available shadow variables
			if uf.Emits != nil {
				for _, fieldBase := range collectUnfuncFields(uf.Emits) {
					next := fmt.Sprintf("%s_available_%d", fieldBase, n+1)
					curr := fmt.Sprintf("%s_available_%d", fieldBase, n)
					smt = append(smt, fmt.Sprintf("(assert (=> %s (= %s true)))", activeVar, next))
					smt = append(smt, fmt.Sprintf("(assert (=> (not %s) (= %s %s)))", activeVar, next, curr))
				}
			}
		}
	}
	return smt
}

// unfuncVarBase returns the SMT base variable name for a ParameterCall,
// using ProcessedName when available, falling back to Spec+Value.
func unfuncVarBase(pc *ast.ParameterCall) string {
	if len(pc.ProcessedName) > 0 {
		return strings.Join(pc.ProcessedName, "_")
	}
	parts := append([]string{pc.Spec}, pc.Value...)
	return strings.Join(parts, "_")
}

// unfuncExprToSMT converts an unfunc requires/emits expression to an SMT
// string, versioning ParameterCall leaves with the given step index.
func unfuncExprToSMT(expr ast.Expression, step int) string {
	switch e := expr.(type) {
	case *ast.ParameterCall:
		return fmt.Sprintf("%s_%d", unfuncVarBase(e), step)
	case *ast.InfixExpression:
		left := unfuncExprToSMT(e.Left, step)
		right := unfuncExprToSMT(e.Right, step)
		switch e.Operator {
		case "&&":
			return fmt.Sprintf("(and %s %s)", left, right)
		case "||":
			return fmt.Sprintf("(or %s %s)", left, right)
		default:
			return fmt.Sprintf("(%s %s %s)", e.Operator, left, right)
		}
	case *ast.PrefixExpression:
		inner := unfuncExprToSMT(e.Right, step)
		return fmt.Sprintf("(not %s)", inner)
	default:
		return ""
	}
}

// unfuncExprToSMTAvail converts an unfunc expression to an SMT string
// referencing the _available shadow Bool variable (e.g. field_available_N).
// Used for requires guards so activation depends on availability, not value.
func unfuncExprToSMTAvail(expr ast.Expression, step int) string {
	switch e := expr.(type) {
	case *ast.ParameterCall:
		return fmt.Sprintf("%s_available_%d", unfuncVarBase(e), step)
	case *ast.InfixExpression:
		left := unfuncExprToSMTAvail(e.Left, step)
		right := unfuncExprToSMTAvail(e.Right, step)
		switch e.Operator {
		case "&&":
			return fmt.Sprintf("(and %s %s)", left, right)
		case "||":
			return fmt.Sprintf("(or %s %s)", left, right)
		default:
			return fmt.Sprintf("(%s %s %s)", e.Operator, left, right)
		}
	case *ast.PrefixExpression:
		inner := unfuncExprToSMTAvail(e.Right, step)
		return fmt.Sprintf("(not %s)", inner)
	default:
		return ""
	}
}

// collectUnfuncFields returns the base SMT variable names (without version
// suffix) for all ParameterCall leaves in an emits expression.
func collectUnfuncFields(expr ast.Expression) []string {
	switch e := expr.(type) {
	case *ast.ParameterCall:
		return []string{unfuncVarBase(e)}
	case *ast.InfixExpression:
		left := collectUnfuncFields(e.Left)
		right := collectUnfuncFields(e.Right)
		return append(left, right...)
	case *ast.PrefixExpression:
		return collectUnfuncFields(e.Right)
	default:
		return nil
	}
}

// collectAllUnfuncFields collects all unique field base names from the Requires
// and Emits expressions of every UnfuncInfo. Deduplication ensures each shadow
// variable is declared exactly once even if the same field appears in multiple
// unfuncs.
func collectAllUnfuncFields(unfuncs []*llvm.UnfuncInfo) []string {
	seen := make(map[string]bool)
	var result []string
	for _, uf := range unfuncs {
		for _, expr := range []ast.Expression{uf.Requires, uf.Emits} {
			if expr == nil {
				continue
			}
			for _, fieldBase := range collectUnfuncFields(expr) {
				if !seen[fieldBase] {
					seen[fieldBase] = true
					result = append(result, fieldBase)
				}
			}
		}
	}
	return result
}

func (g *Generator) SMT() string {
	return strings.Join(g.smt, "\n")
}
