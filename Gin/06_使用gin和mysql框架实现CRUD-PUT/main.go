package _6_使用gin和mysql框架实现CRUD_PUT

import (
	"bytes"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
)
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

	// PUT是更新的操作
	// postman请求方式是：在Body/form-data填写参数
	router.PUT("/food", func(c *gin.Context) {

		var buffer bytes.Buffer
		id := c.Query("id")
		color := c.PostForm("color")
		weight := c.PostForm("weight")

		statement, err := db.Prepare("update food set color= ?, weight= ? where id= ?;")
		if err != nil {
			fmt.Println(err.Error())
		}
		_, err = statement.Exec(color, weight, id)
		if err != nil {
			fmt.Println(err.Error())
		}
		defer statement.Close()

		buffer.WriteString(color)
		buffer.WriteString(weight)
		name := buffer.String()
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("Successfully updated to %s", name),
		})
	})
	router.Run(":9001")

}
