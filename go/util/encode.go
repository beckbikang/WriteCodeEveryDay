package util

import (
	"crypto/md5"
	"fmt"
)

func GetMd5Str(str string) string {
	obj := md5.New()
	obj.Write([]byte(str))
	return fmt.Sprintf("%x", obj.Sum(nil))
}
