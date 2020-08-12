package main

import "github.com/gin-gonic/gin"

func main() {

	r := gin.Default()
	r.LoadHTMLGlob("template/*")
	r.GET("/index", func(context *gin.Context) {
		context.HTML(200, "index.html", gin.H{
			"title": "index.html",
		})
	})
	r.Run()
}
