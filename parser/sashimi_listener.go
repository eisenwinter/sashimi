// Code generated from Sashimi.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // Sashimi

import "github.com/antlr/antlr4/runtime/Go/antlr"

// SashimiListener is a complete listener for a parse tree produced by SashimiParser.
type SashimiListener interface {
	antlr.ParseTreeListener

	// EnterUnion_decl is called when entering the union_decl production.
	EnterUnion_decl(c *Union_declContext)

	// EnterType_decl is called when entering the type_decl production.
	EnterType_decl(c *Type_declContext)

	// EnterList_decl is called when entering the list_decl production.
	EnterList_decl(c *List_declContext)

	// EnterType_def is called when entering the type_def production.
	EnterType_def(c *Type_defContext)

	// EnterAlias_decl is called when entering the alias_decl production.
	EnterAlias_decl(c *Alias_declContext)

	// EnterType_is is called when entering the type_is production.
	EnterType_is(c *Type_isContext)

	// EnterProp_decl is called when entering the prop_decl production.
	EnterProp_decl(c *Prop_declContext)

	// EnterEntity_decl is called when entering the entity_decl production.
	EnterEntity_decl(c *Entity_declContext)

	// EnterExport is called when entering the export production.
	EnterExport(c *ExportContext)

	// ExitUnion_decl is called when exiting the union_decl production.
	ExitUnion_decl(c *Union_declContext)

	// ExitType_decl is called when exiting the type_decl production.
	ExitType_decl(c *Type_declContext)

	// ExitList_decl is called when exiting the list_decl production.
	ExitList_decl(c *List_declContext)

	// ExitType_def is called when exiting the type_def production.
	ExitType_def(c *Type_defContext)

	// ExitAlias_decl is called when exiting the alias_decl production.
	ExitAlias_decl(c *Alias_declContext)

	// ExitType_is is called when exiting the type_is production.
	ExitType_is(c *Type_isContext)

	// ExitProp_decl is called when exiting the prop_decl production.
	ExitProp_decl(c *Prop_declContext)

	// ExitEntity_decl is called when exiting the entity_decl production.
	ExitEntity_decl(c *Entity_declContext)

	// ExitExport is called when exiting the export production.
	ExitExport(c *ExportContext)
}
