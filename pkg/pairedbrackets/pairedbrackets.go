package pairedbrackets

import (
	"golang.org/x/tools/go/analysis"
)

// NewAnalyzer returns Analyzer ensures that bracket starts/ends line or is paired on the same line.
func NewAnalyzer() *analysis.Analyzer {
	return &analysis.Analyzer{
		Name: "pairedbrackets",
		Doc:  "linter ensures that bracket starts/ends line or is paired on the same line",
		Run: func(pass *analysis.Pass) (interface{}, error) {
			return nil, nil
		},
	}
}
