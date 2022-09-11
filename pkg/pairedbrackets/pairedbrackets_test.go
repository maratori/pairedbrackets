package pairedbrackets_test

import (
	"path/filepath"
	"testing"

	"github.com/maratori/pairedbrackets/pkg/pairedbrackets"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestAnalyzer_WithGofmt(t *testing.T) {
	t.Parallel()

	// dir is named `testfiles` not `testdata` to be able to run `golangci-lint` and `go fmt` for these files
	testdata, err := filepath.Abs("testfiles")
	if err != nil {
		t.FailNow()
	}

	analysistest.Run(t, testdata, pairedbrackets.NewAnalyzer(), "./...")
}