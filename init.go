package sugar

import (
	"go.uber.org/zap"
)

var (
	zapLogger *zap.Logger
	zapSugar  *zap.SugaredLogger
	// mainLogger default logger, with unit filed valued "main"
	mainLogger *Logger
)

// GetZapLogger get zap's logger for special using
func GetZapLogger() *zap.Logger {
	return zapLogger
}

func init() {
	config := zap.NewDevelopmentConfig()
	config.OutputPaths = []string{"stderr"}
	config.ErrorOutputPaths = []string{"stderr"}
	SetSugar(&config)
}

// InitProductionSugar for production
func InitProductionSugar(opts ...zap.Option) {
	config := zap.NewProductionConfig()
	config.OutputPaths = []string{"stderr"}
	config.ErrorOutputPaths = []string{"stderr"}
	SetSugar(&config, opts...)
}

// InitDevelopmentSugar for development
// encoding being json instead of console,
// which is practically better for docker's json logger,
func InitDevelopmentSugar(opts ...zap.Option) {
	config := zap.NewDevelopmentConfig()
	config.Encoding = "json"
	config.OutputPaths = []string{"stderr"}
	config.ErrorOutputPaths = []string{"stderr"}
	SetSugar(&config, opts...)
}

// SetSugar set sugar's logger
func SetSugar(config *zap.Config, opts ...zap.Option) {
	opts = append(opts, zap.AddCallerSkip(1))
	var err error
	zapLogger, err = config.Build(opts...)
	if err != nil {
		panic(err)
	}
	zapSugar = zapLogger.Sugar()
	mainLogger = NewLoggerOf("main")
}
