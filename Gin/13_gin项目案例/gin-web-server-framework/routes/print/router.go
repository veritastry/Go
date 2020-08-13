package print

import gin "github.com/gin-gonic/gin"

func Route(r *gin.Engine) {
	r.GET("/", IndexGet)
	r.POST("/", IndexPost)
}
