# errors

Package errors provides handful error handling primitives.

This package is forked from [pingcap/errors](https://github.com/pingcap/errors) which is another derivative of the popular [pkg/errors](https://github.com/pkg/errors) package.

This errors package is different from pkg/errors or pingcap/errors in following ways:

1. A new method `AddStack` is added to avoid the overhead of adding duplicate stacks to the error chain by calling `WithStack`. Generally, `AddStack` should be used instead of `WithStack`.

2. Some helper functions to help migration from [juju/errors](https://github.com/juju/errors) with signature compatibility. This package use different implementation with pingcap/errors.

3. A new error type `withFields` is added to pass context information of the error like logrus. Extra key-value context information can be attached to the error by calling functions `WithFields`, `New`, `AddStack`, `WithStack`, `WithMessage` or `Wrap`. The attached key-value information can be printed by using `fmt.Sprint("%+v", err)`. Also the additional package [logrus_ext](./logrus_ext) can be used to automatically hook the context information when using with logrus.

4. Group errors and multi errors handling primitives are added.

5. Unlike pingcap/errors, the `StackTrace` and `Frame` types in this package are just aliases of the corresponding types from pkg/errors, which makes this package compatible with existing logrus hooks such as gelf hook and sentry hook.

6. **NOTE**: this package does not follow the versioning of either pkg/errors or pingcap/errors.

For users who don't need the additional features from this package, the original [pkg/errors](http://github.com/pkg/errors) package is highly recommended.

For latest information about packages pkg/errors and pingcap/errors, please refer to their websites.

----

[![Travis-CI](https://travis-ci.org/pkg/errors.svg)](https://travis-ci.org/pkg/errors) [![AppVeyor](https://ci.appveyor.com/api/projects/status/b98mptawhudj53ep/branch/master?svg=true)](https://ci.appveyor.com/project/davecheney/errors/branch/master) [![GoDoc](https://godoc.org/github.com/pkg/errors?status.svg)](http://godoc.org/github.com/pkg/errors) [![Report card](https://goreportcard.com/badge/github.com/pkg/errors)](https://goreportcard.com/report/github.com/pkg/errors) [![Sourcegraph](https://sourcegraph.com/github.com/pkg/errors/-/badge.svg)](https://sourcegraph.com/github.com/pkg/errors?badge)

Package errors provides simple error handling primitives.

`go get github.com/pkg/errors`

The traditional error handling idiom in Go is roughly akin to
```go
if err != nil {
        return err
}
```
which applied recursively up the call stack results in error reports without context or debugging information. The errors package allows programmers to add context to the failure path in their code in a way that does not destroy the original value of the error.

## Adding context to an error

The errors.Wrap function returns a new error that adds context to the original error. For example
```go
_, err := ioutil.ReadAll(r)
if err != nil {
        return errors.Wrap(err, "read failed")
}
```
## Retrieving the cause of an error

Using `errors.Wrap` constructs a stack of errors, adding context to the preceding error. Depending on the nature of the error it may be necessary to reverse the operation of errors.Wrap to retrieve the original error for inspection. Any error value which implements this interface can be inspected by `errors.Cause`.
```go
type causer interface {
        Cause() error
}
```
`errors.Cause` will recursively retrieve the topmost error which does not implement `causer`, which is assumed to be the original cause. For example:
```go
switch err := errors.Cause(err).(type) {
case *MyError:
        // handle specifically
default:
        // unknown error
}
```

[Read the package documentation for more information](https://godoc.org/github.com/pkg/errors).

## Roadmap

With the upcoming [Go2 error proposals](https://go.googlesource.com/proposal/+/master/design/go2draft.md) this package is moving into maintenance mode. The roadmap for a 1.0 release is as follows:

- 0.9. Remove pre Go 1.9 support, address outstanding pull requests (if possible)
- 1.0. Final release.

## Contributing

Because of the Go2 errors changes, this package is not accepting proposals for new functionality. With that said, we welcome pull requests, bug fixes and issue reports. 

Before sending a PR, please discuss your change by raising an issue.

## License

BSD-2-Clause
