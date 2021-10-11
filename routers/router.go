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
	web.Get("/", func(ctx *context.Context) {
		ctx.Redirect(http.StatusFound, "/admin/index/index")
	})

	admin := web.NewNamespace("/api",
		//UEditor控制器
		//web.NSRouter("/user", &controllers.UserController{}, "get:All"),
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
	web.AddNamespace(admin)
	web.InsertFilter("*", beego.BeforeExec, filters.JsonParser)
}
