package main

import (
	"ginDemo/controller"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	// http
	router := gin.Default()
	v1 := router.Group("/api/v1")
	{
		// user
		v1.POST("/user/login",controller.Login)
		v1.POST("/user/register",controller.Register)
		v1.GET("/user/get_info",controller.GetUserInfo)
	}
	router.Run(":9000")
}
