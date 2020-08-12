package _1_模型绑定和验证

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Login struct {
	User     string `form:"user" json:"user" xml:"user"  binding:"required"`
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
}

func main() {

	router := gin.Default()

	// json
	router.POST("post/json", func(c *gin.Context) {

		var loginJson Login
		err := c.ShouldBindJSON(&loginJson)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		if loginJson.User != "zzh" || loginJson.Password != "123" {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": "failed",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "succeed"})

	})

	// xml
	router.POST("post/xml", func(c *gin.Context) {

		var loginXml Login
		err := c.ShouldBindXML(loginXml)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		if loginXml.User != "zzh" || loginXml.Password != "123" {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": "failed",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"status": "succeed",
		})
	})

	// form
	router.POST("post/form", func(c *gin.Context) {

		var form Login
		err := c.ShouldBind(&form)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"err": err.Error(),
			})
			return
		}

		if form.User != "zzh" || form.Password != "123" {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": "failed",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "succeed"})

	})

	router.Run()
}
