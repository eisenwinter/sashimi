package build

import (
	"bytes"
	"io"

	"golang.org/x/net/html"
)

type htmlProcessor struct{}

func (*htmlProcessor) extractFromHTML(reader io.Reader, writer io.Writer) error {
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
		if token == html.CommentToken {
			content := tokenizer.Text()
			if bytes.Contains(content, []byte("sashimi:")) {
				_, err := writer.Write(content)
				//len(tokenizer.Raw()) might care for offset later on
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}
