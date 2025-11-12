package log

import "bosh-admin/global"

// Debug debug日志
func Debug(args ...interface{}) {
	global.Logger.Debug(args...)
}

// Info info日志
func Info(args ...interface{}) {
	global.Logger.Info(args...)
}

// Warn warn日志
func Warn(args ...interface{}) {
	global.Logger.Warn(args...)
}

// Error error日志
func Error(args ...interface{}) {
	global.Logger.Error(args...)
}

// Panic panic日志
func Panic(args ...interface{}) {
	global.Logger.Panic(args...)
}

// Fatal fatal日志
func Fatal(args ...interface{}) {
	global.Logger.Fatal(args...)
}

// Debugf debug日志
func Debugf(template string, args ...interface{}) {
	global.Logger.Debugf(template, args...)
}

// Infof info日志
func Infof(template string, args ...interface{}) {
	global.Logger.Infof(template, args...)
}

// Warnf warn日志
func Warnf(template string, args ...interface{}) {
	global.Logger.Warnf(template, args...)
}

// Errorf error日志
func Errorf(template string, args ...interface{}) {
	global.Logger.Errorf(template, args...)
}

// Panicf panic日志
func Panicf(template string, args ...interface{}) {
	global.Logger.Panicf(template, args...)
}

// Fatalf fatal日志
func Fatalf(template string, args ...interface{}) {
	global.Logger.Fatalf(template, args...)
}
