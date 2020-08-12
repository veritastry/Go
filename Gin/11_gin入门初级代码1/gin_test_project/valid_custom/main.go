package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"reflect"
	"time"
	"gopkg.in/go-playground/validator.v8"
)

type Booking struct {
	CheckIn time.Time `form:"check_in" binding
:"required,bookabledate" time_format:"2006-01-02"`
	CheckOut time.Time `form:"check_out" binding
:"required gtfield=CheckIn" time_format:"2006-01-02"`
}

// 定义验证器,官方提供 validator.v8
func bookableDate(v *validator.Validate, topStruct reflect.Value,
	currentStructOrField reflect.Value,
	field reflect.Value, fieldType reflect.Type,
	fieldKind reflect.Kind, param string) bool {
	if date, ok := field.Interface().(time.Time); ok {
		today := time.Now()
		if date.Unix() > today.Unix() {
			return true
		}
	}
	return false
}

func main() {

	r := gin.Default()

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("bookabledate", bookableDate)
	}

	r.GET("/bookable", func(context *gin.Context) {
		var b Booking
		if err := context.ShouldBind(&b); err != nil {
			context.JSON(500, gin.H{"error": err.Error()})
			return
		}
		context.JSON(200, gin.H{"message": "ok!", "booking": b})
	})
	r.Run()
}
