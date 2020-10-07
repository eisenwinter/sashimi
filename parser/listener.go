package parser

import (
	"fmt"
)

type SashimiListeningParser struct {
	*BaseSashimiListener
}

func (l *SashimiListeningParser) EnterExport(ctx *ExportContext) {
	fmt.Println("Parsing sashimi block")
}

func (l *SashimiListeningParser) EnterEntity_def(ctx *Entity_defContext) {
	fmt.Println("block is entity definition")
	fmt.Println("for", ctx.IDENT())
}

func (l *SashimiListeningParser) EnterProp_decl(c *Prop_declContext) {
	fmt.Println("Property declared")
	fmt.Println("-> ", c.IDENT())
}

func (l *SashimiListeningParser) ExitProp_decl(c *Prop_declContext) {
	fmt.Println("Property declared exit")
	fmt.Println("-> ", c.IDENT())
}

func (l *SashimiListeningParser) EnterAlias_decl(c *Alias_declContext) {
	fmt.Println("with alias ", c.ALIAS())
}

func (l *SashimiListeningParser) ExitExport(ctx *ExportContext) {
	fmt.Println("Parsed sashimi block")
}

func (l *SashimiListeningParser) ExitUnion_decl(c *Union_declContext) {
	fmt.Println("is union ", c.AllALIAS())
}

func (l *SashimiListeningParser) ExitType_decl(c *Type_declContext) {
	fmt.Println("type", c.TYPE)
}

func (l *SashimiListeningParser) EnterList_decl(c *List_declContext) {
	fmt.Println("is list of")
}

func (l *SashimiListeningParser) ExitList_decl(c *List_declContext) {
	fmt.Println("end of list")
}

func (l *SashimiListeningParser) ExitCommand_call(c *Command_callContext) {
	fmt.Println("Is command", c.COMMAND())
	fmt.Println("with ident", c.IDENT())
}
