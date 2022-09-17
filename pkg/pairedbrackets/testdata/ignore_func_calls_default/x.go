package x

import (
	"fmt"

	"github.com/stretchr/testify/assert"
	alias "github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

func _() {

	assert.Equal(nil, []int{
		1,
	}, nil)

	assert.MyCustom(nil, []int{
		1,
	}, nil)

	alias.Equal(nil, []int{
		1,
	}, nil)

	require.Equalf(nil, []int{
		1,
	}, nil, "")

	assert.New(nil).JSONEq("", "", []int{
		1,
	}, "")

	require.New(nil).Eventually(func() bool {
		return false
	}, 0, 0)

	fmt.Printf("%s %d", // want `^left parenthesis should either be the last character on a line or be on the same line with the last argument$`
		"xxx", 10,
	)

}

type MyTestSuite struct {
	suite.Suite
}

func (suite *MyTestSuite) TestExample() {
	suite.Zero([]int{
		1,
	}, "")
}
