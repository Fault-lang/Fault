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
	phis.Choices["choice1"] = []string{"test1_branch1", "test1_branch2"}
	phis.AddVar("test1_branch1", "test_value", "test_value_0", forks.NewVar("test_value", false, "choice1", "4"))
	phis.AddVar("test1_branch1", "test_value", "test_value_1", forks.NewVar("test_value", true, "choice1", "4"))
	phis.AddVar("test1_branch2", "test_value", "test_value_2", forks.NewVar("test_value", false, "choice1", "4"))
	phis.AddVar("test1_branch2", "test_value", "test_value_3", forks.NewVar("test_value", true, "choice1", "4"))

	phis.Choices["choice2"] = []string{"test2_branch1", "test2_branch2"}
	phis.AddVar("test2_branch1", "test_value_foo", "test_value_foo_0", forks.NewVar("test_value_foo", false, "choice2", "7"))
	phis.AddVar("test2_branch1", "test_value_foo", "test_value_foo_1", forks.NewVar("test_value_foo", false, "choice2", "7"))
	phis.AddVar("test2_branch1", "test_value_foo", "test_value_foo_2", forks.NewVar("test_value_foo", false, "choice2", "7"))
	phis.AddVar("test2_branch1", "test_value_foo", "test_value_foo_3", forks.NewVar("test_value_foo", false, "choice2", "7"))
	phis.AddVar("test2_branch1", "test_value_foo", "test_value_foo_4", forks.NewVar("test_value_foo", true, "choice2", "7"))
	phis.AddVar("test2_branch2", "test_value_foo", "test_value_foo_5", forks.NewVar("test_value_foo", false, "choice2", "7"))
	phis.AddVar("test2_branch2", "test_value_foo", "test_value_foo_6", forks.NewVar("test_value_foo", true, "choice2", "7"))

	mc := NewModelChecker()

	mc.ResultValues["test_value_0"] = "1.0"
	mc.ResultValues["test_value_1"] = "2.3"
	mc.ResultValues["test_value_2"] = "3.0"
	mc.ResultValues["test_value_3"] = "4.0"
	mc.ResultValues["test_value_4"] = "2.3"

	mc.ResultValues["test_value_foo_0"] = "5.0"
	mc.ResultValues["test_value_foo_1"] = "20.3"
	mc.ResultValues["test_value_foo_3"] = "34.0"
	mc.ResultValues["test_value_foo_4"] = "34.0"
	mc.ResultValues["test_value_foo_5"] = "34.0"
	mc.ResultValues["test_value_foo_6"] = "20.3"
	mc.ResultValues["test_value_foo_7"] = "20.3"
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
		t.Fatal("phi at index 7 was not removed")
	}
}

func TestMultiPhis(t *testing.T) {
	test := make(map[string]Scenario)
	test["test_value"] = &FloatTrace{
		results: map[int16]float64{0: 1.0, 1: 1.5, 2: 2.3, 3: 4.0, 4: 2.3, 5: 3.4, 6: 3.8, 7: 5.0, 8: 5.2, 9: 3.8},
		weights: map[int16]float64{},
	}

	phis := forks.InitFork()
	phis.Choices["choice1"] = []string{"choice1_branch1", "choice1_branch2"}
	phis.AddVar("choice1_branch1", "test_value", "test_value_1", forks.NewVar("test_value", false, "choice1", "4"))
	phis.AddVar("choice1_branch1", "test_value", "test_value_2", forks.NewVar("test_value", true, "choice1", "4"))
	phis.AddVar("choice1_branch2", "test_value", "test_value_3", forks.NewVar("test_value", true, "choice1", "4"))
	phis.Choices["choice2"] = []string{"choice2_branch1", "choice2_branch2"}
	phis.AddVar("choice2_branch1", "test_value", "test_value_5", forks.NewVar("test_value", false, "choice2", "9"))
	phis.AddVar("choice2_branch1", "test_value", "test_value_6", forks.NewVar("test_value", true, "choice2", "9"))
	phis.AddVar("choice2_branch2", "test_value", "test_value_7", forks.NewVar("test_value", false, "choice2", "9"))
	phis.AddVar("choice2_branch2", "test_value", "test_value_8", forks.NewVar("test_value", true, "choice2", "9"))

	mc := NewModelChecker()
	mc.LoadMeta(phis)
	mc.ResultValues["test_value_0"] = "1.0"
	mc.ResultValues["test_value_1"] = "1.5"
	mc.ResultValues["test_value_2"] = "2.3"
	mc.ResultValues["test_value_3"] = "4.0"
	mc.ResultValues["test_value_4"] = "2.3"
	mc.ResultValues["test_value_5"] = "3.4"
	mc.ResultValues["test_value_6"] = "3.8"
	mc.ResultValues["test_value_7"] = "5.0"
	mc.ResultValues["test_value_8"] = "5.2"
	mc.ResultValues["test_value_9"] = "3.8"

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

	if _, ok := v.(*FloatTrace).Index(6); !ok {
		t.Fatal("value at index 6 of variable test_value was removed")
	}

	if _, ok := v.(*FloatTrace).Index(8); ok {
		t.Fatal("value at index 8 of variable test_value was not removed")
	}
	if _, ok := v.(*FloatTrace).Index(9); ok {
		t.Fatal("phi at index 9 was not removed")
	}
}

func TestNestledPhis(t *testing.T) {
	test := make(map[string]Scenario)
	test["test_value"] = &FloatTrace{
		results: map[int16]float64{0: 1.0, 1: 1.5, 2: 2.3,
			3: 4.0, 4: 2.3, 5: 3.4,
			6: 3.8, 7: 5.0, 8: 5.2, 9: 3.4},
		weights: map[int16]float64{},
	}

	phis := forks.InitFork()
	phis.Choices["choice1"] = []string{"choice1_branch1", "choice1_branch2"}
	phis.AddVar("choice1_branch1", "test_value", "test_value_1", forks.NewVar("test_value", false, "choice1", "9"))
	phis.AddVar("choice1_branch1", "test_value", "test_value_2", forks.NewVar("test_value", true, "choice1", "9"))
	phis.AddVar("choice1_branch2", "test_value", "test_value_3", forks.NewVar("test_value", false, "choice1", "9"))
	phis.AddVar("choice1_branch2", "test_value", "test_value_4", forks.NewVar("test_value", false, "choice1", "9"))
	phis.AddVar("choice1_branch2", "test_value", "test_value_5", forks.NewVar("test_value", true, "choice1", "9"))
	phis.AddVar("choice1_branch2", "test_value", "test_value_6", forks.NewVar("test_value", false, "choice1", "9"))
	phis.AddVar("choice1_branch2", "test_value", "test_value_7", forks.NewVar("test_value", false, "choice1", "9"))
	phis.AddVar("choice1_branch2", "test_value", "test_value_8", forks.NewVar("test_value", true, "choice1", "9"))

	phis.Choices["choice2"] = []string{"choice2_branch1", "choice2_branch2"}
	phis.AddVar("choice2_branch1", "test_value", "test_value_3", forks.NewVar("test_value", false, "choice2", "9"))
	phis.AddVar("choice2_branch1", "test_value", "test_value_4", forks.NewVar("test_value", false, "choice2", "9"))
	phis.AddVar("choice2_branch1", "test_value", "test_value_5", forks.NewVar("test_value", true, "choice2", "9"))
	phis.AddVar("choice2_branch2", "test_value", "test_value_6", forks.NewVar("test_value", false, "choice2", "9"))
	phis.AddVar("choice2_branch2", "test_value", "test_value_7", forks.NewVar("test_value", false, "choice2", "9"))
	phis.AddVar("choice2_branch2", "test_value", "test_value_8", forks.NewVar("test_value", true, "choice2", "9"))

	mc := NewModelChecker()
	mc.LoadMeta(phis)
	mc.ResultValues["test_value_0"] = "1.0"
	mc.ResultValues["test_value_1"] = "1.5"
	mc.ResultValues["test_value_2"] = "2.3"
	mc.ResultValues["test_value_3"] = "4.0"
	mc.ResultValues["test_value_4"] = "2.3"
	mc.ResultValues["test_value_5"] = "3.4"
	mc.ResultValues["test_value_6"] = "3.8"
	mc.ResultValues["test_value_7"] = "5.0"
	mc.ResultValues["test_value_8"] = "5.2"
	mc.ResultValues["test_value_9"] = "3.4"

	deadVars := mc.DeadVariables()
	v := deadBranches("test_value", test["test_value"], deadVars)
	if _, ok := v.(*FloatTrace).Index(2); ok {
		t.Fatal("value at index 2 of variable test_value was not removed")
	}
	if _, ok := v.(*FloatTrace).Index(4); !ok {
		t.Fatal("value at index 4 of variable test_value was removed")
	}

	if _, ok := v.(*FloatTrace).Index(7); ok {
		t.Fatal("phi at index 7 was not removed")
	}

	if _, ok := v.(*FloatTrace).Index(5); !ok {
		t.Fatal("value at index 5 of variable test_value was removed")
	}

	if _, ok := v.(*FloatTrace).Index(8); ok {
		t.Fatal("value at index 8 of variable test_value was not removed")
	}
	if _, ok := v.(*FloatTrace).Index(9); ok {
		t.Fatal("phi at index 9 was not removed")
	}
}

func TestMultiVarPhis(t *testing.T) {
	test := make(map[string]Scenario)
	test["test_value"] = &FloatTrace{
		results: map[int16]float64{0: 1.0, 1: 1.5, 2: 2.3, 3: 4.0, 4: 2.3},
		weights: map[int16]float64{},
	}

	test["test_value_foo"] = &FloatTrace{
		results: map[int16]float64{0: 2.0, 1: 2.5, 2: 8.5, 3: 6.0, 4: 5.6, 5: 7.0, 6: 8.5},
		weights: map[int16]float64{},
	}

	phis := forks.InitFork()
	phis.Choices["choice1"] = []string{"choice1_branch1", "choice1_branch2"}
	phis.AddVar("choice1_branch1", "test_value", "test_value_1", forks.NewVar("test_value", false, "choice1", "4"))
	phis.AddVar("choice1_branch1", "test_value", "test_value_2", forks.NewVar("test_value", true, "choice1", "4"))
	phis.AddVar("choice1_branch2", "test_value", "test_value_3", forks.NewVar("test_value", true, "choice1", "4"))

	phis.AddVar("choice1_branch1", "test_value_foo", "test_value_foo_0", forks.NewVar("test_value_foo", false, "choice1", "6"))
	phis.AddVar("choice1_branch1", "test_value_foo", "test_value_foo_1", forks.NewVar("test_value_foo", false, "choice1", "6"))
	phis.AddVar("choice1_branch1", "test_value_foo", "test_value_foo_2", forks.NewVar("test_value_foo", true, "choice1", "6"))
	phis.AddVar("choice1_branch2", "test_value_foo", "test_value_foo_3", forks.NewVar("test_value_foo", false, "choice1", "6"))
	phis.AddVar("choice1_branch2", "test_value_foo", "test_value_foo_4", forks.NewVar("test_value_foo", false, "choice1", "6"))
	phis.AddVar("choice1_branch2", "test_value_foo", "test_value_foo_5", forks.NewVar("test_value_foo", true, "choice1", "6"))

	mc := NewModelChecker()
	mc.LoadMeta(phis)
	mc.ResultValues["test_value_0"] = "1.0"
	mc.ResultValues["test_value_1"] = "1.5"
	mc.ResultValues["test_value_2"] = "2.3"
	mc.ResultValues["test_value_3"] = "4.0"
	mc.ResultValues["test_value_4"] = "2.3"

	mc.ResultValues["test_value_foo_0"] = "2.0"
	mc.ResultValues["test_value_foo_1"] = "2.5"
	mc.ResultValues["test_value_foo_2"] = "8.5"
	mc.ResultValues["test_value_foo_3"] = "6.0"
	mc.ResultValues["test_value_foo_4"] = "5.6"
	mc.ResultValues["test_value_foo_5"] = "7.0"
	mc.ResultValues["test_value_foo_6"] = "8.5"

	deadVars := mc.DeadVariables()
	v := deadBranches("test_value", test["test_value"], deadVars)
	if _, ok := v.(*FloatTrace).Index(3); ok {
		t.Fatal("value at index 3 of variable test_value was not removed")
	}
	if _, ok := v.(*FloatTrace).Index(1); !ok {
		t.Fatal("value at index 4 of variable test_value was removed")
	}

	v = deadBranches("test_value_foo", test["test_value"], deadVars)
	if _, ok := v.(*FloatTrace).Index(4); ok {
		t.Fatal("phi at index 4 was not removed")
	}

	if _, ok := v.(*FloatTrace).Index(0); !ok {
		t.Fatal("value at index 0 of variable test_value was removed")
	}

	if _, ok := v.(*FloatTrace).Index(5); ok {
		t.Fatal("value at index 5 of variable test_value was not removed")
	}
	if _, ok := v.(*FloatTrace).Index(6); ok {
		t.Fatal("phi at index 6 was not removed")
	}
}
