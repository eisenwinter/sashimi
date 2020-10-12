package build

import (
	"testing"

	"github.com/eisenwinter/sashimi/graph"
)

func TestGraphGenEmpty(t *testing.T) {
	e := make(map[string]*defTableEntry)
	graph := createNodes(e)
	if len(graph.Nodes) != 0 {
		t.Error("Unknown node generation in graph")
	}
}

func TestGraphSingleEntitySingleProperty(t *testing.T) {
	e := make(map[string]*defTableEntry)
	e["project"] = &defTableEntry{
		Identifier: "project",
		IsDefined:  true,
		IsUnique:   false,
		Properties: make(map[string]*propertyDef),
	}
	titleProp := &propertyDef{
		Alias: "Title",
		Type: &scalarType{
			name:       "text",
			constraint: make(map[string]string),
		},
	}
	e["project"].Properties["title"] = titleProp
	g := createNodes(e)
	if len(g.Nodes) != 1 {
		t.Error("Unknown node generation in graph")
	} else {
		if g.Nodes[0].Name != "project" {
			t.Error("Wrong node name generated")
		}
		if g.Nodes[0].ConstraintTo != graph.Many {
			t.Error("Wrong node cardinalilty generated")
		}
		if len(g.Nodes[0].Fields) == 1 {
			if g.Nodes[0].Fields[0].Name != "title" {
				t.Error("Unexpected Field Name")
			}
			if g.Nodes[0].Fields[0].Alias != "Title" {
				t.Error("Unexpected Field Alias")
			}
			if g.Nodes[0].Fields[0].Type != "text" {
				t.Error("Unexpected Field Type")
			}
			if len(g.Nodes[0].Fields[0].Atoms) > 0 {
				t.Error("Unexpected Atoms")
			}
			if len(g.Nodes[0].Fields[0].KeyValues) > 0 {
				t.Error("Unexpected KeyValue Pairs")
			}
		} else {
			t.Error("Wrong field count generated")
		}
		if len(g.Nodes[0].Edges) > 0 {
			t.Error("Wrong edges generated")
		}
	}
}

func TestGraphSingleUniqueEntitySingleProperty(t *testing.T) {
	e := make(map[string]*defTableEntry)
	e["project"] = &defTableEntry{
		Identifier: "project",
		IsDefined:  true,
		IsUnique:   true,
		Properties: make(map[string]*propertyDef),
	}
	titleProp := &propertyDef{
		Alias: "Title",
		Type: &scalarType{
			name:       "text",
			constraint: make(map[string]string),
		},
	}
	e["project"].Properties["title"] = titleProp
	g := createNodes(e)
	if len(g.Nodes) != 1 {
		t.Error("Unknown node generation in graph")
	} else {
		if g.Nodes[0].Name != "project" {
			t.Error("Wrong node name generated")
		}
		if g.Nodes[0].ConstraintTo != graph.One {
			t.Error("Wrong node cardinalilty generated")
		}
		if len(g.Nodes[0].Fields) == 1 {
			if g.Nodes[0].Fields[0].Name != "title" {
				t.Error("Unexpected Field Name")
			}
			if g.Nodes[0].Fields[0].Alias != "Title" {
				t.Error("Unexpected Field Alias")
			}
			if g.Nodes[0].Fields[0].Type != "text" {
				t.Error("Unexpected Field Type")
			}
			if len(g.Nodes[0].Fields[0].Atoms) > 0 {
				t.Error("Unexpected Atoms")
			}
			if len(g.Nodes[0].Fields[0].KeyValues) > 0 {
				t.Error("Unexpected KeyValue Pairs")
			}
		} else {
			t.Error("Wrong field count generated")
		}
		if len(g.Nodes[0].Edges) > 0 {
			t.Error("Wrong edges generated")
		}
	}
}

func TestGraphWithOneEdge(t *testing.T) {
	e := make(map[string]*defTableEntry)
	e["project"] = &defTableEntry{
		Identifier: "project",
		IsDefined:  true,
		IsUnique:   false,
		Properties: make(map[string]*propertyDef),
	}
	titleProp := &propertyDef{
		Alias: "Title",
		Type: &scalarType{
			name:       "text",
			constraint: make(map[string]string),
		},
	}
	e["project"].Properties["title"] = titleProp

	e["employee"] = &defTableEntry{
		Identifier: "employee",
		IsDefined:  true,
		IsUnique:   false,
		Properties: make(map[string]*propertyDef),
	}
	nameProp := &propertyDef{
		Alias: "Name",
		Type: &scalarType{
			name:       "text",
			constraint: make(map[string]string),
		},
	}
	e["employee"].Properties["name"] = nameProp
	projectsProp := &propertyDef{
		Alias: "Projects",
		Type:  newList(newRef("project")),
	}
	e["employee"].Properties["projects"] = projectsProp

	g := createNodes(e)
	if len(g.Nodes) != 2 {
		t.Error("Unknown node generation in graph")
	} else {
		var projectNode *graph.Node
		var employeeNode *graph.Node
		if g.Nodes[0].Name == "project" {
			projectNode = g.Nodes[0]
			employeeNode = g.Nodes[1]
		} else {
			projectNode = g.Nodes[1]
			employeeNode = g.Nodes[0]
		}
		if projectNode.Name != "project" {
			t.Error("Wrong node name generated")
		}
		if projectNode.ConstraintTo != graph.Many {
			t.Error("Wrong node cardinalilty generated")
		}
		if len(projectNode.Fields) == 1 {
			if projectNode.Fields[0].Name != "title" {
				t.Error("Unexpected Field Name")
			}
			if projectNode.Fields[0].Alias != "Title" {
				t.Error("Unexpected Field Alias")
			}
			if projectNode.Fields[0].Type != "text" {
				t.Error("Unexpected Field Type")
			}
			if len(projectNode.Fields[0].Atoms) > 0 {
				t.Error("Unexpected Atoms")
			}
			if len(projectNode.Fields[0].KeyValues) > 0 {
				t.Error("Unexpected KeyValue Pairs")
			}
		} else {
			t.Error("Wrong field count generated")
		}
		if len(projectNode.Edges) > 0 {
			t.Error("Wrong edges generated")
		}

		if len(employeeNode.Edges) != 1 {
			t.Error("Edge not created")
		} else {
			if employeeNode.Edges[0].Cardinality != graph.Many {
				t.Error("Unexpected edge cardinality")
			}
			if employeeNode.Edges[0].From != "employee.projects" {
				t.Error("Unexpected edge from")
			}
			if employeeNode.Edges[0].To != "project" {
				t.Error("Unexpected edge to")
			}

		}
	}
}
