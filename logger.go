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
func (l *Logger) With(args ...interface{}) *zap.SugaredLogger {
	return l.zapSugar.With(args...)
}

// Debug package sugar of zap
func (l *Logger) Debug(args ...interface{}) {
	l.zapSugar.Debug(args...)
}

// Debugf package sugar of zap
func (l *Logger) Debugf(template string, args ...interface{}) {
	l.zapSugar.Debugf(template, args...)
}

// Debugw package sugar of zap
func (l *Logger) Debugw(msg string, keysAndValues ...interface{}) {
	l.zapSugar.Debugw(msg, keysAndValues...)
}

// Info package sugar of zap
func (l *Logger) Info(args ...interface{}) {
	l.zapSugar.Info(args...)
}

// Infof package sugar of zap
func (l *Logger) Infof(template string, args ...interface{}) {
	l.zapSugar.Infof(template, args...)
}

// Infow package sugar of zap
func (l *Logger) Infow(msg string, keysAndValues ...interface{}) {
	l.zapSugar.Infow(msg, keysAndValues...)
}

// Warn package sugar of zap
func (l *Logger) Warn(args ...interface{}) {
	l.zapSugar.Warn(args...)
}

// Warnf package sugar of zap
func (l *Logger) Warnf(template string, args ...interface{}) {
	l.zapSugar.Warnf(template, args...)
}

// Warnw package sugar of zap
func (l *Logger) Warnw(msg string, keysAndValues ...interface{}) {
	l.zapSugar.Warnw(msg, keysAndValues...)
}

// Error package sugar of zap
func (l *Logger) Error(args ...interface{}) {
	l.zapSugar.Error(args...)
}

// Errorf package sugar of zap
func (l *Logger) Errorf(template string, args ...interface{}) {
	l.zapSugar.Errorf(template, args...)
}

// Errorw package sugar of zap
func (l *Logger) Errorw(msg string, keysAndValues ...interface{}) {
	l.zapSugar.Errorw(msg, keysAndValues...)
}

// Fatal package sugar of zap
func (l *Logger) Fatal(args ...interface{}) {
	l.zapSugar.Fatal(args...)
}

// Fatalf package sugar of zap
func (l *Logger) Fatalf(template string, args ...interface{}) {
	l.zapSugar.Fatalf(template, args...)
}

// Fatalw package sugar of zap
func (l *Logger) Fatalw(msg string, keysAndValues ...interface{}) {
	l.zapSugar.Fatalw(msg, keysAndValues...)
}

// Panic package sugar of zap
func (l *Logger) Panic(args ...interface{}) {
	l.zapSugar.Panic(args...)
}

// Panicf package sugar of zap
func (l *Logger) Panicf(template string, args ...interface{}) {
	l.zapSugar.Panicf(template, args...)
}

// Panicw package sugar of zap
func (l *Logger) Panicw(msg string, keysAndValues ...interface{}) {
	l.zapSugar.Panicw(msg, keysAndValues...)
}
