// Code generated from C:\tmp\sashimi\sashimi\grammar\Sashimi.g4 by ANTLR 4.8. DO NOT EDIT.

package build // Sashimi
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

// EnterKeyValuePair is called when production keyValuePair is entered.
func (s *BaseSashimiListener) EnterKeyValuePair(ctx *KeyValuePairContext) {}

// ExitKeyValuePair is called when production keyValuePair is exited.
func (s *BaseSashimiListener) ExitKeyValuePair(ctx *KeyValuePairContext) {}

// EnterKeyAtom is called when production keyAtom is entered.
func (s *BaseSashimiListener) EnterKeyAtom(ctx *KeyAtomContext) {}

// ExitKeyAtom is called when production keyAtom is exited.
func (s *BaseSashimiListener) ExitKeyAtom(ctx *KeyAtomContext) {}

// EnterConstraintList is called when production constraintList is entered.
func (s *BaseSashimiListener) EnterConstraintList(ctx *ConstraintListContext) {}

// ExitConstraintList is called when production constraintList is exited.
func (s *BaseSashimiListener) ExitConstraintList(ctx *ConstraintListContext) {}

// EnterUnionDecl is called when production unionDecl is entered.
func (s *BaseSashimiListener) EnterUnionDecl(ctx *UnionDeclContext) {}

// ExitUnionDecl is called when production unionDecl is exited.
func (s *BaseSashimiListener) ExitUnionDecl(ctx *UnionDeclContext) {}

// EnterTypeDecl is called when production typeDecl is entered.
func (s *BaseSashimiListener) EnterTypeDecl(ctx *TypeDeclContext) {}

// ExitTypeDecl is called when production typeDecl is exited.
func (s *BaseSashimiListener) ExitTypeDecl(ctx *TypeDeclContext) {}

// EnterEntityRef is called when production entityRef is entered.
func (s *BaseSashimiListener) EnterEntityRef(ctx *EntityRefContext) {}

// ExitEntityRef is called when production entityRef is exited.
func (s *BaseSashimiListener) ExitEntityRef(ctx *EntityRefContext) {}

// EnterListDecl is called when production listDecl is entered.
func (s *BaseSashimiListener) EnterListDecl(ctx *ListDeclContext) {}

// ExitListDecl is called when production listDecl is exited.
func (s *BaseSashimiListener) ExitListDecl(ctx *ListDeclContext) {}

// EnterTypeDef is called when production typeDef is entered.
func (s *BaseSashimiListener) EnterTypeDef(ctx *TypeDefContext) {}

// ExitTypeDef is called when production typeDef is exited.
func (s *BaseSashimiListener) ExitTypeDef(ctx *TypeDefContext) {}

// EnterAliasDecl is called when production aliasDecl is entered.
func (s *BaseSashimiListener) EnterAliasDecl(ctx *AliasDeclContext) {}

// ExitAliasDecl is called when production aliasDecl is exited.
func (s *BaseSashimiListener) ExitAliasDecl(ctx *AliasDeclContext) {}

// EnterTypeIs is called when production typeIs is entered.
func (s *BaseSashimiListener) EnterTypeIs(ctx *TypeIsContext) {}

// ExitTypeIs is called when production typeIs is exited.
func (s *BaseSashimiListener) ExitTypeIs(ctx *TypeIsContext) {}

// EnterPropDecl is called when production propDecl is entered.
func (s *BaseSashimiListener) EnterPropDecl(ctx *PropDeclContext) {}

// ExitPropDecl is called when production propDecl is exited.
func (s *BaseSashimiListener) ExitPropDecl(ctx *PropDeclContext) {}

// EnterCommandCall is called when production commandCall is entered.
func (s *BaseSashimiListener) EnterCommandCall(ctx *CommandCallContext) {}

// ExitCommandCall is called when production commandCall is exited.
func (s *BaseSashimiListener) ExitCommandCall(ctx *CommandCallContext) {}

// EnterEntityDef is called when production entityDef is entered.
func (s *BaseSashimiListener) EnterEntityDef(ctx *EntityDefContext) {}

// ExitEntityDef is called when production entityDef is exited.
func (s *BaseSashimiListener) ExitEntityDef(ctx *EntityDefContext) {}

// EnterExport is called when production export is entered.
func (s *BaseSashimiListener) EnterExport(ctx *ExportContext) {}

// ExitExport is called when production export is exited.
func (s *BaseSashimiListener) ExitExport(ctx *ExportContext) {}

// EnterBlock is called when production block is entered.
func (s *BaseSashimiListener) EnterBlock(ctx *BlockContext) {}

// ExitBlock is called when production block is exited.
func (s *BaseSashimiListener) ExitBlock(ctx *BlockContext) {}
