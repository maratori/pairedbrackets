package main

import (
	"golang.org/x/tools/go/analysis"

	"github.com/maratori/pairedbrackets/pkg/pairedbrackets"
)

var /* const */ AnalyzerPlugin = Plugin{} //nolint:gochecknoglobals // const

type Plugin struct{}

func (Plugin) GetAnalyzers() []*analysis.Analyzer {
	return []*analysis.Analyzer{
		pairedbrackets.NewAnalyzer(),
	}
}
