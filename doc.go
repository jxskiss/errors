// Package errors is a drop-in replacement of the "errors" package
// in the std library, with additional functions to add and retrieve
// stack information to an error.
//
// Functions New, Is, As, Unwrap, Join are simply aliases of the same
// functions in the "errors" package.
//
// Function WrapNew is similar to New, but it also adds stack frames and optional details.
// Function Wrap wraps an error with stack frames and optional details.
// Function Errorf formats a string error using fmt.Errorf and wraps it with stack frames.
// The frames can be retrieved by function Frames.
// For simple use-case, user may use function Stacktrace to get a formatted stacktrace.
// Function Details gets wrapped details from an error if available.
package errors
