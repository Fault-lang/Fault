// Code generated from FaultParser.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // FaultParser

import "github.com/antlr/antlr4/runtime/Go/antlr"

// FaultParserListener is a complete listener for a parse tree produced by FaultParser.
type FaultParserListener interface {
	antlr.ParseTreeListener

	// EnterSpec is called when entering the spec production.
	EnterSpec(c *SpecContext)

	// EnterSpecClause is called when entering the specClause production.
	EnterSpecClause(c *SpecClauseContext)

	// EnterImportDecl is called when entering the importDecl production.
	EnterImportDecl(c *ImportDeclContext)

	// EnterImportSpec is called when entering the importSpec production.
	EnterImportSpec(c *ImportSpecContext)

	// EnterImportPath is called when entering the importPath production.
	EnterImportPath(c *ImportPathContext)

	// EnterDeclaration is called when entering the declaration production.
	EnterDeclaration(c *DeclarationContext)

	// EnterConstDecl is called when entering the constDecl production.
	EnterConstDecl(c *ConstDeclContext)

	// EnterConstSpec is called when entering the constSpec production.
	EnterConstSpec(c *ConstSpecContext)

	// EnterIdentList is called when entering the identList production.
	EnterIdentList(c *IdentListContext)

	// EnterExpressionList is called when entering the expressionList production.
	EnterExpressionList(c *ExpressionListContext)

	// EnterStructDecl is called when entering the structDecl production.
	EnterStructDecl(c *StructDeclContext)

	// EnterFlow is called when entering the Flow production.
	EnterFlow(c *FlowContext)

	// EnterStock is called when entering the Stock production.
	EnterStock(c *StockContext)

	// EnterPropInt is called when entering the PropInt production.
	EnterPropInt(c *PropIntContext)

	// EnterPropString is called when entering the PropString production.
	EnterPropString(c *PropStringContext)

	// EnterPropFunc is called when entering the PropFunc production.
	EnterPropFunc(c *PropFuncContext)

	// EnterPropVar is called when entering the PropVar production.
	EnterPropVar(c *PropVarContext)

	// EnterInstance is called when entering the instance production.
	EnterInstance(c *InstanceContext)

	// EnterInitDecl is called when entering the initDecl production.
	EnterInitDecl(c *InitDeclContext)

	// EnterBlock is called when entering the block production.
	EnterBlock(c *BlockContext)

	// EnterStatementList is called when entering the statementList production.
	EnterStatementList(c *StatementListContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterSimpleStmt is called when entering the simpleStmt production.
	EnterSimpleStmt(c *SimpleStmtContext)

	// EnterIncDecStmt is called when entering the incDecStmt production.
	EnterIncDecStmt(c *IncDecStmtContext)

	// EnterAccessHistory is called when entering the accessHistory production.
	EnterAccessHistory(c *AccessHistoryContext)

	// EnterAssertion is called when entering the assertion production.
	EnterAssertion(c *AssertionContext)

	// EnterMiscAssign is called when entering the MiscAssign production.
	EnterMiscAssign(c *MiscAssignContext)

	// EnterFaultAssign is called when entering the FaultAssign production.
	EnterFaultAssign(c *FaultAssignContext)

	// EnterEmptyStmt is called when entering the emptyStmt production.
	EnterEmptyStmt(c *EmptyStmtContext)

	// EnterIfStmt is called when entering the ifStmt production.
	EnterIfStmt(c *IfStmtContext)

	// EnterForStmt is called when entering the forStmt production.
	EnterForStmt(c *ForStmtContext)

	// EnterExpr is called when entering the Expr production.
	EnterExpr(c *ExprContext)

	// EnterLrExpr is called when entering the lrExpr production.
	EnterLrExpr(c *LrExprContext)

	// EnterPrefix is called when entering the Prefix production.
	EnterPrefix(c *PrefixContext)

	// EnterOperand is called when entering the operand production.
	EnterOperand(c *OperandContext)

	// EnterOperandName is called when entering the operandName production.
	EnterOperandName(c *OperandNameContext)

	// EnterNumeric is called when entering the numeric production.
	EnterNumeric(c *NumericContext)

	// EnterInteger is called when entering the integer production.
	EnterInteger(c *IntegerContext)

	// EnterNegative is called when entering the negative production.
	EnterNegative(c *NegativeContext)

	// EnterFloat_ is called when entering the float_ production.
	EnterFloat_(c *Float_Context)

	// EnterString_ is called when entering the string_ production.
	EnterString_(c *String_Context)

	// EnterBool_ is called when entering the bool_ production.
	EnterBool_(c *Bool_Context)

	// EnterFunctionLit is called when entering the functionLit production.
	EnterFunctionLit(c *FunctionLitContext)

	// EnterEos is called when entering the eos production.
	EnterEos(c *EosContext)

	// ExitSpec is called when exiting the spec production.
	ExitSpec(c *SpecContext)

	// ExitSpecClause is called when exiting the specClause production.
	ExitSpecClause(c *SpecClauseContext)

	// ExitImportDecl is called when exiting the importDecl production.
	ExitImportDecl(c *ImportDeclContext)

	// ExitImportSpec is called when exiting the importSpec production.
	ExitImportSpec(c *ImportSpecContext)

	// ExitImportPath is called when exiting the importPath production.
	ExitImportPath(c *ImportPathContext)

	// ExitDeclaration is called when exiting the declaration production.
	ExitDeclaration(c *DeclarationContext)

	// ExitConstDecl is called when exiting the constDecl production.
	ExitConstDecl(c *ConstDeclContext)

	// ExitConstSpec is called when exiting the constSpec production.
	ExitConstSpec(c *ConstSpecContext)

	// ExitIdentList is called when exiting the identList production.
	ExitIdentList(c *IdentListContext)

	// ExitExpressionList is called when exiting the expressionList production.
	ExitExpressionList(c *ExpressionListContext)

	// ExitStructDecl is called when exiting the structDecl production.
	ExitStructDecl(c *StructDeclContext)

	// ExitFlow is called when exiting the Flow production.
	ExitFlow(c *FlowContext)

	// ExitStock is called when exiting the Stock production.
	ExitStock(c *StockContext)

	// ExitPropInt is called when exiting the PropInt production.
	ExitPropInt(c *PropIntContext)

	// ExitPropString is called when exiting the PropString production.
	ExitPropString(c *PropStringContext)

	// ExitPropFunc is called when exiting the PropFunc production.
	ExitPropFunc(c *PropFuncContext)

	// ExitPropVar is called when exiting the PropVar production.
	ExitPropVar(c *PropVarContext)

	// ExitInstance is called when exiting the instance production.
	ExitInstance(c *InstanceContext)

	// ExitInitDecl is called when exiting the initDecl production.
	ExitInitDecl(c *InitDeclContext)

	// ExitBlock is called when exiting the block production.
	ExitBlock(c *BlockContext)

	// ExitStatementList is called when exiting the statementList production.
	ExitStatementList(c *StatementListContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitSimpleStmt is called when exiting the simpleStmt production.
	ExitSimpleStmt(c *SimpleStmtContext)

	// ExitIncDecStmt is called when exiting the incDecStmt production.
	ExitIncDecStmt(c *IncDecStmtContext)

	// ExitAccessHistory is called when exiting the accessHistory production.
	ExitAccessHistory(c *AccessHistoryContext)

	// ExitAssertion is called when exiting the assertion production.
	ExitAssertion(c *AssertionContext)

	// ExitMiscAssign is called when exiting the MiscAssign production.
	ExitMiscAssign(c *MiscAssignContext)

	// ExitFaultAssign is called when exiting the FaultAssign production.
	ExitFaultAssign(c *FaultAssignContext)

	// ExitEmptyStmt is called when exiting the emptyStmt production.
	ExitEmptyStmt(c *EmptyStmtContext)

	// ExitIfStmt is called when exiting the ifStmt production.
	ExitIfStmt(c *IfStmtContext)

	// ExitForStmt is called when exiting the forStmt production.
	ExitForStmt(c *ForStmtContext)

	// ExitExpr is called when exiting the Expr production.
	ExitExpr(c *ExprContext)

	// ExitLrExpr is called when exiting the lrExpr production.
	ExitLrExpr(c *LrExprContext)

	// ExitPrefix is called when exiting the Prefix production.
	ExitPrefix(c *PrefixContext)

	// ExitOperand is called when exiting the operand production.
	ExitOperand(c *OperandContext)

	// ExitOperandName is called when exiting the operandName production.
	ExitOperandName(c *OperandNameContext)

	// ExitNumeric is called when exiting the numeric production.
	ExitNumeric(c *NumericContext)

	// ExitInteger is called when exiting the integer production.
	ExitInteger(c *IntegerContext)

	// ExitNegative is called when exiting the negative production.
	ExitNegative(c *NegativeContext)

	// ExitFloat_ is called when exiting the float_ production.
	ExitFloat_(c *Float_Context)

	// ExitString_ is called when exiting the string_ production.
	ExitString_(c *String_Context)

	// ExitBool_ is called when exiting the bool_ production.
	ExitBool_(c *Bool_Context)

	// ExitFunctionLit is called when exiting the functionLit production.
	ExitFunctionLit(c *FunctionLitContext)

	// ExitEos is called when exiting the eos production.
	ExitEos(c *EosContext)
}
