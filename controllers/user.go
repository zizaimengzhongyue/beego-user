package controllers

import (
	"encoding/json"

	"beego.baidu.com/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type UserController struct {
	beego.Controller
}

func (this *UserController) GetAll() {
	users := models.ReadAll()
	bts, err := json.Marshal(users)
	if err != nil {
		this.Error(err)
		return
	}
	this.Return(0, "ok", bts)
}

func (this *UserController) Delete() {
	uid, err := this.GetInt(":uid")
	if err != nil {
		this.Error(err)
		return
	}
	num, err := models.Delete(uid)
	if err != nil {
		this.Error(err)
		return
	}
	this.Return(0, "ok", num)
}

func (this *UserController) Add() {
	uid, err := this.GetInt(":uid")
	if err != nil {
		this.Error(err)
		return
	}
	name := this.GetString(":name")
	user := &models.User{
		Uid:  uid,
		Name: name,
	}
	num, err := models.Add(user)
	if err != nil {
		this.Error(err)
		return
	}
	this.Return(0, "ok", num)
}

func (this *UserController) Update() {
	id, err := this.GetInt(":id")
	if err != nil {
		this.Error(err)
		return
	}
	uid, err := this.GetInt(":uid")
	if err != nil {
		this.Error(err)
		return
	}
	name := this.GetString(":name")
	user := &models.User{
		Id:   id,
		Uid:  uid,
		Name: name,
	}
	num, err := models.Update(user)
	if err != nil {
		this.Error(err)
		return
	}
	this.Return(0, "ok", num)
}

func (this *UserController) Find() {
	uid, err := this.GetInt(":uid")
	if err != nil {
		this.Error(err)
		return
	}
	user, err := models.Find(uid)
	if err != nil {
		this.Error(err)
		return
	}
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

func (this *UserController) Error(err error) {
	logs.Warn(err)
	this.Return(1, "error", "发生了一些错误，我们在紧急处理")
}
