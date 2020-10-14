package build

import (
	"io"
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/eisenwinter/sashimi/graph"
	"github.com/eisenwinter/sashimi/resolver"
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

//Report contains hints to certain parts of the source
type Report interface {
	OnLine() int
	OnPos() int
	InSource() string
	InternalCode() int
}

//CompilerResult contains any warnings or errors generated during a pass
type CompilerResult interface {
	GetErrors() []Report
	GetWarnings() []Report
}

//AnalyizeResult contains the result of a analyze run, it may contain errors or warnings as well as the generated schema graph for the entities
type AnalyizeResult interface {
	CompilerResult
	GetGraph() *graph.SchemaGraph
}

//CompileResult contains the result of the compile run
type CompileResult interface {
	CompilerResult
}

//Compiler is a sashimi compiler
type Compiler interface {
	Analyze(sources []CompilerSource, flags CompilerFlags) (AnalyizeResult, error)
	Compile(sources []CompilerSource, out string, flags CompilerFlags) (CompileResult, error)
}

//NewCompiler returns a new instance of a sashimi compiler
func NewCompiler(nodeResolver resolver.NodeResolver) Compiler {
	return &sashimiCompiler{htmlProc: &htmlProcessor{}, nodeResolver: nodeResolver}
}

type sashimiCompiler struct {
	htmlProc     *htmlProcessor
	nodeResolver resolver.NodeResolver
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

func (c *sashimiCompiler) Compile(sources []CompilerSource, out string, flags CompilerFlags) (CompileResult, error) {
	return nil, nil
}
