// Code generated from /home/vikimaster2/Projects/Go/abacus/Abacus.g4 by ANTLR 4.9.1. DO NOT EDIT.

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 55, 353,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23,
	9, 23, 4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9,
	28, 4, 29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33,
	4, 34, 9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4,
	39, 9, 39, 4, 40, 9, 40, 4, 41, 9, 41, 4, 42, 9, 42, 4, 43, 9, 43, 4, 44,
	9, 44, 4, 45, 9, 45, 4, 46, 9, 46, 4, 47, 9, 47, 4, 48, 9, 48, 4, 49, 9,
	49, 4, 50, 9, 50, 4, 51, 9, 51, 4, 52, 9, 52, 4, 53, 9, 53, 4, 54, 9, 54,
	4, 55, 9, 55, 4, 56, 9, 56, 4, 57, 9, 57, 4, 58, 9, 58, 3, 2, 3, 2, 3,
	2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 3,
	5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 8, 3,
	8, 3, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3,
	11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12,
	3, 12, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 14, 3, 14, 3, 14, 3, 14, 3,
	15, 3, 15, 3, 15, 3, 15, 3, 16, 3, 16, 3, 16, 3, 16, 3, 17, 3, 17, 3, 17,
	3, 17, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 19, 3, 19, 3, 19, 3,
	19, 3, 19, 3, 20, 3, 20, 3, 20, 3, 20, 3, 21, 3, 21, 3, 21, 3, 21, 3, 22,
	3, 22, 3, 22, 3, 22, 3, 23, 3, 23, 3, 23, 3, 23, 3, 24, 3, 24, 3, 24, 3,
	24, 3, 24, 3, 24, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 26, 3, 26, 3, 26,
	3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 27, 3, 27, 3, 27, 3, 27, 3, 28, 3,
	28, 3, 28, 3, 29, 3, 29, 3, 29, 3, 30, 3, 30, 3, 30, 3, 30, 3, 31, 3, 31,
	3, 32, 3, 32, 3, 33, 3, 33, 3, 34, 3, 34, 3, 35, 3, 35, 3, 35, 3, 35, 5,
	35, 260, 10, 35, 3, 36, 3, 36, 3, 36, 5, 36, 265, 10, 36, 3, 37, 3, 37,
	3, 38, 3, 38, 3, 39, 3, 39, 3, 40, 3, 40, 3, 41, 3, 41, 3, 42, 3, 42, 3,
	43, 3, 43, 3, 44, 3, 44, 3, 45, 3, 45, 3, 46, 3, 46, 3, 47, 3, 47, 3, 47,
	3, 47, 3, 47, 3, 47, 5, 47, 293, 10, 47, 3, 48, 3, 48, 3, 48, 5, 48, 298,
	10, 48, 3, 48, 5, 48, 301, 10, 48, 3, 49, 3, 49, 3, 50, 6, 50, 306, 10,
	50, 13, 50, 14, 50, 307, 3, 50, 3, 50, 6, 50, 312, 10, 50, 13, 50, 14,
	50, 313, 5, 50, 316, 10, 50, 3, 51, 3, 51, 7, 51, 320, 10, 51, 12, 51,
	14, 51, 323, 11, 51, 3, 52, 3, 52, 7, 52, 327, 10, 52, 12, 52, 14, 52,
	330, 11, 52, 3, 53, 3, 53, 5, 53, 334, 10, 53, 3, 54, 3, 54, 3, 54, 5,
	54, 339, 10, 54, 3, 55, 3, 55, 3, 56, 3, 56, 3, 57, 3, 57, 3, 58, 6, 58,
	348, 10, 58, 13, 58, 14, 58, 349, 3, 58, 3, 58, 2, 2, 59, 3, 3, 5, 4, 7,
	5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19, 11, 21, 12, 23, 13, 25, 14, 27,
	15, 29, 16, 31, 17, 33, 18, 35, 19, 37, 20, 39, 21, 41, 22, 43, 23, 45,
	24, 47, 25, 49, 26, 51, 27, 53, 28, 55, 29, 57, 30, 59, 31, 61, 32, 63,
	33, 65, 34, 67, 35, 69, 36, 71, 37, 73, 38, 75, 39, 77, 40, 79, 41, 81,
	42, 83, 43, 85, 44, 87, 45, 89, 46, 91, 47, 93, 48, 95, 49, 97, 2, 99,
	2, 101, 50, 103, 51, 105, 2, 107, 2, 109, 52, 111, 53, 113, 54, 115, 55,
	3, 2, 6, 5, 2, 35, 35, 128, 128, 174, 174, 4, 2, 71, 71, 103, 103, 4, 2,
	45, 45, 47, 47, 5, 2, 11, 12, 15, 15, 34, 34, 2, 363, 2, 3, 3, 2, 2, 2,
	2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2,
	2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2,
	2, 2, 21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 2, 25, 3, 2, 2, 2, 2, 27, 3, 2,
	2, 2, 2, 29, 3, 2, 2, 2, 2, 31, 3, 2, 2, 2, 2, 33, 3, 2, 2, 2, 2, 35, 3,
	2, 2, 2, 2, 37, 3, 2, 2, 2, 2, 39, 3, 2, 2, 2, 2, 41, 3, 2, 2, 2, 2, 43,
	3, 2, 2, 2, 2, 45, 3, 2, 2, 2, 2, 47, 3, 2, 2, 2, 2, 49, 3, 2, 2, 2, 2,
	51, 3, 2, 2, 2, 2, 53, 3, 2, 2, 2, 2, 55, 3, 2, 2, 2, 2, 57, 3, 2, 2, 2,
	2, 59, 3, 2, 2, 2, 2, 61, 3, 2, 2, 2, 2, 63, 3, 2, 2, 2, 2, 65, 3, 2, 2,
	2, 2, 67, 3, 2, 2, 2, 2, 69, 3, 2, 2, 2, 2, 71, 3, 2, 2, 2, 2, 73, 3, 2,
	2, 2, 2, 75, 3, 2, 2, 2, 2, 77, 3, 2, 2, 2, 2, 79, 3, 2, 2, 2, 2, 81, 3,
	2, 2, 2, 2, 83, 3, 2, 2, 2, 2, 85, 3, 2, 2, 2, 2, 87, 3, 2, 2, 2, 2, 89,
	3, 2, 2, 2, 2, 91, 3, 2, 2, 2, 2, 93, 3, 2, 2, 2, 2, 95, 3, 2, 2, 2, 2,
	101, 3, 2, 2, 2, 2, 103, 3, 2, 2, 2, 2, 109, 3, 2, 2, 2, 2, 111, 3, 2,
	2, 2, 2, 113, 3, 2, 2, 2, 2, 115, 3, 2, 2, 2, 3, 117, 3, 2, 2, 2, 5, 122,
	3, 2, 2, 2, 7, 128, 3, 2, 2, 2, 9, 130, 3, 2, 2, 2, 11, 132, 3, 2, 2, 2,
	13, 137, 3, 2, 2, 2, 15, 142, 3, 2, 2, 2, 17, 145, 3, 2, 2, 2, 19, 149,
	3, 2, 2, 2, 21, 154, 3, 2, 2, 2, 23, 160, 3, 2, 2, 2, 25, 166, 3, 2, 2,
	2, 27, 171, 3, 2, 2, 2, 29, 175, 3, 2, 2, 2, 31, 179, 3, 2, 2, 2, 33, 183,
	3, 2, 2, 2, 35, 187, 3, 2, 2, 2, 37, 193, 3, 2, 2, 2, 39, 198, 3, 2, 2,
	2, 41, 202, 3, 2, 2, 2, 43, 206, 3, 2, 2, 2, 45, 210, 3, 2, 2, 2, 47, 214,
	3, 2, 2, 2, 49, 220, 3, 2, 2, 2, 51, 225, 3, 2, 2, 2, 53, 233, 3, 2, 2,
	2, 55, 237, 3, 2, 2, 2, 57, 240, 3, 2, 2, 2, 59, 243, 3, 2, 2, 2, 61, 247,
	3, 2, 2, 2, 63, 249, 3, 2, 2, 2, 65, 251, 3, 2, 2, 2, 67, 253, 3, 2, 2,
	2, 69, 259, 3, 2, 2, 2, 71, 264, 3, 2, 2, 2, 73, 266, 3, 2, 2, 2, 75, 268,
	3, 2, 2, 2, 77, 270, 3, 2, 2, 2, 79, 272, 3, 2, 2, 2, 81, 274, 3, 2, 2,
	2, 83, 276, 3, 2, 2, 2, 85, 278, 3, 2, 2, 2, 87, 280, 3, 2, 2, 2, 89, 282,
	3, 2, 2, 2, 91, 284, 3, 2, 2, 2, 93, 292, 3, 2, 2, 2, 95, 294, 3, 2, 2,
	2, 97, 302, 3, 2, 2, 2, 99, 305, 3, 2, 2, 2, 101, 317, 3, 2, 2, 2, 103,
	324, 3, 2, 2, 2, 105, 333, 3, 2, 2, 2, 107, 338, 3, 2, 2, 2, 109, 340,
	3, 2, 2, 2, 111, 342, 3, 2, 2, 2, 113, 344, 3, 2, 2, 2, 115, 347, 3, 2,
	2, 2, 117, 118, 7, 118, 2, 2, 118, 119, 7, 116, 2, 2, 119, 120, 7, 119,
	2, 2, 120, 121, 7, 103, 2, 2, 121, 4, 3, 2, 2, 2, 122, 123, 7, 104, 2,
	2, 123, 124, 7, 99, 2, 2, 124, 125, 7, 110, 2, 2, 125, 126, 7, 117, 2,
	2, 126, 127, 7, 103, 2, 2, 127, 6, 3, 2, 2, 2, 128, 129, 7, 60, 2, 2, 129,
	8, 3, 2, 2, 2, 130, 131, 7, 46, 2, 2, 131, 10, 3, 2, 2, 2, 132, 133, 7,
	117, 2, 2, 133, 134, 7, 115, 2, 2, 134, 135, 7, 116, 2, 2, 135, 136, 7,
	118, 2, 2, 136, 12, 3, 2, 2, 2, 137, 138, 7, 101, 2, 2, 138, 139, 7, 100,
	2, 2, 139, 140, 7, 116, 2, 2, 140, 141, 7, 118, 2, 2, 141, 14, 3, 2, 2,
	2, 142, 143, 7, 110, 2, 2, 143, 144, 7, 112, 2, 2, 144, 16, 3, 2, 2, 2,
	145, 146, 7, 110, 2, 2, 146, 147, 7, 113, 2, 2, 147, 148, 7, 105, 2, 2,
	148, 18, 3, 2, 2, 2, 149, 150, 7, 110, 2, 2, 150, 151, 7, 113, 2, 2, 151,
	152, 7, 105, 2, 2, 152, 153, 7, 52, 2, 2, 153, 20, 3, 2, 2, 2, 154, 155,
	7, 110, 2, 2, 155, 156, 7, 113, 2, 2, 156, 157, 7, 105, 2, 2, 157, 158,
	7, 51, 2, 2, 158, 159, 7, 50, 2, 2, 159, 22, 3, 2, 2, 2, 160, 161, 7, 104,
	2, 2, 161, 162, 7, 110, 2, 2, 162, 163, 7, 113, 2, 2, 163, 164, 7, 113,
	2, 2, 164, 165, 7, 116, 2, 2, 165, 24, 3, 2, 2, 2, 166, 167, 7, 101, 2,
	2, 167, 168, 7, 103, 2, 2, 168, 169, 7, 107, 2, 2, 169, 170, 7, 110, 2,
	2, 170, 26, 3, 2, 2, 2, 171, 172, 7, 103, 2, 2, 172, 173, 7, 122, 2, 2,
	173, 174, 7, 114, 2, 2, 174, 28, 3, 2, 2, 2, 175, 176, 7, 117, 2, 2, 176,
	177, 7, 107, 2, 2, 177, 178, 7, 112, 2, 2, 178, 30, 3, 2, 2, 2, 179, 180,
	7, 101, 2, 2, 180, 181, 7, 113, 2, 2, 181, 182, 7, 117, 2, 2, 182, 32,
	3, 2, 2, 2, 183, 184, 7, 118, 2, 2, 184, 185, 7, 99, 2, 2, 185, 186, 7,
	112, 2, 2, 186, 34, 3, 2, 2, 2, 187, 188, 7, 116, 2, 2, 188, 189, 7, 113,
	2, 2, 189, 190, 7, 119, 2, 2, 190, 191, 7, 112, 2, 2, 191, 192, 7, 102,
	2, 2, 192, 36, 3, 2, 2, 2, 193, 194, 7, 117, 2, 2, 194, 195, 7, 107, 2,
	2, 195, 196, 7, 105, 2, 2, 196, 197, 7, 112, 2, 2, 197, 38, 3, 2, 2, 2,
	198, 199, 7, 99, 2, 2, 199, 200, 7, 100, 2, 2, 200, 201, 7, 117, 2, 2,
	201, 40, 3, 2, 2, 2, 202, 203, 7, 111, 2, 2, 203, 204, 7, 107, 2, 2, 204,
	205, 7, 112, 2, 2, 205, 42, 3, 2, 2, 2, 206, 207, 7, 111, 2, 2, 207, 208,
	7, 99, 2, 2, 208, 209, 7, 122, 2, 2, 209, 44, 3, 2, 2, 2, 210, 211, 7,
	99, 2, 2, 211, 212, 7, 120, 2, 2, 212, 213, 7, 105, 2, 2, 213, 46, 3, 2,
	2, 2, 214, 215, 7, 119, 2, 2, 215, 216, 7, 112, 2, 2, 216, 217, 7, 118,
	2, 2, 217, 218, 7, 107, 2, 2, 218, 219, 7, 110, 2, 2, 219, 48, 3, 2, 2,
	2, 220, 221, 7, 104, 2, 2, 221, 222, 7, 116, 2, 2, 222, 223, 7, 113, 2,
	2, 223, 224, 7, 111, 2, 2, 224, 50, 3, 2, 2, 2, 225, 226, 7, 116, 2, 2,
	226, 227, 7, 103, 2, 2, 227, 228, 7, 120, 2, 2, 228, 229, 7, 103, 2, 2,
	229, 230, 7, 116, 2, 2, 230, 231, 7, 117, 2, 2, 231, 232, 7, 103, 2, 2,
	232, 52, 3, 2, 2, 2, 233, 234, 7, 112, 2, 2, 234, 235, 7, 118, 2, 2, 235,
	236, 7, 106, 2, 2, 236, 54, 3, 2, 2, 2, 237, 238, 7, 40, 2, 2, 238, 239,
	7, 40, 2, 2, 239, 56, 3, 2, 2, 2, 240, 241, 7, 126, 2, 2, 241, 242, 7,
	126, 2, 2, 242, 58, 3, 2, 2, 2, 243, 244, 7, 122, 2, 2, 244, 245, 7, 113,
	2, 2, 245, 246, 7, 116, 2, 2, 246, 60, 3, 2, 2, 2, 247, 248, 9, 2, 2, 2,
	248, 62, 3, 2, 2, 2, 249, 250, 7, 63, 2, 2, 250, 64, 3, 2, 2, 2, 251, 252,
	7, 62, 2, 2, 252, 66, 3, 2, 2, 2, 253, 254, 7, 64, 2, 2, 254, 68, 3, 2,
	2, 2, 255, 256, 7, 47, 2, 2, 256, 260, 7, 64, 2, 2, 257, 258, 7, 63, 2,
	2, 258, 260, 7, 64, 2, 2, 259, 255, 3, 2, 2, 2, 259, 257, 3, 2, 2, 2, 260,
	70, 3, 2, 2, 2, 261, 265, 7, 96, 2, 2, 262, 263, 7, 44, 2, 2, 263, 265,
	7, 44, 2, 2, 264, 261, 3, 2, 2, 2, 264, 262, 3, 2, 2, 2, 265, 72, 3, 2,
	2, 2, 266, 267, 7, 44, 2, 2, 267, 74, 3, 2, 2, 2, 268, 269, 7, 49, 2, 2,
	269, 76, 3, 2, 2, 2, 270, 271, 7, 45, 2, 2, 271, 78, 3, 2, 2, 2, 272, 273,
	7, 47, 2, 2, 273, 80, 3, 2, 2, 2, 274, 275, 7, 39, 2, 2, 275, 82, 3, 2,
	2, 2, 276, 277, 7, 48, 2, 2, 277, 84, 3, 2, 2, 2, 278, 279, 7, 42, 2, 2,
	279, 86, 3, 2, 2, 2, 280, 281, 7, 43, 2, 2, 281, 88, 3, 2, 2, 2, 282, 283,
	7, 93, 2, 2, 283, 90, 3, 2, 2, 2, 284, 285, 7, 95, 2, 2, 285, 92, 3, 2,
	2, 2, 286, 287, 7, 114, 2, 2, 287, 293, 7, 107, 2, 2, 288, 293, 7, 103,
	2, 2, 289, 290, 7, 114, 2, 2, 290, 291, 7, 106, 2, 2, 291, 293, 7, 107,
	2, 2, 292, 286, 3, 2, 2, 2, 292, 288, 3, 2, 2, 2, 292, 289, 3, 2, 2, 2,
	293, 94, 3, 2, 2, 2, 294, 300, 5, 99, 50, 2, 295, 297, 9, 3, 2, 2, 296,
	298, 5, 97, 49, 2, 297, 296, 3, 2, 2, 2, 297, 298, 3, 2, 2, 2, 298, 299,
	3, 2, 2, 2, 299, 301, 5, 99, 50, 2, 300, 295, 3, 2, 2, 2, 300, 301, 3,
	2, 2, 2, 301, 96, 3, 2, 2, 2, 302, 303, 9, 4, 2, 2, 303, 98, 3, 2, 2, 2,
	304, 306, 5, 109, 55, 2, 305, 304, 3, 2, 2, 2, 306, 307, 3, 2, 2, 2, 307,
	305, 3, 2, 2, 2, 307, 308, 3, 2, 2, 2, 308, 315, 3, 2, 2, 2, 309, 311,
	5, 83, 42, 2, 310, 312, 5, 109, 55, 2, 311, 310, 3, 2, 2, 2, 312, 313,
	3, 2, 2, 2, 313, 311, 3, 2, 2, 2, 313, 314, 3, 2, 2, 2, 314, 316, 3, 2,
	2, 2, 315, 309, 3, 2, 2, 2, 315, 316, 3, 2, 2, 2, 316, 100, 3, 2, 2, 2,
	317, 321, 5, 105, 53, 2, 318, 320, 5, 107, 54, 2, 319, 318, 3, 2, 2, 2,
	320, 323, 3, 2, 2, 2, 321, 319, 3, 2, 2, 2, 321, 322, 3, 2, 2, 2, 322,
	102, 3, 2, 2, 2, 323, 321, 3, 2, 2, 2, 324, 328, 5, 111, 56, 2, 325, 327,
	5, 107, 54, 2, 326, 325, 3, 2, 2, 2, 327, 330, 3, 2, 2, 2, 328, 326, 3,
	2, 2, 2, 328, 329, 3, 2, 2, 2, 329, 104, 3, 2, 2, 2, 330, 328, 3, 2, 2,
	2, 331, 334, 5, 113, 57, 2, 332, 334, 7, 97, 2, 2, 333, 331, 3, 2, 2, 2,
	333, 332, 3, 2, 2, 2, 334, 106, 3, 2, 2, 2, 335, 339, 5, 105, 53, 2, 336,
	339, 5, 111, 56, 2, 337, 339, 5, 109, 55, 2, 338, 335, 3, 2, 2, 2, 338,
	336, 3, 2, 2, 2, 338, 337, 3, 2, 2, 2, 339, 108, 3, 2, 2, 2, 340, 341,
	4, 50, 59, 2, 341, 110, 3, 2, 2, 2, 342, 343, 4, 67, 92, 2, 343, 112, 3,
	2, 2, 2, 344, 345, 4, 99, 124, 2, 345, 114, 3, 2, 2, 2, 346, 348, 9, 5,
	2, 2, 347, 346, 3, 2, 2, 2, 348, 349, 3, 2, 2, 2, 349, 347, 3, 2, 2, 2,
	349, 350, 3, 2, 2, 2, 350, 351, 3, 2, 2, 2, 351, 352, 8, 58, 2, 2, 352,
	116, 3, 2, 2, 2, 16, 2, 259, 264, 292, 297, 300, 307, 313, 315, 321, 328,
	333, 338, 349, 3, 8, 2, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'true'", "'false'", "':'", "','", "'sqrt'", "'cbrt'", "'ln'", "'log'",
	"'log2'", "'log10'", "'floor'", "'ceil'", "'exp'", "'sin'", "'cos'", "'tan'",
	"'round'", "'sign'", "'abs'", "'min'", "'max'", "'avg'", "'until'", "'from'",
	"'reverse'", "'nth'", "'&&'", "'||'", "'xor'", "", "'='", "'<'", "'>'",
	"", "", "'*'", "'/'", "'+'", "'-'", "'%'", "'.'", "'('", "')'", "'['",
	"']'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "AND", "OR", "XOR", "NOT", "EQ", "LS",
	"GR", "ARROW", "POW", "MUL", "DIV", "ADD", "SUB", "PER", "POINT", "LPAREN",
	"RPAREN", "LSQPAREN", "RSQPAREN", "CONSTANT", "SCIENTIFIC_NUMBER", "VARIABLE",
	"LAMBDA_VARIABLE", "DIGITS", "UPPERCASE", "LOWERCASE", "WHITESPACE",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
	"T__9", "T__10", "T__11", "T__12", "T__13", "T__14", "T__15", "T__16",
	"T__17", "T__18", "T__19", "T__20", "T__21", "T__22", "T__23", "T__24",
	"T__25", "AND", "OR", "XOR", "NOT", "EQ", "LS", "GR", "ARROW", "POW", "MUL",
	"DIV", "ADD", "SUB", "PER", "POINT", "LPAREN", "RPAREN", "LSQPAREN", "RSQPAREN",
	"CONSTANT", "SCIENTIFIC_NUMBER", "SIGN", "NUMBER", "VARIABLE", "LAMBDA_VARIABLE",
	"VALID_ID_START", "VALID_ID_CHAR", "DIGITS", "UPPERCASE", "LOWERCASE",
	"WHITESPACE",
}

type AbacusLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

// NewAbacusLexer produces a new lexer instance for the optional input antlr.CharStream.
//
// The *AbacusLexer instance produced may be reused by calling the SetInputStream method.
// The initial lexer configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewAbacusLexer(input antlr.CharStream) *AbacusLexer {
	l := new(AbacusLexer)
	lexerDeserializer := antlr.NewATNDeserializer(nil)
	lexerAtn := lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)
	lexerDecisionToDFA := make([]*antlr.DFA, len(lexerAtn.DecisionToState))
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "Abacus.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// AbacusLexer tokens.
const (
	AbacusLexerT__0              = 1
	AbacusLexerT__1              = 2
	AbacusLexerT__2              = 3
	AbacusLexerT__3              = 4
	AbacusLexerT__4              = 5
	AbacusLexerT__5              = 6
	AbacusLexerT__6              = 7
	AbacusLexerT__7              = 8
	AbacusLexerT__8              = 9
	AbacusLexerT__9              = 10
	AbacusLexerT__10             = 11
	AbacusLexerT__11             = 12
	AbacusLexerT__12             = 13
	AbacusLexerT__13             = 14
	AbacusLexerT__14             = 15
	AbacusLexerT__15             = 16
	AbacusLexerT__16             = 17
	AbacusLexerT__17             = 18
	AbacusLexerT__18             = 19
	AbacusLexerT__19             = 20
	AbacusLexerT__20             = 21
	AbacusLexerT__21             = 22
	AbacusLexerT__22             = 23
	AbacusLexerT__23             = 24
	AbacusLexerT__24             = 25
	AbacusLexerT__25             = 26
	AbacusLexerAND               = 27
	AbacusLexerOR                = 28
	AbacusLexerXOR               = 29
	AbacusLexerNOT               = 30
	AbacusLexerEQ                = 31
	AbacusLexerLS                = 32
	AbacusLexerGR                = 33
	AbacusLexerARROW             = 34
	AbacusLexerPOW               = 35
	AbacusLexerMUL               = 36
	AbacusLexerDIV               = 37
	AbacusLexerADD               = 38
	AbacusLexerSUB               = 39
	AbacusLexerPER               = 40
	AbacusLexerPOINT             = 41
	AbacusLexerLPAREN            = 42
	AbacusLexerRPAREN            = 43
	AbacusLexerLSQPAREN          = 44
	AbacusLexerRSQPAREN          = 45
	AbacusLexerCONSTANT          = 46
	AbacusLexerSCIENTIFIC_NUMBER = 47
	AbacusLexerVARIABLE          = 48
	AbacusLexerLAMBDA_VARIABLE   = 49
	AbacusLexerDIGITS            = 50
	AbacusLexerUPPERCASE         = 51
	AbacusLexerLOWERCASE         = 52
	AbacusLexerWHITESPACE        = 53
)
