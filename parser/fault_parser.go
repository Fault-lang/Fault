// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package parser // FaultParser

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type FaultParser struct {
	*antlr.BaseParser
}

var faultparserParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func faultparserParserInit() {
	staticData := &faultparserParserStaticData
	staticData.literalNames = []string{
		"", "'all'", "'assert'", "'assume'", "'clock'", "'const'", "'def'",
		"'else'", "'flow'", "'for'", "'func'", "'if'", "'import'", "'init'",
		"'new'", "'return'", "'run'", "'spec'", "'stock'", "'then'", "'when'",
		"'this'", "'eventually'", "'eventually-always'", "'always'", "'nmt'",
		"'nft'", "'nil'", "'true'", "'false'", "'advance'", "'component'", "'global'",
		"'system'", "'start'", "'states'", "'stay'", "'string'", "'bool'", "'int'",
		"'float'", "'natural'", "'uncertain'", "'unknown'", "", "'='", "'->'",
		"'<-'", "':'", "','", "'.'", "'('", "')'", "'{'", "'}'", "'['", "']'",
		"';'", "'++'", "'--'", "'&'", "'&&'", "'!'", "'=='", "'!='", "'<'",
		"'<='", "'>'", "'>='", "'||'", "'|'", "'+'", "'-'", "'^'", "'**'", "'*'",
		"'/'", "'%'", "'<<'", "'>>'", "'&^'",
	}
	staticData.symbolicNames = []string{
		"", "ALL", "ASSERT", "ASSUME", "CLOCK", "CONST", "DEF", "ELSE", "FLOW",
		"FOR", "FUNC", "IF", "IMPORT", "INIT", "NEW", "RETURN", "RUN", "SPEC",
		"STOCK", "THEN", "WHEN", "THIS", "EVENTUALLY", "EVENTUALLYALWAYS", "ALWAYS",
		"NMT", "NFT", "NIL", "TRUE", "FALSE", "ADVANCE", "COMPONENT", "GLOBAL",
		"SYSTEM", "START", "STATE", "STAY", "TY_STRING", "TY_BOOL", "TY_INT",
		"TY_FLOAT", "TY_NATURAL", "TY_UNCERTAIN", "TY_UNKNOWN", "IDENT", "ASSIGN",
		"ASSIGN_FLOW1", "ASSIGN_FLOW2", "COLON", "COMMA", "DOT", "LPAREN", "RPAREN",
		"LCURLY", "RCURLY", "LBRACE", "RBRACE", "SEMI", "PLUS_PLUS", "MINUS_MINUS",
		"AMPERSAND", "AND", "BANG", "EQUALS", "NOT_EQUALS", "LESS", "LESS_OR_EQUALS",
		"GREATER", "GREATER_OR_EQUALS", "OR", "PIPE", "PLUS", "MINUS", "CARET",
		"EXPO", "MULTI", "DIV", "MOD", "LSHIFT", "RSHIFT", "BIT_CLEAR", "DECIMAL_LIT",
		"OCTAL_LIT", "HEX_LIT", "FLOAT_LIT", "RAW_STRING_LIT", "INTERPRETED_STRING_LIT",
		"WS", "COMMENT", "TERMINATOR", "LINE_COMMENT",
	}
	staticData.ruleNames = []string{
		"sysSpec", "sysClause", "globalDecl", "componentDecl", "startBlock",
		"startPair", "spec", "specClause", "importDecl", "importSpec", "importPath",
		"declaration", "comparison", "constDecl", "constSpec", "identList",
		"constants", "nil", "expressionList", "structDecl", "structType", "sfProperties",
		"comProperties", "structProperties", "initDecl", "block", "statementList",
		"statement", "simpleStmt", "incDecStmt", "stateChange", "accessHistory",
		"assertion", "assumption", "temporal", "invariant", "assignment", "emptyStmt",
		"ifStmt", "ifStmtRun", "ifStmtState", "forStmt", "rounds", "paramCall",
		"stateBlock", "stateStep", "runBlock", "runStep", "faultType", "solvable",
		"expression", "operand", "operandName", "prefix", "numeric", "integer",
		"negative", "float_", "string_", "bool_", "functionLit", "stateLit",
		"eos",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 90, 684, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36, 7, 36,
		2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7, 41, 2,
		42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 2, 45, 7, 45, 2, 46, 7, 46, 2, 47,
		7, 47, 2, 48, 7, 48, 2, 49, 7, 49, 2, 50, 7, 50, 2, 51, 7, 51, 2, 52, 7,
		52, 2, 53, 7, 53, 2, 54, 7, 54, 2, 55, 7, 55, 2, 56, 7, 56, 2, 57, 7, 57,
		2, 58, 7, 58, 2, 59, 7, 59, 2, 60, 7, 60, 2, 61, 7, 61, 2, 62, 7, 62, 1,
		0, 1, 0, 5, 0, 129, 8, 0, 10, 0, 12, 0, 132, 9, 0, 1, 0, 5, 0, 135, 8,
		0, 10, 0, 12, 0, 138, 9, 0, 1, 0, 5, 0, 141, 8, 0, 10, 0, 12, 0, 144, 9,
		0, 1, 0, 5, 0, 147, 8, 0, 10, 0, 12, 0, 150, 9, 0, 1, 0, 1, 0, 3, 0, 154,
		8, 0, 1, 0, 3, 0, 157, 8, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 5, 3,
		177, 8, 3, 10, 3, 12, 3, 180, 9, 3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4,
		1, 4, 1, 4, 5, 4, 190, 8, 4, 10, 4, 12, 4, 193, 9, 4, 1, 4, 1, 4, 1, 4,
		1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 5, 6, 204, 8, 6, 10, 6, 12, 6, 207,
		9, 6, 1, 6, 3, 6, 210, 8, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8,
		1, 8, 5, 8, 220, 8, 8, 10, 8, 12, 8, 223, 9, 8, 1, 8, 3, 8, 226, 8, 8,
		1, 8, 1, 8, 1, 9, 3, 9, 231, 8, 9, 1, 9, 1, 9, 3, 9, 235, 8, 9, 1, 10,
		1, 10, 1, 11, 1, 11, 1, 11, 1, 11, 3, 11, 243, 8, 11, 1, 12, 1, 12, 1,
		13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 5, 13, 253, 8, 13, 10, 13, 12, 13,
		256, 9, 13, 1, 13, 1, 13, 3, 13, 260, 8, 13, 1, 14, 1, 14, 1, 14, 3, 14,
		265, 8, 14, 1, 15, 1, 15, 1, 15, 5, 15, 270, 8, 15, 10, 15, 12, 15, 273,
		9, 15, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 3, 16, 280, 8, 16, 1, 17, 1,
		17, 1, 18, 1, 18, 1, 18, 5, 18, 287, 8, 18, 10, 18, 12, 18, 290, 9, 18,
		1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 20, 1, 20, 1, 20, 1, 20, 1,
		20, 5, 20, 303, 8, 20, 10, 20, 12, 20, 306, 9, 20, 1, 20, 1, 20, 1, 20,
		1, 20, 1, 20, 1, 20, 5, 20, 314, 8, 20, 10, 20, 12, 20, 317, 9, 20, 1,
		20, 3, 20, 320, 8, 20, 1, 21, 1, 21, 1, 21, 1, 21, 3, 21, 326, 8, 21, 1,
		22, 1, 22, 1, 22, 1, 22, 3, 22, 332, 8, 22, 1, 23, 1, 23, 1, 23, 1, 23,
		1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1,
		23, 1, 23, 1, 23, 1, 23, 1, 23, 3, 23, 353, 8, 23, 1, 24, 1, 24, 1, 24,
		1, 24, 1, 25, 1, 25, 3, 25, 361, 8, 25, 1, 25, 1, 25, 1, 26, 4, 26, 366,
		8, 26, 11, 26, 12, 26, 367, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1,
		27, 3, 27, 377, 8, 27, 1, 28, 1, 28, 1, 28, 1, 28, 3, 28, 383, 8, 28, 1,
		29, 1, 29, 1, 29, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30,
		1, 30, 3, 30, 397, 8, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 5,
		30, 405, 8, 30, 10, 30, 12, 30, 408, 9, 30, 1, 31, 1, 31, 1, 31, 1, 31,
		1, 31, 4, 31, 415, 8, 31, 11, 31, 12, 31, 416, 1, 32, 1, 32, 1, 32, 3,
		32, 422, 8, 32, 1, 32, 1, 32, 1, 33, 1, 33, 1, 33, 3, 33, 429, 8, 33, 1,
		33, 1, 33, 1, 34, 1, 34, 1, 34, 3, 34, 436, 8, 34, 1, 35, 1, 35, 1, 35,
		1, 35, 1, 35, 1, 35, 3, 35, 444, 8, 35, 1, 36, 1, 36, 3, 36, 448, 8, 36,
		1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 3, 36, 457, 8, 36, 1,
		37, 1, 37, 1, 38, 1, 38, 1, 38, 1, 38, 3, 38, 465, 8, 38, 1, 38, 1, 38,
		1, 38, 1, 38, 1, 38, 3, 38, 472, 8, 38, 3, 38, 474, 8, 38, 1, 39, 1, 39,
		1, 39, 1, 39, 3, 39, 480, 8, 39, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 3,
		39, 487, 8, 39, 3, 39, 489, 8, 39, 1, 40, 1, 40, 1, 40, 1, 40, 3, 40, 495,
		8, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 3, 40, 502, 8, 40, 3, 40, 504,
		8, 40, 1, 41, 1, 41, 1, 41, 1, 41, 1, 41, 3, 41, 511, 8, 41, 1, 42, 1,
		42, 1, 43, 1, 43, 1, 43, 1, 43, 1, 43, 5, 43, 520, 8, 43, 10, 43, 12, 43,
		523, 9, 43, 1, 44, 1, 44, 5, 44, 527, 8, 44, 10, 44, 12, 44, 530, 9, 44,
		1, 44, 1, 44, 1, 45, 1, 45, 1, 45, 5, 45, 537, 8, 45, 10, 45, 12, 45, 540,
		9, 45, 1, 45, 1, 45, 1, 45, 1, 45, 1, 45, 1, 45, 3, 45, 548, 8, 45, 1,
		46, 1, 46, 5, 46, 552, 8, 46, 10, 46, 12, 46, 555, 9, 46, 1, 46, 1, 46,
		1, 47, 1, 47, 1, 47, 5, 47, 562, 8, 47, 10, 47, 12, 47, 565, 9, 47, 1,
		47, 1, 47, 1, 47, 1, 47, 1, 47, 1, 47, 1, 47, 3, 47, 574, 8, 47, 1, 47,
		1, 47, 1, 47, 1, 47, 1, 47, 3, 47, 581, 8, 47, 1, 48, 1, 48, 1, 49, 1,
		49, 1, 49, 3, 49, 588, 8, 49, 1, 49, 1, 49, 5, 49, 592, 8, 49, 10, 49,
		12, 49, 595, 9, 49, 1, 49, 1, 49, 1, 50, 1, 50, 1, 50, 1, 50, 3, 50, 603,
		8, 50, 1, 50, 1, 50, 1, 50, 1, 50, 1, 50, 1, 50, 1, 50, 1, 50, 1, 50, 1,
		50, 1, 50, 1, 50, 1, 50, 1, 50, 1, 50, 1, 50, 1, 50, 1, 50, 5, 50, 623,
		8, 50, 10, 50, 12, 50, 626, 9, 50, 1, 51, 1, 51, 1, 51, 1, 51, 1, 51, 1,
		51, 1, 51, 1, 51, 1, 51, 1, 51, 3, 51, 638, 8, 51, 1, 52, 1, 52, 1, 52,
		1, 52, 1, 52, 1, 52, 1, 52, 1, 52, 3, 52, 648, 8, 52, 3, 52, 650, 8, 52,
		1, 53, 1, 53, 1, 53, 3, 53, 655, 8, 53, 1, 54, 1, 54, 1, 54, 3, 54, 660,
		8, 54, 1, 55, 1, 55, 1, 56, 1, 56, 1, 56, 1, 56, 3, 56, 668, 8, 56, 1,
		57, 1, 57, 1, 58, 1, 58, 1, 59, 1, 59, 1, 60, 1, 60, 1, 60, 1, 61, 1, 61,
		1, 61, 1, 62, 1, 62, 1, 62, 0, 2, 60, 100, 63, 0, 2, 4, 6, 8, 10, 12, 14,
		16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50,
		52, 54, 56, 58, 60, 62, 64, 66, 68, 70, 72, 74, 76, 78, 80, 82, 84, 86,
		88, 90, 92, 94, 96, 98, 100, 102, 104, 106, 108, 110, 112, 114, 116, 118,
		120, 122, 124, 0, 15, 2, 0, 44, 44, 50, 50, 1, 0, 63, 68, 1, 0, 58, 59,
		1, 0, 22, 24, 1, 0, 25, 26, 3, 0, 60, 60, 71, 73, 75, 80, 1, 0, 46, 47,
		2, 0, 21, 21, 44, 44, 1, 0, 37, 43, 2, 0, 60, 60, 75, 80, 1, 0, 71, 73,
		4, 0, 60, 60, 62, 62, 71, 73, 75, 75, 1, 0, 81, 83, 1, 0, 85, 86, 1, 0,
		28, 29, 723, 0, 126, 1, 0, 0, 0, 2, 158, 1, 0, 0, 0, 4, 162, 1, 0, 0, 0,
		6, 168, 1, 0, 0, 0, 8, 184, 1, 0, 0, 0, 10, 197, 1, 0, 0, 0, 12, 201, 1,
		0, 0, 0, 14, 211, 1, 0, 0, 0, 16, 215, 1, 0, 0, 0, 18, 230, 1, 0, 0, 0,
		20, 236, 1, 0, 0, 0, 22, 242, 1, 0, 0, 0, 24, 244, 1, 0, 0, 0, 26, 246,
		1, 0, 0, 0, 28, 261, 1, 0, 0, 0, 30, 266, 1, 0, 0, 0, 32, 279, 1, 0, 0,
		0, 34, 281, 1, 0, 0, 0, 36, 283, 1, 0, 0, 0, 38, 291, 1, 0, 0, 0, 40, 319,
		1, 0, 0, 0, 42, 325, 1, 0, 0, 0, 44, 331, 1, 0, 0, 0, 46, 352, 1, 0, 0,
		0, 48, 354, 1, 0, 0, 0, 50, 358, 1, 0, 0, 0, 52, 365, 1, 0, 0, 0, 54, 376,
		1, 0, 0, 0, 56, 382, 1, 0, 0, 0, 58, 384, 1, 0, 0, 0, 60, 396, 1, 0, 0,
		0, 62, 409, 1, 0, 0, 0, 64, 418, 1, 0, 0, 0, 66, 425, 1, 0, 0, 0, 68, 435,
		1, 0, 0, 0, 70, 443, 1, 0, 0, 0, 72, 456, 1, 0, 0, 0, 74, 458, 1, 0, 0,
		0, 76, 460, 1, 0, 0, 0, 78, 475, 1, 0, 0, 0, 80, 490, 1, 0, 0, 0, 82, 505,
		1, 0, 0, 0, 84, 512, 1, 0, 0, 0, 86, 514, 1, 0, 0, 0, 88, 524, 1, 0, 0,
		0, 90, 547, 1, 0, 0, 0, 92, 549, 1, 0, 0, 0, 94, 580, 1, 0, 0, 0, 96, 582,
		1, 0, 0, 0, 98, 584, 1, 0, 0, 0, 100, 602, 1, 0, 0, 0, 102, 637, 1, 0,
		0, 0, 104, 649, 1, 0, 0, 0, 106, 654, 1, 0, 0, 0, 108, 659, 1, 0, 0, 0,
		110, 661, 1, 0, 0, 0, 112, 667, 1, 0, 0, 0, 114, 669, 1, 0, 0, 0, 116,
		671, 1, 0, 0, 0, 118, 673, 1, 0, 0, 0, 120, 675, 1, 0, 0, 0, 122, 678,
		1, 0, 0, 0, 124, 681, 1, 0, 0, 0, 126, 130, 3, 2, 1, 0, 127, 129, 3, 16,
		8, 0, 128, 127, 1, 0, 0, 0, 129, 132, 1, 0, 0, 0, 130, 128, 1, 0, 0, 0,
		130, 131, 1, 0, 0, 0, 131, 136, 1, 0, 0, 0, 132, 130, 1, 0, 0, 0, 133,
		135, 3, 4, 2, 0, 134, 133, 1, 0, 0, 0, 135, 138, 1, 0, 0, 0, 136, 134,
		1, 0, 0, 0, 136, 137, 1, 0, 0, 0, 137, 142, 1, 0, 0, 0, 138, 136, 1, 0,
		0, 0, 139, 141, 3, 6, 3, 0, 140, 139, 1, 0, 0, 0, 141, 144, 1, 0, 0, 0,
		142, 140, 1, 0, 0, 0, 142, 143, 1, 0, 0, 0, 143, 148, 1, 0, 0, 0, 144,
		142, 1, 0, 0, 0, 145, 147, 3, 8, 4, 0, 146, 145, 1, 0, 0, 0, 147, 150,
		1, 0, 0, 0, 148, 146, 1, 0, 0, 0, 148, 149, 1, 0, 0, 0, 149, 153, 1, 0,
		0, 0, 150, 148, 1, 0, 0, 0, 151, 154, 3, 64, 32, 0, 152, 154, 3, 66, 33,
		0, 153, 151, 1, 0, 0, 0, 153, 152, 1, 0, 0, 0, 153, 154, 1, 0, 0, 0, 154,
		156, 1, 0, 0, 0, 155, 157, 3, 82, 41, 0, 156, 155, 1, 0, 0, 0, 156, 157,
		1, 0, 0, 0, 157, 1, 1, 0, 0, 0, 158, 159, 5, 33, 0, 0, 159, 160, 5, 44,
		0, 0, 160, 161, 3, 124, 62, 0, 161, 3, 1, 0, 0, 0, 162, 163, 5, 32, 0,
		0, 163, 164, 5, 44, 0, 0, 164, 165, 5, 45, 0, 0, 165, 166, 3, 102, 51,
		0, 166, 167, 3, 124, 62, 0, 167, 5, 1, 0, 0, 0, 168, 169, 5, 31, 0, 0,
		169, 170, 5, 44, 0, 0, 170, 171, 5, 45, 0, 0, 171, 172, 5, 35, 0, 0, 172,
		178, 5, 53, 0, 0, 173, 174, 3, 44, 22, 0, 174, 175, 5, 49, 0, 0, 175, 177,
		1, 0, 0, 0, 176, 173, 1, 0, 0, 0, 177, 180, 1, 0, 0, 0, 178, 176, 1, 0,
		0, 0, 178, 179, 1, 0, 0, 0, 179, 181, 1, 0, 0, 0, 180, 178, 1, 0, 0, 0,
		181, 182, 5, 54, 0, 0, 182, 183, 3, 124, 62, 0, 183, 7, 1, 0, 0, 0, 184,
		185, 5, 34, 0, 0, 185, 191, 5, 53, 0, 0, 186, 187, 3, 10, 5, 0, 187, 188,
		5, 49, 0, 0, 188, 190, 1, 0, 0, 0, 189, 186, 1, 0, 0, 0, 190, 193, 1, 0,
		0, 0, 191, 189, 1, 0, 0, 0, 191, 192, 1, 0, 0, 0, 192, 194, 1, 0, 0, 0,
		193, 191, 1, 0, 0, 0, 194, 195, 5, 54, 0, 0, 195, 196, 3, 124, 62, 0, 196,
		9, 1, 0, 0, 0, 197, 198, 5, 44, 0, 0, 198, 199, 5, 48, 0, 0, 199, 200,
		5, 44, 0, 0, 200, 11, 1, 0, 0, 0, 201, 205, 3, 14, 7, 0, 202, 204, 3, 22,
		11, 0, 203, 202, 1, 0, 0, 0, 204, 207, 1, 0, 0, 0, 205, 203, 1, 0, 0, 0,
		205, 206, 1, 0, 0, 0, 206, 209, 1, 0, 0, 0, 207, 205, 1, 0, 0, 0, 208,
		210, 3, 82, 41, 0, 209, 208, 1, 0, 0, 0, 209, 210, 1, 0, 0, 0, 210, 13,
		1, 0, 0, 0, 211, 212, 5, 17, 0, 0, 212, 213, 5, 44, 0, 0, 213, 214, 3,
		124, 62, 0, 214, 15, 1, 0, 0, 0, 215, 225, 5, 12, 0, 0, 216, 226, 3, 18,
		9, 0, 217, 221, 5, 51, 0, 0, 218, 220, 3, 18, 9, 0, 219, 218, 1, 0, 0,
		0, 220, 223, 1, 0, 0, 0, 221, 219, 1, 0, 0, 0, 221, 222, 1, 0, 0, 0, 222,
		224, 1, 0, 0, 0, 223, 221, 1, 0, 0, 0, 224, 226, 5, 52, 0, 0, 225, 216,
		1, 0, 0, 0, 225, 217, 1, 0, 0, 0, 226, 227, 1, 0, 0, 0, 227, 228, 3, 124,
		62, 0, 228, 17, 1, 0, 0, 0, 229, 231, 7, 0, 0, 0, 230, 229, 1, 0, 0, 0,
		230, 231, 1, 0, 0, 0, 231, 232, 1, 0, 0, 0, 232, 234, 3, 20, 10, 0, 233,
		235, 5, 49, 0, 0, 234, 233, 1, 0, 0, 0, 234, 235, 1, 0, 0, 0, 235, 19,
		1, 0, 0, 0, 236, 237, 3, 116, 58, 0, 237, 21, 1, 0, 0, 0, 238, 243, 3,
		26, 13, 0, 239, 243, 3, 38, 19, 0, 240, 243, 3, 64, 32, 0, 241, 243, 3,
		66, 33, 0, 242, 238, 1, 0, 0, 0, 242, 239, 1, 0, 0, 0, 242, 240, 1, 0,
		0, 0, 242, 241, 1, 0, 0, 0, 243, 23, 1, 0, 0, 0, 244, 245, 7, 1, 0, 0,
		245, 25, 1, 0, 0, 0, 246, 259, 5, 5, 0, 0, 247, 248, 3, 28, 14, 0, 248,
		249, 3, 124, 62, 0, 249, 260, 1, 0, 0, 0, 250, 254, 5, 51, 0, 0, 251, 253,
		3, 28, 14, 0, 252, 251, 1, 0, 0, 0, 253, 256, 1, 0, 0, 0, 254, 252, 1,
		0, 0, 0, 254, 255, 1, 0, 0, 0, 255, 257, 1, 0, 0, 0, 256, 254, 1, 0, 0,
		0, 257, 258, 5, 52, 0, 0, 258, 260, 3, 124, 62, 0, 259, 247, 1, 0, 0, 0,
		259, 250, 1, 0, 0, 0, 260, 27, 1, 0, 0, 0, 261, 264, 3, 30, 15, 0, 262,
		263, 5, 45, 0, 0, 263, 265, 3, 32, 16, 0, 264, 262, 1, 0, 0, 0, 264, 265,
		1, 0, 0, 0, 265, 29, 1, 0, 0, 0, 266, 271, 3, 104, 52, 0, 267, 268, 5,
		49, 0, 0, 268, 270, 3, 104, 52, 0, 269, 267, 1, 0, 0, 0, 270, 273, 1, 0,
		0, 0, 271, 269, 1, 0, 0, 0, 271, 272, 1, 0, 0, 0, 272, 31, 1, 0, 0, 0,
		273, 271, 1, 0, 0, 0, 274, 280, 3, 108, 54, 0, 275, 280, 3, 116, 58, 0,
		276, 280, 3, 118, 59, 0, 277, 280, 3, 98, 49, 0, 278, 280, 3, 34, 17, 0,
		279, 274, 1, 0, 0, 0, 279, 275, 1, 0, 0, 0, 279, 276, 1, 0, 0, 0, 279,
		277, 1, 0, 0, 0, 279, 278, 1, 0, 0, 0, 280, 33, 1, 0, 0, 0, 281, 282, 5,
		27, 0, 0, 282, 35, 1, 0, 0, 0, 283, 288, 3, 100, 50, 0, 284, 285, 5, 49,
		0, 0, 285, 287, 3, 100, 50, 0, 286, 284, 1, 0, 0, 0, 287, 290, 1, 0, 0,
		0, 288, 286, 1, 0, 0, 0, 288, 289, 1, 0, 0, 0, 289, 37, 1, 0, 0, 0, 290,
		288, 1, 0, 0, 0, 291, 292, 5, 6, 0, 0, 292, 293, 5, 44, 0, 0, 293, 294,
		5, 45, 0, 0, 294, 295, 3, 40, 20, 0, 295, 296, 3, 124, 62, 0, 296, 39,
		1, 0, 0, 0, 297, 298, 5, 8, 0, 0, 298, 304, 5, 53, 0, 0, 299, 300, 3, 42,
		21, 0, 300, 301, 5, 49, 0, 0, 301, 303, 1, 0, 0, 0, 302, 299, 1, 0, 0,
		0, 303, 306, 1, 0, 0, 0, 304, 302, 1, 0, 0, 0, 304, 305, 1, 0, 0, 0, 305,
		307, 1, 0, 0, 0, 306, 304, 1, 0, 0, 0, 307, 320, 5, 54, 0, 0, 308, 309,
		5, 18, 0, 0, 309, 315, 5, 53, 0, 0, 310, 311, 3, 42, 21, 0, 311, 312, 5,
		49, 0, 0, 312, 314, 1, 0, 0, 0, 313, 310, 1, 0, 0, 0, 314, 317, 1, 0, 0,
		0, 315, 313, 1, 0, 0, 0, 315, 316, 1, 0, 0, 0, 316, 318, 1, 0, 0, 0, 317,
		315, 1, 0, 0, 0, 318, 320, 5, 54, 0, 0, 319, 297, 1, 0, 0, 0, 319, 308,
		1, 0, 0, 0, 320, 41, 1, 0, 0, 0, 321, 322, 5, 44, 0, 0, 322, 323, 5, 48,
		0, 0, 323, 326, 3, 120, 60, 0, 324, 326, 3, 46, 23, 0, 325, 321, 1, 0,
		0, 0, 325, 324, 1, 0, 0, 0, 326, 43, 1, 0, 0, 0, 327, 328, 5, 44, 0, 0,
		328, 329, 5, 48, 0, 0, 329, 332, 3, 122, 61, 0, 330, 332, 3, 46, 23, 0,
		331, 327, 1, 0, 0, 0, 331, 330, 1, 0, 0, 0, 332, 45, 1, 0, 0, 0, 333, 334,
		5, 44, 0, 0, 334, 335, 5, 48, 0, 0, 335, 353, 3, 108, 54, 0, 336, 337,
		5, 44, 0, 0, 337, 338, 5, 48, 0, 0, 338, 353, 3, 116, 58, 0, 339, 340,
		5, 44, 0, 0, 340, 341, 5, 48, 0, 0, 341, 353, 3, 118, 59, 0, 342, 343,
		5, 44, 0, 0, 343, 344, 5, 48, 0, 0, 344, 353, 3, 104, 52, 0, 345, 346,
		5, 44, 0, 0, 346, 347, 5, 48, 0, 0, 347, 353, 3, 106, 53, 0, 348, 349,
		5, 44, 0, 0, 349, 350, 5, 48, 0, 0, 350, 353, 3, 98, 49, 0, 351, 353, 5,
		44, 0, 0, 352, 333, 1, 0, 0, 0, 352, 336, 1, 0, 0, 0, 352, 339, 1, 0, 0,
		0, 352, 342, 1, 0, 0, 0, 352, 345, 1, 0, 0, 0, 352, 348, 1, 0, 0, 0, 352,
		351, 1, 0, 0, 0, 353, 47, 1, 0, 0, 0, 354, 355, 5, 13, 0, 0, 355, 356,
		3, 102, 51, 0, 356, 357, 3, 124, 62, 0, 357, 49, 1, 0, 0, 0, 358, 360,
		5, 53, 0, 0, 359, 361, 3, 52, 26, 0, 360, 359, 1, 0, 0, 0, 360, 361, 1,
		0, 0, 0, 361, 362, 1, 0, 0, 0, 362, 363, 5, 54, 0, 0, 363, 51, 1, 0, 0,
		0, 364, 366, 3, 54, 27, 0, 365, 364, 1, 0, 0, 0, 366, 367, 1, 0, 0, 0,
		367, 365, 1, 0, 0, 0, 367, 368, 1, 0, 0, 0, 368, 53, 1, 0, 0, 0, 369, 377,
		3, 26, 13, 0, 370, 377, 3, 48, 24, 0, 371, 372, 3, 56, 28, 0, 372, 373,
		3, 124, 62, 0, 373, 377, 1, 0, 0, 0, 374, 377, 3, 50, 25, 0, 375, 377,
		3, 76, 38, 0, 376, 369, 1, 0, 0, 0, 376, 370, 1, 0, 0, 0, 376, 371, 1,
		0, 0, 0, 376, 374, 1, 0, 0, 0, 376, 375, 1, 0, 0, 0, 377, 55, 1, 0, 0,
		0, 378, 383, 3, 100, 50, 0, 379, 383, 3, 58, 29, 0, 380, 383, 3, 72, 36,
		0, 381, 383, 3, 74, 37, 0, 382, 378, 1, 0, 0, 0, 382, 379, 1, 0, 0, 0,
		382, 380, 1, 0, 0, 0, 382, 381, 1, 0, 0, 0, 383, 57, 1, 0, 0, 0, 384, 385,
		3, 100, 50, 0, 385, 386, 7, 2, 0, 0, 386, 59, 1, 0, 0, 0, 387, 388, 6,
		30, -1, 0, 388, 389, 5, 30, 0, 0, 389, 390, 5, 51, 0, 0, 390, 391, 3, 86,
		43, 0, 391, 392, 5, 52, 0, 0, 392, 397, 1, 0, 0, 0, 393, 394, 5, 36, 0,
		0, 394, 395, 5, 51, 0, 0, 395, 397, 5, 52, 0, 0, 396, 387, 1, 0, 0, 0,
		396, 393, 1, 0, 0, 0, 397, 406, 1, 0, 0, 0, 398, 399, 10, 2, 0, 0, 399,
		400, 5, 61, 0, 0, 400, 405, 3, 60, 30, 3, 401, 402, 10, 1, 0, 0, 402, 403,
		5, 69, 0, 0, 403, 405, 3, 60, 30, 2, 404, 398, 1, 0, 0, 0, 404, 401, 1,
		0, 0, 0, 405, 408, 1, 0, 0, 0, 406, 404, 1, 0, 0, 0, 406, 407, 1, 0, 0,
		0, 407, 61, 1, 0, 0, 0, 408, 406, 1, 0, 0, 0, 409, 414, 3, 104, 52, 0,
		410, 411, 5, 55, 0, 0, 411, 412, 3, 100, 50, 0, 412, 413, 5, 56, 0, 0,
		413, 415, 1, 0, 0, 0, 414, 410, 1, 0, 0, 0, 415, 416, 1, 0, 0, 0, 416,
		414, 1, 0, 0, 0, 416, 417, 1, 0, 0, 0, 417, 63, 1, 0, 0, 0, 418, 419, 5,
		2, 0, 0, 419, 421, 3, 70, 35, 0, 420, 422, 3, 68, 34, 0, 421, 420, 1, 0,
		0, 0, 421, 422, 1, 0, 0, 0, 422, 423, 1, 0, 0, 0, 423, 424, 3, 124, 62,
		0, 424, 65, 1, 0, 0, 0, 425, 426, 5, 3, 0, 0, 426, 428, 3, 70, 35, 0, 427,
		429, 3, 68, 34, 0, 428, 427, 1, 0, 0, 0, 428, 429, 1, 0, 0, 0, 429, 430,
		1, 0, 0, 0, 430, 431, 3, 124, 62, 0, 431, 67, 1, 0, 0, 0, 432, 436, 7,
		3, 0, 0, 433, 434, 7, 4, 0, 0, 434, 436, 3, 110, 55, 0, 435, 432, 1, 0,
		0, 0, 435, 433, 1, 0, 0, 0, 436, 69, 1, 0, 0, 0, 437, 444, 3, 100, 50,
		0, 438, 439, 5, 20, 0, 0, 439, 440, 3, 100, 50, 0, 440, 441, 5, 19, 0,
		0, 441, 442, 3, 100, 50, 0, 442, 444, 1, 0, 0, 0, 443, 437, 1, 0, 0, 0,
		443, 438, 1, 0, 0, 0, 444, 71, 1, 0, 0, 0, 445, 447, 3, 36, 18, 0, 446,
		448, 7, 5, 0, 0, 447, 446, 1, 0, 0, 0, 447, 448, 1, 0, 0, 0, 448, 449,
		1, 0, 0, 0, 449, 450, 5, 45, 0, 0, 450, 451, 3, 36, 18, 0, 451, 457, 1,
		0, 0, 0, 452, 453, 3, 36, 18, 0, 453, 454, 7, 6, 0, 0, 454, 455, 3, 36,
		18, 0, 455, 457, 1, 0, 0, 0, 456, 445, 1, 0, 0, 0, 456, 452, 1, 0, 0, 0,
		457, 73, 1, 0, 0, 0, 458, 459, 5, 57, 0, 0, 459, 75, 1, 0, 0, 0, 460, 464,
		5, 11, 0, 0, 461, 462, 3, 56, 28, 0, 462, 463, 5, 57, 0, 0, 463, 465, 1,
		0, 0, 0, 464, 461, 1, 0, 0, 0, 464, 465, 1, 0, 0, 0, 465, 466, 1, 0, 0,
		0, 466, 467, 3, 100, 50, 0, 467, 473, 3, 50, 25, 0, 468, 471, 5, 7, 0,
		0, 469, 472, 3, 76, 38, 0, 470, 472, 3, 50, 25, 0, 471, 469, 1, 0, 0, 0,
		471, 470, 1, 0, 0, 0, 472, 474, 1, 0, 0, 0, 473, 468, 1, 0, 0, 0, 473,
		474, 1, 0, 0, 0, 474, 77, 1, 0, 0, 0, 475, 479, 5, 11, 0, 0, 476, 477,
		3, 56, 28, 0, 477, 478, 5, 57, 0, 0, 478, 480, 1, 0, 0, 0, 479, 476, 1,
		0, 0, 0, 479, 480, 1, 0, 0, 0, 480, 481, 1, 0, 0, 0, 481, 482, 3, 100,
		50, 0, 482, 488, 3, 92, 46, 0, 483, 486, 5, 7, 0, 0, 484, 487, 3, 78, 39,
		0, 485, 487, 3, 92, 46, 0, 486, 484, 1, 0, 0, 0, 486, 485, 1, 0, 0, 0,
		487, 489, 1, 0, 0, 0, 488, 483, 1, 0, 0, 0, 488, 489, 1, 0, 0, 0, 489,
		79, 1, 0, 0, 0, 490, 494, 5, 11, 0, 0, 491, 492, 3, 56, 28, 0, 492, 493,
		5, 57, 0, 0, 493, 495, 1, 0, 0, 0, 494, 491, 1, 0, 0, 0, 494, 495, 1, 0,
		0, 0, 495, 496, 1, 0, 0, 0, 496, 497, 3, 100, 50, 0, 497, 503, 3, 88, 44,
		0, 498, 501, 5, 7, 0, 0, 499, 502, 3, 80, 40, 0, 500, 502, 3, 88, 44, 0,
		501, 499, 1, 0, 0, 0, 501, 500, 1, 0, 0, 0, 502, 504, 1, 0, 0, 0, 503,
		498, 1, 0, 0, 0, 503, 504, 1, 0, 0, 0, 504, 81, 1, 0, 0, 0, 505, 506, 5,
		9, 0, 0, 506, 507, 3, 84, 42, 0, 507, 508, 5, 16, 0, 0, 508, 510, 3, 92,
		46, 0, 509, 511, 3, 124, 62, 0, 510, 509, 1, 0, 0, 0, 510, 511, 1, 0, 0,
		0, 511, 83, 1, 0, 0, 0, 512, 513, 3, 110, 55, 0, 513, 85, 1, 0, 0, 0, 514,
		515, 7, 7, 0, 0, 515, 516, 5, 50, 0, 0, 516, 521, 5, 44, 0, 0, 517, 518,
		5, 50, 0, 0, 518, 520, 5, 44, 0, 0, 519, 517, 1, 0, 0, 0, 520, 523, 1,
		0, 0, 0, 521, 519, 1, 0, 0, 0, 521, 522, 1, 0, 0, 0, 522, 87, 1, 0, 0,
		0, 523, 521, 1, 0, 0, 0, 524, 528, 5, 53, 0, 0, 525, 527, 3, 90, 45, 0,
		526, 525, 1, 0, 0, 0, 527, 530, 1, 0, 0, 0, 528, 526, 1, 0, 0, 0, 528,
		529, 1, 0, 0, 0, 529, 531, 1, 0, 0, 0, 530, 528, 1, 0, 0, 0, 531, 532,
		5, 54, 0, 0, 532, 89, 1, 0, 0, 0, 533, 538, 3, 86, 43, 0, 534, 535, 5,
		70, 0, 0, 535, 537, 3, 86, 43, 0, 536, 534, 1, 0, 0, 0, 537, 540, 1, 0,
		0, 0, 538, 536, 1, 0, 0, 0, 538, 539, 1, 0, 0, 0, 539, 541, 1, 0, 0, 0,
		540, 538, 1, 0, 0, 0, 541, 542, 3, 124, 62, 0, 542, 548, 1, 0, 0, 0, 543,
		544, 3, 60, 30, 0, 544, 545, 3, 124, 62, 0, 545, 548, 1, 0, 0, 0, 546,
		548, 3, 80, 40, 0, 547, 533, 1, 0, 0, 0, 547, 543, 1, 0, 0, 0, 547, 546,
		1, 0, 0, 0, 548, 91, 1, 0, 0, 0, 549, 553, 5, 53, 0, 0, 550, 552, 3, 94,
		47, 0, 551, 550, 1, 0, 0, 0, 552, 555, 1, 0, 0, 0, 553, 551, 1, 0, 0, 0,
		553, 554, 1, 0, 0, 0, 554, 556, 1, 0, 0, 0, 555, 553, 1, 0, 0, 0, 556,
		557, 5, 54, 0, 0, 557, 93, 1, 0, 0, 0, 558, 563, 3, 86, 43, 0, 559, 560,
		5, 70, 0, 0, 560, 562, 3, 86, 43, 0, 561, 559, 1, 0, 0, 0, 562, 565, 1,
		0, 0, 0, 563, 561, 1, 0, 0, 0, 563, 564, 1, 0, 0, 0, 564, 566, 1, 0, 0,
		0, 565, 563, 1, 0, 0, 0, 566, 567, 3, 124, 62, 0, 567, 581, 1, 0, 0, 0,
		568, 569, 5, 44, 0, 0, 569, 570, 5, 45, 0, 0, 570, 573, 5, 14, 0, 0, 571,
		574, 3, 86, 43, 0, 572, 574, 5, 44, 0, 0, 573, 571, 1, 0, 0, 0, 573, 572,
		1, 0, 0, 0, 574, 575, 1, 0, 0, 0, 575, 581, 3, 124, 62, 0, 576, 577, 3,
		56, 28, 0, 577, 578, 3, 124, 62, 0, 578, 581, 1, 0, 0, 0, 579, 581, 3,
		78, 39, 0, 580, 558, 1, 0, 0, 0, 580, 568, 1, 0, 0, 0, 580, 576, 1, 0,
		0, 0, 580, 579, 1, 0, 0, 0, 581, 95, 1, 0, 0, 0, 582, 583, 7, 8, 0, 0,
		583, 97, 1, 0, 0, 0, 584, 585, 3, 96, 48, 0, 585, 587, 5, 51, 0, 0, 586,
		588, 3, 102, 51, 0, 587, 586, 1, 0, 0, 0, 587, 588, 1, 0, 0, 0, 588, 593,
		1, 0, 0, 0, 589, 590, 5, 49, 0, 0, 590, 592, 3, 102, 51, 0, 591, 589, 1,
		0, 0, 0, 592, 595, 1, 0, 0, 0, 593, 591, 1, 0, 0, 0, 593, 594, 1, 0, 0,
		0, 594, 596, 1, 0, 0, 0, 595, 593, 1, 0, 0, 0, 596, 597, 5, 52, 0, 0, 597,
		99, 1, 0, 0, 0, 598, 599, 6, 50, -1, 0, 599, 603, 3, 102, 51, 0, 600, 603,
		3, 98, 49, 0, 601, 603, 3, 106, 53, 0, 602, 598, 1, 0, 0, 0, 602, 600,
		1, 0, 0, 0, 602, 601, 1, 0, 0, 0, 603, 624, 1, 0, 0, 0, 604, 605, 10, 6,
		0, 0, 605, 606, 5, 74, 0, 0, 606, 623, 3, 100, 50, 7, 607, 608, 10, 5,
		0, 0, 608, 609, 7, 9, 0, 0, 609, 623, 3, 100, 50, 6, 610, 611, 10, 4, 0,
		0, 611, 612, 7, 10, 0, 0, 612, 623, 3, 100, 50, 5, 613, 614, 10, 3, 0,
		0, 614, 615, 7, 1, 0, 0, 615, 623, 3, 100, 50, 4, 616, 617, 10, 2, 0, 0,
		617, 618, 5, 61, 0, 0, 618, 623, 3, 100, 50, 3, 619, 620, 10, 1, 0, 0,
		620, 621, 5, 69, 0, 0, 621, 623, 3, 100, 50, 2, 622, 604, 1, 0, 0, 0, 622,
		607, 1, 0, 0, 0, 622, 610, 1, 0, 0, 0, 622, 613, 1, 0, 0, 0, 622, 616,
		1, 0, 0, 0, 622, 619, 1, 0, 0, 0, 623, 626, 1, 0, 0, 0, 624, 622, 1, 0,
		0, 0, 624, 625, 1, 0, 0, 0, 625, 101, 1, 0, 0, 0, 626, 624, 1, 0, 0, 0,
		627, 638, 3, 34, 17, 0, 628, 638, 3, 108, 54, 0, 629, 638, 3, 116, 58,
		0, 630, 638, 3, 118, 59, 0, 631, 638, 3, 104, 52, 0, 632, 638, 3, 62, 31,
		0, 633, 634, 5, 51, 0, 0, 634, 635, 3, 100, 50, 0, 635, 636, 5, 52, 0,
		0, 636, 638, 1, 0, 0, 0, 637, 627, 1, 0, 0, 0, 637, 628, 1, 0, 0, 0, 637,
		629, 1, 0, 0, 0, 637, 630, 1, 0, 0, 0, 637, 631, 1, 0, 0, 0, 637, 632,
		1, 0, 0, 0, 637, 633, 1, 0, 0, 0, 638, 103, 1, 0, 0, 0, 639, 650, 5, 44,
		0, 0, 640, 650, 3, 86, 43, 0, 641, 650, 5, 21, 0, 0, 642, 650, 5, 4, 0,
		0, 643, 644, 5, 14, 0, 0, 644, 647, 5, 44, 0, 0, 645, 646, 5, 50, 0, 0,
		646, 648, 5, 44, 0, 0, 647, 645, 1, 0, 0, 0, 647, 648, 1, 0, 0, 0, 648,
		650, 1, 0, 0, 0, 649, 639, 1, 0, 0, 0, 649, 640, 1, 0, 0, 0, 649, 641,
		1, 0, 0, 0, 649, 642, 1, 0, 0, 0, 649, 643, 1, 0, 0, 0, 650, 105, 1, 0,
		0, 0, 651, 655, 1, 0, 0, 0, 652, 653, 7, 11, 0, 0, 653, 655, 3, 100, 50,
		0, 654, 651, 1, 0, 0, 0, 654, 652, 1, 0, 0, 0, 655, 107, 1, 0, 0, 0, 656,
		660, 3, 110, 55, 0, 657, 660, 3, 112, 56, 0, 658, 660, 3, 114, 57, 0, 659,
		656, 1, 0, 0, 0, 659, 657, 1, 0, 0, 0, 659, 658, 1, 0, 0, 0, 660, 109,
		1, 0, 0, 0, 661, 662, 7, 12, 0, 0, 662, 111, 1, 0, 0, 0, 663, 664, 5, 72,
		0, 0, 664, 668, 3, 110, 55, 0, 665, 666, 5, 72, 0, 0, 666, 668, 3, 114,
		57, 0, 667, 663, 1, 0, 0, 0, 667, 665, 1, 0, 0, 0, 668, 113, 1, 0, 0, 0,
		669, 670, 5, 84, 0, 0, 670, 115, 1, 0, 0, 0, 671, 672, 7, 13, 0, 0, 672,
		117, 1, 0, 0, 0, 673, 674, 7, 14, 0, 0, 674, 119, 1, 0, 0, 0, 675, 676,
		5, 10, 0, 0, 676, 677, 3, 50, 25, 0, 677, 121, 1, 0, 0, 0, 678, 679, 5,
		10, 0, 0, 679, 680, 3, 88, 44, 0, 680, 123, 1, 0, 0, 0, 681, 682, 5, 57,
		0, 0, 682, 125, 1, 0, 0, 0, 70, 130, 136, 142, 148, 153, 156, 178, 191,
		205, 209, 221, 225, 230, 234, 242, 254, 259, 264, 271, 279, 288, 304, 315,
		319, 325, 331, 352, 360, 367, 376, 382, 396, 404, 406, 416, 421, 428, 435,
		443, 447, 456, 464, 471, 473, 479, 486, 488, 494, 501, 503, 510, 521, 528,
		538, 547, 553, 563, 573, 580, 587, 593, 602, 622, 624, 637, 647, 649, 654,
		659, 667,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// FaultParserInit initializes any static state used to implement FaultParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewFaultParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func FaultParserInit() {
	staticData := &faultparserParserStaticData
	staticData.once.Do(faultparserParserInit)
}

// NewFaultParser produces a new parser instance for the optional input antlr.TokenStream.
func NewFaultParser(input antlr.TokenStream) *FaultParser {
	FaultParserInit()
	this := new(FaultParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &faultparserParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "java-escape"

	return this
}

// FaultParser tokens.
const (
	FaultParserEOF                    = antlr.TokenEOF
	FaultParserALL                    = 1
	FaultParserASSERT                 = 2
	FaultParserASSUME                 = 3
	FaultParserCLOCK                  = 4
	FaultParserCONST                  = 5
	FaultParserDEF                    = 6
	FaultParserELSE                   = 7
	FaultParserFLOW                   = 8
	FaultParserFOR                    = 9
	FaultParserFUNC                   = 10
	FaultParserIF                     = 11
	FaultParserIMPORT                 = 12
	FaultParserINIT                   = 13
	FaultParserNEW                    = 14
	FaultParserRETURN                 = 15
	FaultParserRUN                    = 16
	FaultParserSPEC                   = 17
	FaultParserSTOCK                  = 18
	FaultParserTHEN                   = 19
	FaultParserWHEN                   = 20
	FaultParserTHIS                   = 21
	FaultParserEVENTUALLY             = 22
	FaultParserEVENTUALLYALWAYS       = 23
	FaultParserALWAYS                 = 24
	FaultParserNMT                    = 25
	FaultParserNFT                    = 26
	FaultParserNIL                    = 27
	FaultParserTRUE                   = 28
	FaultParserFALSE                  = 29
	FaultParserADVANCE                = 30
	FaultParserCOMPONENT              = 31
	FaultParserGLOBAL                 = 32
	FaultParserSYSTEM                 = 33
	FaultParserSTART                  = 34
	FaultParserSTATE                  = 35
	FaultParserSTAY                   = 36
	FaultParserTY_STRING              = 37
	FaultParserTY_BOOL                = 38
	FaultParserTY_INT                 = 39
	FaultParserTY_FLOAT               = 40
	FaultParserTY_NATURAL             = 41
	FaultParserTY_UNCERTAIN           = 42
	FaultParserTY_UNKNOWN             = 43
	FaultParserIDENT                  = 44
	FaultParserASSIGN                 = 45
	FaultParserASSIGN_FLOW1           = 46
	FaultParserASSIGN_FLOW2           = 47
	FaultParserCOLON                  = 48
	FaultParserCOMMA                  = 49
	FaultParserDOT                    = 50
	FaultParserLPAREN                 = 51
	FaultParserRPAREN                 = 52
	FaultParserLCURLY                 = 53
	FaultParserRCURLY                 = 54
	FaultParserLBRACE                 = 55
	FaultParserRBRACE                 = 56
	FaultParserSEMI                   = 57
	FaultParserPLUS_PLUS              = 58
	FaultParserMINUS_MINUS            = 59
	FaultParserAMPERSAND              = 60
	FaultParserAND                    = 61
	FaultParserBANG                   = 62
	FaultParserEQUALS                 = 63
	FaultParserNOT_EQUALS             = 64
	FaultParserLESS                   = 65
	FaultParserLESS_OR_EQUALS         = 66
	FaultParserGREATER                = 67
	FaultParserGREATER_OR_EQUALS      = 68
	FaultParserOR                     = 69
	FaultParserPIPE                   = 70
	FaultParserPLUS                   = 71
	FaultParserMINUS                  = 72
	FaultParserCARET                  = 73
	FaultParserEXPO                   = 74
	FaultParserMULTI                  = 75
	FaultParserDIV                    = 76
	FaultParserMOD                    = 77
	FaultParserLSHIFT                 = 78
	FaultParserRSHIFT                 = 79
	FaultParserBIT_CLEAR              = 80
	FaultParserDECIMAL_LIT            = 81
	FaultParserOCTAL_LIT              = 82
	FaultParserHEX_LIT                = 83
	FaultParserFLOAT_LIT              = 84
	FaultParserRAW_STRING_LIT         = 85
	FaultParserINTERPRETED_STRING_LIT = 86
	FaultParserWS                     = 87
	FaultParserCOMMENT                = 88
	FaultParserTERMINATOR             = 89
	FaultParserLINE_COMMENT           = 90
)

// FaultParser rules.
const (
	FaultParserRULE_sysSpec          = 0
	FaultParserRULE_sysClause        = 1
	FaultParserRULE_globalDecl       = 2
	FaultParserRULE_componentDecl    = 3
	FaultParserRULE_startBlock       = 4
	FaultParserRULE_startPair        = 5
	FaultParserRULE_spec             = 6
	FaultParserRULE_specClause       = 7
	FaultParserRULE_importDecl       = 8
	FaultParserRULE_importSpec       = 9
	FaultParserRULE_importPath       = 10
	FaultParserRULE_declaration      = 11
	FaultParserRULE_comparison       = 12
	FaultParserRULE_constDecl        = 13
	FaultParserRULE_constSpec        = 14
	FaultParserRULE_identList        = 15
	FaultParserRULE_constants        = 16
	FaultParserRULE_nil              = 17
	FaultParserRULE_expressionList   = 18
	FaultParserRULE_structDecl       = 19
	FaultParserRULE_structType       = 20
	FaultParserRULE_sfProperties     = 21
	FaultParserRULE_comProperties    = 22
	FaultParserRULE_structProperties = 23
	FaultParserRULE_initDecl         = 24
	FaultParserRULE_block            = 25
	FaultParserRULE_statementList    = 26
	FaultParserRULE_statement        = 27
	FaultParserRULE_simpleStmt       = 28
	FaultParserRULE_incDecStmt       = 29
	FaultParserRULE_stateChange      = 30
	FaultParserRULE_accessHistory    = 31
	FaultParserRULE_assertion        = 32
	FaultParserRULE_assumption       = 33
	FaultParserRULE_temporal         = 34
	FaultParserRULE_invariant        = 35
	FaultParserRULE_assignment       = 36
	FaultParserRULE_emptyStmt        = 37
	FaultParserRULE_ifStmt           = 38
	FaultParserRULE_ifStmtRun        = 39
	FaultParserRULE_ifStmtState      = 40
	FaultParserRULE_forStmt          = 41
	FaultParserRULE_rounds           = 42
	FaultParserRULE_paramCall        = 43
	FaultParserRULE_stateBlock       = 44
	FaultParserRULE_stateStep        = 45
	FaultParserRULE_runBlock         = 46
	FaultParserRULE_runStep          = 47
	FaultParserRULE_faultType        = 48
	FaultParserRULE_solvable         = 49
	FaultParserRULE_expression       = 50
	FaultParserRULE_operand          = 51
	FaultParserRULE_operandName      = 52
	FaultParserRULE_prefix           = 53
	FaultParserRULE_numeric          = 54
	FaultParserRULE_integer          = 55
	FaultParserRULE_negative         = 56
	FaultParserRULE_float_           = 57
	FaultParserRULE_string_          = 58
	FaultParserRULE_bool_            = 59
	FaultParserRULE_functionLit      = 60
	FaultParserRULE_stateLit         = 61
	FaultParserRULE_eos              = 62
)

// ISysSpecContext is an interface to support dynamic dispatch.
type ISysSpecContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSysSpecContext differentiates from other interfaces.
	IsSysSpecContext()
}

type SysSpecContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySysSpecContext() *SysSpecContext {
	var p = new(SysSpecContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FaultParserRULE_sysSpec
	return p
}

func (*SysSpecContext) IsSysSpecContext() {}

func NewSysSpecContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SysSpecContext {
	var p = new(SysSpecContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FaultParserRULE_sysSpec

	return p
}

func (s *SysSpecContext) GetParser() antlr.Parser { return s.parser }

func (s *SysSpecContext) SysClause() ISysClauseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISysClauseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISysClauseContext)
}

func (s *SysSpecContext) AllImportDecl() []IImportDeclContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IImportDeclContext); ok {
			len++
		}
	}

	tst := make([]IImportDeclContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IImportDeclContext); ok {
			tst[i] = t.(IImportDeclContext)
			i++
		}
	}

	return tst
}

func (s *SysSpecContext) ImportDecl(i int) IImportDeclContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IImportDeclContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IImportDeclContext)
}

func (s *SysSpecContext) AllGlobalDecl() []IGlobalDeclContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IGlobalDeclContext); ok {
			len++
		}
	}

	tst := make([]IGlobalDeclContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IGlobalDeclContext); ok {
			tst[i] = t.(IGlobalDeclContext)
			i++
		}
	}

	return tst
}

func (s *SysSpecContext) GlobalDecl(i int) IGlobalDeclContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGlobalDeclContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGlobalDeclContext)
}

func (s *SysSpecContext) AllComponentDecl() []IComponentDeclContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IComponentDeclContext); ok {
			len++
		}
	}

	tst := make([]IComponentDeclContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IComponentDeclContext); ok {
			tst[i] = t.(IComponentDeclContext)
			i++
		}
	}

	return tst
}

func (s *SysSpecContext) ComponentDecl(i int) IComponentDeclContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IComponentDeclContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IComponentDeclContext)
}

func (s *SysSpecContext) AllStartBlock() []IStartBlockContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStartBlockContext); ok {
			len++
		}
	}

	tst := make([]IStartBlockContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStartBlockContext); ok {
			tst[i] = t.(IStartBlockContext)
			i++
		}
	}

	return tst
}

func (s *SysSpecContext) StartBlock(i int) IStartBlockContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStartBlockContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStartBlockContext)
}

func (s *SysSpecContext) Assertion() IAssertionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssertionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssertionContext)
}

func (s *SysSpecContext) Assumption() IAssumptionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssumptionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssumptionContext)
}

func (s *SysSpecContext) ForStmt() IForStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IForStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IForStmtContext)
}

func (s *SysSpecContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SysSpecContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SysSpecContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaultParserListener); ok {
		listenerT.EnterSysSpec(s)
	}
}

func (s *SysSpecContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaultParserListener); ok {
		listenerT.ExitSysSpec(s)
	}
}

func (s *SysSpecContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FaultParserVisitor:
		return t.VisitSysSpec(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FaultParser) SysSpec() (localctx ISysSpecContext) {
	this := p
	_ = this

	localctx = NewSysSpecContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, FaultParserRULE_sysSpec)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(126)
		p.SysClause()
	}
	p.SetState(130)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == FaultParserIMPORT {
		{
			p.SetState(127)
			p.ImportDecl()
		}

		p.SetState(132)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(136)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == FaultParserGLOBAL {
		{
			p.SetState(133)
			p.GlobalDecl()
		}

		p.SetState(138)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(142)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == FaultParserCOMPONENT {
		{
			p.SetState(139)
			p.ComponentDecl()
		}

		p.SetState(144)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(148)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == FaultParserSTART {
		{
			p.SetState(145)
			p.StartBlock()
		}

		p.SetState(150)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(153)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case FaultParserASSERT:
		{
			p.SetState(151)
			p.Assertion()
		}

	case FaultParserASSUME:
		{
			p.SetState(152)
			p.Assumption()
		}

	case FaultParserEOF, FaultParserFOR:

	default:
	}
	p.SetState(156)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == FaultParserFOR {
		{
			p.SetState(155)
			p.ForStmt()
		}

	}

	return localctx
}

// ISysClauseContext is an interface to support dynamic dispatch.
type ISysClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSysClauseContext differentiates from other interfaces.
	IsSysClauseContext()
}

type SysClauseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySysClauseContext() *SysClauseContext {
	var p = new(SysClauseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FaultParserRULE_sysClause
	return p
}

func (*SysClauseContext) IsSysClauseContext() {}

func NewSysClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SysClauseContext {
	var p = new(SysClauseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FaultParserRULE_sysClause

	return p
}

func (s *SysClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *SysClauseContext) SYSTEM() antlr.TerminalNode {
	return s.GetToken(FaultParserSYSTEM, 0)
}

func (s *SysClauseContext) IDENT() antlr.TerminalNode {
	return s.GetToken(FaultParserIDENT, 0)
}

func (s *SysClauseContext) Eos() IEosContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEosContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEosContext)
}

func (s *SysClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SysClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SysClauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaultParserListener); ok {
		listenerT.EnterSysClause(s)
	}
}

func (s *SysClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaultParserListener); ok {
		listenerT.ExitSysClause(s)
	}
}

func (s *SysClauseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FaultParserVisitor:
		return t.VisitSysClause(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FaultParser) SysClause() (localctx ISysClauseContext) {
	this := p
	_ = this

	localctx = NewSysClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, FaultParserRULE_sysClause)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(158)
		p.Match(FaultParserSYSTEM)
	}
	{
		p.SetState(159)
		p.Match(FaultParserIDENT)
	}
	{
		p.SetState(160)
		p.Eos()
	}

	return localctx
}

// IGlobalDeclContext is an interface to support dynamic dispatch.
type IGlobalDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsGlobalDeclContext differentiates from other interfaces.
	IsGlobalDeclContext()
}

type GlobalDeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGlobalDeclContext() *GlobalDeclContext {
	var p = new(GlobalDeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FaultParserRULE_globalDecl
	return p
}

func (*GlobalDeclContext) IsGlobalDeclContext() {}

func NewGlobalDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GlobalDeclContext {
	var p = new(GlobalDeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FaultParserRULE_globalDecl

	return p
}

func (s *GlobalDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *GlobalDeclContext) GLOBAL() antlr.TerminalNode {
	return s.GetToken(FaultParserGLOBAL, 0)
}

func (s *GlobalDeclContext) IDENT() antlr.TerminalNode {
	return s.GetToken(FaultParserIDENT, 0)
}

func (s *GlobalDeclContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(FaultParserASSIGN, 0)
}

func (s *GlobalDeclContext) Operand() IOperandContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOperandContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOperandContext)
}

func (s *GlobalDeclContext) Eos() IEosContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEosContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEosContext)
}

func (s *GlobalDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GlobalDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GlobalDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaultParserListener); ok {
		listenerT.EnterGlobalDecl(s)
	}
}

func (s *GlobalDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaultParserListener); ok {
		listenerT.ExitGlobalDecl(s)
	}
}

func (s *GlobalDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FaultParserVisitor:
		return t.VisitGlobalDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FaultParser) GlobalDecl() (localctx IGlobalDeclContext) {
	this := p
	_ = this

	localctx = NewGlobalDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, FaultParserRULE_globalDecl)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(162)
		p.Match(FaultParserGLOBAL)
	}
	{
		p.SetState(163)
		p.Match(FaultParserIDENT)
	}
	{
		p.SetState(164)
		p.Match(FaultParserASSIGN)
	}
	{
		p.SetState(165)
		p.Operand()
	}
	{
		p.SetState(166)
		p.Eos()
	}

	return localctx
}

// IComponentDeclContext is an interface to support dynamic dispatch.
type IComponentDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsComponentDeclContext differentiates from other interfaces.
	IsComponentDeclContext()
}

type ComponentDeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyComponentDeclContext() *ComponentDeclContext {
	var p = new(ComponentDeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FaultParserRULE_componentDecl
	return p
}

func (*ComponentDeclContext) IsComponentDeclContext() {}

func NewComponentDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ComponentDeclContext {
	var p = new(ComponentDeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FaultParserRULE_componentDecl

	return p
}

func (s *ComponentDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *ComponentDeclContext) COMPONENT() antlr.TerminalNode {
	return s.GetToken(FaultParserCOMPONENT, 0)
}

func (s *ComponentDeclContext) IDENT() antlr.TerminalNode {
	return s.GetToken(FaultParserIDENT, 0)
}

func (s *ComponentDeclContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(FaultParserASSIGN, 0)
}

func (s *ComponentDeclContext) STATE() antlr.TerminalNode {
	return s.GetToken(FaultParserSTATE, 0)
}

func (s *ComponentDeclContext) LCURLY() antlr.TerminalNode {
	return s.GetToken(FaultParserLCURLY, 0)
}

func (s *ComponentDeclContext) RCURLY() antlr.TerminalNode {
	return s.GetToken(FaultParserRCURLY, 0)
}

func (s *ComponentDeclContext) Eos() IEosContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEosContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEosContext)
}

func (s *ComponentDeclContext) AllComProperties() []IComPropertiesContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IComPropertiesContext); ok {
			len++
		}
	}

	tst := make([]IComPropertiesContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IComPropertiesContext); ok {
			tst[i] = t.(IComPropertiesContext)
			i++
		}
	}

	return tst
}

func (s *ComponentDeclContext) ComProperties(i int) IComPropertiesContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IComPropertiesContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IComPropertiesContext)
}

func (s *ComponentDeclContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(FaultParserCOMMA)
}

func (s *ComponentDeclContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(FaultParserCOMMA, i)
}

func (s *ComponentDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ComponentDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ComponentDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaultParserListener); ok {
		listenerT.EnterComponentDecl(s)
	}
}

func (s *ComponentDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaultParserListener); ok {
		listenerT.ExitComponentDecl(s)
	}
}

func (s *ComponentDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FaultParserVisitor:
		return t.VisitComponentDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FaultParser) ComponentDecl() (localctx IComponentDeclContext) {
	this := p
	_ = this

	localctx = NewComponentDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, FaultParserRULE_componentDecl)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(168)
		p.Match(FaultParserCOMPONENT)
	}
	{
		p.SetState(169)
		p.Match(FaultParserIDENT)
	}
	{
		p.SetState(170)
		p.Match(FaultParserASSIGN)
	}
	{
		p.SetState(171)
		p.Match(FaultParserSTATE)
	}
	{
		p.SetState(172)
		p.Match(FaultParserLCURLY)
	}
	p.SetState(178)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == FaultParserIDENT {
		{
			p.SetState(173)
			p.ComProperties()
		}
		{
			p.SetState(174)
			p.Match(FaultParserCOMMA)
		}

		p.SetState(180)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(181)
		p.Match(FaultParserRCURLY)
	}
	{
		p.SetState(182)
		p.Eos()
	}

	return localctx
}

// IStartBlockContext is an interface to support dynamic dispatch.
type IStartBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStartBlockContext differentiates from other interfaces.
	IsStartBlockContext()
}

type StartBlockContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStartBlockContext() *StartBlockContext {
	var p = new(StartBlockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FaultParserRULE_startBlock
	return p
}

func (*StartBlockContext) IsStartBlockContext() {}

func NewStartBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StartBlockContext {
	var p = new(StartBlockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FaultParserRULE_startBlock

	return p
}

func (s *StartBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *StartBlockContext) START() antlr.TerminalNode {
	return s.GetToken(FaultParserSTART, 0)
}

func (s *StartBlockContext) LCURLY() antlr.TerminalNode {
	return s.GetToken(FaultParserLCURLY, 0)
}

func (s *StartBlockContext) RCURLY() antlr.TerminalNode {
	return s.GetToken(FaultParserRCURLY, 0)
}

func (s *StartBlockContext) Eos() IEosContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEosContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEosContext)
}

func (s *StartBlockContext) AllStartPair() []IStartPairContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStartPairContext); ok {
			len++
		}
	}

	tst := make([]IStartPairContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStartPairContext); ok {
			tst[i] = t.(IStartPairContext)
			i++
		}
	}

	return tst
}

func (s *StartBlockContext) StartPair(i int) IStartPairContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStartPairContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStartPairContext)
}

func (s *StartBlockContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(FaultParserCOMMA)
}

func (s *StartBlockContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(FaultParserCOMMA, i)
}

func (s *StartBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StartBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StartBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaultParserListener); ok {
		listenerT.EnterStartBlock(s)
	}
}

func (s *StartBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaultParserListener); ok {
		listenerT.ExitStartBlock(s)
	}
}

func (s *StartBlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FaultParserVisitor:
		return t.VisitStartBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FaultParser) StartBlock() (localctx IStartBlockContext) {
	this := p
	_ = this

	localctx = NewStartBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, FaultParserRULE_startBlock)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(184)
		p.Match(FaultParserSTART)
	}
	{
		p.SetState(185)
		p.Match(FaultParserLCURLY)
	}
	p.SetState(191)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == FaultParserIDENT {
		{
			p.SetState(186)
			p.StartPair()
		}
		{
			p.SetState(187)
			p.Match(FaultParserCOMMA)
		}

		p.SetState(193)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(194)
		p.Match(FaultParserRCURLY)
	}
	{
		p.SetState(195)
		p.Eos()
	}

	return localctx
}

// IStartPairContext is an interface to support dynamic dispatch.
type IStartPairContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStartPairContext differentiates from other interfaces.
	IsStartPairContext()
}

type StartPairContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStartPairContext() *StartPairContext {
	var p = new(StartPairContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FaultParserRULE_startPair
	return p
}

func (*StartPairContext) IsStartPairContext() {}

func NewStartPairContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StartPairContext {
	var p = new(StartPairContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FaultParserRULE_startPair

	return p
}

func (s *StartPairContext) GetParser() antlr.Parser { return s.parser }

func (s *StartPairContext) AllIDENT() []antlr.TerminalNode {
	return s.GetTokens(FaultParserIDENT)
}

func (s *StartPairContext) IDENT(i int) antlr.TerminalNode {
	return s.GetToken(FaultParserIDENT, i)
}

func (s *StartPairContext) COLON() antlr.TerminalNode {
	return s.GetToken(FaultParserCOLON, 0)
}

func (s *StartPairContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StartPairContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StartPairContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaultParserListener); ok {
		listenerT.EnterStartPair(s)
	}
}

func (s *StartPairContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaultParserListener); ok {
		listenerT.ExitStartPair(s)
	}
}

func (s *StartPairContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FaultParserVisitor:
		return t.VisitStartPair(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FaultParser) StartPair() (localctx IStartPairContext) {
	this := p
	_ = this

	localctx = NewStartPairContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, FaultParserRULE_startPair)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(197)
		p.Match(FaultParserIDENT)
	}
	{
		p.SetState(198)
		p.Match(FaultParserCOLON)
	}
	{
		p.SetState(199)
		p.Match(FaultParserIDENT)
	}

	return localctx
}

// ISpecContext is an interface to support dynamic dispatch.
type ISpecContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSpecContext differentiates from other interfaces.
	IsSpecContext()
}

type SpecContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySpecContext() *SpecContext {
	var p = new(SpecContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FaultParserRULE_spec
	return p
}

func (*SpecContext) IsSpecContext() {}

func NewSpecContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SpecContext {
	var p = new(SpecContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FaultParserRULE_spec

	return p
}

func (s *SpecContext) GetParser() antlr.Parser { return s.parser }

func (s *SpecContext) SpecClause() ISpecClauseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISpecClauseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISpecClauseContext)
}

func (s *SpecContext) AllDeclaration() []IDeclarationContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDeclarationContext); ok {
			len++
		}
	}

	tst := make([]IDeclarationContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDeclarationContext); ok {
			tst[i] = t.(IDeclarationContext)
			i++
		}
	}

	return tst
}

func (s *SpecContext) Declaration(i int) IDeclarationContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclarationContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclarationContext)
}

func (s *SpecContext) ForStmt() IForStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IForStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IForStmtContext)
}

func (s *SpecContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SpecContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SpecContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaultParserListener); ok {
		listenerT.EnterSpec(s)
	}
}

func (s *SpecContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaultParserListener); ok {
		listenerT.ExitSpec(s)
	}
}

func (s *SpecContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FaultParserVisitor:
		return t.VisitSpec(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FaultParser) Spec() (localctx ISpecContext) {
	this := p
	_ = this

	localctx = NewSpecContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, FaultParserRULE_spec)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(201)
		p.SpecClause()
	}
	p.SetState(205)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&108) != 0 {
		{
			p.SetState(202)
			p.Declaration()
		}

		p.SetState(207)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(209)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == FaultParserFOR {
		{
			p.SetState(208)
			p.ForStmt()
		}

	}

	return localctx
}

// ISpecClauseContext is an interface to support dynamic dispatch.
type ISpecClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSpecClauseContext differentiates from other interfaces.
	IsSpecClauseContext()
}

type SpecClauseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySpecClauseContext() *SpecClauseContext {
	var p = new(SpecClauseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FaultParserRULE_specClause
	return p
}

func (*SpecClauseContext) IsSpecClauseContext() {}

func NewSpecClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SpecClauseContext {
	var p = new(SpecClauseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FaultParserRULE_specClause

	return p
}

func (s *SpecClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *SpecClauseContext) SPEC() antlr.TerminalNode {
	return s.GetToken(FaultParserSPEC, 0)
}

func (s *SpecClauseContext) IDENT() antlr.TerminalNode {
	return s.GetToken(FaultParserIDENT, 0)
}

func (s *SpecClauseContext) Eos() IEosContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEosContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEosContext)
}

func (s *SpecClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SpecClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SpecClauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaultParserListener); ok {
		listenerT.EnterSpecClause(s)
	}
}

func (s *SpecClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaultParserListener); ok {
		listenerT.ExitSpecClause(s)
	}
}

func (s *SpecClauseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FaultParserVisitor:
		return t.VisitSpecClause(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FaultParser) SpecClause() (localctx ISpecClauseContext) {
	this := p
	_ = this

	localctx = NewSpecClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, FaultParserRULE_specClause)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(211)
		p.Match(FaultParserSPEC)
	}
	{
		p.SetState(212)
		p.Match(FaultParserIDENT)
	}
	{
		p.SetState(213)
		p.Eos()
	}

	return localctx
}

// IImportDeclContext is an interface to support dynamic dispatch.
type IImportDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsImportDeclContext differentiates from other interfaces.
	IsImportDeclContext()
}

type ImportDeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyImportDeclContext() *ImportDeclContext {
	var p = new(ImportDeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FaultParserRULE_importDecl
	return p
}

func (*ImportDeclContext) IsImportDeclContext() {}

func NewImportDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ImportDeclContext {
	var p = new(ImportDeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FaultParserRULE_importDecl

	return p
}

func (s *ImportDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *ImportDeclContext) IMPORT() antlr.TerminalNode {
	return s.GetToken(FaultParserIMPORT, 0)
}

func (s *ImportDeclContext) Eos() IEosContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEosContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEosContext)
}

func (s *ImportDeclContext) AllImportSpec() []IImportSpecContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IImportSpecContext); ok {
			len++
		}
	}

	tst := make([]IImportSpecContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IImportSpecContext); ok {
			tst[i] = t.(IImportSpecContext)
			i++
		}
	}

	return tst
}

func (s *ImportDeclContext) ImportSpec(i int) IImportSpecContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IImportSpecContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IImportSpecContext)
}

func (s *ImportDeclContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(FaultParserLPAREN, 0)
}

func (s *ImportDeclContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(FaultParserRPAREN, 0)
}

func (s *ImportDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ImportDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ImportDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaultParserListener); ok {
		listenerT.EnterImportDecl(s)
	}
}

func (s *ImportDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaultParserListener); ok {
		listenerT.ExitImportDecl(s)
	}
}

func (s *ImportDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FaultParserVisitor:
		return t.VisitImportDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FaultParser) ImportDecl() (localctx IImportDeclContext) {
	this := p
	_ = this

	localctx = NewImportDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, FaultParserRULE_importDecl)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(215)
		p.Match(FaultParserIMPORT)
	}
	p.SetState(225)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case FaultParserIDENT, FaultParserDOT, FaultParserRAW_STRING_LIT, FaultParserINTERPRETED_STRING_LIT:
		{
			p.SetState(216)
			p.ImportSpec()
		}

	case FaultParserLPAREN:
		{
			p.SetState(217)
			p.Match(FaultParserLPAREN)
		}
		p.SetState(221)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for (int64((_la-44)) & ^0x3f) == 0 && ((int64(1)<<(_la-44))&6597069766721) != 0 {
			{
				p.SetState(218)
				p.ImportSpec()
			}

			p.SetState(223)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(224)
			p.Match(FaultParserRPAREN)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	{
		p.SetState(227)
		p.Eos()
	}

	return localctx
}

// IImportSpecContext is an interface to support dynamic dispatch.
type IImportSpecContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsImportSpecContext differentiates from other interfaces.
	IsImportSpecContext()
}

type ImportSpecContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyImportSpecContext() *ImportSpecContext {
	var p = new(ImportSpecContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FaultParserRULE_importSpec
	return p
}

func (*ImportSpecContext) IsImportSpecContext() {}

func NewImportSpecContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ImportSpecContext {
	var p = new(ImportSpecContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FaultParserRULE_importSpec

	return p
}

func (s *ImportSpecContext) GetParser() antlr.Parser { return s.parser }

func (s *ImportSpecContext) ImportPath() IImportPathContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IImportPathContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IImportPathContext)
}

func (s *ImportSpecContext) COMMA() antlr.TerminalNode {
	return s.GetToken(FaultParserCOMMA, 0)
}

func (s *ImportSpecContext) DOT() antlr.TerminalNode {
	return s.GetToken(FaultParserDOT, 0)
}

func (s *ImportSpecContext) IDENT() antlr.TerminalNode {
	return s.GetToken(FaultParserIDENT, 0)
}

func (s *ImportSpecContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ImportSpecContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ImportSpecContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaultParserListener); ok {
		listenerT.EnterImportSpec(s)
	}
}

func (s *ImportSpecContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaultParserListener); ok {
		listenerT.ExitImportSpec(s)
	}
}

func (s *ImportSpecContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FaultParserVisitor:
		return t.VisitImportSpec(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FaultParser) ImportSpec() (localctx IImportSpecContext) {
	this := p
	_ = this

	localctx = NewImportSpecContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, FaultParserRULE_importSpec)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(230)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == FaultParserIDENT || _la == FaultParserDOT {
		{
			p.SetState(229)
			_la = p.GetTokenStream().LA(1)

			if !(_la == FaultParserIDENT || _la == FaultParserDOT) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	}
	{
		p.SetState(232)
		p.ImportPath()
	}
	p.SetState(234)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == FaultParserCOMMA {
		{
			p.SetState(233)
			p.Match(FaultParserCOMMA)
		}

	}

	return localctx
}

// IImportPathContext is an interface to support dynamic dispatch.
type IImportPathContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsImportPathContext differentiates from other interfaces.
	IsImportPathContext()
}

type ImportPathContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyImportPathContext() *ImportPathContext {
	var p = new(ImportPathContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FaultParserRULE_importPath
	return p
}

func (*ImportPathContext) IsImportPathContext() {}

func NewImportPathContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ImportPathContext {
	var p = new(ImportPathContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FaultParserRULE_importPath

	return p
}

func (s *ImportPathContext) GetParser() antlr.Parser { return s.parser }

func (s *ImportPathContext) String_() IString_Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IString_Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IString_Context)
}

func (s *ImportPathContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ImportPathContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ImportPathContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaultParserListener); ok {
		listenerT.EnterImportPath(s)
	}
}

func (s *ImportPathContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaultParserListener); ok {
		listenerT.ExitImportPath(s)
	}
}

func (s *ImportPathContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FaultParserVisitor:
		return t.VisitImportPath(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FaultParser) ImportPath() (localctx IImportPathContext) {
	this := p
	_ = this

	localctx = NewImportPathContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, FaultParserRULE_importPath)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(236)
		p.String_()
	}

	return localctx
}

// IDeclarationContext is an interface to support dynamic dispatch.
type IDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDeclarationContext differentiates from other interfaces.
	IsDeclarationContext()
}

type DeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeclarationContext() *DeclarationContext {
	var p = new(DeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FaultParserRULE_declaration
	return p
}

func (*DeclarationContext) IsDeclarationContext() {}

func NewDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclarationContext {
	var p = new(DeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FaultParserRULE_declaration

	return p
}

func (s *DeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclarationContext) ConstDecl() IConstDeclContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConstDeclContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConstDeclContext)
}

func (s *DeclarationContext) StructDecl() IStructDeclContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStructDeclContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStructDeclContext)
}

func (s *DeclarationContext) Assertion() IAssertionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssertionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssertionContext)
}

func (s *DeclarationContext) Assumption() IAssumptionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssumptionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssumptionContext)
}

func (s *DeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaultParserListener); ok {
		listenerT.EnterDeclaration(s)
	}
}

func (s *DeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaultParserListener); ok {
		listenerT.ExitDeclaration(s)
	}
}

func (s *DeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FaultParserVisitor:
		return t.VisitDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FaultParser) Declaration() (localctx IDeclarationContext) {
	this := p
	_ = this

	localctx = NewDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, FaultParserRULE_declaration)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(242)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case FaultParserCONST:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(238)
			p.ConstDecl()
		}

	case FaultParserDEF:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(239)
			p.StructDecl()
		}

	case FaultParserASSERT:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(240)
			p.Assertion()
		}

	case FaultParserASSUME:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(241)
			p.Assumption()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IComparisonContext is an interface to support dynamic dispatch.
type IComparisonContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsComparisonContext differentiates from other interfaces.
	IsComparisonContext()
}

type ComparisonContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyComparisonContext() *ComparisonContext {
	var p = new(ComparisonContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FaultParserRULE_comparison
	return p
}

func (*ComparisonContext) IsComparisonContext() {}

func NewComparisonContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ComparisonContext {
	var p = new(ComparisonContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FaultParserRULE_comparison

	return p
}

func (s *ComparisonContext) GetParser() antlr.Parser { return s.parser }

func (s *ComparisonContext) EQUALS() antlr.TerminalNode {
	return s.GetToken(FaultParserEQUALS, 0)
}

func (s *ComparisonContext) NOT_EQUALS() antlr.TerminalNode {
	return s.GetToken(FaultParserNOT_EQUALS, 0)
}

func (s *ComparisonContext) LESS() antlr.TerminalNode {
	return s.GetToken(FaultParserLESS, 0)
}

func (s *ComparisonContext) LESS_OR_EQUALS() antlr.TerminalNode {
	return s.GetToken(FaultParserLESS_OR_EQUALS, 0)
}

func (s *ComparisonContext) GREATER() antlr.TerminalNode {
	return s.GetToken(FaultParserGREATER, 0)
}

func (s *ComparisonContext) GREATER_OR_EQUALS() antlr.TerminalNode {
	return s.GetToken(FaultParserGREATER_OR_EQUALS, 0)
}

func (s *ComparisonContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ComparisonContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ComparisonContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaultParserListener); ok {
		listenerT.EnterComparison(s)
	}
}

func (s *ComparisonContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaultParserListener); ok {
		listenerT.ExitComparison(s)
	}
}

func (s *ComparisonContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FaultParserVisitor:
		return t.VisitComparison(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FaultParser) Comparison() (localctx IComparisonContext) {
	this := p
	_ = this

	localctx = NewComparisonContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, FaultParserRULE_comparison)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(244)
		_la = p.GetTokenStream().LA(1)

		if !((int64((_la-63)) & ^0x3f) == 0 && ((int64(1)<<(_la-63))&63) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IConstDeclContext is an interface to support dynamic dispatch.
type IConstDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConstDeclContext differentiates from other interfaces.
	IsConstDeclContext()
}

type ConstDeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConstDeclContext() *ConstDeclContext {
	var p = new(ConstDeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FaultParserRULE_constDecl
	return p
}

func (*ConstDeclContext) IsConstDeclContext() {}

func NewConstDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConstDeclContext {
	var p = new(ConstDeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FaultParserRULE_constDecl

	return p
}

func (s *ConstDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *ConstDeclContext) CONST() antlr.TerminalNode {
	return s.GetToken(FaultParserCONST, 0)
}

func (s *ConstDeclContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(FaultParserLPAREN, 0)
}

func (s *ConstDeclContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(FaultParserRPAREN, 0)
}

func (s *ConstDeclContext) Eos() IEosContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEosContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEosContext)
}

func (s *ConstDeclContext) AllConstSpec() []IConstSpecContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IConstSpecContext); ok {
			len++
		}
	}

	tst := make([]IConstSpecContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IConstSpecContext); ok {
			tst[i] = t.(IConstSpecContext)
			i++
		}
	}

	return tst
}

func (s *ConstDeclContext) ConstSpec(i int) IConstSpecContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConstSpecContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConstSpecContext)
}

func (s *ConstDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConstDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaultParserListener); ok {
		listenerT.EnterConstDecl(s)
	}
}

func (s *ConstDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaultParserListener); ok {
		listenerT.ExitConstDecl(s)
	}
}

func (s *ConstDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FaultParserVisitor:
		return t.VisitConstDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FaultParser) ConstDecl() (localctx IConstDeclContext) {
	this := p
	_ = this

	localctx = NewConstDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, FaultParserRULE_constDecl)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(246)
		p.Match(FaultParserCONST)
	}
	p.SetState(259)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case FaultParserCLOCK, FaultParserNEW, FaultParserTHIS, FaultParserIDENT:
		{
			p.SetState(247)
			p.ConstSpec()
		}
		{
			p.SetState(248)
			p.Eos()
		}

	case FaultParserLPAREN:
		{
			p.SetState(250)
			p.Match(FaultParserLPAREN)
		}
		p.SetState(254)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&17592188157968) != 0 {
			{
				p.SetState(251)
				p.ConstSpec()
			}

			p.SetState(256)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(257)
			p.Match(FaultParserRPAREN)
		}
		{
			p.SetState(258)
			p.Eos()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IConstSpecContext is an interface to support dynamic dispatch.
type IConstSpecContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConstSpecContext differentiates from other interfaces.
	IsConstSpecContext()
}

type ConstSpecContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConstSpecContext() *ConstSpecContext {
	var p = new(ConstSpecContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FaultParserRULE_constSpec
	return p
}

func (*ConstSpecContext) IsConstSpecContext() {}

func NewConstSpecContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConstSpecContext {
	var p = new(ConstSpecContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FaultParserRULE_constSpec

	return p
}

func (s *ConstSpecContext) GetParser() antlr.Parser { return s.parser }

func (s *ConstSpecContext) IdentList() IIdentListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentListContext)
}

func (s *ConstSpecContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(FaultParserASSIGN, 0)
}

func (s *ConstSpecContext) Constants() IConstantsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConstantsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConstantsContext)
}

func (s *ConstSpecContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstSpecContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConstSpecContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaultParserListener); ok {
		listenerT.EnterConstSpec(s)
	}
}

func (s *ConstSpecContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaultParserListener); ok {
		listenerT.ExitConstSpec(s)
	}
}

func (s *ConstSpecContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FaultParserVisitor:
		return t.VisitConstSpec(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FaultParser) ConstSpec() (localctx IConstSpecContext) {
	this := p
	_ = this

	localctx = NewConstSpecContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, FaultParserRULE_constSpec)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(261)
		p.IdentList()
	}
	p.SetState(264)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == FaultParserASSIGN {
		{
			p.SetState(262)
			p.Match(FaultParserASSIGN)
		}
		{
			p.SetState(263)
			p.Constants()
		}

	}

	return localctx
}

// IIdentListContext is an interface to support dynamic dispatch.
type IIdentListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIdentListContext differentiates from other interfaces.
	IsIdentListContext()
}

type IdentListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdentListContext() *IdentListContext {
	var p = new(IdentListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FaultParserRULE_identList
	return p
}

func (*IdentListContext) IsIdentListContext() {}

func NewIdentListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdentListContext {
	var p = new(IdentListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FaultParserRULE_identList

	return p
}

func (s *IdentListContext) GetParser() antlr.Parser { return s.parser }

func (s *IdentListContext) AllOperandName() []IOperandNameContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IOperandNameContext); ok {
			len++
		}
	}

	tst := make([]IOperandNameContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IOperandNameContext); ok {
			tst[i] = t.(IOperandNameContext)
			i++
		}
	}

	return tst
}

func (s *IdentListContext) OperandName(i int) IOperandNameContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOperandNameContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOperandNameContext)
}

func (s *IdentListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(FaultParserCOMMA)
}

func (s *IdentListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(FaultParserCOMMA, i)
}

func (s *IdentListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdentListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaultParserListener); ok {
		listenerT.EnterIdentList(s)
	}
}

func (s *IdentListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaultParserListener); ok {
		listenerT.ExitIdentList(s)
	}
}

func (s *IdentListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FaultParserVisitor:
		return t.VisitIdentList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FaultParser) IdentList() (localctx IIdentListContext) {
	this := p
	_ = this

	localctx = NewIdentListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, FaultParserRULE_identList)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(266)
		p.OperandName()
	}
	p.SetState(271)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == FaultParserCOMMA {
		{
			p.SetState(267)
			p.Match(FaultParserCOMMA)
		}
		{
			p.SetState(268)
			p.OperandName()
		}

		p.SetState(273)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IConstantsContext is an interface to support dynamic dispatch.
type IConstantsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConstantsContext differentiates from other interfaces.
	IsConstantsContext()
}

type ConstantsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConstantsContext() *ConstantsContext {
	var p = new(ConstantsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FaultParserRULE_constants
	return p
}

func (*ConstantsContext) IsConstantsContext() {}

func NewConstantsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConstantsContext {
	var p = new(ConstantsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FaultParserRULE_constants

	return p
}

func (s *ConstantsContext) GetParser() antlr.Parser { return s.parser }

func (s *ConstantsContext) Numeric() INumericContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INumericContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INumericContext)
}

func (s *ConstantsContext) String_() IString_Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IString_Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IString_Context)
}

func (s *ConstantsContext) Bool_() IBool_Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBool_Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBool_Context)
}

func (s *ConstantsContext) Solvable() ISolvableContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISolvableContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISolvableContext)
}

func (s *ConstantsContext) Nil_() INilContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INilContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INilContext)
}

func (s *ConstantsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstantsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConstantsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaultParserListener); ok {
		listenerT.EnterConstants(s)
	}
}

func (s *ConstantsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaultParserListener); ok {
		listenerT.ExitConstants(s)
	}
}

func (s *ConstantsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FaultParserVisitor:
		return t.VisitConstants(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FaultParser) Constants() (localctx IConstantsContext) {
	this := p
	_ = this

	localctx = NewConstantsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, FaultParserRULE_constants)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(279)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case FaultParserMINUS, FaultParserDECIMAL_LIT, FaultParserOCTAL_LIT, FaultParserHEX_LIT, FaultParserFLOAT_LIT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(274)
			p.Numeric()
		}

	case FaultParserRAW_STRING_LIT, FaultParserINTERPRETED_STRING_LIT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(275)
			p.String_()
		}

	case FaultParserTRUE, FaultParserFALSE:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(276)
			p.Bool_()
		}

	case FaultParserTY_STRING, FaultParserTY_BOOL, FaultParserTY_INT, FaultParserTY_FLOAT, FaultParserTY_NATURAL, FaultParserTY_UNCERTAIN, FaultParserTY_UNKNOWN:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(277)
			p.Solvable()
		}

	case FaultParserNIL:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(278)
			p.Nil_()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// INilContext is an interface to support dynamic dispatch.
type INilContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNilContext differentiates from other interfaces.
	IsNilContext()
}

type NilContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNilContext() *NilContext {
	var p = new(NilContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FaultParserRULE_nil
	return p
}

func (*NilContext) IsNilContext() {}

func NewNilContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NilContext {
	var p = new(NilContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FaultParserRULE_nil

	return p
}

func (s *NilContext) GetParser() antlr.Parser { return s.parser }

func (s *NilContext) NIL() antlr.TerminalNode {
	return s.GetToken(FaultParserNIL, 0)
}

func (s *NilContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NilContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NilContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaultParserListener); ok {
		listenerT.EnterNil(s)
	}
}

func (s *NilContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaultParserListener); ok {
		listenerT.ExitNil(s)
	}
}

func (s *NilContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FaultParserVisitor:
		return t.VisitNil(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FaultParser) Nil_() (localctx INilContext) {
	this := p
	_ = this

	localctx = NewNilContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, FaultParserRULE_nil)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(281)
		p.Match(FaultParserNIL)
	}

	return localctx
}

// IExpressionListContext is an interface to support dynamic dispatch.
type IExpressionListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressionListContext differentiates from other interfaces.
	IsExpressionListContext()
}

type ExpressionListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionListContext() *ExpressionListContext {
	var p = new(ExpressionListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FaultParserRULE_expressionList
	return p
}

func (*ExpressionListContext) IsExpressionListContext() {}

func NewExpressionListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionListContext {
	var p = new(ExpressionListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FaultParserRULE_expressionList

	return p
}

func (s *ExpressionListContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionListContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ExpressionListContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExpressionListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(FaultParserCOMMA)
}

func (s *ExpressionListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(FaultParserCOMMA, i)
}

func (s *ExpressionListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaultParserListener); ok {
		listenerT.EnterExpressionList(s)
	}
}

func (s *ExpressionListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaultParserListener); ok {
		listenerT.ExitExpressionList(s)
	}
}

func (s *ExpressionListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FaultParserVisitor:
		return t.VisitExpressionList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FaultParser) ExpressionList() (localctx IExpressionListContext) {
	this := p
	_ = this

	localctx = NewExpressionListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, FaultParserRULE_expressionList)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(283)
		p.expression(0)
	}
	p.SetState(288)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == FaultParserCOMMA {
		{
			p.SetState(284)
			p.Match(FaultParserCOMMA)
		}
		{
			p.SetState(285)
			p.expression(0)
		}

		p.SetState(290)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IStructDeclContext is an interface to support dynamic dispatch.
type IStructDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStructDeclContext differentiates from other interfaces.
	IsStructDeclContext()
}

type StructDeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStructDeclContext() *StructDeclContext {
	var p = new(StructDeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FaultParserRULE_structDecl
	return p
}

func (*StructDeclContext) IsStructDeclContext() {}

func NewStructDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StructDeclContext {
	var p = new(StructDeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FaultParserRULE_structDecl

	return p
}

func (s *StructDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *StructDeclContext) DEF() antlr.TerminalNode {
	return s.GetToken(FaultParserDEF, 0)
}

func (s *StructDeclContext) IDENT() antlr.TerminalNode {
	return s.GetToken(FaultParserIDENT, 0)
}

func (s *StructDeclContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(FaultParserASSIGN, 0)
}

func (s *StructDeclContext) StructType() IStructTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStructTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStructTypeContext)
}

func (s *StructDeclContext) Eos() IEosContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEosContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEosContext)
}

func (s *StructDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StructDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaultParserListener); ok {
		listenerT.EnterStructDecl(s)
	}
}

func (s *StructDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaultParserListener); ok {
		listenerT.ExitStructDecl(s)
	}
}

func (s *StructDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FaultParserVisitor:
		return t.VisitStructDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FaultParser) StructDecl() (localctx IStructDeclContext) {
	this := p
	_ = this

	localctx = NewStructDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, FaultParserRULE_structDecl)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(291)
		p.Match(FaultParserDEF)
	}
	{
		p.SetState(292)
		p.Match(FaultParserIDENT)
	}
	{
		p.SetState(293)
		p.Match(FaultParserASSIGN)
	}
	{
		p.SetState(294)
		p.StructType()
	}
	{
		p.SetState(295)
		p.Eos()
	}

	return localctx
}

// IStructTypeContext is an interface to support dynamic dispatch.
type IStructTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStructTypeContext differentiates from other interfaces.
	IsStructTypeContext()
}

type StructTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStructTypeContext() *StructTypeContext {
	var p = new(StructTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FaultParserRULE_structType
	return p
}

func (*StructTypeContext) IsStructTypeContext() {}

func NewStructTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StructTypeContext {
	var p = new(StructTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FaultParserRULE_structType

	return p
}

func (s *StructTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *StructTypeContext) CopyFrom(ctx *StructTypeContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *StructTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type StockContext struct {
	*StructTypeContext
}

func NewStockContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StockContext {
	var p = new(StockContext)

	p.StructTypeContext = NewEmptyStructTypeContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StructTypeContext))

	return p
}

func (s *StockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StockContext) STOCK() antlr.TerminalNode {
	return s.GetToken(FaultParserSTOCK, 0)
}

func (s *StockContext) LCURLY() antlr.TerminalNode {
	return s.GetToken(FaultParserLCURLY, 0)
}

func (s *StockContext) RCURLY() antlr.TerminalNode {
	return s.GetToken(FaultParserRCURLY, 0)
}

func (s *StockContext) AllSfProperties() []ISfPropertiesContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISfPropertiesContext); ok {
			len++
		}
	}

	tst := make([]ISfPropertiesContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISfPropertiesContext); ok {
			tst[i] = t.(ISfPropertiesContext)
			i++
		}
	}

	return tst
}

func (s *StockContext) SfProperties(i int) ISfPropertiesContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISfPropertiesContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISfPropertiesContext)
}

func (s *StockContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(FaultParserCOMMA)
}

func (s *StockContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(FaultParserCOMMA, i)
}

func (s *StockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaultParserListener); ok {
		listenerT.EnterStock(s)
	}
}

func (s *StockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaultParserListener); ok {
		listenerT.ExitStock(s)
	}
}

func (s *StockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FaultParserVisitor:
		return t.VisitStock(s)

	default:
		return t.VisitChildren(s)
	}
}

type FlowContext struct {
	*StructTypeContext
}

func NewFlowContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FlowContext {
	var p = new(FlowContext)

	p.StructTypeContext = NewEmptyStructTypeContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StructTypeContext))

	return p
}

func (s *FlowContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FlowContext) FLOW() antlr.TerminalNode {
	return s.GetToken(FaultParserFLOW, 0)
}

func (s *FlowContext) LCURLY() antlr.TerminalNode {
	return s.GetToken(FaultParserLCURLY, 0)
}

func (s *FlowContext) RCURLY() antlr.TerminalNode {
	return s.GetToken(FaultParserRCURLY, 0)
}

func (s *FlowContext) AllSfProperties() []ISfPropertiesContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISfPropertiesContext); ok {
			len++
		}
	}

	tst := make([]ISfPropertiesContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISfPropertiesContext); ok {
			tst[i] = t.(ISfPropertiesContext)
			i++
		}
	}

	return tst
}

func (s *FlowContext) SfProperties(i int) ISfPropertiesContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISfPropertiesContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISfPropertiesContext)
}

func (s *FlowContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(FaultParserCOMMA)
}

func (s *FlowContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(FaultParserCOMMA, i)
}

func (s *FlowContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaultParserListener); ok {
		listenerT.EnterFlow(s)
	}
}

func (s *FlowContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaultParserListener); ok {
		listenerT.ExitFlow(s)
	}
}

func (s *FlowContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FaultParserVisitor:
		return t.VisitFlow(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FaultParser) StructType() (localctx IStructTypeContext) {
	this := p
	_ = this

	localctx = NewStructTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, FaultParserRULE_structType)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(319)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case FaultParserFLOW:
		localctx = NewFlowContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(297)
			p.Match(FaultParserFLOW)
		}
		{
			p.SetState(298)
			p.Match(FaultParserLCURLY)
		}
		p.SetState(304)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == FaultParserIDENT {
			{
				p.SetState(299)
				p.SfProperties()
			}
			{
				p.SetState(300)
				p.Match(FaultParserCOMMA)
			}

			p.SetState(306)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(307)
			p.Match(FaultParserRCURLY)
		}

	case FaultParserSTOCK:
		localctx = NewStockContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(308)
			p.Match(FaultParserSTOCK)
		}
		{
			p.SetState(309)
			p.Match(FaultParserLCURLY)
		}
		p.SetState(315)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == FaultParserIDENT {
			{
				p.SetState(310)
				p.SfProperties()
			}
			{
				p.SetState(311)
				p.Match(FaultParserCOMMA)
			}

			p.SetState(317)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(318)
			p.Match(FaultParserRCURLY)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ISfPropertiesContext is an interface to support dynamic dispatch.
type ISfPropertiesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSfPropertiesContext differentiates from other interfaces.
	IsSfPropertiesContext()
}

type SfPropertiesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySfPropertiesContext() *SfPropertiesContext {
	var p = new(SfPropertiesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FaultParserRULE_sfProperties
	return p
}

func (*SfPropertiesContext) IsSfPropertiesContext() {}

func NewSfPropertiesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SfPropertiesContext {
	var p = new(SfPropertiesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FaultParserRULE_sfProperties

	return p
}

func (s *SfPropertiesContext) GetParser() antlr.Parser { return s.parser }

func (s *SfPropertiesContext) CopyFrom(ctx *SfPropertiesContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *SfPropertiesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SfPropertiesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type SfMiscContext struct {
	*SfPropertiesContext
}

func NewSfMiscContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SfMiscContext {
	var p = new(SfMiscContext)

	p.SfPropertiesContext = NewEmptySfPropertiesContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SfPropertiesContext))

	return p
}

func (s *SfMiscContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SfMiscContext) StructProperties() IStructPropertiesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStructPropertiesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStructPropertiesContext)
}

func (s *SfMiscContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaultParserListener); ok {
		listenerT.EnterSfMisc(s)
	}
}

func (s *SfMiscContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaultParserListener); ok {
		listenerT.ExitSfMisc(s)
	}
}

func (s *SfMiscContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FaultParserVisitor:
		return t.VisitSfMisc(s)

	default:
		return t.VisitChildren(s)
	}
}

type PropFuncContext struct {
	*SfPropertiesContext
}

func NewPropFuncContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PropFuncContext {
	var p = new(PropFuncContext)

	p.SfPropertiesContext = NewEmptySfPropertiesContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SfPropertiesContext))

	return p
}

func (s *PropFuncContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PropFuncContext) IDENT() antlr.TerminalNode {
	return s.GetToken(FaultParserIDENT, 0)
}

func (s *PropFuncContext) COLON() antlr.TerminalNode {
	return s.GetToken(FaultParserCOLON, 0)
}

func (s *PropFuncContext) FunctionLit() IFunctionLitContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionLitContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionLitContext)
}

func (s *PropFuncContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaultParserListener); ok {
		listenerT.EnterPropFunc(s)
	}
}

func (s *PropFuncContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaultParserListener); ok {
		listenerT.ExitPropFunc(s)
	}
}

func (s *PropFuncContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FaultParserVisitor:
		return t.VisitPropFunc(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FaultParser) SfProperties() (localctx ISfPropertiesContext) {
	this := p
	_ = this

	localctx = NewSfPropertiesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, FaultParserRULE_sfProperties)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(325)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 24, p.GetParserRuleContext()) {
	case 1:
		localctx = NewPropFuncContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(321)
			p.Match(FaultParserIDENT)
		}
		{
			p.SetState(322)
			p.Match(FaultParserCOLON)
		}
		{
			p.SetState(323)
			p.FunctionLit()
		}

	case 2:
		localctx = NewSfMiscContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(324)
			p.StructProperties()
		}

	}

	return localctx
}

// IComPropertiesContext is an interface to support dynamic dispatch.
type IComPropertiesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsComPropertiesContext differentiates from other interfaces.
	IsComPropertiesContext()
}

type ComPropertiesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyComPropertiesContext() *ComPropertiesContext {
	var p = new(ComPropertiesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FaultParserRULE_comProperties
	return p
}

func (*ComPropertiesContext) IsComPropertiesContext() {}

func NewComPropertiesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ComPropertiesContext {
	var p = new(ComPropertiesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FaultParserRULE_comProperties

	return p
}

func (s *ComPropertiesContext) GetParser() antlr.Parser { return s.parser }

func (s *ComPropertiesContext) CopyFrom(ctx *ComPropertiesContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ComPropertiesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ComPropertiesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type StateFuncContext struct {
	*ComPropertiesContext
}

func NewStateFuncContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StateFuncContext {
	var p = new(StateFuncContext)

	p.ComPropertiesContext = NewEmptyComPropertiesContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ComPropertiesContext))

	return p
}

func (s *StateFuncContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StateFuncContext) IDENT() antlr.TerminalNode {
	return s.GetToken(FaultParserIDENT, 0)
}

func (s *StateFuncContext) COLON() antlr.TerminalNode {
	return s.GetToken(FaultParserCOLON, 0)
}

func (s *StateFuncContext) StateLit() IStateLitContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStateLitContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStateLitContext)
}

func (s *StateFuncContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaultParserListener); ok {
		listenerT.EnterStateFunc(s)
	}
}

func (s *StateFuncContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaultParserListener); ok {
		listenerT.ExitStateFunc(s)
	}
}

func (s *StateFuncContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FaultParserVisitor:
		return t.VisitStateFunc(s)

	default:
		return t.VisitChildren(s)
	}
}

type CompMiscContext struct {
	*ComPropertiesContext
}

func NewCompMiscContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CompMiscContext {
	var p = new(CompMiscContext)

	p.ComPropertiesContext = NewEmptyComPropertiesContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ComPropertiesContext))

	return p
}

func (s *CompMiscContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CompMiscContext) StructProperties() IStructPropertiesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStructPropertiesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStructPropertiesContext)
}

func (s *CompMiscContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaultParserListener); ok {
		listenerT.EnterCompMisc(s)
	}
}

func (s *CompMiscContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaultParserListener); ok {
		listenerT.ExitCompMisc(s)
	}
}

func (s *CompMiscContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FaultParserVisitor:
		return t.VisitCompMisc(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FaultParser) ComProperties() (localctx IComPropertiesContext) {
	this := p
	_ = this

	localctx = NewComPropertiesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, FaultParserRULE_comProperties)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(331)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 25, p.GetParserRuleContext()) {
	case 1:
		localctx = NewStateFuncContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(327)
			p.Match(FaultParserIDENT)
		}
		{
			p.SetState(328)
			p.Match(FaultParserCOLON)
		}
		{
			p.SetState(329)
			p.StateLit()
		}

	case 2:
		localctx = NewCompMiscContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(330)
			p.StructProperties()
		}

	}

	return localctx
}

// IStructPropertiesContext is an interface to support dynamic dispatch.
type IStructPropertiesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStructPropertiesContext differentiates from other interfaces.
	IsStructPropertiesContext()
}

type StructPropertiesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStructPropertiesContext() *StructPropertiesContext {
	var p = new(StructPropertiesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FaultParserRULE_structProperties
	return p
}

func (*StructPropertiesContext) IsStructPropertiesContext() {}

func NewStructPropertiesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StructPropertiesContext {
	var p = new(StructPropertiesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FaultParserRULE_structProperties

	return p
}

func (s *StructPropertiesContext) GetParser() antlr.Parser { return s.parser }

func (s *StructPropertiesContext) CopyFrom(ctx *StructPropertiesContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *StructPropertiesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructPropertiesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type PropSolvableContext struct {
	*StructPropertiesContext
}

func NewPropSolvableContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PropSolvableContext {
	var p = new(PropSolvableContext)

	p.StructPropertiesContext = NewEmptyStructPropertiesContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StructPropertiesContext))

	return p
}

func (s *PropSolvableContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PropSolvableContext) IDENT() antlr.TerminalNode {
	return s.GetToken(FaultParserIDENT, 0)
}

func (s *PropSolvableContext) COLON() antlr.TerminalNode {
	return s.GetToken(FaultParserCOLON, 0)
}

func (s *PropSolvableContext) Solvable() ISolvableContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISolvableContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISolvableContext)
}

func (s *PropSolvableContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaultParserListener); ok {
		listenerT.EnterPropSolvable(s)
	}
}

func (s *PropSolvableContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaultParserListener); ok {
		listenerT.ExitPropSolvable(s)
	}
}

func (s *PropSolvableContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FaultParserVisitor:
		return t.VisitPropSolvable(s)

	default:
		return t.VisitChildren(s)
	}
}

type PropBoolContext struct {
	*StructPropertiesContext
}

func NewPropBoolContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PropBoolContext {
	var p = new(PropBoolContext)

	p.StructPropertiesContext = NewEmptyStructPropertiesContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StructPropertiesContext))

	return p
}

func (s *PropBoolContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PropBoolContext) IDENT() antlr.TerminalNode {
	return s.GetToken(FaultParserIDENT, 0)
}

func (s *PropBoolContext) COLON() antlr.TerminalNode {
	return s.GetToken(FaultParserCOLON, 0)
}

func (s *PropBoolContext) Bool_() IBool_Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBool_Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBool_Context)
}

func (s *PropBoolContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaultParserListener); ok {
		listenerT.EnterPropBool(s)
	}
}

func (s *PropBoolContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaultParserListener); ok {
		listenerT.ExitPropBool(s)
	}
}

func (s *PropBoolContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FaultParserVisitor:
		return t.VisitPropBool(s)

	default:
		return t.VisitChildren(s)
	}
}

type PropIntContext struct {
	*StructPropertiesContext
}

func NewPropIntContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PropIntContext {
	var p = new(PropIntContext)

	p.StructPropertiesContext = NewEmptyStructPropertiesContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StructPropertiesContext))

	return p
}

func (s *PropIntContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PropIntContext) IDENT() antlr.TerminalNode {
	return s.GetToken(FaultParserIDENT, 0)
}

func (s *PropIntContext) COLON() antlr.TerminalNode {
	return s.GetToken(FaultParserCOLON, 0)
}

func (s *PropIntContext) Numeric() INumericContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INumericContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INumericContext)
}

func (s *PropIntContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaultParserListener); ok {
		listenerT.EnterPropInt(s)
	}
}

func (s *PropIntContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaultParserListener); ok {
		listenerT.ExitPropInt(s)
	}
}

func (s *PropIntContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FaultParserVisitor:
		return t.VisitPropInt(s)

	default:
		return t.VisitChildren(s)
	}
}

type PropStringContext struct {
	*StructPropertiesContext
}

func NewPropStringContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PropStringContext {
	var p = new(PropStringContext)

	p.StructPropertiesContext = NewEmptyStructPropertiesContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StructPropertiesContext))

	return p
}

func (s *PropStringContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PropStringContext) IDENT() antlr.TerminalNode {
	return s.GetToken(FaultParserIDENT, 0)
}

func (s *PropStringContext) COLON() antlr.TerminalNode {
	return s.GetToken(FaultParserCOLON, 0)
}

func (s *PropStringContext) String_() IString_Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IString_Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IString_Context)
}

func (s *PropStringContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaultParserListener); ok {
		listenerT.EnterPropString(s)
	}
}

func (s *PropStringContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaultParserListener); ok {
		listenerT.ExitPropString(s)
	}
}

func (s *PropStringContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FaultParserVisitor:
		return t.VisitPropString(s)

	default:
		return t.VisitChildren(s)
	}
}

type PropVarContext struct {
	*StructPropertiesContext
}

func NewPropVarContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PropVarContext {
	var p = new(PropVarContext)

	p.StructPropertiesContext = NewEmptyStructPropertiesContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StructPropertiesContext))

	return p
}

func (s *PropVarContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PropVarContext) IDENT() antlr.TerminalNode {
	return s.GetToken(FaultParserIDENT, 0)
}

func (s *PropVarContext) COLON() antlr.TerminalNode {
	return s.GetToken(FaultParserCOLON, 0)
}

func (s *PropVarContext) OperandName() IOperandNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOperandNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOperandNameContext)
}

func (s *PropVarContext) Prefix() IPrefixContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrefixContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrefixContext)
}

func (s *PropVarContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaultParserListener); ok {
		listenerT.EnterPropVar(s)
	}
}

func (s *PropVarContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaultParserListener); ok {
		listenerT.ExitPropVar(s)
	}
}

func (s *PropVarContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FaultParserVisitor:
		return t.VisitPropVar(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FaultParser) StructProperties() (localctx IStructPropertiesContext) {
	this := p
	_ = this

	localctx = NewStructPropertiesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, FaultParserRULE_structProperties)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(352)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 26, p.GetParserRuleContext()) {
	case 1:
		localctx = NewPropIntContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(333)
			p.Match(FaultParserIDENT)
		}
		{
			p.SetState(334)
			p.Match(FaultParserCOLON)
		}
		{
			p.SetState(335)
			p.Numeric()
		}

	case 2:
		localctx = NewPropStringContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(336)
			p.Match(FaultParserIDENT)
		}
		{
			p.SetState(337)
			p.Match(FaultParserCOLON)
		}
		{
			p.SetState(338)
			p.String_()
		}

	case 3:
		localctx = NewPropBoolContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(339)
			p.Match(FaultParserIDENT)
		}
		{
			p.SetState(340)
			p.Match(FaultParserCOLON)
		}
		{
			p.SetState(341)
			p.Bool_()
		}

	case 4:
		localctx = NewPropVarContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(342)
			p.Match(FaultParserIDENT)
		}
		{
			p.SetState(343)
			p.Match(FaultParserCOLON)
		}
		{
			p.SetState(344)
			p.OperandName()
		}

	case 5:
		localctx = NewPropVarContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(345)
			p.Match(FaultParserIDENT)
		}
		{
			p.SetState(346)
			p.Match(FaultParserCOLON)
		}
		{
			p.SetState(347)
			p.Prefix()
		}

	case 6:
		localctx = NewPropSolvableContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(348)
			p.Match(FaultParserIDENT)
		}
		{
			p.SetState(349)
			p.Match(FaultParserCOLON)
		}
		{
			p.SetState(350)
			p.Solvable()
		}

	case 7:
		localctx = NewPropSolvableContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(351)
			p.Match(FaultParserIDENT)
		}

	}

	return localctx
}

// IInitDeclContext is an interface to support dynamic dispatch.
type IInitDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInitDeclContext differentiates from other interfaces.
	IsInitDeclContext()
}

type InitDeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInitDeclContext() *InitDeclContext {
	var p = new(InitDeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FaultParserRULE_initDecl
	return p
}

func (*InitDeclContext) IsInitDeclContext() {}

func NewInitDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InitDeclContext {
	var p = new(InitDeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FaultParserRULE_initDecl

	return p
}

func (s *InitDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *InitDeclContext) INIT() antlr.TerminalNode {
	return s.GetToken(FaultParserINIT, 0)
}

func (s *InitDeclContext) Operand() IOperandContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOperandContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOperandContext)
}

func (s *InitDeclContext) Eos() IEosContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEosContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEosContext)
}

func (s *InitDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InitDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InitDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaultParserListener); ok {
		listenerT.EnterInitDecl(s)
	}
}

func (s *InitDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaultParserListener); ok {
		listenerT.ExitInitDecl(s)
	}
}

func (s *InitDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FaultParserVisitor:
		return t.VisitInitDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FaultParser) InitDecl() (localctx IInitDeclContext) {
	this := p
	_ = this

	localctx = NewInitDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, FaultParserRULE_initDecl)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(354)
		p.Match(FaultParserINIT)
	}
	{
		p.SetState(355)
		p.Operand()
	}
	{
		p.SetState(356)
		p.Eos()
	}

	return localctx
}

// IBlockContext is an interface to support dynamic dispatch.
type IBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBlockContext differentiates from other interfaces.
	IsBlockContext()
}

type BlockContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockContext() *BlockContext {
	var p = new(BlockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FaultParserRULE_block
	return p
}

func (*BlockContext) IsBlockContext() {}

func NewBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockContext {
	var p = new(BlockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FaultParserRULE_block

	return p
}

func (s *BlockContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockContext) LCURLY() antlr.TerminalNode {
	return s.GetToken(FaultParserLCURLY, 0)
}

func (s *BlockContext) RCURLY() antlr.TerminalNode {
	return s.GetToken(FaultParserRCURLY, 0)
}

func (s *BlockContext) StatementList() IStatementListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementListContext)
}

func (s *BlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaultParserListener); ok {
		listenerT.EnterBlock(s)
	}
}

func (s *BlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaultParserListener); ok {
		listenerT.ExitBlock(s)
	}
}

func (s *BlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FaultParserVisitor:
		return t.VisitBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FaultParser) Block() (localctx IBlockContext) {
	this := p
	_ = this

	localctx = NewBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, FaultParserRULE_block)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(358)
		p.Match(FaultParserLCURLY)
	}
	p.SetState(360)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 27, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(359)
			p.StatementList()
		}

	}
	{
		p.SetState(362)
		p.Match(FaultParserRCURLY)
	}

	return localctx
}

// IStatementListContext is an interface to support dynamic dispatch.
type IStatementListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatementListContext differentiates from other interfaces.
	IsStatementListContext()
}

type StatementListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementListContext() *StatementListContext {
	var p = new(StatementListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FaultParserRULE_statementList
	return p
}

func (*StatementListContext) IsStatementListContext() {}

func NewStatementListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementListContext {
	var p = new(StatementListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FaultParserRULE_statementList

	return p
}

func (s *StatementListContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementListContext) AllStatement() []IStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementContext); ok {
			len++
		}
	}

	tst := make([]IStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementContext); ok {
			tst[i] = t.(IStatementContext)
			i++
		}
	}

	return tst
}

func (s *StatementListContext) Statement(i int) IStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *StatementListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaultParserListener); ok {
		listenerT.EnterStatementList(s)
	}
}

func (s *StatementListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaultParserListener); ok {
		listenerT.ExitStatementList(s)
	}
}

func (s *StatementListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FaultParserVisitor:
		return t.VisitStatementList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FaultParser) StatementList() (localctx IStatementListContext) {
	this := p
	_ = this

	localctx = NewStatementListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, FaultParserRULE_statementList)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(365)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(364)
				p.Statement()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(367)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 28, p.GetParserRuleContext())
	}

	return localctx
}

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FaultParserRULE_statement
	return p
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FaultParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) ConstDecl() IConstDeclContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConstDeclContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConstDeclContext)
}

func (s *StatementContext) InitDecl() IInitDeclContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInitDeclContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInitDeclContext)
}

func (s *StatementContext) SimpleStmt() ISimpleStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISimpleStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISimpleStmtContext)
}

func (s *StatementContext) Eos() IEosContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEosContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEosContext)
}

func (s *StatementContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *StatementContext) IfStmt() IIfStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfStmtContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaultParserListener); ok {
		listenerT.EnterStatement(s)
	}
}

func (s *StatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaultParserListener); ok {
		listenerT.ExitStatement(s)
	}
}

func (s *StatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FaultParserVisitor:
		return t.VisitStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FaultParser) Statement() (localctx IStatementContext) {
	this := p
	_ = this

	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, FaultParserRULE_statement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(376)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 29, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(369)
			p.ConstDecl()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(370)
			p.InitDecl()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(371)
			p.SimpleStmt()
		}
		{
			p.SetState(372)
			p.Eos()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(374)
			p.Block()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(375)
			p.IfStmt()
		}

	}

	return localctx
}

// ISimpleStmtContext is an interface to support dynamic dispatch.
type ISimpleStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSimpleStmtContext differentiates from other interfaces.
	IsSimpleStmtContext()
}

type SimpleStmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySimpleStmtContext() *SimpleStmtContext {
	var p = new(SimpleStmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FaultParserRULE_simpleStmt
	return p
}

func (*SimpleStmtContext) IsSimpleStmtContext() {}

func NewSimpleStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SimpleStmtContext {
	var p = new(SimpleStmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FaultParserRULE_simpleStmt

	return p
}

func (s *SimpleStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *SimpleStmtContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *SimpleStmtContext) IncDecStmt() IIncDecStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIncDecStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIncDecStmtContext)
}

func (s *SimpleStmtContext) Assignment() IAssignmentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignmentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignmentContext)
}

func (s *SimpleStmtContext) EmptyStmt() IEmptyStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEmptyStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEmptyStmtContext)
}

func (s *SimpleStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SimpleStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SimpleStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaultParserListener); ok {
		listenerT.EnterSimpleStmt(s)
	}
}

func (s *SimpleStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaultParserListener); ok {
		listenerT.ExitSimpleStmt(s)
	}
}

func (s *SimpleStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FaultParserVisitor:
		return t.VisitSimpleStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FaultParser) SimpleStmt() (localctx ISimpleStmtContext) {
	this := p
	_ = this

	localctx = NewSimpleStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, FaultParserRULE_simpleStmt)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(382)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 30, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(378)
			p.expression(0)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(379)
			p.IncDecStmt()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(380)
			p.Assignment()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(381)
			p.EmptyStmt()
		}

	}

	return localctx
}

// IIncDecStmtContext is an interface to support dynamic dispatch.
type IIncDecStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIncDecStmtContext differentiates from other interfaces.
	IsIncDecStmtContext()
}

type IncDecStmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIncDecStmtContext() *IncDecStmtContext {
	var p = new(IncDecStmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FaultParserRULE_incDecStmt
	return p
}

func (*IncDecStmtContext) IsIncDecStmtContext() {}

func NewIncDecStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IncDecStmtContext {
	var p = new(IncDecStmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FaultParserRULE_incDecStmt

	return p
}

func (s *IncDecStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *IncDecStmtContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *IncDecStmtContext) PLUS_PLUS() antlr.TerminalNode {
	return s.GetToken(FaultParserPLUS_PLUS, 0)
}

func (s *IncDecStmtContext) MINUS_MINUS() antlr.TerminalNode {
	return s.GetToken(FaultParserMINUS_MINUS, 0)
}

func (s *IncDecStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IncDecStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IncDecStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaultParserListener); ok {
		listenerT.EnterIncDecStmt(s)
	}
}

func (s *IncDecStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaultParserListener); ok {
		listenerT.ExitIncDecStmt(s)
	}
}

func (s *IncDecStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FaultParserVisitor:
		return t.VisitIncDecStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FaultParser) IncDecStmt() (localctx IIncDecStmtContext) {
	this := p
	_ = this

	localctx = NewIncDecStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, FaultParserRULE_incDecStmt)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(384)
		p.expression(0)
	}
	{
		p.SetState(385)
		_la = p.GetTokenStream().LA(1)

		if !(_la == FaultParserPLUS_PLUS || _la == FaultParserMINUS_MINUS) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IStateChangeContext is an interface to support dynamic dispatch.
type IStateChangeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStateChangeContext differentiates from other interfaces.
	IsStateChangeContext()
}

type StateChangeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStateChangeContext() *StateChangeContext {
	var p = new(StateChangeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FaultParserRULE_stateChange
	return p
}

func (*StateChangeContext) IsStateChangeContext() {}

func NewStateChangeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StateChangeContext {
	var p = new(StateChangeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FaultParserRULE_stateChange

	return p
}

func (s *StateChangeContext) GetParser() antlr.Parser { return s.parser }

func (s *StateChangeContext) CopyFrom(ctx *StateChangeContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *StateChangeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StateChangeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type BuiltinsContext struct {
	*StateChangeContext
}

func NewBuiltinsContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BuiltinsContext {
	var p = new(BuiltinsContext)

	p.StateChangeContext = NewEmptyStateChangeContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StateChangeContext))

	return p
}

func (s *BuiltinsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BuiltinsContext) ADVANCE() antlr.TerminalNode {
	return s.GetToken(FaultParserADVANCE, 0)
}

func (s *BuiltinsContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(FaultParserLPAREN, 0)
}

func (s *BuiltinsContext) ParamCall() IParamCallContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParamCallContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParamCallContext)
}

func (s *BuiltinsContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(FaultParserRPAREN, 0)
}

func (s *BuiltinsContext) STAY() antlr.TerminalNode {
	return s.GetToken(FaultParserSTAY, 0)
}

func (s *BuiltinsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaultParserListener); ok {
		listenerT.EnterBuiltins(s)
	}
}

func (s *BuiltinsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaultParserListener); ok {
		listenerT.ExitBuiltins(s)
	}
}

func (s *BuiltinsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FaultParserVisitor:
		return t.VisitBuiltins(s)

	default:
		return t.VisitChildren(s)
	}
}

type BuiltinInfixContext struct {
	*StateChangeContext
}

func NewBuiltinInfixContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BuiltinInfixContext {
	var p = new(BuiltinInfixContext)

	p.StateChangeContext = NewEmptyStateChangeContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StateChangeContext))

	return p
}

func (s *BuiltinInfixContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BuiltinInfixContext) AllStateChange() []IStateChangeContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStateChangeContext); ok {
			len++
		}
	}

	tst := make([]IStateChangeContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStateChangeContext); ok {
			tst[i] = t.(IStateChangeContext)
			i++
		}
	}

	return tst
}

func (s *BuiltinInfixContext) StateChange(i int) IStateChangeContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStateChangeContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStateChangeContext)
}

func (s *BuiltinInfixContext) AND() antlr.TerminalNode {
	return s.GetToken(FaultParserAND, 0)
}

func (s *BuiltinInfixContext) OR() antlr.TerminalNode {
	return s.GetToken(FaultParserOR, 0)
}

func (s *BuiltinInfixContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaultParserListener); ok {
		listenerT.EnterBuiltinInfix(s)
	}
}

func (s *BuiltinInfixContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaultParserListener); ok {
		listenerT.ExitBuiltinInfix(s)
	}
}

func (s *BuiltinInfixContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FaultParserVisitor:
		return t.VisitBuiltinInfix(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FaultParser) StateChange() (localctx IStateChangeContext) {
	return p.stateChange(0)
}

func (p *FaultParser) stateChange(_p int) (localctx IStateChangeContext) {
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewStateChangeContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IStateChangeContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 60
	p.EnterRecursionRule(localctx, 60, FaultParserRULE_stateChange, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(396)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case FaultParserADVANCE:
		localctx = NewBuiltinsContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(388)
			p.Match(FaultParserADVANCE)
		}
		{
			p.SetState(389)
			p.Match(FaultParserLPAREN)
		}
		{
			p.SetState(390)
			p.ParamCall()
		}
		{
			p.SetState(391)
			p.Match(FaultParserRPAREN)
		}

	case FaultParserSTAY:
		localctx = NewBuiltinsContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(393)
			p.Match(FaultParserSTAY)
		}
		{
			p.SetState(394)
			p.Match(FaultParserLPAREN)
		}
		{
			p.SetState(395)
			p.Match(FaultParserRPAREN)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(406)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 33, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(404)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 32, p.GetParserRuleContext()) {
			case 1:
				localctx = NewBuiltinInfixContext(p, NewStateChangeContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, FaultParserRULE_stateChange)
				p.SetState(398)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				{
					p.SetState(399)
					p.Match(FaultParserAND)
				}
				{
					p.SetState(400)
					p.stateChange(3)
				}

			case 2:
				localctx = NewBuiltinInfixContext(p, NewStateChangeContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, FaultParserRULE_stateChange)
				p.SetState(401)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
				}
				{
					p.SetState(402)
					p.Match(FaultParserOR)
				}
				{
					p.SetState(403)
					p.stateChange(2)
				}

			}

		}
		p.SetState(408)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 33, p.GetParserRuleContext())
	}

	return localctx
}

// IAccessHistoryContext is an interface to support dynamic dispatch.
type IAccessHistoryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAccessHistoryContext differentiates from other interfaces.
	IsAccessHistoryContext()
}

type AccessHistoryContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAccessHistoryContext() *AccessHistoryContext {
	var p = new(AccessHistoryContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FaultParserRULE_accessHistory
	return p
}

func (*AccessHistoryContext) IsAccessHistoryContext() {}

func NewAccessHistoryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AccessHistoryContext {
	var p = new(AccessHistoryContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FaultParserRULE_accessHistory

	return p
}

func (s *AccessHistoryContext) GetParser() antlr.Parser { return s.parser }

func (s *AccessHistoryContext) OperandName() IOperandNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOperandNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOperandNameContext)
}

func (s *AccessHistoryContext) AllLBRACE() []antlr.TerminalNode {
	return s.GetTokens(FaultParserLBRACE)
}

func (s *AccessHistoryContext) LBRACE(i int) antlr.TerminalNode {
	return s.GetToken(FaultParserLBRACE, i)
}

func (s *AccessHistoryContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *AccessHistoryContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AccessHistoryContext) AllRBRACE() []antlr.TerminalNode {
	return s.GetTokens(FaultParserRBRACE)
}

func (s *AccessHistoryContext) RBRACE(i int) antlr.TerminalNode {
	return s.GetToken(FaultParserRBRACE, i)
}

func (s *AccessHistoryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AccessHistoryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AccessHistoryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaultParserListener); ok {
		listenerT.EnterAccessHistory(s)
	}
}

func (s *AccessHistoryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaultParserListener); ok {
		listenerT.ExitAccessHistory(s)
	}
}

func (s *AccessHistoryContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FaultParserVisitor:
		return t.VisitAccessHistory(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FaultParser) AccessHistory() (localctx IAccessHistoryContext) {
	this := p
	_ = this

	localctx = NewAccessHistoryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, FaultParserRULE_accessHistory)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(409)
		p.OperandName()
	}
	p.SetState(414)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(410)
				p.Match(FaultParserLBRACE)
			}
			{
				p.SetState(411)
				p.expression(0)
			}
			{
				p.SetState(412)
				p.Match(FaultParserRBRACE)
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(416)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 34, p.GetParserRuleContext())
	}

	return localctx
}

// IAssertionContext is an interface to support dynamic dispatch.
type IAssertionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAssertionContext differentiates from other interfaces.
	IsAssertionContext()
}

type AssertionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssertionContext() *AssertionContext {
	var p = new(AssertionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FaultParserRULE_assertion
	return p
}

func (*AssertionContext) IsAssertionContext() {}

func NewAssertionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssertionContext {
	var p = new(AssertionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FaultParserRULE_assertion

	return p
}

func (s *AssertionContext) GetParser() antlr.Parser { return s.parser }

func (s *AssertionContext) ASSERT() antlr.TerminalNode {
	return s.GetToken(FaultParserASSERT, 0)
}

func (s *AssertionContext) Invariant() IInvariantContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInvariantContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInvariantContext)
}

func (s *AssertionContext) Eos() IEosContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEosContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEosContext)
}

func (s *AssertionContext) Temporal() ITemporalContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITemporalContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITemporalContext)
}

func (s *AssertionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssertionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssertionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaultParserListener); ok {
		listenerT.EnterAssertion(s)
	}
}

func (s *AssertionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaultParserListener); ok {
		listenerT.ExitAssertion(s)
	}
}

func (s *AssertionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FaultParserVisitor:
		return t.VisitAssertion(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FaultParser) Assertion() (localctx IAssertionContext) {
	this := p
	_ = this

	localctx = NewAssertionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, FaultParserRULE_assertion)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(418)
		p.Match(FaultParserASSERT)
	}
	{
		p.SetState(419)
		p.Invariant()
	}
	p.SetState(421)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&130023424) != 0 {
		{
			p.SetState(420)
			p.Temporal()
		}

	}
	{
		p.SetState(423)
		p.Eos()
	}

	return localctx
}

// IAssumptionContext is an interface to support dynamic dispatch.
type IAssumptionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAssumptionContext differentiates from other interfaces.
	IsAssumptionContext()
}

type AssumptionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssumptionContext() *AssumptionContext {
	var p = new(AssumptionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FaultParserRULE_assumption
	return p
}

func (*AssumptionContext) IsAssumptionContext() {}

func NewAssumptionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssumptionContext {
	var p = new(AssumptionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FaultParserRULE_assumption

	return p
}

func (s *AssumptionContext) GetParser() antlr.Parser { return s.parser }

func (s *AssumptionContext) ASSUME() antlr.TerminalNode {
	return s.GetToken(FaultParserASSUME, 0)
}

func (s *AssumptionContext) Invariant() IInvariantContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInvariantContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInvariantContext)
}

func (s *AssumptionContext) Eos() IEosContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEosContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEosContext)
}

func (s *AssumptionContext) Temporal() ITemporalContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITemporalContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITemporalContext)
}

func (s *AssumptionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssumptionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssumptionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaultParserListener); ok {
		listenerT.EnterAssumption(s)
	}
}

func (s *AssumptionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaultParserListener); ok {
		listenerT.ExitAssumption(s)
	}
}

func (s *AssumptionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FaultParserVisitor:
		return t.VisitAssumption(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FaultParser) Assumption() (localctx IAssumptionContext) {
	this := p
	_ = this

	localctx = NewAssumptionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, FaultParserRULE_assumption)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(425)
		p.Match(FaultParserASSUME)
	}
	{
		p.SetState(426)
		p.Invariant()
	}
	p.SetState(428)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&130023424) != 0 {
		{
			p.SetState(427)
			p.Temporal()
		}

	}
	{
		p.SetState(430)
		p.Eos()
	}

	return localctx
}

// ITemporalContext is an interface to support dynamic dispatch.
type ITemporalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTemporalContext differentiates from other interfaces.
	IsTemporalContext()
}

type TemporalContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTemporalContext() *TemporalContext {
	var p = new(TemporalContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FaultParserRULE_temporal
	return p
}

func (*TemporalContext) IsTemporalContext() {}

func NewTemporalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TemporalContext {
	var p = new(TemporalContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FaultParserRULE_temporal

	return p
}

func (s *TemporalContext) GetParser() antlr.Parser { return s.parser }

func (s *TemporalContext) EVENTUALLY() antlr.TerminalNode {
	return s.GetToken(FaultParserEVENTUALLY, 0)
}

func (s *TemporalContext) ALWAYS() antlr.TerminalNode {
	return s.GetToken(FaultParserALWAYS, 0)
}

func (s *TemporalContext) EVENTUALLYALWAYS() antlr.TerminalNode {
	return s.GetToken(FaultParserEVENTUALLYALWAYS, 0)
}

func (s *TemporalContext) Integer() IIntegerContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIntegerContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIntegerContext)
}

func (s *TemporalContext) NMT() antlr.TerminalNode {
	return s.GetToken(FaultParserNMT, 0)
}

func (s *TemporalContext) NFT() antlr.TerminalNode {
	return s.GetToken(FaultParserNFT, 0)
}

func (s *TemporalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TemporalContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TemporalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaultParserListener); ok {
		listenerT.EnterTemporal(s)
	}
}

func (s *TemporalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaultParserListener); ok {
		listenerT.ExitTemporal(s)
	}
}

func (s *TemporalContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FaultParserVisitor:
		return t.VisitTemporal(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FaultParser) Temporal() (localctx ITemporalContext) {
	this := p
	_ = this

	localctx = NewTemporalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, FaultParserRULE_temporal)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(435)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case FaultParserEVENTUALLY, FaultParserEVENTUALLYALWAYS, FaultParserALWAYS:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(432)
			_la = p.GetTokenStream().LA(1)

			if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&29360128) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	case FaultParserNMT, FaultParserNFT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(433)
			_la = p.GetTokenStream().LA(1)

			if !(_la == FaultParserNMT || _la == FaultParserNFT) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(434)
			p.Integer()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IInvariantContext is an interface to support dynamic dispatch.
type IInvariantContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInvariantContext differentiates from other interfaces.
	IsInvariantContext()
}

type InvariantContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInvariantContext() *InvariantContext {
	var p = new(InvariantContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FaultParserRULE_invariant
	return p
}

func (*InvariantContext) IsInvariantContext() {}

func NewInvariantContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InvariantContext {
	var p = new(InvariantContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FaultParserRULE_invariant

	return p
}

func (s *InvariantContext) GetParser() antlr.Parser { return s.parser }

func (s *InvariantContext) CopyFrom(ctx *InvariantContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *InvariantContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InvariantContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type InvarContext struct {
	*InvariantContext
}

func NewInvarContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *InvarContext {
	var p = new(InvarContext)

	p.InvariantContext = NewEmptyInvariantContext()
	p.parser = parser
	p.CopyFrom(ctx.(*InvariantContext))

	return p
}

func (s *InvarContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InvarContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *InvarContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaultParserListener); ok {
		listenerT.EnterInvar(s)
	}
}

func (s *InvarContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaultParserListener); ok {
		listenerT.ExitInvar(s)
	}
}

func (s *InvarContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FaultParserVisitor:
		return t.VisitInvar(s)

	default:
		return t.VisitChildren(s)
	}
}

type StageInvariantContext struct {
	*InvariantContext
}

func NewStageInvariantContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StageInvariantContext {
	var p = new(StageInvariantContext)

	p.InvariantContext = NewEmptyInvariantContext()
	p.parser = parser
	p.CopyFrom(ctx.(*InvariantContext))

	return p
}

func (s *StageInvariantContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StageInvariantContext) WHEN() antlr.TerminalNode {
	return s.GetToken(FaultParserWHEN, 0)
}

func (s *StageInvariantContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *StageInvariantContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *StageInvariantContext) THEN() antlr.TerminalNode {
	return s.GetToken(FaultParserTHEN, 0)
}

func (s *StageInvariantContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaultParserListener); ok {
		listenerT.EnterStageInvariant(s)
	}
}

func (s *StageInvariantContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaultParserListener); ok {
		listenerT.ExitStageInvariant(s)
	}
}

func (s *StageInvariantContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FaultParserVisitor:
		return t.VisitStageInvariant(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FaultParser) Invariant() (localctx IInvariantContext) {
	this := p
	_ = this

	localctx = NewInvariantContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, FaultParserRULE_invariant)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(443)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 38, p.GetParserRuleContext()) {
	case 1:
		localctx = NewInvarContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(437)
			p.expression(0)
		}

	case 2:
		localctx = NewStageInvariantContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(438)
			p.Match(FaultParserWHEN)
		}
		{
			p.SetState(439)
			p.expression(0)
		}
		{
			p.SetState(440)
			p.Match(FaultParserTHEN)
		}
		{
			p.SetState(441)
			p.expression(0)
		}

	}

	return localctx
}

// IAssignmentContext is an interface to support dynamic dispatch.
type IAssignmentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAssignmentContext differentiates from other interfaces.
	IsAssignmentContext()
}

type AssignmentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignmentContext() *AssignmentContext {
	var p = new(AssignmentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FaultParserRULE_assignment
	return p
}

func (*AssignmentContext) IsAssignmentContext() {}

func NewAssignmentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignmentContext {
	var p = new(AssignmentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FaultParserRULE_assignment

	return p
}

func (s *AssignmentContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignmentContext) CopyFrom(ctx *AssignmentContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *AssignmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type MiscAssignContext struct {
	*AssignmentContext
}

func NewMiscAssignContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MiscAssignContext {
	var p = new(MiscAssignContext)

	p.AssignmentContext = NewEmptyAssignmentContext()
	p.parser = parser
	p.CopyFrom(ctx.(*AssignmentContext))

	return p
}

func (s *MiscAssignContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MiscAssignContext) AllExpressionList() []IExpressionListContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionListContext); ok {
			len++
		}
	}

	tst := make([]IExpressionListContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionListContext); ok {
			tst[i] = t.(IExpressionListContext)
			i++
		}
	}

	return tst
}

func (s *MiscAssignContext) ExpressionList(i int) IExpressionListContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionListContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionListContext)
}

func (s *MiscAssignContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(FaultParserASSIGN, 0)
}

func (s *MiscAssignContext) PLUS() antlr.TerminalNode {
	return s.GetToken(FaultParserPLUS, 0)
}

func (s *MiscAssignContext) MINUS() antlr.TerminalNode {
	return s.GetToken(FaultParserMINUS, 0)
}

func (s *MiscAssignContext) CARET() antlr.TerminalNode {
	return s.GetToken(FaultParserCARET, 0)
}

func (s *MiscAssignContext) MULTI() antlr.TerminalNode {
	return s.GetToken(FaultParserMULTI, 0)
}

func (s *MiscAssignContext) DIV() antlr.TerminalNode {
	return s.GetToken(FaultParserDIV, 0)
}

func (s *MiscAssignContext) MOD() antlr.TerminalNode {
	return s.GetToken(FaultParserMOD, 0)
}

func (s *MiscAssignContext) LSHIFT() antlr.TerminalNode {
	return s.GetToken(FaultParserLSHIFT, 0)
}

func (s *MiscAssignContext) RSHIFT() antlr.TerminalNode {
	return s.GetToken(FaultParserRSHIFT, 0)
}

func (s *MiscAssignContext) AMPERSAND() antlr.TerminalNode {
	return s.GetToken(FaultParserAMPERSAND, 0)
}

func (s *MiscAssignContext) BIT_CLEAR() antlr.TerminalNode {
	return s.GetToken(FaultParserBIT_CLEAR, 0)
}

func (s *MiscAssignContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaultParserListener); ok {
		listenerT.EnterMiscAssign(s)
	}
}

func (s *MiscAssignContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaultParserListener); ok {
		listenerT.ExitMiscAssign(s)
	}
}

func (s *MiscAssignContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FaultParserVisitor:
		return t.VisitMiscAssign(s)

	default:
		return t.VisitChildren(s)
	}
}

type FaultAssignContext struct {
	*AssignmentContext
}

func NewFaultAssignContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FaultAssignContext {
	var p = new(FaultAssignContext)

	p.AssignmentContext = NewEmptyAssignmentContext()
	p.parser = parser
	p.CopyFrom(ctx.(*AssignmentContext))

	return p
}

func (s *FaultAssignContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FaultAssignContext) AllExpressionList() []IExpressionListContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionListContext); ok {
			len++
		}
	}

	tst := make([]IExpressionListContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionListContext); ok {
			tst[i] = t.(IExpressionListContext)
			i++
		}
	}

	return tst
}

func (s *FaultAssignContext) ExpressionList(i int) IExpressionListContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionListContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionListContext)
}

func (s *FaultAssignContext) ASSIGN_FLOW1() antlr.TerminalNode {
	return s.GetToken(FaultParserASSIGN_FLOW1, 0)
}

func (s *FaultAssignContext) ASSIGN_FLOW2() antlr.TerminalNode {
	return s.GetToken(FaultParserASSIGN_FLOW2, 0)
}

func (s *FaultAssignContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaultParserListener); ok {
		listenerT.EnterFaultAssign(s)
	}
}

func (s *FaultAssignContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaultParserListener); ok {
		listenerT.ExitFaultAssign(s)
	}
}

func (s *FaultAssignContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FaultParserVisitor:
		return t.VisitFaultAssign(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FaultParser) Assignment() (localctx IAssignmentContext) {
	this := p
	_ = this

	localctx = NewAssignmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, FaultParserRULE_assignment)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(456)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 40, p.GetParserRuleContext()) {
	case 1:
		localctx = NewMiscAssignContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(445)
			p.ExpressionList()
		}
		p.SetState(447)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if (int64((_la-60)) & ^0x3f) == 0 && ((int64(1)<<(_la-60))&2078721) != 0 {
			{
				p.SetState(446)
				_la = p.GetTokenStream().LA(1)

				if !((int64((_la-60)) & ^0x3f) == 0 && ((int64(1)<<(_la-60))&2078721) != 0) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}

		}
		{
			p.SetState(449)
			p.Match(FaultParserASSIGN)
		}
		{
			p.SetState(450)
			p.ExpressionList()
		}

	case 2:
		localctx = NewFaultAssignContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(452)
			p.ExpressionList()
		}
		{
			p.SetState(453)
			_la = p.GetTokenStream().LA(1)

			if !(_la == FaultParserASSIGN_FLOW1 || _la == FaultParserASSIGN_FLOW2) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(454)
			p.ExpressionList()
		}

	}

	return localctx
}

// IEmptyStmtContext is an interface to support dynamic dispatch.
type IEmptyStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEmptyStmtContext differentiates from other interfaces.
	IsEmptyStmtContext()
}

type EmptyStmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEmptyStmtContext() *EmptyStmtContext {
	var p = new(EmptyStmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FaultParserRULE_emptyStmt
	return p
}

func (*EmptyStmtContext) IsEmptyStmtContext() {}

func NewEmptyStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EmptyStmtContext {
	var p = new(EmptyStmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FaultParserRULE_emptyStmt

	return p
}

func (s *EmptyStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *EmptyStmtContext) SEMI() antlr.TerminalNode {
	return s.GetToken(FaultParserSEMI, 0)
}

func (s *EmptyStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EmptyStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EmptyStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaultParserListener); ok {
		listenerT.EnterEmptyStmt(s)
	}
}

func (s *EmptyStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaultParserListener); ok {
		listenerT.ExitEmptyStmt(s)
	}
}

func (s *EmptyStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FaultParserVisitor:
		return t.VisitEmptyStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FaultParser) EmptyStmt() (localctx IEmptyStmtContext) {
	this := p
	_ = this

	localctx = NewEmptyStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, FaultParserRULE_emptyStmt)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(458)
		p.Match(FaultParserSEMI)
	}

	return localctx
}

// IIfStmtContext is an interface to support dynamic dispatch.
type IIfStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIfStmtContext differentiates from other interfaces.
	IsIfStmtContext()
}

type IfStmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIfStmtContext() *IfStmtContext {
	var p = new(IfStmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FaultParserRULE_ifStmt
	return p
}

func (*IfStmtContext) IsIfStmtContext() {}

func NewIfStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfStmtContext {
	var p = new(IfStmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FaultParserRULE_ifStmt

	return p
}

func (s *IfStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *IfStmtContext) IF() antlr.TerminalNode {
	return s.GetToken(FaultParserIF, 0)
}

func (s *IfStmtContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *IfStmtContext) AllBlock() []IBlockContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBlockContext); ok {
			len++
		}
	}

	tst := make([]IBlockContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBlockContext); ok {
			tst[i] = t.(IBlockContext)
			i++
		}
	}

	return tst
}

func (s *IfStmtContext) Block(i int) IBlockContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *IfStmtContext) SimpleStmt() ISimpleStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISimpleStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISimpleStmtContext)
}

func (s *IfStmtContext) SEMI() antlr.TerminalNode {
	return s.GetToken(FaultParserSEMI, 0)
}

func (s *IfStmtContext) ELSE() antlr.TerminalNode {
	return s.GetToken(FaultParserELSE, 0)
}

func (s *IfStmtContext) IfStmt() IIfStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfStmtContext)
}

func (s *IfStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IfStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaultParserListener); ok {
		listenerT.EnterIfStmt(s)
	}
}

func (s *IfStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaultParserListener); ok {
		listenerT.ExitIfStmt(s)
	}
}

func (s *IfStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FaultParserVisitor:
		return t.VisitIfStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FaultParser) IfStmt() (localctx IIfStmtContext) {
	this := p
	_ = this

	localctx = NewIfStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, FaultParserRULE_ifStmt)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(460)
		p.Match(FaultParserIF)
	}
	p.SetState(464)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 41, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(461)
			p.SimpleStmt()
		}
		{
			p.SetState(462)
			p.Match(FaultParserSEMI)
		}

	}
	{
		p.SetState(466)
		p.expression(0)
	}
	{
		p.SetState(467)
		p.Block()
	}
	p.SetState(473)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 43, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(468)
			p.Match(FaultParserELSE)
		}
		p.SetState(471)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case FaultParserIF:
			{
				p.SetState(469)
				p.IfStmt()
			}

		case FaultParserLCURLY:
			{
				p.SetState(470)
				p.Block()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

	}

	return localctx
}

// IIfStmtRunContext is an interface to support dynamic dispatch.
type IIfStmtRunContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIfStmtRunContext differentiates from other interfaces.
	IsIfStmtRunContext()
}

type IfStmtRunContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIfStmtRunContext() *IfStmtRunContext {
	var p = new(IfStmtRunContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FaultParserRULE_ifStmtRun
	return p
}

func (*IfStmtRunContext) IsIfStmtRunContext() {}

func NewIfStmtRunContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfStmtRunContext {
	var p = new(IfStmtRunContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FaultParserRULE_ifStmtRun

	return p
}

func (s *IfStmtRunContext) GetParser() antlr.Parser { return s.parser }

func (s *IfStmtRunContext) IF() antlr.TerminalNode {
	return s.GetToken(FaultParserIF, 0)
}

func (s *IfStmtRunContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *IfStmtRunContext) AllRunBlock() []IRunBlockContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IRunBlockContext); ok {
			len++
		}
	}

	tst := make([]IRunBlockContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IRunBlockContext); ok {
			tst[i] = t.(IRunBlockContext)
			i++
		}
	}

	return tst
}

func (s *IfStmtRunContext) RunBlock(i int) IRunBlockContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRunBlockContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRunBlockContext)
}

func (s *IfStmtRunContext) SimpleStmt() ISimpleStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISimpleStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISimpleStmtContext)
}

func (s *IfStmtRunContext) SEMI() antlr.TerminalNode {
	return s.GetToken(FaultParserSEMI, 0)
}

func (s *IfStmtRunContext) ELSE() antlr.TerminalNode {
	return s.GetToken(FaultParserELSE, 0)
}

func (s *IfStmtRunContext) IfStmtRun() IIfStmtRunContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfStmtRunContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfStmtRunContext)
}

func (s *IfStmtRunContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfStmtRunContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IfStmtRunContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaultParserListener); ok {
		listenerT.EnterIfStmtRun(s)
	}
}

func (s *IfStmtRunContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaultParserListener); ok {
		listenerT.ExitIfStmtRun(s)
	}
}

func (s *IfStmtRunContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FaultParserVisitor:
		return t.VisitIfStmtRun(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FaultParser) IfStmtRun() (localctx IIfStmtRunContext) {
	this := p
	_ = this

	localctx = NewIfStmtRunContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, FaultParserRULE_ifStmtRun)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(475)
		p.Match(FaultParserIF)
	}
	p.SetState(479)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 44, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(476)
			p.SimpleStmt()
		}
		{
			p.SetState(477)
			p.Match(FaultParserSEMI)
		}

	}
	{
		p.SetState(481)
		p.expression(0)
	}
	{
		p.SetState(482)
		p.RunBlock()
	}
	p.SetState(488)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 46, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(483)
			p.Match(FaultParserELSE)
		}
		p.SetState(486)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case FaultParserIF:
			{
				p.SetState(484)
				p.IfStmtRun()
			}

		case FaultParserLCURLY:
			{
				p.SetState(485)
				p.RunBlock()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

	}

	return localctx
}

// IIfStmtStateContext is an interface to support dynamic dispatch.
type IIfStmtStateContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIfStmtStateContext differentiates from other interfaces.
	IsIfStmtStateContext()
}

type IfStmtStateContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIfStmtStateContext() *IfStmtStateContext {
	var p = new(IfStmtStateContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FaultParserRULE_ifStmtState
	return p
}

func (*IfStmtStateContext) IsIfStmtStateContext() {}

func NewIfStmtStateContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfStmtStateContext {
	var p = new(IfStmtStateContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FaultParserRULE_ifStmtState

	return p
}

func (s *IfStmtStateContext) GetParser() antlr.Parser { return s.parser }

func (s *IfStmtStateContext) IF() antlr.TerminalNode {
	return s.GetToken(FaultParserIF, 0)
}

func (s *IfStmtStateContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *IfStmtStateContext) AllStateBlock() []IStateBlockContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStateBlockContext); ok {
			len++
		}
	}

	tst := make([]IStateBlockContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStateBlockContext); ok {
			tst[i] = t.(IStateBlockContext)
			i++
		}
	}

	return tst
}

func (s *IfStmtStateContext) StateBlock(i int) IStateBlockContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStateBlockContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStateBlockContext)
}

func (s *IfStmtStateContext) SimpleStmt() ISimpleStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISimpleStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISimpleStmtContext)
}

func (s *IfStmtStateContext) SEMI() antlr.TerminalNode {
	return s.GetToken(FaultParserSEMI, 0)
}

func (s *IfStmtStateContext) ELSE() antlr.TerminalNode {
	return s.GetToken(FaultParserELSE, 0)
}

func (s *IfStmtStateContext) IfStmtState() IIfStmtStateContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfStmtStateContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfStmtStateContext)
}

func (s *IfStmtStateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfStmtStateContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IfStmtStateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaultParserListener); ok {
		listenerT.EnterIfStmtState(s)
	}
}

func (s *IfStmtStateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaultParserListener); ok {
		listenerT.ExitIfStmtState(s)
	}
}

func (s *IfStmtStateContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FaultParserVisitor:
		return t.VisitIfStmtState(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FaultParser) IfStmtState() (localctx IIfStmtStateContext) {
	this := p
	_ = this

	localctx = NewIfStmtStateContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 80, FaultParserRULE_ifStmtState)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(490)
		p.Match(FaultParserIF)
	}
	p.SetState(494)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 47, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(491)
			p.SimpleStmt()
		}
		{
			p.SetState(492)
			p.Match(FaultParserSEMI)
		}

	}
	{
		p.SetState(496)
		p.expression(0)
	}
	{
		p.SetState(497)
		p.StateBlock()
	}
	p.SetState(503)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == FaultParserELSE {
		{
			p.SetState(498)
			p.Match(FaultParserELSE)
		}
		p.SetState(501)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case FaultParserIF:
			{
				p.SetState(499)
				p.IfStmtState()
			}

		case FaultParserLCURLY:
			{
				p.SetState(500)
				p.StateBlock()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

	}

	return localctx
}

// IForStmtContext is an interface to support dynamic dispatch.
type IForStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsForStmtContext differentiates from other interfaces.
	IsForStmtContext()
}

type ForStmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyForStmtContext() *ForStmtContext {
	var p = new(ForStmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FaultParserRULE_forStmt
	return p
}

func (*ForStmtContext) IsForStmtContext() {}

func NewForStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ForStmtContext {
	var p = new(ForStmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FaultParserRULE_forStmt

	return p
}

func (s *ForStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *ForStmtContext) FOR() antlr.TerminalNode {
	return s.GetToken(FaultParserFOR, 0)
}

func (s *ForStmtContext) Rounds() IRoundsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRoundsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRoundsContext)
}

func (s *ForStmtContext) RUN() antlr.TerminalNode {
	return s.GetToken(FaultParserRUN, 0)
}

func (s *ForStmtContext) RunBlock() IRunBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRunBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRunBlockContext)
}

func (s *ForStmtContext) Eos() IEosContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEosContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEosContext)
}

func (s *ForStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ForStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaultParserListener); ok {
		listenerT.EnterForStmt(s)
	}
}

func (s *ForStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaultParserListener); ok {
		listenerT.ExitForStmt(s)
	}
}

func (s *ForStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FaultParserVisitor:
		return t.VisitForStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FaultParser) ForStmt() (localctx IForStmtContext) {
	this := p
	_ = this

	localctx = NewForStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 82, FaultParserRULE_forStmt)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(505)
		p.Match(FaultParserFOR)
	}
	{
		p.SetState(506)
		p.Rounds()
	}
	{
		p.SetState(507)
		p.Match(FaultParserRUN)
	}
	{
		p.SetState(508)
		p.RunBlock()
	}
	p.SetState(510)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == FaultParserSEMI {
		{
			p.SetState(509)
			p.Eos()
		}

	}

	return localctx
}

// IRoundsContext is an interface to support dynamic dispatch.
type IRoundsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRoundsContext differentiates from other interfaces.
	IsRoundsContext()
}

type RoundsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRoundsContext() *RoundsContext {
	var p = new(RoundsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FaultParserRULE_rounds
	return p
}

func (*RoundsContext) IsRoundsContext() {}

func NewRoundsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RoundsContext {
	var p = new(RoundsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FaultParserRULE_rounds

	return p
}

func (s *RoundsContext) GetParser() antlr.Parser { return s.parser }

func (s *RoundsContext) Integer() IIntegerContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIntegerContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIntegerContext)
}

func (s *RoundsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RoundsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RoundsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaultParserListener); ok {
		listenerT.EnterRounds(s)
	}
}

func (s *RoundsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaultParserListener); ok {
		listenerT.ExitRounds(s)
	}
}

func (s *RoundsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FaultParserVisitor:
		return t.VisitRounds(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FaultParser) Rounds() (localctx IRoundsContext) {
	this := p
	_ = this

	localctx = NewRoundsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 84, FaultParserRULE_rounds)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(512)
		p.Integer()
	}

	return localctx
}

// IParamCallContext is an interface to support dynamic dispatch.
type IParamCallContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsParamCallContext differentiates from other interfaces.
	IsParamCallContext()
}

type ParamCallContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParamCallContext() *ParamCallContext {
	var p = new(ParamCallContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FaultParserRULE_paramCall
	return p
}

func (*ParamCallContext) IsParamCallContext() {}

func NewParamCallContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParamCallContext {
	var p = new(ParamCallContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FaultParserRULE_paramCall

	return p
}

func (s *ParamCallContext) GetParser() antlr.Parser { return s.parser }

func (s *ParamCallContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(FaultParserDOT)
}

func (s *ParamCallContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(FaultParserDOT, i)
}

func (s *ParamCallContext) AllIDENT() []antlr.TerminalNode {
	return s.GetTokens(FaultParserIDENT)
}

func (s *ParamCallContext) IDENT(i int) antlr.TerminalNode {
	return s.GetToken(FaultParserIDENT, i)
}

func (s *ParamCallContext) THIS() antlr.TerminalNode {
	return s.GetToken(FaultParserTHIS, 0)
}

func (s *ParamCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParamCallContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParamCallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaultParserListener); ok {
		listenerT.EnterParamCall(s)
	}
}

func (s *ParamCallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaultParserListener); ok {
		listenerT.ExitParamCall(s)
	}
}

func (s *ParamCallContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FaultParserVisitor:
		return t.VisitParamCall(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FaultParser) ParamCall() (localctx IParamCallContext) {
	this := p
	_ = this

	localctx = NewParamCallContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 86, FaultParserRULE_paramCall)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(514)
		_la = p.GetTokenStream().LA(1)

		if !(_la == FaultParserTHIS || _la == FaultParserIDENT) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(515)
		p.Match(FaultParserDOT)
	}
	{
		p.SetState(516)
		p.Match(FaultParserIDENT)
	}
	p.SetState(521)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 51, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(517)
				p.Match(FaultParserDOT)
			}
			{
				p.SetState(518)
				p.Match(FaultParserIDENT)
			}

		}
		p.SetState(523)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 51, p.GetParserRuleContext())
	}

	return localctx
}

// IStateBlockContext is an interface to support dynamic dispatch.
type IStateBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStateBlockContext differentiates from other interfaces.
	IsStateBlockContext()
}

type StateBlockContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStateBlockContext() *StateBlockContext {
	var p = new(StateBlockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FaultParserRULE_stateBlock
	return p
}

func (*StateBlockContext) IsStateBlockContext() {}

func NewStateBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StateBlockContext {
	var p = new(StateBlockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FaultParserRULE_stateBlock

	return p
}

func (s *StateBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *StateBlockContext) LCURLY() antlr.TerminalNode {
	return s.GetToken(FaultParserLCURLY, 0)
}

func (s *StateBlockContext) RCURLY() antlr.TerminalNode {
	return s.GetToken(FaultParserRCURLY, 0)
}

func (s *StateBlockContext) AllStateStep() []IStateStepContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStateStepContext); ok {
			len++
		}
	}

	tst := make([]IStateStepContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStateStepContext); ok {
			tst[i] = t.(IStateStepContext)
			i++
		}
	}

	return tst
}

func (s *StateBlockContext) StateStep(i int) IStateStepContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStateStepContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStateStepContext)
}

func (s *StateBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StateBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StateBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaultParserListener); ok {
		listenerT.EnterStateBlock(s)
	}
}

func (s *StateBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaultParserListener); ok {
		listenerT.ExitStateBlock(s)
	}
}

func (s *StateBlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FaultParserVisitor:
		return t.VisitStateBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FaultParser) StateBlock() (localctx IStateBlockContext) {
	this := p
	_ = this

	localctx = NewStateBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 88, FaultParserRULE_stateBlock)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(524)
		p.Match(FaultParserLCURLY)
	}
	p.SetState(528)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&17661981362176) != 0 {
		{
			p.SetState(525)
			p.StateStep()
		}

		p.SetState(530)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(531)
		p.Match(FaultParserRCURLY)
	}

	return localctx
}

// IStateStepContext is an interface to support dynamic dispatch.
type IStateStepContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStateStepContext differentiates from other interfaces.
	IsStateStepContext()
}

type StateStepContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStateStepContext() *StateStepContext {
	var p = new(StateStepContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FaultParserRULE_stateStep
	return p
}

func (*StateStepContext) IsStateStepContext() {}

func NewStateStepContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StateStepContext {
	var p = new(StateStepContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FaultParserRULE_stateStep

	return p
}

func (s *StateStepContext) GetParser() antlr.Parser { return s.parser }

func (s *StateStepContext) CopyFrom(ctx *StateStepContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *StateStepContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StateStepContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type StateStepExprContext struct {
	*StateStepContext
}

func NewStateStepExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StateStepExprContext {
	var p = new(StateStepExprContext)

	p.StateStepContext = NewEmptyStateStepContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StateStepContext))

	return p
}

func (s *StateStepExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StateStepExprContext) AllParamCall() []IParamCallContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IParamCallContext); ok {
			len++
		}
	}

	tst := make([]IParamCallContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IParamCallContext); ok {
			tst[i] = t.(IParamCallContext)
			i++
		}
	}

	return tst
}

func (s *StateStepExprContext) ParamCall(i int) IParamCallContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParamCallContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParamCallContext)
}

func (s *StateStepExprContext) Eos() IEosContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEosContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEosContext)
}

func (s *StateStepExprContext) AllPIPE() []antlr.TerminalNode {
	return s.GetTokens(FaultParserPIPE)
}

func (s *StateStepExprContext) PIPE(i int) antlr.TerminalNode {
	return s.GetToken(FaultParserPIPE, i)
}

func (s *StateStepExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaultParserListener); ok {
		listenerT.EnterStateStepExpr(s)
	}
}

func (s *StateStepExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaultParserListener); ok {
		listenerT.ExitStateStepExpr(s)
	}
}

func (s *StateStepExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FaultParserVisitor:
		return t.VisitStateStepExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type StateExprContext struct {
	*StateStepContext
}

func NewStateExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StateExprContext {
	var p = new(StateExprContext)

	p.StateStepContext = NewEmptyStateStepContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StateStepContext))

	return p
}

func (s *StateExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StateExprContext) IfStmtState() IIfStmtStateContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfStmtStateContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfStmtStateContext)
}

func (s *StateExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaultParserListener); ok {
		listenerT.EnterStateExpr(s)
	}
}

func (s *StateExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaultParserListener); ok {
		listenerT.ExitStateExpr(s)
	}
}

func (s *StateExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FaultParserVisitor:
		return t.VisitStateExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type StateChainContext struct {
	*StateStepContext
}

func NewStateChainContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StateChainContext {
	var p = new(StateChainContext)

	p.StateStepContext = NewEmptyStateStepContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StateStepContext))

	return p
}

func (s *StateChainContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StateChainContext) StateChange() IStateChangeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStateChangeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStateChangeContext)
}

func (s *StateChainContext) Eos() IEosContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEosContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEosContext)
}

func (s *StateChainContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaultParserListener); ok {
		listenerT.EnterStateChain(s)
	}
}

func (s *StateChainContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaultParserListener); ok {
		listenerT.ExitStateChain(s)
	}
}

func (s *StateChainContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FaultParserVisitor:
		return t.VisitStateChain(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FaultParser) StateStep() (localctx IStateStepContext) {
	this := p
	_ = this

	localctx = NewStateStepContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 90, FaultParserRULE_stateStep)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(547)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case FaultParserTHIS, FaultParserIDENT:
		localctx = NewStateStepExprContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(533)
			p.ParamCall()
		}
		p.SetState(538)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == FaultParserPIPE {
			{
				p.SetState(534)
				p.Match(FaultParserPIPE)
			}
			{
				p.SetState(535)
				p.ParamCall()
			}

			p.SetState(540)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(541)
			p.Eos()
		}

	case FaultParserADVANCE, FaultParserSTAY:
		localctx = NewStateChainContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(543)
			p.stateChange(0)
		}
		{
			p.SetState(544)
			p.Eos()
		}

	case FaultParserIF:
		localctx = NewStateExprContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(546)
			p.IfStmtState()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IRunBlockContext is an interface to support dynamic dispatch.
type IRunBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRunBlockContext differentiates from other interfaces.
	IsRunBlockContext()
}

type RunBlockContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRunBlockContext() *RunBlockContext {
	var p = new(RunBlockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FaultParserRULE_runBlock
	return p
}

func (*RunBlockContext) IsRunBlockContext() {}

func NewRunBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RunBlockContext {
	var p = new(RunBlockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FaultParserRULE_runBlock

	return p
}

func (s *RunBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *RunBlockContext) LCURLY() antlr.TerminalNode {
	return s.GetToken(FaultParserLCURLY, 0)
}

func (s *RunBlockContext) RCURLY() antlr.TerminalNode {
	return s.GetToken(FaultParserRCURLY, 0)
}

func (s *RunBlockContext) AllRunStep() []IRunStepContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IRunStepContext); ok {
			len++
		}
	}

	tst := make([]IRunStepContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IRunStepContext); ok {
			tst[i] = t.(IRunStepContext)
			i++
		}
	}

	return tst
}

func (s *RunBlockContext) RunStep(i int) IRunStepContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRunStepContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRunStepContext)
}

func (s *RunBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RunBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RunBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaultParserListener); ok {
		listenerT.EnterRunBlock(s)
	}
}

func (s *RunBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaultParserListener); ok {
		listenerT.ExitRunBlock(s)
	}
}

func (s *RunBlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FaultParserVisitor:
		return t.VisitRunBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FaultParser) RunBlock() (localctx IRunBlockContext) {
	this := p
	_ = this

	localctx = NewRunBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 92, FaultParserRULE_runBlock)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(549)
		p.Match(FaultParserLCURLY)
	}
	p.SetState(553)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 55, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(550)
				p.RunStep()
			}

		}
		p.SetState(555)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 55, p.GetParserRuleContext())
	}
	{
		p.SetState(556)
		p.Match(FaultParserRCURLY)
	}

	return localctx
}

// IRunStepContext is an interface to support dynamic dispatch.
type IRunStepContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRunStepContext differentiates from other interfaces.
	IsRunStepContext()
}

type RunStepContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRunStepContext() *RunStepContext {
	var p = new(RunStepContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FaultParserRULE_runStep
	return p
}

func (*RunStepContext) IsRunStepContext() {}

func NewRunStepContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RunStepContext {
	var p = new(RunStepContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FaultParserRULE_runStep

	return p
}

func (s *RunStepContext) GetParser() antlr.Parser { return s.parser }

func (s *RunStepContext) CopyFrom(ctx *RunStepContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *RunStepContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RunStepContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type RunStepExprContext struct {
	*RunStepContext
}

func NewRunStepExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RunStepExprContext {
	var p = new(RunStepExprContext)

	p.RunStepContext = NewEmptyRunStepContext()
	p.parser = parser
	p.CopyFrom(ctx.(*RunStepContext))

	return p
}

func (s *RunStepExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RunStepExprContext) AllParamCall() []IParamCallContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IParamCallContext); ok {
			len++
		}
	}

	tst := make([]IParamCallContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IParamCallContext); ok {
			tst[i] = t.(IParamCallContext)
			i++
		}
	}

	return tst
}

func (s *RunStepExprContext) ParamCall(i int) IParamCallContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParamCallContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParamCallContext)
}

func (s *RunStepExprContext) Eos() IEosContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEosContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEosContext)
}

func (s *RunStepExprContext) AllPIPE() []antlr.TerminalNode {
	return s.GetTokens(FaultParserPIPE)
}

func (s *RunStepExprContext) PIPE(i int) antlr.TerminalNode {
	return s.GetToken(FaultParserPIPE, i)
}

func (s *RunStepExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaultParserListener); ok {
		listenerT.EnterRunStepExpr(s)
	}
}

func (s *RunStepExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaultParserListener); ok {
		listenerT.ExitRunStepExpr(s)
	}
}

func (s *RunStepExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FaultParserVisitor:
		return t.VisitRunStepExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type RunInitContext struct {
	*RunStepContext
}

func NewRunInitContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RunInitContext {
	var p = new(RunInitContext)

	p.RunStepContext = NewEmptyRunStepContext()
	p.parser = parser
	p.CopyFrom(ctx.(*RunStepContext))

	return p
}

func (s *RunInitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RunInitContext) AllIDENT() []antlr.TerminalNode {
	return s.GetTokens(FaultParserIDENT)
}

func (s *RunInitContext) IDENT(i int) antlr.TerminalNode {
	return s.GetToken(FaultParserIDENT, i)
}

func (s *RunInitContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(FaultParserASSIGN, 0)
}

func (s *RunInitContext) NEW() antlr.TerminalNode {
	return s.GetToken(FaultParserNEW, 0)
}

func (s *RunInitContext) Eos() IEosContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEosContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEosContext)
}

func (s *RunInitContext) ParamCall() IParamCallContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParamCallContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParamCallContext)
}

func (s *RunInitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaultParserListener); ok {
		listenerT.EnterRunInit(s)
	}
}

func (s *RunInitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaultParserListener); ok {
		listenerT.ExitRunInit(s)
	}
}

func (s *RunInitContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FaultParserVisitor:
		return t.VisitRunInit(s)

	default:
		return t.VisitChildren(s)
	}
}

type RunExprContext struct {
	*RunStepContext
}

func NewRunExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RunExprContext {
	var p = new(RunExprContext)

	p.RunStepContext = NewEmptyRunStepContext()
	p.parser = parser
	p.CopyFrom(ctx.(*RunStepContext))

	return p
}

func (s *RunExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RunExprContext) SimpleStmt() ISimpleStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISimpleStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISimpleStmtContext)
}

func (s *RunExprContext) Eos() IEosContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEosContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEosContext)
}

func (s *RunExprContext) IfStmtRun() IIfStmtRunContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfStmtRunContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfStmtRunContext)
}

func (s *RunExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaultParserListener); ok {
		listenerT.EnterRunExpr(s)
	}
}

func (s *RunExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaultParserListener); ok {
		listenerT.ExitRunExpr(s)
	}
}

func (s *RunExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FaultParserVisitor:
		return t.VisitRunExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FaultParser) RunStep() (localctx IRunStepContext) {
	this := p
	_ = this

	localctx = NewRunStepContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 94, FaultParserRULE_runStep)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(580)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 58, p.GetParserRuleContext()) {
	case 1:
		localctx = NewRunStepExprContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(558)
			p.ParamCall()
		}
		p.SetState(563)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == FaultParserPIPE {
			{
				p.SetState(559)
				p.Match(FaultParserPIPE)
			}
			{
				p.SetState(560)
				p.ParamCall()
			}

			p.SetState(565)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(566)
			p.Eos()
		}

	case 2:
		localctx = NewRunInitContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(568)
			p.Match(FaultParserIDENT)
		}
		{
			p.SetState(569)
			p.Match(FaultParserASSIGN)
		}
		{
			p.SetState(570)
			p.Match(FaultParserNEW)
		}
		p.SetState(573)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 57, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(571)
				p.ParamCall()
			}

		case 2:
			{
				p.SetState(572)
				p.Match(FaultParserIDENT)
			}

		}
		{
			p.SetState(575)
			p.Eos()
		}

	case 3:
		localctx = NewRunExprContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(576)
			p.SimpleStmt()
		}
		{
			p.SetState(577)
			p.Eos()
		}

	case 4:
		localctx = NewRunExprContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(579)
			p.IfStmtRun()
		}

	}

	return localctx
}

// IFaultTypeContext is an interface to support dynamic dispatch.
type IFaultTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFaultTypeContext differentiates from other interfaces.
	IsFaultTypeContext()
}

type FaultTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFaultTypeContext() *FaultTypeContext {
	var p = new(FaultTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FaultParserRULE_faultType
	return p
}

func (*FaultTypeContext) IsFaultTypeContext() {}

func NewFaultTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FaultTypeContext {
	var p = new(FaultTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FaultParserRULE_faultType

	return p
}

func (s *FaultTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *FaultTypeContext) TY_STRING() antlr.TerminalNode {
	return s.GetToken(FaultParserTY_STRING, 0)
}

func (s *FaultTypeContext) TY_BOOL() antlr.TerminalNode {
	return s.GetToken(FaultParserTY_BOOL, 0)
}

func (s *FaultTypeContext) TY_INT() antlr.TerminalNode {
	return s.GetToken(FaultParserTY_INT, 0)
}

func (s *FaultTypeContext) TY_FLOAT() antlr.TerminalNode {
	return s.GetToken(FaultParserTY_FLOAT, 0)
}

func (s *FaultTypeContext) TY_NATURAL() antlr.TerminalNode {
	return s.GetToken(FaultParserTY_NATURAL, 0)
}

func (s *FaultTypeContext) TY_UNCERTAIN() antlr.TerminalNode {
	return s.GetToken(FaultParserTY_UNCERTAIN, 0)
}

func (s *FaultTypeContext) TY_UNKNOWN() antlr.TerminalNode {
	return s.GetToken(FaultParserTY_UNKNOWN, 0)
}

func (s *FaultTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FaultTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FaultTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaultParserListener); ok {
		listenerT.EnterFaultType(s)
	}
}

func (s *FaultTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaultParserListener); ok {
		listenerT.ExitFaultType(s)
	}
}

func (s *FaultTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FaultParserVisitor:
		return t.VisitFaultType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FaultParser) FaultType() (localctx IFaultTypeContext) {
	this := p
	_ = this

	localctx = NewFaultTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 96, FaultParserRULE_faultType)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(582)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&17454747090944) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// ISolvableContext is an interface to support dynamic dispatch.
type ISolvableContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSolvableContext differentiates from other interfaces.
	IsSolvableContext()
}

type SolvableContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySolvableContext() *SolvableContext {
	var p = new(SolvableContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FaultParserRULE_solvable
	return p
}

func (*SolvableContext) IsSolvableContext() {}

func NewSolvableContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SolvableContext {
	var p = new(SolvableContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FaultParserRULE_solvable

	return p
}

func (s *SolvableContext) GetParser() antlr.Parser { return s.parser }

func (s *SolvableContext) FaultType() IFaultTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFaultTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFaultTypeContext)
}

func (s *SolvableContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(FaultParserLPAREN, 0)
}

func (s *SolvableContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(FaultParserRPAREN, 0)
}

func (s *SolvableContext) AllOperand() []IOperandContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IOperandContext); ok {
			len++
		}
	}

	tst := make([]IOperandContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IOperandContext); ok {
			tst[i] = t.(IOperandContext)
			i++
		}
	}

	return tst
}

func (s *SolvableContext) Operand(i int) IOperandContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOperandContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOperandContext)
}

func (s *SolvableContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(FaultParserCOMMA)
}

func (s *SolvableContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(FaultParserCOMMA, i)
}

func (s *SolvableContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SolvableContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SolvableContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaultParserListener); ok {
		listenerT.EnterSolvable(s)
	}
}

func (s *SolvableContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaultParserListener); ok {
		listenerT.ExitSolvable(s)
	}
}

func (s *SolvableContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FaultParserVisitor:
		return t.VisitSolvable(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FaultParser) Solvable() (localctx ISolvableContext) {
	this := p
	_ = this

	localctx = NewSolvableContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 98, FaultParserRULE_solvable)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(584)
		p.FaultType()
	}
	{
		p.SetState(585)
		p.Match(FaultParserLPAREN)
	}
	p.SetState(587)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2269392941367312) != 0 || (int64((_la-72)) & ^0x3f) == 0 && ((int64(1)<<(_la-72))&32257) != 0 {
		{
			p.SetState(586)
			p.Operand()
		}

	}
	p.SetState(593)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == FaultParserCOMMA {
		{
			p.SetState(589)
			p.Match(FaultParserCOMMA)
		}
		{
			p.SetState(590)
			p.Operand()
		}

		p.SetState(595)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(596)
		p.Match(FaultParserRPAREN)
	}

	return localctx
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FaultParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FaultParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) CopyFrom(ctx *ExpressionContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type TypedContext struct {
	*ExpressionContext
}

func NewTypedContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TypedContext {
	var p = new(TypedContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *TypedContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypedContext) Solvable() ISolvableContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISolvableContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISolvableContext)
}

func (s *TypedContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaultParserListener); ok {
		listenerT.EnterTyped(s)
	}
}

func (s *TypedContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaultParserListener); ok {
		listenerT.ExitTyped(s)
	}
}

func (s *TypedContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FaultParserVisitor:
		return t.VisitTyped(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExprContext struct {
	*ExpressionContext
}

func NewExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExprContext {
	var p = new(ExprContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) Operand() IOperandContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOperandContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOperandContext)
}

func (s *ExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaultParserListener); ok {
		listenerT.EnterExpr(s)
	}
}

func (s *ExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaultParserListener); ok {
		listenerT.ExitExpr(s)
	}
}

func (s *ExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FaultParserVisitor:
		return t.VisitExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExprPrefixContext struct {
	*ExpressionContext
}

func NewExprPrefixContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExprPrefixContext {
	var p = new(ExprPrefixContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *ExprPrefixContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprPrefixContext) Prefix() IPrefixContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrefixContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrefixContext)
}

func (s *ExprPrefixContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaultParserListener); ok {
		listenerT.EnterExprPrefix(s)
	}
}

func (s *ExprPrefixContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaultParserListener); ok {
		listenerT.ExitExprPrefix(s)
	}
}

func (s *ExprPrefixContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FaultParserVisitor:
		return t.VisitExprPrefix(s)

	default:
		return t.VisitChildren(s)
	}
}

type LrExprContext struct {
	*ExpressionContext
}

func NewLrExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LrExprContext {
	var p = new(LrExprContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *LrExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LrExprContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *LrExprContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *LrExprContext) EXPO() antlr.TerminalNode {
	return s.GetToken(FaultParserEXPO, 0)
}

func (s *LrExprContext) MULTI() antlr.TerminalNode {
	return s.GetToken(FaultParserMULTI, 0)
}

func (s *LrExprContext) DIV() antlr.TerminalNode {
	return s.GetToken(FaultParserDIV, 0)
}

func (s *LrExprContext) MOD() antlr.TerminalNode {
	return s.GetToken(FaultParserMOD, 0)
}

func (s *LrExprContext) LSHIFT() antlr.TerminalNode {
	return s.GetToken(FaultParserLSHIFT, 0)
}

func (s *LrExprContext) RSHIFT() antlr.TerminalNode {
	return s.GetToken(FaultParserRSHIFT, 0)
}

func (s *LrExprContext) AMPERSAND() antlr.TerminalNode {
	return s.GetToken(FaultParserAMPERSAND, 0)
}

func (s *LrExprContext) BIT_CLEAR() antlr.TerminalNode {
	return s.GetToken(FaultParserBIT_CLEAR, 0)
}

func (s *LrExprContext) PLUS() antlr.TerminalNode {
	return s.GetToken(FaultParserPLUS, 0)
}

func (s *LrExprContext) MINUS() antlr.TerminalNode {
	return s.GetToken(FaultParserMINUS, 0)
}

func (s *LrExprContext) CARET() antlr.TerminalNode {
	return s.GetToken(FaultParserCARET, 0)
}

func (s *LrExprContext) EQUALS() antlr.TerminalNode {
	return s.GetToken(FaultParserEQUALS, 0)
}

func (s *LrExprContext) NOT_EQUALS() antlr.TerminalNode {
	return s.GetToken(FaultParserNOT_EQUALS, 0)
}

func (s *LrExprContext) LESS() antlr.TerminalNode {
	return s.GetToken(FaultParserLESS, 0)
}

func (s *LrExprContext) LESS_OR_EQUALS() antlr.TerminalNode {
	return s.GetToken(FaultParserLESS_OR_EQUALS, 0)
}

func (s *LrExprContext) GREATER() antlr.TerminalNode {
	return s.GetToken(FaultParserGREATER, 0)
}

func (s *LrExprContext) GREATER_OR_EQUALS() antlr.TerminalNode {
	return s.GetToken(FaultParserGREATER_OR_EQUALS, 0)
}

func (s *LrExprContext) AND() antlr.TerminalNode {
	return s.GetToken(FaultParserAND, 0)
}

func (s *LrExprContext) OR() antlr.TerminalNode {
	return s.GetToken(FaultParserOR, 0)
}

func (s *LrExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaultParserListener); ok {
		listenerT.EnterLrExpr(s)
	}
}

func (s *LrExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaultParserListener); ok {
		listenerT.ExitLrExpr(s)
	}
}

func (s *LrExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FaultParserVisitor:
		return t.VisitLrExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FaultParser) Expression() (localctx IExpressionContext) {
	return p.expression(0)
}

func (p *FaultParser) expression(_p int) (localctx IExpressionContext) {
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 100
	p.EnterRecursionRule(localctx, 100, FaultParserRULE_expression, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(602)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 61, p.GetParserRuleContext()) {
	case 1:
		localctx = NewExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(599)
			p.Operand()
		}

	case 2:
		localctx = NewTypedContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(600)
			p.Solvable()
		}

	case 3:
		localctx = NewExprPrefixContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(601)
			p.Prefix()
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(624)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 63, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(622)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 62, p.GetParserRuleContext()) {
			case 1:
				localctx = NewLrExprContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, FaultParserRULE_expression)
				p.SetState(604)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
				}
				{
					p.SetState(605)
					p.Match(FaultParserEXPO)
				}
				{
					p.SetState(606)
					p.expression(7)
				}

			case 2:
				localctx = NewLrExprContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, FaultParserRULE_expression)
				p.SetState(607)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				{
					p.SetState(608)
					_la = p.GetTokenStream().LA(1)

					if !((int64((_la-60)) & ^0x3f) == 0 && ((int64(1)<<(_la-60))&2064385) != 0) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(609)
					p.expression(6)
				}

			case 3:
				localctx = NewLrExprContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, FaultParserRULE_expression)
				p.SetState(610)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(611)
					_la = p.GetTokenStream().LA(1)

					if !((int64((_la-71)) & ^0x3f) == 0 && ((int64(1)<<(_la-71))&7) != 0) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(612)
					p.expression(5)
				}

			case 4:
				localctx = NewLrExprContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, FaultParserRULE_expression)
				p.SetState(613)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(614)
					_la = p.GetTokenStream().LA(1)

					if !((int64((_la-63)) & ^0x3f) == 0 && ((int64(1)<<(_la-63))&63) != 0) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(615)
					p.expression(4)
				}

			case 5:
				localctx = NewLrExprContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, FaultParserRULE_expression)
				p.SetState(616)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				{
					p.SetState(617)
					p.Match(FaultParserAND)
				}
				{
					p.SetState(618)
					p.expression(3)
				}

			case 6:
				localctx = NewLrExprContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, FaultParserRULE_expression)
				p.SetState(619)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
				}
				{
					p.SetState(620)
					p.Match(FaultParserOR)
				}
				{
					p.SetState(621)
					p.expression(2)
				}

			}

		}
		p.SetState(626)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 63, p.GetParserRuleContext())
	}

	return localctx
}

// IOperandContext is an interface to support dynamic dispatch.
type IOperandContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOperandContext differentiates from other interfaces.
	IsOperandContext()
}

type OperandContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOperandContext() *OperandContext {
	var p = new(OperandContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FaultParserRULE_operand
	return p
}

func (*OperandContext) IsOperandContext() {}

func NewOperandContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OperandContext {
	var p = new(OperandContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FaultParserRULE_operand

	return p
}

func (s *OperandContext) GetParser() antlr.Parser { return s.parser }

func (s *OperandContext) Nil_() INilContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INilContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INilContext)
}

func (s *OperandContext) Numeric() INumericContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INumericContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INumericContext)
}

func (s *OperandContext) String_() IString_Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IString_Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IString_Context)
}

func (s *OperandContext) Bool_() IBool_Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBool_Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBool_Context)
}

func (s *OperandContext) OperandName() IOperandNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOperandNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOperandNameContext)
}

func (s *OperandContext) AccessHistory() IAccessHistoryContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAccessHistoryContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAccessHistoryContext)
}

func (s *OperandContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(FaultParserLPAREN, 0)
}

func (s *OperandContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *OperandContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(FaultParserRPAREN, 0)
}

func (s *OperandContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperandContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OperandContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaultParserListener); ok {
		listenerT.EnterOperand(s)
	}
}

func (s *OperandContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaultParserListener); ok {
		listenerT.ExitOperand(s)
	}
}

func (s *OperandContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FaultParserVisitor:
		return t.VisitOperand(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FaultParser) Operand() (localctx IOperandContext) {
	this := p
	_ = this

	localctx = NewOperandContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 102, FaultParserRULE_operand)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(637)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 64, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(627)
			p.Nil_()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(628)
			p.Numeric()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(629)
			p.String_()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(630)
			p.Bool_()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(631)
			p.OperandName()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(632)
			p.AccessHistory()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(633)
			p.Match(FaultParserLPAREN)
		}
		{
			p.SetState(634)
			p.expression(0)
		}
		{
			p.SetState(635)
			p.Match(FaultParserRPAREN)
		}

	}

	return localctx
}

// IOperandNameContext is an interface to support dynamic dispatch.
type IOperandNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOperandNameContext differentiates from other interfaces.
	IsOperandNameContext()
}

type OperandNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOperandNameContext() *OperandNameContext {
	var p = new(OperandNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FaultParserRULE_operandName
	return p
}

func (*OperandNameContext) IsOperandNameContext() {}

func NewOperandNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OperandNameContext {
	var p = new(OperandNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FaultParserRULE_operandName

	return p
}

func (s *OperandNameContext) GetParser() antlr.Parser { return s.parser }

func (s *OperandNameContext) CopyFrom(ctx *OperandNameContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *OperandNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperandNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type OpClockContext struct {
	*OperandNameContext
}

func NewOpClockContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *OpClockContext {
	var p = new(OpClockContext)

	p.OperandNameContext = NewEmptyOperandNameContext()
	p.parser = parser
	p.CopyFrom(ctx.(*OperandNameContext))

	return p
}

func (s *OpClockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OpClockContext) CLOCK() antlr.TerminalNode {
	return s.GetToken(FaultParserCLOCK, 0)
}

func (s *OpClockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaultParserListener); ok {
		listenerT.EnterOpClock(s)
	}
}

func (s *OpClockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaultParserListener); ok {
		listenerT.ExitOpClock(s)
	}
}

func (s *OpClockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FaultParserVisitor:
		return t.VisitOpClock(s)

	default:
		return t.VisitChildren(s)
	}
}

type OpNameContext struct {
	*OperandNameContext
}

func NewOpNameContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *OpNameContext {
	var p = new(OpNameContext)

	p.OperandNameContext = NewEmptyOperandNameContext()
	p.parser = parser
	p.CopyFrom(ctx.(*OperandNameContext))

	return p
}

func (s *OpNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OpNameContext) IDENT() antlr.TerminalNode {
	return s.GetToken(FaultParserIDENT, 0)
}

func (s *OpNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaultParserListener); ok {
		listenerT.EnterOpName(s)
	}
}

func (s *OpNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaultParserListener); ok {
		listenerT.ExitOpName(s)
	}
}

func (s *OpNameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FaultParserVisitor:
		return t.VisitOpName(s)

	default:
		return t.VisitChildren(s)
	}
}

type OpParamContext struct {
	*OperandNameContext
}

func NewOpParamContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *OpParamContext {
	var p = new(OpParamContext)

	p.OperandNameContext = NewEmptyOperandNameContext()
	p.parser = parser
	p.CopyFrom(ctx.(*OperandNameContext))

	return p
}

func (s *OpParamContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OpParamContext) ParamCall() IParamCallContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParamCallContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParamCallContext)
}

func (s *OpParamContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaultParserListener); ok {
		listenerT.EnterOpParam(s)
	}
}

func (s *OpParamContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaultParserListener); ok {
		listenerT.ExitOpParam(s)
	}
}

func (s *OpParamContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FaultParserVisitor:
		return t.VisitOpParam(s)

	default:
		return t.VisitChildren(s)
	}
}

type OpInstanceContext struct {
	*OperandNameContext
}

func NewOpInstanceContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *OpInstanceContext {
	var p = new(OpInstanceContext)

	p.OperandNameContext = NewEmptyOperandNameContext()
	p.parser = parser
	p.CopyFrom(ctx.(*OperandNameContext))

	return p
}

func (s *OpInstanceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OpInstanceContext) NEW() antlr.TerminalNode {
	return s.GetToken(FaultParserNEW, 0)
}

func (s *OpInstanceContext) AllIDENT() []antlr.TerminalNode {
	return s.GetTokens(FaultParserIDENT)
}

func (s *OpInstanceContext) IDENT(i int) antlr.TerminalNode {
	return s.GetToken(FaultParserIDENT, i)
}

func (s *OpInstanceContext) DOT() antlr.TerminalNode {
	return s.GetToken(FaultParserDOT, 0)
}

func (s *OpInstanceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaultParserListener); ok {
		listenerT.EnterOpInstance(s)
	}
}

func (s *OpInstanceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaultParserListener); ok {
		listenerT.ExitOpInstance(s)
	}
}

func (s *OpInstanceContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FaultParserVisitor:
		return t.VisitOpInstance(s)

	default:
		return t.VisitChildren(s)
	}
}

type OpThisContext struct {
	*OperandNameContext
}

func NewOpThisContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *OpThisContext {
	var p = new(OpThisContext)

	p.OperandNameContext = NewEmptyOperandNameContext()
	p.parser = parser
	p.CopyFrom(ctx.(*OperandNameContext))

	return p
}

func (s *OpThisContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OpThisContext) THIS() antlr.TerminalNode {
	return s.GetToken(FaultParserTHIS, 0)
}

func (s *OpThisContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaultParserListener); ok {
		listenerT.EnterOpThis(s)
	}
}

func (s *OpThisContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaultParserListener); ok {
		listenerT.ExitOpThis(s)
	}
}

func (s *OpThisContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FaultParserVisitor:
		return t.VisitOpThis(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FaultParser) OperandName() (localctx IOperandNameContext) {
	this := p
	_ = this

	localctx = NewOperandNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 104, FaultParserRULE_operandName)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(649)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 66, p.GetParserRuleContext()) {
	case 1:
		localctx = NewOpNameContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(639)
			p.Match(FaultParserIDENT)
		}

	case 2:
		localctx = NewOpParamContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(640)
			p.ParamCall()
		}

	case 3:
		localctx = NewOpThisContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(641)
			p.Match(FaultParserTHIS)
		}

	case 4:
		localctx = NewOpClockContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(642)
			p.Match(FaultParserCLOCK)
		}

	case 5:
		localctx = NewOpInstanceContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(643)
			p.Match(FaultParserNEW)
		}
		{
			p.SetState(644)
			p.Match(FaultParserIDENT)
		}
		p.SetState(647)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 65, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(645)
				p.Match(FaultParserDOT)
			}
			{
				p.SetState(646)
				p.Match(FaultParserIDENT)
			}

		}

	}

	return localctx
}

// IPrefixContext is an interface to support dynamic dispatch.
type IPrefixContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPrefixContext differentiates from other interfaces.
	IsPrefixContext()
}

type PrefixContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrefixContext() *PrefixContext {
	var p = new(PrefixContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FaultParserRULE_prefix
	return p
}

func (*PrefixContext) IsPrefixContext() {}

func NewPrefixContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrefixContext {
	var p = new(PrefixContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FaultParserRULE_prefix

	return p
}

func (s *PrefixContext) GetParser() antlr.Parser { return s.parser }

func (s *PrefixContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *PrefixContext) PLUS() antlr.TerminalNode {
	return s.GetToken(FaultParserPLUS, 0)
}

func (s *PrefixContext) MINUS() antlr.TerminalNode {
	return s.GetToken(FaultParserMINUS, 0)
}

func (s *PrefixContext) BANG() antlr.TerminalNode {
	return s.GetToken(FaultParserBANG, 0)
}

func (s *PrefixContext) CARET() antlr.TerminalNode {
	return s.GetToken(FaultParserCARET, 0)
}

func (s *PrefixContext) MULTI() antlr.TerminalNode {
	return s.GetToken(FaultParserMULTI, 0)
}

func (s *PrefixContext) AMPERSAND() antlr.TerminalNode {
	return s.GetToken(FaultParserAMPERSAND, 0)
}

func (s *PrefixContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrefixContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrefixContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaultParserListener); ok {
		listenerT.EnterPrefix(s)
	}
}

func (s *PrefixContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaultParserListener); ok {
		listenerT.ExitPrefix(s)
	}
}

func (s *PrefixContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FaultParserVisitor:
		return t.VisitPrefix(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FaultParser) Prefix() (localctx IPrefixContext) {
	this := p
	_ = this

	localctx = NewPrefixContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 106, FaultParserRULE_prefix)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(654)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 67, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(652)
			_la = p.GetTokenStream().LA(1)

			if !((int64((_la-60)) & ^0x3f) == 0 && ((int64(1)<<(_la-60))&47109) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(653)
			p.expression(0)
		}

	}

	return localctx
}

// INumericContext is an interface to support dynamic dispatch.
type INumericContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNumericContext differentiates from other interfaces.
	IsNumericContext()
}

type NumericContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNumericContext() *NumericContext {
	var p = new(NumericContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FaultParserRULE_numeric
	return p
}

func (*NumericContext) IsNumericContext() {}

func NewNumericContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NumericContext {
	var p = new(NumericContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FaultParserRULE_numeric

	return p
}

func (s *NumericContext) GetParser() antlr.Parser { return s.parser }

func (s *NumericContext) Integer() IIntegerContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIntegerContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIntegerContext)
}

func (s *NumericContext) Negative() INegativeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INegativeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INegativeContext)
}

func (s *NumericContext) Float_() IFloat_Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFloat_Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFloat_Context)
}

func (s *NumericContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumericContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NumericContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaultParserListener); ok {
		listenerT.EnterNumeric(s)
	}
}

func (s *NumericContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaultParserListener); ok {
		listenerT.ExitNumeric(s)
	}
}

func (s *NumericContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FaultParserVisitor:
		return t.VisitNumeric(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FaultParser) Numeric() (localctx INumericContext) {
	this := p
	_ = this

	localctx = NewNumericContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 108, FaultParserRULE_numeric)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(659)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case FaultParserDECIMAL_LIT, FaultParserOCTAL_LIT, FaultParserHEX_LIT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(656)
			p.Integer()
		}

	case FaultParserMINUS:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(657)
			p.Negative()
		}

	case FaultParserFLOAT_LIT:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(658)
			p.Float_()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IIntegerContext is an interface to support dynamic dispatch.
type IIntegerContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIntegerContext differentiates from other interfaces.
	IsIntegerContext()
}

type IntegerContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIntegerContext() *IntegerContext {
	var p = new(IntegerContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FaultParserRULE_integer
	return p
}

func (*IntegerContext) IsIntegerContext() {}

func NewIntegerContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IntegerContext {
	var p = new(IntegerContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FaultParserRULE_integer

	return p
}

func (s *IntegerContext) GetParser() antlr.Parser { return s.parser }

func (s *IntegerContext) DECIMAL_LIT() antlr.TerminalNode {
	return s.GetToken(FaultParserDECIMAL_LIT, 0)
}

func (s *IntegerContext) OCTAL_LIT() antlr.TerminalNode {
	return s.GetToken(FaultParserOCTAL_LIT, 0)
}

func (s *IntegerContext) HEX_LIT() antlr.TerminalNode {
	return s.GetToken(FaultParserHEX_LIT, 0)
}

func (s *IntegerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntegerContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IntegerContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaultParserListener); ok {
		listenerT.EnterInteger(s)
	}
}

func (s *IntegerContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaultParserListener); ok {
		listenerT.ExitInteger(s)
	}
}

func (s *IntegerContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FaultParserVisitor:
		return t.VisitInteger(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FaultParser) Integer() (localctx IIntegerContext) {
	this := p
	_ = this

	localctx = NewIntegerContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 110, FaultParserRULE_integer)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(661)
		_la = p.GetTokenStream().LA(1)

		if !((int64((_la-81)) & ^0x3f) == 0 && ((int64(1)<<(_la-81))&7) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// INegativeContext is an interface to support dynamic dispatch.
type INegativeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNegativeContext differentiates from other interfaces.
	IsNegativeContext()
}

type NegativeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNegativeContext() *NegativeContext {
	var p = new(NegativeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FaultParserRULE_negative
	return p
}

func (*NegativeContext) IsNegativeContext() {}

func NewNegativeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NegativeContext {
	var p = new(NegativeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FaultParserRULE_negative

	return p
}

func (s *NegativeContext) GetParser() antlr.Parser { return s.parser }

func (s *NegativeContext) MINUS() antlr.TerminalNode {
	return s.GetToken(FaultParserMINUS, 0)
}

func (s *NegativeContext) Integer() IIntegerContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIntegerContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIntegerContext)
}

func (s *NegativeContext) Float_() IFloat_Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFloat_Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFloat_Context)
}

func (s *NegativeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NegativeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NegativeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaultParserListener); ok {
		listenerT.EnterNegative(s)
	}
}

func (s *NegativeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaultParserListener); ok {
		listenerT.ExitNegative(s)
	}
}

func (s *NegativeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FaultParserVisitor:
		return t.VisitNegative(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FaultParser) Negative() (localctx INegativeContext) {
	this := p
	_ = this

	localctx = NewNegativeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 112, FaultParserRULE_negative)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(667)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 69, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(663)
			p.Match(FaultParserMINUS)
		}
		{
			p.SetState(664)
			p.Integer()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(665)
			p.Match(FaultParserMINUS)
		}
		{
			p.SetState(666)
			p.Float_()
		}

	}

	return localctx
}

// IFloat_Context is an interface to support dynamic dispatch.
type IFloat_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFloat_Context differentiates from other interfaces.
	IsFloat_Context()
}

type Float_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFloat_Context() *Float_Context {
	var p = new(Float_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FaultParserRULE_float_
	return p
}

func (*Float_Context) IsFloat_Context() {}

func NewFloat_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Float_Context {
	var p = new(Float_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FaultParserRULE_float_

	return p
}

func (s *Float_Context) GetParser() antlr.Parser { return s.parser }

func (s *Float_Context) FLOAT_LIT() antlr.TerminalNode {
	return s.GetToken(FaultParserFLOAT_LIT, 0)
}

func (s *Float_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Float_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Float_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaultParserListener); ok {
		listenerT.EnterFloat_(s)
	}
}

func (s *Float_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaultParserListener); ok {
		listenerT.ExitFloat_(s)
	}
}

func (s *Float_Context) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FaultParserVisitor:
		return t.VisitFloat_(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FaultParser) Float_() (localctx IFloat_Context) {
	this := p
	_ = this

	localctx = NewFloat_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 114, FaultParserRULE_float_)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(669)
		p.Match(FaultParserFLOAT_LIT)
	}

	return localctx
}

// IString_Context is an interface to support dynamic dispatch.
type IString_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsString_Context differentiates from other interfaces.
	IsString_Context()
}

type String_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyString_Context() *String_Context {
	var p = new(String_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FaultParserRULE_string_
	return p
}

func (*String_Context) IsString_Context() {}

func NewString_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *String_Context {
	var p = new(String_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FaultParserRULE_string_

	return p
}

func (s *String_Context) GetParser() antlr.Parser { return s.parser }

func (s *String_Context) RAW_STRING_LIT() antlr.TerminalNode {
	return s.GetToken(FaultParserRAW_STRING_LIT, 0)
}

func (s *String_Context) INTERPRETED_STRING_LIT() antlr.TerminalNode {
	return s.GetToken(FaultParserINTERPRETED_STRING_LIT, 0)
}

func (s *String_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *String_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *String_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaultParserListener); ok {
		listenerT.EnterString_(s)
	}
}

func (s *String_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaultParserListener); ok {
		listenerT.ExitString_(s)
	}
}

func (s *String_Context) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FaultParserVisitor:
		return t.VisitString_(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FaultParser) String_() (localctx IString_Context) {
	this := p
	_ = this

	localctx = NewString_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 116, FaultParserRULE_string_)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(671)
		_la = p.GetTokenStream().LA(1)

		if !(_la == FaultParserRAW_STRING_LIT || _la == FaultParserINTERPRETED_STRING_LIT) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IBool_Context is an interface to support dynamic dispatch.
type IBool_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBool_Context differentiates from other interfaces.
	IsBool_Context()
}

type Bool_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBool_Context() *Bool_Context {
	var p = new(Bool_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FaultParserRULE_bool_
	return p
}

func (*Bool_Context) IsBool_Context() {}

func NewBool_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Bool_Context {
	var p = new(Bool_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FaultParserRULE_bool_

	return p
}

func (s *Bool_Context) GetParser() antlr.Parser { return s.parser }

func (s *Bool_Context) TRUE() antlr.TerminalNode {
	return s.GetToken(FaultParserTRUE, 0)
}

func (s *Bool_Context) FALSE() antlr.TerminalNode {
	return s.GetToken(FaultParserFALSE, 0)
}

func (s *Bool_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Bool_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Bool_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaultParserListener); ok {
		listenerT.EnterBool_(s)
	}
}

func (s *Bool_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaultParserListener); ok {
		listenerT.ExitBool_(s)
	}
}

func (s *Bool_Context) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FaultParserVisitor:
		return t.VisitBool_(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FaultParser) Bool_() (localctx IBool_Context) {
	this := p
	_ = this

	localctx = NewBool_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 118, FaultParserRULE_bool_)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(673)
		_la = p.GetTokenStream().LA(1)

		if !(_la == FaultParserTRUE || _la == FaultParserFALSE) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IFunctionLitContext is an interface to support dynamic dispatch.
type IFunctionLitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctionLitContext differentiates from other interfaces.
	IsFunctionLitContext()
}

type FunctionLitContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionLitContext() *FunctionLitContext {
	var p = new(FunctionLitContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FaultParserRULE_functionLit
	return p
}

func (*FunctionLitContext) IsFunctionLitContext() {}

func NewFunctionLitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionLitContext {
	var p = new(FunctionLitContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FaultParserRULE_functionLit

	return p
}

func (s *FunctionLitContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionLitContext) FUNC() antlr.TerminalNode {
	return s.GetToken(FaultParserFUNC, 0)
}

func (s *FunctionLitContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *FunctionLitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionLitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionLitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaultParserListener); ok {
		listenerT.EnterFunctionLit(s)
	}
}

func (s *FunctionLitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaultParserListener); ok {
		listenerT.ExitFunctionLit(s)
	}
}

func (s *FunctionLitContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FaultParserVisitor:
		return t.VisitFunctionLit(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FaultParser) FunctionLit() (localctx IFunctionLitContext) {
	this := p
	_ = this

	localctx = NewFunctionLitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 120, FaultParserRULE_functionLit)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(675)
		p.Match(FaultParserFUNC)
	}
	{
		p.SetState(676)
		p.Block()
	}

	return localctx
}

// IStateLitContext is an interface to support dynamic dispatch.
type IStateLitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStateLitContext differentiates from other interfaces.
	IsStateLitContext()
}

type StateLitContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStateLitContext() *StateLitContext {
	var p = new(StateLitContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FaultParserRULE_stateLit
	return p
}

func (*StateLitContext) IsStateLitContext() {}

func NewStateLitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StateLitContext {
	var p = new(StateLitContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FaultParserRULE_stateLit

	return p
}

func (s *StateLitContext) GetParser() antlr.Parser { return s.parser }

func (s *StateLitContext) FUNC() antlr.TerminalNode {
	return s.GetToken(FaultParserFUNC, 0)
}

func (s *StateLitContext) StateBlock() IStateBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStateBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStateBlockContext)
}

func (s *StateLitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StateLitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StateLitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaultParserListener); ok {
		listenerT.EnterStateLit(s)
	}
}

func (s *StateLitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaultParserListener); ok {
		listenerT.ExitStateLit(s)
	}
}

func (s *StateLitContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FaultParserVisitor:
		return t.VisitStateLit(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FaultParser) StateLit() (localctx IStateLitContext) {
	this := p
	_ = this

	localctx = NewStateLitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 122, FaultParserRULE_stateLit)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(678)
		p.Match(FaultParserFUNC)
	}
	{
		p.SetState(679)
		p.StateBlock()
	}

	return localctx
}

// IEosContext is an interface to support dynamic dispatch.
type IEosContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEosContext differentiates from other interfaces.
	IsEosContext()
}

type EosContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEosContext() *EosContext {
	var p = new(EosContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FaultParserRULE_eos
	return p
}

func (*EosContext) IsEosContext() {}

func NewEosContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EosContext {
	var p = new(EosContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FaultParserRULE_eos

	return p
}

func (s *EosContext) GetParser() antlr.Parser { return s.parser }

func (s *EosContext) SEMI() antlr.TerminalNode {
	return s.GetToken(FaultParserSEMI, 0)
}

func (s *EosContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EosContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EosContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaultParserListener); ok {
		listenerT.EnterEos(s)
	}
}

func (s *EosContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaultParserListener); ok {
		listenerT.ExitEos(s)
	}
}

func (s *EosContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FaultParserVisitor:
		return t.VisitEos(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FaultParser) Eos() (localctx IEosContext) {
	this := p
	_ = this

	localctx = NewEosContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 124, FaultParserRULE_eos)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(681)
		p.Match(FaultParserSEMI)
	}

	return localctx
}

func (p *FaultParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 30:
		var t *StateChangeContext = nil
		if localctx != nil {
			t = localctx.(*StateChangeContext)
		}
		return p.StateChange_Sempred(t, predIndex)

	case 50:
		var t *ExpressionContext = nil
		if localctx != nil {
			t = localctx.(*ExpressionContext)
		}
		return p.Expression_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *FaultParser) StateChange_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 2)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *FaultParser) Expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 2:
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 2)

	case 7:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
