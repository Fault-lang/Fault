// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package parser // SMTLIBv2

import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

// SMTLIBv2Listener is a complete listener for a parse tree produced by SMTLIBv2Parser.
type SMTLIBv2Listener interface {
	antlr.ParseTreeListener

	// EnterStart is called when entering the start production.
	EnterStart(c *StartContext)

	// EnterGeneralReservedWord is called when entering the generalReservedWord production.
	EnterGeneralReservedWord(c *GeneralReservedWordContext)

	// EnterSimpleSymbol is called when entering the simpleSymbol production.
	EnterSimpleSymbol(c *SimpleSymbolContext)

	// EnterQuotedSymbol is called when entering the quotedSymbol production.
	EnterQuotedSymbol(c *QuotedSymbolContext)

	// EnterPredefSymbol is called when entering the predefSymbol production.
	EnterPredefSymbol(c *PredefSymbolContext)

	// EnterPredefKeyword is called when entering the predefKeyword production.
	EnterPredefKeyword(c *PredefKeywordContext)

	// EnterSymbol is called when entering the symbol production.
	EnterSymbol(c *SymbolContext)

	// EnterNumeral is called when entering the numeral production.
	EnterNumeral(c *NumeralContext)

	// EnterDecimal is called when entering the decimal production.
	EnterDecimal(c *DecimalContext)

	// EnterHexadecimal is called when entering the hexadecimal production.
	EnterHexadecimal(c *HexadecimalContext)

	// EnterBinary is called when entering the binary production.
	EnterBinary(c *BinaryContext)

	// EnterString_ is called when entering the string_ production.
	EnterString_(c *String_Context)

	// EnterKeyword is called when entering the keyword production.
	EnterKeyword(c *KeywordContext)

	// EnterSpec_constant is called when entering the spec_constant production.
	EnterSpec_constant(c *Spec_constantContext)

	// EnterS_expr is called when entering the s_expr production.
	EnterS_expr(c *S_exprContext)

	// EnterIndex is called when entering the index production.
	EnterIndex(c *IndexContext)

	// EnterIdentifier is called when entering the identifier production.
	EnterIdentifier(c *IdentifierContext)

	// EnterAttribute_value is called when entering the attribute_value production.
	EnterAttribute_value(c *Attribute_valueContext)

	// EnterAttribute is called when entering the attribute production.
	EnterAttribute(c *AttributeContext)

	// EnterSort is called when entering the sort production.
	EnterSort(c *SortContext)

	// EnterQual_identifer is called when entering the qual_identifer production.
	EnterQual_identifer(c *Qual_identiferContext)

	// EnterVar_binding is called when entering the var_binding production.
	EnterVar_binding(c *Var_bindingContext)

	// EnterSorted_var is called when entering the sorted_var production.
	EnterSorted_var(c *Sorted_varContext)

	// EnterPattern is called when entering the pattern production.
	EnterPattern(c *PatternContext)

	// EnterMatch_case is called when entering the match_case production.
	EnterMatch_case(c *Match_caseContext)

	// EnterVariable is called when entering the variable production.
	EnterVariable(c *VariableContext)

	// EnterTerm is called when entering the term production.
	EnterTerm(c *TermContext)

	// EnterSort_symbol_decl is called when entering the sort_symbol_decl production.
	EnterSort_symbol_decl(c *Sort_symbol_declContext)

	// EnterMeta_spec_constant is called when entering the meta_spec_constant production.
	EnterMeta_spec_constant(c *Meta_spec_constantContext)

	// EnterFun_symbol_decl is called when entering the fun_symbol_decl production.
	EnterFun_symbol_decl(c *Fun_symbol_declContext)

	// EnterPar_fun_symbol_decl is called when entering the par_fun_symbol_decl production.
	EnterPar_fun_symbol_decl(c *Par_fun_symbol_declContext)

	// EnterTheory_attribute is called when entering the theory_attribute production.
	EnterTheory_attribute(c *Theory_attributeContext)

	// EnterTheory_decl is called when entering the theory_decl production.
	EnterTheory_decl(c *Theory_declContext)

	// EnterLogic_attribue is called when entering the logic_attribue production.
	EnterLogic_attribue(c *Logic_attribueContext)

	// EnterLogic is called when entering the logic production.
	EnterLogic(c *LogicContext)

	// EnterSort_dec is called when entering the sort_dec production.
	EnterSort_dec(c *Sort_decContext)

	// EnterSelector_dec is called when entering the selector_dec production.
	EnterSelector_dec(c *Selector_decContext)

	// EnterConstructor_dec is called when entering the constructor_dec production.
	EnterConstructor_dec(c *Constructor_decContext)

	// EnterDatatype_dec is called when entering the datatype_dec production.
	EnterDatatype_dec(c *Datatype_decContext)

	// EnterFunction_dec is called when entering the function_dec production.
	EnterFunction_dec(c *Function_decContext)

	// EnterFunction_def is called when entering the function_def production.
	EnterFunction_def(c *Function_defContext)

	// EnterProp_literal is called when entering the prop_literal production.
	EnterProp_literal(c *Prop_literalContext)

	// EnterScript is called when entering the script production.
	EnterScript(c *ScriptContext)

	// EnterCmd_assert is called when entering the cmd_assert production.
	EnterCmd_assert(c *Cmd_assertContext)

	// EnterCmd_checkSat is called when entering the cmd_checkSat production.
	EnterCmd_checkSat(c *Cmd_checkSatContext)

	// EnterCmd_checkSatAssuming is called when entering the cmd_checkSatAssuming production.
	EnterCmd_checkSatAssuming(c *Cmd_checkSatAssumingContext)

	// EnterCmd_declareConst is called when entering the cmd_declareConst production.
	EnterCmd_declareConst(c *Cmd_declareConstContext)

	// EnterCmd_declareDatatype is called when entering the cmd_declareDatatype production.
	EnterCmd_declareDatatype(c *Cmd_declareDatatypeContext)

	// EnterCmd_declareDatatypes is called when entering the cmd_declareDatatypes production.
	EnterCmd_declareDatatypes(c *Cmd_declareDatatypesContext)

	// EnterCmd_declareFun is called when entering the cmd_declareFun production.
	EnterCmd_declareFun(c *Cmd_declareFunContext)

	// EnterCmd_declareSort is called when entering the cmd_declareSort production.
	EnterCmd_declareSort(c *Cmd_declareSortContext)

	// EnterCmd_defineFun is called when entering the cmd_defineFun production.
	EnterCmd_defineFun(c *Cmd_defineFunContext)

	// EnterCmd_defineFunRec is called when entering the cmd_defineFunRec production.
	EnterCmd_defineFunRec(c *Cmd_defineFunRecContext)

	// EnterCmd_defineFunsRec is called when entering the cmd_defineFunsRec production.
	EnterCmd_defineFunsRec(c *Cmd_defineFunsRecContext)

	// EnterCmd_defineSort is called when entering the cmd_defineSort production.
	EnterCmd_defineSort(c *Cmd_defineSortContext)

	// EnterCmd_echo is called when entering the cmd_echo production.
	EnterCmd_echo(c *Cmd_echoContext)

	// EnterCmd_exit is called when entering the cmd_exit production.
	EnterCmd_exit(c *Cmd_exitContext)

	// EnterCmd_getAssertions is called when entering the cmd_getAssertions production.
	EnterCmd_getAssertions(c *Cmd_getAssertionsContext)

	// EnterCmd_getAssignment is called when entering the cmd_getAssignment production.
	EnterCmd_getAssignment(c *Cmd_getAssignmentContext)

	// EnterCmd_getInfo is called when entering the cmd_getInfo production.
	EnterCmd_getInfo(c *Cmd_getInfoContext)

	// EnterCmd_getModel is called when entering the cmd_getModel production.
	EnterCmd_getModel(c *Cmd_getModelContext)

	// EnterCmd_getOption is called when entering the cmd_getOption production.
	EnterCmd_getOption(c *Cmd_getOptionContext)

	// EnterCmd_getProof is called when entering the cmd_getProof production.
	EnterCmd_getProof(c *Cmd_getProofContext)

	// EnterCmd_getUnsatAssumptions is called when entering the cmd_getUnsatAssumptions production.
	EnterCmd_getUnsatAssumptions(c *Cmd_getUnsatAssumptionsContext)

	// EnterCmd_getUnsatCore is called when entering the cmd_getUnsatCore production.
	EnterCmd_getUnsatCore(c *Cmd_getUnsatCoreContext)

	// EnterCmd_getValue is called when entering the cmd_getValue production.
	EnterCmd_getValue(c *Cmd_getValueContext)

	// EnterCmd_pop is called when entering the cmd_pop production.
	EnterCmd_pop(c *Cmd_popContext)

	// EnterCmd_push is called when entering the cmd_push production.
	EnterCmd_push(c *Cmd_pushContext)

	// EnterCmd_reset is called when entering the cmd_reset production.
	EnterCmd_reset(c *Cmd_resetContext)

	// EnterCmd_resetAssertions is called when entering the cmd_resetAssertions production.
	EnterCmd_resetAssertions(c *Cmd_resetAssertionsContext)

	// EnterCmd_setInfo is called when entering the cmd_setInfo production.
	EnterCmd_setInfo(c *Cmd_setInfoContext)

	// EnterCmd_setLogic is called when entering the cmd_setLogic production.
	EnterCmd_setLogic(c *Cmd_setLogicContext)

	// EnterCmd_setOption is called when entering the cmd_setOption production.
	EnterCmd_setOption(c *Cmd_setOptionContext)

	// EnterCommand is called when entering the command production.
	EnterCommand(c *CommandContext)

	// EnterB_value is called when entering the b_value production.
	EnterB_value(c *B_valueContext)

	// EnterOption is called when entering the option production.
	EnterOption(c *OptionContext)

	// EnterInfo_flag is called when entering the info_flag production.
	EnterInfo_flag(c *Info_flagContext)

	// EnterError_behaviour is called when entering the error_behaviour production.
	EnterError_behaviour(c *Error_behaviourContext)

	// EnterReason_unknown is called when entering the reason_unknown production.
	EnterReason_unknown(c *Reason_unknownContext)

	// EnterModel_response is called when entering the model_response production.
	EnterModel_response(c *Model_responseContext)

	// EnterInfo_response is called when entering the info_response production.
	EnterInfo_response(c *Info_responseContext)

	// EnterValuation_pair is called when entering the valuation_pair production.
	EnterValuation_pair(c *Valuation_pairContext)

	// EnterT_valuation_pair is called when entering the t_valuation_pair production.
	EnterT_valuation_pair(c *T_valuation_pairContext)

	// EnterCheck_sat_response is called when entering the check_sat_response production.
	EnterCheck_sat_response(c *Check_sat_responseContext)

	// EnterEcho_response is called when entering the echo_response production.
	EnterEcho_response(c *Echo_responseContext)

	// EnterGet_assertions_response is called when entering the get_assertions_response production.
	EnterGet_assertions_response(c *Get_assertions_responseContext)

	// EnterGet_assignment_response is called when entering the get_assignment_response production.
	EnterGet_assignment_response(c *Get_assignment_responseContext)

	// EnterGet_info_response is called when entering the get_info_response production.
	EnterGet_info_response(c *Get_info_responseContext)

	// EnterGet_model_response is called when entering the get_model_response production.
	EnterGet_model_response(c *Get_model_responseContext)

	// EnterGet_option_response is called when entering the get_option_response production.
	EnterGet_option_response(c *Get_option_responseContext)

	// EnterGet_proof_response is called when entering the get_proof_response production.
	EnterGet_proof_response(c *Get_proof_responseContext)

	// EnterGet_unsat_assump_response is called when entering the get_unsat_assump_response production.
	EnterGet_unsat_assump_response(c *Get_unsat_assump_responseContext)

	// EnterGet_unsat_core_response is called when entering the get_unsat_core_response production.
	EnterGet_unsat_core_response(c *Get_unsat_core_responseContext)

	// EnterGet_value_response is called when entering the get_value_response production.
	EnterGet_value_response(c *Get_value_responseContext)

	// EnterSpecific_success_response is called when entering the specific_success_response production.
	EnterSpecific_success_response(c *Specific_success_responseContext)

	// EnterGeneral_response is called when entering the general_response production.
	EnterGeneral_response(c *General_responseContext)

	// ExitStart is called when exiting the start production.
	ExitStart(c *StartContext)

	// ExitGeneralReservedWord is called when exiting the generalReservedWord production.
	ExitGeneralReservedWord(c *GeneralReservedWordContext)

	// ExitSimpleSymbol is called when exiting the simpleSymbol production.
	ExitSimpleSymbol(c *SimpleSymbolContext)

	// ExitQuotedSymbol is called when exiting the quotedSymbol production.
	ExitQuotedSymbol(c *QuotedSymbolContext)

	// ExitPredefSymbol is called when exiting the predefSymbol production.
	ExitPredefSymbol(c *PredefSymbolContext)

	// ExitPredefKeyword is called when exiting the predefKeyword production.
	ExitPredefKeyword(c *PredefKeywordContext)

	// ExitSymbol is called when exiting the symbol production.
	ExitSymbol(c *SymbolContext)

	// ExitNumeral is called when exiting the numeral production.
	ExitNumeral(c *NumeralContext)

	// ExitDecimal is called when exiting the decimal production.
	ExitDecimal(c *DecimalContext)

	// ExitHexadecimal is called when exiting the hexadecimal production.
	ExitHexadecimal(c *HexadecimalContext)

	// ExitBinary is called when exiting the binary production.
	ExitBinary(c *BinaryContext)

	// ExitString_ is called when exiting the string_ production.
	ExitString_(c *String_Context)

	// ExitKeyword is called when exiting the keyword production.
	ExitKeyword(c *KeywordContext)

	// ExitSpec_constant is called when exiting the spec_constant production.
	ExitSpec_constant(c *Spec_constantContext)

	// ExitS_expr is called when exiting the s_expr production.
	ExitS_expr(c *S_exprContext)

	// ExitIndex is called when exiting the index production.
	ExitIndex(c *IndexContext)

	// ExitIdentifier is called when exiting the identifier production.
	ExitIdentifier(c *IdentifierContext)

	// ExitAttribute_value is called when exiting the attribute_value production.
	ExitAttribute_value(c *Attribute_valueContext)

	// ExitAttribute is called when exiting the attribute production.
	ExitAttribute(c *AttributeContext)

	// ExitSort is called when exiting the sort production.
	ExitSort(c *SortContext)

	// ExitQual_identifer is called when exiting the qual_identifer production.
	ExitQual_identifer(c *Qual_identiferContext)

	// ExitVar_binding is called when exiting the var_binding production.
	ExitVar_binding(c *Var_bindingContext)

	// ExitSorted_var is called when exiting the sorted_var production.
	ExitSorted_var(c *Sorted_varContext)

	// ExitPattern is called when exiting the pattern production.
	ExitPattern(c *PatternContext)

	// ExitMatch_case is called when exiting the match_case production.
	ExitMatch_case(c *Match_caseContext)

	// ExitVariable is called when exiting the variable production.
	ExitVariable(c *VariableContext)

	// ExitTerm is called when exiting the term production.
	ExitTerm(c *TermContext)

	// ExitSort_symbol_decl is called when exiting the sort_symbol_decl production.
	ExitSort_symbol_decl(c *Sort_symbol_declContext)

	// ExitMeta_spec_constant is called when exiting the meta_spec_constant production.
	ExitMeta_spec_constant(c *Meta_spec_constantContext)

	// ExitFun_symbol_decl is called when exiting the fun_symbol_decl production.
	ExitFun_symbol_decl(c *Fun_symbol_declContext)

	// ExitPar_fun_symbol_decl is called when exiting the par_fun_symbol_decl production.
	ExitPar_fun_symbol_decl(c *Par_fun_symbol_declContext)

	// ExitTheory_attribute is called when exiting the theory_attribute production.
	ExitTheory_attribute(c *Theory_attributeContext)

	// ExitTheory_decl is called when exiting the theory_decl production.
	ExitTheory_decl(c *Theory_declContext)

	// ExitLogic_attribue is called when exiting the logic_attribue production.
	ExitLogic_attribue(c *Logic_attribueContext)

	// ExitLogic is called when exiting the logic production.
	ExitLogic(c *LogicContext)

	// ExitSort_dec is called when exiting the sort_dec production.
	ExitSort_dec(c *Sort_decContext)

	// ExitSelector_dec is called when exiting the selector_dec production.
	ExitSelector_dec(c *Selector_decContext)

	// ExitConstructor_dec is called when exiting the constructor_dec production.
	ExitConstructor_dec(c *Constructor_decContext)

	// ExitDatatype_dec is called when exiting the datatype_dec production.
	ExitDatatype_dec(c *Datatype_decContext)

	// ExitFunction_dec is called when exiting the function_dec production.
	ExitFunction_dec(c *Function_decContext)

	// ExitFunction_def is called when exiting the function_def production.
	ExitFunction_def(c *Function_defContext)

	// ExitProp_literal is called when exiting the prop_literal production.
	ExitProp_literal(c *Prop_literalContext)

	// ExitScript is called when exiting the script production.
	ExitScript(c *ScriptContext)

	// ExitCmd_assert is called when exiting the cmd_assert production.
	ExitCmd_assert(c *Cmd_assertContext)

	// ExitCmd_checkSat is called when exiting the cmd_checkSat production.
	ExitCmd_checkSat(c *Cmd_checkSatContext)

	// ExitCmd_checkSatAssuming is called when exiting the cmd_checkSatAssuming production.
	ExitCmd_checkSatAssuming(c *Cmd_checkSatAssumingContext)

	// ExitCmd_declareConst is called when exiting the cmd_declareConst production.
	ExitCmd_declareConst(c *Cmd_declareConstContext)

	// ExitCmd_declareDatatype is called when exiting the cmd_declareDatatype production.
	ExitCmd_declareDatatype(c *Cmd_declareDatatypeContext)

	// ExitCmd_declareDatatypes is called when exiting the cmd_declareDatatypes production.
	ExitCmd_declareDatatypes(c *Cmd_declareDatatypesContext)

	// ExitCmd_declareFun is called when exiting the cmd_declareFun production.
	ExitCmd_declareFun(c *Cmd_declareFunContext)

	// ExitCmd_declareSort is called when exiting the cmd_declareSort production.
	ExitCmd_declareSort(c *Cmd_declareSortContext)

	// ExitCmd_defineFun is called when exiting the cmd_defineFun production.
	ExitCmd_defineFun(c *Cmd_defineFunContext)

	// ExitCmd_defineFunRec is called when exiting the cmd_defineFunRec production.
	ExitCmd_defineFunRec(c *Cmd_defineFunRecContext)

	// ExitCmd_defineFunsRec is called when exiting the cmd_defineFunsRec production.
	ExitCmd_defineFunsRec(c *Cmd_defineFunsRecContext)

	// ExitCmd_defineSort is called when exiting the cmd_defineSort production.
	ExitCmd_defineSort(c *Cmd_defineSortContext)

	// ExitCmd_echo is called when exiting the cmd_echo production.
	ExitCmd_echo(c *Cmd_echoContext)

	// ExitCmd_exit is called when exiting the cmd_exit production.
	ExitCmd_exit(c *Cmd_exitContext)

	// ExitCmd_getAssertions is called when exiting the cmd_getAssertions production.
	ExitCmd_getAssertions(c *Cmd_getAssertionsContext)

	// ExitCmd_getAssignment is called when exiting the cmd_getAssignment production.
	ExitCmd_getAssignment(c *Cmd_getAssignmentContext)

	// ExitCmd_getInfo is called when exiting the cmd_getInfo production.
	ExitCmd_getInfo(c *Cmd_getInfoContext)

	// ExitCmd_getModel is called when exiting the cmd_getModel production.
	ExitCmd_getModel(c *Cmd_getModelContext)

	// ExitCmd_getOption is called when exiting the cmd_getOption production.
	ExitCmd_getOption(c *Cmd_getOptionContext)

	// ExitCmd_getProof is called when exiting the cmd_getProof production.
	ExitCmd_getProof(c *Cmd_getProofContext)

	// ExitCmd_getUnsatAssumptions is called when exiting the cmd_getUnsatAssumptions production.
	ExitCmd_getUnsatAssumptions(c *Cmd_getUnsatAssumptionsContext)

	// ExitCmd_getUnsatCore is called when exiting the cmd_getUnsatCore production.
	ExitCmd_getUnsatCore(c *Cmd_getUnsatCoreContext)

	// ExitCmd_getValue is called when exiting the cmd_getValue production.
	ExitCmd_getValue(c *Cmd_getValueContext)

	// ExitCmd_pop is called when exiting the cmd_pop production.
	ExitCmd_pop(c *Cmd_popContext)

	// ExitCmd_push is called when exiting the cmd_push production.
	ExitCmd_push(c *Cmd_pushContext)

	// ExitCmd_reset is called when exiting the cmd_reset production.
	ExitCmd_reset(c *Cmd_resetContext)

	// ExitCmd_resetAssertions is called when exiting the cmd_resetAssertions production.
	ExitCmd_resetAssertions(c *Cmd_resetAssertionsContext)

	// ExitCmd_setInfo is called when exiting the cmd_setInfo production.
	ExitCmd_setInfo(c *Cmd_setInfoContext)

	// ExitCmd_setLogic is called when exiting the cmd_setLogic production.
	ExitCmd_setLogic(c *Cmd_setLogicContext)

	// ExitCmd_setOption is called when exiting the cmd_setOption production.
	ExitCmd_setOption(c *Cmd_setOptionContext)

	// ExitCommand is called when exiting the command production.
	ExitCommand(c *CommandContext)

	// ExitB_value is called when exiting the b_value production.
	ExitB_value(c *B_valueContext)

	// ExitOption is called when exiting the option production.
	ExitOption(c *OptionContext)

	// ExitInfo_flag is called when exiting the info_flag production.
	ExitInfo_flag(c *Info_flagContext)

	// ExitError_behaviour is called when exiting the error_behaviour production.
	ExitError_behaviour(c *Error_behaviourContext)

	// ExitReason_unknown is called when exiting the reason_unknown production.
	ExitReason_unknown(c *Reason_unknownContext)

	// ExitModel_response is called when exiting the model_response production.
	ExitModel_response(c *Model_responseContext)

	// ExitInfo_response is called when exiting the info_response production.
	ExitInfo_response(c *Info_responseContext)

	// ExitValuation_pair is called when exiting the valuation_pair production.
	ExitValuation_pair(c *Valuation_pairContext)

	// ExitT_valuation_pair is called when exiting the t_valuation_pair production.
	ExitT_valuation_pair(c *T_valuation_pairContext)

	// ExitCheck_sat_response is called when exiting the check_sat_response production.
	ExitCheck_sat_response(c *Check_sat_responseContext)

	// ExitEcho_response is called when exiting the echo_response production.
	ExitEcho_response(c *Echo_responseContext)

	// ExitGet_assertions_response is called when exiting the get_assertions_response production.
	ExitGet_assertions_response(c *Get_assertions_responseContext)

	// ExitGet_assignment_response is called when exiting the get_assignment_response production.
	ExitGet_assignment_response(c *Get_assignment_responseContext)

	// ExitGet_info_response is called when exiting the get_info_response production.
	ExitGet_info_response(c *Get_info_responseContext)

	// ExitGet_model_response is called when exiting the get_model_response production.
	ExitGet_model_response(c *Get_model_responseContext)

	// ExitGet_option_response is called when exiting the get_option_response production.
	ExitGet_option_response(c *Get_option_responseContext)

	// ExitGet_proof_response is called when exiting the get_proof_response production.
	ExitGet_proof_response(c *Get_proof_responseContext)

	// ExitGet_unsat_assump_response is called when exiting the get_unsat_assump_response production.
	ExitGet_unsat_assump_response(c *Get_unsat_assump_responseContext)

	// ExitGet_unsat_core_response is called when exiting the get_unsat_core_response production.
	ExitGet_unsat_core_response(c *Get_unsat_core_responseContext)

	// ExitGet_value_response is called when exiting the get_value_response production.
	ExitGet_value_response(c *Get_value_responseContext)

	// ExitSpecific_success_response is called when exiting the specific_success_response production.
	ExitSpecific_success_response(c *Specific_success_responseContext)

	// ExitGeneral_response is called when exiting the general_response production.
	ExitGeneral_response(c *General_responseContext)
}
