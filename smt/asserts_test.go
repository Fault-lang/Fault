package smt

import (
	"testing"
)

func TestSimpleAssert(t *testing.T) {
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

	assert amount.value > 0;

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
	(assert (or (<= test1_t_foo_value_0 0) (<= test1_t_foo_value_1 0)(<= test1_t_foo_value_2 0)(<= test1_t_foo_value_3 0)(<= test1_t_foo_value_4 0)(<= test1_t_foo_value_5 0)))
`

	smt, err := prepTest(test)

	if err != nil {
		t.Fatalf("compilation failed on valid spec. got=%s", err)
	}

	err = compareResults("SimpleAssert", smt, expecting)

	if err != nil {
		t.Fatalf(err.Error())
	}
}

func TestCompoundAssert(t *testing.T) {
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

	assert amount.value > 0 && amount.value <= 10;

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
	(assert (or (<= test1_t_foo_value_0 0)
	(<= test1_t_foo_value_1 0)
	(<= test1_t_foo_value_2 0)
	(<= test1_t_foo_value_3 0)
	(<= test1_t_foo_value_4 0)
	(<= test1_t_foo_value_5 0)))
	(assert (or (> test1_t_foo_value_0 10)
	(> test1_t_foo_value_1 10)
	(> test1_t_foo_value_2 10)
	(> test1_t_foo_value_3 10)
	(> test1_t_foo_value_4 10)
	(> test1_t_foo_value_5 10)))`

	smt, err := prepTest(test)

	if err != nil {
		t.Fatalf("compilation failed on valid spec. got=%s", err)
	}

	err = compareResults("CompoundAssert", smt, expecting)

	if err != nil {
		t.Fatalf(err.Error())
	}
}

func TestAssertInfix(t *testing.T) {
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

	assert amount.value > (2+3-5);

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
	(assert (or (<= test1_t_foo_value_0 0)
	(<= test1_t_foo_value_1 0)
	(<= test1_t_foo_value_2 0)
	(<= test1_t_foo_value_3 0)
	(<= test1_t_foo_value_4 0)
	(<= test1_t_foo_value_5 0)))`

	smt, err := prepTest(test)

	if err != nil {
		t.Fatalf("compilation failed on valid spec. got=%s", err)
	}

	err = compareResults("AssertInfix", smt, expecting)

	if err != nil {
		t.Fatalf(err.Error())
	}
}

func TestMultiVar(t *testing.T) {
	test := `spec test1;
	
	def amount = stock{
		value: 10,
	};
	
	def test = flow{
		foo: new amount,
		fuzz: 5,
		bar: func{
			foo.value -> 2;
		},
	};

	assert amount.value > test.fuzz;

	for 5 run {
		t = new test;
		t.bar;
	};
	`
	expecting := `(declare-fun test1_t_foo_value_0 () Real)
	(declare-fun test1_t_fuzz_0 () Real)
	(declare-fun test1_t_foo_value_1 () Real)
	(declare-fun test1_t_foo_value_2 () Real)
	(declare-fun test1_t_foo_value_3 () Real)
	(declare-fun test1_t_foo_value_4 () Real)
	(declare-fun test1_t_foo_value_5 () Real)
	(assert (= test1_t_foo_value_0 10.0))
	(assert (= test1_t_fuzz_0 5.0))
	(assert (= test1_t_foo_value_1 (- test1_t_foo_value_0 2.0)))
	(assert (= test1_t_foo_value_2 (- test1_t_foo_value_1 2.0)))
	(assert (= test1_t_foo_value_3 (- test1_t_foo_value_2 2.0)))
	(assert (= test1_t_foo_value_4 (- test1_t_foo_value_3 2.0)))
	(assert (= test1_t_foo_value_5 (- test1_t_foo_value_4 2.0)))
	(assert( or (<= test1_t_foo_value_0 test1_t_fuzz_0)
	(<= test1_t_foo_value_1 test1_t_fuzz_0)
	(<= test1_t_foo_value_2 test1_t_fuzz_0)
	(<= test1_t_foo_value_3 test1_t_fuzz_0)
	(<= test1_t_foo_value_4 test1_t_fuzz_0)
	(<= test1_t_foo_value_5 test1_t_fuzz_0)))
	`

	smt, err := prepTest(test)

	if err != nil {
		t.Fatalf("compilation failed on valid spec. got=%s", err)
	}

	err = compareResults("MultiVar", smt, expecting)

	if err != nil {
		t.Fatalf(err.Error())
	}
}

func TestSimpleAssume(t *testing.T) {
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

	assume amount.value > 0;

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
	(assert (and (> test1_t_foo_value_0 0) (> test1_t_foo_value_1 0)(> test1_t_foo_value_2 0)(> test1_t_foo_value_3 0)(> test1_t_foo_value_4 0)(> test1_t_foo_value_5 0)))
`

	smt, err := prepTest(test)

	if err != nil {
		t.Fatalf("compilation failed on valid spec. got=%s", err)
	}

	err = compareResults("SimpleAssume", smt, expecting)

	if err != nil {
		t.Fatalf(err.Error())
	}
}

func TestSpecificStateAssume(t *testing.T) {
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

	assume amount.value[1] > 0;

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
	(assert (> test1_t_foo_value_1 0))
`

	smt, err := prepTest(test)

	if err != nil {
		t.Fatalf("compilation failed on valid spec. got=%s", err)
	}

	err = compareResults("SpecificStateAssume", smt, expecting)

	if err != nil {
		t.Fatalf(err.Error())
	}
}

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

	assume eventually amount.value > 0;

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
