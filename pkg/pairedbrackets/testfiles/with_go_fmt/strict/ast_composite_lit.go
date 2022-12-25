package x

func _() {

	// good - empty
	_ = []int{}

	// good - one line, one element
	_ = []int{1}

	// good - one line, several elements
	_ = []int{1, 2, 3}

	// good - multiline
	_ = []int{
		1,
		2,
		3,
	}

	// good - multiline, elements are not validated (it should be different linter)
	_ = []int{
		1, 2,
		3,
	}

	// good - last item exception
	_ = []any{1, 2, 3, `x
	y`}

	// bad left - good right
	_ = []int{1, 2, // want `^left brace should either be the last character on a line or be on the same line with the last composite element$`
		3,
	}

	// bad left - right is ignored
	_ = []int{1, 2, // want `^left brace should either be the last character on a line or be on the same line with the last composite element$`
		3}

	// bad right - next
	_ = []int{
		1, 2, 3} // want `^right brace should be on the next line$`

	// bad right - previous, one line
	// ../../../testdata/no_go_fmt/strict/ast_composite_lit.go

	// bad right - previous, multiline
	_ = []any{1, 2, 3, `x
	y`,
	} // want `^right brace should be on the previous line$`
}
