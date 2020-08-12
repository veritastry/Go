package main

import "github.com/gin-gonic/gin"

type Person struct {
	Name    string `form:"name" binding:"required"`
	Age     int    `form:"age" binding:"required,gt=10"`  // 验证规则：https://godoc.org/gopkg.in/go-playground/validator.v8
	Address string `form:"address" binding:"required"`
}

func main() {

	r := gin.Default()

	r.GET("/testing", func(context *gin.Context) {
		var p Person
		if err := context.ShouldBind(&p); err != nil {
			context.String(500, "%v", err)
		}
		context.String(200, "%v", p)
	})
	r.Run()
}
