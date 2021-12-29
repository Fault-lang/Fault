// Code generated from FaultParser.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // FaultParser

import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by FaultParser.
type FaultParserVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by FaultParser#spec.
	VisitSpec(ctx *SpecContext) interface{}

	// Visit a parse tree produced by FaultParser#specClause.
	VisitSpecClause(ctx *SpecClauseContext) interface{}

	// Visit a parse tree produced by FaultParser#importDecl.
	VisitImportDecl(ctx *ImportDeclContext) interface{}

	// Visit a parse tree produced by FaultParser#importSpec.
	VisitImportSpec(ctx *ImportSpecContext) interface{}

	// Visit a parse tree produced by FaultParser#importPath.
	VisitImportPath(ctx *ImportPathContext) interface{}

	// Visit a parse tree produced by FaultParser#declaration.
	VisitDeclaration(ctx *DeclarationContext) interface{}

	// Visit a parse tree produced by FaultParser#comparison.
	VisitComparison(ctx *ComparisonContext) interface{}

	// Visit a parse tree produced by FaultParser#constDecl.
	VisitConstDecl(ctx *ConstDeclContext) interface{}

	// Visit a parse tree produced by FaultParser#constSpec.
	VisitConstSpec(ctx *ConstSpecContext) interface{}

	// Visit a parse tree produced by FaultParser#identList.
	VisitIdentList(ctx *IdentListContext) interface{}

	// Visit a parse tree produced by FaultParser#expressionList.
	VisitExpressionList(ctx *ExpressionListContext) interface{}

	// Visit a parse tree produced by FaultParser#structDecl.
	VisitStructDecl(ctx *StructDeclContext) interface{}

	// Visit a parse tree produced by FaultParser#Flow.
	VisitFlow(ctx *FlowContext) interface{}

	// Visit a parse tree produced by FaultParser#Stock.
	VisitStock(ctx *StockContext) interface{}

	// Visit a parse tree produced by FaultParser#PropInt.
	VisitPropInt(ctx *PropIntContext) interface{}

	// Visit a parse tree produced by FaultParser#PropString.
	VisitPropString(ctx *PropStringContext) interface{}

	// Visit a parse tree produced by FaultParser#PropFunc.
	VisitPropFunc(ctx *PropFuncContext) interface{}

	// Visit a parse tree produced by FaultParser#PropVar.
	VisitPropVar(ctx *PropVarContext) interface{}

	// Visit a parse tree produced by FaultParser#initDecl.
	VisitInitDecl(ctx *InitDeclContext) interface{}

	// Visit a parse tree produced by FaultParser#block.
	VisitBlock(ctx *BlockContext) interface{}

	// Visit a parse tree produced by FaultParser#statementList.
	VisitStatementList(ctx *StatementListContext) interface{}

	// Visit a parse tree produced by FaultParser#statement.
	VisitStatement(ctx *StatementContext) interface{}

	// Visit a parse tree produced by FaultParser#simpleStmt.
	VisitSimpleStmt(ctx *SimpleStmtContext) interface{}

	// Visit a parse tree produced by FaultParser#incDecStmt.
	VisitIncDecStmt(ctx *IncDecStmtContext) interface{}

	// Visit a parse tree produced by FaultParser#accessHistory.
	VisitAccessHistory(ctx *AccessHistoryContext) interface{}

	// Visit a parse tree produced by FaultParser#assertion.
	VisitAssertion(ctx *AssertionContext) interface{}

	// Visit a parse tree produced by FaultParser#assumption.
	VisitAssumption(ctx *AssumptionContext) interface{}

	// Visit a parse tree produced by FaultParser#MiscAssign.
	VisitMiscAssign(ctx *MiscAssignContext) interface{}

	// Visit a parse tree produced by FaultParser#FaultAssign.
	VisitFaultAssign(ctx *FaultAssignContext) interface{}

	// Visit a parse tree produced by FaultParser#emptyStmt.
	VisitEmptyStmt(ctx *EmptyStmtContext) interface{}

	// Visit a parse tree produced by FaultParser#ifStmt.
	VisitIfStmt(ctx *IfStmtContext) interface{}

	// Visit a parse tree produced by FaultParser#forStmt.
	VisitForStmt(ctx *ForStmtContext) interface{}

	// Visit a parse tree produced by FaultParser#rounds.
	VisitRounds(ctx *RoundsContext) interface{}

	// Visit a parse tree produced by FaultParser#paramCall.
	VisitParamCall(ctx *ParamCallContext) interface{}

	// Visit a parse tree produced by FaultParser#runBlock.
	VisitRunBlock(ctx *RunBlockContext) interface{}

	// Visit a parse tree produced by FaultParser#runStepExpr.
	VisitRunStepExpr(ctx *RunStepExprContext) interface{}

	// Visit a parse tree produced by FaultParser#runInit.
	VisitRunInit(ctx *RunInitContext) interface{}

	// Visit a parse tree produced by FaultParser#runExpr.
	VisitRunExpr(ctx *RunExprContext) interface{}

	// Visit a parse tree produced by FaultParser#faultType.
	VisitFaultType(ctx *FaultTypeContext) interface{}

	// Visit a parse tree produced by FaultParser#Typed.
	VisitTyped(ctx *TypedContext) interface{}

	// Visit a parse tree produced by FaultParser#Expr.
	VisitExpr(ctx *ExprContext) interface{}

	// Visit a parse tree produced by FaultParser#lrExpr.
	VisitLrExpr(ctx *LrExprContext) interface{}

	// Visit a parse tree produced by FaultParser#Prefix.
	VisitPrefix(ctx *PrefixContext) interface{}

	// Visit a parse tree produced by FaultParser#operand.
	VisitOperand(ctx *OperandContext) interface{}

	// Visit a parse tree produced by FaultParser#OpName.
	VisitOpName(ctx *OpNameContext) interface{}

	// Visit a parse tree produced by FaultParser#OpParam.
	VisitOpParam(ctx *OpParamContext) interface{}

	// Visit a parse tree produced by FaultParser#OpThis.
	VisitOpThis(ctx *OpThisContext) interface{}

	// Visit a parse tree produced by FaultParser#OpClock.
	VisitOpClock(ctx *OpClockContext) interface{}

	// Visit a parse tree produced by FaultParser#OpInstance.
	VisitOpInstance(ctx *OpInstanceContext) interface{}

	// Visit a parse tree produced by FaultParser#numeric.
	VisitNumeric(ctx *NumericContext) interface{}

	// Visit a parse tree produced by FaultParser#integer.
	VisitInteger(ctx *IntegerContext) interface{}

	// Visit a parse tree produced by FaultParser#negative.
	VisitNegative(ctx *NegativeContext) interface{}

	// Visit a parse tree produced by FaultParser#float_.
	VisitFloat_(ctx *Float_Context) interface{}

	// Visit a parse tree produced by FaultParser#string_.
	VisitString_(ctx *String_Context) interface{}

	// Visit a parse tree produced by FaultParser#bool_.
	VisitBool_(ctx *Bool_Context) interface{}

	// Visit a parse tree produced by FaultParser#functionLit.
	VisitFunctionLit(ctx *FunctionLitContext) interface{}

	// Visit a parse tree produced by FaultParser#eos.
	VisitEos(ctx *EosContext) interface{}
}
