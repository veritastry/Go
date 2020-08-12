package _3_使用gin和mysql框架实现CRUD

import (
	"bytes"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
)

func main() {

	/*操作数据库*/
	// 1.打开数据库
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/test")
	if err != nil {
		fmt.Println(err.Error())
	}

	// 2.关闭数据库
	defer db.Close()

	// 3.测试数据库
	err = db.Ping()
	if err != nil {
		fmt.Println(err.Error())
	}

	/*CRUD实现*/
	// 1.定义结构体对象
	type Person struct {
		Id         int    `json:"id"`
		First_Name string `json:"first_name"`
		Last_Name  string `json:"last_name"`
	}

	// GET
	router := gin.Default()
	router.GET("/persons", func(c *gin.Context) {

		// 1.创建对象
		var (
			person  Person
			persons []Person
		)

		// 2.Query查询方式
		rows, err := db.Query("select id,first_name,last_name from person;")
		if err != nil {
			fmt.Println(err.Error())
		}

		// 3.将下一个person存储到persons
		for rows.Next() {
			err = rows.Scan(&person.Id, &person.First_Name, &person.Last_Name)
			persons = append(persons, person)
			if err != nil {
				fmt.Println(err.Error())
			}

			// 4.关闭，延迟调用
			defer rows.Close()

			// 5.返回参数
			c.JSON(http.StatusOK, gin.H{
				"result": persons,
				"count":  len(persons),
			})
		}
	})

	// POST
	router.POST("/person", func(c *gin.Context) {
		var buffer bytes.Buffer
		first_name := c.PostForm("first_name")
		last_name := c.PostForm("last_name")

		stmt, err := db.Prepare("insert into person (first_name, last_name) values(?,?);")
		if err != nil {
			fmt.Println(err.Error())
		}
		_, err = stmt.Exec(first_name, last_name)
		if err != nil {
			fmt.Println(err.Error())
		}

		buffer.WriteString(first_name)
		buffer.WriteString(last_name)
		defer stmt.Close()
		name := buffer.String()

		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("%s successfully created", name),
		})

	})

	// PUT
	router.PUT("/person", func(c *gin.Context) {
		var buffer bytes.Buffer
		id := c.Query("id")
		first_name := c.PostForm("first_name")
		last_name := c.PostForm("last_name")

		stmt, err := db.Prepare("update person set first_name= ?, last_name= ? where id= ?;")
		if err != nil {
			fmt.Println(err.Error())
		}
		_, err = stmt.Exec(first_name, last_name, id)
		if err != nil {
			fmt.Println(err.Error())
		}

		buffer.WriteString(first_name)
		buffer.WriteString(last_name)
		defer stmt.Close()
		name := buffer.String()
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("Successfully updated to %s", name)})
	})

	// DELETE
	router.DELETE("/person", func(c *gin.Context) {
		id := c.Query("id")
		stmt,err := db.Prepare("delete from person where id=?;")
		if err != nil {
			fmt.Println(err.Error())
		}
		_,err = stmt.Exec(id)
		if err != nil {
			fmt.Println(err.Error())
		}
		c.JSON(http.StatusOK,gin.H{
			"message": fmt.Sprintf("Successfully deleted user: %s", id),
		})
	})

	router.Run(":3001")
}
