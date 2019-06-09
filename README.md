# sugar
more simple logger packaged sugared [zap](https://github.com/uber-go/zap).


## something you should know

* sugar use `zap.NewDevelopmentConfig()` as default config, which use console as encoding style,
* `InitDevelopmentSugar()` set encoding being json instead of console,  which is practically better for docker's json logger plugin.
