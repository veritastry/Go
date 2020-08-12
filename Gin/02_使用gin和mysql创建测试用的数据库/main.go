package _2_使用gin和mysql创建测试用的数据库

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main()  {

	// 打开数据库
	db,err := sql.Open("mysql","root:root@tcp(127.0.0.1:3306)/test")
	if err != nil {
		fmt.Println(err.Error())
	}

	// 关闭数据库，延迟调用
	defer db.Close()

	// 验证数据库是否连接
	err = db.Ping()
	if err != nil {
		fmt.Println(err.Error())
	}

	// 准备查询语句
	statement,err := db.Prepare("CREATE TABLE people (id int NOT NULL AUTO_INCREMENT, first_name varchar(40), last_name varchar(40), PRIMARY KEY (id));")
	if err != nil {
		fmt.Println(err.Error())
	}

	// 执行查询语句
	_, err = statement.Exec()
	if err != nil {
		fmt.Println(err.Error())
	}else{
		fmt.Printf("People Table successfully migrated...")
	}






}
