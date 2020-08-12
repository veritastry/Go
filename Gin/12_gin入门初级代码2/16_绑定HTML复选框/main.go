package _6_绑定HTML复选框

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Form struct {
	Colors []string `form:"colors[]"`
}

func formHandler(c *gin.Context) {

	var myForm Form
	err := c.ShouldBind(&myForm)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "succeed",
		"color":  myForm.Colors,
	})

}

func main() {

	router := gin.Default()
	router.Any("/formHandler", formHandler)
	router.Run()
}
