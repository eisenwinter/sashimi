package build

import (
	"io"
	"os"
	"path/filepath"
	"strings"
)

//CompilerSource represents a data input source for the compiler
type CompilerSource interface {
	Load() (rc io.ReadCloser, err error)
	Name() string
	Format() string
}

type fileCompilerSource struct {
	path     string
	fileName string
	ext      string
}

func (f *fileCompilerSource) Load() (rc io.ReadCloser, err error) {
	return os.Open(f.path)
}

func (f *fileCompilerSource) Name() string {
	return f.fileName
}

func (f *fileCompilerSource) Format() string {
	return f.ext
}

//CreateCompilerSourceFromFile creates a new compiler source from the supplied path
func CreateCompilerSourceFromFile(path string) CompilerSource {
	_, file := filepath.Split(path)
	ext := strings.TrimPrefix(strings.ToLower(filepath.Ext(file)), ".")
	return &fileCompilerSource{fileName: file, ext: ext, path: path}
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
func (*stringCompilerSource) Format() string {
	return "plain"
}

//CreateCompilerSourceFromString creates a compiler source based on the supplied string
func CreateCompilerSourceFromString(content string) CompilerSource {
	return &stringCompilerSource{content}
}
