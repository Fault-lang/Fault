package listener

import (
	"fault/parser"
	"fmt"

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

func (l *FaultListener) EnterStructDecl(c *parser.StructDeclContext) {
	l.scope = c.GetChild(1).(antlr.TerminalNode).GetText()
	l.structscope = l.scope
}

func (l *FaultListener) EnterStateFunc(c *parser.StateFuncContext) {
	l.scope = fmt.Sprint(l.scope, ".", c.IDENT().GetText())
}

func (l *FaultListener) EnterStateBlock(c *parser.StateBlockContext) {
	if c.GetChildCount() < 3 {
		panic(fmt.Sprintf("Malformed fspec or fsystem file. A state function cannot be empty: line %d col %d", c.GetStart().GetLine(), c.GetStart().GetColumn()))
	}
}

func (l *FaultListener) EnterPropFunc(c *parser.PropFuncContext) {
	l.scope = fmt.Sprint(l.scope, ".", c.IDENT().GetText())
}
