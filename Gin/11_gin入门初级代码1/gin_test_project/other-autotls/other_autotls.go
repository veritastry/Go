package other_autotls

import (
	"github.com/gin-gonic/autotls"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	r.GET("/test", func(context *gin.Context) {
		context.String(200, "hello test")
	})
	autotls.Run(r, "www.itpp.tk")

}
