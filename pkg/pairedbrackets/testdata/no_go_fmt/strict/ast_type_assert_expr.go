package x

func _() {

	var x any

	// good - one line
	_ = x.(any)

	// good - last item exception
	_ = x.(interface {
	})

	// bad right - next
	_ = x.(
		any) // want `^right parenthesis should be on the next line$`

}
