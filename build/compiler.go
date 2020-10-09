package build

import (
	"io"
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

type CompilerSource interface {
	Load() (sourceName string, reader io.Reader, err error)
}

type Report interface {
	OnLine() int
	OnPos() int
	InSource() string
	InternalCode() int
}

type CompilerResult interface {
	Errors() []Report
	Warnings() []Report
}

type Compiler interface {
	Analyze(sources []CompilerSource, flags CompilerFlags) (CompilerResult, error)
}

type sashimiCompiler struct{}
