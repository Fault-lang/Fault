package smt

import (
	"testing"
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

	for 5 run {
		t = new test;
		t.bar;
	};
	`
	expecting := `(declare-fun test1_t_foo_value_0 () Real)
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

	smt, err := prepTest("", test)

	if err != nil {
		t.Fatalf("compilation failed on valid spec. got=%s", err)
	}

	err = compareResults("Eventually", smt, expecting)

	if err != nil {
		t.Fatalf(err.Error())
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

	for 5 run {
		t = new test;
		t.bar;
	};
	`
	expecting := `(declare-fun test1_t_foo_value_0 () Real)
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
		(and (> test1_t_foo_value_5 0))
		))
`

	smt, err := prepTest("", test)

	if err != nil {
		t.Fatalf("compilation failed on valid spec. got=%s", err)
	}

	err = compareResults("EventuallyAlways", smt, expecting)

	if err != nil {
		t.Fatalf(err.Error())
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

	for 5 run {
		t = new test;
		t.bar;
	};
	`
	expecting := `(declare-fun test1_t_foo_value_0 () Real)
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
		(and (<= test1_t_foo_value_5 0))
		))
`

	smt, err := prepTest("", test)

	if err != nil {
		t.Fatalf("compilation failed on valid spec. got=%s", err)
	}

	err = compareResults("EventuallyAlways2", smt, expecting)

	if err != nil {
		t.Fatalf(err.Error())
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

	for 5 run {
		t = new test;
		t.bar;
	};
	`
	expecting := `(declare-fun test1_t_foo_value_0 () Real)
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

	smt, err := prepTest("", test)

	if err != nil {
		t.Fatalf("compilation failed on valid spec. got=%s", err)
	}

	err = compareResults("Temporal", smt, expecting)

	if err != nil {
		t.Fatalf(err.Error())
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

		for 5 run {
			t = new test;
			t.bar;
		};
	`
	expecting := `
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

	smt, err := prepTest("", test)

	if err != nil {
		t.Fatalf("compilation failed on valid spec. got=%s", err)
	}

	err = compareResults("Temporal2", smt, string(expecting))

	if err != nil {
		t.Fatalf(err.Error())
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
	expecting := `(declare-fun test1_b_bar_2 () Bool)
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

	smt, err := prepTestSys("", test, false)

	if err != nil {
		t.Fatalf("compilation failed on valid spec. got=%s", err)
	}

	err = compareResults("TemporalSys", smt, string(expecting))

	if err != nil {
		t.Fatalf(err.Error())
	}
}

func TestTemporalSys2(t *testing.T) {

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

		assert when a.zoo && b.buzz then b.bar > 3 && a.foo == 4;

		start{
			b: buzz,
			a: zoo,
		};

		for 2 run{};
		`
	expecting := `(declare-fun test1_b_bar_2 () Bool)
	(declare-fun test1_a_foo_2 () Bool)
	(declare-fun test1_a_zoo_3 () Bool)
	(declare-fun test1_a_foo_4 () Bool)
	(declare-fun test1_a_foo_6 () Bool)
	(declare-fun test1_b_buzz_3 () Bool)
	(declare-fun test1_b_bar_4 () Bool)
	(declare-fun test1_a_foo_8 () Bool)
	(declare-fun test1_a_foo_10 () Bool)
	(declare-fun test1_a_zoo_5 () Bool)
	(declare-fun test1_a_foo_12 () Bool)
	(declare-fun test1_b_buzz_5 () Bool)
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
	(declare-fun test1_b_bar_3 () Bool)
	(declare-fun test1_a_foo_7 () Bool)
	(declare-fun test1_a_foo_9 () Bool)
	(declare-fun test1_a_zoo_4 () Bool)
	(declare-fun test1_a_foo_11 () Bool)
	(declare-fun test1_b_buzz_4 () Bool)
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
	(assert (ite (= test1_a_zoo_1 true) (and (= test1_a_zoo_3 test1_a_zoo_2) (= test1_a_foo_4 test1_a_foo_3)) (and (= test1_a_foo_4 test1_a_foo_2) (= test1_a_zoo_3 test1_a_zoo_1))))
	(assert (= test1_a_foo_5 true))
	(assert (= test1_b_buzz_2 false))
	(assert (ite (= test1_b_buzz_1 true) (and (= test1_a_foo_6 test1_a_foo_5) (= test1_b_buzz_3 test1_b_buzz_2)) (and (= test1_a_foo_6 test1_a_foo_4) (= test1_b_buzz_3 test1_b_buzz_1))))
	(assert (= test1_b_bar_3 true))
	(assert (= test1_a_foo_7 false))
	(assert (ite (= test1_a_foo_6 true) (and (= test1_b_bar_4 test1_b_bar_3) (= test1_a_foo_8 test1_a_foo_7)) (and (= test1_b_bar_4 test1_b_bar_2) (= test1_a_foo_8 test1_a_foo_6))))
	(assert (= test1_a_foo_9 true))
	(assert (= test1_a_zoo_4 false))
	(assert (ite (= test1_a_zoo_3 true) (and (= test1_a_foo_10 test1_a_foo_9) (= test1_a_zoo_5 test1_a_zoo_4)) (and (= test1_a_foo_10 test1_a_foo_8) (= test1_a_zoo_5 test1_a_zoo_3))))
	(assert (= test1_a_foo_11 true))
	(assert (= test1_b_buzz_4 false))
	(assert (ite (= test1_b_buzz_3 true) (and (= test1_a_foo_12 test1_a_foo_11) (= test1_b_buzz_5 test1_b_buzz_4)) (and (= test1_b_buzz_5 test1_b_buzz_3) (= test1_a_foo_12 test1_a_foo_10))))(assert (and (or (or test1_a_zoo_0 test1_b_buzz_0) (or (<= test1_b_bar_0 3) (not (= test1_a_foo_0 4)))) (or (or test1_a_zoo_1 test1_b_buzz_0) (or (<= test1_b_bar_0 3) (not (= test1_a_foo_0 4)))) (or (or test1_a_zoo_1 test1_b_buzz_1) (or (<= test1_b_bar_0 3) (not (= test1_a_foo_0 4)))) (or (or test1_a_zoo_2 test1_b_buzz_1) (or (<= test1_b_bar_1 3) (not (= test1_a_foo_1 4)))) (or (or test1_a_zoo_2 test1_b_buzz_1) (or (<= test1_b_bar_1 3) (not (= test1_a_foo_2 4)))) (or (or test1_a_zoo_2 test1_b_buzz_1) (or (<= test1_b_bar_1 3) (not (= test1_a_foo_3 4)))) (or (or test1_a_zoo_2 test1_b_buzz_1) (or (<= test1_b_bar_1 3) (not (= test1_a_foo_4 4)))) (or (or test1_a_zoo_2 test1_b_buzz_1) (or (<= test1_b_bar_1 3) (not (= test1_a_foo_5 4)))) (or (or test1_a_zoo_2 test1_b_buzz_1) (or (<= test1_b_bar_1 3) (not (= test1_a_foo_6 4)))) (or (or test1_a_zoo_2 test1_b_buzz_1) (or (<= test1_b_bar_2 3) (not (= test1_a_foo_1 4)))) (or (or test1_a_zoo_2 test1_b_buzz_1) (or (<= test1_b_bar_2 3) (not (= test1_a_foo_2 4)))) (or (or test1_a_zoo_2 test1_b_buzz_1) (or (<= test1_b_bar_2 3) (not (= test1_a_foo_3 4)))) (or (or test1_a_zoo_2 test1_b_buzz_1) (or (<= test1_b_bar_2 3) (not (= test1_a_foo_4 4)))) (or (or test1_a_zoo_2 test1_b_buzz_1) (or (<= test1_b_bar_2 3) (not (= test1_a_foo_5 4)))) (or (or test1_a_zoo_2 test1_b_buzz_1) (or (<= test1_b_bar_2 3) (not (= test1_a_foo_6 4)))) (or (or test1_a_zoo_3 test1_b_buzz_1) (or (<= test1_b_bar_1 3) (not (= test1_a_foo_1 4)))) (or (or test1_a_zoo_3 test1_b_buzz_1) (or (<= test1_b_bar_1 3) (not (= test1_a_foo_2 4)))) (or (or test1_a_zoo_3 test1_b_buzz_1) (or (<= test1_b_bar_1 3) (not (= test1_a_foo_3 4)))) (or (or test1_a_zoo_3 test1_b_buzz_1) (or (<= test1_b_bar_1 3) (not (= test1_a_foo_4 4)))) (or (or test1_a_zoo_3 test1_b_buzz_1) (or (<= test1_b_bar_1 3) (not (= test1_a_foo_5 4)))) (or (or test1_a_zoo_3 test1_b_buzz_1) (or (<= test1_b_bar_1 3) (not (= test1_a_foo_6 4)))) (or (or test1_a_zoo_3 test1_b_buzz_1) (or (<= test1_b_bar_2 3) (not (= test1_a_foo_1 4)))) (or (or test1_a_zoo_3 test1_b_buzz_1) (or (<= test1_b_bar_2 3) (not (= test1_a_foo_2 4)))) (or (or test1_a_zoo_3 test1_b_buzz_1) (or (<= test1_b_bar_2 3) (not (= test1_a_foo_3 4)))) (or (or test1_a_zoo_3 test1_b_buzz_1) (or (<= test1_b_bar_2 3) (not (= test1_a_foo_4 4)))) (or (or test1_a_zoo_3 test1_b_buzz_1) (or (<= test1_b_bar_2 3) (not (= test1_a_foo_5 4)))) (or (or test1_a_zoo_3 test1_b_buzz_1) (or (<= test1_b_bar_2 3) (not (= test1_a_foo_6 4)))) (or (or test1_a_zoo_3 test1_b_buzz_2) (or (<= test1_b_bar_1 3) (not (= test1_a_foo_1 4)))) (or (or test1_a_zoo_3 test1_b_buzz_2) (or (<= test1_b_bar_1 3) (not (= test1_a_foo_2 4)))) (or (or test1_a_zoo_3 test1_b_buzz_2) (or (<= test1_b_bar_1 3) (not (= test1_a_foo_3 4)))) (or (or test1_a_zoo_3 test1_b_buzz_2) (or (<= test1_b_bar_1 3) (not (= test1_a_foo_4 4)))) (or (or test1_a_zoo_3 test1_b_buzz_2) (or (<= test1_b_bar_1 3) (not (= test1_a_foo_5 4)))) (or (or test1_a_zoo_3 test1_b_buzz_2) (or (<= test1_b_bar_1 3) (not (= test1_a_foo_6 4)))) (or (or test1_a_zoo_3 test1_b_buzz_2) (or (<= test1_b_bar_2 3) (not (= test1_a_foo_1 4)))) (or (or test1_a_zoo_3 test1_b_buzz_2) (or (<= test1_b_bar_2 3) (not (= test1_a_foo_2 4)))) (or (or test1_a_zoo_3 test1_b_buzz_2) (or (<= test1_b_bar_2 3) (not (= test1_a_foo_3 4)))) (or (or test1_a_zoo_3 test1_b_buzz_2) (or (<= test1_b_bar_2 3) (not (= test1_a_foo_4 4)))) (or (or test1_a_zoo_3 test1_b_buzz_2) (or (<= test1_b_bar_2 3) (not (= test1_a_foo_5 4)))) (or (or test1_a_zoo_3 test1_b_buzz_2) (or (<= test1_b_bar_2 3) (not (= test1_a_foo_6 4)))) (or (or test1_a_zoo_3 test1_b_buzz_3) (or (<= test1_b_bar_1 3) (not (= test1_a_foo_1 4)))) (or (or test1_a_zoo_3 test1_b_buzz_3) (or (<= test1_b_bar_1 3) (not (= test1_a_foo_2 4)))) (or (or test1_a_zoo_3 test1_b_buzz_3) (or (<= test1_b_bar_1 3) (not (= test1_a_foo_3 4)))) (or (or test1_a_zoo_3 test1_b_buzz_3) (or (<= test1_b_bar_1 3) (not (= test1_a_foo_4 4)))) (or (or test1_a_zoo_3 test1_b_buzz_3) (or (<= test1_b_bar_1 3) (not (= test1_a_foo_5 4)))) (or (or test1_a_zoo_3 test1_b_buzz_3) (or (<= test1_b_bar_1 3) (not (= test1_a_foo_6 4)))) (or (or test1_a_zoo_3 test1_b_buzz_3) (or (<= test1_b_bar_2 3) (not (= test1_a_foo_1 4)))) (or (or test1_a_zoo_3 test1_b_buzz_3) (or (<= test1_b_bar_2 3) (not (= test1_a_foo_2 4)))) (or (or test1_a_zoo_3 test1_b_buzz_3) (or (<= test1_b_bar_2 3) (not (= test1_a_foo_3 4)))) (or (or test1_a_zoo_3 test1_b_buzz_3) (or (<= test1_b_bar_2 3) (not (= test1_a_foo_4 4)))) (or (or test1_a_zoo_3 test1_b_buzz_3) (or (<= test1_b_bar_2 3) (not (= test1_a_foo_5 4)))) (or (or test1_a_zoo_3 test1_b_buzz_3) (or (<= test1_b_bar_2 3) (not (= test1_a_foo_6 4)))) (or (or test1_a_zoo_4 test1_b_buzz_3) (or (<= test1_b_bar_3 3) (not (= test1_a_foo_7 4)))) (or (or test1_a_zoo_4 test1_b_buzz_3) (or (<= test1_b_bar_3 3) (not (= test1_a_foo_8 4)))) (or (or test1_a_zoo_4 test1_b_buzz_3) (or (<= test1_b_bar_3 3) (not (= test1_a_foo_9 4)))) (or (or test1_a_zoo_4 test1_b_buzz_3) (or (<= test1_b_bar_3 3) (not (= test1_a_foo_10 4)))) (or (or test1_a_zoo_4 test1_b_buzz_3) (or (<= test1_b_bar_3 3) (not (= test1_a_foo_11 4)))) (or (or test1_a_zoo_4 test1_b_buzz_3) (or (<= test1_b_bar_3 3) (not (= test1_a_foo_12 4)))) (or (or test1_a_zoo_4 test1_b_buzz_3) (or (<= test1_b_bar_4 3) (not (= test1_a_foo_7 4)))) (or (or test1_a_zoo_4 test1_b_buzz_3) (or (<= test1_b_bar_4 3) (not (= test1_a_foo_8 4)))) (or (or test1_a_zoo_4 test1_b_buzz_3) (or (<= test1_b_bar_4 3) (not (= test1_a_foo_9 4)))) (or (or test1_a_zoo_4 test1_b_buzz_3) (or (<= test1_b_bar_4 3) (not (= test1_a_foo_10 4)))) (or (or test1_a_zoo_4 test1_b_buzz_3) (or (<= test1_b_bar_4 3) (not (= test1_a_foo_11 4)))) (or (or test1_a_zoo_4 test1_b_buzz_3) (or (<= test1_b_bar_4 3) (not (= test1_a_foo_12 4)))) (or (or test1_a_zoo_5 test1_b_buzz_3) (or (<= test1_b_bar_3 3) (not (= test1_a_foo_7 4)))) (or (or test1_a_zoo_5 test1_b_buzz_3) (or (<= test1_b_bar_3 3) (not (= test1_a_foo_8 4)))) (or (or test1_a_zoo_5 test1_b_buzz_3) (or (<= test1_b_bar_3 3) (not (= test1_a_foo_9 4)))) (or (or test1_a_zoo_5 test1_b_buzz_3) (or (<= test1_b_bar_3 3) (not (= test1_a_foo_10 4)))) (or (or test1_a_zoo_5 test1_b_buzz_3) (or (<= test1_b_bar_3 3) (not (= test1_a_foo_11 4)))) (or (or test1_a_zoo_5 test1_b_buzz_3) (or (<= test1_b_bar_3 3) (not (= test1_a_foo_12 4)))) (or (or test1_a_zoo_5 test1_b_buzz_3) (or (<= test1_b_bar_4 3) (not (= test1_a_foo_7 4)))) (or (or test1_a_zoo_5 test1_b_buzz_3) (or (<= test1_b_bar_4 3) (not (= test1_a_foo_8 4)))) (or (or test1_a_zoo_5 test1_b_buzz_3) (or (<= test1_b_bar_4 3) (not (= test1_a_foo_9 4)))) (or (or test1_a_zoo_5 test1_b_buzz_3) (or (<= test1_b_bar_4 3) (not (= test1_a_foo_10 4)))) (or (or test1_a_zoo_5 test1_b_buzz_3) (or (<= test1_b_bar_4 3) (not (= test1_a_foo_11 4)))) (or (or test1_a_zoo_5 test1_b_buzz_3) (or (<= test1_b_bar_4 3) (not (= test1_a_foo_12 4)))) (or (or test1_a_zoo_5 test1_b_buzz_4) (or (<= test1_b_bar_3 3) (not (= test1_a_foo_7 4)))) (or (or test1_a_zoo_5 test1_b_buzz_4) (or (<= test1_b_bar_3 3) (not (= test1_a_foo_8 4)))) (or (or test1_a_zoo_5 test1_b_buzz_4) (or (<= test1_b_bar_3 3) (not (= test1_a_foo_9 4)))) (or (or test1_a_zoo_5 test1_b_buzz_4) (or (<= test1_b_bar_3 3) (not (= test1_a_foo_10 4)))) (or (or test1_a_zoo_5 test1_b_buzz_4) (or (<= test1_b_bar_3 3) (not (= test1_a_foo_11 4)))) (or (or test1_a_zoo_5 test1_b_buzz_4) (or (<= test1_b_bar_3 3) (not (= test1_a_foo_12 4)))) (or (or test1_a_zoo_5 test1_b_buzz_4) (or (<= test1_b_bar_4 3) (not (= test1_a_foo_7 4)))) (or (or test1_a_zoo_5 test1_b_buzz_4) (or (<= test1_b_bar_4 3) (not (= test1_a_foo_8 4)))) (or (or test1_a_zoo_5 test1_b_buzz_4) (or (<= test1_b_bar_4 3) (not (= test1_a_foo_9 4)))) (or (or test1_a_zoo_5 test1_b_buzz_4) (or (<= test1_b_bar_4 3) (not (= test1_a_foo_10 4)))) (or (or test1_a_zoo_5 test1_b_buzz_4) (or (<= test1_b_bar_4 3) (not (= test1_a_foo_11 4)))) (or (or test1_a_zoo_5 test1_b_buzz_4) (or (<= test1_b_bar_4 3) (not (= test1_a_foo_12 4)))) (or (or test1_a_zoo_5 test1_b_buzz_5) (or (<= test1_b_bar_3 3) (not (= test1_a_foo_7 4)))) (or (or test1_a_zoo_5 test1_b_buzz_5) (or (<= test1_b_bar_3 3) (not (= test1_a_foo_8 4)))) (or (or test1_a_zoo_5 test1_b_buzz_5) (or (<= test1_b_bar_3 3) (not (= test1_a_foo_9 4)))) (or (or test1_a_zoo_5 test1_b_buzz_5) (or (<= test1_b_bar_3 3) (not (= test1_a_foo_10 4)))) (or (or test1_a_zoo_5 test1_b_buzz_5) (or (<= test1_b_bar_3 3) (not (= test1_a_foo_11 4)))) (or (or test1_a_zoo_5 test1_b_buzz_5) (or (<= test1_b_bar_3 3) (not (= test1_a_foo_12 4)))) (or (or test1_a_zoo_5 test1_b_buzz_5) (or (<= test1_b_bar_4 3) (not (= test1_a_foo_7 4)))) (or (or test1_a_zoo_5 test1_b_buzz_5) (or (<= test1_b_bar_4 3) (not (= test1_a_foo_8 4)))) (or (or test1_a_zoo_5 test1_b_buzz_5) (or (<= test1_b_bar_4 3) (not (= test1_a_foo_9 4)))) (or (or test1_a_zoo_5 test1_b_buzz_5) (or (<= test1_b_bar_4 3) (not (= test1_a_foo_10 4)))) (or (or test1_a_zoo_5 test1_b_buzz_5) (or (<= test1_b_bar_4 3) (not (= test1_a_foo_11 4)))) (or (or test1_a_zoo_5 test1_b_buzz_5) (or (<= test1_b_bar_4 3) (not (= test1_a_foo_12 4))))))`

	smt, err := prepTestSys("", test, false)

	if err != nil {
		t.Fatalf("compilation failed on valid spec. got=%s", err)
	}

	err = compareResults("TemporalSys", smt, string(expecting))

	if err != nil {
		t.Fatalf(err.Error())
	}
}
