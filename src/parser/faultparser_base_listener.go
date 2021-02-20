// Code generated from FaultParser.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // FaultParser

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseFaultParserListener is a complete listener for a parse tree produced by FaultParser.
type BaseFaultParserListener struct{}

var _ FaultParserListener = &BaseFaultParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseFaultParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseFaultParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseFaultParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseFaultParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterSpec is called when production spec is entered.
func (s *BaseFaultParserListener) EnterSpec(ctx *SpecContext) {}

// ExitSpec is called when production spec is exited.
func (s *BaseFaultParserListener) ExitSpec(ctx *SpecContext) {}

// EnterSpecClause is called when production specClause is entered.
func (s *BaseFaultParserListener) EnterSpecClause(ctx *SpecClauseContext) {}

// ExitSpecClause is called when production specClause is exited.
func (s *BaseFaultParserListener) ExitSpecClause(ctx *SpecClauseContext) {}

// EnterImportDecl is called when production importDecl is entered.
func (s *BaseFaultParserListener) EnterImportDecl(ctx *ImportDeclContext) {}

// ExitImportDecl is called when production importDecl is exited.
func (s *BaseFaultParserListener) ExitImportDecl(ctx *ImportDeclContext) {}

// EnterImportSpec is called when production importSpec is entered.
func (s *BaseFaultParserListener) EnterImportSpec(ctx *ImportSpecContext) {}

// ExitImportSpec is called when production importSpec is exited.
func (s *BaseFaultParserListener) ExitImportSpec(ctx *ImportSpecContext) {}

// EnterImportPath is called when production importPath is entered.
func (s *BaseFaultParserListener) EnterImportPath(ctx *ImportPathContext) {}

// ExitImportPath is called when production importPath is exited.
func (s *BaseFaultParserListener) ExitImportPath(ctx *ImportPathContext) {}

// EnterDeclaration is called when production declaration is entered.
func (s *BaseFaultParserListener) EnterDeclaration(ctx *DeclarationContext) {}

// ExitDeclaration is called when production declaration is exited.
func (s *BaseFaultParserListener) ExitDeclaration(ctx *DeclarationContext) {}

// EnterConstDecl is called when production constDecl is entered.
func (s *BaseFaultParserListener) EnterConstDecl(ctx *ConstDeclContext) {}

// ExitConstDecl is called when production constDecl is exited.
func (s *BaseFaultParserListener) ExitConstDecl(ctx *ConstDeclContext) {}

// EnterConstSpec is called when production constSpec is entered.
func (s *BaseFaultParserListener) EnterConstSpec(ctx *ConstSpecContext) {}

// ExitConstSpec is called when production constSpec is exited.
func (s *BaseFaultParserListener) ExitConstSpec(ctx *ConstSpecContext) {}

// EnterIdentList is called when production identList is entered.
func (s *BaseFaultParserListener) EnterIdentList(ctx *IdentListContext) {}

// ExitIdentList is called when production identList is exited.
func (s *BaseFaultParserListener) ExitIdentList(ctx *IdentListContext) {}

// EnterExpressionList is called when production expressionList is entered.
func (s *BaseFaultParserListener) EnterExpressionList(ctx *ExpressionListContext) {}

// ExitExpressionList is called when production expressionList is exited.
func (s *BaseFaultParserListener) ExitExpressionList(ctx *ExpressionListContext) {}

// EnterStructDecl is called when production structDecl is entered.
func (s *BaseFaultParserListener) EnterStructDecl(ctx *StructDeclContext) {}

// ExitStructDecl is called when production structDecl is exited.
func (s *BaseFaultParserListener) ExitStructDecl(ctx *StructDeclContext) {}

// EnterFlow is called when production Flow is entered.
func (s *BaseFaultParserListener) EnterFlow(ctx *FlowContext) {}

// ExitFlow is called when production Flow is exited.
func (s *BaseFaultParserListener) ExitFlow(ctx *FlowContext) {}

// EnterStock is called when production Stock is entered.
func (s *BaseFaultParserListener) EnterStock(ctx *StockContext) {}

// ExitStock is called when production Stock is exited.
func (s *BaseFaultParserListener) ExitStock(ctx *StockContext) {}

// EnterPropInt is called when production PropInt is entered.
func (s *BaseFaultParserListener) EnterPropInt(ctx *PropIntContext) {}

// ExitPropInt is called when production PropInt is exited.
func (s *BaseFaultParserListener) ExitPropInt(ctx *PropIntContext) {}

// EnterPropString is called when production PropString is entered.
func (s *BaseFaultParserListener) EnterPropString(ctx *PropStringContext) {}

// ExitPropString is called when production PropString is exited.
func (s *BaseFaultParserListener) ExitPropString(ctx *PropStringContext) {}

// EnterPropFunc is called when production PropFunc is entered.
func (s *BaseFaultParserListener) EnterPropFunc(ctx *PropFuncContext) {}

// ExitPropFunc is called when production PropFunc is exited.
func (s *BaseFaultParserListener) ExitPropFunc(ctx *PropFuncContext) {}

// EnterPropVar is called when production PropVar is entered.
func (s *BaseFaultParserListener) EnterPropVar(ctx *PropVarContext) {}

// ExitPropVar is called when production PropVar is exited.
func (s *BaseFaultParserListener) ExitPropVar(ctx *PropVarContext) {}

// EnterInstance is called when production instance is entered.
func (s *BaseFaultParserListener) EnterInstance(ctx *InstanceContext) {}

// ExitInstance is called when production instance is exited.
func (s *BaseFaultParserListener) ExitInstance(ctx *InstanceContext) {}

// EnterInitDecl is called when production initDecl is entered.
func (s *BaseFaultParserListener) EnterInitDecl(ctx *InitDeclContext) {}

// ExitInitDecl is called when production initDecl is exited.
func (s *BaseFaultParserListener) ExitInitDecl(ctx *InitDeclContext) {}

// EnterBlock is called when production block is entered.
func (s *BaseFaultParserListener) EnterBlock(ctx *BlockContext) {}

// ExitBlock is called when production block is exited.
func (s *BaseFaultParserListener) ExitBlock(ctx *BlockContext) {}

// EnterStatementList is called when production statementList is entered.
func (s *BaseFaultParserListener) EnterStatementList(ctx *StatementListContext) {}

// ExitStatementList is called when production statementList is exited.
func (s *BaseFaultParserListener) ExitStatementList(ctx *StatementListContext) {}

// EnterStatement is called when production statement is entered.
func (s *BaseFaultParserListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaseFaultParserListener) ExitStatement(ctx *StatementContext) {}

// EnterSimpleStmt is called when production simpleStmt is entered.
func (s *BaseFaultParserListener) EnterSimpleStmt(ctx *SimpleStmtContext) {}

// ExitSimpleStmt is called when production simpleStmt is exited.
func (s *BaseFaultParserListener) ExitSimpleStmt(ctx *SimpleStmtContext) {}

// EnterIncDecStmt is called when production incDecStmt is entered.
func (s *BaseFaultParserListener) EnterIncDecStmt(ctx *IncDecStmtContext) {}

// ExitIncDecStmt is called when production incDecStmt is exited.
func (s *BaseFaultParserListener) ExitIncDecStmt(ctx *IncDecStmtContext) {}

// EnterAccessHistory is called when production accessHistory is entered.
func (s *BaseFaultParserListener) EnterAccessHistory(ctx *AccessHistoryContext) {}

// ExitAccessHistory is called when production accessHistory is exited.
func (s *BaseFaultParserListener) ExitAccessHistory(ctx *AccessHistoryContext) {}

// EnterAssertion is called when production assertion is entered.
func (s *BaseFaultParserListener) EnterAssertion(ctx *AssertionContext) {}

// ExitAssertion is called when production assertion is exited.
func (s *BaseFaultParserListener) ExitAssertion(ctx *AssertionContext) {}

// EnterMiscAssign is called when production MiscAssign is entered.
func (s *BaseFaultParserListener) EnterMiscAssign(ctx *MiscAssignContext) {}

// ExitMiscAssign is called when production MiscAssign is exited.
func (s *BaseFaultParserListener) ExitMiscAssign(ctx *MiscAssignContext) {}

// EnterFaultAssign is called when production FaultAssign is entered.
func (s *BaseFaultParserListener) EnterFaultAssign(ctx *FaultAssignContext) {}

// ExitFaultAssign is called when production FaultAssign is exited.
func (s *BaseFaultParserListener) ExitFaultAssign(ctx *FaultAssignContext) {}

// EnterEmptyStmt is called when production emptyStmt is entered.
func (s *BaseFaultParserListener) EnterEmptyStmt(ctx *EmptyStmtContext) {}

// ExitEmptyStmt is called when production emptyStmt is exited.
func (s *BaseFaultParserListener) ExitEmptyStmt(ctx *EmptyStmtContext) {}

// EnterIfStmt is called when production ifStmt is entered.
func (s *BaseFaultParserListener) EnterIfStmt(ctx *IfStmtContext) {}

// ExitIfStmt is called when production ifStmt is exited.
func (s *BaseFaultParserListener) ExitIfStmt(ctx *IfStmtContext) {}

// EnterForStmt is called when production forStmt is entered.
func (s *BaseFaultParserListener) EnterForStmt(ctx *ForStmtContext) {}

// ExitForStmt is called when production forStmt is exited.
func (s *BaseFaultParserListener) ExitForStmt(ctx *ForStmtContext) {}

// EnterExpr is called when production Expr is entered.
func (s *BaseFaultParserListener) EnterExpr(ctx *ExprContext) {}

// ExitExpr is called when production Expr is exited.
func (s *BaseFaultParserListener) ExitExpr(ctx *ExprContext) {}

// EnterLrExpr is called when production lrExpr is entered.
func (s *BaseFaultParserListener) EnterLrExpr(ctx *LrExprContext) {}

// ExitLrExpr is called when production lrExpr is exited.
func (s *BaseFaultParserListener) ExitLrExpr(ctx *LrExprContext) {}

// EnterPrefix is called when production Prefix is entered.
func (s *BaseFaultParserListener) EnterPrefix(ctx *PrefixContext) {}

// ExitPrefix is called when production Prefix is exited.
func (s *BaseFaultParserListener) ExitPrefix(ctx *PrefixContext) {}

// EnterOperand is called when production operand is entered.
func (s *BaseFaultParserListener) EnterOperand(ctx *OperandContext) {}

// ExitOperand is called when production operand is exited.
func (s *BaseFaultParserListener) ExitOperand(ctx *OperandContext) {}

// EnterOperandName is called when production operandName is entered.
func (s *BaseFaultParserListener) EnterOperandName(ctx *OperandNameContext) {}

// ExitOperandName is called when production operandName is exited.
func (s *BaseFaultParserListener) ExitOperandName(ctx *OperandNameContext) {}

// EnterNumeric is called when production numeric is entered.
func (s *BaseFaultParserListener) EnterNumeric(ctx *NumericContext) {}

// ExitNumeric is called when production numeric is exited.
func (s *BaseFaultParserListener) ExitNumeric(ctx *NumericContext) {}

// EnterInteger is called when production integer is entered.
func (s *BaseFaultParserListener) EnterInteger(ctx *IntegerContext) {}

// ExitInteger is called when production integer is exited.
func (s *BaseFaultParserListener) ExitInteger(ctx *IntegerContext) {}

// EnterNegative is called when production negative is entered.
func (s *BaseFaultParserListener) EnterNegative(ctx *NegativeContext) {}

// ExitNegative is called when production negative is exited.
func (s *BaseFaultParserListener) ExitNegative(ctx *NegativeContext) {}

// EnterFloat_ is called when production float_ is entered.
func (s *BaseFaultParserListener) EnterFloat_(ctx *Float_Context) {}

// ExitFloat_ is called when production float_ is exited.
func (s *BaseFaultParserListener) ExitFloat_(ctx *Float_Context) {}

// EnterString_ is called when production string_ is entered.
func (s *BaseFaultParserListener) EnterString_(ctx *String_Context) {}

// ExitString_ is called when production string_ is exited.
func (s *BaseFaultParserListener) ExitString_(ctx *String_Context) {}

// EnterBool_ is called when production bool_ is entered.
func (s *BaseFaultParserListener) EnterBool_(ctx *Bool_Context) {}

// ExitBool_ is called when production bool_ is exited.
func (s *BaseFaultParserListener) ExitBool_(ctx *Bool_Context) {}

// EnterFunctionLit is called when production functionLit is entered.
func (s *BaseFaultParserListener) EnterFunctionLit(ctx *FunctionLitContext) {}

// ExitFunctionLit is called when production functionLit is exited.
func (s *BaseFaultParserListener) ExitFunctionLit(ctx *FunctionLitContext) {}

// EnterEos is called when production eos is entered.
func (s *BaseFaultParserListener) EnterEos(ctx *EosContext) {}

// ExitEos is called when production eos is exited.
func (s *BaseFaultParserListener) ExitEos(ctx *EosContext) {}
