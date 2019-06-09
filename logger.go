package sugar

import (
	"go.uber.org/zap"
)

const (
	unitString = "unit"
)

// Logger 定制化 sugar
type Logger struct {
	zapSugar *zap.SugaredLogger
}

// NewLoggerOf logger with component field
func NewLoggerOf(component string) *Logger {
	return &Logger{
		zapSugar: zapSugar.With(unitString, component),
	}
}

// With adds a variadic number of fields to the logging context.
// see https://github.com/uber-go/zap/blob/v1.10.0/sugar.go#L91
func (l *Logger) With(args ...interface{}) {
	l.zapSugar = l.zapSugar.With(args...)
}

// Debug
func (l *Logger) Debug(args ...interface{}) {
	l.zapSugar.Debug(args...)
}

// Debugf package sugar of zap
func (l *Logger) Debugf(template string, args ...interface{}) {
	l.zapSugar.Debugf(template, args...)
}

// Info package sugar of zap
func (l *Logger) Info(args ...interface{}) {
	l.zapSugar.Info(args...)
}

// Infof package sugar of zap
func (l *Logger) Infof(template string, args ...interface{}) {
	l.zapSugar.Infof(template, args...)
}

// Warn package sugar of zap
func (l *Logger) Warn(args ...interface{}) {
	l.zapSugar.Warn(args...)
}

// Warnf package sugar of zap
func (l *Logger) Warnf(template string, args ...interface{}) {
	l.zapSugar.Warnf(template, args...)
}

// Error package sugar of zap
func (l *Logger) Error(args ...interface{}) {
	l.zapSugar.Warn(args...)
}

// Errorf package sugar of zap
func (l *Logger) Errorf(template string, args ...interface{}) {
	l.zapSugar.Errorf(template, args...)
}

// Fatal package sugar of zap
func (l *Logger) Fatal(args ...interface{}) {
	l.zapSugar.Warn(args...)
}

// Fatalf package sugar of zap
func (l *Logger) Fatalf(template string, args ...interface{}) {
	l.zapSugar.Fatalf(template, args...)
}

// Panic package sugar of zap
func (l *Logger) Panic(args ...interface{}) {
	l.zapSugar.Panic(args...)
}

// Panicf package sugar of zap
func (l *Logger) Panicf(template string, args ...interface{}) {
	l.zapSugar.Panicf(template, args...)
}
