package _1_HTML渲染

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	router := gin.Default()

	// 表示templates目录下的所有文件
	router.LoadHTMLGlob("templates/*")
	// 表示具体要写入哪个文件
	//router.LoadHTMLFiles("templates/index.html","templages/index2.yaml")

	// 获取templates目录下的文件
	router.GET("get/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"status": "succeed",
			"title":  "index.html",
		})
	})

	// 获取同一个目录下，templates/users目录下的文件
	router.GET("get/users/usersIndex", func(c *gin.Context) {
		c.HTML(http.StatusOK, "users/usersIndex.html", gin.H{
			"status": "succeed",
			"title":  "usersIndex.html",
		})
	})

	router.Run(":7000")
}
