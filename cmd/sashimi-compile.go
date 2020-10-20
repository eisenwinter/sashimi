package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/eisenwinter/sashimi/build"
)

func main() {
	fmt.Println("sashimi standalone compiler")
	input := flag.String("in", "../template", "input files directory")
	output := flag.String("out", "out", "output files directory")
	resolverDriver := flag.String("resolver", "noop", "node resolver driver for data")
	flag.Parse()
	fmt.Printf("Using node resolver: %s\r\n", *resolverDriver)
	sources := getCompilerSources(*input)
	c := build.NewCompiler(nil)
	result, err := c.Compile(sources, *output, build.LogToConsole)
	if err != nil {
		fmt.Printf("FATAL: %v", err)
	} else {
		fmt.Println(result)
	}
}

func getCompilerSources(in string) []build.CompilerSource {
	sources := make([]build.CompilerSource, 0)
	items, err := ioutil.ReadDir(in)
	if err != nil {
		fmt.Printf("Error opening directory: %v", err)
	}
	for _, item := range items {
		if item.IsDir() {
			subitems, _ := ioutil.ReadDir(item.Name())
			for _, subitem := range subitems {
				if !subitem.IsDir() {
					fn := in + "/" + item.Name() + "/" + subitem.Name()
					abs, err := filepath.Abs(fn)
					if err == nil {
						if strings.HasSuffix(abs, ".html") || strings.HasSuffix(abs, ".sushi") {
							fmt.Printf("Adding file: %s\r\n", abs)
							sources = append(sources, build.CreateCompilerSourceFromFile(abs))
						}
					} else {
						if strings.HasSuffix(fn, ".html") || strings.HasSuffix(fn, ".sushi") {
							fmt.Printf("Adding file: %s\r\n", fn)
							sources = append(sources, build.CreateCompilerSourceFromFile(fn))
						}
					}

				}
			}
		} else {
			fn := in + "/" + item.Name()
			abs, err := filepath.Abs(fn)
			if err == nil {
				if strings.HasSuffix(abs, ".html") || strings.HasSuffix(abs, ".sushi") {
					fmt.Printf("Adding file: %s\r\n", abs)
					sources = append(sources, build.CreateCompilerSourceFromFile(abs))
				}
			} else {
				if strings.HasSuffix(fn, ".html") || strings.HasSuffix(fn, ".sushi") {
					fmt.Printf("Adding file: %s\r\n", fn)
					sources = append(sources, build.CreateCompilerSourceFromFile(fn))
				}
			}

		}
	}
	return sources
}
