package exercise01_test

import (
	"testing"

	"github.com/gohandson/analysis-ja/solution/section03/exercise01"
	"github.com/gostaticanalysis/testutil"
	"golang.org/x/tools/go/analysis/analysistest"
)

// TestAnalyzer is a test for Analyzer.
func TestAnalyzer(t *testing.T) {
	testdata := testutil.WithModules(t, analysistest.TestData(), nil)
	analysistest.Run(t, testdata, exercise01.Analyzer, "a")
}
