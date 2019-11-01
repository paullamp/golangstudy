package routers

import (
	"ihome/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/api/v1.0/areas", &controllers.AreaController{}, "get:GetArea")
	beego.Router("/api/v1.0/houses/index", &controllers.HousesIndexController{}, "get:GetHousesIndex")
	beego.Router("/api/v1.0/session", &controllers.SessionController{}, "get:GetSessionInfo;delete:DeleteSessionData")
	beego.Router("/api/v1.0/sessions", &controllers.SessionController{}, "post:Login")
	beego.Router("/api/v1.0/users", &controllers.UsersController{}, "post:Reg")
	// 上传用户头像
	beego.Router("/api/v1.0/user/avatar", &controllers.UsersController{}, "post:Postavatar")
	beego.Router("/api/hello", &controllers.TfastdfsControllerT{}, "get:Tfastdfs")
	beego.Router("/api/v1.0/user", &controllers.UsersController{}, "get:GetUserData")
	beego.Router("/api/v1.0/user/name", &controllers.UsersController{}, "put:UpdateName")
	// beego.Router("/api/v1.0/areas", &controllers.AreaController{}, "get:Jtestjson")
	beego.Router("/api/v1.0/user/auth", &controllers.UsersController{}, "get:GetAuth;post:PostAuth")
	beego.Router("/api/v1.0/user/houses", &controllers.HouseController{}, "get:GetHouseData")

	beego.Router("/api/v1.0/houses", &controllers.HouseController{}, "post:PostHouseData")
	beego.Router("/api/v1.0/houses/?:id", &controllers.HouseController{}, "get:GetSingleHouseData")
	beego.Router("/api/v1.0/user/orders", &controllers.OrderController{}, "get:GetUserOrder")
}
