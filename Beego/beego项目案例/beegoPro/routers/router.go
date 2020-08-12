package routers

import (
	"beegoPro/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/register", &controllers.RegisterController{})
}
