// Code generated from C:\gitrepos\sashimi/grammar/Sashimi.g4 by ANTLR 4.8. DO NOT EDIT.

package build // Sashimi
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 42, 240,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 3, 2, 3, 2, 3,
	2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 5, 4, 65, 10, 4, 3, 4, 3,
	4, 3, 4, 5, 4, 70, 10, 4, 7, 4, 72, 10, 4, 12, 4, 14, 4, 75, 11, 4, 3,
	4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 7, 5, 84, 10, 5, 12, 5, 14, 5, 87,
	11, 5, 3, 5, 3, 5, 3, 6, 3, 6, 5, 6, 93, 10, 6, 3, 7, 3, 7, 3, 7, 3, 8,
	3, 8, 3, 8, 5, 8, 101, 10, 8, 3, 9, 3, 9, 3, 9, 3, 9, 5, 9, 107, 10, 9,
	3, 10, 3, 10, 3, 10, 3, 11, 3, 11, 3, 11, 3, 12, 3, 12, 3, 12, 5, 12, 118,
	10, 12, 3, 12, 3, 12, 3, 13, 3, 13, 3, 13, 3, 13, 3, 14, 3, 14, 3, 14,
	7, 14, 129, 10, 14, 12, 14, 14, 14, 132, 11, 14, 3, 15, 3, 15, 3, 15, 3,
	15, 5, 15, 138, 10, 15, 3, 15, 3, 15, 3, 15, 5, 15, 143, 10, 15, 3, 15,
	3, 15, 3, 15, 3, 15, 5, 15, 149, 10, 15, 3, 15, 3, 15, 3, 15, 3, 16, 5,
	16, 155, 10, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 7, 16,
	164, 10, 16, 12, 16, 14, 16, 167, 11, 16, 3, 17, 3, 17, 3, 17, 3, 17, 5,
	17, 173, 10, 17, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 5, 18, 180, 10, 18,
	3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 20, 3, 20, 5, 20, 189, 10, 20, 3,
	20, 3, 20, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 5, 21, 198, 10, 21, 3, 22,
	7, 22, 201, 10, 22, 12, 22, 14, 22, 204, 11, 22, 3, 22, 3, 22, 3, 23, 3,
	23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 5, 23,
	219, 10, 23, 3, 23, 3, 23, 3, 23, 3, 23, 7, 23, 225, 10, 23, 12, 23, 14,
	23, 228, 11, 23, 3, 24, 3, 24, 5, 24, 232, 10, 24, 3, 25, 3, 25, 3, 26,
	3, 26, 3, 27, 3, 27, 3, 27, 2, 3, 44, 28, 2, 4, 6, 8, 10, 12, 14, 16, 18,
	20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 2,
	5, 4, 2, 20, 24, 26, 26, 3, 2, 27, 28, 3, 2, 40, 41, 2, 243, 2, 54, 3,
	2, 2, 2, 4, 58, 3, 2, 2, 2, 6, 61, 3, 2, 2, 2, 8, 78, 3, 2, 2, 2, 10, 90,
	3, 2, 2, 2, 12, 94, 3, 2, 2, 2, 14, 97, 3, 2, 2, 2, 16, 106, 3, 2, 2, 2,
	18, 108, 3, 2, 2, 2, 20, 111, 3, 2, 2, 2, 22, 114, 3, 2, 2, 2, 24, 121,
	3, 2, 2, 2, 26, 125, 3, 2, 2, 2, 28, 133, 3, 2, 2, 2, 30, 154, 3, 2, 2,
	2, 32, 168, 3, 2, 2, 2, 34, 174, 3, 2, 2, 2, 36, 181, 3, 2, 2, 2, 38, 186,
	3, 2, 2, 2, 40, 192, 3, 2, 2, 2, 42, 202, 3, 2, 2, 2, 44, 218, 3, 2, 2,
	2, 46, 231, 3, 2, 2, 2, 48, 233, 3, 2, 2, 2, 50, 235, 3, 2, 2, 2, 52, 237,
	3, 2, 2, 2, 54, 55, 7, 37, 2, 2, 55, 56, 7, 20, 2, 2, 56, 57, 7, 37, 2,
	2, 57, 3, 3, 2, 2, 2, 58, 59, 7, 29, 2, 2, 59, 60, 7, 37, 2, 2, 60, 5,
	3, 2, 2, 2, 61, 64, 7, 18, 2, 2, 62, 65, 5, 4, 3, 2, 63, 65, 5, 2, 2, 2,
	64, 62, 3, 2, 2, 2, 64, 63, 3, 2, 2, 2, 65, 73, 3, 2, 2, 2, 66, 69, 7,
	8, 2, 2, 67, 70, 5, 4, 3, 2, 68, 70, 5, 2, 2, 2, 69, 67, 3, 2, 2, 2, 69,
	68, 3, 2, 2, 2, 70, 72, 3, 2, 2, 2, 71, 66, 3, 2, 2, 2, 72, 75, 3, 2, 2,
	2, 73, 71, 3, 2, 2, 2, 73, 74, 3, 2, 2, 2, 74, 76, 3, 2, 2, 2, 75, 73,
	3, 2, 2, 2, 76, 77, 7, 19, 2, 2, 77, 7, 3, 2, 2, 2, 78, 79, 7, 33, 2, 2,
	79, 80, 7, 16, 2, 2, 80, 85, 7, 35, 2, 2, 81, 82, 7, 8, 2, 2, 82, 84, 7,
	35, 2, 2, 83, 81, 3, 2, 2, 2, 84, 87, 3, 2, 2, 2, 85, 83, 3, 2, 2, 2, 85,
	86, 3, 2, 2, 2, 86, 88, 3, 2, 2, 2, 87, 85, 3, 2, 2, 2, 88, 89, 7, 17,
	2, 2, 89, 9, 3, 2, 2, 2, 90, 92, 7, 32, 2, 2, 91, 93, 5, 6, 4, 2, 92, 91,
	3, 2, 2, 2, 92, 93, 3, 2, 2, 2, 93, 11, 3, 2, 2, 2, 94, 95, 7, 6, 2, 2,
	95, 96, 7, 37, 2, 2, 96, 13, 3, 2, 2, 2, 97, 100, 7, 34, 2, 2, 98, 101,
	5, 10, 6, 2, 99, 101, 5, 12, 7, 2, 100, 98, 3, 2, 2, 2, 100, 99, 3, 2,
	2, 2, 101, 15, 3, 2, 2, 2, 102, 107, 5, 14, 8, 2, 103, 107, 5, 8, 5, 2,
	104, 107, 5, 10, 6, 2, 105, 107, 5, 12, 7, 2, 106, 102, 3, 2, 2, 2, 106,
	103, 3, 2, 2, 2, 106, 104, 3, 2, 2, 2, 106, 105, 3, 2, 2, 2, 107, 17, 3,
	2, 2, 2, 108, 109, 7, 12, 2, 2, 109, 110, 7, 35, 2, 2, 110, 19, 3, 2, 2,
	2, 111, 112, 7, 11, 2, 2, 112, 113, 5, 16, 9, 2, 113, 21, 3, 2, 2, 2, 114,
	115, 7, 9, 2, 2, 115, 117, 7, 37, 2, 2, 116, 118, 5, 18, 10, 2, 117, 116,
	3, 2, 2, 2, 117, 118, 3, 2, 2, 2, 118, 119, 3, 2, 2, 2, 119, 120, 5, 20,
	11, 2, 120, 23, 3, 2, 2, 2, 121, 122, 7, 37, 2, 2, 122, 123, 7, 15, 2,
	2, 123, 124, 5, 44, 23, 2, 124, 25, 3, 2, 2, 2, 125, 130, 7, 37, 2, 2,
	126, 127, 7, 30, 2, 2, 127, 129, 7, 37, 2, 2, 128, 126, 3, 2, 2, 2, 129,
	132, 3, 2, 2, 2, 130, 128, 3, 2, 2, 2, 130, 131, 3, 2, 2, 2, 131, 27, 3,
	2, 2, 2, 132, 130, 3, 2, 2, 2, 133, 134, 7, 5, 2, 2, 134, 137, 7, 16, 2,
	2, 135, 138, 7, 36, 2, 2, 136, 138, 5, 26, 14, 2, 137, 135, 3, 2, 2, 2,
	137, 136, 3, 2, 2, 2, 138, 139, 3, 2, 2, 2, 139, 142, 7, 17, 2, 2, 140,
	141, 7, 12, 2, 2, 141, 143, 7, 37, 2, 2, 142, 140, 3, 2, 2, 2, 142, 143,
	3, 2, 2, 2, 143, 148, 3, 2, 2, 2, 144, 145, 7, 18, 2, 2, 145, 146, 5, 24,
	13, 2, 146, 147, 7, 19, 2, 2, 147, 149, 3, 2, 2, 2, 148, 144, 3, 2, 2,
	2, 148, 149, 3, 2, 2, 2, 149, 150, 3, 2, 2, 2, 150, 151, 7, 3, 2, 2, 151,
	152, 5, 38, 20, 2, 152, 29, 3, 2, 2, 2, 153, 155, 7, 7, 2, 2, 154, 153,
	3, 2, 2, 2, 154, 155, 3, 2, 2, 2, 155, 156, 3, 2, 2, 2, 156, 157, 7, 6,
	2, 2, 157, 158, 7, 16, 2, 2, 158, 159, 7, 37, 2, 2, 159, 160, 7, 17, 2,
	2, 160, 161, 7, 10, 2, 2, 161, 165, 5, 22, 12, 2, 162, 164, 5, 22, 12,
	2, 163, 162, 3, 2, 2, 2, 164, 167, 3, 2, 2, 2, 165, 163, 3, 2, 2, 2, 165,
	166, 3, 2, 2, 2, 166, 31, 3, 2, 2, 2, 167, 165, 3, 2, 2, 2, 168, 172, 7,
	13, 2, 2, 169, 170, 7, 16, 2, 2, 170, 171, 7, 38, 2, 2, 171, 173, 7, 17,
	2, 2, 172, 169, 3, 2, 2, 2, 172, 173, 3, 2, 2, 2, 173, 33, 3, 2, 2, 2,
	174, 175, 7, 3, 2, 2, 175, 179, 7, 14, 2, 2, 176, 177, 7, 16, 2, 2, 177,
	178, 7, 38, 2, 2, 178, 180, 7, 17, 2, 2, 179, 176, 3, 2, 2, 2, 179, 180,
	3, 2, 2, 2, 180, 35, 3, 2, 2, 2, 181, 182, 7, 4, 2, 2, 182, 183, 7, 16,
	2, 2, 183, 184, 5, 26, 14, 2, 184, 185, 7, 17, 2, 2, 185, 37, 3, 2, 2,
	2, 186, 188, 5, 32, 17, 2, 187, 189, 5, 40, 21, 2, 188, 187, 3, 2, 2, 2,
	188, 189, 3, 2, 2, 2, 189, 190, 3, 2, 2, 2, 190, 191, 5, 34, 18, 2, 191,
	39, 3, 2, 2, 2, 192, 197, 7, 3, 2, 2, 193, 198, 5, 36, 19, 2, 194, 198,
	5, 28, 15, 2, 195, 198, 5, 30, 16, 2, 196, 198, 5, 38, 20, 2, 197, 193,
	3, 2, 2, 2, 197, 194, 3, 2, 2, 2, 197, 195, 3, 2, 2, 2, 197, 196, 3, 2,
	2, 2, 198, 41, 3, 2, 2, 2, 199, 201, 5, 40, 21, 2, 200, 199, 3, 2, 2, 2,
	201, 204, 3, 2, 2, 2, 202, 200, 3, 2, 2, 2, 202, 203, 3, 2, 2, 2, 203,
	205, 3, 2, 2, 2, 204, 202, 3, 2, 2, 2, 205, 206, 7, 2, 2, 3, 206, 43, 3,
	2, 2, 2, 207, 208, 8, 23, 1, 2, 208, 209, 7, 16, 2, 2, 209, 210, 5, 44,
	23, 2, 210, 211, 7, 17, 2, 2, 211, 219, 3, 2, 2, 2, 212, 213, 7, 25, 2,
	2, 213, 219, 5, 44, 23, 8, 214, 219, 5, 52, 27, 2, 215, 219, 5, 26, 14,
	2, 216, 219, 7, 35, 2, 2, 217, 219, 7, 39, 2, 2, 218, 207, 3, 2, 2, 2,
	218, 212, 3, 2, 2, 2, 218, 214, 3, 2, 2, 2, 218, 215, 3, 2, 2, 2, 218,
	216, 3, 2, 2, 2, 218, 217, 3, 2, 2, 2, 219, 226, 3, 2, 2, 2, 220, 221,
	12, 7, 2, 2, 221, 222, 5, 46, 24, 2, 222, 223, 5, 44, 23, 8, 223, 225,
	3, 2, 2, 2, 224, 220, 3, 2, 2, 2, 225, 228, 3, 2, 2, 2, 226, 224, 3, 2,
	2, 2, 226, 227, 3, 2, 2, 2, 227, 45, 3, 2, 2, 2, 228, 226, 3, 2, 2, 2,
	229, 232, 5, 50, 26, 2, 230, 232, 5, 48, 25, 2, 231, 229, 3, 2, 2, 2, 231,
	230, 3, 2, 2, 2, 232, 47, 3, 2, 2, 2, 233, 234, 9, 2, 2, 2, 234, 49, 3,
	2, 2, 2, 235, 236, 9, 3, 2, 2, 236, 51, 3, 2, 2, 2, 237, 238, 9, 4, 2,
	2, 238, 53, 3, 2, 2, 2, 24, 64, 69, 73, 85, 92, 100, 106, 117, 130, 137,
	142, 148, 154, 165, 172, 179, 188, 197, 202, 218, 226, 231,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'sashimi:'", "", "'repeat'", "'entity'", "'unique'", "','", "'-'",
	"'of'", "'is'", "'as'", "'begin'", "'end'", "'->'", "'('", "')'", "'['",
	"']'", "'='", "'<='", "'<'", "'>='", "'>'", "'!'", "'<>'", "'&'", "'|'",
	"':'", "'.'", "'@'", "", "", "'list'", "", "", "", "", "", "'true'", "'false'",
}
var symbolicNames = []string{
	"", "DIRECTIVE", "COMMAND", "LOOP", "ENTITY", "UNIQUE", "SEPERATOR", "PROP_START",
	"OF", "IS", "AS", "BEGIN", "END", "ARROW", "LPAREN", "RPAREN", "HLPAREN",
	"HRPAREN", "EQ", "LEQ", "LT", "GEQ", "GT", "NOT", "NEQ", "AND", "OR", "ATOM",
	"DOT", "AT", "TYPE", "UNION", "LIST", "ALIAS", "GLOBAL", "IDENT", "SCOPEIDENT",
	"NUMBER", "TRUE", "FALSE", "WS",
}

var ruleNames = []string{
	"keyValuePair", "keyAtom", "constraintList", "unionDecl", "typeDecl", "entityRef",
	"listDecl", "typeDef", "aliasDecl", "typeIs", "propDecl", "predicate",
	"qualifier", "loopCall", "entityDef", "scopeBegin", "scopeEnd", "commandCall",
	"blockScope", "export", "block", "boolExpression", "op", "comparator",
	"binary", "truth",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type SashimiParser struct {
	*antlr.BaseParser
}

func NewSashimiParser(input antlr.TokenStream) *SashimiParser {
	this := new(SashimiParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Sashimi.g4"

	return this
}

// SashimiParser tokens.
const (
	SashimiParserEOF        = antlr.TokenEOF
	SashimiParserDIRECTIVE  = 1
	SashimiParserCOMMAND    = 2
	SashimiParserLOOP       = 3
	SashimiParserENTITY     = 4
	SashimiParserUNIQUE     = 5
	SashimiParserSEPERATOR  = 6
	SashimiParserPROP_START = 7
	SashimiParserOF         = 8
	SashimiParserIS         = 9
	SashimiParserAS         = 10
	SashimiParserBEGIN      = 11
	SashimiParserEND        = 12
	SashimiParserARROW      = 13
	SashimiParserLPAREN     = 14
	SashimiParserRPAREN     = 15
	SashimiParserHLPAREN    = 16
	SashimiParserHRPAREN    = 17
	SashimiParserEQ         = 18
	SashimiParserLEQ        = 19
	SashimiParserLT         = 20
	SashimiParserGEQ        = 21
	SashimiParserGT         = 22
	SashimiParserNOT        = 23
	SashimiParserNEQ        = 24
	SashimiParserAND        = 25
	SashimiParserOR         = 26
	SashimiParserATOM       = 27
	SashimiParserDOT        = 28
	SashimiParserAT         = 29
	SashimiParserTYPE       = 30
	SashimiParserUNION      = 31
	SashimiParserLIST       = 32
	SashimiParserALIAS      = 33
	SashimiParserGLOBAL     = 34
	SashimiParserIDENT      = 35
	SashimiParserSCOPEIDENT = 36
	SashimiParserNUMBER     = 37
	SashimiParserTRUE       = 38
	SashimiParserFALSE      = 39
	SashimiParserWS         = 40
)

// SashimiParser rules.
const (
	SashimiParserRULE_keyValuePair   = 0
	SashimiParserRULE_keyAtom        = 1
	SashimiParserRULE_constraintList = 2
	SashimiParserRULE_unionDecl      = 3
	SashimiParserRULE_typeDecl       = 4
	SashimiParserRULE_entityRef      = 5
	SashimiParserRULE_listDecl       = 6
	SashimiParserRULE_typeDef        = 7
	SashimiParserRULE_aliasDecl      = 8
	SashimiParserRULE_typeIs         = 9
	SashimiParserRULE_propDecl       = 10
	SashimiParserRULE_predicate      = 11
	SashimiParserRULE_qualifier      = 12
	SashimiParserRULE_loopCall       = 13
	SashimiParserRULE_entityDef      = 14
	SashimiParserRULE_scopeBegin     = 15
	SashimiParserRULE_scopeEnd       = 16
	SashimiParserRULE_commandCall    = 17
	SashimiParserRULE_blockScope     = 18
	SashimiParserRULE_export         = 19
	SashimiParserRULE_block          = 20
	SashimiParserRULE_boolExpression = 21
	SashimiParserRULE_op             = 22
	SashimiParserRULE_comparator     = 23
	SashimiParserRULE_binary         = 24
	SashimiParserRULE_truth          = 25
)

// IKeyValuePairContext is an interface to support dynamic dispatch.
type IKeyValuePairContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsKeyValuePairContext differentiates from other interfaces.
	IsKeyValuePairContext()
}

type KeyValuePairContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyKeyValuePairContext() *KeyValuePairContext {
	var p = new(KeyValuePairContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SashimiParserRULE_keyValuePair
	return p
}

func (*KeyValuePairContext) IsKeyValuePairContext() {}

func NewKeyValuePairContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *KeyValuePairContext {
	var p = new(KeyValuePairContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SashimiParserRULE_keyValuePair

	return p
}

func (s *KeyValuePairContext) GetParser() antlr.Parser { return s.parser }

func (s *KeyValuePairContext) AllIDENT() []antlr.TerminalNode {
	return s.GetTokens(SashimiParserIDENT)
}

func (s *KeyValuePairContext) IDENT(i int) antlr.TerminalNode {
	return s.GetToken(SashimiParserIDENT, i)
}

func (s *KeyValuePairContext) EQ() antlr.TerminalNode {
	return s.GetToken(SashimiParserEQ, 0)
}

func (s *KeyValuePairContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *KeyValuePairContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *KeyValuePairContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SashimiListener); ok {
		listenerT.EnterKeyValuePair(s)
	}
}

func (s *KeyValuePairContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SashimiListener); ok {
		listenerT.ExitKeyValuePair(s)
	}
}

func (p *SashimiParser) KeyValuePair() (localctx IKeyValuePairContext) {
	localctx = NewKeyValuePairContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, SashimiParserRULE_keyValuePair)

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
		p.SetState(52)
		p.Match(SashimiParserIDENT)
	}
	{
		p.SetState(53)
		p.Match(SashimiParserEQ)
	}
	{
		p.SetState(54)
		p.Match(SashimiParserIDENT)
	}

	return localctx
}

// IKeyAtomContext is an interface to support dynamic dispatch.
type IKeyAtomContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsKeyAtomContext differentiates from other interfaces.
	IsKeyAtomContext()
}

type KeyAtomContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyKeyAtomContext() *KeyAtomContext {
	var p = new(KeyAtomContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SashimiParserRULE_keyAtom
	return p
}

func (*KeyAtomContext) IsKeyAtomContext() {}

func NewKeyAtomContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *KeyAtomContext {
	var p = new(KeyAtomContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SashimiParserRULE_keyAtom

	return p
}

func (s *KeyAtomContext) GetParser() antlr.Parser { return s.parser }

func (s *KeyAtomContext) ATOM() antlr.TerminalNode {
	return s.GetToken(SashimiParserATOM, 0)
}

func (s *KeyAtomContext) IDENT() antlr.TerminalNode {
	return s.GetToken(SashimiParserIDENT, 0)
}

func (s *KeyAtomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *KeyAtomContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *KeyAtomContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SashimiListener); ok {
		listenerT.EnterKeyAtom(s)
	}
}

func (s *KeyAtomContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SashimiListener); ok {
		listenerT.ExitKeyAtom(s)
	}
}

func (p *SashimiParser) KeyAtom() (localctx IKeyAtomContext) {
	localctx = NewKeyAtomContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, SashimiParserRULE_keyAtom)

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
		p.SetState(56)
		p.Match(SashimiParserATOM)
	}
	{
		p.SetState(57)
		p.Match(SashimiParserIDENT)
	}

	return localctx
}

// IConstraintListContext is an interface to support dynamic dispatch.
type IConstraintListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConstraintListContext differentiates from other interfaces.
	IsConstraintListContext()
}

type ConstraintListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConstraintListContext() *ConstraintListContext {
	var p = new(ConstraintListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SashimiParserRULE_constraintList
	return p
}

func (*ConstraintListContext) IsConstraintListContext() {}

func NewConstraintListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConstraintListContext {
	var p = new(ConstraintListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SashimiParserRULE_constraintList

	return p
}

func (s *ConstraintListContext) GetParser() antlr.Parser { return s.parser }

func (s *ConstraintListContext) HLPAREN() antlr.TerminalNode {
	return s.GetToken(SashimiParserHLPAREN, 0)
}

func (s *ConstraintListContext) HRPAREN() antlr.TerminalNode {
	return s.GetToken(SashimiParserHRPAREN, 0)
}

func (s *ConstraintListContext) AllKeyAtom() []IKeyAtomContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IKeyAtomContext)(nil)).Elem())
	var tst = make([]IKeyAtomContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IKeyAtomContext)
		}
	}

	return tst
}

func (s *ConstraintListContext) KeyAtom(i int) IKeyAtomContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IKeyAtomContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IKeyAtomContext)
}

func (s *ConstraintListContext) AllKeyValuePair() []IKeyValuePairContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IKeyValuePairContext)(nil)).Elem())
	var tst = make([]IKeyValuePairContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IKeyValuePairContext)
		}
	}

	return tst
}

func (s *ConstraintListContext) KeyValuePair(i int) IKeyValuePairContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IKeyValuePairContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IKeyValuePairContext)
}

func (s *ConstraintListContext) AllSEPERATOR() []antlr.TerminalNode {
	return s.GetTokens(SashimiParserSEPERATOR)
}

func (s *ConstraintListContext) SEPERATOR(i int) antlr.TerminalNode {
	return s.GetToken(SashimiParserSEPERATOR, i)
}

func (s *ConstraintListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstraintListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConstraintListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SashimiListener); ok {
		listenerT.EnterConstraintList(s)
	}
}

func (s *ConstraintListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SashimiListener); ok {
		listenerT.ExitConstraintList(s)
	}
}

func (p *SashimiParser) ConstraintList() (localctx IConstraintListContext) {
	localctx = NewConstraintListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, SashimiParserRULE_constraintList)
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
		p.SetState(59)
		p.Match(SashimiParserHLPAREN)
	}
	p.SetState(62)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SashimiParserATOM:
		{
			p.SetState(60)
			p.KeyAtom()
		}

	case SashimiParserIDENT:
		{
			p.SetState(61)
			p.KeyValuePair()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.SetState(71)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == SashimiParserSEPERATOR {
		{
			p.SetState(64)
			p.Match(SashimiParserSEPERATOR)
		}
		p.SetState(67)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case SashimiParserATOM:
			{
				p.SetState(65)
				p.KeyAtom()
			}

		case SashimiParserIDENT:
			{
				p.SetState(66)
				p.KeyValuePair()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(73)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(74)
		p.Match(SashimiParserHRPAREN)
	}

	return localctx
}

// IUnionDeclContext is an interface to support dynamic dispatch.
type IUnionDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUnionDeclContext differentiates from other interfaces.
	IsUnionDeclContext()
}

type UnionDeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnionDeclContext() *UnionDeclContext {
	var p = new(UnionDeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SashimiParserRULE_unionDecl
	return p
}

func (*UnionDeclContext) IsUnionDeclContext() {}

func NewUnionDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnionDeclContext {
	var p = new(UnionDeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SashimiParserRULE_unionDecl

	return p
}

func (s *UnionDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *UnionDeclContext) UNION() antlr.TerminalNode {
	return s.GetToken(SashimiParserUNION, 0)
}

func (s *UnionDeclContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(SashimiParserLPAREN, 0)
}

func (s *UnionDeclContext) AllALIAS() []antlr.TerminalNode {
	return s.GetTokens(SashimiParserALIAS)
}

func (s *UnionDeclContext) ALIAS(i int) antlr.TerminalNode {
	return s.GetToken(SashimiParserALIAS, i)
}

func (s *UnionDeclContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(SashimiParserRPAREN, 0)
}

func (s *UnionDeclContext) AllSEPERATOR() []antlr.TerminalNode {
	return s.GetTokens(SashimiParserSEPERATOR)
}

func (s *UnionDeclContext) SEPERATOR(i int) antlr.TerminalNode {
	return s.GetToken(SashimiParserSEPERATOR, i)
}

func (s *UnionDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnionDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UnionDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SashimiListener); ok {
		listenerT.EnterUnionDecl(s)
	}
}

func (s *UnionDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SashimiListener); ok {
		listenerT.ExitUnionDecl(s)
	}
}

func (p *SashimiParser) UnionDecl() (localctx IUnionDeclContext) {
	localctx = NewUnionDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, SashimiParserRULE_unionDecl)
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
		p.SetState(76)
		p.Match(SashimiParserUNION)
	}
	{
		p.SetState(77)
		p.Match(SashimiParserLPAREN)
	}
	{
		p.SetState(78)
		p.Match(SashimiParserALIAS)
	}
	p.SetState(83)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == SashimiParserSEPERATOR {
		{
			p.SetState(79)
			p.Match(SashimiParserSEPERATOR)
		}
		{
			p.SetState(80)
			p.Match(SashimiParserALIAS)
		}

		p.SetState(85)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(86)
		p.Match(SashimiParserRPAREN)
	}

	return localctx
}

// ITypeDeclContext is an interface to support dynamic dispatch.
type ITypeDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeDeclContext differentiates from other interfaces.
	IsTypeDeclContext()
}

type TypeDeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeDeclContext() *TypeDeclContext {
	var p = new(TypeDeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SashimiParserRULE_typeDecl
	return p
}

func (*TypeDeclContext) IsTypeDeclContext() {}

func NewTypeDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeDeclContext {
	var p = new(TypeDeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SashimiParserRULE_typeDecl

	return p
}

func (s *TypeDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeDeclContext) TYPE() antlr.TerminalNode {
	return s.GetToken(SashimiParserTYPE, 0)
}

func (s *TypeDeclContext) ConstraintList() IConstraintListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConstraintListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConstraintListContext)
}

func (s *TypeDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SashimiListener); ok {
		listenerT.EnterTypeDecl(s)
	}
}

func (s *TypeDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SashimiListener); ok {
		listenerT.ExitTypeDecl(s)
	}
}

func (p *SashimiParser) TypeDecl() (localctx ITypeDeclContext) {
	localctx = NewTypeDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, SashimiParserRULE_typeDecl)
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
		p.SetState(88)
		p.Match(SashimiParserTYPE)
	}
	p.SetState(90)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SashimiParserHLPAREN {
		{
			p.SetState(89)
			p.ConstraintList()
		}

	}

	return localctx
}

// IEntityRefContext is an interface to support dynamic dispatch.
type IEntityRefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEntityRefContext differentiates from other interfaces.
	IsEntityRefContext()
}

type EntityRefContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEntityRefContext() *EntityRefContext {
	var p = new(EntityRefContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SashimiParserRULE_entityRef
	return p
}

func (*EntityRefContext) IsEntityRefContext() {}

func NewEntityRefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EntityRefContext {
	var p = new(EntityRefContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SashimiParserRULE_entityRef

	return p
}

func (s *EntityRefContext) GetParser() antlr.Parser { return s.parser }

func (s *EntityRefContext) ENTITY() antlr.TerminalNode {
	return s.GetToken(SashimiParserENTITY, 0)
}

func (s *EntityRefContext) IDENT() antlr.TerminalNode {
	return s.GetToken(SashimiParserIDENT, 0)
}

func (s *EntityRefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EntityRefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EntityRefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SashimiListener); ok {
		listenerT.EnterEntityRef(s)
	}
}

func (s *EntityRefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SashimiListener); ok {
		listenerT.ExitEntityRef(s)
	}
}

func (p *SashimiParser) EntityRef() (localctx IEntityRefContext) {
	localctx = NewEntityRefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, SashimiParserRULE_entityRef)

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
		p.SetState(92)
		p.Match(SashimiParserENTITY)
	}
	{
		p.SetState(93)
		p.Match(SashimiParserIDENT)
	}

	return localctx
}

// IListDeclContext is an interface to support dynamic dispatch.
type IListDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsListDeclContext differentiates from other interfaces.
	IsListDeclContext()
}

type ListDeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyListDeclContext() *ListDeclContext {
	var p = new(ListDeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SashimiParserRULE_listDecl
	return p
}

func (*ListDeclContext) IsListDeclContext() {}

func NewListDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListDeclContext {
	var p = new(ListDeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SashimiParserRULE_listDecl

	return p
}

func (s *ListDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *ListDeclContext) LIST() antlr.TerminalNode {
	return s.GetToken(SashimiParserLIST, 0)
}

func (s *ListDeclContext) TypeDecl() ITypeDeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeDeclContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeDeclContext)
}

func (s *ListDeclContext) EntityRef() IEntityRefContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEntityRefContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEntityRefContext)
}

func (s *ListDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SashimiListener); ok {
		listenerT.EnterListDecl(s)
	}
}

func (s *ListDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SashimiListener); ok {
		listenerT.ExitListDecl(s)
	}
}

func (p *SashimiParser) ListDecl() (localctx IListDeclContext) {
	localctx = NewListDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, SashimiParserRULE_listDecl)

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
		p.SetState(95)
		p.Match(SashimiParserLIST)
	}
	p.SetState(98)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SashimiParserTYPE:
		{
			p.SetState(96)
			p.TypeDecl()
		}

	case SashimiParserENTITY:
		{
			p.SetState(97)
			p.EntityRef()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ITypeDefContext is an interface to support dynamic dispatch.
type ITypeDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeDefContext differentiates from other interfaces.
	IsTypeDefContext()
}

type TypeDefContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeDefContext() *TypeDefContext {
	var p = new(TypeDefContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SashimiParserRULE_typeDef
	return p
}

func (*TypeDefContext) IsTypeDefContext() {}

func NewTypeDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeDefContext {
	var p = new(TypeDefContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SashimiParserRULE_typeDef

	return p
}

func (s *TypeDefContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeDefContext) ListDecl() IListDeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IListDeclContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IListDeclContext)
}

func (s *TypeDefContext) UnionDecl() IUnionDeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUnionDeclContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUnionDeclContext)
}

func (s *TypeDefContext) TypeDecl() ITypeDeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeDeclContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeDeclContext)
}

func (s *TypeDefContext) EntityRef() IEntityRefContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEntityRefContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEntityRefContext)
}

func (s *TypeDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SashimiListener); ok {
		listenerT.EnterTypeDef(s)
	}
}

func (s *TypeDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SashimiListener); ok {
		listenerT.ExitTypeDef(s)
	}
}

func (p *SashimiParser) TypeDef() (localctx ITypeDefContext) {
	localctx = NewTypeDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, SashimiParserRULE_typeDef)

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

	p.SetState(104)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SashimiParserLIST:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(100)
			p.ListDecl()
		}

	case SashimiParserUNION:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(101)
			p.UnionDecl()
		}

	case SashimiParserTYPE:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(102)
			p.TypeDecl()
		}

	case SashimiParserENTITY:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(103)
			p.EntityRef()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IAliasDeclContext is an interface to support dynamic dispatch.
type IAliasDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAliasDeclContext differentiates from other interfaces.
	IsAliasDeclContext()
}

type AliasDeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAliasDeclContext() *AliasDeclContext {
	var p = new(AliasDeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SashimiParserRULE_aliasDecl
	return p
}

func (*AliasDeclContext) IsAliasDeclContext() {}

func NewAliasDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AliasDeclContext {
	var p = new(AliasDeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SashimiParserRULE_aliasDecl

	return p
}

func (s *AliasDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *AliasDeclContext) AS() antlr.TerminalNode {
	return s.GetToken(SashimiParserAS, 0)
}

func (s *AliasDeclContext) ALIAS() antlr.TerminalNode {
	return s.GetToken(SashimiParserALIAS, 0)
}

func (s *AliasDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AliasDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AliasDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SashimiListener); ok {
		listenerT.EnterAliasDecl(s)
	}
}

func (s *AliasDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SashimiListener); ok {
		listenerT.ExitAliasDecl(s)
	}
}

func (p *SashimiParser) AliasDecl() (localctx IAliasDeclContext) {
	localctx = NewAliasDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, SashimiParserRULE_aliasDecl)

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
		p.SetState(106)
		p.Match(SashimiParserAS)
	}
	{
		p.SetState(107)
		p.Match(SashimiParserALIAS)
	}

	return localctx
}

// ITypeIsContext is an interface to support dynamic dispatch.
type ITypeIsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeIsContext differentiates from other interfaces.
	IsTypeIsContext()
}

type TypeIsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeIsContext() *TypeIsContext {
	var p = new(TypeIsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SashimiParserRULE_typeIs
	return p
}

func (*TypeIsContext) IsTypeIsContext() {}

func NewTypeIsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeIsContext {
	var p = new(TypeIsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SashimiParserRULE_typeIs

	return p
}

func (s *TypeIsContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeIsContext) IS() antlr.TerminalNode {
	return s.GetToken(SashimiParserIS, 0)
}

func (s *TypeIsContext) TypeDef() ITypeDefContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeDefContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeDefContext)
}

func (s *TypeIsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeIsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeIsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SashimiListener); ok {
		listenerT.EnterTypeIs(s)
	}
}

func (s *TypeIsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SashimiListener); ok {
		listenerT.ExitTypeIs(s)
	}
}

func (p *SashimiParser) TypeIs() (localctx ITypeIsContext) {
	localctx = NewTypeIsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, SashimiParserRULE_typeIs)

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
		p.SetState(109)
		p.Match(SashimiParserIS)
	}
	{
		p.SetState(110)
		p.TypeDef()
	}

	return localctx
}

// IPropDeclContext is an interface to support dynamic dispatch.
type IPropDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPropDeclContext differentiates from other interfaces.
	IsPropDeclContext()
}

type PropDeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPropDeclContext() *PropDeclContext {
	var p = new(PropDeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SashimiParserRULE_propDecl
	return p
}

func (*PropDeclContext) IsPropDeclContext() {}

func NewPropDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PropDeclContext {
	var p = new(PropDeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SashimiParserRULE_propDecl

	return p
}

func (s *PropDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *PropDeclContext) PROP_START() antlr.TerminalNode {
	return s.GetToken(SashimiParserPROP_START, 0)
}

func (s *PropDeclContext) IDENT() antlr.TerminalNode {
	return s.GetToken(SashimiParserIDENT, 0)
}

func (s *PropDeclContext) TypeIs() ITypeIsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeIsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeIsContext)
}

func (s *PropDeclContext) AliasDecl() IAliasDeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAliasDeclContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAliasDeclContext)
}

func (s *PropDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PropDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PropDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SashimiListener); ok {
		listenerT.EnterPropDecl(s)
	}
}

func (s *PropDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SashimiListener); ok {
		listenerT.ExitPropDecl(s)
	}
}

func (p *SashimiParser) PropDecl() (localctx IPropDeclContext) {
	localctx = NewPropDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, SashimiParserRULE_propDecl)
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
		p.Match(SashimiParserPROP_START)
	}
	{
		p.SetState(113)
		p.Match(SashimiParserIDENT)
	}
	p.SetState(115)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SashimiParserAS {
		{
			p.SetState(114)
			p.AliasDecl()
		}

	}
	{
		p.SetState(117)
		p.TypeIs()
	}

	return localctx
}

// IPredicateContext is an interface to support dynamic dispatch.
type IPredicateContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPredicateContext differentiates from other interfaces.
	IsPredicateContext()
}

type PredicateContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPredicateContext() *PredicateContext {
	var p = new(PredicateContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SashimiParserRULE_predicate
	return p
}

func (*PredicateContext) IsPredicateContext() {}

func NewPredicateContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PredicateContext {
	var p = new(PredicateContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SashimiParserRULE_predicate

	return p
}

func (s *PredicateContext) GetParser() antlr.Parser { return s.parser }

func (s *PredicateContext) IDENT() antlr.TerminalNode {
	return s.GetToken(SashimiParserIDENT, 0)
}

func (s *PredicateContext) ARROW() antlr.TerminalNode {
	return s.GetToken(SashimiParserARROW, 0)
}

func (s *PredicateContext) BoolExpression() IBoolExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBoolExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBoolExpressionContext)
}

func (s *PredicateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PredicateContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PredicateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SashimiListener); ok {
		listenerT.EnterPredicate(s)
	}
}

func (s *PredicateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SashimiListener); ok {
		listenerT.ExitPredicate(s)
	}
}

func (p *SashimiParser) Predicate() (localctx IPredicateContext) {
	localctx = NewPredicateContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, SashimiParserRULE_predicate)

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
		p.SetState(119)
		p.Match(SashimiParserIDENT)
	}
	{
		p.SetState(120)
		p.Match(SashimiParserARROW)
	}
	{
		p.SetState(121)
		p.boolExpression(0)
	}

	return localctx
}

// IQualifierContext is an interface to support dynamic dispatch.
type IQualifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsQualifierContext differentiates from other interfaces.
	IsQualifierContext()
}

type QualifierContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyQualifierContext() *QualifierContext {
	var p = new(QualifierContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SashimiParserRULE_qualifier
	return p
}

func (*QualifierContext) IsQualifierContext() {}

func NewQualifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *QualifierContext {
	var p = new(QualifierContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SashimiParserRULE_qualifier

	return p
}

func (s *QualifierContext) GetParser() antlr.Parser { return s.parser }

func (s *QualifierContext) AllIDENT() []antlr.TerminalNode {
	return s.GetTokens(SashimiParserIDENT)
}

func (s *QualifierContext) IDENT(i int) antlr.TerminalNode {
	return s.GetToken(SashimiParserIDENT, i)
}

func (s *QualifierContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(SashimiParserDOT)
}

func (s *QualifierContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(SashimiParserDOT, i)
}

func (s *QualifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QualifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *QualifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SashimiListener); ok {
		listenerT.EnterQualifier(s)
	}
}

func (s *QualifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SashimiListener); ok {
		listenerT.ExitQualifier(s)
	}
}

func (p *SashimiParser) Qualifier() (localctx IQualifierContext) {
	localctx = NewQualifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, SashimiParserRULE_qualifier)

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
		p.SetState(123)
		p.Match(SashimiParserIDENT)
	}
	p.SetState(128)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(124)
				p.Match(SashimiParserDOT)
			}
			{
				p.SetState(125)
				p.Match(SashimiParserIDENT)
			}

		}
		p.SetState(130)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext())
	}

	return localctx
}

// ILoopCallContext is an interface to support dynamic dispatch.
type ILoopCallContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetAlias returns the alias token.
	GetAlias() antlr.Token

	// SetAlias sets the alias token.
	SetAlias(antlr.Token)

	// IsLoopCallContext differentiates from other interfaces.
	IsLoopCallContext()
}

type LoopCallContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	alias  antlr.Token
}

func NewEmptyLoopCallContext() *LoopCallContext {
	var p = new(LoopCallContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SashimiParserRULE_loopCall
	return p
}

func (*LoopCallContext) IsLoopCallContext() {}

func NewLoopCallContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LoopCallContext {
	var p = new(LoopCallContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SashimiParserRULE_loopCall

	return p
}

func (s *LoopCallContext) GetParser() antlr.Parser { return s.parser }

func (s *LoopCallContext) GetAlias() antlr.Token { return s.alias }

func (s *LoopCallContext) SetAlias(v antlr.Token) { s.alias = v }

func (s *LoopCallContext) LOOP() antlr.TerminalNode {
	return s.GetToken(SashimiParserLOOP, 0)
}

func (s *LoopCallContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(SashimiParserLPAREN, 0)
}

func (s *LoopCallContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(SashimiParserRPAREN, 0)
}

func (s *LoopCallContext) DIRECTIVE() antlr.TerminalNode {
	return s.GetToken(SashimiParserDIRECTIVE, 0)
}

func (s *LoopCallContext) BlockScope() IBlockScopeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockScopeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockScopeContext)
}

func (s *LoopCallContext) GLOBAL() antlr.TerminalNode {
	return s.GetToken(SashimiParserGLOBAL, 0)
}

func (s *LoopCallContext) Qualifier() IQualifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IQualifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IQualifierContext)
}

func (s *LoopCallContext) AS() antlr.TerminalNode {
	return s.GetToken(SashimiParserAS, 0)
}

func (s *LoopCallContext) HLPAREN() antlr.TerminalNode {
	return s.GetToken(SashimiParserHLPAREN, 0)
}

func (s *LoopCallContext) Predicate() IPredicateContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPredicateContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPredicateContext)
}

func (s *LoopCallContext) HRPAREN() antlr.TerminalNode {
	return s.GetToken(SashimiParserHRPAREN, 0)
}

func (s *LoopCallContext) IDENT() antlr.TerminalNode {
	return s.GetToken(SashimiParserIDENT, 0)
}

func (s *LoopCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LoopCallContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LoopCallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SashimiListener); ok {
		listenerT.EnterLoopCall(s)
	}
}

func (s *LoopCallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SashimiListener); ok {
		listenerT.ExitLoopCall(s)
	}
}

func (p *SashimiParser) LoopCall() (localctx ILoopCallContext) {
	localctx = NewLoopCallContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, SashimiParserRULE_loopCall)
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
		p.SetState(131)
		p.Match(SashimiParserLOOP)
	}
	{
		p.SetState(132)
		p.Match(SashimiParserLPAREN)
	}
	p.SetState(135)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SashimiParserGLOBAL:
		{
			p.SetState(133)
			p.Match(SashimiParserGLOBAL)
		}

	case SashimiParserIDENT:
		{
			p.SetState(134)
			p.Qualifier()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	{
		p.SetState(137)
		p.Match(SashimiParserRPAREN)
	}
	p.SetState(140)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SashimiParserAS {
		{
			p.SetState(138)
			p.Match(SashimiParserAS)
		}
		{
			p.SetState(139)

			var _m = p.Match(SashimiParserIDENT)

			localctx.(*LoopCallContext).alias = _m
		}

	}
	p.SetState(146)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SashimiParserHLPAREN {
		{
			p.SetState(142)
			p.Match(SashimiParserHLPAREN)
		}
		{
			p.SetState(143)
			p.Predicate()
		}
		{
			p.SetState(144)
			p.Match(SashimiParserHRPAREN)
		}

	}
	{
		p.SetState(148)
		p.Match(SashimiParserDIRECTIVE)
	}
	{
		p.SetState(149)
		p.BlockScope()
	}

	return localctx
}

// IEntityDefContext is an interface to support dynamic dispatch.
type IEntityDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEntityDefContext differentiates from other interfaces.
	IsEntityDefContext()
}

type EntityDefContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEntityDefContext() *EntityDefContext {
	var p = new(EntityDefContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SashimiParserRULE_entityDef
	return p
}

func (*EntityDefContext) IsEntityDefContext() {}

func NewEntityDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EntityDefContext {
	var p = new(EntityDefContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SashimiParserRULE_entityDef

	return p
}

func (s *EntityDefContext) GetParser() antlr.Parser { return s.parser }

func (s *EntityDefContext) ENTITY() antlr.TerminalNode {
	return s.GetToken(SashimiParserENTITY, 0)
}

func (s *EntityDefContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(SashimiParserLPAREN, 0)
}

func (s *EntityDefContext) IDENT() antlr.TerminalNode {
	return s.GetToken(SashimiParserIDENT, 0)
}

func (s *EntityDefContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(SashimiParserRPAREN, 0)
}

func (s *EntityDefContext) OF() antlr.TerminalNode {
	return s.GetToken(SashimiParserOF, 0)
}

func (s *EntityDefContext) AllPropDecl() []IPropDeclContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IPropDeclContext)(nil)).Elem())
	var tst = make([]IPropDeclContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IPropDeclContext)
		}
	}

	return tst
}

func (s *EntityDefContext) PropDecl(i int) IPropDeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPropDeclContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IPropDeclContext)
}

func (s *EntityDefContext) UNIQUE() antlr.TerminalNode {
	return s.GetToken(SashimiParserUNIQUE, 0)
}

func (s *EntityDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EntityDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EntityDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SashimiListener); ok {
		listenerT.EnterEntityDef(s)
	}
}

func (s *EntityDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SashimiListener); ok {
		listenerT.ExitEntityDef(s)
	}
}

func (p *SashimiParser) EntityDef() (localctx IEntityDefContext) {
	localctx = NewEntityDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, SashimiParserRULE_entityDef)
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
	p.SetState(152)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SashimiParserUNIQUE {
		{
			p.SetState(151)
			p.Match(SashimiParserUNIQUE)
		}

	}
	{
		p.SetState(154)
		p.Match(SashimiParserENTITY)
	}
	{
		p.SetState(155)
		p.Match(SashimiParserLPAREN)
	}
	{
		p.SetState(156)
		p.Match(SashimiParserIDENT)
	}
	{
		p.SetState(157)
		p.Match(SashimiParserRPAREN)
	}
	{
		p.SetState(158)
		p.Match(SashimiParserOF)
	}
	{
		p.SetState(159)
		p.PropDecl()
	}
	p.SetState(163)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == SashimiParserPROP_START {
		{
			p.SetState(160)
			p.PropDecl()
		}

		p.SetState(165)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IScopeBeginContext is an interface to support dynamic dispatch.
type IScopeBeginContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsScopeBeginContext differentiates from other interfaces.
	IsScopeBeginContext()
}

type ScopeBeginContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyScopeBeginContext() *ScopeBeginContext {
	var p = new(ScopeBeginContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SashimiParserRULE_scopeBegin
	return p
}

func (*ScopeBeginContext) IsScopeBeginContext() {}

func NewScopeBeginContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ScopeBeginContext {
	var p = new(ScopeBeginContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SashimiParserRULE_scopeBegin

	return p
}

func (s *ScopeBeginContext) GetParser() antlr.Parser { return s.parser }

func (s *ScopeBeginContext) BEGIN() antlr.TerminalNode {
	return s.GetToken(SashimiParserBEGIN, 0)
}

func (s *ScopeBeginContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(SashimiParserLPAREN, 0)
}

func (s *ScopeBeginContext) SCOPEIDENT() antlr.TerminalNode {
	return s.GetToken(SashimiParserSCOPEIDENT, 0)
}

func (s *ScopeBeginContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(SashimiParserRPAREN, 0)
}

func (s *ScopeBeginContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ScopeBeginContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ScopeBeginContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SashimiListener); ok {
		listenerT.EnterScopeBegin(s)
	}
}

func (s *ScopeBeginContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SashimiListener); ok {
		listenerT.ExitScopeBegin(s)
	}
}

func (p *SashimiParser) ScopeBegin() (localctx IScopeBeginContext) {
	localctx = NewScopeBeginContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, SashimiParserRULE_scopeBegin)
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
		p.Match(SashimiParserBEGIN)
	}
	p.SetState(170)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SashimiParserLPAREN {
		{
			p.SetState(167)
			p.Match(SashimiParserLPAREN)
		}
		{
			p.SetState(168)
			p.Match(SashimiParserSCOPEIDENT)
		}
		{
			p.SetState(169)
			p.Match(SashimiParserRPAREN)
		}

	}

	return localctx
}

// IScopeEndContext is an interface to support dynamic dispatch.
type IScopeEndContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsScopeEndContext differentiates from other interfaces.
	IsScopeEndContext()
}

type ScopeEndContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyScopeEndContext() *ScopeEndContext {
	var p = new(ScopeEndContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SashimiParserRULE_scopeEnd
	return p
}

func (*ScopeEndContext) IsScopeEndContext() {}

func NewScopeEndContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ScopeEndContext {
	var p = new(ScopeEndContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SashimiParserRULE_scopeEnd

	return p
}

func (s *ScopeEndContext) GetParser() antlr.Parser { return s.parser }

func (s *ScopeEndContext) DIRECTIVE() antlr.TerminalNode {
	return s.GetToken(SashimiParserDIRECTIVE, 0)
}

func (s *ScopeEndContext) END() antlr.TerminalNode {
	return s.GetToken(SashimiParserEND, 0)
}

func (s *ScopeEndContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(SashimiParserLPAREN, 0)
}

func (s *ScopeEndContext) SCOPEIDENT() antlr.TerminalNode {
	return s.GetToken(SashimiParserSCOPEIDENT, 0)
}

func (s *ScopeEndContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(SashimiParserRPAREN, 0)
}

func (s *ScopeEndContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ScopeEndContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ScopeEndContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SashimiListener); ok {
		listenerT.EnterScopeEnd(s)
	}
}

func (s *ScopeEndContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SashimiListener); ok {
		listenerT.ExitScopeEnd(s)
	}
}

func (p *SashimiParser) ScopeEnd() (localctx IScopeEndContext) {
	localctx = NewScopeEndContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, SashimiParserRULE_scopeEnd)
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
		p.SetState(172)
		p.Match(SashimiParserDIRECTIVE)
	}
	{
		p.SetState(173)
		p.Match(SashimiParserEND)
	}
	p.SetState(177)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SashimiParserLPAREN {
		{
			p.SetState(174)
			p.Match(SashimiParserLPAREN)
		}
		{
			p.SetState(175)
			p.Match(SashimiParserSCOPEIDENT)
		}
		{
			p.SetState(176)
			p.Match(SashimiParserRPAREN)
		}

	}

	return localctx
}

// ICommandCallContext is an interface to support dynamic dispatch.
type ICommandCallContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCommandCallContext differentiates from other interfaces.
	IsCommandCallContext()
}

type CommandCallContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCommandCallContext() *CommandCallContext {
	var p = new(CommandCallContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SashimiParserRULE_commandCall
	return p
}

func (*CommandCallContext) IsCommandCallContext() {}

func NewCommandCallContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CommandCallContext {
	var p = new(CommandCallContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SashimiParserRULE_commandCall

	return p
}

func (s *CommandCallContext) GetParser() antlr.Parser { return s.parser }

func (s *CommandCallContext) COMMAND() antlr.TerminalNode {
	return s.GetToken(SashimiParserCOMMAND, 0)
}

func (s *CommandCallContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(SashimiParserLPAREN, 0)
}

func (s *CommandCallContext) Qualifier() IQualifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IQualifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IQualifierContext)
}

func (s *CommandCallContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(SashimiParserRPAREN, 0)
}

func (s *CommandCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CommandCallContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CommandCallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SashimiListener); ok {
		listenerT.EnterCommandCall(s)
	}
}

func (s *CommandCallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SashimiListener); ok {
		listenerT.ExitCommandCall(s)
	}
}

func (p *SashimiParser) CommandCall() (localctx ICommandCallContext) {
	localctx = NewCommandCallContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, SashimiParserRULE_commandCall)

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
		p.SetState(179)
		p.Match(SashimiParserCOMMAND)
	}
	{
		p.SetState(180)
		p.Match(SashimiParserLPAREN)
	}
	{
		p.SetState(181)
		p.Qualifier()
	}
	{
		p.SetState(182)
		p.Match(SashimiParserRPAREN)
	}

	return localctx
}

// IBlockScopeContext is an interface to support dynamic dispatch.
type IBlockScopeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBlockScopeContext differentiates from other interfaces.
	IsBlockScopeContext()
}

type BlockScopeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockScopeContext() *BlockScopeContext {
	var p = new(BlockScopeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SashimiParserRULE_blockScope
	return p
}

func (*BlockScopeContext) IsBlockScopeContext() {}

func NewBlockScopeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockScopeContext {
	var p = new(BlockScopeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SashimiParserRULE_blockScope

	return p
}

func (s *BlockScopeContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockScopeContext) ScopeBegin() IScopeBeginContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IScopeBeginContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IScopeBeginContext)
}

func (s *BlockScopeContext) ScopeEnd() IScopeEndContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IScopeEndContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IScopeEndContext)
}

func (s *BlockScopeContext) Export() IExportContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExportContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExportContext)
}

func (s *BlockScopeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockScopeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockScopeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SashimiListener); ok {
		listenerT.EnterBlockScope(s)
	}
}

func (s *BlockScopeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SashimiListener); ok {
		listenerT.ExitBlockScope(s)
	}
}

func (p *SashimiParser) BlockScope() (localctx IBlockScopeContext) {
	localctx = NewBlockScopeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, SashimiParserRULE_blockScope)

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
		p.ScopeBegin()
	}
	p.SetState(186)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 16, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(185)
			p.Export()
		}

	}
	{
		p.SetState(188)
		p.ScopeEnd()
	}

	return localctx
}

// IExportContext is an interface to support dynamic dispatch.
type IExportContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExportContext differentiates from other interfaces.
	IsExportContext()
}

type ExportContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExportContext() *ExportContext {
	var p = new(ExportContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SashimiParserRULE_export
	return p
}

func (*ExportContext) IsExportContext() {}

func NewExportContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExportContext {
	var p = new(ExportContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SashimiParserRULE_export

	return p
}

func (s *ExportContext) GetParser() antlr.Parser { return s.parser }

func (s *ExportContext) DIRECTIVE() antlr.TerminalNode {
	return s.GetToken(SashimiParserDIRECTIVE, 0)
}

func (s *ExportContext) CommandCall() ICommandCallContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICommandCallContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICommandCallContext)
}

func (s *ExportContext) LoopCall() ILoopCallContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILoopCallContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILoopCallContext)
}

func (s *ExportContext) EntityDef() IEntityDefContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEntityDefContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEntityDefContext)
}

func (s *ExportContext) BlockScope() IBlockScopeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockScopeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockScopeContext)
}

func (s *ExportContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExportContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExportContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SashimiListener); ok {
		listenerT.EnterExport(s)
	}
}

func (s *ExportContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SashimiListener); ok {
		listenerT.ExitExport(s)
	}
}

func (p *SashimiParser) Export() (localctx IExportContext) {
	localctx = NewExportContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, SashimiParserRULE_export)

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
		p.SetState(190)
		p.Match(SashimiParserDIRECTIVE)
	}
	p.SetState(195)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SashimiParserCOMMAND:
		{
			p.SetState(191)
			p.CommandCall()
		}

	case SashimiParserLOOP:
		{
			p.SetState(192)
			p.LoopCall()
		}

	case SashimiParserENTITY, SashimiParserUNIQUE:
		{
			p.SetState(193)
			p.EntityDef()
		}

	case SashimiParserBEGIN:
		{
			p.SetState(194)
			p.BlockScope()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
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
	p.RuleIndex = SashimiParserRULE_block
	return p
}

func (*BlockContext) IsBlockContext() {}

func NewBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockContext {
	var p = new(BlockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SashimiParserRULE_block

	return p
}

func (s *BlockContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockContext) EOF() antlr.TerminalNode {
	return s.GetToken(SashimiParserEOF, 0)
}

func (s *BlockContext) AllExport() []IExportContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExportContext)(nil)).Elem())
	var tst = make([]IExportContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExportContext)
		}
	}

	return tst
}

func (s *BlockContext) Export(i int) IExportContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExportContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExportContext)
}

func (s *BlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SashimiListener); ok {
		listenerT.EnterBlock(s)
	}
}

func (s *BlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SashimiListener); ok {
		listenerT.ExitBlock(s)
	}
}

func (p *SashimiParser) Block() (localctx IBlockContext) {
	localctx = NewBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, SashimiParserRULE_block)
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
	p.SetState(200)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == SashimiParserDIRECTIVE {
		{
			p.SetState(197)
			p.Export()
		}

		p.SetState(202)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(203)
		p.Match(SashimiParserEOF)
	}

	return localctx
}

// IBoolExpressionContext is an interface to support dynamic dispatch.
type IBoolExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetLeft returns the left rule contexts.
	GetLeft() IBoolExpressionContext

	// GetRight returns the right rule contexts.
	GetRight() IBoolExpressionContext

	// SetLeft sets the left rule contexts.
	SetLeft(IBoolExpressionContext)

	// SetRight sets the right rule contexts.
	SetRight(IBoolExpressionContext)

	// IsBoolExpressionContext differentiates from other interfaces.
	IsBoolExpressionContext()
}

type BoolExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	left   IBoolExpressionContext
	right  IBoolExpressionContext
}

func NewEmptyBoolExpressionContext() *BoolExpressionContext {
	var p = new(BoolExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SashimiParserRULE_boolExpression
	return p
}

func (*BoolExpressionContext) IsBoolExpressionContext() {}

func NewBoolExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BoolExpressionContext {
	var p = new(BoolExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SashimiParserRULE_boolExpression

	return p
}

func (s *BoolExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *BoolExpressionContext) GetLeft() IBoolExpressionContext { return s.left }

func (s *BoolExpressionContext) GetRight() IBoolExpressionContext { return s.right }

func (s *BoolExpressionContext) SetLeft(v IBoolExpressionContext) { s.left = v }

func (s *BoolExpressionContext) SetRight(v IBoolExpressionContext) { s.right = v }

func (s *BoolExpressionContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(SashimiParserLPAREN, 0)
}

func (s *BoolExpressionContext) AllBoolExpression() []IBoolExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IBoolExpressionContext)(nil)).Elem())
	var tst = make([]IBoolExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IBoolExpressionContext)
		}
	}

	return tst
}

func (s *BoolExpressionContext) BoolExpression(i int) IBoolExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBoolExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IBoolExpressionContext)
}

func (s *BoolExpressionContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(SashimiParserRPAREN, 0)
}

func (s *BoolExpressionContext) NOT() antlr.TerminalNode {
	return s.GetToken(SashimiParserNOT, 0)
}

func (s *BoolExpressionContext) Truth() ITruthContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITruthContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITruthContext)
}

func (s *BoolExpressionContext) Qualifier() IQualifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IQualifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IQualifierContext)
}

func (s *BoolExpressionContext) ALIAS() antlr.TerminalNode {
	return s.GetToken(SashimiParserALIAS, 0)
}

func (s *BoolExpressionContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(SashimiParserNUMBER, 0)
}

func (s *BoolExpressionContext) Op() IOpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOpContext)
}

func (s *BoolExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BoolExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BoolExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SashimiListener); ok {
		listenerT.EnterBoolExpression(s)
	}
}

func (s *BoolExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SashimiListener); ok {
		listenerT.ExitBoolExpression(s)
	}
}

func (p *SashimiParser) BoolExpression() (localctx IBoolExpressionContext) {
	return p.boolExpression(0)
}

func (p *SashimiParser) boolExpression(_p int) (localctx IBoolExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewBoolExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IBoolExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 42
	p.EnterRecursionRule(localctx, 42, SashimiParserRULE_boolExpression, _p)

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
	p.SetState(216)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SashimiParserLPAREN:
		{
			p.SetState(206)
			p.Match(SashimiParserLPAREN)
		}
		{
			p.SetState(207)
			p.boolExpression(0)
		}
		{
			p.SetState(208)
			p.Match(SashimiParserRPAREN)
		}

	case SashimiParserNOT:
		{
			p.SetState(210)
			p.Match(SashimiParserNOT)
		}
		{
			p.SetState(211)
			p.boolExpression(6)
		}

	case SashimiParserTRUE, SashimiParserFALSE:
		{
			p.SetState(212)
			p.Truth()
		}

	case SashimiParserIDENT:
		{
			p.SetState(213)
			p.Qualifier()
		}

	case SashimiParserALIAS:
		{
			p.SetState(214)
			p.Match(SashimiParserALIAS)
		}

	case SashimiParserNUMBER:
		{
			p.SetState(215)
			p.Match(SashimiParserNUMBER)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(224)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewBoolExpressionContext(p, _parentctx, _parentState)
			localctx.(*BoolExpressionContext).left = _prevctx
			p.PushNewRecursionContext(localctx, _startState, SashimiParserRULE_boolExpression)
			p.SetState(218)

			if !(p.Precpred(p.GetParserRuleContext(), 5)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
			}
			{
				p.SetState(219)
				p.Op()
			}
			{
				p.SetState(220)

				var _x = p.boolExpression(6)

				localctx.(*BoolExpressionContext).right = _x
			}

		}
		p.SetState(226)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext())
	}

	return localctx
}

// IOpContext is an interface to support dynamic dispatch.
type IOpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOpContext differentiates from other interfaces.
	IsOpContext()
}

type OpContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOpContext() *OpContext {
	var p = new(OpContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SashimiParserRULE_op
	return p
}

func (*OpContext) IsOpContext() {}

func NewOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OpContext {
	var p = new(OpContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SashimiParserRULE_op

	return p
}

func (s *OpContext) GetParser() antlr.Parser { return s.parser }

func (s *OpContext) Binary() IBinaryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBinaryContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBinaryContext)
}

func (s *OpContext) Comparator() IComparatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IComparatorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IComparatorContext)
}

func (s *OpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SashimiListener); ok {
		listenerT.EnterOp(s)
	}
}

func (s *OpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SashimiListener); ok {
		listenerT.ExitOp(s)
	}
}

func (p *SashimiParser) Op() (localctx IOpContext) {
	localctx = NewOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, SashimiParserRULE_op)

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

	p.SetState(229)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SashimiParserAND, SashimiParserOR:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(227)
			p.Binary()
		}

	case SashimiParserEQ, SashimiParserLEQ, SashimiParserLT, SashimiParserGEQ, SashimiParserGT, SashimiParserNEQ:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(228)
			p.Comparator()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IComparatorContext is an interface to support dynamic dispatch.
type IComparatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsComparatorContext differentiates from other interfaces.
	IsComparatorContext()
}

type ComparatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyComparatorContext() *ComparatorContext {
	var p = new(ComparatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SashimiParserRULE_comparator
	return p
}

func (*ComparatorContext) IsComparatorContext() {}

func NewComparatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ComparatorContext {
	var p = new(ComparatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SashimiParserRULE_comparator

	return p
}

func (s *ComparatorContext) GetParser() antlr.Parser { return s.parser }

func (s *ComparatorContext) EQ() antlr.TerminalNode {
	return s.GetToken(SashimiParserEQ, 0)
}

func (s *ComparatorContext) LEQ() antlr.TerminalNode {
	return s.GetToken(SashimiParserLEQ, 0)
}

func (s *ComparatorContext) LT() antlr.TerminalNode {
	return s.GetToken(SashimiParserLT, 0)
}

func (s *ComparatorContext) GEQ() antlr.TerminalNode {
	return s.GetToken(SashimiParserGEQ, 0)
}

func (s *ComparatorContext) GT() antlr.TerminalNode {
	return s.GetToken(SashimiParserGT, 0)
}

func (s *ComparatorContext) NEQ() antlr.TerminalNode {
	return s.GetToken(SashimiParserNEQ, 0)
}

func (s *ComparatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ComparatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ComparatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SashimiListener); ok {
		listenerT.EnterComparator(s)
	}
}

func (s *ComparatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SashimiListener); ok {
		listenerT.ExitComparator(s)
	}
}

func (p *SashimiParser) Comparator() (localctx IComparatorContext) {
	localctx = NewComparatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, SashimiParserRULE_comparator)
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
		p.SetState(231)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<SashimiParserEQ)|(1<<SashimiParserLEQ)|(1<<SashimiParserLT)|(1<<SashimiParserGEQ)|(1<<SashimiParserGT)|(1<<SashimiParserNEQ))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IBinaryContext is an interface to support dynamic dispatch.
type IBinaryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBinaryContext differentiates from other interfaces.
	IsBinaryContext()
}

type BinaryContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBinaryContext() *BinaryContext {
	var p = new(BinaryContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SashimiParserRULE_binary
	return p
}

func (*BinaryContext) IsBinaryContext() {}

func NewBinaryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BinaryContext {
	var p = new(BinaryContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SashimiParserRULE_binary

	return p
}

func (s *BinaryContext) GetParser() antlr.Parser { return s.parser }

func (s *BinaryContext) AND() antlr.TerminalNode {
	return s.GetToken(SashimiParserAND, 0)
}

func (s *BinaryContext) OR() antlr.TerminalNode {
	return s.GetToken(SashimiParserOR, 0)
}

func (s *BinaryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BinaryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BinaryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SashimiListener); ok {
		listenerT.EnterBinary(s)
	}
}

func (s *BinaryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SashimiListener); ok {
		listenerT.ExitBinary(s)
	}
}

func (p *SashimiParser) Binary() (localctx IBinaryContext) {
	localctx = NewBinaryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, SashimiParserRULE_binary)
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
		p.SetState(233)
		_la = p.GetTokenStream().LA(1)

		if !(_la == SashimiParserAND || _la == SashimiParserOR) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// ITruthContext is an interface to support dynamic dispatch.
type ITruthContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTruthContext differentiates from other interfaces.
	IsTruthContext()
}

type TruthContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTruthContext() *TruthContext {
	var p = new(TruthContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SashimiParserRULE_truth
	return p
}

func (*TruthContext) IsTruthContext() {}

func NewTruthContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TruthContext {
	var p = new(TruthContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SashimiParserRULE_truth

	return p
}

func (s *TruthContext) GetParser() antlr.Parser { return s.parser }

func (s *TruthContext) TRUE() antlr.TerminalNode {
	return s.GetToken(SashimiParserTRUE, 0)
}

func (s *TruthContext) FALSE() antlr.TerminalNode {
	return s.GetToken(SashimiParserFALSE, 0)
}

func (s *TruthContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TruthContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TruthContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SashimiListener); ok {
		listenerT.EnterTruth(s)
	}
}

func (s *TruthContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SashimiListener); ok {
		listenerT.ExitTruth(s)
	}
}

func (p *SashimiParser) Truth() (localctx ITruthContext) {
	localctx = NewTruthContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, SashimiParserRULE_truth)
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
		p.SetState(235)
		_la = p.GetTokenStream().LA(1)

		if !(_la == SashimiParserTRUE || _la == SashimiParserFALSE) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

func (p *SashimiParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 21:
		var t *BoolExpressionContext = nil
		if localctx != nil {
			t = localctx.(*BoolExpressionContext)
		}
		return p.BoolExpression_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *SashimiParser) BoolExpression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 5)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
