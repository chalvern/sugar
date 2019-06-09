package sugar

import (
	"go.uber.org/zap"
)

var zapSugar *zap.SugaredLogger

// mainLogger default logger, with unit filed valued "main"
var mainLogger *Logger

func init() {
	InitDevelopmentSugar()
}

// InitProductionSugar for production
func InitProductionSugar(opts ...zap.Option) {
	config := zap.NewProductionConfig()
	initSugarWith(&config, opts...)
}

// InitDevelopmentSugar for development
func InitDevelopmentSugar(opts ...zap.Option) {
	// set sugar for production
	config := zap.NewDevelopmentConfig()
	initSugarWith(&config, opts...)
}

// initSugarWith init Sugar
func initSugarWith(config *zap.Config, opts ...zap.Option) {
	config.OutputPaths = []string{"stderr"}
	config.ErrorOutputPaths = []string{"stderr"}

	opts = append(opts, zap.AddCallerSkip(1))
	zapLogger, err := config.Build(opts...)
	if err != nil {
		panic(err)
	}
	SetSugar(zapLogger.Sugar())
}

// SetSugar set sugar's logger
func SetSugar(zs *zap.SugaredLogger) {
	zapSugar = zs
	mainLogger = NewLoggerOf("main")
}
