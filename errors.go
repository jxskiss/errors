package errors

import "errors"

// ErrUnsupported is an alias name of errors.ErrUnsupported.
var ErrUnsupported = errors.ErrUnsupported

// New is an alias function of errors.New.
func New(text string) error {
	return errors.New(text)
}

// Is is an alias function of errors.Is.
func Is(err, target error) bool {
	return errors.Is(err, target)
}

// As is an alias function of errors.As.
func As(err error, target any) bool {
	return errors.As(err, target)
}

// Unwrap is an alias function of errors.Unwrap.
func Unwrap(err error) error {
	return errors.Unwrap(err)
}

// Join is an alias function of errors.Join.
func Join(errs ...error) error {
	return errors.Join(errs...)
}
