package _3_只绑定GET参数

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Person struct {
	Name    string `form:"name" json:"name"`
	Address string `form:"address" json:"address"`
}

func main() {

	router := gin.Default()
	router.Any("startPage", StartPage) // Any是可以匹配所有的注册路由的方法
	router.Run()
}

func StartPage(c *gin.Context) {

	var person Person

	// 如果是ShouldBind可以使用get请求，也可以使用post请求
	// 如果是ShouldBindQuery，因为带query查询，那可以判断是get请求 err := c.ShouldBindQuery(&person)
	// 如果我ShouldBindWith,只能是get请求 err := c.ShouldBindWith(&person,binding.Query)
	err := c.ShouldBindQuery(&person)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  "succeed",
		"name":    person.Name,
		"address": person.Address,
	})

}
