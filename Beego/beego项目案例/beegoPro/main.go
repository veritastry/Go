package main

import (
	_ "beegoPro/routers"
	"beegoPro/utils/db"
	"github.com/astaxie/beego"
)

func main() {

	db.Mysql()
	beego.Run()
}
