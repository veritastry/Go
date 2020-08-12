package _7_无中间件启动

import "github.com/gin-gonic/gin"

func main() {

	router := gin.New()

	// 这种属于默认启动方式，包含Loggr,Recovery中间件
	//router := gin.Default()

	router.Run()
}
