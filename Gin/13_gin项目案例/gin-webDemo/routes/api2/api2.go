package api2

import (
	"fmt"
	http "net/http"

	gin "github.com/gin-gonic/gin"
)

// Index ...
func Index(c *gin.Context) {
	c.String(http.StatusOK, "hello, index, api2")
}

//Helloworld ...
func Helloworld(c *gin.Context) {
	fmt.Print()
	c.String(http.StatusOK, "hello, world, api2")
}
