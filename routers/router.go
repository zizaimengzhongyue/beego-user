package routers

import (
	"beego.baidu.com/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/user/getAllUser", &controllers.UserController{}, "*:GetAll")
	beego.Router("/user/delete/:uid", &controllers.UserController{}, "*:Delete")
	beego.Router("/user/add/:uid/:name", &controllers.UserController{}, "*:Add")
    beego.Router("/user/update/:id/:uid/:name", &controllers.UserController{}, "*:Update")
    beego.Router("/user/find/:uid", &controllers.UserController{}, "*:Find")
}
