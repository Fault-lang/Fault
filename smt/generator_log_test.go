package smt

// import (
// 	"fault/listener"
// 	"fault/llvm"
// 	"fault/preprocess"
// 	"fault/swaps"
// 	"fault/types"
// 	"fault/util"
// 	gopath "path"
// 	"testing"
// )

// func TestChange(t *testing.T) {
// 	test := `spec test1;
	
// 	def amount = stock{
// 		value: 10,
// 	};
	
// 	def test = flow{
// 		foo: new amount,
// 		bar: func{
// 			foo.value -> 2;
// 		},
// 	};

// 	for 1 init{t = new test;} run {
// 		t.bar;
// 	};
// 	`
// 	expecting := `(set-logic QF_NRA)
// 	(declare-fun test1_t_foo_value_0 () Real)
// 	(declare-fun test1_t_foo_value_1 () Real)
// 	(assert (= test1_t_foo_value_0 10.0))
// 	(assert (= test1_t_foo_value_1 (- test1_t_foo_value_0 2.0)))
// `

// 	generator := prepLogTest("", test, true, false)

// 	err := compareResults("LogChange", generator.SMT(), expecting)

// 	if err != nil {
// 		t.Fatalf(err.Error())
// 	}

// 	if len(generator.Log.Events) != 3 {
// 		t.Fatalf("ResultLog has wrong number of events got=%d, want=2", len(generator.Log.Events))
// 	}

// 	if generator.Log.Events[1].Type != "TRIGGER" && generator.Log.Events[1].Variable != "test1_t_bar" {
// 		t.Fatalf("Event has wrong data got=%s, want=0,TRIGGER,@__run,test1_t_bar,,,", generator.Log.Events[1].String())
// 	}

// 	if generator.Log.Events[1].String() != "0,TRIGGER,@__run,test1_t_bar,,,\n" {
// 		t.Fatalf("Event has wrong data got=%s, want=0,TRIGGER,@__run,test1_t_bar,,,", generator.Log.Events[1].String())
// 	}
// }

// func TestTransition(t *testing.T) {
// 	test := `system test1;
	
// 	component drain = states{
// 		initial: func{
// 			advance(this.open);
// 		},
// 		open: func{
// 			stay();
// 		},
// 	};
	
// 	start { 
// 		drain: initial,
// 	};
	
// 	for 1 run {};
	
// 	`
// 	expecting := `(set-logicQF_NRA)
// 	(declare-fun test1_drain_open_2 () Bool)
// 	(declare-fun test1_drain_initial_3 () Bool)
// 	(declare-fun test1_drain_initial_0 () Bool)
// 	(declare-fun test1_drain_open_0 () Bool)
// 	(declare-fun test1_drain_initial_1 () Bool)
// 	(declare-fun test1_drain_open_1 () Bool)
// 	(declare-fun test1_drain_initial_2 () Bool)
// 	(assert (= test1_drain_initial_0 false))
// 	(assert (= test1_drain_open_0 false))
// 	(assert (= test1_drain_initial_1 true))
// 	(assert (= test1_drain_open_1 true))
// 	(assert (= test1_drain_initial_2 false))
// 	(assert (ite (= test1_drain_initial_1 true) (and (= test1_drain_open_2 test1_drain_open_1)(= test1_drain_initial_3 test1_drain_initial_2))(and (= test1_drain_open_2 test1_drain_open_0)(= test1_drain_initial_3 test1_drain_initial_1))))
// `

// 	generator := prepLogTest("", test, false, false)

// 	err := compareResults("LogTransition", generator.SMT(), expecting)

// 	if err != nil {
// 		t.Fatalf(err.Error())
// 	}

// 	if len(generator.Log.Events) != 10 {
// 		t.Fatalf("ResultLog has wrong number of events got=%d, want=1-", len(generator.Log.Events))
// 	}

// 	if generator.Log.Events[6].String() != "1,TRANSITION,,__state,test1_drain_initial,test1_drain_open,\n" {
// 		t.Fatalf("Event has wrong data got=%s, want=1,TRANSITION,,__state,test1_drain_initial,test1_drain_open,", generator.Log.Events[6].String())
// 	}

// 	if generator.Log.Events[3].String() != "1,TRIGGER,@__run,test1_drain_initial__state,,,\n" {
// 		t.Fatalf("Event has wrong data got=%s, want=1,TRIGGER,@__run,test1_drain_initial__state,,,", generator.Log.Events[3].String())
// 	}
// }

// func prepLogTest(filepath string, test string, specType bool, testRun bool) *Generator {
// 	flags := make(map[string]bool)
// 	flags["specType"] = specType
// 	flags["testing"] = testRun
// 	flags["skipRun"] = false

// 	var path string
// 	if filepath != "" {
// 		filepath = util.Filepath(filepath)
// 		path = gopath.Dir(filepath)
// 	}

// 	l := listener.Execute(test, path, flags)
// 	pre := preprocess.Execute(l)
// 	ty := types.Execute(pre.Processed, pre)
// 	sw := swaps.NewPrecompiler(ty)
// 	tree := sw.Swap(ty.Checked)
// 	compiler := llvm.Execute(tree, ty.SpecStructs, l.Uncertains, l.Unknowns, sw.Alias, true)

// 	//fmt.Println(compiler.GetIR())
// 	generator := Execute(compiler)
// 	return generator
// }
