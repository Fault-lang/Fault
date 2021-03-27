package types

import (
	"fault/ast"
	"fmt"
	"strings"
)

var TYPES = map[string]int{ //Convertible Types
	"STRING":    0, //Not convertible
	"BOOL":      1,
	"NATURAL":   2,
	"FLOAT":     3,
	"INT":       4,
	"UNCERTAIN": 5,
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
	"!":  true, //Prefix
}

type Checker struct {
	SymbolTypes map[string]interface{}
	scope       string
}

func (c *Checker) Check(a *ast.Spec) error {
	c.SymbolTypes = make(map[string]interface{})

	// Pass one, globals and constants
	err := c.assigntype(a, 1)

	if err != nil {
		return err
	}

	// Pass two, stock/flow properties
	err = c.assigntype(a, 2)
	return err
}

func (c *Checker) assigntype(exp interface{}, pass int) error {
	var err error
	switch node := exp.(type) {
	case *ast.Spec:
		for _, v := range node.Statements {
			err = c.assigntype(v, pass)
		}
		return err

	case *ast.SpecDeclStatement:
		return nil

	case *ast.ConstantStatement:
		if pass == 1 {
			id := node.Name.String()
			var valtype string
			if c.isValue(node.Value) {
				valtype, err = c.infer(node.Value, make(map[string]ast.Expression))
			} else {
				valtype, err = c.inferFunction(node.Value, make(map[string]ast.Expression))
			}
			c.SymbolTypes[id] = valtype
		}
		return err

	case *ast.DefStatement:
		c.scope = node.Name.String()
		err = c.assigntype(node.Value, pass)
		return err

	case *ast.StockLiteral:
		if pass == 1 {
			newcontext := make(map[string]string)
			newcontext["__type"] = "STOCK"
			c.SymbolTypes[c.scope] = newcontext
		} else {
			properties := c.preparse(node.Pairs)
			for k, v := range node.Pairs {
				id := k.String()
				var valtype string
				if c.isValue(v) {
					valtype, err = c.infer(v, properties)
				} else {
					valtype, err = c.inferFunction(v, properties)
				}
				c.SymbolTypes[c.scope].(map[string]string)[id] = valtype
			}
		}
		c.scope = ""
		return err

	case *ast.FlowLiteral:
		if pass == 1 {
			newcontext := make(map[string]string)
			newcontext["__type"] = "FLOW"
			c.SymbolTypes[c.scope] = newcontext
		} else {
			properties := c.preparse(node.Pairs)
			for k, v := range node.Pairs {
				id := k.String()
				var valtype string
				if c.isValue(v) {
					valtype, err = c.infer(v, properties)
				} else {
					valtype, err = c.inferFunction(v, properties)
				}
				c.SymbolTypes[c.scope].(map[string]string)[id] = valtype
			}
		}
		c.scope = ""
		return err

	case *ast.AssertionStatement:
		if pass == 1 {
			var valtype string
			if c.isValue(node.Expression) {
				valtype, err = c.infer(node.Expression, make(map[string]ast.Expression))
			} else {
				valtype, err = c.inferFunction(node.Expression, make(map[string]ast.Expression))
			}

			if valtype != "BOOL" {
				return fmt.Errorf("Assert statement not testing a Boolean expression. got=%s", valtype)
			}
		}
		return err

	default:
		return fmt.Errorf("Unimplemented: %T", node)
	}
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
	case *ast.Natural:
		return true
	case *ast.Uncertain:
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
	case *ast.Natural:
		return "NATURAL", nil
	case *ast.Uncertain:
		return "UNCERTAIN", nil
	case *ast.Identifier:
		id := strings.Split(node.Value, ".")

		if s, ok := c.SymbolTypes[id[0]]; ok {
			if ty, ok := s.(string); ok {

				return ty, nil
			}
			return s.(map[string]string)[id[1]], nil
		}
		stock := p[id[0]].String()
		if s, ok := c.SymbolTypes[stock]; ok {
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

func (c *Checker) inferFunction(f ast.Expression, p map[string]ast.Expression) (string, error) {
	var err error
	switch node := f.(type) {
	case *ast.FunctionLiteral:
		var valtype string
		body := node.Body.Statements
		if len(body) == 1 && c.isValue(body[0].(*ast.ExpressionStatement).Expression) {
			valtype, err = c.infer(body[0].(*ast.ExpressionStatement).Expression, p)
			return valtype, err
		}

		for i := 0; i < len(body); i++ {
			valtype, err = c.inferFunction(body[i].(*ast.ExpressionStatement).Expression, p)
		}
		return valtype, err

	case *ast.InstanceExpression:
		return "STOCK", nil

	case *ast.InfixExpression:
		if COMPARE[node.Operator] {
			return "BOOL", err
		}

		var left, right string
		if c.isValue(node.Left) {
			left, err = c.infer(node.Left, p)
		} else {
			left, err = c.inferFunction(node.Left, p)
		}

		if c.isValue(node.Right) {
			right, err = c.infer(node.Right, p)

		} else {
			right, err = c.inferFunction(node.Right, p)
		}

		if left != right {
			if TYPES[left] == 0 || TYPES[right] == 0 {
				return "", fmt.Errorf("type mismatch: got=%s,%s", left, right)
			}
			if TYPES[left] > TYPES[right] {
				return right, err
			} else {
				return left, err
			}
		}
		return left, err

	case *ast.PrefixExpression:
		if COMPARE[node.Operator] {
			return "BOOL", err
		}
		var right string
		if c.isValue(node.Right) {
			right, err = c.infer(node.Right, p)

		} else {
			right, err = c.inferFunction(node.Right, p)
		}
		return right, err
	}
	return "", nil
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
// 			c.SymbolTypes[n.(*ast.Identifier).Value] = "FLOAT"

// 		}
// 	}

// 	return n
//}
