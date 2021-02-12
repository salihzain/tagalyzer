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

	// test ignore field using `tagalyzer:"-"`
	analysistest.Run(t, analysistest.TestData(), Analyzer, "ignorefield.go")

	// test with embedded fields, first ignore embedded
	analysistest.Run(t, analysistest.TestData(), Analyzer, "ignoreembedded.go")

	// test with embedded, this time check
	Analyzer.Flags.Set("checkembedded", "true")
	analysistest.Run(t, analysistest.TestData(), Analyzer, "checkembedded.go")

	// test with multiple tags
	Analyzer.Flags.Set("tag", "gorm")
	analysistest.Run(t, analysistest.TestData(), Analyzer, "multipletags.go")
}
