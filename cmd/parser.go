package main

import (
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/eisenwinter/sashimi/parser"
)

type dummyListener struct {
	*parser.BaseSashimiListener
}

func main() {
	fmt.Println("sashimi standalone parser")

	is := antlr.NewInputStream(`sashimi:entity(project) of
	- title as "Project Title" is text,
	- description as "Description" is text,
	- image as "Main Image" is picture,
	- starCount as "Star Count" is number,
	- category as "Category" is oneOf("Archtiecture,"Misc.")
	- tags as "Tags" is manyOf("Cheap","Fast","Good")
	- gallery as "Gallery" is list of picture`)
	lexer := parser.NewSashimiLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewSashimiParser(stream)
	antlr.ParseTreeWalkerDefault.Walk(&dummyListener{}, p.Export())
}
