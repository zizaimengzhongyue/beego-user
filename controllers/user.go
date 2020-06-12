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

func (this *UserController) Add() {
	uid, _ := this.GetInt(":uid")
	name := this.GetString(":name")
	user := &models.User{
		Uid:  uid,
		Name: name,
	}
	num := models.Add(user)
	this.Return(0, "ok", num)
}

func (this *UserController) Update() {
	id, _ := this.GetInt(":id")
	uid, _ := this.GetInt(":uid")
	name := this.GetString(":name")
	user := &models.User{
		Id:   id,
		Uid:  uid,
		Name: name,
	}
	num := models.Update(user)
	this.Return(0, "ok", num)
}

func (this *UserController) Find() {
	uid, _ := this.GetInt(":uid")
	user := models.Find(uid)
	this.Return(0, "ok", user)
}

func (this *UserController) Return(status int, msg string, data interface{}) {
	type response struct {
		Status int         `json:"status"`
		Msg    string      `json:"msg"`
		Data   interface{} `json:"data"`
	}
	res := response{Status: status, Msg: msg, Data: data}
	bts, _ := json.Marshal(res)
	this.Ctx.WriteString(string(bts))
}
