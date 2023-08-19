package execute

import (
	resultlog "fault/smt/log"
	"fault/smt/variables"
	"testing"
)

func TestSMTOk(t *testing.T) {
	test := `(declare-fun imports_fl3_vault_value_0 () Real)
	(declare-fun imports_fl3_vault_value_1 () Real)
	(declare-fun imports_fl3_vault_value_2 () Real)(assert (= imports_fl3_vault_value_0 30.0))
	(assert (= imports_fl3_vault_value_1 (+ imports_fl3_vault_value_0 10.0)))
	(assert (= imports_fl3_vault_value_2 (+ imports_fl3_vault_value_1 10.0)))
	`
	model := prepTest(test, make(map[string][]float64), []string{}, map[string][]*variables.VarChange{})

	response, err := model.Check()

	if err != nil {
		t.Fatalf("SMT Solver failed on valid expression. got=%s", err)
	}

	if !response {
		t.Fatalf("SMT Solver failed on valid expression.")
	}

	solution, err := model.Solve()

	if err != nil {
		t.Fatalf("SMT Solver failed to provide solution. got=%s", err)
	}

	if solution["imports_fl3_vault_value"] == nil {
		t.Fatal("SMT Solver failed to provide solution.")
	}

	got := solution["imports_fl3_vault_value"].(*FloatTrace).Get()
	expected := map[int64]float64{0: 30.0, 1: 40.0, 2: 50.0}
	if got[0] != expected[0] {
		t.Fatalf("SMT Solver solution not expected. want=%f got=%f", expected[0], got[0])
	}
}

func TestProbability(t *testing.T) {
	test := `(declare-fun imports_fl3_vault_value_0 () Real)
	(declare-fun imports_fl3_vault_value_1 () Real)
	(declare-fun imports_fl3_vault_value_2 () Real)(assert (= imports_fl3_vault_value_0 30.0))
	(assert (= imports_fl3_vault_value_1 (+ imports_fl3_vault_value_0 10.0)))
	(assert (= imports_fl3_vault_value_2 (+ imports_fl3_vault_value_1 10.0)))
	`
	uncertains := make(map[string][]float64)
	uncertains["imports_fl3_vault_value"] = []float64{30.0, 5}

	model := prepTest(test, uncertains, []string{}, map[string][]*variables.VarChange{})

	model.Check()
	solution, _ := model.Solve()
	filter := model.Filter(solution)
	got := filter["imports_fl3_vault_value"].(*FloatTrace).GetWeights()
	expected := map[int64]float64{0: 0.07978845608028654, 1: 0.010798193302637605, 2: 2.6766045152977058e-05}
	if got[0] != expected[0] {
		t.Fatalf("Probability distribution not weighting correctly. want=%f got=%f", expected[0], got[0])
	}

}

func TestEventLog(t *testing.T) {
	test := `(declare-fun imports_fl3_vault_value_0 () Real)
	(declare-fun imports_fl3_vault_value_1 () Real)
	(declare-fun imports_fl3_vault_value_2 () Real)(assert (= imports_fl3_vault_value_0 30.0))
	(assert (= imports_fl3_vault_value_1 (+ imports_fl3_vault_value_0 10.0)))
	(assert (= imports_fl3_vault_value_2 (+ imports_fl3_vault_value_1 10.0)))
	`
	model := prepTest(test, make(map[string][]float64), []string{}, map[string][]*variables.VarChange{})

	model.Log = resultlog.NewLog()
	model.Log.Add(resultlog.NewInit(0, "", "imports_fl3_vault_value_0"))
	model.Log.Add(resultlog.NewInit(0, "", "imports_fl3_vault_value_1"))
	model.Log.Add(resultlog.NewInit(0, "", "imports_fl3_vault_value_2"))
	model.Log.Add(resultlog.NewChange(0, "", "imports_fl3_vault_value_1"))
	model.Log.Add(resultlog.NewChange(0, "", "imports_fl3_vault_value_2"))

	response, err := model.Check()

	if err != nil {
		t.Fatalf("SMT Solver failed on valid expression. got=%s", err)
	}

	if !response {
		t.Fatalf("SMT Solver failed on valid expression.")
	}

	solution, err := model.Solve()

	if err != nil {
		t.Fatalf("SMT Solver failed to provide solution. got=%s", err)
	}

	if solution["imports_fl3_vault_value"] == nil {
		t.Fatal("SMT Solver failed to provide solution.")
	}

	model.mapToLog("imports_fl3_vault_value", solution["imports_fl3_vault_value"])

	if model.Log.Events[0].String() != "0,INIT,,imports_fl3_vault_value_0,,30,\n" {
		t.Fatalf("Incorrect event log format at index 0 got=%s", model.Log.Events[0].String())
	}

	if model.Log.Events[1].String() != "0,INIT,,imports_fl3_vault_value_1,,,\n" {
		t.Fatalf("Incorrect event log format at index 1 got=%s", model.Log.Events[1].String())
	}

	if model.Log.Events[3].String() != "0,CHANGE,,imports_fl3_vault_value_1,,40,\n" {
		t.Fatalf("Incorrect event log format at index 3 got=%s", model.Log.Events[3].String())
	}
}

func TestEval(t *testing.T) {
	mc := NewModelChecker()
	mc.Log = resultlog.NewLog()

	a := &resultlog.Assert{
		Left:  &resultlog.BoolClause{Value: true},
		Right: &resultlog.BoolClause{Value: true},
		Op:    "and",
	}
	if !mc.Eval(a) {
		t.Fatalf("Incorrect evaluation got=%v", mc.Eval(a))
	}

	a1 := &resultlog.Assert{
		Left:  &resultlog.BoolClause{Value: true},
		Right: &resultlog.BoolClause{Value: false},
		Op:    "=",
	}
	if mc.Eval(a1) {
		t.Fatalf("Incorrect evaluation got=%v", mc.Eval(a1))
	}

	a2 := &resultlog.Assert{
		Left:  &resultlog.FlClause{Value: 2.0},
		Right: &resultlog.FlClause{Value: 5.0},
		Op:    ">",
	}
	if mc.Eval(a2) {
		t.Fatalf("Incorrect evaluation got=%v", mc.Eval(a2))
	}
}

func TestEvalAmbiguous(t *testing.T) {
	mc := NewModelChecker()
	mc.Log = resultlog.NewLog()

	a := &resultlog.Assert{
		Left:  &resultlog.BoolClause{Value: true},
		Right: &resultlog.BoolClause{Value: true},
		Op:    "=",
	}
	if !mc.EvalAmbiguous(a) {
		t.Fatalf("Incorrect evaluation got=%v", mc.Eval(a))
	}

	a1 := &resultlog.Assert{
		Left:  &resultlog.FlClause{Value: 2.0},
		Right: &resultlog.FlClause{Value: 2.0},
		Op:    "=",
	}
	if !mc.EvalAmbiguous(a) {
		t.Fatalf("Incorrect evaluation got=%v", mc.Eval(a1))
	}

}

func TestEvalClause(t *testing.T) {
	mc := NewModelChecker()
	mc.Log = resultlog.NewLog()

	mc.ResultValues["test_var_foo"] = "false"

	cf := &resultlog.FlClause{
		Value: 5.0,
	}

	a, err := mc.EvalClause(cf)

	if err == nil {
		t.Fatalf("Incorrect evaluation got=%v", a)
	}

	cb := &resultlog.BoolClause{
		Value: true,
	}

	a1, err := mc.EvalClause(cb)

	if !a1 {
		t.Fatalf("Incorrect evaluation got=%v", a1)
	}

	cs := &resultlog.StringClause{
		Value: "test_var_foo",
	}

	a2, err := mc.EvalClause(cs)

	if a2 {
		t.Fatalf("Incorrect evaluation got=%v", a2)
	}
}

func prepTest(smt string, uncertains map[string][]float64, unknowns []string, results map[string][]*variables.VarChange) *ModelChecker {
	ex := NewModelChecker()
	ex.LoadModel(smt, uncertains, unknowns, results, &resultlog.ResultLog{})
	return ex
}
