package main

import (
	"github.com/go-programming-tour-book/blog-service/internal/routers"
	"net/http"
	"time"
)

func main() {
	router := routers.NewRouter()
	server := &http.Server{ // 自定义http.Server
		Addr:           ":8080",          // 设置监听的TCP 端口
		Handler:        router,           // 处理路由程序
		ReadTimeout:    10 * time.Second, // 允许读取/写入的最大时间
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20, // 请求头的最大字节数
	}
	server.ListenAndServe()
}
