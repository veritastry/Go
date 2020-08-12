package login

import (
	db "gin-demo/db"
	"gin-demo/defs"
	"gin-demo/logger"

	"net/http"

	"github.com/gin-gonic/gin"
)

// Add user godoc
// @Summary user
// @Description add user
// @Tags user
// @Accept  json
// @Produce  json
// @Param  user  body db.TBL_USERS true "add user"
// @Success 200 {object} db.TBL_USERS
// @Failure 400 {object} defs.HTTPError
// @Router /login/register [post]
//AddUser add user to db
func AddUser(c *gin.Context) {
	var userParam db.TBL_USERS
	//bind struct param
	if err := c.ShouldBindJSON(&userParam); err != nil {
		defs.NewError(c, http.StatusBadRequest, err)
		return
	}
	user := db.TBL_USERS{
		Name: userParam.Name,
		Pwd:  userParam.Pwd,
	}
	user.Insert()
	c.JSON(http.StatusOK, user)

}

//GetUserInfo  get user info by name
func GetUserInfo(c *gin.Context) {
	gormConn := db.ConnGormMysql()

	user := db.TBL_USERS{Name: c.Param("name")}

	gormConn.Select("name, pwd").Where(&user).Find(&user)
	logger.Debug(user)

	//返回数据
	c.JSON(http.StatusOK, user)
}
