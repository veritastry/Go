package _8_使用中间件

import "github.com/gin-gonic/gin"

func main() {

	// 创建一个包含中间件的路由器
	router := gin.New()

	// 使用Logger中间件
	router.Use(gin.Logger())
	// 使用Recovery中间件
	router.Use(gin.Recovery())

	// 使用路由添加中间件
	router.GET("get")

	// 使用路由组添加中间件
	v1 := router.Group("/v1")
	//v1.Use(AuthRequired())  // 这里添加中间件
	{
		v1.POST("login")
		v1.POST("register")
	}
	router.Run()

}
