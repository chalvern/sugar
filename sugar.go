package sugar

// Error package mainLogger
func Error(args ...interface{}) {
	mainLogger.Error(args...)
}

// Debug package mainLogger
func Debug(args ...interface{}) {
	mainLogger.Debug(args...)
}

// Debugf package mainLogger
func Debugf(template string, args ...interface{}) {
	mainLogger.Debugf(template, args...)
}

// Info package mainLogger
func Info(args ...interface{}) {
	mainLogger.Info(args...)
}

// Infof package mainLogger
func Infof(template string, args ...interface{}) {
	mainLogger.Infof(template, args...)
}

// Warn package mainLogger
func Warn(args ...interface{}) {
	mainLogger.Warn(args...)
}

// Warnf package mainLogger
func Warnf(template string, args ...interface{}) {
	mainLogger.Warnf(template, args...)
}

// Errorf package mainLogger
func Errorf(template string, args ...interface{}) {
	mainLogger.Errorf(template, args...)
}

// Fatal package mainLogger
func Fatal(args ...interface{}) {
	mainLogger.Fatal(args...)
}

// Fatalf package mainLogger
func Fatalf(template string, args ...interface{}) {
	mainLogger.Fatalf(template, args...)
}

// Panic package mainLogger
func Panic(args ...interface{}) {
	mainLogger.Panic(args...)
}

// Panicf package mainLogger
func Panicf(template string, args ...interface{}) {
	mainLogger.Panicf(template, args...)
}
