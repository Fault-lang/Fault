package reachability

import (
	"fault/ast"
	"fmt"
	"strings"
)

type Tracer struct {
	graph     map[string]bool
	undefined map[string]bool
	last      string
}

func NewTracer() *Tracer {
	return &Tracer{
		graph:     make(map[string]bool),
		undefined: make(map[string]bool),
	}
}

func (t *Tracer) Scan(spec *ast.Spec) error {
	t.walk(spec)
	ch, missing := t.check()
	if !ch {
		return fmt.Errorf("system under specified, states %s are unreachable", strings.Join(missing, ", "))
	}
	return nil
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
				t.undefined[id] = true
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
		if node.Function != "advance" {
			return
		}
		for _, v := range node.Parameters {
			id := v.(ast.Nameable).Id()
			if _, ok := t.graph[id[1]]; !ok {
				t.undefined[id[1]] = true
			} else {
				t.graph[id[1]] = true
				t.removeUndefined(id[1])
			}
		}
	case *ast.InfixExpression:
		t.walk(node.Left)
		t.walk(node.Right)
	case *ast.PrefixExpression:
		t.walk(node.Right)
	}
}

func (t *Tracer) seenBefore(id string) bool {
	return t.undefined[id]
}

func (t *Tracer) removeUndefined(id string) {
	delete(t.undefined, id)
}

func (t *Tracer) check() (bool, []string) {
	for k, v := range t.graph {
		if !v {
			t.undefined[k] = true
		}
	}
	missing := make([]string, 0, len(t.undefined))
	for k := range t.undefined {
		missing = append(missing, k)
	}
	return len(missing) == 0, missing
}
