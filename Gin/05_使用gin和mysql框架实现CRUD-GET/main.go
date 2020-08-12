package _5_使用gin和mysql框架实现CRUD_GET

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
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

	router.GET("/food", func(c *gin.Context) {

		// 创建对象
		var (
			food  Food
			foods []Food
		)

		// GET获取数据
		// 这里用db.Query方法
		rows, err := db.Query("select id,color,weight from food")
		if err != nil {
			fmt.Println(err.Error())
		}
		for rows.Next() {
			err = rows.Scan(&food.Id, &food.Color, &food.Weight)
			if err != nil {
				fmt.Println(err.Error())
			}
			defer rows.Close()

			c.JSON(http.StatusOK, gin.H{
				"result": foods,
				"count":  len(foods),
			})
		}
	})
	router.Run(":9001")

}
