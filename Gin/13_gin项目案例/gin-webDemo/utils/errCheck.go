package utils

import "gin-demo/logger"

/*ErrorCheck check if there is an error  */
func ErrorCheck(errMsg string, err error) {
	if err != nil {
		logger.Error(errMsg, err.Error())
		panic(err.Error())
	}
}
