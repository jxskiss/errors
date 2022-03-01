package errors

import (
	"fmt"
)

type withType struct {
	etype int
	fundamental
}

// Error types.
const (
	timeout = iota
	badRequest
	notFound
	userNotFound
	notSupported
	notValid
	alreadyExists
	unauthorized
	forbidden
	notImplemented
	notProvisioned
	notAssigned
	methodNotAllowed
	quotaLimitExceeded
	notYetAvailable
)

func newTypedError(etype int, format string, args ...interface{}) error {
	return &withType{
		etype: etype,
		fundamental: fundamental{
			msg:   fmt.Sprintf(format, args...),
			stack: callersSkip(5),
		},
	}
}

func isErrorType(err error, etype int) bool {
	switch err := Cause(err).(type) {
	case *withType:
		return err.etype == etype
	}
	return false
}

// ==================== juju adaptor start ========================

// Trace is an alias of AddStack.
var Trace = AddStack

// Annotate adds a message and ensures there is a stack trace.
var Annotate = Wrap

// Annotatef adds a message and ensures there is a stack trace.
var Annotatef = Wrapf

// ErrorStack will format a stack trace if it is available, otherwise it will be Error()
// If the error is nil, the empty string is returned
// Note that this just calls fmt.Sprintf("%+v", err)
func ErrorStack(err error) string {
	if err == nil {
		return ""
	}
	return fmt.Sprintf("%+v", err)
}

// IsTimeout reports whether err is a "timeout" error.
func IsTimeout(err error) bool {
	return isErrorType(err, timeout)
}

// Timeoutf returns a typed error with " timeout" suffix.
func Timeoutf(format string, args ...interface{}) error {
	format += " timeout"
	return newTypedError(timeout, format, args...)
}

// IsBadRequest reports whether err is a "bad request" error.
func IsBadRequest(err error) bool {
	return isErrorType(err, badRequest)
}

// BadRequestf returns a typed error with " bad request" suffix.
func BadRequestf(format string, args ...interface{}) error {
	format += " bad request"
	return newTypedError(badRequest, format, args...)
}

// IsNotFound reports whether err is a "not found" error.
func IsNotFound(err error) bool {
	return isErrorType(err, notFound)
}

// NotFoundf returns a typed error with " not found" suffix.
func NotFoundf(format string, args ...interface{}) error {
	format += " not found"
	return newTypedError(notFound, format, args...)
}

// IsUserNotFound reports whether err is a "user not found" error.
func IsUserNotFound(err error) bool {
	return isErrorType(err, userNotFound)
}

// UserNotFoundf returns a typed error with " user not found" suffix.
func UserNotFoundf(format string, args ...interface{}) error {
	format += " user not found"
	return newTypedError(userNotFound, format, args...)
}

// IsNotSupported reports whether err is a "not supported" error.
func IsNotSupported(err error) bool {
	return isErrorType(err, notSupported)
}

// NotSupportedf returns a typed error with " not supported" suffix.
func NotSupportedf(format string, args ...interface{}) error {
	format += " not supported"
	return newTypedError(notSupported, format, args...)
}

// IsNotValid reports whether err is a "not valid" error.
func IsNotValid(err error) bool {
	return isErrorType(err, notValid)
}

// NotValidf returns a typed error with " not valid" suffix.
func NotValidf(format string, args ...interface{}) error {
	format += " not valid"
	return newTypedError(notValid, format, args...)
}

// IsAlreadyExists reports whether err is an "already exists" error.
func IsAlreadyExists(err error) bool {
	return isErrorType(err, alreadyExists)
}

// AlreadyExistsf returns a typed error with " already exists" suffix.
func AlreadyExistsf(format string, args ...interface{}) error {
	format += " already exists"
	return newTypedError(alreadyExists, format, args...)
}

// IsUnauthorized reports whether err is an "unauthorized" error.
func IsUnauthorized(err error) bool {
	return isErrorType(err, unauthorized)
}

// Unauthorizedf returns a typed error with " unauthorized" suffix.
func Unauthorizedf(format string, args ...interface{}) error {
	format += " unauthorized"
	return newTypedError(unauthorized, format, args...)
}

// IsForbidden reports whether err is a "forbidden" error.
func IsForbidden(err error) bool {
	return isErrorType(err, forbidden)
}

// Forbiddenf returns a typed error with " forbidden" suffix.
func Forbiddenf(format string, args ...interface{}) error {
	format += " forbidden"
	return newTypedError(forbidden, format, args...)
}

// IsNotImplemented reports whether err is a "not implemented" error.
func IsNotImplemented(err error) bool {
	return isErrorType(err, notImplemented)
}

// NotImplementedf returns a typed error with " not implemented" suffix.
func NotImplementedf(format string, args ...interface{}) error {
	format += " not implemented"
	return newTypedError(notImplemented, format, args...)
}

// IsNotProvisioned reports whether err is a "not provisioned" error.
func IsNotProvisioned(err error) bool {
	return isErrorType(err, notProvisioned)
}

// NotProvisionedf returns a typed error with " not provisioned" suffix.
func NotProvisionedf(format string, args ...interface{}) error {
	format += " not provisioned"
	return newTypedError(notProvisioned, format, args...)
}

// IsNotAssigned reports whether err is a "not assigned" error.
func IsNotAssigned(err error) bool {
	return isErrorType(err, notAssigned)
}

// NotAssignedf returns a typed error with " not assigned" suffix.
func NotAssignedf(format string, args ...interface{}) error {
	format += " not assigned"
	return newTypedError(notAssigned, format, args...)
}

// IsMethodNotAllowed reports whether err is a "method not allowed" error.
func IsMethodNotAllowed(err error) bool {
	return isErrorType(err, methodNotAllowed)
}

// MethodNotAllowedf returns a typed error with " method not allowed" suffix.
func MethodNotAllowedf(format string, args ...interface{}) error {
	format += " method not allowed"
	return newTypedError(methodNotAllowed, format, args...)
}

// IsQuotaLimitExceeded reports whether err is a "quota limit exceeded" error.
func IsQuotaLimitExceeded(err error) bool {
	return isErrorType(err, quotaLimitExceeded)
}

// QuotaLimitExceededf returns a typed error with " quota limit exceeded" suffix.
func QuotaLimitExceededf(format string, args ...interface{}) error {
	format += " quota limit exceeded"
	return newTypedError(quotaLimitExceeded, format, args...)
}

// IsNotYetAvailable reports whether err is a "not yet available" error.
func IsNotYetAvailable(err error) bool {
	return isErrorType(err, notYetAvailable)
}

// NotYetAvailablef returns a typed error with " not yet available" suffix.
func NotYetAvailablef(format string, args ...interface{}) error {
	format += " not yet available"
	return newTypedError(notYetAvailable, format, args...)
}

// ==================== juju adaptor end ========================
