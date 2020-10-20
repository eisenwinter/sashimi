package build

import (
	"bytes"
	"fmt"
	"io"
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

type htmlProcessor struct {
	scopeStack        []string
	requireScopeStack bool

	requireDepthStack bool
	depth             int
	depthStack        []int
	sashimiBuffer     strings.Builder
	transformer       *transformer
	skipNextText      bool
}

func (h *htmlProcessor) transform(reader io.Reader, writer io.Writer, skipSushi bool) error {
	h.sashimiBuffer.Reset()
	h.transformer = &transformer{
		attrMod: make(map[string]string),
	}
	tokenizer := html.NewTokenizer(reader)
	for {
		token := tokenizer.Next()
		switch token {
		case html.ErrorToken:
			err := tokenizer.Err()
			if err == io.EOF {
				return nil
			}
			return err
		case html.CommentToken:
			content := tokenizer.Text()
			if bytes.Contains(content, []byte("sashimi:")) {
				h.sashimiBuffer.Write(content)
				//this is rather cheap but for now enough, it could be determined in the transformer as well
				if bytes.Contains(content, []byte("sashimi:repeat")) || bytes.Contains(content, []byte("sashimi:layout(")) {
					if !h.requireDepthStack {
						h.depthStack = make([]int, 0)
					}
					h.requireDepthStack = true
					h.depthStack = append(h.depthStack, h.depth)
				}
			}
			break
		case html.TextToken:
			if h.skipNextText {
				h.skipNextText = false
			} else {
				content := tokenizer.Raw()
				writer.Write(content)
			}
			break
		case html.SelfClosingTagToken:
			content := tokenizer.Raw()
			writer.Write(content)
			break
		case html.StartTagToken:
			h.depth++
			if h.sashimiBuffer.Len() > 0 {
				content := tokenizer.Token()
				is := antlr.NewInputStream(h.sashimiBuffer.String())
				lexer := NewSashimiLexer(is)
				h.sashimiBuffer.Reset()
				stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
				p := NewSashimiParser(stream)
				antlr.ParseTreeWalkerDefault.Walk(h.transformer, p.Block())
				attrMods := h.transformer.FlushAttributeModifiers()
				for k, v := range attrMods {
					found := false
					for i, attr := range content.Attr {
						if attr.Key == k {
							content.Attr[i] = html.Attribute{
								Key: k,
								Val: v,
							}
							found = true
							break
						}
					}
					if !found {
						content.Attr = append(content.Attr, html.Attribute{
							Key: k,
							Val: v,
						})
					}
				}
				pre, post := h.transformer.FlushBuffer()
				if len(pre) > 0 {
					writer.Write([]byte(pre))
				}

				writer.Write([]byte(content.String()))
				if len(post) > 0 {
					writer.Write([]byte(post))
					h.skipNextText = true
				}
			} else {
				content := tokenizer.Raw()
				writer.Write(content)

			}

			break
		case html.EndTagToken:
			h.depth--
			content := tokenizer.Token()
			if !skipSushi && content.DataAtom == atom.Html {
				writer.Write([]byte("<!--🍣-->"))
			}
			writer.Write([]byte(content.String()))
			if h.requireDepthStack {
				n := len(h.depthStack) - 1
				if n >= 0 {
					for n >= 0 && h.depthStack[n] == h.depth {
						h.depthStack = h.depthStack[:n]
						writer.Write([]byte("{{end}}"))
						n = len(h.depthStack) - 1
					}

				} else {
					h.requireDepthStack = false
				}
			}
			break
		case html.DoctypeToken:
			content := tokenizer.Raw()
			writer.Write(content)
			break
		}

	}
}

//extractFromHTML extracts all sashimi directives from the HTML (basically creates a .sushi files from .html)
func (h *htmlProcessor) extractFromHTML(reader io.Reader, writer io.Writer) error {
	tokenizer := html.NewTokenizer(reader)
	for {
		token := tokenizer.Next()
		if token == html.ErrorToken {
			err := tokenizer.Err()
			if err == io.EOF {
				return nil
			}
			return err
		}
		if token == html.StartTagToken {
			if h.requireScopeStack {
				offset := len(tokenizer.Raw())
				name, _ := tokenizer.TagName()
				scope := fmt.Sprintf("%s:%v", name, offset)
				if len(h.scopeStack) > 0 {
					n := len(h.scopeStack) - 1
					innerScope := fmt.Sprintf("%s::%s", h.scopeStack[n], scope)
					h.scopeStack = append(h.scopeStack, innerScope)
					implicitDirective := fmt.Sprintf(" sashimi:begin('%s') ", innerScope)
					_, err := writer.Write([]byte(implicitDirective))
					if err != nil {
						return err
					}
				} else {
					h.scopeStack = append(h.scopeStack, scope)
					implicitDirective := fmt.Sprintf(" sashimi:begin('%s') ", scope)
					_, err := writer.Write([]byte(implicitDirective))
					if err != nil {
						return err
					}
				}
			}
		}
		if token == html.EndTagToken {
			if h.requireScopeStack {
				n := len(h.scopeStack) - 1
				implicitDirective := fmt.Sprintf(" sashimi:end('%s') ", h.scopeStack[n])
				_, err := writer.Write([]byte(implicitDirective))
				if err != nil {
					return err
				}
				h.scopeStack = h.scopeStack[:n]
				if len(h.scopeStack) == 0 {
					h.requireScopeStack = false
				}
			}
		}
		if token == html.CommentToken {
			content := tokenizer.Text()
			if bytes.Contains(content, []byte("sashimi:")) {
				if bytes.Contains(content, []byte("sashimi:repeat")) {
					if !h.requireScopeStack {
						h.scopeStack = make([]string, 0)
					}
					h.requireScopeStack = true
				}
				_, err := writer.Write(content)
				if err != nil {
					return err
				}
			}
		}
	}
}
