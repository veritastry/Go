package _7_使用gin和mysql框架实现CRUD_DELETE

import (
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

	// DELETE
	// postman请求方式：Params写id参数
	router.DELETE("/food", func(c *gin.Context) {
		id := c.Query("id")
		statement, err := db.Prepare("delete from food where id= ?;")
		if err != nil {
			fmt.Println(err.Error())
		}
		_, err = statement.Exec(id)
		if err != nil {
			fmt.Println(err.Error())
		}
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("Successfully deleted food: %s", id),
		})
	})
	router.Run(":9002")

}
