package controllers

import (
	"encoding/json"

	"beego.baidu.com/models"
	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}

func (this *UserController) GetAll() {
	users := models.ReadAll()
	bts, err := json.Marshal(users)
	if err != nil {
	}
	this.Ctx.WriteString(string(bts))
}

func (this *UserController) Delete() {
	uid, _ := this.GetInt(":uid")
    num := models.Delete(uid)
	this.Return(0, "ok", num)
}

func (this *UserController) Return(status int, msg string, data interface{}) {
	type response struct {
		Status int         `json:"status"`
		Msg    string      `json:"msg"`
		Data   interface{} `json:"data"`
	}
	res := response{Status: status, Msg: msg, Data: data}
	bts, err := json.Marshal(res)
	if err != nil {
	}
	this.Ctx.WriteString(string(bts))
}
