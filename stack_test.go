package errors

import (
	"fmt"
	"strings"
	"testing"
)

/*
 * Empty lines to keep the line numbers not changed.
 *
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

func TestGetStacktrace(t *testing.T) {
	err := wrap5()
	if err == nil {
		t.Fatalf("err should not be nil")
	}
	stacktrace := GetStacktrace(err, "")
	wantStrs := []string{
		"jxskiss/errors/stack_test.go:28  (github.com/jxskiss/errors/v2.wrap2)",
		"jxskiss/errors/stack_test.go:32  (github.com/jxskiss/errors/v2.wrap3)",
		"jxskiss/errors/stack_test.go:37  (github.com/jxskiss/errors/v2.wrap4)",
		"jxskiss/errors/stack_test.go:42  (github.com/jxskiss/errors/v2.wrap5)",
		"jxskiss/errors/stack_test.go:46  (github.com/jxskiss/errors/v2.TestGetStacktrace)",
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
		t.Fatalf("err should contains stack frames")
	}
}

func TestWrap(t *testing.T) {
	err := wrap5()
	if err == nil {
		t.Fatalf("err should not be nil")
	}
	var stackErr *withStack
	if !As(err, &stackErr) {
		t.Fatalf("err should contains stack frames")
	}
}

func TestGetFrames(t *testing.T) {
	err := wrap5()
	if err == nil {
		t.Fatalf("err should not be nil")
	}
	if len(GetFrames(err)) == 0 {
		t.Fatalf("err should contains stack frames")
	}
}
