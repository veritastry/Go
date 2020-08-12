package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/cloudflare/cfssl/log"
	"myblog/models"
	"time"
)

type RegisterController struct {
	beego.Controller
}

//func (this *RegisterController) Get() {
//	this.TplName = "register.html"
//}

// 处理注册
func (this *RegisterController) Post() {

	// 获取表单信息
	username := this.GetString("username")
	password := this.GetString("password")
	repassword := this.GetString("repassword")
	fmt.Println(username, password, repassword)
	log.Info(username, password, repassword)

	// 注册之前先判断该用户名是否已经被注册，如果已经注册，返回错误
	id := models.QueryUserWithUsername(username)
	fmt.Println("id", id)
	if id > 0 {
		this.Data["json"] = map[string]interface{}{
			"code":    0,
			"message": "用户名已经存在",
		}
		this.ServeJSON()
		return
	}

	// 注册用户名和密码
	// 存储的密码是md5后的数据，那么在登录的验证的时候，也是需要将用户的密码md5之后和数据库里面的密码进行判断
	//password = utils.MD5(password)
	//fmt.Println("md5后：", password)

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
