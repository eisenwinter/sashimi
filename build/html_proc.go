package build

import (
	"bytes"
	"fmt"
	"io"

	"golang.org/x/net/html"
)

type htmlProcessor struct {
	scopeStack        []string
	requireScopeStack bool
}

//range {{range .Todos}} {{.}} {{end}}
//display {{.Title}} --> add pipe for resolver (link etc)
//sub template {{define "sub-template"}}content{{end}}
func (h *htmlProcessor) transform(reader io.Reader, writer io.Writer, skipSushi bool) error {
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
			break
		case html.TextToken:
			content := tokenizer.Text()
			writer.Write(content)
			break
		case html.SelfClosingTagToken:
			content := tokenizer.Text()
			writer.Write(content)
			break
		case html.StartTagToken:
			content := tokenizer.Text()

			writer.Write(content)
			break
		case html.EndTagToken:
			content := tokenizer.Text()
			writer.Write(content)
			break
		case html.DoctypeToken:
			content := tokenizer.Text()
			writer.Write(content)
			if !skipSushi {
				writer.Write([]byte("<!--🍣-->"))
			}
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
					implicitDirective := fmt.Sprintf("sashimi:begin('%s')", innerScope)
					_, err := writer.Write([]byte(implicitDirective))
					if err != nil {
						return err
					}
				} else {
					h.scopeStack = append(h.scopeStack, scope)
					implicitDirective := fmt.Sprintf("sashimi:begin('%s')", scope)
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
				implicitDirective := fmt.Sprintf("sashimi:end('%s')", h.scopeStack[n])
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
