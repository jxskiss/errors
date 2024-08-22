package errors

import "fmt"

// Cause returns the first error in err's tree, if possible.
// It unwraps err if it contains a method Unwrap or Cause
// returning an error.
// It unwraps error returned by Unwrap or Cause recursively.
//
// If err is nil or does not contain the method, the input err is returned.
//
// Deprecated: errors.Is, errors.As, errors.Unwrap are preferred.
func Cause(err error) error {
	cause := Unwrap(err)
	if cause == nil {
		if x, ok := err.(interface{ Cause() error }); ok {
			cause = x.Cause()
		}
	}
	if cause == nil || cause == err {
		return err
	}
	return Cause(cause)
}

// AddStack is an alias function of Wrap.
// Deprecated: use Wrap instead.
func AddStack(err error) error {
	return skip1.Wrap(err)
}

// WithStack is an alias function of Wrap.
// Deprecated: use Wrap instead.
func WithStack(err error) error {
	return skip1.Wrap(err)
}

// WithMessage is a wrapper function of Errorf.
// Deprecated: use Errorf instead.
func WithMessage(err error, message string) error {
	return skip1.Errorf("%s: %w", message, err)
}

// WithMessagef is a wrapper function of Errorf.
// Deprecated: use Errorf instead.
func WithMessagef(err error, format string, args ...any) error {
	msg := fmt.Sprintf(format, args...)
	return skip1.Errorf("%s: %w", msg, err)
}
