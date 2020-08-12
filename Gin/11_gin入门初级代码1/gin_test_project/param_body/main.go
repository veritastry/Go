package main

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

func main() {

	r := gin.Default()

	r.POST("/post", func(context *gin.Context) {
		bodyByes, err := ioutil.ReadAll(context.Request.Body)
		if err != nil {
			context.String(http.StatusBadRequest, err.Error())
			context.Abort()
		}
		context.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyByes))
		name := context.PostForm("name")
		age := context.DefaultPostForm("age", "defaultAge")
		context.String(http.StatusOK, "%s,%s,%s", name, age,string(bodyByes))
	})
	r.Run()
}
