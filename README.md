# sugar
more simple logger which package sugared [zap](https://github.com/uber-go/zap).


## something you should know

* sugar use `zap.NewDevelopmentConfig()` as default config, which use console as encoding style,
* `InitDevelopmentSugar()` set encoding being json instead of console,  which is practically better for docker's json logger plugin.


## leveled Method

all methods of zap:

* Debug/Debugf
* Info/Infof
* Warn/Warnf
* Error/Errorf
* Fatal/Fatalf
* Panic/Panicf

## example

you can find example in [examples](./examples).

```go
// cat ./examples/default/main.go
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
```