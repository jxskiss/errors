package errors

import (
	"errors"
	"strings"
	"testing"
)

func TestDeprecatedFunctions(t *testing.T) {
	t.Run("Cause", func(t *testing.T) {
		err := wrap5()
		want := "test error"
		if err = Cause(err); err.Error() != want {
			t.Errorf(`want %q, but got %q`, want, err.Error())
		}
	})

	for _, tc := range []struct {
		Name     string
		WrapFunc func(error) error
	}{
		{"AddStack", AddStack},
		{"WithStack", WithStack},
	} {
		t.Run(tc.Name, func(t *testing.T) {
			err := tc.WrapFunc(errors.New("test error"))
			if !errors.As(err, new(*withStack)) {
				t.Errorf("want error type *withStack, but got %T", err)
			}
			stacktrace := Stacktrace(err, "\t")
			t.Logf("stacktrace:\n%s", stacktrace)
			if strings.Contains(stacktrace, "deprecated.go") {
				t.Errorf(`stacktrace contains "deprecatd.go", but should not`)
			}
			if !strings.Contains(stacktrace, "deprecated_test.go") {
				t.Errorf(`stacktrace does not contains "deprecated_test.go"`)
			}
		})
	}

	t.Run("WithMessage", func(t *testing.T) {
		err := WithMessage(errors.New("test error"), "test message")
		if err.Error() != "test message: test error" {
			t.Errorf(`want "test message: test error", but got %q`, err.Error())
		}
		if !errors.As(err, new(*withStack)) {
			t.Errorf("want error type *withStack, but got %T", err)
		}
		stacktrace := Stacktrace(err, "\t")
		t.Logf("stacktrace:\n%s", stacktrace)
		if strings.Contains(stacktrace, "deprecated.go") {
			t.Errorf(`stacktrace contains "deprecatd.go", but should not`)
		}
		if !strings.Contains(stacktrace, "deprecated_test.go") {
			t.Errorf(`stacktrace does not contains "deprecated_test.go"`)
		}
	})

	t.Run("WithMessagef", func(t *testing.T) {
		err := WithMessagef(errors.New("test error"), "test message %d", 123)
		if err.Error() != "test message 123: test error" {
			t.Errorf(`want "test message 123: test error", but got %q`, err.Error())
		}
		if !errors.As(err, new(*withStack)) {
			t.Errorf("want error type *withStack, but got %T", err)
		}
		stacktrace := Stacktrace(err, "\t")
		t.Logf("stacktrace:\n%s", stacktrace)
		if strings.Contains(stacktrace, "deprecated.go") {
			t.Errorf(`stacktrace contains "deprecatd.go", but should not`)
		}
		if !strings.Contains(stacktrace, "deprecated_test.go") {
			t.Errorf(`stacktrace does not contains "deprecated_test.go"`)
		}
	})
}
