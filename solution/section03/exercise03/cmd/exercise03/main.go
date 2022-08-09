package main

import (
	"github.com/gohandson/analysis-ja/solution/section03/exercise03"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(exercise03.Analyzer) }
