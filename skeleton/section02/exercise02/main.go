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

	// TODO: 型チェックをする

	var strconvAtoi types.Object
	for _, p := range pkg.Imports() {
		if p.Path() == "strconv" {
			strconvAtoi = p.Scope().Lookup("Atoi")
			break
		}
	}

	if strconvAtoi == nil {
		// skip
		return nil
	}

	numObj := findCallAtoi(fset, info, strconvAtoi, f)
	findCast(fset, info, numObj, f)

	return nil
}

func findCallAtoi(fset *token.FileSet, info *types.Info, strconvAtoi types.Object, root ast.Node) types.Object {
	var numObj types.Object
	ast.Inspect(root, func(n ast.Node) bool {
		if numObj != nil {
			// already found
			return false
		}

		// TODO: 左辺が2つ、右辺が1つの代入文に絞る
		

		call, _ := assign.Rhs[0].(*ast.CallExpr)
		if call == nil {
			return true
		}

		sel, _ := call.Fun.(*ast.SelectorExpr)
		if sel == nil {
			return true
		}

		fun := info.ObjectOf(sel.Sel)
		if fun == strconvAtoi {
			numVar, _ := assign.Lhs[0].(*ast.Ident)
			if numVar == nil {
				return false
			}

			numObj = info.ObjectOf(numVar)
			return false
		}

		return true
	})

	return numObj
}

func findCast(fset *token.FileSet, info *types.Info, numObj types.Object, root ast.Node) {
	ast.Inspect(root, func(n ast.Node) bool {

		call, _ := n.(*ast.CallExpr)
		if call == nil || len(call.Args) != 1 {
			return true
		}

		numID, _ := call.Args[0].(*ast.Ident)
		if numID == nil || info.ObjectOf(numID) != numObj {
			return true
		}

		typeID, _ := call.Fun.(*ast.Ident)
		if typeID == nil {
			return false
		}

		// TODO: typeIDから取得できるオブジェクトが型名であることを確認する

		if isType && (typeID.Name == "int16" || typeID.Name == "int32") {
			pos := fset.Position(n.Pos())
			fmt.Fprintf(os.Stdout, "%s: find integer overflow\n", pos)
			return false
		}

		return true
	})
}
