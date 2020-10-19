package build

import (
	"fmt"
	"io"
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"golang.org/x/net/html"
)

type htmlLayoutSection struct {
	layout      *html.Node
	contentNode *html.Node
}

type htmlGenerator struct {
	*BaseSashimiListener
	scopeStack        []string
	requireScopeStack bool
	scopeParentNodes  map[string]*html.Node
	in                *html.Node
	sibling           *html.Node
	currentParent     *html.Node
	iteration         int
	callback          func(*html.Node) *html.Node
}

func (h *htmlGenerator) parseDocument(reader io.Reader) ([]*html.Node, error) {
	root, e := html.Parse(reader)
	// roots, e := html.ParseFragment(reader, &html.Node{
	// 	Type:     html.ElementNode,
	// 	Data:     "div",
	// 	DataAtom: atom.Div})
	if e != nil {
		return nil, e
	}
	h.scopeParentNodes = make(map[string]*html.Node)
	return []*html.Node{root}, nil
}

func (h *htmlGenerator) processScope(n *html.Node, scope string) *html.Node {
	ret := h.copyNode(n)
	h.parseScopedBlock(n, ret, scope, 0)
	return ret
}

func (h *htmlGenerator) parseScopedBlock(n, p *html.Node, scope string, i int) {
	if _, ok := h.scopeParentNodes[scope]; !ok {
		h.scopeParentNodes[scope] = n.NextSibling
	}
	if n.Type == html.CommentNode {
		if strings.Contains(n.Data, "sashimi:") {
			//parse
			fmt.Println("parse block")
			if strings.Contains(n.Data, "sashimi:repeat") {
				h.parseScopedBlock(n, p, scope, i+1)
			}
		}
	} else {
		copy := h.copyNode(n)
		copy.Data = "p"
		copy.AppendChild(&html.Node{
			Type: html.TextNode,
			Data: "This is the replacing content",
		})
		p.AppendChild(copy)
	}
	c := n.FirstChild
	for c != nil {
		h.parseScopedBlock(c, p, scope, 0)
		c = c.NextSibling
	}
}

func (h *htmlGenerator) Traverse(nodes []*html.Node) *html.Node {
	rootNode := &html.Node{
		Type: html.DocumentNode,
	}
	for _, n := range nodes {
		h.traverse(n, rootNode)
	}
	return rootNode
}

func (h *htmlGenerator) traverse(n, p *html.Node) {
	h.iteration++
	if n.Type == html.CommentNode {
		if strings.Contains(n.Data, "sashimi:") {
			//parse
			fmt.Println("parse block")
			h.sibling = n.NextSibling
			sashimiCode := n.Data
			if strings.Contains(sashimiCode, "sashimi:repeat") {
				if !h.requireScopeStack {
					h.scopeStack = make([]string, 0)
				}
				h.requireScopeStack = true
				//scope := fmt.Sprintf("%s:%v", h.sibling.Data, h.iteration)
				scope := fmt.Sprintf("%v", h.iteration)
				implicitDirective := fmt.Sprintf("sashimi:begin('%s')", scope)
				sashimiCode += implicitDirective
				//todo: add to loop stack and iterate in loop enter
				h.scopeParentNodes[scope] = n.NextSibling
				//h.parseScopedBlock(h.sibling, scope, 0)
				implicitDirective = fmt.Sprintf("sashimi:end('%s')", scope)
				sashimiCode += implicitDirective
			}
			tree := interpret(sashimiCode)
			fmt.Println(tree)

			antlr.ParseTreeWalkerDefault.Walk(h, tree)
			p.AppendChild(h.copyNode(n))
		}
	} else {
		if h.callback != nil {
			fmt.Println("exec callback")
			p.AppendChild(h.callback(n))
		} else {
			p.AppendChild(h.copyNode(n))
		}
		//p.AppendChild(h.copyNode(n))
	}
	c := n.FirstChild
	for c != nil {
		h.traverse(c, p.LastChild)
		c = c.NextSibling
	}
}

func (*htmlGenerator) copyNode(n *html.Node) *html.Node {
	return &html.Node{
		Type:      n.Type,
		DataAtom:  n.DataAtom,
		Data:      n.Data,
		Namespace: n.Namespace,
		Attr:      n.Attr}
}

func (*htmlGenerator) processLinkTag(n *html.Node) {
	if n.Type == html.ElementNode && n.Data == "a" {

	}
}

func (*htmlGenerator) processImageTag(n *html.Node) {
	if n.Type == html.ElementNode && n.Data == "img" {

	}
}

func (*htmlGenerator) processAbitaryTag(n *html.Node) {
	if n.Type == html.ElementNode {

	}
}

func (h *htmlGenerator) EnterLoopCall(ctx *LoopCallContext) {
	fmt.Println("loop call entered")
}

func (h *htmlGenerator) ExitLoopCall(ctx *LoopCallContext) {

}

func (h *htmlGenerator) EnterCommandCall(ctx *CommandCallContext) {}

func (h *htmlGenerator) ExitCommandCall(ctx *CommandCallContext) {
	//Todo get current scope to resolve wich type
	fmt.Println("Command:", ctx.COMMAND())
	switch ctx.COMMAND().GetText() {
	case "display":
		//resolve type and constraints to validate if next sibling is eligable
		h.processAbitaryTag(h.sibling)
		break
	case "link":
		//resolve type and validate if next sibling is a image
		h.processLinkTag(h.sibling)
		break
	case "layout_section":
		//this should have been processed before hand
		break
	case "layout":
		//resolve layout_section for layout and store it in gen
		//if the value is already set throw an error
		break
	}
	fmt.Println("Received command call for sibling", h.sibling)
}

func (h *htmlGenerator) ExitScopeBegin(ctx *ScopeBeginContext) {
	if ctx.SCOPEIDENT() != nil {
		scopeIdent := strings.Trim(ctx.SCOPEIDENT().GetText(), "'")
		h.scopeStack = append(h.scopeStack, scopeIdent)
		h.callback = func(n *html.Node) *html.Node {
			if n.Type != html.ElementNode {
				return h.copyNode(n)
			}
			r := h.processScope(n, scopeIdent)
			n.FirstChild = nil
			n.LastChild = nil
			h.callback = nil
			return r
		}
	}
}

func (h *htmlGenerator) ExitScopeEnd(ctx *ScopeEndContext) {
	if ctx.SCOPEIDENT() != nil {
		n := len(h.scopeStack) - 1
		h.scopeStack = h.scopeStack[:n]
	}
}

func (h *htmlGenerator) EnterBlockScope(ctx *BlockScopeContext) {

}

func (h *htmlGenerator) ExitBlockScope(ctx *BlockScopeContext) {

}
