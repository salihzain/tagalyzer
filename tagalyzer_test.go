package tagalyzer

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
)

func TestTagalyzer(t *testing.T) {
	// first without any tags
	analysistest.Run(t, analysistest.TestData(), Analyzer, "notag.go")

	// test with a single tag
	Analyzer.Flags.Set("tag", "json")
	analysistest.Run(t, analysistest.TestData(), Analyzer, "onetag.go")

	// test with multiple tags
	Analyzer.Flags.Set("tag", "gorm")
	analysistest.Run(t, analysistest.TestData(), Analyzer, "multipletags.go")
}
