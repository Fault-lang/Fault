package reachability

import (
	"fault/listener"
	"fault/preprocess"
	"fault/swaps"
	"fault/types"
	"testing"
)

func TestSeenBefore(t *testing.T) {
	tracer := NewTracer()
	tracer.undefined = []string{"test"}

	if !tracer.seenBefore("test") {
		t.Fatal("seenBefore function not working")
	}
}

func TestRemoveUndefined(t *testing.T) {
	tracer := NewTracer()
	tracer.undefined = []string{"test", "test2", "test3"}

	tracer.removeUndefined("test2")

	if len(tracer.undefined) != 2 || tracer.undefined[0] != "test" || tracer.undefined[1] != "test3" {
		t.Fatalf("removeUndefined function not working got=%s", tracer.undefined)
	}
}

func TestCorrect(t *testing.T) {
	test := `
	system test;

	component foo = states{
		x: 8,
		initial: func{
			if this.x > 10{
				stay();
			}else{
				advance(this.alarm);
			}
		},
		alarm: func{
			advance(this.close);
		},
		close: func{
			stay();
		},
	};

	start {
		foo: initial,
	};
	`
	check, missing := prepTestSys(test)

	if !check || len(missing) > 0 {
		t.Fatalf("reachability check failed on valid spec got=%s", missing)
	}
}

func TestIncorrect(t *testing.T) {
	test := `
	system test;

	component foo = states{
		initial: func{
			advance(this.alarm);
		},
		alarm: func{
			advance(this.close);
		},
		close: func{
			stay();
		},
		error: func{
			stay();
		},
	};

	start {
		foo: initial,
	};
	`
	check, missing := prepTestSys(test)

	if check {
		t.Fatal("reachability check failed to catch missing state error")
	}

	if len(missing) == 0 || missing[0] != "foo_error" {
		t.Fatalf("reachability check failed to catch missing state got=%s", missing)
	}
}

func TestMultiIncorrect(t *testing.T) {
	test := `
	system test;

	component foo = states{
		initial: func{
			advance(this.alarm);
		},
		alarm: func{
			advance(bar.error);
		},
	};

	component bar = states{
		error: func{
			advance(this.resolve);
		},
		resolve: func {
			advance(foo.initial);
		},
	};

	component fizz = states{
		empty: func{
			advance(bar.error);
		},
		active: func{
			advance(this.empty);
		},
	};

	start {
		foo: initial,
	};
	`
	check, missing := prepTestSys(test)

	if check {
		t.Fatal("reachability check failed to catch missing state error")
	}

	if len(missing) == 0 || missing[0] != "fizz_active" {
		t.Fatalf("reachability check failed to catch missing state got=%s", missing)
	}
}

func TestMultiCorrect(t *testing.T) {
	test := `
	system test;

	component foo = states{
		initial: func{
			advance(this.alarm);
		},
		alarm: func{
			advance(bar.error);
		},
	};

	component bar = states{
		error: func{
			advance(this.resolve);
		},
		resolve: func {
			advance(foo.initial);
		},
	};

	component fizz = states{
		empty: func{
			advance(bar.error);
		},
		active: func{
			advance(this.empty);
		},
	};

	start {
		foo: initial,
		fizz: active,
	};
	`
	check, missing := prepTestSys(test)

	if !check || len(missing) > 0 {
		t.Fatalf("reachability check failed on valid spec got=%s", missing)
	}
}

func TestMultiPath(t *testing.T) {
	test := `
	system test;

	component foo = states{
		initial: func{
			advance(bar.alarm);
		},
	};

	component bar = states{
		initial: func{
			advance(this.alarm);
		},
		alarm: func{
			advance(this.close);
		},
		close: func{
			stay();
		},
	};

	start {
		foo: initial,
		bar: initial,
	};
	`
	check, missing := prepTestSys(test)

	if !check || len(missing) > 0 {
		t.Fatalf("reachability check failed on valid spec got=%s", missing)
	}
}

func prepTestSys(test string) (bool, []string) {
	flags := make(map[string]bool)
	flags["specType"] = false
	flags["testing"] = true
	flags["skipRun"] = false

	var path string

	l := listener.Execute(test, path, flags)
	pre := preprocess.Execute(l)
	ty := types.Execute(pre.Processed, pre.Specs)
	tracer := NewTracer()
	tracer.walk(ty.Checked)
	return tracer.check()
}
