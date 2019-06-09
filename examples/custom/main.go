package main

import (
	"github.com/chalvern/sugar"
	"go.uber.org/zap"
)

func main() {
	config := zap.NewProductionConfig()
	config.OutputPaths = []string{"./production.log"}
	config.ErrorOutputPaths = []string{"./production_err.log"}
	sugar.SetSugar(&config)

	sugar.Info("main info")

	myCustomLogger()
}

func myCustomLogger() {
	myLogger := sugar.NewLoggerOf("my_custom_logger")
	myLogger.Info("log of myCustomLogger info")
	myLogger.Warn("log of myCustomLogger warn")
}
