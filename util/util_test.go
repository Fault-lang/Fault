package util

import "testing"

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
