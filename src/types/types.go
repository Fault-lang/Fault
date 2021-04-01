package types

import (
	"fault/ast"
	"fmt"
	"math"
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

type Type struct {
	Type       string
	Scope      int32
	Parameters []Type
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
			var valtype *Type
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
			newcontext := make(map[string]*Type)
			newcontext["__type"] = &Type{"STOCK", 0, nil}
			c.SymbolTypes[c.scope] = newcontext
		} else {
			properties := c.preparse(node.Pairs)
			for k, v := range node.Pairs {
				id := k.String()
				var valtype *Type
				if c.isValue(v) {
					valtype, err = c.infer(v, properties)
				} else {
					valtype, err = c.inferFunction(v, properties)
				}
				c.SymbolTypes[c.scope].(map[string]*Type)[id] = valtype
			}
		}
		c.scope = ""
		return err

	case *ast.FlowLiteral:
		if pass == 1 {
			newcontext := make(map[string]*Type)
			newcontext["__type"] = &Type{"FLOW", 0, nil}
			c.SymbolTypes[c.scope] = newcontext
		} else {
			properties := c.preparse(node.Pairs)
			for k, v := range node.Pairs {
				id := k.String()
				var valtype *Type
				if c.isValue(v) {
					valtype, err = c.infer(v, properties)
				} else {
					valtype, err = c.inferFunction(v, properties)
				}
				c.SymbolTypes[c.scope].(map[string]*Type)[id] = valtype
			}
		}
		c.scope = ""
		return err

	case *ast.AssertionStatement:
		if pass == 1 {
			var valtype *Type
			if c.isValue(node.Expression) {
				valtype, err = c.infer(node.Expression, make(map[string]ast.Expression))
			} else {
				valtype, err = c.inferFunction(node.Expression, make(map[string]ast.Expression))
			}

			if valtype.Type != "BOOL" {
				return fmt.Errorf("Assert statement not testing a Boolean expression. got=%s", valtype.Type)
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

func (c *Checker) infer(exp interface{}, p map[string]ast.Expression) (*Type, error) {
	switch node := exp.(type) {
	case *ast.IntegerLiteral:
		return &Type{"INT", 1, nil}, nil
	case *ast.Boolean:
		return &Type{"BOOL", 0, nil}, nil
	case *ast.FloatLiteral:
		scope := c.inferScope(node.Value)
		return &Type{"FLOAT", scope, nil}, nil
	case *ast.StringLiteral:
		return &Type{"STRING", 0, nil}, nil
	case *ast.Natural:
		return &Type{"NATURAL", 1, nil}, nil
	case *ast.Uncertain:
		params := c.inferUncertain(node)
		return &Type{"UNCERTAIN", 0, params}, nil
	case *ast.Identifier:
		id := strings.Split(node.Value, ".")

		if s, ok := c.SymbolTypes[id[0]]; ok {
			if ty, ok := s.(*Type); ok {

				return ty, nil
			}
			return s.(map[string]*Type)[id[1]], nil
		}
		stock := p[id[0]].String()
		if s, ok := c.SymbolTypes[stock]; ok {
			if ty, ok := s.(*Type); ok {

				return ty, nil
			}
			return s.(map[string]*Type)[id[1]], nil
		}

		pos := node.Position()
		return nil, fmt.Errorf("Unrecognized type: line %d col %d got=%T", pos[0], pos[1], node)
	default:
		pos := node.(ast.Node).Position()
		return nil, fmt.Errorf("Unrecognized type: line %d col %d got=%T", pos[0], pos[1], node)
	}
}

func (c *Checker) inferFunction(f ast.Expression, p map[string]ast.Expression) (*Type, error) {
	var err error
	switch node := f.(type) {
	case *ast.FunctionLiteral:
		var valtype *Type
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
		return &Type{"STOCK", 0, nil}, nil

	case *ast.InfixExpression:
		if COMPARE[node.Operator] {
			return &Type{"BOOL", 0, nil}, err
		}

		var left, right *Type
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
			if TYPES[left.Type] == 0 || TYPES[right.Type] == 0 {
				return nil, fmt.Errorf("type mismatch: got=%s,%s", left.Type, right.Type)
			}
			if TYPES[left.Type] > TYPES[right.Type] {
				return right, err
			} else {
				return left, err
			}
		}
		return left, err

	case *ast.PrefixExpression:
		if COMPARE[node.Operator] {
			return &Type{"BOOL", 0, nil}, err
		}
		var right *Type
		if c.isValue(node.Right) {
			right, err = c.infer(node.Right, p)

		} else {
			right, err = c.inferFunction(node.Right, p)
		}
		return right, err
	}
	return nil, nil
}

func (c *Checker) inferScope(fl float64) int32 {
	s := strings.Split(fmt.Sprintf("%f", fl), ".")
	base := c.calculateBase(s[1])
	return int32(base)
}

func (c *Checker) inferUncertain(node *ast.Uncertain) []Type {
	return []Type{
		{"MEAN", c.inferScope(node.Mean), nil},
		{"SIGMA", c.inferScope(node.Sigma), nil},
	}
}
func (c *Checker) calculateBase(s string) int32 {
	rns := []rune(s) // convert to rune
	zero := []rune("0")
	for i := len(rns) - 1; i >= 0; i = i - 1 {
		if rns[i] != zero[0] {
			base := math.Pow10(i + 1)
			return int32(base)
		}
	}
	return 1
}
