package listener

import (
	"fault/parser"
	"fmt"
	"regexp"

	"github.com/antlr4-go/antlr/v4"
)

var alphanumeric = regexp.MustCompile(`^[a-zA-Z0-9]+$`)

func validVarName(varname string) bool {
	return alphanumeric.MatchString(varname)
}

func assertValidVarName(varname string, token antlr.Token) {
	if !validVarName(varname) {
		panic(fmt.Sprintf("Variable names must be only letters or numbers: line %d col %d", token.GetLine(), token.GetColumn()))
	}
}

func (l *FaultListener) EnterSpecClause(c *parser.SpecClauseContext) {
	if l.currSpec == "" { //on import we may override the declared name
		l.currSpec = c.IDENT().GetText()
	}
	l.specs = append(l.specs, l.currSpec)
}

func (l *FaultListener) EnterSysClause(c *parser.SysClauseContext) {
	l.currSpec = c.IDENT().GetText()
}

func (l *FaultListener) EnterGlobalDecl(c *parser.GlobalDeclContext) {
	assertValidVarName(c.GetChild(1).(antlr.TerminalNode).GetText(), c.GetStart())
}

func (l *FaultListener) EnterStructDecl(c *parser.StructDeclContext) {
	l.scope = c.GetChild(1).(antlr.TerminalNode).GetText()
	assertValidVarName(l.scope, c.GetStart())
	l.structscope = l.scope
}

func (l *FaultListener) EnterComponentDecl(c *parser.ComponentDeclContext) {
	assertValidVarName(c.IDENT().GetText(), c.GetStart())
}

func (l *FaultListener) EnterStringDecl(c *parser.StringDeclContext) {
	assertValidVarName(c.IDENT().GetText(), c.GetStart())
}

func (l *FaultListener) EnterConstSpec(c *parser.ConstSpecContext) {
	identlist, ok := c.GetChild(0).(*parser.IdentListContext)
	if !ok {
		return
	}
	for _, name := range identlist.AllOperandName() {
		assertValidVarName(name.GetText(), c.GetStart())
	}
}

func (l *FaultListener) EnterStateFunc(c *parser.StateFuncContext) {
	varname := c.IDENT().GetText()
	assertValidVarName(varname, c.GetStart())
	l.scope = fmt.Sprint(l.scope, ".", varname)
}

func (l *FaultListener) EnterFunctionLit(c *parser.FunctionLitContext) {
	if c.Block().GetChildCount() < 3 {
		panic(fmt.Sprintf("Malformed fspec or fsystem file. A function cannot be empty: line %d col %d", c.GetStart().GetLine(), c.GetStart().GetColumn()))
	}
}

func (l *FaultListener) EnterStateBlock(c *parser.StateBlockContext) {
	if c.GetChildCount() < 3 {
		panic(fmt.Sprintf("Malformed fspec or fsystem file. A state function cannot be empty: line %d col %d", c.GetStart().GetLine(), c.GetStart().GetColumn()))
	}
}

func (l *FaultListener) EnterPropFunc(c *parser.PropFuncContext) {
	varname := c.IDENT().GetText()
	assertValidVarName(varname, c.GetStart())
	l.scope = fmt.Sprint(l.scope, ".", varname)
}

func (l *FaultListener) EnterPropInt(c *parser.PropIntContext) {
	assertValidVarName(c.IDENT().GetText(), c.GetStart())
}

func (l *FaultListener) EnterPropString(c *parser.PropStringContext) {
	assertValidVarName(c.IDENT().GetText(), c.GetStart())
}

func (l *FaultListener) EnterPropBool(c *parser.PropBoolContext) {
	assertValidVarName(c.IDENT().GetText(), c.GetStart())
}

func (l *FaultListener) EnterPropVar(c *parser.PropVarContext) {
	assertValidVarName(c.IDENT().GetText(), c.GetStart())
}

func (l *FaultListener) EnterPropSolvable(c *parser.PropSolvableContext) {
	// Skip the bare IDENT case (reference, not a declaration)
	if c.GetChildCount() == 1 {
		return
	}
	assertValidVarName(c.IDENT().GetText(), c.GetStart())
}

func (l *FaultListener) EnterRunInit(c *parser.RunInitContext) {
	// IDENT(0) is the variable being declared; IDENT(1) (if present) is the type reference
	assertValidVarName(c.IDENT(0).GetText(), c.GetStart())
}
