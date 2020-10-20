package build

import (
	"strings"
)

//Transforms sashimi code to html/tempalte code

type transformer struct {
	*BaseSashimiListener
	preBuffer  strings.Builder
	postBuffer strings.Builder
	attrMod    map[string]string
}

func (t *transformer) FlushBuffer() (string, string) {
	pre := t.preBuffer.String()
	t.preBuffer.Reset()
	post := t.postBuffer.String()
	t.postBuffer.Reset()
	return pre, post
}

func (t *transformer) FlushAttributeModifiers() map[string]string {
	copy := make(map[string]string)
	for k, v := range t.attrMod {
		copy[k] = v
		delete(t.attrMod, k)
	}
	return copy
}

func (t *transformer) ExitCommandCall(ctx *CommandCallContext) {
	qualifier := ctx.Qualifier().GetText()
	switch ctx.COMMAND().GetText() {
	case "display":
		t.postBuffer.WriteString("{{")
		t.postBuffer.WriteString(qualifier)
		t.postBuffer.WriteString("}}")
		break
	case "link":
		t.postBuffer.WriteString("{{")
		t.postBuffer.WriteString(qualifier)
		t.postBuffer.WriteString(" | linkTextt}}")
		t.attrMod["href"] = "{{" + qualifier + " | link" + "}}"
		break
	case "layout_section":
		//this should have been processed before hand
		break
	case "layout":
		t.preBuffer.WriteString("{{define \"")
		t.preBuffer.WriteString(qualifier)
		t.preBuffer.WriteString("\"}}")
		break
	}
}

func (t *transformer) EnterLoopCall(ctx *LoopCallContext) {
	t.preBuffer.WriteString("{{range .}}")
}
