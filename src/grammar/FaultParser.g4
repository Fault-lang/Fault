parser grammar FaultParser;

options {
    tokenVocab=FaultLexer;
}

spec
    : specClause importDecl* declaration* forStmt? eos
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
    | IDENT ':' functionLit #PropFunc
    | IDENT ':' operandName #PropVar
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
    | emptyStmt
    ;

incDecStmt
    : expression (PLUS_PLUS | MINUS_MINUS)
    ;

accessHistory
    : operandName ('[' expression ']')+
    ;

assertion
    : 'assert' expression eos
    ;

assumption
    : 'assume' expression eos
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
    : IDENT '.' IDENT ('.' IDENT)*
    ;

runBlock
    : '{' runStep* '}'
    ;

runStep
    : paramCall ('|' paramCall)* eos              #runStepExpr
    | IDENT '=' 'new' IDENT ('.' IDENT)? eos      #runInit
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
    ;

expression
    : operand                                                            #Expr
    | faultType '(' operand (',' operand)* ')'                           #Typed
    | ('+' | '-' | '!' | '^' | '*' | '&') expression                     #Prefix
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

