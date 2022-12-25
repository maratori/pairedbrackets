package x

func _() {

	// other tests: ../../../testfiles/with_go_fmt/strict/ast_index_list_expr.go
	type x[T, V any] struct{}

	// bad right - previous, one line
	type _ x[int, string,
	] // want `^right bracket should be on the previous line$`

}
