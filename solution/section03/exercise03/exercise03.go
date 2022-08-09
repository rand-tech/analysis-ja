package exercise03

import (
	"go/ast"
	"go/types"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

const doc = "exercise03 finds ignored errors"

// Analyzer finds ignored errors
var Analyzer = &analysis.Analyzer{
	Name: "exercise03",
	Doc:  doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
}

func run(pass *analysis.Pass) (any, error) {
	errType := types.Universe.Lookup("error").Type()
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.AssignStmt)(nil),
	}

	inspect.Preorder(nodeFilter, func(n ast.Node) {
		assign, _ := n.(*ast.AssignStmt)
		if assign == nil || len(assign.Rhs) != 1 {
			return
		}

		lastVar, _ := assign.Lhs[len(assign.Lhs)-1].(*ast.Ident)
		if lastVar == nil || lastVar.Name != "_" {
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

		fun, _ := pass.TypesInfo.ObjectOf(sel.Sel).(*types.Func)
		if fun == nil {
			return
		}

		sig, _ := fun.Type().(*types.Signature)
		if sig == nil {
			return
		}

		rets := sig.Results()
		if rets.Len() == 0 {
			return
		}

		lastRetType := rets.At(rets.Len() - 1).Type()
		if types.Identical(lastRetType, errType) {
			pass.Reportf(n.Pos(), "ignore error")
		}
	})

	return nil, nil
}
