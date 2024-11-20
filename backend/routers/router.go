package routers

import (
	"backend/controllers"
	"github.com/beego/beego/v2/server/web"
)

func init() {

	web.SetStaticPath("", "static")
	web.Router("/*", &controllers.ReactController{}, "get:ServeReact")


	web.Router("/api/users/:id", &controllers.UserController{})
}
