package x

func _() {

	var x []int

	// good - empty
	for range x {}

	// good - one line, one statement
	for range x { print() }

	// good - one line, two statements
	for range x { print(); print() }

	// good - multiline
	for range x {
		print()
		print()
	}

	// good - multiline, statements are not validated (it should be different linter)
	for range x {
		print(); print()
		print()
	}

	// good - last item exception
	for range x { print(); print(struct {
	}{}) }

	// bad left - good right
	for range x { print() // want `^left brace should either be the last character on a line or be on the same line with the last statement$`
		print()
	}

	// bad left - right is ignored
	for range x { print() // want `^left brace should either be the last character on a line or be on the same line with the last statement$`
		print() }

	// bad right - next
	for range x {
		print()
		print() } // want `^right brace should be on the next line$`

	// bad right - previous, one line
	for range x { print(); print()
	} // want `^right brace should be on the previous line$`

	// bad right - previous, multiline
	for range x { print(); print(struct {
	}{})
	} // want `^right brace should be on the previous line$`

}
