package errors

import "fmt"

// Skip helps to skip stack frames when wrapping errors,
// for example, in middlewares or util functions.
type Skip int

func (skip Skip) Errorf(format string, a ...any) error {
	err := fmt.Errorf(format, a...)
	return wrapError(int(skip), err, nil)
}

func (skip Skip) New(text string, details ...any) error {
	err := New(text)
	return wrapError(int(skip), err, details)
}

func (skip Skip) Wrap(err error, details ...any) error {
	return wrapError(int(skip), err, details)
}
