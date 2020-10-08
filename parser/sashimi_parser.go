// Code generated from Sashimi.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // Sashimi

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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 23, 136,
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
	12, 16, 14, 16, 134, 11, 16, 3, 16, 2, 2, 17, 2, 4, 6, 8, 10, 12, 14, 16,
	18, 20, 22, 24, 26, 28, 30, 2, 2, 2, 134, 2, 32, 3, 2, 2, 2, 4, 36, 3,
	2, 2, 2, 6, 39, 3, 2, 2, 2, 8, 56, 3, 2, 2, 2, 10, 68, 3, 2, 2, 2, 12,
	72, 3, 2, 2, 2, 14, 75, 3, 2, 2, 2, 16, 84, 3, 2, 2, 2, 18, 86, 3, 2, 2,
	2, 20, 89, 3, 2, 2, 2, 22, 92, 3, 2, 2, 2, 24, 99, 3, 2, 2, 2, 26, 111,
	3, 2, 2, 2, 28, 123, 3, 2, 2, 2, 30, 128, 3, 2, 2, 2, 32, 33, 7, 22, 2,
	2, 33, 34, 7, 15, 2, 2, 34, 35, 7, 22, 2, 2, 35, 3, 3, 2, 2, 2, 36, 37,
	7, 16, 2, 2, 37, 38, 7, 22, 2, 2, 38, 5, 3, 2, 2, 2, 39, 42, 7, 13, 2,
	2, 40, 43, 5, 4, 3, 2, 41, 43, 5, 2, 2, 2, 42, 40, 3, 2, 2, 2, 42, 41,
	3, 2, 2, 2, 43, 51, 3, 2, 2, 2, 44, 47, 7, 6, 2, 2, 45, 48, 5, 4, 3, 2,
	46, 48, 5, 2, 2, 2, 47, 45, 3, 2, 2, 2, 47, 46, 3, 2, 2, 2, 48, 50, 3,
	2, 2, 2, 49, 44, 3, 2, 2, 2, 50, 53, 3, 2, 2, 2, 51, 49, 3, 2, 2, 2, 51,
	52, 3, 2, 2, 2, 52, 54, 3, 2, 2, 2, 53, 51, 3, 2, 2, 2, 54, 55, 7, 14,
	2, 2, 55, 7, 3, 2, 2, 2, 56, 57, 7, 19, 2, 2, 57, 58, 7, 11, 2, 2, 58,
	63, 7, 21, 2, 2, 59, 60, 7, 6, 2, 2, 60, 62, 7, 21, 2, 2, 61, 59, 3, 2,
	2, 2, 62, 65, 3, 2, 2, 2, 63, 61, 3, 2, 2, 2, 63, 64, 3, 2, 2, 2, 64, 66,
	3, 2, 2, 2, 65, 63, 3, 2, 2, 2, 66, 67, 7, 12, 2, 2, 67, 9, 3, 2, 2, 2,
	68, 70, 7, 18, 2, 2, 69, 71, 5, 6, 4, 2, 70, 69, 3, 2, 2, 2, 70, 71, 3,
	2, 2, 2, 71, 11, 3, 2, 2, 2, 72, 73, 7, 5, 2, 2, 73, 74, 7, 22, 2, 2, 74,
	13, 3, 2, 2, 2, 75, 78, 7, 20, 2, 2, 76, 79, 5, 10, 6, 2, 77, 79, 5, 12,
	7, 2, 78, 76, 3, 2, 2, 2, 78, 77, 3, 2, 2, 2, 79, 15, 3, 2, 2, 2, 80, 85,
	5, 14, 8, 2, 81, 85, 5, 8, 5, 2, 82, 85, 5, 10, 6, 2, 83, 85, 5, 12, 7,
	2, 84, 80, 3, 2, 2, 2, 84, 81, 3, 2, 2, 2, 84, 82, 3, 2, 2, 2, 84, 83,
	3, 2, 2, 2, 85, 17, 3, 2, 2, 2, 86, 87, 7, 10, 2, 2, 87, 88, 7, 21, 2,
	2, 88, 19, 3, 2, 2, 2, 89, 90, 7, 9, 2, 2, 90, 91, 5, 16, 9, 2, 91, 21,
	3, 2, 2, 2, 92, 93, 7, 7, 2, 2, 93, 95, 7, 22, 2, 2, 94, 96, 5, 18, 10,
	2, 95, 94, 3, 2, 2, 2, 95, 96, 3, 2, 2, 2, 96, 97, 3, 2, 2, 2, 97, 98,
	5, 20, 11, 2, 98, 23, 3, 2, 2, 2, 99, 100, 7, 4, 2, 2, 100, 101, 7, 11,
	2, 2, 101, 106, 7, 22, 2, 2, 102, 103, 7, 17, 2, 2, 103, 105, 7, 22, 2,
	2, 104, 102, 3, 2, 2, 2, 105, 108, 3, 2, 2, 2, 106, 104, 3, 2, 2, 2, 106,
	107, 3, 2, 2, 2, 107, 109, 3, 2, 2, 2, 108, 106, 3, 2, 2, 2, 109, 110,
	7, 12, 2, 2, 110, 25, 3, 2, 2, 2, 111, 112, 7, 5, 2, 2, 112, 113, 7, 11,
	2, 2, 113, 114, 7, 22, 2, 2, 114, 115, 7, 12, 2, 2, 115, 116, 7, 8, 2,
	2, 116, 120, 5, 22, 12, 2, 117, 119, 5, 22, 12, 2, 118, 117, 3, 2, 2, 2,
	119, 122, 3, 2, 2, 2, 120, 118, 3, 2, 2, 2, 120, 121, 3, 2, 2, 2, 121,
	27, 3, 2, 2, 2, 122, 120, 3, 2, 2, 2, 123, 126, 7, 3, 2, 2, 124, 127, 5,
	24, 13, 2, 125, 127, 5, 26, 14, 2, 126, 124, 3, 2, 2, 2, 126, 125, 3, 2,
	2, 2, 127, 29, 3, 2, 2, 2, 128, 132, 5, 28, 15, 2, 129, 131, 5, 28, 15,
	2, 130, 129, 3, 2, 2, 2, 131, 134, 3, 2, 2, 2, 132, 130, 3, 2, 2, 2, 132,
	133, 3, 2, 2, 2, 133, 31, 3, 2, 2, 2, 134, 132, 3, 2, 2, 2, 14, 42, 47,
	51, 63, 70, 78, 84, 95, 106, 120, 126, 132,
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
	"key_value_pair", "key_atom", "constraint_list", "union_decl", "type_decl",
	"entity_ref", "list_decl", "type_def", "alias_decl", "type_is", "prop_decl",
	"command_call", "entity_def", "export", "block",
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
	SashimiParserRULE_key_value_pair  = 0
	SashimiParserRULE_key_atom        = 1
	SashimiParserRULE_constraint_list = 2
	SashimiParserRULE_union_decl      = 3
	SashimiParserRULE_type_decl       = 4
	SashimiParserRULE_entity_ref      = 5
	SashimiParserRULE_list_decl       = 6
	SashimiParserRULE_type_def        = 7
	SashimiParserRULE_alias_decl      = 8
	SashimiParserRULE_type_is         = 9
	SashimiParserRULE_prop_decl       = 10
	SashimiParserRULE_command_call    = 11
	SashimiParserRULE_entity_def      = 12
	SashimiParserRULE_export          = 13
	SashimiParserRULE_block           = 14
)

// IKey_value_pairContext is an interface to support dynamic dispatch.
type IKey_value_pairContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsKey_value_pairContext differentiates from other interfaces.
	IsKey_value_pairContext()
}

type Key_value_pairContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyKey_value_pairContext() *Key_value_pairContext {
	var p = new(Key_value_pairContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SashimiParserRULE_key_value_pair
	return p
}

func (*Key_value_pairContext) IsKey_value_pairContext() {}

func NewKey_value_pairContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Key_value_pairContext {
	var p = new(Key_value_pairContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SashimiParserRULE_key_value_pair

	return p
}

func (s *Key_value_pairContext) GetParser() antlr.Parser { return s.parser }

func (s *Key_value_pairContext) AllIDENT() []antlr.TerminalNode {
	return s.GetTokens(SashimiParserIDENT)
}

func (s *Key_value_pairContext) IDENT(i int) antlr.TerminalNode {
	return s.GetToken(SashimiParserIDENT, i)
}

func (s *Key_value_pairContext) EQ() antlr.TerminalNode {
	return s.GetToken(SashimiParserEQ, 0)
}

func (s *Key_value_pairContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Key_value_pairContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Key_value_pairContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SashimiListener); ok {
		listenerT.EnterKey_value_pair(s)
	}
}

func (s *Key_value_pairContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SashimiListener); ok {
		listenerT.ExitKey_value_pair(s)
	}
}

func (p *SashimiParser) Key_value_pair() (localctx IKey_value_pairContext) {
	localctx = NewKey_value_pairContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, SashimiParserRULE_key_value_pair)

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

// IKey_atomContext is an interface to support dynamic dispatch.
type IKey_atomContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsKey_atomContext differentiates from other interfaces.
	IsKey_atomContext()
}

type Key_atomContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyKey_atomContext() *Key_atomContext {
	var p = new(Key_atomContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SashimiParserRULE_key_atom
	return p
}

func (*Key_atomContext) IsKey_atomContext() {}

func NewKey_atomContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Key_atomContext {
	var p = new(Key_atomContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SashimiParserRULE_key_atom

	return p
}

func (s *Key_atomContext) GetParser() antlr.Parser { return s.parser }

func (s *Key_atomContext) ATOM() antlr.TerminalNode {
	return s.GetToken(SashimiParserATOM, 0)
}

func (s *Key_atomContext) IDENT() antlr.TerminalNode {
	return s.GetToken(SashimiParserIDENT, 0)
}

func (s *Key_atomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Key_atomContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Key_atomContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SashimiListener); ok {
		listenerT.EnterKey_atom(s)
	}
}

func (s *Key_atomContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SashimiListener); ok {
		listenerT.ExitKey_atom(s)
	}
}

func (p *SashimiParser) Key_atom() (localctx IKey_atomContext) {
	localctx = NewKey_atomContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, SashimiParserRULE_key_atom)

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

// IConstraint_listContext is an interface to support dynamic dispatch.
type IConstraint_listContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConstraint_listContext differentiates from other interfaces.
	IsConstraint_listContext()
}

type Constraint_listContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConstraint_listContext() *Constraint_listContext {
	var p = new(Constraint_listContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SashimiParserRULE_constraint_list
	return p
}

func (*Constraint_listContext) IsConstraint_listContext() {}

func NewConstraint_listContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Constraint_listContext {
	var p = new(Constraint_listContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SashimiParserRULE_constraint_list

	return p
}

func (s *Constraint_listContext) GetParser() antlr.Parser { return s.parser }

func (s *Constraint_listContext) HLPAREN() antlr.TerminalNode {
	return s.GetToken(SashimiParserHLPAREN, 0)
}

func (s *Constraint_listContext) HRPAREN() antlr.TerminalNode {
	return s.GetToken(SashimiParserHRPAREN, 0)
}

func (s *Constraint_listContext) AllKey_atom() []IKey_atomContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IKey_atomContext)(nil)).Elem())
	var tst = make([]IKey_atomContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IKey_atomContext)
		}
	}

	return tst
}

func (s *Constraint_listContext) Key_atom(i int) IKey_atomContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IKey_atomContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IKey_atomContext)
}

func (s *Constraint_listContext) AllKey_value_pair() []IKey_value_pairContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IKey_value_pairContext)(nil)).Elem())
	var tst = make([]IKey_value_pairContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IKey_value_pairContext)
		}
	}

	return tst
}

func (s *Constraint_listContext) Key_value_pair(i int) IKey_value_pairContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IKey_value_pairContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IKey_value_pairContext)
}

func (s *Constraint_listContext) AllSEPERATOR() []antlr.TerminalNode {
	return s.GetTokens(SashimiParserSEPERATOR)
}

func (s *Constraint_listContext) SEPERATOR(i int) antlr.TerminalNode {
	return s.GetToken(SashimiParserSEPERATOR, i)
}

func (s *Constraint_listContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Constraint_listContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Constraint_listContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SashimiListener); ok {
		listenerT.EnterConstraint_list(s)
	}
}

func (s *Constraint_listContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SashimiListener); ok {
		listenerT.ExitConstraint_list(s)
	}
}

func (p *SashimiParser) Constraint_list() (localctx IConstraint_listContext) {
	localctx = NewConstraint_listContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, SashimiParserRULE_constraint_list)
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
			p.Key_atom()
		}

	case SashimiParserIDENT:
		{
			p.SetState(39)
			p.Key_value_pair()
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
				p.Key_atom()
			}

		case SashimiParserIDENT:
			{
				p.SetState(44)
				p.Key_value_pair()
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

// IUnion_declContext is an interface to support dynamic dispatch.
type IUnion_declContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUnion_declContext differentiates from other interfaces.
	IsUnion_declContext()
}

type Union_declContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnion_declContext() *Union_declContext {
	var p = new(Union_declContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SashimiParserRULE_union_decl
	return p
}

func (*Union_declContext) IsUnion_declContext() {}

func NewUnion_declContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Union_declContext {
	var p = new(Union_declContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SashimiParserRULE_union_decl

	return p
}

func (s *Union_declContext) GetParser() antlr.Parser { return s.parser }

func (s *Union_declContext) UNION() antlr.TerminalNode {
	return s.GetToken(SashimiParserUNION, 0)
}

func (s *Union_declContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(SashimiParserLPAREN, 0)
}

func (s *Union_declContext) AllALIAS() []antlr.TerminalNode {
	return s.GetTokens(SashimiParserALIAS)
}

func (s *Union_declContext) ALIAS(i int) antlr.TerminalNode {
	return s.GetToken(SashimiParserALIAS, i)
}

func (s *Union_declContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(SashimiParserRPAREN, 0)
}

func (s *Union_declContext) AllSEPERATOR() []antlr.TerminalNode {
	return s.GetTokens(SashimiParserSEPERATOR)
}

func (s *Union_declContext) SEPERATOR(i int) antlr.TerminalNode {
	return s.GetToken(SashimiParserSEPERATOR, i)
}

func (s *Union_declContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Union_declContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Union_declContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SashimiListener); ok {
		listenerT.EnterUnion_decl(s)
	}
}

func (s *Union_declContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SashimiListener); ok {
		listenerT.ExitUnion_decl(s)
	}
}

func (p *SashimiParser) Union_decl() (localctx IUnion_declContext) {
	localctx = NewUnion_declContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, SashimiParserRULE_union_decl)
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

// IType_declContext is an interface to support dynamic dispatch.
type IType_declContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsType_declContext differentiates from other interfaces.
	IsType_declContext()
}

type Type_declContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyType_declContext() *Type_declContext {
	var p = new(Type_declContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SashimiParserRULE_type_decl
	return p
}

func (*Type_declContext) IsType_declContext() {}

func NewType_declContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Type_declContext {
	var p = new(Type_declContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SashimiParserRULE_type_decl

	return p
}

func (s *Type_declContext) GetParser() antlr.Parser { return s.parser }

func (s *Type_declContext) TYPE() antlr.TerminalNode {
	return s.GetToken(SashimiParserTYPE, 0)
}

func (s *Type_declContext) Constraint_list() IConstraint_listContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConstraint_listContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConstraint_listContext)
}

func (s *Type_declContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Type_declContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Type_declContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SashimiListener); ok {
		listenerT.EnterType_decl(s)
	}
}

func (s *Type_declContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SashimiListener); ok {
		listenerT.ExitType_decl(s)
	}
}

func (p *SashimiParser) Type_decl() (localctx IType_declContext) {
	localctx = NewType_declContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, SashimiParserRULE_type_decl)
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
			p.Constraint_list()
		}

	}

	return localctx
}

// IEntity_refContext is an interface to support dynamic dispatch.
type IEntity_refContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEntity_refContext differentiates from other interfaces.
	IsEntity_refContext()
}

type Entity_refContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEntity_refContext() *Entity_refContext {
	var p = new(Entity_refContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SashimiParserRULE_entity_ref
	return p
}

func (*Entity_refContext) IsEntity_refContext() {}

func NewEntity_refContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Entity_refContext {
	var p = new(Entity_refContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SashimiParserRULE_entity_ref

	return p
}

func (s *Entity_refContext) GetParser() antlr.Parser { return s.parser }

func (s *Entity_refContext) ENTITY() antlr.TerminalNode {
	return s.GetToken(SashimiParserENTITY, 0)
}

func (s *Entity_refContext) IDENT() antlr.TerminalNode {
	return s.GetToken(SashimiParserIDENT, 0)
}

func (s *Entity_refContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Entity_refContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Entity_refContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SashimiListener); ok {
		listenerT.EnterEntity_ref(s)
	}
}

func (s *Entity_refContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SashimiListener); ok {
		listenerT.ExitEntity_ref(s)
	}
}

func (p *SashimiParser) Entity_ref() (localctx IEntity_refContext) {
	localctx = NewEntity_refContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, SashimiParserRULE_entity_ref)

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

// IList_declContext is an interface to support dynamic dispatch.
type IList_declContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsList_declContext differentiates from other interfaces.
	IsList_declContext()
}

type List_declContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyList_declContext() *List_declContext {
	var p = new(List_declContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SashimiParserRULE_list_decl
	return p
}

func (*List_declContext) IsList_declContext() {}

func NewList_declContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *List_declContext {
	var p = new(List_declContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SashimiParserRULE_list_decl

	return p
}

func (s *List_declContext) GetParser() antlr.Parser { return s.parser }

func (s *List_declContext) LIST() antlr.TerminalNode {
	return s.GetToken(SashimiParserLIST, 0)
}

func (s *List_declContext) Type_decl() IType_declContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IType_declContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IType_declContext)
}

func (s *List_declContext) Entity_ref() IEntity_refContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEntity_refContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEntity_refContext)
}

func (s *List_declContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *List_declContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *List_declContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SashimiListener); ok {
		listenerT.EnterList_decl(s)
	}
}

func (s *List_declContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SashimiListener); ok {
		listenerT.ExitList_decl(s)
	}
}

func (p *SashimiParser) List_decl() (localctx IList_declContext) {
	localctx = NewList_declContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, SashimiParserRULE_list_decl)

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
			p.Type_decl()
		}

	case SashimiParserENTITY:
		{
			p.SetState(75)
			p.Entity_ref()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IType_defContext is an interface to support dynamic dispatch.
type IType_defContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsType_defContext differentiates from other interfaces.
	IsType_defContext()
}

type Type_defContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyType_defContext() *Type_defContext {
	var p = new(Type_defContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SashimiParserRULE_type_def
	return p
}

func (*Type_defContext) IsType_defContext() {}

func NewType_defContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Type_defContext {
	var p = new(Type_defContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SashimiParserRULE_type_def

	return p
}

func (s *Type_defContext) GetParser() antlr.Parser { return s.parser }

func (s *Type_defContext) List_decl() IList_declContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IList_declContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IList_declContext)
}

func (s *Type_defContext) Union_decl() IUnion_declContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUnion_declContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUnion_declContext)
}

func (s *Type_defContext) Type_decl() IType_declContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IType_declContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IType_declContext)
}

func (s *Type_defContext) Entity_ref() IEntity_refContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEntity_refContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEntity_refContext)
}

func (s *Type_defContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Type_defContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Type_defContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SashimiListener); ok {
		listenerT.EnterType_def(s)
	}
}

func (s *Type_defContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SashimiListener); ok {
		listenerT.ExitType_def(s)
	}
}

func (p *SashimiParser) Type_def() (localctx IType_defContext) {
	localctx = NewType_defContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, SashimiParserRULE_type_def)

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
			p.List_decl()
		}

	case SashimiParserUNION:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(79)
			p.Union_decl()
		}

	case SashimiParserTYPE:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(80)
			p.Type_decl()
		}

	case SashimiParserENTITY:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(81)
			p.Entity_ref()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IAlias_declContext is an interface to support dynamic dispatch.
type IAlias_declContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAlias_declContext differentiates from other interfaces.
	IsAlias_declContext()
}

type Alias_declContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAlias_declContext() *Alias_declContext {
	var p = new(Alias_declContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SashimiParserRULE_alias_decl
	return p
}

func (*Alias_declContext) IsAlias_declContext() {}

func NewAlias_declContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Alias_declContext {
	var p = new(Alias_declContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SashimiParserRULE_alias_decl

	return p
}

func (s *Alias_declContext) GetParser() antlr.Parser { return s.parser }

func (s *Alias_declContext) AS() antlr.TerminalNode {
	return s.GetToken(SashimiParserAS, 0)
}

func (s *Alias_declContext) ALIAS() antlr.TerminalNode {
	return s.GetToken(SashimiParserALIAS, 0)
}

func (s *Alias_declContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Alias_declContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Alias_declContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SashimiListener); ok {
		listenerT.EnterAlias_decl(s)
	}
}

func (s *Alias_declContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SashimiListener); ok {
		listenerT.ExitAlias_decl(s)
	}
}

func (p *SashimiParser) Alias_decl() (localctx IAlias_declContext) {
	localctx = NewAlias_declContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, SashimiParserRULE_alias_decl)

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

// IType_isContext is an interface to support dynamic dispatch.
type IType_isContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsType_isContext differentiates from other interfaces.
	IsType_isContext()
}

type Type_isContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyType_isContext() *Type_isContext {
	var p = new(Type_isContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SashimiParserRULE_type_is
	return p
}

func (*Type_isContext) IsType_isContext() {}

func NewType_isContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Type_isContext {
	var p = new(Type_isContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SashimiParserRULE_type_is

	return p
}

func (s *Type_isContext) GetParser() antlr.Parser { return s.parser }

func (s *Type_isContext) IS() antlr.TerminalNode {
	return s.GetToken(SashimiParserIS, 0)
}

func (s *Type_isContext) Type_def() IType_defContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IType_defContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IType_defContext)
}

func (s *Type_isContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Type_isContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Type_isContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SashimiListener); ok {
		listenerT.EnterType_is(s)
	}
}

func (s *Type_isContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SashimiListener); ok {
		listenerT.ExitType_is(s)
	}
}

func (p *SashimiParser) Type_is() (localctx IType_isContext) {
	localctx = NewType_isContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, SashimiParserRULE_type_is)

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
		p.Type_def()
	}

	return localctx
}

// IProp_declContext is an interface to support dynamic dispatch.
type IProp_declContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsProp_declContext differentiates from other interfaces.
	IsProp_declContext()
}

type Prop_declContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProp_declContext() *Prop_declContext {
	var p = new(Prop_declContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SashimiParserRULE_prop_decl
	return p
}

func (*Prop_declContext) IsProp_declContext() {}

func NewProp_declContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Prop_declContext {
	var p = new(Prop_declContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SashimiParserRULE_prop_decl

	return p
}

func (s *Prop_declContext) GetParser() antlr.Parser { return s.parser }

func (s *Prop_declContext) PROP_START() antlr.TerminalNode {
	return s.GetToken(SashimiParserPROP_START, 0)
}

func (s *Prop_declContext) IDENT() antlr.TerminalNode {
	return s.GetToken(SashimiParserIDENT, 0)
}

func (s *Prop_declContext) Type_is() IType_isContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IType_isContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IType_isContext)
}

func (s *Prop_declContext) Alias_decl() IAlias_declContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAlias_declContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAlias_declContext)
}

func (s *Prop_declContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Prop_declContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Prop_declContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SashimiListener); ok {
		listenerT.EnterProp_decl(s)
	}
}

func (s *Prop_declContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SashimiListener); ok {
		listenerT.ExitProp_decl(s)
	}
}

func (p *SashimiParser) Prop_decl() (localctx IProp_declContext) {
	localctx = NewProp_declContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, SashimiParserRULE_prop_decl)
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
			p.Alias_decl()
		}

	}
	{
		p.SetState(95)
		p.Type_is()
	}

	return localctx
}

// ICommand_callContext is an interface to support dynamic dispatch.
type ICommand_callContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCommand_callContext differentiates from other interfaces.
	IsCommand_callContext()
}

type Command_callContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCommand_callContext() *Command_callContext {
	var p = new(Command_callContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SashimiParserRULE_command_call
	return p
}

func (*Command_callContext) IsCommand_callContext() {}

func NewCommand_callContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Command_callContext {
	var p = new(Command_callContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SashimiParserRULE_command_call

	return p
}

func (s *Command_callContext) GetParser() antlr.Parser { return s.parser }

func (s *Command_callContext) COMMAND() antlr.TerminalNode {
	return s.GetToken(SashimiParserCOMMAND, 0)
}

func (s *Command_callContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(SashimiParserLPAREN, 0)
}

func (s *Command_callContext) AllIDENT() []antlr.TerminalNode {
	return s.GetTokens(SashimiParserIDENT)
}

func (s *Command_callContext) IDENT(i int) antlr.TerminalNode {
	return s.GetToken(SashimiParserIDENT, i)
}

func (s *Command_callContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(SashimiParserRPAREN, 0)
}

func (s *Command_callContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(SashimiParserDOT)
}

func (s *Command_callContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(SashimiParserDOT, i)
}

func (s *Command_callContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Command_callContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Command_callContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SashimiListener); ok {
		listenerT.EnterCommand_call(s)
	}
}

func (s *Command_callContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SashimiListener); ok {
		listenerT.ExitCommand_call(s)
	}
}

func (p *SashimiParser) Command_call() (localctx ICommand_callContext) {
	localctx = NewCommand_callContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, SashimiParserRULE_command_call)
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

// IEntity_defContext is an interface to support dynamic dispatch.
type IEntity_defContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEntity_defContext differentiates from other interfaces.
	IsEntity_defContext()
}

type Entity_defContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEntity_defContext() *Entity_defContext {
	var p = new(Entity_defContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SashimiParserRULE_entity_def
	return p
}

func (*Entity_defContext) IsEntity_defContext() {}

func NewEntity_defContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Entity_defContext {
	var p = new(Entity_defContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SashimiParserRULE_entity_def

	return p
}

func (s *Entity_defContext) GetParser() antlr.Parser { return s.parser }

func (s *Entity_defContext) ENTITY() antlr.TerminalNode {
	return s.GetToken(SashimiParserENTITY, 0)
}

func (s *Entity_defContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(SashimiParserLPAREN, 0)
}

func (s *Entity_defContext) IDENT() antlr.TerminalNode {
	return s.GetToken(SashimiParserIDENT, 0)
}

func (s *Entity_defContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(SashimiParserRPAREN, 0)
}

func (s *Entity_defContext) OF() antlr.TerminalNode {
	return s.GetToken(SashimiParserOF, 0)
}

func (s *Entity_defContext) AllProp_decl() []IProp_declContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IProp_declContext)(nil)).Elem())
	var tst = make([]IProp_declContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IProp_declContext)
		}
	}

	return tst
}

func (s *Entity_defContext) Prop_decl(i int) IProp_declContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IProp_declContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IProp_declContext)
}

func (s *Entity_defContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Entity_defContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Entity_defContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SashimiListener); ok {
		listenerT.EnterEntity_def(s)
	}
}

func (s *Entity_defContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SashimiListener); ok {
		listenerT.ExitEntity_def(s)
	}
}

func (p *SashimiParser) Entity_def() (localctx IEntity_defContext) {
	localctx = NewEntity_defContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, SashimiParserRULE_entity_def)
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
		p.Prop_decl()
	}
	p.SetState(118)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == SashimiParserPROP_START {
		{
			p.SetState(115)
			p.Prop_decl()
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

func (s *ExportContext) Command_call() ICommand_callContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICommand_callContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICommand_callContext)
}

func (s *ExportContext) Entity_def() IEntity_defContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEntity_defContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEntity_defContext)
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
			p.Command_call()
		}

	case SashimiParserENTITY:
		{
			p.SetState(123)
			p.Entity_def()
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

	return localctx
}
