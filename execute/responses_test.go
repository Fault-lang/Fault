package execute

import (
	"fault/execute/parser"
	"testing"

	"github.com/antlr4-go/antlr/v4"

)

func TestParseReals(t *testing.T) {
	test := `(model 
		(define-fun imports_fl3_vault_value_2 () Real
		  50.0)
	  )
	  `
	response := prepTestParser(test)

	if response["imports_fl3_vault_value"] == nil {
		t.Fatalf("SMT parser failed to parse reals in solution returned. got=%s", response)
	}
}

func TestParsePrecise(t *testing.T) {
	test := `(model 
		(define-fun imports_fl3_vault_value_2 () Real
		  (/ 3.0 20.0))
	  )
	  `
	response := prepTestParser(test)

	if response["imports_fl3_vault_value"] == nil {
		t.Fatalf("SMT parser failed to parse reals in solution returned. got=%s", response)
	}
}

func TestParseNegPrecise(t *testing.T) {
	test := `(model 
		(define-fun imports_fl3_vault_value_2 () Real
		  (-(/ 3.0 20.0)))
	  )
	  `
	response := prepTestParser(test)

	if response["imports_fl3_vault_value"] == nil {
		t.Fatalf("SMT parser failed to parse reals in solution returned. got=%s", response)
	}
}

func TestParseNeg(t *testing.T) {
	test := `(model 
		(define-fun imports_fl3_vault_value_2 () Real
		  (- 20.0))
	  )
	  `
	response := prepTestParser(test)

	if response["imports_fl3_vault_value"] == nil {
		t.Fatalf("SMT parser failed to parse reals in solution returned. got=%s", response)
	}
}

func TestParseInts(t *testing.T) {
	test := `(model 
		(define-fun imports_fl3_vault_value_2 () Int
		  50)
	  )
	  `
	response := prepTestParser(test)

	if response["imports_fl3_vault_value"] == nil {
		t.Fatalf("SMT parser failed to parse int in solution returned. got=%s", response)
	}
}

func TestParseBools(t *testing.T) {
	test := `(model 
		(define-fun imports_fl3_vault_value_2 () Bool
		  true)
	  )
	  `
	response := prepTestParser(test)

	if response["imports_fl3_vault_value"] == nil {
		t.Fatalf("SMT parser failed to parse bools in solution returned. got=%s", response)
	}
}

func prepTestParser(response string) map[string]Scenario {
	is := antlr.NewInputStream(response)
	lexer := parser.NewSMTLIBv2Lexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	p := parser.NewSMTLIBv2Parser(stream)
	l := NewSMTListener()
	antlr.ParseTreeWalkerDefault.Walk(l, p.Start_())
	return l.Results
}

func prepTestListener(response string) *SMTListener {
	is := antlr.NewInputStream(response)
	lexer := parser.NewSMTLIBv2Lexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewSMTLIBv2Parser(stream)
	l := NewSMTListener()
	antlr.ParseTreeWalkerDefault.Walk(l, p.Start_())
	return l
}

func TestMergeTermPartsOne(t *testing.T) {
	result, err := mergeTermParts([]string{"3.0"})
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
	if result != "3.0" {
		t.Fatalf("expected 3.0, got %s", result)
	}
}

func TestMergeTermPartsDivision(t *testing.T) {
	result, err := mergeTermParts([]string{"3.0", "4.0"})
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
	if result != "0.750000" {
		t.Fatalf("expected 0.750000, got %s", result)
	}
}

func TestMergeTermPartsTooMany(t *testing.T) {
	_, err := mergeTermParts([]string{"1.0", "2.0", "3.0"})
	if err == nil {
		t.Fatal("expected error for too many term parts, got nil")
	}
}

func TestMergeTermPartsBadSecond(t *testing.T) {
	_, err := mergeTermParts([]string{"1.0", "notanumber"})
	if err == nil {
		t.Fatal("expected error for non-numeric second part, got nil")
	}
}

func TestConvertTermBadReal(t *testing.T) {
	_, err := convertTerm("Real", "notanumber")
	if err == nil {
		t.Fatal("expected error for invalid Real value, got nil")
	}
}

func TestConvertTermBadBool(t *testing.T) {
	_, err := convertTerm("Bool", "notabool")
	if err == nil {
		t.Fatal("expected error for invalid Bool value, got nil")
	}
}

func TestConvertTermBadInt(t *testing.T) {
	_, err := convertTerm("Int", "notanint")
	if err == nil {
		t.Fatal("expected error for invalid Int value, got nil")
	}
}

func TestConvertTermUnknownSort(t *testing.T) {
	v, err := convertTerm("String", "hello")
	if err != nil {
		t.Fatalf("unexpected error for unknown sort: %s", err)
	}
	if v != "hello" {
		t.Fatalf("expected passthrough value 'hello', got %v", v)
	}
}

func TestListenerErrPropagatedToSolve(t *testing.T) {
	// A model with a Bool variable that has a non-boolean value should
	// record an error on the listener rather than panicking.
	bad := `(model
		(define-fun test_var_1 () Bool
		  notabool)
	)`
	l := prepTestListener(bad)
	if l.err == nil {
		t.Fatal("expected listener error for invalid Bool value, got nil")
	}
}

// --- Multiple-round accumulation ---

func TestMultipleRoundsAccumulateInOneTrace(t *testing.T) {
	// Three rounds of the same variable should accumulate into a single
	// FloatTrace rather than creating three separate traces.
	model := `(model
		(define-fun vault_value_0 () Real 10.0)
		(define-fun vault_value_1 () Real 20.0)
		(define-fun vault_value_2 () Real 30.0)
	)`
	l := prepTestListener(model)
	if l.err != nil {
		t.Fatalf("unexpected error: %s", l.err)
	}

	trace, ok := l.Results["vault_value"].(*FloatTrace)
	if !ok || trace == nil {
		t.Fatalf("expected *FloatTrace for vault_value, got %T", l.Results["vault_value"])
	}
	if len(trace.Get()) != 3 {
		t.Fatalf("expected 3 entries in trace, got %d", len(trace.Get()))
	}
	if v, _ := trace.Index(0); v != 10.0 {
		t.Errorf("round 0: expected 10.0, got %f", v)
	}
	if v, _ := trace.Index(1); v != 20.0 {
		t.Errorf("round 1: expected 20.0, got %f", v)
	}
	if v, _ := trace.Index(2); v != 30.0 {
		t.Errorf("round 2: expected 30.0, got %f", v)
	}
}

// --- Multiple variables in one model ---

func TestMultipleVariablesInModel(t *testing.T) {
	model := `(model
		(define-fun foo_0 () Real 1.0)
		(define-fun bar_0 () Bool true)
	)`
	l := prepTestListener(model)
	if l.err != nil {
		t.Fatalf("unexpected error: %s", l.err)
	}
	if l.Results["foo"] == nil {
		t.Error("expected foo in Results, got nil")
	}
	if l.Results["bar"] == nil {
		t.Error("expected bar in Results, got nil")
	}
}

// --- Values map is populated with raw strings ---

func TestValuesMapPopulated(t *testing.T) {
	model := `(model
		(define-fun vault_value_0 () Real 42.0)
	)`
	l := prepTestListener(model)
	if l.err != nil {
		t.Fatalf("unexpected error: %s", l.err)
	}
	if l.Values["vault_value_0"] != "42.0" {
		t.Fatalf("expected Values[vault_value_0] = 42.0, got %q", l.Values["vault_value_0"])
	}
}

// --- Malformed symbol: non-numeric suffix sets l.err ---

func TestMalformedSymbolNonNumericSuffix(t *testing.T) {
	// splitIdent("first_var_bad") returns key="bad" which fails ParseInt.
	bad := `(model
		(define-fun first_var_bad () Real 50.0)
	)`
	l := prepTestListener(bad)
	if l.err == nil {
		t.Fatal("expected error for non-numeric symbol suffix, got nil")
	}
}

func TestMalformedSymbolNoUnderscore(t *testing.T) {
	// splitIdent("foo") returns key="foo" which fails ParseInt.
	bad := `(model
		(define-fun foo () Real 5.0)
	)`
	l := prepTestListener(bad)
	if l.err == nil {
		t.Fatal("expected error for symbol with no underscore, got nil")
	}
}

// --- Error short-circuits processing of subsequent entries ---

func TestErrorShortCircuitsSubsequentEntries(t *testing.T) {
	// First define-fun has a non-numeric suffix → sets l.err.
	// Second define-fun is valid but must not appear in Values or Results.
	model := `(model
		(define-fun first_var_bad () Real 50.0)
		(define-fun second_var_1 () Real 99.0)
	)`
	l := prepTestListener(model)
	if l.err == nil {
		t.Fatal("expected error for non-numeric symbol suffix, got nil")
	}
	if _, ok := l.Values["second_var_1"]; ok {
		t.Fatal("second entry should not have been written to Values after an error")
	}
	if l.Results["second_var"] != nil {
		t.Fatal("second entry should not have been written to Results after an error")
	}
}

// --- mergeTermParts: negative-prefix fast path ---

func TestMergeTermPartsNegativePrefix(t *testing.T) {
	// When parts[0] is not a valid float (e.g. the sign character "-"),
	// mergeTermParts joins the parts directly to form a negative literal.
	result, err := mergeTermParts([]string{"-", "3.0"})
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
	if result != "-3.0" {
		t.Fatalf("expected -3.0, got %s", result)
	}
}
