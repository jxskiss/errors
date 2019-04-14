package errors

import "testing"

var jujuAdaptorTestcases = []struct {
	maker   func(format string, args ...interface{}) error
	checker func(err error) bool
}{
	{Timeoutf, IsTimeout},
	{BadRequestf, IsBadRequest},
	{NotFoundf, IsNotFound},
	{UserNotFoundf, IsUserNotFound},
	{NotSupportedf, IsNotSupported},
	{NotValidf, IsNotValid},
	{AlreadyExistsf, IsAlreadyExists},
	{Unauthorizedf, IsUnauthorized},
	{Forbiddenf, IsForbidden},
	{NotImplementedf, IsNotImplemented},
	{NotProvisionedf, IsNotProvisioned},
	{NotAssignedf, IsNotAssigned},
	{MethodNotAllowedf, IsMethodNotAllowed},
}

func TestJujuAdaptor(t *testing.T) {
	for _, c := range jujuAdaptorTestcases {
		err := c.maker("test error %v", "something")
		if !c.checker(err) {
			t.Errorf("failed check error: %v", err)
		}
	}
}
