package types

import (
	"fault/ast"
	"fault/walker"
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
	Scope      int64
	Parameters []Type
}

type StockFlow map[string]map[string]ast.Node

func (sf StockFlow) Add(strname string, prpname string, node ast.Node) StockFlow {
	if sf[strname] != nil {
		sf[strname][prpname] = node
	} else {
		sf[strname] = make(map[string]ast.Node)
		sf[strname][prpname] = node
	}
	return sf
}

func (sf StockFlow) Bulk(strname string, nodes map[string]ast.Node) StockFlow {
	if sf[strname] != nil {
		sf[strname] = nodes
	} else {
		sf[strname] = make(map[string]ast.Node)
		sf[strname] = nodes
	}
	return sf
}

func (sf StockFlow) Get(strname string, prpname string) ast.Node {
	if sf[strname] != nil && sf[strname][prpname] != nil {
		return sf[strname][prpname]
	}
	panic(fmt.Sprintf("No variable named %s.%s", strname, prpname))
}

func (sf StockFlow) GetStruct(strname string) map[string]ast.Node {
	if sf[strname] != nil {
		return sf[strname]
	}
	panic(fmt.Sprintf("No stock or flow named %s", strname))
}

type importTrail []string

func (i importTrail) CurrentSpec() string {
	if len(i) == 0 {
		panic(fmt.Sprintln("import trail is empty"))
	}
	return i[len(i)-1]
}

func (i importTrail) PushSpec(spec string) []string {
	i = append(i, spec)
	return i
}

func (i importTrail) PopSpec() (string, []string) {
	if len(i) == 0 {
		panic(fmt.Sprintln("import trail is empty"))
	}
	spec := i[len(i)-1]
	i = i[0 : len(i)-1]
	return spec, i
}

type Checker struct {
	SymbolTypes map[string]map[string]interface{}
	scope       string
	pass        int8
	SpecStructs map[string]StockFlow
	trail       importTrail
}

func (c *Checker) Check(a *ast.Spec) error {
	c.SymbolTypes = make(map[string]map[string]interface{})
	c.SpecStructs = make(map[string]StockFlow)

	// Pass one, globals and constants
	c.pass = 1
	err := c.assigntype(a)

	if err != nil {
		return err
	}

	// Pass two, stock/flow properties
	c.pass = 2
	err = c.assigntype(a)
	return err
}

func (c *Checker) assigntype(exp interface{}) error {
	var err error
	switch node := exp.(type) {
	case *ast.Spec:
		for _, v := range node.Statements {
			err = c.assigntype(v)
		}
		return err

	case *ast.SpecDeclStatement:
		if c.pass == 1 {
			c.SymbolTypes[node.Name.Value] = make(map[string]interface{})
			c.SpecStructs[node.Name.Value] = StockFlow{}
			c.trail = c.trail.PushSpec(node.Name.Value)
		}
		return nil

	case *ast.ImportStatement:
		im := c.assigntype(node.Tree)
		_, c.trail = c.trail.PopSpec()
		return im

	case *ast.ConstantStatement:
		if c.pass == 1 {
			id := node.Name.String()
			var valtype *Type
			if c.isValue(node.Value) {
				valtype, err = c.infer(node.Value, make(map[string]ast.Node))
			} else {
				valtype, err = c.inferFunction(node.Value, make(map[string]ast.Node))
			}
			c.SymbolTypes[c.trail.CurrentSpec()][id] = valtype
		}
		return err

	case *ast.DefStatement:
		c.scope = node.Name.String()
		err = c.assigntype(node.Value)
		return err

	case *ast.StockLiteral:
		if c.pass == 1 {
			newcontext := make(map[string]*Type)
			newcontext["__type"] = &Type{"STOCK", 0, nil}
			c.SymbolTypes[c.trail.CurrentSpec()][c.scope] = newcontext

			structs := c.SpecStructs[c.trail.CurrentSpec()]
			nodes := walker.Preparse(node.Pairs)
			c.SpecStructs[c.trail.CurrentSpec()] = structs.Bulk(c.scope, nodes)
		} else {
			properties := walker.Preparse(node.Pairs)
			for k, v := range node.Pairs {
				id := k.String()
				var valtype *Type
				if c.isValue(v) {
					valtype, err = c.infer(v, properties)
				} else {
					valtype, err = c.inferFunction(v, properties)
				}
				c.SymbolTypes[c.trail.CurrentSpec()][c.scope].(map[string]*Type)[id] = valtype
			}
		}
		c.scope = ""
		return err

	case *ast.FlowLiteral:
		if c.pass == 1 {
			newcontext := make(map[string]*Type)
			newcontext["__type"] = &Type{"FLOW", 0, nil}
			c.SymbolTypes[c.trail.CurrentSpec()][c.scope] = newcontext

			structs := c.SpecStructs[c.trail.CurrentSpec()]
			nodes := walker.Preparse(node.Pairs)
			c.SpecStructs[c.trail.CurrentSpec()] = structs.Bulk(c.scope, nodes)
			fmt.Print(c.SpecStructs[c.trail.CurrentSpec()])
		} else {
			properties := walker.Preparse(node.Pairs)
			for k, v := range node.Pairs {
				id := k.String()
				var valtype *Type
				if c.isValue(v) {
					valtype, err = c.infer(v, properties)
				} else {
					valtype, err = c.inferFunction(v, properties)
				}
				c.SymbolTypes[c.trail.CurrentSpec()][c.scope].(map[string]*Type)[id] = valtype
			}
		}
		c.scope = ""
		fmt.Print(c.SpecStructs[c.trail.CurrentSpec()])
		return err

	case *ast.AssertionStatement:
		if c.pass == 1 {
			var valtype *Type
			if c.isValue(node.Expression) {
				valtype, err = c.infer(node.Expression, make(map[string]ast.Node))
			} else {
				valtype, err = c.inferFunction(node.Expression, make(map[string]ast.Node))
			}

			if valtype.Type != "BOOL" {
				return fmt.Errorf("Assert statement not testing a Boolean expression. got=%s", valtype.Type)
			}
		}
		return err
	case *ast.ForStatement:
		return err
	default:
		return fmt.Errorf("Unimplemented: %T", node)
	}
}

func (c *Checker) isValue(exp interface{}) bool {
	switch exp.(type) {
	case int64:
		return true
	case float64:
		return true
	case string:
		return true
	case bool:
		return true
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
	case *ast.ParameterCall:
		return true
	case *ast.Natural:
		return true
	case *ast.Uncertain:
		return true
	default:
		return false
	}
}

func (c *Checker) infer(exp interface{}, p map[string]ast.Node) (*Type, error) {
	switch node := exp.(type) {
	case int64:
		return &Type{"INT", 1, nil}, nil
	case float64:
		scope := c.inferScope(node)
		return &Type{"FLOAT", scope, nil}, nil
	case string:
		return &Type{"STRING", 0, nil}, nil
	case bool:
		return &Type{"BOOL", 0, nil}, nil
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
		return c.lookupType(node, p)
	case *ast.Instance:
		return c.lookupType(node, p)
	case *ast.ParameterCall:
		return c.lookupType(node, p)

	default:
		pos := node.(ast.Node).Position()
		return nil, fmt.Errorf("unrecognized type: line %d col %d got=%T", pos[0], pos[1], node)
	}
}

func (c *Checker) lookupType(node ast.Node, p map[string]ast.Node) (*Type, error) {

	//Prepare ID
	var id []string
	switch n := node.(type) {
	case *ast.Identifier:
		id = append(id, n.Value)
	case *ast.Instance:
		return &Type{"STOCK", 0, nil}, nil
	case *ast.ParameterCall:
		id = n.Value
	}
	var structIdent string //

	// Check local vars
	if s, ok := c.SymbolTypes[c.trail.CurrentSpec()][c.scope]; ok {
		valtype := s.(map[string]*Type)[id[0]]
		if valtype != nil {
			return valtype, nil
		}
	}

	// Check local preparse
	if s, ok := p[id[0]]; ok {
		switch ty := s.(type) {
		case *ast.Instance:
			if len(id) > 1 { //Must be a parameter call
				structIdent = ty.Value.Value
			} else {
				return c.SymbolTypes[c.trail.CurrentSpec()][ty.Value.Value].(*Type), nil
			}
		case *ast.ParameterCall:
			structIdent = ty.Value[0]
		case *ast.FunctionLiteral:
			var ret *Type
			var err error
			body := ty.Body.Statements
			for i := 0; i < len(body); i++ {
				exp := body[i].(*ast.ExpressionStatement).Expression
				ret, err = c.inferFunction(exp, p)
				if err != nil {
					panic(err)
				}
			}
			return ret, err

		default:
			if c.isValue(p[id[0]]) {
				return c.infer(p[id[0]], p)
			}
			return c.inferFunction(p[id[0]].(ast.Expression), p)
		}
	}

	// Check global preparse
	currSpec := c.trail.CurrentSpec()
	if s, ok := c.SpecStructs[currSpec][structIdent]; ok {
		switch ty := s[id[1]].(type) {
		case *ast.Instance:
			return c.SymbolTypes[c.trail.CurrentSpec()][ty.Value.Value].(*Type), nil
		case *ast.FunctionLiteral:
			var ret *Type
			var err error
			body := ty.Body.Statements
			for i := 0; i < len(body); i++ {
				exp := body[i].(*ast.ExpressionStatement).Expression
				ret, err = c.inferFunction(exp, p)
				if err != nil {
					panic(err)
				}
			}
			return ret, err

		default:
			if c.isValue(s[id[1]]) {
				return c.infer(s[id[1]], p)
			}
			return c.inferFunction(s[id[1]].(ast.Expression), p)
		}

	}

	// Check global
	if s, ok := c.SymbolTypes[c.trail.CurrentSpec()][id[0]]; ok {
		switch ty := s.(type) {
		case *Type:
			return ty, nil
		case map[string]*Type:
			return ty[id[1]], nil
		}
	}

	return nil, nil
}

func (c *Checker) inferFunction(f ast.Expression, p map[string]ast.Node) (*Type, error) {
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

	case *ast.Instance:
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
		if err != nil {
			return left, err
		}

		if c.isValue(node.Right) {
			right, err = c.infer(node.Right, p)

		} else {
			right, err = c.inferFunction(node.Right, p)
		}

		if err != nil {
			return right, err
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

func (c *Checker) inferScope(fl float64) int64 {
	s := strings.Split(fmt.Sprintf("%f", fl), ".")
	base := c.calculateBase(s[1])
	return int64(base)
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
