package controllers

type MainController struct {
    BaseController
}

func (this *MainController) Get() {
	this.Data["Website"] = "beego.baidu.com"
	this.Data["Email"] = "test@baidu.com"
	this.TplName = "index.tpl"
}
