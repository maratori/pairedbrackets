package pairedbrackets

import (
	"fmt"
	"go/ast"
	"go/token"
	"testing"

	"github.com/stretchr/testify/assert"
)

// This test is written to get 100% code coverage.
func TestGenDeclElement(t *testing.T) {
	t.Parallel()
	tests := map[token.Token]Element{
		token.AND:    Unknown,
		token.CONST:  Constant,
		token.IMPORT: Import,
		token.TYPE:   Type,
		token.VAR:    Variable,
	}
	for tk, el := range tests {
		tk, el := tk, el
		t.Run(fmt.Sprintf("%s -> %s", tk, el), func(t *testing.T) {
			t.Parallel()
			actual := genDeclElement(&ast.GenDecl{Tok: tk})
			assert.Equal(t, el, actual)
		})
	}
}

// This test is written to get 100% code coverage.
func TestBoundariesNotSorted(t *testing.T) {
	t.Parallel()
	firstPos, lastPos, lastEnd, ok := boundaries([]ast.Node{
		&ast.Ident{
			NamePos: 100,
			Name:    "xxx",
		},
		&ast.Ident{
			NamePos: 10,
		},
	})
	assert.EqualValues(t, 10, firstPos)
	assert.EqualValues(t, 100, lastPos)
	assert.EqualValues(t, 102, lastEnd)
	assert.True(t, ok)
}
