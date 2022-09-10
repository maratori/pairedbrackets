package x

func _() {

	// other tests: ../../testfiles/with_go_fmt/ast_composite_lit.go

	// bad right - previous, one line
	_ = []int{1, 2, 3,
	} // want `^right brace should be on the previous line$`

}
