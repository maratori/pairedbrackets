package x

func _() {

	type x[T, V, R any] struct{}

	// good - one line
	type _ x[int, string, bool]

	// good - multiline
	type _ x[
		int,
		string,
		bool,
	]

	// good - multiline, types are not validated (it should be different linter)
	type _ x[
		int, string,
		bool,
	]

	// good - last item exception
	type _ x[int, string, struct {
	}]

	// bad left - good right
	type _ x[int, string, // want `^left bracket should either be the last character on a line or be on the same line with the last element$`
		bool,
	]

	// bad left - right is ignored
	type _ x[int, string, // want `^left bracket should either be the last character on a line or be on the same line with the last element$`
		bool]

	// bad right - next
	type _ x[
		int, string,
		bool] // want `^right bracket should be on the next line$`

	// bad right - previous, one line
	// ../../../testdata/no_go_fmt/strict/ast_index_list_expr.go

	// bad right - previous, multiline
	type _ x[int, string, struct {
	},
	] // want `^right bracket should be on the previous line$`

}
