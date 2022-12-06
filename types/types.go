package types

import (
	"fault/ast"
	"fault/preprocess"
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

type Checker struct {
	SpecStructs map[string]*preprocess.SpecRecord
	Constants   map[string]map[string]ast.Node
	inStock     string
	temps       map[string]*ast.Type
}

func NewTypeChecker(specs map[string]*preprocess.SpecRecord) *Checker {
	return &Checker{
		SpecStructs: specs,
		Constants:   make(map[string]map[string]ast.Node),
		temps:       make(map[string]*ast.Type),
	}
}

func (c *Checker) Check(a *ast.Spec) (*ast.Spec, error) {
	// Break down the AST into constants and structs
	n, err := c.typecheck(a)

	a = n.(*ast.Spec)

	if err != nil {
		return a, err
	}

	return a, err
}

func (c *Checker) typecheck(n ast.Node) (ast.Node, error) {
	if n == nil {
		panic("nil value")
	}
	var tnode ast.Node
	var err error
	switch node := n.(type) {
	case *ast.Spec:
		var st []ast.Statement
		for _, v := range node.Statements {
			tnode, err = c.typecheck(v)
			if err != nil {
				return node, err
			}
			st = append(st, tnode.(ast.Statement))
		}
		node.Statements = st
		return node, err
	case *ast.SpecDeclStatement:
		return node, err
	case *ast.SysDeclStatement:
		return node, err
	case *ast.ImportStatement:
		tnode, err = c.typecheck(node.Tree)
		node.Tree = tnode.(*ast.Spec)
		return node, err
	case *ast.ConstantStatement:
		var n ast.Node
		if c.isValue(node.Value) {
			n, err = c.infer(node.Value)
		} else {
			n, err = c.inferFunction(node.Value)
		}
		if err != nil {
			return node, err
		}
		node.Value = n.(ast.Expression)
		return node, err
	case *ast.Identifier:
		n, err := c.infer(node)
		if err != nil {
			return node, err
		}
		return n, err
	case *ast.DefStatement:
		tnode, err = c.typecheck(node.Value)
		if err != nil {
			return node, err
		}
		node.Value = tnode.(ast.Expression)
		return node, err
	case *ast.StockLiteral:
		c.inStock = node.IdString()
		rawid := node.RawId()
		spec := c.SpecStructs[rawid[0]]
		node.InferredType = &ast.Type{Type: "STOCK", Scope: 0, Parameters: nil}
		for _, key := range node.Order {
			propid := node.GetPropertyIdent(key)
			v := node.Pairs[propid]
			tnode, err = c.typecheck(v)
			if err != nil {
				return node, err
			}
			node.Pairs[propid] = tnode.(ast.Expression)
			name := append(rawid, key)
			spec.UpdateVar(name, "STOCK", tnode)
		}
		c.inStock = ""
		return node, err
	case *ast.FlowLiteral:
		node.InferredType = &ast.Type{Type: "FLOW",
			Scope:      0,
			Parameters: nil}
		rawid := node.RawId()
		spec := c.SpecStructs[rawid[0]]
		for _, key := range node.Order {
			propid := node.GetPropertyIdent(key)
			v := node.Pairs[propid]
			tnode, err = c.typecheck(v)
			if err != nil {
				return node, err
			}
			node.Pairs[propid] = tnode.(ast.Expression)
			name := append(rawid, key)
			spec.UpdateVar(name, "FLOW", tnode)
		}

		return node, err
	case *ast.ComponentLiteral:
		node.InferredType = &ast.Type{Type: "COMPONENT",
			Scope:      0,
			Parameters: nil}
		rawid := node.RawId()
		spec := c.SpecStructs[rawid[0]]
		for _, key := range node.Order {
			propid := node.GetPropertyIdent(key)
			v := node.Pairs[propid]
			tnode, err = c.typecheck(v)
			if err != nil {
				return node, err
			}
			node.Pairs[propid] = tnode.(ast.Expression)
			name := append(rawid, key)
			spec.UpdateVar(name, "COMPONENT", tnode)
		}

		return node, err
	case *ast.AssertionStatement:
		n, err := c.inferFunction(node.Constraints)
		valtype := typeable(n)
		if err != nil {
			return node, err
		}
		if valtype.Type != "BOOL" {
			return nil, fmt.Errorf("assert statement not testing a Boolean expression. got=%s", valtype.Type)
		}
		return node, err

	case *ast.AssumptionStatement:
		n, err := c.inferFunction(node.Constraints)
		valtype := typeable(n)
		if err != nil {
			return node, err
		}
		if valtype.Type != "BOOL" {
			return nil, fmt.Errorf("assume statement not testing a Boolean expression. got=%s", valtype.Type)
		}
		return node, err
	case *ast.ForStatement:
		return node, err
	case *ast.StartStatement:
		return node, err
	case *ast.Instance:
		return node, err
	case *ast.StructInstance:
		cnode, err := c.complexInstances(node)
		node = cnode
		return node, err
	case *ast.FunctionLiteral:
		oldTemps := c.temps
		node2, err := c.typecheck(node.Body)
		if err != nil {
			return nil, err
		}
		node.Body = node2.(*ast.BlockStatement)
		c.temps = oldTemps
		return node, err
	case *ast.BlockStatement:
		var valtype *ast.Type
		for i := 0; i < len(node.Statements); i++ {
			exp := node.Statements[i].(*ast.ExpressionStatement).Expression
			typedNode, err := c.inferFunction(exp)
			if err != nil {
				return nil, err

			}
			node.Statements[i].(*ast.ExpressionStatement).Expression = typedNode
			valtype = typeable(typedNode)
			node.Statements[i].(*ast.ExpressionStatement).InferredType = valtype
		}
		node.InferredType = valtype
		return node, err
	case *ast.IntegerLiteral:
		return c.infer(node)
	case *ast.FloatLiteral:
		return c.infer(node)
	case *ast.Boolean:
		return c.infer(node)
	case *ast.StringLiteral:
		return c.infer(node)
	case *ast.ParameterCall:
		return c.infer(node)
	case *ast.ExpressionStatement:
		tnode, err = c.typecheck(node.Expression)
		if err != nil {
			return node, err
		}
		node.Expression = tnode.(ast.Expression)
		return node, err
	case *ast.Natural:
		return c.infer(node)
	case *ast.Uncertain:
		return c.infer(node)
	case *ast.Unknown:
		return c.infer(node)
	case *ast.PrefixExpression:
		return c.inferFunction(node)
	case *ast.InfixExpression:
		return c.inferFunction(node)
	case *ast.This:
		return c.infer(node)
	case *ast.Clock:
		return node, err
	case *ast.Nil:
		return c.infer(node)
	case *ast.ParallelFunctions:
		return node, err
	case *ast.InitExpression:
		return node, err
	case *ast.IfExpression:
		return c.inferFunction(node)
	case *ast.IndexExpression:
		tnode, err = c.typecheck(node.Left)
		if err != nil {
			return node, err
		}
		node.Left = tnode.(ast.Expression)
		return node, err
	case *ast.InvariantClause:
		return c.inferFunction(node)
	default:
		return node, fmt.Errorf("unimplemented: %s type %T", node, node)
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

func (c *Checker) infer(exp interface{}) (ast.Node, error) {
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
			t, err := c.lookupType(node)
			if err != nil {
				return nil, err
			}
			node.InferredType = t
		}
		return node, nil
	case *ast.Instance:
		if node.InferredType == nil {
			t, err := c.lookupType(node)
			node.InferredType = t
			if err != nil {
				return nil, err
			}
		}
		return node, nil
	case *ast.StructInstance:
		if node.InferredType == nil {
			t, err := c.lookupType(node)
			node.InferredType = t
			if err != nil {
				return nil, err
			}
		}
		return node, nil
	case *ast.ParameterCall:
		if node.InferredType == nil {
			t, err := c.lookupType(node)
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
	case *ast.IndexExpression:
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

func (c *Checker) lookupType(node ast.Node) (*ast.Type, error) {
	var err error
	pos := node.Position()

	if t := typeable(node); t != nil {
		return t, nil
	}

	if t, ok := c.temps[node.String()]; ok {
		return t, nil
	}
	
	//Prepare ID
	var rawid []string
	switch n := node.(type) {
	case *ast.Identifier:
		rawid = n.RawId()
	case *ast.StructInstance:
		return &ast.Type{Type: n.Type(),
			Scope:      0,
			Parameters: nil}, nil
	case *ast.Instance:
		return &ast.Type{Type: n.Type(),
			Scope:      0,
			Parameters: nil}, nil
	case *ast.ParameterCall:
		return c.lookupCallType(n)
	}

	spec := c.SpecStructs[rawid[0]]
	ty, _ := spec.GetStructType(rawid)
	v, err := spec.FetchVar(rawid, ty)
	if err != nil {
		return nil, fmt.Errorf("can't find node %s line:%d, col:%d", rawid, pos[0], pos[1])
	}
	return typeable(v), err
}

func (c *Checker) inferFunction(f ast.Expression) (ast.Expression, error) {
	var err error
	switch node := f.(type) {
	case *ast.FunctionLiteral:
		body := node.Body.Statements
		if len(body) == 1 && c.isValue(body[0].(*ast.ExpressionStatement).Expression) {
			typedNode, err := c.infer(body[0].(*ast.ExpressionStatement).Expression)
			tn, ok := typedNode.(ast.Expression)
			if !ok {
				pos := typedNode.Position()
				return nil, fmt.Errorf("node %T not an valid expression line: %d, col: %d", typedNode, pos[0], pos[1])
			}
			node.Body.Statements[0].(*ast.ExpressionStatement).Expression = tn
			return node, err
		}

		for i := 0; i < len(body); i++ {
			node.Body.Statements[i].(*ast.ExpressionStatement).Expression, err = c.inferFunction(body[i].(*ast.ExpressionStatement).Expression)
		}
		return node, err
	case *ast.BuiltIn:
		return node, err

	case *ast.Instance:
		node.InferredType = &ast.Type{Type: node.Type(),
			Scope:      0,
			Parameters: nil}
		return node, err
	case *ast.StructInstance:
		node.InferredType = &ast.Type{Type: node.Type(),
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
			if c.inStock != "" {
				return nil, fmt.Errorf("stock is the store of values, stock %s should be a flow", c.inStock)
			}
		}

		var nl, nr ast.Node
		if c.isValue(node.Right) {
			nr, err = c.infer(node.Right)

		} else {
			nr, err = c.inferFunction(node.Right)
		}
		if err != nil {
			return nil, err
		}
		right := typeable(nr)

		if node.Token.Type == "ASSIGN" { //In case of temp values
			ty, _ := c.lookupType(node.Left)
			if ty != nil && !isConvertible(ty, right) {
				return node, fmt.Errorf("cannot redeclare variable %s is type %s got %s", node.Left.String(), ty.Type, right.Type)
			}
			node.InferredType = right
			node.Left.SetType(right)
			c.temps[node.Left.String()] = right
			return node, nil
		}

		if c.isValue(node.Left) {
			nl, err = c.infer(node.Left)
		} else {
			nl, err = c.inferFunction(node.Left)
		}

		if err != nil {
			return nil, err
		}
		left := typeable(nl)

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
			nl, err = c.infer(node.Left)
		} else {
			nl, err = c.inferFunction(node.Left)
		}
		if err != nil {
			return nil, err
		}
		left := typeable(nl)

		if c.isValue(node.Right) {
			nr, err = c.infer(node.Right)

		} else {
			nr, err = c.inferFunction(node.Right)
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
			ncond, err = c.infer(node.Condition)
		} else {
			ncond, err = c.inferFunction(node.Condition)
		}
		node.Condition = ncond.(ast.Expression)

		for i := 0; i < len(node.Consequence.Statements); i++ {
			exp := node.Consequence.Statements[i].(*ast.ExpressionStatement).Expression
			if c.isValue(exp) {
				typedNode, err = c.infer(exp)
			} else {
				typedNode, err = c.inferFunction(exp)
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
					typedNode, err = c.infer(exp)
				} else {
					typedNode, err = c.inferFunction(exp)
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
			typedNode, err = c.inferFunction(node.Elif)
			if err != nil {
				return nil, err
			}
			node.Elif = typedNode.(*ast.IfExpression)
		}
		node.InferredType = node.Consequence.InferredType // This is probably an incorrect approach. Need to think about it.
		return node, err

	case *ast.IndexExpression:
		return node, err
	case *ast.PrefixExpression:
		var nr ast.Node
		if c.isValue(node.Right) {
			nr, err = c.infer(node.Right)
			if err != nil {
				return node, err
			}
			node.Right = nr.(ast.Expression)

		} else {
			node.Right, err = c.inferFunction(node.Right)
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

func (c *Checker) lookupCallType(base ast.Node) (*ast.Type, error) {
	var err error
	switch b := base.(type) {
	case *ast.ParameterCall:
		rawid := b.RawId()
		spec := c.SpecStructs[rawid[0]]
		ty, _ := spec.GetStructType(rawid[0 : len(rawid)-1])
		p, err := spec.FetchVar(rawid, ty)
		if err != nil {
			return nil, err
		}
		return c.lookupCallType(p)
	default:
		var n ast.Node
		if c.isValue(base) {
			n, err = c.infer(base)
		} else {
			n, err = c.inferFunction(base.(ast.Expression))
		}
		return typeable(n), err
	}
}

func (c *Checker) complexInstances(base *ast.StructInstance) (*ast.StructInstance, error) {
	var err error
	var rawid []string
	ret := make(map[string]*ast.StructProperty)
	for k, v := range base.Properties {
		b, ok := v.Value.(*ast.StructInstance)
		if !ok {
			cnode, err := c.typecheck(v.Value)
			if err != nil {
				return nil, err
			}
			v.Value = cnode
			ret[k] = v
			continue
		}

		rawid = b.RawId()

		cnode, err := c.complexInstances(b)
		props := util.ExtractBranches(cnode.Properties)
		if err != nil {
			return nil, err
		}

		spec := c.SpecStructs[rawid[0]]
		spec.Update(rawid, props)

		b = cnode
		v.Value = b
		ret[k] = v
	}
	base.Properties = ret
	return base, err
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
	} else if left == nil || right == nil {
		return nil, fmt.Errorf("improperly formatted infix")
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

func isConvertible(t1 *ast.Type, t2 *ast.Type) bool {
	if t1.Type == t2.Type {
		return true
	}
	if isNumeric(t1) && isNumeric(t2) {
		return true
	}
	return false
}

func isNumeric(t *ast.Type) bool {
	switch t.Type {
	case "INT":
		return true
	case "FLOAT":
		return true
	}
	return false
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
	case *ast.StructInstance:
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
