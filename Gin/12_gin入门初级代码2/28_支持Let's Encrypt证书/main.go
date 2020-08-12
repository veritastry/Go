package _8_支持Let_s_Encrypt证书

import (
	"github.com/gin-gonic/autotls"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/acme/autocert"
	"log"
	"net/http"
)

func main() {

	router := gin.Default()

	router.GET("get", func(c *gin.Context) {

		c.JSON(http.StatusOK, gin.H{
			"message": "succeed",
		})
	})
	//log.Fatal(autotls.Run(router,"example.com"))

	m := autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		HostPolicy: autocert.HostWhitelist(),
		Cache:      autocert.DirCache("/var/www/.cache"),
	}
	log.Fatal(autotls.RunWithManager(router, &m))
}
