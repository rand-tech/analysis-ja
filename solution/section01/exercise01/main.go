package main

import (
	"errors"
	"fmt"
	"go/parser"
	"go/token"
	"os"
	"strconv"
)

func main() {
	if err := run(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(args []string) error {
	if len(args) < 1 {
		return errors.New("source code must be specified")
	}

	fname := args[0]
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, fname, nil, 0)
	if err != nil {
		return err
	}

	for _, spec := range f.Imports {
		path, err := strconv.Unquote(spec.Path.Value)
		if err != nil {
			return err
		}

		if path == "unsafe" {
			pos := fset.Position(spec.Pos())
			fmt.Fprintf(os.Stderr, "%s: import unsafe\n", pos)
		}
	}

	return nil
}
