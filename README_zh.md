# sugar

《[English README](./README.md)》

Sugar 封装了 [zap](https://github.com/uber-go/zap) 日志库，让开发者能够更方便、快捷地使用这个插件。
本库只是简单的封装，如果使用过程中遇到问题，推荐阅读 [zap](https://github.com/uber-go/zap) 的相关文档寻找答案。

当然，欢迎大家提 issue 来一起完善这个仓库。


## 一些注意点

* Sugar 使用 zap 仓库的 `zap.NewDevelopmentConfig()` 方法返回的配置作为默认配置（开发环境），这个配置使用 console 类型的输出格式（平铺的那种，区别于 json 类型的日志）
* 我曾经的工作阅历，认为服务容器化会是未来的趋势，因此觉得开发环境也有必要配置 json 格式的输出，通过 `InitDevelopmentSugar()` 可以达到目的。这样大家写容器化的服务时可以直接使用这个库而不需要定制日志路径。
* 生产环境默认就是 json 格式，主要是为了方便地把日志收集到 ELK 中去。


## 方法

Sugar 封装了 sugared-zap 的所有方法，如下（如果有疏漏，可以提 issue 通知我）：

* Debug/Debugf/Debugw
* Info/Infof/Infow
* Warn/Warnf/Warnw
* Error/Errorf/Errorw
* Fatal/Fatalf/Fatalw
* Panic/Panicf/Panicw


## logs 格式预先看

我之所以喜欢使用 zap，主要因为它漂亮的输出格式（根据官方的描述，它的输出效率也很高，不过我还没有遇到这个瓶颈）

### 默认的开发模式

这种模式的输出是 console 样式，就是说日志平铺展示（开发模式也可以配置 json 格式的输出，见 [一些注意点](#一些注意点)。

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

生产模式默认是 json 格式的输出，从而方便把日志收集起来集中处理。

```json
{"level":"info","ts":1560129183.672966,"caller":"default/main.go:13","msg":"default production sugar of chalvern","unit":"main"}
{"level":"info","ts":1560129183.673015,"caller":"default/main.go:20","msg":"log of myCustomLogger info","unit":"my_custom_logger"}
{"level":"warn","ts":1560129183.6730392,"caller":"default/main.go:21","msg":"log of myCustomLogger warn","unit":"my_custom_logger"}
{"level":"info","ts":1560129183.673059,"caller":"default/main.go:26","msg":"log of myCustomLogger2 info","unit":"my_custom_logger_2"}
{"level":"warn","ts":1560129183.673182,"caller":"default/main.go:27","msg":"log of myCustomLogger2 warn","unit":"my_custom_logger_2"}
```

## 例子

下面的例子也可以在 [examples](./examples) 目录找到。


### 默认配置
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

### 定制化配置

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