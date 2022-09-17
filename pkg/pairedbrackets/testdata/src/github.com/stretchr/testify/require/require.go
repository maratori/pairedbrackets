package require

import "time"

type TestingT interface {
	Errorf(format string, args ...interface{})
}

func Equalf(t TestingT, expected interface{}, actual interface{}, msg string, args ...interface{}) {}

type Assertions struct {
	t TestingT
}

func New(t TestingT) *Assertions {
	return &Assertions{
		t: t,
	}
}

func (a *Assertions) Eventually(condition func() bool, waitFor time.Duration, tick time.Duration, msgAndArgs ...interface{}) {
}
