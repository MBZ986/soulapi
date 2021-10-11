package filters

import (
	"github.com/beego/beego/v2/server/web/context"
)

var JsonParser = func(ctx *context.Context) {
	//fmt.Println("过滤器")
	//if strings.Contains(ctx.Input.Header("content-type"), "application/json") {
	//	if json, err := simplejson.NewJson(ctx.Input.RequestBody); err == nil {
	//		ctx.Input.SetData("json", json)
	//		ctx.Input.RunController = controllers.TitleController
	//	} else {
	//		global.Logger.Debug("解析请求JSON参数失败：%v", err)
	//	}
	//}
}
