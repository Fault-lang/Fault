package unroll

import (
	"fault/generator/rules"
	"fault/llvm"
	"testing"

	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	irtypes "github.com/llir/llvm/ir/types"
	"github.com/stretchr/testify/assert"
)

func TestNewEnv(t *testing.T) {
	env := NewEnv(llvm.NewRawInputs())
	assert.NotNil(t, env)
	assert.Empty(t, env.VarLoads)
	assert.Empty(t, env.VarTypes)
}

func TestNewLLFunc(t *testing.T) {
	env := NewEnv(llvm.NewRawInputs())
	irf := &ir.Func{}
	llFunc := NewLLFunc(env, make(map[string]*ir.Func), irf)
	assert.NotNil(t, llFunc)
	assert.Equal(t, env, llFunc.Env)
	assert.Empty(t, llFunc.Rules)
	assert.Nil(t, llFunc.Start)
	assert.Empty(t, llFunc.localCallstack)
	assert.Empty(t, llFunc.functions)
	assert.Empty(t, llFunc.rawFunctions)
	assert.Equal(t, irf, llFunc.rawIR)
}

func TestNewConstants(t *testing.T) {
	// Spec-level constants (no version digit suffix) are always declared in SMT.
	e := NewEnv(llvm.NewRawInputs())
	globals := []*ir.Global{
		ir.NewGlobalDef("spec_constA", constant.NewFloat(irtypes.Double, 10)),
		ir.NewGlobalDef("spec_constB", constant.NewInt(irtypes.I1, 0)),
		ir.NewGlobalDef("spec_constC", constant.NewFloat(irtypes.Double, 30)),
	}
	rawInputs := &llvm.RawInputs{}

	expected := []rules.Rule{
		declareVar("spec_constA", "Real", &rules.Wrap{Value: "10.0"}, false),
		declareVar("spec_constB", "Bool", &rules.Wrap{Value: "0"}, false),
		declareVar("spec_constC", "Real", &rules.Wrap{Value: "30.0"}, false),
	}

	result := NewConstants(e, globals, rawInputs)
	assert.Equal(t, expected, result)
}

func TestNewConstantsInlinesUnmodifiedStocks(t *testing.T) {
	// Versioned stock properties (digit suffix) that have no local alloca
	// (i.e. no flow modifies them) should be inlined into ConstantVals rather
	// than declared as free SMT variables.
	e := NewEnv(llvm.NewRawInputs())
	// s_v1 is a versioned stock property; MutableVars["s_v"] is unset → constant stock
	globals := []*ir.Global{
		ir.NewGlobalDef("spec_s_v1", constant.NewFloat(irtypes.Double, 5)),
	}
	rawInputs := &llvm.RawInputs{}

	result := NewConstants(e, globals, rawInputs)
	assert.Empty(t, result, "constant stock should be inlined, not declared")

	litVal, ok := e.ConstantVals["spec_s_v1"]
	assert.True(t, ok, "constant stock literal should be stored in ConstantVals")
	assert.NotNil(t, litVal)
}

func TestNewConstantsDeclaresModifiedStocks(t *testing.T) {
	// Versioned stock properties that DO have a local alloca (flow modifies them)
	// must still be declared as SMT variables.
	e := NewEnv(llvm.NewRawInputs())
	e.MutableVars["spec_s_v"] = true // flow creates spec_s_v2 alloca
	globals := []*ir.Global{
		ir.NewGlobalDef("spec_s_v1", constant.NewFloat(irtypes.Double, 5)),
	}
	rawInputs := &llvm.RawInputs{}

	result := NewConstants(e, globals, rawInputs)
	assert.Len(t, result, 1, "mutable stock's initial value should be declared in SMT")
	assert.Empty(t, e.ConstantVals, "mutable stock should not be in ConstantVals")
}
func TestFindUsedVars(t *testing.T) {
	// Non-__run function: loads from alloca "a", stores to alloca "b".
	// Both should appear in UsedVars.
	flowFunc := ir.NewFunc("flow_fn", irtypes.Void)
	entry := flowFunc.NewBlock("entry")
	allocaA := entry.NewAlloca(irtypes.Double)
	allocaA.SetName("a")
	allocaB := entry.NewAlloca(irtypes.Double)
	allocaB.SetName("b")
	loadedA := entry.NewLoad(irtypes.Double, allocaA)
	entry.NewStore(loadedA, allocaB)

	// __run function: stores to alloca "c". Should be excluded from UsedVars.
	runFunc := ir.NewFunc("__run", irtypes.Void)
	runEntry := runFunc.NewBlock("entry")
	allocaC := runEntry.NewAlloca(irtypes.Double)
	allocaC.SetName("c")
	runEntry.NewStore(constant.NewFloat(irtypes.Double, 0), allocaC)

	used := FindUsedVars([]*ir.Func{flowFunc, runFunc})

	assert.True(t, used["a"], "a (loaded in flow_fn) should be in UsedVars")
	assert.True(t, used["b"], "b (stored in flow_fn) should be in UsedVars")
	assert.False(t, used["c"], "c (only stored in __run) should not be in UsedVars")
}

func TestUnroll(t *testing.T) {
	env := NewEnv(llvm.NewRawInputs())
	irf := ir.NewFunc("test", irtypes.Void)
	llFunc := NewLLFunc(env, make(map[string]*ir.Func), irf)

	// Test case 1: Empty function
	llFunc.Unroll()
	assert.Nil(t, llFunc.Start)

	// Test case 2: Function with one block
	block := ir.NewBlock("test1")
	block.Insts = []ir.Instruction{ir.NewFAdd(constant.NewInt(irtypes.I32, 0), constant.NewInt(irtypes.I32, 1))}
	irf.Blocks = append(irf.Blocks, block)
	llFunc.Unroll()
	assert.NotNil(t, llFunc.Start)
	assert.Equal(t, block, llFunc.Start.rawIR)

	// Test case 3: Function with multiple blocks
	b2 := ir.NewBlock("test2")
	block2 := NewLLBlock(env, llFunc.rawFunctions, b2)
	llFunc.Start.After = block2
	irf.Blocks = append(irf.Blocks, b2)
	llFunc.Unroll()
	assert.NotNil(t, llFunc.Start)
	assert.Equal(t, block, llFunc.Start.rawIR)
	assert.NotNil(t, llFunc.Start.After)
	assert.Equal(t, block2.rawIR, llFunc.Start.After.rawIR)
}
