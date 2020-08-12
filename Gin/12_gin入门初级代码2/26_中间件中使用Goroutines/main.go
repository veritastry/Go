package _6_中间件中使用Goroutines

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func main() {

	router := gin.Default()

	// 不使用goroutine
	router.GET("/get", func(c *gin.Context) {
		time.Sleep(5 * time.Second)
		log.Println("Done! in path " + c.Request.URL.Path)
	})

	// 使用goroutine
	router.GET("/get/goroutine", func(c *gin.Context) {
		cCp := c.Copy()
		go func() {
			time.Sleep(5 * time.Second)
			log.Println("Done! in path " + cCp.Request.URL.Path)
		}()
	})

	router.Run(":5000")
}
