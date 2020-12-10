// Package log provides a log interface
package logger

import "os"

var (
	// Default logger
	DefaultLogger ILog
)

// Logger is a generic logging interface
type ILog interface {
	// Init initialises options
	Init(options ...Option) error
	// The Logger options
	Options() Options
	// Fields set fields to always be logged
	Fields(fields map[string]interface{}) ILog
	// Log writes a log entry
	Log(level Level, v ...interface{})
	// Logf writes a formatted log entry
	Logf(level Level, format string, v ...interface{})
	// String returns the name of logger
	String() string
}

func Init(opts ...Option) error {
	return DefaultLogger.Init(opts...)
}

func Fields(fields map[string]interface{}) ILog {
	return DefaultLogger.Fields(fields)
}

func Log(level Level, v ...interface{}) {
	DefaultLogger.Log(level, v...)
}

func Logf(level Level, format string, v ...interface{}) {
	DefaultLogger.Logf(level, format, v...)
}

func String() string {
	return DefaultLogger.String()
}

// Logger 通用log个性化实现
type Logger struct {
	ILog
}

// Info info级日志输出
func (l *Logger) Info(args ...interface{}) {
	l.Log(InfoLevel, args...)
}

// Infof info级日志输出
func (l *Logger) Infof(template string, args ...interface{}) {
	l.Logf(InfoLevel, template, args...)
}

// Trace trace级日志输出
func (l *Logger) Trace(args ...interface{}) {
	l.Log(InfoLevel, args...)
}

// Tracef trace级日志输出
func (l *Logger) Tracef(template string, args ...interface{}) {
	l.Logf(InfoLevel, template, args...)
}

// Debug debug级日志输出
func (l *Logger) Debug(args ...interface{}) {
	l.Log(InfoLevel, args...)
}

// Debugf debug级日志输出
func (l *Logger) Debugf(template string, args ...interface{}) {
	l.Logf(InfoLevel, template, args...)
}

// Warn warn级日志输出
func (l *Logger) Warn(args ...interface{}) {
	l.Log(InfoLevel, args...)
}

// Warnf warn级日志输出
func (l *Logger) Warnf(template string, args ...interface{}) {
	l.Logf(InfoLevel, template, args...)
}

// Error error级日志输出
func (l *Logger) Error(args ...interface{}) {
	l.Log(InfoLevel, args...)
}

// Errorf error级日志输出
func (l *Logger) Errorf(template string, args ...interface{}) {
	l.Logf(InfoLevel, template, args...)
}

// Fatal fatal级日志输出
func (l *Logger) Fatal(args ...interface{}) {
	l.Log(InfoLevel, args...)
	os.Exit(1)
}

// Fatalf fatal级日志输出
func (l *Logger) Fatalf(template string, args ...interface{}) {
	l.Logf(InfoLevel, template, args...)
	os.Exit(1)
}
