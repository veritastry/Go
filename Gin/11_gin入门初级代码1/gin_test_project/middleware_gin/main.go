package main

import (
	"github.com/gin-gonic/gin"
	"io"
	"os"
)

func main() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)
	gin.DefaultErrorWriter = io.MultiWriter(f)
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())
	r.GET("/get", func(context *gin.Context) {
		name := context.DefaultQuery("name", "DefaultName")
		panic("test panic")
		context.String(200, "%s", name)
	})
	r.Run()
}
