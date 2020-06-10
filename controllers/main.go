package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	this.Data["Website"] = "beego.baidu.com"
	this.Data["Email"] = "test@baidu.com"
	this.TplName = "index.tpl"
}
