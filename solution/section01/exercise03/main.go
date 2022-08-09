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
		decl, _ := decl.(*ast.GenDecl)
		if decl == nil || decl.Tok != token.VAR {
			continue
		}

		for _, spec := range decl.Specs {
			spec, _ := spec.(*ast.ValueSpec)
			if spec == nil {
				continue
			}

			for _, value := range spec.Values {
				call, _ := value.(*ast.CallExpr)
				if call == nil {
					continue
				}

				funlit, _ := call.Fun.(*ast.FuncLit)
				if funlit == nil {
					continue
				}

				findCallCommand(fset, funlit.Body)
			}
		}
	}

	return nil
}

func findCallCommand(fset *token.FileSet, root ast.Node) {
	ast.Inspect(root, func(n ast.Node) bool {
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
			fmt.Fprintf(os.Stdout, "%s: find exec.Command in package variable initializer\n", pos)
			return false
		}

		return true
	})
}
