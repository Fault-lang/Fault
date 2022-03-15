package smt

import (
	"fmt"
	"os"
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

	smt, err := prepTest(test)

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

	smt, err := prepTest(test)

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

	smt, err := prepTest(test)

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

	smt, err := prepTest(test)

	if err != nil {
		t.Fatalf("compilation failed on valid spec. got=%s", err)
	}

	err = compareResults("Temporal", smt, expecting)

	if err != nil {
		t.Fatalf(err.Error())
	}
}

func TestTemporal2(t *testing.T) {
	// nft (and (or ))
	// nmt (or (and
	//(or
	//(and ))))

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
	expecting, err := os.ReadFile("testdata/temporal2.smt2")
	if err != nil {
		panic(fmt.Sprintf("compiled spec %s is not valid", "testdata/temporal2.smt2"))
	}

	smt, err := prepTest(test)

	if err != nil {
		t.Fatalf("compilation failed on valid spec. got=%s", err)
	}

	err = compareResults("Temporal2", smt, string(expecting))

	if err != nil {
		t.Fatalf(err.Error())
	}
}
