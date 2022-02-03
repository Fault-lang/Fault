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
	if c.currScope != "" && c.contextFuncName != "__run" &&
		c.currScope != c.contextFuncName {
		return append([]string{c.currScope}, id...)
	} else {
		return id
	}
}

func (c *Compiler) getVariableName(id []string) string {
	id, _ = c.GetSpec(id)
	return strings.Join(id, "_")
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
	//name := c.getVariableStateName(id)
	name := c.getVariableName(id)
	var alloc *ir.InstAlloca
	var store *ir.InstStore

	switch v := val.(type) {
	case *constant.CharArray:
		l := uint64(len(v.X))
		alloc = c.contextBlock.NewAlloca(&irtypes.ArrayType{TypeName: "string", Len: l, ElemType: irtypes.I8})
		alloc.SetName(name)
		store = c.contextBlock.NewStore(v, alloc)
	case *constant.Int:
		alloc = c.contextBlock.NewAlloca(irtypes.I1)
		alloc.SetName(name)
		store = c.contextBlock.NewStore(v, alloc)
	case *constant.Float:
		alloc = c.contextBlock.NewAlloca(irtypes.Double)
		alloc.SetName(name)
		store = c.contextBlock.NewStore(v, alloc)
	case *constant.Null:
		return //Figure out what to do here
	case *ir.InstFAdd:
		alloc = c.contextBlock.NewAlloca(irtypes.Double)
		alloc.SetName(name)
		store = c.contextBlock.NewStore(v, alloc)
	case *ir.InstFSub:
		alloc = c.contextBlock.NewAlloca(irtypes.Double)
		alloc.SetName(name)
		store = c.contextBlock.NewStore(v, alloc)
	case *ir.InstFMul:
		alloc = c.contextBlock.NewAlloca(irtypes.Double)
		alloc.SetName(name)
		store = c.contextBlock.NewStore(v, alloc)
	case *ir.InstFDiv:
		alloc = c.contextBlock.NewAlloca(irtypes.Double)
		alloc.SetName(name)
		store = c.contextBlock.NewStore(v, alloc)
	case *ir.InstFRem:
		alloc = c.contextBlock.NewAlloca(irtypes.Double)
		alloc.SetName(name)
		store = c.contextBlock.NewStore(v, alloc)
	case *ir.InstFCmp:
		alloc = c.contextBlock.NewAlloca(irtypes.I1)
		alloc.SetName(name)
		store = c.contextBlock.NewStore(v, alloc)
	case *ir.Func:
		return
	default:
		panic(fmt.Sprintf("unknown variable type %T line: %d col: %d", v, pos[0], pos[1]))
	}

	//Add round metadata
	/*round := &metadata.Attachment{
		Name: fmt.Sprintf("round-%d", c.runRound),
		Node: &metadata.DIBasicType{
			MetadataID: -1,
			Tag:        enum.DwarfTagStringType,
		}}
	store.Metadata = append(store.Metadata, round)*/

	//Other metadata
	if c.contextMetadata != nil {
		store.Metadata = append(store.Metadata, c.contextMetadata)
	}
	c.storeAllocation(name, id, alloc)
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
	case *constant.Null:
		alloc := c.module.NewGlobalDef(name, val.(constant.Constant))
		c.storeGlobal(name, alloc)
	case *ir.InstFAdd:
		c.allocVariable(id, val, pos)
	case *ir.InstFSub:
		c.allocVariable(id, val, pos)
	case *ir.InstFMul:
		c.allocVariable(id, val, pos)
	case *ir.InstFDiv:
		c.allocVariable(id, val, pos)
	case *ir.InstFRem:
		c.allocVariable(id, val, pos)
	case *ir.InstFCmp:
		c.allocVariable(id, val, pos)
	case *ir.Func:
	default:
		panic(fmt.Sprintf("unknown variable type %T line: %d col: %d", v, pos[0], pos[1]))
	}
}

func (c *Compiler) storeAllocation(name string, id []string, alloc *ir.InstAlloca) {
	id, s := c.GetSpec(id)
	s.vars.IncrState(id)
	s.vars.Store(id, name, alloc)
}

func (c *Compiler) storeGlobal(name string, alloc *ir.Global) {
	c.specGlobals[name] = alloc
}

func (c *Compiler) fetchGlobal(name string) *ir.Global {
	return c.specGlobals[name]
}
