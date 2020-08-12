package api2

import gin "github.com/gin-gonic/gin"

func Route(r *gin.Engine) {
	router := r.Group("/api2")
	{
		router.GET("/", Index)
		router.GET("/h1", Helloworld)
	}
}
