package preprocess

import (
	"fault/ast"
	"fault/listener"
	"fault/parser"
	"testing"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func TestConstants(t *testing.T) {
	test := `spec test1;
	const x = 2;
	const y = 2+3.1;
	const z = unknown(a);`

	process, err := prepTest(test)

	if err != nil {
		t.Fatalf("Type checking failed on valid expression. got=%s", err)
	}

	consts := process.Specs["test1"]

	x := consts.FetchConstant("x")

	if _, ok := x.(*ast.IntegerLiteral); !ok {
		t.Fatalf("Constant x does not have the right type. got=%T", x)
	}

	if x.(*ast.IntegerLiteral).Value != 2 {
		t.Fatalf("Constant x does not have the right value. got=%d", x)
	}

	y := consts.FetchConstant("y")
	if _, ok := y.(*ast.InfixExpression); !ok {
		t.Fatalf("Constant y not the right type. got=%T", y)
	}

	if _, ok1 := y.(*ast.InfixExpression).Right.(*ast.FloatLiteral); !ok1 {
		t.Fatalf("right y node does not have the right type. got=%T", y.(*ast.InfixExpression).Right)
	}

	if _, ok2 := y.(*ast.InfixExpression).Left.(*ast.IntegerLiteral); !ok2 {
		t.Fatalf("left y node does not have the right type. got=%T", y.(*ast.InfixExpression).Left)
	}

	z := consts.FetchConstant("z")
	if _, ok3 := z.(*ast.Unknown); !ok3 {
		t.Fatalf("Constant z does not have the right type. got=%T", z)
	}
}

func TestStructDef(t *testing.T) {
	test := `spec test1;
	const a = 2.3;
			
			def foo = stock{
				foosh: 3,
				bar: "hello!",
				fizz: a,
				fizz2: -a,
			};

			def zoo = flow{
				con: new foo,
				rate: func{
					con.foosh + 2;
				},
				rate2: func{
					2 - a;
					a - 2;
				},
			};
	`

	process, err := prepTest(test)

	if err != nil {
		t.Fatalf("Type checking failed on valid expression. got=%s", err)
	}

	variables := process.Specs["test1"]

	foo := variables.FetchStock("foo")
	if len(foo) != 4 {
		t.Fatalf("stock foo returns the wrong number of properties got=%d want=4", len(foo))
	}

	zoo := variables.FetchFlow("zoo")
	if len(zoo) != 3 {
		t.Fatalf("flow zoo returns the wrong number of properties got=%d want=4", len(zoo))
	}

}

func TestInstances(t *testing.T) {
	test := `spec test1;
	def str = stock{
		foo: 3,
	};

	def str2 = stock{
		bar: new str,
	};

	def str3 = stock{
		foosh: new str2,
	};

	def fl = flow{
		buzz: new str3,
		fizz: func{
			buzz.foosh.bar.foo <- 5;
		},
	};
	`

	process, err := prepTest(test)

	if err != nil {
		t.Fatalf("Type checking failed on valid expression. got=%s", err)
	}

	variables := process.Specs["test1"]

	bar := variables.FetchStock("str2_bar")
	if bar == nil {
		t.Fatal("stock names str2_bar not found")
	}

	if len(bar) != 1 {
		t.Fatalf("stock str2_bar returns the wrong number of properties got=%d want=4", len(bar))
	}

	zoo := variables.FetchStock("fl_buzz_foosh_bar")
	if zoo == nil {
		t.Fatal("stock named fl_buzz_foosh_bar not found")
	}

	if len(zoo) != 1 {
		t.Fatalf("flow fl_buzz_foosh_bar returns the wrong number of properties got=%d want=4", len(zoo))
	}

}

func prepTest(test string) (*Processor, error) {
	path := ""
	is := antlr.NewInputStream(test)
	lexer := parser.NewFaultLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	p := parser.NewFaultParser(stream)
	l := listener.NewListener(path, true, false)
	antlr.ParseTreeWalkerDefault.Walk(l, p.Spec())
	pro := NewProcesser()
	err := pro.Run(l.AST)
	return pro, err
}

func prepSysTest(test string) (*Processor, error) {
	path := ""
	is := antlr.NewInputStream(test)
	lexer := parser.NewFaultLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	p := parser.NewFaultParser(stream)
	l := listener.NewListener(path, true, false)
	antlr.ParseTreeWalkerDefault.Walk(l, p.SysSpec())
	pro := NewProcesser()
	err := pro.Run(l.AST)
	return pro, err
}
