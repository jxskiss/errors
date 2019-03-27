package logrus_ext

import (
	"github.com/jxskiss/errors"
	"github.com/sirupsen/logrus"
)

type withFields struct {
	error
	logrus.Fields
}

func (w *withFields) Cause() error { return w.error }

// Fields returns attached fields with the given error if available, else nil.
func Fields(err error) logrus.Fields {
	var data logrus.Fields
	errors.WalkDeep(err, func(err error) bool {
		if fieldsErr, ok := err.(*withFields); ok && fieldsErr != nil {
			data = make(logrus.Fields, len(fieldsErr.Fields))
			for k, v := range fieldsErr.Fields {
				data[k] = v
			}
			return true
		}
		return false
	})
	return data
}

// WithField add a field to the given error. If you want multiple fields, use `WithFields`.
func WithField(err error, key string, value interface{}) *withFields {
	return WithFields(err, logrus.Fields{key: value})
}

// WithFields add multiple fields to the given error.
func WithFields(err error, fields ...logrus.Fields) *withFields {
	var length int
	for _, fs := range fields {
		length += len(fs)
	}
	fieldsErr, ok := err.(*withFields)
	if ok {
		length += len(fieldsErr.Fields)
		err = fieldsErr.error
	}
	data := make(logrus.Fields, length)
	if ok {
		for k, v := range fieldsErr.Fields {
			data[k] = v
		}
	}
	for _, fs := range fields {
		for k, v := range fs {
			data[k] = v
		}
	}
	return &withFields{
		error:  err,
		Fields: data,
	}
}

// NewErrFieldsHook returns a new logrus hook which will check attached
// error with the entry, if the error or it's cause error is of type *withFields
// (created by `WithField` and `WithFields` methods), then the extra fields
// attached with the error object will be appended to the log entry.
func NewErrFieldsHook() *errFieldsHook {
	return &errFieldsHook{}
}

type errFieldsHook struct {
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
	var data logrus.Fields
	errors.WalkDeep(err, func(err error) bool {
		if fieldsErr, ok := err.(*withFields); ok && fieldsErr != nil {
			data = fieldsErr.Fields
			return true
		}
		return false
	})
	for k, v := range data {
		// add "err_" prefix to reduce key conflict
		entry.Data["err_"+k] = v
	}
	return nil
}
