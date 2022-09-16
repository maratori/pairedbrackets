package x

import (
	"fmt"
)

func _() {

	// other tests: ../../testfiles/with_go_fmt/ast_call_expr.go

	// bad right - previous, one line
	fmt.Printf("%s %d", "xxx", 10,
	) // want `^right parenthesis should be on the previous line$`

}
