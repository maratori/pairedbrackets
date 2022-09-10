package x

func _() {

	type x[T any] struct{}

	// good - one line
	type _ x[int]

	// good - multiline
	type _ x[
		struct {
		},
	]

	// good - last item exception
	type _ x[struct {
	}]

	// bad right - next
	type _ x[
		int] // want `^right bracket should be on the next line$`

	// bad right - previous, one line
	type _ x[int,
	] // want `^right bracket should be on the previous line$`

	// bad right - previous, multiline
	type _ x[struct {
	},
	] // want `^right bracket should be on the previous line$`

}
