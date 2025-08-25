package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello, WriteCodeEveryDay!")
	
	// 这里是应用程序的入口点
	// 你可以在这里添加你的程序逻辑
	
	if len(os.Args) > 1 {
		fmt.Printf("Arguments received: %v\n", os.Args[1:])
	}
}
