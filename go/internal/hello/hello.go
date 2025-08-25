package hello

import "fmt"

// Hello 返回问候语
func Hello(name string) string {
	return fmt.Sprintf("Hello, %s!", name)
}
