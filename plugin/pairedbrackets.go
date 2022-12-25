package main

import (
	"github.com/maratori/pairedbrackets/pkg/pairedbrackets"
	"golang.org/x/tools/go/analysis"
)

var /* const */ AnalyzerPlugin = Plugin{} //nolint:gochecknoglobals // const

type Plugin struct{}

func (Plugin) GetAnalyzers() []*analysis.Analyzer {
	return []*analysis.Analyzer{
		pairedbrackets.NewAnalyzer(),
	}
}
