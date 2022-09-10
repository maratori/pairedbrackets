package x

func _() {

	// good - empty
	type _ interface{}

	// good - one line, one method
	type _ interface{ Method() }

	// good - one line, two methods
	type _ interface{ Method1(); Method2() }

	// good - multiline
	type _ interface {
		Method1()
		Method2()
	}

	// good - multiline, methods are not validated (it should be different linter)
	type _ interface {
		Method1(); Method2()
		Method3()
	}

	// good - last item exception
	type _ interface { Method1(); Method2(struct {
	}) }

	// bad left - good right
	type _ interface { Method1() // want `^left brace should either be the last character on a line or be on the same line with the last method$`
		Method2()
	}

	// bad left - right is ignored
	type _ interface { Method1() // want `^left brace should either be the last character on a line or be on the same line with the last method$`
		Method2() }

	// bad right - next
	type _ interface {
		Method1()
		Method2() } // want `^right brace should be on the next line$`

	// bad right - previous, one line
	type _ interface { Method1(); Method2()
	} // want `^right brace should be on the previous line$`

	// bad right - previous, multiline
	type _ interface { Method1(); Method2(struct {
	})
	} // want `^right brace should be on the previous line$`

}
