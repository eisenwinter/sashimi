package build

import (
	"bytes"
	"fmt"
	"strings"
	"testing"

	"golang.org/x/net/html"
)

func TestValidHtmlWithoutDirective(t *testing.T) {
	val := TestValidHTMLUpper + TestValidHTMLLower
	htmlGen := &htmlGenerator{}
	r := strings.NewReader(val)
	dom, err := htmlGen.parseDocument(r)
	if err != nil {
		t.Errorf("Could not parse HTML document: %s", val)
	}
	rootNode := htmlGen.copyNode(dom)
	htmlGen.traverse(dom, rootNode)
	buf := new(bytes.Buffer)
	err = html.Render(buf, rootNode)
	if err != nil {
		t.Errorf("Could not render processed HTML docuemnt: %s", val)
	}
	result := strings.TrimSpace(buf.String())
	if len(result) == 0 {
		t.Error("Render output is empty")
	}
}

func TestValidHtmlNestedWithoutDirective(t *testing.T) {
	val := TestValidHTMLUpper + "<div><div><div><div><h1>hello</h1></div></div></div></div>" + TestValidHTMLLower
	htmlGen := &htmlGenerator{}
	r := strings.NewReader(val)
	dom, err := htmlGen.parseDocument(r)
	if err != nil {
		t.Errorf("Could not parse HTML document: %s", val)
	}
	rootNode := htmlGen.copyNode(dom)
	htmlGen.traverse(dom, rootNode)
	buf := new(bytes.Buffer)
	err = html.Render(buf, rootNode)
	if err != nil {
		t.Errorf("Could not render processed HTML docuemnt: %s", val)
	}
	result := strings.TrimSpace(buf.String())
	if len(result) == 0 {
		t.Error("Render output is empty")
	}
	if !strings.Contains(result, "<div><div><div><div><h1>hello</h1></div></div></div></div>") {
		t.Error("Render output is missing nested elements")
	}
}
func TestValidHtmlNestedTwiceWithoutDirective(t *testing.T) {
	val := TestValidHTMLUpper + "<div><div><div><div><h1>hello</h1></div></div></div></div><div><div><div><div><h1>hello</h1></div></div></div></div>" + TestValidHTMLLower
	htmlGen := &htmlGenerator{}
	r := strings.NewReader(val)
	dom, err := htmlGen.parseDocument(r)
	if err != nil {
		t.Errorf("Could not parse HTML document: %s", val)
	}
	rootNode := htmlGen.copyNode(dom)
	htmlGen.traverse(dom, rootNode)
	buf := new(bytes.Buffer)
	err = html.Render(buf, rootNode)
	if err != nil {
		t.Errorf("Could not render processed HTML docuemnt: %s", val)
	}
	result := strings.TrimSpace(buf.String())
	if len(result) == 0 {
		t.Error("Render output is empty")
	}
	if !strings.Contains(result, "<div><div><div><div><h1>hello</h1></div></div></div></div><div><div><div><div><h1>hello</h1></div></div></div></div>") {
		t.Error("Render output is missing nested elements")
	}
}

func TestProjectSite(t *testing.T) {
	val, ok := testDummyProject["project.html"]
	if !ok {
		t.Error("Could not load dummy project data `project.html`")
	}
	htmlGen := &htmlGenerator{}
	r := strings.NewReader(val)
	dom, err := htmlGen.parseDocument(r)
	if err != nil {
		t.Errorf("Could not parse HTML document: %s", val)
	}
	rootNode := htmlGen.copyNode(dom)
	htmlGen.traverse(dom, rootNode)
	buf := new(bytes.Buffer)
	err = html.Render(buf, rootNode)
	if err != nil {
		t.Errorf("Could not render processed HTML docuemnt: %s", val)
	}
	result := strings.TrimSpace(buf.String())
	if len(result) == 0 {
		t.Error("Render output is empty")
	}

}

func TestProjectsSite(t *testing.T) {
	val, ok := testDummyProject["projects.html"]
	if !ok {
		t.Error("Could not load dummy project data `projects.html`")
	}
	htmlGen := &htmlGenerator{}
	r := strings.NewReader(val)
	dom, err := htmlGen.parseDocument(r)
	if err != nil {
		t.Errorf("Could not parse HTML document: %s", val)
	}
	rootNode := htmlGen.copyNode(dom)
	htmlGen.traverse(dom, rootNode)
	buf := new(bytes.Buffer)
	err = html.Render(buf, rootNode)
	if err != nil {
		t.Errorf("Could not render processed HTML docuemnt: %s", val)
	}
	result := strings.TrimSpace(buf.String())
	fmt.Println(result)
	if len(result) == 0 {
		t.Error("Render output is empty")
	}

}
