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

func prepAssertTest(test string) (*Compiler, error) {
	is := antlr.NewInputStream(test)
	lexer := parser.NewFaultLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	p := parser.NewFaultParser(stream)
	l := &listener.FaultListener{}
	antlr.ParseTreeWalkerDefault.Walk(l, p.Spec())
	ty := &types.Checker{}
	err := ty.Check(l.AST)
	if err != nil {
		return nil, err
	}
	compiler := NewCompiler(ty.SpecStructs)
	err = compiler.Compile(l.AST)
	if err != nil {
		return nil, err
	}
	//fmt.Println(compiler.GetIR())
	return compiler, err
}
