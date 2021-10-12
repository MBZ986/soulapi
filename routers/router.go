package routers

import (
	beego "github.com/beego/beego/v2/adapter"
	"github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
	"net/http"
	"soulapi/controllers"
	"soulapi/filters"
)

func init() {
	//过滤器
	web.InsertFilter("*", beego.BeforeExec, filters.JsonParser)
	//简单路由
	web.Get("/", func(ctx *context.Context) {
		ctx.Redirect(http.StatusFound, "/admin/index/index")
	})
	//组路由(核心路由规则)
	apiRouter := web.NewNamespace("/api",
		//UEditor控制器
		web.NSNamespace("/user",
			web.NSRouter("/all", &controllers.UserController{}, "get:All"),
			web.NSRouter("/add", &controllers.UserController{}, "post:Add"),
			web.NSRouter("/del", &controllers.UserController{}, "delete:Del"),
			web.NSRouter("/upd", &controllers.UserController{}, "put:Upd"),
			web.NSRouter("/get", &controllers.UserController{}, "get:Get"),
			web.NSRouter("/page", &controllers.UserController{}, "get:Page"),
			web.NSRouter("/count", &controllers.UserController{}, "get:Count"),
			web.NSRouter("/find", &controllers.UserController{}, "post,get:Find"),
			web.NSRouter("/insertTitle", &controllers.UserController{}, "put:InsertTitle"),
		),
		web.NSNamespace("/title",
			web.NSRouter("/all", &controllers.TitleController{}, "get:All"),
			web.NSRouter("/add", &controllers.TitleController{}, "post:Add"),
			web.NSRouter("/del", &controllers.TitleController{}, "delete:Del"),
			web.NSRouter("/upd", &controllers.TitleController{}, "put:Upd"),
			web.NSRouter("/get", &controllers.TitleController{}, "get:Get"),
			web.NSRouter("/page", &controllers.TitleController{}, "get:Page"),
			web.NSRouter("/count", &controllers.TitleController{}, "get:Count"),
			web.NSRouter("/find", &controllers.TitleController{}, "post,get:Find"),
			web.NSRouter("/getUsersById", &controllers.TitleController{}, "get:GetUsersById"),
		),
	)
	web.AddNamespace(apiRouter)
}
