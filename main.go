package main

import (
	_ "github.com/zizaimengzhongyue/beego-user/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.SetStaticPath("/static", "static")
	beego.Run(":8087")
}
