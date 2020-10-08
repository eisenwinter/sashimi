package parser

import (
	"fmt"
)

type FirstPass struct {
	*BaseSashimiListener
	ctx     *parserContext
	current *defTableEntry
	p       *propertyDef
	builder typeBuilder
}

func NewFirstPassParser() *FirstPass {
	return &FirstPass{
		ctx: &parserContext{
			Def:      make(map[string]*defTableEntry),
			Errors:   make([]*lineReporter, 0),
			Warnings: make([]*lineReporter, 0),
			Calls:    make(map[string]map[string]bool),
		},
		builder: newTypeBuilder(),
	}
}

func (l *FirstPass) Dump() {
	fmt.Printf("%+v", l.ctx)
	for k, v := range l.ctx.Def {
		fmt.Printf("\r\n  %s = %+v", k, v)
		for t, p := range v.Properties {
			fmt.Printf("\r\n    %s = %+v", t, p)
		}
	}
	for k, v := range l.ctx.Calls {
		fmt.Printf("Calls for %s", k)
		for c := range v {
			fmt.Printf("Property %s", c)
		}
	}
}

func (l *FirstPass) EnterExport(ctx *ExportContext) {
	fmt.Println("Parsing sashimi block")
}

func (l *FirstPass) EnterCommand_call(ctx *Command_callContext) {
	_, ok := l.ctx.Calls[ctx.COMMAND().GetText()]
	if !ok {
		l.ctx.Calls[ctx.COMMAND().GetText()] = make(map[string]bool)
	}
	qualifier := ""
	for _, v := range ctx.AllIDENT() {
		if qualifier != "" {
			qualifier += "."
		}
		qualifier += v.GetText()
	}
	l.ctx.Calls[ctx.COMMAND().GetText()][qualifier] = true
}

func (l *FirstPass) EnterEntity_def(ctx *Entity_defContext) {
	identifier := ctx.IDENT().GetText()
	if val, ok := l.ctx.Def[identifier]; ok {
		start := ctx.IDENT().GetSymbol().GetStart()
		ln := ctx.IDENT().GetSymbol().GetLine()
		warning := &lineReporter{
			ErrorMarkerPos: start,
			Line:           ln,
			Message:        fmt.Sprintf("Entity with name `%s` has already been declared, reusing known definiton. This should be avoided.", identifier),
		}
		l.ctx.Warnings = append(l.ctx.Warnings, warning)
		l.current = val
	} else {
		table := &defTableEntry{
			Identfier:  identifier,
			Properties: make(map[string]*propertyDef),
		}
		l.ctx.Def[identifier] = table
		l.current = table
	}

}

func (l *FirstPass) EnterProp_decl(c *Prop_declContext) {
	identifier := c.IDENT().GetText()
	if _, ok := l.current.Properties[identifier]; ok {
		start := c.IDENT().GetSymbol().GetStart()
		ln := c.IDENT().GetSymbol().GetLine()
		error := &lineReporter{
			ErrorMarkerPos: start,
			Line:           ln,
			Message:        fmt.Sprintf("Property with name `%s` has already been declared.", identifier),
		}
		l.ctx.Errors = append(l.ctx.Errors, error)
	} else {
		def := &propertyDef{}
		l.current.Properties[identifier] = def
		l.p = def
	}
}

func (l *FirstPass) EnterAlias_decl(c *Alias_declContext) {
	if l.p != nil {
		l.p.Alias = c.ALIAS().GetText()
	} else {
		start := c.AS().GetSymbol().GetStart()
		ln := c.AS().GetSymbol().GetLine()
		error := &lineReporter{
			ErrorMarkerPos: start,
			Line:           ln,
			Message:        "Invalid alias declaration at this point.",
		}
		l.ctx.Errors = append(l.ctx.Errors, error)
	}
}

func (l *FirstPass) ExitExport(ctx *ExportContext) {
	fmt.Println("Parsed sashimi block")
	if l.current != nil {
		l.current.IsDefined = true
	}
}

func (l *FirstPass) ExitUnion_decl(c *Union_declContext) {
	tags := make([]string, 0)
	for _, v := range c.AllALIAS() {
		tags = append(tags, v.GetText())
	}
	isExcl := c.UNION().GetText() == "oneOf"
	l.builder.Union(isExcl, tags)
}

func (l *FirstPass) EnterType_decl(ctx *Type_declContext) {
	l.builder.Scalar(ctx.TYPE().GetText())
}

func (l *FirstPass) ExitKey_value_pair(ctx *Key_value_pairContext) {
	if l.builder.IsAwaitingConstraint() {
		kvp := ctx.AllIDENT()
		l.builder.Constraint(kvp[0].GetText(), kvp[1].GetText())
	}
}

func (l *FirstPass) ExitKey_atom(ctx *Key_atomContext) {
	if l.builder.IsAwaitingConstraint() {
		l.builder.Constraint(ctx.IDENT().GetText(), "atomic")
	}
}

func (l *FirstPass) ExitType_is(ctx *Type_isContext) {
	l.p.Type = l.builder.Build()
	l.builder.Reset()
}

func (l *FirstPass) ExitEntity_ref(ctx *Entity_refContext) {
	l.builder.Ref(ctx.IDENT().GetText())
}

func (l *FirstPass) EnterType_def(ctx *Type_defContext) {
	l.builder.Reset()
}

func (l *FirstPass) EnterList_decl(c *List_declContext) {
	l.builder.List()
}
