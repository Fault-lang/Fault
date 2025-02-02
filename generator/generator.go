package generator

import (
	"fault/generator/unroll"
	"fault/llvm"
	"fault/smt/rules"

	"github.com/llir/llvm/asm"
	"github.com/llir/llvm/ir"
)

// Take LL IR and Generate SMTLib2
//
// Step 1: Translate LL IR into Rules by LLBlock
// and Function
//
// Step 2: Add Phi values necessary for the SMT model
// and flatten to a single set of rules

type Generator struct {
	constants []rules.Rule
	functions map[string]*ir.Func
	Env       *unroll.Env
	RawInputs *llvm.RawInputs
	RunBlock  *unroll.LLFunc
}

func NewGenerator() *Generator {
	return &Generator{
		functions: make(map[string]*ir.Func),
		Env:       unroll.NewEnv(),
	}
}
func Execute(compiler *llvm.Compiler) *Generator {
	generator := NewGenerator()
	//generator.LoadMeta(compiler)
	//generator.States = compiler.States
	generator.Run(compiler.GetIR())
	generator.LoadStringRules(compiler.StringRules) // Do last to get SSA values
	return generator
}

func (g *Generator) LoadStringRules(sr map[string]string) {
	// g.Log.StringRules = sr
	// for k := range sr {
	// 	num := g.variables.GetSSANum(k)
	// 	for i := 0; i < int(num)+1; i++ {
	// 		state := fmt.Sprintf("%s_%v", k, i)
	// 		g.Log.IsStringRule[state] = true
	// 	}
	// }
}

func (g *Generator) Run(llopt string) {
	m, err := asm.ParseString("", llopt) //"" because ParseString has a path variable
	if err != nil {
		panic(err)
	}
	g.newCallgraph(m)

}

func (g *Generator) newCallgraph(m *ir.Module) {
	g.constants = unroll.NewConstants(g.Env, m.Globals, g.RawInputs)
	g.sortFuncs(m.Funcs)

	g.RunBlock = unroll.NewLLFunc(g.Env, g.functions["@__run"])
	g.RunBlock.Unroll()

	// g.processAsserts()
	// g.newAsserts(g.RawInputs.Asserts)
	// g.newAssumes(g.RawInputs.Assumes)

}

func (g *Generator) sortFuncs(funcs []*ir.Func) {
	//Iterate through all the function blocks and store them by
	// function call name.
	for _, f := range funcs {
		// Get function name.
		fname := f.Ident()

		if fname != "@__run" {
			g.functions[f.Ident()] = f
			continue
		}
	}
}
