package pairedbrackets_test

import (
	"path/filepath"
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"

	"github.com/maratori/pairedbrackets/pkg/pairedbrackets"
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

func TestAnalyzer_NoGofmt(t *testing.T) {
	t.Parallel()

	// `golangci-lint` and `go fmt` ignore `testdata` dir
	testdata, err := filepath.Abs("testdata")
	if err != nil {
		t.FailNow()
	}

	analysistest.Run(t, testdata, pairedbrackets.NewAnalyzer(), "./no_go_fmt")
}
