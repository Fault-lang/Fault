// Code generated from SMTLIBv2.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // SMTLIBv2

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseSMTLIBv2Listener is a complete listener for a parse tree produced by SMTLIBv2Parser.
type BaseSMTLIBv2Listener struct{}

var _ SMTLIBv2Listener = &BaseSMTLIBv2Listener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseSMTLIBv2Listener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseSMTLIBv2Listener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseSMTLIBv2Listener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseSMTLIBv2Listener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterStart is called when production start is entered.
func (s *BaseSMTLIBv2Listener) EnterStart(ctx *StartContext) {}

// ExitStart is called when production start is exited.
func (s *BaseSMTLIBv2Listener) ExitStart(ctx *StartContext) {}

// EnterGeneralReservedWord is called when production generalReservedWord is entered.
func (s *BaseSMTLIBv2Listener) EnterGeneralReservedWord(ctx *GeneralReservedWordContext) {}

// ExitGeneralReservedWord is called when production generalReservedWord is exited.
func (s *BaseSMTLIBv2Listener) ExitGeneralReservedWord(ctx *GeneralReservedWordContext) {}

// EnterSimpleSymbol is called when production simpleSymbol is entered.
func (s *BaseSMTLIBv2Listener) EnterSimpleSymbol(ctx *SimpleSymbolContext) {}

// ExitSimpleSymbol is called when production simpleSymbol is exited.
func (s *BaseSMTLIBv2Listener) ExitSimpleSymbol(ctx *SimpleSymbolContext) {}

// EnterQuotedSymbol is called when production quotedSymbol is entered.
func (s *BaseSMTLIBv2Listener) EnterQuotedSymbol(ctx *QuotedSymbolContext) {}

// ExitQuotedSymbol is called when production quotedSymbol is exited.
func (s *BaseSMTLIBv2Listener) ExitQuotedSymbol(ctx *QuotedSymbolContext) {}

// EnterPredefSymbol is called when production predefSymbol is entered.
func (s *BaseSMTLIBv2Listener) EnterPredefSymbol(ctx *PredefSymbolContext) {}

// ExitPredefSymbol is called when production predefSymbol is exited.
func (s *BaseSMTLIBv2Listener) ExitPredefSymbol(ctx *PredefSymbolContext) {}

// EnterPredefKeyword is called when production predefKeyword is entered.
func (s *BaseSMTLIBv2Listener) EnterPredefKeyword(ctx *PredefKeywordContext) {}

// ExitPredefKeyword is called when production predefKeyword is exited.
func (s *BaseSMTLIBv2Listener) ExitPredefKeyword(ctx *PredefKeywordContext) {}

// EnterSymbol is called when production symbol is entered.
func (s *BaseSMTLIBv2Listener) EnterSymbol(ctx *SymbolContext) {}

// ExitSymbol is called when production symbol is exited.
func (s *BaseSMTLIBv2Listener) ExitSymbol(ctx *SymbolContext) {}

// EnterNumeral is called when production numeral is entered.
func (s *BaseSMTLIBv2Listener) EnterNumeral(ctx *NumeralContext) {}

// ExitNumeral is called when production numeral is exited.
func (s *BaseSMTLIBv2Listener) ExitNumeral(ctx *NumeralContext) {}

// EnterDecimal is called when production decimal is entered.
func (s *BaseSMTLIBv2Listener) EnterDecimal(ctx *DecimalContext) {}

// ExitDecimal is called when production decimal is exited.
func (s *BaseSMTLIBv2Listener) ExitDecimal(ctx *DecimalContext) {}

// EnterHexadecimal is called when production hexadecimal is entered.
func (s *BaseSMTLIBv2Listener) EnterHexadecimal(ctx *HexadecimalContext) {}

// ExitHexadecimal is called when production hexadecimal is exited.
func (s *BaseSMTLIBv2Listener) ExitHexadecimal(ctx *HexadecimalContext) {}

// EnterBinary is called when production binary is entered.
func (s *BaseSMTLIBv2Listener) EnterBinary(ctx *BinaryContext) {}

// ExitBinary is called when production binary is exited.
func (s *BaseSMTLIBv2Listener) ExitBinary(ctx *BinaryContext) {}

// EnterString_ is called when production string_ is entered.
func (s *BaseSMTLIBv2Listener) EnterString_(ctx *String_Context) {}

// ExitString_ is called when production string_ is exited.
func (s *BaseSMTLIBv2Listener) ExitString_(ctx *String_Context) {}

// EnterKeyword is called when production keyword is entered.
func (s *BaseSMTLIBv2Listener) EnterKeyword(ctx *KeywordContext) {}

// ExitKeyword is called when production keyword is exited.
func (s *BaseSMTLIBv2Listener) ExitKeyword(ctx *KeywordContext) {}

// EnterSpec_constant is called when production spec_constant is entered.
func (s *BaseSMTLIBv2Listener) EnterSpec_constant(ctx *Spec_constantContext) {}

// ExitSpec_constant is called when production spec_constant is exited.
func (s *BaseSMTLIBv2Listener) ExitSpec_constant(ctx *Spec_constantContext) {}

// EnterS_expr is called when production s_expr is entered.
func (s *BaseSMTLIBv2Listener) EnterS_expr(ctx *S_exprContext) {}

// ExitS_expr is called when production s_expr is exited.
func (s *BaseSMTLIBv2Listener) ExitS_expr(ctx *S_exprContext) {}

// EnterIndex is called when production index is entered.
func (s *BaseSMTLIBv2Listener) EnterIndex(ctx *IndexContext) {}

// ExitIndex is called when production index is exited.
func (s *BaseSMTLIBv2Listener) ExitIndex(ctx *IndexContext) {}

// EnterIdentifier is called when production identifier is entered.
func (s *BaseSMTLIBv2Listener) EnterIdentifier(ctx *IdentifierContext) {}

// ExitIdentifier is called when production identifier is exited.
func (s *BaseSMTLIBv2Listener) ExitIdentifier(ctx *IdentifierContext) {}

// EnterAttribute_value is called when production attribute_value is entered.
func (s *BaseSMTLIBv2Listener) EnterAttribute_value(ctx *Attribute_valueContext) {}

// ExitAttribute_value is called when production attribute_value is exited.
func (s *BaseSMTLIBv2Listener) ExitAttribute_value(ctx *Attribute_valueContext) {}

// EnterAttribute is called when production attribute is entered.
func (s *BaseSMTLIBv2Listener) EnterAttribute(ctx *AttributeContext) {}

// ExitAttribute is called when production attribute is exited.
func (s *BaseSMTLIBv2Listener) ExitAttribute(ctx *AttributeContext) {}

// EnterSort is called when production sort is entered.
func (s *BaseSMTLIBv2Listener) EnterSort(ctx *SortContext) {}

// ExitSort is called when production sort is exited.
func (s *BaseSMTLIBv2Listener) ExitSort(ctx *SortContext) {}

// EnterQual_identifer is called when production qual_identifer is entered.
func (s *BaseSMTLIBv2Listener) EnterQual_identifer(ctx *Qual_identiferContext) {}

// ExitQual_identifer is called when production qual_identifer is exited.
func (s *BaseSMTLIBv2Listener) ExitQual_identifer(ctx *Qual_identiferContext) {}

// EnterVar_binding is called when production var_binding is entered.
func (s *BaseSMTLIBv2Listener) EnterVar_binding(ctx *Var_bindingContext) {}

// ExitVar_binding is called when production var_binding is exited.
func (s *BaseSMTLIBv2Listener) ExitVar_binding(ctx *Var_bindingContext) {}

// EnterSorted_var is called when production sorted_var is entered.
func (s *BaseSMTLIBv2Listener) EnterSorted_var(ctx *Sorted_varContext) {}

// ExitSorted_var is called when production sorted_var is exited.
func (s *BaseSMTLIBv2Listener) ExitSorted_var(ctx *Sorted_varContext) {}

// EnterPattern is called when production pattern is entered.
func (s *BaseSMTLIBv2Listener) EnterPattern(ctx *PatternContext) {}

// ExitPattern is called when production pattern is exited.
func (s *BaseSMTLIBv2Listener) ExitPattern(ctx *PatternContext) {}

// EnterMatch_case is called when production match_case is entered.
func (s *BaseSMTLIBv2Listener) EnterMatch_case(ctx *Match_caseContext) {}

// ExitMatch_case is called when production match_case is exited.
func (s *BaseSMTLIBv2Listener) ExitMatch_case(ctx *Match_caseContext) {}

// EnterVariable is called when production variable is entered.
func (s *BaseSMTLIBv2Listener) EnterVariable(ctx *VariableContext) {}

// ExitVariable is called when production variable is exited.
func (s *BaseSMTLIBv2Listener) ExitVariable(ctx *VariableContext) {}

// EnterTerm is called when production term is entered.
func (s *BaseSMTLIBv2Listener) EnterTerm(ctx *TermContext) {}

// ExitTerm is called when production term is exited.
func (s *BaseSMTLIBv2Listener) ExitTerm(ctx *TermContext) {}

// EnterSort_symbol_decl is called when production sort_symbol_decl is entered.
func (s *BaseSMTLIBv2Listener) EnterSort_symbol_decl(ctx *Sort_symbol_declContext) {}

// ExitSort_symbol_decl is called when production sort_symbol_decl is exited.
func (s *BaseSMTLIBv2Listener) ExitSort_symbol_decl(ctx *Sort_symbol_declContext) {}

// EnterMeta_spec_constant is called when production meta_spec_constant is entered.
func (s *BaseSMTLIBv2Listener) EnterMeta_spec_constant(ctx *Meta_spec_constantContext) {}

// ExitMeta_spec_constant is called when production meta_spec_constant is exited.
func (s *BaseSMTLIBv2Listener) ExitMeta_spec_constant(ctx *Meta_spec_constantContext) {}

// EnterFun_symbol_decl is called when production fun_symbol_decl is entered.
func (s *BaseSMTLIBv2Listener) EnterFun_symbol_decl(ctx *Fun_symbol_declContext) {}

// ExitFun_symbol_decl is called when production fun_symbol_decl is exited.
func (s *BaseSMTLIBv2Listener) ExitFun_symbol_decl(ctx *Fun_symbol_declContext) {}

// EnterPar_fun_symbol_decl is called when production par_fun_symbol_decl is entered.
func (s *BaseSMTLIBv2Listener) EnterPar_fun_symbol_decl(ctx *Par_fun_symbol_declContext) {}

// ExitPar_fun_symbol_decl is called when production par_fun_symbol_decl is exited.
func (s *BaseSMTLIBv2Listener) ExitPar_fun_symbol_decl(ctx *Par_fun_symbol_declContext) {}

// EnterTheory_attribute is called when production theory_attribute is entered.
func (s *BaseSMTLIBv2Listener) EnterTheory_attribute(ctx *Theory_attributeContext) {}

// ExitTheory_attribute is called when production theory_attribute is exited.
func (s *BaseSMTLIBv2Listener) ExitTheory_attribute(ctx *Theory_attributeContext) {}

// EnterTheory_decl is called when production theory_decl is entered.
func (s *BaseSMTLIBv2Listener) EnterTheory_decl(ctx *Theory_declContext) {}

// ExitTheory_decl is called when production theory_decl is exited.
func (s *BaseSMTLIBv2Listener) ExitTheory_decl(ctx *Theory_declContext) {}

// EnterLogic_attribue is called when production logic_attribue is entered.
func (s *BaseSMTLIBv2Listener) EnterLogic_attribue(ctx *Logic_attribueContext) {}

// ExitLogic_attribue is called when production logic_attribue is exited.
func (s *BaseSMTLIBv2Listener) ExitLogic_attribue(ctx *Logic_attribueContext) {}

// EnterLogic is called when production logic is entered.
func (s *BaseSMTLIBv2Listener) EnterLogic(ctx *LogicContext) {}

// ExitLogic is called when production logic is exited.
func (s *BaseSMTLIBv2Listener) ExitLogic(ctx *LogicContext) {}

// EnterSort_dec is called when production sort_dec is entered.
func (s *BaseSMTLIBv2Listener) EnterSort_dec(ctx *Sort_decContext) {}

// ExitSort_dec is called when production sort_dec is exited.
func (s *BaseSMTLIBv2Listener) ExitSort_dec(ctx *Sort_decContext) {}

// EnterSelector_dec is called when production selector_dec is entered.
func (s *BaseSMTLIBv2Listener) EnterSelector_dec(ctx *Selector_decContext) {}

// ExitSelector_dec is called when production selector_dec is exited.
func (s *BaseSMTLIBv2Listener) ExitSelector_dec(ctx *Selector_decContext) {}

// EnterConstructor_dec is called when production constructor_dec is entered.
func (s *BaseSMTLIBv2Listener) EnterConstructor_dec(ctx *Constructor_decContext) {}

// ExitConstructor_dec is called when production constructor_dec is exited.
func (s *BaseSMTLIBv2Listener) ExitConstructor_dec(ctx *Constructor_decContext) {}

// EnterDatatype_dec is called when production datatype_dec is entered.
func (s *BaseSMTLIBv2Listener) EnterDatatype_dec(ctx *Datatype_decContext) {}

// ExitDatatype_dec is called when production datatype_dec is exited.
func (s *BaseSMTLIBv2Listener) ExitDatatype_dec(ctx *Datatype_decContext) {}

// EnterFunction_dec is called when production function_dec is entered.
func (s *BaseSMTLIBv2Listener) EnterFunction_dec(ctx *Function_decContext) {}

// ExitFunction_dec is called when production function_dec is exited.
func (s *BaseSMTLIBv2Listener) ExitFunction_dec(ctx *Function_decContext) {}

// EnterFunction_def is called when production function_def is entered.
func (s *BaseSMTLIBv2Listener) EnterFunction_def(ctx *Function_defContext) {}

// ExitFunction_def is called when production function_def is exited.
func (s *BaseSMTLIBv2Listener) ExitFunction_def(ctx *Function_defContext) {}

// EnterProp_literal is called when production prop_literal is entered.
func (s *BaseSMTLIBv2Listener) EnterProp_literal(ctx *Prop_literalContext) {}

// ExitProp_literal is called when production prop_literal is exited.
func (s *BaseSMTLIBv2Listener) ExitProp_literal(ctx *Prop_literalContext) {}

// EnterScript is called when production script is entered.
func (s *BaseSMTLIBv2Listener) EnterScript(ctx *ScriptContext) {}

// ExitScript is called when production script is exited.
func (s *BaseSMTLIBv2Listener) ExitScript(ctx *ScriptContext) {}

// EnterCmd_assert is called when production cmd_assert is entered.
func (s *BaseSMTLIBv2Listener) EnterCmd_assert(ctx *Cmd_assertContext) {}

// ExitCmd_assert is called when production cmd_assert is exited.
func (s *BaseSMTLIBv2Listener) ExitCmd_assert(ctx *Cmd_assertContext) {}

// EnterCmd_checkSat is called when production cmd_checkSat is entered.
func (s *BaseSMTLIBv2Listener) EnterCmd_checkSat(ctx *Cmd_checkSatContext) {}

// ExitCmd_checkSat is called when production cmd_checkSat is exited.
func (s *BaseSMTLIBv2Listener) ExitCmd_checkSat(ctx *Cmd_checkSatContext) {}

// EnterCmd_checkSatAssuming is called when production cmd_checkSatAssuming is entered.
func (s *BaseSMTLIBv2Listener) EnterCmd_checkSatAssuming(ctx *Cmd_checkSatAssumingContext) {}

// ExitCmd_checkSatAssuming is called when production cmd_checkSatAssuming is exited.
func (s *BaseSMTLIBv2Listener) ExitCmd_checkSatAssuming(ctx *Cmd_checkSatAssumingContext) {}

// EnterCmd_declareConst is called when production cmd_declareConst is entered.
func (s *BaseSMTLIBv2Listener) EnterCmd_declareConst(ctx *Cmd_declareConstContext) {}

// ExitCmd_declareConst is called when production cmd_declareConst is exited.
func (s *BaseSMTLIBv2Listener) ExitCmd_declareConst(ctx *Cmd_declareConstContext) {}

// EnterCmd_declareDatatype is called when production cmd_declareDatatype is entered.
func (s *BaseSMTLIBv2Listener) EnterCmd_declareDatatype(ctx *Cmd_declareDatatypeContext) {}

// ExitCmd_declareDatatype is called when production cmd_declareDatatype is exited.
func (s *BaseSMTLIBv2Listener) ExitCmd_declareDatatype(ctx *Cmd_declareDatatypeContext) {}

// EnterCmd_declareDatatypes is called when production cmd_declareDatatypes is entered.
func (s *BaseSMTLIBv2Listener) EnterCmd_declareDatatypes(ctx *Cmd_declareDatatypesContext) {}

// ExitCmd_declareDatatypes is called when production cmd_declareDatatypes is exited.
func (s *BaseSMTLIBv2Listener) ExitCmd_declareDatatypes(ctx *Cmd_declareDatatypesContext) {}

// EnterCmd_declareFun is called when production cmd_declareFun is entered.
func (s *BaseSMTLIBv2Listener) EnterCmd_declareFun(ctx *Cmd_declareFunContext) {}

// ExitCmd_declareFun is called when production cmd_declareFun is exited.
func (s *BaseSMTLIBv2Listener) ExitCmd_declareFun(ctx *Cmd_declareFunContext) {}

// EnterCmd_declareSort is called when production cmd_declareSort is entered.
func (s *BaseSMTLIBv2Listener) EnterCmd_declareSort(ctx *Cmd_declareSortContext) {}

// ExitCmd_declareSort is called when production cmd_declareSort is exited.
func (s *BaseSMTLIBv2Listener) ExitCmd_declareSort(ctx *Cmd_declareSortContext) {}

// EnterCmd_defineFun is called when production cmd_defineFun is entered.
func (s *BaseSMTLIBv2Listener) EnterCmd_defineFun(ctx *Cmd_defineFunContext) {}

// ExitCmd_defineFun is called when production cmd_defineFun is exited.
func (s *BaseSMTLIBv2Listener) ExitCmd_defineFun(ctx *Cmd_defineFunContext) {}

// EnterCmd_defineFunRec is called when production cmd_defineFunRec is entered.
func (s *BaseSMTLIBv2Listener) EnterCmd_defineFunRec(ctx *Cmd_defineFunRecContext) {}

// ExitCmd_defineFunRec is called when production cmd_defineFunRec is exited.
func (s *BaseSMTLIBv2Listener) ExitCmd_defineFunRec(ctx *Cmd_defineFunRecContext) {}

// EnterCmd_defineFunsRec is called when production cmd_defineFunsRec is entered.
func (s *BaseSMTLIBv2Listener) EnterCmd_defineFunsRec(ctx *Cmd_defineFunsRecContext) {}

// ExitCmd_defineFunsRec is called when production cmd_defineFunsRec is exited.
func (s *BaseSMTLIBv2Listener) ExitCmd_defineFunsRec(ctx *Cmd_defineFunsRecContext) {}

// EnterCmd_defineSort is called when production cmd_defineSort is entered.
func (s *BaseSMTLIBv2Listener) EnterCmd_defineSort(ctx *Cmd_defineSortContext) {}

// ExitCmd_defineSort is called when production cmd_defineSort is exited.
func (s *BaseSMTLIBv2Listener) ExitCmd_defineSort(ctx *Cmd_defineSortContext) {}

// EnterCmd_echo is called when production cmd_echo is entered.
func (s *BaseSMTLIBv2Listener) EnterCmd_echo(ctx *Cmd_echoContext) {}

// ExitCmd_echo is called when production cmd_echo is exited.
func (s *BaseSMTLIBv2Listener) ExitCmd_echo(ctx *Cmd_echoContext) {}

// EnterCmd_exit is called when production cmd_exit is entered.
func (s *BaseSMTLIBv2Listener) EnterCmd_exit(ctx *Cmd_exitContext) {}

// ExitCmd_exit is called when production cmd_exit is exited.
func (s *BaseSMTLIBv2Listener) ExitCmd_exit(ctx *Cmd_exitContext) {}

// EnterCmd_getAssertions is called when production cmd_getAssertions is entered.
func (s *BaseSMTLIBv2Listener) EnterCmd_getAssertions(ctx *Cmd_getAssertionsContext) {}

// ExitCmd_getAssertions is called when production cmd_getAssertions is exited.
func (s *BaseSMTLIBv2Listener) ExitCmd_getAssertions(ctx *Cmd_getAssertionsContext) {}

// EnterCmd_getAssignment is called when production cmd_getAssignment is entered.
func (s *BaseSMTLIBv2Listener) EnterCmd_getAssignment(ctx *Cmd_getAssignmentContext) {}

// ExitCmd_getAssignment is called when production cmd_getAssignment is exited.
func (s *BaseSMTLIBv2Listener) ExitCmd_getAssignment(ctx *Cmd_getAssignmentContext) {}

// EnterCmd_getInfo is called when production cmd_getInfo is entered.
func (s *BaseSMTLIBv2Listener) EnterCmd_getInfo(ctx *Cmd_getInfoContext) {}

// ExitCmd_getInfo is called when production cmd_getInfo is exited.
func (s *BaseSMTLIBv2Listener) ExitCmd_getInfo(ctx *Cmd_getInfoContext) {}

// EnterCmd_getModel is called when production cmd_getModel is entered.
func (s *BaseSMTLIBv2Listener) EnterCmd_getModel(ctx *Cmd_getModelContext) {}

// ExitCmd_getModel is called when production cmd_getModel is exited.
func (s *BaseSMTLIBv2Listener) ExitCmd_getModel(ctx *Cmd_getModelContext) {}

// EnterCmd_getOption is called when production cmd_getOption is entered.
func (s *BaseSMTLIBv2Listener) EnterCmd_getOption(ctx *Cmd_getOptionContext) {}

// ExitCmd_getOption is called when production cmd_getOption is exited.
func (s *BaseSMTLIBv2Listener) ExitCmd_getOption(ctx *Cmd_getOptionContext) {}

// EnterCmd_getProof is called when production cmd_getProof is entered.
func (s *BaseSMTLIBv2Listener) EnterCmd_getProof(ctx *Cmd_getProofContext) {}

// ExitCmd_getProof is called when production cmd_getProof is exited.
func (s *BaseSMTLIBv2Listener) ExitCmd_getProof(ctx *Cmd_getProofContext) {}

// EnterCmd_getUnsatAssumptions is called when production cmd_getUnsatAssumptions is entered.
func (s *BaseSMTLIBv2Listener) EnterCmd_getUnsatAssumptions(ctx *Cmd_getUnsatAssumptionsContext) {}

// ExitCmd_getUnsatAssumptions is called when production cmd_getUnsatAssumptions is exited.
func (s *BaseSMTLIBv2Listener) ExitCmd_getUnsatAssumptions(ctx *Cmd_getUnsatAssumptionsContext) {}

// EnterCmd_getUnsatCore is called when production cmd_getUnsatCore is entered.
func (s *BaseSMTLIBv2Listener) EnterCmd_getUnsatCore(ctx *Cmd_getUnsatCoreContext) {}

// ExitCmd_getUnsatCore is called when production cmd_getUnsatCore is exited.
func (s *BaseSMTLIBv2Listener) ExitCmd_getUnsatCore(ctx *Cmd_getUnsatCoreContext) {}

// EnterCmd_getValue is called when production cmd_getValue is entered.
func (s *BaseSMTLIBv2Listener) EnterCmd_getValue(ctx *Cmd_getValueContext) {}

// ExitCmd_getValue is called when production cmd_getValue is exited.
func (s *BaseSMTLIBv2Listener) ExitCmd_getValue(ctx *Cmd_getValueContext) {}

// EnterCmd_pop is called when production cmd_pop is entered.
func (s *BaseSMTLIBv2Listener) EnterCmd_pop(ctx *Cmd_popContext) {}

// ExitCmd_pop is called when production cmd_pop is exited.
func (s *BaseSMTLIBv2Listener) ExitCmd_pop(ctx *Cmd_popContext) {}

// EnterCmd_push is called when production cmd_push is entered.
func (s *BaseSMTLIBv2Listener) EnterCmd_push(ctx *Cmd_pushContext) {}

// ExitCmd_push is called when production cmd_push is exited.
func (s *BaseSMTLIBv2Listener) ExitCmd_push(ctx *Cmd_pushContext) {}

// EnterCmd_reset is called when production cmd_reset is entered.
func (s *BaseSMTLIBv2Listener) EnterCmd_reset(ctx *Cmd_resetContext) {}

// ExitCmd_reset is called when production cmd_reset is exited.
func (s *BaseSMTLIBv2Listener) ExitCmd_reset(ctx *Cmd_resetContext) {}

// EnterCmd_resetAssertions is called when production cmd_resetAssertions is entered.
func (s *BaseSMTLIBv2Listener) EnterCmd_resetAssertions(ctx *Cmd_resetAssertionsContext) {}

// ExitCmd_resetAssertions is called when production cmd_resetAssertions is exited.
func (s *BaseSMTLIBv2Listener) ExitCmd_resetAssertions(ctx *Cmd_resetAssertionsContext) {}

// EnterCmd_setInfo is called when production cmd_setInfo is entered.
func (s *BaseSMTLIBv2Listener) EnterCmd_setInfo(ctx *Cmd_setInfoContext) {}

// ExitCmd_setInfo is called when production cmd_setInfo is exited.
func (s *BaseSMTLIBv2Listener) ExitCmd_setInfo(ctx *Cmd_setInfoContext) {}

// EnterCmd_setLogic is called when production cmd_setLogic is entered.
func (s *BaseSMTLIBv2Listener) EnterCmd_setLogic(ctx *Cmd_setLogicContext) {}

// ExitCmd_setLogic is called when production cmd_setLogic is exited.
func (s *BaseSMTLIBv2Listener) ExitCmd_setLogic(ctx *Cmd_setLogicContext) {}

// EnterCmd_setOption is called when production cmd_setOption is entered.
func (s *BaseSMTLIBv2Listener) EnterCmd_setOption(ctx *Cmd_setOptionContext) {}

// ExitCmd_setOption is called when production cmd_setOption is exited.
func (s *BaseSMTLIBv2Listener) ExitCmd_setOption(ctx *Cmd_setOptionContext) {}

// EnterCommand is called when production command is entered.
func (s *BaseSMTLIBv2Listener) EnterCommand(ctx *CommandContext) {}

// ExitCommand is called when production command is exited.
func (s *BaseSMTLIBv2Listener) ExitCommand(ctx *CommandContext) {}

// EnterB_value is called when production b_value is entered.
func (s *BaseSMTLIBv2Listener) EnterB_value(ctx *B_valueContext) {}

// ExitB_value is called when production b_value is exited.
func (s *BaseSMTLIBv2Listener) ExitB_value(ctx *B_valueContext) {}

// EnterOption is called when production option is entered.
func (s *BaseSMTLIBv2Listener) EnterOption(ctx *OptionContext) {}

// ExitOption is called when production option is exited.
func (s *BaseSMTLIBv2Listener) ExitOption(ctx *OptionContext) {}

// EnterInfo_flag is called when production info_flag is entered.
func (s *BaseSMTLIBv2Listener) EnterInfo_flag(ctx *Info_flagContext) {}

// ExitInfo_flag is called when production info_flag is exited.
func (s *BaseSMTLIBv2Listener) ExitInfo_flag(ctx *Info_flagContext) {}

// EnterError_behaviour is called when production error_behaviour is entered.
func (s *BaseSMTLIBv2Listener) EnterError_behaviour(ctx *Error_behaviourContext) {}

// ExitError_behaviour is called when production error_behaviour is exited.
func (s *BaseSMTLIBv2Listener) ExitError_behaviour(ctx *Error_behaviourContext) {}

// EnterReason_unknown is called when production reason_unknown is entered.
func (s *BaseSMTLIBv2Listener) EnterReason_unknown(ctx *Reason_unknownContext) {}

// ExitReason_unknown is called when production reason_unknown is exited.
func (s *BaseSMTLIBv2Listener) ExitReason_unknown(ctx *Reason_unknownContext) {}

// EnterModel_response is called when production model_response is entered.
func (s *BaseSMTLIBv2Listener) EnterModel_response(ctx *Model_responseContext) {}

// ExitModel_response is called when production model_response is exited.
func (s *BaseSMTLIBv2Listener) ExitModel_response(ctx *Model_responseContext) {}

// EnterInfo_response is called when production info_response is entered.
func (s *BaseSMTLIBv2Listener) EnterInfo_response(ctx *Info_responseContext) {}

// ExitInfo_response is called when production info_response is exited.
func (s *BaseSMTLIBv2Listener) ExitInfo_response(ctx *Info_responseContext) {}

// EnterValuation_pair is called when production valuation_pair is entered.
func (s *BaseSMTLIBv2Listener) EnterValuation_pair(ctx *Valuation_pairContext) {}

// ExitValuation_pair is called when production valuation_pair is exited.
func (s *BaseSMTLIBv2Listener) ExitValuation_pair(ctx *Valuation_pairContext) {}

// EnterT_valuation_pair is called when production t_valuation_pair is entered.
func (s *BaseSMTLIBv2Listener) EnterT_valuation_pair(ctx *T_valuation_pairContext) {}

// ExitT_valuation_pair is called when production t_valuation_pair is exited.
func (s *BaseSMTLIBv2Listener) ExitT_valuation_pair(ctx *T_valuation_pairContext) {}

// EnterCheck_sat_response is called when production check_sat_response is entered.
func (s *BaseSMTLIBv2Listener) EnterCheck_sat_response(ctx *Check_sat_responseContext) {}

// ExitCheck_sat_response is called when production check_sat_response is exited.
func (s *BaseSMTLIBv2Listener) ExitCheck_sat_response(ctx *Check_sat_responseContext) {}

// EnterEcho_response is called when production echo_response is entered.
func (s *BaseSMTLIBv2Listener) EnterEcho_response(ctx *Echo_responseContext) {}

// ExitEcho_response is called when production echo_response is exited.
func (s *BaseSMTLIBv2Listener) ExitEcho_response(ctx *Echo_responseContext) {}

// EnterGet_assertions_response is called when production get_assertions_response is entered.
func (s *BaseSMTLIBv2Listener) EnterGet_assertions_response(ctx *Get_assertions_responseContext) {}

// ExitGet_assertions_response is called when production get_assertions_response is exited.
func (s *BaseSMTLIBv2Listener) ExitGet_assertions_response(ctx *Get_assertions_responseContext) {}

// EnterGet_assignment_response is called when production get_assignment_response is entered.
func (s *BaseSMTLIBv2Listener) EnterGet_assignment_response(ctx *Get_assignment_responseContext) {}

// ExitGet_assignment_response is called when production get_assignment_response is exited.
func (s *BaseSMTLIBv2Listener) ExitGet_assignment_response(ctx *Get_assignment_responseContext) {}

// EnterGet_info_response is called when production get_info_response is entered.
func (s *BaseSMTLIBv2Listener) EnterGet_info_response(ctx *Get_info_responseContext) {}

// ExitGet_info_response is called when production get_info_response is exited.
func (s *BaseSMTLIBv2Listener) ExitGet_info_response(ctx *Get_info_responseContext) {}

// EnterGet_model_response is called when production get_model_response is entered.
func (s *BaseSMTLIBv2Listener) EnterGet_model_response(ctx *Get_model_responseContext) {}

// ExitGet_model_response is called when production get_model_response is exited.
func (s *BaseSMTLIBv2Listener) ExitGet_model_response(ctx *Get_model_responseContext) {}

// EnterGet_option_response is called when production get_option_response is entered.
func (s *BaseSMTLIBv2Listener) EnterGet_option_response(ctx *Get_option_responseContext) {}

// ExitGet_option_response is called when production get_option_response is exited.
func (s *BaseSMTLIBv2Listener) ExitGet_option_response(ctx *Get_option_responseContext) {}

// EnterGet_proof_response is called when production get_proof_response is entered.
func (s *BaseSMTLIBv2Listener) EnterGet_proof_response(ctx *Get_proof_responseContext) {}

// ExitGet_proof_response is called when production get_proof_response is exited.
func (s *BaseSMTLIBv2Listener) ExitGet_proof_response(ctx *Get_proof_responseContext) {}

// EnterGet_unsat_assump_response is called when production get_unsat_assump_response is entered.
func (s *BaseSMTLIBv2Listener) EnterGet_unsat_assump_response(ctx *Get_unsat_assump_responseContext) {
}

// ExitGet_unsat_assump_response is called when production get_unsat_assump_response is exited.
func (s *BaseSMTLIBv2Listener) ExitGet_unsat_assump_response(ctx *Get_unsat_assump_responseContext) {}

// EnterGet_unsat_core_response is called when production get_unsat_core_response is entered.
func (s *BaseSMTLIBv2Listener) EnterGet_unsat_core_response(ctx *Get_unsat_core_responseContext) {}

// ExitGet_unsat_core_response is called when production get_unsat_core_response is exited.
func (s *BaseSMTLIBv2Listener) ExitGet_unsat_core_response(ctx *Get_unsat_core_responseContext) {}

// EnterGet_value_response is called when production get_value_response is entered.
func (s *BaseSMTLIBv2Listener) EnterGet_value_response(ctx *Get_value_responseContext) {}

// ExitGet_value_response is called when production get_value_response is exited.
func (s *BaseSMTLIBv2Listener) ExitGet_value_response(ctx *Get_value_responseContext) {}

// EnterSpecific_success_response is called when production specific_success_response is entered.
func (s *BaseSMTLIBv2Listener) EnterSpecific_success_response(ctx *Specific_success_responseContext) {
}

// ExitSpecific_success_response is called when production specific_success_response is exited.
func (s *BaseSMTLIBv2Listener) ExitSpecific_success_response(ctx *Specific_success_responseContext) {}

// EnterGeneral_response is called when production general_response is entered.
func (s *BaseSMTLIBv2Listener) EnterGeneral_response(ctx *General_responseContext) {}

// ExitGeneral_response is called when production general_response is exited.
func (s *BaseSMTLIBv2Listener) ExitGeneral_response(ctx *General_responseContext) {}
