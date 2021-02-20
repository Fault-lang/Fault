parser grammar FaultParser;

options {
    tokenVocab=FaultLexer;
}

spec
    : specClause importDecl? declaration* forStmt? eos
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
    : 'def' identList '=' structType eos
    ;

structType
    : 'flow' '{' (structProperties ',')* '}'    #Flow
    | 'stock' '{' (structProperties ',')* '}'   #Stock
    ;

structProperties
    : IDENT ':' numeric #PropInt 
    | IDENT ':' string_ #PropString
    | IDENT ':' functionLit #PropFunc
    | IDENT ':' instance #PropVar
    ;

instance
    : 'new' operandName
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
    : 'assert' expression
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
    : 'for' integer 'run' block eos
    ;

expression
    : operand                                                            #Expr
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
    | instance
    | operandName
    | accessHistory
    | '(' expression ')'
    ;

operandName
    : IDENT
    | IDENT ('.' IDENT)?
    | 'new' IDENT
    | THIS
    | CLOCK
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

