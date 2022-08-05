package llvm

import (
	"fault/llvm/variables"
	"fmt"
	"unicode"

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

func (s *spec) DefineSpecVar(id []string, val value.Value) {
	if s.GetSpecVar(id) != nil {
		s.vars.Update(id, val)
	} else {
		s.vars.Add(id, val)
	}
}

func (s *spec) GetSpecVar(id []string) value.Value {
	return s.vars.Get(id)
}

func (s *spec) GetSpecVarState(id []string) int16 {
	return s.vars.GetState(id)
}

func (s *spec) GetSpecVarPointer(name string) *ir.InstAlloca {
	return s.vars.GetPointer(name)
}

func (s *spec) GetParams(id []string) []value.Value {
	return s.vars.GetParams(id)
}

func (s *spec) AddParam(id []string, p value.Value) {
	s.vars.AddParam(id, p)
}

func (s *spec) DefineSpecType(id []string, ty irtypes.Type) {
	s.vars.Type(id, ty)
}

func (s *spec) GetSpecType(name string, inSamePackage bool) (irtypes.Type, bool) {
	if unicode.IsLower([]rune(name)[0]) && !inSamePackage {
		panic(fmt.Sprintf("Can't use %s from outside of %s", name, s.name))
	}

	ty := s.vars.GetType(name)
	if ty != nil {
		return ty, true
	}
	return nil, false
}
