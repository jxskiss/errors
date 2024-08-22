package errors

import "fmt"

var skip1 = StackSkip(1)

// StackSkip helps to skip stack frames when wrapping errors,
// for example, in middlewares or util functions.
type StackSkip int

func (skip StackSkip) Errorf(format string, a ...any) error {
	err := fmt.Errorf(format, a...)
	return wrapError(int(skip), err, nil)
}

func (skip StackSkip) New(text string, details ...any) error {
	err := New(text)
	return wrapError(int(skip), err, details)
}

func (skip StackSkip) Wrap(err error, details ...any) error {
	return wrapError(int(skip), err, details)
}
