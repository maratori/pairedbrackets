package x

// good - empty
var _ = func() {}

// good - one line, one statement
var _ = func() { print() }

// good - one line, two statements
var _ = func() { print(); print() }

// good - multiline
var _ = func() {
	print()
	print()
}

// good - multiline, statements are not validated (it should be different linter)
var _ = func() {
	print(); print()
	print()
}

// good - last item exception
var _ = func() { print(); print(struct {
}{}) }

// bad left - good right
var _ = func() { print() // want `^left brace should either be the last character on a line or be on the same line with the last statement$`
	print()
}

// bad left - right is ignored
var _ = func() { print() // want `^left brace should either be the last character on a line or be on the same line with the last statement$`
	print() }

// bad right - next
var _ = func() {
	print()
	print() } // want `^right brace should be on the next line$`

// bad right - previous, one line
var _ = func() { print(); print()
} // want `^right brace should be on the previous line$`

// bad right - previous, multiline
var _ = func() { print(); print(struct {
}{})
} // want `^right brace should be on the previous line$`
