package api1

import (
	http "net/http"

	gin "github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	c.String(http.StatusOK, "hello, index, api1")
}

func Helloworld(c *gin.Context) {
	c.String(http.StatusOK, "hello, world, api1")
}
