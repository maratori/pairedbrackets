package x

func _() {

	var x []int
	var f func(struct{}) int

	// good - no indexes
	_ = x[:]

	// good - one line, several indexes
	_ = x[:10:20]

	// good - multiline
	_ = x[
		1:
	]

	// good - last item exception
	_ = x[1:2:f(struct{}{
	})]

	// bad left
	_ = x[1:2: // want `^left bracket should either be the last character on a line or be on the same line with the last index$`
	3]

	// bad right - next
	_ = x[
		:10:20] // want `^right bracket should be on the next line$`

	// bad right - previous, one line
	_ = x[1:
	] // want `^right bracket should be on the previous line$`

	// bad right - previous, multiline
	_ = x[f(struct{}{
	}):
	] // want `^right bracket should be on the previous line$`

}
