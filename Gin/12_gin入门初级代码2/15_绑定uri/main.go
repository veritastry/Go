package _5_绑定uri

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Person struct {
	Uri   string `uri:"uri" binding:"required,uuid"`
	Name string `uri:"name" binding:"required"`
}

func main() {

	router := gin.Default()
	router.Any("/:name/:uri", func(c *gin.Context) {
		var person Person
		err := c.ShouldBindUri(&person)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"status": "succeed",
			"id":     person.Uri,
			"name":   person.Name,
		})
	})
	router.Run()
}
