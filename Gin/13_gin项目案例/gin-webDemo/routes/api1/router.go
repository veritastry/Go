package api1

import gin "github.com/gin-gonic/gin"

func Route(r *gin.Engine) {
	router := r.Group("/api1")
	{
		router.GET("/", Index)
		router.GET("/h1", Helloworld)
	}
}
