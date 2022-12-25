package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPlugin(t *testing.T) {
	analyzers := AnalyzerPlugin.GetAnalyzers()
	require.Len(t, analyzers, 1)
	require.Equal(t, "pairedbrackets", analyzers[0].Name)
}
