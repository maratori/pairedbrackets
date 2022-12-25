package x

func _() {

	// good - [ast.TypeSpec].TypeParams == nil
	type _ int

	// good - one line, one type parameter
	type _[T int] int

	// good - one line, several elements
	type _[T int, V string] int

	// good - multiline
	type _[
		T int,
		V string,
	] int

	// good - multiline, type parameters are not validated (it should be different linter)
	type _[
		T int, V string,
		R bool,
	] int

	// good - last item exception
	type _[T int, V struct {
	}] int

	// bad left - good right
	type _[T int, V string, // want `^left bracket should either be the last character on a line or be on the same line with the last type parameter$`
		R bool,
	] int

	// bad left - right is ignored
	type _[T int, V string, // want `^left bracket should either be the last character on a line or be on the same line with the last type parameter$`
		R bool] int

	// bad right - next
	type _[
		T int,
		V string] int // want `^right bracket should be on the next line$`

	// bad right - previous, one line
	type _[T int, V string,
	] int // want `^right bracket should be on the previous line$`

	// bad right - previous, multiline
	type _[T int, V struct {
	},
	] int // want `^right bracket should be on the previous line$`

}
