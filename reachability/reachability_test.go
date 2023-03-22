package reachability

import (
	"fault/listener"
	"fault/parser"
	"fault/preprocess"
	"fault/types"
	"testing"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
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
	path := ""
	is := antlr.NewInputStream(test)
	lexer := parser.NewFaultLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	p := parser.NewFaultParser(stream)
	l := listener.NewListener(path, true, false)
	antlr.ParseTreeWalkerDefault.Walk(l, p.SysSpec())
	pre := preprocess.NewProcesser()
	pre.StructsPropertyOrder = l.StructsPropertyOrder
	tree := pre.Run(l.AST)

	ty := types.NewTypeChecker(pre)
	tree, err := ty.Check(tree)
	if err != nil {
		panic(err)
	}
	tracer := NewTracer()
	tracer.walk(tree)
	return tracer.check()
}
