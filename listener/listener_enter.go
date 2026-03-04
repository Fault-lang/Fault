package listener

import (
	"fault/parser"
	"fmt"
	"strings"

	"github.com/antlr4-go/antlr/v4"
)

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
	varname := c.GetChild(1).(antlr.TerminalNode).GetText()
	if strings.Contains(varname, "_") {
		panic(fmt.Sprintf("Variable names may not have underscores: line %d col %d", c.GetStart().GetLine(), c.GetStart().GetColumn()))
	}
}

func (l *FaultListener) EnterStructDecl(c *parser.StructDeclContext) {
	l.scope = c.GetChild(1).(antlr.TerminalNode).GetText()
	if strings.Contains(l.scope, "_") {
		panic(fmt.Sprintf("Variable names may not have underscores: line %d col %d", c.GetStart().GetLine(), c.GetStart().GetColumn()))
	}

	l.structscope = l.scope
}

func (l *FaultListener) EnterComponentDecl(c *parser.ComponentDeclContext) {
	varname := c.IDENT().GetText()
	if strings.Contains(varname, "_") {
		panic(fmt.Sprintf("Variable names may not have underscores: line %d col %d", c.GetStart().GetLine(), c.GetStart().GetColumn()))
	}
}

func (l *FaultListener) EnterStringDecl(c *parser.StringDeclContext) {
	varname := c.IDENT().GetText()
	if strings.Contains(varname, "_") {
		panic(fmt.Sprintf("Variable names may not have underscores: line %d col %d", c.GetStart().GetLine(), c.GetStart().GetColumn()))
	}
}

func (l *FaultListener) EnterConstSpec(c *parser.ConstSpecContext) {
	identlist, ok := c.GetChild(0).(*parser.IdentListContext)
	if !ok {
		return
	}
	for _, name := range identlist.AllOperandName() {
		if strings.Contains(name.GetText(), "_") {
			panic(fmt.Sprintf("Variable names may not have underscores: line %d col %d", c.GetStart().GetLine(), c.GetStart().GetColumn()))
		}
	}
}

func (l *FaultListener) EnterStateFunc(c *parser.StateFuncContext) {
	varname := c.IDENT().GetText()
	if strings.Contains(varname, "_") {
		panic(fmt.Sprintf("Variable names may not have underscores: line %d col %d", c.GetStart().GetLine(), c.GetStart().GetColumn()))
	}
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
	if strings.Contains(varname, "_") {
		panic(fmt.Sprintf("Variable names may not have underscores: line %d col %d", c.GetStart().GetLine(), c.GetStart().GetColumn()))
	}
	l.scope = fmt.Sprint(l.scope, ".", varname)
}

func (l *FaultListener) EnterPropInt(c *parser.PropIntContext) {
	if strings.Contains(c.IDENT().GetText(), "_") {
		panic(fmt.Sprintf("Variable names may not have underscores: line %d col %d", c.GetStart().GetLine(), c.GetStart().GetColumn()))
	}
}

func (l *FaultListener) EnterPropString(c *parser.PropStringContext) {
	if strings.Contains(c.IDENT().GetText(), "_") {
		panic(fmt.Sprintf("Variable names may not have underscores: line %d col %d", c.GetStart().GetLine(), c.GetStart().GetColumn()))
	}
}

func (l *FaultListener) EnterPropBool(c *parser.PropBoolContext) {
	if strings.Contains(c.IDENT().GetText(), "_") {
		panic(fmt.Sprintf("Variable names may not have underscores: line %d col %d", c.GetStart().GetLine(), c.GetStart().GetColumn()))
	}
}

func (l *FaultListener) EnterPropVar(c *parser.PropVarContext) {
	if strings.Contains(c.IDENT().GetText(), "_") {
		panic(fmt.Sprintf("Variable names may not have underscores: line %d col %d", c.GetStart().GetLine(), c.GetStart().GetColumn()))
	}
}

func (l *FaultListener) EnterPropSolvable(c *parser.PropSolvableContext) {
	// Skip the bare IDENT case (reference, not a declaration)
	if c.GetChildCount() == 1 {
		return
	}
	if strings.Contains(c.IDENT().GetText(), "_") {
		panic(fmt.Sprintf("Variable names may not have underscores: line %d col %d", c.GetStart().GetLine(), c.GetStart().GetColumn()))
	}
}

func (l *FaultListener) EnterRunInit(c *parser.RunInitContext) {
	// IDENT(0) is the variable being declared; IDENT(1) (if present) is the type reference
	if strings.Contains(c.IDENT(0).GetText(), "_") {
		panic(fmt.Sprintf("Variable names may not have underscores: line %d col %d", c.GetStart().GetLine(), c.GetStart().GetColumn()))
	}
}
