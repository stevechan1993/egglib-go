package filters

import (
	"github.com/astaxie/beego/context"
	"github.com/stevechan1993/egglib-go/log"
)

func CreateRequstLogFilter(logger log.Logger) func(ctx *context.Context) {
	return func(ctx *context.Context) {
		var append = make(map[string]interface{})
		append["framework"] = "beego"
		append["method"] = ctx.Input.Method()
		append["url"] = ctx.Input.URL()
		logger.Info("http请求", append)
		if ctx.Input.Is("GET") {
			logger.Debug("http请求", append)
		}
		if ctx.Input.Is("POST") {
			append["inputData"] = string(ctx.Input.GetData("requestBody").([]byte))
			logger.Debug("http请求", append)
		}
	}
}
