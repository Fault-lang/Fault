package listener

// From antlr parse tree create Fault AST

import (
	"fault/ast"
	"fault/parser"
	"fmt"
	"strconv"
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type FaultListener struct {
	*parser.BaseFaultParserListener
	stack []interface{}
	AST   *ast.Spec
	scope string
}

func (l *FaultListener) push(n interface{}) {
	l.stack = append(l.stack, n)
}

func (l *FaultListener) pop() interface{} {
	var s interface{}
	s, l.stack = l.stack[len(l.stack)-1], l.stack[:len(l.stack)-1]
	return s
}

func (l *FaultListener) ExitSpec(c *parser.SpecContext) {
	var spec = &ast.Spec{}
	for _, v := range l.stack {
		spec.Statements = append(spec.Statements, v.(ast.Statement))
	}
	l.AST = spec
}

func (l *FaultListener) ExitSpecClause(c *parser.SpecClauseContext) {
	token := ast.Token{
		Type:    "SPEC_DECL",
		Literal: "SPEC_DECL",
		Position: []int{c.GetStart().GetLine(),
			c.GetStart().GetColumn(),
			c.GetStop().GetLine(),
			c.GetStop().GetColumn(),
		},
	}

	iden_token := ast.Token{
		Type:    "IDENT",
		Literal: "IDENT",
		Position: []int{c.GetStart().GetLine(),
			c.GetStart().GetColumn(),
			c.GetStop().GetLine(),
			c.GetStop().GetColumn(),
		},
	}

	l.push(
		&ast.SpecDeclStatement{
			Token: token,
			Name: &ast.Identifier{
				Token: iden_token,
				Value: c.IDENT().GetText(),
			},
		},
	)
}

func (l *FaultListener) ExitImportDecl(c *parser.ImportDeclContext) {
	items := len(c.AllImportSpec())

	var itemList []interface{}
	for i := 0; i < items; i++ {
		right := l.pop()

		var temp []interface{}
		temp = append(temp, right)

		itemList = append(temp, itemList...)
	}

	for _, v := range itemList {
		l.push(v)
	}
}

func (l *FaultListener) ExitImportSpec(c *parser.ImportSpecContext) {
	token := ast.Token{
		Type:    "IMPORT_DECL",
		Literal: "IMPORT_DECL",
		Position: []int{c.GetStart().GetLine(),
			c.GetStart().GetColumn(),
			c.GetStop().GetLine(),
			c.GetStop().GetColumn(),
		},
	}

	right := l.pop()
	val := right
	if val == nil {
		panic(fmt.Sprintf("top of stack not an expression: line %d col %d type %T", c.GetStart().GetLine(), c.GetStart().GetColumn(), right))
	}

	path, ok := val.(*ast.StringLiteral)
	if !ok {
		panic(fmt.Sprintf("import path not a string: line %d col %d type %T", c.GetStart().GetLine(), c.GetStart().GetColumn(), right))
	}

	// If no ident, create one from import path
	var ident *ast.Identifier
	if len(c.GetChildren()) == 2 {
		ident = &ast.Identifier{
			Token: token,
			Value: c.IDENT().GetText(),
		}
	} else {
		ident = &ast.Identifier{
			Token: token,
			Value: pathToIdent(path.String()),
		}
	}

	l.push(&ast.ImportStatement{
		Token: token,
		Name:  ident,
		Path:  path,
	})
}

func (l *FaultListener) ExitConstSpec(c *parser.ConstSpecContext) {
	token := ast.Token{
		Type:    "CONST_DECL",
		Literal: "CONST_DECL",
		Position: []int{c.GetStart().GetLine(),
			c.GetStart().GetColumn(),
			c.GetStop().GetLine(),
			c.GetStop().GetColumn(),
		},
	}

	right := l.pop()
	val := right
	if val == nil {
		panic(fmt.Sprintf("top of stack not an expression: line %d col %d type %T", c.GetStart().GetLine(), c.GetStart().GetColumn(), right))
	}
	var items int
	identlist, ok := c.GetChild(0).(*parser.IdentListContext)
	if ok {
		items = len(identlist.AllOperandName())
	} else {
		items = 1
	}

	var itemList []interface{}
	for i := 0; i < items; i++ {
		left := l.pop()
		ident, ok := left.(*ast.Identifier)
		if ident == nil {
			panic(fmt.Sprintf("top of stack not an expression: line %d col %d type %T", c.GetStart().GetLine(), c.GetStart().GetColumn(), left))
		}
		if !ok {
			panic(fmt.Sprintf("top of stack not an identifier: line %d col %d type %T", c.GetStart().GetLine(), c.GetStart().GetColumn(), left))
		}
		var temp []interface{}
		temp = append(temp, &ast.ConstantStatement{
			Token: token,
			Name:  ident,
			Value: val.(ast.Expression),
		})

		itemList = append( //Prepend to get in correct order
			temp, itemList...,
		)
	}
	for _, v := range itemList {
		l.push(v)
	}

}

func (l *FaultListener) EnterStructDecl(c *parser.StructDeclContext) {
	l.scope = c.GetChild(1).(antlr.TerminalNode).GetText()
}

func (l *FaultListener) ExitStructDecl(c *parser.StructDeclContext) {
	token := ast.Token{
		Type:    "ASSIGN",
		Literal: "=",
		Position: []int{c.GetStart().GetLine(),
			c.GetStart().GetColumn(),
			c.GetStop().GetLine(),
			c.GetStop().GetColumn(),
		},
	}
	right := l.pop()
	var val ast.Expression
	switch right.(type) {
	case *ast.StockLiteral, *ast.FlowLiteral:
		val = right.(ast.Expression)
	default:
		if right == nil {
			panic(fmt.Sprintf("top of stack not an expression: line %d col %d type %T", c.GetStart().GetLine(), c.GetStart().GetColumn(), right))
		}
		panic(fmt.Sprintf("def can only be used to define a valid stock or flow: line %d col %d", c.GetStart().GetLine(), c.GetStart().GetColumn()))
	}

	token2 := ast.Token{
		Type:    "IDENT",
		Literal: "IDENT",
		Position: []int{c.GetStart().GetLine(),
			c.GetStart().GetColumn(),
			c.GetStop().GetLine(),
			c.GetStop().GetColumn(),
		},
	}
	ident := &ast.Identifier{
		Token: token2,
		Value: c.IDENT().GetText(),
	}

	l.push(
		&ast.DefStatement{
			Token: token,
			Name:  ident,
			Value: val,
		})
	l.scope = ""
}

func (l *FaultListener) ExitStock(c *parser.StockContext) {
	pairs := c.AllStructProperties()
	token := ast.Token{
		Type:    "STOCK",
		Literal: "stock",
		Position: []int{c.GetStart().GetLine(),
			c.GetStart().GetColumn(),
			c.GetStop().GetLine(),
			c.GetStop().GetColumn(),
		},
	}

	l.push(
		&ast.StockLiteral{
			Token: token,
			Pairs: l.getPairs(len(pairs), []int{c.GetStart().GetLine(), c.GetStart().GetColumn()}),
		})
}

func (l *FaultListener) ExitFlow(c *parser.FlowContext) {
	pairs := c.AllStructProperties()
	token := ast.Token{
		Type:    "FLOW",
		Literal: "flow",
		Position: []int{c.GetStart().GetLine(),
			c.GetStart().GetColumn(),
			c.GetStop().GetLine(),
			c.GetStop().GetColumn(),
		},
	}

	l.push(
		&ast.FlowLiteral{
			Token: token,
			Pairs: l.getPairs(len(pairs), []int{c.GetStart().GetLine(), c.GetStart().GetColumn()}),
		},
	)
}

func (l *FaultListener) ExitPropInt(c *parser.PropIntContext) {
	val := l.pop()

	token2 := ast.Token{
		Type:    "IDENT",
		Literal: "IDENT",
		Position: []int{c.GetStart().GetLine(),
			c.GetStart().GetColumn(),
			c.GetStop().GetLine(),
			c.GetStop().GetColumn(),
		},
	}
	l.push(&ast.Identifier{
		Token: token2,
		Value: c.IDENT().GetText(),
	})

	l.push(val)
}

func (l *FaultListener) ExitPropString(c *parser.PropStringContext) {
	val := l.pop()

	token2 := ast.Token{
		Type:    "IDENT",
		Literal: "IDENT",
		Position: []int{c.GetStart().GetLine(),
			c.GetStart().GetColumn(),
			c.GetStop().GetLine(),
			c.GetStop().GetColumn(),
		},
	}
	l.push(&ast.Identifier{
		Token: token2,
		Value: c.IDENT().GetText(),
	},
	)

	l.push(val)
}

func (l *FaultListener) ExitPropVar(c *parser.PropVarContext) {
	f := l.pop()

	token2 := ast.Token{
		Type:    "IDENT",
		Literal: "IDENT",
		Position: []int{c.GetStart().GetLine(),
			c.GetStart().GetColumn(),
			c.GetStop().GetLine(),
			c.GetStop().GetColumn(),
		},
	}
	l.push(&ast.Identifier{
		Token: token2,
		Value: c.IDENT().GetText(),
	},
	)
	switch v := f.(type) {
	case *ast.Instance:
		v.Name = c.IDENT().GetText()
		l.push(v)
	case *ast.Identifier:
		l.push(v)
	default:
		panic(fmt.Sprintf("top of stack not an identifier: line %d col %d type %T", c.GetStart().GetLine(), c.GetStart().GetColumn(), f))
	}
}

// func (l *FaultListener) ExitInstance(c *parser.InstanceContext) {
// 	node := c.GetText()
// 	if len(node) > 3 && node[0:3] == "new" {
// 		val := l.pop()
// 		token := ast.Token{
// 			Type:    "FUNCTION",
// 			Literal: "FUNCTION",
// 			Position: []int{c.GetStart().GetLine(),
// 				c.GetStart().GetColumn(),
// 				c.GetStop().GetLine(),
// 				c.GetStop().GetColumn(),
// 			},
// 		}

// 		f := &ast.InstanceExpression{
// 			Token: token,
// 			Stock: val.(ast.Expression),
// 		}
// 		l.push(f)
// 	}

//}

func (l *FaultListener) EnterPropFunc(c *parser.PropFuncContext) {
	l.scope = fmt.Sprint(l.scope, ".", c.IDENT().GetText())
}

func (l *FaultListener) ExitPropFunc(c *parser.PropFuncContext) {
	val := l.pop()
	token1 := ast.Token{
		Type:    "IDENT",
		Literal: "IDENT",
		Position: []int{c.GetStart().GetLine(),
			c.GetStart().GetColumn(),
			c.GetStop().GetLine(),
			c.GetStop().GetColumn(),
		},
	}
	l.push(&ast.Identifier{
		Token: token1,
		Value: c.IDENT().GetText(),
	},
	)
	l.push(val)
}

func (l *FaultListener) ExitFunctionLit(c *parser.FunctionLitContext) {
	token := ast.Token{
		Type:    "FUNCTION",
		Literal: "FUNCTION",
		Position: []int{c.GetStart().GetLine(),
			c.GetStart().GetColumn(),
			c.GetStop().GetLine(),
			c.GetStop().GetColumn(),
		},
	}

	b := l.pop()

	f := &ast.FunctionLiteral{
		Token: token,
		Body:  b.(*ast.BlockStatement),
	}
	l.push(f)
}

func (l *FaultListener) ExitStatementList(c *parser.StatementListContext) {
	token := ast.Token{
		Type:    "FUNCTION",
		Literal: "FUNCTION",
		Position: []int{c.GetStart().GetLine(),
			c.GetStart().GetColumn(),
			c.GetStop().GetLine(),
			c.GetStop().GetColumn(),
		},
	}

	sl := &ast.BlockStatement{Token: token}
	for _, v := range c.GetChildren() {
		ex := l.pop()
		switch ex.(type) {
		case ast.Statement:
			sl.Statements = append([]ast.Statement{ex.(ast.Statement)}, sl.Statements...)
		case ast.Expression:
			token2 := ast.Token{
				Type:    "FUNCTION",
				Literal: "FUNCTION",
				Position: []int{v.(*parser.StatementContext).GetStart().GetLine(),
					v.(*parser.StatementContext).GetStart().GetColumn(),
					v.(*parser.StatementContext).GetStop().GetLine(),
					v.(*parser.StatementContext).GetStop().GetColumn(),
				},
			}
			s := &ast.ExpressionStatement{
				Token:      token2,
				Expression: ex.(ast.Expression),
			}
			sl.Statements = append([]ast.Statement{s}, sl.Statements...)
		default:
			panic(fmt.Sprintf("Neither statement nor expression got=%T", v))
		}
	}
	l.push(sl)
}

func (l *FaultListener) ExitFaultAssign(c *parser.FaultAssignContext) {
	operator := c.GetChild(1).(antlr.TerminalNode).GetText()
	token := ast.Token{
		Type:    "ASSIGN",
		Literal: operator,
		Position: []int{c.GetStart().GetLine(),
			c.GetStart().GetColumn(),
			c.GetStop().GetLine(),
			c.GetStop().GetColumn(),
		},
	}

	var receiver ast.Expression
	var sender ast.Expression
	var ok bool

	right := l.pop()
	left := l.pop()
	if operator == "->" {
		switch right.(type) {
		case *ast.ParameterCall:
			receiver = right.(ast.Expression)
		case *ast.Identifier:
			receiver = right.(ast.Expression)
		default:
			panic(fmt.Sprintf("right side of expression should be a parameter call or identifier: line %d col %d type %T", c.GetStart().GetLine(), c.GetStart().GetColumn(), right))
		}

		sender, ok = left.(ast.Expression)
		if !ok {
			panic(fmt.Sprintf("left side of expression should be an expression: line %d col %d type %T", c.GetStart().GetLine(), c.GetStart().GetColumn(), left))
		}

	} else if operator == "<-" {
		switch left.(type) {
		case *ast.ParameterCall:
			receiver = left.(ast.Expression)
		case *ast.Identifier:
			receiver = left.(ast.Expression)
		default:
			panic(fmt.Sprintf("left side of expression should be a parameter call or identifier: line %d col %d type %T", c.GetStart().GetLine(), c.GetStart().GetColumn(), left))
		}
		sender, ok = right.(ast.Expression)
		if !ok {
			panic(fmt.Sprintf("right side of expression should be an expression: line %d col %d type %T", c.GetStart().GetLine(), c.GetStart().GetColumn(), right))
		}

	} else {
		panic(fmt.Sprintf("Invalid operator %s in expression", operator))
	}

	l.push(
		&ast.InfixExpression{
			Token:    token,
			Left:     receiver,
			Operator: "<-", // "->" converted automatically by swapping order
			Right:    sender,
		})

}

func (l *FaultListener) ExitMiscAssign(c *parser.MiscAssignContext) {
	token := ast.Token{
		Type:    "ASSIGN",
		Literal: c.GetChild(1).(antlr.TerminalNode).GetText(),
		Position: []int{c.GetStart().GetLine(),
			c.GetStart().GetColumn(),
			c.GetStop().GetLine(),
			c.GetStop().GetColumn(),
		},
	}

	right := l.pop()
	if right == nil {
		panic(fmt.Sprintf("top of stack not an expression: line %d col %d type %T", c.GetStart().GetLine(), c.GetStart().GetColumn(), right))
	}

	left := l.pop()
	ident, ok := left.(*ast.Identifier)
	if !ok {
		panic(fmt.Sprintf("left side of expression should be an identifier: line %d col %d type %T", c.GetStart().GetLine(), c.GetStart().GetColumn(), left))
	}

	// If a new instance is initialized in the run block
	// the listener needs to add the name
	switch inst := right.(type) {
	case *ast.Instance:
		inst.Name = ident.Value
		right = inst
	}

	l.push(
		&ast.InfixExpression{
			Token:    token,
			Left:     ident,
			Operator: c.GetChild(1).(antlr.TerminalNode).GetText(),
			Right:    right.(ast.Expression),
		})
}

func (l *FaultListener) ExitLrExpr(c *parser.LrExprContext) {
	token := ast.Token{
		Type:    ast.OPS[c.GetChild(1).(antlr.TerminalNode).GetText()],
		Literal: c.GetChild(1).(antlr.TerminalNode).GetText(),
		Position: []int{c.GetStart().GetLine(),
			c.GetStart().GetColumn(),
			c.GetStop().GetLine(),
			c.GetStop().GetColumn(),
		},
	}

	rght := l.pop()
	lft := l.pop()
	e := &ast.InfixExpression{
		Token:    token,
		Left:     lft.(ast.Expression),
		Operator: c.GetChild(1).(antlr.TerminalNode).GetText(),
		Right:    rght.(ast.Expression),
	}
	l.push(e)
}

func (l *FaultListener) ExitRunStepExpr(c *parser.RunStepExprContext) {
	token := ast.Token{
		Type:    ast.OPS[c.GetChild(1).(antlr.TerminalNode).GetText()],
		Literal: c.GetChild(1).(antlr.TerminalNode).GetText(),
		Position: []int{c.GetStart().GetLine(),
			c.GetStart().GetColumn(),
			c.GetStop().GetLine(),
			c.GetStop().GetColumn(),
		},
	}

	rght := l.pop()
	lft := l.pop()
	e := &ast.InfixExpression{
		Token:    token,
		Left:     lft.(ast.Expression),
		Operator: c.GetChild(1).(antlr.TerminalNode).GetText(),
		Right:    rght.(ast.Expression),
	}
	l.push(e)
}

func (l *FaultListener) ExitPrefix(c *parser.PrefixContext) {
	token := ast.Token{
		Type:    ast.OPS[c.GetChild(0).(antlr.TerminalNode).GetText()],
		Literal: c.GetChild(0).(antlr.TerminalNode).GetText(),
		Position: []int{c.GetStart().GetLine(),
			c.GetStart().GetColumn(),
			c.GetStop().GetLine(),
			c.GetStop().GetColumn(),
		},
	}

	rght := l.pop()
	e := &ast.PrefixExpression{
		Token:    token,
		Operator: c.GetChild(0).(antlr.TerminalNode).GetText(),
		Right:    rght.(ast.Expression),
	}
	l.push(e)
}

func (l *FaultListener) ExitTyped(c *parser.TypedContext) {
	switch c.FaultType().GetText() {
	case "natural":
		token := ast.Token{
			Type:    "NATURAL",
			Literal: "NATURAL",
			Position: []int{c.GetStart().GetLine(),
				c.GetStart().GetColumn(),
				c.GetStop().GetLine(),
				c.GetStop().GetColumn(),
			},
		}

		value := l.pop()
		nat, ok := value.(*ast.IntegerLiteral)

		if !ok {
			panic(fmt.Sprintf("Invalid value cast to type natural. got=%T at line %d col %d", value, c.GetStart().GetLine(), c.GetStart().GetColumn()))
		}

		l.push(&ast.Natural{
			Token: token,
			Value: nat.Value,
		})

	case "uncertain":
		token := ast.Token{
			Type:    "UNCERTAIN",
			Literal: "UNCERTAIN",
			Position: []int{c.GetStart().GetLine(),
				c.GetStart().GetColumn(),
				c.GetStop().GetLine(),
				c.GetStop().GetColumn(),
			},
		}

		v1 := l.pop()
		sigma, err := l.intOrFloatOk(v1)
		if err != nil {
			panic(fmt.Sprintf("Invalid value for sigma of type uncertain. got=%T at: line %d col %d", v1, c.GetStart().GetLine(), c.GetStart().GetColumn()))
		}

		v2 := l.pop()
		mean, err := l.intOrFloatOk(v2)
		if err != nil {
			panic(fmt.Sprintf("Invalid value for mean of type uncertain. got=%T at: line %d col %d", v2, c.GetStart().GetLine(), c.GetStart().GetColumn()))
		}

		l.push(&ast.Uncertain{
			Token: token,
			Mean:  mean,
			Sigma: sigma,
		})
	}
}

func (l *FaultListener) ExitIncDecStmt(c *parser.IncDecStmtContext) {
	var tType ast.TokenType
	var tLit string
	if c.PLUS_PLUS() != nil {
		tType = "PLUS"
		tLit = "+"
	} else if c.MINUS_MINUS() != nil {
		tType = "MINUS"
		tLit = "-"
	} else {
		panic(fmt.Sprintf("Illegal operation: line %d col %d", c.GetStart().GetLine(), c.GetStart().GetColumn()))
	}

	token := ast.Token{
		Type:    tType,
		Literal: tLit,
		Position: []int{c.GetStart().GetLine(),
			c.GetStart().GetColumn(),
			c.GetStop().GetLine(),
			c.GetStop().GetColumn(),
		},
	}

	ident := l.pop()

	token2 := ast.Token{
		Type:    "INT",
		Literal: "INT",
		Position: []int{c.GetStart().GetLine(),
			c.GetStart().GetColumn(),
			c.GetStop().GetLine(),
			c.GetStop().GetColumn(),
		},
	}
	e := &ast.InfixExpression{
		Token:    token,
		Left:     ident.(ast.Expression),
		Operator: tLit,
		Right:    &ast.IntegerLiteral{Token: token2, Value: 1},
	}
	l.push(e)
}

func (l *FaultListener) ExitIfStmt(c *parser.IfStmtContext) {
	token := ast.Token{
		Type:    "IF",
		Literal: "IF",
		Position: []int{c.GetStart().GetLine(),
			c.GetStart().GetColumn(),
			c.GetStop().GetLine(),
			c.GetStop().GetColumn(),
		},
	}
	a := l.pop()
	csq := l.pop()
	cond := l.pop()

	e := &ast.IfExpression{
		Token:       token,
		Condition:   cond.(ast.Expression),
		Consequence: csq.(*ast.BlockStatement),
		Alternative: a.(*ast.BlockStatement),
	}

	l.push(e)
}

func (l *FaultListener) ExitAccessHistory(c *parser.AccessHistoryContext) {
	token := ast.Token{
		Type:    "HISTORY",
		Literal: "HISTORY",
		Position: []int{c.GetStart().GetLine(),
			c.GetStart().GetColumn(),
			c.GetStop().GetLine(),
			c.GetStop().GetColumn(),
		},
	}
	var exp []ast.Expression
	for i := 0; i < len(c.AllExpression()); i++ {
		idx := l.pop()
		exp = append([]ast.Expression{idx.(ast.Expression)}, exp...)
	}
	for i := 0; i < len(exp); i++ {
		ident := l.pop()
		left := ident

		right := exp[i]
		l.push(&ast.IndexExpression{
			Token: token,
			Left:  left.(ast.Expression),
			Index: right,
		})
	}

}

func (l *FaultListener) ExitOpName(c *parser.OpNameContext) {
	token := ast.Token{
		Type:    "IDENT",
		Literal: "IDENT",
		Position: []int{c.GetStart().GetLine(),
			c.GetStart().GetColumn(),
			c.GetStop().GetLine(),
			c.GetStop().GetColumn(),
		},
	}
	l.push(&ast.Identifier{
		Token: token,
		Value: c.GetText(),
	},
	)
}

func (l *FaultListener) ExitOpParam(c *parser.OpParamContext) {
	token := ast.Token{
		Type:    "IDENT",
		Literal: "IDENT",
		Position: []int{c.GetStart().GetLine(),
			c.GetStart().GetColumn(),
			c.GetStop().GetLine(),
			c.GetStop().GetColumn(),
		},
	}
	v := c.GetText()
	param := strings.Split(v, ".")

	l.push(&ast.ParameterCall{
		Token: token,
		Value: param,
	},
	)
}

func (l *FaultListener) ExitOpInstance(c *parser.OpInstanceContext) {
	token := ast.Token{
		Type:    "IDENT",
		Literal: "IDENT",
		Position: []int{c.GetStart().GetLine(),
			c.GetStart().GetColumn(),
			c.GetStop().GetLine(),
			c.GetStop().GetColumn(),
		},
	}
	ident := &ast.Identifier{
		Token: token,
		Value: c.IDENT().GetText(),
	}

	l.push(&ast.Instance{
		Token: token,
		Value: ident,
	},
	)
}

func (l *FaultListener) ExitOpThis(c *parser.OpThisContext) {
	token := ast.Token{
		Type:    "IDENT",
		Literal: "IDENT",
		Position: []int{c.GetStart().GetLine(),
			c.GetStart().GetColumn(),
			c.GetStop().GetLine(),
			c.GetStop().GetColumn(),
		},
	}
	l.push(&ast.This{
		Token: token,
		Value: strings.Split(l.scope, "."),
	},
	)
}

func (l *FaultListener) ExitOpClock(c *parser.OpClockContext) {
	token := ast.Token{
		Type:    "IDENT",
		Literal: "IDENT",
		Position: []int{c.GetStart().GetLine(),
			c.GetStart().GetColumn(),
			c.GetStop().GetLine(),
			c.GetStop().GetColumn(),
		},
	}
	l.push(&ast.Clock{
		Token: token,
		Value: l.scope,
	},
	)
}

func (l *FaultListener) ExitOperand(c *parser.OperandContext) {
	if c.GetText() == "nil" {
		token := ast.Token{
			Type:    "NIL",
			Literal: "NIL",
			Position: []int{c.GetStart().GetLine(),
				c.GetStart().GetColumn(),
				c.GetStop().GetLine(),
				c.GetStop().GetColumn(),
			},
		}
		l.push(&ast.Nil{
			Token: token,
		})

	}
}

func (l *FaultListener) ExitInteger(c *parser.IntegerContext) {
	token := ast.Token{
		Type:    "INT",
		Literal: "INT",
		Position: []int{c.GetStart().GetLine(),
			c.GetStart().GetColumn(),
			c.GetStop().GetLine(),
			c.GetStop().GetColumn(),
		},
	}

	v, err := strconv.ParseInt(c.GetText(), 10, 64)
	if err != nil {
		panic(fmt.Sprintf("integer value detected but not parsable: line %d col %d got=%s", c.GetStart().GetLine(), c.GetStart().GetColumn(), c.GetText()))
	}

	l.push(&ast.IntegerLiteral{
		Token: token,
		Value: v,
	})
}

func (l *FaultListener) ExitNegative(c *parser.NegativeContext) {
	base := l.pop()

	var token ast.Token
	switch base.(type) {
	case *ast.IntegerLiteral:
		token = ast.Token{
			Type:    "INT",
			Literal: "INT",
			Position: []int{c.GetStart().GetLine(),
				c.GetStart().GetColumn(),
				c.GetStop().GetLine(),
				c.GetStop().GetColumn(),
			},
		}
		i := base.(*ast.IntegerLiteral)
		i.Token = token
		i.Value = -i.Value

		l.push(i)

	case *ast.FloatLiteral:
		token = ast.Token{
			Type:    "FLOAT",
			Literal: "FLOAT",
			Position: []int{c.GetStart().GetLine(),
				c.GetStart().GetColumn(),
				c.GetStop().GetLine(),
				c.GetStop().GetColumn(),
			},
		}

		i := base.(*ast.FloatLiteral)
		i.Token = token
		i.Value = -i.Value

		l.push(i)

	default:
		panic(fmt.Sprintf("top of stack not an integer or a float got=%T", base))
	}

}

func (l *FaultListener) ExitFloat_(c *parser.Float_Context) {
	token := ast.Token{
		Type:    "FLOAT",
		Literal: "FLOAT",
		Position: []int{c.GetStart().GetLine(),
			c.GetStart().GetColumn(),
			c.GetStop().GetLine(),
			c.GetStop().GetColumn(),
		},
	}

	v, err := strconv.ParseFloat(c.GetText(), 64)
	if err != nil {
		panic(fmt.Sprintf("float value detected but not parsable: line %d col %d got=%s", c.GetStart().GetLine(), c.GetStart().GetColumn(), c.GetText()))
	}

	l.push(&ast.FloatLiteral{
		Token: token,
		Value: v,
	})
}

func (l *FaultListener) ExitString_(c *parser.String_Context) {
	token := ast.Token{
		Type:    "STRING",
		Literal: "STRING",
		Position: []int{c.GetStart().GetLine(),
			c.GetStart().GetColumn(),
			c.GetStop().GetLine(),
			c.GetStop().GetColumn(),
		},
	}

	v := c.GetText()

	l.push(&ast.StringLiteral{
		Token: token,
		Value: v,
	})
}

func (l *FaultListener) ExitBool_(c *parser.Bool_Context) {
	token := ast.Token{
		Type:    "BOOL",
		Literal: "BOOL",
		Position: []int{c.GetStart().GetLine(),
			c.GetStart().GetColumn(),
			c.GetStop().GetLine(),
			c.GetStop().GetColumn(),
		},
	}

	v, err := strconv.ParseBool(c.GetText())
	if err != nil {
		panic(fmt.Sprintf("Detected boolean will not parse: line %d col %d", c.GetStart().GetLine(), c.GetStart().GetColumn()))
	}

	l.push(&ast.Boolean{
		Token: token,
		Value: v,
	})
}

func (l *FaultListener) ExitBlock(c *parser.BlockContext) {
	if len(c.GetChildren()) == 2 { // If 2 this is an empty block
		token := ast.Token{
			Type:    "FUNCTION",
			Literal: "FUNCTION",
			Position: []int{c.GetStart().GetLine(),
				c.GetStart().GetColumn(),
				c.GetStop().GetLine(),
				c.GetStop().GetColumn(),
			},
		}

		l.push(&ast.BlockStatement{
			Token: token,
		})
	}
}

func (l *FaultListener) ExitInitDecl(c *parser.InitDeclContext) {
	token := ast.Token{
		Type:    "ASSIGN",
		Literal: "init",
		Position: []int{c.GetStart().GetLine(),
			c.GetStart().GetColumn(),
			c.GetStop().GetLine(),
			c.GetStop().GetColumn(),
		},
	}
	init := l.pop()
	if init == nil {
		panic(fmt.Sprintf("top of stack not an expression: line %d col %d type %T", c.GetStart().GetLine(), c.GetStart().GetColumn(), init))
	}
	l.push(&ast.InitExpression{
		Token:      token,
		Expression: init.(ast.Expression),
	})

}

func (l *FaultListener) ExitForStmt(c *parser.ForStmtContext) {
	token := ast.Token{
		Type:    "FOR",
		Literal: "for",
		Position: []int{c.GetStart().GetLine(),
			c.GetStart().GetColumn(),
			c.GetStop().GetLine(),
			c.GetStop().GetColumn(),
		},
	}

	rg := l.pop()
	lf := l.pop()

	if rg == nil {
		panic(fmt.Sprintf("top of stack not an expression: line %d col %d type %T", c.GetStart().GetLine(), c.GetStart().GetColumn(), rg))
	}

	block, ok := rg.(*ast.BlockStatement)
	if !ok {
		panic(fmt.Sprintf("top of stack not a block statement: line %d col %d type %T", c.GetStart().GetLine(), c.GetStart().GetColumn(), rg))
	}

	if lf == nil {
		panic(fmt.Sprintf("top of stack not an statement: line %d col %d type %T", c.GetStart().GetLine(), c.GetStart().GetColumn(), lf))
	}

	rounds, ok := lf.(*ast.IntegerLiteral)
	if !ok {
		panic(fmt.Sprintf("top of stack not an integer literal: line %d col %d type %T", c.GetStart().GetLine(), c.GetStart().GetColumn(), lf))
	}

	forSt := &ast.ForStatement{
		Token:  token,
		Rounds: rounds,
		Body:   block,
	}

	l.push(forSt)
}

func (l *FaultListener) ExitAssertion(c *parser.AssertionContext) {
	token := ast.Token{
		Type:    "ASSERT",
		Literal: "assert",
		Position: []int{c.GetStart().GetLine(),
			c.GetStart().GetColumn(),
			c.GetStop().GetLine(),
			c.GetStop().GetColumn(),
		},
	}

	expr := l.pop()
	l.push(&ast.AssertionStatement{
		Token:      token,
		Expression: expr.(ast.Expression),
	})
}

func (l *FaultListener) getPairs(p int, pos []int) map[ast.Expression]ast.Expression {
	pairs := make(map[ast.Expression]ast.Expression)
	for i := 0; i < p; i++ {
		right := l.pop()
		if right == nil {
			panic(fmt.Sprintf("top of stack not an expression: line %d col %d type %T", pos[0], pos[1], right))
		}

		left := l.pop()
		ident, ok := left.(*ast.Identifier)
		if ident == nil {
			panic(fmt.Sprintf("top of stack not an expression: line %d col %d type %T", pos[0], pos[1], left))
		}
		if !ok {
			panic(fmt.Sprintf("top of stack not an identifier: line %d col %d type %T", pos[0], pos[1], left))
		}
		pairs[ident] = right.(ast.Expression)
	}
	return pairs
}

func (l *FaultListener) intOrFloatOk(v interface{}) (float64, error) {
	switch val := v.(type) {
	case *ast.FloatLiteral:
		return val.Value, nil
	case *ast.IntegerLiteral:
		return float64(val.Value), nil
	default:
		return 0, fmt.Errorf("Invalid input type. Should be float or int got=%T", v)
	}
}

func pathToIdent(path string) string {
	s1 := strings.ReplaceAll(path, ".", "")
	s2 := strings.ReplaceAll(s1, "~", "")
	s3 := strings.ReplaceAll(s2, "\\", "")
	s4 := strings.ReplaceAll(s3, `"`, "")
	return strings.ReplaceAll(s4, "/", "")
}
