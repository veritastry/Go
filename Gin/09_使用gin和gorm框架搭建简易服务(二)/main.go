package _9_使用gin和gorm框架搭建简易服务_二_

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"net/http"
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
	db.AutoMigrate(&Person{})
	defer db.Close()

	router := gin.Default()
	router.GET("/getAllPerson", GetAllPerson)
	router.GET("/getPersonById/:id", GetPersonById)
	router.POST("/addPerson", AddPerson)
	router.PUT("/updatePerson/:id", UpdatePerson)
	router.DELETE("/deletePerson/:id", DeletePerson)
	router.Run()

}

// 增加person信息
func AddPerson(c *gin.Context) {
	var person Person
	c.BindJSON(&person)
	db.Create(&person)
	c.JSON(http.StatusOK, person)
}

// 覆盖更新person信息
func UpdatePerson(c *gin.Context) {
	var person Person
	id := c.Params.ByName("id")
	err := db.Where("id = ?", id).First(&person).Error
	if err != nil {
		c.AbortWithStatus(403)
		fmt.Println(err)
	}
	c.BindJSON(&person)
	db.Save(&person)
	c.JSON(http.StatusOK, person)

}

// 通过id获取某一条person的信息
func GetPersonById(c *gin.Context) {
	id := c.Params.ByName("id")
	var person Person
	err := db.Where("id = ?", id).First(&person).Error
	if err != nil {
		c.AbortWithStatus(403)
		fmt.Println(err)
	}
	c.JSON(http.StatusOK, person)

}

// 获取所有person信息
func GetAllPerson(c *gin.Context) {
	var person []Person
	// Find: records that match given conditions
	//找到找到给定的条件相匹配的记录
	err := db.Find(&person)
	if err != nil {
		c.AbortWithStatus(403)
		fmt.Println(err)
	}
	c.JSON(http.StatusOK, person)

}

// 通过id删除person信息
func DeletePerson(c *gin.Context) {
	id := c.Params.ByName("id")
	var person Person
	err := db.Where("id = ?", id).Delete(&person)
	if err != nil {
		c.AbortWithStatus(403)
		fmt.Println(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"id #" + id: "id is deleted",
	})

}
