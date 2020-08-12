package _3_将请求体绑定到不同的结构体中

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type formA struct {
	Name string `json:"name" xml:"name" binding:"required"`
}

type formB struct {
	Hobby string `json:"hobby" xml:"hobby" binding:"required"`
}

func Handler(c *gin.Context) {
	objA := formA{}
	objB := formB{}

	if err := c.ShouldBind(&objA); err == nil {
		c.JSON(http.StatusOK, gin.H{
			"status": "formA succeed",
		})
	} else if err := c.ShouldBind(&objB); err == nil {
		c.JSON(http.StatusOK, gin.H{
			"status": "formB succeed",
		})
	}

}

func main() {
	router := gin.Default()

	router.GET("/", Handler)

	router.Run(":6767")
}
