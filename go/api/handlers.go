package api

import (
	"net/http"
	"fmt"
)

// Handler 处理HTTP请求
func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the WriteCodeEveryDay API!")
}
