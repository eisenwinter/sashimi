// Code generated from C:\tmp\sashimi\sashimi\grammar\Sashimi.g4 by ANTLR 4.8. DO NOT EDIT.

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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 23, 138,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 3, 2, 3, 2, 3, 2, 3, 2,
	3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 5, 4, 43, 10, 4, 3, 4, 3, 4, 3, 4,
	5, 4, 48, 10, 4, 7, 4, 50, 10, 4, 12, 4, 14, 4, 53, 11, 4, 3, 4, 3, 4,
	3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 7, 5, 62, 10, 5, 12, 5, 14, 5, 65, 11, 5,
	3, 5, 3, 5, 3, 6, 3, 6, 5, 6, 71, 10, 6, 3, 7, 3, 7, 3, 7, 3, 8, 3, 8,
	3, 8, 5, 8, 79, 10, 8, 3, 9, 3, 9, 3, 9, 3, 9, 5, 9, 85, 10, 9, 3, 10,
	3, 10, 3, 10, 3, 11, 3, 11, 3, 11, 3, 12, 3, 12, 3, 12, 5, 12, 96, 10,
	12, 3, 12, 3, 12, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 7, 13, 105, 10, 13,
	12, 13, 14, 13, 108, 11, 13, 3, 13, 3, 13, 3, 14, 3, 14, 3, 14, 3, 14,
	3, 14, 3, 14, 3, 14, 7, 14, 119, 10, 14, 12, 14, 14, 14, 122, 11, 14, 3,
	15, 3, 15, 3, 15, 5, 15, 127, 10, 15, 3, 16, 3, 16, 7, 16, 131, 10, 16,
	12, 16, 14, 16, 134, 11, 16, 3, 16, 3, 16, 3, 16, 2, 2, 17, 2, 4, 6, 8,
	10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 2, 2, 2, 136, 2, 32, 3, 2,
	2, 2, 4, 36, 3, 2, 2, 2, 6, 39, 3, 2, 2, 2, 8, 56, 3, 2, 2, 2, 10, 68,
	3, 2, 2, 2, 12, 72, 3, 2, 2, 2, 14, 75, 3, 2, 2, 2, 16, 84, 3, 2, 2, 2,
	18, 86, 3, 2, 2, 2, 20, 89, 3, 2, 2, 2, 22, 92, 3, 2, 2, 2, 24, 99, 3,
	2, 2, 2, 26, 111, 3, 2, 2, 2, 28, 123, 3, 2, 2, 2, 30, 128, 3, 2, 2, 2,
	32, 33, 7, 22, 2, 2, 33, 34, 7, 15, 2, 2, 34, 35, 7, 22, 2, 2, 35, 3, 3,
	2, 2, 2, 36, 37, 7, 16, 2, 2, 37, 38, 7, 22, 2, 2, 38, 5, 3, 2, 2, 2, 39,
	42, 7, 13, 2, 2, 40, 43, 5, 4, 3, 2, 41, 43, 5, 2, 2, 2, 42, 40, 3, 2,
	2, 2, 42, 41, 3, 2, 2, 2, 43, 51, 3, 2, 2, 2, 44, 47, 7, 6, 2, 2, 45, 48,
	5, 4, 3, 2, 46, 48, 5, 2, 2, 2, 47, 45, 3, 2, 2, 2, 47, 46, 3, 2, 2, 2,
	48, 50, 3, 2, 2, 2, 49, 44, 3, 2, 2, 2, 50, 53, 3, 2, 2, 2, 51, 49, 3,
	2, 2, 2, 51, 52, 3, 2, 2, 2, 52, 54, 3, 2, 2, 2, 53, 51, 3, 2, 2, 2, 54,
	55, 7, 14, 2, 2, 55, 7, 3, 2, 2, 2, 56, 57, 7, 19, 2, 2, 57, 58, 7, 11,
	2, 2, 58, 63, 7, 21, 2, 2, 59, 60, 7, 6, 2, 2, 60, 62, 7, 21, 2, 2, 61,
	59, 3, 2, 2, 2, 62, 65, 3, 2, 2, 2, 63, 61, 3, 2, 2, 2, 63, 64, 3, 2, 2,
	2, 64, 66, 3, 2, 2, 2, 65, 63, 3, 2, 2, 2, 66, 67, 7, 12, 2, 2, 67, 9,
	3, 2, 2, 2, 68, 70, 7, 18, 2, 2, 69, 71, 5, 6, 4, 2, 70, 69, 3, 2, 2, 2,
	70, 71, 3, 2, 2, 2, 71, 11, 3, 2, 2, 2, 72, 73, 7, 5, 2, 2, 73, 74, 7,
	22, 2, 2, 74, 13, 3, 2, 2, 2, 75, 78, 7, 20, 2, 2, 76, 79, 5, 10, 6, 2,
	77, 79, 5, 12, 7, 2, 78, 76, 3, 2, 2, 2, 78, 77, 3, 2, 2, 2, 79, 15, 3,
	2, 2, 2, 80, 85, 5, 14, 8, 2, 81, 85, 5, 8, 5, 2, 82, 85, 5, 10, 6, 2,
	83, 85, 5, 12, 7, 2, 84, 80, 3, 2, 2, 2, 84, 81, 3, 2, 2, 2, 84, 82, 3,
	2, 2, 2, 84, 83, 3, 2, 2, 2, 85, 17, 3, 2, 2, 2, 86, 87, 7, 10, 2, 2, 87,
	88, 7, 21, 2, 2, 88, 19, 3, 2, 2, 2, 89, 90, 7, 9, 2, 2, 90, 91, 5, 16,
	9, 2, 91, 21, 3, 2, 2, 2, 92, 93, 7, 7, 2, 2, 93, 95, 7, 22, 2, 2, 94,
	96, 5, 18, 10, 2, 95, 94, 3, 2, 2, 2, 95, 96, 3, 2, 2, 2, 96, 97, 3, 2,
	2, 2, 97, 98, 5, 20, 11, 2, 98, 23, 3, 2, 2, 2, 99, 100, 7, 4, 2, 2, 100,
	101, 7, 11, 2, 2, 101, 106, 7, 22, 2, 2, 102, 103, 7, 17, 2, 2, 103, 105,
	7, 22, 2, 2, 104, 102, 3, 2, 2, 2, 105, 108, 3, 2, 2, 2, 106, 104, 3, 2,
	2, 2, 106, 107, 3, 2, 2, 2, 107, 109, 3, 2, 2, 2, 108, 106, 3, 2, 2, 2,
	109, 110, 7, 12, 2, 2, 110, 25, 3, 2, 2, 2, 111, 112, 7, 5, 2, 2, 112,
	113, 7, 11, 2, 2, 113, 114, 7, 22, 2, 2, 114, 115, 7, 12, 2, 2, 115, 116,
	7, 8, 2, 2, 116, 120, 5, 22, 12, 2, 117, 119, 5, 22, 12, 2, 118, 117, 3,
	2, 2, 2, 119, 122, 3, 2, 2, 2, 120, 118, 3, 2, 2, 2, 120, 121, 3, 2, 2,
	2, 121, 27, 3, 2, 2, 2, 122, 120, 3, 2, 2, 2, 123, 126, 7, 3, 2, 2, 124,
	127, 5, 24, 13, 2, 125, 127, 5, 26, 14, 2, 126, 124, 3, 2, 2, 2, 126, 125,
	3, 2, 2, 2, 127, 29, 3, 2, 2, 2, 128, 132, 5, 28, 15, 2, 129, 131, 5, 28,
	15, 2, 130, 129, 3, 2, 2, 2, 131, 134, 3, 2, 2, 2, 132, 130, 3, 2, 2, 2,
	132, 133, 3, 2, 2, 2, 133, 135, 3, 2, 2, 2, 134, 132, 3, 2, 2, 2, 135,
	136, 7, 2, 2, 3, 136, 31, 3, 2, 2, 2, 14, 42, 47, 51, 63, 70, 78, 84, 95,
	106, 120, 126, 132,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'sashimi:'", "", "'entity'", "','", "'-'", "'of'", "'is'", "'as'",
	"'('", "')'", "'['", "']'", "'='", "':'", "'.'", "", "", "'list'",
}
var symbolicNames = []string{
	"", "DIRECTIVE", "COMMAND", "ENTITY", "SEPERATOR", "PROP_START", "OF",
	"IS", "AS", "LPAREN", "RPAREN", "HLPAREN", "HRPAREN", "EQ", "ATOM", "DOT",
	"TYPE", "UNION", "LIST", "ALIAS", "IDENT", "WS",
}

var ruleNames = []string{
	"keyValuePair", "keyAtom", "constraintList", "unionDecl", "typeDecl", "entityRef",
	"listDecl", "typeDef", "aliasDecl", "typeIs", "propDecl", "commandCall",
	"entityDef", "export", "block",
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
	SashimiParserENTITY     = 3
	SashimiParserSEPERATOR  = 4
	SashimiParserPROP_START = 5
	SashimiParserOF         = 6
	SashimiParserIS         = 7
	SashimiParserAS         = 8
	SashimiParserLPAREN     = 9
	SashimiParserRPAREN     = 10
	SashimiParserHLPAREN    = 11
	SashimiParserHRPAREN    = 12
	SashimiParserEQ         = 13
	SashimiParserATOM       = 14
	SashimiParserDOT        = 15
	SashimiParserTYPE       = 16
	SashimiParserUNION      = 17
	SashimiParserLIST       = 18
	SashimiParserALIAS      = 19
	SashimiParserIDENT      = 20
	SashimiParserWS         = 21
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
	SashimiParserRULE_commandCall    = 11
	SashimiParserRULE_entityDef      = 12
	SashimiParserRULE_export         = 13
	SashimiParserRULE_block          = 14
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
		p.SetState(30)
		p.Match(SashimiParserIDENT)
	}
	{
		p.SetState(31)
		p.Match(SashimiParserEQ)
	}
	{
		p.SetState(32)
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
		p.SetState(34)
		p.Match(SashimiParserATOM)
	}
	{
		p.SetState(35)
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
		p.SetState(37)
		p.Match(SashimiParserHLPAREN)
	}
	p.SetState(40)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SashimiParserATOM:
		{
			p.SetState(38)
			p.KeyAtom()
		}

	case SashimiParserIDENT:
		{
			p.SetState(39)
			p.KeyValuePair()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.SetState(49)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == SashimiParserSEPERATOR {
		{
			p.SetState(42)
			p.Match(SashimiParserSEPERATOR)
		}
		p.SetState(45)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case SashimiParserATOM:
			{
				p.SetState(43)
				p.KeyAtom()
			}

		case SashimiParserIDENT:
			{
				p.SetState(44)
				p.KeyValuePair()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(51)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(52)
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
		p.SetState(54)
		p.Match(SashimiParserUNION)
	}
	{
		p.SetState(55)
		p.Match(SashimiParserLPAREN)
	}
	{
		p.SetState(56)
		p.Match(SashimiParserALIAS)
	}
	p.SetState(61)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == SashimiParserSEPERATOR {
		{
			p.SetState(57)
			p.Match(SashimiParserSEPERATOR)
		}
		{
			p.SetState(58)
			p.Match(SashimiParserALIAS)
		}

		p.SetState(63)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(64)
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
		p.SetState(66)
		p.Match(SashimiParserTYPE)
	}
	p.SetState(68)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SashimiParserHLPAREN {
		{
			p.SetState(67)
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
		p.SetState(70)
		p.Match(SashimiParserENTITY)
	}
	{
		p.SetState(71)
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
		p.SetState(73)
		p.Match(SashimiParserLIST)
	}
	p.SetState(76)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SashimiParserTYPE:
		{
			p.SetState(74)
			p.TypeDecl()
		}

	case SashimiParserENTITY:
		{
			p.SetState(75)
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

	p.SetState(82)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SashimiParserLIST:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(78)
			p.ListDecl()
		}

	case SashimiParserUNION:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(79)
			p.UnionDecl()
		}

	case SashimiParserTYPE:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(80)
			p.TypeDecl()
		}

	case SashimiParserENTITY:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(81)
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
		p.SetState(84)
		p.Match(SashimiParserAS)
	}
	{
		p.SetState(85)
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
		p.SetState(87)
		p.Match(SashimiParserIS)
	}
	{
		p.SetState(88)
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
		p.SetState(90)
		p.Match(SashimiParserPROP_START)
	}
	{
		p.SetState(91)
		p.Match(SashimiParserIDENT)
	}
	p.SetState(93)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SashimiParserAS {
		{
			p.SetState(92)
			p.AliasDecl()
		}

	}
	{
		p.SetState(95)
		p.TypeIs()
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

func (s *CommandCallContext) AllIDENT() []antlr.TerminalNode {
	return s.GetTokens(SashimiParserIDENT)
}

func (s *CommandCallContext) IDENT(i int) antlr.TerminalNode {
	return s.GetToken(SashimiParserIDENT, i)
}

func (s *CommandCallContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(SashimiParserRPAREN, 0)
}

func (s *CommandCallContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(SashimiParserDOT)
}

func (s *CommandCallContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(SashimiParserDOT, i)
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
	p.EnterRule(localctx, 22, SashimiParserRULE_commandCall)
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
		p.SetState(97)
		p.Match(SashimiParserCOMMAND)
	}
	{
		p.SetState(98)
		p.Match(SashimiParserLPAREN)
	}
	{
		p.SetState(99)
		p.Match(SashimiParserIDENT)
	}
	p.SetState(104)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == SashimiParserDOT {
		{
			p.SetState(100)
			p.Match(SashimiParserDOT)
		}
		{
			p.SetState(101)
			p.Match(SashimiParserIDENT)
		}

		p.SetState(106)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(107)
		p.Match(SashimiParserRPAREN)
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
	p.EnterRule(localctx, 24, SashimiParserRULE_entityDef)
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
		p.SetState(109)
		p.Match(SashimiParserENTITY)
	}
	{
		p.SetState(110)
		p.Match(SashimiParserLPAREN)
	}
	{
		p.SetState(111)
		p.Match(SashimiParserIDENT)
	}
	{
		p.SetState(112)
		p.Match(SashimiParserRPAREN)
	}
	{
		p.SetState(113)
		p.Match(SashimiParserOF)
	}
	{
		p.SetState(114)
		p.PropDecl()
	}
	p.SetState(118)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == SashimiParserPROP_START {
		{
			p.SetState(115)
			p.PropDecl()
		}

		p.SetState(120)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
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
	p.EnterRule(localctx, 26, SashimiParserRULE_export)

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
		p.SetState(121)
		p.Match(SashimiParserDIRECTIVE)
	}
	p.SetState(124)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SashimiParserCOMMAND:
		{
			p.SetState(122)
			p.CommandCall()
		}

	case SashimiParserENTITY:
		{
			p.SetState(123)
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

func (s *BlockContext) EOF() antlr.TerminalNode {
	return s.GetToken(SashimiParserEOF, 0)
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
	p.EnterRule(localctx, 28, SashimiParserRULE_block)
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
		p.Export()
	}
	p.SetState(130)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == SashimiParserDIRECTIVE {
		{
			p.SetState(127)
			p.Export()
		}

		p.SetState(132)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(133)
		p.Match(SashimiParserEOF)
	}

	return localctx
}
