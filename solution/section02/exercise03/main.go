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

	_, err = config.Check("main", fset, []*ast.File{f}, info)
	if err != nil {
		return err
	}

	errType := types.Universe.Lookup("error").Type()
	findIgnoreError(fset, info, errType, f)

	return nil
}

func findIgnoreError(fset *token.FileSet, info *types.Info, errType types.Type, root ast.Node) {
	ast.Inspect(root, func(n ast.Node) bool {

		assign, _ := n.(*ast.AssignStmt)
		if assign == nil || len(assign.Rhs) != 1 {
			return true
		}

		lastVar, _ := assign.Lhs[len(assign.Lhs)-1].(*ast.Ident)
		if lastVar == nil || lastVar.Name != "_" {
			return true
		}

		call, _ := assign.Rhs[0].(*ast.CallExpr)
		if call == nil {
			return true
		}

		sel, _ := call.Fun.(*ast.SelectorExpr)
		if sel == nil {
			return true
		}

		fun, _ := info.ObjectOf(sel.Sel).(*types.Func)
		if fun == nil {
			return false
		}

		sig, _ := fun.Type().(*types.Signature)
		if sig == nil {
			return false
		}

		rets := sig.Results()
		if rets.Len() == 0 {
			return false
		}

		lastRetType := rets.At(rets.Len() - 1).Type()
		if types.Identical(lastRetType, errType) {
			pos := fset.Position(n.Pos())
			fmt.Fprintf(os.Stdout, "%s: find ignored error\n", pos)
			return false
		}

		return false
	})
}
