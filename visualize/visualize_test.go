package visualize

import (
	"fault/ast"
	"fault/listener"
	"fault/parser"
	"fault/preprocess"
	"fault/types"
	"strings"
	"testing"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

func TestSpec(t *testing.T) {
	test := `spec test1;
		def foo = stock{
				foosh: 3,
			};

			def zoo = flow{
				con: new foo,
				rate: func{
					con.foosh + 2;
				},
			};

		for 2 run {
			bar = new zoo;
		};
		
	`

	vis := prepTest(test)

	got := vis.Render()

	if string(got[0]) == "\n" {
		t.Fatal("rough line break detected")
	}

	expected := `flowchart TD
	test1_bar{{test1_bar}}-->test1_bar_con[test1_bar_con]`

	if stripAndEscape(got) != stripAndEscape(expected) {
		t.Fatalf("incorrect visualization generated got=%s want=%s", got, expected)
	}

}

func TestSys(t *testing.T) {
	test := `system test1;
		component foo = states{
			idle: func{
				advance(this.step1);
			},
			step1: func{
				stay();
			},
		};
	`

	vis := prepSysTest(test, true)

	got := vis.Render()

	expected := `stateDiagram
	state foo {
		foo_idle --> foo_step1
	}`

	if stripAndEscape(got) != stripAndEscape(expected) {
		t.Fatalf("incorrect visualization generated got=%s want=%s", got, expected)
	}

}

func TestCombined(t *testing.T) {
	test := `system test1;
		import "../smt/testdata/simple.fspec"
		
		global f = new simple.fl; 

		component foo = states{
			idle: func{
				advance(this.step1);
			},
			step1: func{
				stay();
			},
		};
	`

	vis := prepSysTest(test, false)

	got := vis.Render()

	expected := `stateDiagram
	state foo {
		foo_idle --> foo_step1
	}
	
	flowchart TD
		test1_f{{test1_f}}-->test1_f_vault[test1_f_vault] 
`

	if stripAndEscape(got) != stripAndEscape(expected) {
		t.Fatalf("incorrect visualization generated got=%s want=%s", got, expected)
	}

}

func TestError(t *testing.T) {
	token := ast.Token{Type: "BAD", Position: []int{0, 0, 0, 0}}
	test := &ast.StructInstance{Token: token}

	vis := NewVisual(test)
	err := vis.walk(vis.tree)
	if err == nil {
		t.Fatal("visualizer did not error on bad tree")
	}
}

func stripAndEscape(str string) string {
	var output strings.Builder
	output.Grow(len(str))
	for _, ch := range str {
		if !unicode.IsSpace(ch) {
			if ch == '%' {
				output.WriteString("%%")
			} else {
				output.WriteRune(ch)
			}
		}
	}
	return output.String()
}

func prepTest(test string) *Visual {
	path := ""
	is := antlr.NewInputStream(test)
	lexer := parser.NewFaultLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	p := parser.NewFaultParser(stream)
	l := listener.NewListener(path, true, false)
	antlr.ParseTreeWalkerDefault.Walk(l, p.Spec())
	pre := preprocess.NewProcesser()
	pre.StructsPropertyOrder = l.StructsPropertyOrder
	tree := pre.Run(l.AST)

	ty := types.NewTypeChecker(pre)
	tree, _ = ty.Check(tree)

	vis := NewVisual(tree)
	vis.Build()

	return vis
}

func prepSysTest(test string, im bool) *Visual {
	path := ""
	is := antlr.NewInputStream(test)
	lexer := parser.NewFaultLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	p := parser.NewFaultParser(stream)
	l := listener.NewListener(path, im, false)
	antlr.ParseTreeWalkerDefault.Walk(l, p.SysSpec())
	pre := preprocess.NewProcesser()
	pre.StructsPropertyOrder = l.StructsPropertyOrder
	tree := pre.Run(l.AST)
	ty := types.NewTypeChecker(pre)
	tree, _ = ty.Check(tree)

	vis := NewVisual(tree)
	vis.Build()

	return vis
}
