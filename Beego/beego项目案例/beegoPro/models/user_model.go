package models

import (
	"beegoPro/utils/db"
	"fmt"
)

type User struct {
	Id         int64  `json:"id"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	Status     int64  `json:"status"`
	Createtime int64  `json:"createtime"`
}

// 插入
func InserUser(user User) (int64, error) {
	return db.ModifyDB("insert into users(username,password,status,createtime) values (?,?,?,?)",
		user.Username, user.Password, user.Status, user.Createtime)
}

// 通过id查询
func QueryUserById(idname string) int {
	sql := fmt.Sprintf("select id from users %s", idname)
	queryRes := db.Query(sql)
	id := 0
	queryRes.Scan(&id)
	return id
}

// 通过用户名查询
func QueryUserByUsername(username string) int {
	sql := fmt.Sprintf("where username = `%s`", username)
	return QueryUserById(sql)
}
