package errors

import (
	"strings"
	"testing"
)

func errWrapper(err error) error {
	return StackSkip(1).Wrap(err)
}

func TestStackSkip(t *testing.T) {
	err1 := errWrapper(inner())
	stacktrace1 := Stacktrace(err1, "")
	if strings.Contains(stacktrace1, "errWrapper") {
		t.Fatalf("err1 stacktrace should not contain frame errWrapper")
	}

	err2 := errWrapper(wrap5())
	stacktrace2 := Stacktrace(err2, "")
	if strings.Contains(stacktrace2, "errWrapper") {
		t.Fatalf("err2 stacktrace should not contain frame errWrapper")
	}
}
