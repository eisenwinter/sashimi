package main

import (
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/eisenwinter/sashimi/parser"
)

func main() {
	fmt.Println("sashimi standalone parser")

	is := antlr.NewInputStream(`sashimi:entity(project) of
	- title as "Project Title" is text
	- description as "Description" is text[:uppercase]
	- image as "Main Image" is picture[res=400x200]
	- starCount as "Star Count" is number[:int,max=5]
	- category as "Category" is oneOf("Archtiecture","Misc")
	- tags as "Tags" is manyOf("Cheap","Fast","Good")
	- gallery as "Gallery" is list picture
	`)
	//is := antlr.NewInputStream(`sashimi:layout_section(main)`)
	lexer := parser.NewSashimiLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewSashimiParser(stream)
	firstPass := parser.NewFirstPassParser()
	antlr.ParseTreeWalkerDefault.Walk(firstPass, p.Block())
	firstPass.Dump()
}
