package reassign

import (
	"golang.org/x/tools/go/analysis"

	"github.com/curioswitch/go-reassign/internal/analyzer"
)

const FlagPattern = analyzer.FlagPattern

// NewAnalyzer returns an analyzer for checking that package variables are not reassigned.
func NewAnalyzer() *analysis.Analyzer {
	return analyzer.New()
}
