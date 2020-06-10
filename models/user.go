package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

const (
	getAllUser = "SELECT * FROM beego_user"
	deleteUser = "DELETE FROM beego_user WHERE uid = ?"
)

type User struct {
	Id   int
	Uid  int
	Name string
}

func ReadAll() []User {
	users := []User{}
	o := orm.NewOrm()
	num, err := o.Raw(getAllUser).QueryRows(&users)
	if err != nil {
		fmt.Println(num, err)
	}
	return users
}

func Delete(uid int) int {
	o := orm.NewOrm()
	result, _ := o.Raw(deleteUser, uid).Exec()
	num, _ := result.RowsAffected()
	return int(num)
}
