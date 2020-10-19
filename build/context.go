package build

import (
	"fmt"
	"strings"

	"github.com/eisenwinter/sashimi/graph"
)

type parserContext struct {
	Def      map[string]*defTableEntry
	Errors   []*lineReporter
	Warnings []*lineReporter
	Calls    map[string]map[string]map[string][]string
	//Hokay this is very primitive - it doesnt accout for any scoping right now
	//so it needds to be improved once the ting works
	KnownTypeAlias   map[string][]*aliasEntry
	KnownGlobals     map[string]bool
	ProcessedSources map[string]*sourceReport
}

func (c *parserContext) GetErrors() []Report {
	ret := make([]Report, 0)
	for _, r := range c.Errors {
		ret = append(ret, r)
	}
	return ret
}
func (c *parserContext) GetWarnings() []Report {
	ret := make([]Report, 0)
	for _, r := range c.Warnings {
		ret = append(ret, r)
	}
	return ret
}

func (c *parserContext) GetGraph() *graph.SchemaGraph {
	return nil
}

type sourceReport struct {
	isMany           bool
	manyBy           []string
	requiredEntities []*requiredEntity
}

type requiredEntity struct {
	name      string
	many      bool
	predicate string
	scope     string
}

type lineReporter struct {
	Line           int
	ErrorMarkerPos int
	Message        string
	Source         string
	Code           SashimiErrorCode
}

func (l *lineReporter) OnLine() int {
	return l.Line
}

func (l *lineReporter) OnPos() int {
	return l.ErrorMarkerPos
}

func (l *lineReporter) InSource() string {
	return l.Source
}

func (l *lineReporter) InternalCode() int {
	return int(l.Code)
}

func (*parserContext) isGlobal(path string) bool {
	return strings.HasPrefix(path, "@")
}

func (c *parserContext) getUnderlyingEntityType(path string, callSource string, callScopes []string) string {
	propParts := strings.Split(path, ".")
	if len(propParts) > 0 {
		if _, ok := c.KnownGlobals[propParts[0]]; ok {
			return propParts[0]
		}
		if alias, ok := c.KnownTypeAlias[propParts[0]]; ok {
			for _, av := range alias {
				if av.Source == callSource {
					for _, scope := range callScopes {
						if isValidForScope(av.Scope, scope) {
							return av.UnderlyingType
						}
					}
				}
			}
		}
		return propParts[0]
	}
	return ""
}

func (c *parserContext) propertyExists(path string, callSource string, callScopes []string) bool {
	propParts := strings.Split(path, ".")
	if len(propParts) > 0 {
		if _, ok := c.KnownGlobals[propParts[0]]; ok {
			return true
		}
		if val, ok := c.Def[propParts[0]]; ok {
			for i := 1; i < len(propParts); i++ {
				t, ok := val.Properties[propParts[i]]
				if !ok {
					return false
				}
				kind := t.Type.Kind()
				switch kind {
				case SashimiScalar, SashimiUnion:
					return true
				case SashimiList:
					hkt := t.Type.HKT()
					if hkt == nil {
						return true
					}
					derefType := hkt.ResolveTypeName()
					val, ok = c.Def[derefType]
					if !ok {
						return false
					}
					break
				case SashimiReference:
					derefType := t.Type.ResolveTypeName()
					val, ok = c.Def[derefType]
					if !ok {
						return false
					}
					break
				}
			}
			return true
		}
		if alias, ok := c.KnownTypeAlias[propParts[0]]; ok {
			for _, av := range alias {
				if av.Source == callSource {
					for _, scope := range callScopes {
						if isValidForScope(av.Scope, scope) {
							propParts[0] = av.UnderlyingType
							return c.propertyExists(strings.Join(propParts, "."), callSource, callScopes)
						}
					}
				}
			}
		}
	}
	return false
}

func isValidForScope(aliasScope, contextScope string) bool {
	return strings.HasPrefix(contextScope, aliasScope)
}

func (c *parserContext) Consolidate() {
	for call, prop := range c.Calls {
		for propName, callCtx := range prop {
			for source, scopes := range callCtx {
				switch call {
				case "link", "repeat", "display":
					if !c.propertyExists(propName, source, scopes) {
						c.Errors = append(c.Errors, &lineReporter{
							Line:           0,
							ErrorMarkerPos: 0,
							Message:        fmt.Sprintf("Unknown property path: `%s`", propName),
							Source:         source,
							Code:           SashminiUnknownPropertyPath,
						})
					}
					break
				case "layout_section":
					if layout, ok := c.Calls["layout"]; ok {
						if _, ok := layout[propName]; !ok {
							c.Warnings = append(c.Warnings, &lineReporter{
								Line:           0,
								ErrorMarkerPos: 0,
								Message:        fmt.Sprintf("Unused layout section `%s`", propName),
								Source:         source,
								Code:           SashminiUnusedLayoutSection,
							})
						}
					} else {
						c.Warnings = append(c.Warnings, &lineReporter{
							Line:           0,
							ErrorMarkerPos: 0,
							Message:        fmt.Sprintf("Unused layout section `%s`", propName),
							Code:           SashminiUnusedLayoutSection,
							Source:         source,
						})
					}
					break
				case "layout":
					if layout, ok := c.Calls["layout_section"]; ok {
						if _, ok := layout[propName]; !ok {
							c.Errors = append(c.Errors, &lineReporter{
								Line:           0,
								ErrorMarkerPos: 0,
								Message:        fmt.Sprintf("Undefined layout section `%s`", propName),
								Source:         source,
								Code:           SashminiUndefinedLayoutSection,
							})
						}
					} else {
						c.Errors = append(c.Errors, &lineReporter{
							Line:           0,
							ErrorMarkerPos: 0,
							Message:        fmt.Sprintf("Undefined layout section `%s`", propName),
							Source:         source,
							Code:           SashminiUndefinedLayoutSection,
						})
					}
					break
				}
			}

		}

	}
	for name, source := range c.ProcessedSources {
		if source.requiredEntities != nil && len(source.requiredEntities) > 0 {
			for _, e := range source.requiredEntities {
				t := c.getUnderlyingEntityType(e.name, name, []string{e.scope})
				if c.isGlobal(t) {
					continue
				}
				val, ok := c.Def[t]
				if ok {
					if val.IsUnique {
						e.many = false
					}
				} else {
					c.Errors = append(c.Errors, &lineReporter{
						Line:           0,
						ErrorMarkerPos: 0,
						Message:        fmt.Sprintf("Unknown entity: `%s`", e.name),
						Source:         name,
						Code:           SashimiUnknownEntity,
					})
				}
			}
		}
	}
}
