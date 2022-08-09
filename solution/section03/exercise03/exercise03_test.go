package exercise03_test

import (
	"testing"

	"github.com/gohandson/analysis-ja/solution/section03/exercise03"
	"github.com/gostaticanalysis/testutil"
	"golang.org/x/tools/go/analysis/analysistest"
)

// TestAnalyzer is a test for Analyzer.
func TestAnalyzer(t *testing.T) {
	testdata := testutil.WithModules(t, analysistest.TestData(), nil)
	analysistest.Run(t, testdata, exercise03.Analyzer, "a")
}
