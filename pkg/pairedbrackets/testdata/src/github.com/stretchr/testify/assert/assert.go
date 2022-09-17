package assert

type TestingT interface {
	Errorf(format string, args ...interface{})
}

func Equal(t TestingT, expected, actual interface{}, msgAndArgs ...interface{}) bool {
	return true
}

type Assertions struct {
	t TestingT
}

func New(t TestingT) *Assertions {
	return &Assertions{
		t: t,
	}
}

func (a *Assertions) JSONEq(expected string, actual string, msgAndArgs ...interface{}) bool {
	return true
}

func (a *Assertions) Zero(i interface{}, msgAndArgs ...interface{}) bool {
	return true
}

func MyCustom(t TestingT, expected, actual interface{}, msgAndArgs ...interface{}) bool {
	return true
}
