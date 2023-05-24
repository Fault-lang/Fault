// Code generated from FaultParser.g4 by ANTLR 4.12.0. DO NOT EDIT.

package parser // FaultParser

import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

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

// EnterSysSpec is called when production sysSpec is entered.
func (s *BaseFaultParserListener) EnterSysSpec(ctx *SysSpecContext) {}

// ExitSysSpec is called when production sysSpec is exited.
func (s *BaseFaultParserListener) ExitSysSpec(ctx *SysSpecContext) {}

// EnterSysClause is called when production sysClause is entered.
func (s *BaseFaultParserListener) EnterSysClause(ctx *SysClauseContext) {}

// ExitSysClause is called when production sysClause is exited.
func (s *BaseFaultParserListener) ExitSysClause(ctx *SysClauseContext) {}

// EnterGlobalDecl is called when production globalDecl is entered.
func (s *BaseFaultParserListener) EnterGlobalDecl(ctx *GlobalDeclContext) {}

// ExitGlobalDecl is called when production globalDecl is exited.
func (s *BaseFaultParserListener) ExitGlobalDecl(ctx *GlobalDeclContext) {}

// EnterSwap is called when production swap is entered.
func (s *BaseFaultParserListener) EnterSwap(ctx *SwapContext) {}

// ExitSwap is called when production swap is exited.
func (s *BaseFaultParserListener) ExitSwap(ctx *SwapContext) {}

// EnterComponentDecl is called when production componentDecl is entered.
func (s *BaseFaultParserListener) EnterComponentDecl(ctx *ComponentDeclContext) {}

// ExitComponentDecl is called when production componentDecl is exited.
func (s *BaseFaultParserListener) ExitComponentDecl(ctx *ComponentDeclContext) {}

// EnterStartBlock is called when production startBlock is entered.
func (s *BaseFaultParserListener) EnterStartBlock(ctx *StartBlockContext) {}

// ExitStartBlock is called when production startBlock is exited.
func (s *BaseFaultParserListener) ExitStartBlock(ctx *StartBlockContext) {}

// EnterStartPair is called when production startPair is entered.
func (s *BaseFaultParserListener) EnterStartPair(ctx *StartPairContext) {}

// ExitStartPair is called when production startPair is exited.
func (s *BaseFaultParserListener) ExitStartPair(ctx *StartPairContext) {}

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

// EnterComparison is called when production comparison is entered.
func (s *BaseFaultParserListener) EnterComparison(ctx *ComparisonContext) {}

// ExitComparison is called when production comparison is exited.
func (s *BaseFaultParserListener) ExitComparison(ctx *ComparisonContext) {}

// EnterConstDecl is called when production constDecl is entered.
func (s *BaseFaultParserListener) EnterConstDecl(ctx *ConstDeclContext) {}

// ExitConstDecl is called when production constDecl is exited.
func (s *BaseFaultParserListener) ExitConstDecl(ctx *ConstDeclContext) {}

// EnterConstSpec is called when production constSpec is entered.
func (s *BaseFaultParserListener) EnterConstSpec(ctx *ConstSpecContext) {}

// ExitConstSpec is called when production constSpec is exited.
func (s *BaseFaultParserListener) ExitConstSpec(ctx *ConstSpecContext) {}

// EnterStringDecl is called when production stringDecl is entered.
func (s *BaseFaultParserListener) EnterStringDecl(ctx *StringDeclContext) {}

// ExitStringDecl is called when production stringDecl is exited.
func (s *BaseFaultParserListener) ExitStringDecl(ctx *StringDeclContext) {}

// EnterCompoundString is called when production compoundString is entered.
func (s *BaseFaultParserListener) EnterCompoundString(ctx *CompoundStringContext) {}

// ExitCompoundString is called when production compoundString is exited.
func (s *BaseFaultParserListener) ExitCompoundString(ctx *CompoundStringContext) {}

// EnterIdentList is called when production identList is entered.
func (s *BaseFaultParserListener) EnterIdentList(ctx *IdentListContext) {}

// ExitIdentList is called when production identList is exited.
func (s *BaseFaultParserListener) ExitIdentList(ctx *IdentListContext) {}

// EnterConstants is called when production constants is entered.
func (s *BaseFaultParserListener) EnterConstants(ctx *ConstantsContext) {}

// ExitConstants is called when production constants is exited.
func (s *BaseFaultParserListener) ExitConstants(ctx *ConstantsContext) {}

// EnterNil is called when production nil is entered.
func (s *BaseFaultParserListener) EnterNil(ctx *NilContext) {}

// ExitNil is called when production nil is exited.
func (s *BaseFaultParserListener) ExitNil(ctx *NilContext) {}

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

// EnterPropFunc is called when production PropFunc is entered.
func (s *BaseFaultParserListener) EnterPropFunc(ctx *PropFuncContext) {}

// ExitPropFunc is called when production PropFunc is exited.
func (s *BaseFaultParserListener) ExitPropFunc(ctx *PropFuncContext) {}

// EnterSfMisc is called when production sfMisc is entered.
func (s *BaseFaultParserListener) EnterSfMisc(ctx *SfMiscContext) {}

// ExitSfMisc is called when production sfMisc is exited.
func (s *BaseFaultParserListener) ExitSfMisc(ctx *SfMiscContext) {}

// EnterStateFunc is called when production StateFunc is entered.
func (s *BaseFaultParserListener) EnterStateFunc(ctx *StateFuncContext) {}

// ExitStateFunc is called when production StateFunc is exited.
func (s *BaseFaultParserListener) ExitStateFunc(ctx *StateFuncContext) {}

// EnterCompMisc is called when production compMisc is entered.
func (s *BaseFaultParserListener) EnterCompMisc(ctx *CompMiscContext) {}

// ExitCompMisc is called when production compMisc is exited.
func (s *BaseFaultParserListener) ExitCompMisc(ctx *CompMiscContext) {}

// EnterPropInt is called when production PropInt is entered.
func (s *BaseFaultParserListener) EnterPropInt(ctx *PropIntContext) {}

// ExitPropInt is called when production PropInt is exited.
func (s *BaseFaultParserListener) ExitPropInt(ctx *PropIntContext) {}

// EnterPropString is called when production PropString is entered.
func (s *BaseFaultParserListener) EnterPropString(ctx *PropStringContext) {}

// ExitPropString is called when production PropString is exited.
func (s *BaseFaultParserListener) ExitPropString(ctx *PropStringContext) {}

// EnterPropBool is called when production PropBool is entered.
func (s *BaseFaultParserListener) EnterPropBool(ctx *PropBoolContext) {}

// ExitPropBool is called when production PropBool is exited.
func (s *BaseFaultParserListener) ExitPropBool(ctx *PropBoolContext) {}

// EnterPropVar is called when production PropVar is entered.
func (s *BaseFaultParserListener) EnterPropVar(ctx *PropVarContext) {}

// ExitPropVar is called when production PropVar is exited.
func (s *BaseFaultParserListener) ExitPropVar(ctx *PropVarContext) {}

// EnterPropSolvable is called when production PropSolvable is entered.
func (s *BaseFaultParserListener) EnterPropSolvable(ctx *PropSolvableContext) {}

// ExitPropSolvable is called when production PropSolvable is exited.
func (s *BaseFaultParserListener) ExitPropSolvable(ctx *PropSolvableContext) {}

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

// EnterBuiltins is called when production builtins is entered.
func (s *BaseFaultParserListener) EnterBuiltins(ctx *BuiltinsContext) {}

// ExitBuiltins is called when production builtins is exited.
func (s *BaseFaultParserListener) ExitBuiltins(ctx *BuiltinsContext) {}

// EnterBuiltinInfix is called when production builtinInfix is entered.
func (s *BaseFaultParserListener) EnterBuiltinInfix(ctx *BuiltinInfixContext) {}

// ExitBuiltinInfix is called when production builtinInfix is exited.
func (s *BaseFaultParserListener) ExitBuiltinInfix(ctx *BuiltinInfixContext) {}

// EnterAccessHistory is called when production accessHistory is entered.
func (s *BaseFaultParserListener) EnterAccessHistory(ctx *AccessHistoryContext) {}

// ExitAccessHistory is called when production accessHistory is exited.
func (s *BaseFaultParserListener) ExitAccessHistory(ctx *AccessHistoryContext) {}

// EnterAssertion is called when production assertion is entered.
func (s *BaseFaultParserListener) EnterAssertion(ctx *AssertionContext) {}

// ExitAssertion is called when production assertion is exited.
func (s *BaseFaultParserListener) ExitAssertion(ctx *AssertionContext) {}

// EnterAssumption is called when production assumption is entered.
func (s *BaseFaultParserListener) EnterAssumption(ctx *AssumptionContext) {}

// ExitAssumption is called when production assumption is exited.
func (s *BaseFaultParserListener) ExitAssumption(ctx *AssumptionContext) {}

// EnterTemporal is called when production temporal is entered.
func (s *BaseFaultParserListener) EnterTemporal(ctx *TemporalContext) {}

// ExitTemporal is called when production temporal is exited.
func (s *BaseFaultParserListener) ExitTemporal(ctx *TemporalContext) {}

// EnterInvar is called when production invar is entered.
func (s *BaseFaultParserListener) EnterInvar(ctx *InvarContext) {}

// ExitInvar is called when production invar is exited.
func (s *BaseFaultParserListener) ExitInvar(ctx *InvarContext) {}

// EnterStageInvariant is called when production stageInvariant is entered.
func (s *BaseFaultParserListener) EnterStageInvariant(ctx *StageInvariantContext) {}

// ExitStageInvariant is called when production stageInvariant is exited.
func (s *BaseFaultParserListener) ExitStageInvariant(ctx *StageInvariantContext) {}

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

// EnterIfStmtRun is called when production ifStmtRun is entered.
func (s *BaseFaultParserListener) EnterIfStmtRun(ctx *IfStmtRunContext) {}

// ExitIfStmtRun is called when production ifStmtRun is exited.
func (s *BaseFaultParserListener) ExitIfStmtRun(ctx *IfStmtRunContext) {}

// EnterIfStmtState is called when production ifStmtState is entered.
func (s *BaseFaultParserListener) EnterIfStmtState(ctx *IfStmtStateContext) {}

// ExitIfStmtState is called when production ifStmtState is exited.
func (s *BaseFaultParserListener) ExitIfStmtState(ctx *IfStmtStateContext) {}

// EnterForStmt is called when production forStmt is entered.
func (s *BaseFaultParserListener) EnterForStmt(ctx *ForStmtContext) {}

// ExitForStmt is called when production forStmt is exited.
func (s *BaseFaultParserListener) ExitForStmt(ctx *ForStmtContext) {}

// EnterRounds is called when production rounds is entered.
func (s *BaseFaultParserListener) EnterRounds(ctx *RoundsContext) {}

// ExitRounds is called when production rounds is exited.
func (s *BaseFaultParserListener) ExitRounds(ctx *RoundsContext) {}

// EnterParamCall is called when production paramCall is entered.
func (s *BaseFaultParserListener) EnterParamCall(ctx *ParamCallContext) {}

// ExitParamCall is called when production paramCall is exited.
func (s *BaseFaultParserListener) ExitParamCall(ctx *ParamCallContext) {}

// EnterStateBlock is called when production stateBlock is entered.
func (s *BaseFaultParserListener) EnterStateBlock(ctx *StateBlockContext) {}

// ExitStateBlock is called when production stateBlock is exited.
func (s *BaseFaultParserListener) ExitStateBlock(ctx *StateBlockContext) {}

// EnterStateStepExpr is called when production stateStepExpr is entered.
func (s *BaseFaultParserListener) EnterStateStepExpr(ctx *StateStepExprContext) {}

// ExitStateStepExpr is called when production stateStepExpr is exited.
func (s *BaseFaultParserListener) ExitStateStepExpr(ctx *StateStepExprContext) {}

// EnterStateChain is called when production stateChain is entered.
func (s *BaseFaultParserListener) EnterStateChain(ctx *StateChainContext) {}

// ExitStateChain is called when production stateChain is exited.
func (s *BaseFaultParserListener) ExitStateChain(ctx *StateChainContext) {}

// EnterStateExpr is called when production stateExpr is entered.
func (s *BaseFaultParserListener) EnterStateExpr(ctx *StateExprContext) {}

// ExitStateExpr is called when production stateExpr is exited.
func (s *BaseFaultParserListener) ExitStateExpr(ctx *StateExprContext) {}

// EnterRunBlock is called when production runBlock is entered.
func (s *BaseFaultParserListener) EnterRunBlock(ctx *RunBlockContext) {}

// ExitRunBlock is called when production runBlock is exited.
func (s *BaseFaultParserListener) ExitRunBlock(ctx *RunBlockContext) {}

// EnterInitBlock is called when production initBlock is entered.
func (s *BaseFaultParserListener) EnterInitBlock(ctx *InitBlockContext) {}

// ExitInitBlock is called when production initBlock is exited.
func (s *BaseFaultParserListener) ExitInitBlock(ctx *InitBlockContext) {}

// EnterRunInit is called when production runInit is entered.
func (s *BaseFaultParserListener) EnterRunInit(ctx *RunInitContext) {}

// ExitRunInit is called when production runInit is exited.
func (s *BaseFaultParserListener) ExitRunInit(ctx *RunInitContext) {}

// EnterRunStepExpr is called when production runStepExpr is entered.
func (s *BaseFaultParserListener) EnterRunStepExpr(ctx *RunStepExprContext) {}

// ExitRunStepExpr is called when production runStepExpr is exited.
func (s *BaseFaultParserListener) ExitRunStepExpr(ctx *RunStepExprContext) {}

// EnterRunExpr is called when production runExpr is entered.
func (s *BaseFaultParserListener) EnterRunExpr(ctx *RunExprContext) {}

// ExitRunExpr is called when production runExpr is exited.
func (s *BaseFaultParserListener) ExitRunExpr(ctx *RunExprContext) {}

// EnterFaultType is called when production faultType is entered.
func (s *BaseFaultParserListener) EnterFaultType(ctx *FaultTypeContext) {}

// ExitFaultType is called when production faultType is exited.
func (s *BaseFaultParserListener) ExitFaultType(ctx *FaultTypeContext) {}

// EnterSolvable is called when production solvable is entered.
func (s *BaseFaultParserListener) EnterSolvable(ctx *SolvableContext) {}

// ExitSolvable is called when production solvable is exited.
func (s *BaseFaultParserListener) ExitSolvable(ctx *SolvableContext) {}

// EnterTyped is called when production Typed is entered.
func (s *BaseFaultParserListener) EnterTyped(ctx *TypedContext) {}

// ExitTyped is called when production Typed is exited.
func (s *BaseFaultParserListener) ExitTyped(ctx *TypedContext) {}

// EnterExpr is called when production Expr is entered.
func (s *BaseFaultParserListener) EnterExpr(ctx *ExprContext) {}

// ExitExpr is called when production Expr is exited.
func (s *BaseFaultParserListener) ExitExpr(ctx *ExprContext) {}

// EnterExprPrefix is called when production ExprPrefix is entered.
func (s *BaseFaultParserListener) EnterExprPrefix(ctx *ExprPrefixContext) {}

// ExitExprPrefix is called when production ExprPrefix is exited.
func (s *BaseFaultParserListener) ExitExprPrefix(ctx *ExprPrefixContext) {}

// EnterLrExpr is called when production lrExpr is entered.
func (s *BaseFaultParserListener) EnterLrExpr(ctx *LrExprContext) {}

// ExitLrExpr is called when production lrExpr is exited.
func (s *BaseFaultParserListener) ExitLrExpr(ctx *LrExprContext) {}

// EnterOperand is called when production operand is entered.
func (s *BaseFaultParserListener) EnterOperand(ctx *OperandContext) {}

// ExitOperand is called when production operand is exited.
func (s *BaseFaultParserListener) ExitOperand(ctx *OperandContext) {}

// EnterOpName is called when production OpName is entered.
func (s *BaseFaultParserListener) EnterOpName(ctx *OpNameContext) {}

// ExitOpName is called when production OpName is exited.
func (s *BaseFaultParserListener) ExitOpName(ctx *OpNameContext) {}

// EnterOpParam is called when production OpParam is entered.
func (s *BaseFaultParserListener) EnterOpParam(ctx *OpParamContext) {}

// ExitOpParam is called when production OpParam is exited.
func (s *BaseFaultParserListener) ExitOpParam(ctx *OpParamContext) {}

// EnterOpThis is called when production OpThis is entered.
func (s *BaseFaultParserListener) EnterOpThis(ctx *OpThisContext) {}

// ExitOpThis is called when production OpThis is exited.
func (s *BaseFaultParserListener) ExitOpThis(ctx *OpThisContext) {}

// EnterOpClock is called when production OpClock is entered.
func (s *BaseFaultParserListener) EnterOpClock(ctx *OpClockContext) {}

// ExitOpClock is called when production OpClock is exited.
func (s *BaseFaultParserListener) ExitOpClock(ctx *OpClockContext) {}

// EnterOpInstance is called when production OpInstance is entered.
func (s *BaseFaultParserListener) EnterOpInstance(ctx *OpInstanceContext) {}

// ExitOpInstance is called when production OpInstance is exited.
func (s *BaseFaultParserListener) ExitOpInstance(ctx *OpInstanceContext) {}

// EnterPrefix is called when production prefix is entered.
func (s *BaseFaultParserListener) EnterPrefix(ctx *PrefixContext) {}

// ExitPrefix is called when production prefix is exited.
func (s *BaseFaultParserListener) ExitPrefix(ctx *PrefixContext) {}

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

// EnterStateLit is called when production stateLit is entered.
func (s *BaseFaultParserListener) EnterStateLit(ctx *StateLitContext) {}

// ExitStateLit is called when production stateLit is exited.
func (s *BaseFaultParserListener) ExitStateLit(ctx *StateLitContext) {}

// EnterEos is called when production eos is entered.
func (s *BaseFaultParserListener) EnterEos(ctx *EosContext) {}

// ExitEos is called when production eos is exited.
func (s *BaseFaultParserListener) ExitEos(ctx *EosContext) {}
