// Code generated from C:\gitrepos\sashimi/grammar/Sashimi.g4 by ANTLR 4.8. DO NOT EDIT.

package build // Sashimi
import "github.com/antlr/antlr4/runtime/Go/antlr"

// SashimiListener is a complete listener for a parse tree produced by SashimiParser.
type SashimiListener interface {
	antlr.ParseTreeListener

	// EnterKeyValuePair is called when entering the keyValuePair production.
	EnterKeyValuePair(c *KeyValuePairContext)

	// EnterKeyAtom is called when entering the keyAtom production.
	EnterKeyAtom(c *KeyAtomContext)

	// EnterConstraintList is called when entering the constraintList production.
	EnterConstraintList(c *ConstraintListContext)

	// EnterUnionDecl is called when entering the unionDecl production.
	EnterUnionDecl(c *UnionDeclContext)

	// EnterTypeDecl is called when entering the typeDecl production.
	EnterTypeDecl(c *TypeDeclContext)

	// EnterEntityRef is called when entering the entityRef production.
	EnterEntityRef(c *EntityRefContext)

	// EnterListDecl is called when entering the listDecl production.
	EnterListDecl(c *ListDeclContext)

	// EnterTypeDef is called when entering the typeDef production.
	EnterTypeDef(c *TypeDefContext)

	// EnterAliasDecl is called when entering the aliasDecl production.
	EnterAliasDecl(c *AliasDeclContext)

	// EnterTypeIs is called when entering the typeIs production.
	EnterTypeIs(c *TypeIsContext)

	// EnterPropDecl is called when entering the propDecl production.
	EnterPropDecl(c *PropDeclContext)

	// EnterPredicate is called when entering the predicate production.
	EnterPredicate(c *PredicateContext)

	// EnterQualifier is called when entering the qualifier production.
	EnterQualifier(c *QualifierContext)

	// EnterLoopCall is called when entering the loopCall production.
	EnterLoopCall(c *LoopCallContext)

	// EnterEntityDef is called when entering the entityDef production.
	EnterEntityDef(c *EntityDefContext)

	// EnterScopeBegin is called when entering the scopeBegin production.
	EnterScopeBegin(c *ScopeBeginContext)

	// EnterScopeEnd is called when entering the scopeEnd production.
	EnterScopeEnd(c *ScopeEndContext)

	// EnterCommandCall is called when entering the commandCall production.
	EnterCommandCall(c *CommandCallContext)

	// EnterBlockScope is called when entering the blockScope production.
	EnterBlockScope(c *BlockScopeContext)

	// EnterExport is called when entering the export production.
	EnterExport(c *ExportContext)

	// EnterBlock is called when entering the block production.
	EnterBlock(c *BlockContext)

	// EnterBoolExpression is called when entering the boolExpression production.
	EnterBoolExpression(c *BoolExpressionContext)

	// EnterOp is called when entering the op production.
	EnterOp(c *OpContext)

	// EnterComparator is called when entering the comparator production.
	EnterComparator(c *ComparatorContext)

	// EnterBinary is called when entering the binary production.
	EnterBinary(c *BinaryContext)

	// EnterTruth is called when entering the truth production.
	EnterTruth(c *TruthContext)

	// ExitKeyValuePair is called when exiting the keyValuePair production.
	ExitKeyValuePair(c *KeyValuePairContext)

	// ExitKeyAtom is called when exiting the keyAtom production.
	ExitKeyAtom(c *KeyAtomContext)

	// ExitConstraintList is called when exiting the constraintList production.
	ExitConstraintList(c *ConstraintListContext)

	// ExitUnionDecl is called when exiting the unionDecl production.
	ExitUnionDecl(c *UnionDeclContext)

	// ExitTypeDecl is called when exiting the typeDecl production.
	ExitTypeDecl(c *TypeDeclContext)

	// ExitEntityRef is called when exiting the entityRef production.
	ExitEntityRef(c *EntityRefContext)

	// ExitListDecl is called when exiting the listDecl production.
	ExitListDecl(c *ListDeclContext)

	// ExitTypeDef is called when exiting the typeDef production.
	ExitTypeDef(c *TypeDefContext)

	// ExitAliasDecl is called when exiting the aliasDecl production.
	ExitAliasDecl(c *AliasDeclContext)

	// ExitTypeIs is called when exiting the typeIs production.
	ExitTypeIs(c *TypeIsContext)

	// ExitPropDecl is called when exiting the propDecl production.
	ExitPropDecl(c *PropDeclContext)

	// ExitPredicate is called when exiting the predicate production.
	ExitPredicate(c *PredicateContext)

	// ExitQualifier is called when exiting the qualifier production.
	ExitQualifier(c *QualifierContext)

	// ExitLoopCall is called when exiting the loopCall production.
	ExitLoopCall(c *LoopCallContext)

	// ExitEntityDef is called when exiting the entityDef production.
	ExitEntityDef(c *EntityDefContext)

	// ExitScopeBegin is called when exiting the scopeBegin production.
	ExitScopeBegin(c *ScopeBeginContext)

	// ExitScopeEnd is called when exiting the scopeEnd production.
	ExitScopeEnd(c *ScopeEndContext)

	// ExitCommandCall is called when exiting the commandCall production.
	ExitCommandCall(c *CommandCallContext)

	// ExitBlockScope is called when exiting the blockScope production.
	ExitBlockScope(c *BlockScopeContext)

	// ExitExport is called when exiting the export production.
	ExitExport(c *ExportContext)

	// ExitBlock is called when exiting the block production.
	ExitBlock(c *BlockContext)

	// ExitBoolExpression is called when exiting the boolExpression production.
	ExitBoolExpression(c *BoolExpressionContext)

	// ExitOp is called when exiting the op production.
	ExitOp(c *OpContext)

	// ExitComparator is called when exiting the comparator production.
	ExitComparator(c *ComparatorContext)

	// ExitBinary is called when exiting the binary production.
	ExitBinary(c *BinaryContext)

	// ExitTruth is called when exiting the truth production.
	ExitTruth(c *TruthContext)
}
