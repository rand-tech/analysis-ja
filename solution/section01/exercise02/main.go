package main

import (
	"errors"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
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

	for _, decl := range f.Decls {
		decl, _ := decl.(*ast.FuncDecl)
		if decl == nil || decl.Recv != nil || decl.Name.Name != "init" || decl.Body == nil {
			continue
		}

		ast.Inspect(decl.Body, func(n ast.Node) bool {
			call, _ := n.(*ast.CallExpr)
			if call == nil {
				return true
			}

			sel, _ := call.Fun.(*ast.SelectorExpr)
			if sel == nil {
				return false
			}

			pkgname, _ := sel.X.(*ast.Ident)
			if pkgname == nil {
				return false
			}

			if pkgname.Name == "exec" && sel.Sel.Name == "Command" {
				pos := fset.Position(n.Pos())
				fmt.Fprintf(os.Stdout, "%s: find exec.Command in init\n", pos)
				return false
			}

			return true
		})
	}

	return nil
}
