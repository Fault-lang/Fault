package preprocess

import "testing"

func TestFetchVarError(t *testing.T) {
	sr := NewSpecRecord()
	sr.SpecName = "test"
	test := []string{"test", "should", "fail"}
	_, err := sr.FetchVar(test, "STOCK")
	if err == nil {
		t.Fatal("test failed to produce an error")
	}

	if err.Error() != "no stock found with name should in spec test" {
		t.Fatalf("error message did not match got=%s", err.Error())
	}

	_, err = sr.FetchVar(test, "FLOW")
	if err == nil {
		t.Fatal("test failed to produce an error")
	}

	if err.Error() != "no flow found with name should in spec test" {
		t.Fatalf("error message did not match got=%s", err.Error())
	}

	_, err = sr.FetchVar(test, "COMPONENT")
	if err == nil {
		t.Fatal("test failed to produce an error")
	}

	if err.Error() != "no component found with name should in spec test" {
		t.Fatalf("error message did not match got=%s", err.Error())
	}

	_, err = sr.FetchVar(test, "CONSTANT")
	if err == nil {
		t.Fatal("test failed to produce an error")
	}

	if err.Error() != "no constant found with name should in spec test" {
		t.Fatalf("error message did not match got=%s", err.Error())
	}

	_, err = sr.FetchVar(test, "INVALID")
	if err == nil {
		t.Fatal("test failed to produce an error")
	}

	if err.Error() != "cannot fetch a variable [test should fail] of type INVALID" {
		t.Fatalf("error message did not match got=%s", err.Error())
	}
}
