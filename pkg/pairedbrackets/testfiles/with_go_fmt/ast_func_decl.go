package x

type methodReceiver struct{}

func (methodReceiver) goodOneLine() {}

func (
	methodReceiver,
) goodMultiline() {
}

func (
	methodReceiver) badRightNext() { // want `^right parenthesis should be on the next line$`
}

func (methodReceiver,
) badRightPrevious() { // want `^right parenthesis should be on the previous line$`
}

// tests for body: ../../testdata/no_go_fmt/ast_func_decl.go
