package defs

import "github.com/gin-gonic/gin"

var (
	// ParseToJSONErr parse struct to map err
	ParseToJSONErr = "Errored when parsing to json"
	// CallFuncErr call function err
	CallFuncErr = "Errored when calling a function"
	// HTTPRequestErr send request err
	HTTPRequestErr = "Errored when sending a request"
	// ReadRespBodyErr read responsebody err
	ReadRespBodyErr = "Errored when reading resp body"
	// ConnDBErr connect db err
	ConnDBErr = "Errored when connecting to db"
)

// HTTPError example
type HTTPError struct {
	Code    int    `json:"code" example:"400"`
	Message string `json:"message" example:"status bad request"`
}

// NewError example
func NewError(ctx *gin.Context, status int, err error) {
	er := HTTPError{
		Code:    status,
		Message: err.Error(),
	}
	ctx.JSON(status, er)
}
