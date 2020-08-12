package _8_使用gin和gorm框架搭建简易服务_一_

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type Person struct {
	Id        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

var (
	db  *gorm.DB
	err error
)

func main() {

	db, err = gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/test")
	// AutoMigrate运行自动迁移对于给定的模型,只会增加缺失的领域,不会删除/修改当前数据
	db.AutoMigrate(&Person{})
	/*
	  备注：
	  为`name`列添加索引`age`
	  db.Set("age","name").AutoMigrate(&Person{})
	*/
	defer db.Close()

	p1 := Person{FirstName: "John", LastName: "Michle"}
	p2 := Person{FirstName: "Tom", LastName: "Smith"}

	// Create insert the value into database
	// 创建的值插入数据库
	db.Create(&p1)

	// First find first record that match given conditions, order by primary key
	// 按主键首先找到首先给定条件相匹配的记录
	var p3 Person
	db.First(&p3)
	fmt.Println(p1.FirstName)
	fmt.Println(p2.LastName)
	fmt.Println(p3.FirstName)

}
