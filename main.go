package main

import (
	"github.com/maratori/pairedbrackets/pkg/pairedbrackets"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(pairedbrackets.NewAnalyzer())
}
