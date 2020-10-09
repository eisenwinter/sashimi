package build

import "testing"

func TestGraphGenEmpty(t *testing.T) {
	e := make(map[string]*defTableEntry)
	nodes := createNodes(e)
	if len(nodes) != 0 {
		t.Error("Unknown node generation in graph")
	}
}
