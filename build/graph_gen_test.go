package build

import "testing"

func TestGraphGenEmpty(t *testing.T) {
	e := make(map[string]*defTableEntry)
	graph := createNodes(e)
	if len(graph.Nodes) != 0 {
		t.Error("Unknown node generation in graph")
	}
}
