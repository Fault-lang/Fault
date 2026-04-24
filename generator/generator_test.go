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

	for 5 init{t = new test;} run {
		t.bar;
	};
	`
	expecting := `(set-logic QF_NRA)
	(declare-fun test1_t_foo_value_0 () Real)
	(declare-fun test1_t_foo_value_1 () Real)
	(declare-fun test1_t_foo_value_2 () Real)
	(declare-fun test1_t_foo_value_3 () Real)
	(declare-fun test1_t_foo_value_4 () Real)
	(declare-fun test1_t_foo_value_5 () Real)
	(assert (= test1_t_foo_value_0 10.0))
	(assert (= test1_t_foo_value_1 (- test1_t_foo_value_0 2.0)))
	(assert (= test1_t_foo_value_2 (- test1_t_foo_value_1 2.0)))
	(assert (= test1_t_foo_value_3 (- test1_t_foo_value_2 2.0)))
	(assert (= test1_t_foo_value_4 (- test1_t_foo_value_3 2.0)))
	(assert (= test1_t_foo_value_5 (- test1_t_foo_value_4 2.0)))
	(assert (or (> test1_t_foo_value_0 0) (> test1_t_foo_value_1 0)(> test1_t_foo_value_2 0)(> test1_t_foo_value_3 0)(> test1_t_foo_value_4 0)(> test1_t_foo_value_5 0)))
`

	g := prepTest("", test, true, false)

	err := compareResults("Eventually", g.SMT(), expecting)

	if err != nil {
		t.Fatal(err.Error())
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

	for 5 init{t = new test;} run {
		t.bar;
	};
	`
	expecting := `(set-logic QF_NRA)
	(declare-fun test1_t_foo_value_0 () Real)
	(declare-fun test1_t_foo_value_1 () Real)
	(declare-fun test1_t_foo_value_2 () Real)
	(declare-fun test1_t_foo_value_3 () Real)
	(declare-fun test1_t_foo_value_4 () Real)
	(declare-fun test1_t_foo_value_5 () Real)
	(assert (= test1_t_foo_value_0 10.0))
	(assert (= test1_t_foo_value_1 (- test1_t_foo_value_0 2.0)))
	(assert (= test1_t_foo_value_2 (- test1_t_foo_value_1 2.0)))
	(assert (= test1_t_foo_value_3 (- test1_t_foo_value_2 2.0)))
	(assert (= test1_t_foo_value_4 (- test1_t_foo_value_3 2.0)))
	(assert (= test1_t_foo_value_5 (- test1_t_foo_value_4 2.0)))
	(assert (or
		(and (> test1_t_foo_value_0 0) (> test1_t_foo_value_1 0)(> test1_t_foo_value_2 0)(> test1_t_foo_value_3 0)(> test1_t_foo_value_4 0)(> test1_t_foo_value_5 0))
		(and (> test1_t_foo_value_1 0)(> test1_t_foo_value_2 0)(> test1_t_foo_value_3 0)(> test1_t_foo_value_4 0)(> test1_t_foo_value_5 0))
		(and (> test1_t_foo_value_2 0)(> test1_t_foo_value_3 0)(> test1_t_foo_value_4 0)(> test1_t_foo_value_5 0))
		(and (> test1_t_foo_value_3 0)(> test1_t_foo_value_4 0)(> test1_t_foo_value_5 0))
		(and (> test1_t_foo_value_4 0)(> test1_t_foo_value_5 0))
		(> test1_t_foo_value_5 0)
		))
`

	g := prepTest("", test, true, false)

	err := compareResults("EventuallyAlways", g.SMT(), expecting)

	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestEventuallyAlways2(t *testing.T) {
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

	for 5 init{t = new test;} run {
		t.bar;
	};
	`
	expecting := `(set-logic QF_NRA)
	(declare-fun test1_t_foo_value_0 () Real)
	(declare-fun test1_t_foo_value_1 () Real)
	(declare-fun test1_t_foo_value_2 () Real)
	(declare-fun test1_t_foo_value_3 () Real)
	(declare-fun test1_t_foo_value_4 () Real)
	(declare-fun test1_t_foo_value_5 () Real)
	(assert (= test1_t_foo_value_0 10.0))
	(assert (= test1_t_foo_value_1 (- test1_t_foo_value_0 2.0)))
	(assert (= test1_t_foo_value_2 (- test1_t_foo_value_1 2.0)))
	(assert (= test1_t_foo_value_3 (- test1_t_foo_value_2 2.0)))
	(assert (= test1_t_foo_value_4 (- test1_t_foo_value_3 2.0)))
	(assert (= test1_t_foo_value_5 (- test1_t_foo_value_4 2.0)))
	(assert (or
		(and (<= test1_t_foo_value_0 0) (<= test1_t_foo_value_1 0)(<= test1_t_foo_value_2 0)(<= test1_t_foo_value_3 0)(<= test1_t_foo_value_4 0)(<= test1_t_foo_value_5 0))
		(and (<= test1_t_foo_value_1 0)(<= test1_t_foo_value_2 0)(<= test1_t_foo_value_3 0)(<= test1_t_foo_value_4 0)(<= test1_t_foo_value_5 0))
		(and (<= test1_t_foo_value_2 0)(<= test1_t_foo_value_3 0)(<= test1_t_foo_value_4 0)(<= test1_t_foo_value_5 0))
		(and (<= test1_t_foo_value_3 0)(<= test1_t_foo_value_4 0)(<= test1_t_foo_value_5 0))
		(and (<= test1_t_foo_value_4 0)(<= test1_t_foo_value_5 0))
		(<= test1_t_foo_value_5 0)
		))
`

	g := prepTest("", test, true, false)

	err := compareResults("EventuallyAlways2", g.SMT(), expecting)

	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestTemporal(t *testing.T) {
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

	for 5 init{t = new test;} run {
		t.bar;
	};
	`
	expecting := `(set-logic QF_NRA)
	(declare-fun test1_t_foo_value_0 () Real)
	(declare-fun test1_t_foo_value_1 () Real)
	(declare-fun test1_t_foo_value_2 () Real)
	(declare-fun test1_t_foo_value_3 () Real)
	(declare-fun test1_t_foo_value_4 () Real)
	(declare-fun test1_t_foo_value_5 () Real)(assert (= test1_t_foo_value_0 4.0))
	(assert (= test1_t_foo_value_1 (- test1_t_foo_value_0 2.0)))
	(assert (= test1_t_foo_value_2 (- test1_t_foo_value_1 2.0)))
	(assert (= test1_t_foo_value_3 (- test1_t_foo_value_2 2.0)))
	(assert (= test1_t_foo_value_4 (- test1_t_foo_value_3 2.0)))
	(assert (= test1_t_foo_value_5 (- test1_t_foo_value_4 2.0)))
	(assert (or (and (<= test1_t_foo_value_0 0) (<= test1_t_foo_value_1 0)) (and (<= test1_t_foo_value_0 0) (<= test1_t_foo_value_2 0)) (and (<= test1_t_foo_value_0 0) (<= test1_t_foo_value_3 0)) (and (<= test1_t_foo_value_0 0) (<= test1_t_foo_value_4 0)) (and (<= test1_t_foo_value_0 0) (<= test1_t_foo_value_5 0)) (and (<= test1_t_foo_value_1 0) (<= test1_t_foo_value_2 0)) (and (<= test1_t_foo_value_1 0) (<= test1_t_foo_value_3 0)) (and (<= test1_t_foo_value_1 0) (<= test1_t_foo_value_4 0)) (and (<= test1_t_foo_value_1 0) (<= test1_t_foo_value_5 0)) (and (<= test1_t_foo_value_2 0) (<= test1_t_foo_value_3 0)) (and (<= test1_t_foo_value_2 0) (<= test1_t_foo_value_4 0)) (and (<= test1_t_foo_value_2 0) (<= test1_t_foo_value_5 0)) (and (<= test1_t_foo_value_3 0) (<= test1_t_foo_value_4 0)) (and (<= test1_t_foo_value_3 0) (<= test1_t_foo_value_5 0)) (and (<= test1_t_foo_value_4 0) (<= test1_t_foo_value_5 0))))
	`

	g := prepTest("", test, true, false)

	err := compareResults("Temporal", g.SMT(), expecting)

	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestTemporal2(t *testing.T) {

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

		for 5 init{t = new test;} run {
			t.bar;
		};
	`
	expecting := `(set-logic QF_NRA)
	(declare-fun test1_a_0 () Real)
(declare-fun test1_b_0 () Real)
(declare-fun test1_t_u_x_0 () Real)
(declare-fun test1_t_u_x_1 () Real)
(declare-fun test1_t_u_x_2 () Real)
(declare-fun test1_t_u_x_3 () Real)
(declare-fun test1_t_u_x_4 () Real)
(declare-fun test1_t_u_x_5 () Real)
(assert (= test1_t_u_x_1 (+ test1_t_u_x_0 (+ test1_a_0 test1_b_0))))
(assert (= test1_t_u_x_2 (+ test1_t_u_x_1 (+ test1_a_0 test1_b_0))))
(assert (= test1_t_u_x_3 (+ test1_t_u_x_2 (+ test1_a_0 test1_b_0))))
(assert (= test1_t_u_x_4 (+ test1_t_u_x_3 (+ test1_a_0 test1_b_0))))
(assert (= test1_t_u_x_5 (+ test1_t_u_x_4 (+ test1_a_0 test1_b_0))))(assert (and (not (= test1_t_u_x_0 11)) (not (= test1_t_u_x_1 11)) (not (= test1_t_u_x_2 11)) (not (= test1_t_u_x_3 11)) (not (= test1_t_u_x_4 11)) (not (= test1_t_u_x_5 11))))
(assert (or (and (>= test1_t_u_x_0 2) (< test1_t_u_x_0 10)) (and (>= test1_t_u_x_1 2) (< test1_t_u_x_1 10)) (and (>= test1_t_u_x_2 2) (< test1_t_u_x_2 10)) (and (>= test1_t_u_x_3 2) (< test1_t_u_x_3 10)) (and (>= test1_t_u_x_4 2) (< test1_t_u_x_4 10)) (and (>= test1_t_u_x_5 2) (< test1_t_u_x_5 10))))
(assert (or (and (or (= test1_t_u_x_0 2) (= test1_t_u_x_1 2)) (and (not (= test1_t_u_x_2 2)) (not (= test1_t_u_x_3 2)) (not (= test1_t_u_x_4 2)) (not (= test1_t_u_x_5 2)))) (and (or (= test1_t_u_x_0 2) (= test1_t_u_x_2 2)) (and (not (= test1_t_u_x_1 2)) (not (= test1_t_u_x_3 2)) (not (= test1_t_u_x_4 2)) (not (= test1_t_u_x_5 2)))) (and (or (= test1_t_u_x_0 2) (= test1_t_u_x_3 2)) (and (not (= test1_t_u_x_1 2)) (not (= test1_t_u_x_2 2)) (not (= test1_t_u_x_4 2)) (not (= test1_t_u_x_5 2)))) (and (or (= test1_t_u_x_0 2) (= test1_t_u_x_4 2)) (and (not (= test1_t_u_x_1 2)) (not (= test1_t_u_x_2 2)) (not (= test1_t_u_x_3 2)) (not (= test1_t_u_x_5 2)))) (and (or (= test1_t_u_x_0 2) (= test1_t_u_x_5 2)) (and (not (= test1_t_u_x_1 2)) (not (= test1_t_u_x_2 2)) (not (= test1_t_u_x_3 2)) (not (= test1_t_u_x_4 2)))) (and (or (= test1_t_u_x_1 2) (= test1_t_u_x_2 2)) (and (not (= test1_t_u_x_0 2)) (not (= test1_t_u_x_3 2)) (not (= test1_t_u_x_4 2)) (not (= test1_t_u_x_5 2)))) (and (or (= test1_t_u_x_1 2) (= test1_t_u_x_3 2)) (and (not (= test1_t_u_x_0 2)) (not (= test1_t_u_x_2 2)) (not (= test1_t_u_x_4 2)) (not (= test1_t_u_x_5 2)))) (and (or (= test1_t_u_x_1 2) (= test1_t_u_x_4 2)) (and (not (= test1_t_u_x_0 2)) (not (= test1_t_u_x_2 2)) (not (= test1_t_u_x_3 2)) (not (= test1_t_u_x_5 2)))) (and (or (= test1_t_u_x_1 2) (= test1_t_u_x_5 2)) (and (not (= test1_t_u_x_0 2)) (not (= test1_t_u_x_2 2)) (not (= test1_t_u_x_3 2)) (not (= test1_t_u_x_4 2)))) (and (or (= test1_t_u_x_2 2) (= test1_t_u_x_3 2)) (and (not (= test1_t_u_x_0 2)) (not (= test1_t_u_x_1 2)) (not (= test1_t_u_x_4 2)) (not (= test1_t_u_x_5 2)))) (and (or (= test1_t_u_x_2 2) (= test1_t_u_x_4 2)) (and (not (= test1_t_u_x_0 2)) (not (= test1_t_u_x_1 2)) (not (= test1_t_u_x_3 2)) (not (= test1_t_u_x_5 2)))) (and (or (= test1_t_u_x_2 2) (= test1_t_u_x_5 2)) (and (not (= test1_t_u_x_0 2)) (not (= test1_t_u_x_1 2)) (not (= test1_t_u_x_3 2)) (not (= test1_t_u_x_4 2)))) (and (or (= test1_t_u_x_3 2) (= test1_t_u_x_4 2)) (and (not (= test1_t_u_x_0 2)) (not (= test1_t_u_x_1 2)) (not (= test1_t_u_x_2 2)) (not (= test1_t_u_x_5 2)))) (and (or (= test1_t_u_x_3 2) (= test1_t_u_x_5 2)) (and (not (= test1_t_u_x_0 2)) (not (= test1_t_u_x_1 2)) (not (= test1_t_u_x_2 2)) (not (= test1_t_u_x_4 2)))) (and (or (= test1_t_u_x_4 2) (= test1_t_u_x_5 2)) (and (not (= test1_t_u_x_0 2)) (not (= test1_t_u_x_1 2)) (not (= test1_t_u_x_2 2)) (not (= test1_t_u_x_3 2))))))`

	g := prepTest("", test, true, false)

	err := compareResults("Temporal2", g.SMT(), string(expecting))

	if err != nil {
		t.Fatal(err.Error())
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

		start{
			b: buzz,
			a: zoo,
		};
		`
	// Execution order follows transition chains from start states (a:zoo, b:buzz):
	// a.zoo → a.foo → b.bar (stay, stop); then b.buzz → a.foo (already visited, stop)
	expecting := `(set-logicQF_NRA)(declare-funtest1_a_foo_0()Bool)(declare-funtest1_a_zoo_0()Bool)(declare-funtest1_b_buzz_0()Bool)(declare-funtest1_b_bar_0()Bool)(declare-funtest1_a_zoo_1()Bool)(declare-funtest1_b_buzz_1()Bool)(declare-funtest1_a_foo_1()Bool)(declare-funtest1_a_foo_2()Bool)(declare-funblocktrue_0()Bool)(declare-funblockfalse_0()Bool)(declare-funtest1_b_bar_1()Bool)(declare-funtest1_b_bar_2()Bool)(declare-funblocktrue_0()Bool)(declare-funblockfalse_0()Bool)(declare-funtest1_b_bar_3()Bool)(declare-funtest1_b_bar_4()Bool)(declare-funblocktrue_0()Bool)(declare-funblockfalse_0()Bool)(declare-funtest1_a_foo_3()Bool)(declare-funtest1_a_foo_4()Bool)(declare-funblocktrue_0()Bool)(declare-funblockfalse_0()Bool)(assert(=test1_a_foo_0false))(assert(=test1_a_zoo_0false))(assert(=test1_b_buzz_0false))(assert(=test1_b_bar_0false))(assert(=test1_a_zoo_1true))(assert(=test1_b_buzz_1true))(assert(=test1_a_foo_1true))(assert(ite(=test1_a_zoo_1true)(and(=blocktrue_0true)(=blockfalse_0false)(=test1_a_foo_2test1_a_foo_1))(and(=blocktrue_0false)(=blockfalse_0true)(=test1_a_foo_2test1_a_foo_0))))(assert(or(andblocktrue_0(notblockfalse_0))(and(notblocktrue_0)blockfalse_0)))(assert(=test1_b_bar_1true))(assert(ite(=test1_a_foo_2true)(and(=blocktrue_0true)(=blockfalse_0false)(=test1_b_bar_2test1_b_bar_1))(and(=blocktrue_0false)(=blockfalse_0true)(=test1_b_bar_2test1_b_bar_0))))(assert(or(andblocktrue_0(notblockfalse_0))(and(notblocktrue_0)blockfalse_0)))(assert(=test1_b_bar_3true))(assert(ite(=test1_b_bar_2true)(and(=blocktrue_0true)(=blockfalse_0false)(=test1_b_bar_4test1_b_bar_3))(and(=blocktrue_0false)(=blockfalse_0true)(=test1_b_bar_4test1_b_bar_2))))(assert(or(andblocktrue_0(notblockfalse_0))(and(notblocktrue_0)blockfalse_0)))(assert(=test1_a_foo_3true))(assert(ite(=test1_b_buzz_1true)(and(=blocktrue_0true)(=blockfalse_0false)(=test1_a_foo_4test1_a_foo_3))(and(=blocktrue_0false)(=blockfalse_0true)(=test1_a_foo_4test1_a_foo_2))))(assert(or(andblocktrue_0(notblockfalse_0))(and(notblocktrue_0)blockfalse_0)))(assert(or(andtest1_a_zoo_0(nottest1_b_bar_0))(andtest1_a_zoo_1(nottest1_b_bar_0))(andtest1_a_zoo_1(nottest1_b_bar_2))(andtest1_a_zoo_1(nottest1_b_bar_4))))`

	g := prepTest("", test, false, false)

	err := compareResults("TemporalSys", g.SMT(), string(expecting))
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestCrossRoundWhenThen(t *testing.T) {
	// Verify that `assert when A then B` fires correctly across rounds.
	//
	// The Fault model works as follows:
	//   - Variables carry their round-N phi value into round N+1 unchanged.
	//   - Combos are generated at every SSA Init (new-value assignment).
	//   - The cross-round combo (a_active_5, b_on_3) detects:
	//       "a.active is still true in round 2 AND b.on was false at end of round 1"
	//   - The combo (a_active_3, b_on_3) detects the same-round case:
	//       "both active=true, on=false at end of round 1"
	//
	// Together these cover: if A=true at end of round 1 and B=false at
	// the beginning of round 2 (= end of round 1 phi value), the assertion fires.
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

		for 2 run{};
		`

	g := prepTest("", test, false, false)
	smt := g.SMT()
	stripped := stripAndEscape(smt)

	if !strings.Contains(smt, "(declare-fun") {
		t.Fatal("no SMT generated")
	}

	// The assertion must contain the key cross-round combo:
	// (and test1_a_active_5 (not test1_b_on_3))
	// where _5 is round-2's a.active phi and _3 is round-1's b.on phi.
	// This combo fires when a is still active in round 2 but b was already
	// inactive at the end of round 1 (= beginning of round 2).
	if !strings.Contains(stripped, "(andtest1_a_active_5(nottest1_b_on_3))") {
		t.Fatalf("missing cross-round combo (a_active_5, b_on_3) in assertion.\ngot SMT:\n%s", smt)
	}

	// Also verify the same-round combo for completeness.
	if !strings.Contains(stripped, "(andtest1_a_active_3(nottest1_b_on_3))") {
		t.Fatalf("missing same-round combo (a_active_3, b_on_3) in assertion.\ngot SMT:\n%s", smt)
	}
}

//func TestTemporalSys2(t *testing.T) {
// 	test := `system test1;

// 		component a = states{
// 			foo: func{
// 				advance(b.bar);
// 			},
// 			zoo: func{
// 				advance(this.foo);
// 			},
// 		};

// 		component b = states{
// 			buzz: func{
// 				advance(a.foo);
// 			},
// 			bar: func{
// 				stay();
// 			},
// 		};

// 		assert when a.zoo && b.buzz then b.bar > 3 && a.foo == 4;

// 		start{
// 			b: buzz,
// 			a: zoo,
// 		};

// 		for 2 run{};
// 		`
// 	expecting := `(set-logic QF_NRA)
// 	(declare-fun test1_b_bar_2 () Bool)
// 	(declare-fun test1_a_foo_2 () Bool)
// 	(declare-fun test1_a_zoo_3 () Bool)
// 	(declare-fun test1_a_foo_4 () Bool)
// 	(declare-fun test1_a_foo_6 () Bool)
// 	(declare-fun test1_b_buzz_3 () Bool)
// 	(declare-fun test1_b_bar_4 () Bool)
// 	(declare-fun test1_a_foo_8 () Bool)
// 	(declare-fun test1_a_foo_10 () Bool)
// 	(declare-fun test1_a_zoo_5 () Bool)
// 	(declare-fun test1_a_foo_12 () Bool)
// 	(declare-fun test1_b_buzz_5 () Bool)
// 	(declare-fun test1_a_foo_0 () Bool)
// 	(declare-fun test1_a_zoo_0 () Bool)
// 	(declare-fun test1_b_buzz_0 () Bool)
// 	(declare-fun test1_b_bar_0 () Bool)
// 	(declare-fun test1_a_zoo_1 () Bool)
// 	(declare-fun test1_b_buzz_1 () Bool)
// 	(declare-fun test1_b_bar_1 () Bool)
// 	(declare-fun test1_a_foo_1 () Bool)
// 	(declare-fun test1_a_foo_3 () Bool)
// 	(declare-fun test1_a_zoo_2 () Bool)
// 	(declare-fun test1_a_foo_5 () Bool)
// 	(declare-fun test1_b_buzz_2 () Bool)
// 	(declare-fun test1_b_bar_3 () Bool)
// 	(declare-fun test1_a_foo_7 () Bool)
// 	(declare-fun test1_a_foo_9 () Bool)
// 	(declare-fun test1_a_zoo_4 () Bool)
// 	(declare-fun test1_a_foo_11 () Bool)
// 	(declare-fun test1_b_buzz_4 () Bool)
// 	(assert (= test1_a_foo_0 false))
// 	(assert (= test1_a_zoo_0 false))
// 	(assert (= test1_b_buzz_0 false))
// 	(assert (= test1_b_bar_0 false))
// 	(assert (= test1_a_zoo_1 true))
// 	(assert (= test1_b_buzz_1 true))
// 	(assert (= test1_b_bar_1 true))
// 	(assert (= test1_a_foo_1 false))
// 	(assert (ite (= test1_a_foo_0 true) (and (= test1_b_bar_2 test1_b_bar_1) (= test1_a_foo_2 test1_a_foo_1)) (and (= test1_b_bar_2 test1_b_bar_0) (= test1_a_foo_2 test1_a_foo_0))))
// 	(assert (= test1_a_foo_3 true))
// 	(assert (= test1_a_zoo_2 false))
// 	(assert (ite (= test1_a_zoo_1 true) (and (= test1_a_zoo_3 test1_a_zoo_2) (= test1_a_foo_4 test1_a_foo_3)) (and (= test1_a_foo_4 test1_a_foo_2) (= test1_a_zoo_3 test1_a_zoo_1))))
// 	(assert (= test1_a_foo_5 true))
// 	(assert (= test1_b_buzz_2 false))
// 	(assert (ite (= test1_b_buzz_1 true) (and (= test1_a_foo_6 test1_a_foo_5) (= test1_b_buzz_3 test1_b_buzz_2)) (and (= test1_a_foo_6 test1_a_foo_4) (= test1_b_buzz_3 test1_b_buzz_1))))
// 	(assert (= test1_b_bar_3 true))
// 	(assert (= test1_a_foo_7 false))
// 	(assert (ite (= test1_a_foo_6 true) (and (= test1_b_bar_4 test1_b_bar_3) (= test1_a_foo_8 test1_a_foo_7)) (and (= test1_b_bar_4 test1_b_bar_2) (= test1_a_foo_8 test1_a_foo_6))))
// 	(assert (= test1_a_foo_9 true))
// 	(assert (= test1_a_zoo_4 false))
// 	(assert (ite (= test1_a_zoo_3 true) (and (= test1_a_foo_10 test1_a_foo_9) (= test1_a_zoo_5 test1_a_zoo_4)) (and (= test1_a_foo_10 test1_a_foo_8) (= test1_a_zoo_5 test1_a_zoo_3))))
// 	(assert (= test1_a_foo_11 true))
// 	(assert (= test1_b_buzz_4 false))
// 	(assert (ite (= test1_b_buzz_3 true) (and (= test1_a_foo_12 test1_a_foo_11) (= test1_b_buzz_5 test1_b_buzz_4)) (and (= test1_b_buzz_5 test1_b_buzz_3) (= test1_a_foo_12 test1_a_foo_10))))(assert (and (or (or test1_a_zoo_0 test1_b_buzz_0) (or (<= test1_b_bar_0 3) (not (= test1_a_foo_0 4)))) (or (or test1_a_zoo_1 test1_b_buzz_0) (or (<= test1_b_bar_0 3) (not (= test1_a_foo_0 4)))) (or (or test1_a_zoo_1 test1_b_buzz_1) (or (<= test1_b_bar_0 3) (not (= test1_a_foo_0 4)))) (or (or test1_a_zoo_2 test1_b_buzz_1) (or (<= test1_b_bar_1 3) (not (= test1_a_foo_1 4)))) (or (or test1_a_zoo_2 test1_b_buzz_1) (or (<= test1_b_bar_1 3) (not (= test1_a_foo_2 4)))) (or (or test1_a_zoo_2 test1_b_buzz_1) (or (<= test1_b_bar_1 3) (not (= test1_a_foo_3 4)))) (or (or test1_a_zoo_2 test1_b_buzz_1) (or (<= test1_b_bar_1 3) (not (= test1_a_foo_4 4)))) (or (or test1_a_zoo_2 test1_b_buzz_1) (or (<= test1_b_bar_1 3) (not (= test1_a_foo_5 4)))) (or (or test1_a_zoo_2 test1_b_buzz_1) (or (<= test1_b_bar_1 3) (not (= test1_a_foo_6 4)))) (or (or test1_a_zoo_2 test1_b_buzz_1) (or (<= test1_b_bar_2 3) (not (= test1_a_foo_1 4)))) (or (or test1_a_zoo_2 test1_b_buzz_1) (or (<= test1_b_bar_2 3) (not (= test1_a_foo_2 4)))) (or (or test1_a_zoo_2 test1_b_buzz_1) (or (<= test1_b_bar_2 3) (not (= test1_a_foo_3 4)))) (or (or test1_a_zoo_2 test1_b_buzz_1) (or (<= test1_b_bar_2 3) (not (= test1_a_foo_4 4)))) (or (or test1_a_zoo_2 test1_b_buzz_1) (or (<= test1_b_bar_2 3) (not (= test1_a_foo_5 4)))) (or (or test1_a_zoo_2 test1_b_buzz_1) (or (<= test1_b_bar_2 3) (not (= test1_a_foo_6 4)))) (or (or test1_a_zoo_3 test1_b_buzz_1) (or (<= test1_b_bar_1 3) (not (= test1_a_foo_1 4)))) (or (or test1_a_zoo_3 test1_b_buzz_1) (or (<= test1_b_bar_1 3) (not (= test1_a_foo_2 4)))) (or (or test1_a_zoo_3 test1_b_buzz_1) (or (<= test1_b_bar_1 3) (not (= test1_a_foo_3 4)))) (or (or test1_a_zoo_3 test1_b_buzz_1) (or (<= test1_b_bar_1 3) (not (= test1_a_foo_4 4)))) (or (or test1_a_zoo_3 test1_b_buzz_1) (or (<= test1_b_bar_1 3) (not (= test1_a_foo_5 4)))) (or (or test1_a_zoo_3 test1_b_buzz_1) (or (<= test1_b_bar_1 3) (not (= test1_a_foo_6 4)))) (or (or test1_a_zoo_3 test1_b_buzz_1) (or (<= test1_b_bar_2 3) (not (= test1_a_foo_1 4)))) (or (or test1_a_zoo_3 test1_b_buzz_1) (or (<= test1_b_bar_2 3) (not (= test1_a_foo_2 4)))) (or (or test1_a_zoo_3 test1_b_buzz_1) (or (<= test1_b_bar_2 3) (not (= test1_a_foo_3 4)))) (or (or test1_a_zoo_3 test1_b_buzz_1) (or (<= test1_b_bar_2 3) (not (= test1_a_foo_4 4)))) (or (or test1_a_zoo_3 test1_b_buzz_1) (or (<= test1_b_bar_2 3) (not (= test1_a_foo_5 4)))) (or (or test1_a_zoo_3 test1_b_buzz_1) (or (<= test1_b_bar_2 3) (not (= test1_a_foo_6 4)))) (or (or test1_a_zoo_3 test1_b_buzz_2) (or (<= test1_b_bar_1 3) (not (= test1_a_foo_1 4)))) (or (or test1_a_zoo_3 test1_b_buzz_2) (or (<= test1_b_bar_1 3) (not (= test1_a_foo_2 4)))) (or (or test1_a_zoo_3 test1_b_buzz_2) (or (<= test1_b_bar_1 3) (not (= test1_a_foo_3 4)))) (or (or test1_a_zoo_3 test1_b_buzz_2) (or (<= test1_b_bar_1 3) (not (= test1_a_foo_4 4)))) (or (or test1_a_zoo_3 test1_b_buzz_2) (or (<= test1_b_bar_1 3) (not (= test1_a_foo_5 4)))) (or (or test1_a_zoo_3 test1_b_buzz_2) (or (<= test1_b_bar_1 3) (not (= test1_a_foo_6 4)))) (or (or test1_a_zoo_3 test1_b_buzz_2) (or (<= test1_b_bar_2 3) (not (= test1_a_foo_1 4)))) (or (or test1_a_zoo_3 test1_b_buzz_2) (or (<= test1_b_bar_2 3) (not (= test1_a_foo_2 4)))) (or (or test1_a_zoo_3 test1_b_buzz_2) (or (<= test1_b_bar_2 3) (not (= test1_a_foo_3 4)))) (or (or test1_a_zoo_3 test1_b_buzz_2) (or (<= test1_b_bar_2 3) (not (= test1_a_foo_4 4)))) (or (or test1_a_zoo_3 test1_b_buzz_2) (or (<= test1_b_bar_2 3) (not (= test1_a_foo_5 4)))) (or (or test1_a_zoo_3 test1_b_buzz_2) (or (<= test1_b_bar_2 3) (not (= test1_a_foo_6 4)))) (or (or test1_a_zoo_3 test1_b_buzz_3) (or (<= test1_b_bar_1 3) (not (= test1_a_foo_1 4)))) (or (or test1_a_zoo_3 test1_b_buzz_3) (or (<= test1_b_bar_1 3) (not (= test1_a_foo_2 4)))) (or (or test1_a_zoo_3 test1_b_buzz_3) (or (<= test1_b_bar_1 3) (not (= test1_a_foo_3 4)))) (or (or test1_a_zoo_3 test1_b_buzz_3) (or (<= test1_b_bar_1 3) (not (= test1_a_foo_4 4)))) (or (or test1_a_zoo_3 test1_b_buzz_3) (or (<= test1_b_bar_1 3) (not (= test1_a_foo_5 4)))) (or (or test1_a_zoo_3 test1_b_buzz_3) (or (<= test1_b_bar_1 3) (not (= test1_a_foo_6 4)))) (or (or test1_a_zoo_3 test1_b_buzz_3) (or (<= test1_b_bar_2 3) (not (= test1_a_foo_1 4)))) (or (or test1_a_zoo_3 test1_b_buzz_3) (or (<= test1_b_bar_2 3) (not (= test1_a_foo_2 4)))) (or (or test1_a_zoo_3 test1_b_buzz_3) (or (<= test1_b_bar_2 3) (not (= test1_a_foo_3 4)))) (or (or test1_a_zoo_3 test1_b_buzz_3) (or (<= test1_b_bar_2 3) (not (= test1_a_foo_4 4)))) (or (or test1_a_zoo_3 test1_b_buzz_3) (or (<= test1_b_bar_2 3) (not (= test1_a_foo_5 4)))) (or (or test1_a_zoo_3 test1_b_buzz_3) (or (<= test1_b_bar_2 3) (not (= test1_a_foo_6 4)))) (or (or test1_a_zoo_4 test1_b_buzz_3) (or (<= test1_b_bar_3 3) (not (= test1_a_foo_7 4)))) (or (or test1_a_zoo_4 test1_b_buzz_3) (or (<= test1_b_bar_3 3) (not (= test1_a_foo_8 4)))) (or (or test1_a_zoo_4 test1_b_buzz_3) (or (<= test1_b_bar_3 3) (not (= test1_a_foo_9 4)))) (or (or test1_a_zoo_4 test1_b_buzz_3) (or (<= test1_b_bar_3 3) (not (= test1_a_foo_10 4)))) (or (or test1_a_zoo_4 test1_b_buzz_3) (or (<= test1_b_bar_3 3) (not (= test1_a_foo_11 4)))) (or (or test1_a_zoo_4 test1_b_buzz_3) (or (<= test1_b_bar_3 3) (not (= test1_a_foo_12 4)))) (or (or test1_a_zoo_4 test1_b_buzz_3) (or (<= test1_b_bar_4 3) (not (= test1_a_foo_7 4)))) (or (or test1_a_zoo_4 test1_b_buzz_3) (or (<= test1_b_bar_4 3) (not (= test1_a_foo_8 4)))) (or (or test1_a_zoo_4 test1_b_buzz_3) (or (<= test1_b_bar_4 3) (not (= test1_a_foo_9 4)))) (or (or test1_a_zoo_4 test1_b_buzz_3) (or (<= test1_b_bar_4 3) (not (= test1_a_foo_10 4)))) (or (or test1_a_zoo_4 test1_b_buzz_3) (or (<= test1_b_bar_4 3) (not (= test1_a_foo_11 4)))) (or (or test1_a_zoo_4 test1_b_buzz_3) (or (<= test1_b_bar_4 3) (not (= test1_a_foo_12 4)))) (or (or test1_a_zoo_5 test1_b_buzz_3) (or (<= test1_b_bar_3 3) (not (= test1_a_foo_7 4)))) (or (or test1_a_zoo_5 test1_b_buzz_3) (or (<= test1_b_bar_3 3) (not (= test1_a_foo_8 4)))) (or (or test1_a_zoo_5 test1_b_buzz_3) (or (<= test1_b_bar_3 3) (not (= test1_a_foo_9 4)))) (or (or test1_a_zoo_5 test1_b_buzz_3) (or (<= test1_b_bar_3 3) (not (= test1_a_foo_10 4)))) (or (or test1_a_zoo_5 test1_b_buzz_3) (or (<= test1_b_bar_3 3) (not (= test1_a_foo_11 4)))) (or (or test1_a_zoo_5 test1_b_buzz_3) (or (<= test1_b_bar_3 3) (not (= test1_a_foo_12 4)))) (or (or test1_a_zoo_5 test1_b_buzz_3) (or (<= test1_b_bar_4 3) (not (= test1_a_foo_7 4)))) (or (or test1_a_zoo_5 test1_b_buzz_3) (or (<= test1_b_bar_4 3) (not (= test1_a_foo_8 4)))) (or (or test1_a_zoo_5 test1_b_buzz_3) (or (<= test1_b_bar_4 3) (not (= test1_a_foo_9 4)))) (or (or test1_a_zoo_5 test1_b_buzz_3) (or (<= test1_b_bar_4 3) (not (= test1_a_foo_10 4)))) (or (or test1_a_zoo_5 test1_b_buzz_3) (or (<= test1_b_bar_4 3) (not (= test1_a_foo_11 4)))) (or (or test1_a_zoo_5 test1_b_buzz_3) (or (<= test1_b_bar_4 3) (not (= test1_a_foo_12 4)))) (or (or test1_a_zoo_5 test1_b_buzz_4) (or (<= test1_b_bar_3 3) (not (= test1_a_foo_7 4)))) (or (or test1_a_zoo_5 test1_b_buzz_4) (or (<= test1_b_bar_3 3) (not (= test1_a_foo_8 4)))) (or (or test1_a_zoo_5 test1_b_buzz_4) (or (<= test1_b_bar_3 3) (not (= test1_a_foo_9 4)))) (or (or test1_a_zoo_5 test1_b_buzz_4) (or (<= test1_b_bar_3 3) (not (= test1_a_foo_10 4)))) (or (or test1_a_zoo_5 test1_b_buzz_4) (or (<= test1_b_bar_3 3) (not (= test1_a_foo_11 4)))) (or (or test1_a_zoo_5 test1_b_buzz_4) (or (<= test1_b_bar_3 3) (not (= test1_a_foo_12 4)))) (or (or test1_a_zoo_5 test1_b_buzz_4) (or (<= test1_b_bar_4 3) (not (= test1_a_foo_7 4)))) (or (or test1_a_zoo_5 test1_b_buzz_4) (or (<= test1_b_bar_4 3) (not (= test1_a_foo_8 4)))) (or (or test1_a_zoo_5 test1_b_buzz_4) (or (<= test1_b_bar_4 3) (not (= test1_a_foo_9 4)))) (or (or test1_a_zoo_5 test1_b_buzz_4) (or (<= test1_b_bar_4 3) (not (= test1_a_foo_10 4)))) (or (or test1_a_zoo_5 test1_b_buzz_4) (or (<= test1_b_bar_4 3) (not (= test1_a_foo_11 4)))) (or (or test1_a_zoo_5 test1_b_buzz_4) (or (<= test1_b_bar_4 3) (not (= test1_a_foo_12 4)))) (or (or test1_a_zoo_5 test1_b_buzz_5) (or (<= test1_b_bar_3 3) (not (= test1_a_foo_7 4)))) (or (or test1_a_zoo_5 test1_b_buzz_5) (or (<= test1_b_bar_3 3) (not (= test1_a_foo_8 4)))) (or (or test1_a_zoo_5 test1_b_buzz_5) (or (<= test1_b_bar_3 3) (not (= test1_a_foo_9 4)))) (or (or test1_a_zoo_5 test1_b_buzz_5) (or (<= test1_b_bar_3 3) (not (= test1_a_foo_10 4)))) (or (or test1_a_zoo_5 test1_b_buzz_5) (or (<= test1_b_bar_3 3) (not (= test1_a_foo_11 4)))) (or (or test1_a_zoo_5 test1_b_buzz_5) (or (<= test1_b_bar_3 3) (not (= test1_a_foo_12 4)))) (or (or test1_a_zoo_5 test1_b_buzz_5) (or (<= test1_b_bar_4 3) (not (= test1_a_foo_7 4)))) (or (or test1_a_zoo_5 test1_b_buzz_5) (or (<= test1_b_bar_4 3) (not (= test1_a_foo_8 4)))) (or (or test1_a_zoo_5 test1_b_buzz_5) (or (<= test1_b_bar_4 3) (not (= test1_a_foo_9 4)))) (or (or test1_a_zoo_5 test1_b_buzz_5) (or (<= test1_b_bar_4 3) (not (= test1_a_foo_10 4)))) (or (or test1_a_zoo_5 test1_b_buzz_5) (or (<= test1_b_bar_4 3) (not (= test1_a_foo_11 4)))) (or (or test1_a_zoo_5 test1_b_buzz_5) (or (<= test1_b_bar_4 3) (not (= test1_a_foo_12 4))))))`

// 	g := prepTest("", test, false, false)

// 	err := compareResults("TemporalSys", g.SMT(), string(expecting))

// 	if err != nil {
// 		t.Fatal(err.Error())
// 	}
// }

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
			//fmt.Println(g.SMT())
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

func TestBadSpecs(t *testing.T) {
	type specCase struct {
		path        string
		specType    bool
		expectedErr string // non-empty: expect an error containing this string
	}
	specs := []specCase{
		{"testdata/badspecs/nodefs.fspec", true, "Missing run block or start block"},
		{"testdata/badspecs/doubleswap.fspec", true, "swapped more than once"},
		{"testdata/badspecs/sharedstate.fspec", true, ""},
		{"testdata/badspecs/deep.fspec", true, ""},
		{"testdata/badspecs/zerounds.fspec", true, "zero-round loop"},
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
				ty := types.Execute(pre.Processed, pre)
				sw := swaps.NewPrecompiler(ty)
				tree := sw.Swap(ty.Checked)
				compiler, err := llvm.Execute(tree, ty.SpecStructs, l.Uncertains, l.Unknowns, sw.Alias, false)
				if err != nil {
					pipelineErr = fmt.Errorf("llvm: %w", err)
					return
				}
				result = Execute(compiler).SMT()
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
	// Normalize block variable names: block<n>true/block<n>false → blocktrue/blockfalse
	// so that LLVM IR block renumbering (e.g. from optimization passes) doesn't break tests.
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

func TestSwapsNested(t *testing.T) {
	// A flow whose target is a stock containing a nested stock.
	// After swap, accesses to target.sub.x must resolve to the base
	// stock variable (s_sub_x), not the flow-local name (f_target_sub_x).
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

for 2 init{
	s = new outstock;
	f = new f1;
	f.target = s;
} run {
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
	// A single flow with two swaps: both properties must resolve
	// to their respective base stock variables after a single
	// swapDeepNames call (not one per swap).
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

for 2 init{
	sa = new stock1;
	sb = new stock2;
	f = new f1;
	f.addtarget = sa;
	f.subtarget = sb;
} run {
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
	// A stock with two properties: only "used" is accessed in the flow function.
	// The "unused" property should be absent from the generated SMT.
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

	for 1 init{ inst = new f; } run { inst.fn; };
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
	// True branch modifies snap.a; false branch modifies snap.b.
	// With the phi completeness fix, both variables must be fully constrained
	// in both branches of the ite assertion (identity rules added for the branch
	// that doesn't directly modify a variable).
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

	for 1 init{ inst = new f; } run { inst.fn; };
	`
	g := prepTest("", test, true, false)
	smt := stripAndEscape(g.SMT())

	iteIdx := strings.Index(smt, "(assert(ite")
	if iteIdx < 0 {
		t.Fatal("no ite assertion found in SMT")
	}
	iteExpr := smt[iteIdx:]

	// snap.b is only modified in the false branch. Without the phi completeness fix,
	// snap.b_ would only appear in the false side of the ite.
	// With the fix, snap.b_ also appears in the true side as an identity rule
	// (= b_phi b_entry), so the count across the whole ite expression is higher.
	// False branch alone: b_phi(LHS) + b_1(assign) + b_phi(LHS) + b_1(RHS) → varies
	// Both branches: additional b_phi(LHS) + b_entry(RHS) in true branch.
	bCount := strings.Count(iteExpr, "phitest_inst_snap_b_")
	if bCount < 4 {
		t.Fatalf("phitest_inst_snap_b_ appears %d times in ite expression (want >= 4: constrained in both branches).\nite=%s", bCount, iteExpr)
	}

	// snap.a is only modified in the true branch. Without the fix, snap.a_ would
	// only appear in the condition and true side. With the fix, snap.a_ also appears
	// in the false side as an identity rule.
	aCount := strings.Count(iteExpr, "phitest_inst_snap_a_")
	if aCount < 5 {
		t.Fatalf("phitest_inst_snap_a_ appears %d times in ite expression (want >= 5: constrained in both branches).\nite=%s", aCount, iteExpr)
	}
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
	ty := types.Execute(pre.Processed, pre)
	sw := swaps.NewPrecompiler(ty)
	tree := sw.Swap(ty.Checked)
	compiler, err := llvm.Execute(tree, ty.SpecStructs, l.Uncertains, l.Unknowns, sw.Alias, true)
	if err != nil {
		panic(err)
	}

	//fmt.Println(compiler.GetIR())
	generator := Execute(compiler)
	return generator
}

func notStrictlyOrdered(want string, got string) bool {
	// Fixing cases where lines of SMT end up in slightly
	// different orders. Only runs when shallow string
	// compare fails

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
