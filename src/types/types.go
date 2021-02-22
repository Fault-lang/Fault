package types

import (
	"fault/ast"
	"fmt"
)

type Checker struct {
	Symbols map[string]interface{}
	AST     *ast.Spec
	scope   string
}

func (c *Checker) Check(a *ast.Spec) error {
	c.AST = a
	c.Symbols = make(map[string]interface{})
	err := c.assigntype(c.Symbols, a)
	return err
}

func (c *Checker) assigntype(context map[string]interface{}, exp interface{}) error {
	switch node := exp.(type) {
	case *ast.Spec:
		for _, v := range node.Statements {
			err := c.assigntype(context, v)
			if err != nil {
				return err
			}
		}
	case *ast.SpecDeclStatement:
		return nil

	case *ast.ConstantStatement:
		id := node.Name.String()
		valtype, err := c.infer(node.Value)
		if err != nil {
			return err
		}
		context[id] = valtype
	case *ast.DefStatement:
		c.scope = node.Name.String()
		err := c.assigntype(context, node.Value)
		if err != nil {
			return err
		}

	case *ast.StockLiteral:
		newcontext := make(map[string]interface{})
		for k, v := range node.Pairs {
			id := k.String()
			valtype, err := c.infer(v)
			if err != nil {
				return err
			}
			newcontext[id] = valtype
		}
		context[c.scope] = newcontext
		c.scope = ""

	case *ast.FlowLiteral:
		newcontext := make(map[string]interface{})
		for k, v := range node.Pairs {
			id := k.String()
			valtype, err := c.infer(v)
			if err != nil {
				return err
			}
			newcontext[id] = valtype
		}
		context[c.scope] = newcontext
		c.scope = ""
	default:
		return fmt.Errorf("Unimplemented: %T", node)
	}
	c.Symbols = context
	return nil
}

func (c *Checker) infer(exp interface{}) (string, error) {
	switch node := exp.(type) {
	case *ast.IntegerLiteral:
		return "INT", nil
	case *ast.Boolean:
		return "BOOL", nil
	case *ast.FloatLiteral:
		return "FLOAT", nil
	case *ast.StringLiteral:
		return "STRING", nil
	case *ast.Identifier:
		return "REF", nil
	case *ast.InfixExpression:
		left, err := c.infer(node.Left)
		if err != nil {
			return "", err
		}
		right, err := c.infer(node.Right)
		if err != nil {
			return "", err
		}
		if left != right {
			//Union
			return "", fmt.Errorf("type mismatch: got=%T,%T", left, right)
		}
		return left, nil

	default:
		//pos := node.(ast.Node).Position()
		//return "", fmt.Errorf("Unrecognized type: line %d col %d", pos[0], pos[1])
		return "", nil
	}
}
