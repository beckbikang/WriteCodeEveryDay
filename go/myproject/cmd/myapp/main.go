package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/yourusername/myproject/internal/app"
)

func main() {
	log.Println("Starting myapp service...")
	
	// 初始化应用
	application := app.NewApp()
	
	// 设置路由
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to MyProject!")
	})
	
	// 启动服务
	log.Println("Server listening on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
