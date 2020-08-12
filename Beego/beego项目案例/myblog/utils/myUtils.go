package utils

import (
	"crypto/md5"
	"fmt"
)

func MD5(str string) string {
	md5str := fmt.Sprintf("%s", md5.Sum([]byte(str)))
	return md5str
}
