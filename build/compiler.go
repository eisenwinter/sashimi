package build

import (
	"fmt"
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
	//LogToConsole enables direct log to console
	LogToConsole
)

func (c CompilerFlags) isSet(flag CompilerFlags) bool {
	return c&flag != 0
}

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
	Analyze(sources []CompilerSource) (AnalyizeResult, error)
	Compile(sources []CompilerSource, out string, flags CompilerFlags) (CompileResult, error)
}

//NewCompiler returns a new instance of a sashimi compiler
func NewCompiler(nodeResolver resolver.NodeResolver) Compiler {
	return &sashimiCompiler{htmlProc: &htmlProcessor{}, nodeResolver: nodeResolver}
}

type sashimiCompiler struct {
	htmlProc     *htmlProcessor
	nodeResolver resolver.NodeResolver
	ctx          *parserContext
}

func (c *sashimiCompiler) Analyze(sources []CompilerSource) (AnalyizeResult, error) {
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
	c.ctx = fp.ctx
	return fp.ctx, nil
}

func (c *sashimiCompiler) Compile(sources []CompilerSource, out string, flags CompilerFlags) (CompileResult, error) {
	res, err := c.Analyze(sources)
	if err != nil {
		return res, err
	}
	if len(res.GetErrors()) > 0 {
		return res, err
	}
	if flags.isSet(HyperCritical) && len(res.GetWarnings()) > 0 {
		return res, err
	}
	for qualifier, layout := range c.ctx.Calls["layout_section"] {
		if flags.isSet(LogToConsole) {
			fmt.Printf("Processing layout: %s", qualifier)
		}
		for source := range layout {
			fmt.Printf("Layout source: %s", source)
		}
		//ToDo: basically we transform it to html template layout and store it
	}
	for n, s := range c.ctx.ProcessedSources {
		if flags.isSet(LogToConsole) {
			fmt.Printf("Processing: %s", n)
			fmt.Printf("Requires multiple runs: %v", s.isMany)
		}
		for _, ed := range s.requiredEntities {
			if flags.isSet(LogToConsole) {
				fmt.Printf("Processing required entity: %s", ed.name)
				fmt.Printf("Requires multiple runs: %v", ed.many)
				fmt.Printf("Predicate: %s", ed.predicate)
				fmt.Printf("Scoped to: %s", ed.scope)
			}
			if ed.predicate != "" {
				//ToDo build predicate and call Resolve
			} else {
				c.nodeResolver.ResolveAll(ed.name)
			}
		}
		neededTypes := strings.Join(s.manyBy, "_")
		if flags.isSet(LogToConsole) {
			fmt.Printf("Constructed postfix: %s", neededTypes)
		}
	}
	return nil, nil
}
