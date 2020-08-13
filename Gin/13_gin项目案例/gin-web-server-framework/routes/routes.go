package routes

import (
	print "gin-demo/routes/print"
	api1 "gin-demo/routes/api1"
	api2 "gin-demo/routes/api2"
	login "gin-demo/routes/login"
	doc "gin-demo/routes/swagger"

	gin "github.com/gin-gonic/gin"
)

// InitRouter ...
func InitRouter() *gin.Engine {
	r := gin.Default()
	doc.Route(r)

	print.Route(r)
	login.Route(r)
	api1.Route(r)
	api2.Route(r)
	return r
}
