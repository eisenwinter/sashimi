// Code generated from sashimi.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // sashimi

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasesashimiListener is a complete listener for a parse tree produced by sashimiParser.
type BasesashimiListener struct{}

var _ sashimiListener = &BasesashimiListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasesashimiListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasesashimiListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasesashimiListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasesashimiListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterUnion_decl is called when production union_decl is entered.
func (s *BasesashimiListener) EnterUnion_decl(ctx *Union_declContext) {}

// ExitUnion_decl is called when production union_decl is exited.
func (s *BasesashimiListener) ExitUnion_decl(ctx *Union_declContext) {}

// EnterType_decl is called when production type_decl is entered.
func (s *BasesashimiListener) EnterType_decl(ctx *Type_declContext) {}

// ExitType_decl is called when production type_decl is exited.
func (s *BasesashimiListener) ExitType_decl(ctx *Type_declContext) {}

// EnterList_decl is called when production list_decl is entered.
func (s *BasesashimiListener) EnterList_decl(ctx *List_declContext) {}

// ExitList_decl is called when production list_decl is exited.
func (s *BasesashimiListener) ExitList_decl(ctx *List_declContext) {}

// EnterType_def is called when production type_def is entered.
func (s *BasesashimiListener) EnterType_def(ctx *Type_defContext) {}

// ExitType_def is called when production type_def is exited.
func (s *BasesashimiListener) ExitType_def(ctx *Type_defContext) {}

// EnterAlias_decl is called when production alias_decl is entered.
func (s *BasesashimiListener) EnterAlias_decl(ctx *Alias_declContext) {}

// ExitAlias_decl is called when production alias_decl is exited.
func (s *BasesashimiListener) ExitAlias_decl(ctx *Alias_declContext) {}

// EnterType_is is called when production type_is is entered.
func (s *BasesashimiListener) EnterType_is(ctx *Type_isContext) {}

// ExitType_is is called when production type_is is exited.
func (s *BasesashimiListener) ExitType_is(ctx *Type_isContext) {}

// EnterProp_decl is called when production prop_decl is entered.
func (s *BasesashimiListener) EnterProp_decl(ctx *Prop_declContext) {}

// ExitProp_decl is called when production prop_decl is exited.
func (s *BasesashimiListener) ExitProp_decl(ctx *Prop_declContext) {}

// EnterEntity_decl is called when production entity_decl is entered.
func (s *BasesashimiListener) EnterEntity_decl(ctx *Entity_declContext) {}

// ExitEntity_decl is called when production entity_decl is exited.
func (s *BasesashimiListener) ExitEntity_decl(ctx *Entity_declContext) {}

// EnterExport is called when production export is entered.
func (s *BasesashimiListener) EnterExport(ctx *ExportContext) {}

// ExitExport is called when production export is exited.
func (s *BasesashimiListener) ExitExport(ctx *ExportContext) {}
