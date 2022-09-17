package pairedbrackets_test

import (
	"path/filepath"
	"testing"

	"github.com/maratori/pairedbrackets/pkg/pairedbrackets"
	"github.com/stretchr/testify/assert"
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

func TestAnalyzer_NoGofmt(t *testing.T) {
	t.Parallel()

	// `golangci-lint` and `go fmt` ignore `testdata` dir
	testdata, err := filepath.Abs("testdata")
	if err != nil {
		t.FailNow()
	}

	analysistest.Run(t, testdata, pairedbrackets.NewAnalyzer(), "./no_go_fmt")
}

func TestAnalyzer_IgnoreFuncCallsDefault(t *testing.T) {
	t.Parallel()

	// `golangci-lint` and `go fmt` ignore `testdata` dir
	testdata, err := filepath.Abs("testdata")
	if err != nil {
		t.FailNow()
	}

	analysistest.Run(t, testdata, pairedbrackets.NewAnalyzer(), "./ignore_func_calls_default")
}

func TestAnalyzer_IgnoreFuncCallsCustom(t *testing.T) {
	t.Parallel()

	analyzer := pairedbrackets.NewAnalyzer()
	err := analyzer.Flags.Set(pairedbrackets.IgnoreFuncCallsFlagName, `github.com/stretchr/testify/require,^fmt\.`)
	if err != nil {
		t.FailNow()
	}

	// `golangci-lint` and `go fmt` ignore `testdata` dir
	testdata, err := filepath.Abs("testdata")
	if err != nil {
		t.FailNow()
	}

	analysistest.Run(t, testdata, analyzer, "./ignore_func_calls_custom")
}

func TestAnalyzer_InvalidRegexp(t *testing.T) {
	t.Parallel()

	invalid := `\Ca`
	analyzer := pairedbrackets.NewAnalyzer()
	err := analyzer.Flags.Set(pairedbrackets.IgnoreFuncCallsFlagName, invalid)
	assert.EqualError(t, err, "error parsing regexp: invalid escape sequence: `\\C`")
}
