package reachability

import (
	"fault/listener"
	"fault/preprocess"
	"fault/types"
	"testing"
)

func TestSeenBefore(t *testing.T) {
	tracer := NewTracer()
	tracer.undefined = map[string]bool{"test": true}

	if !tracer.seenBefore("test") {
		t.Fatal("seenBefore function not working")
	}
}

func TestRemoveUndefined(t *testing.T) {
	tracer := NewTracer()
	tracer.undefined = map[string]bool{"test": true, "test2": true, "test3": true}

	tracer.removeUndefined("test2")

	if len(tracer.undefined) != 2 || tracer.undefined["test2"] {
		t.Fatalf("removeUndefined function not working got=%v", tracer.undefined)
	}
}

func TestCorrect(t *testing.T) {
	test := `
	system test;

	component foo = states{
		x: 8,
		initial: sfunc{
			if this.x > 10{
				stay();
			}else{
				advance(this.alarm);
			}
		},
		alarm: sfunc{
			advance(this.close);
		},
		close: sfunc{
			stay();
		},
	};

	run {
		foo.initial;
	}
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
		initial: sfunc{
			advance(this.alarm);
		},
		alarm: sfunc{
			advance(this.close);
		},
		close: sfunc{
			stay();
		},
		error: sfunc{
			stay();
		},
	};

	run {
		foo.initial;
	}
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
		initial: sfunc{
			advance(this.alarm);
		},
		alarm: sfunc{
			advance(bar.error);
		},
	};

	component bar = states{
		error: sfunc{
			advance(this.resolve);
		},
		resolve: sfunc {
			advance(foo.initial);
		},
	};

	component fizz = states{
		empty: sfunc{
			advance(bar.error);
		},
		active: sfunc{
			advance(this.empty);
		},
	};

	run {
		foo.initial;
	}
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
		initial: sfunc{
			advance(this.alarm);
		},
		alarm: sfunc{
			advance(bar.error);
		},
	};

	component bar = states{
		error: sfunc{
			advance(this.resolve);
		},
		resolve: sfunc {
			advance(foo.initial);
		},
	};

	component fizz = states{
		empty: sfunc{
			advance(bar.error);
		},
		active: sfunc{
			advance(this.empty);
		},
	};

	run {
		foo.initial && fizz.active;
	}
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
		initial: sfunc{
			advance(bar.alarm);
		},
	};

	component bar = states{
		initial: sfunc{
			advance(this.alarm);
		},
		alarm: sfunc{
			advance(this.close);
		},
		close: sfunc{
			stay();
		},
	};

	run {
		foo.initial && bar.initial;
	}
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

	l, _ := listener.Execute(test, path, flags)
	pre, err := preprocess.Execute(l)
	if err != nil {
		panic(err)
	}
	ty, err := types.Execute(pre.Processed, pre)
	if err != nil {
		panic(err)
	}
	tracer := NewTracer()
	tracer.walk(ty.Checked)
	return tracer.check()
}
