package ding

// Logger 日志接口
type Logger interface {
	Info(v ...interface{})

	Infof(template string, args ...interface{})

	Warn(args ...interface{})

	Warnf(template string, args ...interface{})

	Debug(args ...interface{})

	Debugf(template string, args ...interface{})

	Error(args ...interface{})

	Errorf(template string, args ...interface{})

	Fatal(args ...interface{})

	Fatalf(template string, args ...interface{})
}
