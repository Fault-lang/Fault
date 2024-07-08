// Code generated from SMTLIBv2.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // SMTLIBv2

import "github.com/antlr4-go/antlr/v4"

type BaseSMTLIBv2Visitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseSMTLIBv2Visitor) VisitStart(ctx *StartContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSMTLIBv2Visitor) VisitGeneralReservedWord(ctx *GeneralReservedWordContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSMTLIBv2Visitor) VisitSimpleSymbol(ctx *SimpleSymbolContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSMTLIBv2Visitor) VisitQuotedSymbol(ctx *QuotedSymbolContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSMTLIBv2Visitor) VisitPredefSymbol(ctx *PredefSymbolContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSMTLIBv2Visitor) VisitPredefKeyword(ctx *PredefKeywordContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSMTLIBv2Visitor) VisitSymbol(ctx *SymbolContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSMTLIBv2Visitor) VisitNumeral(ctx *NumeralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSMTLIBv2Visitor) VisitDecimal(ctx *DecimalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSMTLIBv2Visitor) VisitHexadecimal(ctx *HexadecimalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSMTLIBv2Visitor) VisitBinary(ctx *BinaryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSMTLIBv2Visitor) VisitString_(ctx *String_Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSMTLIBv2Visitor) VisitKeyword(ctx *KeywordContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSMTLIBv2Visitor) VisitSpec_constant(ctx *Spec_constantContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSMTLIBv2Visitor) VisitS_expr(ctx *S_exprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSMTLIBv2Visitor) VisitIndex(ctx *IndexContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSMTLIBv2Visitor) VisitIdentifier(ctx *IdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSMTLIBv2Visitor) VisitAttribute_value(ctx *Attribute_valueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSMTLIBv2Visitor) VisitAttribute(ctx *AttributeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSMTLIBv2Visitor) VisitSort(ctx *SortContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSMTLIBv2Visitor) VisitQual_identifer(ctx *Qual_identiferContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSMTLIBv2Visitor) VisitVar_binding(ctx *Var_bindingContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSMTLIBv2Visitor) VisitSorted_var(ctx *Sorted_varContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSMTLIBv2Visitor) VisitPattern(ctx *PatternContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSMTLIBv2Visitor) VisitMatch_case(ctx *Match_caseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSMTLIBv2Visitor) VisitVariable(ctx *VariableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSMTLIBv2Visitor) VisitTerm(ctx *TermContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSMTLIBv2Visitor) VisitSort_symbol_decl(ctx *Sort_symbol_declContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSMTLIBv2Visitor) VisitMeta_spec_constant(ctx *Meta_spec_constantContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSMTLIBv2Visitor) VisitFun_symbol_decl(ctx *Fun_symbol_declContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSMTLIBv2Visitor) VisitPar_fun_symbol_decl(ctx *Par_fun_symbol_declContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSMTLIBv2Visitor) VisitTheory_attribute(ctx *Theory_attributeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSMTLIBv2Visitor) VisitTheory_decl(ctx *Theory_declContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSMTLIBv2Visitor) VisitLogic_attribue(ctx *Logic_attribueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSMTLIBv2Visitor) VisitLogic(ctx *LogicContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSMTLIBv2Visitor) VisitSort_dec(ctx *Sort_decContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSMTLIBv2Visitor) VisitSelector_dec(ctx *Selector_decContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSMTLIBv2Visitor) VisitConstructor_dec(ctx *Constructor_decContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSMTLIBv2Visitor) VisitDatatype_dec(ctx *Datatype_decContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSMTLIBv2Visitor) VisitFunction_dec(ctx *Function_decContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSMTLIBv2Visitor) VisitFunction_def(ctx *Function_defContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSMTLIBv2Visitor) VisitProp_literal(ctx *Prop_literalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSMTLIBv2Visitor) VisitScript(ctx *ScriptContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSMTLIBv2Visitor) VisitCmd_assert(ctx *Cmd_assertContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSMTLIBv2Visitor) VisitCmd_checkSat(ctx *Cmd_checkSatContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSMTLIBv2Visitor) VisitCmd_checkSatAssuming(ctx *Cmd_checkSatAssumingContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSMTLIBv2Visitor) VisitCmd_declareConst(ctx *Cmd_declareConstContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSMTLIBv2Visitor) VisitCmd_declareDatatype(ctx *Cmd_declareDatatypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSMTLIBv2Visitor) VisitCmd_declareDatatypes(ctx *Cmd_declareDatatypesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSMTLIBv2Visitor) VisitCmd_declareFun(ctx *Cmd_declareFunContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSMTLIBv2Visitor) VisitCmd_declareSort(ctx *Cmd_declareSortContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSMTLIBv2Visitor) VisitCmd_defineFun(ctx *Cmd_defineFunContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSMTLIBv2Visitor) VisitCmd_defineFunRec(ctx *Cmd_defineFunRecContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSMTLIBv2Visitor) VisitCmd_defineFunsRec(ctx *Cmd_defineFunsRecContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSMTLIBv2Visitor) VisitCmd_defineSort(ctx *Cmd_defineSortContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSMTLIBv2Visitor) VisitCmd_echo(ctx *Cmd_echoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSMTLIBv2Visitor) VisitCmd_exit(ctx *Cmd_exitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSMTLIBv2Visitor) VisitCmd_getAssertions(ctx *Cmd_getAssertionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSMTLIBv2Visitor) VisitCmd_getAssignment(ctx *Cmd_getAssignmentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSMTLIBv2Visitor) VisitCmd_getInfo(ctx *Cmd_getInfoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSMTLIBv2Visitor) VisitCmd_getModel(ctx *Cmd_getModelContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSMTLIBv2Visitor) VisitCmd_getOption(ctx *Cmd_getOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSMTLIBv2Visitor) VisitCmd_getProof(ctx *Cmd_getProofContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSMTLIBv2Visitor) VisitCmd_getUnsatAssumptions(ctx *Cmd_getUnsatAssumptionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSMTLIBv2Visitor) VisitCmd_getUnsatCore(ctx *Cmd_getUnsatCoreContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSMTLIBv2Visitor) VisitCmd_getValue(ctx *Cmd_getValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSMTLIBv2Visitor) VisitCmd_pop(ctx *Cmd_popContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSMTLIBv2Visitor) VisitCmd_push(ctx *Cmd_pushContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSMTLIBv2Visitor) VisitCmd_reset(ctx *Cmd_resetContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSMTLIBv2Visitor) VisitCmd_resetAssertions(ctx *Cmd_resetAssertionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSMTLIBv2Visitor) VisitCmd_setInfo(ctx *Cmd_setInfoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSMTLIBv2Visitor) VisitCmd_setLogic(ctx *Cmd_setLogicContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSMTLIBv2Visitor) VisitCmd_setOption(ctx *Cmd_setOptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSMTLIBv2Visitor) VisitCommand(ctx *CommandContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSMTLIBv2Visitor) VisitB_value(ctx *B_valueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSMTLIBv2Visitor) VisitOption(ctx *OptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSMTLIBv2Visitor) VisitInfo_flag(ctx *Info_flagContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSMTLIBv2Visitor) VisitError_behaviour(ctx *Error_behaviourContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSMTLIBv2Visitor) VisitReason_unknown(ctx *Reason_unknownContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSMTLIBv2Visitor) VisitModel_response(ctx *Model_responseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSMTLIBv2Visitor) VisitInfo_response(ctx *Info_responseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSMTLIBv2Visitor) VisitValuation_pair(ctx *Valuation_pairContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSMTLIBv2Visitor) VisitT_valuation_pair(ctx *T_valuation_pairContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSMTLIBv2Visitor) VisitCheck_sat_response(ctx *Check_sat_responseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSMTLIBv2Visitor) VisitEcho_response(ctx *Echo_responseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSMTLIBv2Visitor) VisitGet_assertions_response(ctx *Get_assertions_responseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSMTLIBv2Visitor) VisitGet_assignment_response(ctx *Get_assignment_responseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSMTLIBv2Visitor) VisitGet_info_response(ctx *Get_info_responseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSMTLIBv2Visitor) VisitGet_model_response(ctx *Get_model_responseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSMTLIBv2Visitor) VisitGet_option_response(ctx *Get_option_responseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSMTLIBv2Visitor) VisitGet_proof_response(ctx *Get_proof_responseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSMTLIBv2Visitor) VisitGet_unsat_assump_response(ctx *Get_unsat_assump_responseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSMTLIBv2Visitor) VisitGet_unsat_core_response(ctx *Get_unsat_core_responseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSMTLIBv2Visitor) VisitGet_value_response(ctx *Get_value_responseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSMTLIBv2Visitor) VisitSpecific_success_response(ctx *Specific_success_responseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSMTLIBv2Visitor) VisitGeneral_response(ctx *General_responseContext) interface{} {
	return v.VisitChildren(ctx)
}
