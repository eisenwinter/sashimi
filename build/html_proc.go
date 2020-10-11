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
