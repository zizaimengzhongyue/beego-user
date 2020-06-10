package routers

import (
	"beego.baidu.com/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/user/getAllUser", &controllers.UserController{}, "*:GetAll")
	beego.Router("/user/delete/:uid", &controllers.UserController{}, "*:Delete")
}
