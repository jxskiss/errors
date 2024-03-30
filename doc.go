// Package errors is a drop-in replacement of the "errors" package
// in the std library, with additional functions to add and retrieve
// stack information to an error.
//
// Functions New, Is, As, Unwrap, Join are simply aliases of the same
// functions in the "errors" package.
//
// Function Wrap wraps and error with stack frames, function Errorf
// formats a string error using fmt.Errorf and wraps it with stack frames.
// The frames can be retrieved by function GetFrames.
// For simple use-case, user may use function GetStacktrace to get
// a formatted stacktrace.
package errors
