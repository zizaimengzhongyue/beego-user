package controllers

import (
	"encoding/json"

	"github.com/zizaimengzhongyue/beego-user/models"
)

type UserController struct {
    BaseController
}

func (this *UserController) GetAll() {
	var (
		users []models.User
		bts   []byte
		err   error
	)

	users = models.ReadAll()
	bts, err = json.Marshal(users)
	if err != nil {
		goto end
	}

end:
	if err != nil {
		this.Error(err)
	} else {
		this.Return(0, "ok", bts)
	}
}

func (this *UserController) Delete() {
	var (
		uid int
		num int
		err error
	)
	uid, err = this.GetInt(":uid")
	if err != nil {
		goto end
	}
	num, err = models.Delete(uid)
	if err != nil {
		goto end
	}
	this.Return(0, "ok", num)

end:
	if err != nil {
		this.Error(err)
	} else {
		this.Return(0, "ok", num)
	}
}

func (this *UserController) Add() {
	var (
		uid  int
		name string
		user *models.User
		num  int
		err  error
	)

	uid, err = this.GetInt(":uid")
	if err != nil {
		goto end
	}
	name = this.GetString(":name")
	user = &models.User{
		Uid:  uid,
		Name: name,
	}
	num, err = models.Add(user)
	if err != nil {
		goto end
	}

end:
	if err != nil {
		this.Error(err)
	} else {
		this.Return(0, "ok", num)
	}
}

func (this *UserController) Update() {
	var (
		id   int
		uid  int
		name string
		user *models.User
		num  int
		err  error
	)

	id, err = this.GetInt(":id")
	if err != nil {
		goto end
	}
	uid, err = this.GetInt(":uid")
	if err != nil {
		goto end
	}
	name = this.GetString(":name")
	user = &models.User{
		Id:   id,
		Uid:  uid,
		Name: name,
	}
	num, err = models.Update(user)
	if err != nil {
		goto end
	}
	this.Return(0, "ok", num)

end:
	if err != nil {
		this.Error(err)
	} else {
		this.Return(0, "ok", user)
	}
}

func (this *UserController) Find() {
	var (
		user models.User
		uid  int
		err  error
	)

	uid, err = this.GetInt(":uid")
	if err != nil {
		goto end
	}
	user, err = models.Find(uid)
	if err != nil {
		goto end
	}

end:
	if err != nil {
		this.Error(err)
	} else {
		this.Return(0, "ok", user)
	}

}
