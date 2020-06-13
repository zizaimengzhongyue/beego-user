package controllers

import (
	"encoding/json"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type BaseController struct {
	beego.Controller
}

func (this *BaseController) Return(status int, msg string, data interface{}) {
	type response struct {
		Status int         `json:"status"`
		Msg    string      `json:"msg"`
		Data   interface{} `json:"data"`
	}
	var (
		res response
		bts []byte
	)

	res = response{Status: status, Msg: msg, Data: data}
	bts, _ = json.Marshal(res)
	this.Ctx.WriteString(string(bts))
}

func (this *BaseController) Error(err error) {
	logs.Warn(err)
	this.Return(1, "error", "发生了一些错误，我们在紧急处理")
}
