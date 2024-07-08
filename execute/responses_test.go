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
