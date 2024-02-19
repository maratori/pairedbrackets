package x

import (
	"fmt"
	"net/http"
)

func _() {

	// good - empty
	fmt.Println()

	// good - one line, one element
	fmt.Println("xxx")

	// good - one line, several elements
	fmt.Printf("%s %d", "xxx", 10)

	// good - multiline
	fmt.Printf(
		"%s %d",
		"xxx",
		10,
	)

	// good - multiline, arguments are not validated (it should be different linter)
	fmt.Printf(
		"%s %d",
		"xxx", 10,
	)

	// good - last item exception
	http.HandleFunc("/api/v1", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	// bad left - good right
	fmt.Printf("%s %d", // want `^left parenthesis should either be the last character on a line or be on the same line with the last argument$`
		"xxx", 10,
	)

	// bad left - right is ignored
	fmt.Printf("%s %d", // want `^left parenthesis should either be the last character on a line or be on the same line with the last argument$`
		"xxx", 10)

	// bad right - next
	fmt.Printf(
		"%s %d",
		"xxx",
		10) // want `^right parenthesis should be on the next line$`

	// bad right - previous, one line
	// ../../testdata/no_go_fmt/ast_call_expr.go

	// bad right - previous, multiline
	http.HandleFunc("/api/v1", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
	},
	) // want `^right parenthesis should be on the previous line$`

}
