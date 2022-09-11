package x

// parameters
func _() {

	type goodEmpty func()

	type goodOneLineOneParam func(int)

	type goodOneLineTwoParams func(int, string)

	type goodMultiline func(
		int,
		string,
	)

	type goodMultilineParamsNotValidated func(
		int, bool,
		string,
	)

	type goodLastItemException func(int, struct {
	})

	type badLeftGoodRight func(int, // want `^left parenthesis should either be the last character on a line or be on the same line with the last parameter$`
		string,
	)

	type badLeftRightIsIgnored func(int, // want `^left parenthesis should either be the last character on a line or be on the same line with the last parameter$`
		string)

	type badRightNext func(
		int,
		string) // want `^right parenthesis should be on the next line$`

	type badRightPreviousOneLine func(int, string,
	) // want `^right parenthesis should be on the previous line$`

	type badRightPreviousMultiline func(int, struct {
	},
	) // want `^right parenthesis should be on the previous line$`

}

// results
func _() {

	type goodNoParensEmpty func()

	type goodNoParensOne func() int

	type goodEmpty func()

	type goodOneLineOneResult func() int

	type goodOneLineTwoResults func() (int, string)

	type goodMultiline func() (
		int,
		string,
	)

	type goodMultilineResultsNotValidated func() (
		int, bool,
		string,
	)

	type goodLastItemException func() (int, struct {
	})

	type badLeftGoodRight func() (int, // want `^left parenthesis should either be the last character on a line or be on the same line with the last result$`
		string,
	)

	type badLeftRightIsIgnored func() (int, // want `^left parenthesis should either be the last character on a line or be on the same line with the last result$`
		string)

	type badRightNext func() (
		int,
		string) // want `^right parenthesis should be on the next line$`

	type badRightPreviousOneLine func() (int, string,
	) // want `^right parenthesis should be on the previous line$`

	type badRightPreviousMultiline func() (int, struct {
	},
	) // want `^right parenthesis should be on the previous line$`

}

// good - empty
func _() {}

// good - one line, one type
func _[T int]() {}

// good - one line, two types
func _[T int, V string]() {}

// good - multiline
func _[
	T int,
	V string,
]() {
}

// good - multiline, types are not validated (it should be different linter)
func _[
	T int, V string,
	W bool,
]() {
}

// good - last item exception
func _[T int, V struct {
}]() {
}

// bad left - good right
func _[T int, // want `^left bracket should either be the last character on a line or be on the same line with the last type parameter$`
	V string]() {
}

// bad left - right is ignored
func _[T int, // want `^left bracket should either be the last character on a line or be on the same line with the last type parameter$`
	V string,
]() {
}

// bad right - next
func _[
	T int,
	V string]() { // want `^right bracket should be on the next line$`
}

// bad right - previous, one lint
func _[T int, V string,
]() { // want `^right bracket should be on the previous line$`
}

// bad right - previous, multiline
func _[T int, V struct {
},
]() { // want `^right bracket should be on the previous line$`
}
