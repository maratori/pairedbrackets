package x

func _() {

	var x any

	// good - empty
	switch x.(type) {}

	// good - one line, one statement
	switch x.(type) { default: }

	// good - one line, two statements
	switch x.(type) { case any: print(); default: }

	// good - multiline
	switch x.(type) {
	case any:
		print()
	default:
		print()
	}

	// good - last item exception
	switch x.(type) { default: print(); print(struct {
	}{}) }

	// bad left - good right
	switch x.(type) { case any: print() // want `^left brace should either be the last character on a line or be on the same line with the last statement$`
	default: print()
	}

	// bad left - right is ignored
	switch x.(type) { case any: print() // want `^left brace should either be the last character on a line or be on the same line with the last statement$`
	default: print() }

	// bad right - next
	switch x.(type) {
	case any: print()
	default: print() } // want `^right brace should be on the next line$`

	// bad right - previous, one line
	switch x.(type) { case any: print(); default: print()
	} // want `^right brace should be on the previous line$`

	// bad right - previous, multiline
	switch x.(type) { default: print(); print(struct {
	}{})
	} // want `^right brace should be on the previous line$`

}
