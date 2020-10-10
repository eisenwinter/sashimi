// Code generated from C:\gitrepos\sashimi/grammar/Sashimi.g4 by ANTLR 4.8. DO NOT EDIT.

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

// EnterPredicate is called when production predicate is entered.
func (s *BaseSashimiListener) EnterPredicate(ctx *PredicateContext) {}

// ExitPredicate is called when production predicate is exited.
func (s *BaseSashimiListener) ExitPredicate(ctx *PredicateContext) {}

// EnterQualifier is called when production qualifier is entered.
func (s *BaseSashimiListener) EnterQualifier(ctx *QualifierContext) {}

// ExitQualifier is called when production qualifier is exited.
func (s *BaseSashimiListener) ExitQualifier(ctx *QualifierContext) {}

// EnterLoopCall is called when production loopCall is entered.
func (s *BaseSashimiListener) EnterLoopCall(ctx *LoopCallContext) {}

// ExitLoopCall is called when production loopCall is exited.
func (s *BaseSashimiListener) ExitLoopCall(ctx *LoopCallContext) {}

// EnterEntityDef is called when production entityDef is entered.
func (s *BaseSashimiListener) EnterEntityDef(ctx *EntityDefContext) {}

// ExitEntityDef is called when production entityDef is exited.
func (s *BaseSashimiListener) ExitEntityDef(ctx *EntityDefContext) {}

// EnterCommandCall is called when production commandCall is entered.
func (s *BaseSashimiListener) EnterCommandCall(ctx *CommandCallContext) {}

// ExitCommandCall is called when production commandCall is exited.
func (s *BaseSashimiListener) ExitCommandCall(ctx *CommandCallContext) {}

// EnterExport is called when production export is entered.
func (s *BaseSashimiListener) EnterExport(ctx *ExportContext) {}

// ExitExport is called when production export is exited.
func (s *BaseSashimiListener) ExitExport(ctx *ExportContext) {}

// EnterBlock is called when production block is entered.
func (s *BaseSashimiListener) EnterBlock(ctx *BlockContext) {}

// ExitBlock is called when production block is exited.
func (s *BaseSashimiListener) ExitBlock(ctx *BlockContext) {}

// EnterBoolExpression is called when production boolExpression is entered.
func (s *BaseSashimiListener) EnterBoolExpression(ctx *BoolExpressionContext) {}

// ExitBoolExpression is called when production boolExpression is exited.
func (s *BaseSashimiListener) ExitBoolExpression(ctx *BoolExpressionContext) {}

// EnterOp is called when production op is entered.
func (s *BaseSashimiListener) EnterOp(ctx *OpContext) {}

// ExitOp is called when production op is exited.
func (s *BaseSashimiListener) ExitOp(ctx *OpContext) {}

// EnterComparator is called when production comparator is entered.
func (s *BaseSashimiListener) EnterComparator(ctx *ComparatorContext) {}

// ExitComparator is called when production comparator is exited.
func (s *BaseSashimiListener) ExitComparator(ctx *ComparatorContext) {}

// EnterBinary is called when production binary is entered.
func (s *BaseSashimiListener) EnterBinary(ctx *BinaryContext) {}

// ExitBinary is called when production binary is exited.
func (s *BaseSashimiListener) ExitBinary(ctx *BinaryContext) {}

// EnterTruth is called when production truth is entered.
func (s *BaseSashimiListener) EnterTruth(ctx *TruthContext) {}

// ExitTruth is called when production truth is exited.
func (s *BaseSashimiListener) ExitTruth(ctx *TruthContext) {}
