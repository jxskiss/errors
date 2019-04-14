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

// IsTimeout reports whether err was timeout error.
func IsTimeout(err error) bool {
	return isErrorType(err, timeout)
}

// Timeoutf represents an error with timeout message.
func Timeoutf(format string, args ...interface{}) error {
	format += " timeout"
	return newTypedError(timeout, format, args...)
}

// IsBadRequest reports whether err was bad request error.
func IsBadRequest(err error) bool {
	return isErrorType(err, badRequest)
}

// BadRequestf represents an error with bad request message.
func BadRequestf(format string, args ...interface{}) error {
	format += " bad request"
	return newTypedError(badRequest, format, args...)
}

// IsNotFound reports whether err was not found error.
func IsNotFound(err error) bool {
	return isErrorType(err, notFound)
}

// NotFoundf represents an error with not found message.
func NotFoundf(format string, args ...interface{}) error {
	format += " not found"
	return newTypedError(notFound, format, args...)
}

// IsUserNotFound reports whether err was not found error.
func IsUserNotFound(err error) bool {
	return isErrorType(err, userNotFound)
}

// UserNotFoundf represents an error with user not found message.
func UserNotFoundf(format string, args ...interface{}) error {
	format += " user not found"
	return newTypedError(userNotFound, format, args...)
}

// IsNotSupported reports whether err was not supported error.
func IsNotSupported(err error) bool {
	return isErrorType(err, notSupported)
}

// NotSupportedf represents an error with not supported message.
func NotSupportedf(format string, args ...interface{}) error {
	format += " not supported"
	return newTypedError(notSupported, format, args...)
}

// IsNotValid reports whether err was not valid error.
func IsNotValid(err error) bool {
	return isErrorType(err, notValid)
}

// NotValidf represents an error with not valid message.
func NotValidf(format string, args ...interface{}) error {
	format += " not valid"
	return newTypedError(notValid, format, args...)
}

// IsAlreadyExists reports whether err was already exists error.
func IsAlreadyExists(err error) bool {
	return isErrorType(err, alreadyExists)
}

// AlreadyExistsf represents an error with already exists message.
func AlreadyExistsf(format string, args ...interface{}) error {
	format += " already exists"
	return newTypedError(alreadyExists, format, args...)
}

// IsUnauthorized reports whether err was unauthorized error.
func IsUnauthorized(err error) bool {
	return isErrorType(err, unauthorized)
}

// Unauthorizedf represents an error with unauthorized message.
func Unauthorizedf(format string, args ...interface{}) error {
	format += " unauthorized"
	return newTypedError(unauthorized, format, args...)
}

// IsForbidden reports whether err was forbidden error.
func IsForbidden(err error) bool {
	return isErrorType(err, forbidden)
}

// Forbiddenf represents an error with forbidden message.
func Forbiddenf(format string, args ...interface{}) error {
	format += " forbidden"
	return newTypedError(forbidden, format, args...)
}

// IsNotImplemented reports whether err was not implemented error.
func IsNotImplemented(err error) bool {
	return isErrorType(err, notImplemented)
}

// NotImplementedf represents an error with not implemented message.
func NotImplementedf(format string, args ...interface{}) error {
	format += " not implemented"
	return newTypedError(notImplemented, format, args...)
}

// IsNotProvisioned reports whether err was not provisioned error.
func IsNotProvisioned(err error) bool {
	return isErrorType(err, notProvisioned)
}

// NotProvisionedf represents an error with not provisioned message.
func NotProvisionedf(format string, args ...interface{}) error {
	format += " not provisioned"
	return newTypedError(notProvisioned, format, args...)
}

// IsNotAssigned reports whether err was not assigned error.
func IsNotAssigned(err error) bool {
	return isErrorType(err, notAssigned)
}

// NotAssignedf represents an error with not assigned message.
func NotAssignedf(format string, args ...interface{}) error {
	format += " not assigned"
	return newTypedError(notAssigned, format, args...)
}

// IsMethodNotAllowed reports whether err was method not allowed error.
func IsMethodNotAllowed(err error) bool {
	return isErrorType(err, methodNotAllowed)
}

// MethodNotAllowedf represents an error with method not allowed message.
func MethodNotAllowedf(format string, args ...interface{}) error {
	format += " method not allowed"
	return newTypedError(methodNotAllowed, format, args...)
}

// ==================== juju adaptor end ========================
