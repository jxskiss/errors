# errors

Package errors is a drop-in replacement of the "errors" package
in the std library, with additional functions to add and retrieve
stack frames or extra details with an error.

This package is inspired by the famous package "github.com/pkg/errors".

Functions New, Is, As, Unwrap, Join are alias functions to package "errors".

Function Wrap wraps an error with stack frames.

Function WrapNew is similar to New, but it also adds stack frames.

Function Errorf formats a string error using fmt.Errorf and wraps it with stack frames,
the frames can be retrieved by function Frames.
For simple use-case, user may use function Stacktrace to get
a formatted stacktrace.

Extra details attached to an error can be retrieved by function Details.

Functions Cause, AddStack, WithStack, WithMessage, WithMessagef are deprecated,
they are kept not removed to help migrating from the v1 package.

`go get github.com/jxskiss/errors/v2`

[Read the package documentation for more information](https://godoc.org/github.com/jxkiss/errors/v2).

## License

BSD-2-Clause
