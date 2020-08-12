package utils

import (
	"database/sql"
	"fmt"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func InitMysql() {

	fmt.Println("initMysql....")
	driveName := beego.AppConfig.String("driveName")

	// 连接数据库
	user := beego.AppConfig.String("mysqluser")
	pwd := beego.AppConfig.String("mysqlpwd")
	host := beego.AppConfig.String("host")
	port := beego.AppConfig.String("port")
	dbname := beego.AppConfig.String("dbname")

	// root:root@tcp{127.0.0.1:3306}/myblog?charset=utf8
	dbConn := user + ":" + pwd + "@tcp(" + host + ":" + port + ")/" + dbname + "?charset=utf8"

	// 打开数据库
	database, err := sql.Open(driveName, dbConn)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		db = database
		//CreateTableWithUser()
	}
}

// 操作数据库
func ModifyDB(sql string, args ...interface{}) (int64, error) {
	result, err := db.Exec(sql, args...)
	if err != nil {
		fmt.Println(err.Error())
		return 0, err
	}
	count, err := result.RowsAffected()
	if err != nil {
		fmt.Println(err.Error())
		return 0, err
	}
	return count, nil
}

// 创建用户表
//func CreateTableWithUser() {
//
//	sql := `CREATE TABLE IF NOT EXISTS users(
//		id int(4) PRIMARY KEY NOT NULL AUTO_INCREMENT,
//		username varchar(64),
//		password varchar(64),
//		status int(4),
//		createtime int(10)
//	);`
//	ModifyDB(sql)
//}

//查询
func QueryRowDB(sql string) *sql.Row {
	return db.QueryRow(sql)
}
