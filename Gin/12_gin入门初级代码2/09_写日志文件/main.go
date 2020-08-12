package _9_写日志文件

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
)

func main() {

	// 禁用制台颜色
	gin.DisableConsoleColor()

	// 创建记录日志的文件
	file, err := os.Create("bbb.txt")
	if err != nil {
		fmt.Println(err.Error())
	}
	//gin.DefaultWriter = io.MultiWriter(file)

	// 可以同时讲日志写入文件和控制台
	gin.DefaultWriter = io.MultiWriter(file, os.Stdout)

	router := gin.Default()
	router.GET("/get", func(c *gin.Context) {

		c.String(http.StatusOK, "successful")
	})
	router.Run()
}
