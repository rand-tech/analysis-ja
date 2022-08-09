package exercise01

import (
	"go/types"

	"github.com/gostaticanalysis/analysisutil"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/buildssa"
	"golang.org/x/tools/go/ssa"
)

const doc = "exercise01 finds integer overflow"

// Analyzer finds integer overflow
var Analyzer = &analysis.Analyzer{
	Name: "exercise01",
	Doc:  doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		buildssa.Analyzer,
	},
}

func run(pass *analysis.Pass) (any, error) {
	strconvAtoi := analysisutil.LookupFromImports(pass.Pkg.Imports(), "strconv", "Atoi")
	if strconvAtoi == nil {
		// skip
		return nil, nil
	}

	s := pass.ResultOf[buildssa.Analyzer].(*buildssa.SSA)
	for _, f := range s.SrcFuncs {
		for _, b := range f.Blocks {
			for _, instr := range b.Instrs {
				instr, _ := instr.(*ssa.Convert)
				if instr == nil {
					continue
				}

				typ := instr.Type()
				if !types.Identical(typ, types.Typ[types.Int16]) &&
					!types.Identical(typ, types.Typ[types.Int32]) {
					continue
				}

				extract, _ := instr.X.(*ssa.Extract)
				if extract == nil || extract.Index != 0 {
					continue
				}

				call, _ := extract.Tuple.(*ssa.Call)
				if call == nil {
					continue
				}

				fun, _ := call.Call.Value.(*ssa.Function)
				if fun == nil {
					continue
				}

				if fun.Object() == strconvAtoi {
					pass.Reportf(instr.Pos(), "integer overflow")
				}
			}
		}
	}
	return nil, nil
}
