# sugar

《[中文文档](./README_zh.md)》

more simple logger which package sugared [zap](https://github.com/uber-go/zap).
Please learn [zap](https://github.com/uber-go/zap) first if you want advanced using.


## something you should know

* sugar use `zap.NewDevelopmentConfig()` as default config, which use console as encoding style,
* `InitDevelopmentSugar()` set encoding being json instead of console,  which is practically better for docker's json logger plugin.


## Leveled method

all methods of zap:

* Debug/Debugf
* Info/Infof
* Warn/Warnf
* Error/Errorf
* Fatal/Fatalf
* Panic/Panicf

## logs looking
I recommend zap mostly because of its beautiful logs looking.

### looking #1
Development mode with console style printing.

```bash
2019-06-10T09:13:03.672+0800    DEBUG   default/main.go:7       default development sugar of chalvern   {"unit": "main"}
2019-06-10T09:13:03.672+0800    INFO    default/main.go:20      log of myCustomLogger info      {"unit": "my_custom_logger"}
2019-06-10T09:13:03.672+0800    WARN    default/main.go:21      log of myCustomLogger warn      {"unit": "my_custom_logger"}
github.com/chalvern/sugar.(*Logger).Warn
        /Users/Chalvern/developer/golang/src/github.com/chalvern/sugar/logger.go:51
main.myCustomLogger
        /Users/Chalvern/developer/golang/src/github.com/chalvern/sugar/examples/default/main.go:21
main.main
        /Users/Chalvern/developer/golang/src/github.com/chalvern/sugar/examples/default/main.go:8
runtime.main
        /usr/local/go/src/runtime/proc.go:200
2019-06-10T09:13:03.672+0800    INFO    default/main.go:26      log of myCustomLogger2 info     {"unit": "my_custom_logger_2"}
2019-06-10T09:13:03.672+0800    WARN    default/main.go:27      log of myCustomLogger2 warn     {"unit": "my_custom_logger_2"}
github.com/chalvern/sugar.(*Logger).Warn
        /Users/Chalvern/developer/golang/src/github.com/chalvern/sugar/logger.go:51
main.myCustomLogger2
        /Users/Chalvern/developer/golang/src/github.com/chalvern/sugar/examples/default/main.go:27
main.main
        /Users/Chalvern/developer/golang/src/github.com/chalvern/sugar/examples/default/main.go:9
runtime.main
        /usr/local/go/src/runtime/proc.go:200
```

### looking #2

Production mode with json style printing (Development mode can also has json stype printing, more see [something you should know](#something-you-should-know) )

```json
{"level":"info","ts":1560129183.672966,"caller":"default/main.go:13","msg":"default production sugar of chalvern","unit":"main"}
{"level":"info","ts":1560129183.673015,"caller":"default/main.go:20","msg":"log of myCustomLogger info","unit":"my_custom_logger"}
{"level":"warn","ts":1560129183.6730392,"caller":"default/main.go:21","msg":"log of myCustomLogger warn","unit":"my_custom_logger"}
{"level":"info","ts":1560129183.673059,"caller":"default/main.go:26","msg":"log of myCustomLogger2 info","unit":"my_custom_logger_2"}
{"level":"warn","ts":1560129183.673182,"caller":"default/main.go:27","msg":"log of myCustomLogger2 warn","unit":"my_custom_logger_2"}
```

## example

you can find example in [examples](./examples).


### default sugar
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

### custom sugar

```go
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

```