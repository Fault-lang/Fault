package util

import (
	"os"
	"testing"
)

func TestFilepath(t *testing.T) {
	var host string
	var ok bool

	if host, ok = os.LookupEnv("FAULT_HOST"); !ok {
		host = ""
	}
	os.Setenv("FAULT_HOST", "/Users/test/file/system")
	filepath1 := "../test.spec"
	filepath1a := Filepath(filepath1)
	if filepath1a != "/Users/test/file/test.spec" {
		t.Fatalf("filepath not correct. want=/Users/test/file/test.spec got=%s", filepath1a)
	}

	filepath2 := "../../test.spec"
	filepath2a := Filepath(filepath2)
	if filepath2a != "/Users/test/test.spec" {
		t.Fatalf("filepath not correct. want=/Users/test/test.spec got=%s", filepath2a)
	}

	filepath3 := "../../../test.spec"
	filepath3a := Filepath(filepath3)
	if filepath3a != "/Users/test.spec" {
		t.Fatalf("filepath not correct. want=/Users/test.spec got=%s", filepath3a)
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
