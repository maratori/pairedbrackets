package imports

import (
	_ "github.com/stretchr/testify/assert"
	_ "github.com/stretchr/testify/require"
	_ "golang.org/x/tools/go/analysis"
	_ "golang.org/x/tools/go/analysis/analysistest"
	_ "golang.org/x/tools/go/analysis/passes/inspect"
	_ "golang.org/x/tools/go/analysis/singlechecker"
	_ "golang.org/x/tools/go/ast/inspector"
)
