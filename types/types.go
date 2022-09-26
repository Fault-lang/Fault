package types

import (
	"fault/ast"
	"fault/util"
	"fmt"
	"math"
	"strings"
)

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
	scope       string
	SpecStructs map[string]StockFlow
	Constants   map[string]map[string]ast.Node
	trail       importTrail
}

func (c *Checker) Check(a *ast.Spec) error {
	c.SpecStructs = make(map[string]StockFlow)
	c.Constants = make(map[string]map[string]ast.Node)

	// Break down the AST into constants and structs
	err := c.pass1(a)

	if err != nil {
		return err
	}

	// Pass two add types
	for k, v := range c.SpecStructs {
		for k2, v2 := range v {
			c.scope = k2
			for k3, v3 := range v2 {
				c.SpecStructs[k][k2][k3], err = c.pass2(v3, c.SpecStructs[k][k2])
				if err != nil {
					return err
				}
			}
			c.scope = ""
		}
	}
	return err
}

func (c *Checker) pass1(n ast.Node) error {
	var err error
	switch node := n.(type) {
	case *ast.Spec:
		for _, v := range node.Statements {
			err = c.pass1(v)
		}
		return err
	case *ast.SpecDeclStatement:
		c.SpecStructs[node.Name.Value] = StockFlow{}
		c.Constants[node.Name.Value] = make(map[string]ast.Node)
		c.trail = c.trail.PushSpec(node.Name.Value)
		return nil
	case *ast.SysDeclStatement:
		c.SpecStructs[node.Name.Value] = StockFlow{}
		c.Constants[node.Name.Value] = make(map[string]ast.Node)
		c.trail = c.trail.PushSpec(node.Name.Value)
		return nil
	case *ast.ImportStatement:
		err = c.pass1(node.Tree)
		_, c.trail = c.trail.PopSpec()
		return err
	case *ast.ConstantStatement:
		var n ast.Node
		if c.isValue(node.Value) {
			n, err = c.infer(node.Value, make(map[string]ast.Node))
		} else {
			n, err = c.inferFunction(node.Value, make(map[string]ast.Node))
		}
		c.Constants[c.trail.CurrentSpec()][node.Name.Value] = n
		return err
	case *ast.DefStatement:
		c.scope = strings.TrimSpace(node.Name.String())
		err = c.pass1(node.Value)
		return err
	case *ast.StockLiteral:
		node.InferredType = &ast.Type{Type: "STOCK", Scope: 0, Parameters: nil}
		structs := c.SpecStructs[c.trail.CurrentSpec()]
		nodes := util.Preparse(node.Pairs)
		c.SpecStructs[c.trail.CurrentSpec()] = structs.Bulk(c.scope, nodes)
		c.SpecStructs[c.trail.CurrentSpec()][c.scope]["___base"] = node //Don't lose the typing for the stocks and flows themselves
		c.scope = ""
		return err
	case *ast.FlowLiteral:
		node.InferredType = &ast.Type{Type: "FLOW",
			Scope:      0,
			Parameters: nil}
		structs := c.SpecStructs[c.trail.CurrentSpec()]
		nodes := util.Preparse(node.Pairs)
		c.SpecStructs[c.trail.CurrentSpec()] = structs.Bulk(c.scope, nodes)
		c.SpecStructs[c.trail.CurrentSpec()][c.scope]["___base"] = node
		c.scope = ""
		return err
	case *ast.ComponentLiteral:
		node.InferredType = &ast.Type{Type: "COMPONENT",
			Scope:      0,
			Parameters: nil}
		structs := c.SpecStructs[c.trail.CurrentSpec()]
		nodes := util.Preparse(node.Pairs)
		c.SpecStructs[c.trail.CurrentSpec()] = structs.Bulk(c.scope, nodes)
		c.SpecStructs[c.trail.CurrentSpec()][c.scope]["___base"] = node
		c.scope = ""
		return err
	case *ast.AssertionStatement:
		n, err := c.inferFunction(node.Constraints, make(map[string]ast.Node))
		valtype := typeable(n)
		if err != nil {
			return err
		}
		if valtype.Type != "BOOL" {
			return fmt.Errorf("assert statement not testing a Boolean expression. got=%s", valtype.Type)
		}
		return err

	case *ast.AssumptionStatement:
		n, err := c.inferFunction(node.Constraints, make(map[string]ast.Node))
		valtype := typeable(n)
		if err != nil {
			return err
		}
		if valtype.Type != "BOOL" {
			return fmt.Errorf("assume statement not testing a Boolean expression. got=%s", valtype.Type)
		}
		return err
	case *ast.ForStatement:
		return err
	case *ast.StartStatement:
		return err
	default:
		return fmt.Errorf("unimplemented: %T", node)
	}
}

func (c *Checker) pass2(n ast.Node, properties map[string]ast.Node) (ast.Node, error) {
	var err error
	if c.isValue(n) {
		n, err = c.infer(n, properties)
		return n, err
	} else {
		switch node := n.(type) {
		case *ast.BlockStatement:
			var valtype *ast.Type
			for i := 0; i < len(node.Statements); i++ {
				exp := node.Statements[i].(*ast.ExpressionStatement).Expression
				typedNode, err := c.inferFunction(exp, properties)
				if err != nil {
					return nil, err
				}
				node.Statements[i].(*ast.ExpressionStatement).Expression = typedNode
				valtype = typeable(typedNode)
				node.Statements[i].(*ast.ExpressionStatement).InferredType = valtype
			}
			node.InferredType = valtype
			return node, err
		default:
			n, err = c.inferFunction(n.(ast.Expression), properties)
			return n, err
		}
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
	case *ast.Unknown:
		return true
	case *ast.Nil:
		return true
	case *ast.StockLiteral:
		return true
	case *ast.FlowLiteral:
		return true
	case *ast.ComponentLiteral:
		return true
	default:
		return false
	}
}

func (c *Checker) infer(exp interface{}, p map[string]ast.Node) (ast.Node, error) {
	switch node := exp.(type) {
	/*case int64:
		return &ast.Type{"INT", 1, nil}, nil
	case float64:
		scope := c.inferScope(node)
		return &ast.Type{"FLOAT", scope, nil}, nil
	case string:
		return &ast.Type{"STRING", 0, nil}, nil
	case bool:
		return &ast.Type{"BOOL", 0, nil}, nil*/
	case *ast.IntegerLiteral:
		if node.InferredType == nil {
			node.InferredType = &ast.Type{Type: "INT",
				Scope:      1,
				Parameters: nil}
		}
		return node, nil
	case *ast.Boolean:
		if node.InferredType == nil {
			node.InferredType = &ast.Type{Type: "BOOL",
				Scope:      0,
				Parameters: nil}
		}
		return node, nil
	case *ast.FloatLiteral:
		if node.InferredType == nil {
			scope := c.inferScope(node.Value)
			node.InferredType = &ast.Type{Type: "FLOAT", Scope: scope, Parameters: nil}
		}
		return node, nil
	case *ast.StringLiteral:
		if node.InferredType == nil {
			node.InferredType = &ast.Type{Type: "STRING", Scope: 0, Parameters: nil}
		}
		return node, nil
	case *ast.Natural:
		if node.InferredType == nil {
			node.InferredType = &ast.Type{Type: "NATURAL", Scope: 1, Parameters: nil}
		}
		return node, nil
	case *ast.Uncertain:
		if node.InferredType == nil {
			params := c.inferUncertain(node)
			node.InferredType = &ast.Type{Type: "UNCERTAIN", Scope: 0, Parameters: params}
		}
		return node, nil

	case *ast.Unknown:
		if node.InferredType == nil {
			node.InferredType = &ast.Type{Type: "UNKNOWN", Scope: 0, Parameters: nil}
		}
		return node, nil
	case *ast.Identifier:
		if node.InferredType == nil {
			t, err := c.lookupType(node, p)
			node.InferredType = t
			if err != nil {
				return nil, err
			}
		}
		return node, nil
	case *ast.Instance:
		if node.InferredType == nil {
			t, err := c.lookupType(node, p)
			node.InferredType = t
			if err != nil {
				return nil, err
			}
		}
		return node, nil
	case *ast.ParameterCall:
		if node.InferredType == nil {
			t, err := c.lookupType(node, p)
			node.InferredType = t
			if err != nil {
				return nil, err
			}
		}
		return node, nil
	case *ast.Nil:
		if node.InferredType == nil {
			node.InferredType = &ast.Type{Type: "NIL", Scope: 0, Parameters: nil}
		}
		return node, nil
	case *ast.StockLiteral:
		if node.InferredType == nil {
			node.InferredType = &ast.Type{Type: "STOCK", Scope: 0, Parameters: nil}
		}
		return node, nil
	case *ast.FlowLiteral:
		if node.InferredType == nil {
			node.InferredType = &ast.Type{Type: "FLOW", Scope: 0, Parameters: nil}
		}
		return node, nil
	case *ast.ComponentLiteral:
		if node.InferredType == nil {
			node.InferredType = &ast.Type{Type: "COMPONENT", Scope: 0, Parameters: nil}
		}
		return node, nil
	default:
		pos := node.(ast.Node).Position()
		return nil, fmt.Errorf("unrecognized type: line %d col %d got=%T", pos[0], pos[1], node)
	}
}

func (c *Checker) lookupType(node ast.Node, p map[string]ast.Node) (*ast.Type, error) {
	if t := typeable(node); t != nil {
		return t, nil
	}
	//Prepare ID
	var id []string
	switch n := node.(type) {
	case *ast.Identifier:
		id = append(id, n.Value)
	case *ast.Instance:
		return &ast.Type{Type: "STOCK",
			Scope:      0,
			Parameters: nil}, nil
	case *ast.ParameterCall:
		id = n.Value
	}

	//Is this imported from another spec?
	if spec, ok := c.SpecStructs[id[0]]; ok {
		if len(id) == 3 {
			n, err := c.pass2(spec[id[1]][id[2]], p)
			c.SpecStructs[id[0]][id[1]][id[2]] = n
			return typeable(n), err
		}
	}

	// Check local preparse
	if s, ok := p[id[0]]; ok {
		switch ty := s.(type) {
		case *ast.Instance:
			structIdent := ty.Value.Value
			if len(id) > 2 {
				base := c.SpecStructs[ty.Value.Spec][structIdent][id[1]]
				n, err := c.complexInstances(base, id, p)
				return typeable(n), err
			} else {
				if len(id) == 1 {
					pos := ty.Position()
					return nil, fmt.Errorf("struct %s missing property, line:%d, col:%d", id[0], pos[0], pos[1])
				}
				n, err := c.pass2(c.SpecStructs[ty.Value.Spec][structIdent][id[1]], p)
				c.SpecStructs[ty.Value.Spec][structIdent][id[1]] = n
				return typeable(n), err
			}
		default:
			n, err := c.pass2(s, p)
			s = n
			return typeable(n), err
		}

	}
	// Check global variables
	if len(id) == 1 {
		//Assume current spec
		if v, ok := c.Constants[c.trail.CurrentSpec()][id[0]]; ok {
			return typeable(v), nil
		}
	} else if len(id) > 1 {
		if v, ok := c.Constants[id[0]][id[1]]; ok {
			return typeable(v), nil
		}
	}
	return nil, nil
}

func (c *Checker) complexInstances(base ast.Node, id []string, p map[string]ast.Node) (ast.Node, error) {
	b, ok := base.(*ast.Instance)
	if !ok {
		n, err := c.pass2(base, p)
		base = n
		return base, err
	}
	switch len(id) {
	case 1: //Do nothing, keep id as is
	case 2:
		b.Complex = false
		id = id[1:]
	default:
		b.Complex = true
		id = id[2:]
	}
	return c.complexInstances(c.SpecStructs[b.Value.Spec][b.Value.Value][id[0]], id, p)
}

func (c *Checker) inferFunction(f ast.Expression, p map[string]ast.Node) (ast.Expression, error) {
	var err error
	switch node := f.(type) {
	case *ast.FunctionLiteral:
		body := node.Body.Statements
		if len(body) == 1 && c.isValue(body[0].(*ast.ExpressionStatement).Expression) {
			typedNode, err := c.infer(body[0].(*ast.ExpressionStatement).Expression, p)
			tn, ok := typedNode.(ast.Expression)
			if !ok {
				pos := typedNode.Position()
				return nil, fmt.Errorf("node %T not an valid expression line: %d, col: %d", typedNode, pos[0], pos[1])
			}
			node.Body.Statements[0].(*ast.ExpressionStatement).Expression = tn
			return node, err
		}

		for i := 0; i < len(body); i++ {
			node.Body.Statements[i].(*ast.ExpressionStatement).Expression, err = c.inferFunction(body[i].(*ast.ExpressionStatement).Expression, p)
		}
		return node, err

	case *ast.StateLiteral:
		body := node.Body.Statements
		if len(body) == 1 && c.isValue(body[0].(*ast.ExpressionStatement).Expression) {
			typedNode, err := c.infer(body[0].(*ast.ExpressionStatement).Expression, p)
			tn, ok := typedNode.(ast.Expression)
			if !ok {
				pos := typedNode.Position()
				return nil, fmt.Errorf("node %T not an valid expression line: %d, col: %d", typedNode, pos[0], pos[1])
			}
			node.Body.Statements[0].(*ast.ExpressionStatement).Expression = tn
			return node, err
		}

		for i := 0; i < len(body); i++ {
			node.Body.Statements[i].(*ast.ExpressionStatement).Expression, err = c.inferFunction(body[i].(*ast.ExpressionStatement).Expression, p)
		}
		return node, err

	case *ast.BuiltIn:
		return node, err

	case *ast.Instance:
		node.InferredType = &ast.Type{Type: "STOCK",
			Scope:      0,
			Parameters: nil}
		return node, err

	case *ast.InfixExpression:
		if COMPARE[node.Operator] {
			node.InferredType = &ast.Type{Type: "BOOL",
				Scope:      0,
				Parameters: nil}
			return node, err
		}

		if node.Operator == "<-" {
			_, ok := c.SpecStructs[c.trail.CurrentSpec()][c.scope]["___base"].(*ast.StockLiteral)
			if ok {
				return nil, fmt.Errorf("stock is the store of values please use a flow for %s.%s", c.trail.CurrentSpec(), c.scope)
			}
		}

		var nl, nr ast.Node
		if c.isValue(node.Left) {
			nl, err = c.infer(node.Left, p)
		} else {
			nl, err = c.inferFunction(node.Left, p)
		}

		if err != nil {
			return nil, err
		}
		left := typeable(nl)

		if c.isValue(node.Right) {
			nr, err = c.infer(node.Right, p)

		} else {
			nr, err = c.inferFunction(node.Right, p)
		}
		if err != nil {
			return nil, err
		}
		right := typeable(nr)
		node.Left = nl.(ast.Expression)
		node.Right = nr.(ast.Expression)

		// If either value is Nil, return other type.
		// If both are Nil, return Nil.
		// Nils are kind of vestigial in Fault, unclear when
		// they will ever actually come up in the context of SMT
		_, nilL := node.Left.(*ast.Nil)
		_, nilR := node.Right.(*ast.Nil)
		if nilL && nilR {
			node.InferredType = &ast.Type{Type: "NIL",
				Scope:      0,
				Parameters: nil}
			return node, err
		} else if nilL {
			node.InferredType = right
			return node, err

		} else if nilR {
			node.InferredType = left
			return node, err
		}

		ty, err := typeAdju(left, right, node.Operator)
		if err != nil {
			return nil, err
		}
		node.InferredType = ty
		return node, err

	case *ast.InvariantClause:
		var nl, nr ast.Node
		if c.isValue(node.Left) {
			nl, err = c.infer(node.Left, p)
		} else {
			nl, err = c.inferFunction(node.Left, p)
		}
		if err != nil {
			return nil, err
		}
		left := typeable(nl)

		if c.isValue(node.Right) {
			nr, err = c.infer(node.Right, p)

		} else {
			nr, err = c.inferFunction(node.Right, p)
		}
		if err != nil {
			return nil, err
		}
		right := typeable(nr)

		node.Left = nl.(ast.Expression)
		node.Right = nr.(ast.Expression)

		if COMPARE[node.Operator] {
			node.InferredType = &ast.Type{Type: "BOOL",
				Scope:      0,
				Parameters: nil}
			return node, err
		}

		_, nilL := node.Left.(*ast.Nil)
		_, nilR := node.Right.(*ast.Nil)
		if nilL && nilR {
			node.InferredType = &ast.Type{Type: "NIL",
				Scope:      0,
				Parameters: nil}
			return node, err
		} else if nilL {
			node.InferredType = right
			return node, err

		} else if nilR {
			node.InferredType = left
			return node, err
		}
		ty, err := typeAdju(left, right, node.Operator)
		if err != nil {
			return nil, err
		}
		node.InferredType = ty
		return node, err

	case *ast.IfExpression:
		var ncond ast.Node
		var typedNode ast.Node
		var valtype *ast.Type
		if c.isValue(node.Condition) {
			ncond, err = c.infer(node.Condition, p)
		} else {
			ncond, err = c.inferFunction(node.Condition, p)
		}
		node.Condition = ncond.(ast.Expression)

		for i := 0; i < len(node.Consequence.Statements); i++ {
			exp := node.Consequence.Statements[i].(*ast.ExpressionStatement).Expression
			if c.isValue(exp) {
				typedNode, err = c.infer(exp, p)
			} else {
				typedNode, err = c.inferFunction(exp, p)
			}
			if err != nil {
				return nil, err
			}
			node.Consequence.Statements[i].(*ast.ExpressionStatement).Expression = typedNode.(ast.Expression)
			valtype = typeable(typedNode)
		}
		node.Consequence.InferredType = valtype

		if node.Alternative != nil {
			for i := 0; i < len(node.Alternative.Statements); i++ {
				exp := node.Alternative.Statements[i].(*ast.ExpressionStatement).Expression
				if c.isValue(exp) {
					typedNode, err = c.infer(exp, p)
				} else {
					typedNode, err = c.inferFunction(exp, p)
				}
				if err != nil {
					return nil, err
				}
				node.Alternative.Statements[i].(*ast.ExpressionStatement).Expression = typedNode.(ast.Expression)
				valtype = typeable(typedNode)
			}
			node.Alternative.InferredType = valtype
		}

		if node.Elif != nil {
			typedNode, err = c.inferFunction(node.Elif, p)
			if err != nil {
				return nil, err
			}
			node.Elif = typedNode.(*ast.IfExpression)
		}
		node.InferredType = node.Consequence.InferredType // This is probably an incorrect approach. Need to think about it.
		return node, err

	case *ast.PrefixExpression:
		var nr ast.Node
		if c.isValue(node.Right) {
			nr, err = c.infer(node.Right, p)
			node.Right = nr.(ast.Expression)

		} else {
			node.Right, err = c.inferFunction(node.Right, p)
		}

		if COMPARE[node.Operator] {
			node.InferredType = &ast.Type{Type: "BOOL",
				Scope:      0,
				Parameters: nil}
			return node, err
		}

		node.InferredType = typeable(node.Right)
		return node, err
	default:
		pos := node.(ast.Node).Position()
		return nil, fmt.Errorf("unrecognized type: line %d col %d got=%T", pos[0], pos[1], node)
	}
}

func (c *Checker) inferScope(fl float64) int64 {
	s := strings.Split(fmt.Sprintf("%f", fl), ".")
	base := c.calculateBase(s[1])
	return int64(base)
}

func (c *Checker) inferUncertain(node *ast.Uncertain) []ast.Type {
	return []ast.Type{
		{Type: "MEAN", Scope: c.inferScope(node.Mean), Parameters: nil},
		{Type: "SIGMA", Scope: c.inferScope(node.Sigma), Parameters: nil},
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

func typeAdju(left *ast.Type, right *ast.Type, op string) (*ast.Type, error) {
	if op == "=" && left == nil { //Allow variables local to functions
		return right, nil
	} else if op == "=" && right.Type == left.Type { // Allow functions to change variables
		return right, nil
	} else if op == "=" && right.Type != left.Type { // ...as long as they're the same type
		return nil, fmt.Errorf("cannot assign value of type %s to variable declared type %s", left.Type, right.Type)
	}

	if !COMPARE[op] && (left.Type == "BOOL" || right.Type == "BOOL") {
		return nil, fmt.Errorf("invalid expression: got=%s %s %s", left.Type, op, right.Type)
	}
	if left != right {
		if ast.TYPES[left.Type] == 0 || ast.TYPES[right.Type] == 0 {
			return nil, fmt.Errorf("type mismatch: got=%s,%s", left.Type, right.Type)
		}
		if ast.TYPES[left.Type] > ast.TYPES[right.Type] {
			return right, nil
		}
	}
	return left, nil
}

func typeable(node ast.Node) *ast.Type {
	switch n := node.(type) {
	case *ast.ConstantStatement:
		return n.InferredType
	case *ast.Identifier:
		return n.InferredType
	case *ast.ParameterCall:
		return n.InferredType
	case *ast.Instance:
		return n.InferredType
	case *ast.ExpressionStatement:
		return n.InferredType
	case *ast.IntegerLiteral:
		return n.InferredType
	case *ast.FloatLiteral:
		return n.InferredType
	case *ast.Natural:
		return n.InferredType
	case *ast.Uncertain:
		return n.InferredType
	case *ast.Unknown:
		return n.InferredType
	case *ast.PrefixExpression:
		return n.InferredType
	case *ast.InfixExpression:
		return n.InferredType
	case *ast.Boolean:
		return n.InferredType
	case *ast.This:
		return n.InferredType
	case *ast.Clock:
		return n.InferredType
	case *ast.Nil:
		return n.InferredType
	case *ast.BlockStatement:
		return n.InferredType
	case *ast.ParallelFunctions:
		return n.InferredType
	case *ast.InitExpression:
		return n.InferredType
	case *ast.IfExpression:
		return n.InferredType
	case *ast.StringLiteral:
		return n.InferredType
	case *ast.IndexExpression:
		return n.InferredType
	case *ast.StockLiteral:
		return n.InferredType
	case *ast.FlowLiteral:
		return n.InferredType
	case *ast.InvariantClause:
		return n.InferredType
	default:
		return nil
	}
}
