package name

import (
	"crypto/md5"
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func TestBlock(t *testing.T) {
	b := Block()
	b2 := Block()

	s := strings.Split(b, "-")
	s2 := strings.Split(b2, "-")

	if s[0] != "block" || s2[0] != "block" {
		t.Fatal("Block function does not return a correct block name")
	}

	i, err := strconv.Atoi(s[1])
	if err != nil {
		t.Fatalf("Block function does not return a correct block name got=%s want=block-1", b)
	}

	i2, err := strconv.Atoi(s2[1])
	if err != nil {
		t.Fatalf("Block function does not return a correct block name got=%s want=block-2", b2)
	}

	if i2 <= i {
		t.Fatal("Block function does iterate block name correctly")
	}
}

func TestAnonFunc(t *testing.T) {
	b := AnonFunc()
	b2 := AnonFunc()

	s := strings.Split(b, "-")
	s2 := strings.Split(b2, "-")

	if s[0] != "fn" || s2[0] != "fn" {
		t.Fatal("AnonFunc function does not return a correct name")
	}

	i, err := strconv.Atoi(s[1])
	if err != nil {
		t.Fatalf("AnonFunc function does not return a correct name got=%s want=fn-1", b)
	}

	i2, err := strconv.Atoi(s2[1])
	if err != nil {
		t.Fatalf("AnonFunc function does not return a correct name got=%s want=fn-2", b2)
	}

	if i2 <= i {
		t.Fatal("AnonFunc function does iterate name correctly")
	}
}

func TestAssert(t *testing.T) {
	b := Assert()
	b2 := Assert()

	s := strings.Split(b, "-")
	s2 := strings.Split(b2, "-")

	if s[0] != "__assert" || s2[0] != "__assert" {
		t.Fatal("Assert function does not return a correct name")
	}

	i, err := strconv.Atoi(s[1])
	if err != nil {
		t.Fatalf("Assert function does not return a correct name got=%s want=__assert-1", b)
	}

	i2, err := strconv.Atoi(s2[1])
	if err != nil {
		t.Fatalf("Assert function does not return a correct name got=%s want=__assert-2", b2)
	}

	if i2 <= i {
		t.Fatal("Assert function does iterate name correctly")
	}
}
func TestVar(t *testing.T) {
	b := Var("test")
	b2 := Var("test")

	s := strings.Split(b, "-")
	s2 := strings.Split(b2, "-")

	if s[0] != "test" || s2[0] != "test" {
		t.Fatal("Var function does not return a correct name")
	}

	i, err := strconv.Atoi(s[1])
	if err != nil {
		t.Fatalf("Var function does not return a correct name got=%s want=test-3", b)
	}

	i2, err := strconv.Atoi(s2[1])
	if err != nil {
		t.Fatalf("Var function does not return a correct name got=%s want=test-4", b2)
	}

	if i2 <= i {
		t.Fatal("Var function does iterate name correctly")
	}
}

func TestParallelGroup(t *testing.T) {
	data := []byte(fmt.Sprint("test", 0))
	expecting := fmt.Sprintf("%x", md5.Sum(data))
	results := ParallelGroup("test")

	if expecting != results {
		t.Fatalf("ParallelGroup does not return correct hash got=%s want=%s", results, expecting)
	}

	data2 := []byte(fmt.Sprint("test", 1))
	expecting2 := fmt.Sprintf("%x", md5.Sum(data2))
	results2 := ParallelGroup("test")

	if expecting2 != results2 {
		t.Fatal("ParallelGroup does not iterate correctly")
	}
}
