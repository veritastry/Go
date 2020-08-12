package controller

import (
	"ginDemo/bean"
	"github.com/gin-gonic/gin"
)

// post form
func Login(c *gin.Context) {
	// 获取表单数据 tel，password
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
	} else {
		var user bean.User
		db.Where("tel=?", tel).First(&user)
		if user.Tel == "" {
			response := bean.ResponseData{
				Code:    50001,
				Status:  "error",
				Message: "用户不存在",
				Data:    "",
			}
			c.JSON(200, response)
		} else {
			if user.Password == psd {
				response := bean.ResponseData{
					Code:    200,
					Status:  "success",
					Message: "登录成功",
					Data:    "",
				}
				c.JSON(200, response)
			} else {
				response := bean.ResponseData{
					Code:    5001,
					Status:  "error",
					Message: "密码错误",
					Data:    "",
				}
				c.JSON(200, response)
			}
		}
	}
}
