package checkcustomerror_test

import (
	"flag"
	"testing"

	"github.com/gostaticanalysis/testutil"
	"github.com/wheatandcat/memoir-static-analytics/checkcustomerror"
	"golang.org/x/tools/go/analysis/analysistest"
)

// TestAnalyzer is a test for Analyzer.
func TestAnalyzer(t *testing.T) {
	flag.CommandLine.Set("exclude_regex", "_test.go")

	testdata := testutil.WithModules(t, analysistest.TestData(), nil)
	analysistest.Run(t, testdata, checkcustomerror.Analyzer, "a")
}
