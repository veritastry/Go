package _5_上传单个和多个文件

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"net/http"
)

func main() {

	router := gin.Default()

	// 为 mutipart 表单 设置一个较低的内存限制（默认是 32 MiB）
	router.MaxMultipartMemory = 1
	router.POST("/oneFileUpload", oneFileUpload)
	router.POST("/twoFileUpload", twoFileUpload)

	router.Run(":3000")

}

func oneFileUpload(c *gin.Context) {
	file, _ := c.FormFile("file")
	log.Println(file.Filename)

	c.String(http.StatusOK, fmt.Sprint("'%s' uploaded"), file.Filename)

}

func twoFileUpload(c *gin.Context) {

	form, _ := c.MultipartForm()
	files := form.File["twoFileUpload[]"]
	for _, file := range files {
		log.Println(file.Filename)
	}

	//c.JSON(http.StatusOK, gin.H{
	//	"code":"200",
	//	"message":"成功",
	//	"length":len(files),
	//})
	c.String(http.StatusOK, fmt.Sprintf("%d files uploaded!", len(files)))


}
