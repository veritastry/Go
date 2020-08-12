package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	r := gin.Default()

	r.GET("/get", func(context *gin.Context) {
		name := context.Query("name")
		age := context.DefaultQuery("age", "defaultAge")
		context.String(http.StatusOK, "%s，%s", name, age)
	})
	r.Run(":8085")
}
