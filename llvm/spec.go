package llvm

import (
	"fault/llvm/variables"
	"fmt"
	"strings"

	"github.com/llir/llvm/ir"
	irtypes "github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
)

// Representation of a spec
type spec struct {
	name string
	vars *variables.LookupTable
}

func NewCompiledSpec(name string) *spec {
	return &spec{
		name: name,
		vars: variables.NewTable(),
	}
}

func (s *spec) DefineSpecVar(rawid []string, val value.Value) {
	if s.GetSpecVar(rawid) != nil {
		s.vars.Update(rawid, val)
	} else {
		s.vars.Add(rawid, val)
	}
}

func (s *spec) GetSpecVar(rawid []string) value.Value {
	return s.vars.Get(rawid)
}

func (s *spec) GetSpecVarState(rawid []string) int16 {
	return s.vars.GetState(rawid)
}

func (s *spec) GetSpecVarPointer(rawid []string) *ir.InstAlloca {
	name := strings.Join(rawid, "_")
	return s.vars.GetPointer(name)
}

func (s *spec) GetParams(rawid []string) []value.Value {
	return s.vars.GetParams(rawid)
}

func (s *spec) AddParam(rawid []string, p value.Value) {
	s.vars.AddParam(rawid, p)
}

func (s *spec) AddParams(rawid []string, p []value.Value) {
	s.vars.AddParams(rawid, p)
}

func (s *spec) DefineSpecType(rawid []string, ty irtypes.Type) {
	s.vars.Type(rawid, ty)
}

func (s *spec) GetSpecType(name string) irtypes.Type {
	return s.vars.GetType(name)
}

func (s *spec) GetPointerType(name string) irtypes.Type {
	ty := s.vars.GetType(name)
	if ty != nil {
		switch ty {
		case irtypes.Double:
			return DoubleP
		case irtypes.I1:
			return I1P
		default:
			panic(fmt.Sprintf("invalid pointer type %T for variable %s", ty, name))
		}
	}
	return DoubleP //Should reconsider this at some point and err here instead
}
