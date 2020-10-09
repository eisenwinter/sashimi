package build

import "testing"

func TestAnalyzeEmptySourceRun(t *testing.T) {
	c := NewCompiler()
	es := make([]CompilerSource, 0)
	r, err := c.Analyze(es, HyperCritical&OptimizeOff)
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}
	if r == nil {
		t.Errorf("Unexpected NIL result for result")
	}
}
