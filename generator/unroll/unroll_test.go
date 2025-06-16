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
	e := NewEnv(llvm.NewRawInputs())
	globals := []*ir.Global{
		ir.NewGlobalDef("test_global1", constant.NewFloat(irtypes.Double, 10)),
		ir.NewGlobalDef("test_global2", constant.NewInt(irtypes.I1, 0)),
		ir.NewGlobalDef("test_global3", constant.NewFloat(irtypes.Double, 30)),
	}
	rawInputs := &llvm.RawInputs{}

	expected := []rules.Rule{
		declareVar("test_global1", "Real", &rules.Wrap{Value: "10.0"}, false),
		declareVar("test_global2", "Bool", &rules.Wrap{Value: "0"}, false),
		declareVar("test_global3", "Real", &rules.Wrap{Value: "30.0"}, false),
	}

	result := NewConstants(e, globals, rawInputs)
	assert.Equal(t, expected, result)
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
