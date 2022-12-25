package x

// tests for receiver: ../../../testfiles/with_go_fmt/strict/ast_func_decl.go

// good - external non-Go function produces [ast.FuncDecl].Body == nil
func _()

// good - empty
func _() {}

// good - one line, one statement
func _() { print() }

// good - one line, two statements
func _() { print(); print() }

// good - multiline
func _() {
	print()
	print()
}

// good - multiline, statements are not validated (it should be different linter)
func _() {
	print(); print()
	print()
}

// good - last item exception
func _() { print(); print(struct {
}{}) }

// bad left - good right
func _() { print() // want `^left brace should either be the last character on a line or be on the same line with the last statement$`
	print()
}

// bad left - right is ignored
func _() { print() // want `^left brace should either be the last character on a line or be on the same line with the last statement$`
	print() }

// bad right - next
func _() {
	print()
	print() } // want `^right brace should be on the next line$`

// bad right - previous, one line
func _() { print(); print()
} // want `^right brace should be on the previous line$`

// bad right - previous, multiline
func _() { print(); print(struct {
}{})
} // want `^right brace should be on the previous line$`
