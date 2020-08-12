package main

import "github.com/gin-gonic/gin"

func main() {

	r := gin.Default()

	r.GET("/user/*action", func(context *gin.Context) {
		context.String(200, "get")
	})
	r.Run()
}
