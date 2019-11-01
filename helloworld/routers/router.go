package routers

import (
	"helloworld/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/user", &controllers.UserController{}) //新增加/user处理的路由
	// beego.Router("/user/?:id", &controllers.UserController{}, "get:GetInfo")
	beego.Router("/user/:id([0-9]+)", &controllers.UserController{}, "get:GetNum")
	beego.Router("/download/*.*", &controllers.UserController{}, "get:GetFilename")
	beego.Router("/down2/*", &controllers.UserController{}, "get:GetAllinfo;post:PostData")
}
