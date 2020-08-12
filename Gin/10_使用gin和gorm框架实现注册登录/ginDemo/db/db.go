package db

import (
	"fmt"
	"ginDemo/bean"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"

)

var (
	db  *gorm.DB
	err error
)

func Database() {
	// db
	db, err = gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/test")
	checkErr()
	//db.SingularTable(true)
	db.AutoMigrate(&bean.User{})
	defer db.Close()
}

func checkErr() {
	if err != nil {
		fmt.Println(err)
		err = nil
	}
}
