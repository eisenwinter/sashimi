// Code generated from C:\gitrepos\sashimi/grammar/Sashimi.g4 by ANTLR 4.8. DO NOT EDIT.

package build

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 40, 302,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23,
	9, 23, 4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9,
	28, 4, 29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33,
	4, 34, 9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4,
	39, 9, 39, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 5, 3, 120, 10, 3, 3, 4, 3, 4, 3, 4,
	3, 4, 3, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 6,
	3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 8, 3, 8, 3, 9, 3, 9,
	3, 9, 3, 10, 3, 10, 3, 10, 3, 11, 3, 11, 3, 11, 3, 12, 3, 12, 3, 12, 3,
	12, 3, 12, 3, 12, 3, 13, 3, 13, 3, 13, 3, 13, 3, 14, 3, 14, 3, 14, 3, 15,
	3, 15, 3, 16, 3, 16, 3, 17, 3, 17, 3, 18, 3, 18, 3, 19, 3, 19, 3, 20, 3,
	20, 3, 20, 3, 21, 3, 21, 3, 22, 3, 22, 3, 22, 3, 23, 3, 23, 3, 24, 3, 24,
	3, 25, 3, 25, 3, 25, 3, 26, 3, 26, 3, 27, 3, 27, 3, 28, 3, 28, 3, 29, 3,
	29, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30,
	3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3,
	30, 3, 30, 3, 30, 3, 30, 3, 30, 5, 30, 227, 10, 30, 3, 31, 3, 31, 3, 31,
	3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 5, 31, 240, 10,
	31, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 33, 3, 33, 7, 33, 249, 10, 33,
	12, 33, 14, 33, 252, 11, 33, 3, 33, 3, 33, 3, 34, 6, 34, 257, 10, 34, 13,
	34, 14, 34, 258, 3, 35, 3, 35, 6, 35, 263, 10, 35, 13, 35, 14, 35, 264,
	3, 35, 3, 35, 3, 36, 5, 36, 270, 10, 36, 3, 36, 6, 36, 273, 10, 36, 13,
	36, 14, 36, 274, 3, 36, 3, 36, 6, 36, 279, 10, 36, 13, 36, 14, 36, 280,
	5, 36, 283, 10, 36, 3, 37, 3, 37, 3, 37, 3, 37, 3, 37, 3, 38, 3, 38, 3,
	38, 3, 38, 3, 38, 3, 38, 3, 39, 6, 39, 297, 10, 39, 13, 39, 14, 39, 298,
	3, 39, 3, 39, 2, 2, 40, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17,
	10, 19, 11, 21, 12, 23, 13, 25, 14, 27, 15, 29, 16, 31, 17, 33, 18, 35,
	19, 37, 20, 39, 21, 41, 22, 43, 23, 45, 24, 47, 25, 49, 26, 51, 27, 53,
	28, 55, 29, 57, 30, 59, 31, 61, 32, 63, 33, 65, 34, 67, 35, 69, 36, 71,
	37, 73, 38, 75, 39, 77, 40, 3, 2, 7, 5, 2, 12, 12, 15, 15, 36, 36, 6, 2,
	50, 59, 67, 92, 97, 97, 99, 124, 5, 2, 50, 60, 67, 92, 99, 124, 3, 2, 50,
	59, 5, 2, 11, 12, 15, 15, 34, 34, 2, 317, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2,
	2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3,
	2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21,
	3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 2, 25, 3, 2, 2, 2, 2, 27, 3, 2, 2, 2, 2,
	29, 3, 2, 2, 2, 2, 31, 3, 2, 2, 2, 2, 33, 3, 2, 2, 2, 2, 35, 3, 2, 2, 2,
	2, 37, 3, 2, 2, 2, 2, 39, 3, 2, 2, 2, 2, 41, 3, 2, 2, 2, 2, 43, 3, 2, 2,
	2, 2, 45, 3, 2, 2, 2, 2, 47, 3, 2, 2, 2, 2, 49, 3, 2, 2, 2, 2, 51, 3, 2,
	2, 2, 2, 53, 3, 2, 2, 2, 2, 55, 3, 2, 2, 2, 2, 57, 3, 2, 2, 2, 2, 59, 3,
	2, 2, 2, 2, 61, 3, 2, 2, 2, 2, 63, 3, 2, 2, 2, 2, 65, 3, 2, 2, 2, 2, 67,
	3, 2, 2, 2, 2, 69, 3, 2, 2, 2, 2, 71, 3, 2, 2, 2, 2, 73, 3, 2, 2, 2, 2,
	75, 3, 2, 2, 2, 2, 77, 3, 2, 2, 2, 3, 79, 3, 2, 2, 2, 5, 119, 3, 2, 2,
	2, 7, 121, 3, 2, 2, 2, 9, 128, 3, 2, 2, 2, 11, 135, 3, 2, 2, 2, 13, 142,
	3, 2, 2, 2, 15, 144, 3, 2, 2, 2, 17, 146, 3, 2, 2, 2, 19, 149, 3, 2, 2,
	2, 21, 152, 3, 2, 2, 2, 23, 155, 3, 2, 2, 2, 25, 161, 3, 2, 2, 2, 27, 165,
	3, 2, 2, 2, 29, 168, 3, 2, 2, 2, 31, 170, 3, 2, 2, 2, 33, 172, 3, 2, 2,
	2, 35, 174, 3, 2, 2, 2, 37, 176, 3, 2, 2, 2, 39, 178, 3, 2, 2, 2, 41, 181,
	3, 2, 2, 2, 43, 183, 3, 2, 2, 2, 45, 186, 3, 2, 2, 2, 47, 188, 3, 2, 2,
	2, 49, 190, 3, 2, 2, 2, 51, 193, 3, 2, 2, 2, 53, 195, 3, 2, 2, 2, 55, 197,
	3, 2, 2, 2, 57, 199, 3, 2, 2, 2, 59, 226, 3, 2, 2, 2, 61, 239, 3, 2, 2,
	2, 63, 241, 3, 2, 2, 2, 65, 246, 3, 2, 2, 2, 67, 256, 3, 2, 2, 2, 69, 260,
	3, 2, 2, 2, 71, 269, 3, 2, 2, 2, 73, 284, 3, 2, 2, 2, 75, 289, 3, 2, 2,
	2, 77, 296, 3, 2, 2, 2, 79, 80, 7, 117, 2, 2, 80, 81, 7, 99, 2, 2, 81,
	82, 7, 117, 2, 2, 82, 83, 7, 106, 2, 2, 83, 84, 7, 107, 2, 2, 84, 85, 7,
	111, 2, 2, 85, 86, 7, 107, 2, 2, 86, 87, 7, 60, 2, 2, 87, 4, 3, 2, 2, 2,
	88, 89, 7, 102, 2, 2, 89, 90, 7, 107, 2, 2, 90, 91, 7, 117, 2, 2, 91, 92,
	7, 114, 2, 2, 92, 93, 7, 110, 2, 2, 93, 94, 7, 99, 2, 2, 94, 120, 7, 123,
	2, 2, 95, 96, 7, 110, 2, 2, 96, 97, 7, 99, 2, 2, 97, 98, 7, 123, 2, 2,
	98, 99, 7, 113, 2, 2, 99, 100, 7, 119, 2, 2, 100, 101, 7, 118, 2, 2, 101,
	102, 7, 97, 2, 2, 102, 103, 7, 117, 2, 2, 103, 104, 7, 103, 2, 2, 104,
	105, 7, 101, 2, 2, 105, 106, 7, 118, 2, 2, 106, 107, 7, 107, 2, 2, 107,
	108, 7, 113, 2, 2, 108, 120, 7, 112, 2, 2, 109, 110, 7, 110, 2, 2, 110,
	111, 7, 99, 2, 2, 111, 112, 7, 123, 2, 2, 112, 113, 7, 113, 2, 2, 113,
	114, 7, 119, 2, 2, 114, 120, 7, 118, 2, 2, 115, 116, 7, 110, 2, 2, 116,
	117, 7, 107, 2, 2, 117, 118, 7, 112, 2, 2, 118, 120, 7, 109, 2, 2, 119,
	88, 3, 2, 2, 2, 119, 95, 3, 2, 2, 2, 119, 109, 3, 2, 2, 2, 119, 115, 3,
	2, 2, 2, 120, 6, 3, 2, 2, 2, 121, 122, 7, 116, 2, 2, 122, 123, 7, 103,
	2, 2, 123, 124, 7, 114, 2, 2, 124, 125, 7, 103, 2, 2, 125, 126, 7, 99,
	2, 2, 126, 127, 7, 118, 2, 2, 127, 8, 3, 2, 2, 2, 128, 129, 7, 103, 2,
	2, 129, 130, 7, 112, 2, 2, 130, 131, 7, 118, 2, 2, 131, 132, 7, 107, 2,
	2, 132, 133, 7, 118, 2, 2, 133, 134, 7, 123, 2, 2, 134, 10, 3, 2, 2, 2,
	135, 136, 7, 119, 2, 2, 136, 137, 7, 112, 2, 2, 137, 138, 7, 107, 2, 2,
	138, 139, 7, 115, 2, 2, 139, 140, 7, 119, 2, 2, 140, 141, 7, 103, 2, 2,
	141, 12, 3, 2, 2, 2, 142, 143, 7, 46, 2, 2, 143, 14, 3, 2, 2, 2, 144, 145,
	7, 47, 2, 2, 145, 16, 3, 2, 2, 2, 146, 147, 7, 113, 2, 2, 147, 148, 7,
	104, 2, 2, 148, 18, 3, 2, 2, 2, 149, 150, 7, 107, 2, 2, 150, 151, 7, 117,
	2, 2, 151, 20, 3, 2, 2, 2, 152, 153, 7, 99, 2, 2, 153, 154, 7, 117, 2,
	2, 154, 22, 3, 2, 2, 2, 155, 156, 7, 100, 2, 2, 156, 157, 7, 103, 2, 2,
	157, 158, 7, 105, 2, 2, 158, 159, 7, 107, 2, 2, 159, 160, 7, 112, 2, 2,
	160, 24, 3, 2, 2, 2, 161, 162, 7, 103, 2, 2, 162, 163, 7, 112, 2, 2, 163,
	164, 7, 102, 2, 2, 164, 26, 3, 2, 2, 2, 165, 166, 7, 47, 2, 2, 166, 167,
	7, 64, 2, 2, 167, 28, 3, 2, 2, 2, 168, 169, 7, 42, 2, 2, 169, 30, 3, 2,
	2, 2, 170, 171, 7, 43, 2, 2, 171, 32, 3, 2, 2, 2, 172, 173, 7, 93, 2, 2,
	173, 34, 3, 2, 2, 2, 174, 175, 7, 95, 2, 2, 175, 36, 3, 2, 2, 2, 176, 177,
	7, 63, 2, 2, 177, 38, 3, 2, 2, 2, 178, 179, 7, 62, 2, 2, 179, 180, 7, 63,
	2, 2, 180, 40, 3, 2, 2, 2, 181, 182, 7, 62, 2, 2, 182, 42, 3, 2, 2, 2,
	183, 184, 7, 64, 2, 2, 184, 185, 7, 63, 2, 2, 185, 44, 3, 2, 2, 2, 186,
	187, 7, 64, 2, 2, 187, 46, 3, 2, 2, 2, 188, 189, 7, 35, 2, 2, 189, 48,
	3, 2, 2, 2, 190, 191, 7, 62, 2, 2, 191, 192, 7, 64, 2, 2, 192, 50, 3, 2,
	2, 2, 193, 194, 7, 40, 2, 2, 194, 52, 3, 2, 2, 2, 195, 196, 7, 126, 2,
	2, 196, 54, 3, 2, 2, 2, 197, 198, 7, 60, 2, 2, 198, 56, 3, 2, 2, 2, 199,
	200, 7, 48, 2, 2, 200, 58, 3, 2, 2, 2, 201, 202, 7, 118, 2, 2, 202, 203,
	7, 103, 2, 2, 203, 204, 7, 122, 2, 2, 204, 227, 7, 118, 2, 2, 205, 206,
	7, 114, 2, 2, 206, 207, 7, 107, 2, 2, 207, 208, 7, 101, 2, 2, 208, 209,
	7, 118, 2, 2, 209, 210, 7, 119, 2, 2, 210, 211, 7, 116, 2, 2, 211, 227,
	7, 103, 2, 2, 212, 213, 7, 112, 2, 2, 213, 214, 7, 119, 2, 2, 214, 215,
	7, 111, 2, 2, 215, 216, 7, 100, 2, 2, 216, 217, 7, 103, 2, 2, 217, 227,
	7, 116, 2, 2, 218, 219, 7, 100, 2, 2, 219, 220, 7, 113, 2, 2, 220, 221,
	7, 113, 2, 2, 221, 227, 7, 110, 2, 2, 222, 223, 7, 102, 2, 2, 223, 224,
	7, 99, 2, 2, 224, 225, 7, 118, 2, 2, 225, 227, 7, 103, 2, 2, 226, 201,
	3, 2, 2, 2, 226, 205, 3, 2, 2, 2, 226, 212, 3, 2, 2, 2, 226, 218, 3, 2,
	2, 2, 226, 222, 3, 2, 2, 2, 227, 60, 3, 2, 2, 2, 228, 229, 7, 111, 2, 2,
	229, 230, 7, 99, 2, 2, 230, 231, 7, 112, 2, 2, 231, 232, 7, 123, 2, 2,
	232, 233, 7, 81, 2, 2, 233, 240, 7, 104, 2, 2, 234, 235, 7, 113, 2, 2,
	235, 236, 7, 112, 2, 2, 236, 237, 7, 103, 2, 2, 237, 238, 7, 81, 2, 2,
	238, 240, 7, 104, 2, 2, 239, 228, 3, 2, 2, 2, 239, 234, 3, 2, 2, 2, 240,
	62, 3, 2, 2, 2, 241, 242, 7, 110, 2, 2, 242, 243, 7, 107, 2, 2, 243, 244,
	7, 117, 2, 2, 244, 245, 7, 118, 2, 2, 245, 64, 3, 2, 2, 2, 246, 250, 7,
	36, 2, 2, 247, 249, 10, 2, 2, 2, 248, 247, 3, 2, 2, 2, 249, 252, 3, 2,
	2, 2, 250, 248, 3, 2, 2, 2, 250, 251, 3, 2, 2, 2, 251, 253, 3, 2, 2, 2,
	252, 250, 3, 2, 2, 2, 253, 254, 7, 36, 2, 2, 254, 66, 3, 2, 2, 2, 255,
	257, 9, 3, 2, 2, 256, 255, 3, 2, 2, 2, 257, 258, 3, 2, 2, 2, 258, 256,
	3, 2, 2, 2, 258, 259, 3, 2, 2, 2, 259, 68, 3, 2, 2, 2, 260, 262, 7, 41,
	2, 2, 261, 263, 9, 4, 2, 2, 262, 261, 3, 2, 2, 2, 263, 264, 3, 2, 2, 2,
	264, 262, 3, 2, 2, 2, 264, 265, 3, 2, 2, 2, 265, 266, 3, 2, 2, 2, 266,
	267, 7, 41, 2, 2, 267, 70, 3, 2, 2, 2, 268, 270, 7, 47, 2, 2, 269, 268,
	3, 2, 2, 2, 269, 270, 3, 2, 2, 2, 270, 272, 3, 2, 2, 2, 271, 273, 9, 5,
	2, 2, 272, 271, 3, 2, 2, 2, 273, 274, 3, 2, 2, 2, 274, 272, 3, 2, 2, 2,
	274, 275, 3, 2, 2, 2, 275, 282, 3, 2, 2, 2, 276, 278, 7, 48, 2, 2, 277,
	279, 9, 5, 2, 2, 278, 277, 3, 2, 2, 2, 279, 280, 3, 2, 2, 2, 280, 278,
	3, 2, 2, 2, 280, 281, 3, 2, 2, 2, 281, 283, 3, 2, 2, 2, 282, 276, 3, 2,
	2, 2, 282, 283, 3, 2, 2, 2, 283, 72, 3, 2, 2, 2, 284, 285, 7, 118, 2, 2,
	285, 286, 7, 116, 2, 2, 286, 287, 7, 119, 2, 2, 287, 288, 7, 103, 2, 2,
	288, 74, 3, 2, 2, 2, 289, 290, 7, 104, 2, 2, 290, 291, 7, 99, 2, 2, 291,
	292, 7, 110, 2, 2, 292, 293, 7, 117, 2, 2, 293, 294, 7, 103, 2, 2, 294,
	76, 3, 2, 2, 2, 295, 297, 9, 6, 2, 2, 296, 295, 3, 2, 2, 2, 297, 298, 3,
	2, 2, 2, 298, 296, 3, 2, 2, 2, 298, 299, 3, 2, 2, 2, 299, 300, 3, 2, 2,
	2, 300, 301, 8, 39, 2, 2, 301, 78, 3, 2, 2, 2, 14, 2, 119, 226, 239, 250,
	258, 264, 269, 274, 280, 282, 298, 3, 8, 2, 2,
}

var lexerDeserializer = antlr.NewATNDeserializer(nil)
var lexerAtn = lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'sashimi:'", "", "'repeat'", "'entity'", "'unique'", "','", "'-'",
	"'of'", "'is'", "'as'", "'begin'", "'end'", "'->'", "'('", "')'", "'['",
	"']'", "'='", "'<='", "'<'", "'>='", "'>'", "'!'", "'<>'", "'&'", "'|'",
	"':'", "'.'", "", "", "'list'", "", "", "", "", "'true'", "'false'",
}

var lexerSymbolicNames = []string{
	"", "DIRECTIVE", "COMMAND", "LOOP", "ENTITY", "UNIQUE", "SEPERATOR", "PROP_START",
	"OF", "IS", "AS", "BEGIN", "END", "ARROW", "LPAREN", "RPAREN", "HLPAREN",
	"HRPAREN", "EQ", "LEQ", "LT", "GEQ", "GT", "NOT", "NEQ", "AND", "OR", "ATOM",
	"DOT", "TYPE", "UNION", "LIST", "ALIAS", "IDENT", "SCOPEIDENT", "NUMBER",
	"TRUE", "FALSE", "WS",
}

var lexerRuleNames = []string{
	"DIRECTIVE", "COMMAND", "LOOP", "ENTITY", "UNIQUE", "SEPERATOR", "PROP_START",
	"OF", "IS", "AS", "BEGIN", "END", "ARROW", "LPAREN", "RPAREN", "HLPAREN",
	"HRPAREN", "EQ", "LEQ", "LT", "GEQ", "GT", "NOT", "NEQ", "AND", "OR", "ATOM",
	"DOT", "TYPE", "UNION", "LIST", "ALIAS", "IDENT", "SCOPEIDENT", "NUMBER",
	"TRUE", "FALSE", "WS",
}

type SashimiLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var lexerDecisionToDFA = make([]*antlr.DFA, len(lexerAtn.DecisionToState))

func init() {
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

func NewSashimiLexer(input antlr.CharStream) *SashimiLexer {

	l := new(SashimiLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "Sashimi.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// SashimiLexer tokens.
const (
	SashimiLexerDIRECTIVE  = 1
	SashimiLexerCOMMAND    = 2
	SashimiLexerLOOP       = 3
	SashimiLexerENTITY     = 4
	SashimiLexerUNIQUE     = 5
	SashimiLexerSEPERATOR  = 6
	SashimiLexerPROP_START = 7
	SashimiLexerOF         = 8
	SashimiLexerIS         = 9
	SashimiLexerAS         = 10
	SashimiLexerBEGIN      = 11
	SashimiLexerEND        = 12
	SashimiLexerARROW      = 13
	SashimiLexerLPAREN     = 14
	SashimiLexerRPAREN     = 15
	SashimiLexerHLPAREN    = 16
	SashimiLexerHRPAREN    = 17
	SashimiLexerEQ         = 18
	SashimiLexerLEQ        = 19
	SashimiLexerLT         = 20
	SashimiLexerGEQ        = 21
	SashimiLexerGT         = 22
	SashimiLexerNOT        = 23
	SashimiLexerNEQ        = 24
	SashimiLexerAND        = 25
	SashimiLexerOR         = 26
	SashimiLexerATOM       = 27
	SashimiLexerDOT        = 28
	SashimiLexerTYPE       = 29
	SashimiLexerUNION      = 30
	SashimiLexerLIST       = 31
	SashimiLexerALIAS      = 32
	SashimiLexerIDENT      = 33
	SashimiLexerSCOPEIDENT = 34
	SashimiLexerNUMBER     = 35
	SashimiLexerTRUE       = 36
	SashimiLexerFALSE      = 37
	SashimiLexerWS         = 38
)
