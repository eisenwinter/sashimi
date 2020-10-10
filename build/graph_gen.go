package build

import (
	"fmt"
	"strings"

	"github.com/eisenwinter/sashimi/graph"
)

func createNodes(def map[string]*defTableEntry) *graph.SchemaGraph {
	nodes := make([]*graph.Node, 0)
	for k, v := range def {
		if v.IsDefined {
			node := &graph.Node{
				Name:   k,
				Fields: make([]*graph.Field, 0),
				Edges:  make([]*graph.Edge, 0),
			}
			for n, p := range v.Properties {
				typeInfo := p.Type
				switch typeInfo.Kind() {
				case SashimiScalar:
					field := &graph.Field{
						Name:      n,
						Alias:     strings.Trim(p.Alias, "\""),
						Type:      typeInfo.ResolveTypeName(),
						Atoms:     make([]string, 0),
						KeyValues: make(map[string]string),
					}
					meta := typeInfo.MetaMap()
					if meta != nil {
						for j, metaval := range meta {
							if k == "atomic" && strings.HasPrefix(j, ":") {
								field.Atoms = append(field.Atoms, strings.TrimPrefix(j, ":"))
							} else {
								field.KeyValues[j] = metaval
							}
						}
					}
					node.Fields = append(node.Fields, field)
					break
				case SashimiUnion:
					field := &graph.Field{
						Name:      n,
						Alias:     strings.Trim(p.Alias, "\""),
						Type:      typeInfo.ResolveTypeName(),
						Atoms:     make([]string, 0),
						KeyValues: make(map[string]string),
					}
					meta := typeInfo.MetaMap()
					_, ok := meta["@@union_excl"]
					delete(meta, "@@union_excl")
					if ok {
						field.KeyValues["excl"] = "true"
					}
					edgeName := fmt.Sprintf("@@%s__%s", k, n)
					node.Edges = append(node.Edges, &graph.Edge{
						To:          edgeName,
						Cardinality: graph.Many,
					})
					for _, metaval := range meta {
						field.Atoms = append(field.Atoms, metaval)
					}
					valueField := &graph.Field{
						Name:  "value",
						Alias: "",
						Type:  string(SashimiTypeText),
					}
					nodes = append(nodes, &graph.Node{
						Name:   edgeName,
						Fields: []*graph.Field{valueField},
					})
					node.Fields = append(node.Fields, field)
					break
				case SashimiReference:
					node.Edges = append(node.Edges, &graph.Edge{
						To:          fmt.Sprintf("%s.%s->%s", k, n, typeInfo.ResolveTypeName()),
						Cardinality: graph.One,
					})
					break
				case SashimiList:
					hkt := typeInfo.HKT()
					if hkt.Kind() == SashimiScalar {
						edgeName := fmt.Sprintf("@@%s__%s", k, n)
						node.Edges = append(node.Edges, &graph.Edge{
							To:          edgeName,
							Cardinality: graph.Many,
						})
						valueField := &graph.Field{
							Name:  "value",
							Alias: "",
							Type:  string(hkt.Type()),
						}
						nodes = append(nodes, &graph.Node{
							Name:   edgeName,
							Fields: []*graph.Field{valueField},
						})
					}
					if hkt.Kind() == SashimiReference {
						node.Edges = append(node.Edges, &graph.Edge{
							To:          fmt.Sprintf("%s.%s->%s", k, n, hkt.ResolveTypeName()),
							Cardinality: graph.Many,
						})
					}
					break
				default:
					panic("Invalid type kind in graph generation (this should not happend. at all)")
				}
			}
			nodes = append(nodes, node)
		}

	}
	return &graph.SchemaGraph{nodes}
}
