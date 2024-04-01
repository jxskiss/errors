package errors

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

/*
 * Empty lines to keep the line numbers not changed.
 *
 *
 *
 */

func inner() error {
	return New("test error")
}

func wrap1() error {
	err := inner()
	return fmt.Errorf("wrap1: %w", err)
}

func wrap2() error {
	err := wrap1()
	return Wrap(err)
}

func wrap3() error {
	err := wrap2()
	return Errorf("wrap3: %w", err)
}

func wrap4() error {
	err := wrap3()
	return Wrap(err)
}

func wrap5() error {
	return wrap4()
}

func TestStacktrace(t *testing.T) {
	err := wrap5()
	if err == nil {
		t.Fatalf("err should not be nil")
	}
	stacktrace := Stacktrace(err, "")
	wantStrs := []string{
		"jxskiss/errors/stack_test.go:28  (github.com/jxskiss/errors/v2.wrap2)",
		"jxskiss/errors/stack_test.go:32  (github.com/jxskiss/errors/v2.wrap3)",
		"jxskiss/errors/stack_test.go:37  (github.com/jxskiss/errors/v2.wrap4)",
		"jxskiss/errors/stack_test.go:42  (github.com/jxskiss/errors/v2.wrap5)",
		"jxskiss/errors/stack_test.go:46  (github.com/jxskiss/errors/v2.TestStacktrace)",
	}
	t.Logf("\n%s\n", stacktrace)
	for _, str := range wantStrs {
		if !strings.Contains(stacktrace, str) {
			t.Errorf("not found but want %q in stacktrace", str)
		}
	}
	notWant := "jxskiss/errors/stack.go:"
	if strings.Contains(stacktrace, notWant) {
		t.Errorf("found but not want %q in stacktrace", notWant)
	}
}

func TestErrorf(t *testing.T) {
	err := wrap5()
	if err == nil {
		t.Fatalf("err should not be nil")
	}
	var stackErr *withStack
	if !As(err, &stackErr) {
		t.Fatalf("err should contain stack frames")
	}
}

func TestWrap(t *testing.T) {
	err := wrap5()
	if err == nil {
		t.Fatalf("err should not be nil")
	}
	var stackErr *withStack
	if !As(err, &stackErr) {
		t.Fatalf("err should contain stack frames")
	}
}

func TestWrapNew(t *testing.T) {
	f1 := func() error {
		return WrapNew("test error from WrapNew")
	}
	f2 := func() error {
		return fmt.Errorf("fmt.Errorf: %w", f1())
	}
	err := f2()
	want := "fmt.Errorf: test error from WrapNew"
	if err.Error() != want {
		t.Fatalf("want %q but got %q", want, err.Error())
	}
	if len(Frames(err)) == 0 {
		t.Fatalf("err should contain stack frames")
	}
}

func TestDetails(t *testing.T) {
	f1 := func() error {
		return WrapNew("test error from WrapNew", 1, "abc")
	}
	f2 := func() error {
		return fmt.Errorf("fmt.Errorf: %w", f1())
	}
	f3 := func() error {
		return Wrap(f2(), 2, "def")
	}
	err := f3()
	wantErrMsg := "fmt.Errorf: test error from WrapNew"
	if err.Error() != wantErrMsg {
		t.Fatalf("want %q but got %q", wantErrMsg, err.Error())
	}
	if len(Frames(err)) == 0 {
		t.Fatalf("err should contain stack frames")
	}
	wantDetails := []any{2, "def", 1, "abc"}
	if !reflect.DeepEqual(Details(err), wantDetails) {
		t.Fatalf("got unexpected error details: %q", Details(err))
	}
}

func TestFrames(t *testing.T) {
	err := wrap5()
	if err == nil {
		t.Fatalf("err should not be nil")
	}
	if len(Frames(err)) == 0 {
		t.Fatalf("err should contain stack frames")
	}
}
