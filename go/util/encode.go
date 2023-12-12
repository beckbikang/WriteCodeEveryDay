package util

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func GetMd5Str(str string) string {
	obj := md5.New()
	obj.Write([]byte(str))
	return fmt.Sprintf("%x", obj.Sum(nil))
}

func GetMd5Str2(str string) string {
	bytes := md5.Sum([]byte(str))
	return hex.EncodeToString(bytes[:])
}
