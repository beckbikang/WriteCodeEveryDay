package utils

import (
	"fmt"
	"time"
)

// FormatTime 将时间格式化为标准格式
func FormatTime(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

// GenerateGreeting 生成问候语
func GenerateGreeting(name string) string {
	return fmt.Sprintf("Hello, %s! Welcome to MyProject.", name)
}
