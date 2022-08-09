package main

import (
	"errors"
	"fmt"
	"go/ast"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
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

	config := &types.Config{
		Importer: importer.Default(),
	}

	info := &types.Info{
		Types:      make(map[ast.Expr]types.TypeAndValue),
		Instances:  make(map[*ast.Ident]types.Instance),
		Defs:       make(map[*ast.Ident]types.Object),
		Uses:       make(map[*ast.Ident]types.Object),
		Implicits:  make(map[ast.Node]types.Object),
		Selections: make(map[*ast.SelectorExpr]*types.Selection),
		Scopes:     make(map[ast.Node]*types.Scope),
		InitOrder:  []*types.Initializer{},
	}

	pkg, err := config.Check("lib", fset, []*ast.File{f}, info)
	if err != nil {
		return err
	}

	var execCommand types.Object
	// TODO: インポートしているパッケージからexecパッケージを探して、
	// パッケージスコープからCommand関数のオブジェクトを取得する
	

	if execCommand == nil {
		// skip
		return nil
	}

	for _, decl := range f.Decls {
		decl, _ := decl.(*ast.FuncDecl)
		if decl == nil || decl.Recv != nil || decl.Name.Name != "init" || decl.Body == nil {
			continue
		}

		findCallCommand(fset, info, execCommand, decl.Body)
	}

	return nil
}

func findCallCommand(fset *token.FileSet, info *types.Info, execCommand types.Object, root ast.Node) {
	ast.Inspect(root, func(n ast.Node) bool {
		call, _ := n.(*ast.CallExpr)
		if call == nil {
			return true
		}

		sel, _ := call.Fun.(*ast.SelectorExpr)
		if sel == nil {
			return false
		}

		// TODO: 呼び出している関数のオブジェクトを取得する
		// ヒント：*ast.Ident型の値はast.SelectorExprのSelフィールドから取れる

		if fun == execCommand {
			pos := fset.Position(n.Pos())
			fmt.Fprintf(os.Stdout, "%s: find exec.Command in init\n", pos)
			return false
		}

		return true
	})
}
