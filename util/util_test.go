package util

import (
	"fault/ast"
	"os"
	"testing"
)

func TestPreparse(t *testing.T) {
	token := ast.Token{Literal: "test", Position: []int{1, 2, 3, 4}}
	pairs := make(map[ast.Expression]ast.Expression)
	pairs[&ast.Identifier{Token: token, Value: "foo"}] = &ast.IntegerLiteral{Token: token, Value: 3}
	pairs[&ast.Identifier{Token: token, Value: "bash"}] = &ast.FunctionLiteral{Token: token,
		Parameters: []*ast.Identifier{{Token: token, Value: "foo"}},
		Body: &ast.BlockStatement{Token: token,
			Statements: []ast.Statement{&ast.ConstantStatement{Token: token, Name: &ast.Identifier{Token: token, Value: "buzz"}, Value: &ast.IntegerLiteral{Token: token, Value: 20}}}}}

	ret := Preparse(pairs)

	if len(ret) != 2 {
		t.Fatalf("item removed from map. got=%s", ret)
	}

	for k, v := range ret {
		if k == "foo" {
			ty, ok := v.(*ast.IntegerLiteral)
			if !ok {
				t.Fatalf("pair type incorrect. want=IntegerLiteral got=%T", v)
			}

			if ty.Value != 3 {
				t.Fatalf("pair value incorrect. want=3 got=%d", ty.Value)
			}
		} else if k == "bash" {
			ty, ok := v.(*ast.BlockStatement)
			if !ok {
				t.Fatalf("pair type incorrect. want=BlockStatement got=%T", v)
			}

			if ty.Statements[0].(*ast.ConstantStatement).Value.(*ast.IntegerLiteral).Value != 20 {
				t.Fatalf("pair value incorrect. want=20 got=%d", ty.Statements[0].(*ast.ConstantStatement).Value.(*ast.IntegerLiteral).Value)
			}
		} else {
			t.Fatalf("pair key unrecognized. got=%s", k)
		}
	}
}

func TestFilepath(t *testing.T) {
	var host string
	var ok bool

	if host, ok = os.LookupEnv("FAULT_HOST"); !ok {
		host = ""
	}
	os.Setenv("FAULT_HOST", "/host")
	filepath1 := "foo/test/file/system../test.spec"
	filepath1a := Filepath(filepath1)
	if filepath1a != "/host/foo/test/file/test.spec" {
		t.Fatalf("filepath not correct. want=/host/foo/test/file/test.spec got=%s", filepath1a)
	}

	filepath2 := "foo/test/file/system../../test.spec"
	filepath2a := Filepath(filepath2)
	if filepath2a != "/host/foo/test/test.spec" {
		t.Fatalf("filepath not correct. want=/host/foo/test/test.spec got=%s", filepath2a)
	}

	filepath3 := "foo/test/file/system../../../test.spec"
	filepath3a := Filepath(filepath3)
	if filepath3a != "/host/foo/test.spec" {
		t.Fatalf("filepath not correct. want=/host/foo/test.spec got=%s", filepath3a)
	}

	filepath4 := "foo/test/file/system/~/test.spec"
	filepath4a := Filepath(filepath4)
	if filepath4a != "/host/test.spec" {
		t.Fatalf("filepath not correct. want=/host/test.spec got=%s", filepath4a)
	}

	filepath5 := "foo/test/file/system/~test.spec"
	filepath5a := Filepath(filepath5)
	if filepath5a != "/host/test.spec" {
		t.Fatalf("filepath not correct. want=/host/test.spec got=%s", filepath5a)
	}

	filepath6 := "test.spec"
	filepath6a := Filepath(filepath6)
	if filepath6a != "/host/test.spec" {
		t.Fatalf("filepath not correct. want=/host/test.spec got=%s", filepath6a)
	}

	filepath7 := "/.."
	filepath7a := Filepath(filepath7)
	if filepath7a != "/host/" {
		t.Fatalf("filepath not correct. want=/host/ got=%s", filepath7a)
	}

	os.Setenv("FAULT_HOST", host)
}

func TestCartesian(t *testing.T) {
	list1 := []string{"a", "b", "c"}
	list2 := []string{"1", "2"}
	r := Cartesian(list1, list2)

	if r[0][0] != "a" || r[0][1] != "1" {
		t.Fatalf("cartesian product not correct. got=%s", r)
	}

	if r[1][0] != "a" || r[1][1] != "2" {
		t.Fatalf("cartesian product not correct. got=%s", r)
	}

	if r[2][0] != "b" || r[2][1] != "1" {
		t.Fatalf("cartesian product not correct. got=%s", r)
	}

	if r[3][0] != "b" || r[3][1] != "2" {
		t.Fatalf("cartesian product not correct. got=%s", r)
	}

}

func TestCartesianMulti(t *testing.T) {
	list1 := []string{"a", "b", "c"}
	list2 := []string{"1", "2"}
	list3 := []string{"3", "4"}
	r := CartesianMulti([][]string{list1, list2, list3})

	if r[0][0] != "a" || r[0][1] != "1" || r[0][2] != "3" {
		t.Fatalf("cartesian product not correct. got=%s", r)
	}

	if r[1][0] != "a" || r[1][1] != "1" || r[1][2] != "4" {
		t.Fatalf("cartesian product not correct. got=%s", r)
	}

	if r[2][0] != "a" || r[2][1] != "2" || r[2][2] != "3" {
		t.Fatalf("cartesian product not correct. got=%s", r)
	}

	if r[3][0] != "a" || r[3][1] != "2" || r[3][2] != "4" {
		t.Fatalf("cartesian product not correct. got=%s", r)
	}

}
