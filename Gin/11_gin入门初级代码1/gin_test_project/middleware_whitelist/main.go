package main

import "github.com/gin-gonic/gin"

func IPAuthMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		ipList := []string{
			"127.0.0.1",
		}
		flag := false
		clientIP := context.ClientIP()
		for _, host := range ipList {
			if clientIP == host {
				flag = true
				break
			}
		}
		if !flag {
			context.String(401, "%s", "not in iplist", clientIP)
			context.Abort()
		}
	}
}

func main() {

	r := gin.Default()
	r.Use(IPAuthMiddleware())
	r.GET("/test", func(context *gin.Context) {
		context.String(200, "hello world")
	})
	r.Run()
}
