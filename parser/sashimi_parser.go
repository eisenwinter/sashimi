// Code generated from sashimi.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // sashimi

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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 21, 84, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 4,
	8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2,
	3, 2, 3, 2, 3, 2, 7, 2, 30, 10, 2, 12, 2, 14, 2, 33, 11, 2, 3, 2, 3, 2,
	3, 3, 3, 3, 3, 3, 3, 3, 5, 3, 41, 10, 3, 3, 4, 3, 4, 3, 4, 3, 5, 3, 5,
	3, 5, 5, 5, 49, 10, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7,
	3, 8, 3, 8, 3, 8, 5, 8, 62, 10, 8, 3, 8, 3, 8, 3, 8, 3, 9, 3, 9, 6, 9,
	69, 10, 9, 13, 9, 14, 9, 70, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10,
	7, 10, 79, 10, 10, 12, 10, 14, 10, 82, 11, 10, 3, 10, 2, 2, 11, 2, 4, 6,
	8, 10, 12, 14, 16, 18, 2, 2, 2, 81, 2, 20, 3, 2, 2, 2, 4, 36, 3, 2, 2,
	2, 6, 42, 3, 2, 2, 2, 8, 48, 3, 2, 2, 2, 10, 50, 3, 2, 2, 2, 12, 55, 3,
	2, 2, 2, 14, 58, 3, 2, 2, 2, 16, 66, 3, 2, 2, 2, 18, 72, 3, 2, 2, 2, 20,
	21, 7, 20, 2, 2, 21, 22, 7, 3, 2, 2, 22, 23, 7, 4, 2, 2, 23, 24, 7, 16,
	2, 2, 24, 31, 7, 4, 2, 2, 25, 26, 7, 18, 2, 2, 26, 27, 7, 4, 2, 2, 27,
	28, 7, 16, 2, 2, 28, 30, 7, 4, 2, 2, 29, 25, 3, 2, 2, 2, 30, 33, 3, 2,
	2, 2, 31, 29, 3, 2, 2, 2, 31, 32, 3, 2, 2, 2, 32, 34, 3, 2, 2, 2, 33, 31,
	3, 2, 2, 2, 34, 35, 7, 5, 2, 2, 35, 3, 3, 2, 2, 2, 36, 40, 7, 19, 2, 2,
	37, 38, 7, 6, 2, 2, 38, 39, 7, 17, 2, 2, 39, 41, 7, 7, 2, 2, 40, 37, 3,
	2, 2, 2, 40, 41, 3, 2, 2, 2, 41, 5, 3, 2, 2, 2, 42, 43, 7, 8, 2, 2, 43,
	44, 5, 4, 3, 2, 44, 7, 3, 2, 2, 2, 45, 49, 5, 6, 4, 2, 46, 49, 5, 2, 2,
	2, 47, 49, 5, 4, 3, 2, 48, 45, 3, 2, 2, 2, 48, 46, 3, 2, 2, 2, 48, 47,
	3, 2, 2, 2, 49, 9, 3, 2, 2, 2, 50, 51, 7, 9, 2, 2, 51, 52, 7, 4, 2, 2,
	52, 53, 7, 16, 2, 2, 53, 54, 7, 4, 2, 2, 54, 11, 3, 2, 2, 2, 55, 56, 7,
	10, 2, 2, 56, 57, 5, 8, 5, 2, 57, 13, 3, 2, 2, 2, 58, 59, 7, 11, 2, 2,
	59, 61, 7, 15, 2, 2, 60, 62, 5, 10, 6, 2, 61, 60, 3, 2, 2, 2, 61, 62, 3,
	2, 2, 2, 62, 63, 3, 2, 2, 2, 63, 64, 7, 10, 2, 2, 64, 65, 5, 8, 5, 2, 65,
	15, 3, 2, 2, 2, 66, 68, 7, 12, 2, 2, 67, 69, 5, 14, 8, 2, 68, 67, 3, 2,
	2, 2, 69, 70, 3, 2, 2, 2, 70, 68, 3, 2, 2, 2, 70, 71, 3, 2, 2, 2, 71, 17,
	3, 2, 2, 2, 72, 73, 7, 13, 2, 2, 73, 74, 7, 14, 2, 2, 74, 75, 7, 3, 2,
	2, 75, 76, 7, 15, 2, 2, 76, 80, 7, 5, 2, 2, 77, 79, 5, 14, 8, 2, 78, 77,
	3, 2, 2, 2, 79, 82, 3, 2, 2, 2, 80, 78, 3, 2, 2, 2, 80, 81, 3, 2, 2, 2,
	81, 19, 3, 2, 2, 2, 82, 80, 3, 2, 2, 2, 8, 31, 40, 48, 61, 70, 80,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'('", "'\"'", "')'", "'['", "']'", "'list of'", "'as'", "'is'", "'-'",
	"'of'", "'sashimi:'", "", "", "", "", "','",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "DIRECTIVE", "COMMAND", "IDENT",
	"ALIAS", "CONSTRAINT", "SEPERATOR", "TYPE", "UNION", "WS",
}

var ruleNames = []string{
	"union_decl", "type_decl", "list_decl", "type_def", "alias_decl", "type_is",
	"prop_decl", "entity_decl", "export",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type sashimiParser struct {
	*antlr.BaseParser
}

func NewsashimiParser(input antlr.TokenStream) *sashimiParser {
	this := new(sashimiParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "sashimi.g4"

	return this
}

// sashimiParser tokens.
const (
	sashimiParserEOF        = antlr.TokenEOF
	sashimiParserT__0       = 1
	sashimiParserT__1       = 2
	sashimiParserT__2       = 3
	sashimiParserT__3       = 4
	sashimiParserT__4       = 5
	sashimiParserT__5       = 6
	sashimiParserT__6       = 7
	sashimiParserT__7       = 8
	sashimiParserT__8       = 9
	sashimiParserT__9       = 10
	sashimiParserDIRECTIVE  = 11
	sashimiParserCOMMAND    = 12
	sashimiParserIDENT      = 13
	sashimiParserALIAS      = 14
	sashimiParserCONSTRAINT = 15
	sashimiParserSEPERATOR  = 16
	sashimiParserTYPE       = 17
	sashimiParserUNION      = 18
	sashimiParserWS         = 19
)

// sashimiParser rules.
const (
	sashimiParserRULE_union_decl  = 0
	sashimiParserRULE_type_decl   = 1
	sashimiParserRULE_list_decl   = 2
	sashimiParserRULE_type_def    = 3
	sashimiParserRULE_alias_decl  = 4
	sashimiParserRULE_type_is     = 5
	sashimiParserRULE_prop_decl   = 6
	sashimiParserRULE_entity_decl = 7
	sashimiParserRULE_export      = 8
)

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
	p.RuleIndex = sashimiParserRULE_union_decl
	return p
}

func (*Union_declContext) IsUnion_declContext() {}

func NewUnion_declContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Union_declContext {
	var p = new(Union_declContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = sashimiParserRULE_union_decl

	return p
}

func (s *Union_declContext) GetParser() antlr.Parser { return s.parser }

func (s *Union_declContext) UNION() antlr.TerminalNode {
	return s.GetToken(sashimiParserUNION, 0)
}

func (s *Union_declContext) AllALIAS() []antlr.TerminalNode {
	return s.GetTokens(sashimiParserALIAS)
}

func (s *Union_declContext) ALIAS(i int) antlr.TerminalNode {
	return s.GetToken(sashimiParserALIAS, i)
}

func (s *Union_declContext) AllSEPERATOR() []antlr.TerminalNode {
	return s.GetTokens(sashimiParserSEPERATOR)
}

func (s *Union_declContext) SEPERATOR(i int) antlr.TerminalNode {
	return s.GetToken(sashimiParserSEPERATOR, i)
}

func (s *Union_declContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Union_declContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Union_declContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(sashimiListener); ok {
		listenerT.EnterUnion_decl(s)
	}
}

func (s *Union_declContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(sashimiListener); ok {
		listenerT.ExitUnion_decl(s)
	}
}

func (p *sashimiParser) Union_decl() (localctx IUnion_declContext) {
	localctx = NewUnion_declContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, sashimiParserRULE_union_decl)
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
		p.SetState(18)
		p.Match(sashimiParserUNION)
	}
	{
		p.SetState(19)
		p.Match(sashimiParserT__0)
	}
	{
		p.SetState(20)
		p.Match(sashimiParserT__1)
	}
	{
		p.SetState(21)
		p.Match(sashimiParserALIAS)
	}
	{
		p.SetState(22)
		p.Match(sashimiParserT__1)
	}
	p.SetState(29)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == sashimiParserSEPERATOR {
		{
			p.SetState(23)
			p.Match(sashimiParserSEPERATOR)
		}
		{
			p.SetState(24)
			p.Match(sashimiParserT__1)
		}
		{
			p.SetState(25)
			p.Match(sashimiParserALIAS)
		}
		{
			p.SetState(26)
			p.Match(sashimiParserT__1)
		}

		p.SetState(31)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(32)
		p.Match(sashimiParserT__2)
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
	p.RuleIndex = sashimiParserRULE_type_decl
	return p
}

func (*Type_declContext) IsType_declContext() {}

func NewType_declContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Type_declContext {
	var p = new(Type_declContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = sashimiParserRULE_type_decl

	return p
}

func (s *Type_declContext) GetParser() antlr.Parser { return s.parser }

func (s *Type_declContext) TYPE() antlr.TerminalNode {
	return s.GetToken(sashimiParserTYPE, 0)
}

func (s *Type_declContext) CONSTRAINT() antlr.TerminalNode {
	return s.GetToken(sashimiParserCONSTRAINT, 0)
}

func (s *Type_declContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Type_declContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Type_declContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(sashimiListener); ok {
		listenerT.EnterType_decl(s)
	}
}

func (s *Type_declContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(sashimiListener); ok {
		listenerT.ExitType_decl(s)
	}
}

func (p *sashimiParser) Type_decl() (localctx IType_declContext) {
	localctx = NewType_declContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, sashimiParserRULE_type_decl)
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
		p.SetState(34)
		p.Match(sashimiParserTYPE)
	}
	p.SetState(38)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == sashimiParserT__3 {
		{
			p.SetState(35)
			p.Match(sashimiParserT__3)
		}
		{
			p.SetState(36)
			p.Match(sashimiParserCONSTRAINT)
		}
		{
			p.SetState(37)
			p.Match(sashimiParserT__4)
		}

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
	p.RuleIndex = sashimiParserRULE_list_decl
	return p
}

func (*List_declContext) IsList_declContext() {}

func NewList_declContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *List_declContext {
	var p = new(List_declContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = sashimiParserRULE_list_decl

	return p
}

func (s *List_declContext) GetParser() antlr.Parser { return s.parser }

func (s *List_declContext) Type_decl() IType_declContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IType_declContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IType_declContext)
}

func (s *List_declContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *List_declContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *List_declContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(sashimiListener); ok {
		listenerT.EnterList_decl(s)
	}
}

func (s *List_declContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(sashimiListener); ok {
		listenerT.ExitList_decl(s)
	}
}

func (p *sashimiParser) List_decl() (localctx IList_declContext) {
	localctx = NewList_declContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, sashimiParserRULE_list_decl)

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
		p.SetState(40)
		p.Match(sashimiParserT__5)
	}
	{
		p.SetState(41)
		p.Type_decl()
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
	p.RuleIndex = sashimiParserRULE_type_def
	return p
}

func (*Type_defContext) IsType_defContext() {}

func NewType_defContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Type_defContext {
	var p = new(Type_defContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = sashimiParserRULE_type_def

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

func (s *Type_defContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Type_defContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Type_defContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(sashimiListener); ok {
		listenerT.EnterType_def(s)
	}
}

func (s *Type_defContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(sashimiListener); ok {
		listenerT.ExitType_def(s)
	}
}

func (p *sashimiParser) Type_def() (localctx IType_defContext) {
	localctx = NewType_defContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, sashimiParserRULE_type_def)

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

	p.SetState(46)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case sashimiParserT__5:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(43)
			p.List_decl()
		}

	case sashimiParserUNION:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(44)
			p.Union_decl()
		}

	case sashimiParserTYPE:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(45)
			p.Type_decl()
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
	p.RuleIndex = sashimiParserRULE_alias_decl
	return p
}

func (*Alias_declContext) IsAlias_declContext() {}

func NewAlias_declContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Alias_declContext {
	var p = new(Alias_declContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = sashimiParserRULE_alias_decl

	return p
}

func (s *Alias_declContext) GetParser() antlr.Parser { return s.parser }

func (s *Alias_declContext) ALIAS() antlr.TerminalNode {
	return s.GetToken(sashimiParserALIAS, 0)
}

func (s *Alias_declContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Alias_declContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Alias_declContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(sashimiListener); ok {
		listenerT.EnterAlias_decl(s)
	}
}

func (s *Alias_declContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(sashimiListener); ok {
		listenerT.ExitAlias_decl(s)
	}
}

func (p *sashimiParser) Alias_decl() (localctx IAlias_declContext) {
	localctx = NewAlias_declContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, sashimiParserRULE_alias_decl)

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
		p.SetState(48)
		p.Match(sashimiParserT__6)
	}
	{
		p.SetState(49)
		p.Match(sashimiParserT__1)
	}
	{
		p.SetState(50)
		p.Match(sashimiParserALIAS)
	}
	{
		p.SetState(51)
		p.Match(sashimiParserT__1)
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
	p.RuleIndex = sashimiParserRULE_type_is
	return p
}

func (*Type_isContext) IsType_isContext() {}

func NewType_isContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Type_isContext {
	var p = new(Type_isContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = sashimiParserRULE_type_is

	return p
}

func (s *Type_isContext) GetParser() antlr.Parser { return s.parser }

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
	if listenerT, ok := listener.(sashimiListener); ok {
		listenerT.EnterType_is(s)
	}
}

func (s *Type_isContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(sashimiListener); ok {
		listenerT.ExitType_is(s)
	}
}

func (p *sashimiParser) Type_is() (localctx IType_isContext) {
	localctx = NewType_isContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, sashimiParserRULE_type_is)

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
		p.Match(sashimiParserT__7)
	}
	{
		p.SetState(54)
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
	p.RuleIndex = sashimiParserRULE_prop_decl
	return p
}

func (*Prop_declContext) IsProp_declContext() {}

func NewProp_declContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Prop_declContext {
	var p = new(Prop_declContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = sashimiParserRULE_prop_decl

	return p
}

func (s *Prop_declContext) GetParser() antlr.Parser { return s.parser }

func (s *Prop_declContext) IDENT() antlr.TerminalNode {
	return s.GetToken(sashimiParserIDENT, 0)
}

func (s *Prop_declContext) Type_def() IType_defContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IType_defContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IType_defContext)
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
	if listenerT, ok := listener.(sashimiListener); ok {
		listenerT.EnterProp_decl(s)
	}
}

func (s *Prop_declContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(sashimiListener); ok {
		listenerT.ExitProp_decl(s)
	}
}

func (p *sashimiParser) Prop_decl() (localctx IProp_declContext) {
	localctx = NewProp_declContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, sashimiParserRULE_prop_decl)
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
		p.SetState(56)
		p.Match(sashimiParserT__8)
	}
	{
		p.SetState(57)
		p.Match(sashimiParserIDENT)
	}
	p.SetState(59)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == sashimiParserT__6 {
		{
			p.SetState(58)
			p.Alias_decl()
		}

	}
	{
		p.SetState(61)
		p.Match(sashimiParserT__7)
	}
	{
		p.SetState(62)
		p.Type_def()
	}

	return localctx
}

// IEntity_declContext is an interface to support dynamic dispatch.
type IEntity_declContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEntity_declContext differentiates from other interfaces.
	IsEntity_declContext()
}

type Entity_declContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEntity_declContext() *Entity_declContext {
	var p = new(Entity_declContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = sashimiParserRULE_entity_decl
	return p
}

func (*Entity_declContext) IsEntity_declContext() {}

func NewEntity_declContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Entity_declContext {
	var p = new(Entity_declContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = sashimiParserRULE_entity_decl

	return p
}

func (s *Entity_declContext) GetParser() antlr.Parser { return s.parser }

func (s *Entity_declContext) AllProp_decl() []IProp_declContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IProp_declContext)(nil)).Elem())
	var tst = make([]IProp_declContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IProp_declContext)
		}
	}

	return tst
}

func (s *Entity_declContext) Prop_decl(i int) IProp_declContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IProp_declContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IProp_declContext)
}

func (s *Entity_declContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Entity_declContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Entity_declContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(sashimiListener); ok {
		listenerT.EnterEntity_decl(s)
	}
}

func (s *Entity_declContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(sashimiListener); ok {
		listenerT.ExitEntity_decl(s)
	}
}

func (p *sashimiParser) Entity_decl() (localctx IEntity_declContext) {
	localctx = NewEntity_declContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, sashimiParserRULE_entity_decl)
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
		p.SetState(64)
		p.Match(sashimiParserT__9)
	}
	p.SetState(66)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == sashimiParserT__8 {
		{
			p.SetState(65)
			p.Prop_decl()
		}

		p.SetState(68)
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
	p.RuleIndex = sashimiParserRULE_export
	return p
}

func (*ExportContext) IsExportContext() {}

func NewExportContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExportContext {
	var p = new(ExportContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = sashimiParserRULE_export

	return p
}

func (s *ExportContext) GetParser() antlr.Parser { return s.parser }

func (s *ExportContext) DIRECTIVE() antlr.TerminalNode {
	return s.GetToken(sashimiParserDIRECTIVE, 0)
}

func (s *ExportContext) COMMAND() antlr.TerminalNode {
	return s.GetToken(sashimiParserCOMMAND, 0)
}

func (s *ExportContext) IDENT() antlr.TerminalNode {
	return s.GetToken(sashimiParserIDENT, 0)
}

func (s *ExportContext) AllProp_decl() []IProp_declContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IProp_declContext)(nil)).Elem())
	var tst = make([]IProp_declContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IProp_declContext)
		}
	}

	return tst
}

func (s *ExportContext) Prop_decl(i int) IProp_declContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IProp_declContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IProp_declContext)
}

func (s *ExportContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExportContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExportContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(sashimiListener); ok {
		listenerT.EnterExport(s)
	}
}

func (s *ExportContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(sashimiListener); ok {
		listenerT.ExitExport(s)
	}
}

func (p *sashimiParser) Export() (localctx IExportContext) {
	localctx = NewExportContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, sashimiParserRULE_export)
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
		p.Match(sashimiParserDIRECTIVE)
	}
	{
		p.SetState(71)
		p.Match(sashimiParserCOMMAND)
	}
	{
		p.SetState(72)
		p.Match(sashimiParserT__0)
	}
	{
		p.SetState(73)
		p.Match(sashimiParserIDENT)
	}
	{
		p.SetState(74)
		p.Match(sashimiParserT__2)
	}
	p.SetState(78)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == sashimiParserT__8 {
		{
			p.SetState(75)
			p.Prop_decl()
		}

		p.SetState(80)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}
