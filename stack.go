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
	return wrapError(err)
}

// Wrap wraps and error with stack frames.
// If the error already has stack frames, it does not add duplicate
// stack frames.
func Wrap(err error) error {
	if err == nil {
		return nil
	}
	return wrapError(err)
}

// GetFrames gets stack frames from err if any error in err's tree
// is wrapped by Wrap or Errorf.
//
// The only valid use for the return value is as an argument to
// [runtime.CallersFrames]. In particular, it must not be passed to
// [runtime.FuncForPC].
func GetFrames(err error) []uintptr {

	// For long-running programs, the number of different error frames
	// is usually small, and they don't change during program running.
	// Here, we choose to return []uintptr instead of [runtime.CallerFrames]
	// so that the caller can decide whether to cache the formatted
	// stacktrace to reduce performance overhead at runtime.

	var frames []uintptr
	var stackErr *withStack
	if As(err, &stackErr) && stackErr.stack != nil {
		pcs := *stackErr.stack
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

// GetStacktrace returns a formatted stacktrace if err contains
// stack frames.
func GetStacktrace(err error, indent string) string {
	callers := GetFrames(err)
	if len(callers) == 0 {
		return ""
	}
	frames := runtime.CallersFrames(callers)
	if frames == nil {
		return ""
	}

	var buf strings.Builder
	frame, more := frames.Next()
	for i := 0; more; i++ {
		if i > 0 {
			buf.WriteByte('\n')
		}
		buf.WriteString(indent)
		formatFrame(&buf, frame)
		frame, more = frames.Next()
	}
	return buf.String()
}

func formatFrame(b *strings.Builder, frame runtime.Frame) {
	file, line, fnName := frame.File, frame.Line, frame.Function
	setDefault(&file, "unknown")
	setDefault(&fnName, "unknown")
	fmt.Fprintf(b, "%s:%d  (%s)", file, line, fnName)
}

func wrapError(err error) error {
	var stackErr *withStack
	if As(err, &stackErr) {
		return err
	}
	var pcs stack
	runtime.Callers(3, pcs[:])
	return &withStack{
		inner: err,
		stack: &pcs,
	}
}

type withStack struct {
	inner error
	stack *stack
}

func (e *withStack) Error() string { return e.inner.Error() }

func (e *withStack) Unwrap() error { return e.inner }

type stack = [20]uintptr

func setDefault[T comparable](dst *T, val T) {
	var zero T
	if *dst == zero {
		*dst = val
	}
}
