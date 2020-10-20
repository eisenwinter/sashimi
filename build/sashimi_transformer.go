package build

import (
	"fmt"
	"strings"
)

//Transforms sashimi code to html/tempalte code

type transformer struct {
	*BaseSashimiListener
	buffer strings.Builder
}

func (t *transformer) ExitCommandCall(ctx *CommandCallContext) {
	qualifier := ctx.Qualifier().GetText()
	switch ctx.COMMAND().GetText() {
	case "display":
		//resolve type and constraints to validate if next sibling is eligable
		t.buffer.WriteString("{{")
		t.buffer.WriteString(qualifier)
		t.buffer.WriteString("}}")
		break
	case "link":
		//resolve type and validate if next sibling is a image
		break
	case "layout_section":
		//this should have been processed before hand
		break
	case "layout":
		//resolve layout_section for layout and store it in gen
		//if the value is already set throw an error
		t.buffer.WriteString("{{define \"")
		t.buffer.WriteString(qualifier)
		t.buffer.WriteString("\"}}content{{end}}")
		break
	}
}

func (t *transformer) EnterLoopCall(ctx *LoopCallContext) {
	fmt.Println("loop call entered")
	t.buffer.WriteString("{{range .}}")
}

func (t *transformer) ExitLoopCall(ctx *LoopCallContext) {
	t.buffer.WriteString("{{end}}")
}
