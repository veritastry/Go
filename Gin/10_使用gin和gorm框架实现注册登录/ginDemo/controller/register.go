package controller

import (
	"ginDemo/bean"
	"github.com/gin-gonic/gin"
)

// post form
func Register(c *gin.Context) {
	tel := c.PostForm("tel")
	psd := c.PostForm("password")
	if len(tel) == 0 || len(psd) == 0 {
		response := bean.ResponseData{
			Code:    50001,
			Status:  "error",
			Message: "账号或密码不能为空",
			Data:    "",
		}
		c.JSON(200, response)
		return
	}
	var user bean.User
	db.Where("tel=?", tel).First(&user)
	if user.Tel == tel {
		response := bean.ResponseData{
			Code:    50001,
			Status:  "error",
			Message: "手机号已注册",
			Data:    "",
		}
		c.JSON(200, response)
	} else {
		newUser := bean.User{Tel: tel, Password: psd}
		db.Create(&newUser)
		response := bean.ResponseData{
			Code:    200,
			Status:  "success",
			Message: "注册成功",
			Data:    "",
		}
		c.JSON(200, response)
	}
}

