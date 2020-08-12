package _4_使用gin和mysql框架实现CRUD_POST

import (
	"bytes"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Food struct {
	Id     int    `json:"id"`
	Color  string `json:"color"`
	Weight string `json:"weight"`
}

func main() {

	driverName := "mysql"
	sqlUser := "root"
	sqlPwd := "root"

	// 打开数据库
	db, err := sql.Open(driverName, sqlUser+":"+sqlPwd+"@tcp(127.0.0.1:3306)/test")
	if err != nil {
		panic(err.Error())
	}
	// 关闭数据库，延迟调用
	defer db.Close()
	// 测试数据库是否连接
	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	router := gin.Default()

	// POST
	// post请求是写入数据，所以这里用buffer写入数据，所以要用db.Prepare而不能用db.Query
	// c.PostForm的效果与c.Request.PostFormValue的效果一致
	// 插入数据到数据库用insert而不是用select
	router.POST("/food", func(c *gin.Context) {

		var buffer bytes.Buffer

		color := c.Request.PostFormValue("color")
		weight := c.Request.PostFormValue("weight")

		statement, err := db.Prepare("insert into food (color, weight) values(?,?);")
		if err != nil {
			fmt.Println(err.Error())
		}
		_, err = statement.Exec(color, weight)
		if err != nil {
			fmt.Println(err.Error())
		}
		defer statement.Close()

		buffer.WriteString(color)
		buffer.WriteString(weight)
		//name := buffer.String()
		c.JSON(http.StatusOK, gin.H{
			"message": "success",
			//"message":fmt.Sprintf(" %s successfully created", name),
		})
	})
	router.Run(":9001")

}
