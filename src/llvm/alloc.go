package llvm

import (
	"fmt"
	"strings"

	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	irtypes "github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
)

func (c *Compiler) getFullVariableName(id []string) []string {
	if c.currScope != "" && c.contextFuncName != "__run" {
		return append([]string{c.currScope}, id...)
	} else {
		return id
	}
}

func (c *Compiler) getVariableStateName(id []string) string {
	id, s := c.GetSpec(id)
	incr := s.GetSpecVarState(id)
	return fmt.Sprint(strings.Join(id, "_"), incr)
}

func (c *Compiler) updateVariableStateName(id []string) string {
	id, s := c.GetSpec(id)
	if len(id) == 2 { // This is a constant, doesn't change
		return strings.Join(id, "_")
	}

	incr := s.GetSpecVarState(id)
	return fmt.Sprint(strings.Join(id, "_"), incr+1)
}

func (c *Compiler) allocVariable(id []string, val value.Value, pos []int) {
	id, _ = c.GetSpec(id)
	name := c.updateVariableStateName(id)

	switch v := val.(type) {
	case *constant.CharArray:
		l := uint64(len(v.X))
		alloc := c.contextBlock.NewAlloca(&irtypes.ArrayType{"string", l, irtypes.I8})
		alloc.SetName(name)
		c.contextBlock.NewStore(v, alloc)
		c.storeAllocation(name, id, alloc)
	case *constant.Int:
		alloc := c.contextBlock.NewAlloca(irtypes.I1)
		alloc.SetName(name)
		c.contextBlock.NewStore(v, alloc)
		c.storeAllocation(name, id, alloc)
	case *constant.Float:
		alloc := c.contextBlock.NewAlloca(irtypes.Double)
		alloc.SetName(name)
		c.contextBlock.NewStore(v, alloc)
		c.storeAllocation(name, id, alloc)
	case *ir.InstFAdd:
		alloc := c.contextBlock.NewAlloca(irtypes.Double)
		alloc.SetName(name)
		c.contextBlock.NewStore(v, alloc)
		c.storeAllocation(name, id, alloc)
	case *ir.InstFSub:
		alloc := c.contextBlock.NewAlloca(irtypes.Double)
		alloc.SetName(name)
		c.contextBlock.NewStore(v, alloc)
		c.storeAllocation(name, id, alloc)
	case *ir.InstFMul:
		alloc := c.contextBlock.NewAlloca(irtypes.Double)
		alloc.SetName(name)
		c.contextBlock.NewStore(v, alloc)
		c.storeAllocation(name, id, alloc)
	case *ir.InstFDiv:
		alloc := c.contextBlock.NewAlloca(irtypes.Double)
		alloc.SetName(name)
		c.contextBlock.NewStore(v, alloc)
		c.storeAllocation(name, id, alloc)
	case *ir.InstFRem:
		alloc := c.contextBlock.NewAlloca(irtypes.Double)
		alloc.SetName(name)
		c.contextBlock.NewStore(v, alloc)
		c.storeAllocation(name, id, alloc)
	case *ir.InstFCmp:
		alloc := c.contextBlock.NewAlloca(irtypes.I1)
		alloc.SetName(name)
		c.contextBlock.NewStore(v, alloc)
		c.storeAllocation(name, id, alloc)
	case *ir.Func:

	default:
		panic(fmt.Sprintf("unknown variable type %T line: %d col: %d", v, pos[0], pos[1]))
	}
}

func (c *Compiler) globalVariable(id []string, val value.Value, pos []int) {
	id, _ = c.GetSpec(id)
	name := c.updateVariableStateName(id)

	switch v := val.(type) {
	case *constant.CharArray:
		alloc := c.module.NewGlobalDef(name, val.(constant.Constant))
		c.storeGlobal(name, alloc)
	case *constant.Int:
		alloc := c.module.NewGlobalDef(name, val.(constant.Constant))
		c.storeGlobal(name, alloc)
	case *constant.Float:
		alloc := c.module.NewGlobalDef(name, val.(constant.Constant))
		c.storeGlobal(name, alloc)
	case *ir.InstFAdd:
		c.allocVariable(id, val, pos)
		//alloc := c.module.NewGlobalDef(name, val.(constant.Constant))
		//c.storeGlobal(name, alloc)
	case *ir.InstFSub:
		c.allocVariable(id, val, pos)
		//alloc := c.module.NewGlobalDef(name, val.(constant.Constant))
		//c.storeGlobal(name, alloc)
	case *ir.InstFMul:
		c.allocVariable(id, val, pos)
		//alloc := c.module.NewGlobalDef(name, val.(constant.Constant))
		//c.storeGlobal(name, alloc)
	case *ir.InstFDiv:
		c.allocVariable(id, val, pos)
		//alloc := c.module.NewGlobalDef(name, val.(constant.Constant))
		//c.storeGlobal(name, alloc)
	case *ir.InstFRem:
		c.allocVariable(id, val, pos)
		//alloc := c.module.NewGlobalDef(name, val.(constant.Constant))
		//c.storeGlobal(name, alloc)
	case *ir.InstFCmp:
		c.allocVariable(id, val, pos)
		//alloc := c.module.NewGlobalDef(name, val.(constant.Constant))
		//c.storeGlobal(name, alloc)
	case *ir.Func:
	default:
		panic(fmt.Sprintf("unknown variable type %T line: %d col: %d", v, pos[0], pos[1]))
	}
}

func (c *Compiler) storeAllocation(name string, id []string, alloc *ir.InstAlloca) {
	c.specs[c.currentSpecName].vars.Store(id, name, alloc)
}

func (c *Compiler) fetchAllocation(id []string) *ir.InstAlloca {
	id, s := c.GetSpec(id)
	name := c.getVariableStateName(id)
	return s.vars.GetPointer(name)
}

func (c *Compiler) storeGlobal(name string, alloc *ir.Global) {
	c.specGlobals[name] = alloc
}

func (c *Compiler) fetchGlobal(name string) *ir.Global {
	return c.specGlobals[name]
}
