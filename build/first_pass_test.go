package build

import (
	"testing"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func TestSuccessfulEntityRun(t *testing.T) {
	t.Log("Testing a first pass on valid entity declaration")
	is := antlr.NewInputStream(`sashimi:entity(project) of
	- title as "Project Title" is text
	- description as "Description" is text[:uppercase]
	- image as "Main Image" is picture[res=400x200]
	- starCount as "Star Count" is number[:int,max=5]
	- category as "Category" is oneOf("Archtiecture","Misc")
	- tags as "Tags" is manyOf("Cheap","Fast","Good")
	- gallery as "Gallery" is list picture
	`)
	lexer := NewSashimiLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := NewSashimiParser(stream)
	firstPass := plainFirstPassParser()
	antlr.ParseTreeWalkerDefault.Walk(firstPass, p.Block())

	if len(firstPass.ctx.Errors) > 0 {
		t.Errorf("Sematic errors occured, test failed.")
	}
	if len(firstPass.ctx.Warnings) > 0 {
		t.Errorf("Sematic warnings occured, test failed.")
	}
	if val, ok := firstPass.ctx.Def["project"]; ok {
		if !val.IsDefined {
			t.Error("IsDefined not set on defined entity")
		}
		if val.Identifier != "project" {
			t.Error("Identifier not set correctly")
		}
		if len(val.Properties) != 7 {
			t.Error("Incorrect number of properties parsed")
		}
		if val.Properties["title"].Alias != `"Project Title"` {
			t.Errorf("Incorrect alias: %s", val.Properties["title"].Alias)
		}
	} else {
		t.Error("project entity has not been parsed")
	}
}

func TestSuccessfulEntitiesRun(t *testing.T) {
	t.Log("Testing a first pass on two valid entity declarations")
	is := antlr.NewInputStream(`sashimi:entity(project) of
	- title as "Project Title" is text
	sashimi:entity(employee) of
	- name as "Emp. Name" is text
	`)
	lexer := NewSashimiLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := NewSashimiParser(stream)
	firstPass := plainFirstPassParser()
	antlr.ParseTreeWalkerDefault.Walk(firstPass, p.Block())

	if len(firstPass.ctx.Errors) > 0 {
		t.Errorf("Sematic errors occured, test failed.")
	}
	if len(firstPass.ctx.Warnings) > 0 {
		t.Errorf("Sematic warnings occured, test failed.")
	}
	if val, ok := firstPass.ctx.Def["project"]; ok {
		if !val.IsDefined {
			t.Error("IsDefined not set on defined entity")
		}
		if val.Identifier != "project" {
			t.Error("Identifier not set correctly")
		}
		if len(val.Properties) != 1 {
			t.Error("Incorrect number of properties parsed")
		}
		if val.Properties["title"].Alias != `"Project Title"` {
			t.Errorf("Incorrect alias: %s", val.Properties["title"].Alias)
		}
	} else {
		t.Error("project entity has not been parsed")
	}
	if val, ok := firstPass.ctx.Def["employee"]; ok {
		if !val.IsDefined {
			t.Error("IsDefined not set on defined entity")
		}
		if val.Identifier != "employee" {
			t.Error("Identifier not set correctly")
		}
		if len(val.Properties) != 1 {
			t.Error("Incorrect number of properties parsed")
		}
		if val.Properties["name"].Alias != `"Emp. Name"` {
			t.Errorf("Incorrect alias: %s", val.Properties["name"].Alias)
		}
	} else {
		t.Error("employee entity has not been parsed")
	}
}

func TestDoubleDeclaredProperty(t *testing.T) {
	t.Log("Testing a first pass on valid entity declaration")
	is := antlr.NewInputStream(`sashimi:entity(project) of
	- title as "Project Title" is text
	- description as "Description" is text[:uppercase]
	- image as "Main Image" is picture[res=400x200]
	- image as "Second Image" is picture
	- starCount as "Star Count" is number[:int,max=5]
	- category as "Category" is oneOf("Archtiecture","Misc")
	- tags as "Tags" is manyOf("Cheap","Fast","Good")
	- gallery as "Gallery" is list picture
	`)
	lexer := NewSashimiLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := NewSashimiParser(stream)
	firstPass := plainFirstPassParser()
	antlr.ParseTreeWalkerDefault.Walk(firstPass, p.Block())

	if len(firstPass.ctx.Errors) != 1 {
		t.Errorf("Sematic error for duplicate property reporting failed.")
	}
	if len(firstPass.ctx.Warnings) > 0 {
		t.Errorf("Sematic warnings occured, test failed.")
	}
}

func TestDoubleDeclaredEntity(t *testing.T) {
	is := antlr.NewInputStream(`
	sashimi:entity(project) of
		- title as "Project Title" is text
	sashimi:entity(project) of
		- description as "Desc" is text
	`)
	lexer := NewSashimiLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := NewSashimiParser(stream)
	firstPass := plainFirstPassParser()
	antlr.ParseTreeWalkerDefault.Walk(firstPass, p.Block())

	if len(firstPass.ctx.Errors) != 0 {
		t.Errorf("Sematic error for duplicate property reporting failed.")
	}
	if len(firstPass.ctx.Warnings) != 1 {
		t.Errorf("Sematic warnings missing for double declared entity!")
	}
}

func TestConsolidation(t *testing.T) {
	is := antlr.NewInputStream(`sashimi:entity(page) of
	- title as "Pagetitle" is text
	- description as "Description" is text[:markdown]
	sashimi:display(page.description)
	`)
	lexer := NewSashimiLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := NewSashimiParser(stream)
	firstPass := plainFirstPassParser()
	antlr.ParseTreeWalkerDefault.Walk(firstPass, p.Block())
	firstPass.ctx.Consolidate()
	if len(firstPass.ctx.Errors) > 0 {
		t.Errorf("Consolidation created unwanted error. %v", firstPass.ctx.Errors[0])
	}
	if len(firstPass.ctx.Warnings) > 0 {
		t.Errorf("Consolidation created unwanted warning.")
	}
	if len(firstPass.ctx.ProcessedSources) != 1 {
		t.Error("Wrong source count encountered")
	} else {
		val, ok := firstPass.ctx.ProcessedSources["not-set"]
		if !ok {
			t.Error("Processed source not collected")
		}
		if (len(val.requiredEntities)) != 1 {
			t.Error("Required entities not collected correctly")
		} else {
			if val.requiredEntities[0].name != "page" {
				t.Errorf("Required entities wrongly collected, expected `page` but got `%s`", val.requiredEntities[0].name)
			}
			if !val.requiredEntities[0].many {
				t.Error("Required entities with wrong cardinality")
			}
			if val.requiredEntities[0].predicate != "" {
				t.Error("Required entities with wrong predicate")
			}
		}
	}
}

func TestConsolidationWithUniqueEntity(t *testing.T) {
	is := antlr.NewInputStream(`sashimi:unique entity(page) of
	- title as "Pagetitle" is text
	- description as "Description" is text[:markdown]
	sashimi:display(page.description)
	`)
	lexer := NewSashimiLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := NewSashimiParser(stream)
	firstPass := plainFirstPassParser()
	antlr.ParseTreeWalkerDefault.Walk(firstPass, p.Block())
	firstPass.ctx.Consolidate()
	if len(firstPass.ctx.Errors) > 0 {
		t.Errorf("Consolidation created unwanted error. %v", firstPass.ctx.Errors[0])
	}
	if len(firstPass.ctx.Warnings) > 0 {
		t.Errorf("Consolidation created unwanted warning.")
	}
	if len(firstPass.ctx.ProcessedSources) != 1 {
		t.Error("Wrong source count encountered")
	} else {
		val, ok := firstPass.ctx.ProcessedSources["not-set"]
		if !ok {
			t.Error("Processed source not collected")
		}
		if (len(val.requiredEntities)) != 1 {
			t.Error("Required entities not collected correctly")
		} else {
			if val.requiredEntities[0].name != "page" {
				t.Errorf("Required entities wrongly collected, expected `page` but got `%s`", val.requiredEntities[0].name)
			}
			if val.requiredEntities[0].many {
				t.Error("Required entities with wrong cardinality")
			}
			if val.requiredEntities[0].predicate != "" {
				t.Error("Required entities with wrong predicate")
			}
		}
	}
}

func TestConsolidationReference(t *testing.T) {
	is := antlr.NewInputStream(`
	sashimi:entity(pagepart) of
		- description as "Description" is text
	sashimi:entity(page) of
		- title as "Pagetitle" is text
		- part as "Description" is entity pagepart
	sashimi:display(page.part.description)
	`)
	lexer := NewSashimiLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := NewSashimiParser(stream)
	firstPass := plainFirstPassParser()
	antlr.ParseTreeWalkerDefault.Walk(firstPass, p.Block())
	firstPass.ctx.Consolidate()
	if len(firstPass.ctx.Errors) > 0 {
		t.Errorf("Consolidation created unwanted error. %v", firstPass.ctx.Errors[0])
	}
	if len(firstPass.ctx.Warnings) > 0 {
		t.Errorf("Consolidation created unwanted warning.")
	}
}

func TestRepeatExplicitHTMLScope(t *testing.T) {
	is := antlr.NewInputStream(`
	sashimi:entity(something) of
		- desc as "Description" is text
		sashimi:repeat(something) as a sashimi:begin('div:5') sashimi:begin('div:5::span:6') sashimi:display(a.desc) sashimi:end('div:5::span:6')sashimi:end('div:5')
	`)
	lexer := NewSashimiLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := NewSashimiParser(stream)
	firstPass := plainFirstPassParser()
	antlr.ParseTreeWalkerDefault.Walk(firstPass, p.Block())
	firstPass.ctx.Consolidate()
	if len(firstPass.ctx.Errors) > 0 {
		t.Errorf("Consolidation created unwanted error. %v", firstPass.ctx.Errors[0])
	}
	if len(firstPass.ctx.Warnings) > 0 {
		t.Errorf("Consolidation created unwanted warning.")
	}
	if len(firstPass.ctx.ProcessedSources) != 1 {
		t.Error("Processed sources didnt not get tracked")
	} else {
		val, ok := firstPass.ctx.ProcessedSources["not-set"]
		if !ok {
			t.Error("Wrong source tracked")
		} else if !val.isMany {
			t.Error("Wrong cardinality tracked")
		}
	}
}

func TestConsolidationListReference(t *testing.T) {
	is := antlr.NewInputStream(`
	sashimi:entity(pagepart) of
		- description as "Description" is text
	sashimi:entity(page) of
		- title as "Pagetitle" is text
		- part as "Description" is list entity pagepart
	sashimi:repeat(page.part.description) 
	sashimi:begin
	sashimi:end
	`)
	lexer := NewSashimiLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := NewSashimiParser(stream)
	firstPass := plainFirstPassParser()
	antlr.ParseTreeWalkerDefault.Walk(firstPass, p.Block())
	firstPass.ctx.Consolidate()
	if len(firstPass.ctx.Errors) > 0 {
		t.Errorf("Consolidation created unwanted error. %v", firstPass.ctx.Errors[0])
	}
	if len(firstPass.ctx.Warnings) > 0 {
		t.Errorf("Consolidation created unwanted warning.")
	}
}

func TestConsolidationListReferenceError(t *testing.T) {
	is := antlr.NewInputStream(`
	sashimi:entity(pagepart) of
		- description as "Description" is text
	sashimi:entity(page) of
		- title as "Pagetitle" is text
		- part as "Description" is list entity pagepart
	sashimi:repeat(page.part.nonexistent)
	sashimi:begin
	sashimi:end
	`)
	lexer := NewSashimiLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := NewSashimiParser(stream)
	firstPass := plainFirstPassParser()
	antlr.ParseTreeWalkerDefault.Walk(firstPass, p.Block())
	firstPass.ctx.Consolidate()
	if len(firstPass.ctx.Errors) == 1 {
		if firstPass.ctx.Errors[0].Message != "Unknown property path: `page.part.nonexistent`" {
			t.Errorf("Consolidation created unwanted error. %v", firstPass.ctx.Errors[0])
		}
	} else {
		t.Errorf("Consolidation created NO error when it should have.")
	}
	if len(firstPass.ctx.Warnings) > 0 {
		t.Errorf("Consolidation created unwanted warning.")
	}
}

func TestRepeatAliasNoScopeError(t *testing.T) {
	t.Skip("done on lexer base now")
	is := antlr.NewInputStream(`
	sashimi:entity(pagepart) of
		- title is text
		- description as "Description" is text
	sashimi:repeat(pagepart) as pp
	sashimi:display(pp.title)
	`)
	lexer := NewSashimiLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := NewSashimiParser(stream)
	firstPass := plainFirstPassParser()
	antlr.ParseTreeWalkerDefault.Walk(firstPass, p.Block())
	firstPass.ctx.Consolidate()
	if len(firstPass.ctx.Errors) > 0 {
		t.Errorf("Consolidation created unwanted error. %v", firstPass.ctx.Errors[0])
	}
	if len(firstPass.ctx.Warnings) > 0 {
		t.Errorf("Consolidation created unwanted warning.")
	}
}

func TestRepeatAliasOutOfScopeError(t *testing.T) {
	is := antlr.NewInputStream(`
	sashimi:entity(pagepart) of
		- title is text
		- description as "Description" is text
	sashimi:repeat(pagepart) as pp
	sashimi:begin
	sashimi:end
	sashimi:display(pp.title)
	`)
	lexer := NewSashimiLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := NewSashimiParser(stream)
	firstPass := plainFirstPassParser()
	antlr.ParseTreeWalkerDefault.Walk(firstPass, p.Block())
	firstPass.ctx.Consolidate()
	if len(firstPass.ctx.Errors) == 2 {
		if firstPass.ctx.Errors[0].Message != "Unknown property path: `pp.title`" {
			t.Errorf("Consolidation created unwanted error. %v", firstPass.ctx.Errors[0])
		}
	} else {
		t.Errorf("Consolidation created NO error when it should have.")
	}
	if len(firstPass.ctx.Warnings) > 0 {
		t.Errorf("Consolidation created unwanted warning.")
	}
}

func TestRepeatAliasExplicitScope(t *testing.T) {
	is := antlr.NewInputStream(`
	sashimi:entity(pagepart) of
		- title is text
		- description as "Description" is text
	sashimi:repeat(pagepart) as pp
	sashimi:begin('r:f')
	sashimi:display(pp.title)
	sashimi:end('r:f')
	`)
	lexer := NewSashimiLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := NewSashimiParser(stream)
	firstPass := plainFirstPassParser()
	antlr.ParseTreeWalkerDefault.Walk(firstPass, p.Block())
	firstPass.ctx.Consolidate()
	if len(firstPass.ctx.Errors) > 0 {
		t.Errorf("Consolidation created unwanted error. %v", firstPass.ctx.Errors[0])
	}
	if len(firstPass.ctx.Warnings) > 0 {
		t.Errorf("Consolidation created unwanted warning.")
	}
}

func TestRepeatAliasGlobalExplicitScope(t *testing.T) {
	is := antlr.NewInputStream(`
	sashimi:entity(pagepart) of
		- title is text
		- description as "Description" is text
	sashimi:repeat(@pages) as pp
	sashimi:begin('r:f')
	sashimi:display(pp.title)
	sashimi:end('r:f')
	`)
	lexer := NewSashimiLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := NewSashimiParser(stream)
	firstPass := plainFirstPassParser()
	antlr.ParseTreeWalkerDefault.Walk(firstPass, p.Block())
	firstPass.ctx.Consolidate()
	if len(firstPass.ctx.Errors) > 0 {
		t.Errorf("Consolidation created unwanted error. %v", firstPass.ctx.Errors[0])
	}
	if len(firstPass.ctx.Warnings) > 0 {
		t.Errorf("Consolidation created unwanted warning.")
	}
}

func TestRepeatAliasExplicitNestedScope(t *testing.T) {
	is := antlr.NewInputStream(`
	sashimi:entity(pagepart) of
		- title is text
		- description as "Description" is text
	sashimi:repeat(pagepart) as pp
	sashimi:begin('r')
	sashimi:begin('r:a')
	sashimi:begin('r:a:b')
	sashimi:display(pp.title)
	sashimi:end('r:a:b')
	sashimi:end('r:a')
	sashimi:end('r')
	`)
	lexer := NewSashimiLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := NewSashimiParser(stream)
	firstPass := plainFirstPassParser()
	antlr.ParseTreeWalkerDefault.Walk(firstPass, p.Block())
	firstPass.ctx.Consolidate()
	if len(firstPass.ctx.Errors) > 0 {
		t.Errorf("Consolidation created unwanted error. %v", firstPass.ctx.Errors[0])
	}
	if len(firstPass.ctx.Warnings) > 0 {
		t.Errorf("Consolidation created unwanted warning.")
	}
}

func TestRepeatAliasExplicitScopeNoEndTag(t *testing.T) {
	t.Skip("happends on token level need to fix up")
	is := antlr.NewInputStream(`
	sashimi:entity(pagepart) of
		- title is text
		- description as "Description" is text
	sashimi:repeat(pagepart) as pp
	sashimi:begin('r')
	sashimi:display(pp.title)
	`)
	lexer := NewSashimiLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := NewSashimiParser(stream)
	firstPass := plainFirstPassParser()
	antlr.ParseTreeWalkerDefault.Walk(firstPass, p.Block())
	firstPass.ctx.Consolidate()
	if len(firstPass.ctx.Errors) > 0 {
		t.Errorf("Consolidation created unwanted error. %v", firstPass.ctx.Errors[0])
	}
	if len(firstPass.ctx.Warnings) > 0 {
		t.Errorf("Consolidation created unwanted warning.")
	}
}

func TestRepeatAliasImplicitScope(t *testing.T) {
	is := antlr.NewInputStream(`
	sashimi:entity(pagepart) of
		- title is text
		- description as "Description" is text
	sashimi:repeat(pagepart) as pp
	sashimi:begin
	sashimi:display(pp.title)
	sashimi:end
	`)
	lexer := NewSashimiLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := NewSashimiParser(stream)
	firstPass := plainFirstPassParser()
	antlr.ParseTreeWalkerDefault.Walk(firstPass, p.Block())
	firstPass.ctx.Consolidate()
	if len(firstPass.ctx.Errors) > 0 {
		t.Errorf("Consolidation created unwanted error. %v", firstPass.ctx.Errors[0])
	}
	if len(firstPass.ctx.Warnings) > 0 {
		t.Errorf("Consolidation created unwanted warning.")
	}
}

func TestRepeatAliasPredicate(t *testing.T) {

	is := antlr.NewInputStream(`
	sashimi:entity(pagepart) of
		- description as "Description" is text
		- visible is bool
	sashimi:repeat(pagepart) as pp [x -> x]
	sashimi:begin
	sashimi:display(pp.description)
	sashimi:end
	`)
	lexer := NewSashimiLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := NewSashimiParser(stream)
	firstPass := plainFirstPassParser()
	antlr.ParseTreeWalkerDefault.Walk(firstPass, p.Block())
	firstPass.ctx.Consolidate()
	if len(firstPass.ctx.Errors) > 0 {
		t.Errorf("Consolidation created unwanted error. %v", firstPass.ctx.Errors[0])
	}
	if len(firstPass.ctx.Warnings) > 0 {
		t.Errorf("Consolidation created unwanted warning.")
	}

}

func TestConsolidationFailedEntity(t *testing.T) {
	is := antlr.NewInputStream(`sashimi:entity(page) of
	- title as "Pagetitle" is text
	- description as "Description" is text[:markdown]
	sashimi:display(nonexistent.title)
	`)
	lexer := NewSashimiLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := NewSashimiParser(stream)
	firstPass := plainFirstPassParser()
	antlr.ParseTreeWalkerDefault.Walk(firstPass, p.Block())
	firstPass.ctx.Consolidate()
	if len(firstPass.ctx.Errors) == 2 {
		if firstPass.ctx.Errors[0].Message != "Unknown property path: `nonexistent.title`" {
			t.Errorf("Consolidation created unwanted error. %v", firstPass.ctx.Errors[0])
		}
	} else {
		t.Errorf("Consolidation created NO error when it should have.")
	}
	if len(firstPass.ctx.Warnings) > 0 {
		t.Errorf("Consolidation created unwanted warning.")
	}
}

func TestConsolidationFailedProperty(t *testing.T) {
	is := antlr.NewInputStream(`sashimi:entity(page) of
	- title as "Pagetitle" is text
	- description as "Description" is text[:markdown]
	sashimi:display(page.nonexistent)
	`)
	lexer := NewSashimiLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := NewSashimiParser(stream)
	firstPass := plainFirstPassParser()
	antlr.ParseTreeWalkerDefault.Walk(firstPass, p.Block())
	firstPass.ctx.Consolidate()
	if len(firstPass.ctx.Errors) == 1 {
		if firstPass.ctx.Errors[0].Message != "Unknown property path: `page.nonexistent`" {
			t.Errorf("Consolidation created WRONG error.")
		}
	} else {
		t.Errorf("Consolidation created NO error when it should have.")
	}
	if len(firstPass.ctx.Warnings) > 0 {
		t.Errorf("Consolidation created unwanted warning.")
	}
}

func TestConsolidationUndefinedDisplay(t *testing.T) {
	is := antlr.NewInputStream(`sashimi:display(project.nonexistent)`)
	lexer := NewSashimiLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := NewSashimiParser(stream)
	firstPass := plainFirstPassParser()
	antlr.ParseTreeWalkerDefault.Walk(firstPass, p.Block())
	firstPass.ctx.Consolidate()
	if len(firstPass.ctx.Errors) == 2 {
		if firstPass.ctx.Errors[0].Message != "Unknown property path: `project.nonexistent`" {
			t.Errorf("Consolidation created WRONG error (%s).", firstPass.ctx.Errors[0].Message)
		}
	} else {
		t.Errorf("Consolidation created NO error when it should have.")
	}
	if len(firstPass.ctx.Warnings) > 0 {
		t.Errorf("Consolidation created unwanted warning.")
	}
}

func TestConsolidationUndefinedRepeat(t *testing.T) {
	is := antlr.NewInputStream(`sashimi:repeat(project) sashimi:begin sashimi:end`)
	lexer := NewSashimiLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := NewSashimiParser(stream)
	firstPass := plainFirstPassParser()
	antlr.ParseTreeWalkerDefault.Walk(firstPass, p.Block())
	firstPass.ctx.Consolidate()
	if len(firstPass.ctx.Errors) == 2 {
		if firstPass.ctx.Errors[0].Message != "Unknown property path: `project`" {
			t.Errorf("Consolidation created WRONG error (%s).", firstPass.ctx.Errors[0].Message)
		}
	} else {
		t.Errorf("Consolidation created NO error when it should have.")
	}
	if len(firstPass.ctx.Warnings) > 0 {
		t.Errorf("Consolidation created unwanted warning.")
	}
}

func TestConsolidationUndefinedLayout(t *testing.T) {
	is := antlr.NewInputStream(`sashimi:layout(main)`)
	lexer := NewSashimiLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := NewSashimiParser(stream)
	firstPass := plainFirstPassParser()
	antlr.ParseTreeWalkerDefault.Walk(firstPass, p.Block())
	firstPass.ctx.Consolidate()
	if len(firstPass.ctx.Errors) == 1 {
		if firstPass.ctx.Errors[0].Message != "Undefined layout section `main`" {
			t.Errorf("Consolidation created WRONG error (%s).", firstPass.ctx.Errors[0].Message)
		}
	} else {
		t.Errorf("Consolidation created NO error when it should have.")
	}
	if len(firstPass.ctx.Warnings) > 0 {
		t.Errorf("Consolidation created unwanted warning.")
	}
}

func TestConsolidationUnusedLayoutSection(t *testing.T) {
	is := antlr.NewInputStream(`sashimi:layout_section(main)`)
	lexer := NewSashimiLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := NewSashimiParser(stream)
	firstPass := plainFirstPassParser()
	antlr.ParseTreeWalkerDefault.Walk(firstPass, p.Block())
	firstPass.ctx.Consolidate()
	if len(firstPass.ctx.Errors) != 0 {
		t.Errorf("Consolidation created an error where it should not have.")
	}
	if len(firstPass.ctx.Warnings) > 0 {
		if firstPass.ctx.Warnings[0].Message != "Unused layout section `main`" {
			t.Errorf("Consolidation created WRONG warning (%s).", firstPass.ctx.Warnings[0].Message)
		}
	} else {
		t.Errorf("Consolidation created NO warning where it should have.")
	}
}

func TestConsolidationUndefinedLink(t *testing.T) {
	is := antlr.NewInputStream(`sashimi:link(user)`)
	lexer := NewSashimiLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := NewSashimiParser(stream)
	firstPass := plainFirstPassParser()
	antlr.ParseTreeWalkerDefault.Walk(firstPass, p.Block())
	firstPass.ctx.Consolidate()
	if len(firstPass.ctx.Errors) == 2 {
		if firstPass.ctx.Errors[0].Message != "Unknown property path: `user`" {
			t.Errorf("Consolidation created WRONG error (%s).", firstPass.ctx.Errors[0].Message)
		}
	} else {
		t.Errorf("Consolidation created NO error when it should have.")
	}
	if len(firstPass.ctx.Warnings) > 0 {
		t.Errorf("Consolidation created unwanted warning.")
	}
}

func TestConsolidationLayoutAndLayoutSectioInSameBlock(t *testing.T) {
	is := antlr.NewInputStream(`sashimi:layout_section(main) sashimi:layout(main)`)
	lexer := NewSashimiLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := NewSashimiParser(stream)
	firstPass := plainFirstPassParser()
	antlr.ParseTreeWalkerDefault.Walk(firstPass, p.Block())
	firstPass.ctx.Consolidate()
	if len(firstPass.ctx.Errors) > 0 {
		t.Errorf("Consolidate created a unwanted error.")
	}
	if len(firstPass.ctx.Warnings) > 0 {
		t.Errorf("Consolidate created a unwanted warning.")
	}
}
