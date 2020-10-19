package build

import (
	"io"
	"strings"
	"testing"
)

type testMockCompilerSource struct {
	f       string
	n       string
	content string
}

func (t *testMockCompilerSource) Load() (rc io.ReadCloser, err error) {
	return &stringReaderNoopCloser{strings.NewReader(t.content)}, nil
}

func (t *testMockCompilerSource) Name() string {
	return t.n
}
func (t *testMockCompilerSource) Format() string {
	return t.f
}

func TestAnalyzeEmptySourceRun(t *testing.T) {
	c := NewCompiler(&mockNodeResolver{})
	es := make([]CompilerSource, 0)
	r, err := c.Analyze(es)
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}
	if r == nil {
		t.Errorf("Unexpected NIL result for result")
	}
}

func TestAnalyzeSushieRun(t *testing.T) {
	c := NewCompiler(&mockNodeResolver{})
	es := make([]CompilerSource, 0)
	es = append(es, &testMockCompilerSource{
		f: "sushi",
		n: "testing-dummy.sushi",
		content: `sashimi:entity(project) of
		- title as "Project Title" is text
		sashimi:entity(employee) of
		- name as "Emp. Name" is text`,
	})
	r, err := c.Analyze(es)
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}
	if r == nil {
		t.Errorf("Unexpected NIL result for result")
	} else {
		errors := r.GetErrors()
		warnings := r.GetWarnings()
		if len(warnings) > 0 {
			t.Errorf("Unknown warning triggered (%v)", len(warnings))
		}
		if len(errors) > 0 {
			t.Errorf("Unkown errors triggered (%v)", len(errors))
		}
	}
}

func TestAnalyzeTriggerHTMLExtraction(t *testing.T) {
	c := NewCompiler(&mockNodeResolver{})
	es := make([]CompilerSource, 0)
	es = append(es, &testMockCompilerSource{
		f:       "html",
		n:       "testing-dummy.html",
		content: ``,
	})
	r, err := c.Analyze(es)
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}
	if r == nil {
		t.Errorf("Unexpected NIL result for result")
	} else {
		errors := r.GetErrors()
		warnings := r.GetWarnings()
		if len(warnings) > 0 {
			t.Errorf("Unknown warning triggered (%v)", len(warnings))
		}
		if len(errors) > 0 {
			t.Errorf("Unkown errors triggered (%v)", len(errors))
		}
	}
}

func TestAnalyzeHTMLExtraction(t *testing.T) {
	c := NewCompiler(&mockNodeResolver{})
	es := make([]CompilerSource, 0)
	es = append(es, &testMockCompilerSource{
		f:       "html",
		n:       "testing-dummy.html",
		content: TestValidHTMLUpper + `<!-- sashimi:display(something)--><br><h1><!-- sashimi:display(somethingElse)--></h1>` + TestValidHTMLLower,
	})
	r, err := c.Analyze(es)
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}
	if r == nil {
		t.Errorf("Unexpected NIL result for result")
	} else {
		errors := r.GetErrors()
		warnings := r.GetWarnings()
		if len(warnings) > 0 {
			t.Errorf("Unknown warning triggered (%v)", len(warnings))
		}
		if len(errors) != 4 {
			t.Errorf("Unkown errors triggered (%v)", len(errors))
			for _, e := range errors {
				t.Errorf("Unknown eror: %+v", e)
			}

		}
	}
}
