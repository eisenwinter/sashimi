package build

import (
	"io"
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/eisenwinter/sashimi/graph"
)

//CompilerFlags represent possible flags the compiler takes
type CompilerFlags uint8

const (
	//HyperCritical flag treats all warnings as errors
	HyperCritical CompilerFlags = 1 << iota
	//NoSushi omits the sushi emoji on code generation
	NoSushi
	//OptimizeOff do not minify and inline output
	OptimizeOff
)

type Report interface {
	OnLine() int
	OnPos() int
	InSource() string
	InternalCode() int
}

type CompilerResult interface {
	GetErrors() []Report
	GetWarnings() []Report
}

type AnalyizeResult interface {
	CompilerResult
	GetGraph() *graph.SchemaGraph
}

type Compiler interface {
	Analyze(sources []CompilerSource, flags CompilerFlags) (AnalyizeResult, error)
}

func NewCompiler() Compiler {
	return &sashimiCompiler{htmlProc: &htmlProcessor{}}
}

type sashimiCompiler struct {
	htmlProc *htmlProcessor
}

func (c *sashimiCompiler) Analyze(sources []CompilerSource, flags CompilerFlags) (AnalyizeResult, error) {
	fp := plainFirstPassParser()
	for _, v := range sources {
		var sb strings.Builder
		rc, err := v.Load()
		if err != nil {
			return nil, err
		}
		defer rc.Close()
		if v.Format() == "html" {
			err := c.htmlProc.extractFromHTML(rc, &sb)
			if err != nil {
				return nil, err
			}
		} else {
			_, err = io.Copy(&sb, rc)
			if err != nil {
				return nil, err
			}
		}
		fp.source = v.Name()
		is := antlr.NewInputStream(sb.String())
		lexer := NewSashimiLexer(is)
		stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
		p := NewSashimiParser(stream)
		antlr.ParseTreeWalkerDefault.Walk(fp, p.Block())
	}
	fp.ctx.Consolidate()
	return fp.ctx, nil
}
