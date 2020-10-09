// Code generated from C:\tmp\sashimi\sashimi\grammar\Sashimi.g4 by ANTLR 4.8. DO NOT EDIT.

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

	// EnterCommandCall is called when entering the commandCall production.
	EnterCommandCall(c *CommandCallContext)

	// EnterEntityDef is called when entering the entityDef production.
	EnterEntityDef(c *EntityDefContext)

	// EnterExport is called when entering the export production.
	EnterExport(c *ExportContext)

	// EnterBlock is called when entering the block production.
	EnterBlock(c *BlockContext)

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

	// ExitCommandCall is called when exiting the commandCall production.
	ExitCommandCall(c *CommandCallContext)

	// ExitEntityDef is called when exiting the entityDef production.
	ExitEntityDef(c *EntityDefContext)

	// ExitExport is called when exiting the export production.
	ExitExport(c *ExportContext)

	// ExitBlock is called when exiting the block production.
	ExitBlock(c *BlockContext)
}
