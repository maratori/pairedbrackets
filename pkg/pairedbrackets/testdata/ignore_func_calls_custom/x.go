package x

import (
	"fmt"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func _() {

	assert.Equal(nil, []int{ // want `^left parenthesis should either be the last character on a line or be on the same line with the last argument$`
		1,
	}, nil)

	require.Equalf(nil, []int{
		1,
	}, nil, "")

	require.New(nil).Eventually(func() bool {
		return false
	}, 0, 0)

	fmt.Printf("%s %d",
		"xxx", 10,
	)

}
