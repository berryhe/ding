package ding

// Logger 日志接口
type Logger interface {
	Info(v ...interface{})

	Infof(template string, v ...interface{})

	Warn(v ...interface{})

	Warnf(template string, v ...interface{})

	Debug(v ...interface{})

	Debugf(template string, args ...interface{})

	Error(v ...interface{})

	Errorf(template string, v ...interface{})

	Fatal(v ...interface{})

	Fatalf(template string, v ...interface{})
}
