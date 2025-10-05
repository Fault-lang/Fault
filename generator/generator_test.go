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
(assert (= test1_t_u_x_5 (+ test1_t_u_x_4 (+ test1_a_0 test1_b_0))))
(assert (and (not (= test1_t_u_x_0 11)) (not (= test1_t_u_x_1 11)) (not (= test1_t_u_x_2 11)) (not (= test1_t_u_x_3 11)) (not (= test1_t_u_x_4 11)) (not (= test1_t_u_x_5 11))))
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
	expecting := `(set-logic QF_NRA)
	(declare-fun test1_b_bar_2 () Bool)
	(declare-fun test1_a_foo_2 () Bool)
	(declare-fun test1_a_foo_4 () Bool)
	(declare-fun test1_a_zoo_3 () Bool)
	(declare-fun test1_a_foo_6 () Bool)
	(declare-fun test1_b_buzz_3 () Bool)
	(declare-fun test1_a_foo_0 () Bool)
	(declare-fun test1_a_zoo_0 () Bool)
	(declare-fun test1_b_buzz_0 () Bool)
	(declare-fun test1_b_bar_0 () Bool)
	(declare-fun test1_a_zoo_1 () Bool)
	(declare-fun test1_b_buzz_1 () Bool)
	(declare-fun test1_b_bar_1 () Bool)
	(declare-fun test1_a_foo_1 () Bool)
	(declare-fun test1_a_foo_3 () Bool)
	(declare-fun test1_a_zoo_2 () Bool)
	(declare-fun test1_a_foo_5 () Bool)
	(declare-fun test1_b_buzz_2 () Bool)
	(assert (= test1_a_foo_0 false))
	(assert (= test1_a_zoo_0 false))
	(assert (= test1_b_buzz_0 false))
	(assert (= test1_b_bar_0 false))
	(assert (= test1_a_zoo_1 true))
	(assert (= test1_b_buzz_1 true))
	(assert (= test1_b_bar_1 true))
	(assert (= test1_a_foo_1 false))
	(assert (ite (= test1_a_foo_0 true) (and (= test1_b_bar_2 test1_b_bar_1) (= test1_a_foo_2 test1_a_foo_1)) (and (= test1_b_bar_2 test1_b_bar_0) (= test1_a_foo_2 test1_a_foo_0))))
	(assert (= test1_a_foo_3 true))
	(assert (= test1_a_zoo_2 false))
	(assert (ite (= test1_a_zoo_1 true) (and (= test1_a_foo_4 test1_a_foo_3) (= test1_a_zoo_3 test1_a_zoo_2)) (and (= test1_a_foo_4 test1_a_foo_2) (= test1_a_zoo_3 test1_a_zoo_1))))
	(assert (= test1_a_foo_5 true))
	(assert (= test1_b_buzz_2 false))
	(assert (ite (= test1_b_buzz_1 true) (and (= test1_a_foo_6 test1_a_foo_5) (= test1_b_buzz_3 test1_b_buzz_2)) (and (= test1_a_foo_6 test1_a_foo_4) (= test1_b_buzz_3 test1_b_buzz_1))))(assert (and (or test1_a_zoo_0 test1_b_bar_0) (or test1_a_zoo_1 test1_b_bar_0) (or test1_a_zoo_1 test1_b_bar_1) (or test1_a_zoo_1 test1_b_bar_2) (or test1_a_zoo_2 test1_b_bar_2) (or test1_a_zoo_3 test1_b_bar_2)))
	`

	g := prepTest("", test, false, false)

	err := compareResults("TemporalSys", g.SMT(), string(expecting))

	if err != nil {
		t.Fatal(err.Error())
	}
}

// func TestTemporalSys2(t *testing.T) {

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
		{"testdata/statecharts/triggerfunc.fsystem", "0"},
	}
	smt2s := []string{
		"testdata/statecharts/statechart.smt2",
		"testdata/statecharts/advanceor.smt2",
		"testdata/statecharts/multioradvance.smt2",
		"testdata/statecharts/advanceand.smt2",
		"testdata/statecharts/mixedcalls.smt2",
		"testdata/statecharts/triggerfunc.smt2",
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
			t.Fatal(err.Error())
		}
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

	l := listener.Execute(test, path, flags)
	pre := preprocess.Execute(l)
	ty := types.Execute(pre.Processed, pre)
	sw := swaps.NewPrecompiler(ty)
	tree := sw.Swap(ty.Checked)
	compiler := llvm.Execute(tree, ty.SpecStructs, l.Uncertains, l.Unknowns, sw.Alias, true)

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
