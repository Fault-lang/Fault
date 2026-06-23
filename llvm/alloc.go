package llvm

import (
	"fmt"
	"strings"

	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	irtypes "github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
)

func (c *Compiler) updateVariableStateName(id []string) string {
	if len(id) == 2 { // This is a constant, doesn't change
		return strings.Join(id, "_")
	}
	s := c.specs[id[0]]

	incr := s.GetSpecVarState(id[1:])
	return fmt.Sprint(strings.Join(id, "_"), incr+1)
}

func (c *Compiler) allocVariable(id []string, val value.Value, loc string) {
	name := strings.Join(id, "_")

	var alloc *ir.InstAlloca

	switch v := val.(type) {
	case *constant.CharArray:
		l := uint64(len(v.X))
		alloc = c.contextBlock.NewAlloca(&irtypes.ArrayType{TypeName: "string", Len: l, ElemType: irtypes.I8})
		alloc.SetName(name)
		c.contextBlock.NewStore(v, alloc)
	case *constant.Int:
		alloc = c.contextBlock.NewAlloca(irtypes.I1)
		alloc.SetName(name)
		c.contextBlock.NewStore(v, alloc)
	case *constant.Float:
		alloc = c.contextBlock.NewAlloca(irtypes.Double)
		alloc.SetName(name)
		c.contextBlock.NewStore(v, alloc)
	case *constant.Null:
		return //Figure out what to do here
	case *ir.InstFAdd:
		alloc = c.contextBlock.NewAlloca(irtypes.Double)
		alloc.SetName(name)
		c.contextBlock.NewStore(v, alloc)
	case *ir.InstFSub:
		alloc = c.contextBlock.NewAlloca(irtypes.Double)
		alloc.SetName(name)
		c.contextBlock.NewStore(v, alloc)
	case *ir.InstFMul:
		alloc = c.contextBlock.NewAlloca(irtypes.Double)
		alloc.SetName(name)
		c.contextBlock.NewStore(v, alloc)
	case *ir.InstFDiv:
		alloc = c.contextBlock.NewAlloca(irtypes.Double)
		alloc.SetName(name)
		c.contextBlock.NewStore(v, alloc)
	case *ir.InstFRem:
		alloc = c.contextBlock.NewAlloca(irtypes.Double)
		alloc.SetName(name)
		c.contextBlock.NewStore(v, alloc)
	case *ir.InstICmp: //Needed for if true {} constructions
		alloc = c.contextBlock.NewAlloca(irtypes.I1)
		alloc.SetName(name)
		c.contextBlock.NewStore(v, alloc)
	case *ir.InstFCmp:
		alloc = c.contextBlock.NewAlloca(irtypes.I1)
		alloc.SetName(name)
		c.contextBlock.NewStore(v, alloc)
	case *ir.InstAnd:
		alloc = c.contextBlock.NewAlloca(irtypes.I1)
		alloc.SetName(name)
		if v.Type() == nil {
			v.Typ = irtypes.I1
		}
		c.contextBlock.NewStore(v, alloc)
	case *ir.InstOr:
		alloc = c.contextBlock.NewAlloca(irtypes.I1)
		alloc.SetName(name)
		if v.Type() == nil {
			v.Typ = irtypes.I1
		}
		c.contextBlock.NewStore(v, alloc)
	case *ir.Func:
		return
	default:
		panic(fmt.Sprintf("unknown variable type %T %s", v, loc))
	}

	c.storeAllocation(name, id, alloc)
}

func (c *Compiler) globalVariable(id []string, val value.Value, loc string) {
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
		c.allocVariable(id, val, loc)
	case *ir.InstFSub:
		c.allocVariable(id, val, loc)
	case *ir.InstFMul:
		c.allocVariable(id, val, loc)
	case *ir.InstFDiv:
		c.allocVariable(id, val, loc)
	case *ir.InstFRem:
		c.allocVariable(id, val, loc)
	case *ir.InstICmp:
		c.allocVariable(id, val, loc)
	case *ir.InstFCmp:
		c.allocVariable(id, val, loc)
	case *ir.Func:
	case *ir.InstAnd:
		placeholder := constant.NewAnd(v.X.(constant.Expression), v.Y.(constant.Expression))
		alloc := c.module.NewGlobalDef(name, placeholder)
		c.storeGlobal(name, alloc)
	case *ir.InstOr:
		placeholder := constant.NewOr(v.X.(constant.Expression), v.Y.(constant.Expression))
		alloc := c.module.NewGlobalDef(name, placeholder)
		c.storeGlobal(name, alloc)
	default:
		panic(fmt.Sprintf("unknown variable type %T %s", v, loc))
	}

}

func (c *Compiler) storeAllocation(name string, id []string, alloc *ir.InstAlloca) {
	s := c.specs[id[0]]
	s.vars.IncrState(id)
	s.vars.Store(id, name, alloc)
}

func (c *Compiler) storeGlobal(name string, alloc *ir.Global) {
	c.specGlobals[name] = alloc
}
