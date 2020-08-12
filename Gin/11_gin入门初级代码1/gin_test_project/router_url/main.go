package main

import "github.com/gin-gonic/gin"

func main() {

	r := gin.Default()

	r.GET("/:name/:age", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"name": context.Param("name"),
			"age":  context.Param("age"),
		})
	})
	r.Run("9090")
}
