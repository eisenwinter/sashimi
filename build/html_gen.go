package build

import (
	"fmt"
	"io"
	"strings"

	"golang.org/x/net/html"
)

type htmlLayoutSection struct {
	layout      *html.Node
	contentNode *html.Node
}

type htmlGenerator struct {
	scopeStack        []string
	requireScopeStack bool
	in                *html.Node
	out               *html.Node
}

func (h *htmlGenerator) parseDocument(reader io.Reader) (*html.Node, error) {
	root, e := html.Parse(reader)
	if e != nil {
		return nil, e
	}
	return root, nil
}

func (h *htmlGenerator) traverse(n, p *html.Node) {
	if n.Type == html.CommentNode {
		if strings.Contains(n.Data, "sashimi:") {
			//parse
			fmt.Println("parse block")
			//decide wheter or not we need the next child or not
			//check if its scoped or not
		}
	} else {
		//process this first the append to parent node
		p.AppendChild(h.copyNode(n))
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
