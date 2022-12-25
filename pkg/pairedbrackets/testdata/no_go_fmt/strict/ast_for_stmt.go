package x

func _() {

	// good - empty
	for {}

	// good - one line, one statement
	for { print() }

	// good - one line, two statements
	for { print(); print() }

	// good - multiline
	for {
		print()
		print()
	}

	// good - multiline, statements are not validated (it should be different linter)
	for {
		print(); print()
		print()
	}

	// good - last item exception
	for { print(); print(struct {
	}{}) }

	// bad left - good right
	for { print() // want `^left brace should either be the last character on a line or be on the same line with the last statement$`
		print()
	}

	// bad left - right is ignored
	for { print() // want `^left brace should either be the last character on a line or be on the same line with the last statement$`
		print() }

	// bad right - next
	for {
		print()
		print() } // want `^right brace should be on the next line$`

	// bad right - previous, one line
	for { print(); print()
	} // want `^right brace should be on the previous line$`

	// bad right - previous, multiline
	for { print(); print(struct {
	}{})
	} // want `^right brace should be on the previous line$`

}
