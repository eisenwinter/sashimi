package parser

import (
	"fmt"
	"strings"
)

type parserContext struct {
	Def      map[string]*defTableEntry
	Errors   []*lineReporter
	Warnings []*lineReporter
	Calls    map[string]map[string]bool
}

type lineReporter struct {
	Line           int
	ErrorMarkerPos int
	Message        string
}

func (c *parserContext) propertyExists(path string) bool {
	propParts := strings.Split(path, ".")
	if len(propParts) > 1 {
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
	}
	return false
}

func (c *parserContext) Consolidate() {
	for call, prop := range c.Calls {
		for propName := range prop {
			switch call {
			case "link", "repeat", "display":
				if !c.propertyExists(propName) {
					c.Errors = append(c.Errors, &lineReporter{
						Line:           0,
						ErrorMarkerPos: 0,
						Message:        fmt.Sprintf("Unknown property path: `%s`", propName),
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
						})
					}
				} else {
					c.Warnings = append(c.Warnings, &lineReporter{
						Line:           0,
						ErrorMarkerPos: 0,
						Message:        fmt.Sprintf("Unused layout section `%s`", propName),
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
						})
					}
				} else {
					c.Errors = append(c.Errors, &lineReporter{
						Line:           0,
						ErrorMarkerPos: 0,
						Message:        fmt.Sprintf("Undefined layout section `%s`", propName),
					})
				}
				break
			}
		}

	}
}
