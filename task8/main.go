package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// 设置响应头
	w.Header().Set("Content-Type", "text/plain")
	// 写入响应内容
	fmt.Fprint(w, "Hello, this is a simple back end!")
}

func main() {
	// 注册路由和处理函数
	http.HandleFunc("/", handler)

	// 启动服务器
	fmt.Println("Server is running on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Failed to start server:", err)
	}
}
