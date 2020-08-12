package controller

import (
	"ginDemo/bean"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

var (
	db *gorm.DB
)

// get url
func GetUserInfo(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		response := bean.ResponseData{
			Code:    50001,
			Status:  "error",
			Message: "参数错误",
			Data:    "",
		}
		c.JSON(200, response)
		return
	}
	user := bean.User{}
	db.First(&user, id)
	if user.Id > 0 {
		response := bean.ResponseData{
			Code:    200,
			Status:  "success",
			Message: "",
			Data:    user,
		}
		c.JSON(200, response)
	} else {
		response := bean.ResponseData{
			Code:    50001,
			Status:  "error",
			Message: "用户不存在",
			Data:    "",
		}
		c.JSON(200, response)
	}
}
