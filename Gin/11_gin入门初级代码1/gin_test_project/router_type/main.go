package main

import "github.com/gin-gonic/gin"

func main() {

	r := gin.Default()
	r.GET("/get", func(context *gin.Context) {
		context.String(200, "get")
	})
	r.POST("/post", func(context *gin.Context) {
		context.String(200, "post")
	})
	r.Handle("DELETE", "/delete", func(context *gin.Context) {
		context.String(200, "delete")
	})
	r.Any("/any", func(context *gin.Context) {
		context.String(200, "any")
	})
	r.Run()
}
