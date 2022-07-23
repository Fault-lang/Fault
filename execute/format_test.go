package execute

import (
	"fault/smt"
	"testing"
)

func TestPhis(t *testing.T) {
	test := make(map[string]Scenario)
	test["test_value"] = &FloatTrace{
		results: map[int16]float64{0: 1.0, 1: 2.3, 3: 4.0, 4: 2.3},
		weights: map[int16]float64{},
	}
	test["test_value_foo"] = &FloatTrace{
		results: map[int16]float64{0: 5.0, 1: 20.3, 3: 34.0, 4: 34.0, 5: 34.0, 6: 20.3, 7: 20.3},
		weights: map[int16]float64{},
	}

	phis := []smt.Fork{{
		"test_value": []*smt.Choice{
			{
				Base:   "test_value",
				Values: []int16{1},
			},
			{
				Base:   "test_value",
				Values: []int16{3},
			},
		},
		"test_value_foo": []*smt.Choice{
			{
				Base:   "test_value_foo",
				Values: []int16{1, 3, 4},
			},
			{
				Base:   "test_value_foo",
				Values: []int16{5, 6},
			},
		}}}

	mc := NewModelChecker()
	mc.LoadMeta(phis)

	v := deadBranches("test_value", test["test_value"], mc.forks)
	if _, ok := v.(*FloatTrace).Index(3); ok {
		t.Fatal("value at index 3 of variable test_value was not removed")
	}
	if _, ok := v.(*FloatTrace).Index(1); !ok {
		t.Fatal("value at index 1 of variable test_value was removed")
	}

	if _, ok := v.(*FloatTrace).Index(4); ok {
		t.Fatal("phi at index 4 not was removed")
	}

	v2 := deadBranches("test_value_foo", test["test_value_foo"], mc.forks)
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

/*func TestPlainTrail(t *testing.T) {
	test := make(map[string]Scenario)
	test["test_value"] = &FloatTrace{
		results: map[int64]float64{0: 1.0, 1: 2.3, 3: 4.0, 7: 2.3},
		weights: map[int64]float64{0: .01, 1: .30, 3: .00, 7: .20},
	}
	test["test_value_foo"] = &FloatTrace{
		results: map[int64]float64{0: 5.0, 1: 20.3, 3: 34.0, 7: 20.3},
		weights: map[int64]float64{0: .001, 1: .30, 3: .001, 7: .20},
	}

	row1 := generateRows(test["test_value"])
	if row1[0] != "-> 1.000000 (0.010000)" || row1[len(row1)-1] != "-> 2.300000 (0.200000)" {
		t.Fatalf(fmt.Sprintf("incorrect row returned. got=%s", row1))
	}

	row2 := generateRows(test["test_value_foo"])
	if row2[0] != "-> 5.000000 (0.001000)" || row2[len(row2)-1] != "-> 20.300000 (0.200000)" {
		t.Fatalf(fmt.Sprintf("incorrect row returned. got=%s", row2))
	}

}

func TestPlainTrailNoWeights(t *testing.T) {
	test := make(map[string]Scenario)
	test["test_value"] = &FloatTrace{
		results: map[int64]float64{0: 1.0, 1: 2.3, 3: 4.0, 7: 2.3},
		weights: make(map[int64]float64),
	}
	test["test_value_foo"] = &FloatTrace{
		results: map[int64]float64{0: 5.0, 1: 20.3, 3: 34.0, 7: 20.3},
		weights: make(map[int64]float64),
	}

	row1 := generateRows(test["test_value"])
	if row1[0] != "-> 1.000000" || row1[len(row1)-1] != "-> 2.300000" {
		t.Fatalf(fmt.Sprintf("incorrect row returned. got=%s", row1))
	}

	row2 := generateRows(test["test_value_foo"])
	if row2[0] != "-> 5.000000" || row2[len(row2)-1] != "-> 20.300000" {
		t.Fatalf(fmt.Sprintf("incorrect row returned. got=%s", row2))
	}

}

func TestEndStateBranch(t *testing.T) {
	trails := map[string][]string{"cond_true": []string{"test_value_foo_1", "test_value_1"},
		"cond_false": []string{"test_value_foo_3", "test_value_3"}}

	end1 := endStatesBranch(trails["cond_true"])
	if end1["test_value_foo"] != 1 {
		t.Fatalf(fmt.Sprintf("incorrect end state returned. got=%d", end1["test_value_foo"]))
	}

	if end1["test_value"] != 1 {
		t.Fatalf(fmt.Sprintf("incorrect end state returned. got=%d", end1["test_value"]))
	}

	end2 := endStatesBranch(trails["cond_false"])
	if end2["test_value_foo"] != 3 {
		t.Fatalf(fmt.Sprintf("incorrect end state returned. got=%d", end2["test_value_foo"]))
	}

	if end2["test_value"] != 3 {
		t.Fatalf(fmt.Sprintf("incorrect end state returned. got=%d", end2["test_value"]))
	}

}

func TestFilterBranch(t *testing.T) {
	test := make(map[string]Scenario)
	test["test_value"] = &FloatTrace{
		results: map[int64]float64{0: 1.0, 1: 2.3, 3: 4.0, 4: 2.3, 7: 2.3},
		weights: map[int64]float64{0: .01, 1: .30, 3: .00, 4: .30, 7: .20},
	}
	test["test_value_foo"] = &FloatTrace{
		results: map[int64]float64{0: 5.0, 1: 20.3, 3: 34.0, 4: 20.3, 7: 20.3},
		weights: map[int64]float64{0: .001, 1: .30, 3: .001, 4: .30, 7: .20},
	}

	trails := map[string]map[string][]string{"cond": {"true": []string{"test_value_foo_1", "test_value_1"},
		"false": []string{"test_value_foo_3", "test_value_3"}}}

	filter := definePath(test, trails)

	if v, ok := filter["test_value"].(*FloatTrace).Index(3); ok {
		t.Fatalf(fmt.Sprintf("value for test_value not removed from scenario. got=%f", v))
	}

	if v, ok := filter["test_value_foo"].(*FloatTrace).Index(3); ok {
		t.Fatalf(fmt.Sprintf("value for test_value_foo not removed from scenario. got=%f", v))
	}

	if v, ok := filter["test_value"].(*FloatTrace).Index(1); !ok {
		t.Fatalf(fmt.Sprintf("value for test_value not removed from scenario. got=%f", v))
	}

	if v, ok := filter["test_value_foo"].(*FloatTrace).Index(1); !ok {
		t.Fatalf(fmt.Sprintf("value for test_value_foo not removed from scenario. got=%f", v))
	}

}

/*func TestFilterBranchParallels(t *testing.T) {
	test := make(map[string]Scenario)
	test["test_value"] = &FloatTrace{
		results: map[int64]float64{0: 1.0, 1: 2.3, 3: 4.0, 4: 2.3, 5: 1.0, 7: 2.3},
		weights: map[int64]float64{0: .01, 1: .30, 3: .00, 4: .30, 5: 0, 7: .20},
	}
	test["test_value_foo"] = &FloatTrace{
		results: map[int64]float64{0: 5.0, 1: 20.3, 3: 34.0, 4: 20.3, 5: 1.5, 7: 20.3},
		weights: map[int64]float64{0: .001, 1: .30, 3: .001, 4: .30, 5: .1, 7: .20},
	}

	trails := map[string]map[string][]string{"cond": {"option_1": []string{"test_value_foo_1", "test_value_1"},
		"option_2": []string{"test_value_foo_3", "test_value_3"}, "option_3": []string{"test_value_foo_5", "test_value_5"}}}

	filter := definePath(test, trails)

	if v, ok := filter["test_value"].(*FloatTrace).Index(3); ok {
		t.Fatalf(fmt.Sprintf("value for test_value third position not removed from scenario. got=%f", v))
	}

	if v, ok := filter["test_value_foo"].(*FloatTrace).Index(3); ok {
		t.Fatalf(fmt.Sprintf("value for test_value_foo third position not removed from scenario. got=%f", v))
	}

	if v, ok := filter["test_value"].(*FloatTrace).Index(5); ok {
		t.Fatalf(fmt.Sprintf("value for test_value fifth position not removed from scenario. got=%f", v))
	}

	if v, ok := filter["test_value_foo"].(*FloatTrace).Index(5); ok {
		t.Fatalf(fmt.Sprintf("value for test_value_foo fifth position not removed from scenario. got=%f", v))
	}

	if _, ok := filter["test_value"].(*FloatTrace).Index(1); !ok {
		t.Fatalf("value for test_value was removed from scenario.")
	}

	if _, ok := filter["test_value_foo"].(*FloatTrace).Index(1); !ok {
		t.Fatalf("value for test_value_foo was removed from scenario")
	}

}*/
