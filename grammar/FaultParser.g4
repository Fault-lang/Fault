parser grammar FaultParser;

options {
    tokenVocab=FaultLexer;
}

/*
    State charts of the whole system
*/

sysSpec
    : sysClause importDecl* globalDecl* componentDecl* startBlock* (assertion | assumption)? forStmt? eos
    ;

sysClause
    : 'system' IDENT eos
    ;

globalDecl
    : 'global' IDENT '=' operand eos
    ;

componentDecl
    : 'component' IDENT '=' 'states' '{' (structProperties ',')* '}' eos
    ;

startBlock
    : 'start' '{' (startPair ',')* '}' eos
    ;

startPair
    : IDENT ':' IDENT
    ;
/*
    Individual specs of state changes
*/

spec
    : specClause declaration* forStmt? eos
    ;

specClause
    : 'spec' IDENT eos
    ;

importDecl
    : 'import' (importSpec | '(' importSpec* ')') eos
    ;

importSpec
    : ('.' | IDENT)? importPath
    ;

importPath
    : string_
    ;

declaration
    : constDecl
    | structDecl
    | assertion
    | assumption
    ;

comparison
    : EQUALS
    | NOT_EQUALS
    | LESS 
    | LESS_OR_EQUALS
    | GREATER
    | GREATER_OR_EQUALS
    ;

constDecl
    : 'const' ((constSpec eos) | '(' (constSpec eos)* ')')
    ;

constSpec
    : identList ('=' expressionList)?
    ;

identList
    : operandName (',' operandName)*
    ;

expressionList
    : expression (',' expression)*
    ;

structDecl
    : 'def' IDENT '=' structType eos
    ;

structType
    : 'flow' '{' (structProperties ',')* '}'    #Flow
    | 'stock' '{' (structProperties ',')* '}'   #Stock
    ;

structProperties
    : IDENT ':' numeric #PropInt 
    | IDENT ':' string_ #PropString
    | IDENT ':' bool_ #PropBool
    | IDENT ':' functionLit #PropFunc
    | IDENT ':' operandName #PropVar
    | IDENT ':' prefix #PropVar
    | IDENT ':' solvable #PropSolvable
    | IDENT              #PropSolvable
    ;

initDecl
    : 'init' operand eos
    ;

block
    : '{' statementList? '}'
    ;

statementList
    : statement+
    ;

statement
    : constDecl
    | initDecl
    | simpleStmt eos
    | block
    | ifStmt
    ;

simpleStmt
    : expression
    | incDecStmt
    | assignment
    | builtins
    | emptyStmt
    ;

incDecStmt
    : expression (PLUS_PLUS | MINUS_MINUS)
    ;

builtins
    : 'advance' '(' paramCall ')'
    | 'stay' '(' ')'
    ;

accessHistory
    : operandName ('[' expression ']')+
    ;

assertion
    : 'assert' invariant temporal? eos
    ;

assumption
    : 'assume' invariant temporal? eos
    ;

temporal
    : ('eventually' | 'always' | 'eventually-always' )
    | ('nmt' | 'nft') integer
    | 'then' expression
    ;

invariant
    : expression
    ;

assignment
    : expressionList ('+' | '-' | '^' | '*' | '/' | '%' | '<<' | '>>' | '&' | '&^')? '=' expressionList #MiscAssign
    | expressionList ('->' | '<-') expressionList #FaultAssign
    ;

emptyStmt
    : ';'
    ;

ifStmt
    : 'if' (simpleStmt ';')? expression block ('else' (ifStmt | block))?
    ;

forStmt
    : 'for' rounds 'run' runBlock eos
    ;

rounds
    : integer
    ;

paramCall
    : (IDENT|THIS) '.' IDENT ('.' IDENT)*
    ;

runBlock
    : '{' runStep* '}'
    ;

runStep
    : paramCall ('|' paramCall)* eos              #runStepExpr
    | IDENT '=' 'new' (paramCall | IDENT) eos      #runInit
    | simpleStmt eos                              #runExpr
    | ifStmt                                      #runExpr
    ;

faultType
    : TY_STRING
    | TY_BOOL
    | TY_INT
    | TY_FLOAT
    | TY_NATURAL
    | TY_UNCERTAIN
    | TY_UNKNOWN
    ;

solvable
    : faultType '(' operand? (',' operand)* ')' 
    ;

expression
    : operand                                                            #Expr
    | solvable                                                           #Typed
    | prefix                                                             #ExprPrefix
    | expression '**' expression                                         #lrExpr
    | expression ('*' | '/' | '%' | '<<' | '>>' | '&' | '&^') expression #lrExpr
    | expression ('+' | '-' | '^') expression                            #lrExpr
    | expression ('==' | '!=' | '<' | '<=' | '>' | '>=') expression      #lrExpr
    | expression '&&' expression                                         #lrExpr
    | expression '||' expression                                         #lrExpr
    ;

operand
    : NIL
    | numeric
    | string_
    | bool_
    | operandName
    | accessHistory
    | '(' expression ')'
    ;

operandName
    : IDENT                     #OpName
    | paramCall                 #OpParam
    | THIS                      #OpThis
    | CLOCK                     #OpClock
    | 'new' IDENT ('.' IDENT)?  #OpInstance
    ;

prefix
    :
    | ('+' | '-' | '!' | '^' | '*' | '&' ) expression
    ;

numeric
    : integer
    | negative
    | float_
    ;

integer
    : DECIMAL_LIT
    | OCTAL_LIT
    | HEX_LIT
    ;

negative
    : '-' integer
    | '-' float_
    ;

float_
    : FLOAT_LIT
    ;

string_
    : RAW_STRING_LIT
    | INTERPRETED_STRING_LIT
    ;

bool_
    : TRUE
    | FALSE
    ;

functionLit
    : 'func' block
    ;

eos
    : ';'
    | EOF
    ;

