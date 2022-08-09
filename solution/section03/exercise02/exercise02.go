package exercise02

import (
	"go/ast"
	"go/types"

	"github.com/gostaticanalysis/analysisutil"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

const doc = "exercise02 finds integer overflow"

// Analyzer finds integer overflow
var Analyzer = &analysis.Analyzer{
	Name: "exercise02",
	Doc:  doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
}

func run(pass *analysis.Pass) (any, error) {
	strconvAtoi := analysisutil.LookupFromImports(pass.Pkg.Imports(), "strconv", "Atoi")
	if strconvAtoi == nil {
		// skip
		return nil, nil
	}

	numObj := findCallAtoi(pass, strconvAtoi)
	findCast(pass, numObj)

	return nil, nil
}

func findCallAtoi(pass *analysis.Pass, strconvAtoi types.Object) types.Object {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.AssignStmt)(nil),
	}

	var numObj types.Object
	inspect.Preorder(nodeFilter, func(n ast.Node) {
		if numObj != nil {
			// already found
			return
		}

		assign, _ := n.(*ast.AssignStmt)
		if assign == nil || len(assign.Lhs) != 2 || len(assign.Rhs) != 1 {
			return
		}

		call, _ := assign.Rhs[0].(*ast.CallExpr)
		if call == nil {
			return
		}

		sel, _ := call.Fun.(*ast.SelectorExpr)
		if sel == nil {
			return
		}

		fun := pass.TypesInfo.ObjectOf(sel.Sel)
		if fun == strconvAtoi {
			numVar, _ := assign.Lhs[0].(*ast.Ident)
			if numVar == nil {
				return
			}
			numObj = pass.TypesInfo.ObjectOf(numVar)
		}
	})

	return numObj
}

func findCast(pass *analysis.Pass, numObj types.Object) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.CallExpr)(nil),
	}

	inspect.Preorder(nodeFilter, func(n ast.Node) {
		call, _ := n.(*ast.CallExpr)
		if call == nil || len(call.Args) != 1 {
			return
		}

		numID, _ := call.Args[0].(*ast.Ident)
		if numID == nil || pass.TypesInfo.ObjectOf(numID) != numObj {
			return
		}

		typeID, _ := call.Fun.(*ast.Ident)
		if typeID == nil {
			return
		}

		_, isType := pass.TypesInfo.ObjectOf(typeID).(*types.TypeName)
		if isType && (typeID.Name == "int16" || typeID.Name == "int32") {
			pass.Reportf(n.Pos(), "integer overflow")
		}
	})
}
