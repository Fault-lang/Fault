// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package parser

import (
	"fmt"
	"sync"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type FaultLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var faultlexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	channelNames           []string
	modeNames              []string
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func faultlexerLexerInit() {
	staticData := &faultlexerLexerStaticData
	staticData.channelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.modeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.literalNames = []string{
		"", "'all'", "'assert'", "'assume'", "'now'", "'const'", "'def'", "'else'",
		"'flow'", "'for'", "'func'", "'if'", "'import'", "'init'", "'new'",
		"'return'", "'run'", "'spec'", "'stock'", "'then'", "'when'", "'this'",
		"'eventually'", "'eventually-always'", "'always'", "'nmt'", "'nft'",
		"'nil'", "'true'", "'false'", "'advance'", "'component'", "'global'",
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
		"ALL", "ASSERT", "ASSUME", "CLOCK", "CONST", "DEF", "ELSE", "FLOW",
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
		"WS", "COMMENT", "TERMINATOR", "LINE_COMMENT", "ESCAPED_VALUE", "DECIMALS",
		"OCTAL_DIGIT", "HEX_DIGIT", "EXPONENT", "LETTER", "UNICODE_DIGIT", "UNICODE_LETTER",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 90, 705, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36,
		7, 36, 2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7,
		41, 2, 42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 2, 45, 7, 45, 2, 46, 7, 46,
		2, 47, 7, 47, 2, 48, 7, 48, 2, 49, 7, 49, 2, 50, 7, 50, 2, 51, 7, 51, 2,
		52, 7, 52, 2, 53, 7, 53, 2, 54, 7, 54, 2, 55, 7, 55, 2, 56, 7, 56, 2, 57,
		7, 57, 2, 58, 7, 58, 2, 59, 7, 59, 2, 60, 7, 60, 2, 61, 7, 61, 2, 62, 7,
		62, 2, 63, 7, 63, 2, 64, 7, 64, 2, 65, 7, 65, 2, 66, 7, 66, 2, 67, 7, 67,
		2, 68, 7, 68, 2, 69, 7, 69, 2, 70, 7, 70, 2, 71, 7, 71, 2, 72, 7, 72, 2,
		73, 7, 73, 2, 74, 7, 74, 2, 75, 7, 75, 2, 76, 7, 76, 2, 77, 7, 77, 2, 78,
		7, 78, 2, 79, 7, 79, 2, 80, 7, 80, 2, 81, 7, 81, 2, 82, 7, 82, 2, 83, 7,
		83, 2, 84, 7, 84, 2, 85, 7, 85, 2, 86, 7, 86, 2, 87, 7, 87, 2, 88, 7, 88,
		2, 89, 7, 89, 2, 90, 7, 90, 2, 91, 7, 91, 2, 92, 7, 92, 2, 93, 7, 93, 2,
		94, 7, 94, 2, 95, 7, 95, 2, 96, 7, 96, 2, 97, 7, 97, 1, 0, 1, 0, 1, 0,
		1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4,
		1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7,
		1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9,
		1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1,
		12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 1, 13, 1, 14, 1, 14,
		1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 1, 15, 1, 16, 1,
		16, 1, 16, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 18,
		1, 18, 1, 18, 1, 18, 1, 18, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 20, 1,
		20, 1, 20, 1, 20, 1, 20, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21,
		1, 21, 1, 21, 1, 21, 1, 21, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1,
		22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22,
		1, 22, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 24, 1, 24, 1,
		24, 1, 24, 1, 25, 1, 25, 1, 25, 1, 25, 1, 26, 1, 26, 1, 26, 1, 26, 1, 27,
		1, 27, 1, 27, 1, 27, 1, 27, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1,
		29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 30, 1, 30, 1, 30,
		1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 31, 1, 31, 1, 31, 1,
		31, 1, 31, 1, 31, 1, 31, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32,
		1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 34, 1, 34, 1, 34, 1, 34, 1,
		34, 1, 34, 1, 34, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 1, 36, 1, 36, 1, 36,
		1, 36, 1, 36, 1, 36, 1, 36, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 38, 1,
		38, 1, 38, 1, 38, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 1, 40, 1, 40,
		1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 41, 1, 41, 1, 41, 1, 41, 1,
		41, 1, 41, 1, 41, 1, 41, 1, 41, 1, 41, 1, 42, 1, 42, 1, 42, 1, 42, 1, 42,
		1, 42, 1, 42, 1, 42, 1, 43, 1, 43, 1, 43, 5, 43, 465, 8, 43, 10, 43, 12,
		43, 468, 9, 43, 1, 44, 1, 44, 1, 45, 1, 45, 1, 45, 1, 46, 1, 46, 1, 46,
		1, 47, 1, 47, 1, 48, 1, 48, 1, 49, 1, 49, 1, 50, 1, 50, 1, 51, 1, 51, 1,
		52, 1, 52, 1, 53, 1, 53, 1, 54, 1, 54, 1, 55, 1, 55, 1, 56, 1, 56, 1, 57,
		1, 57, 1, 57, 1, 58, 1, 58, 1, 58, 1, 59, 1, 59, 1, 60, 1, 60, 1, 60, 1,
		61, 1, 61, 1, 62, 1, 62, 1, 62, 1, 63, 1, 63, 1, 63, 1, 64, 1, 64, 1, 65,
		1, 65, 1, 65, 1, 66, 1, 66, 1, 67, 1, 67, 1, 67, 1, 68, 1, 68, 1, 68, 1,
		69, 1, 69, 1, 70, 1, 70, 1, 71, 1, 71, 1, 72, 1, 72, 1, 73, 1, 73, 1, 73,
		1, 74, 1, 74, 1, 75, 1, 75, 1, 76, 1, 76, 1, 77, 1, 77, 1, 77, 1, 78, 1,
		78, 1, 78, 1, 79, 1, 79, 1, 79, 1, 80, 1, 80, 5, 80, 558, 8, 80, 10, 80,
		12, 80, 561, 9, 80, 1, 81, 1, 81, 5, 81, 565, 8, 81, 10, 81, 12, 81, 568,
		9, 81, 1, 82, 1, 82, 1, 82, 4, 82, 573, 8, 82, 11, 82, 12, 82, 574, 1,
		83, 1, 83, 1, 83, 3, 83, 580, 8, 83, 1, 83, 3, 83, 583, 8, 83, 1, 83, 3,
		83, 586, 8, 83, 1, 83, 1, 83, 1, 83, 3, 83, 591, 8, 83, 3, 83, 593, 8,
		83, 1, 84, 1, 84, 5, 84, 597, 8, 84, 10, 84, 12, 84, 600, 9, 84, 1, 84,
		1, 84, 1, 85, 1, 85, 1, 85, 5, 85, 607, 8, 85, 10, 85, 12, 85, 610, 9,
		85, 1, 85, 1, 85, 1, 86, 4, 86, 615, 8, 86, 11, 86, 12, 86, 616, 1, 86,
		1, 86, 1, 87, 1, 87, 1, 87, 1, 87, 5, 87, 625, 8, 87, 10, 87, 12, 87, 628,
		9, 87, 1, 87, 1, 87, 1, 87, 1, 87, 1, 87, 1, 88, 4, 88, 636, 8, 88, 11,
		88, 12, 88, 637, 1, 88, 1, 88, 1, 89, 1, 89, 1, 89, 1, 89, 5, 89, 646,
		8, 89, 10, 89, 12, 89, 649, 9, 89, 1, 89, 1, 89, 1, 90, 1, 90, 1, 90, 1,
		90, 1, 90, 1, 90, 1, 90, 1, 90, 1, 90, 1, 90, 1, 90, 1, 90, 1, 90, 1, 90,
		1, 90, 1, 90, 1, 90, 1, 90, 1, 90, 1, 90, 1, 90, 1, 90, 1, 90, 1, 90, 1,
		90, 1, 90, 3, 90, 679, 8, 90, 1, 91, 4, 91, 682, 8, 91, 11, 91, 12, 91,
		683, 1, 92, 1, 92, 1, 93, 1, 93, 1, 94, 1, 94, 3, 94, 692, 8, 94, 1, 94,
		1, 94, 1, 95, 1, 95, 3, 95, 698, 8, 95, 1, 96, 3, 96, 701, 8, 96, 1, 97,
		3, 97, 704, 8, 97, 1, 626, 0, 98, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6,
		13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 31,
		16, 33, 17, 35, 18, 37, 19, 39, 20, 41, 21, 43, 22, 45, 23, 47, 24, 49,
		25, 51, 26, 53, 27, 55, 28, 57, 29, 59, 30, 61, 31, 63, 32, 65, 33, 67,
		34, 69, 35, 71, 36, 73, 37, 75, 38, 77, 39, 79, 40, 81, 41, 83, 42, 85,
		43, 87, 44, 89, 45, 91, 46, 93, 47, 95, 48, 97, 49, 99, 50, 101, 51, 103,
		52, 105, 53, 107, 54, 109, 55, 111, 56, 113, 57, 115, 58, 117, 59, 119,
		60, 121, 61, 123, 62, 125, 63, 127, 64, 129, 65, 131, 66, 133, 67, 135,
		68, 137, 69, 139, 70, 141, 71, 143, 72, 145, 73, 147, 74, 149, 75, 151,
		76, 153, 77, 155, 78, 157, 79, 159, 80, 161, 81, 163, 82, 165, 83, 167,
		84, 169, 85, 171, 86, 173, 87, 175, 88, 177, 89, 179, 90, 181, 0, 183,
		0, 185, 0, 187, 0, 189, 0, 191, 0, 193, 0, 195, 0, 1, 0, 14, 1, 0, 49,
		57, 1, 0, 48, 57, 2, 0, 88, 88, 120, 120, 1, 0, 96, 96, 2, 0, 34, 34, 92,
		92, 2, 0, 9, 9, 32, 32, 2, 0, 10, 10, 13, 13, 9, 0, 34, 34, 39, 39, 92,
		92, 97, 98, 102, 102, 110, 110, 114, 114, 116, 116, 118, 118, 1, 0, 48,
		55, 3, 0, 48, 57, 65, 70, 97, 102, 2, 0, 69, 69, 101, 101, 2, 0, 43, 43,
		45, 45, 20, 0, 48, 57, 1632, 1641, 1776, 1785, 2406, 2415, 2534, 2543,
		2662, 2671, 2790, 2799, 2918, 2927, 3047, 3055, 3174, 3183, 3302, 3311,
		3430, 3439, 3664, 3673, 3792, 3801, 3872, 3881, 4160, 4169, 4969, 4977,
		6112, 6121, 6160, 6169, 65296, 65305, 258, 0, 65, 90, 97, 122, 170, 170,
		181, 181, 186, 186, 192, 214, 216, 246, 248, 543, 546, 563, 592, 685, 688,
		696, 699, 705, 720, 721, 736, 740, 750, 750, 890, 890, 902, 902, 904, 906,
		908, 908, 910, 929, 931, 974, 976, 983, 986, 1011, 1024, 1153, 1164, 1220,
		1223, 1224, 1227, 1228, 1232, 1269, 1272, 1273, 1329, 1366, 1369, 1369,
		1377, 1415, 1488, 1514, 1520, 1522, 1569, 1594, 1600, 1610, 1649, 1747,
		1749, 1749, 1765, 1766, 1786, 1788, 1808, 1808, 1810, 1836, 1920, 1957,
		2309, 2361, 2365, 2365, 2384, 2384, 2392, 2401, 2437, 2444, 2447, 2448,
		2451, 2472, 2474, 2480, 2482, 2482, 2486, 2489, 2524, 2525, 2527, 2529,
		2544, 2545, 2565, 2570, 2575, 2576, 2579, 2600, 2602, 2608, 2610, 2611,
		2613, 2614, 2616, 2617, 2649, 2652, 2654, 2654, 2674, 2676, 2693, 2699,
		2701, 2701, 2703, 2705, 2707, 2728, 2730, 2736, 2738, 2739, 2741, 2745,
		2749, 2749, 2768, 2768, 2784, 2784, 2821, 2828, 2831, 2832, 2835, 2856,
		2858, 2864, 2866, 2867, 2870, 2873, 2877, 2877, 2908, 2909, 2911, 2913,
		2949, 2954, 2958, 2960, 2962, 2965, 2969, 2970, 2972, 2972, 2974, 2975,
		2979, 2980, 2984, 2986, 2990, 2997, 2999, 3001, 3077, 3084, 3086, 3088,
		3090, 3112, 3114, 3123, 3125, 3129, 3168, 3169, 3205, 3212, 3214, 3216,
		3218, 3240, 3242, 3251, 3253, 3257, 3294, 3294, 3296, 3297, 3333, 3340,
		3342, 3344, 3346, 3368, 3370, 3385, 3424, 3425, 3461, 3478, 3482, 3505,
		3507, 3515, 3517, 3517, 3520, 3526, 3585, 3632, 3634, 3635, 3648, 3654,
		3713, 3714, 3716, 3716, 3719, 3720, 3722, 3722, 3725, 3725, 3732, 3735,
		3737, 3743, 3745, 3747, 3749, 3749, 3751, 3751, 3754, 3755, 3757, 3760,
		3762, 3763, 3773, 3780, 3782, 3782, 3804, 3805, 3840, 3840, 3904, 3946,
		3976, 3979, 4096, 4129, 4131, 4135, 4137, 4138, 4176, 4181, 4256, 4293,
		4304, 4342, 4352, 4441, 4447, 4514, 4520, 4601, 4608, 4614, 4616, 4678,
		4680, 4680, 4682, 4685, 4688, 4694, 4696, 4696, 4698, 4701, 4704, 4742,
		4744, 4744, 4746, 4749, 4752, 4782, 4784, 4784, 4786, 4789, 4792, 4798,
		4800, 4800, 4802, 4805, 4808, 4814, 4816, 4822, 4824, 4846, 4848, 4878,
		4880, 4880, 4882, 4885, 4888, 4894, 4896, 4934, 4936, 4954, 5024, 5108,
		5121, 5750, 5761, 5786, 5792, 5866, 6016, 6067, 6176, 6263, 6272, 6312,
		7680, 7835, 7840, 7929, 7936, 7957, 7960, 7965, 7968, 8005, 8008, 8013,
		8016, 8023, 8025, 8025, 8027, 8027, 8029, 8029, 8031, 8061, 8064, 8116,
		8118, 8124, 8126, 8126, 8130, 8132, 8134, 8140, 8144, 8147, 8150, 8155,
		8160, 8172, 8178, 8180, 8182, 8188, 8319, 8319, 8450, 8450, 8455, 8455,
		8458, 8467, 8469, 8469, 8473, 8477, 8484, 8484, 8486, 8486, 8488, 8488,
		8490, 8493, 8495, 8497, 8499, 8505, 8544, 8579, 12293, 12295, 12321, 12329,
		12337, 12341, 12344, 12346, 12353, 12436, 12445, 12446, 12449, 12538, 12540,
		12542, 12549, 12588, 12593, 12686, 12704, 12727, 13312, 13312, 19893, 19893,
		19968, 19968, 40869, 40869, 40960, 42124, 44032, 44032, 55203, 55203, 63744,
		64045, 64256, 64262, 64275, 64279, 64285, 64285, 64287, 64296, 64298, 64310,
		64312, 64316, 64318, 64318, 64320, 64321, 64323, 64324, 64326, 64433, 64467,
		64829, 64848, 64911, 64914, 64967, 65008, 65019, 65136, 65138, 65140, 65140,
		65142, 65276, 65313, 65338, 65345, 65370, 65382, 65470, 65474, 65479, 65482,
		65487, 65490, 65495, 65498, 65500, 720, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0,
		0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0,
		0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0,
		0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1,
		0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35,
		1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0,
		43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0, 47, 1, 0, 0, 0, 0, 49, 1, 0, 0, 0,
		0, 51, 1, 0, 0, 0, 0, 53, 1, 0, 0, 0, 0, 55, 1, 0, 0, 0, 0, 57, 1, 0, 0,
		0, 0, 59, 1, 0, 0, 0, 0, 61, 1, 0, 0, 0, 0, 63, 1, 0, 0, 0, 0, 65, 1, 0,
		0, 0, 0, 67, 1, 0, 0, 0, 0, 69, 1, 0, 0, 0, 0, 71, 1, 0, 0, 0, 0, 73, 1,
		0, 0, 0, 0, 75, 1, 0, 0, 0, 0, 77, 1, 0, 0, 0, 0, 79, 1, 0, 0, 0, 0, 81,
		1, 0, 0, 0, 0, 83, 1, 0, 0, 0, 0, 85, 1, 0, 0, 0, 0, 87, 1, 0, 0, 0, 0,
		89, 1, 0, 0, 0, 0, 91, 1, 0, 0, 0, 0, 93, 1, 0, 0, 0, 0, 95, 1, 0, 0, 0,
		0, 97, 1, 0, 0, 0, 0, 99, 1, 0, 0, 0, 0, 101, 1, 0, 0, 0, 0, 103, 1, 0,
		0, 0, 0, 105, 1, 0, 0, 0, 0, 107, 1, 0, 0, 0, 0, 109, 1, 0, 0, 0, 0, 111,
		1, 0, 0, 0, 0, 113, 1, 0, 0, 0, 0, 115, 1, 0, 0, 0, 0, 117, 1, 0, 0, 0,
		0, 119, 1, 0, 0, 0, 0, 121, 1, 0, 0, 0, 0, 123, 1, 0, 0, 0, 0, 125, 1,
		0, 0, 0, 0, 127, 1, 0, 0, 0, 0, 129, 1, 0, 0, 0, 0, 131, 1, 0, 0, 0, 0,
		133, 1, 0, 0, 0, 0, 135, 1, 0, 0, 0, 0, 137, 1, 0, 0, 0, 0, 139, 1, 0,
		0, 0, 0, 141, 1, 0, 0, 0, 0, 143, 1, 0, 0, 0, 0, 145, 1, 0, 0, 0, 0, 147,
		1, 0, 0, 0, 0, 149, 1, 0, 0, 0, 0, 151, 1, 0, 0, 0, 0, 153, 1, 0, 0, 0,
		0, 155, 1, 0, 0, 0, 0, 157, 1, 0, 0, 0, 0, 159, 1, 0, 0, 0, 0, 161, 1,
		0, 0, 0, 0, 163, 1, 0, 0, 0, 0, 165, 1, 0, 0, 0, 0, 167, 1, 0, 0, 0, 0,
		169, 1, 0, 0, 0, 0, 171, 1, 0, 0, 0, 0, 173, 1, 0, 0, 0, 0, 175, 1, 0,
		0, 0, 0, 177, 1, 0, 0, 0, 0, 179, 1, 0, 0, 0, 1, 197, 1, 0, 0, 0, 3, 201,
		1, 0, 0, 0, 5, 208, 1, 0, 0, 0, 7, 215, 1, 0, 0, 0, 9, 219, 1, 0, 0, 0,
		11, 225, 1, 0, 0, 0, 13, 229, 1, 0, 0, 0, 15, 234, 1, 0, 0, 0, 17, 239,
		1, 0, 0, 0, 19, 243, 1, 0, 0, 0, 21, 248, 1, 0, 0, 0, 23, 251, 1, 0, 0,
		0, 25, 258, 1, 0, 0, 0, 27, 263, 1, 0, 0, 0, 29, 267, 1, 0, 0, 0, 31, 274,
		1, 0, 0, 0, 33, 278, 1, 0, 0, 0, 35, 283, 1, 0, 0, 0, 37, 289, 1, 0, 0,
		0, 39, 294, 1, 0, 0, 0, 41, 299, 1, 0, 0, 0, 43, 304, 1, 0, 0, 0, 45, 315,
		1, 0, 0, 0, 47, 333, 1, 0, 0, 0, 49, 340, 1, 0, 0, 0, 51, 344, 1, 0, 0,
		0, 53, 348, 1, 0, 0, 0, 55, 352, 1, 0, 0, 0, 57, 357, 1, 0, 0, 0, 59, 363,
		1, 0, 0, 0, 61, 371, 1, 0, 0, 0, 63, 381, 1, 0, 0, 0, 65, 388, 1, 0, 0,
		0, 67, 395, 1, 0, 0, 0, 69, 401, 1, 0, 0, 0, 71, 408, 1, 0, 0, 0, 73, 413,
		1, 0, 0, 0, 75, 420, 1, 0, 0, 0, 77, 425, 1, 0, 0, 0, 79, 429, 1, 0, 0,
		0, 81, 435, 1, 0, 0, 0, 83, 443, 1, 0, 0, 0, 85, 453, 1, 0, 0, 0, 87, 461,
		1, 0, 0, 0, 89, 469, 1, 0, 0, 0, 91, 471, 1, 0, 0, 0, 93, 474, 1, 0, 0,
		0, 95, 477, 1, 0, 0, 0, 97, 479, 1, 0, 0, 0, 99, 481, 1, 0, 0, 0, 101,
		483, 1, 0, 0, 0, 103, 485, 1, 0, 0, 0, 105, 487, 1, 0, 0, 0, 107, 489,
		1, 0, 0, 0, 109, 491, 1, 0, 0, 0, 111, 493, 1, 0, 0, 0, 113, 495, 1, 0,
		0, 0, 115, 497, 1, 0, 0, 0, 117, 500, 1, 0, 0, 0, 119, 503, 1, 0, 0, 0,
		121, 505, 1, 0, 0, 0, 123, 508, 1, 0, 0, 0, 125, 510, 1, 0, 0, 0, 127,
		513, 1, 0, 0, 0, 129, 516, 1, 0, 0, 0, 131, 518, 1, 0, 0, 0, 133, 521,
		1, 0, 0, 0, 135, 523, 1, 0, 0, 0, 137, 526, 1, 0, 0, 0, 139, 529, 1, 0,
		0, 0, 141, 531, 1, 0, 0, 0, 143, 533, 1, 0, 0, 0, 145, 535, 1, 0, 0, 0,
		147, 537, 1, 0, 0, 0, 149, 540, 1, 0, 0, 0, 151, 542, 1, 0, 0, 0, 153,
		544, 1, 0, 0, 0, 155, 546, 1, 0, 0, 0, 157, 549, 1, 0, 0, 0, 159, 552,
		1, 0, 0, 0, 161, 555, 1, 0, 0, 0, 163, 562, 1, 0, 0, 0, 165, 569, 1, 0,
		0, 0, 167, 592, 1, 0, 0, 0, 169, 594, 1, 0, 0, 0, 171, 603, 1, 0, 0, 0,
		173, 614, 1, 0, 0, 0, 175, 620, 1, 0, 0, 0, 177, 635, 1, 0, 0, 0, 179,
		641, 1, 0, 0, 0, 181, 652, 1, 0, 0, 0, 183, 681, 1, 0, 0, 0, 185, 685,
		1, 0, 0, 0, 187, 687, 1, 0, 0, 0, 189, 689, 1, 0, 0, 0, 191, 697, 1, 0,
		0, 0, 193, 700, 1, 0, 0, 0, 195, 703, 1, 0, 0, 0, 197, 198, 5, 97, 0, 0,
		198, 199, 5, 108, 0, 0, 199, 200, 5, 108, 0, 0, 200, 2, 1, 0, 0, 0, 201,
		202, 5, 97, 0, 0, 202, 203, 5, 115, 0, 0, 203, 204, 5, 115, 0, 0, 204,
		205, 5, 101, 0, 0, 205, 206, 5, 114, 0, 0, 206, 207, 5, 116, 0, 0, 207,
		4, 1, 0, 0, 0, 208, 209, 5, 97, 0, 0, 209, 210, 5, 115, 0, 0, 210, 211,
		5, 115, 0, 0, 211, 212, 5, 117, 0, 0, 212, 213, 5, 109, 0, 0, 213, 214,
		5, 101, 0, 0, 214, 6, 1, 0, 0, 0, 215, 216, 5, 110, 0, 0, 216, 217, 5,
		111, 0, 0, 217, 218, 5, 119, 0, 0, 218, 8, 1, 0, 0, 0, 219, 220, 5, 99,
		0, 0, 220, 221, 5, 111, 0, 0, 221, 222, 5, 110, 0, 0, 222, 223, 5, 115,
		0, 0, 223, 224, 5, 116, 0, 0, 224, 10, 1, 0, 0, 0, 225, 226, 5, 100, 0,
		0, 226, 227, 5, 101, 0, 0, 227, 228, 5, 102, 0, 0, 228, 12, 1, 0, 0, 0,
		229, 230, 5, 101, 0, 0, 230, 231, 5, 108, 0, 0, 231, 232, 5, 115, 0, 0,
		232, 233, 5, 101, 0, 0, 233, 14, 1, 0, 0, 0, 234, 235, 5, 102, 0, 0, 235,
		236, 5, 108, 0, 0, 236, 237, 5, 111, 0, 0, 237, 238, 5, 119, 0, 0, 238,
		16, 1, 0, 0, 0, 239, 240, 5, 102, 0, 0, 240, 241, 5, 111, 0, 0, 241, 242,
		5, 114, 0, 0, 242, 18, 1, 0, 0, 0, 243, 244, 5, 102, 0, 0, 244, 245, 5,
		117, 0, 0, 245, 246, 5, 110, 0, 0, 246, 247, 5, 99, 0, 0, 247, 20, 1, 0,
		0, 0, 248, 249, 5, 105, 0, 0, 249, 250, 5, 102, 0, 0, 250, 22, 1, 0, 0,
		0, 251, 252, 5, 105, 0, 0, 252, 253, 5, 109, 0, 0, 253, 254, 5, 112, 0,
		0, 254, 255, 5, 111, 0, 0, 255, 256, 5, 114, 0, 0, 256, 257, 5, 116, 0,
		0, 257, 24, 1, 0, 0, 0, 258, 259, 5, 105, 0, 0, 259, 260, 5, 110, 0, 0,
		260, 261, 5, 105, 0, 0, 261, 262, 5, 116, 0, 0, 262, 26, 1, 0, 0, 0, 263,
		264, 5, 110, 0, 0, 264, 265, 5, 101, 0, 0, 265, 266, 5, 119, 0, 0, 266,
		28, 1, 0, 0, 0, 267, 268, 5, 114, 0, 0, 268, 269, 5, 101, 0, 0, 269, 270,
		5, 116, 0, 0, 270, 271, 5, 117, 0, 0, 271, 272, 5, 114, 0, 0, 272, 273,
		5, 110, 0, 0, 273, 30, 1, 0, 0, 0, 274, 275, 5, 114, 0, 0, 275, 276, 5,
		117, 0, 0, 276, 277, 5, 110, 0, 0, 277, 32, 1, 0, 0, 0, 278, 279, 5, 115,
		0, 0, 279, 280, 5, 112, 0, 0, 280, 281, 5, 101, 0, 0, 281, 282, 5, 99,
		0, 0, 282, 34, 1, 0, 0, 0, 283, 284, 5, 115, 0, 0, 284, 285, 5, 116, 0,
		0, 285, 286, 5, 111, 0, 0, 286, 287, 5, 99, 0, 0, 287, 288, 5, 107, 0,
		0, 288, 36, 1, 0, 0, 0, 289, 290, 5, 116, 0, 0, 290, 291, 5, 104, 0, 0,
		291, 292, 5, 101, 0, 0, 292, 293, 5, 110, 0, 0, 293, 38, 1, 0, 0, 0, 294,
		295, 5, 119, 0, 0, 295, 296, 5, 104, 0, 0, 296, 297, 5, 101, 0, 0, 297,
		298, 5, 110, 0, 0, 298, 40, 1, 0, 0, 0, 299, 300, 5, 116, 0, 0, 300, 301,
		5, 104, 0, 0, 301, 302, 5, 105, 0, 0, 302, 303, 5, 115, 0, 0, 303, 42,
		1, 0, 0, 0, 304, 305, 5, 101, 0, 0, 305, 306, 5, 118, 0, 0, 306, 307, 5,
		101, 0, 0, 307, 308, 5, 110, 0, 0, 308, 309, 5, 116, 0, 0, 309, 310, 5,
		117, 0, 0, 310, 311, 5, 97, 0, 0, 311, 312, 5, 108, 0, 0, 312, 313, 5,
		108, 0, 0, 313, 314, 5, 121, 0, 0, 314, 44, 1, 0, 0, 0, 315, 316, 5, 101,
		0, 0, 316, 317, 5, 118, 0, 0, 317, 318, 5, 101, 0, 0, 318, 319, 5, 110,
		0, 0, 319, 320, 5, 116, 0, 0, 320, 321, 5, 117, 0, 0, 321, 322, 5, 97,
		0, 0, 322, 323, 5, 108, 0, 0, 323, 324, 5, 108, 0, 0, 324, 325, 5, 121,
		0, 0, 325, 326, 5, 45, 0, 0, 326, 327, 5, 97, 0, 0, 327, 328, 5, 108, 0,
		0, 328, 329, 5, 119, 0, 0, 329, 330, 5, 97, 0, 0, 330, 331, 5, 121, 0,
		0, 331, 332, 5, 115, 0, 0, 332, 46, 1, 0, 0, 0, 333, 334, 5, 97, 0, 0,
		334, 335, 5, 108, 0, 0, 335, 336, 5, 119, 0, 0, 336, 337, 5, 97, 0, 0,
		337, 338, 5, 121, 0, 0, 338, 339, 5, 115, 0, 0, 339, 48, 1, 0, 0, 0, 340,
		341, 5, 110, 0, 0, 341, 342, 5, 109, 0, 0, 342, 343, 5, 116, 0, 0, 343,
		50, 1, 0, 0, 0, 344, 345, 5, 110, 0, 0, 345, 346, 5, 102, 0, 0, 346, 347,
		5, 116, 0, 0, 347, 52, 1, 0, 0, 0, 348, 349, 5, 110, 0, 0, 349, 350, 5,
		105, 0, 0, 350, 351, 5, 108, 0, 0, 351, 54, 1, 0, 0, 0, 352, 353, 5, 116,
		0, 0, 353, 354, 5, 114, 0, 0, 354, 355, 5, 117, 0, 0, 355, 356, 5, 101,
		0, 0, 356, 56, 1, 0, 0, 0, 357, 358, 5, 102, 0, 0, 358, 359, 5, 97, 0,
		0, 359, 360, 5, 108, 0, 0, 360, 361, 5, 115, 0, 0, 361, 362, 5, 101, 0,
		0, 362, 58, 1, 0, 0, 0, 363, 364, 5, 97, 0, 0, 364, 365, 5, 100, 0, 0,
		365, 366, 5, 118, 0, 0, 366, 367, 5, 97, 0, 0, 367, 368, 5, 110, 0, 0,
		368, 369, 5, 99, 0, 0, 369, 370, 5, 101, 0, 0, 370, 60, 1, 0, 0, 0, 371,
		372, 5, 99, 0, 0, 372, 373, 5, 111, 0, 0, 373, 374, 5, 109, 0, 0, 374,
		375, 5, 112, 0, 0, 375, 376, 5, 111, 0, 0, 376, 377, 5, 110, 0, 0, 377,
		378, 5, 101, 0, 0, 378, 379, 5, 110, 0, 0, 379, 380, 5, 116, 0, 0, 380,
		62, 1, 0, 0, 0, 381, 382, 5, 103, 0, 0, 382, 383, 5, 108, 0, 0, 383, 384,
		5, 111, 0, 0, 384, 385, 5, 98, 0, 0, 385, 386, 5, 97, 0, 0, 386, 387, 5,
		108, 0, 0, 387, 64, 1, 0, 0, 0, 388, 389, 5, 115, 0, 0, 389, 390, 5, 121,
		0, 0, 390, 391, 5, 115, 0, 0, 391, 392, 5, 116, 0, 0, 392, 393, 5, 101,
		0, 0, 393, 394, 5, 109, 0, 0, 394, 66, 1, 0, 0, 0, 395, 396, 5, 115, 0,
		0, 396, 397, 5, 116, 0, 0, 397, 398, 5, 97, 0, 0, 398, 399, 5, 114, 0,
		0, 399, 400, 5, 116, 0, 0, 400, 68, 1, 0, 0, 0, 401, 402, 5, 115, 0, 0,
		402, 403, 5, 116, 0, 0, 403, 404, 5, 97, 0, 0, 404, 405, 5, 116, 0, 0,
		405, 406, 5, 101, 0, 0, 406, 407, 5, 115, 0, 0, 407, 70, 1, 0, 0, 0, 408,
		409, 5, 115, 0, 0, 409, 410, 5, 116, 0, 0, 410, 411, 5, 97, 0, 0, 411,
		412, 5, 121, 0, 0, 412, 72, 1, 0, 0, 0, 413, 414, 5, 115, 0, 0, 414, 415,
		5, 116, 0, 0, 415, 416, 5, 114, 0, 0, 416, 417, 5, 105, 0, 0, 417, 418,
		5, 110, 0, 0, 418, 419, 5, 103, 0, 0, 419, 74, 1, 0, 0, 0, 420, 421, 5,
		98, 0, 0, 421, 422, 5, 111, 0, 0, 422, 423, 5, 111, 0, 0, 423, 424, 5,
		108, 0, 0, 424, 76, 1, 0, 0, 0, 425, 426, 5, 105, 0, 0, 426, 427, 5, 110,
		0, 0, 427, 428, 5, 116, 0, 0, 428, 78, 1, 0, 0, 0, 429, 430, 5, 102, 0,
		0, 430, 431, 5, 108, 0, 0, 431, 432, 5, 111, 0, 0, 432, 433, 5, 97, 0,
		0, 433, 434, 5, 116, 0, 0, 434, 80, 1, 0, 0, 0, 435, 436, 5, 110, 0, 0,
		436, 437, 5, 97, 0, 0, 437, 438, 5, 116, 0, 0, 438, 439, 5, 117, 0, 0,
		439, 440, 5, 114, 0, 0, 440, 441, 5, 97, 0, 0, 441, 442, 5, 108, 0, 0,
		442, 82, 1, 0, 0, 0, 443, 444, 5, 117, 0, 0, 444, 445, 5, 110, 0, 0, 445,
		446, 5, 99, 0, 0, 446, 447, 5, 101, 0, 0, 447, 448, 5, 114, 0, 0, 448,
		449, 5, 116, 0, 0, 449, 450, 5, 97, 0, 0, 450, 451, 5, 105, 0, 0, 451,
		452, 5, 110, 0, 0, 452, 84, 1, 0, 0, 0, 453, 454, 5, 117, 0, 0, 454, 455,
		5, 110, 0, 0, 455, 456, 5, 107, 0, 0, 456, 457, 5, 110, 0, 0, 457, 458,
		5, 111, 0, 0, 458, 459, 5, 119, 0, 0, 459, 460, 5, 110, 0, 0, 460, 86,
		1, 0, 0, 0, 461, 466, 3, 191, 95, 0, 462, 465, 3, 191, 95, 0, 463, 465,
		3, 193, 96, 0, 464, 462, 1, 0, 0, 0, 464, 463, 1, 0, 0, 0, 465, 468, 1,
		0, 0, 0, 466, 464, 1, 0, 0, 0, 466, 467, 1, 0, 0, 0, 467, 88, 1, 0, 0,
		0, 468, 466, 1, 0, 0, 0, 469, 470, 5, 61, 0, 0, 470, 90, 1, 0, 0, 0, 471,
		472, 5, 45, 0, 0, 472, 473, 5, 62, 0, 0, 473, 92, 1, 0, 0, 0, 474, 475,
		5, 60, 0, 0, 475, 476, 5, 45, 0, 0, 476, 94, 1, 0, 0, 0, 477, 478, 5, 58,
		0, 0, 478, 96, 1, 0, 0, 0, 479, 480, 5, 44, 0, 0, 480, 98, 1, 0, 0, 0,
		481, 482, 5, 46, 0, 0, 482, 100, 1, 0, 0, 0, 483, 484, 5, 40, 0, 0, 484,
		102, 1, 0, 0, 0, 485, 486, 5, 41, 0, 0, 486, 104, 1, 0, 0, 0, 487, 488,
		5, 123, 0, 0, 488, 106, 1, 0, 0, 0, 489, 490, 5, 125, 0, 0, 490, 108, 1,
		0, 0, 0, 491, 492, 5, 91, 0, 0, 492, 110, 1, 0, 0, 0, 493, 494, 5, 93,
		0, 0, 494, 112, 1, 0, 0, 0, 495, 496, 5, 59, 0, 0, 496, 114, 1, 0, 0, 0,
		497, 498, 5, 43, 0, 0, 498, 499, 5, 43, 0, 0, 499, 116, 1, 0, 0, 0, 500,
		501, 5, 45, 0, 0, 501, 502, 5, 45, 0, 0, 502, 118, 1, 0, 0, 0, 503, 504,
		5, 38, 0, 0, 504, 120, 1, 0, 0, 0, 505, 506, 5, 38, 0, 0, 506, 507, 5,
		38, 0, 0, 507, 122, 1, 0, 0, 0, 508, 509, 5, 33, 0, 0, 509, 124, 1, 0,
		0, 0, 510, 511, 5, 61, 0, 0, 511, 512, 5, 61, 0, 0, 512, 126, 1, 0, 0,
		0, 513, 514, 5, 33, 0, 0, 514, 515, 5, 61, 0, 0, 515, 128, 1, 0, 0, 0,
		516, 517, 5, 60, 0, 0, 517, 130, 1, 0, 0, 0, 518, 519, 5, 60, 0, 0, 519,
		520, 5, 61, 0, 0, 520, 132, 1, 0, 0, 0, 521, 522, 5, 62, 0, 0, 522, 134,
		1, 0, 0, 0, 523, 524, 5, 62, 0, 0, 524, 525, 5, 61, 0, 0, 525, 136, 1,
		0, 0, 0, 526, 527, 5, 124, 0, 0, 527, 528, 5, 124, 0, 0, 528, 138, 1, 0,
		0, 0, 529, 530, 5, 124, 0, 0, 530, 140, 1, 0, 0, 0, 531, 532, 5, 43, 0,
		0, 532, 142, 1, 0, 0, 0, 533, 534, 5, 45, 0, 0, 534, 144, 1, 0, 0, 0, 535,
		536, 5, 94, 0, 0, 536, 146, 1, 0, 0, 0, 537, 538, 5, 42, 0, 0, 538, 539,
		5, 42, 0, 0, 539, 148, 1, 0, 0, 0, 540, 541, 5, 42, 0, 0, 541, 150, 1,
		0, 0, 0, 542, 543, 5, 47, 0, 0, 543, 152, 1, 0, 0, 0, 544, 545, 5, 37,
		0, 0, 545, 154, 1, 0, 0, 0, 546, 547, 5, 60, 0, 0, 547, 548, 5, 60, 0,
		0, 548, 156, 1, 0, 0, 0, 549, 550, 5, 62, 0, 0, 550, 551, 5, 62, 0, 0,
		551, 158, 1, 0, 0, 0, 552, 553, 5, 38, 0, 0, 553, 554, 5, 94, 0, 0, 554,
		160, 1, 0, 0, 0, 555, 559, 7, 0, 0, 0, 556, 558, 7, 1, 0, 0, 557, 556,
		1, 0, 0, 0, 558, 561, 1, 0, 0, 0, 559, 557, 1, 0, 0, 0, 559, 560, 1, 0,
		0, 0, 560, 162, 1, 0, 0, 0, 561, 559, 1, 0, 0, 0, 562, 566, 5, 48, 0, 0,
		563, 565, 3, 185, 92, 0, 564, 563, 1, 0, 0, 0, 565, 568, 1, 0, 0, 0, 566,
		564, 1, 0, 0, 0, 566, 567, 1, 0, 0, 0, 567, 164, 1, 0, 0, 0, 568, 566,
		1, 0, 0, 0, 569, 570, 5, 48, 0, 0, 570, 572, 7, 2, 0, 0, 571, 573, 3, 187,
		93, 0, 572, 571, 1, 0, 0, 0, 573, 574, 1, 0, 0, 0, 574, 572, 1, 0, 0, 0,
		574, 575, 1, 0, 0, 0, 575, 166, 1, 0, 0, 0, 576, 585, 3, 183, 91, 0, 577,
		579, 5, 46, 0, 0, 578, 580, 3, 183, 91, 0, 579, 578, 1, 0, 0, 0, 579, 580,
		1, 0, 0, 0, 580, 582, 1, 0, 0, 0, 581, 583, 3, 189, 94, 0, 582, 581, 1,
		0, 0, 0, 582, 583, 1, 0, 0, 0, 583, 586, 1, 0, 0, 0, 584, 586, 3, 189,
		94, 0, 585, 577, 1, 0, 0, 0, 585, 584, 1, 0, 0, 0, 586, 593, 1, 0, 0, 0,
		587, 588, 5, 46, 0, 0, 588, 590, 3, 183, 91, 0, 589, 591, 3, 189, 94, 0,
		590, 589, 1, 0, 0, 0, 590, 591, 1, 0, 0, 0, 591, 593, 1, 0, 0, 0, 592,
		576, 1, 0, 0, 0, 592, 587, 1, 0, 0, 0, 593, 168, 1, 0, 0, 0, 594, 598,
		5, 96, 0, 0, 595, 597, 8, 3, 0, 0, 596, 595, 1, 0, 0, 0, 597, 600, 1, 0,
		0, 0, 598, 596, 1, 0, 0, 0, 598, 599, 1, 0, 0, 0, 599, 601, 1, 0, 0, 0,
		600, 598, 1, 0, 0, 0, 601, 602, 5, 96, 0, 0, 602, 170, 1, 0, 0, 0, 603,
		608, 5, 34, 0, 0, 604, 607, 8, 4, 0, 0, 605, 607, 3, 181, 90, 0, 606, 604,
		1, 0, 0, 0, 606, 605, 1, 0, 0, 0, 607, 610, 1, 0, 0, 0, 608, 606, 1, 0,
		0, 0, 608, 609, 1, 0, 0, 0, 609, 611, 1, 0, 0, 0, 610, 608, 1, 0, 0, 0,
		611, 612, 5, 34, 0, 0, 612, 172, 1, 0, 0, 0, 613, 615, 7, 5, 0, 0, 614,
		613, 1, 0, 0, 0, 615, 616, 1, 0, 0, 0, 616, 614, 1, 0, 0, 0, 616, 617,
		1, 0, 0, 0, 617, 618, 1, 0, 0, 0, 618, 619, 6, 86, 0, 0, 619, 174, 1, 0,
		0, 0, 620, 621, 5, 47, 0, 0, 621, 622, 5, 42, 0, 0, 622, 626, 1, 0, 0,
		0, 623, 625, 9, 0, 0, 0, 624, 623, 1, 0, 0, 0, 625, 628, 1, 0, 0, 0, 626,
		627, 1, 0, 0, 0, 626, 624, 1, 0, 0, 0, 627, 629, 1, 0, 0, 0, 628, 626,
		1, 0, 0, 0, 629, 630, 5, 42, 0, 0, 630, 631, 5, 47, 0, 0, 631, 632, 1,
		0, 0, 0, 632, 633, 6, 87, 1, 0, 633, 176, 1, 0, 0, 0, 634, 636, 7, 6, 0,
		0, 635, 634, 1, 0, 0, 0, 636, 637, 1, 0, 0, 0, 637, 635, 1, 0, 0, 0, 637,
		638, 1, 0, 0, 0, 638, 639, 1, 0, 0, 0, 639, 640, 6, 88, 1, 0, 640, 178,
		1, 0, 0, 0, 641, 642, 5, 47, 0, 0, 642, 643, 5, 47, 0, 0, 643, 647, 1,
		0, 0, 0, 644, 646, 8, 6, 0, 0, 645, 644, 1, 0, 0, 0, 646, 649, 1, 0, 0,
		0, 647, 645, 1, 0, 0, 0, 647, 648, 1, 0, 0, 0, 648, 650, 1, 0, 0, 0, 649,
		647, 1, 0, 0, 0, 650, 651, 6, 89, 1, 0, 651, 180, 1, 0, 0, 0, 652, 678,
		5, 92, 0, 0, 653, 654, 5, 117, 0, 0, 654, 655, 3, 187, 93, 0, 655, 656,
		3, 187, 93, 0, 656, 657, 3, 187, 93, 0, 657, 658, 3, 187, 93, 0, 658, 679,
		1, 0, 0, 0, 659, 660, 5, 85, 0, 0, 660, 661, 3, 187, 93, 0, 661, 662, 3,
		187, 93, 0, 662, 663, 3, 187, 93, 0, 663, 664, 3, 187, 93, 0, 664, 665,
		3, 187, 93, 0, 665, 666, 3, 187, 93, 0, 666, 667, 3, 187, 93, 0, 667, 668,
		3, 187, 93, 0, 668, 679, 1, 0, 0, 0, 669, 679, 7, 7, 0, 0, 670, 671, 3,
		185, 92, 0, 671, 672, 3, 185, 92, 0, 672, 673, 3, 185, 92, 0, 673, 679,
		1, 0, 0, 0, 674, 675, 5, 120, 0, 0, 675, 676, 3, 187, 93, 0, 676, 677,
		3, 187, 93, 0, 677, 679, 1, 0, 0, 0, 678, 653, 1, 0, 0, 0, 678, 659, 1,
		0, 0, 0, 678, 669, 1, 0, 0, 0, 678, 670, 1, 0, 0, 0, 678, 674, 1, 0, 0,
		0, 679, 182, 1, 0, 0, 0, 680, 682, 7, 1, 0, 0, 681, 680, 1, 0, 0, 0, 682,
		683, 1, 0, 0, 0, 683, 681, 1, 0, 0, 0, 683, 684, 1, 0, 0, 0, 684, 184,
		1, 0, 0, 0, 685, 686, 7, 8, 0, 0, 686, 186, 1, 0, 0, 0, 687, 688, 7, 9,
		0, 0, 688, 188, 1, 0, 0, 0, 689, 691, 7, 10, 0, 0, 690, 692, 7, 11, 0,
		0, 691, 690, 1, 0, 0, 0, 691, 692, 1, 0, 0, 0, 692, 693, 1, 0, 0, 0, 693,
		694, 3, 183, 91, 0, 694, 190, 1, 0, 0, 0, 695, 698, 3, 195, 97, 0, 696,
		698, 5, 95, 0, 0, 697, 695, 1, 0, 0, 0, 697, 696, 1, 0, 0, 0, 698, 192,
		1, 0, 0, 0, 699, 701, 7, 12, 0, 0, 700, 699, 1, 0, 0, 0, 701, 194, 1, 0,
		0, 0, 702, 704, 7, 13, 0, 0, 703, 702, 1, 0, 0, 0, 704, 196, 1, 0, 0, 0,
		24, 0, 464, 466, 559, 566, 574, 579, 582, 585, 590, 592, 598, 606, 608,
		616, 626, 637, 647, 678, 683, 691, 697, 700, 703, 2, 6, 0, 0, 0, 1, 0,
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

// FaultLexerInit initializes any static state used to implement FaultLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewFaultLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func FaultLexerInit() {
	staticData := &faultlexerLexerStaticData
	staticData.once.Do(faultlexerLexerInit)
}

// NewFaultLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewFaultLexer(input antlr.CharStream) *FaultLexer {
	FaultLexerInit()
	l := new(FaultLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &faultlexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	l.channelNames = staticData.channelNames
	l.modeNames = staticData.modeNames
	l.RuleNames = staticData.ruleNames
	l.LiteralNames = staticData.literalNames
	l.SymbolicNames = staticData.symbolicNames
	l.GrammarFileName = "FaultLexer.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// FaultLexer tokens.
const (
	FaultLexerALL                    = 1
	FaultLexerASSERT                 = 2
	FaultLexerASSUME                 = 3
	FaultLexerCLOCK                  = 4
	FaultLexerCONST                  = 5
	FaultLexerDEF                    = 6
	FaultLexerELSE                   = 7
	FaultLexerFLOW                   = 8
	FaultLexerFOR                    = 9
	FaultLexerFUNC                   = 10
	FaultLexerIF                     = 11
	FaultLexerIMPORT                 = 12
	FaultLexerINIT                   = 13
	FaultLexerNEW                    = 14
	FaultLexerRETURN                 = 15
	FaultLexerRUN                    = 16
	FaultLexerSPEC                   = 17
	FaultLexerSTOCK                  = 18
	FaultLexerTHEN                   = 19
	FaultLexerWHEN                   = 20
	FaultLexerTHIS                   = 21
	FaultLexerEVENTUALLY             = 22
	FaultLexerEVENTUALLYALWAYS       = 23
	FaultLexerALWAYS                 = 24
	FaultLexerNMT                    = 25
	FaultLexerNFT                    = 26
	FaultLexerNIL                    = 27
	FaultLexerTRUE                   = 28
	FaultLexerFALSE                  = 29
	FaultLexerADVANCE                = 30
	FaultLexerCOMPONENT              = 31
	FaultLexerGLOBAL                 = 32
	FaultLexerSYSTEM                 = 33
	FaultLexerSTART                  = 34
	FaultLexerSTATE                  = 35
	FaultLexerSTAY                   = 36
	FaultLexerTY_STRING              = 37
	FaultLexerTY_BOOL                = 38
	FaultLexerTY_INT                 = 39
	FaultLexerTY_FLOAT               = 40
	FaultLexerTY_NATURAL             = 41
	FaultLexerTY_UNCERTAIN           = 42
	FaultLexerTY_UNKNOWN             = 43
	FaultLexerIDENT                  = 44
	FaultLexerASSIGN                 = 45
	FaultLexerASSIGN_FLOW1           = 46
	FaultLexerASSIGN_FLOW2           = 47
	FaultLexerCOLON                  = 48
	FaultLexerCOMMA                  = 49
	FaultLexerDOT                    = 50
	FaultLexerLPAREN                 = 51
	FaultLexerRPAREN                 = 52
	FaultLexerLCURLY                 = 53
	FaultLexerRCURLY                 = 54
	FaultLexerLBRACE                 = 55
	FaultLexerRBRACE                 = 56
	FaultLexerSEMI                   = 57
	FaultLexerPLUS_PLUS              = 58
	FaultLexerMINUS_MINUS            = 59
	FaultLexerAMPERSAND              = 60
	FaultLexerAND                    = 61
	FaultLexerBANG                   = 62
	FaultLexerEQUALS                 = 63
	FaultLexerNOT_EQUALS             = 64
	FaultLexerLESS                   = 65
	FaultLexerLESS_OR_EQUALS         = 66
	FaultLexerGREATER                = 67
	FaultLexerGREATER_OR_EQUALS      = 68
	FaultLexerOR                     = 69
	FaultLexerPIPE                   = 70
	FaultLexerPLUS                   = 71
	FaultLexerMINUS                  = 72
	FaultLexerCARET                  = 73
	FaultLexerEXPO                   = 74
	FaultLexerMULTI                  = 75
	FaultLexerDIV                    = 76
	FaultLexerMOD                    = 77
	FaultLexerLSHIFT                 = 78
	FaultLexerRSHIFT                 = 79
	FaultLexerBIT_CLEAR              = 80
	FaultLexerDECIMAL_LIT            = 81
	FaultLexerOCTAL_LIT              = 82
	FaultLexerHEX_LIT                = 83
	FaultLexerFLOAT_LIT              = 84
	FaultLexerRAW_STRING_LIT         = 85
	FaultLexerINTERPRETED_STRING_LIT = 86
	FaultLexerWS                     = 87
	FaultLexerCOMMENT                = 88
	FaultLexerTERMINATOR             = 89
	FaultLexerLINE_COMMENT           = 90
)
