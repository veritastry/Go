package server

import (
	"singo/api"
	"singo/middleware"
	"os"

	"github.com/gin-gonic/gin"
)

// NewRouter 路由配置
func NewRouter() *gin.Engine {
	r := gin.Default()

	// 中间件, 顺序不能改
	r.Use(middleware.Session(os.Getenv("SESSION_SECRET")))
	r.Use(middleware.Cors())
	r.Use(middleware.CurrentUser())

	// 路由
	v1 := r.Group("/api/v1")
	{
		v1.POST("ping", controller.Ping)

		// 用户登录
		v1.POST("user/register", controller.UserRegister)

		// 用户登录
		v1.POST("user/login", controller.UserLogin)

		// 需要登录保护的
		auth := v1.Group("")
		auth.Use(middleware.AuthRequired())
		{
			// User Routing
			auth.GET("user/me", controller.UserMe)
			auth.DELETE("user/logout", controller.UserLogout)
		}
	}
	return r
}
