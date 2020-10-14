package build

import "github.com/eisenwinter/sashimi/resolver"

type mockNodeResolver struct{}

func (m *mockNodeResolver) ResolveByID(nodeType string, id int) (map[string]interface{}, error) {
	panic("nan")
}

func (m *mockNodeResolver) ResolveAll(nodeType string) ([]map[string]interface{}, error) {
	panic("nan")
}

func (m *mockNodeResolver) Resolve(nodeType string, predicate resolver.Predicate) ([]map[string]interface{}, error) {
	panic("nan")
}

var dummyProject = map[string]string{
	"page.html":           "page",
	"projects.html":       "projects",
	"project.html":        "project",
	"shared/_layout.html": "layout"}
