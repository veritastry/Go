package _7_绑定POST参数

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Login struct {
	Name     string `form:"name" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func main() {

	router := gin.Default()

	router.POST("login", func(c *gin.Context) {

		var login Login

		err := c.ShouldBind(&login)
		if err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}
		if login.Name != "zzh" || login.Password != "123" {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":   "failed",
				"message":  "账号或者密码有误",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"status":   "succeed",
			"message":  "登录成功",
			"name":     login.Name,
			"password": login.Password,
		})
	})
	router.Run()
}
