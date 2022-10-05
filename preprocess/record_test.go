package preprocess

import (
	"fault/ast"
	"testing"
)

func TestSpecRecord(t *testing.T) {
	sr := NewSpecRecord()

	testNodes := map[string]ast.Node{
		"bar": &ast.IntegerLiteral{Value: 1}}

	sr.AddStock("test", testNodes)
	t1 := sr.FetchStock("test")
	if t1 == nil {
		t.Fatal("stock not found in spec record")
	}

	i, ok := t1["bar"].(*ast.IntegerLiteral)
	if !ok {
		t.Fatalf("property bar is incorrect type got=%T want=IntegerLiteral", t1["bar"])
	}

	if i.Value != 1 {
		t.Fatalf("property bar has incorrect value got=%d want=1", i.Value)
	}

	sr.AddFlow("test", testNodes)
	t2 := sr.FetchFlow("test")
	if t2 == nil {
		t.Fatal("stock not found in spec record")
	}

	i2, ok := t2["bar"].(*ast.IntegerLiteral)
	if !ok {
		t.Fatalf("property bar is incorrect type got=%T want=IntegerLiteral", t2["bar"])
	}

	if i2.Value != 1 {
		t.Fatalf("property bar has incorrect value got=%d want=1", i2.Value)
	}

	sr.AddComponent("test", testNodes)
	t3 := sr.FetchComponent("test")
	if t3 == nil {
		t.Fatal("stock not found in spec record")
	}

	i3, ok := t3["bar"].(*ast.IntegerLiteral)
	if !ok {
		t.Fatalf("property bar is incorrect type got=%T want=IntegerLiteral", t3["bar"])
	}

	if i3.Value != 1 {
		t.Fatalf("property bar has incorrect value got=%d want=1", i3.Value)
	}
}
