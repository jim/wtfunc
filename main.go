package main

import (
	"flag"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	os.Exit(run(os.Stdin, os.Stdout, os.Stderr, os.Args[1:]))
}

func run(in io.Reader, out io.Writer, err io.Writer, args []string) int {
	var line int

	flagSet := flag.NewFlagSet("wtfunc", flag.ExitOnError)
	flagSet.IntVar(&line, "line", 0, "line of file to check")
	flagSet.Parse(args)

	fset := token.NewFileSet()
	var parsedFile *ast.File
	var parseErr error

	if f := flagSet.Arg(0); f != "" {
		parsedFile, parseErr = parser.ParseFile(fset, f, nil, parser.ParseComments)
		if parseErr != nil {
			fmt.Printf("error reading file %s: %s", f, parseErr)
			return 1
		}
	} else {
		content, readErr := ioutil.ReadAll(in)
		if readErr != nil {
			fmt.Println("error reading file:", readErr)
			return 1
		}
		parsedFile, parseErr = parser.ParseFile(fset, "stdin", content, parser.ParseComments)
		if parseErr != nil {
			fmt.Println("error reading from standard in:", parseErr)
			return 1
		}
	}

	for _, decl := range parsedFile.Decls {
		if fn, ok := decl.(*ast.FuncDecl); ok {

			pos := fset.Position(fn.Pos())
			end := fset.Position((fn.End()))

			if pos.Line <= line && end.Line >= line {
				fmt.Fprint(out, fn.Name.Name)
				return 0
			}
		}
	}

	return 2
}
