package x

func _() {

	var x chan int

	// good - empty
	select {}

	// good - one line, one statement
	select { default: }

	// good - one line, two statements
	select { case <-x: print(); default: }

	// good - multiline
	select {
	case <-x:
		print()
	default:
		print()
	}

	// good - last item exception
	select { default: print(); print(struct {
	}{}) }

	// bad left - good right
	select { case <-x: print() // want `^left brace should either be the last character on a line or be on the same line with the last statement$`
	default: print()
	}

	// bad left - right is ignored
	select { case <-x: print() // want `^left brace should either be the last character on a line or be on the same line with the last statement$`
	default: print() }

	// bad right - next
	select {
	case <-x: print()
	default: print() } // want `^right brace should be on the next line$`

	// bad right - previous, one line
	select { case <-x: print(); default: print()
	} // want `^right brace should be on the previous line$`

	// bad right - previous, multiline
	select { default: print(); print(struct {
	}{})
	} // want `^right brace should be on the previous line$`

}
