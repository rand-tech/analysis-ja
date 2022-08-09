package exercise02_test

import (
	"testing"

	"github.com/gohandson/analysis-ja/solution/section03/exercise02"
	"github.com/gostaticanalysis/testutil"
	"golang.org/x/tools/go/analysis/analysistest"
)

// TestAnalyzer is a test for Analyzer.
func TestAnalyzer(t *testing.T) {
	testdata := testutil.WithModules(t, analysistest.TestData(), nil)
	analysistest.Run(t, testdata, exercise02.Analyzer, "a")
}
