package logrus_ext

import (
	"fmt"

	"github.com/jxskiss/errors"
	"github.com/sirupsen/logrus"
)

const stacktraceKey = "stacktrace"

// NewStacktraceHook returns a new logrus hook which will check attached
// error with the entry, if the error implements `errors.StackTracker` interface,
// then a formatted stacktrace with be attached to the entry using a given key.
//
// Example:
//   logrus.AddHook(NewStacktraceHook())
//
// The stacktrace info key and levels can be customized.
//   hook := NewStacktraceHook()
//   hook.StacktraceKey = "stack"
//   hook.StackLevels = []logrus.Level{logrus.PanicLevel, logrus.FatalLevel}
//   logrus.AddHook(hook)
//
func NewStacktraceHook() *stacktraceHook {
	return &stacktraceHook{
		StacktraceKey: stacktraceKey,
		StackLevels: []logrus.Level{
			logrus.PanicLevel,
			logrus.FatalLevel,
			logrus.ErrorLevel,
			logrus.WarnLevel,
			logrus.InfoLevel,
			logrus.DebugLevel,
		},
	}
}

type stacktraceHook struct {
	StacktraceKey string
	StackLevels   []logrus.Level
}

func (hook *stacktraceHook) Levels() []logrus.Level {
	return hook.StackLevels
}

func (hook *stacktraceHook) Fire(entry *logrus.Entry) error {
	err, ok := entry.Data[logrus.ErrorKey].(error)
	if !ok || err == nil {
		return nil
	}
	stackTracker := errors.GetStackTracer(err)
	if stackTracker == nil {
		return nil
	}
	entry.Data[hook.StacktraceKey] = fmt.Sprintf("%+v", stackTracker.StackTrace())
	return nil
}
