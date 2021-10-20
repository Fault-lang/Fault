// Code generated from FaultParser.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // FaultParser

import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseFaultParserVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseFaultParserVisitor) VisitSpec(ctx *SpecContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFaultParserVisitor) VisitSpecClause(ctx *SpecClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFaultParserVisitor) VisitImportDecl(ctx *ImportDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFaultParserVisitor) VisitImportSpec(ctx *ImportSpecContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFaultParserVisitor) VisitImportPath(ctx *ImportPathContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFaultParserVisitor) VisitDeclaration(ctx *DeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFaultParserVisitor) VisitConstDecl(ctx *ConstDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFaultParserVisitor) VisitConstSpec(ctx *ConstSpecContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFaultParserVisitor) VisitIdentList(ctx *IdentListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFaultParserVisitor) VisitExpressionList(ctx *ExpressionListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFaultParserVisitor) VisitStructDecl(ctx *StructDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFaultParserVisitor) VisitFlow(ctx *FlowContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFaultParserVisitor) VisitStock(ctx *StockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFaultParserVisitor) VisitPropInt(ctx *PropIntContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFaultParserVisitor) VisitPropString(ctx *PropStringContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFaultParserVisitor) VisitPropFunc(ctx *PropFuncContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFaultParserVisitor) VisitPropVar(ctx *PropVarContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFaultParserVisitor) VisitInitDecl(ctx *InitDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFaultParserVisitor) VisitBlock(ctx *BlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFaultParserVisitor) VisitStatementList(ctx *StatementListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFaultParserVisitor) VisitStatement(ctx *StatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFaultParserVisitor) VisitSimpleStmt(ctx *SimpleStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFaultParserVisitor) VisitIncDecStmt(ctx *IncDecStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFaultParserVisitor) VisitAccessHistory(ctx *AccessHistoryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFaultParserVisitor) VisitAssertion(ctx *AssertionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFaultParserVisitor) VisitMiscAssign(ctx *MiscAssignContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFaultParserVisitor) VisitFaultAssign(ctx *FaultAssignContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFaultParserVisitor) VisitEmptyStmt(ctx *EmptyStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFaultParserVisitor) VisitIfStmt(ctx *IfStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFaultParserVisitor) VisitForStmt(ctx *ForStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFaultParserVisitor) VisitRounds(ctx *RoundsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFaultParserVisitor) VisitParamCall(ctx *ParamCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFaultParserVisitor) VisitRunBlock(ctx *RunBlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFaultParserVisitor) VisitRunStepExpr(ctx *RunStepExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFaultParserVisitor) VisitRunInit(ctx *RunInitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFaultParserVisitor) VisitFaultType(ctx *FaultTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFaultParserVisitor) VisitTyped(ctx *TypedContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFaultParserVisitor) VisitExpr(ctx *ExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFaultParserVisitor) VisitLrExpr(ctx *LrExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFaultParserVisitor) VisitPrefix(ctx *PrefixContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFaultParserVisitor) VisitOperand(ctx *OperandContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFaultParserVisitor) VisitOpName(ctx *OpNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFaultParserVisitor) VisitOpParam(ctx *OpParamContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFaultParserVisitor) VisitOpThis(ctx *OpThisContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFaultParserVisitor) VisitOpClock(ctx *OpClockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFaultParserVisitor) VisitOpInstance(ctx *OpInstanceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFaultParserVisitor) VisitNumeric(ctx *NumericContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFaultParserVisitor) VisitInteger(ctx *IntegerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFaultParserVisitor) VisitNegative(ctx *NegativeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFaultParserVisitor) VisitFloat_(ctx *Float_Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFaultParserVisitor) VisitString_(ctx *String_Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFaultParserVisitor) VisitBool_(ctx *Bool_Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFaultParserVisitor) VisitFunctionLit(ctx *FunctionLitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseFaultParserVisitor) VisitEos(ctx *EosContext) interface{} {
	return v.VisitChildren(ctx)
}
