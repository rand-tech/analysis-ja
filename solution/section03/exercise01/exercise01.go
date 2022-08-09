package exercise01

import (
	"strconv"

	"golang.org/x/tools/go/analysis"
)

const doc = "exercise01 finds import spec which imports unsafe package"

// Analyzer finds import spec which imports unsafe package
var Analyzer = &analysis.Analyzer{
	Name: "exercise01",
	Doc:  doc,
	Run:  run,
}

func run(pass *analysis.Pass) (any, error) {
	for _, f := range pass.Files {
		for _, spec := range f.Imports {
			path, err := strconv.Unquote(spec.Path.Value)
			if err != nil {
				return nil, err
			}

			if path == "unsafe" {
				pass.Reportf(spec.Pos(), "import unsafe")
			}
		}
	}
	return nil, nil
}
