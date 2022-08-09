package main

import (
	"github.com/gohandson/analysis-ja/solution/section03/exercise02"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(exercise02.Analyzer) }
