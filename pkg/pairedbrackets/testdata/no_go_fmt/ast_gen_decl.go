package x

// good - no parentheses
import _ "os"

// good - empty
import ()

// good - one line, one import
import ( _ "fmt" )

// good - one line, two imports
import ( _ "bufio"; _ "bytes" )

// good - multiline
import (
	_ "flag"
	_ "image"
)

// good - multiline, elements are not validated (it should be different linter)
import (
	_ "reflect"; _"go/build"
	_ "go/ast"
)

// bad left - good right
import ( _ "html" // want `^left parenthesis should either be the last character on a line or be on the same line with the last import$`
	_ "hash"
)

// bad left - right is ignored
import ( _ "context" // want `^left parenthesis should either be the last character on a line or be on the same line with the last import$`
	_ "crypto" )

// bad right - next
import (
	_ "embed"
	_ "encoding" )// want `^right parenthesis should be on the next line$`

// bad right - previous
import ( _ "errors"; _ "expvar"
) // want `^right parenthesis should be on the previous line$`


// constant
func _() {

	// good - no parentheses
	const _ = 10

	// good - empty
	const ()

	// good - one line, one constant
	const ( _ = 10 )

	// good - one line, two constants
	const ( _ = 10; _ = 20 )

	// good - multiline
	const (
		_ = 10
		_ = 20
	)

	// good - multiline, constants are not validated (it should be different linter)
	const (
		_ = 10; _ = 20
		_ = 30
	)

	// good - last item exception
	const ( _ = 10; _ = `x
	y`)

	// bad left - good right
	const ( _ = 10 // want `^left parenthesis should either be the last character on a line or be on the same line with the last constant$`
		_ = 20
	)

	// bad left - right is ignored
	const ( _ = 10 // want `^left parenthesis should either be the last character on a line or be on the same line with the last constant$`
		_ = 20 )

	// bad right - next
	const (
		_ = 10
		_ = 20 ) // want `^right parenthesis should be on the next line$`

	// bad right - previous, one line
	const ( _ = 10; _ = 20
	) // want `^right parenthesis should be on the previous line$`

	// bad right - previous, multiline
	const ( _ = 10; _ = `x
	y`
	) // want `^right parenthesis should be on the previous line$`

}

// type
func _() {

	// good - no parentheses
	type _ int

	// good - empty
	type ()

	// good - one line, one type
	type ( _ int )

	// good - one line, two types
	type ( _ int; _ string )

	// good - multiline
	type (
		_ int
		_ string
	)

	// good - multiline, types are not validated (it should be different linter)
	type (
		_ int; _ string
		_ bool
	)

	// good - last item exception
	type ( _ int; _ struct {
	} )

	// bad left - good right
	type ( _ int // want `^left parenthesis should either be the last character on a line or be on the same line with the last type$`
		_ string
	)

	// bad left - right is ignored
	type ( _ int // want `^left parenthesis should either be the last character on a line or be on the same line with the last type$`
		_ string )

	// bad right - next
	type (
		_ int
		_ string ) // want `^right parenthesis should be on the next line$`

	// bad right - previous, one line
	type ( _ int; _ string
	) // want `^right parenthesis should be on the previous line$`

	// bad right - previous, multiline
	type ( _ int; _ struct {
	}
	) // want `^right parenthesis should be on the previous line$`

}

// variable
func _() {

	// good - no parentheses
	var _ int

	// good - empty
	var ()

	// good - one line, one variable
	var ( _ int )

	// good - one line, two variables
	var ( _ int; _ string )

	// good - multiline
	var (
		_ int
		_ string
	)

	// good - multiline, variables are not validated (it should be different linter)
	var (
		_ int; _ string
		_ bool
	)

	// good - last item exception
	var ( _ int; _ struct {
	} )

	// bad left - good right
	var ( _ int // want `^left parenthesis should either be the last character on a line or be on the same line with the last variable$`
		_ string
	)

	// bad left - right is ignored
	var ( _ int // want `^left parenthesis should either be the last character on a line or be on the same line with the last variable$`
		_ string )

	// bad right - next
	var (
		_ int
		_ string ) // want `^right parenthesis should be on the next line$`

	// bad right - previous, one line
	var ( _ int; _ string
	) // want `^right parenthesis should be on the previous line$`

	// bad right - previous, multiline
	var ( _ int; _ struct {
	}
	) // want `^right parenthesis should be on the previous line$`

}
