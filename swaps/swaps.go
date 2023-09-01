package swaps

import (
	"fault/ast"
	"fault/types"
	"fmt"

	"github.com/barkimedes/go-deepcopy"
)

type Precompiler struct {
	checker *types.Checker
	Alias   map[string]string
}

func NewPrecompiler(check *types.Checker) *Precompiler {
	return &Precompiler{
		checker: check,
		Alias:   make(map[string]string),
	}
}

func (c *Precompiler) Swap(n *ast.Spec) *ast.Spec {
	s := c.walk(n)
	return s.(*ast.Spec)
}

func (c *Precompiler) walk(n ast.Node) ast.Node {
	var err error
	switch node := n.(type) {
	case *ast.StructInstance:
		node, err = c.swapValues(node)
		if err != nil {
			panic(err)
		}
		return node
	case *ast.Spec:
		var st []ast.Statement
		for _, v := range node.Statements {
			snode := c.walk(v)
			st = append(st, snode.(ast.Statement))
		}
		node.Statements = st
		return node
	case *ast.SpecDeclStatement:
		return node
	case *ast.SysDeclStatement:
		return node
	case *ast.ImportStatement:
		snode := c.walk(node.Tree)
		node.Tree = snode.(*ast.Spec)
		return node
	case *ast.ConstantStatement:
		return node
	case *ast.Identifier:
		return node
	case *ast.DefStatement:
		return node
	case *ast.StockLiteral:
		return node
	case *ast.FlowLiteral:
		return node
	case *ast.ComponentLiteral:
		return node
	case *ast.AssertionStatement:
		return node
	case *ast.ForStatement:
		var st []ast.Statement
		for _, v := range node.Inits.Statements {
			snode := c.walk(v)
			st = append(st, snode.(ast.Statement))
		}
		node.Inits.Statements = st
		return node
	case *ast.StartStatement:
		return node
	case *ast.FunctionLiteral:
		return node
	case *ast.BlockStatement:
		if node == nil {
			return node
		}
		for i := 0; i < len(node.Statements); i++ {
			if e, ok := node.Statements[i].(*ast.ExpressionStatement); ok {
				snode := c.walk(e.Expression)
				node.Statements[i].(*ast.ExpressionStatement).Expression = snode.(ast.Expression)
			}
		}
		return node
	case *ast.BuiltIn:
		return node
	case *ast.IntegerLiteral:
		return node
	case *ast.FloatLiteral:
		return node
	case *ast.Boolean:
		return node
	case *ast.StringLiteral:
		return node
	case *ast.ParameterCall:
		return node
	case *ast.ExpressionStatement:
		snode := c.walk(node.Expression)
		node.Expression = snode.(ast.Expression)
		return node
	case *ast.Natural:
		return node
	case *ast.Uncertain:
		return node
	case *ast.Unknown:
		return node
	case *ast.PrefixExpression:
		return node
	case *ast.InfixExpression:
		return node
	case *ast.This:
		return node
	case *ast.Clock:
		return node
	case *ast.Nil:
		return node
	case *ast.ParallelFunctions:
		return node
	case *ast.InitExpression:
		return node
	case *ast.IfExpression:
		if node == nil {
			return node
		}
		//Not sure to allow this
		con := c.walk(node.Consequence)
		alt := c.walk(node.Alternative)
		elif := c.walk(node.Elif)
		node.Consequence = con.(*ast.BlockStatement)
		node.Alternative = alt.(*ast.BlockStatement)
		node.Elif = elif.(*ast.IfExpression)
		return node
	case *ast.IndexExpression:
		return node
	case *ast.InvariantClause:
		return node
	default:
		panic(fmt.Errorf("unimplemented: %s type %T", node, node))
	}
}

func (c *Precompiler) swapValues(base *ast.StructInstance) (*ast.StructInstance, error) {
	for _, s := range base.Swaps {
		infix := s.(*ast.InfixExpression)
		rawid := infix.Left.(ast.Nameable).RawId()
		key := rawid[len(rawid)-1]
		val, err := c.checker.Reference(infix.Right)
		if err != nil {
			return base, err
		}

		// Because part of what we're doing here is renaming
		// these nodes. We need to do a deep copy to separate
		// the swapped nodes from their original reference values
		copyVal, err := deepcopy.Anything(val)
		if err != nil {
			return base, err
		}

		val = copyVal.(ast.Node)

		switch v := val.(type) {
		case *ast.ParameterCall, *ast.Identifier:
			c.Alias[infix.Left.(ast.Nameable).IdString()] = infix.Right.(ast.Nameable).IdString()
		case *ast.StructInstance:
			for k, v2 := range v.Properties {
				aliasKey := fmt.Sprintf("%s_%s", infix.Left.(ast.Nameable).IdString(), k)
				c.Alias[aliasKey] = v2.IdString()
			}

			v.Name = key
			val = v
		}

		if len(val.(ast.Nameable).RawId()) == 0 {
			val.(ast.Nameable).SetId(rawid)
		}

		base.Properties[key].Value = val
		base = c.swapDeepNames(base)

	}
	return base, nil
}

func (c *Precompiler) swapDeepNames(val *ast.StructInstance) *ast.StructInstance {
	rawid := val.RawId()
	err := c.checker.SpecStructs[rawid[0]].Update(rawid, ast.ExtractBranches(val.Properties))
	if err != nil {
		panic(fmt.Sprintf("failed to update spec record on swap %s: %s", val.String(), err))
	}

	node, err := c.checker.Preprocesser.Partial(rawid[0], val)
	if err != nil {
		panic(fmt.Sprintf("failed to update process ids on swap %s: %s", val.String(), err))
	}
	return node.(*ast.StructInstance)
}
