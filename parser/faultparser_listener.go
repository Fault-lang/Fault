// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package parser // FaultParser

import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

// FaultParserListener is a complete listener for a parse tree produced by FaultParser.
type FaultParserListener interface {
	antlr.ParseTreeListener

	// EnterSysSpec is called when entering the sysSpec production.
	EnterSysSpec(c *SysSpecContext)

	// EnterSysClause is called when entering the sysClause production.
	EnterSysClause(c *SysClauseContext)

	// EnterGlobalDecl is called when entering the globalDecl production.
	EnterGlobalDecl(c *GlobalDeclContext)

	// EnterComponentDecl is called when entering the componentDecl production.
	EnterComponentDecl(c *ComponentDeclContext)

	// EnterStartBlock is called when entering the startBlock production.
	EnterStartBlock(c *StartBlockContext)

	// EnterStartPair is called when entering the startPair production.
	EnterStartPair(c *StartPairContext)

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

	// EnterComparison is called when entering the comparison production.
	EnterComparison(c *ComparisonContext)

	// EnterConstDecl is called when entering the constDecl production.
	EnterConstDecl(c *ConstDeclContext)

	// EnterConstSpec is called when entering the constSpec production.
	EnterConstSpec(c *ConstSpecContext)

	// EnterIdentList is called when entering the identList production.
	EnterIdentList(c *IdentListContext)

	// EnterConstants is called when entering the constants production.
	EnterConstants(c *ConstantsContext)

	// EnterNil is called when entering the nil production.
	EnterNil(c *NilContext)

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

	// EnterPropBool is called when entering the PropBool production.
	EnterPropBool(c *PropBoolContext)

	// EnterPropFunc is called when entering the PropFunc production.
	EnterPropFunc(c *PropFuncContext)

	// EnterPropVar is called when entering the PropVar production.
	EnterPropVar(c *PropVarContext)

	// EnterPropSolvable is called when entering the PropSolvable production.
	EnterPropSolvable(c *PropSolvableContext)

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

	// EnterBuiltins is called when entering the builtins production.
	EnterBuiltins(c *BuiltinsContext)

	// EnterBuiltinInfix is called when entering the builtinInfix production.
	EnterBuiltinInfix(c *BuiltinInfixContext)

	// EnterAccessHistory is called when entering the accessHistory production.
	EnterAccessHistory(c *AccessHistoryContext)

	// EnterAssertion is called when entering the assertion production.
	EnterAssertion(c *AssertionContext)

	// EnterAssumption is called when entering the assumption production.
	EnterAssumption(c *AssumptionContext)

	// EnterTemporal is called when entering the temporal production.
	EnterTemporal(c *TemporalContext)

	// EnterInvariant is called when entering the invariant production.
	EnterInvariant(c *InvariantContext)

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

	// EnterRounds is called when entering the rounds production.
	EnterRounds(c *RoundsContext)

	// EnterParamCall is called when entering the paramCall production.
	EnterParamCall(c *ParamCallContext)

	// EnterRunBlock is called when entering the runBlock production.
	EnterRunBlock(c *RunBlockContext)

	// EnterRunStepExpr is called when entering the runStepExpr production.
	EnterRunStepExpr(c *RunStepExprContext)

	// EnterRunInit is called when entering the runInit production.
	EnterRunInit(c *RunInitContext)

	// EnterRunExpr is called when entering the runExpr production.
	EnterRunExpr(c *RunExprContext)

	// EnterFaultType is called when entering the faultType production.
	EnterFaultType(c *FaultTypeContext)

	// EnterSolvable is called when entering the solvable production.
	EnterSolvable(c *SolvableContext)

	// EnterTyped is called when entering the Typed production.
	EnterTyped(c *TypedContext)

	// EnterExpr is called when entering the Expr production.
	EnterExpr(c *ExprContext)

	// EnterExprPrefix is called when entering the ExprPrefix production.
	EnterExprPrefix(c *ExprPrefixContext)

	// EnterLrExpr is called when entering the lrExpr production.
	EnterLrExpr(c *LrExprContext)

	// EnterOperand is called when entering the operand production.
	EnterOperand(c *OperandContext)

	// EnterOpName is called when entering the OpName production.
	EnterOpName(c *OpNameContext)

	// EnterOpParam is called when entering the OpParam production.
	EnterOpParam(c *OpParamContext)

	// EnterOpThis is called when entering the OpThis production.
	EnterOpThis(c *OpThisContext)

	// EnterOpClock is called when entering the OpClock production.
	EnterOpClock(c *OpClockContext)

	// EnterOpInstance is called when entering the OpInstance production.
	EnterOpInstance(c *OpInstanceContext)

	// EnterPrefix is called when entering the prefix production.
	EnterPrefix(c *PrefixContext)

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

	// ExitSysSpec is called when exiting the sysSpec production.
	ExitSysSpec(c *SysSpecContext)

	// ExitSysClause is called when exiting the sysClause production.
	ExitSysClause(c *SysClauseContext)

	// ExitGlobalDecl is called when exiting the globalDecl production.
	ExitGlobalDecl(c *GlobalDeclContext)

	// ExitComponentDecl is called when exiting the componentDecl production.
	ExitComponentDecl(c *ComponentDeclContext)

	// ExitStartBlock is called when exiting the startBlock production.
	ExitStartBlock(c *StartBlockContext)

	// ExitStartPair is called when exiting the startPair production.
	ExitStartPair(c *StartPairContext)

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

	// ExitComparison is called when exiting the comparison production.
	ExitComparison(c *ComparisonContext)

	// ExitConstDecl is called when exiting the constDecl production.
	ExitConstDecl(c *ConstDeclContext)

	// ExitConstSpec is called when exiting the constSpec production.
	ExitConstSpec(c *ConstSpecContext)

	// ExitIdentList is called when exiting the identList production.
	ExitIdentList(c *IdentListContext)

	// ExitConstants is called when exiting the constants production.
	ExitConstants(c *ConstantsContext)

	// ExitNil is called when exiting the nil production.
	ExitNil(c *NilContext)

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

	// ExitPropBool is called when exiting the PropBool production.
	ExitPropBool(c *PropBoolContext)

	// ExitPropFunc is called when exiting the PropFunc production.
	ExitPropFunc(c *PropFuncContext)

	// ExitPropVar is called when exiting the PropVar production.
	ExitPropVar(c *PropVarContext)

	// ExitPropSolvable is called when exiting the PropSolvable production.
	ExitPropSolvable(c *PropSolvableContext)

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

	// ExitBuiltins is called when exiting the builtins production.
	ExitBuiltins(c *BuiltinsContext)

	// ExitBuiltinInfix is called when exiting the builtinInfix production.
	ExitBuiltinInfix(c *BuiltinInfixContext)

	// ExitAccessHistory is called when exiting the accessHistory production.
	ExitAccessHistory(c *AccessHistoryContext)

	// ExitAssertion is called when exiting the assertion production.
	ExitAssertion(c *AssertionContext)

	// ExitAssumption is called when exiting the assumption production.
	ExitAssumption(c *AssumptionContext)

	// ExitTemporal is called when exiting the temporal production.
	ExitTemporal(c *TemporalContext)

	// ExitInvariant is called when exiting the invariant production.
	ExitInvariant(c *InvariantContext)

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

	// ExitRounds is called when exiting the rounds production.
	ExitRounds(c *RoundsContext)

	// ExitParamCall is called when exiting the paramCall production.
	ExitParamCall(c *ParamCallContext)

	// ExitRunBlock is called when exiting the runBlock production.
	ExitRunBlock(c *RunBlockContext)

	// ExitRunStepExpr is called when exiting the runStepExpr production.
	ExitRunStepExpr(c *RunStepExprContext)

	// ExitRunInit is called when exiting the runInit production.
	ExitRunInit(c *RunInitContext)

	// ExitRunExpr is called when exiting the runExpr production.
	ExitRunExpr(c *RunExprContext)

	// ExitFaultType is called when exiting the faultType production.
	ExitFaultType(c *FaultTypeContext)

	// ExitSolvable is called when exiting the solvable production.
	ExitSolvable(c *SolvableContext)

	// ExitTyped is called when exiting the Typed production.
	ExitTyped(c *TypedContext)

	// ExitExpr is called when exiting the Expr production.
	ExitExpr(c *ExprContext)

	// ExitExprPrefix is called when exiting the ExprPrefix production.
	ExitExprPrefix(c *ExprPrefixContext)

	// ExitLrExpr is called when exiting the lrExpr production.
	ExitLrExpr(c *LrExprContext)

	// ExitOperand is called when exiting the operand production.
	ExitOperand(c *OperandContext)

	// ExitOpName is called when exiting the OpName production.
	ExitOpName(c *OpNameContext)

	// ExitOpParam is called when exiting the OpParam production.
	ExitOpParam(c *OpParamContext)

	// ExitOpThis is called when exiting the OpThis production.
	ExitOpThis(c *OpThisContext)

	// ExitOpClock is called when exiting the OpClock production.
	ExitOpClock(c *OpClockContext)

	// ExitOpInstance is called when exiting the OpInstance production.
	ExitOpInstance(c *OpInstanceContext)

	// ExitPrefix is called when exiting the prefix production.
	ExitPrefix(c *PrefixContext)

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
