package helloWord

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	router := gin.Default()

	router.GET("get", func(c *gin.Context) {

		// 第一种表达
		//c.String(http.StatusOK, "successful")

		// 第二种表达
		c.JSON(http.StatusOK, gin.H{
			"message": "succeed",
		})
	})

	router.Run(":5000")

}
