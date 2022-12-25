package x

func _() {

	// good - empty
	type _ struct{}

	// good - one line, one fields
	type _ struct{ _ int }

	// good - one line, two fields
	type _ struct{ _ int; _ string }

	// good - multiline
	type _ struct {
		_ int
		_ string
	}

	// good - multiline, methods are not validated (it should be different linter)
	type _ struct {
		_ int; _ string
		_ bool
	}

	// good - last item exception
	type _ struct { _ int; _ struct {
	} }

	// bad left - good right
	type _ struct { _ int // want `^left brace should either be the last character on a line or be on the same line with the last field$`
		_ string
	}

	// bad left - right is ignored
	type _ struct { _ int // want `^left brace should either be the last character on a line or be on the same line with the last field$`
		_ string }

	// bad right - next
	type _ struct {
		_ int
		_ string } // want `^right brace should be on the next line$`

	// bad right - previous, one line
	type _ struct { _ int; _ string
	} // want `^right brace should be on the previous line$`

	// bad right - previous, multiline
	type _ struct { _ int; _ struct {
	}
	} // want `^right brace should be on the previous line$`

}
