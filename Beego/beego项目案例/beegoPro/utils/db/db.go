package db

import (
	"database/sql"
	"fmt"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func Mysql() {

	driveName := beego.AppConfig.String("driveName")

	// 1. 连接数据库
	user := beego.AppConfig.String("mysqluser")
	pwd := beego.AppConfig.String("mysqlpwd")
	host := beego.AppConfig.String("host")
	port := beego.AppConfig.String("port")
	dbname := beego.AppConfig.String("dbname")

	// root:root@tcp(127.0.0.1:3306)/myblog?charset=utf8
	dbConn := user + ":" + pwd + "@tcp(" + host + ":" + port + ")/" + dbname + "?charset=utf8"

	// 2. 打开数据库
	database, err := sql.Open(driveName, dbConn)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		db = database
	}
}

// 操作数据库(增)
func ModifyDB(sql string, args ...interface{}) (int64, error) {
	result, err := db.Exec(sql, args...)
	if err != nil {
		fmt.Println(err.Error())
		return 0, err
	}
	response, err := result.RowsAffected()
	if err != nil {
		fmt.Println(err.Error())
		return 0, err
	}
	return response, nil
}

// 操作数据库(查)
func Query(sql string) *sql.Row {
	return db.QueryRow(sql)
}
