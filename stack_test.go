package errors

import (
	"fmt"
	"runtime"
	"testing"

	pkgerr "github.com/pkg/errors"
)

var initpc, _, _, _ = runtime.Caller(0)

type X struct{}

func (x X) val() pkgerr.Frame {
	var pc, _, _, _ = runtime.Caller(0)
	return pkgerr.Frame(pc)
}

func (x *X) ptr() pkgerr.Frame {
	var pc, _, _, _ = runtime.Caller(0)
	return pkgerr.Frame(pc)
}

func TestFrameFormat(t *testing.T) {
	var tests = []struct {
		pkgerr.Frame
		format string
		want   string
	}{{
		pkgerr.Frame(initpc),
		"%s",
		"stack_test.go",
	}, {
		pkgerr.Frame(initpc),
		"%+s",
		"github.com/jxskiss/errors.init\n" +
			"\t.+/github.com/jxskiss/errors/stack_test.go",
	}, {
		pkgerr.Frame(0),
		"%s",
		"unknown",
	}, {
		pkgerr.Frame(0),
		"%+s",
		"unknown",
	}, {
		pkgerr.Frame(initpc),
		"%d",
		"11",
	}, {
		pkgerr.Frame(0),
		"%d",
		"0",
	}, {
		pkgerr.Frame(initpc),
		"%n",
		"init",
	}, {
		func() pkgerr.Frame {
			var x X
			return x.ptr()
		}(),
		"%n",
		`\(\*X\).ptr`,
	}, {
		func() pkgerr.Frame {
			var x X
			return x.val()
		}(),
		"%n",
		"X.val",
	}, {
		pkgerr.Frame(0),
		"%n",
		"",
	}, {
		pkgerr.Frame(initpc),
		"%v",
		"stack_test.go:11",
	}, {
		pkgerr.Frame(initpc),
		"%+v",
		"github.com/jxskiss/errors.init\n" +
			"\t.+/github.com/jxskiss/errors/stack_test.go:11",
	}, {
		pkgerr.Frame(0),
		"%v",
		"unknown:0",
	}}

	for i, tt := range tests {
		testFormatRegexp(t, i, tt.Frame, tt.format, tt.want)
	}
}

func TestFuncname(t *testing.T) {
	tests := []struct {
		name, want string
	}{
		{"", ""},
		{"runtime.main", "main"},
		{"github.com/jxskiss/errors.funcname", "funcname"},
		{"funcname", "funcname"},
		{"io.copyBuffer", "copyBuffer"},
		{"main.(*R).Write", "(*R).Write"},
	}

	for _, tt := range tests {
		got := funcname(tt.name)
		want := tt.want
		if got != want {
			t.Errorf("funcname(%q): want: %q, got %q", tt.name, want, got)
		}
	}
}

func TestStackTrace(t *testing.T) {
	tests := []struct {
		err  error
		want []string
	}{{
		New("ooh"), []string{
			"github.com/jxskiss/errors.TestStackTrace\n" +
				"\t.+/github.com/jxskiss/errors/stack_test.go:123",
		},
	}, {
		Annotate(New("ooh"), "ahh"), []string{
			"github.com/jxskiss/errors.TestStackTrace\n" +
				"\t.+/github.com/jxskiss/errors/stack_test.go:128", // this is the stack of Wrap, not New
		},
	}, {
		Cause(Annotate(New("ooh"), "ahh")), []string{
			"github.com/jxskiss/errors.TestStackTrace\n" +
				"\t.+/github.com/jxskiss/errors/stack_test.go:133", // this is the stack of New
		},
	}, {
		func() error { return New("ooh") }(), []string{
			`github.com/jxskiss/errors.(func·009|TestStackTrace.func1)` +
				"\n\t.+/github.com/jxskiss/errors/stack_test.go:138", // this is the stack of New
			"github.com/jxskiss/errors.TestStackTrace\n" +
				"\t.+/github.com/jxskiss/errors/stack_test.go:138", // this is the stack of New's caller
		},
	}, {
		Cause(func() error {
			return func() error {
				return Errorf("hello %s", fmt.Sprintf("world"))
			}()
		}()), []string{
			`github.com/jxskiss/errors.(func·010|TestStackTrace.func2.1)` +
				"\n\t.+/github.com/jxskiss/errors/stack_test.go:147", // this is the stack of Errorf
			`github.com/jxskiss/errors.(func·011|TestStackTrace.func2)` +
				"\n\t.+/github.com/jxskiss/errors/stack_test.go:148", // this is the stack of Errorf's caller
			"github.com/jxskiss/errors.TestStackTrace\n" +
				"\t.+/github.com/jxskiss/errors/stack_test.go:149", // this is the stack of Errorf's caller's caller
		},
	}}
	for i, tt := range tests {
		ste, ok := tt.err.(interface {
			StackTrace() pkgerr.StackTrace
		})
		if !ok {
			ste = tt.err.(interface {
				Cause() error
			}).Cause().(interface {
				StackTrace() pkgerr.StackTrace
			})
		}
		st := ste.StackTrace()
		for j, want := range tt.want {
			testFormatRegexp(t, i, st[j], "%+v", want)
		}
	}
}

// This comment helps to maintain original line numbers
// Perhaps this test is too fragile :)
func stackTrace() pkgerr.StackTrace {
	return NewStack(0).StackTrace()
	// This comment helps to maintain original line numbers
	// Perhaps this test is too fragile :)
}

func TestStackTraceFormat(t *testing.T) {
	tests := []struct {
		pkgerr.StackTrace
		format string
		want   string
	}{{
		nil,
		"%s",
		`\[\]`,
	}, {
		nil,
		"%v",
		`\[\]`,
	}, {
		nil,
		"%+v",
		"",
	}, {
		nil,
		"%#v",
		`\[\]errors.Frame\(nil\)`,
	}, {
		make(pkgerr.StackTrace, 0),
		"%s",
		`\[\]`,
	}, {
		make(pkgerr.StackTrace, 0),
		"%v",
		`\[\]`,
	}, {
		make(pkgerr.StackTrace, 0),
		"%+v",
		"",
	}, {
		make(pkgerr.StackTrace, 0),
		"%#v",
		`\[\]errors.Frame{}`,
	}, {
		stackTrace()[:2],
		"%s",
		`\[stack_test.go stack_test.go\]`,
	}, {
		stackTrace()[:2],
		"%v",
		`[stack_test.go:207 stack_test.go:254]`,
	}, {
		stackTrace()[:2],
		"%+v",
		"\n" +
			"github.com/jxskiss/errors.stackTrace\n" +
			"\t.+/github.com/jxskiss/errors/stack_test.go:179\n" +
			"github.com/jxskiss/errors.TestStackTraceFormat\n" +
			"\t.+/github.com/jxskiss/errors/stack_test.go:230",
	}, {
		stackTrace()[:2],
		"%#v",
		`\[\]errors.Frame{stack_test.go:179, stack_test.go:238}`,
	}}

	for i, tt := range tests {
		testFormatRegexp(t, i, tt.StackTrace, tt.format, tt.want)
	}
}

func TestNewStack(t *testing.T) {
	got := NewStack(1).StackTrace()
	want := NewStack(1).StackTrace()
	if got[0] != want[0] {
		t.Errorf("NewStack(remove NewStack): want: %v, got: %v", want, got)
	}
	gotFirst := fmt.Sprintf("%+v", got[0])[0:15]
	if gotFirst != "testing.tRunner" {
		t.Errorf("NewStack(): want: %v, got: %+v", "testing.tRunner", gotFirst)
	}
}
