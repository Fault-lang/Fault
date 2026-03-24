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
