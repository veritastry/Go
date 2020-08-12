package _3_重定向

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	router := gin.Default()

	router.GET("/redict", func(c *gin.Context) {
		// 重定向到百度的页面
		c.Redirect(http.StatusMovedPermanently, "http://www.baidu.com")
	})

	// 重定向到test2请求的页面
	router.GET("/test", func(c *gin.Context) {
		c.Request.URL.Path = "/test2"
		router.HandleContext(c)
	})
	router.GET("/test2", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "helloWorld"})
	})

	router.Run(":9000")
}
