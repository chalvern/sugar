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
}
