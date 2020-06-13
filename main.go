package main

import (
	_ "beego.baidu.com/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.Debug = true
	orm.RegisterDataBase("default", "mysql", "root:123456@tcp(127.0.0.1:3306)/beego?charset=utf8", 30)
	logs.SetLogger("file", `{"filename": "log/beego_user.log"}`)
}

func main() {
	beego.SetStaticPath("/static", "static")
	beego.Run(":8087")
}
