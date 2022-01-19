package smt

import (
	"testing"
)

func TestGetStartsEnds(t *testing.T) {
	g := NewGenerator()
	data := make(map[int][]int)
	data[0] = []int{2, 3, 4}
	data[1] = []int{7, 8}
	data[2] = []int{-2, 6}

	if g.getStarts(data) != -2 {
		t.Fatalf("getStarts returned the wrong result. got=%d", g.getStarts(data))
	}

	if g.getEnds(data) != 8 {
		t.Fatalf("getEnds returned the wrong result. got=%d", g.getEnds(data))
	}
}

func TestParallelPermutations(t *testing.T) {
	g := NewGenerator()
	test1 := []string{"foo", "bar"}
	results1 := g.parallelPermutations(test1)

	test2 := []string{"foo", "bar", "fizz", "buzz"}
	results2 := g.parallelPermutations(test2)

	test3 := []string{"foo", "bar", "fizz", "buzz", "foosh"}
	results3 := g.parallelPermutations(test3)

	if len(results1) != 2 {
		t.Fatalf("wrong number of permutations on set 1got=%d", len(results1))
	}

	if len(results2) != 24 {
		t.Fatalf("wrong number of permutations on set 2 got=%d", len(results2))
	}

	if len(results3) != 120 {
		t.Fatalf("wrong number of permutations on set 3 got=%d", len(results3))
	}
}
