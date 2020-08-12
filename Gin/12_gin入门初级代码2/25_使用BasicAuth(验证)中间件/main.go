package _5_使用BasicAuth_验证_中间件

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var persons = gin.H{
	"lulu": gin.H{"phone": "13798583666", "wechat": "luluwechat"},
	"lala": gin.H{"phone": "8883677", "wechat": "lalawechat"},
}

func getPersons(c *gin.Context) {

	user := c.MustGet(gin.AuthUserKey).(string)
	if person, ok := persons[user]; ok {
		c.JSON(http.StatusOK, gin.H{"user": user, "person": person})
	} else {
		c.JSON(http.StatusOK, gin.H{"user": user, "person": "NO PERSON"})
	}

}

func main() {

	router := gin.Default()

	// gin.BasicAuth就是中间件,
	// 它的参数gin.Accounts其实是一个map[string]string类型的映射，这里是用来记录用户名和密码。
	authorized := router.Group("/v1", gin.BasicAuth(gin.Accounts{
		"lulu": "pswlulu", // 用户名：密码
		"lala": "pswlala",
	}))

	authorized.GET("/get", getPersons)
	router.Run(":8000")
}
