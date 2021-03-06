package types

import (
	"fault/ast"
	"fmt"
	"strings"
)

var TYPES = map[string]int{ //Convertible Types
	"STRING":    0, //Not convertible
	"BOOL":      1,
	"UNCERTAIN": 2,
	"NATURAL":   3,
	"FLOAT":     4,
	"INT":       5,
}

var COMPARE = map[string]bool{
	">":  true,
	"<":  true,
	"==": true,
	"!=": true,
	"<=": true,
	">=": true,
	"&&": true,
	"||": true,
}

type Checker struct {
	Symbols map[string]interface{}
	AST     *ast.Spec
	scope   string
}

func (c *Checker) Check(a *ast.Spec) error {
	c.Symbols = make(map[string]interface{})
	spec, err := c.assigntype(a)
	c.AST = spec.(*ast.Spec)
	return err
}

func (c *Checker) assigntype(exp interface{}) (interface{}, error) {
	var err error
	switch node := exp.(type) {
	case *ast.Spec:
		for k, v := range node.Statements {
			var statement interface{}
			statement, err = c.assigntype(v)
			node.Statements[k] = statement.(ast.Statement)
		}
		return node, err

	case *ast.SpecDeclStatement:
		return node, nil

	case *ast.ConstantStatement:
		id := node.Name.String()
		var valtype string
		if c.isValue(node.Value) {
			valtype, err = c.infer(node.Value, make(map[string]ast.Expression))
		} else {
			var n interface{}
			n, valtype, err = c.inferFunction(node.Value, make(map[string]ast.Expression))
			node.Value = n.(ast.Expression)
		}
		c.Symbols[id] = valtype
		return node, err

	case *ast.DefStatement:
		c.scope = node.Name.String()
		value, err := c.assigntype(node.Value)
		node.Value = value.(ast.Expression)
		return node, err

	case *ast.StockLiteral:
		newcontext := make(map[string]string)
		newcontext["__type"] = "STOCK"
		c.Symbols[c.scope] = newcontext
		properties := c.preparse(node.Pairs)
		for k, v := range node.Pairs {
			id := k.String()
			var valtype string
			if c.isValue(v) {
				valtype, err = c.infer(v, properties)
			} else {
				var n interface{}
				n, valtype, err = c.inferFunction(v, properties)
				node.Pairs[k] = n.(ast.Expression)
			}
			c.Symbols[c.scope].(map[string]string)[id] = valtype
		}
		c.scope = ""
		return node, err

	case *ast.FlowLiteral:
		newcontext := make(map[string]string)
		newcontext["__type"] = "FLOW"
		c.Symbols[c.scope] = newcontext
		properties := c.preparse(node.Pairs)
		for k, v := range node.Pairs {
			id := k.String()
			var valtype string
			if c.isValue(v) {
				valtype, err = c.infer(v, properties)
			} else {
				var n interface{}
				n, valtype, err = c.inferFunction(v, properties)
				node.Pairs[k] = n.(ast.Expression)
			}
			c.Symbols[c.scope].(map[string]string)[id] = valtype
		}
		c.scope = ""
		return node, err

	case *ast.AssertionStatement:
		var valtype string
		if c.isValue(node.Expression) {
			valtype, err = c.infer(node.Expression, make(map[string]ast.Expression))
		} else {
			_, valtype, err = c.inferFunction(node.Expression, make(map[string]ast.Expression))
		}

		if valtype != "BOOLEAN" {
			return node, fmt.Errorf("Assert statement not testing a Boolean expression. got=%s", valtype)
		}

		return node, err

	default:
		return node, fmt.Errorf("Unimplemented: %T", node)
	}
	return exp, nil
}

func (c *Checker) isValue(exp interface{}) bool {
	switch exp.(type) {
	case *ast.IntegerLiteral:
		return true
	case *ast.Boolean:
		return true
	case *ast.FloatLiteral:
		return true
	case *ast.StringLiteral:
		return true
	case *ast.Identifier:
		return true
	default:
		return false
	}
}

func (c *Checker) preparse(pairs map[ast.Expression]ast.Expression) map[string]ast.Expression {
	properties := make(map[string]ast.Expression)
	for k, v := range pairs {
		id := k.String()
		switch tree := v.(type) {
		case *ast.FunctionLiteral:
			properties[id] = c.preparseWalk(tree)
		case *ast.InstanceExpression:
			properties[id] = tree.Stock.(*ast.Identifier)
		}
	}
	return properties
}

func (c *Checker) preparseWalk(tree *ast.FunctionLiteral) ast.Expression {
	if len(tree.Body.Statements) == 1 {
		return tree.Body.Statements[0].(*ast.ExpressionStatement).Expression
	}
	return nil
}

func (c *Checker) infer(exp interface{}, p map[string]ast.Expression) (string, error) {
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
		id := strings.Split(node.Value, ".")

		if s, ok := c.Symbols[id[0]]; ok {
			if ty, ok := s.(string); ok {

				return ty, nil
			}
			return s.(map[string]string)[id[1]], nil
		}
		stock := p[id[0]].String()
		if s, ok := c.Symbols[stock]; ok {
			if ty, ok := s.(string); ok {

				return ty, nil
			}
			return s.(map[string]string)[id[1]], nil
		}

		pos := node.Position()
		return "", fmt.Errorf("Unrecognized type: line %d col %d got=%T", pos[0], pos[1], node)
	default:
		pos := node.(ast.Node).Position()
		return "", fmt.Errorf("Unrecognized type: line %d col %d got=%T", pos[0], pos[1], node)
	}
}

func (c *Checker) inferFunction(f ast.Expression, p map[string]ast.Expression) (ast.Expression, string, error) {
	var err error
	switch node := f.(type) {
	case *ast.FunctionLiteral:
		var valtype string
		body := node.Body.Statements
		if len(body) == 1 && c.isValue(body[0].(*ast.ExpressionStatement).Expression) {
			valtype, err = c.infer(body[0].(*ast.ExpressionStatement).Expression, p)
			return node, valtype, err
		}

		var n ast.Expression
		for i := 0; i < len(body); i++ {
			n, valtype, err = c.inferFunction(body[i].(*ast.ExpressionStatement).Expression, p)
		}
		return n, valtype, err

	case *ast.InstanceExpression:
		return node, "STOCK", nil

	case *ast.InfixExpression:
		if COMPARE[node.Operator] {
			return node, "BOOLEAN", err
		}

		var left, right string
		if c.isValue(node.Left) {
			left, err = c.infer(node.Left, p)
		} else {
			var n ast.Expression
			n, left, err = c.inferFunction(node.Left, p)
			node.Left = n.(ast.Expression)
		}

		if c.isValue(node.Right) {
			right, err = c.infer(node.Right, p)

		} else {
			var n ast.Expression
			n, right, err = c.inferFunction(node.Right, p)
			node.Right = n.(ast.Expression)
		}

		if left != right {
			if TYPES[left] == 0 || TYPES[right] == 0 {
				return node, "", fmt.Errorf("type mismatch: got=%s,%s", left, right)
			}
			if TYPES[left] > TYPES[right] {
				return node, right, err
			} else {
				return node, left, err
			}
		}
		return node, left, err
	}
	return f, "", nil
}

// func (c *Checker) convert(n ast.Expression, newType string) ast.Expression {
// 	//Needs to handle complex expressions too
// 	if !c.isValue(n) {
// 		switch node := n.(type) {
// 		case *ast.InfixExpression:
// 			node.Left = c.convert(node.Left, newType)
// 			node.Right = c.convert(node.Right, newType)
// 			return node
// 		case *ast.PrefixExpression:
// 			node.Right = c.convert(node.Right, newType)
// 			return node
// 		}
// 	}

// 	switch newType {
// 	case "FLOAT":
// 		node, ok := n.(*ast.IntegerLiteral)
// 		if ok {
// 			return &ast.FloatLiteral{
// 				Token: node.Token,
// 				Value: float64(node.Value),
// 			}
// 		} else { //Otherwise Identifier
// 			c.Symbols[n.(*ast.Identifier).Value] = "FLOAT"

// 		}
// 	}

// 	return n
//}
