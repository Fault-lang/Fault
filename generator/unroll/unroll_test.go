package unroll

import (
	"fault/llvm"
	"fault/smt/rules"
	"testing"

	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	irtypes "github.com/llir/llvm/ir/types"
	"github.com/stretchr/testify/assert"
)

func TestNewEnv(t *testing.T) {
	env := NewEnv()
	assert.NotNil(t, env)
	assert.Empty(t, env.VarLoads)
	assert.Empty(t, env.VarTypes)
}

func TestNewLLFunc(t *testing.T) {
	env := NewEnv()
	irf := &ir.Func{}
	llFunc := NewLLFunc(env, irf)
	assert.NotNil(t, llFunc)
	assert.Equal(t, env, llFunc.Env)
	assert.Empty(t, llFunc.Rules)
	assert.Nil(t, llFunc.Start)
	assert.Empty(t, llFunc.localCallstack)
	assert.Empty(t, llFunc.functions)
	assert.Empty(t, llFunc.rawFunctions)
	assert.Equal(t, irf, llFunc.rawIR)
}

// Add more tests for other functions and types...
func TestGenerateCallstack(t *testing.T) {
	// Test case 1: Callstack with one function name
	llf := NewLLFunc(NewEnv(), ir.NewFunc("test", irtypes.Void))
	callstack := []string{"foo"}
	functions := make(map[string]*LLFunc)
	functions["foo"] = &LLFunc{
		Env:          llf.Env,
		Rules:        []rules.Rule{},
		functions:    make(map[string]*LLFunc),
		rawFunctions: make(map[string]*ir.Func),
		rawIR:        ir.NewFunc("foo", irtypes.Void),
	}
	llf.functions = functions
	result := GenerateCallstack(llf, callstack)
	assert.Equal(t, functions["foo"].String(), result.String())

	// Test case 2: Callstack with one block name
	llb := NewLLBlock(NewEnv(), ir.NewBlock("test"))
	llb.functions = functions
	callstack = []string{"foo"}
	result = GenerateCallstack(llb, callstack)
	assert.Equal(t, functions["foo"].String(), result.String())
}

func TestNewConstants(t *testing.T) {
	e := NewEnv()
	globals := []*ir.Global{
		ir.NewGlobalDef("test_global1", constant.NewFloat(irtypes.Double, 10)),
		ir.NewGlobalDef("test_global2", constant.NewInt(irtypes.I1, 0)),
		ir.NewGlobalDef("test_global3", constant.NewFloat(irtypes.Double, 30)),
	}
	rawInputs := &llvm.RawInputs{}

	expected := []rules.Rule{
		declareVar("test_global1", "Real", "10.0"),
		declareVar("test_global2", "Bool", "0"),
		declareVar("test_global3", "Real", "30.0"),
	}

	result := NewConstants(e, globals, rawInputs)
	assert.Equal(t, expected, result)
}
