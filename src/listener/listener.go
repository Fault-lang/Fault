package listener

// From antlr parse tree create Fault AST

import (
	"fault/ast"
	"fault/parser"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type FaultListener struct {
	*parser.BaseFaultParserListener
	stack []interface{}
	AST   *ast.Spec
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
		log.Fatal(
			fmt.Errorf("top of stack not an expression: line %d col %d type %T", c.GetStart().GetLine(), c.GetStart().GetColumn(), right),
		)
	}

	path, ok := val.(*ast.StringLiteral)
	if !ok {
		log.Fatal(
			fmt.Errorf("import path not a string: line %d col %d type %T", c.GetStart().GetLine(), c.GetStart().GetColumn(), right),
		)
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
		log.Fatal(
			fmt.Errorf("top of stack not an expression: line %d col %d type %T", c.GetStart().GetLine(), c.GetStart().GetColumn(), right),
		)
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
			log.Fatal(
				fmt.Errorf("top of stack not an expression: line %d col %d type %T", c.GetStart().GetLine(), c.GetStart().GetColumn(), left),
			)
		}
		if !ok {
			log.Fatal(
				fmt.Errorf("top of stack not an identifier: line %d col %d type %T", c.GetStart().GetLine(), c.GetStart().GetColumn(), left),
			)
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
			log.Fatal(
				fmt.Errorf("top of stack not an expression: line %d col %d type %T", c.GetStart().GetLine(), c.GetStart().GetColumn(), right),
			)
		}
		log.Fatal(
			fmt.Errorf("def can only be used to define a valid stock or flow: line %d col %d", c.GetStart().GetLine(), c.GetStart().GetColumn()),
		)
	}

	left := l.pop()
	ident, ok := left.(*ast.Identifier)
	if ident == nil {
		log.Fatal(
			fmt.Errorf("top of stack not an expression: line %d col %d type %T", c.GetStart().GetLine(), c.GetStart().GetColumn(), left),
		)
	}
	if !ok {
		log.Fatal(
			fmt.Errorf("top of stack not an identifier: line %d col %d type %T", c.GetStart().GetLine(), c.GetStart().GetColumn(), left),
		)
	}

	l.push(
		&ast.DefStatement{
			Token: token,
			Name:  ident,
			Value: val,
		})

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
	token1 := ast.Token{
		Type:    "FUNCTION",
		Literal: "FUNCTION",
		Position: []int{c.GetStart().GetLine(),
			c.GetStart().GetColumn(),
			c.GetStop().GetLine(),
			c.GetStop().GetColumn(),
		},
	}

	statement := &ast.ExpressionStatement{
		Token:      token1,
		Expression: val.(ast.Expression),
	}

	f := &ast.FunctionLiteral{
		Token: token1,
		Body: &ast.BlockStatement{
			Token:      token1,
			Statements: []ast.Statement{statement},
		},
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
	l.push(&ast.Identifier{
		Token: token2,
		Value: c.IDENT().GetText(),
	})

	l.push(f)
}

func (l *FaultListener) ExitPropString(c *parser.PropStringContext) {
	val := l.pop()
	token1 := ast.Token{
		Type:    "FUNCTION",
		Literal: "FUNCTION",
		Position: []int{c.GetStart().GetLine(),
			c.GetStart().GetColumn(),
			c.GetStop().GetLine(),
			c.GetStop().GetColumn(),
		},
	}

	statement := &ast.ExpressionStatement{
		Token:      token1,
		Expression: val.(ast.Expression),
	}

	f := &ast.FunctionLiteral{
		Token: token1,
		Body: &ast.BlockStatement{
			Token:      token1,
			Statements: []ast.Statement{statement},
		},
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
	l.push(&ast.Identifier{
		Token: token2,
		Value: c.IDENT().GetText(),
	},
	)

	l.push(f)
}

func (l *FaultListener) ExitPropVar(c *parser.PropVarContext) {
	val := l.pop()
	token1 := ast.Token{
		Type:    "FUNCTION",
		Literal: "FUNCTION",
		Position: []int{c.GetStart().GetLine(),
			c.GetStart().GetColumn(),
			c.GetStop().GetLine(),
			c.GetStop().GetColumn(),
		},
	}

	f := &ast.InstanceExpression{
		Token: token1,
		Stock: val.(ast.Expression),
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
	l.push(&ast.Identifier{
		Token: token2,
		Value: c.IDENT().GetText(),
	},
	)

	l.push(f)
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
			log.Fatal(
				fmt.Errorf("Neither statement nor expression got=%T", v),
			)
		}
	}
	l.push(sl)
}

func (l *FaultListener) ExitFaultAssign(c *parser.FaultAssignContext) {
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
	receiver, ok := right.(*ast.Identifier)
	if receiver == nil {
		log.Fatal(
			fmt.Errorf("top of stack not an expression: line %d col %d type %T", c.GetStart().GetLine(), c.GetStart().GetColumn(), right),
		)
	}
	if !ok {
		log.Fatal(
			fmt.Errorf("top of stack not an identifier: line %d col %d type %T", c.GetStart().GetLine(), c.GetStart().GetColumn(), right),
		)
	}

	left := l.pop()
	sender, ok := left.(*ast.Identifier)
	if sender == nil {
		log.Fatal(
			fmt.Errorf("top of stack not an expression: line %d col %d type %T", c.GetStart().GetLine(), c.GetStart().GetColumn(), left),
		)
	}
	if !ok {
		log.Fatal(
			fmt.Errorf("top of stack not an identifier: line %d col %d type %T", c.GetStart().GetLine(), c.GetStart().GetColumn(), left),
		)
	}
	l.push(
		&ast.InfixExpression{
			Token:    token,
			Left:     sender,
			Operator: c.GetChild(1).(antlr.TerminalNode).GetText(),
			Right:    receiver,
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
		log.Fatal(
			fmt.Errorf("top of stack not an expression: line %d col %d type %T", c.GetStart().GetLine(), c.GetStart().GetColumn(), right),
		)
	}

	left := l.pop()
	ident, ok := left.(*ast.Identifier)
	if ident == nil {
		log.Fatal(
			fmt.Errorf("top of stack not an expression: line %d col %d type %T", c.GetStart().GetLine(), c.GetStart().GetColumn(), left),
		)
	}
	if !ok {
		log.Fatal(
			fmt.Errorf("top of stack not an identifier: line %d col %d type %T", c.GetStart().GetLine(), c.GetStart().GetColumn(), left),
		)
	}

	l.push(
		&ast.DefStatement{
			Token: token,
			Name:  ident,
			Value: right.(ast.Expression),
		},
	)
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
		log.Fatal(
			fmt.Errorf("Illegal operation: line %d col %d", c.GetStart().GetLine(), c.GetStart().GetColumn()),
		)
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

func (l *FaultListener) ExitOperandName(c *parser.OperandNameContext) {
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
		log.Fatal(
			fmt.Errorf("integer value detected but not parsable: line %d col %d got=%s", c.GetStart().GetLine(), c.GetStart().GetColumn(), c.GetText()),
		)
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
		log.Fatal(
			fmt.Errorf("top of stack not an integer or a float got=%T", base),
		)
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
		log.Fatal(
			fmt.Errorf("float value detected but not parsable: line %d col %d got=%s", c.GetStart().GetLine(), c.GetStart().GetColumn(), c.GetText()),
		)
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
		log.Fatal(
			fmt.Errorf("Detected boolean will not parse: line %d col %d", c.GetStart().GetLine(), c.GetStart().GetColumn()),
		)
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
		log.Fatal(
			fmt.Errorf("top of stack not an expression: line %d col %d type %T", c.GetStart().GetLine(), c.GetStart().GetColumn(), init),
		)
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
		log.Fatal(
			fmt.Errorf("top of stack not an expression: line %d col %d type %T", c.GetStart().GetLine(), c.GetStart().GetColumn(), rg),
		)
	}

	block, ok := rg.(*ast.BlockStatement)
	if !ok {
		log.Fatal(
			fmt.Errorf("top of stack not a block statement: line %d col %d type %T", c.GetStart().GetLine(), c.GetStart().GetColumn(), rg),
		)
	}

	if lf == nil {
		log.Fatal(
			fmt.Errorf("top of stack not an statement: line %d col %d type %T", c.GetStart().GetLine(), c.GetStart().GetColumn(), lf),
		)
	}

	rounds, ok := lf.(*ast.IntegerLiteral)
	if !ok {
		log.Fatal(
			fmt.Errorf("top of stack not an integer literal: line %d col %d type %T", c.GetStart().GetLine(), c.GetStart().GetColumn(), lf),
		)
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
			log.Fatal(
				fmt.Errorf("top of stack not an expression: line %d col %d type %T", pos[0], pos[1], right),
			)
		}

		left := l.pop()
		ident, ok := left.(*ast.Identifier)
		if ident == nil {
			log.Fatal(
				fmt.Errorf("top of stack not an expression: line %d col %d type %T", pos[0], pos[1], left),
			)
		}
		if !ok {
			log.Fatal(
				fmt.Errorf("top of stack not an identifier: line %d col %d type %T", pos[0], pos[1], left),
			)
		}
		pairs[ident] = right.(ast.Expression)
	}
	return pairs
}

func pathToIdent(path string) string {
	s1 := strings.ReplaceAll(path, ".", "")
	s2 := strings.ReplaceAll(s1, "~", "")
	s3 := strings.ReplaceAll(s2, "\\", "")
	s4 := strings.ReplaceAll(s3, `"`, "")
	return strings.ReplaceAll(s4, "/", "")
}
