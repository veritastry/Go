package _1_获取Get参数

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	router := gin.Default()

	// GET
	router.GET("get/:name/:age", func(c *gin.Context) {

		/*1.传参查询*/
		//name := c.Param("name")
		//age := c.Param("age")

		/*2.获取自定义参数*/
		name := c.DefaultQuery("name", "tom")
		age := c.DefaultQuery("age", "16")

		/*3.是 c.Request.URL.Query().Get("lastname") 的简写
		这里因为没有写数据，所以query不到数据*/
		//name := c.Query("name")
		//age := c.Query("age")
		c.String(http.StatusOK, "get方法：name是%s，age是%s", name, age)

	})

	router.Run(":9000")
}
