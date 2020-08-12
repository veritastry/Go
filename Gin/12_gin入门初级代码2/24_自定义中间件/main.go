package _4_自定义中间件

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func main() {

	router := gin.New()

	router.Use(Logger())
	router.GET("get", func(c *gin.Context) {

		example := c.MustGet("example").(string)
		c.JSON(http.StatusOK, gin.H{
			"example": example,
		})
	})
	router.Run(":9000")
}

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		// Set example variable
		c.Set("example", "12345")

		// before request

		c.Next()

		// after request
		latency := time.Since(t)
		log.Print(latency)

		// access the status we are sending
		status := c.Writer.Status()
		log.Println(status)
	}
}
