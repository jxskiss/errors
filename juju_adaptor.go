package errors

import (
	"fmt"
)

// ==================== juju adaptor start ========================

// Trace is an alias of AddStack.
var Trace = AddStack

// Annotate adds a message and ensures there is a stack trace.
func Annotate(err error, message string) error {
	if err == nil {
		return nil
	}
	hasStack := HasStack(err)
	err = &withMessage{
		cause:         err,
		msg:           message,
		causeHasStack: hasStack,
	}
	if hasStack {
		return err
	}
	return &withStack{
		error: err,
		stack: callers(),
	}
}

// Annotatef adds a message and ensures there is a stack trace.
func Annotatef(err error, format string, args ...interface{}) error {
	if err == nil {
		return nil
	}
	hasStack := HasStack(err)
	err = &withMessage{
		cause:         err,
		msg:           fmt.Sprintf(format, args...),
		causeHasStack: hasStack,
	}
	if hasStack {
		return err
	}
	return &withStack{
		error: err,
		stack: callers(),
	}
}

// ErrorStack will format a stack trace if it is available, otherwise it will be Error()
// If the error is nil, the empty string is returned
// Note that this just calls fmt.Sprintf("%+v", err)
func ErrorStack(err error) string {
	if err == nil {
		return ""
	}
	return fmt.Sprintf("%+v", err)
}

// Error types.
const (
	unknown = iota
	timeout
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

func isErrorType(err error, etype int) bool {
	switch err := Cause(err).(type) {
	case *fundamental:
		return err.etype == etype
	}
	return false
}

// IsTimeout reports whether err was timeout error.
func IsTimeout(err error) bool {
	return isErrorType(err, timeout)
}

// Timeoutf represents an error with timeout message.
func Timeoutf(format string, args ...interface{}) error {
	return &fundamental{
		etype: timeout,
		msg:   fmt.Sprintf(format+" timeout", args...),
		stack: callers(),
	}
}

// IsBadRequest reports whether err was bad request error.
func IsBadRequest(err error) bool {
	return isErrorType(err, badRequest)
}

// BadRequestf represents an error with bad request message.
func BadRequestf(format string, args ...interface{}) error {
	return &fundamental{
		etype: badRequest,
		msg:   fmt.Sprintf(format+" bad request", args...),
		stack: callers(),
	}
}

// IsNotFound reports whether err was not found error.
func IsNotFound(err error) bool {
	return isErrorType(err, notFound)
}

// NotFoundf represents an error with not found message.
func NotFoundf(format string, args ...interface{}) error {
	return &fundamental{
		etype: notFound,
		msg:   fmt.Sprintf(format+" not found", args...),
		stack: callers(),
	}
}

// IsUserNotFound reports whether err was not found error.
func IsUserNotFound(err error) bool {
	return isErrorType(err, userNotFound)
}

// UserNotFoundf represents an error with user not found message.
func UserNotFoundf(format string, args ...interface{}) error {
	return &fundamental{
		etype: userNotFound,
		msg:   fmt.Sprintf(format+" user not found", args...),
		stack: callers(),
	}
}

// IsNotSupported reports whether err was not supported error.
func IsNotSupported(err error) bool {
	return isErrorType(err, notSupported)
}

// NotSupportedf represents an error with not supported message.
func NotSupportedf(format string, args ...interface{}) error {
	return &fundamental{
		etype: notSupported,
		msg:   fmt.Sprintf(format+" not supported", args...),
		stack: callers(),
	}
}

// IsNotValid reports whether err was not valid error.
func IsNotValid(err error) bool {
	return isErrorType(err, notValid)
}

// NotValidf represents an error with not valid message.
func NotValidf(format string, args ...interface{}) error {
	return &fundamental{
		etype: notValid,
		msg:   fmt.Sprintf(format+" not valid", args...),
		stack: callers(),
	}
}

// IsAlreadyExists reports whether err was already exists error.
func IsAlreadyExists(err error) bool {
	return isErrorType(err, alreadyExists)
}

// AlreadyExistsf represents an error with already exists message.
func AlreadyExistsf(format string, args ...interface{}) error {
	return &fundamental{
		etype: alreadyExists,
		msg:   fmt.Sprintf(format+" already exists", args...),
		stack: callers(),
	}
}

// IsUnauthorized reports whether err was unauthorized error.
func IsUnauthorized(err error) bool {
	return isErrorType(err, unauthorized)
}

// Unauthorizedf represents an error with unauthorized message.
func Unauthorizedf(format string, args ...interface{}) error {
	return &fundamental{
		etype: unauthorized,
		msg:   fmt.Sprintf(format+" unauthorized", args...),
		stack: callers(),
	}
}

// IsForbidden reports whether err was forbidden error.
func IsForbidden(err error) bool {
	return isErrorType(err, forbidden)
}

// Forbiddenf represents an error with forbidden message.
func Forbiddenf(format string, args ...interface{}) error {
	return &fundamental{
		etype: forbidden,
		msg:   fmt.Sprintf(format+" forbidden", args...),
		stack: callers(),
	}
}

// IsNotImplemented reports whether err was not implemented error.
func IsNotImplemented(err error) bool {
	return isErrorType(err, notImplemented)
}

// NotImplementedf represents an error with not implemented message.
func NotImplementedf(format string, args ...interface{}) error {
	return &fundamental{
		etype: notImplemented,
		msg:   fmt.Sprintf(format+" not implemented", args...),
		stack: callers(),
	}
}

// IsNotProvisioned reports whether err was not provisioned error.
func IsNotProvisioned(err error) bool {
	return isErrorType(err, notProvisioned)
}

// NotProvisionedf represents an error with not provisioned message.
func NotProvisionedf(format string, args ...interface{}) error {
	return &fundamental{
		etype: notProvisioned,
		msg:   fmt.Sprintf(format+" not provisioned", args...),
		stack: callers(),
	}
}

// IsNotAssigned reports whether err was not assigned error.
func IsNotAssigned(err error) bool {
	return isErrorType(err, notAssigned)
}

// NotAssignedf represents an error with not assigned message.
func NotAssignedf(format string, args ...interface{}) error {
	return &fundamental{
		etype: notAssigned,
		msg:   fmt.Sprintf(format+" not assigned", args...),
		stack: callers(),
	}
}

// IsMethodNotAllowed reports whether err was method not allowed error.
func IsMethodNotAllowed(err error) bool {
	return isErrorType(err, methodNotAllowed)
}

// MethodNotAllowedf represents an error with method not allowed message.
func MethodNotAllowedf(format string, args ...interface{}) error {
	return &fundamental{
		etype: methodNotAllowed,
		msg:   fmt.Sprintf(format+" method not allowed", args...),
		stack: callers(),
	}
}

// ==================== juju adaptor end ========================
