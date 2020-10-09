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
	antlr.ParseTreeWalkerDefault.Walk(firstPass, p.Export())

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
		if val.Identfier != "project" {
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
		if val.Identfier != "project" {
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
		if val.Identfier != "employee" {
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
	antlr.ParseTreeWalkerDefault.Walk(firstPass, p.Export())

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

func TestConsolidationListReference(t *testing.T) {
	is := antlr.NewInputStream(`
	sashimi:entity(pagepart) of
		- description as "Description" is text
	sashimi:entity(page) of
		- title as "Pagetitle" is text
		- part as "Description" is list entity pagepart
	sashimi:repeat(page.part.description)
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
	if len(firstPass.ctx.Errors) == 1 {
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
	if len(firstPass.ctx.Errors) == 1 {
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
	is := antlr.NewInputStream(`sashimi:repeat(project)`)
	lexer := NewSashimiLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := NewSashimiParser(stream)
	firstPass := plainFirstPassParser()
	antlr.ParseTreeWalkerDefault.Walk(firstPass, p.Block())
	firstPass.ctx.Consolidate()
	if len(firstPass.ctx.Errors) == 1 {
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
	if len(firstPass.ctx.Errors) == 1 {
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
