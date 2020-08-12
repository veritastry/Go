package routers

import (
	"github.com/astaxie/beego"
	"myblog/controllers"
)

func init() {
	// 注册
	beego.Router("/register", &controllers.RegisterController{})
}
