package llvm

import (
	"fault/ast"
	"fault/llvm/name"
	"strings"
	"testing"

	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
)

func TestDefineSpecVar(t *testing.T) {
	id := []string{"test", "this", "func"}
	s := initSpec(id)

	if s.name != "test" {
		t.Fatal("spec this not created")
	}
	v := s.GetSpecVar(id)
	if v.(*constant.Int).X.Int64() != int64(0) {
		t.Fatal("spec var this.func not added correctly")
	}
}

func TestSpecState(t *testing.T) {
	id := []string{"test", "this", "func"}
	s := initSpec(id)

	state := s.GetSpecVarState(id)
	if state != 0 {
		t.Fatalf("spec var this.func has the wrong state label. got=%d want=0", state)
	}

	val := constant.NewInt(types.I32, 5)
	s.DefineSpecVar(id, val)

	state2 := s.GetSpecVarState(id)
	if state2 != 1 {
		t.Fatalf("spec var this.func has the wrong state label. got=%d want=1", state2)
	}

	v := s.GetSpecVar(id)
	if v.(*constant.Int).X.Int64() != int64(5) {
		t.Fatalf("spec var this.func was not updated. got=%d want=5", v.(*constant.Int).X.Int64())
	}
}

func TestSpecPointer(t *testing.T) {
	id := []string{"test", "this", "func"}
	s := initSpec(id)
	fvn := strings.Join(id, "_")
	b := ir.NewBlock(name.Block())
	alloc := b.NewAlloca(types.I1)
	alloc.SetName(fvn)
	s.vars.Store(id, fvn, alloc)

	pointer := s.GetSpecVarPointer(fvn)
	if pointer.LocalName != "test_this_func" {
		t.Fatal("spec var this.func is missing a pointer")
	}
}

func TestParams(t *testing.T) {
	id := []string{"test", "this", "func"}
	s := initSpec(id)
	param := constant.NewInt(types.I32, 5)

	s.AddParam(id, param)
	p := s.GetParams(id)

	if p[0].(*constant.Int).X.Int64() != int64(5) {
		t.Fatal("spec var this.func is missing parameters")
	}

}

func TestSpecTypes(t *testing.T) {
	id := []string{"test", "this", "func"}
	s := initSpec(id)

	fvn := strings.Join(id, "_")
	s.DefineSpecType(fvn, ast.Type{})

	if _, ok := s.GetSpecType(fvn, true); !ok {
		t.Fatal("spec var this.func is missing type")
	}

}

func initSpec(id []string) *spec {
	s := NewCompiledSpec("test")
	val := constant.NewInt(types.I32, 0)
	s.DefineSpecVar(id, val)
	return s
}
