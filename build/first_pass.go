package build

import (
	"fmt"
)

type firstPass struct {
	*BaseSashimiListener
	ctx     *parserContext
	current *defTableEntry
	p       *propertyDef
	builder typeBuilder
	source  string
}

func plainFirstPassParser() *firstPass {
	return &firstPass{
		ctx: &parserContext{
			Def:      make(map[string]*defTableEntry),
			Errors:   make([]*lineReporter, 0),
			Warnings: make([]*lineReporter, 0),
			Calls:    make(map[string]map[string]string),
		},
		source:  "not-set",
		builder: newTypeBuilder(),
	}
}

func (l *firstPass) Dump() {
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

func (l *firstPass) EnterExport(ctx *ExportContext) {
	fmt.Println("Parsing sashimi block")

}

func (l *firstPass) EnterCommandCall(ctx *CommandCallContext) {
	_, ok := l.ctx.Calls[ctx.COMMAND().GetText()]
	if !ok {
		l.ctx.Calls[ctx.COMMAND().GetText()] = make(map[string]string)
	}
	qualifier := ""
	for _, v := range ctx.AllIDENT() {
		if qualifier != "" {
			qualifier += "."
		}
		qualifier += v.GetText()
	}
	l.ctx.Calls[ctx.COMMAND().GetText()][qualifier] = l.source
}

func (l *firstPass) EnterEntityDef(ctx *EntityDefContext) {
	identifier := ctx.IDENT().GetText()
	if val, ok := l.ctx.Def[identifier]; ok {
		start := ctx.IDENT().GetSymbol().GetStart()
		ln := ctx.IDENT().GetSymbol().GetLine()
		warning := &lineReporter{
			ErrorMarkerPos: start,
			Line:           ln,
			Message:        fmt.Sprintf("Entity with name `%s` has already been declared, reusing known definiton. This should be avoided", identifier),
			Source:         l.source,
			Code:           SashimiEntityAlreadyDeclared,
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

func (l *firstPass) EnterPropDecl(c *PropDeclContext) {
	identifier := c.IDENT().GetText()
	if _, ok := l.current.Properties[identifier]; ok {
		start := c.IDENT().GetSymbol().GetStart()
		ln := c.IDENT().GetSymbol().GetLine()
		error := &lineReporter{
			ErrorMarkerPos: start,
			Line:           ln,
			Message:        fmt.Sprintf("Property with name `%s` has already been declared", identifier),
			Source:         l.source,
			Code:           SashimiPropertyAlreadyDeclared,
		}
		l.ctx.Errors = append(l.ctx.Errors, error)
	} else {
		def := &propertyDef{}
		l.current.Properties[identifier] = def
		l.p = def
	}
}

func (l *firstPass) EnterAliasDecl(c *AliasDeclContext) {
	if l.p != nil {
		l.p.Alias = c.ALIAS().GetText()
	} else {
		start := c.AS().GetSymbol().GetStart()
		ln := c.AS().GetSymbol().GetLine()
		error := &lineReporter{
			ErrorMarkerPos: start,
			Line:           ln,
			Message:        "Invalid alias declaration at this point",
			Source:         l.source,
			Code:           SashimiInvalidAlias,
		}
		l.ctx.Errors = append(l.ctx.Errors, error)
	}
}

func (l *firstPass) ExitExport(ctx *ExportContext) {
	fmt.Println("Parsed sashimi block")
	if l.current != nil {
		l.current.IsDefined = true
	}
}

func (l *firstPass) ExitUnionDecl(c *UnionDeclContext) {
	tags := make([]string, 0)
	for _, v := range c.AllALIAS() {
		tags = append(tags, v.GetText())
	}
	isExcl := c.UNION().GetText() == "oneOf"
	l.builder.Union(isExcl, tags)
}

func (l *firstPass) EnterTypeDecl(ctx *TypeDeclContext) {
	l.builder.Scalar(ctx.TYPE().GetText())
}

func (l *firstPass) ExitKeyValuePair(ctx *KeyValuePairContext) {
	if l.builder.IsAwaitingConstraint() {
		kvp := ctx.AllIDENT()
		l.builder.Constraint(kvp[0].GetText(), kvp[1].GetText())
	}
}

func (l *firstPass) ExitKeyAtom(ctx *KeyAtomContext) {
	if l.builder.IsAwaitingConstraint() {
		l.builder.Constraint(ctx.IDENT().GetText(), "atomic")
	}
}

func (l *firstPass) ExitTypeIs(ctx *TypeIsContext) {
	l.p.Type = l.builder.Build()
	l.builder.Reset()
}

func (l *firstPass) ExitEntityRef(ctx *EntityRefContext) {
	l.builder.Ref(ctx.IDENT().GetText())
}

func (l *firstPass) EnterTypeDef(ctx *TypeDefContext) {
	l.builder.Reset()
}

func (l *firstPass) EnterListDecl(c *ListDeclContext) {
	l.builder.List()
}
