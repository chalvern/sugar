package main

import "github.com/chalvern/sugar"

func main() {

	sugar.Debug("default development sugar of chalvern")
	sugar.InitProductionSugar()
	sugar.Debug("should not be printed")
	sugar.Info("default production sugar of chalvern")

	myCustomLogger()
}

func myCustomLogger() {
	myLogger := sugar.NewLoggerOf("my_custom_logger")
	myLogger.Info("log of myCustomLogger")
	myLogger.Warn("log of myCustomLogger")
}
