package greeter

import "fmt"

// Greet 返回问候语
func Greet(name string) string {
	return fmt.Sprintf("Greetings, %s!", name)
}
