package print

import (
	http "net/http"
	logger "gin-demo/logger"
	gin "github.com/gin-gonic/gin"
	"io/ioutil"
)

func IndexGet(c *gin.Context) {
	query := c.Request.URL.Query()
	logger.Info(query)
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func IndexPost(c *gin.Context) {
	query := c.Request.URL.Query()
	logger.Infof("Query: %v", query)
	c.Request.ParseForm()
	logger.Infof("JsonForm: %v", c.Request.PostForm )
	data, _ := ioutil.ReadAll(c.Request.Body)
	logger.Infof("BodyForm: %v", string(data))
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
