package visualize

// import (
// 	"fault/ast"
// 	"fault/listener"
// 	"fault/preprocess"
// 	"fault/types"
// 	"strings"
// 	"testing"
// 	"unicode"
// )

// func TestSpec(t *testing.T) {
// 	test := `spec test1;
// 		def foo = stock{
// 				foosh: 3,
// 			};

// 			def zoo = flow{
// 				con: new foo,
// 				rate: func{
// 					con.foosh + 2;
// 				},
// 			};

// 		for 2 run {
// 			bar = new zoo;
// 		};

// 	`
// 	flags := make(map[string]bool)
// 	flags["specType"] = true
// 	flags["testing"] = false
// 	flags["skipRun"] = false
// 	vis := prepTest(test, flags)

// 	got := vis.Render()

// 	if string(got[0]) == "\n" {
// 		t.Fatal("rough line break detected")
// 	}

// 	expected := `flowchart TD
// 	test1_bar{{test1_bar}}-->test1_bar_con[test1_bar_con]`

// 	if stripAndEscape(got) != stripAndEscape(expected) {
// 		t.Fatalf("incorrect visualization generated got=%s want=%s", got, expected)
// 	}

// }

// func TestSys(t *testing.T) {
// 	test := `system test1;
// 		component foo = states{
// 			idle: func{
// 				advance(this.step1);
// 			},
// 			step1: func{
// 				stay();
// 			},
// 		};
// 	`

// 	flags := make(map[string]bool)
// 	flags["specType"] = false
// 	flags["testing"] = true
// 	flags["skipRun"] = false
// 	vis := prepTest(test, flags)

// 	got := vis.Render()

// 	expected := `stateDiagram
// 	state foo {
// 		foo_idle --> foo_step1
// 	}`

// 	if stripAndEscape(got) != stripAndEscape(expected) {
// 		t.Fatalf("incorrect visualization generated got=%s want=%s", got, expected)
// 	}

// }

// func TestCombined(t *testing.T) {
// 	test := `system test1;
// 		import "../smt/testdata/simple.fspec";

// 		global f = new simple.fl;

// 		component foo = states{
// 			idle: func{
// 				advance(this.step1);
// 			},
// 			step1: func{
// 				stay();
// 			},
// 		};
// 	`

// 	flags := make(map[string]bool)
// 	flags["specType"] = false
// 	flags["testing"] = false
// 	flags["skipRun"] = false
// 	vis := prepTest(test, flags)

// 	got := vis.Render()

// 	expected := `stateDiagram
// 	state foo {
// 		foo_idle --> foo_step1
// 	}

// 	flowchart TD
// 		test1_f{{test1_f}}-->test1_f_vault[test1_f_vault]
// `

// 	if stripAndEscape(got) != stripAndEscape(expected) {
// 		t.Fatalf("incorrect visualization generated got=%s \nwant=%s", got, expected)
// 	}

// }

// func TestError(t *testing.T) {
// 	token := ast.Token{Type: "BAD", Position: []int{0, 0, 0, 0}}
// 	test := &ast.StructInstance{Token: token}

// 	vis := NewVisual(test)
// 	err := vis.walk(vis.tree)
// 	if err == nil {
// 		t.Fatal("visualizer did not error on bad tree")
// 	}
// }

// func stripAndEscape(str string) string {
// 	var output strings.Builder
// 	output.Grow(len(str))
// 	for _, ch := range str {
// 		if !unicode.IsSpace(ch) {
// 			if ch == '%' {
// 				output.WriteString("%%")
// 			} else {
// 				output.WriteRune(ch)
// 			}
// 		}
// 	}
// 	return output.String()
// }

// func prepTest(test string, flags map[string]bool) *Visual {
// 	var path string
// 	l := listener.Execute(test, path, flags)
// 	pre := preprocess.Execute(l)
// 	ty := types.Execute(pre.Processed, pre)
// 	vis := NewVisual(ty.Checked)
// 	vis.Build()

// 	return vis
// }
