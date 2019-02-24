package errors

import (
	"bytes"
	"fmt"
	"strings"
)

func Append(err error, errs ...error) error {
	switch err := err.(type) {
	case *MultiError:
		var merr MultiError
		// Typed nils can reach here, initialize in case of nil.
		if err == nil {
			merr = make([]error, 0, len(errs))
		} else {
			merr = *err
		}

		// Go through each error and flatten.
		for _, e := range errs {
			switch e := e.(type) {
			case *MultiError:
				if e != nil {
					merr = append(merr, e.Errors()...)
				}
			case MultiError:
				if len(e) > 0 {
					merr = append(merr, e...)
				}
			default:
				if e != nil {
					merr = append(merr, e)
				}
			}
		}
		return merr

	case MultiError:
		for _, e := range errs {
			switch e := e.(type) {
			case *MultiError:
				if e != nil {
					err = append(err, e.Errors()...)
				}
			case MultiError:
				if len(e) > 0 {
					err = append(err, e...)
				}
			default:
				if e != nil {
					err = append(err, e)
				}
			}
		}
		return err

	default:
		newErrs := make([]error, 0, len(errs)+1)
		if err != nil {
			newErrs = append(newErrs, err)
		}
		newErrs = append(newErrs, errs...)
		return MultiError(newErrs)
	}
}

func ErrOrNil(err error) error {
	if err == nil {
		return nil
	}
	switch err := err.(type) {
	case *MultiError:
		if err == nil || len(*err) == 0 {
			return nil
		}
	case MultiError:
		if len(err) == 0 {
			return nil
		}
	case *sizedError:
		if err == nil || err.count == 0 {
			return nil
		}
	}
	return err
}

// MultiError wraps a slice of errors and implements the error interface.
// This can be used to collect a bunch of errors (such as during form validation)
// and then return them all together as a single error.
type MultiError []error

func (E MultiError) Error() string {
	if len(E) == 0 {
		return "<nil>"
	}
	return string(formatSingleLine(E.Errors()))
}

func (E MultiError) Errors() []error {
	return []error(E)
}

func (E MultiError) Format(f fmt.State, c rune) {
	if c == 'v' && f.Flag('+') {
		f.Write(formatMultiLine(E.Errors()))
	} else {
		f.Write(formatSingleLine(E.Errors()))
	}
}

func NewSizedError(size int) *sizedError {
	return &sizedError{
		errs: make([]error, size),
		size: size,
	}
}

type sizedError struct {
	errs  []error
	size  int
	count int
}

func (E *sizedError) Append(errs ...error) {
	for _, err := range errs {
		if err != nil {
			E.errs[E.count%E.size] = err
			E.count++
		}
	}
}

func (E *sizedError) Error() string {
	if E == nil || E.count == 0 {
		return "<nil>"
	}
	var buf bytes.Buffer
	var first = true
	for _, err := range E.Errors() {
		if first {
			first = false
		} else {
			buf.Write(_singlelineSeparator)
		}
		buf.WriteString(err.Error())
	}
	return string(buf.Bytes())
}

// Errors returns the errors as a slice in reversed order, if the underlying
// errors are more than size, only size errors will be returned, plus an
// additional error indicates the omitted error count.
func (E *sizedError) Errors() (errors []error) {
	if E.count == 0 {
		return nil
	}
	if E.count <= E.size {
		errors = make([]error, 0, E.count)
		for i := E.count - 1; i >= 0; i-- {
			errors = append(errors, E.errs[i])
		}
		return errors
	}
	errors = make([]error, 0, E.count+1)
	for i := E.count%E.size - 1; i >= 0; i-- {
		errors = append(errors, E.errs[i])
	}
	for i := E.size - 1; i >= E.count%E.size; i-- {
		errors = append(errors, E.errs[i])
	}
	errors = append(errors, fmt.Errorf("and %d more errors omitted", E.count-E.size))
	return errors
}

func (E *sizedError) Format(f fmt.State, c rune) {
	if c == 'v' && f.Flag('+') {
		f.Write(formatMultiLine(E.Errors()))
	} else {
		f.Write(formatSingleLine(E.Errors()))
	}
}

var (
	// Separator for single-line error messages.
	_singlelineSeparator = []byte("; ")

	_newline = []byte("\n")

	// Prefix for multi-line messages
	_multilinePrefix = []byte("the following errors occurred:")

	// Prefix for the first and following lines of an item in a list of
	// multi-line error messages.
	//
	// For example, if a single item is:
	//
	// 	foo
	// 	bar
	//
	// It will become,
	//
	// 	 -  foo
	// 	    bar
	_multilineSeparator = []byte("\n -  ")
	_multilineIndent    = []byte("    ")
)

func formatSingleLine(errs []error) []byte {
	var buf bytes.Buffer
	var first = true
	for _, err := range errs {
		if first {
			first = false
		} else {
			buf.Write(_singlelineSeparator)
		}
		buf.WriteString(err.Error())
	}
	return buf.Bytes()
}

func formatMultiLine(errs []error) []byte {
	var buf bytes.Buffer
	buf.Write(_multilinePrefix)
	for _, err := range errs {
		buf.Write(_multilineSeparator)
		s := fmt.Sprintf("%+v", err)
		first := true
		for len(s) > 0 {
			if first {
				first = false
			} else {
				buf.Write(_multilineIndent)
			}
			idx := strings.IndexByte(s, '\n')
			if idx < 0 {
				idx = len(s) - 1
			}
			buf.WriteString(s[:idx+1])
			s = s[idx+1:]
		}
	}
	return buf.Bytes()
}
