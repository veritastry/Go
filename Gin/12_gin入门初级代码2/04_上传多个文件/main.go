package _4_上传多个文件

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	router := gin.Default()

	router.POST("/upload", func(c *gin.Context) {

		// 1.获取解析后的表单
		form, err := c.MultipartForm()
		if err != nil {
			c.String(http.StatusBadRequest, "请求失败")
		}
		return

		// 2.获取所有文件
		files := form.File["upload[]"]

		// 3.保存文件到服务器本地
		for _, file := range files {
			c.SaveUploadedFile(file, file.Filename)
		}
		c.String(http.StatusOK, fmt.Sprintf("%d 个文件被上传成功!", len(files)))

	})
	router.Run()

}
