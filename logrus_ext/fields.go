package logrus_ext

import (
	"github.com/jxskiss/errors"
	"github.com/sirupsen/logrus"
)

// NewErrFieldsHook returns a new logrus hook which will check attached
// error with the entry, if the error or it's cause error has attached fields,
// (created by methods in the errors package), then the extra fields
// attached with the error object will be appended to the log entry.
func NewErrFieldsHook(keyPrefix string) *errFieldsHook {
	return &errFieldsHook{prefix: keyPrefix}
}

type errFieldsHook struct {
	prefix string
}

func (hook *errFieldsHook) Levels() []logrus.Level {
	return []logrus.Level{
		logrus.PanicLevel,
		logrus.FatalLevel,
		logrus.ErrorLevel,
		logrus.WarnLevel,
		logrus.InfoLevel,
		logrus.DebugLevel,
	}
}

func (hook *errFieldsHook) Fire(entry *logrus.Entry) error {
	err, ok := entry.Data[logrus.ErrorKey].(error)
	if !ok || err == nil {
		return nil
	}
	fields := errors.Fields(err)
	if hook.prefix != "" {
		for k, v := range fields {
			entry.Data[hook.prefix+k] = v
		}
	} else {
		for k, v := range fields {
			entry.Data[k] = v
		}
	}
	return nil
}
