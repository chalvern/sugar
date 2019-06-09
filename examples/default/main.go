package main

import "github.com/chalvern/sugar"

func main() {

	sugar.Debug("default development sugar of chalvern")
	myCustomLogger()
	myCustomLogger2()

	sugar.InitProductionSugar()
	sugar.Debug("should not be printed")
	sugar.Info("default production sugar of chalvern")
	myCustomLogger()
	myCustomLogger2()
}

func myCustomLogger() {
	myLogger := sugar.NewLoggerOf("my_custom_logger")
	myLogger.Info("log of myCustomLogger info")
	myLogger.Warn("log of myCustomLogger warn")
}

func myCustomLogger2() {
	myLogger := sugar.NewLoggerOf("my_custom_logger_2")
	myLogger.Info("log of myCustomLogger2 info")
	myLogger.Warn("log of myCustomLogger2 warn")
}
