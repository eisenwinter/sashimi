package build

import (
	"strings"
	"testing"
)

const TestValidHTMLUpper = `<!doctype html>
<html class="no-js" lang="">
<head>
  <meta charset="utf-8">
  <title></title>
  <meta name="description" content="">
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <meta property="og:title" content="">
  <meta property="og:type" content="">
  <meta property="og:url" content="">
  <meta property="og:image" content="">
</head>
<body>`

const TestValidHTMLLower = `
</body>
</html>
`

func TestValidHTML(t *testing.T) {
	r := strings.NewReader(TestValidHTMLLower + TestValidHTMLLower)
	p := &htmlProcessor{}
	var builder strings.Builder
	err := p.extractFromHTML(r, &builder)
	if err != nil {
		t.Error(err)
	}
	if builder.Len() > 0 {
		t.Errorf("Unexpected builder length of %v", builder.Len())
	}
}

func TestMalformedHTML(t *testing.T) {
	t.Skip("Not done yet")
	r := strings.NewReader("lel this aint html </div>")
	p := &htmlProcessor{}
	var builder strings.Builder
	err := p.extractFromHTML(r, &builder)
	if err == nil {
		t.Error("Expected error did not happend.")
	}
}

func TestBasicCommentExtraction(t *testing.T) {
	r := strings.NewReader(TestValidHTMLLower + `<!-- this is a comment -->` + TestValidHTMLLower)
	p := &htmlProcessor{}
	var builder strings.Builder
	err := p.extractFromHTML(r, &builder)
	if err != nil {
		t.Error(err)
	}
	if builder.Len() > 0 {
		t.Errorf("Unexpected builder length of %v", builder.Len())
	}
}

func TestBasicCommentWithSashimiDirectiveExtraction(t *testing.T) {
	r := strings.NewReader(TestValidHTMLLower + `<!-- sashimi:display(something)-->` + TestValidHTMLLower)
	p := &htmlProcessor{}
	var builder strings.Builder
	err := p.extractFromHTML(r, &builder)
	if err != nil {
		t.Error(err)
	}
	if builder.Len() != 27 {
		t.Errorf("Unexpected builder length of %v", builder.Len())
	} else {
		str := builder.String()
		if strings.TrimSpace(str) != "sashimi:display(something)" {
			t.Errorf("Unexpected extraction result %v", str)
		}
	}
}

func TestMultipleCommentsWithSashimiDirectiveExtraction(t *testing.T) {
	r := strings.NewReader(TestValidHTMLLower + `<!-- sashimi:display(something)--><br><h1><!-- sashimi:display(somethingElse)--></h1>` + TestValidHTMLLower)
	p := &htmlProcessor{}
	var builder strings.Builder
	err := p.extractFromHTML(r, &builder)
	if err != nil {
		t.Error(err)
	}
	if builder.Len() != 58 {
		t.Errorf("Unexpected builder length of %v", builder.Len())
	} else {
		str := builder.String()
		if strings.TrimSpace(str) != "sashimi:display(something) sashimi:display(somethingElse)" {
			t.Errorf("Unexpected extraction result %v", str)
		}
	}
}

func TestBasicRepeatScope(t *testing.T) {
	r := strings.NewReader(TestValidHTMLLower + `<!--  sashimi:repeat(something) as a --><div><span><!-- sashimi:display(a) --></span></div>` + TestValidHTMLLower)
	p := &htmlProcessor{}
	var builder strings.Builder
	err := p.extractFromHTML(r, &builder)
	if err != nil {
		t.Error(err)
	}
	if builder.Len() != 161 {
		t.Errorf("Unexpected builder length of %v", builder.Len())
		t.Errorf("Unexpected builder content %v", builder.String())
	}
}

func TestTransformBasicSiteProjects(t *testing.T) {
	val, ok := testDummyProject["projects.html"]
	if !ok {
		t.Error("Could not load dummy project data `projects.html`")
	}
	r := strings.NewReader(val)
	p := &htmlProcessor{}
	var builder strings.Builder
	err := p.transform(r, &builder, false)
	if err != nil {
		t.Error(err)
	}
	if builder.Len() != 212 {
		t.Errorf("Unexpected builder length of %v", builder.Len())
		t.Errorf("Input file used: %s", val)
		t.Errorf("Unexpected builder content %v", builder.String())
	}
}

func TestTransformBasicSiteProject(t *testing.T) {
	val, ok := testDummyProject["project.html"]
	if !ok {
		t.Error("Could not load dummy project data `project.html`")
	}
	r := strings.NewReader(val)
	p := &htmlProcessor{}
	var builder strings.Builder
	err := p.transform(r, &builder, false)
	if err != nil {
		t.Error(err)
	}
	if builder.Len() != 143 {
		t.Errorf("Unexpected builder length of %v", builder.Len())
		t.Errorf("Input file used: %s", val)
		t.Errorf("Unexpected builder content %v", builder.String())
	}
}
