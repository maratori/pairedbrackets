package x

func _() {

	// good - one line
	_ = ("x")

	// good - last item exception
	_ = (struct {
	}{})

	// bad right - next
	_ = (
		"x") // want `^right parenthesis should be on the next line$`

}
