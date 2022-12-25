package x

// if
func _() {

	// good - empty
	if true {}

	// good - one line, one statement
	if true { print() }

	// good - one line, two statements
	if true { print(); print() }

	// good - multiline
	if true {
		print()
		print()
	}

	// good - multiline, statements are not validated (it should be different linter)
	if true {
		print(); print()
		print()
	}

	// good - last item exception
	if true { print(); print(struct {
	}{}) }

	// bad left - good right
	if true { print() // want `^left brace should either be the last character on a line or be on the same line with the last statement$`
		print()
	}

	// bad left - right is ignored
	if true { print() // want `^left brace should either be the last character on a line or be on the same line with the last statement$`
		print() }

	// bad right - next
	if true {
		print()
		print() } // want `^right brace should be on the next line$`

	// bad right - previous, one line
	if true { print(); print()
	} // want `^right brace should be on the previous line$`

	// bad right - previous, multiline
	if true { print(); print(struct {
	}{})
	} // want `^right brace should be on the previous line$`

}

// else
func _() {

	// else if produces [ast.IfStmt].Else == [ast.IfStmt], not [ast.BlockStmt]
	if true {} else if true {}

	// good - empty
	if true {} else {}

	// good - one line, one statement
	if true {} else { print() }

	// good - one line, two statements
	if true {} else { print(); print() }

	// good - multiline
	if true {} else {
		print()
		print()
	}

	// good - multiline, statements are not validated (it should be different linter)
	if true {} else {
		print(); print()
		print()
	}

	// good - last item exception
	if true {} else { print(); print(struct {
	}{}) }

	// bad left - good right
	if true {} else { print() // want `^left brace should either be the last character on a line or be on the same line with the last statement$`
		print()
	}

	// bad left - right is ignored
	if true {} else { print() // want `^left brace should either be the last character on a line or be on the same line with the last statement$`
		print() }

	// bad right - next
	if true {} else {
		print()
		print() } // want `^right brace should be on the next line$`

	// bad right - previous, one line
	if true {} else { print(); print()
	} // want `^right brace should be on the previous line$`

	// bad right - previous, multiline
	if true {} else { print(); print(struct {
	}{})
	} // want `^right brace should be on the previous line$`

}
