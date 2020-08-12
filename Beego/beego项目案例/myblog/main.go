package main

import (
	"github.com/astaxie/beego"
	_ "myblog/routers"
	"myblog/utils"
)

func main() {

	utils.InitMysql()
	beego.Run()
}
