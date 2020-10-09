package build

import (
	"io"
	"os"
	"strings"
)

type CompilerSource interface {
	Load() (rc io.ReadCloser, err error)
	Name() string
}

type fileCompilerSource struct {
	path     string
	fileName string
}

func (f *fileCompilerSource) Load() (rc io.ReadCloser, err error) {
	return os.Open(f.path)
}

func (f *fileCompilerSource) Name() string {
	return f.fileName
}

//CreateCompilerSourceFromFile creates a new compiler source from the supplied path
func CreateCompilerSourceFromFile(path string) CompilerSource {
	return &fileCompilerSource{}
}

type stringReaderNoopCloser struct {
	reader *strings.Reader
}

func (c *stringReaderNoopCloser) Read(p []byte) (n int, err error) {
	return c.reader.Read(p)
}

func (c *stringReaderNoopCloser) Close() error {
	return nil
}

type stringCompilerSource struct {
	content string
}

func (s *stringCompilerSource) Load() (rc io.ReadCloser, err error) {
	return &stringReaderNoopCloser{strings.NewReader(s.content)}, nil
}

func (*stringCompilerSource) Name() string {
	return "direct input"
}

//CreateCompilerSourceFromString creates a compiler source based on the supplied string
func CreateCompilerSourceFromString(content string) CompilerSource {
	return &stringCompilerSource{content}
}
