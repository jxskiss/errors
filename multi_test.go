package errors

import (
	"errors"
	"fmt"
	"github.com/matryer/is"
	"strings"
	"testing"
)

func Test_sizedError(t *testing.T) {
	is := is.New(t)
	merr := NewSizedError(5)

	merr.Append(errors.New("error 1"))
	is.Equal(merr.count, 1)
	is.Equal(len(merr.Errors()), 1)
	is.Equal(merr.Errors()[0].Error(), "error 1")

	merr.Append(errors.New("error 2"))
	is.Equal(merr.count, 2)
	merrors := merr.Errors()
	is.Equal(len(merrors), 2)
	is.Equal(merrors[0].Error(), "error 2")
	is.Equal(merrors[1].Error(), "error 1")

	merr.Append(errors.New("error 3"))
	merr.Append(errors.New("error 4"))
	merr.Append(errors.New("error 5"))
	merr.Append(errors.New("error 6"))
	merr.Append(errors.New("error 7"))
	is.Equal(merr.count, 7)
	merrors = merr.Errors()
	is.Equal(len(merrors), 6)
	is.Equal(merrors[0].Error(), "error 7")
	is.Equal(merrors[1].Error(), "error 6")
	is.Equal(merrors[2].Error(), "error 5")
	is.Equal(merrors[3].Error(), "error 4")
	is.Equal(merrors[4].Error(), "error 3")
	is.Equal(merrors[5].Error(), "and 2 more errors omitted")

	// many many errors
	for i := 1; i <= 100; i++ {
		merr.Append(fmt.Errorf("error %d", i))
	}
	is.Equal(merr.count, 107)
	merrors = merr.Errors()
	is.Equal(len(merrors), 6)
	is.Equal(merrors[0].Error(), "error 100")
	is.Equal(merrors[1].Error(), "error 99")
	is.Equal(merrors[2].Error(), "error 98")
	is.Equal(merrors[3].Error(), "error 97")
	is.Equal(merrors[4].Error(), "error 96")
	is.Equal(merrors[5].Error(), "and 102 more errors omitted")
}

func Test_MultiError(t *testing.T) {
	is := is.New(t)
	var err error

	strs := []string{
		"error 0",
		"error 1",
		"error 2",
		"error 3",
		"error 4",
		"error 5",
		"error 6",
		"error 7",
	}
	for _, msg := range strs {
		err = Append(err, errors.New(msg))
	}

	merr, ok := err.(MultiError)
	is.True(ok)
	is.Equal(len(merr), len(strs))

	singleLine := fmt.Sprintf("%v", err)
	is.Equal(singleLine, strings.Join(strs, "; "))

	multiLine := fmt.Sprintf("%+v", err)
	is.Equal(multiLine, "the following errors occurred:\n -  "+strings.Join(strs, "\n -  "))
}
