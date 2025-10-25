parser grammar FaultParser;

options {
    tokenVocab=FaultLexer;
}

/*
    State charts of the whole system
*/

sysSpec
    : sysClause importDecl* globalDecl* componentDecl* (assertion | assumption | stringDecl)* startBlock? forStmt?
    ;

sysClause
    : 'system' IDENT eos
    ;

globalDecl
    : 'global' IDENT '=' operand eos (swap eos)*
    ;

swap
    : paramCall '=' (functionLit | numeric | string_ | bool_ | operandName | prefix | solvable)
    ;

componentDecl
    : 'component' IDENT '=' 'states' '{' (comProperties ',')* '}' eos
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
    : specClause declaration* forStmt?
    ;

specClause
    : 'spec' IDENT eos
    ;

importDecl
    : 'import' (importSpec | '(' importSpec* ')') eos
    ;

importSpec
    : ('.' | IDENT)? importPath ','?
    ;

importPath
    : string_
    ;

declaration
    : importDecl
    | constDecl
    | structDecl
    | assertion
    | assumption
    | stringDecl
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
    : 'const' ((constSpec eos) | '(' constSpec* ')' eos)
    ;

constSpec
    : identList ('=' constants)?
    ;

stringDecl
    : IDENT '=' string_ eos
    | IDENT '=' compoundString eos
    ;

compoundString
    : operandName
    | '!' operandName
    | '(' compoundString ')'
    | compoundString '&&' compoundString
    | compoundString '||' compoundString
    ;

identList
    : operandName (',' operandName)*
    ;

constants
    : numeric
    | string_
    | bool_
    | solvable
    | nil
    ;

nil
: NIL
;


expressionList
    : expression (',' expression)*
    ;

structDecl
    : 'def' IDENT '=' structType eos
    ;

structType
    : 'flow' '{' (sfProperties ',')* '}'    #Flow
    | 'stock' '{' (sfProperties ',')* '}'   #Stock
    ;

sfProperties
    : IDENT ':' functionLit #PropFunc
    | structProperties      #sfMisc
    ;

comProperties
    : IDENT ':' stateLit #StateFunc
    | structProperties   #compMisc
    ;

structProperties
    : IDENT ':' numeric #PropInt 
    | IDENT ':' string_ #PropString
    | IDENT ':' bool_ #PropBool
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
    | emptyStmt
    ;

incDecStmt
    : expression (PLUS_PLUS | MINUS_MINUS)
    ;

boolExpression
    : boolCompound
    ;

boolCompound
    : boolCompound '||' boolAnd
    | boolAnd
    ;

boolAnd
    : boolAnd '&&' boolPrimary
    | boolPrimary
    ;

boolPrimary
    : stateChange
    | '(' boolCompound ')'
    ;

stateChange
    : ('advance' | 'leave') '(' paramCall ')' #builtins
    | ('stay' | 'leave') '(' ')'              #builtins
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
    ;

invariant
    : operand '=' expression  # defInvariant
    | expression                            # invar
    | 'when' expression 'then' expression   # stageInvariant
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

ifStmtRun
    : 'if' (simpleStmt ';')? expression runBlock ('else' (ifStmtRun | runBlock))?
    ;

ifStmtState
    : 'if' (simpleStmt ';')? expression stateBlock ('else' (ifStmtState | stateBlock))?
    ;

forStmt
    : 'for' rounds ('init' initBlock)? 'run' runBlock eos?
    ;

rounds
    : integer
    ;

paramCall
    : (IDENT|THIS) '.' IDENT ('.' IDENT)*
    ;

stateBlock
    : '{' stateStep* '}'
    ;

stateStep
    : paramCall ('|' paramCall)? eos              #stateStepExpr
    | 'choose'? boolExpression eos                   #builtinInfix
    | stateChange eos                                #stateChain
    | ifStmtState                                 #stateExpr
    ;

runBlock
    : '{' runStep* '}'
    ;

initBlock
    : '{' initStep* '}'
    ;

initStep
    : IDENT '=' 'new' (paramCall | IDENT) eos (swap eos)*  #runInit                               
    ;

runStep
    : paramCall ('|' paramCall)* eos              #runStepExpr
    | simpleStmt eos                              #runExpr
    | ifStmtRun                                     #runExpr
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

postfix
    : operand
    | solvable
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
    : nil
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
    ('+' | '-' | '!' | '^' | '*' | '&' ) postfix
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

stateLit
    : 'func' stateBlock
    ;

eos
    : ';'
    ;

