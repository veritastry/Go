package _2_使用自定义结构绑定表单数据

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type StructA struct {
	FieldA string `form:"field_a"`
}

type StructB struct {
	Struct1 StructA
	FieldB       string `form:"field_b"`
}

type StructC struct {
	Struct2 *StructA
	FieldC              string `form:"field_c"`
}

type StructD struct {
	Struct3 struct {
		FieldX string `form:"field_x"`
	}
	FieldD string `form:"field_d"`
}

func GetDataB(c *gin.Context) {
	var b StructB
	c.Bind(&b)
	c.JSON(http.StatusOK, gin.H{
		"a": b.Struct1,
		"b": b.FieldB,
	})
}

func GetDataC(c *gin.Context) {
	var b StructC
	c.Bind(&b)
	c.JSON(http.StatusOK, gin.H{
		"a": b.Struct2,
		"c": b.FieldC,
	})
}

func GetDataD(c *gin.Context) {
	var b StructD
	c.Bind(&b)
	c.JSON(http.StatusOK, gin.H{
		"x": b.Struct3,
		"d": b.FieldD,
	})
}

func main() {
	router := gin.Default()

	router.GET("/getb", GetDataB)
	router.GET("/getc", GetDataC)
	router.GET("/getd", GetDataD)

	router.Run(":7979")
}
