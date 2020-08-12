package _2_获取Post参数

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	router := gin.Default()

	// POST
	router.POST("post", func(c *gin.Context) {

		/*1.post方法用PostForm，无法用c.Param("name")获取数据*/
		//name := c.PostForm("name")
		//age := c.PostForm("age")
		//c.String(http.StatusOK, "post方法: name是%s,age是%s", name, age)

		/*2.使用gin.H{}的方式来获取自定义的数据*/
		name := c.PostForm("name")
		age := c.DefaultPostForm("age", "44")
		c.JSON(http.StatusOK, gin.H{
			"status": "666",
			"name":   name,
			"age":    age,
		})

	})
	router.Run(":9000")

}
