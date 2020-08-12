package _8_XML_JSON_YAML和ProtoBuf渲染_输出格式_

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/testdata/protoexample"
	"net/http"
)

func main() {

	router := gin.Default()

	// someJSON
	router.GET("/someJSON", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "hey",
			"status":  http.StatusOK,
		})
	})

	// moreJSON
	router.GET("/moreJSON", func(c *gin.Context) {
		var msg struct {
			Name     string `json:"name"`
			Password string `json:"password"`
		}

		msg.Name = "lala"
		msg.Password = "123"

		c.JSON(http.StatusOK, gin.H{
			"status": "succeed",
			"msg":    msg,
		})
	})

	// someXML
	router.GET("/someXML", func(c *gin.Context) {
		c.XML(http.StatusOK, gin.H{
			"message": "hey",
			"status":  http.StatusOK,
		})
	})

	// someYAML
	router.GET("/someYAML", func(c *gin.Context) {
		c.YAML(http.StatusOK, gin.H{
			"message": "hey",
			"status":  http.StatusOK,
		})
	})

	// someProtoBuf
	router.GET("/someProtoBuf", func(c *gin.Context) {
		response := []int64{int64(1), int64(2)}
		label := "test"
		data := &protoexample.Test{
			Label: &label,
			Reps:  response,
		}
		c.ProtoBuf(http.StatusOK, data)
	})
	router.Run(":8081")

}
