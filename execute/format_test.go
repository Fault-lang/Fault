package execute

import (
	"fault/smt/forks"
	"testing"
)

func TestPhis(t *testing.T) {
	test := make(map[string]Scenario)
	test["test_value"] = &FloatTrace{
		results: map[int16]float64{0: 1.0, 1: 2.3, 2: 3.0, 3: 4.0, 4: 2.3},
		weights: map[int16]float64{},
	}
	test["test_value_foo"] = &FloatTrace{
		results: map[int16]float64{0: 5.0, 1: 20.3, 3: 34.0, 4: 34.0, 5: 34.0, 6: 20.3, 7: 20.3},
		weights: map[int16]float64{},
	}

	phis := forks.InitFork()
	phis.Choices["choice1"] = []string{"branch1", "branch2"}
	phis.AddVar("test1_branch1", "test_value", "test_value_0", &forks.Var{Base: "test_value", Last: false, Phi: "4"})
	phis.AddVar("test1_branch1", "test_value", "test_value_1", &forks.Var{Base: "test_value", Last: true, Phi: "4"})
	phis.AddVar("test1_branch2", "test_value", "test_value_2", &forks.Var{Base: "test_value", Last: false, Phi: "4"})
	phis.AddVar("test1_branch2", "test_value", "test_value_3", &forks.Var{Base: "test_value", Last: true, Phi: "4"})

	phis.Choices["choice2"] = []string{"branch1", "branch2"}
	phis.AddVar("test2_branch1", "test_value_foo", "test_value_foo_0", &forks.Var{Base: "test_value_foo", Last: false, Phi: "7"})
	phis.AddVar("test2_branch1", "test_value_foo", "test_value_foo_1", &forks.Var{Base: "test_value_foo", Last: false, Phi: "7"})
	phis.AddVar("test2_branch1", "test_value_foo", "test_value_foo_2", &forks.Var{Base: "test_value_foo", Last: false, Phi: "7"})
	phis.AddVar("test2_branch1", "test_value_foo", "test_value_foo_3", &forks.Var{Base: "test_value_foo", Last: false, Phi: "7"})
	phis.AddVar("test2_branch1", "test_value_foo", "test_value_foo_4", &forks.Var{Base: "test_value_foo", Last: false, Phi: "7"})
	phis.AddVar("test2_branch2", "test_value_foo", "test_value_foo_5", &forks.Var{Base: "test_value_foo", Last: false, Phi: "7"})
	phis.AddVar("test2_branch2", "test_value_foo", "test_value_foo_6", &forks.Var{Base: "test_value_foo", Last: false, Phi: "7"})

	mc := NewModelChecker()
	mc.LoadMeta(phis)

	deadVars := mc.DeadVariables()
	v := deadBranches("test_value", test["test_value"], deadVars)
	if _, ok := v.(*FloatTrace).Index(3); ok {
		t.Fatal("value at index 3 of variable test_value was not removed")
	}
	if _, ok := v.(*FloatTrace).Index(1); !ok {
		t.Fatal("value at index 1 of variable test_value was removed")
	}

	if _, ok := v.(*FloatTrace).Index(4); ok {
		t.Fatal("phi at index 4 was not removed")
	}

	v2 := deadBranches("test_value_foo", test["test_value_foo"], deadVars)
	if _, ok := v2.(*FloatTrace).Index(1); ok {
		t.Fatal("value at index 1 of variable test_value_foo was not removed")
	}
	if _, ok := v2.(*FloatTrace).Index(5); !ok {
		t.Fatal("value at index 5 of variable test_value_foo was removed")
	}
	if _, ok := v2.(*FloatTrace).Index(7); ok {
		t.Fatal("phi at index 7 not was removed")
	}
}

// func TestMultiPhis(t *testing.T) {
// 	test := make(map[string]Scenario)
// 	test["test_value"] = &FloatTrace{
// 		results: map[int16]float64{0: 1.0, 1: 1.5, 2: 2.3, 3: 4.0, 4: 2.3, 5: 3.4, 6: 3.8, 7: 5.0, 8: 5.2, 9: 3.8},
// 		weights: map[int16]float64{},
// 	}

// 	phis := []forks.Fork{{
// 		"test_value": []*forks.Choice{
// 			{
// 				Base:   "test_value",
// 				Values: []int16{1, 2},
// 			},
// 			{
// 				Base:   "test_value",
// 				Values: []int16{3},
// 			},
// 		}},
// 		{
// 			"test_value": []*forks.Choice{
// 				{
// 					Base:   "test_value",
// 					Values: []int16{5, 6},
// 				},
// 				{
// 					Base:   "test_value",
// 					Values: []int16{7, 8},
// 				},
// 			}},
// 	}

// 	mc := NewModelChecker()
// 	mc.LoadMeta(phis)

// 	v := deadBranches("test_value", test["test_value"], mc.forks)
// 	if _, ok := v.(*FloatTrace).Index(3); ok {
// 		t.Fatal("value at index 3 of variable test_value was not removed")
// 	}
// 	if _, ok := v.(*FloatTrace).Index(1); !ok {
// 		t.Fatal("value at index 1 of variable test_value was removed")
// 	}

// 	if _, ok := v.(*FloatTrace).Index(4); ok {
// 		t.Fatal("phi at index 4 not was removed")
// 	}

// 	if _, ok := v.(*FloatTrace).Index(6); !ok {
// 		t.Fatal("value at index 6 of variable test_value was removed")
// 	}

// 	if _, ok := v.(*FloatTrace).Index(8); ok {
// 		t.Fatal("value at index 8 of variable test_value was not removed")
// 	}
// 	if _, ok := v.(*FloatTrace).Index(9); ok {
// 		t.Fatal("phi at index 9 not was removed")
// 	}
// }
