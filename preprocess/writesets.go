package preprocess

import (
	"fault/ast"
	"strings"
)

// WriteSets maps fully-qualified function ID (e.g. "myspec_myflow_fn1") to the
// set of base variable names (SSA-index stripped) that the function writes to.
// Computed from the post-swap, post-alias-resolution AST — call after swaps.Swap.
type WriteSets map[string]map[string]bool

// ComputeWriteSets walks the resolved AST and returns a WriteSets for every
// FunctionLiteral found. Safe to call after swaps.Swap has resolved all aliases.
func ComputeWriteSets(tree *ast.Spec) WriteSets {
	ws := make(WriteSets)
	if tree == nil {
		return ws
	}
	for _, stmt := range tree.Statements {
		collectFunctions(stmt, ws)
	}
	return ws
}

// collectFunctions recurses into nodes that can contain FunctionLiterals.
func collectFunctions(node ast.Node, ws WriteSets) {
	switch n := node.(type) {
	case *ast.DefStatement:
		collectFunctions(n.Value, ws)
	case *ast.FlowLiteral:
		for _, v := range n.Pairs {
			collectFunctions(v, ws)
		}
	case *ast.ComponentLiteral:
		for _, v := range n.Pairs {
			collectFunctions(v, ws)
		}
	case *ast.FunctionLiteral:
		id := n.IdString()
		ws[id] = collectWrites(n.Body)
	}
}

// collectWrites returns the set of base variable names written within a function body.
// Recurses into nested blocks and if/else branches.
func collectWrites(body *ast.BlockStatement) map[string]bool {
	writes := make(map[string]bool)
	if body == nil {
		return writes
	}
	for _, stmt := range body.Statements {
		collectWritesFromStatement(stmt, writes)
	}
	return writes
}

func collectWritesFromStatement(stmt ast.Statement, writes map[string]bool) {
	switch s := stmt.(type) {
	case *ast.ExpressionStatement:
		collectWritesFromExpr(s.Expression, writes)
	}
}

func collectWritesFromExpr(expr ast.Expression, writes map[string]bool) {
	if expr == nil {
		return
	}
	switch e := expr.(type) {
	case *ast.InfixExpression:
		if e.Token.Type == "ASSIGN" {
			if nameable, ok := e.Left.(ast.Nameable); ok {
				writes[baseVarName(nameable.RawId())] = true
			}
		}
		// Recurse into both sides to catch nested assignments
		collectWritesFromExpr(e.Left, writes)
		collectWritesFromExpr(e.Right, writes)
	case *ast.IfExpression:
		collectWritesFromIfExpr(e, writes)
	}
}

func collectWritesFromIfExpr(e *ast.IfExpression, writes map[string]bool) {
	if e.Consequence != nil {
		for _, stmt := range e.Consequence.Statements {
			collectWritesFromStatement(stmt, writes)
		}
	}
	if e.Alternative != nil {
		for _, stmt := range e.Alternative.Statements {
			collectWritesFromStatement(stmt, writes)
		}
	}
	if e.Elif != nil {
		collectWritesFromIfExpr(e.Elif, writes)
	}
}

// baseVarName strips the SSA index suffix from a variable's raw ID parts and
// returns a single underscore-joined string.
// e.g. ["myspec", "myflow", "x"] -> "myspec_myflow_x"
func baseVarName(parts []string) string {
	return strings.Join(parts, "_")
}
