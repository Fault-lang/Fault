// Code generated from FaultParser.g4 by ANTLR 4.9.3. DO NOT EDIT.

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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 91, 591,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4,
	29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33, 4, 34,
	9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4, 39, 9,
	39, 4, 40, 9, 40, 4, 41, 9, 41, 4, 42, 9, 42, 4, 43, 9, 43, 4, 44, 9, 44,
	4, 45, 9, 45, 4, 46, 9, 46, 4, 47, 9, 47, 4, 48, 9, 48, 4, 49, 9, 49, 4,
	50, 9, 50, 4, 51, 9, 51, 4, 52, 9, 52, 4, 53, 9, 53, 4, 54, 9, 54, 4, 55,
	9, 55, 4, 56, 9, 56, 4, 57, 9, 57, 3, 2, 3, 2, 7, 2, 117, 10, 2, 12, 2,
	14, 2, 120, 11, 2, 3, 2, 7, 2, 123, 10, 2, 12, 2, 14, 2, 126, 11, 2, 3,
	2, 7, 2, 129, 10, 2, 12, 2, 14, 2, 132, 11, 2, 3, 2, 7, 2, 135, 10, 2,
	12, 2, 14, 2, 138, 11, 2, 3, 2, 3, 2, 5, 2, 142, 10, 2, 3, 2, 5, 2, 145,
	10, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4,
	3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 7, 5, 166, 10, 5, 12, 5,
	14, 5, 169, 11, 5, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6, 5, 6, 176, 10, 6, 3, 7,
	3, 7, 3, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 7, 8, 188, 10, 8,
	12, 8, 14, 8, 191, 11, 8, 3, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3,
	10, 3, 10, 7, 10, 202, 10, 10, 12, 10, 14, 10, 205, 11, 10, 3, 10, 5, 10,
	208, 10, 10, 3, 10, 3, 10, 3, 11, 3, 11, 3, 11, 3, 11, 3, 12, 3, 12, 3,
	12, 3, 12, 7, 12, 220, 10, 12, 12, 12, 14, 12, 223, 11, 12, 3, 12, 5, 12,
	226, 10, 12, 3, 12, 3, 12, 3, 13, 5, 13, 231, 10, 13, 3, 13, 3, 13, 3,
	14, 3, 14, 3, 15, 3, 15, 3, 15, 3, 15, 5, 15, 241, 10, 15, 3, 16, 3, 16,
	3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 7, 17, 253, 10,
	17, 12, 17, 14, 17, 256, 11, 17, 3, 17, 5, 17, 259, 10, 17, 3, 18, 3, 18,
	3, 18, 5, 18, 264, 10, 18, 3, 19, 3, 19, 3, 19, 7, 19, 269, 10, 19, 12,
	19, 14, 19, 272, 11, 19, 3, 20, 3, 20, 3, 20, 7, 20, 277, 10, 20, 12, 20,
	14, 20, 280, 11, 20, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 22, 3,
	22, 3, 22, 3, 22, 3, 22, 7, 22, 293, 10, 22, 12, 22, 14, 22, 296, 11, 22,
	3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 7, 22, 304, 10, 22, 12, 22, 14,
	22, 307, 11, 22, 3, 22, 5, 22, 310, 10, 22, 3, 23, 3, 23, 3, 23, 3, 23,
	3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3,
	23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 5, 23, 334, 10, 23,
	3, 24, 3, 24, 3, 24, 3, 24, 3, 25, 3, 25, 5, 25, 342, 10, 25, 3, 25, 3,
	25, 3, 26, 6, 26, 347, 10, 26, 13, 26, 14, 26, 348, 3, 27, 3, 27, 3, 27,
	3, 27, 3, 27, 3, 27, 3, 27, 5, 27, 358, 10, 27, 3, 28, 3, 28, 3, 28, 3,
	28, 3, 28, 5, 28, 365, 10, 28, 3, 29, 3, 29, 3, 29, 3, 30, 3, 30, 3, 30,
	3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 5, 30, 378, 10, 30, 3, 31, 3, 31, 3,
	31, 3, 31, 3, 31, 6, 31, 385, 10, 31, 13, 31, 14, 31, 386, 3, 32, 3, 32,
	3, 32, 5, 32, 392, 10, 32, 3, 32, 3, 32, 3, 33, 3, 33, 3, 33, 5, 33, 399,
	10, 33, 3, 33, 3, 33, 3, 34, 3, 34, 3, 34, 3, 34, 3, 34, 5, 34, 408, 10,
	34, 3, 35, 3, 35, 3, 36, 3, 36, 5, 36, 414, 10, 36, 3, 36, 3, 36, 3, 36,
	3, 36, 3, 36, 3, 36, 3, 36, 5, 36, 423, 10, 36, 3, 37, 3, 37, 3, 38, 3,
	38, 3, 38, 3, 38, 5, 38, 431, 10, 38, 3, 38, 3, 38, 3, 38, 3, 38, 3, 38,
	5, 38, 438, 10, 38, 5, 38, 440, 10, 38, 3, 39, 3, 39, 3, 39, 3, 39, 3,
	39, 3, 39, 3, 40, 3, 40, 3, 41, 3, 41, 3, 41, 3, 41, 3, 41, 7, 41, 455,
	10, 41, 12, 41, 14, 41, 458, 11, 41, 3, 42, 3, 42, 7, 42, 462, 10, 42,
	12, 42, 14, 42, 465, 11, 42, 3, 42, 3, 42, 3, 43, 3, 43, 3, 43, 7, 43,
	472, 10, 43, 12, 43, 14, 43, 475, 11, 43, 3, 43, 3, 43, 3, 43, 3, 43, 3,
	43, 3, 43, 3, 43, 5, 43, 484, 10, 43, 3, 43, 3, 43, 3, 43, 3, 43, 3, 43,
	5, 43, 491, 10, 43, 3, 44, 3, 44, 3, 45, 3, 45, 3, 45, 5, 45, 498, 10,
	45, 3, 45, 3, 45, 7, 45, 502, 10, 45, 12, 45, 14, 45, 505, 11, 45, 3, 45,
	3, 45, 3, 46, 3, 46, 3, 46, 3, 46, 5, 46, 513, 10, 46, 3, 46, 3, 46, 3,
	46, 3, 46, 3, 46, 3, 46, 3, 46, 3, 46, 3, 46, 3, 46, 3, 46, 3, 46, 3, 46,
	3, 46, 3, 46, 3, 46, 3, 46, 3, 46, 7, 46, 533, 10, 46, 12, 46, 14, 46,
	536, 11, 46, 3, 47, 3, 47, 3, 47, 3, 47, 3, 47, 3, 47, 3, 47, 3, 47, 3,
	47, 3, 47, 5, 47, 548, 10, 47, 3, 48, 3, 48, 3, 48, 3, 48, 3, 48, 3, 48,
	3, 48, 3, 48, 5, 48, 558, 10, 48, 5, 48, 560, 10, 48, 3, 49, 3, 49, 3,
	49, 5, 49, 565, 10, 49, 3, 50, 3, 50, 3, 50, 5, 50, 570, 10, 50, 3, 51,
	3, 51, 3, 52, 3, 52, 3, 52, 3, 52, 5, 52, 578, 10, 52, 3, 53, 3, 53, 3,
	54, 3, 54, 3, 55, 3, 55, 3, 56, 3, 56, 3, 56, 3, 57, 3, 57, 3, 57, 2, 3,
	90, 58, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34,
	36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68, 70,
	72, 74, 76, 78, 80, 82, 84, 86, 88, 90, 92, 94, 96, 98, 100, 102, 104,
	106, 108, 110, 112, 2, 18, 4, 2, 45, 45, 51, 51, 3, 2, 64, 69, 3, 2, 59,
	60, 3, 2, 23, 25, 3, 2, 26, 27, 5, 2, 61, 61, 72, 74, 76, 81, 3, 2, 47,
	48, 4, 2, 22, 22, 45, 45, 3, 2, 38, 44, 4, 2, 61, 61, 76, 81, 3, 2, 72,
	74, 6, 2, 61, 61, 63, 63, 72, 74, 76, 76, 3, 2, 82, 84, 3, 2, 86, 87, 3,
	2, 29, 30, 3, 3, 58, 58, 2, 620, 2, 114, 3, 2, 2, 2, 4, 148, 3, 2, 2, 2,
	6, 152, 3, 2, 2, 2, 8, 157, 3, 2, 2, 2, 10, 175, 3, 2, 2, 2, 12, 177, 3,
	2, 2, 2, 14, 182, 3, 2, 2, 2, 16, 195, 3, 2, 2, 2, 18, 199, 3, 2, 2, 2,
	20, 211, 3, 2, 2, 2, 22, 215, 3, 2, 2, 2, 24, 230, 3, 2, 2, 2, 26, 234,
	3, 2, 2, 2, 28, 240, 3, 2, 2, 2, 30, 242, 3, 2, 2, 2, 32, 244, 3, 2, 2,
	2, 34, 260, 3, 2, 2, 2, 36, 265, 3, 2, 2, 2, 38, 273, 3, 2, 2, 2, 40, 281,
	3, 2, 2, 2, 42, 309, 3, 2, 2, 2, 44, 333, 3, 2, 2, 2, 46, 335, 3, 2, 2,
	2, 48, 339, 3, 2, 2, 2, 50, 346, 3, 2, 2, 2, 52, 357, 3, 2, 2, 2, 54, 364,
	3, 2, 2, 2, 56, 366, 3, 2, 2, 2, 58, 377, 3, 2, 2, 2, 60, 379, 3, 2, 2,
	2, 62, 388, 3, 2, 2, 2, 64, 395, 3, 2, 2, 2, 66, 407, 3, 2, 2, 2, 68, 409,
	3, 2, 2, 2, 70, 422, 3, 2, 2, 2, 72, 424, 3, 2, 2, 2, 74, 426, 3, 2, 2,
	2, 76, 441, 3, 2, 2, 2, 78, 447, 3, 2, 2, 2, 80, 449, 3, 2, 2, 2, 82, 459,
	3, 2, 2, 2, 84, 490, 3, 2, 2, 2, 86, 492, 3, 2, 2, 2, 88, 494, 3, 2, 2,
	2, 90, 512, 3, 2, 2, 2, 92, 547, 3, 2, 2, 2, 94, 559, 3, 2, 2, 2, 96, 564,
	3, 2, 2, 2, 98, 569, 3, 2, 2, 2, 100, 571, 3, 2, 2, 2, 102, 577, 3, 2,
	2, 2, 104, 579, 3, 2, 2, 2, 106, 581, 3, 2, 2, 2, 108, 583, 3, 2, 2, 2,
	110, 585, 3, 2, 2, 2, 112, 588, 3, 2, 2, 2, 114, 118, 5, 4, 3, 2, 115,
	117, 5, 22, 12, 2, 116, 115, 3, 2, 2, 2, 117, 120, 3, 2, 2, 2, 118, 116,
	3, 2, 2, 2, 118, 119, 3, 2, 2, 2, 119, 124, 3, 2, 2, 2, 120, 118, 3, 2,
	2, 2, 121, 123, 5, 6, 4, 2, 122, 121, 3, 2, 2, 2, 123, 126, 3, 2, 2, 2,
	124, 122, 3, 2, 2, 2, 124, 125, 3, 2, 2, 2, 125, 130, 3, 2, 2, 2, 126,
	124, 3, 2, 2, 2, 127, 129, 5, 8, 5, 2, 128, 127, 3, 2, 2, 2, 129, 132,
	3, 2, 2, 2, 130, 128, 3, 2, 2, 2, 130, 131, 3, 2, 2, 2, 131, 136, 3, 2,
	2, 2, 132, 130, 3, 2, 2, 2, 133, 135, 5, 14, 8, 2, 134, 133, 3, 2, 2, 2,
	135, 138, 3, 2, 2, 2, 136, 134, 3, 2, 2, 2, 136, 137, 3, 2, 2, 2, 137,
	141, 3, 2, 2, 2, 138, 136, 3, 2, 2, 2, 139, 142, 5, 62, 32, 2, 140, 142,
	5, 64, 33, 2, 141, 139, 3, 2, 2, 2, 141, 140, 3, 2, 2, 2, 141, 142, 3,
	2, 2, 2, 142, 144, 3, 2, 2, 2, 143, 145, 5, 76, 39, 2, 144, 143, 3, 2,
	2, 2, 144, 145, 3, 2, 2, 2, 145, 146, 3, 2, 2, 2, 146, 147, 5, 112, 57,
	2, 147, 3, 3, 2, 2, 2, 148, 149, 7, 34, 2, 2, 149, 150, 7, 45, 2, 2, 150,
	151, 5, 112, 57, 2, 151, 5, 3, 2, 2, 2, 152, 153, 7, 33, 2, 2, 153, 154,
	7, 45, 2, 2, 154, 155, 7, 46, 2, 2, 155, 156, 5, 92, 47, 2, 156, 7, 3,
	2, 2, 2, 157, 158, 7, 8, 2, 2, 158, 159, 7, 45, 2, 2, 159, 160, 7, 46,
	2, 2, 160, 161, 7, 32, 2, 2, 161, 167, 7, 54, 2, 2, 162, 163, 5, 10, 6,
	2, 163, 164, 7, 50, 2, 2, 164, 166, 3, 2, 2, 2, 165, 162, 3, 2, 2, 2, 166,
	169, 3, 2, 2, 2, 167, 165, 3, 2, 2, 2, 167, 168, 3, 2, 2, 2, 168, 170,
	3, 2, 2, 2, 169, 167, 3, 2, 2, 2, 170, 171, 7, 55, 2, 2, 171, 172, 5, 112,
	57, 2, 172, 9, 3, 2, 2, 2, 173, 176, 5, 44, 23, 2, 174, 176, 5, 12, 7,
	2, 175, 173, 3, 2, 2, 2, 175, 174, 3, 2, 2, 2, 176, 11, 3, 2, 2, 2, 177,
	178, 7, 45, 2, 2, 178, 179, 7, 49, 2, 2, 179, 180, 7, 36, 2, 2, 180, 181,
	5, 48, 25, 2, 181, 13, 3, 2, 2, 2, 182, 183, 7, 35, 2, 2, 183, 189, 7,
	54, 2, 2, 184, 185, 5, 16, 9, 2, 185, 186, 7, 50, 2, 2, 186, 188, 3, 2,
	2, 2, 187, 184, 3, 2, 2, 2, 188, 191, 3, 2, 2, 2, 189, 187, 3, 2, 2, 2,
	189, 190, 3, 2, 2, 2, 190, 192, 3, 2, 2, 2, 191, 189, 3, 2, 2, 2, 192,
	193, 7, 55, 2, 2, 193, 194, 5, 112, 57, 2, 194, 15, 3, 2, 2, 2, 195, 196,
	7, 45, 2, 2, 196, 197, 7, 49, 2, 2, 197, 198, 7, 45, 2, 2, 198, 17, 3,
	2, 2, 2, 199, 203, 5, 20, 11, 2, 200, 202, 5, 28, 15, 2, 201, 200, 3, 2,
	2, 2, 202, 205, 3, 2, 2, 2, 203, 201, 3, 2, 2, 2, 203, 204, 3, 2, 2, 2,
	204, 207, 3, 2, 2, 2, 205, 203, 3, 2, 2, 2, 206, 208, 5, 76, 39, 2, 207,
	206, 3, 2, 2, 2, 207, 208, 3, 2, 2, 2, 208, 209, 3, 2, 2, 2, 209, 210,
	5, 112, 57, 2, 210, 19, 3, 2, 2, 2, 211, 212, 7, 19, 2, 2, 212, 213, 7,
	45, 2, 2, 213, 214, 5, 112, 57, 2, 214, 21, 3, 2, 2, 2, 215, 225, 7, 14,
	2, 2, 216, 226, 5, 24, 13, 2, 217, 221, 7, 52, 2, 2, 218, 220, 5, 24, 13,
	2, 219, 218, 3, 2, 2, 2, 220, 223, 3, 2, 2, 2, 221, 219, 3, 2, 2, 2, 221,
	222, 3, 2, 2, 2, 222, 224, 3, 2, 2, 2, 223, 221, 3, 2, 2, 2, 224, 226,
	7, 53, 2, 2, 225, 216, 3, 2, 2, 2, 225, 217, 3, 2, 2, 2, 226, 227, 3, 2,
	2, 2, 227, 228, 5, 112, 57, 2, 228, 23, 3, 2, 2, 2, 229, 231, 9, 2, 2,
	2, 230, 229, 3, 2, 2, 2, 230, 231, 3, 2, 2, 2, 231, 232, 3, 2, 2, 2, 232,
	233, 5, 26, 14, 2, 233, 25, 3, 2, 2, 2, 234, 235, 5, 106, 54, 2, 235, 27,
	3, 2, 2, 2, 236, 241, 5, 32, 17, 2, 237, 241, 5, 40, 21, 2, 238, 241, 5,
	62, 32, 2, 239, 241, 5, 64, 33, 2, 240, 236, 3, 2, 2, 2, 240, 237, 3, 2,
	2, 2, 240, 238, 3, 2, 2, 2, 240, 239, 3, 2, 2, 2, 241, 29, 3, 2, 2, 2,
	242, 243, 9, 3, 2, 2, 243, 31, 3, 2, 2, 2, 244, 258, 7, 7, 2, 2, 245, 246,
	5, 34, 18, 2, 246, 247, 5, 112, 57, 2, 247, 259, 3, 2, 2, 2, 248, 254,
	7, 52, 2, 2, 249, 250, 5, 34, 18, 2, 250, 251, 5, 112, 57, 2, 251, 253,
	3, 2, 2, 2, 252, 249, 3, 2, 2, 2, 253, 256, 3, 2, 2, 2, 254, 252, 3, 2,
	2, 2, 254, 255, 3, 2, 2, 2, 255, 257, 3, 2, 2, 2, 256, 254, 3, 2, 2, 2,
	257, 259, 7, 53, 2, 2, 258, 245, 3, 2, 2, 2, 258, 248, 3, 2, 2, 2, 259,
	33, 3, 2, 2, 2, 260, 263, 5, 36, 19, 2, 261, 262, 7, 46, 2, 2, 262, 264,
	5, 38, 20, 2, 263, 261, 3, 2, 2, 2, 263, 264, 3, 2, 2, 2, 264, 35, 3, 2,
	2, 2, 265, 270, 5, 94, 48, 2, 266, 267, 7, 50, 2, 2, 267, 269, 5, 94, 48,
	2, 268, 266, 3, 2, 2, 2, 269, 272, 3, 2, 2, 2, 270, 268, 3, 2, 2, 2, 270,
	271, 3, 2, 2, 2, 271, 37, 3, 2, 2, 2, 272, 270, 3, 2, 2, 2, 273, 278, 5,
	90, 46, 2, 274, 275, 7, 50, 2, 2, 275, 277, 5, 90, 46, 2, 276, 274, 3,
	2, 2, 2, 277, 280, 3, 2, 2, 2, 278, 276, 3, 2, 2, 2, 278, 279, 3, 2, 2,
	2, 279, 39, 3, 2, 2, 2, 280, 278, 3, 2, 2, 2, 281, 282, 7, 8, 2, 2, 282,
	283, 7, 45, 2, 2, 283, 284, 7, 46, 2, 2, 284, 285, 5, 42, 22, 2, 285, 286,
	5, 112, 57, 2, 286, 41, 3, 2, 2, 2, 287, 288, 7, 10, 2, 2, 288, 294, 7,
	54, 2, 2, 289, 290, 5, 44, 23, 2, 290, 291, 7, 50, 2, 2, 291, 293, 3, 2,
	2, 2, 292, 289, 3, 2, 2, 2, 293, 296, 3, 2, 2, 2, 294, 292, 3, 2, 2, 2,
	294, 295, 3, 2, 2, 2, 295, 297, 3, 2, 2, 2, 296, 294, 3, 2, 2, 2, 297,
	310, 7, 55, 2, 2, 298, 299, 7, 20, 2, 2, 299, 305, 7, 54, 2, 2, 300, 301,
	5, 44, 23, 2, 301, 302, 7, 50, 2, 2, 302, 304, 3, 2, 2, 2, 303, 300, 3,
	2, 2, 2, 304, 307, 3, 2, 2, 2, 305, 303, 3, 2, 2, 2, 305, 306, 3, 2, 2,
	2, 306, 308, 3, 2, 2, 2, 307, 305, 3, 2, 2, 2, 308, 310, 7, 55, 2, 2, 309,
	287, 3, 2, 2, 2, 309, 298, 3, 2, 2, 2, 310, 43, 3, 2, 2, 2, 311, 312, 7,
	45, 2, 2, 312, 313, 7, 49, 2, 2, 313, 334, 5, 98, 50, 2, 314, 315, 7, 45,
	2, 2, 315, 316, 7, 49, 2, 2, 316, 334, 5, 106, 54, 2, 317, 318, 7, 45,
	2, 2, 318, 319, 7, 49, 2, 2, 319, 334, 5, 108, 55, 2, 320, 321, 7, 45,
	2, 2, 321, 322, 7, 49, 2, 2, 322, 334, 5, 110, 56, 2, 323, 324, 7, 45,
	2, 2, 324, 325, 7, 49, 2, 2, 325, 334, 5, 94, 48, 2, 326, 327, 7, 45, 2,
	2, 327, 328, 7, 49, 2, 2, 328, 334, 5, 96, 49, 2, 329, 330, 7, 45, 2, 2,
	330, 331, 7, 49, 2, 2, 331, 334, 5, 88, 45, 2, 332, 334, 7, 45, 2, 2, 333,
	311, 3, 2, 2, 2, 333, 314, 3, 2, 2, 2, 333, 317, 3, 2, 2, 2, 333, 320,
	3, 2, 2, 2, 333, 323, 3, 2, 2, 2, 333, 326, 3, 2, 2, 2, 333, 329, 3, 2,
	2, 2, 333, 332, 3, 2, 2, 2, 334, 45, 3, 2, 2, 2, 335, 336, 7, 15, 2, 2,
	336, 337, 5, 92, 47, 2, 337, 338, 5, 112, 57, 2, 338, 47, 3, 2, 2, 2, 339,
	341, 7, 54, 2, 2, 340, 342, 5, 50, 26, 2, 341, 340, 3, 2, 2, 2, 341, 342,
	3, 2, 2, 2, 342, 343, 3, 2, 2, 2, 343, 344, 7, 55, 2, 2, 344, 49, 3, 2,
	2, 2, 345, 347, 5, 52, 27, 2, 346, 345, 3, 2, 2, 2, 347, 348, 3, 2, 2,
	2, 348, 346, 3, 2, 2, 2, 348, 349, 3, 2, 2, 2, 349, 51, 3, 2, 2, 2, 350,
	358, 5, 32, 17, 2, 351, 358, 5, 46, 24, 2, 352, 353, 5, 54, 28, 2, 353,
	354, 5, 112, 57, 2, 354, 358, 3, 2, 2, 2, 355, 358, 5, 48, 25, 2, 356,
	358, 5, 74, 38, 2, 357, 350, 3, 2, 2, 2, 357, 351, 3, 2, 2, 2, 357, 352,
	3, 2, 2, 2, 357, 355, 3, 2, 2, 2, 357, 356, 3, 2, 2, 2, 358, 53, 3, 2,
	2, 2, 359, 365, 5, 90, 46, 2, 360, 365, 5, 56, 29, 2, 361, 365, 5, 70,
	36, 2, 362, 365, 5, 58, 30, 2, 363, 365, 5, 72, 37, 2, 364, 359, 3, 2,
	2, 2, 364, 360, 3, 2, 2, 2, 364, 361, 3, 2, 2, 2, 364, 362, 3, 2, 2, 2,
	364, 363, 3, 2, 2, 2, 365, 55, 3, 2, 2, 2, 366, 367, 5, 90, 46, 2, 367,
	368, 9, 4, 2, 2, 368, 57, 3, 2, 2, 2, 369, 370, 7, 31, 2, 2, 370, 371,
	7, 52, 2, 2, 371, 372, 5, 80, 41, 2, 372, 373, 7, 53, 2, 2, 373, 378, 3,
	2, 2, 2, 374, 375, 7, 37, 2, 2, 375, 376, 7, 52, 2, 2, 376, 378, 7, 53,
	2, 2, 377, 369, 3, 2, 2, 2, 377, 374, 3, 2, 2, 2, 378, 59, 3, 2, 2, 2,
	379, 384, 5, 94, 48, 2, 380, 381, 7, 56, 2, 2, 381, 382, 5, 90, 46, 2,
	382, 383, 7, 57, 2, 2, 383, 385, 3, 2, 2, 2, 384, 380, 3, 2, 2, 2, 385,
	386, 3, 2, 2, 2, 386, 384, 3, 2, 2, 2, 386, 387, 3, 2, 2, 2, 387, 61, 3,
	2, 2, 2, 388, 389, 7, 4, 2, 2, 389, 391, 5, 68, 35, 2, 390, 392, 5, 66,
	34, 2, 391, 390, 3, 2, 2, 2, 391, 392, 3, 2, 2, 2, 392, 393, 3, 2, 2, 2,
	393, 394, 5, 112, 57, 2, 394, 63, 3, 2, 2, 2, 395, 396, 7, 5, 2, 2, 396,
	398, 5, 68, 35, 2, 397, 399, 5, 66, 34, 2, 398, 397, 3, 2, 2, 2, 398, 399,
	3, 2, 2, 2, 399, 400, 3, 2, 2, 2, 400, 401, 5, 112, 57, 2, 401, 65, 3,
	2, 2, 2, 402, 408, 9, 5, 2, 2, 403, 404, 9, 6, 2, 2, 404, 408, 5, 100,
	51, 2, 405, 406, 7, 21, 2, 2, 406, 408, 5, 90, 46, 2, 407, 402, 3, 2, 2,
	2, 407, 403, 3, 2, 2, 2, 407, 405, 3, 2, 2, 2, 408, 67, 3, 2, 2, 2, 409,
	410, 5, 90, 46, 2, 410, 69, 3, 2, 2, 2, 411, 413, 5, 38, 20, 2, 412, 414,
	9, 7, 2, 2, 413, 412, 3, 2, 2, 2, 413, 414, 3, 2, 2, 2, 414, 415, 3, 2,
	2, 2, 415, 416, 7, 46, 2, 2, 416, 417, 5, 38, 20, 2, 417, 423, 3, 2, 2,
	2, 418, 419, 5, 38, 20, 2, 419, 420, 9, 8, 2, 2, 420, 421, 5, 38, 20, 2,
	421, 423, 3, 2, 2, 2, 422, 411, 3, 2, 2, 2, 422, 418, 3, 2, 2, 2, 423,
	71, 3, 2, 2, 2, 424, 425, 7, 58, 2, 2, 425, 73, 3, 2, 2, 2, 426, 430, 7,
	13, 2, 2, 427, 428, 5, 54, 28, 2, 428, 429, 7, 58, 2, 2, 429, 431, 3, 2,
	2, 2, 430, 427, 3, 2, 2, 2, 430, 431, 3, 2, 2, 2, 431, 432, 3, 2, 2, 2,
	432, 433, 5, 90, 46, 2, 433, 439, 5, 48, 25, 2, 434, 437, 7, 9, 2, 2, 435,
	438, 5, 74, 38, 2, 436, 438, 5, 48, 25, 2, 437, 435, 3, 2, 2, 2, 437, 436,
	3, 2, 2, 2, 438, 440, 3, 2, 2, 2, 439, 434, 3, 2, 2, 2, 439, 440, 3, 2,
	2, 2, 440, 75, 3, 2, 2, 2, 441, 442, 7, 11, 2, 2, 442, 443, 5, 78, 40,
	2, 443, 444, 7, 18, 2, 2, 444, 445, 5, 82, 42, 2, 445, 446, 5, 112, 57,
	2, 446, 77, 3, 2, 2, 2, 447, 448, 5, 100, 51, 2, 448, 79, 3, 2, 2, 2, 449,
	450, 9, 9, 2, 2, 450, 451, 7, 51, 2, 2, 451, 456, 7, 45, 2, 2, 452, 453,
	7, 51, 2, 2, 453, 455, 7, 45, 2, 2, 454, 452, 3, 2, 2, 2, 455, 458, 3,
	2, 2, 2, 456, 454, 3, 2, 2, 2, 456, 457, 3, 2, 2, 2, 457, 81, 3, 2, 2,
	2, 458, 456, 3, 2, 2, 2, 459, 463, 7, 54, 2, 2, 460, 462, 5, 84, 43, 2,
	461, 460, 3, 2, 2, 2, 462, 465, 3, 2, 2, 2, 463, 461, 3, 2, 2, 2, 463,
	464, 3, 2, 2, 2, 464, 466, 3, 2, 2, 2, 465, 463, 3, 2, 2, 2, 466, 467,
	7, 55, 2, 2, 467, 83, 3, 2, 2, 2, 468, 473, 5, 80, 41, 2, 469, 470, 7,
	71, 2, 2, 470, 472, 5, 80, 41, 2, 471, 469, 3, 2, 2, 2, 472, 475, 3, 2,
	2, 2, 473, 471, 3, 2, 2, 2, 473, 474, 3, 2, 2, 2, 474, 476, 3, 2, 2, 2,
	475, 473, 3, 2, 2, 2, 476, 477, 5, 112, 57, 2, 477, 491, 3, 2, 2, 2, 478,
	479, 7, 45, 2, 2, 479, 480, 7, 46, 2, 2, 480, 483, 7, 16, 2, 2, 481, 484,
	5, 80, 41, 2, 482, 484, 7, 45, 2, 2, 483, 481, 3, 2, 2, 2, 483, 482, 3,
	2, 2, 2, 484, 485, 3, 2, 2, 2, 485, 491, 5, 112, 57, 2, 486, 487, 5, 54,
	28, 2, 487, 488, 5, 112, 57, 2, 488, 491, 3, 2, 2, 2, 489, 491, 5, 74,
	38, 2, 490, 468, 3, 2, 2, 2, 490, 478, 3, 2, 2, 2, 490, 486, 3, 2, 2, 2,
	490, 489, 3, 2, 2, 2, 491, 85, 3, 2, 2, 2, 492, 493, 9, 10, 2, 2, 493,
	87, 3, 2, 2, 2, 494, 495, 5, 86, 44, 2, 495, 497, 7, 52, 2, 2, 496, 498,
	5, 92, 47, 2, 497, 496, 3, 2, 2, 2, 497, 498, 3, 2, 2, 2, 498, 503, 3,
	2, 2, 2, 499, 500, 7, 50, 2, 2, 500, 502, 5, 92, 47, 2, 501, 499, 3, 2,
	2, 2, 502, 505, 3, 2, 2, 2, 503, 501, 3, 2, 2, 2, 503, 504, 3, 2, 2, 2,
	504, 506, 3, 2, 2, 2, 505, 503, 3, 2, 2, 2, 506, 507, 7, 53, 2, 2, 507,
	89, 3, 2, 2, 2, 508, 509, 8, 46, 1, 2, 509, 513, 5, 92, 47, 2, 510, 513,
	5, 88, 45, 2, 511, 513, 5, 96, 49, 2, 512, 508, 3, 2, 2, 2, 512, 510, 3,
	2, 2, 2, 512, 511, 3, 2, 2, 2, 513, 534, 3, 2, 2, 2, 514, 515, 12, 8, 2,
	2, 515, 516, 7, 75, 2, 2, 516, 533, 5, 90, 46, 9, 517, 518, 12, 7, 2, 2,
	518, 519, 9, 11, 2, 2, 519, 533, 5, 90, 46, 8, 520, 521, 12, 6, 2, 2, 521,
	522, 9, 12, 2, 2, 522, 533, 5, 90, 46, 7, 523, 524, 12, 5, 2, 2, 524, 525,
	9, 3, 2, 2, 525, 533, 5, 90, 46, 6, 526, 527, 12, 4, 2, 2, 527, 528, 7,
	62, 2, 2, 528, 533, 5, 90, 46, 5, 529, 530, 12, 3, 2, 2, 530, 531, 7, 70,
	2, 2, 531, 533, 5, 90, 46, 4, 532, 514, 3, 2, 2, 2, 532, 517, 3, 2, 2,
	2, 532, 520, 3, 2, 2, 2, 532, 523, 3, 2, 2, 2, 532, 526, 3, 2, 2, 2, 532,
	529, 3, 2, 2, 2, 533, 536, 3, 2, 2, 2, 534, 532, 3, 2, 2, 2, 534, 535,
	3, 2, 2, 2, 535, 91, 3, 2, 2, 2, 536, 534, 3, 2, 2, 2, 537, 548, 7, 28,
	2, 2, 538, 548, 5, 98, 50, 2, 539, 548, 5, 106, 54, 2, 540, 548, 5, 108,
	55, 2, 541, 548, 5, 94, 48, 2, 542, 548, 5, 60, 31, 2, 543, 544, 7, 52,
	2, 2, 544, 545, 5, 90, 46, 2, 545, 546, 7, 53, 2, 2, 546, 548, 3, 2, 2,
	2, 547, 537, 3, 2, 2, 2, 547, 538, 3, 2, 2, 2, 547, 539, 3, 2, 2, 2, 547,
	540, 3, 2, 2, 2, 547, 541, 3, 2, 2, 2, 547, 542, 3, 2, 2, 2, 547, 543,
	3, 2, 2, 2, 548, 93, 3, 2, 2, 2, 549, 560, 7, 45, 2, 2, 550, 560, 5, 80,
	41, 2, 551, 560, 7, 22, 2, 2, 552, 560, 7, 6, 2, 2, 553, 554, 7, 16, 2,
	2, 554, 557, 7, 45, 2, 2, 555, 556, 7, 51, 2, 2, 556, 558, 7, 45, 2, 2,
	557, 555, 3, 2, 2, 2, 557, 558, 3, 2, 2, 2, 558, 560, 3, 2, 2, 2, 559,
	549, 3, 2, 2, 2, 559, 550, 3, 2, 2, 2, 559, 551, 3, 2, 2, 2, 559, 552,
	3, 2, 2, 2, 559, 553, 3, 2, 2, 2, 560, 95, 3, 2, 2, 2, 561, 565, 3, 2,
	2, 2, 562, 563, 9, 13, 2, 2, 563, 565, 5, 90, 46, 2, 564, 561, 3, 2, 2,
	2, 564, 562, 3, 2, 2, 2, 565, 97, 3, 2, 2, 2, 566, 570, 5, 100, 51, 2,
	567, 570, 5, 102, 52, 2, 568, 570, 5, 104, 53, 2, 569, 566, 3, 2, 2, 2,
	569, 567, 3, 2, 2, 2, 569, 568, 3, 2, 2, 2, 570, 99, 3, 2, 2, 2, 571, 572,
	9, 14, 2, 2, 572, 101, 3, 2, 2, 2, 573, 574, 7, 73, 2, 2, 574, 578, 5,
	100, 51, 2, 575, 576, 7, 73, 2, 2, 576, 578, 5, 104, 53, 2, 577, 573, 3,
	2, 2, 2, 577, 575, 3, 2, 2, 2, 578, 103, 3, 2, 2, 2, 579, 580, 7, 85, 2,
	2, 580, 105, 3, 2, 2, 2, 581, 582, 9, 15, 2, 2, 582, 107, 3, 2, 2, 2, 583,
	584, 9, 16, 2, 2, 584, 109, 3, 2, 2, 2, 585, 586, 7, 12, 2, 2, 586, 587,
	5, 48, 25, 2, 587, 111, 3, 2, 2, 2, 588, 589, 9, 17, 2, 2, 589, 113, 3,
	2, 2, 2, 56, 118, 124, 130, 136, 141, 144, 167, 175, 189, 203, 207, 221,
	225, 230, 240, 254, 258, 263, 270, 278, 294, 305, 309, 333, 341, 348, 357,
	364, 377, 386, 391, 398, 407, 413, 422, 430, 437, 439, 456, 463, 473, 483,
	490, 497, 503, 512, 532, 534, 547, 557, 559, 564, 569, 577,
}
var literalNames = []string{
	"", "'all'", "'assert'", "'assume'", "'clock'", "'const'", "'def'", "'else'",
	"'flow'", "'for'", "'func'", "'if'", "'import'", "'init'", "'new'", "'return'",
	"'run'", "'spec'", "'stock'", "'then'", "'this'", "'eventually'", "'eventually-always'",
	"'always'", "'nmt'", "'nft'", "'nil'", "'true'", "'false'", "'advance'",
	"'component'", "'global'", "'system'", "'start'", "'state'", "'stay'",
	"'string'", "'bool'", "'int'", "'float'", "'natural'", "'uncertain'", "'unknown'",
	"", "'='", "'->'", "'<-'", "':'", "','", "'.'", "'('", "')'", "'{'", "'}'",
	"'['", "']'", "';'", "'++'", "'--'", "'&'", "'&&'", "'!'", "'=='", "'!='",
	"'<'", "'<='", "'>'", "'>='", "'||'", "'|'", "'+'", "'-'", "'^'", "'**'",
	"'*'", "'/'", "'%'", "'<<'", "'>>'", "'&^'",
}
var symbolicNames = []string{
	"", "ALL", "ASSERT", "ASSUME", "CLOCK", "CONST", "DEF", "ELSE", "FLOW",
	"FOR", "FUNC", "IF", "IMPORT", "INIT", "NEW", "RETURN", "RUN", "SPEC",
	"STOCK", "THEN", "THIS", "EVENTUALLY", "EVENTUALLYALWAYS", "ALWAYS", "NMT",
	"NFT", "NIL", "TRUE", "FALSE", "ADVANCE", "COMPONENT", "GLOBAL", "SYSTEM",
	"START", "STATE", "STAY", "TY_STRING", "TY_BOOL", "TY_INT", "TY_FLOAT",
	"TY_NATURAL", "TY_UNCERTAIN", "TY_UNKNOWN", "IDENT", "ASSIGN", "ASSIGN_FLOW1",
	"ASSIGN_FLOW2", "COLON", "COMMA", "DOT", "LPAREN", "RPAREN", "LCURLY",
	"RCURLY", "LBRACE", "RBRACE", "SEMI", "PLUS_PLUS", "MINUS_MINUS", "AMPERSAND",
	"AND", "BANG", "EQUALS", "NOT_EQUALS", "LESS", "LESS_OR_EQUALS", "GREATER",
	"GREATER_OR_EQUALS", "OR", "PIPE", "PLUS", "MINUS", "CARET", "EXPO", "MULTI",
	"DIV", "MOD", "LSHIFT", "RSHIFT", "BIT_CLEAR", "DECIMAL_LIT", "OCTAL_LIT",
	"HEX_LIT", "FLOAT_LIT", "RAW_STRING_LIT", "INTERPRETED_STRING_LIT", "WS",
	"COMMENT", "TERMINATOR", "LINE_COMMENT",
}

var ruleNames = []string{
	"sysSpec", "sysClause", "globalDecl", "componentDecl", "stateLists", "stateInit",
	"startBlock", "startPair", "spec", "specClause", "importDecl", "importSpec",
	"importPath", "declaration", "comparison", "constDecl", "constSpec", "identList",
	"expressionList", "structDecl", "structType", "structProperties", "initDecl",
	"block", "statementList", "statement", "simpleStmt", "incDecStmt", "builtins",
	"accessHistory", "assertion", "assumption", "temporal", "invariant", "assignment",
	"emptyStmt", "ifStmt", "forStmt", "rounds", "paramCall", "runBlock", "runStep",
	"faultType", "solvable", "expression", "operand", "operandName", "prefix",
	"numeric", "integer", "negative", "float_", "string_", "bool_", "functionLit",
	"eos",
}

type FaultParser struct {
	*antlr.BaseParser
}

// NewFaultParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *FaultParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewFaultParser(input antlr.TokenStream) *FaultParser {
	this := new(FaultParser)
	deserializer := antlr.NewATNDeserializer(nil)
	deserializedATN := deserializer.DeserializeFromUInt16(parserATN)
	decisionToDFA := make([]*antlr.DFA, len(deserializedATN.DecisionToState))
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
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
	FaultParserTHEN                   = 19
	FaultParserTHIS                   = 20
	FaultParserEVENTUALLY             = 21
	FaultParserEVENTUALLYALWAYS       = 22
	FaultParserALWAYS                 = 23
	FaultParserNMT                    = 24
	FaultParserNFT                    = 25
	FaultParserNIL                    = 26
	FaultParserTRUE                   = 27
	FaultParserFALSE                  = 28
	FaultParserADVANCE                = 29
	FaultParserCOMPONENT              = 30
	FaultParserGLOBAL                 = 31
	FaultParserSYSTEM                 = 32
	FaultParserSTART                  = 33
	FaultParserSTATE                  = 34
	FaultParserSTAY                   = 35
	FaultParserTY_STRING              = 36
	FaultParserTY_BOOL                = 37
	FaultParserTY_INT                 = 38
	FaultParserTY_FLOAT               = 39
	FaultParserTY_NATURAL             = 40
	FaultParserTY_UNCERTAIN           = 41
	FaultParserTY_UNKNOWN             = 42
	FaultParserIDENT                  = 43
	FaultParserASSIGN                 = 44
	FaultParserASSIGN_FLOW1           = 45
	FaultParserASSIGN_FLOW2           = 46
	FaultParserCOLON                  = 47
	FaultParserCOMMA                  = 48
	FaultParserDOT                    = 49
	FaultParserLPAREN                 = 50
	FaultParserRPAREN                 = 51
	FaultParserLCURLY                 = 52
	FaultParserRCURLY                 = 53
	FaultParserLBRACE                 = 54
	FaultParserRBRACE                 = 55
	FaultParserSEMI                   = 56
	FaultParserPLUS_PLUS              = 57
	FaultParserMINUS_MINUS            = 58
	FaultParserAMPERSAND              = 59
	FaultParserAND                    = 60
	FaultParserBANG                   = 61
	FaultParserEQUALS                 = 62
	FaultParserNOT_EQUALS             = 63
	FaultParserLESS                   = 64
	FaultParserLESS_OR_EQUALS         = 65
	FaultParserGREATER                = 66
	FaultParserGREATER_OR_EQUALS      = 67
	FaultParserOR                     = 68
	FaultParserPIPE                   = 69
	FaultParserPLUS                   = 70
	FaultParserMINUS                  = 71
	FaultParserCARET                  = 72
	FaultParserEXPO                   = 73
	FaultParserMULTI                  = 74
	FaultParserDIV                    = 75
	FaultParserMOD                    = 76
	FaultParserLSHIFT                 = 77
	FaultParserRSHIFT                 = 78
	FaultParserBIT_CLEAR              = 79
	FaultParserDECIMAL_LIT            = 80
	FaultParserOCTAL_LIT              = 81
	FaultParserHEX_LIT                = 82
	FaultParserFLOAT_LIT              = 83
	FaultParserRAW_STRING_LIT         = 84
	FaultParserINTERPRETED_STRING_LIT = 85
	FaultParserWS                     = 86
	FaultParserCOMMENT                = 87
	FaultParserTERMINATOR             = 88
	FaultParserLINE_COMMENT           = 89
)

// FaultParser rules.
const (
	FaultParserRULE_sysSpec          = 0
	FaultParserRULE_sysClause        = 1
	FaultParserRULE_globalDecl       = 2
	FaultParserRULE_componentDecl    = 3
	FaultParserRULE_stateLists       = 4
	FaultParserRULE_stateInit        = 5
	FaultParserRULE_startBlock       = 6
	FaultParserRULE_startPair        = 7
	FaultParserRULE_spec             = 8
	FaultParserRULE_specClause       = 9
	FaultParserRULE_importDecl       = 10
	FaultParserRULE_importSpec       = 11
	FaultParserRULE_importPath       = 12
	FaultParserRULE_declaration      = 13
	FaultParserRULE_comparison       = 14
	FaultParserRULE_constDecl        = 15
	FaultParserRULE_constSpec        = 16
	FaultParserRULE_identList        = 17
	FaultParserRULE_expressionList   = 18
	FaultParserRULE_structDecl       = 19
	FaultParserRULE_structType       = 20
	FaultParserRULE_structProperties = 21
	FaultParserRULE_initDecl         = 22
	FaultParserRULE_block            = 23
	FaultParserRULE_statementList    = 24
	FaultParserRULE_statement        = 25
	FaultParserRULE_simpleStmt       = 26
	FaultParserRULE_incDecStmt       = 27
	FaultParserRULE_builtins         = 28
	FaultParserRULE_accessHistory    = 29
	FaultParserRULE_assertion        = 30
	FaultParserRULE_assumption       = 31
	FaultParserRULE_temporal         = 32
	FaultParserRULE_invariant        = 33
	FaultParserRULE_assignment       = 34
	FaultParserRULE_emptyStmt        = 35
	FaultParserRULE_ifStmt           = 36
	FaultParserRULE_forStmt          = 37
	FaultParserRULE_rounds           = 38
	FaultParserRULE_paramCall        = 39
	FaultParserRULE_runBlock         = 40
	FaultParserRULE_runStep          = 41
	FaultParserRULE_faultType        = 42
	FaultParserRULE_solvable         = 43
	FaultParserRULE_expression       = 44
	FaultParserRULE_operand          = 45
	FaultParserRULE_operandName      = 46
	FaultParserRULE_prefix           = 47
	FaultParserRULE_numeric          = 48
	FaultParserRULE_integer          = 49
	FaultParserRULE_negative         = 50
	FaultParserRULE_float_           = 51
	FaultParserRULE_string_          = 52
	FaultParserRULE_bool_            = 53
	FaultParserRULE_functionLit      = 54
	FaultParserRULE_eos              = 55
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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISysClauseContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISysClauseContext)
}

func (s *SysSpecContext) Eos() IEosContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEosContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEosContext)
}

func (s *SysSpecContext) AllImportDecl() []IImportDeclContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IImportDeclContext)(nil)).Elem())
	var tst = make([]IImportDeclContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IImportDeclContext)
		}
	}

	return tst
}

func (s *SysSpecContext) ImportDecl(i int) IImportDeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IImportDeclContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IImportDeclContext)
}

func (s *SysSpecContext) AllGlobalDecl() []IGlobalDeclContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IGlobalDeclContext)(nil)).Elem())
	var tst = make([]IGlobalDeclContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IGlobalDeclContext)
		}
	}

	return tst
}

func (s *SysSpecContext) GlobalDecl(i int) IGlobalDeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IGlobalDeclContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IGlobalDeclContext)
}

func (s *SysSpecContext) AllComponentDecl() []IComponentDeclContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IComponentDeclContext)(nil)).Elem())
	var tst = make([]IComponentDeclContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IComponentDeclContext)
		}
	}

	return tst
}

func (s *SysSpecContext) ComponentDecl(i int) IComponentDeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IComponentDeclContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IComponentDeclContext)
}

func (s *SysSpecContext) AllStartBlock() []IStartBlockContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStartBlockContext)(nil)).Elem())
	var tst = make([]IStartBlockContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStartBlockContext)
		}
	}

	return tst
}

func (s *SysSpecContext) StartBlock(i int) IStartBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStartBlockContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStartBlockContext)
}

func (s *SysSpecContext) Assertion() IAssertionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAssertionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAssertionContext)
}

func (s *SysSpecContext) Assumption() IAssumptionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAssumptionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAssumptionContext)
}

func (s *SysSpecContext) ForStmt() IForStmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IForStmtContext)(nil)).Elem(), 0)

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
		p.SetState(112)
		p.SysClause()
	}
	p.SetState(116)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == FaultParserIMPORT {
		{
			p.SetState(113)
			p.ImportDecl()
		}

		p.SetState(118)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(122)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == FaultParserGLOBAL {
		{
			p.SetState(119)
			p.GlobalDecl()
		}

		p.SetState(124)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(128)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == FaultParserDEF {
		{
			p.SetState(125)
			p.ComponentDecl()
		}

		p.SetState(130)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(134)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == FaultParserSTART {
		{
			p.SetState(131)
			p.StartBlock()
		}

		p.SetState(136)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(139)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case FaultParserASSERT:
		{
			p.SetState(137)
			p.Assertion()
		}

	case FaultParserASSUME:
		{
			p.SetState(138)
			p.Assumption()
		}

	case FaultParserEOF, FaultParserFOR, FaultParserSEMI:

	default:
	}
	p.SetState(142)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == FaultParserFOR {
		{
			p.SetState(141)
			p.ForStmt()
		}

	}
	{
		p.SetState(144)
		p.Eos()
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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEosContext)(nil)).Elem(), 0)

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
		p.SetState(146)
		p.Match(FaultParserSYSTEM)
	}
	{
		p.SetState(147)
		p.Match(FaultParserIDENT)
	}
	{
		p.SetState(148)
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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOperandContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOperandContext)
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
		p.SetState(150)
		p.Match(FaultParserGLOBAL)
	}
	{
		p.SetState(151)
		p.Match(FaultParserIDENT)
	}
	{
		p.SetState(152)
		p.Match(FaultParserASSIGN)
	}
	{
		p.SetState(153)
		p.Operand()
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

func (s *ComponentDeclContext) DEF() antlr.TerminalNode {
	return s.GetToken(FaultParserDEF, 0)
}

func (s *ComponentDeclContext) IDENT() antlr.TerminalNode {
	return s.GetToken(FaultParserIDENT, 0)
}

func (s *ComponentDeclContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(FaultParserASSIGN, 0)
}

func (s *ComponentDeclContext) COMPONENT() antlr.TerminalNode {
	return s.GetToken(FaultParserCOMPONENT, 0)
}

func (s *ComponentDeclContext) LCURLY() antlr.TerminalNode {
	return s.GetToken(FaultParserLCURLY, 0)
}

func (s *ComponentDeclContext) RCURLY() antlr.TerminalNode {
	return s.GetToken(FaultParserRCURLY, 0)
}

func (s *ComponentDeclContext) Eos() IEosContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEosContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEosContext)
}

func (s *ComponentDeclContext) AllStateLists() []IStateListsContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStateListsContext)(nil)).Elem())
	var tst = make([]IStateListsContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStateListsContext)
		}
	}

	return tst
}

func (s *ComponentDeclContext) StateLists(i int) IStateListsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStateListsContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStateListsContext)
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
		p.SetState(155)
		p.Match(FaultParserDEF)
	}
	{
		p.SetState(156)
		p.Match(FaultParserIDENT)
	}
	{
		p.SetState(157)
		p.Match(FaultParserASSIGN)
	}
	{
		p.SetState(158)
		p.Match(FaultParserCOMPONENT)
	}
	{
		p.SetState(159)
		p.Match(FaultParserLCURLY)
	}
	p.SetState(165)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == FaultParserIDENT {
		{
			p.SetState(160)
			p.StateLists()
		}
		{
			p.SetState(161)
			p.Match(FaultParserCOMMA)
		}

		p.SetState(167)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(168)
		p.Match(FaultParserRCURLY)
	}
	{
		p.SetState(169)
		p.Eos()
	}

	return localctx
}

// IStateListsContext is an interface to support dynamic dispatch.
type IStateListsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStateListsContext differentiates from other interfaces.
	IsStateListsContext()
}

type StateListsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStateListsContext() *StateListsContext {
	var p = new(StateListsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FaultParserRULE_stateLists
	return p
}

func (*StateListsContext) IsStateListsContext() {}

func NewStateListsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StateListsContext {
	var p = new(StateListsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FaultParserRULE_stateLists

	return p
}

func (s *StateListsContext) GetParser() antlr.Parser { return s.parser }

func (s *StateListsContext) StructProperties() IStructPropertiesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStructPropertiesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStructPropertiesContext)
}

func (s *StateListsContext) StateInit() IStateInitContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStateInitContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStateInitContext)
}

func (s *StateListsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StateListsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StateListsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaultParserListener); ok {
		listenerT.EnterStateLists(s)
	}
}

func (s *StateListsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaultParserListener); ok {
		listenerT.ExitStateLists(s)
	}
}

func (s *StateListsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FaultParserVisitor:
		return t.VisitStateLists(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FaultParser) StateLists() (localctx IStateListsContext) {
	this := p
	_ = this

	localctx = NewStateListsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, FaultParserRULE_stateLists)

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

	p.SetState(173)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(171)
			p.StructProperties()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(172)
			p.StateInit()
		}

	}

	return localctx
}

// IStateInitContext is an interface to support dynamic dispatch.
type IStateInitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStateInitContext differentiates from other interfaces.
	IsStateInitContext()
}

type StateInitContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStateInitContext() *StateInitContext {
	var p = new(StateInitContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FaultParserRULE_stateInit
	return p
}

func (*StateInitContext) IsStateInitContext() {}

func NewStateInitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StateInitContext {
	var p = new(StateInitContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FaultParserRULE_stateInit

	return p
}

func (s *StateInitContext) GetParser() antlr.Parser { return s.parser }

func (s *StateInitContext) IDENT() antlr.TerminalNode {
	return s.GetToken(FaultParserIDENT, 0)
}

func (s *StateInitContext) COLON() antlr.TerminalNode {
	return s.GetToken(FaultParserCOLON, 0)
}

func (s *StateInitContext) STATE() antlr.TerminalNode {
	return s.GetToken(FaultParserSTATE, 0)
}

func (s *StateInitContext) Block() IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *StateInitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StateInitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StateInitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaultParserListener); ok {
		listenerT.EnterStateInit(s)
	}
}

func (s *StateInitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaultParserListener); ok {
		listenerT.ExitStateInit(s)
	}
}

func (s *StateInitContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case FaultParserVisitor:
		return t.VisitStateInit(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *FaultParser) StateInit() (localctx IStateInitContext) {
	this := p
	_ = this

	localctx = NewStateInitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, FaultParserRULE_stateInit)

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
		p.SetState(175)
		p.Match(FaultParserIDENT)
	}
	{
		p.SetState(176)
		p.Match(FaultParserCOLON)
	}
	{
		p.SetState(177)
		p.Match(FaultParserSTATE)
	}
	{
		p.SetState(178)
		p.Block()
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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEosContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEosContext)
}

func (s *StartBlockContext) AllStartPair() []IStartPairContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStartPairContext)(nil)).Elem())
	var tst = make([]IStartPairContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStartPairContext)
		}
	}

	return tst
}

func (s *StartBlockContext) StartPair(i int) IStartPairContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStartPairContext)(nil)).Elem(), i)

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
	p.EnterRule(localctx, 12, FaultParserRULE_startBlock)
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
		p.SetState(180)
		p.Match(FaultParserSTART)
	}
	{
		p.SetState(181)
		p.Match(FaultParserLCURLY)
	}
	p.SetState(187)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == FaultParserIDENT {
		{
			p.SetState(182)
			p.StartPair()
		}
		{
			p.SetState(183)
			p.Match(FaultParserCOMMA)
		}

		p.SetState(189)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(190)
		p.Match(FaultParserRCURLY)
	}
	{
		p.SetState(191)
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
	p.EnterRule(localctx, 14, FaultParserRULE_startPair)

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
		p.SetState(193)
		p.Match(FaultParserIDENT)
	}
	{
		p.SetState(194)
		p.Match(FaultParserCOLON)
	}
	{
		p.SetState(195)
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
	this := p
	_ = this

	localctx = NewSpecContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, FaultParserRULE_spec)
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
		p.SetState(197)
		p.SpecClause()
	}
	p.SetState(201)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<FaultParserASSERT)|(1<<FaultParserASSUME)|(1<<FaultParserCONST)|(1<<FaultParserDEF))) != 0 {
		{
			p.SetState(198)
			p.Declaration()
		}

		p.SetState(203)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(205)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == FaultParserFOR {
		{
			p.SetState(204)
			p.ForStmt()
		}

	}
	{
		p.SetState(207)
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
	this := p
	_ = this

	localctx = NewSpecClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, FaultParserRULE_specClause)

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
		p.SetState(209)
		p.Match(FaultParserSPEC)
	}
	{
		p.SetState(210)
		p.Match(FaultParserIDENT)
	}
	{
		p.SetState(211)
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
	this := p
	_ = this

	localctx = NewImportDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, FaultParserRULE_importDecl)
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
		p.SetState(213)
		p.Match(FaultParserIMPORT)
	}
	p.SetState(223)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case FaultParserIDENT, FaultParserDOT, FaultParserRAW_STRING_LIT, FaultParserINTERPRETED_STRING_LIT:
		{
			p.SetState(214)
			p.ImportSpec()
		}

	case FaultParserLPAREN:
		{
			p.SetState(215)
			p.Match(FaultParserLPAREN)
		}
		p.SetState(219)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == FaultParserIDENT || _la == FaultParserDOT || _la == FaultParserRAW_STRING_LIT || _la == FaultParserINTERPRETED_STRING_LIT {
			{
				p.SetState(216)
				p.ImportSpec()
			}

			p.SetState(221)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(222)
			p.Match(FaultParserRPAREN)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	{
		p.SetState(225)
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
	this := p
	_ = this

	localctx = NewImportSpecContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, FaultParserRULE_importSpec)
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
	p.SetState(228)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == FaultParserIDENT || _la == FaultParserDOT {
		{
			p.SetState(227)
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
		p.SetState(230)
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
	this := p
	_ = this

	localctx = NewImportPathContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, FaultParserRULE_importPath)

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
		p.SetState(232)
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
	this := p
	_ = this

	localctx = NewDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, FaultParserRULE_declaration)

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

	p.SetState(238)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case FaultParserCONST:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(234)
			p.ConstDecl()
		}

	case FaultParserDEF:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(235)
			p.StructDecl()
		}

	case FaultParserASSERT:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(236)
			p.Assertion()
		}

	case FaultParserASSUME:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(237)
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
	p.EnterRule(localctx, 28, FaultParserRULE_comparison)
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
		p.SetState(240)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-62)&-(0x1f+1)) == 0 && ((1<<uint((_la-62)))&((1<<(FaultParserEQUALS-62))|(1<<(FaultParserNOT_EQUALS-62))|(1<<(FaultParserLESS-62))|(1<<(FaultParserLESS_OR_EQUALS-62))|(1<<(FaultParserGREATER-62))|(1<<(FaultParserGREATER_OR_EQUALS-62)))) != 0) {
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
	this := p
	_ = this

	localctx = NewConstDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, FaultParserRULE_constDecl)
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
		p.SetState(242)
		p.Match(FaultParserCONST)
	}
	p.SetState(256)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case FaultParserCLOCK, FaultParserNEW, FaultParserTHIS, FaultParserIDENT:
		{
			p.SetState(243)
			p.ConstSpec()
		}
		{
			p.SetState(244)
			p.Eos()
		}

	case FaultParserLPAREN:
		{
			p.SetState(246)
			p.Match(FaultParserLPAREN)
		}
		p.SetState(252)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<FaultParserCLOCK)|(1<<FaultParserNEW)|(1<<FaultParserTHIS))) != 0) || _la == FaultParserIDENT {
			{
				p.SetState(247)
				p.ConstSpec()
			}
			{
				p.SetState(248)
				p.Eos()
			}

			p.SetState(254)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(255)
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
	this := p
	_ = this

	localctx = NewConstSpecContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, FaultParserRULE_constSpec)
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
		p.SetState(258)
		p.IdentList()
	}
	p.SetState(261)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == FaultParserASSIGN {
		{
			p.SetState(259)
			p.Match(FaultParserASSIGN)
		}
		{
			p.SetState(260)
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
	this := p
	_ = this

	localctx = NewIdentListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, FaultParserRULE_identList)
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
		p.OperandName()
	}
	p.SetState(268)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == FaultParserCOMMA {
		{
			p.SetState(264)
			p.Match(FaultParserCOMMA)
		}
		{
			p.SetState(265)
			p.OperandName()
		}

		p.SetState(270)
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
		p.SetState(271)
		p.expression(0)
	}
	p.SetState(276)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == FaultParserCOMMA {
		{
			p.SetState(272)
			p.Match(FaultParserCOMMA)
		}
		{
			p.SetState(273)
			p.expression(0)
		}

		p.SetState(278)
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
		p.SetState(279)
		p.Match(FaultParserDEF)
	}
	{
		p.SetState(280)
		p.Match(FaultParserIDENT)
	}
	{
		p.SetState(281)
		p.Match(FaultParserASSIGN)
	}
	{
		p.SetState(282)
		p.StructType()
	}
	{
		p.SetState(283)
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

	p.SetState(307)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case FaultParserFLOW:
		localctx = NewFlowContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(285)
			p.Match(FaultParserFLOW)
		}
		{
			p.SetState(286)
			p.Match(FaultParserLCURLY)
		}
		p.SetState(292)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == FaultParserIDENT {
			{
				p.SetState(287)
				p.StructProperties()
			}
			{
				p.SetState(288)
				p.Match(FaultParserCOMMA)
			}

			p.SetState(294)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(295)
			p.Match(FaultParserRCURLY)
		}

	case FaultParserSTOCK:
		localctx = NewStockContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(296)
			p.Match(FaultParserSTOCK)
		}
		{
			p.SetState(297)
			p.Match(FaultParserLCURLY)
		}
		p.SetState(303)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == FaultParserIDENT {
			{
				p.SetState(298)
				p.StructProperties()
			}
			{
				p.SetState(299)
				p.Match(FaultParserCOMMA)
			}

			p.SetState(305)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(306)
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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBool_Context)(nil)).Elem(), 0)

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
	this := p
	_ = this

	localctx = NewStructPropertiesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, FaultParserRULE_structProperties)

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
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 23, p.GetParserRuleContext()) {
	case 1:
		localctx = NewPropIntContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(309)
			p.Match(FaultParserIDENT)
		}
		{
			p.SetState(310)
			p.Match(FaultParserCOLON)
		}
		{
			p.SetState(311)
			p.Numeric()
		}

	case 2:
		localctx = NewPropStringContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(312)
			p.Match(FaultParserIDENT)
		}
		{
			p.SetState(313)
			p.Match(FaultParserCOLON)
		}
		{
			p.SetState(314)
			p.String_()
		}

	case 3:
		localctx = NewPropBoolContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(315)
			p.Match(FaultParserIDENT)
		}
		{
			p.SetState(316)
			p.Match(FaultParserCOLON)
		}
		{
			p.SetState(317)
			p.Bool_()
		}

	case 4:
		localctx = NewPropFuncContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(318)
			p.Match(FaultParserIDENT)
		}
		{
			p.SetState(319)
			p.Match(FaultParserCOLON)
		}
		{
			p.SetState(320)
			p.FunctionLit()
		}

	case 5:
		localctx = NewPropVarContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
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
			p.OperandName()
		}

	case 6:
		localctx = NewPropVarContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(324)
			p.Match(FaultParserIDENT)
		}
		{
			p.SetState(325)
			p.Match(FaultParserCOLON)
		}
		{
			p.SetState(326)
			p.Prefix()
		}

	case 7:
		localctx = NewPropSolvableContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
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
			p.Solvable()
		}

	case 8:
		localctx = NewPropSolvableContext(p, localctx)
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(330)
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
	this := p
	_ = this

	localctx = NewInitDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, FaultParserRULE_initDecl)

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
		p.SetState(333)
		p.Match(FaultParserINIT)
	}
	{
		p.SetState(334)
		p.Operand()
	}
	{
		p.SetState(335)
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
	this := p
	_ = this

	localctx = NewBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, FaultParserRULE_block)

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
		p.SetState(337)
		p.Match(FaultParserLCURLY)
	}
	p.SetState(339)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 24, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(338)
			p.StatementList()
		}

	}
	{
		p.SetState(341)
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
	this := p
	_ = this

	localctx = NewStatementListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, FaultParserRULE_statementList)

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
	p.SetState(344)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(343)
				p.Statement()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(346)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 25, p.GetParserRuleContext())
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
	this := p
	_ = this

	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, FaultParserRULE_statement)

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

	p.SetState(355)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 26, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(348)
			p.ConstDecl()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(349)
			p.InitDecl()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(350)
			p.SimpleStmt()
		}
		{
			p.SetState(351)
			p.Eos()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(353)
			p.Block()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(354)
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

func (s *SimpleStmtContext) Builtins() IBuiltinsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBuiltinsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBuiltinsContext)
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
	this := p
	_ = this

	localctx = NewSimpleStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, FaultParserRULE_simpleStmt)

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

	p.SetState(362)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 27, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(357)
			p.expression(0)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(358)
			p.IncDecStmt()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(359)
			p.Assignment()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(360)
			p.Builtins()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(361)
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
	this := p
	_ = this

	localctx = NewIncDecStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, FaultParserRULE_incDecStmt)
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
		p.SetState(364)
		p.expression(0)
	}
	{
		p.SetState(365)
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

// IBuiltinsContext is an interface to support dynamic dispatch.
type IBuiltinsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBuiltinsContext differentiates from other interfaces.
	IsBuiltinsContext()
}

type BuiltinsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBuiltinsContext() *BuiltinsContext {
	var p = new(BuiltinsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FaultParserRULE_builtins
	return p
}

func (*BuiltinsContext) IsBuiltinsContext() {}

func NewBuiltinsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BuiltinsContext {
	var p = new(BuiltinsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FaultParserRULE_builtins

	return p
}

func (s *BuiltinsContext) GetParser() antlr.Parser { return s.parser }

func (s *BuiltinsContext) ADVANCE() antlr.TerminalNode {
	return s.GetToken(FaultParserADVANCE, 0)
}

func (s *BuiltinsContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(FaultParserLPAREN, 0)
}

func (s *BuiltinsContext) ParamCall() IParamCallContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParamCallContext)(nil)).Elem(), 0)

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

func (s *BuiltinsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BuiltinsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
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

func (p *FaultParser) Builtins() (localctx IBuiltinsContext) {
	this := p
	_ = this

	localctx = NewBuiltinsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, FaultParserRULE_builtins)

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

	p.SetState(375)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case FaultParserADVANCE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(367)
			p.Match(FaultParserADVANCE)
		}
		{
			p.SetState(368)
			p.Match(FaultParserLPAREN)
		}
		{
			p.SetState(369)
			p.ParamCall()
		}
		{
			p.SetState(370)
			p.Match(FaultParserRPAREN)
		}

	case FaultParserSTAY:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(372)
			p.Match(FaultParserSTAY)
		}
		{
			p.SetState(373)
			p.Match(FaultParserLPAREN)
		}
		{
			p.SetState(374)
			p.Match(FaultParserRPAREN)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
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
	this := p
	_ = this

	localctx = NewAccessHistoryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, FaultParserRULE_accessHistory)

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
		p.SetState(377)
		p.OperandName()
	}
	p.SetState(382)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(378)
				p.Match(FaultParserLBRACE)
			}
			{
				p.SetState(379)
				p.expression(0)
			}
			{
				p.SetState(380)
				p.Match(FaultParserRBRACE)
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(384)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 29, p.GetParserRuleContext())
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
	this := p
	_ = this

	localctx = NewAssertionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, FaultParserRULE_assertion)
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
		p.SetState(386)
		p.Match(FaultParserASSERT)
	}
	{
		p.SetState(387)
		p.Invariant()
	}
	p.SetState(389)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<FaultParserTHEN)|(1<<FaultParserEVENTUALLY)|(1<<FaultParserEVENTUALLYALWAYS)|(1<<FaultParserALWAYS)|(1<<FaultParserNMT)|(1<<FaultParserNFT))) != 0 {
		{
			p.SetState(388)
			p.Temporal()
		}

	}
	{
		p.SetState(391)
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
	this := p
	_ = this

	localctx = NewAssumptionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, FaultParserRULE_assumption)
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
		p.SetState(393)
		p.Match(FaultParserASSUME)
	}
	{
		p.SetState(394)
		p.Invariant()
	}
	p.SetState(396)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<FaultParserTHEN)|(1<<FaultParserEVENTUALLY)|(1<<FaultParserEVENTUALLYALWAYS)|(1<<FaultParserALWAYS)|(1<<FaultParserNMT)|(1<<FaultParserNFT))) != 0 {
		{
			p.SetState(395)
			p.Temporal()
		}

	}
	{
		p.SetState(398)
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
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIntegerContext)(nil)).Elem(), 0)

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

func (s *TemporalContext) THEN() antlr.TerminalNode {
	return s.GetToken(FaultParserTHEN, 0)
}

func (s *TemporalContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
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
	p.EnterRule(localctx, 64, FaultParserRULE_temporal)
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

	p.SetState(405)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case FaultParserEVENTUALLY, FaultParserEVENTUALLYALWAYS, FaultParserALWAYS:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(400)
			_la = p.GetTokenStream().LA(1)

			if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<FaultParserEVENTUALLY)|(1<<FaultParserEVENTUALLYALWAYS)|(1<<FaultParserALWAYS))) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	case FaultParserNMT, FaultParserNFT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(401)
			_la = p.GetTokenStream().LA(1)

			if !(_la == FaultParserNMT || _la == FaultParserNFT) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(402)
			p.Integer()
		}

	case FaultParserTHEN:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(403)
			p.Match(FaultParserTHEN)
		}
		{
			p.SetState(404)
			p.expression(0)
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

func (s *InvariantContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
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
	this := p
	_ = this

	localctx = NewInvariantContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, FaultParserRULE_invariant)

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
		p.SetState(407)
		p.expression(0)
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
	this := p
	_ = this

	localctx = NewAssignmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, FaultParserRULE_assignment)
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

	p.SetState(420)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 34, p.GetParserRuleContext()) {
	case 1:
		localctx = NewMiscAssignContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(409)
			p.ExpressionList()
		}
		p.SetState(411)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if ((_la-59)&-(0x1f+1)) == 0 && ((1<<uint((_la-59)))&((1<<(FaultParserAMPERSAND-59))|(1<<(FaultParserPLUS-59))|(1<<(FaultParserMINUS-59))|(1<<(FaultParserCARET-59))|(1<<(FaultParserMULTI-59))|(1<<(FaultParserDIV-59))|(1<<(FaultParserMOD-59))|(1<<(FaultParserLSHIFT-59))|(1<<(FaultParserRSHIFT-59))|(1<<(FaultParserBIT_CLEAR-59)))) != 0 {
			{
				p.SetState(410)
				_la = p.GetTokenStream().LA(1)

				if !(((_la-59)&-(0x1f+1)) == 0 && ((1<<uint((_la-59)))&((1<<(FaultParserAMPERSAND-59))|(1<<(FaultParserPLUS-59))|(1<<(FaultParserMINUS-59))|(1<<(FaultParserCARET-59))|(1<<(FaultParserMULTI-59))|(1<<(FaultParserDIV-59))|(1<<(FaultParserMOD-59))|(1<<(FaultParserLSHIFT-59))|(1<<(FaultParserRSHIFT-59))|(1<<(FaultParserBIT_CLEAR-59)))) != 0) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}

		}
		{
			p.SetState(413)
			p.Match(FaultParserASSIGN)
		}
		{
			p.SetState(414)
			p.ExpressionList()
		}

	case 2:
		localctx = NewFaultAssignContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(416)
			p.ExpressionList()
		}
		{
			p.SetState(417)
			_la = p.GetTokenStream().LA(1)

			if !(_la == FaultParserASSIGN_FLOW1 || _la == FaultParserASSIGN_FLOW2) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(418)
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
	p.EnterRule(localctx, 70, FaultParserRULE_emptyStmt)

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
		p.SetState(422)
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
	this := p
	_ = this

	localctx = NewIfStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, FaultParserRULE_ifStmt)

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
		p.SetState(424)
		p.Match(FaultParserIF)
	}
	p.SetState(428)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 35, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(425)
			p.SimpleStmt()
		}
		{
			p.SetState(426)
			p.Match(FaultParserSEMI)
		}

	}
	{
		p.SetState(430)
		p.expression(0)
	}
	{
		p.SetState(431)
		p.Block()
	}
	p.SetState(437)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 37, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(432)
			p.Match(FaultParserELSE)
		}
		p.SetState(435)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case FaultParserIF:
			{
				p.SetState(433)
				p.IfStmt()
			}

		case FaultParserLCURLY:
			{
				p.SetState(434)
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
	this := p
	_ = this

	localctx = NewForStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, FaultParserRULE_forStmt)

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
		p.SetState(439)
		p.Match(FaultParserFOR)
	}
	{
		p.SetState(440)
		p.Rounds()
	}
	{
		p.SetState(441)
		p.Match(FaultParserRUN)
	}
	{
		p.SetState(442)
		p.RunBlock()
	}
	{
		p.SetState(443)
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
	this := p
	_ = this

	localctx = NewRoundsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, FaultParserRULE_rounds)

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
		p.SetState(445)
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
	p.EnterRule(localctx, 78, FaultParserRULE_paramCall)
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
		p.SetState(447)
		_la = p.GetTokenStream().LA(1)

		if !(_la == FaultParserTHIS || _la == FaultParserIDENT) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(448)
		p.Match(FaultParserDOT)
	}
	{
		p.SetState(449)
		p.Match(FaultParserIDENT)
	}
	p.SetState(454)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 38, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(450)
				p.Match(FaultParserDOT)
			}
			{
				p.SetState(451)
				p.Match(FaultParserIDENT)
			}

		}
		p.SetState(456)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 38, p.GetParserRuleContext())
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
	this := p
	_ = this

	localctx = NewRunBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 80, FaultParserRULE_runBlock)

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
		p.SetState(457)
		p.Match(FaultParserLCURLY)
	}
	p.SetState(461)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 39, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(458)
				p.RunStep()
			}

		}
		p.SetState(463)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 39, p.GetParserRuleContext())
	}
	{
		p.SetState(464)
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

func (s *RunInitContext) ParamCall() IParamCallContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParamCallContext)(nil)).Elem(), 0)

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
	this := p
	_ = this

	localctx = NewRunStepContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 82, FaultParserRULE_runStep)
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

	p.SetState(488)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 42, p.GetParserRuleContext()) {
	case 1:
		localctx = NewRunStepExprContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(466)
			p.ParamCall()
		}
		p.SetState(471)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == FaultParserPIPE {
			{
				p.SetState(467)
				p.Match(FaultParserPIPE)
			}
			{
				p.SetState(468)
				p.ParamCall()
			}

			p.SetState(473)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(474)
			p.Eos()
		}

	case 2:
		localctx = NewRunInitContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(476)
			p.Match(FaultParserIDENT)
		}
		{
			p.SetState(477)
			p.Match(FaultParserASSIGN)
		}
		{
			p.SetState(478)
			p.Match(FaultParserNEW)
		}
		p.SetState(481)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 41, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(479)
				p.ParamCall()
			}

		case 2:
			{
				p.SetState(480)
				p.Match(FaultParserIDENT)
			}

		}
		{
			p.SetState(483)
			p.Eos()
		}

	case 3:
		localctx = NewRunExprContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(484)
			p.SimpleStmt()
		}
		{
			p.SetState(485)
			p.Eos()
		}

	case 4:
		localctx = NewRunExprContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(487)
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
	this := p
	_ = this

	localctx = NewFaultTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 84, FaultParserRULE_faultType)
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
		_la = p.GetTokenStream().LA(1)

		if !(((_la-36)&-(0x1f+1)) == 0 && ((1<<uint((_la-36)))&((1<<(FaultParserTY_STRING-36))|(1<<(FaultParserTY_BOOL-36))|(1<<(FaultParserTY_INT-36))|(1<<(FaultParserTY_FLOAT-36))|(1<<(FaultParserTY_NATURAL-36))|(1<<(FaultParserTY_UNCERTAIN-36))|(1<<(FaultParserTY_UNKNOWN-36)))) != 0) {
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
	this := p
	_ = this

	localctx = NewSolvableContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 86, FaultParserRULE_solvable)
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
		p.SetState(492)
		p.FaultType()
	}
	{
		p.SetState(493)
		p.Match(FaultParserLPAREN)
	}
	p.SetState(495)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<FaultParserCLOCK)|(1<<FaultParserNEW)|(1<<FaultParserTHIS)|(1<<FaultParserNIL)|(1<<FaultParserTRUE)|(1<<FaultParserFALSE))) != 0) || (((_la-43)&-(0x1f+1)) == 0 && ((1<<uint((_la-43)))&((1<<(FaultParserIDENT-43))|(1<<(FaultParserLPAREN-43))|(1<<(FaultParserMINUS-43)))) != 0) || (((_la-80)&-(0x1f+1)) == 0 && ((1<<uint((_la-80)))&((1<<(FaultParserDECIMAL_LIT-80))|(1<<(FaultParserOCTAL_LIT-80))|(1<<(FaultParserHEX_LIT-80))|(1<<(FaultParserFLOAT_LIT-80))|(1<<(FaultParserRAW_STRING_LIT-80))|(1<<(FaultParserINTERPRETED_STRING_LIT-80)))) != 0) {
		{
			p.SetState(494)
			p.Operand()
		}

	}
	p.SetState(501)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == FaultParserCOMMA {
		{
			p.SetState(497)
			p.Match(FaultParserCOMMA)
		}
		{
			p.SetState(498)
			p.Operand()
		}

		p.SetState(503)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(504)
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
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 88
	p.EnterRecursionRule(localctx, 88, FaultParserRULE_expression, _p)
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
	p.SetState(510)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 45, p.GetParserRuleContext()) {
	case 1:
		localctx = NewExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(507)
			p.Operand()
		}

	case 2:
		localctx = NewTypedContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(508)
			p.Solvable()
		}

	case 3:
		localctx = NewExprPrefixContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(509)
			p.Prefix()
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(532)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 47, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(530)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 46, p.GetParserRuleContext()) {
			case 1:
				localctx = NewLrExprContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, FaultParserRULE_expression)
				p.SetState(512)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
				}
				{
					p.SetState(513)
					p.Match(FaultParserEXPO)
				}
				{
					p.SetState(514)
					p.expression(7)
				}

			case 2:
				localctx = NewLrExprContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, FaultParserRULE_expression)
				p.SetState(515)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				{
					p.SetState(516)
					_la = p.GetTokenStream().LA(1)

					if !(((_la-59)&-(0x1f+1)) == 0 && ((1<<uint((_la-59)))&((1<<(FaultParserAMPERSAND-59))|(1<<(FaultParserMULTI-59))|(1<<(FaultParserDIV-59))|(1<<(FaultParserMOD-59))|(1<<(FaultParserLSHIFT-59))|(1<<(FaultParserRSHIFT-59))|(1<<(FaultParserBIT_CLEAR-59)))) != 0) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(517)
					p.expression(6)
				}

			case 3:
				localctx = NewLrExprContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, FaultParserRULE_expression)
				p.SetState(518)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(519)
					_la = p.GetTokenStream().LA(1)

					if !(((_la-70)&-(0x1f+1)) == 0 && ((1<<uint((_la-70)))&((1<<(FaultParserPLUS-70))|(1<<(FaultParserMINUS-70))|(1<<(FaultParserCARET-70)))) != 0) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(520)
					p.expression(5)
				}

			case 4:
				localctx = NewLrExprContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, FaultParserRULE_expression)
				p.SetState(521)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(522)
					_la = p.GetTokenStream().LA(1)

					if !(((_la-62)&-(0x1f+1)) == 0 && ((1<<uint((_la-62)))&((1<<(FaultParserEQUALS-62))|(1<<(FaultParserNOT_EQUALS-62))|(1<<(FaultParserLESS-62))|(1<<(FaultParserLESS_OR_EQUALS-62))|(1<<(FaultParserGREATER-62))|(1<<(FaultParserGREATER_OR_EQUALS-62)))) != 0) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(523)
					p.expression(4)
				}

			case 5:
				localctx = NewLrExprContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, FaultParserRULE_expression)
				p.SetState(524)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				{
					p.SetState(525)
					p.Match(FaultParserAND)
				}
				{
					p.SetState(526)
					p.expression(3)
				}

			case 6:
				localctx = NewLrExprContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, FaultParserRULE_expression)
				p.SetState(527)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
				}
				{
					p.SetState(528)
					p.Match(FaultParserOR)
				}
				{
					p.SetState(529)
					p.expression(2)
				}

			}

		}
		p.SetState(534)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 47, p.GetParserRuleContext())
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
	this := p
	_ = this

	localctx = NewOperandContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 90, FaultParserRULE_operand)

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

	p.SetState(545)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 48, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(535)
			p.Match(FaultParserNIL)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(536)
			p.Numeric()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(537)
			p.String_()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(538)
			p.Bool_()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(539)
			p.OperandName()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(540)
			p.AccessHistory()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(541)
			p.Match(FaultParserLPAREN)
		}
		{
			p.SetState(542)
			p.expression(0)
		}
		{
			p.SetState(543)
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
	this := p
	_ = this

	localctx = NewOperandNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 92, FaultParserRULE_operandName)

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

	p.SetState(557)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 50, p.GetParserRuleContext()) {
	case 1:
		localctx = NewOpNameContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(547)
			p.Match(FaultParserIDENT)
		}

	case 2:
		localctx = NewOpParamContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(548)
			p.ParamCall()
		}

	case 3:
		localctx = NewOpThisContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(549)
			p.Match(FaultParserTHIS)
		}

	case 4:
		localctx = NewOpClockContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(550)
			p.Match(FaultParserCLOCK)
		}

	case 5:
		localctx = NewOpInstanceContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(551)
			p.Match(FaultParserNEW)
		}
		{
			p.SetState(552)
			p.Match(FaultParserIDENT)
		}
		p.SetState(555)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 49, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(553)
				p.Match(FaultParserDOT)
			}
			{
				p.SetState(554)
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
	this := p
	_ = this

	localctx = NewPrefixContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 94, FaultParserRULE_prefix)
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

	p.SetState(562)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 51, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(560)
			_la = p.GetTokenStream().LA(1)

			if !(((_la-59)&-(0x1f+1)) == 0 && ((1<<uint((_la-59)))&((1<<(FaultParserAMPERSAND-59))|(1<<(FaultParserBANG-59))|(1<<(FaultParserPLUS-59))|(1<<(FaultParserMINUS-59))|(1<<(FaultParserCARET-59))|(1<<(FaultParserMULTI-59)))) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(561)
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
	this := p
	_ = this

	localctx = NewNumericContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 96, FaultParserRULE_numeric)

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

	p.SetState(567)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case FaultParserDECIMAL_LIT, FaultParserOCTAL_LIT, FaultParserHEX_LIT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(564)
			p.Integer()
		}

	case FaultParserMINUS:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(565)
			p.Negative()
		}

	case FaultParserFLOAT_LIT:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(566)
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
	p.EnterRule(localctx, 98, FaultParserRULE_integer)
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
		p.SetState(569)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-80)&-(0x1f+1)) == 0 && ((1<<uint((_la-80)))&((1<<(FaultParserDECIMAL_LIT-80))|(1<<(FaultParserOCTAL_LIT-80))|(1<<(FaultParserHEX_LIT-80)))) != 0) {
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
	this := p
	_ = this

	localctx = NewNegativeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 100, FaultParserRULE_negative)

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

	p.SetState(575)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 53, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(571)
			p.Match(FaultParserMINUS)
		}
		{
			p.SetState(572)
			p.Integer()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(573)
			p.Match(FaultParserMINUS)
		}
		{
			p.SetState(574)
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
	p.EnterRule(localctx, 102, FaultParserRULE_float_)

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
		p.SetState(577)
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
	p.EnterRule(localctx, 104, FaultParserRULE_string_)
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
		p.SetState(579)
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
	p.EnterRule(localctx, 106, FaultParserRULE_bool_)
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
		p.SetState(581)
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
	this := p
	_ = this

	localctx = NewFunctionLitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 108, FaultParserRULE_functionLit)

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
		p.SetState(583)
		p.Match(FaultParserFUNC)
	}
	{
		p.SetState(584)
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
	this := p
	_ = this

	localctx = NewEosContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 110, FaultParserRULE_eos)
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
		p.SetState(586)
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
	case 44:
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
	this := p
	_ = this

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
