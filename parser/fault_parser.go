// Code generated from FaultParser.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // FaultParser

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 82, 481,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4,
	29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33, 4, 34,
	9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4, 39, 9,
	39, 4, 40, 9, 40, 4, 41, 9, 41, 4, 42, 9, 42, 4, 43, 9, 43, 4, 44, 9, 44,
	4, 45, 9, 45, 4, 46, 9, 46, 4, 47, 9, 47, 4, 48, 9, 48, 3, 2, 3, 2, 7,
	2, 99, 10, 2, 12, 2, 14, 2, 102, 11, 2, 3, 2, 7, 2, 105, 10, 2, 12, 2,
	14, 2, 108, 11, 2, 3, 2, 5, 2, 111, 10, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 4, 3, 4, 3, 4, 3, 4, 7, 4, 123, 10, 4, 12, 4, 14, 4, 126, 11,
	4, 3, 4, 5, 4, 129, 10, 4, 3, 4, 3, 4, 3, 5, 5, 5, 134, 10, 5, 3, 5, 3,
	5, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 7, 5, 7, 144, 10, 7, 3, 8, 3, 8, 3,
	9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 7, 9, 156, 10, 9, 12, 9, 14,
	9, 159, 11, 9, 3, 9, 5, 9, 162, 10, 9, 3, 10, 3, 10, 3, 10, 5, 10, 167,
	10, 10, 3, 11, 3, 11, 3, 11, 7, 11, 172, 10, 11, 12, 11, 14, 11, 175, 11,
	11, 3, 12, 3, 12, 3, 12, 7, 12, 180, 10, 12, 12, 12, 14, 12, 183, 11, 12,
	3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 14, 3, 14, 3, 14, 3, 14, 3,
	14, 7, 14, 196, 10, 14, 12, 14, 14, 14, 199, 11, 14, 3, 14, 3, 14, 3, 14,
	3, 14, 3, 14, 3, 14, 7, 14, 207, 10, 14, 12, 14, 14, 14, 210, 11, 14, 3,
	14, 5, 14, 213, 10, 14, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15,
	3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3,
	15, 3, 15, 5, 15, 234, 10, 15, 3, 16, 3, 16, 3, 16, 3, 16, 3, 17, 3, 17,
	5, 17, 242, 10, 17, 3, 17, 3, 17, 3, 18, 6, 18, 247, 10, 18, 13, 18, 14,
	18, 248, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 5, 19, 258, 10,
	19, 3, 20, 3, 20, 3, 20, 3, 20, 5, 20, 264, 10, 20, 3, 21, 3, 21, 3, 21,
	3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 6, 22, 274, 10, 22, 13, 22, 14, 22,
	275, 3, 23, 3, 23, 5, 23, 280, 10, 23, 3, 23, 3, 23, 3, 23, 3, 24, 3, 24,
	5, 24, 287, 10, 24, 3, 24, 3, 24, 3, 24, 3, 25, 3, 25, 3, 26, 3, 26, 3,
	26, 3, 26, 3, 26, 5, 26, 299, 10, 26, 3, 27, 3, 27, 5, 27, 303, 10, 27,
	3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 5, 27, 312, 10, 27, 3,
	28, 3, 28, 3, 29, 3, 29, 3, 29, 3, 29, 5, 29, 320, 10, 29, 3, 29, 3, 29,
	3, 29, 3, 29, 3, 29, 5, 29, 327, 10, 29, 5, 29, 329, 10, 29, 3, 30, 3,
	30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 31, 3, 31, 3, 32, 3, 32, 3, 32, 3, 32,
	3, 32, 7, 32, 344, 10, 32, 12, 32, 14, 32, 347, 11, 32, 3, 33, 3, 33, 7,
	33, 351, 10, 33, 12, 33, 14, 33, 354, 11, 33, 3, 33, 3, 33, 3, 34, 3, 34,
	3, 34, 7, 34, 361, 10, 34, 12, 34, 14, 34, 364, 11, 34, 3, 34, 3, 34, 3,
	34, 3, 34, 3, 34, 3, 34, 3, 34, 3, 34, 5, 34, 374, 10, 34, 3, 34, 3, 34,
	3, 34, 3, 34, 3, 34, 5, 34, 381, 10, 34, 3, 35, 3, 35, 3, 36, 3, 36, 3,
	36, 5, 36, 388, 10, 36, 3, 36, 3, 36, 7, 36, 392, 10, 36, 12, 36, 14, 36,
	395, 11, 36, 3, 36, 3, 36, 3, 37, 3, 37, 3, 37, 3, 37, 5, 37, 403, 10,
	37, 3, 37, 3, 37, 3, 37, 3, 37, 3, 37, 3, 37, 3, 37, 3, 37, 3, 37, 3, 37,
	3, 37, 3, 37, 3, 37, 3, 37, 3, 37, 3, 37, 3, 37, 3, 37, 7, 37, 423, 10,
	37, 12, 37, 14, 37, 426, 11, 37, 3, 38, 3, 38, 3, 38, 3, 38, 3, 38, 3,
	38, 3, 38, 3, 38, 3, 38, 3, 38, 5, 38, 438, 10, 38, 3, 39, 3, 39, 3, 39,
	3, 39, 3, 39, 3, 39, 3, 39, 3, 39, 5, 39, 448, 10, 39, 5, 39, 450, 10,
	39, 3, 40, 3, 40, 3, 40, 5, 40, 455, 10, 40, 3, 41, 3, 41, 3, 41, 5, 41,
	460, 10, 41, 3, 42, 3, 42, 3, 43, 3, 43, 3, 43, 3, 43, 5, 43, 468, 10,
	43, 3, 44, 3, 44, 3, 45, 3, 45, 3, 46, 3, 46, 3, 47, 3, 47, 3, 47, 3, 48,
	3, 48, 3, 48, 2, 3, 72, 49, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24,
	26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60,
	62, 64, 66, 68, 70, 72, 74, 76, 78, 80, 82, 84, 86, 88, 90, 92, 94, 2,
	17, 4, 2, 36, 36, 42, 42, 3, 2, 55, 60, 3, 2, 50, 51, 3, 2, 22, 23, 3,
	2, 24, 25, 5, 2, 52, 52, 63, 65, 67, 72, 3, 2, 38, 39, 3, 2, 29, 35, 4,
	2, 52, 52, 67, 72, 3, 2, 63, 65, 6, 2, 52, 52, 54, 54, 63, 65, 67, 67,
	3, 2, 73, 75, 3, 2, 77, 78, 3, 2, 27, 28, 3, 3, 49, 49, 2, 506, 2, 96,
	3, 2, 2, 2, 4, 114, 3, 2, 2, 2, 6, 118, 3, 2, 2, 2, 8, 133, 3, 2, 2, 2,
	10, 137, 3, 2, 2, 2, 12, 143, 3, 2, 2, 2, 14, 145, 3, 2, 2, 2, 16, 147,
	3, 2, 2, 2, 18, 163, 3, 2, 2, 2, 20, 168, 3, 2, 2, 2, 22, 176, 3, 2, 2,
	2, 24, 184, 3, 2, 2, 2, 26, 212, 3, 2, 2, 2, 28, 233, 3, 2, 2, 2, 30, 235,
	3, 2, 2, 2, 32, 239, 3, 2, 2, 2, 34, 246, 3, 2, 2, 2, 36, 257, 3, 2, 2,
	2, 38, 263, 3, 2, 2, 2, 40, 265, 3, 2, 2, 2, 42, 268, 3, 2, 2, 2, 44, 277,
	3, 2, 2, 2, 46, 284, 3, 2, 2, 2, 48, 291, 3, 2, 2, 2, 50, 298, 3, 2, 2,
	2, 52, 311, 3, 2, 2, 2, 54, 313, 3, 2, 2, 2, 56, 315, 3, 2, 2, 2, 58, 330,
	3, 2, 2, 2, 60, 336, 3, 2, 2, 2, 62, 338, 3, 2, 2, 2, 64, 348, 3, 2, 2,
	2, 66, 380, 3, 2, 2, 2, 68, 382, 3, 2, 2, 2, 70, 384, 3, 2, 2, 2, 72, 402,
	3, 2, 2, 2, 74, 437, 3, 2, 2, 2, 76, 449, 3, 2, 2, 2, 78, 454, 3, 2, 2,
	2, 80, 459, 3, 2, 2, 2, 82, 461, 3, 2, 2, 2, 84, 467, 3, 2, 2, 2, 86, 469,
	3, 2, 2, 2, 88, 471, 3, 2, 2, 2, 90, 473, 3, 2, 2, 2, 92, 475, 3, 2, 2,
	2, 94, 478, 3, 2, 2, 2, 96, 100, 5, 4, 3, 2, 97, 99, 5, 6, 4, 2, 98, 97,
	3, 2, 2, 2, 99, 102, 3, 2, 2, 2, 100, 98, 3, 2, 2, 2, 100, 101, 3, 2, 2,
	2, 101, 106, 3, 2, 2, 2, 102, 100, 3, 2, 2, 2, 103, 105, 5, 12, 7, 2, 104,
	103, 3, 2, 2, 2, 105, 108, 3, 2, 2, 2, 106, 104, 3, 2, 2, 2, 106, 107,
	3, 2, 2, 2, 107, 110, 3, 2, 2, 2, 108, 106, 3, 2, 2, 2, 109, 111, 5, 58,
	30, 2, 110, 109, 3, 2, 2, 2, 110, 111, 3, 2, 2, 2, 111, 112, 3, 2, 2, 2,
	112, 113, 5, 94, 48, 2, 113, 3, 3, 2, 2, 2, 114, 115, 7, 19, 2, 2, 115,
	116, 7, 36, 2, 2, 116, 117, 5, 94, 48, 2, 117, 5, 3, 2, 2, 2, 118, 128,
	7, 14, 2, 2, 119, 129, 5, 8, 5, 2, 120, 124, 7, 43, 2, 2, 121, 123, 5,
	8, 5, 2, 122, 121, 3, 2, 2, 2, 123, 126, 3, 2, 2, 2, 124, 122, 3, 2, 2,
	2, 124, 125, 3, 2, 2, 2, 125, 127, 3, 2, 2, 2, 126, 124, 3, 2, 2, 2, 127,
	129, 7, 44, 2, 2, 128, 119, 3, 2, 2, 2, 128, 120, 3, 2, 2, 2, 129, 130,
	3, 2, 2, 2, 130, 131, 5, 94, 48, 2, 131, 7, 3, 2, 2, 2, 132, 134, 9, 2,
	2, 2, 133, 132, 3, 2, 2, 2, 133, 134, 3, 2, 2, 2, 134, 135, 3, 2, 2, 2,
	135, 136, 5, 10, 6, 2, 136, 9, 3, 2, 2, 2, 137, 138, 5, 88, 45, 2, 138,
	11, 3, 2, 2, 2, 139, 144, 5, 16, 9, 2, 140, 144, 5, 24, 13, 2, 141, 144,
	5, 44, 23, 2, 142, 144, 5, 46, 24, 2, 143, 139, 3, 2, 2, 2, 143, 140, 3,
	2, 2, 2, 143, 141, 3, 2, 2, 2, 143, 142, 3, 2, 2, 2, 144, 13, 3, 2, 2,
	2, 145, 146, 9, 3, 2, 2, 146, 15, 3, 2, 2, 2, 147, 161, 7, 7, 2, 2, 148,
	149, 5, 18, 10, 2, 149, 150, 5, 94, 48, 2, 150, 162, 3, 2, 2, 2, 151, 157,
	7, 43, 2, 2, 152, 153, 5, 18, 10, 2, 153, 154, 5, 94, 48, 2, 154, 156,
	3, 2, 2, 2, 155, 152, 3, 2, 2, 2, 156, 159, 3, 2, 2, 2, 157, 155, 3, 2,
	2, 2, 157, 158, 3, 2, 2, 2, 158, 160, 3, 2, 2, 2, 159, 157, 3, 2, 2, 2,
	160, 162, 7, 44, 2, 2, 161, 148, 3, 2, 2, 2, 161, 151, 3, 2, 2, 2, 162,
	17, 3, 2, 2, 2, 163, 166, 5, 20, 11, 2, 164, 165, 7, 37, 2, 2, 165, 167,
	5, 22, 12, 2, 166, 164, 3, 2, 2, 2, 166, 167, 3, 2, 2, 2, 167, 19, 3, 2,
	2, 2, 168, 173, 5, 76, 39, 2, 169, 170, 7, 41, 2, 2, 170, 172, 5, 76, 39,
	2, 171, 169, 3, 2, 2, 2, 172, 175, 3, 2, 2, 2, 173, 171, 3, 2, 2, 2, 173,
	174, 3, 2, 2, 2, 174, 21, 3, 2, 2, 2, 175, 173, 3, 2, 2, 2, 176, 181, 5,
	72, 37, 2, 177, 178, 7, 41, 2, 2, 178, 180, 5, 72, 37, 2, 179, 177, 3,
	2, 2, 2, 180, 183, 3, 2, 2, 2, 181, 179, 3, 2, 2, 2, 181, 182, 3, 2, 2,
	2, 182, 23, 3, 2, 2, 2, 183, 181, 3, 2, 2, 2, 184, 185, 7, 8, 2, 2, 185,
	186, 7, 36, 2, 2, 186, 187, 7, 37, 2, 2, 187, 188, 5, 26, 14, 2, 188, 189,
	5, 94, 48, 2, 189, 25, 3, 2, 2, 2, 190, 191, 7, 10, 2, 2, 191, 197, 7,
	45, 2, 2, 192, 193, 5, 28, 15, 2, 193, 194, 7, 41, 2, 2, 194, 196, 3, 2,
	2, 2, 195, 192, 3, 2, 2, 2, 196, 199, 3, 2, 2, 2, 197, 195, 3, 2, 2, 2,
	197, 198, 3, 2, 2, 2, 198, 200, 3, 2, 2, 2, 199, 197, 3, 2, 2, 2, 200,
	213, 7, 46, 2, 2, 201, 202, 7, 20, 2, 2, 202, 208, 7, 45, 2, 2, 203, 204,
	5, 28, 15, 2, 204, 205, 7, 41, 2, 2, 205, 207, 3, 2, 2, 2, 206, 203, 3,
	2, 2, 2, 207, 210, 3, 2, 2, 2, 208, 206, 3, 2, 2, 2, 208, 209, 3, 2, 2,
	2, 209, 211, 3, 2, 2, 2, 210, 208, 3, 2, 2, 2, 211, 213, 7, 46, 2, 2, 212,
	190, 3, 2, 2, 2, 212, 201, 3, 2, 2, 2, 213, 27, 3, 2, 2, 2, 214, 215, 7,
	36, 2, 2, 215, 216, 7, 40, 2, 2, 216, 234, 5, 80, 41, 2, 217, 218, 7, 36,
	2, 2, 218, 219, 7, 40, 2, 2, 219, 234, 5, 88, 45, 2, 220, 221, 7, 36, 2,
	2, 221, 222, 7, 40, 2, 2, 222, 234, 5, 92, 47, 2, 223, 224, 7, 36, 2, 2,
	224, 225, 7, 40, 2, 2, 225, 234, 5, 76, 39, 2, 226, 227, 7, 36, 2, 2, 227,
	228, 7, 40, 2, 2, 228, 234, 5, 78, 40, 2, 229, 230, 7, 36, 2, 2, 230, 231,
	7, 40, 2, 2, 231, 234, 5, 70, 36, 2, 232, 234, 7, 36, 2, 2, 233, 214, 3,
	2, 2, 2, 233, 217, 3, 2, 2, 2, 233, 220, 3, 2, 2, 2, 233, 223, 3, 2, 2,
	2, 233, 226, 3, 2, 2, 2, 233, 229, 3, 2, 2, 2, 233, 232, 3, 2, 2, 2, 234,
	29, 3, 2, 2, 2, 235, 236, 7, 15, 2, 2, 236, 237, 5, 74, 38, 2, 237, 238,
	5, 94, 48, 2, 238, 31, 3, 2, 2, 2, 239, 241, 7, 45, 2, 2, 240, 242, 5,
	34, 18, 2, 241, 240, 3, 2, 2, 2, 241, 242, 3, 2, 2, 2, 242, 243, 3, 2,
	2, 2, 243, 244, 7, 46, 2, 2, 244, 33, 3, 2, 2, 2, 245, 247, 5, 36, 19,
	2, 246, 245, 3, 2, 2, 2, 247, 248, 3, 2, 2, 2, 248, 246, 3, 2, 2, 2, 248,
	249, 3, 2, 2, 2, 249, 35, 3, 2, 2, 2, 250, 258, 5, 16, 9, 2, 251, 258,
	5, 30, 16, 2, 252, 253, 5, 38, 20, 2, 253, 254, 5, 94, 48, 2, 254, 258,
	3, 2, 2, 2, 255, 258, 5, 32, 17, 2, 256, 258, 5, 56, 29, 2, 257, 250, 3,
	2, 2, 2, 257, 251, 3, 2, 2, 2, 257, 252, 3, 2, 2, 2, 257, 255, 3, 2, 2,
	2, 257, 256, 3, 2, 2, 2, 258, 37, 3, 2, 2, 2, 259, 264, 5, 72, 37, 2, 260,
	264, 5, 40, 21, 2, 261, 264, 5, 52, 27, 2, 262, 264, 5, 54, 28, 2, 263,
	259, 3, 2, 2, 2, 263, 260, 3, 2, 2, 2, 263, 261, 3, 2, 2, 2, 263, 262,
	3, 2, 2, 2, 264, 39, 3, 2, 2, 2, 265, 266, 5, 72, 37, 2, 266, 267, 9, 4,
	2, 2, 267, 41, 3, 2, 2, 2, 268, 273, 5, 76, 39, 2, 269, 270, 7, 47, 2,
	2, 270, 271, 5, 72, 37, 2, 271, 272, 7, 48, 2, 2, 272, 274, 3, 2, 2, 2,
	273, 269, 3, 2, 2, 2, 274, 275, 3, 2, 2, 2, 275, 273, 3, 2, 2, 2, 275,
	276, 3, 2, 2, 2, 276, 43, 3, 2, 2, 2, 277, 279, 7, 4, 2, 2, 278, 280, 5,
	48, 25, 2, 279, 278, 3, 2, 2, 2, 279, 280, 3, 2, 2, 2, 280, 281, 3, 2,
	2, 2, 281, 282, 5, 50, 26, 2, 282, 283, 5, 94, 48, 2, 283, 45, 3, 2, 2,
	2, 284, 286, 7, 5, 2, 2, 285, 287, 5, 48, 25, 2, 286, 285, 3, 2, 2, 2,
	286, 287, 3, 2, 2, 2, 287, 288, 3, 2, 2, 2, 288, 289, 5, 50, 26, 2, 289,
	290, 5, 94, 48, 2, 290, 47, 3, 2, 2, 2, 291, 292, 9, 5, 2, 2, 292, 49,
	3, 2, 2, 2, 293, 299, 5, 72, 37, 2, 294, 295, 5, 72, 37, 2, 295, 296, 9,
	6, 2, 2, 296, 297, 5, 82, 42, 2, 297, 299, 3, 2, 2, 2, 298, 293, 3, 2,
	2, 2, 298, 294, 3, 2, 2, 2, 299, 51, 3, 2, 2, 2, 300, 302, 5, 22, 12, 2,
	301, 303, 9, 7, 2, 2, 302, 301, 3, 2, 2, 2, 302, 303, 3, 2, 2, 2, 303,
	304, 3, 2, 2, 2, 304, 305, 7, 37, 2, 2, 305, 306, 5, 22, 12, 2, 306, 312,
	3, 2, 2, 2, 307, 308, 5, 22, 12, 2, 308, 309, 9, 8, 2, 2, 309, 310, 5,
	22, 12, 2, 310, 312, 3, 2, 2, 2, 311, 300, 3, 2, 2, 2, 311, 307, 3, 2,
	2, 2, 312, 53, 3, 2, 2, 2, 313, 314, 7, 49, 2, 2, 314, 55, 3, 2, 2, 2,
	315, 319, 7, 13, 2, 2, 316, 317, 5, 38, 20, 2, 317, 318, 7, 49, 2, 2, 318,
	320, 3, 2, 2, 2, 319, 316, 3, 2, 2, 2, 319, 320, 3, 2, 2, 2, 320, 321,
	3, 2, 2, 2, 321, 322, 5, 72, 37, 2, 322, 328, 5, 32, 17, 2, 323, 326, 7,
	9, 2, 2, 324, 327, 5, 56, 29, 2, 325, 327, 5, 32, 17, 2, 326, 324, 3, 2,
	2, 2, 326, 325, 3, 2, 2, 2, 327, 329, 3, 2, 2, 2, 328, 323, 3, 2, 2, 2,
	328, 329, 3, 2, 2, 2, 329, 57, 3, 2, 2, 2, 330, 331, 7, 11, 2, 2, 331,
	332, 5, 60, 31, 2, 332, 333, 7, 18, 2, 2, 333, 334, 5, 64, 33, 2, 334,
	335, 5, 94, 48, 2, 335, 59, 3, 2, 2, 2, 336, 337, 5, 82, 42, 2, 337, 61,
	3, 2, 2, 2, 338, 339, 7, 36, 2, 2, 339, 340, 7, 42, 2, 2, 340, 345, 7,
	36, 2, 2, 341, 342, 7, 42, 2, 2, 342, 344, 7, 36, 2, 2, 343, 341, 3, 2,
	2, 2, 344, 347, 3, 2, 2, 2, 345, 343, 3, 2, 2, 2, 345, 346, 3, 2, 2, 2,
	346, 63, 3, 2, 2, 2, 347, 345, 3, 2, 2, 2, 348, 352, 7, 45, 2, 2, 349,
	351, 5, 66, 34, 2, 350, 349, 3, 2, 2, 2, 351, 354, 3, 2, 2, 2, 352, 350,
	3, 2, 2, 2, 352, 353, 3, 2, 2, 2, 353, 355, 3, 2, 2, 2, 354, 352, 3, 2,
	2, 2, 355, 356, 7, 46, 2, 2, 356, 65, 3, 2, 2, 2, 357, 362, 5, 62, 32,
	2, 358, 359, 7, 62, 2, 2, 359, 361, 5, 62, 32, 2, 360, 358, 3, 2, 2, 2,
	361, 364, 3, 2, 2, 2, 362, 360, 3, 2, 2, 2, 362, 363, 3, 2, 2, 2, 363,
	365, 3, 2, 2, 2, 364, 362, 3, 2, 2, 2, 365, 366, 5, 94, 48, 2, 366, 381,
	3, 2, 2, 2, 367, 368, 7, 36, 2, 2, 368, 369, 7, 37, 2, 2, 369, 370, 7,
	16, 2, 2, 370, 373, 7, 36, 2, 2, 371, 372, 7, 42, 2, 2, 372, 374, 7, 36,
	2, 2, 373, 371, 3, 2, 2, 2, 373, 374, 3, 2, 2, 2, 374, 375, 3, 2, 2, 2,
	375, 381, 5, 94, 48, 2, 376, 377, 5, 38, 20, 2, 377, 378, 5, 94, 48, 2,
	378, 381, 3, 2, 2, 2, 379, 381, 5, 56, 29, 2, 380, 357, 3, 2, 2, 2, 380,
	367, 3, 2, 2, 2, 380, 376, 3, 2, 2, 2, 380, 379, 3, 2, 2, 2, 381, 67, 3,
	2, 2, 2, 382, 383, 9, 9, 2, 2, 383, 69, 3, 2, 2, 2, 384, 385, 5, 68, 35,
	2, 385, 387, 7, 43, 2, 2, 386, 388, 5, 74, 38, 2, 387, 386, 3, 2, 2, 2,
	387, 388, 3, 2, 2, 2, 388, 393, 3, 2, 2, 2, 389, 390, 7, 41, 2, 2, 390,
	392, 5, 74, 38, 2, 391, 389, 3, 2, 2, 2, 392, 395, 3, 2, 2, 2, 393, 391,
	3, 2, 2, 2, 393, 394, 3, 2, 2, 2, 394, 396, 3, 2, 2, 2, 395, 393, 3, 2,
	2, 2, 396, 397, 7, 44, 2, 2, 397, 71, 3, 2, 2, 2, 398, 399, 8, 37, 1, 2,
	399, 403, 5, 74, 38, 2, 400, 403, 5, 70, 36, 2, 401, 403, 5, 78, 40, 2,
	402, 398, 3, 2, 2, 2, 402, 400, 3, 2, 2, 2, 402, 401, 3, 2, 2, 2, 403,
	424, 3, 2, 2, 2, 404, 405, 12, 8, 2, 2, 405, 406, 7, 66, 2, 2, 406, 423,
	5, 72, 37, 9, 407, 408, 12, 7, 2, 2, 408, 409, 9, 10, 2, 2, 409, 423, 5,
	72, 37, 8, 410, 411, 12, 6, 2, 2, 411, 412, 9, 11, 2, 2, 412, 423, 5, 72,
	37, 7, 413, 414, 12, 5, 2, 2, 414, 415, 9, 3, 2, 2, 415, 423, 5, 72, 37,
	6, 416, 417, 12, 4, 2, 2, 417, 418, 7, 53, 2, 2, 418, 423, 5, 72, 37, 5,
	419, 420, 12, 3, 2, 2, 420, 421, 7, 61, 2, 2, 421, 423, 5, 72, 37, 4, 422,
	404, 3, 2, 2, 2, 422, 407, 3, 2, 2, 2, 422, 410, 3, 2, 2, 2, 422, 413,
	3, 2, 2, 2, 422, 416, 3, 2, 2, 2, 422, 419, 3, 2, 2, 2, 423, 426, 3, 2,
	2, 2, 424, 422, 3, 2, 2, 2, 424, 425, 3, 2, 2, 2, 425, 73, 3, 2, 2, 2,
	426, 424, 3, 2, 2, 2, 427, 438, 7, 26, 2, 2, 428, 438, 5, 80, 41, 2, 429,
	438, 5, 88, 45, 2, 430, 438, 5, 90, 46, 2, 431, 438, 5, 76, 39, 2, 432,
	438, 5, 42, 22, 2, 433, 434, 7, 43, 2, 2, 434, 435, 5, 72, 37, 2, 435,
	436, 7, 44, 2, 2, 436, 438, 3, 2, 2, 2, 437, 427, 3, 2, 2, 2, 437, 428,
	3, 2, 2, 2, 437, 429, 3, 2, 2, 2, 437, 430, 3, 2, 2, 2, 437, 431, 3, 2,
	2, 2, 437, 432, 3, 2, 2, 2, 437, 433, 3, 2, 2, 2, 438, 75, 3, 2, 2, 2,
	439, 450, 7, 36, 2, 2, 440, 450, 5, 62, 32, 2, 441, 450, 7, 21, 2, 2, 442,
	450, 7, 6, 2, 2, 443, 444, 7, 16, 2, 2, 444, 447, 7, 36, 2, 2, 445, 446,
	7, 42, 2, 2, 446, 448, 7, 36, 2, 2, 447, 445, 3, 2, 2, 2, 447, 448, 3,
	2, 2, 2, 448, 450, 3, 2, 2, 2, 449, 439, 3, 2, 2, 2, 449, 440, 3, 2, 2,
	2, 449, 441, 3, 2, 2, 2, 449, 442, 3, 2, 2, 2, 449, 443, 3, 2, 2, 2, 450,
	77, 3, 2, 2, 2, 451, 455, 3, 2, 2, 2, 452, 453, 9, 12, 2, 2, 453, 455,
	5, 72, 37, 2, 454, 451, 3, 2, 2, 2, 454, 452, 3, 2, 2, 2, 455, 79, 3, 2,
	2, 2, 456, 460, 5, 82, 42, 2, 457, 460, 5, 84, 43, 2, 458, 460, 5, 86,
	44, 2, 459, 456, 3, 2, 2, 2, 459, 457, 3, 2, 2, 2, 459, 458, 3, 2, 2, 2,
	460, 81, 3, 2, 2, 2, 461, 462, 9, 13, 2, 2, 462, 83, 3, 2, 2, 2, 463, 464,
	7, 64, 2, 2, 464, 468, 5, 82, 42, 2, 465, 466, 7, 64, 2, 2, 466, 468, 5,
	86, 44, 2, 467, 463, 3, 2, 2, 2, 467, 465, 3, 2, 2, 2, 468, 85, 3, 2, 2,
	2, 469, 470, 7, 76, 2, 2, 470, 87, 3, 2, 2, 2, 471, 472, 9, 14, 2, 2, 472,
	89, 3, 2, 2, 2, 473, 474, 9, 15, 2, 2, 474, 91, 3, 2, 2, 2, 475, 476, 7,
	12, 2, 2, 476, 477, 5, 32, 17, 2, 477, 93, 3, 2, 2, 2, 478, 479, 9, 16,
	2, 2, 479, 95, 3, 2, 2, 2, 47, 100, 106, 110, 124, 128, 133, 143, 157,
	161, 166, 173, 181, 197, 208, 212, 233, 241, 248, 257, 263, 275, 279, 286,
	298, 302, 311, 319, 326, 328, 345, 352, 362, 373, 380, 387, 393, 402, 422,
	424, 437, 447, 449, 454, 459, 467,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'all'", "'assert'", "'assume'", "'clock'", "'const'", "'def'", "'else'",
	"'flow'", "'for'", "'func'", "'if'", "'import'", "'init'", "'new'", "'return'",
	"'run'", "'spec'", "'stock'", "'this'", "'eventually'", "'always'", "'nmt'",
	"'nft'", "'nil'", "'true'", "'false'", "'string'", "'bool'", "'int'", "'float'",
	"'natural'", "'uncertain'", "'unknown'", "", "'='", "'->'", "'<-'", "':'",
	"','", "'.'", "'('", "')'", "'{'", "'}'", "'['", "']'", "';'", "'++'",
	"'--'", "'&'", "'&&'", "'!'", "'=='", "'!='", "'<'", "'<='", "'>'", "'>='",
	"'||'", "'|'", "'+'", "'-'", "'^'", "'**'", "'*'", "'/'", "'%'", "'<<'",
	"'>>'", "'&^'",
}
var symbolicNames = []string{
	"", "ALL", "ASSERT", "ASSUME", "CLOCK", "CONST", "DEF", "ELSE", "FLOW",
	"FOR", "FUNC", "IF", "IMPORT", "INIT", "NEW", "RETURN", "RUN", "SPEC",
	"STOCK", "THIS", "EVENTUALLY", "ALWAYS", "NMT", "NFT", "NIL", "TRUE", "FALSE",
	"TY_STRING", "TY_BOOL", "TY_INT", "TY_FLOAT", "TY_NATURAL", "TY_UNCERTAIN",
	"TY_UNKNOWN", "IDENT", "ASSIGN", "ASSIGN_FLOW1", "ASSIGN_FLOW2", "COLON",
	"COMMA", "DOT", "LPAREN", "RPAREN", "LCURLY", "RCURLY", "LBRACE", "RBRACE",
	"SEMI", "PLUS_PLUS", "MINUS_MINUS", "AMPERSAND", "AND", "BANG", "EQUALS",
	"NOT_EQUALS", "LESS", "LESS_OR_EQUALS", "GREATER", "GREATER_OR_EQUALS",
	"OR", "PIPE", "PLUS", "MINUS", "CARET", "EXPO", "MULTI", "DIV", "MOD",
	"LSHIFT", "RSHIFT", "BIT_CLEAR", "DECIMAL_LIT", "OCTAL_LIT", "HEX_LIT",
	"FLOAT_LIT", "RAW_STRING_LIT", "INTERPRETED_STRING_LIT", "WS", "COMMENT",
	"TERMINATOR", "LINE_COMMENT",
}

var ruleNames = []string{
	"spec", "specClause", "importDecl", "importSpec", "importPath", "declaration",
	"comparison", "constDecl", "constSpec", "identList", "expressionList",
	"structDecl", "structType", "structProperties", "initDecl", "block", "statementList",
	"statement", "simpleStmt", "incDecStmt", "accessHistory", "assertion",
	"assumption", "temporal", "invariant", "assignment", "emptyStmt", "ifStmt",
	"forStmt", "rounds", "paramCall", "runBlock", "runStep", "faultType", "solvable",
	"expression", "operand", "operandName", "prefix", "numeric", "integer",
	"negative", "float_", "string_", "bool_", "functionLit", "eos",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type FaultParser struct {
	*antlr.BaseParser
}

func NewFaultParser(input antlr.TokenStream) *FaultParser {
	this := new(FaultParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "FaultParser.g4"

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
	FaultParserTHIS                   = 19
	FaultParserEVENTUALLY             = 20
	FaultParserALWAYS                 = 21
	FaultParserNMT                    = 22
	FaultParserNFT                    = 23
	FaultParserNIL                    = 24
	FaultParserTRUE                   = 25
	FaultParserFALSE                  = 26
	FaultParserTY_STRING              = 27
	FaultParserTY_BOOL                = 28
	FaultParserTY_INT                 = 29
	FaultParserTY_FLOAT               = 30
	FaultParserTY_NATURAL             = 31
	FaultParserTY_UNCERTAIN           = 32
	FaultParserTY_UNKNOWN             = 33
	FaultParserIDENT                  = 34
	FaultParserASSIGN                 = 35
	FaultParserASSIGN_FLOW1           = 36
	FaultParserASSIGN_FLOW2           = 37
	FaultParserCOLON                  = 38
	FaultParserCOMMA                  = 39
	FaultParserDOT                    = 40
	FaultParserLPAREN                 = 41
	FaultParserRPAREN                 = 42
	FaultParserLCURLY                 = 43
	FaultParserRCURLY                 = 44
	FaultParserLBRACE                 = 45
	FaultParserRBRACE                 = 46
	FaultParserSEMI                   = 47
	FaultParserPLUS_PLUS              = 48
	FaultParserMINUS_MINUS            = 49
	FaultParserAMPERSAND              = 50
	FaultParserAND                    = 51
	FaultParserBANG                   = 52
	FaultParserEQUALS                 = 53
	FaultParserNOT_EQUALS             = 54
	FaultParserLESS                   = 55
	FaultParserLESS_OR_EQUALS         = 56
	FaultParserGREATER                = 57
	FaultParserGREATER_OR_EQUALS      = 58
	FaultParserOR                     = 59
	FaultParserPIPE                   = 60
	FaultParserPLUS                   = 61
	FaultParserMINUS                  = 62
	FaultParserCARET                  = 63
	FaultParserEXPO                   = 64
	FaultParserMULTI                  = 65
	FaultParserDIV                    = 66
	FaultParserMOD                    = 67
	FaultParserLSHIFT                 = 68
	FaultParserRSHIFT                 = 69
	FaultParserBIT_CLEAR              = 70
	FaultParserDECIMAL_LIT            = 71
	FaultParserOCTAL_LIT              = 72
	FaultParserHEX_LIT                = 73
	FaultParserFLOAT_LIT              = 74
	FaultParserRAW_STRING_LIT         = 75
	FaultParserINTERPRETED_STRING_LIT = 76
	FaultParserWS                     = 77
	FaultParserCOMMENT                = 78
	FaultParserTERMINATOR             = 79
	FaultParserLINE_COMMENT           = 80
)

// FaultParser rules.
const (
	FaultParserRULE_spec             = 0
	FaultParserRULE_specClause       = 1
	FaultParserRULE_importDecl       = 2
	FaultParserRULE_importSpec       = 3
	FaultParserRULE_importPath       = 4
	FaultParserRULE_declaration      = 5
	FaultParserRULE_comparison       = 6
	FaultParserRULE_constDecl        = 7
	FaultParserRULE_constSpec        = 8
	FaultParserRULE_identList        = 9
	FaultParserRULE_expressionList   = 10
	FaultParserRULE_structDecl       = 11
	FaultParserRULE_structType       = 12
	FaultParserRULE_structProperties = 13
	FaultParserRULE_initDecl         = 14
	FaultParserRULE_block            = 15
	FaultParserRULE_statementList    = 16
	FaultParserRULE_statement        = 17
	FaultParserRULE_simpleStmt       = 18
	FaultParserRULE_incDecStmt       = 19
	FaultParserRULE_accessHistory    = 20
	FaultParserRULE_assertion        = 21
	FaultParserRULE_assumption       = 22
	FaultParserRULE_temporal         = 23
	FaultParserRULE_invariant        = 24
	FaultParserRULE_assignment       = 25
	FaultParserRULE_emptyStmt        = 26
	FaultParserRULE_ifStmt           = 27
	FaultParserRULE_forStmt          = 28
	FaultParserRULE_rounds           = 29
	FaultParserRULE_paramCall        = 30
	FaultParserRULE_runBlock         = 31
	FaultParserRULE_runStep          = 32
	FaultParserRULE_faultType        = 33
	FaultParserRULE_solvable         = 34
	FaultParserRULE_expression       = 35
	FaultParserRULE_operand          = 36
	FaultParserRULE_operandName      = 37
	FaultParserRULE_prefix           = 38
	FaultParserRULE_numeric          = 39
	FaultParserRULE_integer          = 40
	FaultParserRULE_negative         = 41
	FaultParserRULE_float_           = 42
	FaultParserRULE_string_          = 43
	FaultParserRULE_bool_            = 44
	FaultParserRULE_functionLit      = 45
	FaultParserRULE_eos              = 46
)

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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISpecClauseContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISpecClauseContext)
}

func (s *SpecContext) Eos() IEosContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEosContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEosContext)
}

func (s *SpecContext) AllImportDecl() []IImportDeclContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IImportDeclContext)(nil)).Elem())
	var tst = make([]IImportDeclContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IImportDeclContext)
		}
	}

	return tst
}

func (s *SpecContext) ImportDecl(i int) IImportDeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IImportDeclContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IImportDeclContext)
}

func (s *SpecContext) AllDeclaration() []IDeclarationContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDeclarationContext)(nil)).Elem())
	var tst = make([]IDeclarationContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDeclarationContext)
		}
	}

	return tst
}

func (s *SpecContext) Declaration(i int) IDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDeclarationContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDeclarationContext)
}

func (s *SpecContext) ForStmt() IForStmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IForStmtContext)(nil)).Elem(), 0)

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
	localctx = NewSpecContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, FaultParserRULE_spec)
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
		p.SetState(94)
		p.SpecClause()
	}
	p.SetState(98)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == FaultParserIMPORT {
		{
			p.SetState(95)
			p.ImportDecl()
		}

		p.SetState(100)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(104)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<FaultParserASSERT)|(1<<FaultParserASSUME)|(1<<FaultParserCONST)|(1<<FaultParserDEF))) != 0 {
		{
			p.SetState(101)
			p.Declaration()
		}

		p.SetState(106)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(108)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == FaultParserFOR {
		{
			p.SetState(107)
			p.ForStmt()
		}

	}
	{
		p.SetState(110)
		p.Eos()
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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEosContext)(nil)).Elem(), 0)

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
	localctx = NewSpecClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, FaultParserRULE_specClause)

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
		p.SetState(112)
		p.Match(FaultParserSPEC)
	}
	{
		p.SetState(113)
		p.Match(FaultParserIDENT)
	}
	{
		p.SetState(114)
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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEosContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEosContext)
}

func (s *ImportDeclContext) AllImportSpec() []IImportSpecContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IImportSpecContext)(nil)).Elem())
	var tst = make([]IImportSpecContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IImportSpecContext)
		}
	}

	return tst
}

func (s *ImportDeclContext) ImportSpec(i int) IImportSpecContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IImportSpecContext)(nil)).Elem(), i)

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
	localctx = NewImportDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, FaultParserRULE_importDecl)
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
		p.SetState(116)
		p.Match(FaultParserIMPORT)
	}
	p.SetState(126)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case FaultParserIDENT, FaultParserDOT, FaultParserRAW_STRING_LIT, FaultParserINTERPRETED_STRING_LIT:
		{
			p.SetState(117)
			p.ImportSpec()
		}

	case FaultParserLPAREN:
		{
			p.SetState(118)
			p.Match(FaultParserLPAREN)
		}
		p.SetState(122)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == FaultParserIDENT || _la == FaultParserDOT || _la == FaultParserRAW_STRING_LIT || _la == FaultParserINTERPRETED_STRING_LIT {
			{
				p.SetState(119)
				p.ImportSpec()
			}

			p.SetState(124)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(125)
			p.Match(FaultParserRPAREN)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	{
		p.SetState(128)
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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IImportPathContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IImportPathContext)
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
	localctx = NewImportSpecContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, FaultParserRULE_importSpec)
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
	p.SetState(131)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == FaultParserIDENT || _la == FaultParserDOT {
		{
			p.SetState(130)
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
		p.SetState(133)
		p.ImportPath()
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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IString_Context)(nil)).Elem(), 0)

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
	localctx = NewImportPathContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, FaultParserRULE_importPath)

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
		p.SetState(135)
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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConstDeclContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConstDeclContext)
}

func (s *DeclarationContext) StructDecl() IStructDeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStructDeclContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStructDeclContext)
}

func (s *DeclarationContext) Assertion() IAssertionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAssertionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAssertionContext)
}

func (s *DeclarationContext) Assumption() IAssumptionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAssumptionContext)(nil)).Elem(), 0)

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
	localctx = NewDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, FaultParserRULE_declaration)

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

	p.SetState(141)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case FaultParserCONST:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(137)
			p.ConstDecl()
		}

	case FaultParserDEF:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(138)
			p.StructDecl()
		}

	case FaultParserASSERT:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(139)
			p.Assertion()
		}

	case FaultParserASSUME:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(140)
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
	localctx = NewComparisonContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, FaultParserRULE_comparison)
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
		p.SetState(143)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-53)&-(0x1f+1)) == 0 && ((1<<uint((_la-53)))&((1<<(FaultParserEQUALS-53))|(1<<(FaultParserNOT_EQUALS-53))|(1<<(FaultParserLESS-53))|(1<<(FaultParserLESS_OR_EQUALS-53))|(1<<(FaultParserGREATER-53))|(1<<(FaultParserGREATER_OR_EQUALS-53)))) != 0) {
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

func (s *ConstDeclContext) AllConstSpec() []IConstSpecContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IConstSpecContext)(nil)).Elem())
	var tst = make([]IConstSpecContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IConstSpecContext)
		}
	}

	return tst
}

func (s *ConstDeclContext) ConstSpec(i int) IConstSpecContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConstSpecContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IConstSpecContext)
}

func (s *ConstDeclContext) AllEos() []IEosContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IEosContext)(nil)).Elem())
	var tst = make([]IEosContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IEosContext)
		}
	}

	return tst
}

func (s *ConstDeclContext) Eos(i int) IEosContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEosContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IEosContext)
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
	localctx = NewConstDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, FaultParserRULE_constDecl)
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
		p.SetState(145)
		p.Match(FaultParserCONST)
	}
	p.SetState(159)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case FaultParserCLOCK, FaultParserNEW, FaultParserTHIS, FaultParserIDENT:
		{
			p.SetState(146)
			p.ConstSpec()
		}
		{
			p.SetState(147)
			p.Eos()
		}

	case FaultParserLPAREN:
		{
			p.SetState(149)
			p.Match(FaultParserLPAREN)
		}
		p.SetState(155)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ((_la-4)&-(0x1f+1)) == 0 && ((1<<uint((_la-4)))&((1<<(FaultParserCLOCK-4))|(1<<(FaultParserNEW-4))|(1<<(FaultParserTHIS-4))|(1<<(FaultParserIDENT-4)))) != 0 {
			{
				p.SetState(150)
				p.ConstSpec()
			}
			{
				p.SetState(151)
				p.Eos()
			}

			p.SetState(157)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(158)
			p.Match(FaultParserRPAREN)
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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIdentListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIdentListContext)
}

func (s *ConstSpecContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(FaultParserASSIGN, 0)
}

func (s *ConstSpecContext) ExpressionList() IExpressionListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionListContext)
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
	localctx = NewConstSpecContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, FaultParserRULE_constSpec)
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
		p.SetState(161)
		p.IdentList()
	}
	p.SetState(164)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == FaultParserASSIGN {
		{
			p.SetState(162)
			p.Match(FaultParserASSIGN)
		}
		{
			p.SetState(163)
			p.ExpressionList()
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
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IOperandNameContext)(nil)).Elem())
	var tst = make([]IOperandNameContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IOperandNameContext)
		}
	}

	return tst
}

func (s *IdentListContext) OperandName(i int) IOperandNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOperandNameContext)(nil)).Elem(), i)

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
	localctx = NewIdentListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, FaultParserRULE_identList)
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
		p.SetState(166)
		p.OperandName()
	}
	p.SetState(171)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == FaultParserCOMMA {
		{
			p.SetState(167)
			p.Match(FaultParserCOMMA)
		}
		{
			p.SetState(168)
			p.OperandName()
		}

		p.SetState(173)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
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
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *ExpressionListContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

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
	localctx = NewExpressionListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, FaultParserRULE_expressionList)
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
		p.SetState(174)
		p.expression(0)
	}
	p.SetState(179)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == FaultParserCOMMA {
		{
			p.SetState(175)
			p.Match(FaultParserCOMMA)
		}
		{
			p.SetState(176)
			p.expression(0)
		}

		p.SetState(181)
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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStructTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStructTypeContext)
}

func (s *StructDeclContext) Eos() IEosContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEosContext)(nil)).Elem(), 0)

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
	localctx = NewStructDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, FaultParserRULE_structDecl)

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
		p.SetState(182)
		p.Match(FaultParserDEF)
	}
	{
		p.SetState(183)
		p.Match(FaultParserIDENT)
	}
	{
		p.SetState(184)
		p.Match(FaultParserASSIGN)
	}
	{
		p.SetState(185)
		p.StructType()
	}
	{
		p.SetState(186)
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

func (s *StockContext) AllStructProperties() []IStructPropertiesContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStructPropertiesContext)(nil)).Elem())
	var tst = make([]IStructPropertiesContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStructPropertiesContext)
		}
	}

	return tst
}

func (s *StockContext) StructProperties(i int) IStructPropertiesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStructPropertiesContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStructPropertiesContext)
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

func (s *FlowContext) AllStructProperties() []IStructPropertiesContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStructPropertiesContext)(nil)).Elem())
	var tst = make([]IStructPropertiesContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStructPropertiesContext)
		}
	}

	return tst
}

func (s *FlowContext) StructProperties(i int) IStructPropertiesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStructPropertiesContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStructPropertiesContext)
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
	localctx = NewStructTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, FaultParserRULE_structType)
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

	p.SetState(210)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case FaultParserFLOW:
		localctx = NewFlowContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(188)
			p.Match(FaultParserFLOW)
		}
		{
			p.SetState(189)
			p.Match(FaultParserLCURLY)
		}
		p.SetState(195)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == FaultParserIDENT {
			{
				p.SetState(190)
				p.StructProperties()
			}
			{
				p.SetState(191)
				p.Match(FaultParserCOMMA)
			}

			p.SetState(197)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(198)
			p.Match(FaultParserRCURLY)
		}

	case FaultParserSTOCK:
		localctx = NewStockContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(199)
			p.Match(FaultParserSTOCK)
		}
		{
			p.SetState(200)
			p.Match(FaultParserLCURLY)
		}
		p.SetState(206)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == FaultParserIDENT {
			{
				p.SetState(201)
				p.StructProperties()
			}
			{
				p.SetState(202)
				p.Match(FaultParserCOMMA)
			}

			p.SetState(208)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(209)
			p.Match(FaultParserRCURLY)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISolvableContext)(nil)).Elem(), 0)

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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INumericContext)(nil)).Elem(), 0)

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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IString_Context)(nil)).Elem(), 0)

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

type PropFuncContext struct {
	*StructPropertiesContext
}

func NewPropFuncContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PropFuncContext {
	var p = new(PropFuncContext)

	p.StructPropertiesContext = NewEmptyStructPropertiesContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StructPropertiesContext))

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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionLitContext)(nil)).Elem(), 0)

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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOperandNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOperandNameContext)
}

func (s *PropVarContext) Prefix() IPrefixContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrefixContext)(nil)).Elem(), 0)

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
	localctx = NewStructPropertiesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, FaultParserRULE_structProperties)

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

	p.SetState(231)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext()) {
	case 1:
		localctx = NewPropIntContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(212)
			p.Match(FaultParserIDENT)
		}
		{
			p.SetState(213)
			p.Match(FaultParserCOLON)
		}
		{
			p.SetState(214)
			p.Numeric()
		}

	case 2:
		localctx = NewPropStringContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(215)
			p.Match(FaultParserIDENT)
		}
		{
			p.SetState(216)
			p.Match(FaultParserCOLON)
		}
		{
			p.SetState(217)
			p.String_()
		}

	case 3:
		localctx = NewPropFuncContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(218)
			p.Match(FaultParserIDENT)
		}
		{
			p.SetState(219)
			p.Match(FaultParserCOLON)
		}
		{
			p.SetState(220)
			p.FunctionLit()
		}

	case 4:
		localctx = NewPropVarContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(221)
			p.Match(FaultParserIDENT)
		}
		{
			p.SetState(222)
			p.Match(FaultParserCOLON)
		}
		{
			p.SetState(223)
			p.OperandName()
		}

	case 5:
		localctx = NewPropVarContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(224)
			p.Match(FaultParserIDENT)
		}
		{
			p.SetState(225)
			p.Match(FaultParserCOLON)
		}
		{
			p.SetState(226)
			p.Prefix()
		}

	case 6:
		localctx = NewPropSolvableContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(227)
			p.Match(FaultParserIDENT)
		}
		{
			p.SetState(228)
			p.Match(FaultParserCOLON)
		}
		{
			p.SetState(229)
			p.Solvable()
		}

	case 7:
		localctx = NewPropSolvableContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(230)
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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOperandContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOperandContext)
}

func (s *InitDeclContext) Eos() IEosContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEosContext)(nil)).Elem(), 0)

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
	localctx = NewInitDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, FaultParserRULE_initDecl)

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
		p.SetState(233)
		p.Match(FaultParserINIT)
	}
	{
		p.SetState(234)
		p.Operand()
	}
	{
		p.SetState(235)
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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementListContext)(nil)).Elem(), 0)

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
	localctx = NewBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, FaultParserRULE_block)

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
		p.SetState(237)
		p.Match(FaultParserLCURLY)
	}
	p.SetState(239)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 16, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(238)
			p.StatementList()
		}

	}
	{
		p.SetState(241)
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
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatementContext)(nil)).Elem())
	var tst = make([]IStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatementContext)
		}
	}

	return tst
}

func (s *StatementListContext) Statement(i int) IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), i)

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
	localctx = NewStatementListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, FaultParserRULE_statementList)

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
	p.SetState(244)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(243)
				p.Statement()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(246)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 17, p.GetParserRuleContext())
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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConstDeclContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConstDeclContext)
}

func (s *StatementContext) InitDecl() IInitDeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInitDeclContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInitDeclContext)
}

func (s *StatementContext) SimpleStmt() ISimpleStmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISimpleStmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISimpleStmtContext)
}

func (s *StatementContext) Eos() IEosContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEosContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEosContext)
}

func (s *StatementContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *StatementContext) IfStmt() IIfStmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIfStmtContext)(nil)).Elem(), 0)

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
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, FaultParserRULE_statement)

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

	p.SetState(255)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 18, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(248)
			p.ConstDecl()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(249)
			p.InitDecl()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(250)
			p.SimpleStmt()
		}
		{
			p.SetState(251)
			p.Eos()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(253)
			p.Block()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(254)
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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *SimpleStmtContext) IncDecStmt() IIncDecStmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIncDecStmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIncDecStmtContext)
}

func (s *SimpleStmtContext) Assignment() IAssignmentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAssignmentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAssignmentContext)
}

func (s *SimpleStmtContext) EmptyStmt() IEmptyStmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEmptyStmtContext)(nil)).Elem(), 0)

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
	localctx = NewSimpleStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, FaultParserRULE_simpleStmt)

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

	p.SetState(261)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 19, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(257)
			p.expression(0)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(258)
			p.IncDecStmt()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(259)
			p.Assignment()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(260)
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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

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
	localctx = NewIncDecStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, FaultParserRULE_incDecStmt)
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
		p.SetState(263)
		p.expression(0)
	}
	{
		p.SetState(264)
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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOperandNameContext)(nil)).Elem(), 0)

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
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *AccessHistoryContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

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
	localctx = NewAccessHistoryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, FaultParserRULE_accessHistory)

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
		p.SetState(266)
		p.OperandName()
	}
	p.SetState(271)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(267)
				p.Match(FaultParserLBRACE)
			}
			{
				p.SetState(268)
				p.expression(0)
			}
			{
				p.SetState(269)
				p.Match(FaultParserRBRACE)
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(273)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext())
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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInvariantContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInvariantContext)
}

func (s *AssertionContext) Eos() IEosContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEosContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEosContext)
}

func (s *AssertionContext) Temporal() ITemporalContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITemporalContext)(nil)).Elem(), 0)

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
	localctx = NewAssertionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, FaultParserRULE_assertion)

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
		p.SetState(275)
		p.Match(FaultParserASSERT)
	}
	p.SetState(277)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 21, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(276)
			p.Temporal()
		}

	}
	{
		p.SetState(279)
		p.Invariant()
	}
	{
		p.SetState(280)
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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInvariantContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInvariantContext)
}

func (s *AssumptionContext) Eos() IEosContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEosContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEosContext)
}

func (s *AssumptionContext) Temporal() ITemporalContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITemporalContext)(nil)).Elem(), 0)

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
	localctx = NewAssumptionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, FaultParserRULE_assumption)

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
		p.SetState(282)
		p.Match(FaultParserASSUME)
	}
	p.SetState(284)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 22, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(283)
			p.Temporal()
		}

	}
	{
		p.SetState(286)
		p.Invariant()
	}
	{
		p.SetState(287)
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
	localctx = NewTemporalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, FaultParserRULE_temporal)
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
		p.SetState(289)
		_la = p.GetTokenStream().LA(1)

		if !(_la == FaultParserEVENTUALLY || _la == FaultParserALWAYS) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
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

func (s *InvariantContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *InvariantContext) Integer() IIntegerContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIntegerContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIntegerContext)
}

func (s *InvariantContext) NMT() antlr.TerminalNode {
	return s.GetToken(FaultParserNMT, 0)
}

func (s *InvariantContext) NFT() antlr.TerminalNode {
	return s.GetToken(FaultParserNFT, 0)
}

func (s *InvariantContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InvariantContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InvariantContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaultParserListener); ok {
		listenerT.EnterInvariant(s)
	}
}

func (s *InvariantContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaultParserListener); ok {
		listenerT.ExitInvariant(s)
	}
}

func (s *InvariantContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FaultParserVisitor:
		return t.VisitInvariant(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FaultParser) Invariant() (localctx IInvariantContext) {
	localctx = NewInvariantContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, FaultParserRULE_invariant)
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

	p.SetState(296)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 23, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(291)
			p.expression(0)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(292)
			p.expression(0)
		}
		{
			p.SetState(293)
			_la = p.GetTokenStream().LA(1)

			if !(_la == FaultParserNMT || _la == FaultParserNFT) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(294)
			p.Integer()
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
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionListContext)(nil)).Elem())
	var tst = make([]IExpressionListContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionListContext)
		}
	}

	return tst
}

func (s *MiscAssignContext) ExpressionList(i int) IExpressionListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionListContext)(nil)).Elem(), i)

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
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionListContext)(nil)).Elem())
	var tst = make([]IExpressionListContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionListContext)
		}
	}

	return tst
}

func (s *FaultAssignContext) ExpressionList(i int) IExpressionListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionListContext)(nil)).Elem(), i)

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
	localctx = NewAssignmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, FaultParserRULE_assignment)
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

	p.SetState(309)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 25, p.GetParserRuleContext()) {
	case 1:
		localctx = NewMiscAssignContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(298)
			p.ExpressionList()
		}
		p.SetState(300)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if ((_la-50)&-(0x1f+1)) == 0 && ((1<<uint((_la-50)))&((1<<(FaultParserAMPERSAND-50))|(1<<(FaultParserPLUS-50))|(1<<(FaultParserMINUS-50))|(1<<(FaultParserCARET-50))|(1<<(FaultParserMULTI-50))|(1<<(FaultParserDIV-50))|(1<<(FaultParserMOD-50))|(1<<(FaultParserLSHIFT-50))|(1<<(FaultParserRSHIFT-50))|(1<<(FaultParserBIT_CLEAR-50)))) != 0 {
			{
				p.SetState(299)
				_la = p.GetTokenStream().LA(1)

				if !(((_la-50)&-(0x1f+1)) == 0 && ((1<<uint((_la-50)))&((1<<(FaultParserAMPERSAND-50))|(1<<(FaultParserPLUS-50))|(1<<(FaultParserMINUS-50))|(1<<(FaultParserCARET-50))|(1<<(FaultParserMULTI-50))|(1<<(FaultParserDIV-50))|(1<<(FaultParserMOD-50))|(1<<(FaultParserLSHIFT-50))|(1<<(FaultParserRSHIFT-50))|(1<<(FaultParserBIT_CLEAR-50)))) != 0) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}

		}
		{
			p.SetState(302)
			p.Match(FaultParserASSIGN)
		}
		{
			p.SetState(303)
			p.ExpressionList()
		}

	case 2:
		localctx = NewFaultAssignContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(305)
			p.ExpressionList()
		}
		{
			p.SetState(306)
			_la = p.GetTokenStream().LA(1)

			if !(_la == FaultParserASSIGN_FLOW1 || _la == FaultParserASSIGN_FLOW2) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(307)
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
	localctx = NewEmptyStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, FaultParserRULE_emptyStmt)

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
		p.SetState(311)
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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *IfStmtContext) AllBlock() []IBlockContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IBlockContext)(nil)).Elem())
	var tst = make([]IBlockContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IBlockContext)
		}
	}

	return tst
}

func (s *IfStmtContext) Block(i int) IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *IfStmtContext) SimpleStmt() ISimpleStmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISimpleStmtContext)(nil)).Elem(), 0)

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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIfStmtContext)(nil)).Elem(), 0)

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
	localctx = NewIfStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, FaultParserRULE_ifStmt)

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
		p.SetState(313)
		p.Match(FaultParserIF)
	}
	p.SetState(317)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 26, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(314)
			p.SimpleStmt()
		}
		{
			p.SetState(315)
			p.Match(FaultParserSEMI)
		}

	}
	{
		p.SetState(319)
		p.expression(0)
	}
	{
		p.SetState(320)
		p.Block()
	}
	p.SetState(326)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 28, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(321)
			p.Match(FaultParserELSE)
		}
		p.SetState(324)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case FaultParserIF:
			{
				p.SetState(322)
				p.IfStmt()
			}

		case FaultParserLCURLY:
			{
				p.SetState(323)
				p.Block()
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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRoundsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRoundsContext)
}

func (s *ForStmtContext) RUN() antlr.TerminalNode {
	return s.GetToken(FaultParserRUN, 0)
}

func (s *ForStmtContext) RunBlock() IRunBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRunBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRunBlockContext)
}

func (s *ForStmtContext) Eos() IEosContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEosContext)(nil)).Elem(), 0)

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
	localctx = NewForStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, FaultParserRULE_forStmt)

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
		p.SetState(328)
		p.Match(FaultParserFOR)
	}
	{
		p.SetState(329)
		p.Rounds()
	}
	{
		p.SetState(330)
		p.Match(FaultParserRUN)
	}
	{
		p.SetState(331)
		p.RunBlock()
	}
	{
		p.SetState(332)
		p.Eos()
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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIntegerContext)(nil)).Elem(), 0)

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
	localctx = NewRoundsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, FaultParserRULE_rounds)

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
		p.SetState(334)
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

func (s *ParamCallContext) AllIDENT() []antlr.TerminalNode {
	return s.GetTokens(FaultParserIDENT)
}

func (s *ParamCallContext) IDENT(i int) antlr.TerminalNode {
	return s.GetToken(FaultParserIDENT, i)
}

func (s *ParamCallContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(FaultParserDOT)
}

func (s *ParamCallContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(FaultParserDOT, i)
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
	localctx = NewParamCallContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, FaultParserRULE_paramCall)

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
		p.SetState(336)
		p.Match(FaultParserIDENT)
	}
	{
		p.SetState(337)
		p.Match(FaultParserDOT)
	}
	{
		p.SetState(338)
		p.Match(FaultParserIDENT)
	}
	p.SetState(343)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 29, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(339)
				p.Match(FaultParserDOT)
			}
			{
				p.SetState(340)
				p.Match(FaultParserIDENT)
			}

		}
		p.SetState(345)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 29, p.GetParserRuleContext())
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
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IRunStepContext)(nil)).Elem())
	var tst = make([]IRunStepContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IRunStepContext)
		}
	}

	return tst
}

func (s *RunBlockContext) RunStep(i int) IRunStepContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRunStepContext)(nil)).Elem(), i)

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
	localctx = NewRunBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, FaultParserRULE_runBlock)

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
		p.SetState(346)
		p.Match(FaultParserLCURLY)
	}
	p.SetState(350)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 30, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(347)
				p.RunStep()
			}

		}
		p.SetState(352)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 30, p.GetParserRuleContext())
	}
	{
		p.SetState(353)
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
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IParamCallContext)(nil)).Elem())
	var tst = make([]IParamCallContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IParamCallContext)
		}
	}

	return tst
}

func (s *RunStepExprContext) ParamCall(i int) IParamCallContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParamCallContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IParamCallContext)
}

func (s *RunStepExprContext) Eos() IEosContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEosContext)(nil)).Elem(), 0)

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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEosContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEosContext)
}

func (s *RunInitContext) DOT() antlr.TerminalNode {
	return s.GetToken(FaultParserDOT, 0)
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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISimpleStmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISimpleStmtContext)
}

func (s *RunExprContext) Eos() IEosContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEosContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEosContext)
}

func (s *RunExprContext) IfStmt() IIfStmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIfStmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIfStmtContext)
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
	localctx = NewRunStepContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, FaultParserRULE_runStep)
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

	p.SetState(378)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 33, p.GetParserRuleContext()) {
	case 1:
		localctx = NewRunStepExprContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(355)
			p.ParamCall()
		}
		p.SetState(360)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == FaultParserPIPE {
			{
				p.SetState(356)
				p.Match(FaultParserPIPE)
			}
			{
				p.SetState(357)
				p.ParamCall()
			}

			p.SetState(362)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(363)
			p.Eos()
		}

	case 2:
		localctx = NewRunInitContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(365)
			p.Match(FaultParserIDENT)
		}
		{
			p.SetState(366)
			p.Match(FaultParserASSIGN)
		}
		{
			p.SetState(367)
			p.Match(FaultParserNEW)
		}
		{
			p.SetState(368)
			p.Match(FaultParserIDENT)
		}
		p.SetState(371)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == FaultParserDOT {
			{
				p.SetState(369)
				p.Match(FaultParserDOT)
			}
			{
				p.SetState(370)
				p.Match(FaultParserIDENT)
			}

		}
		{
			p.SetState(373)
			p.Eos()
		}

	case 3:
		localctx = NewRunExprContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(374)
			p.SimpleStmt()
		}
		{
			p.SetState(375)
			p.Eos()
		}

	case 4:
		localctx = NewRunExprContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(377)
			p.IfStmt()
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
	localctx = NewFaultTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, FaultParserRULE_faultType)
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
		p.SetState(380)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-27)&-(0x1f+1)) == 0 && ((1<<uint((_la-27)))&((1<<(FaultParserTY_STRING-27))|(1<<(FaultParserTY_BOOL-27))|(1<<(FaultParserTY_INT-27))|(1<<(FaultParserTY_FLOAT-27))|(1<<(FaultParserTY_NATURAL-27))|(1<<(FaultParserTY_UNCERTAIN-27))|(1<<(FaultParserTY_UNKNOWN-27)))) != 0) {
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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFaultTypeContext)(nil)).Elem(), 0)

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
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IOperandContext)(nil)).Elem())
	var tst = make([]IOperandContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IOperandContext)
		}
	}

	return tst
}

func (s *SolvableContext) Operand(i int) IOperandContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOperandContext)(nil)).Elem(), i)

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
	localctx = NewSolvableContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, FaultParserRULE_solvable)
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
		p.SetState(382)
		p.FaultType()
	}
	{
		p.SetState(383)
		p.Match(FaultParserLPAREN)
	}
	p.SetState(385)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<FaultParserCLOCK)|(1<<FaultParserNEW)|(1<<FaultParserTHIS)|(1<<FaultParserNIL)|(1<<FaultParserTRUE)|(1<<FaultParserFALSE))) != 0) || (((_la-34)&-(0x1f+1)) == 0 && ((1<<uint((_la-34)))&((1<<(FaultParserIDENT-34))|(1<<(FaultParserLPAREN-34))|(1<<(FaultParserMINUS-34)))) != 0) || (((_la-71)&-(0x1f+1)) == 0 && ((1<<uint((_la-71)))&((1<<(FaultParserDECIMAL_LIT-71))|(1<<(FaultParserOCTAL_LIT-71))|(1<<(FaultParserHEX_LIT-71))|(1<<(FaultParserFLOAT_LIT-71))|(1<<(FaultParserRAW_STRING_LIT-71))|(1<<(FaultParserINTERPRETED_STRING_LIT-71)))) != 0) {
		{
			p.SetState(384)
			p.Operand()
		}

	}
	p.SetState(391)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == FaultParserCOMMA {
		{
			p.SetState(387)
			p.Match(FaultParserCOMMA)
		}
		{
			p.SetState(388)
			p.Operand()
		}

		p.SetState(393)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(394)
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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISolvableContext)(nil)).Elem(), 0)

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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOperandContext)(nil)).Elem(), 0)

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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrefixContext)(nil)).Elem(), 0)

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
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *LrExprContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

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
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 70
	p.EnterRecursionRule(localctx, 70, FaultParserRULE_expression, _p)
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
	p.SetState(400)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 36, p.GetParserRuleContext()) {
	case 1:
		localctx = NewExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(397)
			p.Operand()
		}

	case 2:
		localctx = NewTypedContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(398)
			p.Solvable()
		}

	case 3:
		localctx = NewExprPrefixContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(399)
			p.Prefix()
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(422)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 38, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(420)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 37, p.GetParserRuleContext()) {
			case 1:
				localctx = NewLrExprContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, FaultParserRULE_expression)
				p.SetState(402)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
				}
				{
					p.SetState(403)
					p.Match(FaultParserEXPO)
				}
				{
					p.SetState(404)
					p.expression(7)
				}

			case 2:
				localctx = NewLrExprContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, FaultParserRULE_expression)
				p.SetState(405)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				{
					p.SetState(406)
					_la = p.GetTokenStream().LA(1)

					if !(((_la-50)&-(0x1f+1)) == 0 && ((1<<uint((_la-50)))&((1<<(FaultParserAMPERSAND-50))|(1<<(FaultParserMULTI-50))|(1<<(FaultParserDIV-50))|(1<<(FaultParserMOD-50))|(1<<(FaultParserLSHIFT-50))|(1<<(FaultParserRSHIFT-50))|(1<<(FaultParserBIT_CLEAR-50)))) != 0) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(407)
					p.expression(6)
				}

			case 3:
				localctx = NewLrExprContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, FaultParserRULE_expression)
				p.SetState(408)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(409)
					_la = p.GetTokenStream().LA(1)

					if !(((_la-61)&-(0x1f+1)) == 0 && ((1<<uint((_la-61)))&((1<<(FaultParserPLUS-61))|(1<<(FaultParserMINUS-61))|(1<<(FaultParserCARET-61)))) != 0) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(410)
					p.expression(5)
				}

			case 4:
				localctx = NewLrExprContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, FaultParserRULE_expression)
				p.SetState(411)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(412)
					_la = p.GetTokenStream().LA(1)

					if !(((_la-53)&-(0x1f+1)) == 0 && ((1<<uint((_la-53)))&((1<<(FaultParserEQUALS-53))|(1<<(FaultParserNOT_EQUALS-53))|(1<<(FaultParserLESS-53))|(1<<(FaultParserLESS_OR_EQUALS-53))|(1<<(FaultParserGREATER-53))|(1<<(FaultParserGREATER_OR_EQUALS-53)))) != 0) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(413)
					p.expression(4)
				}

			case 5:
				localctx = NewLrExprContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, FaultParserRULE_expression)
				p.SetState(414)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				{
					p.SetState(415)
					p.Match(FaultParserAND)
				}
				{
					p.SetState(416)
					p.expression(3)
				}

			case 6:
				localctx = NewLrExprContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, FaultParserRULE_expression)
				p.SetState(417)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
				}
				{
					p.SetState(418)
					p.Match(FaultParserOR)
				}
				{
					p.SetState(419)
					p.expression(2)
				}

			}

		}
		p.SetState(424)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 38, p.GetParserRuleContext())
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

func (s *OperandContext) NIL() antlr.TerminalNode {
	return s.GetToken(FaultParserNIL, 0)
}

func (s *OperandContext) Numeric() INumericContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INumericContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INumericContext)
}

func (s *OperandContext) String_() IString_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IString_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IString_Context)
}

func (s *OperandContext) Bool_() IBool_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBool_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBool_Context)
}

func (s *OperandContext) OperandName() IOperandNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOperandNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOperandNameContext)
}

func (s *OperandContext) AccessHistory() IAccessHistoryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAccessHistoryContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAccessHistoryContext)
}

func (s *OperandContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(FaultParserLPAREN, 0)
}

func (s *OperandContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

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
	localctx = NewOperandContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, FaultParserRULE_operand)

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
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 39, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(425)
			p.Match(FaultParserNIL)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(426)
			p.Numeric()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(427)
			p.String_()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(428)
			p.Bool_()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(429)
			p.OperandName()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(430)
			p.AccessHistory()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(431)
			p.Match(FaultParserLPAREN)
		}
		{
			p.SetState(432)
			p.expression(0)
		}
		{
			p.SetState(433)
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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParamCallContext)(nil)).Elem(), 0)

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
	localctx = NewOperandNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, FaultParserRULE_operandName)

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

	p.SetState(447)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 41, p.GetParserRuleContext()) {
	case 1:
		localctx = NewOpNameContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(437)
			p.Match(FaultParserIDENT)
		}

	case 2:
		localctx = NewOpParamContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(438)
			p.ParamCall()
		}

	case 3:
		localctx = NewOpThisContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(439)
			p.Match(FaultParserTHIS)
		}

	case 4:
		localctx = NewOpClockContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(440)
			p.Match(FaultParserCLOCK)
		}

	case 5:
		localctx = NewOpInstanceContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(441)
			p.Match(FaultParserNEW)
		}
		{
			p.SetState(442)
			p.Match(FaultParserIDENT)
		}
		p.SetState(445)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 40, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(443)
				p.Match(FaultParserDOT)
			}
			{
				p.SetState(444)
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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

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
	localctx = NewPrefixContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, FaultParserRULE_prefix)
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

	p.SetState(452)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 42, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(450)
			_la = p.GetTokenStream().LA(1)

			if !(((_la-50)&-(0x1f+1)) == 0 && ((1<<uint((_la-50)))&((1<<(FaultParserAMPERSAND-50))|(1<<(FaultParserBANG-50))|(1<<(FaultParserPLUS-50))|(1<<(FaultParserMINUS-50))|(1<<(FaultParserCARET-50))|(1<<(FaultParserMULTI-50)))) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(451)
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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIntegerContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIntegerContext)
}

func (s *NumericContext) Negative() INegativeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INegativeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INegativeContext)
}

func (s *NumericContext) Float_() IFloat_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFloat_Context)(nil)).Elem(), 0)

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
	localctx = NewNumericContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, FaultParserRULE_numeric)

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

	p.SetState(457)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case FaultParserDECIMAL_LIT, FaultParserOCTAL_LIT, FaultParserHEX_LIT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(454)
			p.Integer()
		}

	case FaultParserMINUS:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(455)
			p.Negative()
		}

	case FaultParserFLOAT_LIT:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(456)
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
	localctx = NewIntegerContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 80, FaultParserRULE_integer)
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
		p.SetState(459)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-71)&-(0x1f+1)) == 0 && ((1<<uint((_la-71)))&((1<<(FaultParserDECIMAL_LIT-71))|(1<<(FaultParserOCTAL_LIT-71))|(1<<(FaultParserHEX_LIT-71)))) != 0) {
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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIntegerContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIntegerContext)
}

func (s *NegativeContext) Float_() IFloat_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFloat_Context)(nil)).Elem(), 0)

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
	localctx = NewNegativeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 82, FaultParserRULE_negative)

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

	p.SetState(465)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 44, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(461)
			p.Match(FaultParserMINUS)
		}
		{
			p.SetState(462)
			p.Integer()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(463)
			p.Match(FaultParserMINUS)
		}
		{
			p.SetState(464)
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
	localctx = NewFloat_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 84, FaultParserRULE_float_)

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
		p.SetState(467)
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
	localctx = NewString_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 86, FaultParserRULE_string_)
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
		p.SetState(469)
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
	localctx = NewBool_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 88, FaultParserRULE_bool_)
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
		p.SetState(471)
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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

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
	localctx = NewFunctionLitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 90, FaultParserRULE_functionLit)

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
		p.SetState(473)
		p.Match(FaultParserFUNC)
	}
	{
		p.SetState(474)
		p.Block()
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

func (s *EosContext) EOF() antlr.TerminalNode {
	return s.GetToken(FaultParserEOF, 0)
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
	localctx = NewEosContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 92, FaultParserRULE_eos)
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
		p.SetState(476)
		_la = p.GetTokenStream().LA(1)

		if !(_la == FaultParserEOF || _la == FaultParserSEMI) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

func (p *FaultParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 35:
		var t *ExpressionContext = nil
		if localctx != nil {
			t = localctx.(*ExpressionContext)
		}
		return p.Expression_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *FaultParser) Expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 2)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
