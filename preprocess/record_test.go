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
	t1, _ := sr.FetchStock("test")
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
	t2, _ := sr.FetchFlow("test")
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
	t3, _ := sr.FetchComponent("test")
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

	sr.AddConstant("test", &ast.IntegerLiteral{Value: 1})
	t4, _ := sr.FetchConstant("test")
	if t4 == nil {
		t.Fatal("constant not found in spec record")
	}

	i4, ok := t4.(*ast.IntegerLiteral)
	if !ok {
		t.Fatalf("constant is incorrect type got=%T want=IntegerLiteral", t4)
	}

	if i4.Value != 1 {
		t.Fatalf("constant has incorrect value got=%d want=1", i4.Value)
	}
}

func TestSpecRecordUpdate(t *testing.T) {
	sr := NewSpecRecord()

	testNodes := map[string]ast.Node{
		"bar": &ast.IntegerLiteral{Value: 1}}

	testNodes2 := map[string]ast.Node{
		"foo": &ast.FloatLiteral{Value: .01},
		"bar": &ast.IntegerLiteral{Value: 2}}

	sr.AddStock("test1", testNodes)
	t1, _ := sr.FetchStock("test1")
	sr.UpdateStock("test1", testNodes2)
	t1a, _ := sr.FetchStock("test1")
	if len(t1) == len(t1a) {
		t.Fatal("stock not updated correctly spec record")
	}

	i, ok := t1a["bar"].(*ast.IntegerLiteral)
	if !ok {
		t.Fatalf("property bar is incorrect type got=%T want=IntegerLiteral", t1a["bar"])
	}

	if i.Value != 2 {
		t.Fatalf("property bar has incorrect value got=%d want=2", i.Value)
	}

	sr.AddFlow("test2", testNodes)
	t2, _ := sr.FetchFlow("test2")
	sr.UpdateFlow("test2", testNodes2)
	t2a, _ := sr.FetchFlow("test2")
	if len(t2) == len(t2a) {
		t.Fatal("flow not updated correctly")
	}

	i2, ok := t2a["bar"].(*ast.IntegerLiteral)
	if !ok {
		t.Fatalf("property bar is incorrect type got=%T want=IntegerLiteral", t2a["bar"])
	}

	if i2.Value != 2 {
		t.Fatalf("property bar has incorrect value got=%d want=2", i2.Value)
	}

	sr.AddComponent("test3", testNodes)
	t3, _ := sr.FetchComponent("test3")
	sr.UpdateComponent("test3", testNodes2)
	t3a, _ := sr.FetchComponent("test3")
	if len(t3) == len(t3a) {
		t.Fatal("component not found in spec record")
	}

	i3, ok := t3a["bar"].(*ast.IntegerLiteral)
	if !ok {
		t.Fatalf("property bar is incorrect type got=%T want=IntegerLiteral", t3a["bar"])
	}

	if i3.Value != 2 {
		t.Fatalf("property bar has incorrect value got=%d want=2", i3.Value)
	}
}

func TestUpdateVar(t *testing.T) {
	sr := NewSpecRecord()

	testNodes := map[string]ast.Node{
		"bar": &ast.IntegerLiteral{Value: 1}}

	testNodes3 := map[string]ast.Node{
		"foo": &ast.FloatLiteral{Value: .01},
		"bar": &ast.IntegerLiteral{Value: 3}}

	sr.AddStock("test1", testNodes)
	sr.UpdateVar([]string{"foo", "test1", "bar"}, "STOCK", testNodes3["bar"])
	t1b, _ := sr.FetchStock("test1")

	ib, ok := t1b["bar"].(*ast.IntegerLiteral)
	if !ok {
		t.Fatalf("property bar is incorrect type got=%T want=IntegerLiteral", t1b["bar"])
	}

	if ib.Value != 3 {
		t.Fatalf("property bar has incorrect value got=%d want=3", ib.Value)
	}

	sr.AddFlow("test2", testNodes)
	sr.UpdateVar([]string{"foo", "test2", "bar"}, "FLOW", testNodes3["bar"])
	t2b, _ := sr.FetchFlow("test2")

	i2b, ok := t2b["bar"].(*ast.IntegerLiteral)
	if !ok {
		t.Fatalf("property bar is incorrect type got=%T want=IntegerLiteral", t2b["bar"])
	}

	if i2b.Value != 3 {
		t.Fatalf("property bar has incorrect value got=%d want=3", i2b.Value)
	}

	sr.AddComponent("test3", testNodes)
	sr.UpdateVar([]string{"foo", "test3", "bar"}, "COMPONENT", testNodes3["bar"])
	t3b, _ := sr.FetchComponent("test3")

	i3b, ok := t3b["bar"].(*ast.IntegerLiteral)
	if !ok {
		t.Fatalf("property bar is incorrect type got=%T want=IntegerLiteral", t3b["bar"])
	}

	if i3b.Value != 3 {
		t.Fatalf("property bar has incorrect value got=%d want=3", i3b.Value)
	}

	sr.AddGlobal("test4", &ast.IntegerLiteral{Value: 1})
	sr.UpdateVar([]string{"foo", "test4"}, "GLOBAL", &ast.IntegerLiteral{Value: 3})
	t4b, _ := sr.FetchGlobal("test4")

	i4b, ok := t4b.(*ast.IntegerLiteral)
	if !ok {
		t.Fatalf("constant is incorrect type got=%T want=IntegerLiteral", t4b)
	}

	if i4b.Value != 3 {
		t.Fatalf("constant has incorrect value got=%d want=3", i4b.Value)
	}
}

func TestSpecRecordInstance(t *testing.T) {
	sr := NewSpecRecord()

	testNodes := map[string]ast.Node{
		"bar": &ast.IntegerLiteral{Value: 1}}

	sr.AddStock("test", testNodes)
	t1, _ := sr.FetchStock("test")
	sr.AddInstance("test2", t1, "STOCK")
	t2, _ := sr.FetchStock("test2")
	t2["bar"] = &ast.IntegerLiteral{Value: 0}
	if t1["bar"] == t2["bar"] {
		t.Fatal("instance not added correctly")
	}
}

func TestSpecRecordFetching(t *testing.T) {
	sr := NewSpecRecord()

	testNodes := map[string]ast.Node{
		"bar": &ast.IntegerLiteral{Value: 1}}
	testNodes1 := map[string]ast.Node{
		"bar1": &ast.IntegerLiteral{Value: 1}}
	testNodes2 := map[string]ast.Node{
		"bar2": &ast.IntegerLiteral{Value: 1}}
	testNodes3 := map[string]ast.Node{
		"bar3": &ast.IntegerLiteral{Value: 1}}

	testNodes4 := map[string]ast.Node{
		"bar4": &ast.IntegerLiteral{Value: 1}}

	sr.AddStock("test", testNodes)
	s2, _ := sr.FetchVar([]string{"here", "test", "bar"}, "STOCK")

	if _, ok := s2.(*ast.IntegerLiteral); !ok {
		t.Fatalf("var not an IntegerLiteral got=%T", s2)
	}

	sr.AddFlow("test1", testNodes1)
	sr.AddStock("test2", testNodes2)
	sr.AddStock("test3", testNodes3)
	sr.AddComponent("test4", testNodes4)

	sr.AddConstant("foo", &ast.IntegerLiteral{Value: 1})
	sr.AddConstant("foo2", &ast.IntegerLiteral{Value: 1})

	all := sr.FetchAll()
	i := 0
	for k := range all {
		if k == "bar" || k == "bar1" || k == "bar2" || k == "bar3" || k == "bar4" || k == "foo" || k == "foo1" {
			i = i + 1
		}
	}
	if i != 6 {
		t.Fatalf("fetch all returned wrong number of vars got=%d", i)
	}
}

func TestSpecRecordTypes(t *testing.T) {
	sr := NewSpecRecord()

	ty0, _ := sr.GetStructType([]string{"this", "does", "not", "exist"})
	if ty0 != "NIL" {
		t.Fatalf("spec record did not return the correct type for a nil variable got=%s", ty0)
	}

	testNodes := map[string]ast.Node{
		"bar": &ast.IntegerLiteral{Value: 1}}

	sr.AddStock("test", testNodes)
	sr.Index("STOCK", "test")
	ty1, _ := sr.GetStructType([]string{"this", "test"})
	if ty1 != "STOCK" {
		t.Fatalf("spec record did not return the correct type for stock got=%s", ty1)
	}

	sr.AddFlow("test1", testNodes)
	sr.Index("FLOW", "test1")
	ty2, _ := sr.GetStructType([]string{"this", "test1"})
	if ty2 != "FLOW" {
		t.Fatalf("spec record did not return the correct type for flow got=%s", ty2)
	}

	sr.AddComponent("test2", testNodes)
	sr.Index("COMPONENT", "test2")
	ty3, _ := sr.GetStructType([]string{"this", "test2"})
	if ty3 != "COMPONENT" {
		t.Fatalf("spec record did not return the correct type for component got=%s", ty3)
	}

	sr.AddConstant("test3", testNodes["bar"])
	ty4, _ := sr.GetStructType([]string{"this", "test3"})
	if ty4 != "CONSTANT" {
		t.Fatalf("spec record did not return the correct type for constant got=%s", ty4)
	}

}

func TestSpecRecordIndex(t *testing.T) {
	sr := NewSpecRecord()

	testNodes := map[string]ast.Node{
		"bar": &ast.IntegerLiteral{Value: 1}}

	sr.AddStock("test", testNodes)
	sr.Index("STOCK", "test")

	if len(sr.Order) != 1 {
		t.Fatalf("SpecRecord has an incorrect index length got=%s", sr.Order)
	}

	if sr.Order[0][0] != "STOCK" || sr.Order[0][1] != "test" {
		t.Fatalf("SpecRecord has an incorrect index  got=%s", sr.Order[0])
	}
}
