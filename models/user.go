package models

import (
	"fmt"
	"github.com/astaxie/beego/logs"
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

func Delete(uid int) (int, error) {
	o := orm.NewOrm()
	result, err := o.Raw(deleteUser, uid).Exec()
	if err != nil {
		logs.Warn(err)
		return 0, err
	}
	num, _ := result.RowsAffected()
	return int(num), nil
}

func Add(user *User) (int, error) {
	o := orm.NewOrm()
	num, err := o.Insert(user)
	if err != nil {
		logs.Warn(err)
		return 0, err
	}
	return int(num), nil
}

func Update(user *User) (int, error) {
	o := orm.NewOrm()
	num, err := o.Update(user)
	if err != nil {
		logs.Warn(err)
		return 0, err
	}
	return int(num), nil
}

func Find(uid int) (User, error) {
	user := User{}
	o := orm.NewOrm()
	err := o.Raw(selectUser, uid).QueryRow(&user)
	if err != nil {
		logs.Warn(err)
		return user, err
	}
	return user, nil
}

func init() {
	orm.RegisterModelWithPrefix("beego_", new(User))
}
