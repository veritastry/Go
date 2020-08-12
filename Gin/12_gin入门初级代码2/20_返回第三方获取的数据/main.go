package _0_返回第三方获取的数据

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	router := gin.Default()

	router.GET("/get", func(c *gin.Context) {
		response, err := http.Get("http://tupian.baike.com/ipad/a0_16_12_01300535031999137270128786964_jpg.html")
		if err != nil || response.StatusCode != http.StatusOK {
			c.Status(http.StatusServiceUnavailable)
			return
		}


		reader := response.Body
		contentLength := response.ContentLength
		contentType := response.Header.Get("Content-Type")
		headers := map[string]string{"Content-Disposition": `attachment; filename="a0_16_12_01300535031999137270128786964_jpg.html"`}

		c.DataFromReader(http.StatusOK, contentLength, contentType, reader, headers)
	})
	router.Run(":9000")
}
