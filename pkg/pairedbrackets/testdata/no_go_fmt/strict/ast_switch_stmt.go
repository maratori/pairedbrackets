package x

func _() {

	// good - empty
	switch {}

	// good - one line, one statement
	switch { default: }

	// good - one line, two statements
	switch { case true: print(); default: }

	// good - multiline
	switch {
	case true:
		print()
	default:
		print()
	}

	// good - last item exception
	switch { default: print(); print(struct {
	}{}) }

	// bad left - good right
	switch { case true: print() // want `^left brace should either be the last character on a line or be on the same line with the last statement$`
	default: print()
	}

	// bad left - right is ignored
	switch { case true: print() // want `^left brace should either be the last character on a line or be on the same line with the last statement$`
	default: print() }

	// bad right - next
	switch {
	case true: print()
	default: print() } // want `^right brace should be on the next line$`

	// bad right - previous, one line
	switch { case true: print(); default: print()
	} // want `^right brace should be on the previous line$`

	// bad right - previous, multiline
	switch { default: print(); print(struct {
	}{})
	} // want `^right brace should be on the previous line$`

}
