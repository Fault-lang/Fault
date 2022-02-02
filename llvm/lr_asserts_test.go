package llvm

// Named xasserts to make go test to run these tests AFTER the main
// tests in llvm_test

import (
	"fault/ast"
	"fault/listener"
	"fault/parser"
	"fault/types"
	"testing"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func TestSimpleAssert(t *testing.T) {
	test := `spec test1;
			const hello = false;
			assert hello == true;
	`

	llvm, err := prepAssertTest(test)

	if err != nil {
		t.Fatalf("compilation failed on valid spec. got=%s", err)
	}

	for _, v := range llvm.AssertAssume {

		c := v.(*ast.AssertionStatement).Constraints
		av := c.Variable.(*ast.AssertVar)
		if av.Instances[0] != "test1_hello" {
			t.Fatalf("assert assigned to wrong variable. got=%s", av.Instances[0])
		}

		if c.Comparison != "!=" {
			t.Fatalf("assert has wrong comparison. got=%s", c.Comparison)
		}

		if _, ok := c.Expression.(*ast.Boolean); !ok {
			t.Fatalf("assert has wrong operator. got=%s", c.Expression)
		}

	}
}

func TestAssertWConjunc(t *testing.T) {
	test := `spec test1;
			const hello = false;
			assert hello == true && 5 > 2;
	`

	llvm, err := prepAssertTest(test)

	if err != nil {
		t.Fatalf("compilation failed on valid spec. got=%s", err)
	}

	for _, v := range llvm.AssertAssume {

		c := v.(*ast.AssertionStatement).Constraints
		av := c.Variable.(*ast.InfixExpression).Left.(*ast.AssertVar)
		if av.Instances[0] != "test1_hello" {
			t.Fatalf("assert assigned to wrong variable. got=%s", av.Instances[0])
		}

		if c.Conjuction != "||" {
			t.Fatalf("assert has wrong comparison. got=%s", c.Conjuction)
		}

		if c.Variable.(*ast.InfixExpression).Operator != "!=" {
			t.Fatalf("assert has wrong operator. got=%s", c.Variable.(*ast.InfixExpression).Operator)
		}

		right := c.Expression.(*ast.InfixExpression)
		if right.Operator != "<=" {
			t.Fatalf("assert has wrong operator. got=%s", right.Operator)
		}

	}
}

// THIS TEST IS ~AWFUL~ BUT SOME OF THIS SYNTAX HASN'T ACTUALLY BEEN
// IMPLEMENTED. SO WILL INVESTIGATE THE FIX LATER
func TestAssertState(t *testing.T) {
	test := `spec test1;
			def fl = flow{
				value: 30,
				scope: 10,
				rate: func{
					value + 2;
				},
			};

			assert fl.value > fl.scope;
			assert fl.scope > -fl.value;
			assert fl.value[1] > fl.scope;

			for 2 run{
				x = new fl;
				x.rate; 
			}
	`

	llvm, err := prepAssertTest(test)

	if err != nil {
		t.Fatalf("compilation failed on valid spec. got=%s", err)
	}

	for i, v := range llvm.AssertAssume {
		c := v.(*ast.AssertionStatement).Constraints
		compareAsserts(t, c.Variable, i)

		if c.Comparison != "<=" {
			t.Fatalf("assert %d has wrong comparison. got=%s", i, c.Comparison)
		}
		compareAsserts(t, c.Expression, i)
	}
}

// HELPER FUNCTION FOR AWFUL TEST
func compareAsserts(t *testing.T, e ast.Expression, i int) {
	switch ex := e.(type) {
	case *ast.AssertVar:
		var check int
		for _, v := range ex.Instances {
			switch v {
			case "test1_x_scope":
				check++
			case "test1_x_value":
				check++
			case "test1_x_rate":
				check++
			default:
				t.Fatalf("assert %d has wrong expression. got=%s", i, v)
			}
		}

		if check != 1 {
			t.Fatalf("assert %d has wrong number of instances. got=%d want=3", i, check)
		}
	case *ast.IndexExpression:
		compareAsserts(t, ex.Left, i)
		if ex.Index.(*ast.IntegerLiteral).Value != 1 {
			t.Fatalf("assert %d has wrong expression. got=%d", i, ex.Index.(*ast.IntegerLiteral).Value)
		}
	}
}

// END AWFUL TEST, NOTHING TO SEE HERE... MOVE ALONG ;)

func prepAssertTest(test string) (*Compiler, error) {
	is := antlr.NewInputStream(test)
	lexer := parser.NewFaultLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	p := parser.NewFaultParser(stream)
	l := listener.NewListener(true, false)
	antlr.ParseTreeWalkerDefault.Walk(l, p.Spec())
	ty := &types.Checker{}
	err := ty.Check(l.AST)
	if err != nil {
		return nil, err
	}
	compiler := NewCompiler()
	compiler.LoadMeta(ty.SpecStructs, l.Uncertains, l.Unknowns)
	err = compiler.Compile(l.AST)
	if err != nil {
		return nil, err
	}
	return compiler, err
}
