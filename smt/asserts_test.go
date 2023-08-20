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
	(assert (or (<= test1_t_foo_value_0 0) (<= test1_t_foo_value_1 0)(<= test1_t_foo_value_2 0)(<= test1_t_foo_value_3 0)(<= test1_t_foo_value_4 0)(<= test1_t_foo_value_5 0)))
`

	g := prepTest("", test, true, false)

	if len(g.Log.Asserts) != 6 {
		t.Fatalf("wrong number of asserts in the event log")
	}

	if g.Log.Asserts[0].Right.Type() != "FLOAT" || g.Log.Asserts[0].Right.GetFloat() != 0 {
		t.Fatalf("wrong right value in the first assert")
	}

	err := compareResults("SimpleAssert", g.SMT(), expecting)

	if err != nil {
		t.Fatalf(err.Error())
	}
}

func TestCompoundAssertAnd(t *testing.T) {
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
	(assert (or (<= test1_t_foo_value_0 0)
	(<= test1_t_foo_value_1 0)
	(<= test1_t_foo_value_2 0)
	(<= test1_t_foo_value_3 0)
	(<= test1_t_foo_value_4 0)
	(<= test1_t_foo_value_5 0)
	(> test1_t_foo_value_0 10)
	(> test1_t_foo_value_1 10)
	(> test1_t_foo_value_2 10)
	(> test1_t_foo_value_3 10)
	(> test1_t_foo_value_4 10)
	(> test1_t_foo_value_5 10)))`

	g := prepTest("", test, true, false)

	err := compareResults("CompoundAssert", g.SMT(), expecting)

	if err != nil {
		t.Fatalf(err.Error())
	}
}

func TestCompoundAssertOr(t *testing.T) {
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

	assert amount.value > 0 || amount.value <= 10;

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
	(assert (or (and (<= test1_t_foo_value_5 0) (> test1_t_foo_value_5 10)) (and (<= test1_t_foo_value_0 0) (> test1_t_foo_value_0 10)) (and (<= test1_t_foo_value_1 0) (> test1_t_foo_value_1 10)) (and (<= test1_t_foo_value_2 0) (> test1_t_foo_value_2 10)) (and (<= test1_t_foo_value_3 0) (> test1_t_foo_value_3 10)) (and (<= test1_t_foo_value_4 0) (> test1_t_foo_value_4 10))))`

	g := prepTest("", test, true, false)

	err := compareResults("CompoundAssert", g.SMT(), expecting)

	if err != nil {
		t.Fatalf(err.Error())
	}
}

func TestMultiAssert(t *testing.T) {
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
	assert amount.value <= 10;

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
		(or 
			(<= test1_t_foo_value_0 0)
			(<= test1_t_foo_value_1 0)
			(<= test1_t_foo_value_2 0)
			(<= test1_t_foo_value_3 0)
			(<= test1_t_foo_value_4 0)
			(<= test1_t_foo_value_5 0)
		)
		(or 
			(> test1_t_foo_value_0 10)
			(> test1_t_foo_value_1 10)
			(> test1_t_foo_value_2 10)
			(> test1_t_foo_value_3 10)
			(> test1_t_foo_value_4 10)
			(> test1_t_foo_value_5 10)
		)
	))`

	g := prepTest("", test, true, false)

	err := compareResults("CompoundAssert", g.SMT(), expecting)

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
	(assert (or (<= test1_t_foo_value_0 0)
	(<= test1_t_foo_value_1 0)
	(<= test1_t_foo_value_2 0)
	(<= test1_t_foo_value_3 0)
	(<= test1_t_foo_value_4 0)
	(<= test1_t_foo_value_5 0)))`

	g := prepTest("", test, true, false)

	err := compareResults("AssertInfix", g.SMT(), expecting)

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

	for 5 init{t = new test;}run {
		t.bar;
	};
	`
	expecting := `(set-logic QF_NRA)
	(declare-fun test1_t_foo_value_0 () Real)
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

	g := prepTest("", test, true, false)

	err := compareResults("MultiVar", g.SMT(), expecting)

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
	(assert (and (> test1_t_foo_value_0 0) (> test1_t_foo_value_1 0)(> test1_t_foo_value_2 0)(> test1_t_foo_value_3 0)(> test1_t_foo_value_4 0)(> test1_t_foo_value_5 0)))
`

	g := prepTest("", test, true, false)

	err := compareResults("SimpleAssume", g.SMT(), expecting)

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
	(assert (> test1_t_foo_value_1 0))
`

	g := prepTest("", test, true, false)

	err := compareResults("SpecificStateAssume", g.SMT(), expecting)

	if err != nil {
		t.Fatalf(err.Error())
	}
}

func TestIndexes(t *testing.T) {
	test := `spec indexes;

	def foo = stock{
		a: 10,
	};
	
	def bar = flow{
		bash: new foo,
		fizz: func{
			bash.a <- 2;
		},
	};
	
	assert foo.a[1] == 8;
	
	for 2 init{
		gee = new bar;
	}run{
		gee.fizz;
	};`

	expecting := `(set-logic QF_NRA)
	(declare-fun indexes_gee_bash_a_0 () Real)
	(declare-fun indexes_gee_bash_a_1 () Real)
	(declare-fun indexes_gee_bash_a_2 () Real)
	(assert (= indexes_gee_bash_a_0 10.0))
	(assert (= indexes_gee_bash_a_1 (+ indexes_gee_bash_a_0 2.0)))
	(assert (= indexes_gee_bash_a_2 (+ indexes_gee_bash_a_1 2.0)))
	(assert (not (= indexes_gee_bash_a_1 8)))`

	g := prepTest("", test, true, false)

	err := compareResults("SpecificStateAssume", g.SMT(), expecting)

	if err != nil {
		t.Fatalf(err.Error())
	}
}
