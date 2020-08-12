package JSON输出

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	router := gin.Default()

	// SecureJSON
	/*
	输出结果：
	while(1);[
	   "lala",
	   "lulu",
	   "lele"
	]
	*/
	router.GET("someJson", func(c *gin.Context) {
		names := []string{"lala", "lulu", "lele"}
		c.SecureJSON(http.StatusOK, names)
	})

	// JSONP
	/*
	输出结果：
		{
   			 "message": "oh yeah"
		}
	*/
	router.GET("/jsonP", func(c *gin.Context) {
		data := map[string]interface{}{
			"message": "oh yeah",
		}
		c.JSONP(http.StatusOK, data)
	})

	// AsciiJSON
	/*
	输出结果：
	{
    "lang": "GO语言",
    "tag": "<br>"
}
	*/
	router.GET("/someJSON", func(c *gin.Context) {
		data := map[string]interface{}{
			"lang": "GO语言",
			"tag":  "<br>",
		}
		c.AsciiJSON(http.StatusOK, data)
	})

	// PureJSON
	/*
	输出结果
	{
        "html": "<b>Hello, world!</b>"
    }
	*/
	router.GET("/pureJSON", func(c *gin.Context) {
		c.PureJSON(http.StatusOK, gin.H{
			"html": "<b>Hello, world!</b>",
		})
	})
	router.Run(":8082")

}

