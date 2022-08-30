package sugar

// Debug package mainLogger
func Debug(args ...interface{}) {
	mainLogger.zapSugar.Debug(args...)
}

// Debugf package mainLogger
func Debugf(template string, args ...interface{}) {
	mainLogger.zapSugar.Debugf(template, args...)
}

// Debugw package mainLogger
func Debugw(msg string, keysAndValues ...interface{}) {
	mainLogger.zapSugar.Debugw(msg, keysAndValues...)
}

// Info package mainLogger
func Info(args ...interface{}) {
	mainLogger.zapSugar.Info(args...)
}

// Infof package mainLogger
func Infof(template string, args ...interface{}) {
	mainLogger.zapSugar.Infof(template, args...)
}

// Infow package mainLogger
func Infow(msg string, keysAndValues ...interface{}) {
	mainLogger.zapSugar.Infow(msg, keysAndValues...)
}

// Warn package mainLogger
func Warn(args ...interface{}) {
	mainLogger.zapSugar.Warn(args...)
}

// Warnf package mainLogger
func Warnf(template string, args ...interface{}) {
	mainLogger.zapSugar.Warnf(template, args...)
}

// Warnw package mainLogger
func Warnw(msg string, keysAndValues ...interface{}) {
	mainLogger.zapSugar.Warnw(msg, keysAndValues...)
}

// Error package mainLogger
func Error(args ...interface{}) {
	mainLogger.zapSugar.Error(args...)
}

// Errorf package mainLogger
func Errorf(template string, args ...interface{}) {
	mainLogger.zapSugar.Errorf(template, args...)
}

// Errorw package mainLogger
func Errorw(msg string, keysAndValues ...interface{}) {
	mainLogger.zapSugar.Errorw(msg, keysAndValues...)
}

// Fatal package mainLogger
func Fatal(args ...interface{}) {
	mainLogger.zapSugar.Fatal(args...)
}

// Fatalf package mainLogger
func Fatalf(template string, args ...interface{}) {
	mainLogger.zapSugar.Fatalf(template, args...)
}

// Fatalw package mainLogger
func Fatalw(msg string, keysAndValues ...interface{}) {
	mainLogger.zapSugar.Fatalw(msg, keysAndValues...)
}

// Panic package mainLogger
func Panic(args ...interface{}) {
	mainLogger.zapSugar.Panic(args...)
}

// Panicf package mainLogger
func Panicf(template string, args ...interface{}) {
	mainLogger.zapSugar.Panicf(template, args...)
}

// Panicw package mainLogger
func Panicw(msg string, keysAndValues ...interface{}) {
	mainLogger.zapSugar.Panicw(msg, keysAndValues...)
}
