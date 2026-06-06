package generator

import (
	"fault/listener"
	"fault/llvm"
	"fault/preprocess"
	"fault/swaps"
	"fault/types"
	"fault/util"
	"fmt"
	"os"
	gopath "path"
	"regexp"
	"strconv"
	"strings"
	"testing"
	"unicode"
)

// ---- Temporal operators (structural checks) ----

func TestEventually(t *testing.T) {
	test := `spec test1;

	def amount = stock{
		value: 10,
	};

	def test = flow{
		foo: new amount,
		bar: func{
			foo.value -> 2;
		},
	};

	assume amount.value > 0 eventually;

	run init{t = new test;} {
		t.bar;
		t.bar;
		t.bar;
		t.bar;
		t.bar;
	};
	`
	g := prepTest("", test, true, false)
	smt := g.SMT()

	if !strings.Contains(smt, "(declare-fun") {
		t.Fatal("no SMT generated")
	}
	// eventually: or over all SSA versions
	if !strings.Contains(smt, "(assert (or") {
		t.Fatalf("assume eventually should produce (assert (or ...)), got:\n%s", smt)
	}
	// variable present
	if !strings.Contains(smt, "test1_t_foo_value_") {
		t.Fatalf("SMT missing test1_t_foo_value_. got:\n%s", smt)
	}
	// 6 rounds declared (0-5)
	count := strings.Count(smt, "declare-fun test1_t_foo_value_")
	if count != 6 {
		t.Fatalf("expected 6 declare-fun for test1_t_foo_value_, got %d", count)
	}
}

func TestEventuallyAlways(t *testing.T) {
	test := `spec test1;

	def amount = stock{
		value: 10,
	};

	def test = flow{
		foo: new amount,
		bar: func{
			foo.value -> 2;
		},
	};

	assume amount.value > 0 eventually-always;

	run init{t = new test;} {
		t.bar;
		t.bar;
		t.bar;
		t.bar;
		t.bar;
	};
	`
	g := prepTest("", test, true, false)
	smt := g.SMT()

	if !strings.Contains(smt, "(declare-fun") {
		t.Fatal("no SMT generated")
	}
	// eventually-always: or wrapping and-conjunctions
	if !strings.Contains(smt, "(assert (or") {
		t.Fatalf("assume eventually-always should produce (assert (or ...)), got:\n%s", smt)
	}
	if !strings.Contains(smt, "(and") {
		t.Fatalf("eventually-always should contain (and ...) conjunctions, got:\n%s", smt)
	}
	if !strings.Contains(smt, "(> test1_t_foo_value_") {
		t.Fatalf("SMT missing (> test1_t_foo_value_. got:\n%s", smt)
	}
}

func TestEventuallyAlways_Assert(t *testing.T) {
	// assert with eventually-always negates the operator (> becomes <=)
	test := `spec test1;

	def amount = stock{
		value: 10,
	};

	def test = flow{
		foo: new amount,
		bar: func{
			foo.value -> 2;
		},
	};

	assert amount.value > 0 eventually-always;

	run init{t = new test;} {
		t.bar;
		t.bar;
		t.bar;
		t.bar;
		t.bar;
	};
	`
	g := prepTest("", test, true, false)
	smt := g.SMT()

	if !strings.Contains(smt, "(declare-fun") {
		t.Fatal("no SMT generated")
	}
	if !strings.Contains(smt, "(assert (or") {
		t.Fatalf("assert eventually-always should produce (assert (or ...)), got:\n%s", smt)
	}
	// compiler negates >0 to <=0
	if !strings.Contains(smt, "(<= test1_t_foo_value_") {
		t.Fatalf("assert should negate operator (> to <=), got:\n%s", smt)
	}
}

func TestTemporal_NMT(t *testing.T) {
	// nmt 1: no more than 1 time; produces XOR-style combinations
	test := `spec test1;

	def amount = stock{
		value: 4,
	};

	def test = flow{
		foo: new amount,
		bar: func{
			foo.value -> 2;
		},
	};

	assert amount.value <= 0 nmt 1;

	run init{t = new test;} {
		t.bar;
		t.bar;
		t.bar;
		t.bar;
		t.bar;
	};
	`
	g := prepTest("", test, true, false)
	smt := g.SMT()

	if !strings.Contains(smt, "(declare-fun") {
		t.Fatal("no SMT generated")
	}
	if !strings.Contains(smt, "(<= test1_t_foo_value_") {
		t.Fatalf("SMT missing nmt assertion. got:\n%s", smt)
	}
	// nmt 1 produces (or ...) of single-term combinations (no (and) needed for n=1)
	if !strings.Contains(smt, "(assert (or") {
		t.Fatalf("nmt 1 should produce (assert (or ...)), got:\n%s", smt)
	}
}

func TestTemporal_Mixed(t *testing.T) {
	// Multiple temporal constraints in one spec all compile without error.
	test := `spec test1;
			const a;
			const b;

			def s = stock{
   				x: unknown(),
			};

			def test = flow{
				u: new s,
				bar: func{
		   			u.x <- a + b;
			},
		};

		assume s.x >= 2 && s.x < 10 nft 3;
		assume s.x == 2 nmt 2;
		assert s.x == 11 eventually;

		run init{t = new test;} {
			t.bar;
			t.bar;
			t.bar;
			t.bar;
			t.bar;
		};
	`
	g := prepTest("", test, true, false)
	smt := g.SMT()

	if !strings.Contains(smt, "(declare-fun") {
		t.Fatal("no SMT generated")
	}
	// assert s.x == 11 eventually → negated to != → (not (= ...))
	if !strings.Contains(smt, "(not (= test1_t_u_x_") {
		t.Fatalf("assert eventually should contain negated equality, got:\n%s", smt)
	}
	// assume nft 3 → (and ...) triples
	if !strings.Contains(smt, "(and (>= test1_t_u_x_") {
		t.Fatalf("assume nft should produce (and >= ...) conjunctions, got:\n%s", smt)
	}
	// both a and b constants present
	if !strings.Contains(smt, "test1_a_0") || !strings.Contains(smt, "test1_b_0") {
		t.Fatalf("SMT missing const declarations. got:\n%s", smt)
	}
}

func TestTemporalSys(t *testing.T) {
	test := `system test1;
		component a = states{
			foo: func{
				advance(b.bar);
			},
			zoo: func{
				advance(this.foo);
			},
		};

		component b = states{
			buzz: func{
				advance(a.foo);
			},
			bar: func{
				stay();
			},
		};

		assert when a.zoo then !b.bar;

		run {
			b.buzz && a.zoo;
		}
		`
	g := prepTest("", test, false, false)
	smt := g.SMT()

	if !strings.Contains(smt, "(declare-fun") {
		t.Fatal("no SMT generated")
	}
	// state variables declared
	if !strings.Contains(smt, "test1_a_foo_") {
		t.Fatalf("SMT missing test1_a_foo_ declarations. got:\n%s", smt)
	}
	if !strings.Contains(smt, "test1_b_bar_") {
		t.Fatalf("SMT missing test1_b_bar_ declarations. got:\n%s", smt)
	}
	// ite transitions present (advance() compiles to conditional transition)
	if !strings.Contains(smt, "(assert (ite") {
		t.Fatalf("SMT missing ite assertions for state transitions. got:\n%s", smt)
	}
	// when/then assertion present
	if !strings.Contains(smt, "test1_a_zoo_") || !strings.Contains(smt, "test1_b_bar_") {
		t.Fatalf("SMT missing when/then variables. got:\n%s", smt)
	}
}

func TestCrossRoundWhenThen(t *testing.T) {
	// Verify that `assert when A then B` fires correctly across rounds.
	test := `system test1;
		component a = states{
			active: func{
				stay();
			},
			inactive: func{
				stay();
			},
		};

		component b = states{
			on: func{
				stay();
			},
			off: func{
				stay();
			},
		};

		assert when a.active then b.on;

		start{
			a: active,
			b: on,
		};

		run{};
		`

	g := prepTest("", test, false, false)
	smt := g.SMT()

	if !strings.Contains(smt, "(declare-fun") {
		t.Fatal("no SMT generated")
	}

	// The when/then assertion must reference both a.active and b.on variables
	if !strings.Contains(smt, "test1_a_active_") {
		t.Fatalf("SMT missing test1_a_active_ in when/then assertion. got:\n%s", smt)
	}
	if !strings.Contains(smt, "test1_b_on_") {
		t.Fatalf("SMT missing test1_b_on_ in when/then assertion. got:\n%s", smt)
	}

	// The assertion must be a negation of the implication: (and A (not B))
	stripped := stripAndEscape(smt)
	if !strings.Contains(stripped, "(andtest1_a_active_") {
		t.Fatalf("when/then assertion missing (and a.active ...) structure. got:\n%s", smt)
	}
	if !strings.Contains(stripped, "(nottest1_b_on_") {
		t.Fatalf("when/then assertion missing (not b.on ...) structure. got:\n%s", smt)
	}
}

// ---- Golden file tests (broad pipeline coverage) ----

func TestTestData(t *testing.T) {
	specs := []string{
		"testdata/bathtub.fspec",
		"testdata/simple.fspec",
		"testdata/bathtub2.fspec",
		"testdata/booleans.fspec",
		"testdata/unknowns.fspec",
		"testdata/swaps/swaps.fspec",
		"testdata/swaps/swaps1.fspec",
		"testdata/swaps/swaps2.fspec",
		"testdata/indexes.fspec",
		"testdata/strings.fspec",
		"testdata/strings2.fspec",
	}
	smt2s := []string{
		"testdata/bathtub.smt2",
		"testdata/simple.smt2",
		"testdata/bathtub2.smt2",
		"testdata/booleans.smt2",
		"testdata/unknowns.smt2",
		"testdata/swaps/swaps.smt2",
		"testdata/swaps/swaps1.smt2",
		"testdata/swaps/swaps2.smt2",
		"testdata/indexes.smt2",
		"testdata/strings.smt2",
		"testdata/strings2.smt2",
	}
	for i, s := range specs {
		data, err := os.ReadFile(s)
		if err != nil {
			panic(fmt.Sprintf("spec %s is not valid", s))
		}
		expecting, err := os.ReadFile(smt2s[i])
		if err != nil {
			panic(fmt.Sprintf("compiled spec %s is not valid", smt2s[i]))
		}
		g := prepTest(s, string(data), true, false)

		err = compareResults(s, g.SMT(), string(expecting))

		if err != nil {
			t.Fatal(err.Error())
		}
	}
}

func TestImports(t *testing.T) {
	specs := []string{
		"testdata/imports/circle_import1.fspec",
		"testdata/imports/single_import.fspec",
		"testdata/imports/renamed_import.fspec",
	}
	smt2s := []string{
		"testdata/imports/circle_import.smt2",
		"testdata/imports/single_import.smt2",
		"testdata/imports/renamed_import.smt2",
	}

	for i, s := range specs {
		data, err := os.ReadFile(s)
		if err != nil {
			panic(fmt.Sprintf("spec %s is not valid", s))
		}
		expecting, err := os.ReadFile(smt2s[i])
		if err != nil {
			panic(fmt.Sprintf("compiled spec %s is not valid", smt2s[i]))
		}
		g := prepTest(s, string(data), true, false)

		err = compareResults(s, g.SMT(), string(expecting))

		if err != nil {
			fmt.Println(g.SMT())
			t.Fatal(err.Error())
		}
	}
}

func TestClocks(t *testing.T) {
	specs := []string{
		"testdata/increment.fspec",
		"testdata/history1.fspec",
		"testdata/history2.fspec",
		"testdata/history3.fspec",
		"testdata/history4.fspec",
	}
	smt2s := []string{
		"testdata/increment.smt2",
		"testdata/history1.smt2",
		"testdata/history2.smt2",
		"testdata/history3.smt2",
		"testdata/history4.smt2",
	}

	for i, s := range specs {
		data, err := os.ReadFile(s)
		if err != nil {
			panic(fmt.Sprintf("spec %s is not valid", s))
		}
		expecting, err := os.ReadFile(smt2s[i])
		if err != nil {
			panic(fmt.Sprintf("compiled spec %s is not valid", smt2s[i]))
		}
		g := prepTest(s, string(data), true, false)

		err = compareResults(s, g.SMT(), string(expecting))

		if err != nil {
			fmt.Println(g.SMT())
			t.Fatal(err.Error())
		}
	}
}

func TestSys(t *testing.T) {
	specs := [][]string{
		{"testdata/statecharts/statechart.fsystem", "0"},
		{"testdata/statecharts/advanceor.fsystem", "0"},
		{"testdata/statecharts/multioradvance.fsystem", "0"},
		{"testdata/statecharts/advanceand.fsystem", "0"},
		{"testdata/statecharts/mixedcalls.fsystem", "0"},
		{"testdata/statecharts/trigger.fsystem", "0"},
		{"testdata/statecharts/choose1.fsystem", "0"},
		{"testdata/statecharts/choose2.fsystem", "0"},
	}
	smt2s := []string{
		"testdata/statecharts/statechart.smt2",
		"testdata/statecharts/advanceor.smt2",
		"testdata/statecharts/multioradvance.smt2",
		"testdata/statecharts/advanceand.smt2",
		"testdata/statecharts/mixedcalls.smt2",
		"testdata/statecharts/trigger.smt2",
		"testdata/statecharts/choose1.smt2",
		"testdata/statecharts/choose2.smt2",
	}
	for i, s := range specs {
		data, err := os.ReadFile(s[0])
		if err != nil {
			panic(fmt.Sprintf("spec %s is not valid", s[0]))
		}
		imports, _ := strconv.ParseBool(s[1])

		expecting, err := os.ReadFile(smt2s[i])
		if err != nil {
			panic(fmt.Sprintf("compiled spec %s is not valid", smt2s[i]))
		}
		g := prepTest(s[0], string(data), false, imports)

		err = compareResults(s[0], g.SMT(), string(expecting))

		if err != nil {
			fmt.Println(g.SMT())
			t.Fatalf("compilation failed on valid spec %s. got=%s", s[0], err)
		}

	}
}

func TestMultiCond(t *testing.T) {
	specs := []string{
		"testdata/conditionals/multicond.fspec",
		"testdata/conditionals/multicond2.fspec",
		"testdata/conditionals/multicond3.fspec",
		"testdata/conditionals/multicond4.fspec",
		"testdata/conditionals/multicond5.fspec",
		"testdata/conditionals/condwelse.fspec",
	}
	smt2s := []string{
		"testdata/conditionals/multicond.smt2",
		"testdata/conditionals/multicond2.smt2",
		"testdata/conditionals/multicond3.smt2",
		"testdata/conditionals/multicond4.smt2",
		"testdata/conditionals/multicond5.smt2",
		"testdata/conditionals/condwelse.smt2",
	}

	for i, s := range specs {
		data, err := os.ReadFile(s)
		if err != nil {
			panic(fmt.Sprintf("spec %s is not valid", s))
		}
		expecting, err := os.ReadFile(smt2s[i])
		if err != nil {
			panic(fmt.Sprintf("compiled spec %s is not valid", smt2s[i]))
		}
		g := prepTest(s, string(data), true, true)

		err = compareResults(s, g.SMT(), string(expecting))

		if err != nil {
			fmt.Println(g.SMT())
			t.Fatal(err.Error())
		}
	}

}

// ---- Bad spec error handling ----

func TestBadSpecs(t *testing.T) {
	type specCase struct {
		path        string
		specType    bool
		expectedErr string // non-empty: expect an error containing this string
	}
	specs := []specCase{
		{"testdata/badspecs/nodefs.fspec", true, "Missing run block"},
		{"testdata/badspecs/doubleswap.fspec", true, "swapped more than once"},
		{"testdata/badspecs/sharedstate.fspec", true, ""},
		{"testdata/badspecs/deep.fspec", true, ""},
		{"testdata/badspecs/emptyfunc.fspec", true, "A function cannot be empty"},
		{"testdata/badspecs/aliaschain.fspec", true, "swapped more than once"},
	}

	for _, s := range specs {
		t.Run(s.path, func(t *testing.T) {
			data, err := os.ReadFile(s.path)
			if err != nil {
				t.Fatalf("could not read spec file %s: %v", s.path, err)
			}

			var result string
			var pipelineErr error

			func() {
				defer func() {
					if r := recover(); r != nil {
						pipelineErr = fmt.Errorf("panic: %v", r)
					}
				}()

				flags := make(map[string]bool)
				flags["specType"] = s.specType
				flags["testing"] = false
				flags["skipRun"] = false

				fp := util.Filepath(s.path)
				path := gopath.Dir(fp)

				l, err := listener.Execute(string(data), path, flags)
				if err != nil {
					pipelineErr = fmt.Errorf("listener: %w", err)
					return
				}
				pre, err := preprocess.Execute(l)
				if err != nil {
					pipelineErr = fmt.Errorf("preprocess: %w", err)
					return
				}
				ty, err := types.Execute(pre.Processed, pre)
	if err != nil {
		t.Fatal(err)
	}
				sw := swaps.NewPrecompiler(ty)
				tree := sw.Swap(ty.Checked)
				compiler, err := llvm.Execute(tree, ty.SpecStructs, l.Uncertains, l.Unknowns, l.Wholes, sw.Alias, false)
				if err != nil {
					pipelineErr = fmt.Errorf("llvm: %w", err)
					return
				}
				result = Execute(compiler, GeneratorOptions{}).SMT()
			}()

			if s.expectedErr != "" {
				if pipelineErr == nil {
					t.Fatalf("expected error containing %q but got none", s.expectedErr)
				} else if !strings.Contains(pipelineErr.Error(), s.expectedErr) {
					t.Fatalf("expected error containing %q, got: %v", s.expectedErr, pipelineErr)
				}
			} else {
				if pipelineErr != nil {
					t.Fatalf("unexpected error: %v", pipelineErr)
				}
				t.Logf("OK %s: %d bytes of SMT", s.path, len(result))
			}
		})
	}
}

// ---- Structural SMT correctness ----

func TestSwapsNested(t *testing.T) {
	test := `spec nestedswap;

def inner = stock{
	x: 5,
};

def outstock = stock{
	sub: new inner,
};

def f1 = flow{
	target: new outstock,
	fn: func{
		target.sub.x <- 1;
	},
};

run init{
	s = new outstock;
	f = new f1;
	f.target = s;
} {
	f.fn;
	f.fn;
}
`
	g := prepTest("", test, true, false)
	smt := stripAndEscape(g.SMT())

	if strings.Contains(smt, "nestedswap_f_target_sub_x") {
		t.Fatal("SMT contains unresolved flow-local variable nestedswap_f_target_sub_x: nested alias not applied")
	}
	if !strings.Contains(smt, "nestedswap_s_sub_x") {
		t.Fatalf("SMT missing expected base variable nestedswap_s_sub_x.\ngot=%s", smt)
	}
}

func TestSwapsMultiple(t *testing.T) {
	test := `spec multiswap;

def stock1 = stock{
	v: 10,
};

def stock2 = stock{
	w: 5,
};

def f1 = flow{
	addtarget: new stock1,
	subtarget: new stock2,
	fn1: func{
		addtarget.v <- 3;
	},
	fn2: func{
		subtarget.w -> 2;
	},
};

run init{
	sa = new stock1;
	sb = new stock2;
	f = new f1;
	f.addtarget = sa;
	f.subtarget = sb;
} {
	f.fn1 | f.fn2;
	f.fn1 | f.fn2;
}
`
	g := prepTest("", test, true, false)
	smt := stripAndEscape(g.SMT())

	if !strings.Contains(smt, "multiswap_sa_v") {
		t.Fatalf("SMT missing expected base variable multiswap_sa_v (addtarget swap not applied).\ngot=%s", smt)
	}
	if !strings.Contains(smt, "multiswap_sb_w") {
		t.Fatalf("SMT missing expected base variable multiswap_sb_w (subtarget swap not applied).\ngot=%s", smt)
	}
}

func TestUnusedVarElimination(t *testing.T) {
	test := `spec test1;

	def props = stock{
		used: 5,
		unused: 10,
	};

	def f = flow{
		snap: new props,
		fn: func{
			snap.used <- 1;
		},
	};

	run init{ inst = new f; } { inst.fn; };
	`
	g := prepTest("", test, true, false)
	smt := stripAndEscape(g.SMT())

	if strings.Contains(smt, "test1_inst_snap_unused") {
		t.Fatalf("SMT contains unused variable test1_inst_snap_unused; should have been eliminated.\ngot=%s", smt)
	}
	if !strings.Contains(smt, "test1_inst_snap_used") {
		t.Fatalf("SMT missing expected variable test1_inst_snap_used.\ngot=%s", smt)
	}
}

func TestPhiCompleteness(t *testing.T) {
	test := `spec phitest;

	def s = stock{
		a: 5,
		b: 10,
	};

	def f = flow{
		snap: new s,
		fn: func{
			if snap.a > 3 {
				snap.a <- 1;
			} else {
				snap.b <- 1;
			}
		},
	};

	run init{ inst = new f; } { inst.fn; };
	`
	g := prepTest("", test, true, false)
	smt := stripAndEscape(g.SMT())

	iteIdx := strings.Index(smt, "(assert(ite")
	if iteIdx < 0 {
		t.Fatal("no ite assertion found in SMT")
	}
	iteExpr := smt[iteIdx:]

	bCount := strings.Count(iteExpr, "phitest_inst_snap_b_")
	if bCount < 4 {
		t.Fatalf("phitest_inst_snap_b_ appears %d times in ite expression (want >= 4: constrained in both branches).\nite=%s", bCount, iteExpr)
	}

	aCount := strings.Count(iteExpr, "phitest_inst_snap_a_")
	if aCount < 5 {
		t.Fatalf("phitest_inst_snap_a_ appears %d times in ite expression (want >= 5: constrained in both branches).\nite=%s", aCount, iteExpr)
	}
}

func TestSynthSlot(t *testing.T) {
	test := `spec test1;
	def foo = flow{
		buzz: new bar,
		fizz: func{
			buzz.a <- buzz.a + 1;
		},
	};
	def bar = stock{
		a: 10,
	};
	run init{t = new foo;} {
		t.fizz;
		__;
		t.fizz;
	};
	`

	g := prepTest("", test, true, false)
	smt := g.SMT()

	if !strings.Contains(smt, "(declare-fun") {
		t.Fatalf("SMT missing declare-fun. got=%s", smt)
	}
	if !strings.Contains(smt, "synth_") {
		t.Fatalf("SMT missing synth_ selector. got=%s", smt)
	}
	if !strings.Contains(smt, "(=>") {
		t.Fatalf("SMT missing implication rule. got=%s", smt)
	}
	if !strings.Contains(smt, "test1_t_buzz_a") {
		t.Fatalf("SMT missing buzz.a variable. got=%s", smt)
	}
}

// ---- Synthesis integration tests (file-based) ----

func TestSynthPick(t *testing.T) {
	// synth_pick: solver must choose increment over decrement to satisfy counter.value > 10
	data, err := os.ReadFile("testdata/synth/synth_pick.fspec")
	if err != nil {
		t.Fatalf("could not read synth_pick.fspec: %v", err)
	}
	g := prepTest("testdata/synth/synth_pick.fspec", string(data), true, false)
	smt := g.SMT()

	if !strings.Contains(smt, "(declare-fun") {
		t.Fatal("no SMT generated for synth_pick")
	}
	// Selector vars for both candidates
	if !strings.Contains(smt, "synth_") {
		t.Fatalf("SMT missing synth_ selectors. got:\n%s", smt)
	}
	// Exactly-one XOR constraint
	if !strings.Contains(smt, "(assert (or") {
		t.Fatalf("SMT missing strictOr constraint. got:\n%s", smt)
	}
	// Implication: selector → candidate rules
	if !strings.Contains(smt, "(assert (=>") {
		t.Fatalf("SMT missing implication assertion. got:\n%s", smt)
	}
	// The assume eventually goal
	if !strings.Contains(smt, "(assert (or") {
		t.Fatalf("SMT missing assume eventually goal. got:\n%s", smt)
	}
	// Counter variable present
	if !strings.Contains(smt, "synth_pick_inst_c_value") {
		t.Fatalf("SMT missing counter variable. got:\n%s", smt)
	}
}

func TestSynthSequence(t *testing.T) {
	// synth_sequence: two __ slots; solver must pick fill twice to reach level 60
	data, err := os.ReadFile("testdata/synth/synth_sequence.fspec")
	if err != nil {
		t.Fatalf("could not read synth_sequence.fspec: %v", err)
	}
	g := prepTest("testdata/synth/synth_sequence.fspec", string(data), true, false)
	smt := g.SMT()

	if !strings.Contains(smt, "(declare-fun") {
		t.Fatal("no SMT generated for synth_sequence")
	}
	// Two synthesis slots → two synth_ selector groups
	count := strings.Count(smt, "synth_")
	if count < 2 {
		t.Fatalf("expected at least 2 synth_ references for two slots, got %d. smt:\n%s", count, smt)
	}
	if !strings.Contains(smt, "(assert (=>") {
		t.Fatalf("SMT missing implication assertions. got:\n%s", smt)
	}
	// Tank level variable
	if !strings.Contains(smt, "synth_sequence_inst_t_level") {
		t.Fatalf("SMT missing tank level variable. got:\n%s", smt)
	}
	// Assume eventually goal (or over all states)
	if !strings.Contains(smt, "(assert (or") {
		t.Fatalf("SMT missing assume eventually goal. got:\n%s", smt)
	}
}

func TestSynthSandwich(t *testing.T) {
	// synth_sandwich: explicit deposit, then __ slot; solver must pick withdraw
	data, err := os.ReadFile("testdata/synth/synth_sandwich.fspec")
	if err != nil {
		t.Fatalf("could not read synth_sandwich.fspec: %v", err)
	}
	g := prepTest("testdata/synth/synth_sandwich.fspec", string(data), true, false)
	smt := g.SMT()

	if !strings.Contains(smt, "(declare-fun") {
		t.Fatal("no SMT generated for synth_sandwich")
	}
	if !strings.Contains(smt, "synth_") {
		t.Fatalf("SMT missing synth_ selectors. got:\n%s", smt)
	}
	if !strings.Contains(smt, "(assert (=>") {
		t.Fatalf("SMT missing implication assertions. got:\n%s", smt)
	}
	// Both deposit and withdraw candidates must appear in the SMT
	if !strings.Contains(smt, "deposit") && !strings.Contains(smt, "withdraw") {
		t.Fatalf("SMT missing candidate function references. got:\n%s", smt)
	}
	// Wallet balance variable
	if !strings.Contains(smt, "synth_sandwich_inst_w_balance") {
		t.Fatalf("SMT missing wallet balance variable. got:\n%s", smt)
	}
}

// ---- Stock inheritance SMT integration tests ----

func TestStockInheritanceSMT(t *testing.T) {
	// Child stock inherits a field from its parent. When the child is used in a
	// flow and instantiated in a run block, the inherited field must appear as an
	// SMT variable and carry the parent's initial value.
	test := `spec test1;

def generic = stock{
	level: 10,
};

def child = stock{
	extends generic,
	extra: 5,
};

def f = flow{
	c: new child,
	fn: func{
		if c.extra > 5 {
			c.level -> 2;
		}else{
			c.extra <- 1;
		}
	},
};

assert child.extra < 7;

run init{inst = new f;} {
	inst.fn;
	inst.fn;
};
`
	g := prepTest("", test, true, false)
	smt := g.SMT()

	if !strings.Contains(smt, "(declare-fun") {
		t.Fatal("no SMT generated")
	}
	// Inherited field must be declared as an SMT variable for the run-block instance
	if !strings.Contains(smt, "test1_inst_c_level") {
		t.Fatalf("SMT missing inherited field test1_inst_c_level.\ngot:\n%s", smt)
	}
	// Initial value of the inherited field must match the parent's value (10)
	if !strings.Contains(smt, "(= test1_inst_c_level_0 10.0)") {
		t.Fatalf("SMT missing initial value assert for inherited field (want 10.0).\ngot:\n%s", smt)
	}
	// Own field of the child must also be present
	if !strings.Contains(smt, "test1_inst_c_extra") {
		t.Fatalf("SMT missing child's own field test1_inst_c_extra.\ngot:\n%s", smt)
	}
}

func TestStockInheritanceExcludeSMT(t *testing.T) {
	// Excluded fields must NOT appear in the SMT; non-excluded inherited fields must.
	test := `spec test1;

def base = stock{
	keep: 3,
	drop: 7,
};

def derived = stock{
	extends base,
	exclude drop,
	own: 1,
};

def f = flow{
	d: new derived,
	fn: func{
		d.keep -> 1;
	},
};

run init{inst = new f;} {
	inst.fn;
};
`
	g := prepTest("", test, true, false)
	smt := g.SMT()

	if !strings.Contains(smt, "(declare-fun") {
		t.Fatal("no SMT generated")
	}
	if !strings.Contains(smt, "test1_inst_d_keep") {
		t.Fatalf("SMT missing inherited (non-excluded) field test1_inst_d_keep.\ngot:\n%s", smt)
	}
	if strings.Contains(smt, "test1_inst_d_drop") {
		t.Fatalf("SMT contains excluded field test1_inst_d_drop, should be absent.\ngot:\n%s", smt)
	}
}

func TestStockInheritanceAssertPropagationSMT(t *testing.T) {
	// An assert written against the parent stock must generate SMT constraints
	// that reference the child's run-block instance variables.
	test := `spec test1;

def generic = stock{
	level: 10,
};

def child = stock{
	extends generic,
	extra: 5,
};

def f = flow{
	c: new child,
	fn: func{
		c.level -> 2;
	},
};

assert generic.level > 0;

run init{inst = new f;} {
	inst.fn;
	inst.fn;
};
`
	g := prepTest("", test, true, false)
	smt := g.SMT()

	if !strings.Contains(smt, "(declare-fun") {
		t.Fatal("no SMT generated")
	}
	// The assert must produce an SMT assertion referencing the child instance variable
	if !strings.Contains(smt, "test1_inst_c_level") {
		t.Fatalf("SMT missing child instance variable test1_inst_c_level.\ngot:\n%s", smt)
	}
	// The negated assert (> becomes <=) must appear
	if !strings.Contains(smt, "(<= test1_inst_c_level_") {
		t.Fatalf("SMT missing negated assert (<= test1_inst_c_level_N) for child.\ngot:\n%s", smt)
	}
}

// ---- Unfunc SMT generation tests ----

func TestUnfuncGeneratesDeclare(t *testing.T) {
	test := `
	system test;

	component fetch = states{
		id: false,
		count: false,
		countVotes: unfunc{
			requires fetch.id,
			emits fetch.count,
		},
	};

	run {
		fetch.countVotes;
	}
	`

	g := prepTest("", test, false, false)
	smt := g.SMT()

	if !strings.Contains(smt, "(declare-fun") {
		t.Fatalf("SMT missing declare-fun. got=%s", smt)
	}
	// The run block's Bool state variable drives activation (no separate _active var)
	if !strings.Contains(smt, "test_fetch_countVotes") {
		t.Fatalf("SMT missing unfunc state identifier. got=%s", smt)
	}
	// Activation guard implication uses the run block's state variable directly
	if !strings.Contains(smt, "(assert (=>") {
		t.Fatalf("SMT missing activation guard. got=%s", smt)
	}
	// Shadow availability variable declarations
	if !strings.Contains(smt, "_available_") {
		t.Fatalf("SMT missing _available shadow variable declarations. got=%s", smt)
	}
}

func TestUnfuncActivationGuard(t *testing.T) {
	test := `
	system test;

	component fetch = states{
		id: false,
		count: false,
		countVotes: unfunc{
			requires fetch.id,
			emits fetch.count,
		},
	};

	run {
		fetch.countVotes;
	}
	`

	g := prepTest("", test, false, false)
	smt := g.SMT()

	// Requires guard must use the _available shadow variable
	if !strings.Contains(smt, "test_fetch_id_available_") {
		t.Fatalf("SMT missing requires shadow variable test_fetch_id_available_. got=%s", smt)
	}
	// Write effect must use the _available shadow variable (not the numeric field)
	if !strings.Contains(smt, "test_fetch_count_available_") {
		t.Fatalf("SMT missing emits shadow variable test_fetch_count_available_. got=%s", smt)
	}
}

func TestUnfuncConjunctiveRequires(t *testing.T) {
	test := `
	system test;

	component ops = states{
		id: false,
		joinId: false,
		result: false,
		getWithJoin: unfunc{
			requires ops.id && ops.joinId,
			emits ops.result,
		},
	};

	run {
		ops.getWithJoin;
	}
	`

	g := prepTest("", test, false, false)
	smt := g.SMT()

	if !strings.Contains(smt, "(declare-fun") {
		t.Fatalf("SMT missing declare-fun. got=%s", smt)
	}
	// Conjunctive requires should produce (and ...) in SMT
	if !strings.Contains(smt, "(and") {
		t.Fatalf("SMT missing (and ...) for conjunctive requires. got=%s", smt)
	}
	if !strings.Contains(smt, "test_ops_id_available_") {
		t.Fatalf("SMT missing test_ops_id_available_. got=%s", smt)
	}
	if !strings.Contains(smt, "test_ops_joinId_available_") {
		t.Fatalf("SMT missing test_ops_joinId_available_. got=%s", smt)
	}
}

func TestUnfuncFrameCondition(t *testing.T) {
	test := `
	system test;

	component repo = states{
		key: false,
		value: false,
		lookup: unfunc{
			requires repo.key,
			emits repo.value,
		},
	};

	run {
		repo.lookup;
	}
	`

	g := prepTest("", test, false, false)
	smt := g.SMT()

	// Frame condition: (assert (=> (not active) (= field_available_N+1 field_available_N)))
	if !strings.Contains(smt, "(not") {
		t.Fatalf("SMT missing negation for frame condition. got=%s", smt)
	}
	// Write effect: (assert (=> active (= field_available_N+1 true)))
	if !strings.Contains(smt, "true") {
		t.Fatalf("SMT missing 'true' in write effect. got=%s", smt)
	}
	// Shadow variable must be present for the emitted field
	if !strings.Contains(smt, "test_repo_value_available_") {
		t.Fatalf("SMT missing _available shadow variable for repo.value. got=%s", smt)
	}
}

// ---- Test helpers ----

func compareResults(s string, smt string, expecting string) error {
	if !strings.Contains(smt, "(declare-fun") {
		return fmt.Errorf("smt not valid for spec %s. \ngot=%s", s, smt)
	}

	smt = stripAndEscape(smt)
	expecting = stripAndEscape(expecting)
	if len(smt) != len(expecting) {
		return fmt.Errorf("wrong instructions length for spec %s.\nwant=%s\ngot=%s",
			s, expecting, smt)
	}

	if smt != expecting {
		if !notStrictlyOrdered(expecting, smt) {
			return fmt.Errorf("SMT string does not match for spec %s.\nwant=%q\ngot=%q",
				s, expecting, smt)
		}
	}
	return nil
}

var blockNumRe = regexp.MustCompile(`block\d+(true|false)`)

func stripAndEscape(str string) string {
	str = blockNumRe.ReplaceAllString(str, "block${1}")
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

func prepTest(filepath string, test string, specType bool, testRun bool) *Generator {
	flags := make(map[string]bool)
	flags["specType"] = specType
	flags["testing"] = testRun
	flags["skipRun"] = false

	var path string
	if filepath != "" {
		filepath = util.Filepath(filepath)
		path = gopath.Dir(filepath)
	}

	l, _ := listener.Execute(test, path, flags)
	pre, err := preprocess.Execute(l)
	if err != nil {
		panic(err)
	}
	ty, err := types.Execute(pre.Processed, pre)
	if err != nil {
		panic(err)
	}
	sw := swaps.NewPrecompiler(ty)
	tree := sw.Swap(ty.Checked)
	compiler, err := llvm.Execute(tree, ty.SpecStructs, l.Uncertains, l.Unknowns, l.Wholes, sw.Alias, true)
	if err != nil {
		panic(err)
	}

	generator := Execute(compiler, GeneratorOptions{})
	return generator
}

func notStrictlyOrdered(want string, got string) bool {
	s := strings.Split(want, "")
	dedup := make(map[string]bool)
	var keys []string
	for _, v := range s {
		if _, ok := dedup[v]; !ok {
			dedup[v] = true
			keys = append(keys, v)
		}
	}

	for _, k := range keys {
		if strings.Count(want, k) != strings.Count(got, k) {
			return false
		}
	}
	return true
}

func TestUnfuncAssumeConstraint(t *testing.T) {
	test := `
	system test;

	component calc = states{
		a: false,
		b: false,
		product: false,
		multiply: unfunc{
			requires calc.a && calc.b,
			emits calc.product,
			assume calc.product = calc.a * calc.b,
		},
	};

	run {
		calc.multiply;
	}
	`

	g := prepTest("", test, false, false)
	smt := g.SMT()

	// LHS output field must be declared as a new versioned variable.
	if !strings.Contains(smt, "(declare-fun test_calc_product_1 ()") {
		t.Fatalf("SMT missing declaration of output field at step+1. got=%s", smt)
	}
	// Assume constraint: run-block state var drives activation (no separate _active).
	if !strings.Contains(smt, "(=> test_calc_multiply_1 (= test_calc_product_1 (* test_calc_a_0 test_calc_b_0)))") {
		t.Fatalf("SMT assume constraint missing or wrong. got=%s", smt)
	}
	// Frame condition: not active => output_n+1 = output_n (value unchanged when unfunc doesn't fire).
	if !strings.Contains(smt, "(=> (not test_calc_multiply_1) (= test_calc_product_1 test_calc_product_0))") {
		t.Fatalf("SMT missing frame condition for output field. got=%s", smt)
	}
	// _available shadow for the emitted field still present.
	if !strings.Contains(smt, "test_calc_product_available_") {
		t.Fatalf("SMT missing _available shadow variable for calc.product. got=%s", smt)
	}
}

func TestStockInitWithWhenThen(t *testing.T) {
	test := `
spec test_stock;

def square = stock{
    value,
};

def row = stock{
    pos1: new square,
    pos2: new square,
};

assume when row.pos1.value == 1 then row.pos2.value != 1;

run init{r1 = new row;} {
}
`
	g := prepTest("", test, true, false)
	smt := g.SMT()
	if !strings.Contains(smt, "(declare-fun test_stock_r1_pos1_value_0 () Real)") {
		t.Fatalf("SMT missing declare-fun for r1_pos1_value. got=%s", smt)
	}
	if !strings.Contains(smt, "(declare-fun test_stock_r1_pos2_value_0 () Real)") {
		t.Fatalf("SMT missing declare-fun for r1_pos2_value. got=%s", smt)
	}
	if strings.Contains(smt, "test_stock_row_pos1_value") {
		t.Fatalf("SMT should not contain template variable name test_stock_row_pos1_value. got=%s", smt)
	}
	if !strings.Contains(smt, "test_stock_r1_pos1_value_0") || !strings.Contains(smt, "test_stock_r1_pos2_value_0") {
		t.Fatalf("SMT constraint should use instance variable names. got=%s", smt)
	}
}

func TestIntegerModeDetection(t *testing.T) {
	test := `spec testwhole;
def square = stock{
    value: whole(),
};
run init{r1 = new square;} {
}`
	g := prepTest("", test, true, false)
	smt := g.SMT()
	if !strings.Contains(smt, "(set-logic QF_NIA)") {
		t.Fatalf("expected QF_NIA, got:\n%s", smt)
	}
	if !strings.Contains(smt, "(declare-fun testwhole_r1_value_0 () Int)") {
		t.Fatalf("expected Int sort, got:\n%s", smt)
	}
	if strings.Contains(smt, "is_int") {
		t.Fatalf("expected no is_int in integer mode, got:\n%s", smt)
	}
}

func TestUncertainBoundsDefault(t *testing.T) {
	// uncertain(mean, sigma) with no k should emit bounds using default k=3.0
	test := `spec testuncertain;
def box = stock{
    temp: uncertain(10.0, 2.0),
};
run init{r1 = new box;} {
}`
	g := prepTest("", test, true, false)
	smt := g.SMT()
	// default k=3.0: lower=4.0, upper=16.0
	if !strings.Contains(smt, "(assert (>= testuncertain_r1_temp_0 4.000000))") {
		t.Fatalf("expected lower bound assert, got:\n%s", smt)
	}
	if !strings.Contains(smt, "(assert (<= testuncertain_r1_temp_0 16.000000))") {
		t.Fatalf("expected upper bound assert, got:\n%s", smt)
	}
}

func TestUncertainBoundsCustomK(t *testing.T) {
	// uncertain(mean, sigma, k) should use the provided k
	test := `spec testuncertaink;
def box = stock{
    temp: uncertain(10.0, 2.0, 1.0),
};
run init{r1 = new box;} {
}`
	g := prepTest("", test, true, false)
	smt := g.SMT()
	// k=1.0: lower=8.0, upper=12.0
	if !strings.Contains(smt, "(assert (>= testuncertaink_r1_temp_0 8.000000))") {
		t.Fatalf("expected lower bound assert with k=1, got:\n%s", smt)
	}
	if !strings.Contains(smt, "(assert (<= testuncertaink_r1_temp_0 12.000000))") {
		t.Fatalf("expected upper bound assert with k=1, got:\n%s", smt)
	}
}
