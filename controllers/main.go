package controllers

type MainController struct {
    BaseController
}

func (this *MainController) Get() {
	this.Data["Website"] = "github.com/zizaimengzhongyue/beego-user"
	this.Data["Email"] = "luyanfeng1992@qq.com"
	this.TplName = "index.tpl"
}
