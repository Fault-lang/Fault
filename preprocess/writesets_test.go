package preprocess

import (
	"testing"
)

func prepWriteSets(test string) WriteSets {
	pre := prepTest(test, true)
	return ComputeWriteSets(pre.Processed)
}

func TestWriteSets_SimpleAssignment(t *testing.T) {
	test := `spec test1;
	def stock1 = stock{
		x: 0,
	};
	def foo = flow{
		s: new stock1,
		bar: func{
			s.x <- 5;
		},
	};`

	ws := prepWriteSets(test)

	fn, ok := ws["test1_foo_bar"]
	if !ok {
		t.Fatalf("expected write set for test1_foo_bar, got keys: %v", wsKeys(ws))
	}
	if !fn["test1_foo_s_x"] {
		t.Fatalf("expected test1_foo_s_x in write set, got: %v", fn)
	}
}

func TestWriteSets_MultipleWrites(t *testing.T) {
	test := `spec test1;
	def stock1 = stock{
		x: 0,
		y: 0,
	};
	def foo = flow{
		s: new stock1,
		bar: func{
			s.x <- 1;
			s.y <- 2;
		},
	};`

	ws := prepWriteSets(test)

	fn, ok := ws["test1_foo_bar"]
	if !ok {
		t.Fatalf("expected write set for test1_foo_bar, got keys: %v", wsKeys(ws))
	}
	if !fn["test1_foo_s_x"] {
		t.Fatalf("expected test1_foo_s_x in write set, got: %v", fn)
	}
	if !fn["test1_foo_s_y"] {
		t.Fatalf("expected test1_foo_s_y in write set, got: %v", fn)
	}
}

func TestWriteSets_DisjointFunctions(t *testing.T) {
	test := `spec test1;
	def stock1 = stock{
		x: 0,
		y: 0,
	};
	def foo = flow{
		s: new stock1,
		inc: func{
			s.x <- s.x + 1;
		},
		dec: func{
			s.y <- s.y - 1;
		},
	};`

	ws := prepWriteSets(test)

	inc, ok := ws["test1_foo_inc"]
	if !ok {
		t.Fatalf("expected write set for test1_foo_inc, got keys: %v", wsKeys(ws))
	}
	if !inc["test1_foo_s_x"] {
		t.Fatalf("expected test1_foo_s_x in inc write set, got: %v", inc)
	}
	if inc["test1_foo_s_y"] {
		t.Fatalf("test1_foo_s_y should not be in inc write set, got: %v", inc)
	}

	dec, ok := ws["test1_foo_dec"]
	if !ok {
		t.Fatalf("expected write set for test1_foo_dec, got keys: %v", wsKeys(ws))
	}
	if !dec["test1_foo_s_y"] {
		t.Fatalf("expected test1_foo_s_y in dec write set, got: %v", dec)
	}
	if dec["test1_foo_s_x"] {
		t.Fatalf("test1_foo_s_x should not be in dec write set, got: %v", dec)
	}
}

func TestWriteSets_BranchWrites(t *testing.T) {
	// Both branches write different variables — both should appear in the write set
	test := `spec test1;
	def stock1 = stock{
		x: 0,
		y: 0,
		z: 0,
	};
	def foo = flow{
		s: new stock1,
		toggle: func{
			if s.x > 0 {
				s.y <- s.y + 1;
			} else {
				s.z <- s.z - 1;
			}
		},
	};`

	ws := prepWriteSets(test)

	fn, ok := ws["test1_foo_toggle"]
	if !ok {
		t.Fatalf("expected write set for test1_foo_toggle, got keys: %v", wsKeys(ws))
	}
	if !fn["test1_foo_s_y"] {
		t.Fatalf("expected test1_foo_s_y in write set (true branch), got: %v", fn)
	}
	if !fn["test1_foo_s_z"] {
		t.Fatalf("expected test1_foo_s_z in write set (false branch), got: %v", fn)
	}
}

func TestWriteSets_ReadOnlyFunction(t *testing.T) {
	// A function with no assignments — write set should be empty
	test := `spec test1;
	def stock1 = stock{
		x: 0,
	};
	def foo = flow{
		s: new stock1,
		check: func{
			s.x + 1;
		},
	};`

	ws := prepWriteSets(test)

	fn, ok := ws["test1_foo_check"]
	if !ok {
		t.Fatalf("expected write set for test1_foo_check, got keys: %v", wsKeys(ws))
	}
	if len(fn) != 0 {
		t.Fatalf("expected empty write set for read-only function, got: %v", fn)
	}
}

func wsKeys(ws WriteSets) []string {
	ks := make([]string, 0, len(ws))
	for k := range ws {
		ks = append(ks, k)
	}
	return ks
}
