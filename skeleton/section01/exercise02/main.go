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

		for _, stmt := range /* TODO: 関数のボディを取得しfor文で回る */ {

			// TODO: 代入文か確かめる

			if assign == nil || len(assign.Rhs) != 1 {
				continue
			}

			// TODO: 代入文の右辺の1つめが関数呼び出しか調べる

			if call == nil {
				continue
			}

			sel, _ := call.Fun.(*ast.SelectorExpr)
			if sel == nil {
				continue
			}

			pkgname, _ := sel.X.(*ast.Ident)
			if pkgname == nil {
				continue
			}

			if /* TODO: パッケージ名がexecで関数名がCommandの関数を呼んでるか調べる*/ {
				pos := fset.Position(assign.Rhs[0].Pos())
				fmt.Fprintf(os.Stdout, "%s: find exec.Command in init\n", pos)
			}
		}
	}

	return nil
}
