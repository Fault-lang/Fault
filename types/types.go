package types

import (
	"fault/ast"
	"fault/preprocess"
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
	SpecStructs  map[string]*preprocess.SpecRecord
	Instances    map[string]*ast.StructInstance
	inStock      string
	temps        map[string]*ast.Type
	Checked      *ast.Spec
	Preprocesser *preprocess.Processor
}

func NewTypeChecker(Processer *preprocess.Processor) *Checker {
	return &Checker{
		SpecStructs:  Processer.Specs,
		Instances:    Processer.Instances,
		temps:        make(map[string]*ast.Type),
		Preprocesser: Processer,
	}
}

func Execute(tree *ast.Spec, processor *preprocess.Processor) *Checker {
	ty := NewTypeChecker(processor)
	tree, err := ty.Check(tree)
	if err != nil {
		panic(err)
	}
	ty.Checked = tree
	return ty
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

func (c *Checker) Reference(n ast.Node) (ast.Node, error) {
	return c.lookupReference(n)
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
		n2, err := c.lookupReference(node)
		if err != nil {
			return node, err
		}

		tnode, err = c.typecheck(n2)
		if err != nil {
			return node, err
		}

		node.InferredType = &ast.Type{Type: tnode.Type(),
			Scope:      0,
			Parameters: nil}
		return node, err
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
		n, err := c.inferFunction(node.Constraint)
		valtype := typeable(n)
		if err != nil {
			return node, err
		}
		if valtype.Type != "BOOL" {
			return nil, fmt.Errorf("assert statement not testing a Boolean expression. got=%s", valtype.Type)
		}
		return node, err

	case *ast.ForStatement:
		var st1 []ast.Statement
		var st2 []ast.Statement
		for _, v := range node.Inits.Statements {
			tnode, err = c.typecheck(v)
			if err != nil {
				return node, err
			}
			st1 = append(st1, tnode.(ast.Statement))
		}
		node.Inits.Statements = st1

		for _, v := range node.Body.Statements {
			tnode, err = c.typecheck(v)
			if err != nil {
				return node, err
			}
			st2 = append(st2, tnode.(ast.Statement))
		}
		node.Body.Statements = st2
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
			switch e := node.Statements[i].(type) {
			case *ast.ExpressionStatement:
				exp := e.Expression
				typedNode, err := c.inferFunction(exp)
				if err != nil {
					return nil, err

				}
				node.Statements[i].(*ast.ExpressionStatement).Expression = typedNode
				valtype = typeable(typedNode)
				node.Statements[i].(*ast.ExpressionStatement).InferredType = valtype
			case *ast.ParallelFunctions:
				var exp []ast.Expression
				var valtype *ast.Type
				for _, f := range e.Expressions {
					typedNode, err := c.inferFunction(f)
					if err != nil {
						return nil, err

					}
					exp = append(exp, typedNode)
					valtype = typeable(typedNode)
				}
				node.Statements[i].(*ast.ParallelFunctions).Expressions = exp
				node.Statements[i].(*ast.ParallelFunctions).InferredType = valtype
			}
		}
		node.InferredType = valtype
		return node, err
	case *ast.BuiltIn:
		return node, nil
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
	case *ast.BuiltIn:
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
	case *ast.IndexExpression:
		return true
	default:
		return false
	}
}

func (c *Checker) infer(exp interface{}) (ast.Node, error) {
	switch node := exp.(type) {
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
	case *ast.BuiltIn:
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
			t, err := c.LookupType(node)
			if err != nil {
				return nil, err
			}
			node.InferredType = t
		}
		return node, nil
	case *ast.Instance:
		if node.InferredType == nil {
			t, err := c.LookupType(node)
			node.InferredType = t
			if err != nil {
				return nil, err
			}
		}
		return node, nil
	case *ast.StructInstance:
		if node.InferredType == nil {
			t, err := c.LookupType(node)
			node.InferredType = t
			if err != nil {
				return nil, err
			}
		}
		return node, nil
	case *ast.ParameterCall:
		if node.InferredType == nil {
			t, err := c.LookupType(node)
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
		rawid := node.Left.(ast.Nameable).RawId()
		spec := c.SpecStructs[rawid[0]]
		con, _ := spec.FetchConstant(rawid[1])
		if con != nil {
			return nil, fmt.Errorf("variable %s is a constant cannot access by index", node.Left.String())
		}

		if node.InferredType == nil {
			t, err := c.LookupType(node.Left)
			if err != nil {
				return nil, err
			}
			node.InferredType = t
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

func (c *Checker) LookupType(node ast.Node) (*ast.Type, error) {
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
	case *ast.BuiltIn:
		return &ast.Type{Type: "BOOL",
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

	if v.TokenLiteral() == "COMPOUND_STRING" {
		return &ast.Type{Type: "BOOL"}, err
	}

	ret := typeable(v)
	if ret == nil {
		v2, err := c.typecheck(v)
		if err != nil {
			return nil, err
		}
		spec.UpdateVar(rawid, ty, v2)
		return typeable(v2), err
	}
	return ret, err
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

	case *ast.ParameterCall:
		exp, err := c.infer(node)
		return exp.(ast.Expression), err
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

		if node.TokenLiteral() == "SWAP" {
			return c.inferSwap(node)
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
		if right == nil {
			nr, err = c.infer(node.Right)
		}

		if node.Token.Type == "ASSIGN" || //In case of temp values
			node.Token.Type == "COMPOUND_STRING" { //In case of compound string based rules
			ty, _ := c.LookupType(node.Left)
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
			switch exp := node.Consequence.Statements[i].(type) {
			case *ast.ExpressionStatement:
				if c.isValue(exp.Expression) {
					typedNode, err = c.infer(exp.Expression)
				} else {
					typedNode, err = c.inferFunction(exp.Expression)
				}
				if err != nil {
					return nil, err
				}
				node.Consequence.Statements[i].(*ast.ExpressionStatement).Expression = typedNode.(ast.Expression)
				valtype = typeable(typedNode)
				node.Consequence.InferredType = valtype
			case *ast.ParallelFunctions:
				for idx, p := range exp.Expressions {
					typedNode, err = c.inferFunction(p)
					exp.Expressions[idx] = typedNode.(ast.Expression)
				}
			}
		}

		if node.Alternative != nil {
			for i := 0; i < len(node.Alternative.Statements); i++ {
				switch exp := node.Alternative.Statements[i].(type) {
				case *ast.ExpressionStatement:
					if c.isValue(exp.Expression) {
						typedNode, err = c.infer(exp.Expression)
					} else {
						typedNode, err = c.inferFunction(exp.Expression)
					}
					if err != nil {
						return nil, err
					}
					node.Alternative.Statements[i].(*ast.ExpressionStatement).Expression = typedNode.(ast.Expression)
					valtype = typeable(typedNode)
					node.Alternative.InferredType = valtype
				case *ast.ParallelFunctions:
					for idx, p := range exp.Expressions {
						typedNode, err = c.inferFunction(p)
						exp.Expressions[idx] = typedNode.(ast.Expression)
					}
				}
			}
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

func (c *Checker) inferSwap(node *ast.InfixExpression) (ast.Expression, error) {
	var left, right ast.Node
	var err error

	left, err = c.inferSwapNode(node.Left)
	if err != nil {
		return node, err
	}

	right, err = c.inferSwapNode(node.Right)
	if err != nil {
		return node, err
	}

	if left.Type() != right.Type() {
		return node, fmt.Errorf("cannot redeclare variable %s is type %s got %s", node.Left.String(), left.Type(), right.Type())
	}

	if c.InstanceOf(left) != c.InstanceOf(right) {
		return node, fmt.Errorf("cannot redeclare variable %s is instance of %s got %s", node.Left.String(), c.InstanceOf(left), c.InstanceOf(right))
	}

	node.Left = left.(ast.Expression)
	node.Right = right.(ast.Expression)

	return node, err
}

func (c *Checker) inferSwapNode(node ast.Expression) (ast.Node, error) {
	switch n := node.(type) {
	case *ast.ParameterCall:
		rawid := n.RawId()
		spec := c.SpecStructs[rawid[0]]
		ty, _ := spec.GetStructType(rawid)
		p, err := spec.FetchVar(rawid, ty)
		if err != nil {
			n.InferredType = &ast.Type{Type: ty,
				Scope:      0,
				Parameters: nil}
			return n, nil
		}

		ref, err := c.lookupReference(p)
		if err != nil {
			return node, err
		}
		n.InferredType = &ast.Type{Type: ref.Type(),
			Scope:      0,
			Parameters: nil}
		return n, nil

	case *ast.Identifier:
		rawid := n.RawId()
		spec := c.SpecStructs[rawid[0]]
		ty, _ := spec.GetStructType(rawid)
		if ty != "" {
			n.InferredType = &ast.Type{Type: ty,
				Scope:      0,
				Parameters: nil}
			return n, nil
		}

		ref, err := c.lookupReference(n)
		if err != nil {
			return node, err
		}
		n.InferredType = &ast.Type{Type: ref.Type(),
			Scope:      0,
			Parameters: nil}
		return n, nil

	default:
		return c.lookupReference(node)
	}
}

func (c *Checker) swapValues(base *ast.StructInstance) (*ast.StructInstance, error) {
	var swaps []ast.Node
	for _, s := range base.Swaps {
		n, err := c.typecheck(s)
		if err != nil {
			return base, err
		}
		swaps = append(swaps, n)
		// infix := n.(*ast.InfixExpression)
		// rawid := infix.Left.(ast.Nameable).RawId()
		// key := rawid[len(rawid)-1]
		// val, err := c.lookupReference(infix.Right)
		// if err != nil {
		// 	return base, err
		// }

		// // Because part of what we're doing here is renaming
		// // these nodes. We need to do a deep copy to separate
		// // the swapped nodes from their original reference values
		// copyVal, err := deepcopy.Anything(val)
		// if err != nil {
		// 	return base, err
		// }

		// val = copyVal.(ast.Node)

		// switch v := val.(type) {
		// case *ast.StructInstance:
		// 	v.Name = key
		// 	val = v
		// }

		// base.Properties[key].Value = val
		// base = c.swapDeepNames(base)

	}
	base.Swaps = swaps
	return base, nil
}

func (c *Checker) swapDeepNames(val *ast.StructInstance) *ast.StructInstance {
	rawid := val.RawId()
	node, err := c.Preprocesser.Partial(rawid[0], val)
	if err != nil {
		panic(fmt.Sprintf("failed to update process ids on swap %s", val.String()))
	}
	return node.(*ast.StructInstance)
}

func (c *Checker) InstanceOf(node ast.Node) string {
	switch n := node.(type) {
	case *ast.StructInstance:
		return strings.Join(n.Parent, ".")

	case *ast.ParameterCall:
		key := n.IdString()
		value, ok := c.Instances[key]
		if !ok {
			return ""
		}

		return strings.Join(value.Parent, ".")
	case *ast.Identifier:
		key := n.IdString()
		value, ok := c.Instances[key]
		if !ok {
			return ""
		}

		return strings.Join(value.Parent, ".")
	default:
		return ""
	}
}

func (c *Checker) lookupReference(base ast.Node) (ast.Node, error) {
	switch b := base.(type) {
	case *ast.ParameterCall:
		rawid := b.RawId()
		spec := c.SpecStructs[rawid[0]]
		ty, _ := spec.GetStructType(rawid)
		p, err := spec.FetchVar(rawid, ty)
		if err == nil {
			return c.lookupReference(p)
		}
		id := b.Id()
		return c.lookupStruct(id, ty)
	case *ast.Identifier:
		// Check to see if this variable is referencing
		// a local variable
		rawid := b.RawId()
		spec := c.SpecStructs[rawid[0]]
		ty, _ := spec.GetStructType(rawid)
		p, err := spec.FetchVar(rawid, ty)

		if p == nil {
			id := b.Id()
			return c.lookupStruct(id, ty)
		}

		if _, ok := p.(*ast.Identifier); !ok {
			return c.infer(p)
		}

		// Assume it's referencing a constant then
		n, _ := spec.FetchConstant(b.Value)
		if n != nil {
			return n, err
		}
		return nil, fmt.Errorf("cannot establish node %s", b.IdString())
	default:
		if c.isValue(base) {
			return c.infer(base)
		} else {
			return c.inferFunction(base.(ast.Expression))
		}
	}
}

func (c *Checker) lookupStruct(id []string, ty string) (*ast.StructInstance, error) {
	spec := c.SpecStructs[id[0]]
	prop, err := spec.Fetch(id[1], ty)
	if err != nil {
		return nil, err
	}
	st := c.Instances[strings.Join(id, "_")]
	st.Properties = ast.WrapBranches(prop)
	return st, err
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
		props := ast.ExtractBranches(cnode.Properties)
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
	rawid = base.RawId()
	swappedBase, err := c.swapValues(base)
	if err != nil {
		return base, err
	}
	spec := c.SpecStructs[rawid[0]]
	prop := ast.ExtractBranches(swappedBase.Properties)
	spec.Update(base.RawId(), prop)

	return swappedBase, err
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
	if op == "then" {
		return &ast.Type{Type: "BOOL",
			Scope:      0,
			Parameters: nil}, nil
	}
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
	if IsNumeric(t1) && IsNumeric(t2) {
		return true
	}
	return false
}

func IsNumeric(t *ast.Type) bool {
	switch t.Type {
	case "INT":
		return true
	case "FLOAT":
		return true
	case "UNKNOWN":
		return true
	case "UNCERTAIN":
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
