package preprocess

import (
	"fault/listener"
	"fmt"
	"testing"
)

func TestPreProbe(t *testing.T) {
	test := `spec test1;
def generic = stock{
	level: 10,
};
def child = stock{
	extends generic,
	extra: 5,
};
def f = flow{
	c: new child,
	fn: func{
		if c.extra > 5 {
			c.level -> 2;
		} else {
			c.extra <- 1;
		}
	},
};
run init{inst = new f;} {
	inst.fn;
};
`
	flags := map[string]bool{"specType": true, "testing": false, "skipRun": false}
	l, _ := listener.Execute(test, "", flags)
	pre, err := Execute(l)
	if err != nil {
		t.Fatalf("preprocess failed: %s", err)
	}
	spec := pre.Specs["test1"]
	fmt.Println("=== Stocks ===")
	for name, fields := range spec.Stocks {
		var keys []string
		for k := range fields {
			keys = append(keys, k)
		}
		fmt.Printf("  %s: %v\n", name, keys)
	}
	fmt.Println("=== Order ===")
	for _, v := range spec.Order {
		fmt.Printf("  %v\n", v)
	}
	t.Log("preprocess succeeded")
}
