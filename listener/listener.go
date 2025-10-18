package listener

// From antlr parse tree create Fault AST

import (
	"fault/ast"
	"fault/parser"
	"fault/util"
	"fmt"
	"log"
	"os"
	gopath "path"
	"strconv"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	"github.com/barkimedes/go-deepcopy"
)

type FaultListener struct {
	*parser.BaseFaultParserListener
	stack                []ast.Node
	AST                  *ast.Spec
	structscope          string
	scope                string
	currSpec             string
	specs                []string
	skipRun              bool
	Path                 string // The location of the main spec
	testing              bool   // bypass imports when we're running unit tests
	Uncertains           map[string][]float64
	Unknowns             []string
	StructsPropertyOrder map[string][]string
	instances            map[string]*ast.Instance
	swaps                map[string][]ast.Node
}

func NewListener(path string, testing bool, skipRun bool) *FaultListener {
	return &FaultListener{
		Path:                 path,
		testing:              testing,
		skipRun:              skipRun,
		Uncertains:           make(map[string][]float64),
		StructsPropertyOrder: make(map[string][]string),
		instances:            make(map[string]*ast.Instance),
		swaps:                make(map[string][]ast.Node),
	}
}

func Execute(spec string, path string, flags map[string]bool /*specType bool, testing bool*/) *FaultListener {
	is := antlr.NewInputStream(spec)
	lexer := parser.NewFaultLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	p := parser.NewFaultParser(stream)
	l := NewListener(path, flags["testing"], flags["skipRun"])

	if flags["specType"] {
		antlr.ParseTreeWalkerDefault.Walk(l, p.Spec())
	} else {
		antlr.ParseTreeWalkerDefault.Walk(l, p.SysSpec())
	}
	return l
}

func (l *FaultListener) validate() {
	if l.testing { //will allow invalid specs during testing
		return
	}

	if len(l.stack) < 2 {
		fmt.Println(l.stack)
		fmt.Println("Malformed fspec or fsystem file. Too few statements.")
		os.Exit(1)
	}

	for _, v := range l.stack {

		if _, ok := v.(*ast.AssertionStatement); ok {
			return
		}

		if _, ok := v.(*ast.DefStatement); ok {
			return
		}

		if forS, ok := v.(*ast.ForStatement); ok {
			if len(forS.Inits.Statements) > 0 {
				return
			}
		}
	}

	fmt.Println("Malformed fspec or fsystem file. No model possible.")
	os.Exit(1)
}

func (l *FaultListener) push(n ast.Node) {
	l.stack = append(l.stack, n)
}

func (l *FaultListener) pushN(n []ast.Node) {
	l.stack = append(l.stack, n...)
}

func (l *FaultListener) pop() ast.Node {
	var s ast.Node
	s, l.stack = l.stack[len(l.stack)-1], l.stack[:len(l.stack)-1]
	return s
}

func (l *FaultListener) peek() ast.Node {
	return l.stack[len(l.stack)-1]
}

func (l *FaultListener) ExitSpec(c *parser.SpecContext) {
	var spec = &ast.Spec{}
	spec.Ext = "fspec"
	l.validate()
	for _, v := range l.stack {
		spec.Statements = append(spec.Statements, v.(ast.Statement))
	}
	l.addSwaps()
	l.AST = spec
}

func (l *FaultListener) EnterSpecClause(c *parser.SpecClauseContext) {
	if l.currSpec == "" { //on import we may override the declared name
		l.currSpec = c.IDENT().GetText()
	}
	l.specs = append(l.specs, l.currSpec)
}

func (l *FaultListener) ExitSpecClause(c *parser.SpecClauseContext) {
	token := ast.GenerateToken("SPEC_DECL", "SPEC_DECL", c.GetStart(), c.GetStop())

	iden_token := ast.GenerateToken("IDENT", "IDENT", c.GetStart(), c.GetStop())

	l.push(
		&ast.SpecDeclStatement{
			Token: token,
			Name: &ast.Identifier{
				Token: iden_token,
				Value: c.IDENT().GetText(),
				Spec:  l.currSpec,
			},
		},
	)
}

func (l *FaultListener) ExitImportDecl(c *parser.ImportDeclContext) {
	items := len(c.AllImportSpec())

	var itemList []ast.Node
	for i := 0; i < items; i++ {
		right := l.pop()

		var temp []ast.Node
		temp = append(temp, right)

		itemList = append(temp, itemList...)
	}

	for _, v := range itemList {
		l.push(v)
	}
}

func (l *FaultListener) ExitImportSpec(c *parser.ImportSpecContext) {
	token := ast.GenerateToken("IMPORT_DECL", "IMPORT_DECL", c.GetStart(), c.GetStop())

	val := l.pop()
	if val == nil {
		panic(fmt.Sprintf("top of stack not an expression: line %d col %d type %T", c.GetStart().GetLine(), c.GetStart().GetColumn(), val))
	}

	fpath, ok := val.(*ast.StringLiteral)
	if !ok {
		panic(fmt.Sprintf("import path not a string: line %d col %d type %T", c.GetStart().GetLine(), c.GetStart().GetColumn(), val))
	}

	// If no ident, create one from import path
	var importId string
	txt := c.GetText()
	if len(c.GetChildren()) > 2 {
		importId = c.IDENT().GetText()
	} else if len(c.GetChildren()) == 2 && string(txt[len(txt)-1]) != "," {
		importId = c.IDENT().GetText()
	} else {
		importId = pathToIdent(fpath.String())
	}

	var tree *ast.Spec
	if !l.testing {
		//Have we already imported this spec?
		if util.InStringSlice(l.specs, importId) {
			return
		}
		//Remove quotes
		trimmedFP := fpath.Value[1 : len(fpath.Value)-1]
		//Does file exist?
		fp := gopath.Join(l.Path, trimmedFP)
		fp = util.Filepath(fp)
		importFile, err := os.ReadFile(fp)
		if err != nil {
			panic(fmt.Sprintf("spec file %s not found\n", fpath))
		}
		tree = l.parseImport(importId, string(importFile))
	}

	ident := &ast.Identifier{
		Token: token,
		Value: pathToIdent(importId),
		Spec:  l.currSpec,
	}

	l.specs = append(l.specs, importId)

	l.push(&ast.ImportStatement{
		Token: token,
		Name:  ident,
		Path:  fpath,
		Tree:  tree,
	})
}

func (l *FaultListener) ExitConstSpec(c *parser.ConstSpecContext) {
	token := ast.GenerateToken("CONST_DECL", "CONST_DECL", c.GetStart(), c.GetStop())

	var items int
	identlist, ok := c.GetChild(0).(*parser.IdentListContext)
	if !ok {
		panic(fmt.Sprintf("can't find ident list: line %d col %d type %T", c.GetStart().GetLine(), c.GetStart().GetColumn(), c.GetChild(0)))
	}
	items = len(identlist.AllOperandName())

	var val ast.Node
	if (c.GetChildCount() - items) > 0 {
		val = l.pop()
		if val == nil {
			panic(fmt.Sprintf("top of stack not an expression: line %d col %d type %T", c.GetStart().GetLine(), c.GetStart().GetColumn(), val))
		}

	} else {
		token2 := ast.GenerateToken("UNKNOWN", "UNKNOWN", c.GetStart(), c.GetStop())
		val = &ast.Unknown{Token: token2, Name: nil}
	}

	var itemList []ast.Node
	for i := 0; i < items; i++ {
		left := l.pop()
		ident, ok := left.(*ast.Identifier)
		if !ok {
			panic(fmt.Sprintf("top of stack not an identifier: line %d col %d type %T", c.GetStart().GetLine(), c.GetStart().GetColumn(), left))
		}

		switch inst := val.(type) {
		case *ast.Unknown:
			inst.Name = ident
			val = inst
			l.Unknowns = append(l.Unknowns, strings.Join([]string{l.currSpec, ident.Value}, "_"))
		case *ast.Uncertain:
			l.Uncertains[strings.Join([]string{l.currSpec, ident.Value}, "_")] = []float64{inst.Mean, inst.Sigma}
		}
		var temp []ast.Node
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
	l.structscope = l.scope
}

func (l *FaultListener) ExitStructDecl(c *parser.StructDeclContext) {
	token2 := ast.GenerateToken("IDENT", "IDENT", c.GetStart(), c.GetStop())

	ident := &ast.Identifier{
		Token: token2,
		Value: c.IDENT().GetText(),
		Spec:  l.currSpec,
	}

	key := strings.Join([]string{ident.Spec, ident.Value}, "_")

	right := l.pop()
	var val ast.Expression
	var token ast.Token
	switch r := right.(type) {
	case *ast.StockLiteral:
		token = ast.GenerateToken("STOCK", "STOCK", c.GetStart(), c.GetStop())
		l.StructsPropertyOrder[key] = r.Order
		val = right.(ast.Expression)
	case *ast.FlowLiteral:
		token = ast.GenerateToken("FLOW", "FLOW", c.GetStart(), c.GetStop())
		l.StructsPropertyOrder[key] = r.Order
		val = right.(ast.Expression)
	default:
		if right == nil {
			panic(fmt.Sprintf("top of stack not an expression: line %d col %d type %T", c.GetStart().GetLine(), c.GetStart().GetColumn(), right))
		}
		panic(fmt.Sprintf("def can only be used to define a valid stock or flow: line %d col %d", c.GetStart().GetLine(), c.GetStart().GetColumn()))
	}

	l.push(
		&ast.DefStatement{
			Token: token,
			Name:  ident,
			Value: val,
		})
	l.scope = ""
	l.structscope = ""
}

func (l *FaultListener) ExitStock(c *parser.StockContext) {
	pairs := c.AllSfProperties()
	token := ast.GenerateToken("STOCK", "STOCK", c.GetStart(), c.GetStop())

	p, order := l.getPairs(len(pairs), []int{c.GetStart().GetLine(), c.GetStart().GetColumn()})

	l.push(
		&ast.StockLiteral{
			Token: token,
			Order: order,
			Pairs: p,
		})
}

func (l *FaultListener) ExitFlow(c *parser.FlowContext) {
	pairs := c.AllSfProperties()
	token := ast.GenerateToken("FLOW", "FLOW", c.GetStart(), c.GetStop())

	p, order := l.getPairs(len(pairs), []int{c.GetStart().GetLine(), c.GetStart().GetColumn()})
	l.push(
		&ast.FlowLiteral{
			Token: token,
			Order: order,
			Pairs: p,
		},
	)
}

func (l *FaultListener) ExitPropInt(c *parser.PropIntContext) {
	val := l.pop()
	token := ast.GenerateToken("IDENT", "IDENT", c.GetStart(), c.GetStop())

	l.push(&ast.Identifier{
		Token: token,
		Value: c.IDENT().GetText(),
		Spec:  l.currSpec,
	})

	l.push(val)
}

func (l *FaultListener) ExitPropBool(c *parser.PropBoolContext) {
	val := l.pop()
	token := ast.GenerateToken("IDENT", "IDENT", c.GetStart(), c.GetStop())

	l.push(&ast.Identifier{
		Token: token,
		Value: c.IDENT().GetText(),
		Spec:  l.currSpec,
	})

	l.push(val)
}

func (l *FaultListener) ExitPropString(c *parser.PropStringContext) {
	val := l.pop()

	token := ast.GenerateToken("IDENT", "IDENT", c.GetStart(), c.GetStop())

	l.push(&ast.Identifier{
		Token: token,
		Value: c.IDENT().GetText(),
		Spec:  l.currSpec,
	},
	)

	l.push(val)
}

func (l *FaultListener) ExitPropVar(c *parser.PropVarContext) {
	f := l.pop()

	token := ast.GenerateToken("IDENT", "IDENT", c.GetStart(), c.GetStop())

	l.push(&ast.Identifier{
		Token: token,
		Value: c.IDENT().GetText(),
		Spec:  l.currSpec,
	},
	)
	switch v := f.(type) {
	case *ast.Instance:
		v.Name = c.IDENT().GetText()
		l.push(v)
	case *ast.Identifier:
		l.push(v)
	case *ast.ParameterCall:
		l.push(v)
	case *ast.PrefixExpression:
		l.push(v)
	default:
		panic(fmt.Sprintf("top of stack not an identifier: line %d col %d type %T", c.GetStart().GetLine(), c.GetStart().GetColumn(), f))
	}
}

func (l *FaultListener) ExitPropSolvable(c *parser.PropSolvableContext) {
	var val ast.Node
	var keyValuePair bool
	if c.GetChildCount() != 1 {
		val = l.pop()
		keyValuePair = true
	} else {
		keyValuePair = false
	}
	token := ast.GenerateToken("IDENT", "IDENT", c.GetStart(), c.GetStop())

	ident := &ast.Identifier{
		Token: token,
		Value: c.IDENT().GetText(),
		Spec:  l.currSpec,
	}
	l.push(ident)

	if keyValuePair {
		unknown, ok := val.(*ast.Unknown)
		if ok {
			unknown.Name = ident
			l.push(unknown)
		} else {
			l.push(val)
		}
	} else {
		token2 := ast.GenerateToken("UNKNOWN", "UNKNOWN", c.GetStart(), c.GetStop())
		unknown := &ast.Unknown{Token: token2, Name: ident}
		l.push(unknown)
	}
}

func (l *FaultListener) EnterStateFunc(c *parser.StateFuncContext) {
	l.scope = fmt.Sprint(l.scope, ".", c.IDENT().GetText())
}

func (l *FaultListener) ExitStateFunc(c *parser.StateFuncContext) {
	val := l.pop()
	token := ast.GenerateToken("IDENT", "IDENT", c.GetStart(), c.GetStop())

	l.push(&ast.Identifier{
		Token: token,
		Value: c.IDENT().GetText(),
		Spec:  l.currSpec,
	},
	)
	l.push(val)

	scope := strings.Split(l.scope, ".")
	l.scope = strings.Join(scope[0:len(scope)-1], ".")

}

func (l *FaultListener) ExitStateLit(c *parser.StateLitContext) {
	token := ast.GenerateToken("FUNCTION", "FUNCTION", c.GetStart(), c.GetStop())

	b := l.pop()

	f := &ast.FunctionLiteral{
		Token: token,
		Body:  b.(*ast.BlockStatement),
	}
	l.push(f)
}

func (l *FaultListener) EnterPropFunc(c *parser.PropFuncContext) {
	l.scope = fmt.Sprint(l.scope, ".", c.IDENT().GetText())
}

func (l *FaultListener) ExitPropFunc(c *parser.PropFuncContext) {
	val := l.pop()
	token := ast.GenerateToken("IDENT", "IDENT", c.GetStart(), c.GetStop())

	l.push(&ast.Identifier{
		Token: token,
		Value: c.IDENT().GetText(),
		Spec:  l.currSpec,
	},
	)
	l.push(val)

	scope := strings.Split(l.scope, ".")
	l.scope = strings.Join(scope[0:len(scope)-1], ".")

}

func (l *FaultListener) ExitFunctionLit(c *parser.FunctionLitContext) {
	token := ast.GenerateToken("FUNCTION", "FUNCTION", c.GetStart(), c.GetStop())

	b := l.pop()

	f := &ast.FunctionLiteral{
		Token: token,
		Body:  b.(*ast.BlockStatement),
	}
	l.push(f)
}

func (l *FaultListener) ExitStatementList(c *parser.StatementListContext) {
	token := ast.GenerateToken("FUNCTION", "FUNCTION", c.GetStart(), c.GetStop())

	sl := &ast.BlockStatement{Token: token}
	for _, v := range c.GetChildren() {
		ex := l.pop()
		switch e := ex.(type) {
		case ast.Statement:
			sl.Statements = append([]ast.Statement{e}, sl.Statements...)
		case ast.Expression:
			token2 := ast.GenerateToken("FUNCTION", "FUNCTION", v.(*parser.StatementContext).GetStart(), v.(*parser.StatementContext).GetStop())

			s := &ast.ExpressionStatement{
				Token:      token2,
				Expression: e,
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
	token := ast.GenerateToken("ASSIGN", operator, c.GetStart(), c.GetStop())

	var valChange ast.Expression

	right := l.pop()
	left := l.pop()
	if operator == "->" {
		token2 := ast.GenerateToken("MINUS", "-", c.GetStart(), c.GetStop())
		valChange = &ast.InfixExpression{
			Token:    token2,
			Left:     left.(ast.Expression),
			Operator: "-",
			Right:    right.(ast.Expression)}
	} else if operator == "<-" {
		token2 := ast.GenerateToken("ADD", "+", c.GetStart(), c.GetStop())

		valChange = &ast.InfixExpression{
			Token:    token2,
			Left:     left.(ast.Expression),
			Operator: "+",
			Right:    right.(ast.Expression)}
	} else {
		panic(fmt.Sprintf("Invalid operator %s in expression", operator))
	}

	l.push(
		&ast.InfixExpression{
			Token:    token,
			Left:     left.(ast.Expression),
			Operator: "<-",
			Right:    valChange,
		})

}

func (l *FaultListener) ExitMiscAssign(c *parser.MiscAssignContext) {
	token := ast.GenerateToken("ASSIGN", c.GetChild(1).(antlr.TerminalNode).GetText(), c.GetStart(), c.GetStop())

	right := l.pop()
	if right == nil {
		panic(fmt.Sprintf("top of stack not an expression: line %d col %d type %T", c.GetStart().GetLine(), c.GetStart().GetColumn(), right))
	}

	var assign *ast.InfixExpression

	left := l.pop()
	switch ident := left.(type) {
	case *ast.Identifier: // This may not ever happen?
		// If a new instance is initialized in the run block
		// the listener needs to add the name
		switch inst := right.(type) {
		case *ast.Instance:
			inst.Name = ident.Value
			right = inst
		case *ast.Unknown:
			if inst.Name == nil {
				inst.Name = ident
				right = inst
			}
			l.Unknowns = append(l.Unknowns, strings.Join([]string{l.currSpec, l.scope, ident.Value}, "_"))
		case *ast.Uncertain:
			l.Uncertains[strings.Join([]string{l.currSpec, l.scope, ident.Value}, "_")] = []float64{inst.Mean, inst.Sigma}
		}

		assign = &ast.InfixExpression{
			Token:    token,
			Left:     ident,
			Operator: c.GetChild(1).(antlr.TerminalNode).GetText(),
			Right:    right.(ast.Expression),
		}
	case *ast.ParameterCall:
		switch inst := right.(type) {
		case *ast.Instance:
			inst.Name = ident.Value[0]
			right = inst
		}

		assign = &ast.InfixExpression{
			Token:    token,
			Left:     ident,
			Operator: c.GetChild(1).(antlr.TerminalNode).GetText(),
			Right:    right.(ast.Expression),
		}

	default:
		panic(fmt.Sprintf("left side of expression should be an identifier: line %d col %d type %T", c.GetStart().GetLine(), c.GetStart().GetColumn(), left))
	}

	l.push(assign)
}

func (l *FaultListener) ExitLrExpr(c *parser.LrExprContext) {
	token := ast.GenerateToken(string(ast.OPS[c.GetChild(1).(antlr.TerminalNode).GetText()]), c.GetChild(1).(antlr.TerminalNode).GetText(), c.GetStart(), c.GetStop())

	rght := l.pop()
	lft := l.pop()
	// If left is an empty Prefix, correct the parsing error
	if pre, ok := lft.(*ast.PrefixExpression); ok {
		if pre.Operator == "" {
			pre.Token = token
			pre.Operator = c.GetChild(1).(antlr.TerminalNode).GetText()
			pre.Right = rght.(ast.Expression)
			l.push(pre)
			return
		}
	}
	e := &ast.InfixExpression{
		Token:    token,
		Left:     lft.(ast.Expression),
		Operator: c.GetChild(1).(antlr.TerminalNode).GetText(),
		Right:    rght.(ast.Expression),
	}
	l.push(e)
}

func (l *FaultListener) ExitParamCall(c *parser.ParamCallContext) {
	token := ast.GenerateToken("IDENT", "IDENT", c.GetStart(), c.GetStop())

	v := c.GetText()
	param := strings.Split(v, ".")

	// This is a variable from another spec, not a call
	if util.InStringSlice(l.specs, param[0]) && len(param) == 2 {
		ident := &ast.Identifier{
			Token: token,
			Spec:  param[0],
			Value: param[1],
		}
		l.push(ident)
		return
	}

	pc := &ast.ParameterCall{
		Token: token,
		Value: param,
		Scope: l.structscope,
	}
	if util.InStringSlice(l.specs, param[0]) {
		pc.Spec = param[0]
	} else {
		pc.Spec = l.currSpec
	}

	l.push(pc)
}

func (l *FaultListener) ExitInitBlock(c *parser.InitBlockContext) {
	//var swaps, orphanSwaps []ast.Node
	var swaps []ast.Node
	token := ast.GenerateToken("FUNCTION", "FUNCTION", c.GetStart(), c.GetStop())

	sl := &ast.BlockStatement{
		Token: token,
	}
	swaps = l.getSwaps()

	steps := c.AllInitStep()
	for i := len(steps) - 1; i >= 0; i-- {
		ex := l.pop()

		if sw, ok := ex.(*ast.InfixExpression); ok && sw.TokenLiteral() == "SWAP" {
			i++
			swaps = append(swaps, sw)
			continue
		}

		if t, ok := ex.(*ast.Instance); ok {
			t.Swaps = append(t.Swaps, swaps...)

			token2 := ex.GetToken()

			s := &ast.ExpressionStatement{
				Token:      token2,
				Expression: t,
			}
			sl.Statements = append([]ast.Statement{s}, sl.Statements...)
		}
	}
	l.push(sl)
}

func (l *FaultListener) ExitRunBlock(c *parser.RunBlockContext) {
	token := ast.GenerateToken("FUNCTION", "FUNCTION", c.GetStart(), c.GetStop())

	sl := &ast.BlockStatement{
		Token: token,
	}

	steps := c.AllRunStep()
	for i := len(steps) - 1; i >= 0; i-- {
		ex := l.pop()

		switch t := ex.(type) {
		case *ast.ParallelFunctions:
			sl.Statements = append([]ast.Statement{t}, sl.Statements...)
		case ast.Expression:
			token2 := ex.GetToken()
			n := l.packageCallsAsRunSteps(t)

			s := &ast.ExpressionStatement{
				Token:      token2,
				Expression: n.(ast.Expression),
			}
			sl.Statements = append([]ast.Statement{s}, sl.Statements...)
		case *ast.BlockStatement:
			n := l.packageCallsAsRunSteps(t)
			t = n.(*ast.BlockStatement)
			sl.Statements = append(t.Statements, sl.Statements...)
		case *ast.ExpressionStatement:
			n := l.packageCallsAsRunSteps(t)
			t = n.(*ast.ExpressionStatement)
			sl.Statements = append([]ast.Statement{t}, sl.Statements...)
		default:
			panic(fmt.Sprintf("Neither statement nor expression got=%T", ex))
		}
	}
	l.push(sl)
}

func (l *FaultListener) ExitStateBlock(c *parser.StateBlockContext) {
	token := ast.GenerateToken("FUNCTION", "FUNCTION", c.GetStart(), c.GetStop())

	sl := &ast.BlockStatement{
		Token: token,
	}
	steps := c.AllStateStep()
	for i := len(steps) - 1; i >= 0; i-- {
		ex := l.pop()
		switch t := ex.(type) {
		case *ast.ParallelFunctions:
			sl.Statements = append([]ast.Statement{t}, sl.Statements...)
		case *ast.BlockStatement:
			n := l.packageCallsAsRunSteps(t)
			t = n.(*ast.BlockStatement)
			sl.Statements = append(t.Statements, sl.Statements...)
		case *ast.ExpressionStatement:
			n := l.packageCallsAsRunSteps(t)
			t = n.(*ast.ExpressionStatement)
			sl.Statements = append([]ast.Statement{t}, sl.Statements...)
		case *ast.BuiltIn:
			sl.Statements = append([]ast.Statement{&ast.ExpressionStatement{Expression: t}}, sl.Statements...)
		case *ast.InfixExpression:
			sl.Statements = append([]ast.Statement{&ast.ExpressionStatement{Expression: t}}, sl.Statements...)
		case *ast.PrefixExpression:
			sl.Statements = append([]ast.Statement{&ast.ExpressionStatement{Expression: t}}, sl.Statements...)

		default:
			panic(fmt.Sprintf("Neither statement nor expression got=%T", ex))
		}
	}
	l.push(sl)
}

func (l *FaultListener) ExitRunInit(c *parser.RunInitContext) {
	//var swaps, orphanSwaps []ast.Node
	var swaps []ast.Node
	txt := c.AllIDENT()
	var right string

	token2 := ast.GenerateToken("IDENT", "IDENT", c.GetStart(), c.GetStop())

	// Check for swaps
	swaps = l.getSwaps()

	ident := &ast.Identifier{Token: token2}
	switch len(txt) {
	case 1:
		pc := l.pop()
		switch r := pc.(type) {
		case *ast.Identifier:
			ident = r
			right = txt[0].GetText()
		case *ast.ParameterCall:
			ident.Spec = r.Value[0]
			ident.Value = r.Value[1]
			right = txt[0].GetText()
		default:
			panic(fmt.Sprintf("%s is an invalid identifier line: %d col:%d", txt, c.GetStart().GetLine(), c.GetStart().GetColumn()))
		}

	case 2:
		ident.Spec = l.currSpec
		ident.Value = txt[1].GetText()
		right = txt[0].GetText()
	case 3:
		ident.Spec = txt[1].GetText()
		ident.Value = txt[2].GetText() // Not sure why the parser flips the order
		right = txt[0].GetText()
	default:
		panic(fmt.Sprintf("%s is an invalid identifier line: %d col:%d", txt, c.GetStart().GetLine(), c.GetStart().GetColumn()))
	}

	key := strings.Join([]string{ident.Spec, ident.Value}, "_")
	order := l.StructsPropertyOrder[key]

	//swaps, orphanSwaps = l.filterSwaps(right, orphanSwaps)
	inst := &ast.Instance{
		Value: ident,
		Name:  right,
		Order: order,
	}

	l.instances[right] = inst
	l.sortSwaps(swaps)

	//l.pushN(orphanSwaps)
	l.push(inst)
}

func (l *FaultListener) ExitRunSwap(c *parser.SwapContext) {
	token := ast.GenerateToken("SWAP", "SWAP", c.GetStart(), c.GetStop())

	right := l.pop()
	left := l.pop()

	l.push(&ast.InfixExpression{
		Token:    token,
		Left:     left.(ast.Expression),
		Operator: "=",
		Right:    right.(ast.Expression),
	})
}

func (l *FaultListener) ExitRunStepExpr(c *parser.RunStepExprContext) {
	token := ast.GenerateToken("PARALLEL", c.GetText(), c.GetStart(), c.GetStop())

	var exp []ast.Expression
	for i := 0; i < len(c.AllParamCall()); i++ {
		idx := l.pop()
		exp = append([]ast.Expression{idx.(ast.Expression)}, exp...)
	}

	e := &ast.ParallelFunctions{
		Token:       token,
		Expressions: exp,
	}
	l.push(e)
}

func (l *FaultListener) ExitStateStepExpr(c *parser.StateStepExprContext) {
	token := ast.GenerateToken("PARALLEL", c.GetText(), c.GetStart(), c.GetStop())

	var exp []ast.Expression
	for i := 0; i < len(c.AllParamCall()); i++ {
		idx := l.pop()
		exp = append([]ast.Expression{idx.(ast.Expression)}, exp...)
	}

	e := &ast.ParallelFunctions{
		Token:       token,
		Expressions: exp,
	}
	l.push(e)
}

func (l *FaultListener) ExitRunExpr(c *parser.RunExprContext) {
	token := ast.GenerateToken("CODE", c.GetText(), c.GetStart(), c.GetStop())

	x := l.pop()
	exp, ok := x.(ast.Expression)
	if !ok {
		panic(fmt.Sprintf("top of stack is not a expression. got=%T", x))
	}

	e := &ast.ExpressionStatement{
		Token:      token,
		Expression: exp,
	}
	l.push(e)
}

func (l *FaultListener) ExitStateExpr(c *parser.StateExprContext) {
	token := ast.GenerateToken("CODE", c.GetText(), c.GetStart(), c.GetStop())

	x := l.pop()
	exp, ok := x.(ast.Expression)
	if !ok {
		panic(fmt.Sprintf("top of stack is not a expression. got=%T", x))
	}

	e := &ast.ExpressionStatement{
		Token:      token,
		Expression: exp,
	}
	l.push(e)
}

func (l *FaultListener) ExitPrefix(c *parser.PrefixContext) {
	if c.GetChild(0) == nil { //Bug in the grammar concerning
		// prefixes involving MINUS and idents
		e := &ast.PrefixExpression{}
		l.push(e)
		return
	}

	token := ast.GenerateToken(string(ast.OPS[c.GetChild(0).(antlr.TerminalNode).GetText()]), c.GetChild(0).(antlr.TerminalNode).GetText(), c.GetStart(), c.GetStop())

	rght := l.pop()
	e := &ast.PrefixExpression{
		Token:    token,
		Operator: c.GetChild(0).(antlr.TerminalNode).GetText(),
		Right:    rght.(ast.Expression),
	}
	l.push(e)
}

func (l *FaultListener) ExitSolvable(c *parser.SolvableContext) {
	switch c.FaultType().GetText() {
	case "natural":
		token := ast.GenerateToken("NATURAL", "NATURAL", c.GetStart(), c.GetStop())

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
		token := ast.GenerateToken("UNCERTAIN", "UNCERTAIN", c.GetStart(), c.GetStop())

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
	case "unknown":
		token := ast.GenerateToken("UNKNOWN", "UNKNOWN", c.GetStart(), c.GetStop())

		var ident *ast.Identifier
		if c.GetChildCount() > 3 {
			ident, _ = l.pop().(*ast.Identifier)
		}
		l.push(&ast.Unknown{
			Token: token,
			Name:  ident,
		})
	default:
		log.Fatalf("Unimplemented: %s", c.FaultType().GetText())
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

	token := ast.GenerateToken(string(tType), tLit, c.GetStart(), c.GetStop())

	ident := l.pop()

	token2 := ast.GenerateToken("INT", "INT", c.GetStart(), c.GetStop())

	e := &ast.InfixExpression{
		Token:    token,
		Left:     ident.(ast.Expression),
		Operator: tLit,
		Right:    &ast.IntegerLiteral{Token: token2, Value: 1},
	}
	l.push(e)
}

func (l *FaultListener) assembleIf(token ast.Token, children []antlr.Tree) *ast.IfExpression {
	var a *ast.BlockStatement
	var b *ast.IfExpression
	if len(children) > 3 {
		ra := l.pop()
		switch x := ra.(type) {
		case *ast.BlockStatement:
			a = x
		case *ast.IfExpression:
			b = x
		case *ast.ParallelFunctions:
			a = &ast.BlockStatement{Statements: []ast.Statement{
				x,
			}}
		default:
			panic(fmt.Sprintf("improper type in conditional got=%T", ra))
		}

	}
	csq := l.pop()
	cond := l.pop()

	//Delete this in a minute
	if _, ok := cond.(*ast.ParallelFunctions); ok {
		cond1 := csq
		csq = cond
		cond = cond1
	}

	if c, ok := cond.(ast.Operand); ok {
		//Modifying the construction if a {} to be more specific
		cond = &ast.InfixExpression{
			Left:     c,
			Operator: "==",
			Right:    &ast.Boolean{Value: true},
		}
	}

	if p, ok := csq.(*ast.ParallelFunctions); ok {
		csq = &ast.BlockStatement{Statements: []ast.Statement{
			p,
		}}
	}

	e := &ast.IfExpression{
		Token:       token,
		Condition:   cond.(ast.Expression),
		Consequence: csq.(*ast.BlockStatement),
		Alternative: a,
		Elif:        b,
	}
	return e
}

func (l *FaultListener) ExitIfStmt(c *parser.IfStmtContext) {
	token := ast.GenerateToken("IF", "IF", c.GetStart(), c.GetStop())

	e := l.assembleIf(token, c.GetChildren())

	l.push(e)
}

func (l *FaultListener) ExitIfStmtState(c *parser.IfStmtStateContext) {
	token := ast.GenerateToken("IF", "IF", c.GetStart(), c.GetStop())
	e := l.assembleIf(token, c.GetChildren())
	l.push(e)
}

func (l *FaultListener) ExitIfStmtRun(c *parser.IfStmtRunContext) {
	token := ast.GenerateToken("IF", "IF", c.GetStart(), c.GetStop())
	e := l.assembleIf(token, c.GetChildren())
	l.push(e)
}

func (l *FaultListener) ExitAccessHistory(c *parser.AccessHistoryContext) {
	token := ast.GenerateToken("HISTORY", "HISTORY", c.GetStart(), c.GetStop())

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
	token := ast.GenerateToken("IDENT", "IDENT", c.GetStart(), c.GetStop())

	l.push(&ast.Identifier{
		Token: token,
		Value: c.GetText(),
		Spec:  l.currSpec,
	},
	)
}

func (l *FaultListener) ExitOpInstance(c *parser.OpInstanceContext) {
	ident := &ast.Identifier{}
	id := c.AllIDENT()
	switch len(id) {
	case 1:
		ident.Spec = l.currSpec
		ident.Value = id[0].GetText()
	case 2:
		ident.Spec = id[0].GetText()
		ident.Value = id[1].GetText()
	default:
		panic(fmt.Sprintf("%s is an invalid identifier line: %d col:%d", id, c.GetStart().GetLine(), c.GetStart().GetColumn()))
	}

	key := strings.Join([]string{ident.Spec, ident.Value}, "_")
	order := l.StructsPropertyOrder[key]

	l.push(&ast.Instance{
		Value: ident,
		Order: order,
	},
	)
}

func (l *FaultListener) ExitOpThis(c *parser.OpThisContext) {
	token := ast.GenerateToken("IDENT", "IDENT", c.GetStart(), c.GetStop())

	l.push(&ast.This{
		Token: token,
		Value: strings.Split(l.scope, "."),
	},
	)
}

func (l *FaultListener) ExitOpClock(c *parser.OpClockContext) {
	token := ast.GenerateToken("IDENT", "IDENT", c.GetStart(), c.GetStop())
	l.push(&ast.Clock{
		Token: token,
		Value: l.scope,
	},
	)
}

func (l *FaultListener) ExitNil(c *parser.NilContext) {
	token := ast.GenerateToken("NIL", "NIL", c.GetStart(), c.GetStop())

	l.push(&ast.Nil{
		Token: token,
	})
}

func (l *FaultListener) ExitInteger(c *parser.IntegerContext) {
	token := ast.GenerateToken("INT", "INT", c.GetStart(), c.GetStop())

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

	switch i := base.(type) {
	case *ast.IntegerLiteral:
		token := ast.GenerateToken("INT", "INT", c.GetStart(), c.GetStop())
		i.Token = token
		i.Value = -i.Value

		l.push(i)

	case *ast.FloatLiteral:
		token := ast.GenerateToken("FLOAT", "FLOAT", c.GetStart(), c.GetStop())

		i.Token = token
		i.Value = -i.Value

		l.push(i)
	case *ast.Identifier:
		token := ast.GenerateToken("IDENT", "IDENT", c.GetStart(), c.GetStop())

		e := &ast.PrefixExpression{
			Token:    token,
			Operator: "-",
			Right:    i,
		}
		l.push(e)

	default:
		panic(fmt.Sprintf("top of stack not an integer or a float got=%T", base))
	}

}

func (l *FaultListener) ExitFloat_(c *parser.Float_Context) {
	token := ast.GenerateToken("FLOAT", "FLOAT", c.GetStart(), c.GetStop())

	v, err := strconv.ParseFloat(c.GetText(), 64)
	if err != nil {
		panic(fmt.Sprintf("float value detected but not parsable: line %d col %d got=%s", c.GetStart().GetLine(), c.GetStart().GetColumn(), c.GetText()))
	}

	l.push(&ast.FloatLiteral{
		Token: token,
		Value: v,
	})
}

func (l *FaultListener) ExitCompoundString(c *parser.CompoundStringContext) {
	if c.GetChildCount() < 2 {
		return //Single operand, no actions
	}

	if pre, ok := c.GetChild(0).(antlr.TerminalNode); ok && pre.GetText() == "!" {
		token := ast.GenerateToken("IDENT", "IDENT", c.GetStart(), c.GetStop())
		exp := l.pop()
		e := &ast.PrefixExpression{
			Token:    token,
			Operator: "-",
			Right:    exp.(ast.Expression),
		}
		l.push(e)
		return
	}

	if op, ok := c.GetChild(1).(antlr.TerminalNode); ok {
		operator := op.GetText()
		token := ast.GenerateToken(string(ast.OPS[operator]), operator, c.GetStart(), c.GetStop())

		rght := l.pop()
		lft := l.pop()
		e := &ast.InfixExpression{
			Token:    token,
			Left:     lft.(ast.Expression),
			Operator: operator,
			Right:    rght.(ast.Expression),
		}
		l.push(e)
	}
}

func (l *FaultListener) ExitStringDecl(c *parser.StringDeclContext) {
	token := ast.GenerateToken("GLOBAL", "GLOBAL", c.GetStart(), c.GetStop())

	val := l.pop()
	token2 := ast.GenerateToken("IDENT", "IDENT", c.GetStart(), c.GetStop())

	ident := &ast.Identifier{
		Token: token2,
		Value: c.IDENT().GetText(),
		Spec:  l.currSpec,
	}
	switch val.(type) {
	case *ast.StringLiteral:
		l.push(&ast.DefStatement{Token: token, Name: ident, Value: val.(ast.Expression)})
	case *ast.InfixExpression:
		token2 := ast.GenerateToken("COMPOUND_STRING", "COMPOUND_STRING", c.GetStart(), c.GetStop())
		val.(*ast.InfixExpression).Token = token2
		l.push(&ast.DefStatement{Token: token, Name: ident, Value: val.(ast.Expression)})
	case *ast.PrefixExpression:
		token2 := ast.GenerateToken("COMPOUND_STRING", "COMPOUND_STRING", c.GetStart(), c.GetStop())
		val.(*ast.PrefixExpression).Token = token2
		l.push(&ast.DefStatement{Token: token, Name: ident, Value: val.(ast.Expression)})
	default:
		panic(fmt.Sprintf("top of the stack is not a string got %T: line %d col %d", val, c.GetStart().GetLine(), c.GetStart().GetColumn()))
	}
}

func (l *FaultListener) ExitString_(c *parser.String_Context) {
	token := ast.GenerateToken("STRING", "STRING", c.GetStart(), c.GetStop())

	v := c.GetText()

	l.push(&ast.StringLiteral{
		Token: token,
		Value: v,
	})
}

func (l *FaultListener) ExitBool_(c *parser.Bool_Context) {
	token := ast.GenerateToken("BOOL", "BOOL", c.GetStart(), c.GetStop())

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
		token := ast.GenerateToken("FUNCTION", "FUNCTION", c.GetStart(), c.GetStop())

		l.push(&ast.BlockStatement{
			Token: token,
		})
	}
}

func (l *FaultListener) ExitInitDecl(c *parser.InitDeclContext) {
	token := ast.GenerateToken("ASSIGN", "init", c.GetStart(), c.GetStop())

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
	token := ast.GenerateToken("FOR", "for", c.GetStart(), c.GetStop())

	run := l.pop()
	init := l.pop()
	var rounds *ast.IntegerLiteral
	var block2 *ast.BlockStatement

	if run == nil {
		panic(fmt.Sprintf("top of stack not an expression: line %d col %d type %T", c.GetStart().GetLine(), c.GetStart().GetColumn(), run))
	}

	block, ok := run.(*ast.BlockStatement)
	if !ok {
		panic(fmt.Sprintf("top of stack not a block statement: line %d col %d type %T", c.GetStart().GetLine(), c.GetStart().GetColumn(), run))
	}

	if init == nil {
		panic(fmt.Sprintf("top of stack not an expression: line %d col %d type %T", c.GetStart().GetLine(), c.GetStart().GetColumn(), init))
	}

	switch x := init.(type) {
	case *ast.BlockStatement:
		num := l.pop()
		rounds, ok = num.(*ast.IntegerLiteral)
		if !ok {
			panic(fmt.Sprintf("top of stack not an integer literal: line %d col %d type %T", c.GetStart().GetLine(), c.GetStart().GetColumn(), num))
		}
		block2 = x

	case *ast.IntegerLiteral:
		rounds = x
		block2 = &ast.BlockStatement{}

	default:
		panic(fmt.Sprintf("top of stack not a block statement or integer: line %d col %d type %T", c.GetStart().GetLine(), c.GetStart().GetColumn(), init))
	}

	forSt := &ast.ForStatement{
		Token:  token,
		Rounds: rounds,
		Body:   block,
		Inits:  block2,
	}

	if !l.skipRun {
		l.push(forSt)
	}
}

func (l *FaultListener) ExitDefInvariant(c *parser.DefInvariantContext) {
	right := l.pop()
	left := l.pop()
	token := ast.GenerateToken("ASSERT", "assert", c.GetStart(), c.GetStop())
	l.push(&ast.InvariantClause{
		Token:    token,
		Left:     left.(ast.Expression),
		Operator: "==",
		Right:    right.(ast.Expression),
	})
}

func (l *FaultListener) ExitStageInvariant(c *parser.StageInvariantContext) {
	right := l.pop()
	left := l.pop()
	token := ast.GenerateToken("ASSERT", "assert", c.GetStart(), c.GetStop())
	l.push(&ast.InvariantClause{
		Token:    token,
		Left:     left.(ast.Expression),
		Operator: "then",
		Right:    right.(ast.Expression),
	})
}

func (l *FaultListener) ExitAssertion(c *parser.AssertionContext) {
	token := ast.GenerateToken("ASSERT", "assert", c.GetStart(), c.GetStop())

	var temporal string
	var temporalFilter string
	var temporalN int
	if c.Temporal() != nil {
		if c.Temporal().GetChildCount() == 2 {
			i := l.pop()
			temporalRaw := c.Temporal().GetText()
			temporal = ""
			temporalFilter = temporalRaw[0 : len(temporalRaw)-len(i.(*ast.IntegerLiteral).String())]
			temporalN = int(i.(*ast.IntegerLiteral).Value)
		} else {
			temporal = c.Temporal().GetText()
		}
	}

	expr := l.pop()
	var con *ast.InvariantClause
	switch e := expr.(type) {
	default:
		panic(fmt.Sprintf("invariant unusable. Must be expression not %T line: %d, col: %d", e, c.GetStart().GetLine(), c.GetStart().GetColumn()))
	case *ast.IntegerLiteral:
		// Disregard, this is part of the temporal filter
	case *ast.ParameterCall:
		con = &ast.InvariantClause{
			Token:    e.Token,
			Left:     e,
			Operator: "==",
			Right:    &ast.Boolean{Value: true},
		}
	case *ast.Identifier:
		con = &ast.InvariantClause{
			Token:    e.Token,
			Left:     e,
			Operator: "==",
			Right:    &ast.Boolean{Value: true},
		}
	case *ast.PrefixExpression:
		if e.Operator == "!" {
			e.Operator = "!="
			con = &ast.InvariantClause{
				Token:    e.Token,
				Left:     e.Right,
				Operator: e.Operator,
				Right:    &ast.Boolean{Value: true},
			}
		} else {
			panic("illegal prefix operator in assertion")
		}
	case *ast.InfixExpression:

		con = &ast.InvariantClause{
			Token:    e.Token,
			Left:     e.Left,
			Operator: e.Operator,
			Right:    e.Right,
		}
	case *ast.InvariantClause:
		con = e
	case *ast.IndexExpression:
		con = &ast.InvariantClause{
			Token:    e.Token,
			Left:     e,
			Operator: "==",
			Right:    &ast.Boolean{Value: true},
		}
	}

	l.push(&ast.AssertionStatement{
		Token:          token,
		Constraint:     con,
		Temporal:       temporal,
		TemporalFilter: temporalFilter,
		TemporalN:      temporalN,
		Assume:         false,
	})
}

func (l *FaultListener) ExitAssumption(c *parser.AssumptionContext) {
	token := ast.GenerateToken("ASSUME", "assume", c.GetStart(), c.GetStop())
	var temporal string
	var temporalFilter string
	var temporalN int
	if c.Temporal() != nil {
		if c.Temporal().GetChildCount() == 2 {
			i := l.pop()
			temporalRaw := c.Temporal().GetText()
			temporal = ""
			temporalFilter = temporalRaw[0 : len(temporalRaw)-len(i.(*ast.IntegerLiteral).String())]
			temporalN = int(i.(*ast.IntegerLiteral).Value)
		} else {
			temporal = c.Temporal().GetText()
		}
	}

	expr := l.pop()
	var con *ast.InvariantClause
	switch e := expr.(type) {
	default:
		panic(fmt.Sprintf("invariant unusable. Must be expression not %T line: %d, col: %d", e, c.GetStart().GetLine(), c.GetStart().GetColumn()))
	case *ast.ParameterCall:
		con = &ast.InvariantClause{
			Token:    e.Token,
			Left:     e,
			Operator: "==",
			Right:    &ast.Boolean{Value: true},
		}
	case *ast.Identifier:
		con = &ast.InvariantClause{
			Token:    e.Token,
			Left:     e,
			Operator: "==",
			Right:    &ast.Boolean{Value: true},
		}
	case *ast.PrefixExpression:
		if e.Operator == "!" {
			e.Operator = "!="
			con = &ast.InvariantClause{
				Token:    e.Token,
				Left:     e.Right,
				Operator: e.Operator,
				Right:    &ast.Boolean{Value: true},
			}
		} else {
			panic("illegal prefix operator in assumption")
		}
	case *ast.InfixExpression:
		if e.Operator == "!=" {
			con = &ast.InvariantClause{
				Token:    e.Token,
				Left:     &ast.Boolean{Value: true},
				Operator: "!=",
				Right:    e,
			}
		} else {
			con = &ast.InvariantClause{
				Token:    e.Token,
				Left:     e.Left,
				Operator: e.Operator,
				Right:    e.Right,
			}
		}
	case *ast.InvariantClause:
		con = e
	}

	l.push(&ast.AssertionStatement{
		Token:          token,
		Constraint:     con,
		Temporal:       temporal,
		TemporalFilter: temporalFilter,
		TemporalN:      temporalN,
		Assume:         true,
	})
}

func (l *FaultListener) parseImport(id string, spec string) *ast.Spec {
	is := antlr.NewInputStream(spec)
	lexer := parser.NewFaultLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	p := parser.NewFaultParser(stream)
	listener := NewListener("", false, true)
	listener.currSpec = id
	listener.specs = l.specs
	listener.Path = l.Path
	antlr.ParseTreeWalkerDefault.Walk(listener, p.Spec())

	l.Uncertains, l.Unknowns, l.StructsPropertyOrder, l.specs = mergeListeners(l, listener)
	return listener.AST
}

func mergeListeners(l1 *FaultListener, l2 *FaultListener) (map[string][]float64, []string, map[string][]string, []string) {
	for k, v := range l2.Uncertains {
		l1.Uncertains[k] = v
	}

	l1.Unknowns = append(l1.Unknowns, l2.Unknowns...)
	l1.specs = append(l1.specs, l2.specs...)

	for k, v := range l2.StructsPropertyOrder {
		l1.StructsPropertyOrder[k] = v
	}
	return l1.Uncertains, l1.Unknowns, l1.StructsPropertyOrder, l1.specs
}

func (l *FaultListener) getPairs(p int, pos []int) (map[*ast.Identifier]ast.Expression, []string) {
	var order []string
	pairs := make(map[*ast.Identifier]ast.Expression)
	for i := 0; i < p; i++ {
		right := l.pop()
		if right == nil {
			panic(fmt.Sprintf("top of stack not an expression: line %d col %d type %T", pos[0], pos[1], right))
		}

		left := l.pop()
		if left == nil {
			panic(fmt.Sprintf("top of stack not an expression: line %d col %d type %T", pos[0], pos[1], left))
		}

		ident, ok := left.(*ast.Identifier)
		if !ok {
			panic(fmt.Sprintf("top of stack not an identifier: line %d col %d type %T", pos[0], pos[1], left))
		}

		switch inst := right.(type) {
		case *ast.Unknown:
			inst.Name = ident
			right = inst
			l.Unknowns = append(l.Unknowns, strings.Join([]string{l.currSpec, l.scope, ident.Value}, "_"))
		case *ast.Uncertain:
			l.Uncertains[strings.Join([]string{l.currSpec, l.scope, ident.Value}, "_")] = []float64{inst.Mean, inst.Sigma}
		}
		order = append([]string{ident.Value}, order...)
		pairs[ident] = right.(ast.Expression)
	}
	return pairs, order
}

func (l *FaultListener) componentPairs(pairs map[*ast.Identifier]ast.Expression) map[*ast.Identifier]ast.Expression {
	p := make(map[*ast.Identifier]ast.Expression)
	for k, v := range pairs {
		switch f := v.(type) {
		case *ast.FunctionLiteral:
			// If the only thing inside is a stay();
			// move on.
			var bi *ast.BuiltIn
			if es, ok := f.Body.Statements[0].(*ast.ExpressionStatement); ok {
				bi, _ = es.Expression.(*ast.BuiltIn)
			}
			if len(f.Body.Statements) == 1 && bi != nil && l.builtInType(bi) == "stay" {
				p[k] = v
				continue
			}

			//Wrap inner function in conditional so that only
			// executes if the state is active
			this := &ast.ParameterCall{Spec: k.Spec, Value: []string{"this", k.Value}}
			cond := &ast.InfixExpression{Left: this, Operator: "==", Right: &ast.Boolean{Value: true}}
			con := &ast.IfExpression{Condition: cond, Consequence: f.Body}
			f.Body = &ast.BlockStatement{Statements: []ast.Statement{&ast.ExpressionStatement{Expression: con}}}
			p[k] = f
		default:
			p[k] = v
		}
	}
	return p
}

func (l *FaultListener) intOrFloatOk(v ast.Node) (float64, error) {
	switch val := v.(type) {
	case *ast.FloatLiteral:
		return val.Value, nil
	case *ast.IntegerLiteral:
		return float64(val.Value), nil
	default:
		return 0, fmt.Errorf("invalid input type. Should be float or int got=%T", v)
	}
}

func pathToIdent(path string) string {
	base := strings.Split(path, string(os.PathSeparator))

	s1 := strings.ReplaceAll(base[len(base)-1], ".fspec", "")
	s2 := strings.ReplaceAll(s1, "~", "")
	s3 := strings.ReplaceAll(s2, "\\", "")
	s4 := strings.ReplaceAll(s3, `"`, "")
	s5 := strings.ReplaceAll(s4, ".", "")
	return strings.ReplaceAll(s5, "/", "")
}

//////////////////////////////////////////////
//  State Charts
//////////////////////////////////////////////

func (l *FaultListener) ExitSysSpec(c *parser.SysSpecContext) {
	var spec = &ast.Spec{}
	spec.Ext = "fsystem"
	l.validate()
	for _, v := range l.stack {
		spec.Statements = append(spec.Statements, v.(ast.Statement))
	}
	l.addSwaps()
	l.AST = spec
}

func (l *FaultListener) EnterSysClause(c *parser.SysClauseContext) {
	l.currSpec = c.IDENT().GetText()
}

func (l *FaultListener) ExitSysClause(c *parser.SysClauseContext) {
	token := ast.GenerateToken("SYS_DECL", "SYS_DECL", c.GetStart(), c.GetStop())

	iden_token := ast.GenerateToken("IDENT", "IDENT", c.GetStart(), c.GetStop())

	l.push(
		&ast.SysDeclStatement{
			Token: token,
			Name: &ast.Identifier{
				Token: iden_token,
				Value: c.IDENT().GetText(),
				Spec:  l.currSpec,
			},
		},
	)
}

func (l *FaultListener) getSwaps() []ast.Node {
	var swaps []ast.Node
	loop := true

	for loop {
		peek := l.peek()
		if swap, ok := peek.(*ast.InfixExpression); ok && swap.TokenLiteral() == "SWAP" {
			l.pop()
			swaps = append(swaps, swap)
		} else {
			loop = false
		}
	}
	return swaps
}

func (l *FaultListener) filterSwaps(id string, swaps []ast.Node) ([]ast.Node, []ast.Node) {
	var filtered, orphaned []ast.Node
	for _, s := range swaps {
		if infx, ok := s.(*ast.InfixExpression); ok {
			var id2 string
			switch n := infx.Left.(type) {
			case *ast.ParameterCall:
				id2 = n.Value[0]
			default:
				panic(fmt.Sprintf("malformed swap got=%s", infx.String()))
			}

			if id == id2 {
				filtered = append(filtered, s)
			} else {
				orphaned = append(orphaned, s)
			}

		}
	}
	return filtered, orphaned
}

func (l *FaultListener) ExitGlobalDecl(c *parser.GlobalDeclContext) {
	var swaps, orphanSwaps []ast.Node

	token := ast.GenerateToken("GLOBAL", "GLOBAL", c.GetStart(), c.GetStop())

	orphanSwaps = l.getSwaps()

	instance := l.pop()

	token2 := ast.GenerateToken("IDENT", "IDENT", c.GetStart(), c.GetStop())

	ident := &ast.Identifier{
		Token: token2,
		Value: c.IDENT().GetText(),
		Spec:  l.currSpec,
	}

	var importSpec string
	var importStruct string
	switch parent := instance.(type) {
	case *ast.Instance:
		swaps, orphanSwaps = l.filterSwaps(ident.Value, orphanSwaps)
		importSpec = parent.Value.Spec
		importStruct = parent.Value.Value
		parent.Name = ident.Value
		key := strings.Join([]string{importSpec, importStruct}, "_")
		order := l.StructsPropertyOrder[key]
		parent.Order = order

		l.instances[parent.Name] = parent
		l.sortSwaps(swaps)
	}

	l.push(&ast.DefStatement{
		Token: token,
		Name:  ident,
		Value: instance.(ast.Expression),
	})

	l.pushN(orphanSwaps)

}

func (l *FaultListener) ExitSwap(c *parser.SwapContext) {
	token := ast.GenerateToken("SWAP", "SWAP", c.GetStart(), c.GetStop())

	right := l.pop()
	left := l.pop()

	l.push(&ast.InfixExpression{
		Token:    token,
		Left:     left.(ast.Expression),
		Operator: "=",
		Right:    right.(ast.Expression),
	})
}

func (l *FaultListener) ExitComponentDecl(c *parser.ComponentDeclContext) {
	pairs := c.AllComProperties()
	token := ast.GenerateToken("COMPONENT", "COMPONENT", c.GetStart(), c.GetStop())

	p, order := l.getPairs(len(pairs), []int{c.GetStart().GetLine(), c.GetStart().GetColumn()})

	p2 := l.componentPairs(p)

	val :=
		&ast.ComponentLiteral{
			Token: token,
			Order: order,
			Pairs: p2,
		}

	token2 := ast.GenerateToken("IDENT", "IDENT", c.GetStart(), c.GetStop())

	ident := &ast.Identifier{
		Token: token2,
		Value: c.IDENT().GetText(),
		Spec:  l.currSpec,
	}

	l.push(
		&ast.DefStatement{
			Token: token,
			Name:  ident,
			Value: val,
		})
}

func (l *FaultListener) ExitBuiltins(c *parser.BuiltinsContext) {
	token := ast.GenerateToken("BUILTIN", "BUILTIN", c.GetStart(), c.GetStop())

	f := &ast.BuiltIn{
		Token: token,
	}

	f.Function = c.GetChild(0).(antlr.TerminalNode).GetText()
	f.Parameters = make(map[string]ast.Operand)

	if f.Function == "advance" {
		p := l.pop()
		f.Parameters["toState"] = p.(ast.Operand)
	}

	l.push(f)
}

func (l *FaultListener) ExitBoolCompound(c *parser.BoolCompoundContext) {
	token := ast.GenerateToken(string(ast.OPS[c.GetChild(1).(antlr.TerminalNode).GetText()]), c.GetChild(1).(antlr.TerminalNode).GetText(), c.GetStart(), c.GetStop())

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

func (l *FaultListener) ExitBuiltinInfix(c *parser.BuiltinInfixContext) {
	infix := l.pop().(*ast.InfixExpression)
	switch c.GetChild(0).(type) {
	case *antlr.TerminalNodeImpl:
		e := &ast.PrefixExpression{
			Token:    infix.Token,
			Operator: "choose",
			Right:    infix,
		}
		l.push(e)
	case *parser.BoolCompoundContext:
		l.push(infix)
	default:
		panic("unknown child type in builtin infix")
	}

}

func (l *FaultListener) ExitStartPair(c *parser.StartPairContext) {
	idents := c.AllIDENT()
	start := &ast.InfixExpression{
		Left:     &ast.StringLiteral{Value: idents[0].GetText()},
		Operator: ":",
		Right:    &ast.StringLiteral{Value: idents[1].GetText()},
	}
	l.push(start)

}

func (l *FaultListener) ExitStartBlock(c *parser.StartBlockContext) {
	token := ast.GenerateToken("START", "START", c.GetStart(), c.GetStop())
	var pairs [][]string
	for i := 0; i < len(c.AllStartPair()); i++ {
		p := l.pop()
		pair := p.(*ast.InfixExpression)
		pairs = append(pairs, []string{pair.Left.String(), pair.Right.String()})
	}

	l.push(&ast.StartStatement{Token: token, Pairs: pairs})
}

func (l *FaultListener) packageCallsAsRunSteps(node ast.Node) ast.Node {
	switch n := node.(type) {
	case *ast.BlockStatement:
		st := []ast.Statement{}
		for _, s := range n.Statements {
			t := l.packageCallsAsRunSteps(s)
			st = append(st, t.(ast.Statement))
		}
		n.Statements = st
		return n
	case *ast.ExpressionStatement:
		e := l.packageCallsAsRunSteps(n.Expression)
		switch e.(type) {
		case ast.Statement:
			return e
		default:
			n.Expression = e.(ast.Expression)
			return n
		}
	case *ast.IfExpression:
		con := l.packageCallsAsRunSteps(n.Consequence)
		n.Consequence = con.(*ast.BlockStatement)

		if n.Alternative != nil {
			alt := l.packageCallsAsRunSteps(n.Alternative)
			n.Alternative = alt.(*ast.BlockStatement)
		}
		if n.Elif != nil {
			el := l.packageCallsAsRunSteps(n.Elif)
			n.Elif = el.(*ast.IfExpression)
		}
		return n
	case *ast.ParameterCall:
		return &ast.ParallelFunctions{
			Token:        n.Token,
			InferredType: n.InferredType,
			Expressions:  []ast.Expression{n},
		}

	case *ast.InfixExpression:
		left := l.packageCallsAsRunSteps(n.Left)
		pfLeft, okl := left.(*ast.ParallelFunctions)
		right := l.packageCallsAsRunSteps(n.Right)
		pfRight, okr := right.(*ast.ParallelFunctions)
		if n.Operator == "|" && okl && okr {
			return &ast.ParallelFunctions{
				Token:        n.Token,
				InferredType: n.InferredType,
				Expressions:  append(pfLeft.Expressions, pfRight.Expressions...),
			}
		}
		if n.Operator == "|" && okl {
			return &ast.ParallelFunctions{
				Token:        n.Token,
				InferredType: n.InferredType,
				Expressions:  append(pfLeft.Expressions, right.(ast.Expression)),
			}
		}

		if n.Operator == "|" && okr {
			return &ast.ParallelFunctions{
				Token:        n.Token,
				InferredType: n.InferredType,
				Expressions:  append(pfRight.Expressions, left.(ast.Expression)),
			}
		}

		if n.Operator == "|" {
			return &ast.ParallelFunctions{
				Token:        n.Token,
				InferredType: n.InferredType,
				Expressions:  []ast.Expression{left.(ast.Expression), right.(ast.Expression)},
			}
		}
		return n
	default:
		return n
	}
}

func (l *FaultListener) sortSwaps(swaps []ast.Node) {
	for _, s := range swaps {
		if node, ok := s.(*ast.InfixExpression).Left.(*ast.ParameterCall); ok {
			if _, ok2 := l.swaps[node.Value[0]]; !ok2 {
				l.swaps[node.Value[0]] = []ast.Node{}
			}
			l.swaps[node.Value[0]] = append(l.swaps[node.Value[0]], s)
		}
	}
}

func (l *FaultListener) addSwaps() {
	for key, inst := range l.instances {
		if sw, ok := l.swaps[key]; ok {
			c, err := deepcopy.Anything(sw)
			if err != nil {
				panic(err)
			}
			inst.Swaps = c.([]ast.Node)
			l.swaps[key] = []ast.Node{}
		}
	}
}

func (l *FaultListener) builtInType(b *ast.BuiltIn) string {
	return b.Function
}
