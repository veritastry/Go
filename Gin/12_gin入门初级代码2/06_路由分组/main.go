package _6_路由分组

import "github.com/gin-gonic/gin"

func main() {

	router := gin.Default()

	v1 := router.Group("/v1")
	{
		v1.POST("register")
		v1.POST("login")
		v1.POST("query")
	}

	v2 := router.Group("/v2")
	{
		v2.POST("register")
		v2.POST("login")
		v2.POST("query")
	}

	router.Run()
}
