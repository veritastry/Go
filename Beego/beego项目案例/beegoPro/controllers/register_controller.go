package controllers

import (
	"beegoPro/models"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/cloudflare/cfssl/log"
	"time"
)

type RegisterController struct {
	beego.Controller
}

// 注册
func (this *RegisterController) Post() {

	// 获取表单信息
	username := this.GetString("username")
	password := this.GetString("password")
	repassword := this.GetString("repassword")
	fmt.Println(username, password, repassword)
	log.Info(username, password, repassword)

	// 1. 判断用户名是否已经被注册
	id := models.QueryUserByUsername(username)
	if id > 0 {
		this.Data["json"] = map[string]interface{}{
			"code":    500001,
			"message": "用户名已经被注册",
		}
		this.ServeJSON()
		return
	}

	// 2. 注册用户名和密码
	//Md5Psw := utils.MD5(password)

	user := models.User{0, username, password, 0, time.Now().Unix()}

	_, err := models.InserUser(user)
	if err != nil {
		this.Data["json"] = map[string]interface{}{
			"code":    0,
			"message": "注册失败",
		}
	} else {
		this.Data["json"] = map[string]interface{}{
			"code":    1,
			"message": "注册成功",
		}
	}
	this.ServeJSON()
}
