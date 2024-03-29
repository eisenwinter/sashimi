package build

import (
	"fmt"
	"strings"

	"github.com/eisenwinter/sashimi/util"
)

type firstPass struct {
	*BaseSashimiListener
	ctx                   *parserContext
	current               *defTableEntry
	p                     *propertyDef
	builder               typeBuilder
	source                string
	scopeStack            []string
	scopeAutoIdent        int
	loopAliasScopePending bool
	pendingAliasType      string
	pendingAlias          string
}

func plainFirstPassParser() *firstPass {
	return &firstPass{
		ctx: &parserContext{
			Def:              make(map[string]*defTableEntry),
			Errors:           make([]*lineReporter, 0),
			Warnings:         make([]*lineReporter, 0),
			Calls:            make(map[string]map[string]map[string][]string),
			KnownTypeAlias:   make(map[string][]*aliasEntry),
			KnownGlobals:     map[string]bool{"@pages": true},
			ProcessedSources: make(map[string]*sourceReport),
		},
		source:     "not-set",
		builder:    newTypeBuilder(),
		scopeStack: make([]string, 0),
	}
}

func firstPassParserWithSource(source string) *firstPass {
	return &firstPass{
		ctx: &parserContext{
			Def:              make(map[string]*defTableEntry),
			Errors:           make([]*lineReporter, 0),
			Warnings:         make([]*lineReporter, 0),
			Calls:            make(map[string]map[string]map[string][]string),
			KnownTypeAlias:   make(map[string][]*aliasEntry),
			KnownGlobals:     map[string]bool{"@pages": true},
			ProcessedSources: make(map[string]*sourceReport),
		},
		source:     source,
		builder:    newTypeBuilder(),
		scopeStack: make([]string, 0),
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

}

func (l *firstPass) ExitBlock(ctx *BlockContext) {
	if len(l.scopeStack) > 0 {
		error := &lineReporter{
			ErrorMarkerPos: 0,
			Line:           0,
			Message:        "The found start tag does not have a matching end tag. (Explicit scoping only.)",
			Source:         l.source,
			Code:           SashimiInvalidScopingStart,
		}
		l.ctx.Errors = append(l.ctx.Errors, error)
	}
	if _, ok := l.ctx.ProcessedSources[l.source]; !ok {
		l.ctx.ProcessedSources[l.source] = &sourceReport{}
	}

}

func (l *firstPass) getCurrentScope() string {
	if len(l.scopeStack) == 0 {
		return ""
	}
	n := len(l.scopeStack) - 1
	return fmt.Sprintf("%s::%v", l.source, l.scopeStack[n])
}

func (l *firstPass) PushScope(scope string) {
	l.scopeStack = append(l.scopeStack, scope)
}

func (l *firstPass) PopScope() string {
	if len(l.scopeStack) == 0 {
		error := &lineReporter{
			ErrorMarkerPos: 0,
			Line:           0,
			Message:        "The found end tag does not have a matching start tag. (Explicit scoping only.)",
			Source:         l.source,
			Code:           SashimiInvalidScopingEnd,
		}
		l.ctx.Errors = append(l.ctx.Errors, error)
		return ""
	}
	n := len(l.scopeStack) - 1
	val := l.scopeStack[n]
	l.scopeStack = l.scopeStack[:n]
	return val

}

func (l *firstPass) EnterCommandCall(ctx *CommandCallContext) {

	if _, ok := l.ctx.Calls[ctx.COMMAND().GetText()]; !ok {
		l.ctx.Calls[ctx.COMMAND().GetText()] = make(map[string]map[string][]string)
	}
	if _, ok := l.ctx.Calls[ctx.COMMAND().GetText()][ctx.Qualifier().GetText()]; !ok {
		l.ctx.Calls[ctx.COMMAND().GetText()][ctx.Qualifier().GetText()] = make(map[string][]string, 0)
		l.ctx.Calls[ctx.COMMAND().GetText()][ctx.Qualifier().GetText()][l.source] = make([]string, 0)
	}
	entries := l.ctx.Calls[ctx.COMMAND().GetText()][ctx.Qualifier().GetText()][l.source]
	l.ctx.Calls[ctx.COMMAND().GetText()][ctx.Qualifier().GetText()][l.source] = append(entries, l.getCurrentScope())
	if !strings.Contains(ctx.COMMAND().GetText(), "layout") {
		firstPart := strings.Split(ctx.Qualifier().GetText(), ".")[0]
		entity := &requiredEntity{
			name:  firstPart,
			many:  true,
			scope: l.getCurrentScope(),
		}
		if _, ok := l.ctx.ProcessedSources[l.source]; !ok {
			l.ctx.ProcessedSources[l.source] = &sourceReport{}
			l.ctx.ProcessedSources[l.source].requiredEntities = make([]*requiredEntity, 0)
		}
		l.ctx.ProcessedSources[l.source].requiredEntities = append(l.ctx.ProcessedSources[l.source].requiredEntities, entity)

	}
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
			Identifier: identifier,
			Properties: make(map[string]*propertyDef),
			IsUnique:   ctx.UNIQUE() != nil,
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

func (l *firstPass) EnterLoopCall(ctx *LoopCallContext) {
	if ctx.alias != nil {
		l.loopAliasScopePending = true
		l.pendingAlias = ctx.alias.GetText()
		if ctx.GLOBAL() != nil {
			l.pendingAliasType = ctx.GLOBAL().GetText()
		} else {
			l.pendingAliasType = ctx.Qualifier().GetText()
		}

	}
}

func (l *firstPass) ExitLoopCall(ctx *LoopCallContext) {
	//Todo we should prolly trace and validate the predicate if it even makes sense ...
	// if !ctx.Predicate().IsEmpty() {
	// 	fmt.Println(ctx.Predicate().GetText())
	// }

	if _, ok := l.ctx.Calls[ctx.LOOP().GetText()]; !ok {
		l.ctx.Calls[ctx.LOOP().GetText()] = make(map[string]map[string][]string)
	}
	var ref string
	if ctx.GLOBAL() != nil {
		ref = ctx.GLOBAL().GetText()
	} else {
		ref = ctx.Qualifier().GetText()
		firstPart := strings.Split(ref, ".")[0]
		if _, ok := l.ctx.ProcessedSources[l.source]; !ok {
			l.ctx.ProcessedSources[l.source] = &sourceReport{}
			l.ctx.ProcessedSources[l.source].requiredEntities = make([]*requiredEntity, 0)
		}
		l.ctx.ProcessedSources[l.source].isMany = true
		if l.ctx.ProcessedSources[l.source].manyBy == nil {
			l.ctx.ProcessedSources[l.source].manyBy = make([]string, 0)
		}
		l.ctx.ProcessedSources[l.source].manyBy = util.AddIfNotContained(l.ctx.ProcessedSources[l.source].manyBy, ref)
		if ctx.Predicate() != nil && !ctx.Predicate().IsEmpty() {
			entity := &requiredEntity{
				name:      firstPart,
				many:      false,
				predicate: ctx.Predicate().GetText(),
				scope:     l.getCurrentScope(),
			}
			l.ctx.ProcessedSources[l.source].requiredEntities = append(l.ctx.ProcessedSources[l.source].requiredEntities, entity)
		} else {
			entity := &requiredEntity{
				name:  firstPart,
				many:  false,
				scope: l.getCurrentScope(),
			}
			l.ctx.ProcessedSources[l.source].requiredEntities = append(l.ctx.ProcessedSources[l.source].requiredEntities, entity)
		}

	}

	if _, ok := l.ctx.Calls[ctx.LOOP().GetText()][ref]; !ok {
		l.ctx.Calls[ctx.LOOP().GetText()][ref] = make(map[string][]string, 0)
		l.ctx.Calls[ctx.LOOP().GetText()][ref][l.source] = make([]string, 0)
	}

	//set awaiting with scope
	entries := l.ctx.Calls[ctx.LOOP().GetText()][ref][l.source]
	l.ctx.Calls[ctx.LOOP().GetText()][ref][l.source] = append(entries, l.getCurrentScope())

}

func (l *firstPass) ExitScopeBegin(ctx *ScopeBeginContext) {
	var scopeIdent string
	//decide if its explicit or implicit scoping
	if ctx.SCOPEIDENT() != nil {
		scopeIdent = strings.Trim(ctx.SCOPEIDENT().GetText(), "'")
	} else {
		l.scopeAutoIdent++
		scopeIdent = fmt.Sprintf("~:%v", l.scopeAutoIdent)
	}
	l.PushScope(scopeIdent)

	if l.loopAliasScopePending {
		if _, ok := l.ctx.KnownTypeAlias[l.pendingAlias]; !ok {
			l.ctx.KnownTypeAlias[l.pendingAlias] = make([]*aliasEntry, 0)
		}
		l.ctx.KnownTypeAlias[l.pendingAlias] = append(l.ctx.KnownTypeAlias[l.pendingAlias], &aliasEntry{
			UnderlyingType: l.pendingAliasType,
			Scope:          l.getCurrentScope(),
			Source:         l.source,
		})
		l.loopAliasScopePending = false
	}

}

func (l *firstPass) ExitScopeEnd(ctx *ScopeEndContext) {
	scopeIdent := l.PopScope()
	if ctx.SCOPEIDENT() != nil {
		trimmed := strings.Trim(ctx.SCOPEIDENT().GetText(), "'")
		if trimmed != scopeIdent {
			error := &lineReporter{
				ErrorMarkerPos: 0,
				Line:           0,
				Message:        "The found end tag has a not expected scope identifier. (Explicit scoping only.)",
				Source:         l.source,
				Code:           SashimiInvalidScopingEnd,
			}
			l.ctx.Errors = append(l.ctx.Errors, error)
		}
	}
}
