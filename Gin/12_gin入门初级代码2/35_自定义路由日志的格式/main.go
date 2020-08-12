package _5_自定义路由日志的格式

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {

	router := gin.Default()

	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
		log.Printf("endpoint %v %v %v %v\n", httpMethod, absolutePath, handlerName, nuHandlers)
	}

	router.POST("/post", func(c *gin.Context) {
		c.JSON(http.StatusOK, "post")
	})

	router.GET("/get", func(c *gin.Context) {
		c.JSON(http.StatusOK, "get")
	})

	router.GET("/status", func(c *gin.Context) {
		c.JSON(http.StatusOK, "ok")
	})
	router.Run(":7979")

}
