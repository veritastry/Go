package main

import (
	config "gin-demo/config"
	logger "gin-demo/logger"
	routes "gin-demo/routes"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

// @title Demo Gin Server API
// @version 1.0

// @contact.name API Support
// @contact.url https://github.com/luxuze/demo-gin-server
// @contact.email luxuze1994@gmial.com

// @license.name DO WHAT THE FUCK YOU WANT TO PUBLIC LICENSE
// @license.url https://github.com/luxuze/demo-gin-server/blob/master/LICENSE

// @host localhost:8080
// @BasePath /
func main() {
	/* loading toml configs */
	cfg := config.Config()

	if !cfg.Debug {
		gin.SetMode(gin.ReleaseMode)
	}
	logger.Noticef("gin is debug mode ? %v", cfg.Debug)

	s := &http.Server{
		Addr:           cfg.Server.Port,
		Handler:        routes.InitRouter(),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	logger.Noticef("Server Listing %s", cfg.Server.Port)
	s.ListenAndServe()
}
