// Code generated from Sashimi.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // Sashimi

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseSashimiListener is a complete listener for a parse tree produced by SashimiParser.
type BaseSashimiListener struct{}

var _ SashimiListener = &BaseSashimiListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseSashimiListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseSashimiListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseSashimiListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseSashimiListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterKey_value_pair is called when production key_value_pair is entered.
func (s *BaseSashimiListener) EnterKey_value_pair(ctx *Key_value_pairContext) {}

// ExitKey_value_pair is called when production key_value_pair is exited.
func (s *BaseSashimiListener) ExitKey_value_pair(ctx *Key_value_pairContext) {}

// EnterKey_atom is called when production key_atom is entered.
func (s *BaseSashimiListener) EnterKey_atom(ctx *Key_atomContext) {}

// ExitKey_atom is called when production key_atom is exited.
func (s *BaseSashimiListener) ExitKey_atom(ctx *Key_atomContext) {}

// EnterConstraint_list is called when production constraint_list is entered.
func (s *BaseSashimiListener) EnterConstraint_list(ctx *Constraint_listContext) {}

// ExitConstraint_list is called when production constraint_list is exited.
func (s *BaseSashimiListener) ExitConstraint_list(ctx *Constraint_listContext) {}

// EnterUnion_decl is called when production union_decl is entered.
func (s *BaseSashimiListener) EnterUnion_decl(ctx *Union_declContext) {}

// ExitUnion_decl is called when production union_decl is exited.
func (s *BaseSashimiListener) ExitUnion_decl(ctx *Union_declContext) {}

// EnterType_decl is called when production type_decl is entered.
func (s *BaseSashimiListener) EnterType_decl(ctx *Type_declContext) {}

// ExitType_decl is called when production type_decl is exited.
func (s *BaseSashimiListener) ExitType_decl(ctx *Type_declContext) {}

// EnterEntity_ref is called when production entity_ref is entered.
func (s *BaseSashimiListener) EnterEntity_ref(ctx *Entity_refContext) {}

// ExitEntity_ref is called when production entity_ref is exited.
func (s *BaseSashimiListener) ExitEntity_ref(ctx *Entity_refContext) {}

// EnterList_decl is called when production list_decl is entered.
func (s *BaseSashimiListener) EnterList_decl(ctx *List_declContext) {}

// ExitList_decl is called when production list_decl is exited.
func (s *BaseSashimiListener) ExitList_decl(ctx *List_declContext) {}

// EnterType_def is called when production type_def is entered.
func (s *BaseSashimiListener) EnterType_def(ctx *Type_defContext) {}

// ExitType_def is called when production type_def is exited.
func (s *BaseSashimiListener) ExitType_def(ctx *Type_defContext) {}

// EnterAlias_decl is called when production alias_decl is entered.
func (s *BaseSashimiListener) EnterAlias_decl(ctx *Alias_declContext) {}

// ExitAlias_decl is called when production alias_decl is exited.
func (s *BaseSashimiListener) ExitAlias_decl(ctx *Alias_declContext) {}

// EnterType_is is called when production type_is is entered.
func (s *BaseSashimiListener) EnterType_is(ctx *Type_isContext) {}

// ExitType_is is called when production type_is is exited.
func (s *BaseSashimiListener) ExitType_is(ctx *Type_isContext) {}

// EnterProp_decl is called when production prop_decl is entered.
func (s *BaseSashimiListener) EnterProp_decl(ctx *Prop_declContext) {}

// ExitProp_decl is called when production prop_decl is exited.
func (s *BaseSashimiListener) ExitProp_decl(ctx *Prop_declContext) {}

// EnterCommand_call is called when production command_call is entered.
func (s *BaseSashimiListener) EnterCommand_call(ctx *Command_callContext) {}

// ExitCommand_call is called when production command_call is exited.
func (s *BaseSashimiListener) ExitCommand_call(ctx *Command_callContext) {}

// EnterEntity_def is called when production entity_def is entered.
func (s *BaseSashimiListener) EnterEntity_def(ctx *Entity_defContext) {}

// ExitEntity_def is called when production entity_def is exited.
func (s *BaseSashimiListener) ExitEntity_def(ctx *Entity_defContext) {}

// EnterExport is called when production export is entered.
func (s *BaseSashimiListener) EnterExport(ctx *ExportContext) {}

// ExitExport is called when production export is exited.
func (s *BaseSashimiListener) ExitExport(ctx *ExportContext) {}

// EnterBlock is called when production block is entered.
func (s *BaseSashimiListener) EnterBlock(ctx *BlockContext) {}

// ExitBlock is called when production block is exited.
func (s *BaseSashimiListener) ExitBlock(ctx *BlockContext) {}
