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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 37, 208,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 3, 2, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3,
	4, 5, 4, 59, 10, 4, 3, 4, 3, 4, 3, 4, 5, 4, 64, 10, 4, 7, 4, 66, 10, 4,
	12, 4, 14, 4, 69, 11, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 7, 5,
	78, 10, 5, 12, 5, 14, 5, 81, 11, 5, 3, 5, 3, 5, 3, 6, 3, 6, 5, 6, 87, 10,
	6, 3, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 5, 8, 95, 10, 8, 3, 9, 3, 9, 3,
	9, 3, 9, 5, 9, 101, 10, 9, 3, 10, 3, 10, 3, 10, 3, 11, 3, 11, 3, 11, 3,
	12, 3, 12, 3, 12, 5, 12, 112, 10, 12, 3, 12, 3, 12, 3, 13, 3, 13, 3, 13,
	3, 13, 3, 14, 3, 14, 3, 14, 7, 14, 123, 10, 14, 12, 14, 14, 14, 126, 11,
	14, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 5, 15, 134, 10, 15, 3, 15,
	3, 15, 3, 15, 3, 15, 5, 15, 140, 10, 15, 3, 16, 5, 16, 143, 10, 16, 3,
	16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 7, 16, 152, 10, 16, 12, 16,
	14, 16, 155, 11, 16, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 18, 3, 18, 3,
	18, 3, 18, 5, 18, 166, 10, 18, 3, 19, 7, 19, 169, 10, 19, 12, 19, 14, 19,
	172, 11, 19, 3, 19, 3, 19, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3,
	20, 3, 20, 3, 20, 3, 20, 3, 20, 5, 20, 187, 10, 20, 3, 20, 3, 20, 3, 20,
	3, 20, 7, 20, 193, 10, 20, 12, 20, 14, 20, 196, 11, 20, 3, 21, 3, 21, 5,
	21, 200, 10, 21, 3, 22, 3, 22, 3, 23, 3, 23, 3, 24, 3, 24, 3, 24, 2, 3,
	38, 25, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34,
	36, 38, 40, 42, 44, 46, 2, 5, 4, 2, 18, 22, 24, 24, 3, 2, 25, 26, 3, 2,
	35, 36, 2, 209, 2, 48, 3, 2, 2, 2, 4, 52, 3, 2, 2, 2, 6, 55, 3, 2, 2, 2,
	8, 72, 3, 2, 2, 2, 10, 84, 3, 2, 2, 2, 12, 88, 3, 2, 2, 2, 14, 91, 3, 2,
	2, 2, 16, 100, 3, 2, 2, 2, 18, 102, 3, 2, 2, 2, 20, 105, 3, 2, 2, 2, 22,
	108, 3, 2, 2, 2, 24, 115, 3, 2, 2, 2, 26, 119, 3, 2, 2, 2, 28, 127, 3,
	2, 2, 2, 30, 142, 3, 2, 2, 2, 32, 156, 3, 2, 2, 2, 34, 161, 3, 2, 2, 2,
	36, 170, 3, 2, 2, 2, 38, 186, 3, 2, 2, 2, 40, 199, 3, 2, 2, 2, 42, 201,
	3, 2, 2, 2, 44, 203, 3, 2, 2, 2, 46, 205, 3, 2, 2, 2, 48, 49, 7, 33, 2,
	2, 49, 50, 7, 18, 2, 2, 50, 51, 7, 33, 2, 2, 51, 3, 3, 2, 2, 2, 52, 53,
	7, 27, 2, 2, 53, 54, 7, 33, 2, 2, 54, 5, 3, 2, 2, 2, 55, 58, 7, 16, 2,
	2, 56, 59, 5, 4, 3, 2, 57, 59, 5, 2, 2, 2, 58, 56, 3, 2, 2, 2, 58, 57,
	3, 2, 2, 2, 59, 67, 3, 2, 2, 2, 60, 63, 7, 8, 2, 2, 61, 64, 5, 4, 3, 2,
	62, 64, 5, 2, 2, 2, 63, 61, 3, 2, 2, 2, 63, 62, 3, 2, 2, 2, 64, 66, 3,
	2, 2, 2, 65, 60, 3, 2, 2, 2, 66, 69, 3, 2, 2, 2, 67, 65, 3, 2, 2, 2, 67,
	68, 3, 2, 2, 2, 68, 70, 3, 2, 2, 2, 69, 67, 3, 2, 2, 2, 70, 71, 7, 17,
	2, 2, 71, 7, 3, 2, 2, 2, 72, 73, 7, 30, 2, 2, 73, 74, 7, 14, 2, 2, 74,
	79, 7, 32, 2, 2, 75, 76, 7, 8, 2, 2, 76, 78, 7, 32, 2, 2, 77, 75, 3, 2,
	2, 2, 78, 81, 3, 2, 2, 2, 79, 77, 3, 2, 2, 2, 79, 80, 3, 2, 2, 2, 80, 82,
	3, 2, 2, 2, 81, 79, 3, 2, 2, 2, 82, 83, 7, 15, 2, 2, 83, 9, 3, 2, 2, 2,
	84, 86, 7, 29, 2, 2, 85, 87, 5, 6, 4, 2, 86, 85, 3, 2, 2, 2, 86, 87, 3,
	2, 2, 2, 87, 11, 3, 2, 2, 2, 88, 89, 7, 6, 2, 2, 89, 90, 7, 33, 2, 2, 90,
	13, 3, 2, 2, 2, 91, 94, 7, 31, 2, 2, 92, 95, 5, 10, 6, 2, 93, 95, 5, 12,
	7, 2, 94, 92, 3, 2, 2, 2, 94, 93, 3, 2, 2, 2, 95, 15, 3, 2, 2, 2, 96, 101,
	5, 14, 8, 2, 97, 101, 5, 8, 5, 2, 98, 101, 5, 10, 6, 2, 99, 101, 5, 12,
	7, 2, 100, 96, 3, 2, 2, 2, 100, 97, 3, 2, 2, 2, 100, 98, 3, 2, 2, 2, 100,
	99, 3, 2, 2, 2, 101, 17, 3, 2, 2, 2, 102, 103, 7, 12, 2, 2, 103, 104, 7,
	32, 2, 2, 104, 19, 3, 2, 2, 2, 105, 106, 7, 11, 2, 2, 106, 107, 5, 16,
	9, 2, 107, 21, 3, 2, 2, 2, 108, 109, 7, 9, 2, 2, 109, 111, 7, 33, 2, 2,
	110, 112, 5, 18, 10, 2, 111, 110, 3, 2, 2, 2, 111, 112, 3, 2, 2, 2, 112,
	113, 3, 2, 2, 2, 113, 114, 5, 20, 11, 2, 114, 23, 3, 2, 2, 2, 115, 116,
	7, 33, 2, 2, 116, 117, 7, 13, 2, 2, 117, 118, 5, 38, 20, 2, 118, 25, 3,
	2, 2, 2, 119, 124, 7, 33, 2, 2, 120, 121, 7, 28, 2, 2, 121, 123, 7, 33,
	2, 2, 122, 120, 3, 2, 2, 2, 123, 126, 3, 2, 2, 2, 124, 122, 3, 2, 2, 2,
	124, 125, 3, 2, 2, 2, 125, 27, 3, 2, 2, 2, 126, 124, 3, 2, 2, 2, 127, 128,
	7, 5, 2, 2, 128, 129, 7, 14, 2, 2, 129, 130, 5, 26, 14, 2, 130, 133, 7,
	15, 2, 2, 131, 132, 7, 12, 2, 2, 132, 134, 7, 33, 2, 2, 133, 131, 3, 2,
	2, 2, 133, 134, 3, 2, 2, 2, 134, 139, 3, 2, 2, 2, 135, 136, 7, 16, 2, 2,
	136, 137, 5, 24, 13, 2, 137, 138, 7, 17, 2, 2, 138, 140, 3, 2, 2, 2, 139,
	135, 3, 2, 2, 2, 139, 140, 3, 2, 2, 2, 140, 29, 3, 2, 2, 2, 141, 143, 7,
	7, 2, 2, 142, 141, 3, 2, 2, 2, 142, 143, 3, 2, 2, 2, 143, 144, 3, 2, 2,
	2, 144, 145, 7, 6, 2, 2, 145, 146, 7, 14, 2, 2, 146, 147, 7, 33, 2, 2,
	147, 148, 7, 15, 2, 2, 148, 149, 7, 10, 2, 2, 149, 153, 5, 22, 12, 2, 150,
	152, 5, 22, 12, 2, 151, 150, 3, 2, 2, 2, 152, 155, 3, 2, 2, 2, 153, 151,
	3, 2, 2, 2, 153, 154, 3, 2, 2, 2, 154, 31, 3, 2, 2, 2, 155, 153, 3, 2,
	2, 2, 156, 157, 7, 4, 2, 2, 157, 158, 7, 14, 2, 2, 158, 159, 5, 26, 14,
	2, 159, 160, 7, 15, 2, 2, 160, 33, 3, 2, 2, 2, 161, 165, 7, 3, 2, 2, 162,
	166, 5, 32, 17, 2, 163, 166, 5, 28, 15, 2, 164, 166, 5, 30, 16, 2, 165,
	162, 3, 2, 2, 2, 165, 163, 3, 2, 2, 2, 165, 164, 3, 2, 2, 2, 166, 35, 3,
	2, 2, 2, 167, 169, 5, 34, 18, 2, 168, 167, 3, 2, 2, 2, 169, 172, 3, 2,
	2, 2, 170, 168, 3, 2, 2, 2, 170, 171, 3, 2, 2, 2, 171, 173, 3, 2, 2, 2,
	172, 170, 3, 2, 2, 2, 173, 174, 7, 2, 2, 3, 174, 37, 3, 2, 2, 2, 175, 176,
	8, 20, 1, 2, 176, 177, 7, 14, 2, 2, 177, 178, 5, 38, 20, 2, 178, 179, 7,
	15, 2, 2, 179, 187, 3, 2, 2, 2, 180, 181, 7, 23, 2, 2, 181, 187, 5, 38,
	20, 8, 182, 187, 5, 46, 24, 2, 183, 187, 5, 26, 14, 2, 184, 187, 7, 32,
	2, 2, 185, 187, 7, 34, 2, 2, 186, 175, 3, 2, 2, 2, 186, 180, 3, 2, 2, 2,
	186, 182, 3, 2, 2, 2, 186, 183, 3, 2, 2, 2, 186, 184, 3, 2, 2, 2, 186,
	185, 3, 2, 2, 2, 187, 194, 3, 2, 2, 2, 188, 189, 12, 7, 2, 2, 189, 190,
	5, 40, 21, 2, 190, 191, 5, 38, 20, 8, 191, 193, 3, 2, 2, 2, 192, 188, 3,
	2, 2, 2, 193, 196, 3, 2, 2, 2, 194, 192, 3, 2, 2, 2, 194, 195, 3, 2, 2,
	2, 195, 39, 3, 2, 2, 2, 196, 194, 3, 2, 2, 2, 197, 200, 5, 44, 23, 2, 198,
	200, 5, 42, 22, 2, 199, 197, 3, 2, 2, 2, 199, 198, 3, 2, 2, 2, 200, 41,
	3, 2, 2, 2, 201, 202, 9, 2, 2, 2, 202, 43, 3, 2, 2, 2, 203, 204, 9, 3,
	2, 2, 204, 45, 3, 2, 2, 2, 205, 206, 9, 4, 2, 2, 206, 47, 3, 2, 2, 2, 20,
	58, 63, 67, 79, 86, 94, 100, 111, 124, 133, 139, 142, 153, 165, 170, 186,
	194, 199,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'sashimi:'", "", "'repeat'", "'entity'", "'unique'", "','", "'-'",
	"'of'", "'is'", "'as'", "'->'", "'('", "')'", "'['", "']'", "'='", "'<='",
	"'<'", "'>='", "'>'", "'!'", "'<>'", "'&'", "'|'", "':'", "'.'", "", "",
	"'list'", "", "", "", "'true'", "'false'",
}
var symbolicNames = []string{
	"", "DIRECTIVE", "COMMAND", "LOOP", "ENTITY", "UNIQUE", "SEPERATOR", "PROP_START",
	"OF", "IS", "AS", "ARROW", "LPAREN", "RPAREN", "HLPAREN", "HRPAREN", "EQ",
	"LEQ", "LT", "GEQ", "GT", "NOT", "NEQ", "AND", "OR", "ATOM", "DOT", "TYPE",
	"UNION", "LIST", "ALIAS", "IDENT", "NUMBER", "TRUE", "FALSE", "WS",
}

var ruleNames = []string{
	"keyValuePair", "keyAtom", "constraintList", "unionDecl", "typeDecl", "entityRef",
	"listDecl", "typeDef", "aliasDecl", "typeIs", "propDecl", "predicate",
	"qualifier", "loopCall", "entityDef", "commandCall", "export", "block",
	"boolExpression", "op", "comparator", "binary", "truth",
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
	SashimiParserARROW      = 11
	SashimiParserLPAREN     = 12
	SashimiParserRPAREN     = 13
	SashimiParserHLPAREN    = 14
	SashimiParserHRPAREN    = 15
	SashimiParserEQ         = 16
	SashimiParserLEQ        = 17
	SashimiParserLT         = 18
	SashimiParserGEQ        = 19
	SashimiParserGT         = 20
	SashimiParserNOT        = 21
	SashimiParserNEQ        = 22
	SashimiParserAND        = 23
	SashimiParserOR         = 24
	SashimiParserATOM       = 25
	SashimiParserDOT        = 26
	SashimiParserTYPE       = 27
	SashimiParserUNION      = 28
	SashimiParserLIST       = 29
	SashimiParserALIAS      = 30
	SashimiParserIDENT      = 31
	SashimiParserNUMBER     = 32
	SashimiParserTRUE       = 33
	SashimiParserFALSE      = 34
	SashimiParserWS         = 35
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
	SashimiParserRULE_commandCall    = 15
	SashimiParserRULE_export         = 16
	SashimiParserRULE_block          = 17
	SashimiParserRULE_boolExpression = 18
	SashimiParserRULE_op             = 19
	SashimiParserRULE_comparator     = 20
	SashimiParserRULE_binary         = 21
	SashimiParserRULE_truth          = 22
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
		p.SetState(46)
		p.Match(SashimiParserIDENT)
	}
	{
		p.SetState(47)
		p.Match(SashimiParserEQ)
	}
	{
		p.SetState(48)
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
		p.SetState(50)
		p.Match(SashimiParserATOM)
	}
	{
		p.SetState(51)
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
		p.SetState(53)
		p.Match(SashimiParserHLPAREN)
	}
	p.SetState(56)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SashimiParserATOM:
		{
			p.SetState(54)
			p.KeyAtom()
		}

	case SashimiParserIDENT:
		{
			p.SetState(55)
			p.KeyValuePair()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.SetState(65)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == SashimiParserSEPERATOR {
		{
			p.SetState(58)
			p.Match(SashimiParserSEPERATOR)
		}
		p.SetState(61)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case SashimiParserATOM:
			{
				p.SetState(59)
				p.KeyAtom()
			}

		case SashimiParserIDENT:
			{
				p.SetState(60)
				p.KeyValuePair()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(67)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(68)
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
		p.SetState(70)
		p.Match(SashimiParserUNION)
	}
	{
		p.SetState(71)
		p.Match(SashimiParserLPAREN)
	}
	{
		p.SetState(72)
		p.Match(SashimiParserALIAS)
	}
	p.SetState(77)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == SashimiParserSEPERATOR {
		{
			p.SetState(73)
			p.Match(SashimiParserSEPERATOR)
		}
		{
			p.SetState(74)
			p.Match(SashimiParserALIAS)
		}

		p.SetState(79)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(80)
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
		p.SetState(82)
		p.Match(SashimiParserTYPE)
	}
	p.SetState(84)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SashimiParserHLPAREN {
		{
			p.SetState(83)
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
		p.SetState(86)
		p.Match(SashimiParserENTITY)
	}
	{
		p.SetState(87)
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
		p.SetState(89)
		p.Match(SashimiParserLIST)
	}
	p.SetState(92)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SashimiParserTYPE:
		{
			p.SetState(90)
			p.TypeDecl()
		}

	case SashimiParserENTITY:
		{
			p.SetState(91)
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

	p.SetState(98)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SashimiParserLIST:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(94)
			p.ListDecl()
		}

	case SashimiParserUNION:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(95)
			p.UnionDecl()
		}

	case SashimiParserTYPE:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(96)
			p.TypeDecl()
		}

	case SashimiParserENTITY:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(97)
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
		p.SetState(100)
		p.Match(SashimiParserAS)
	}
	{
		p.SetState(101)
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
		p.SetState(103)
		p.Match(SashimiParserIS)
	}
	{
		p.SetState(104)
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
		p.SetState(106)
		p.Match(SashimiParserPROP_START)
	}
	{
		p.SetState(107)
		p.Match(SashimiParserIDENT)
	}
	p.SetState(109)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SashimiParserAS {
		{
			p.SetState(108)
			p.AliasDecl()
		}

	}
	{
		p.SetState(111)
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
		p.SetState(113)
		p.Match(SashimiParserIDENT)
	}
	{
		p.SetState(114)
		p.Match(SashimiParserARROW)
	}
	{
		p.SetState(115)
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
		p.SetState(117)
		p.Match(SashimiParserIDENT)
	}
	p.SetState(122)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(118)
				p.Match(SashimiParserDOT)
			}
			{
				p.SetState(119)
				p.Match(SashimiParserIDENT)
			}

		}
		p.SetState(124)
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

func (s *LoopCallContext) Qualifier() IQualifierContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IQualifierContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IQualifierContext)
}

func (s *LoopCallContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(SashimiParserRPAREN, 0)
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
		p.SetState(125)
		p.Match(SashimiParserLOOP)
	}
	{
		p.SetState(126)
		p.Match(SashimiParserLPAREN)
	}
	{
		p.SetState(127)
		p.Qualifier()
	}
	{
		p.SetState(128)
		p.Match(SashimiParserRPAREN)
	}
	p.SetState(131)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SashimiParserAS {
		{
			p.SetState(129)
			p.Match(SashimiParserAS)
		}
		{
			p.SetState(130)

			var _m = p.Match(SashimiParserIDENT)

			localctx.(*LoopCallContext).alias = _m
		}

	}
	p.SetState(137)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SashimiParserHLPAREN {
		{
			p.SetState(133)
			p.Match(SashimiParserHLPAREN)
		}
		{
			p.SetState(134)
			p.Predicate()
		}
		{
			p.SetState(135)
			p.Match(SashimiParserHRPAREN)
		}

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
	p.SetState(140)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SashimiParserUNIQUE {
		{
			p.SetState(139)
			p.Match(SashimiParserUNIQUE)
		}

	}
	{
		p.SetState(142)
		p.Match(SashimiParserENTITY)
	}
	{
		p.SetState(143)
		p.Match(SashimiParserLPAREN)
	}
	{
		p.SetState(144)
		p.Match(SashimiParserIDENT)
	}
	{
		p.SetState(145)
		p.Match(SashimiParserRPAREN)
	}
	{
		p.SetState(146)
		p.Match(SashimiParserOF)
	}
	{
		p.SetState(147)
		p.PropDecl()
	}
	p.SetState(151)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == SashimiParserPROP_START {
		{
			p.SetState(148)
			p.PropDecl()
		}

		p.SetState(153)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
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
	p.EnterRule(localctx, 30, SashimiParserRULE_commandCall)

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
		p.SetState(154)
		p.Match(SashimiParserCOMMAND)
	}
	{
		p.SetState(155)
		p.Match(SashimiParserLPAREN)
	}
	{
		p.SetState(156)
		p.Qualifier()
	}
	{
		p.SetState(157)
		p.Match(SashimiParserRPAREN)
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
	p.EnterRule(localctx, 32, SashimiParserRULE_export)

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
		p.SetState(159)
		p.Match(SashimiParserDIRECTIVE)
	}
	p.SetState(163)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SashimiParserCOMMAND:
		{
			p.SetState(160)
			p.CommandCall()
		}

	case SashimiParserLOOP:
		{
			p.SetState(161)
			p.LoopCall()
		}

	case SashimiParserENTITY, SashimiParserUNIQUE:
		{
			p.SetState(162)
			p.EntityDef()
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
	p.EnterRule(localctx, 34, SashimiParserRULE_block)
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
	p.SetState(168)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == SashimiParserDIRECTIVE {
		{
			p.SetState(165)
			p.Export()
		}

		p.SetState(170)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(171)
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
	_startState := 36
	p.EnterRecursionRule(localctx, 36, SashimiParserRULE_boolExpression, _p)

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
	p.SetState(184)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SashimiParserLPAREN:
		{
			p.SetState(174)
			p.Match(SashimiParserLPAREN)
		}
		{
			p.SetState(175)
			p.boolExpression(0)
		}
		{
			p.SetState(176)
			p.Match(SashimiParserRPAREN)
		}

	case SashimiParserNOT:
		{
			p.SetState(178)
			p.Match(SashimiParserNOT)
		}
		{
			p.SetState(179)
			p.boolExpression(6)
		}

	case SashimiParserTRUE, SashimiParserFALSE:
		{
			p.SetState(180)
			p.Truth()
		}

	case SashimiParserIDENT:
		{
			p.SetState(181)
			p.Qualifier()
		}

	case SashimiParserALIAS:
		{
			p.SetState(182)
			p.Match(SashimiParserALIAS)
		}

	case SashimiParserNUMBER:
		{
			p.SetState(183)
			p.Match(SashimiParserNUMBER)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(192)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 16, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewBoolExpressionContext(p, _parentctx, _parentState)
			localctx.(*BoolExpressionContext).left = _prevctx
			p.PushNewRecursionContext(localctx, _startState, SashimiParserRULE_boolExpression)
			p.SetState(186)

			if !(p.Precpred(p.GetParserRuleContext(), 5)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
			}
			{
				p.SetState(187)
				p.Op()
			}
			{
				p.SetState(188)

				var _x = p.boolExpression(6)

				localctx.(*BoolExpressionContext).right = _x
			}

		}
		p.SetState(194)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 16, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 38, SashimiParserRULE_op)

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

	p.SetState(197)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SashimiParserAND, SashimiParserOR:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(195)
			p.Binary()
		}

	case SashimiParserEQ, SashimiParserLEQ, SashimiParserLT, SashimiParserGEQ, SashimiParserGT, SashimiParserNEQ:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(196)
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
	p.EnterRule(localctx, 40, SashimiParserRULE_comparator)
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
		p.SetState(199)
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
	p.EnterRule(localctx, 42, SashimiParserRULE_binary)
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
	p.EnterRule(localctx, 44, SashimiParserRULE_truth)
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
		p.SetState(203)
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
	case 18:
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
