package main

import (
	"github.com/gin-gonic/gin"
	"time"
)

type Person struct {
	Name     string    `form:"name"`
	Age      int       `form:"age"`
	Birthday time.Time `form:"birthday" time_format:"2006-01-02"`
}

func testing(context *gin.Context) {
	var p Person
	err := context.ShouldBind(&p)
	if err == nil {
		context.String(200, "%v", p)
	} else {
		context.String(200, "p bind error:%v", err)
	}
}

func main() {
	r := gin.Default()
	r.GET("/testing", testing)
	r.POST("/testing", testing)

	r.Run()
}
