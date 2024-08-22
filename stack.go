package errors

import (
	"fmt"
	"runtime"
	"strings"
)

// Errorf formats a string error using fmt.Errorf and wraps
// it with stack frames.
// If the formatted error already has stack frames,
// it does not add duplicate stack frames.
func Errorf(format string, a ...any) error {
	err := fmt.Errorf(format, a...)
	return wrapError(0, err, nil)
}

// Wrap wraps an error with stack frames.
// If the error already has stack frames, it does not add duplicate
// stack frames.
func Wrap(err error, details ...any) error {
	return wrapError(0, err, details)
}

// WrapNew returns an error that formats as the given text,
// it also wraps the error with stack frames.
func WrapNew(text string, details ...any) error {
	err := New(text)
	return wrapError(0, err, details)
}

// Details gets wrapped details from err if available.
func Details(err error) []any {
	var detailsErr interface{ Details() []any }
	if As(err, &detailsErr) {
		return detailsErr.Details()
	}
	return nil
}

// Frames gets stack frames from err if any error in err's tree
// is wrapped by functions in this package.
//
// The only valid use for the return value is as an argument to
// [runtime.CallersFrames]. In particular, it must not be passed to
// [runtime.FuncForPC].
func Frames(err error) []uintptr {

	// For long-running programs, the number of different error frames
	// is usually small, and they don't change during program running.
	// Here, we choose to return []uintptr instead of [runtime.CallerFrames]
	// so that the caller can decide whether to cache the formatted
	// stacktrace to reduce performance overhead at runtime.

	var frames []uintptr
	var stackErr interface{ Frames() []uintptr }
	if As(err, &stackErr) {
		pcs := stackErr.Frames()
		frames = pcs[:]
		for i, pc := range pcs {
			if pc == 0 {
				frames = pcs[:i]
				break
			}
		}
	}
	return frames
}

// Stacktrace returns a formatted stacktrace if err contains stack frames.
func Stacktrace(err error, indent string) string {
	callers := Frames(err)
	if len(callers) == 0 {
		return ""
	}
	frames := runtime.CallersFrames(callers)
	if frames == nil {
		return ""
	}

	var buf strings.Builder
	for {
		frame, more := frames.Next()
		if buf.Len() > 0 {
			buf.WriteByte('\n')
		}
		buf.WriteString(indent)
		formatFrame(&buf, frame)
		if !more { // no more frames to
			break
		}
	}
	return buf.String()
}

func formatFrame(b *strings.Builder, frame runtime.Frame) {
	file, line, fnName := frame.File, frame.Line, frame.Function
	setDefault(&file, "unknown")
	setDefault(&fnName, "unknown")
	fmt.Fprintf(b, "%s:%d  (%s)", file, line, fnName)
}

func wrapError(skip int, err error, details []any) error {
	if err == nil {
		return nil
	}
	origErr := err
	if len(details) > 0 {
		innerDetails := Details(origErr)
		details = append(details, innerDetails...)
		err = &withDetails{
			error:   err,
			details: details,
		}
	}
	var stackErr *withStack
	if !As(origErr, &stackErr) {
		var pcs stack
		runtime.Callers(skip+3, pcs[:])
		err = &withStack{
			error: err,
			stack: &pcs,
		}
	}
	return err
}

type stack = [20]uintptr

type withStack struct {
	error
	stack *stack
}

func (e *withStack) Unwrap() error { return e.error }

func (e *withStack) Frames() []uintptr {
	if e.stack == nil {
		return nil
	}
	return e.stack[:]
}

type withDetails struct {
	error
	details []any
}

func (e *withDetails) Unwrap() error { return e.error }

func (e *withDetails) Details() []any { return e.details }

func setDefault[T comparable](dst *T, val T) {
	var zero T
	if *dst == zero {
		*dst = val
	}
}
