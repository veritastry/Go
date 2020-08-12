package models

import (
	"fmt"
	"myblog/utils"
)

type User struct {
	Id         int    `json:"id"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	Status     int64  `json:"status"`
	Createtime int64  `json:"createtime"`
}

// --------------数据库操作---------------

// 插入
func InserUser(user User) (int64, error) {
	return utils.ModifyDB("insert into users(username,password,status,createtime) values (?,?,?,?)",
		user.Username, user.Password, user.Status, user.Createtime)
}

// 按条件查询
func QueryUserWithCon(con string) int {
	sql := fmt.Sprintf("select id from users %s", con)
	fmt.Println(sql)
	row := utils.QueryRowDB(sql)
	id := 0
	row.Scan(&id)
	return id
}

// 根据用户名查询id
func QueryUserWithUsername(username string) int {
	sql := fmt.Sprintf("where username=`%s`", username)
	return QueryUserWithCon(sql)
}

//根据用户名和密码，查询id
//func QueryUserWithParam(username, password string) int {
//	sql := fmt.Sprintf("where username=`%s` and password=`%s`", username, password)
//	return QueryUserWithCon(sql)
//}
