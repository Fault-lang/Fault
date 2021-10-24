// Code generated from SMTLIBv2.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // SMTLIBv2

import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by SMTLIBv2Parser.
type SMTLIBv2Visitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by SMTLIBv2Parser#start.
	VisitStart(ctx *StartContext) interface{}

	// Visit a parse tree produced by SMTLIBv2Parser#generalReservedWord.
	VisitGeneralReservedWord(ctx *GeneralReservedWordContext) interface{}

	// Visit a parse tree produced by SMTLIBv2Parser#simpleSymbol.
	VisitSimpleSymbol(ctx *SimpleSymbolContext) interface{}

	// Visit a parse tree produced by SMTLIBv2Parser#quotedSymbol.
	VisitQuotedSymbol(ctx *QuotedSymbolContext) interface{}

	// Visit a parse tree produced by SMTLIBv2Parser#predefSymbol.
	VisitPredefSymbol(ctx *PredefSymbolContext) interface{}

	// Visit a parse tree produced by SMTLIBv2Parser#predefKeyword.
	VisitPredefKeyword(ctx *PredefKeywordContext) interface{}

	// Visit a parse tree produced by SMTLIBv2Parser#symbol.
	VisitSymbol(ctx *SymbolContext) interface{}

	// Visit a parse tree produced by SMTLIBv2Parser#numeral.
	VisitNumeral(ctx *NumeralContext) interface{}

	// Visit a parse tree produced by SMTLIBv2Parser#decimal.
	VisitDecimal(ctx *DecimalContext) interface{}

	// Visit a parse tree produced by SMTLIBv2Parser#hexadecimal.
	VisitHexadecimal(ctx *HexadecimalContext) interface{}

	// Visit a parse tree produced by SMTLIBv2Parser#binary.
	VisitBinary(ctx *BinaryContext) interface{}

	// Visit a parse tree produced by SMTLIBv2Parser#string_.
	VisitString_(ctx *String_Context) interface{}

	// Visit a parse tree produced by SMTLIBv2Parser#keyword.
	VisitKeyword(ctx *KeywordContext) interface{}

	// Visit a parse tree produced by SMTLIBv2Parser#spec_constant.
	VisitSpec_constant(ctx *Spec_constantContext) interface{}

	// Visit a parse tree produced by SMTLIBv2Parser#s_expr.
	VisitS_expr(ctx *S_exprContext) interface{}

	// Visit a parse tree produced by SMTLIBv2Parser#index.
	VisitIndex(ctx *IndexContext) interface{}

	// Visit a parse tree produced by SMTLIBv2Parser#identifier.
	VisitIdentifier(ctx *IdentifierContext) interface{}

	// Visit a parse tree produced by SMTLIBv2Parser#attribute_value.
	VisitAttribute_value(ctx *Attribute_valueContext) interface{}

	// Visit a parse tree produced by SMTLIBv2Parser#attribute.
	VisitAttribute(ctx *AttributeContext) interface{}

	// Visit a parse tree produced by SMTLIBv2Parser#sort.
	VisitSort(ctx *SortContext) interface{}

	// Visit a parse tree produced by SMTLIBv2Parser#qual_identifer.
	VisitQual_identifer(ctx *Qual_identiferContext) interface{}

	// Visit a parse tree produced by SMTLIBv2Parser#var_binding.
	VisitVar_binding(ctx *Var_bindingContext) interface{}

	// Visit a parse tree produced by SMTLIBv2Parser#sorted_var.
	VisitSorted_var(ctx *Sorted_varContext) interface{}

	// Visit a parse tree produced by SMTLIBv2Parser#pattern.
	VisitPattern(ctx *PatternContext) interface{}

	// Visit a parse tree produced by SMTLIBv2Parser#match_case.
	VisitMatch_case(ctx *Match_caseContext) interface{}

	// Visit a parse tree produced by SMTLIBv2Parser#term.
	VisitTerm(ctx *TermContext) interface{}

	// Visit a parse tree produced by SMTLIBv2Parser#sort_symbol_decl.
	VisitSort_symbol_decl(ctx *Sort_symbol_declContext) interface{}

	// Visit a parse tree produced by SMTLIBv2Parser#meta_spec_constant.
	VisitMeta_spec_constant(ctx *Meta_spec_constantContext) interface{}

	// Visit a parse tree produced by SMTLIBv2Parser#fun_symbol_decl.
	VisitFun_symbol_decl(ctx *Fun_symbol_declContext) interface{}

	// Visit a parse tree produced by SMTLIBv2Parser#par_fun_symbol_decl.
	VisitPar_fun_symbol_decl(ctx *Par_fun_symbol_declContext) interface{}

	// Visit a parse tree produced by SMTLIBv2Parser#theory_attribute.
	VisitTheory_attribute(ctx *Theory_attributeContext) interface{}

	// Visit a parse tree produced by SMTLIBv2Parser#theory_decl.
	VisitTheory_decl(ctx *Theory_declContext) interface{}

	// Visit a parse tree produced by SMTLIBv2Parser#logic_attribue.
	VisitLogic_attribue(ctx *Logic_attribueContext) interface{}

	// Visit a parse tree produced by SMTLIBv2Parser#logic.
	VisitLogic(ctx *LogicContext) interface{}

	// Visit a parse tree produced by SMTLIBv2Parser#sort_dec.
	VisitSort_dec(ctx *Sort_decContext) interface{}

	// Visit a parse tree produced by SMTLIBv2Parser#selector_dec.
	VisitSelector_dec(ctx *Selector_decContext) interface{}

	// Visit a parse tree produced by SMTLIBv2Parser#constructor_dec.
	VisitConstructor_dec(ctx *Constructor_decContext) interface{}

	// Visit a parse tree produced by SMTLIBv2Parser#datatype_dec.
	VisitDatatype_dec(ctx *Datatype_decContext) interface{}

	// Visit a parse tree produced by SMTLIBv2Parser#function_dec.
	VisitFunction_dec(ctx *Function_decContext) interface{}

	// Visit a parse tree produced by SMTLIBv2Parser#function_def.
	VisitFunction_def(ctx *Function_defContext) interface{}

	// Visit a parse tree produced by SMTLIBv2Parser#prop_literal.
	VisitProp_literal(ctx *Prop_literalContext) interface{}

	// Visit a parse tree produced by SMTLIBv2Parser#script.
	VisitScript(ctx *ScriptContext) interface{}

	// Visit a parse tree produced by SMTLIBv2Parser#cmd_assert.
	VisitCmd_assert(ctx *Cmd_assertContext) interface{}

	// Visit a parse tree produced by SMTLIBv2Parser#cmd_checkSat.
	VisitCmd_checkSat(ctx *Cmd_checkSatContext) interface{}

	// Visit a parse tree produced by SMTLIBv2Parser#cmd_checkSatAssuming.
	VisitCmd_checkSatAssuming(ctx *Cmd_checkSatAssumingContext) interface{}

	// Visit a parse tree produced by SMTLIBv2Parser#cmd_declareConst.
	VisitCmd_declareConst(ctx *Cmd_declareConstContext) interface{}

	// Visit a parse tree produced by SMTLIBv2Parser#cmd_declareDatatype.
	VisitCmd_declareDatatype(ctx *Cmd_declareDatatypeContext) interface{}

	// Visit a parse tree produced by SMTLIBv2Parser#cmd_declareDatatypes.
	VisitCmd_declareDatatypes(ctx *Cmd_declareDatatypesContext) interface{}

	// Visit a parse tree produced by SMTLIBv2Parser#cmd_declareFun.
	VisitCmd_declareFun(ctx *Cmd_declareFunContext) interface{}

	// Visit a parse tree produced by SMTLIBv2Parser#cmd_declareSort.
	VisitCmd_declareSort(ctx *Cmd_declareSortContext) interface{}

	// Visit a parse tree produced by SMTLIBv2Parser#cmd_defineFun.
	VisitCmd_defineFun(ctx *Cmd_defineFunContext) interface{}

	// Visit a parse tree produced by SMTLIBv2Parser#cmd_defineFunRec.
	VisitCmd_defineFunRec(ctx *Cmd_defineFunRecContext) interface{}

	// Visit a parse tree produced by SMTLIBv2Parser#cmd_defineFunsRec.
	VisitCmd_defineFunsRec(ctx *Cmd_defineFunsRecContext) interface{}

	// Visit a parse tree produced by SMTLIBv2Parser#cmd_defineSort.
	VisitCmd_defineSort(ctx *Cmd_defineSortContext) interface{}

	// Visit a parse tree produced by SMTLIBv2Parser#cmd_echo.
	VisitCmd_echo(ctx *Cmd_echoContext) interface{}

	// Visit a parse tree produced by SMTLIBv2Parser#cmd_exit.
	VisitCmd_exit(ctx *Cmd_exitContext) interface{}

	// Visit a parse tree produced by SMTLIBv2Parser#cmd_getAssertions.
	VisitCmd_getAssertions(ctx *Cmd_getAssertionsContext) interface{}

	// Visit a parse tree produced by SMTLIBv2Parser#cmd_getAssignment.
	VisitCmd_getAssignment(ctx *Cmd_getAssignmentContext) interface{}

	// Visit a parse tree produced by SMTLIBv2Parser#cmd_getInfo.
	VisitCmd_getInfo(ctx *Cmd_getInfoContext) interface{}

	// Visit a parse tree produced by SMTLIBv2Parser#cmd_getModel.
	VisitCmd_getModel(ctx *Cmd_getModelContext) interface{}

	// Visit a parse tree produced by SMTLIBv2Parser#cmd_getOption.
	VisitCmd_getOption(ctx *Cmd_getOptionContext) interface{}

	// Visit a parse tree produced by SMTLIBv2Parser#cmd_getProof.
	VisitCmd_getProof(ctx *Cmd_getProofContext) interface{}

	// Visit a parse tree produced by SMTLIBv2Parser#cmd_getUnsatAssumptions.
	VisitCmd_getUnsatAssumptions(ctx *Cmd_getUnsatAssumptionsContext) interface{}

	// Visit a parse tree produced by SMTLIBv2Parser#cmd_getUnsatCore.
	VisitCmd_getUnsatCore(ctx *Cmd_getUnsatCoreContext) interface{}

	// Visit a parse tree produced by SMTLIBv2Parser#cmd_getValue.
	VisitCmd_getValue(ctx *Cmd_getValueContext) interface{}

	// Visit a parse tree produced by SMTLIBv2Parser#cmd_pop.
	VisitCmd_pop(ctx *Cmd_popContext) interface{}

	// Visit a parse tree produced by SMTLIBv2Parser#cmd_push.
	VisitCmd_push(ctx *Cmd_pushContext) interface{}

	// Visit a parse tree produced by SMTLIBv2Parser#cmd_reset.
	VisitCmd_reset(ctx *Cmd_resetContext) interface{}

	// Visit a parse tree produced by SMTLIBv2Parser#cmd_resetAssertions.
	VisitCmd_resetAssertions(ctx *Cmd_resetAssertionsContext) interface{}

	// Visit a parse tree produced by SMTLIBv2Parser#cmd_setInfo.
	VisitCmd_setInfo(ctx *Cmd_setInfoContext) interface{}

	// Visit a parse tree produced by SMTLIBv2Parser#cmd_setLogic.
	VisitCmd_setLogic(ctx *Cmd_setLogicContext) interface{}

	// Visit a parse tree produced by SMTLIBv2Parser#cmd_setOption.
	VisitCmd_setOption(ctx *Cmd_setOptionContext) interface{}

	// Visit a parse tree produced by SMTLIBv2Parser#command.
	VisitCommand(ctx *CommandContext) interface{}

	// Visit a parse tree produced by SMTLIBv2Parser#b_value.
	VisitB_value(ctx *B_valueContext) interface{}

	// Visit a parse tree produced by SMTLIBv2Parser#option.
	VisitOption(ctx *OptionContext) interface{}

	// Visit a parse tree produced by SMTLIBv2Parser#info_flag.
	VisitInfo_flag(ctx *Info_flagContext) interface{}

	// Visit a parse tree produced by SMTLIBv2Parser#error_behaviour.
	VisitError_behaviour(ctx *Error_behaviourContext) interface{}

	// Visit a parse tree produced by SMTLIBv2Parser#reason_unknown.
	VisitReason_unknown(ctx *Reason_unknownContext) interface{}

	// Visit a parse tree produced by SMTLIBv2Parser#model_response.
	VisitModel_response(ctx *Model_responseContext) interface{}

	// Visit a parse tree produced by SMTLIBv2Parser#info_response.
	VisitInfo_response(ctx *Info_responseContext) interface{}

	// Visit a parse tree produced by SMTLIBv2Parser#valuation_pair.
	VisitValuation_pair(ctx *Valuation_pairContext) interface{}

	// Visit a parse tree produced by SMTLIBv2Parser#t_valuation_pair.
	VisitT_valuation_pair(ctx *T_valuation_pairContext) interface{}

	// Visit a parse tree produced by SMTLIBv2Parser#check_sat_response.
	VisitCheck_sat_response(ctx *Check_sat_responseContext) interface{}

	// Visit a parse tree produced by SMTLIBv2Parser#echo_response.
	VisitEcho_response(ctx *Echo_responseContext) interface{}

	// Visit a parse tree produced by SMTLIBv2Parser#get_assertions_response.
	VisitGet_assertions_response(ctx *Get_assertions_responseContext) interface{}

	// Visit a parse tree produced by SMTLIBv2Parser#get_assignment_response.
	VisitGet_assignment_response(ctx *Get_assignment_responseContext) interface{}

	// Visit a parse tree produced by SMTLIBv2Parser#get_info_response.
	VisitGet_info_response(ctx *Get_info_responseContext) interface{}

	// Visit a parse tree produced by SMTLIBv2Parser#get_model_response.
	VisitGet_model_response(ctx *Get_model_responseContext) interface{}

	// Visit a parse tree produced by SMTLIBv2Parser#get_option_response.
	VisitGet_option_response(ctx *Get_option_responseContext) interface{}

	// Visit a parse tree produced by SMTLIBv2Parser#get_proof_response.
	VisitGet_proof_response(ctx *Get_proof_responseContext) interface{}

	// Visit a parse tree produced by SMTLIBv2Parser#get_unsat_assump_response.
	VisitGet_unsat_assump_response(ctx *Get_unsat_assump_responseContext) interface{}

	// Visit a parse tree produced by SMTLIBv2Parser#get_unsat_core_response.
	VisitGet_unsat_core_response(ctx *Get_unsat_core_responseContext) interface{}

	// Visit a parse tree produced by SMTLIBv2Parser#get_value_response.
	VisitGet_value_response(ctx *Get_value_responseContext) interface{}

	// Visit a parse tree produced by SMTLIBv2Parser#specific_success_response.
	VisitSpecific_success_response(ctx *Specific_success_responseContext) interface{}

	// Visit a parse tree produced by SMTLIBv2Parser#general_response.
	VisitGeneral_response(ctx *General_responseContext) interface{}
}
