package reachability

import (
	"fault/ast"
	"fmt"
	"os"
	"strings"
)

type Tracer struct {
	graph     map[string]bool
	undefined []string
}

func NewTracer() *Tracer {
	return &Tracer{graph: make(map[string]bool)}
}

func (t *Tracer) Scan(spec *ast.Spec) {
	t.walk(spec)
	ch, missing := t.check()
	if !ch {
		fmt.Fprintf(os.Stderr, "error: system under specified, states %s are unreachable\n", missing)
		os.Exit(1)
	}
}

func (t *Tracer) walk(n ast.Node) {
	switch node := n.(type) {
	case *ast.Spec:
		for _, v := range node.Statements {
			t.walk(v)
		}
	case *ast.DefStatement:
		t.walk(node.Value)
	case *ast.ComponentLiteral:
		nid := node.Id()
		for k, v := range node.Pairs {
			if f, ok := v.(*ast.FunctionLiteral); ok {
				pid := k.Id()
				id := fmt.Sprintf("%s_%s", nid[1], pid[1])
				if _, ok := t.graph[id]; !ok {
					t.graph[id] = false
				}

				if t.seenBefore(id) {
					t.graph[id] = true
					t.removeUndefined(id)
				}

				t.walk(f)
			}
		}
	case *ast.StartStatement:
		for _, v := range node.Pairs {
			id := strings.Join(v, "_")
			if _, ok := t.graph[id]; !ok {
				t.undefined = append(t.undefined, id)
			} else {
				t.graph[id] = true
				t.removeUndefined(id)
			}
		}
	case *ast.FunctionLiteral:
		t.walk(node.Body)
	case *ast.BlockStatement:
		for _, v := range node.Statements {
			t.walk(v)
		}
	case *ast.ExpressionStatement:
		t.walk(node.Expression)
	case *ast.IfExpression:
		t.walk(node.Consequence)

		if node.Elif != nil {
			t.walk(node.Elif)
		}
		if node.Alternative != nil {
			t.walk(node.Alternative)
		}
	case *ast.BuiltIn:
		for _, v := range node.Parameters {
			id := v.(ast.Nameable).Id()
			if _, ok := t.graph[id[1]]; !ok {
				t.undefined = append(t.undefined, id[1])
			} else {
				t.graph[id[1]] = true
				t.removeUndefined(id[1])
			}
		}
	}
}

func (t *Tracer) seenBefore(id string) bool {
	for _, v := range t.undefined {
		if v == id {
			return true
		}
	}
	return false
}

func (t *Tracer) removeUndefined(id string) {
	var new []string
	for _, v := range t.undefined {
		if v != id {
			new = append(new, v)
		}
	}
	t.undefined = new
}

func (t *Tracer) check() (bool, []string) {
	for k, v := range t.graph {
		if !v {
			t.undefined = append(t.undefined, k)
		}
	}
	return len(t.undefined) == 0, t.undefined
}
