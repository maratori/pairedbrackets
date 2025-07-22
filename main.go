package main

import (
	"context"
	"embed"
	"errors"
	"flag"
	"fmt"
	"io"
	"io/fs"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"regexp"
	"syscall"

	"golang.org/x/tools/go/analysis/singlechecker"

	"github.com/maratori/pairedbrackets/pkg/pairedbrackets"
)

//go:embed pkg
//go:embed plugin
//go:embed go.*
var self embed.FS

const (
	buildFlagName    = "build-golangci-lint-plugin"
	buildFlagDefault = false
	buildFlagUsage   = "build plugin for your version of golangci-lint, see other flags"

	goPathFlagName    = "go-path"
	goPathFlagDefault = "go"
	goPathFlagUsage   = "path to go executable, is used only with -" + buildFlagName

	linterFlagName    = "golangci-lint"
	linterFlagDefault = "golangci-lint"
	linterFlagUsage   = "path to golangci-lint executable, is used only with -" + buildFlagName

	outputFlagName    = "plugin-output"
	outputFlagDefault = "pairedbrackets.so"
	outputFlagUsage   = "plugin output path, is used only with -" + buildFlagName
)

func main() {
	set := flag.NewFlagSet("", flag.ContinueOnError)
	set.SetOutput(io.Discard)
	build := set.Bool(buildFlagName, buildFlagDefault, buildFlagUsage)
	goPath := set.String(goPathFlagName, goPathFlagDefault, goPathFlagUsage)
	linter := set.String(linterFlagName, linterFlagDefault, linterFlagUsage)
	output := set.String(outputFlagName, outputFlagDefault, outputFlagUsage)
	if set.Parse(os.Args[1:]) == nil && *build {
		err := buildPlugin(*goPath, *linter, *output)
		if err != nil {
			fmt.Println(err) //nolint:forbidigo // better than log
			os.Exit(1)
		}
		os.Exit(0)
	}

	analyzer := pairedbrackets.NewAnalyzer()
	set.VisitAll(func(f *flag.Flag) {
		analyzer.Flags.Var(f.Value, f.Name, f.Usage) // for documentation
	})
	singlechecker.Main(analyzer)
}

func buildPlugin(goPath string, linterPath string, outputPath string) error {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer stop()

	goPath, linterPath, outputPath, err := resolvePaths(goPath, linterPath, outputPath)
	if err != nil {
		return err
	}

	output, err := exec.CommandContext(ctx, goPath, "version", "-m", linterPath).CombinedOutput()
	if err != nil {
		return fmt.Errorf("can't get version of golangci-lint dependencies: %w", err)
	}

	match := regexp.MustCompile(`dep\s+golang.org/x/tools\s+(v\S+)`).FindSubmatch(output)
	if len(match) == 0 {
		return errors.New("golang.org/x/tools not found in golangci-lint dependencies")
	}
	lib := "golang.org/x/tools@" + string(match[1])

	temp, err := os.MkdirTemp("", "")
	if err != nil {
		return fmt.Errorf("temp dir not created: %w", err)
	}
	defer os.RemoveAll(temp)

	err = writeSelfSourceCodeToDisk(temp)
	if err != nil {
		return fmt.Errorf("can't write pairedbrackets source code: %w", err)
	}

	err = os.Chdir(temp)
	if err != nil {
		return fmt.Errorf("can't change wording directory: %w", err)
	}

	output, err = exec.CommandContext(ctx, goPath, "get", lib).CombinedOutput()
	if err != nil {
		return fmt.Errorf("can't get %s: %w\n%s", lib, err, output)
	}

	output, err = exec.CommandContext(
		ctx,
		goPath,
		"build",
		"-buildmode=plugin",
		"-o",
		outputPath,
		"plugin/pairedbrackets.go",
	).CombinedOutput()
	if err != nil {
		return fmt.Errorf("can't build plugin: %w\n%s", err, output)
	}

	fmt.Printf("pairedbrackets.so is built with %s\n", lib) //nolint:forbidigo // better than log

	return nil
}

func resolvePaths(goPath string, linterPath string, outputPath string) (string, string, string, error) {
	goPath, err := exec.LookPath(goPath)
	if err != nil {
		return "", "", "", fmt.Errorf("go not found: %w", err)
	}

	linterPath, err = exec.LookPath(linterPath)
	if err != nil {
		return "", "", "", fmt.Errorf("golangci-lint not found: %w", err)
	}

	outputPath, err = filepath.Abs(outputPath)
	if err != nil {
		return "", "", "", fmt.Errorf("can't get absolute output path: %w", err)
	}

	return goPath, linterPath, outputPath, nil
}

func writeSelfSourceCodeToDisk(temp string) error {
	return fs.WalkDir(self, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		fullPath := filepath.Join(temp, path)

		if d.IsDir() {
			return os.MkdirAll(fullPath, 0700)
		}

		content, err := fs.ReadFile(self, path)
		if err != nil {
			return err
		}

		err = os.WriteFile(fullPath, content, 0600)
		if err != nil {
			return err
		}

		return nil
	})
}
