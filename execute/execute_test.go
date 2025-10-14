package execute

import (
	"fault/generator"
	"fault/listener"
	"fault/llvm"
	"fault/preprocess"
	"fault/smt/variables"
	"fault/swaps"
	"fault/types"
	"fmt"
	"log"
	"os"
	gopath "path"
	"path/filepath"
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

	err = model.Solve()

	if err != nil {
		t.Fatalf("SMT Solver failed to provide solution. got=%s", err)
	}

	if model.ResultValues["imports_fl3_vault_value_0"] == "" {
		t.Fatal("SMT Solver failed to provide solution.")
	}

	if model.ResultValues["imports_fl3_vault_value_0"] != "30.0" {
		t.Fatalf("Value of imports_fl3_vault_value_0 is incorrect. got=%s", model.ResultValues["imports_fl3_vault_value_0"])
	}
	if model.ResultValues["imports_fl3_vault_value_1"] != "40.0" {
		t.Fatalf("Value of imports_fl3_vault_value_1 is incorrect. got=%s", model.ResultValues["imports_fl3_vault_value_1"])
	}
	if model.ResultValues["imports_fl3_vault_value_2"] != "50.0" {
		t.Fatalf("Value of imports_fl3_vault_value_2 is incorrect. got=%s", model.ResultValues["imports_fl3_vault_value_2"])
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
	err := model.Solve()
	if err != nil {
		t.Fatalf("SMT Solver failed to provide solution. got=%s", err)
	}

}

func prepTest(smt string, uncertains map[string][]float64, unknowns []string, results map[string][]*variables.VarChange) *ModelChecker {
	ex := NewModelChecker()
	ex.LoadModel(smt, uncertains, unknowns)
	return ex
}

func TestFullSuite(t *testing.T) {
	// Run through all the tests in generator/testdata to check for errors
	var run = func(path string, fileInfo os.FileInfo, inpErr error) (err error) {
		fmt.Println(path)

		uncertains := make(map[string][]float64)
		unknowns := []string{}
		//extract the extension from the path
		filetype := filepath.Ext(path)
		if filetype != ".fspec" && filetype != ".fsystem" {
			return nil
		}
		data, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		d := string(data)
		fpath := gopath.Dir(path)
		flags := make(map[string]bool)
		flags["specType"] = (filetype == ".fspec")
		flags["testing"] = false
		flags["skipRun"] = false
		lstnr := listener.Execute(d, fpath, flags)
		if lstnr == nil {
			log.Fatal("Fault parser returned nil")
		}

		pre := preprocess.Execute(lstnr)

		ty := types.Execute(pre.Processed, pre)

		sw := swaps.NewPrecompiler(ty)
		tree := sw.Swap(ty.Checked)
		compiler := llvm.Execute(tree, ty.SpecStructs, lstnr.Uncertains, lstnr.Unknowns, sw.Alias, false)
		uncertains = compiler.RawInputs.Uncertains
		unknowns = compiler.RawInputs.Unknowns
		if !compiler.IsValid {
			return fmt.Errorf("Fault found nothing to run. Missing run block or start block.")
		}

		g := generator.Execute(compiler)
		ex := NewModelChecker()
		ex.LoadModel(g.SMT(), uncertains, unknowns)
		ok, err := ex.Check()
		if err != nil {
			return fmt.Errorf("model checker has failed: %s", err)
		}
		if !ok {
			return fmt.Errorf("Fault could not find a failure case.")
		}
		err = ex.Solve()
		if err != nil {
			return fmt.Errorf("error found fetching solution from solver: %s", err)
		}
		g.ResultLog.Results = ex.ResultValues
		g.ResultLog.Trace()
		g.ResultLog.Kill()
		g.ResultLog.Print()
		return nil
	}

	err := filepath.Walk("../generator/testdata", run)
	if err != nil {
		t.Fatalf("Error in full test suite: %s", err)
	}
}
