package _3_上传单个文件

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	router := gin.Default()

	// 给表单限制上传大小（默认为32M）
	router.MaxMultipartMemory = 8 << 20
	router.POST("/upload", func(c *gin.Context) {

		// 1.获取文件头
		file, err := c.FormFile("upload")
		if err != nil {
			fmt.Println(err.Error())
			//c.String(http.StatusBadRequest, "请求失败")
			//return
		}

		// 2.获取文件名
		fileName := file.Filename
		fmt.Println("文件名：", fileName)

		// 3.保存文件到服务器本地
		err = c.SaveUploadedFile(file, fileName)
		if err != nil {
			c.String(http.StatusBadRequest, "保存失败 Error:%s", err.Error())
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"code":    "200",
			"message": "succeed",
			"file":    file.Filename,
		})

		//c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))

	})
	router.Run()

}
