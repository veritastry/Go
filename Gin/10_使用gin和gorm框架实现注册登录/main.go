package _10_使用gin和gorm框架实现注册登录

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type User struct {
	Id       int    `json:"id",gorm:"auto-increment"`
	Name     string `json:"name"`
	Tel      string `json:"tel"`
	Password string `json:"password"`
}

type ResponseData struct {
	Code    int         `json:"code"`
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

var (
	db  *gorm.DB
	err error
)

func main() {
	// db
	db, err = gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/test")
	checkErr()
	//db.SingularTable(true)
	db.AutoMigrate(&User{})
	defer db.Close()
	// http
	router := gin.Default()
	v1 := router.Group("/api/v1")
	{
		// user
		v1.POST("/user/login", login)
		v1.POST("/user/register", register)
		v1.GET("/user/get_info", getUserInfo)
	}
	router.Run("5000")
}

func checkErr() {
	if err != nil {
		fmt.Println(err)
		err = nil
	}
}

// post form
func login(c *gin.Context) {
	// 获取表单数据 tel，password
	tel := c.PostForm("tel")
	psd := c.PostForm("password")
	if len(tel) == 0 || len(psd) == 0 {
		response := ResponseData{
			Code:    50001,
			Status:  "error",
			Message: "账号或密码不能为空",
			Data:    "",
		}
		c.JSON(200, response)
	} else {
		var user User
		db.Where("tel=?", tel).First(&user)
		if user.Tel == "" {
			response := ResponseData{
				Code:    50001,
				Status:  "error",
				Message: "用户不存在",
				Data:    "",
			}
			c.JSON(200, response)
		} else {
			if user.Password == psd {
				response := ResponseData{
					Code:    200,
					Status:  "success",
					Message: "登录成功",
					Data:    "",
				}
				c.JSON(200, response)
			} else {
				response := ResponseData{
					Code:    5001,
					Status:  "error",
					Message: "密码错误",
					Data:    "",
				}
				c.JSON(200, response)
			}
		}
	}
}

// post form
func register(c *gin.Context) {
	tel := c.PostForm("tel")
	psd := c.PostForm("password")
	if len(tel) == 0 || len(psd) == 0 {
		response := ResponseData{
			Code:    50001,
			Status:  "error",
			Message: "账号或密码不能为空",
			Data:    "",
		}
		c.JSON(200, response)
		return
	}
	var user User
	db.Where("tel=?", tel).First(&user)
	if user.Tel == tel {
		response := ResponseData{
			Code:    50001,
			Status:  "error",
			Message: "手机号已注册",
			Data:    "",
		}
		c.JSON(200, response)
	} else {
		newUser := User{Tel: tel, Password: psd}
		db.Create(&newUser)
		response := ResponseData{
			Code:    200,
			Status:  "success",
			Message: "注册成功",
			Data:    "",
		}
		c.JSON(200, response)
	}
}

// get url
func getUserInfo(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		response := ResponseData{
			Code:    50001,
			Status:  "error",
			Message: "参数错误",
			Data:    "",
		}
		c.JSON(200, response)
		return
	}
	var user User
	db.First(&user, id)
	if user.Id > 0 {
		response := ResponseData{
			Code:    200,
			Status:  "success",
			Message: "",
			Data:    user,
		}
		c.JSON(200, response)
	} else {
		response := ResponseData{
			Code:    50001,
			Status:  "error",
			Message: "用户不存在",
			Data:    "",
		}
		c.JSON(200, response)
	}
}
