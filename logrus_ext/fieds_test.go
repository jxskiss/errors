package logrus_ext

import (
	"bytes"
	stderr "errors"
	"io/ioutil"
	"sync"
	"testing"

	"github.com/jxskiss/errors"
	"github.com/sirupsen/logrus"
)

func Test_WithFields(t *testing.T) {
	var b bytes.Buffer
	var logger = logrus.New()
	logger.Out = &b
	logger.AddHook(NewErrFieldsHook("err_"))

	err1 := stderr.New("dummy std error")
	err1 = errors.WithFields(err1, errors.F{"key1": "values1"})
	err1 = errors.WithFields(err1, logrus.Fields{
		"key2": "value2",
		"key3": "value3",
	})
	logger.WithError(err1).Info("test std error")
	logrus.Println(string(b.Bytes())) // stderr

	b.Reset()
	err2 := errors.New("dummy utils error")
	err2 = errors.WithFields(err2, errors.F{"key1": "value1"})
	err2 = errors.WithFields(err2, logrus.Fields{
		"key2": "value2",
		"key3": "value3",
	})
	logger.WithError(err2).Info("test utils error")
	logrus.Println(string(b.Bytes())) // stderr
}

func Test_Concurrent(t *testing.T) {
	var logger = logrus.New()
	logger.Out = ioutil.Discard
	logger.AddHook(NewErrFieldsHook("err_"))
	logger.AddHook(NewStacktraceHook())

	err1 := errors.New("dummy utils error")
	err1 = errors.WithFields(err1, errors.F{"key1": "value1"})
	err1 = errors.WithFields(err1, logrus.Fields{
		"key2": "value2",
		"key3": "value3",
	})

	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			logger.WithError(err1).Infof("test concurrent %d", i)
		}(i)
	}
	wg.Wait()
}

func Benchmark_Raw(b *testing.B) {
	logger := logrus.New()
	logger.Out = ioutil.Discard

	err := errForBenchmark()
	for i := 0; i < b.N; i++ {
		logger.WithError(err).Info("benchmark")
	}
}

func Benchmark_WithFields(b *testing.B) {
	logger := logrus.New()
	logger.Out = ioutil.Discard
	logger.AddHook(NewErrFieldsHook("err_"))

	err := errForBenchmark()
	for i := 0; i < b.N; i++ {
		logger.WithError(err).Info("benchmark")
	}
}

func Benchmark_Stacktrace(b *testing.B) {
	logger := logrus.New()
	logger.Out = ioutil.Discard
	logger.AddHook(NewStacktraceHook())

	err := errForBenchmark()
	for i := 0; i < b.N; i++ {
		logger.WithError(err).Info("benchmark")
	}
}

func Benchmark_WithFields_Stacktrace(b *testing.B) {
	logger := logrus.New()
	logger.Out = ioutil.Discard
	logger.AddHook(NewErrFieldsHook("err_"))
	logger.AddHook(NewStacktraceHook())

	err := errForBenchmark()
	for i := 0; i < b.N; i++ {
		logger.WithError(err).Info("benchmark")
	}
}

func errForBenchmark() error {
	err := errors.New("dummy error for benchmark")
	err = errors.WithFields(err, logrus.Fields{
		"key1": "value1",
		"key2": "value2",
		"key3": "value3",
	})
	return err
}
