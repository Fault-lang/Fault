package preprocess

import (
	"fault/ast"
	"testing"
)

func TestConstantsErr(t *testing.T) {
	p := NewProcesser()
	p.trail = p.trail.PushSpec("test")
	p.Specs["test"] = NewSpecRecord()
	p.Specs["test"].SpecName = "test"
	p.Specs["test"].AddConstant("foo", &ast.IntegerLiteral{Value: 2})

	test := &ast.Spec{Statements: []ast.Statement{&ast.ConstantStatement{Name: &ast.Identifier{Spec: "test", Value: "foo"},
		Value: &ast.IntegerLiteral{Value: 2},
	}}}

	_, err := p.walk(test)
	if err == nil {
		t.Fatal("failed to error on constant redeclare")
	}

	if err.Error() != "variable foo is a constant and cannot be modified" {
		t.Fatalf("error message on constant redeclare incorrect got=%s", err.Error())
	}
}

func TestInstanceErr(t *testing.T) {
	p := NewProcesser()
	p.trail = p.trail.PushSpec("test")
	p.Specs["test"] = NewSpecRecord()
	p.Specs["test"].SpecName = "test"
	p.initialPass = false

	test := &ast.Instance{Value: &ast.Identifier{Spec: "test", Value: "bash"}, Name: "foo"}

	_, err := p.walk(test)
	if err == nil {
		t.Fatal("failed to error on unknown instance")
	}

	if err.Error() != "can't find an instance named bash" {
		t.Fatalf("error message on unknown instance incorrect got=%s", err.Error())
	}
}

func TestStructInstanceErr(t *testing.T) {
	p := NewProcesser()
	p.trail = p.trail.PushSpec("test")
	p.Specs["test"] = NewSpecRecord()
	p.Specs["test"].SpecName = "test"

	stockdata := make(map[*ast.Identifier]ast.Expression)
	stockdata[&ast.Identifier{Spec: "test", Value: "foo"}] = &ast.StructInstance{Spec: "test", Name: "foo", Parent: []string{"test", "bar"}}
	test := &ast.DefStatement{Name: &ast.Identifier{Spec: "test", Value: "zoo"}, Value: &ast.StockLiteral{Pairs: stockdata}}

	_, err := p.walk(test)
	if err == nil {
		t.Fatal("failed to error on unknown instance")
	}

	if err.Error() != "can't find a struct instance named [test bar]" {
		t.Fatalf("error message on unknown instance incorrect got=%s", err.Error())
	}
}

// func TestParamCallErr(t *testing.T) {
// 	p := NewProcesser()
// 	p.trail = p.trail.PushSpec("test")
// 	p.Specs["test"] = NewSpecRecord()
// 	p.Specs["test"].SpecName = "test"
// 	p.initialPass = false

// 	test := &ast.ParameterCall{Spec: "test", Value: []string{"foo", "bar"}}

// 	_, err := p.walk(test)
// 	if err == nil {
// 		t.Fatal("failed to error on unknown instance")
// 	}

// 	if err.Error() != "can't find a struct instance named [test bar]" {
// 		t.Fatalf("error message on unknown instance incorrect got=%s", err.Error())
// 	}
// }
