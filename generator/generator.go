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

func (g *Generator) SMT() string {
	return strings.Join(g.smt, "\n")
}
