package errors

import (
	"bytes"
	"fmt"
	"io"
	"runtime"
	"strings"

	// Use StackTrace and Frame from pkg/errors to be compatible with logrus/graylog/sentry hooks.
	pkgerr "github.com/pkg/errors"
)

type (
	Frame      = pkgerr.Frame
	StackTrace = pkgerr.StackTrace
)

// StackTracer retrieves the StackTrace
// Generally you would want to use the GetStackTracer function to do that.
type StackTracer interface {
	StackTrace() StackTrace
}

// GetStackTracer will return the first StackTracer in the causer chain.
// This function is used by AddStack to avoid creating redundant stack traces.
//
// You can also use the StackTracer interface on the returned error to get the stack trace.
func GetStackTracer(origErr error) StackTracer {
	var stacked StackTracer
	WalkDeep(origErr, func(err error) bool {
		if stackTracer, ok := err.(StackTracer); ok {
			stacked = stackTracer
			return true
		}
		return false
	})
	return stacked
}

// stackMinLen is a best-guess at the minimum length of a stack trace. It
// doesn't need to be exact, just give a good enough head start for the buffer
// to avoid the expensive early growth.
const stackMinLen = 96

// stack represents a stack of program counters.
type stack []uintptr

func (s *stack) Format(st fmt.State, verb rune) {
	switch verb {
	case 'v':
		switch {
		case st.Flag('+'):
			var b bytes.Buffer
			b.Grow(len(*s) * stackMinLen)
			for _, pc := range *s {
				f := Frame(pc)
				fmt.Fprintf(&b, "\n%+v", f)
			}
			io.Copy(st, &b)
		}
	}
}

func (s *stack) StackTrace() StackTrace {
	f := make([]Frame, len(*s))
	for i := 0; i < len(f); i++ {
		f[i] = Frame((*s)[i])
	}
	return f
}

func callers() *stack {
	return callersSkip(4)
}

func callersSkip(skip int) *stack {
	const depth = 32
	var pcs [depth]uintptr
	n := runtime.Callers(skip, pcs[:])
	var st stack = pcs[0:n]
	return &st
}

// funcname removes the path prefix component of a function's name reported by func.Name().
func funcname(name string) string {
	i := strings.LastIndex(name, "/")
	name = name[i+1:]
	i = strings.Index(name, ".")
	return name[i+1:]
}

// NewStack is for library implementers that want to generate a stack trace.
// Normally you should insted use AddStack to get an error with a stack trace.
//
// The result of this function can be turned into a stack trace by calling .StackTrace()
//
// This function takes an argument for the number of stack frames to skip.
// This avoids putting stack generation function calls like this one in the stack trace.
// A value of 0 will give you the line that called NewStack(0)
// A library author wrapping this in their own function will want to use a value of at least 1.
func NewStack(skip int) StackTracer {
	return callersSkip(skip + 3)
}
