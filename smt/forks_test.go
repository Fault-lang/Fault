package smt

import "testing"

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
