package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

const (
	getAllUser = "SELECT * FROM beego_user"
	deleteUser = "DELETE FROM beego_user WHERE uid = ?"
    selectUser = "SELECT * FROM beego_user WHERE uid = ?"
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

func Add(user *User) int {
	o := orm.NewOrm()
	num, _ := o.Insert(user)
	return int(num)
}

func Update(user *User) int {
	o := orm.NewOrm()
	num, _ := o.Update(user)
	return int(num)
}

func Find(uid int) User {
    user := User{}
    o := orm.NewOrm()
    _ = o.Raw(selectUser, uid).QueryRow(&user)
    return user
}

func init() {
	orm.RegisterModelWithPrefix("beego_", new(User))
}
