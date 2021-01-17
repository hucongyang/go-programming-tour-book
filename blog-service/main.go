package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-programming-tour-book/blog-service/global"
	"github.com/go-programming-tour-book/blog-service/internal/routers"
	"github.com/go-programming-tour-book/blog-service/pkg/setting"
	"log"
	"net/http"
	"time"
)

// 程序执行顺序：全局变量初始化 -> init方法 -> main方法

func init() {
	err := setupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting err: %v", err)
	}
}

func main() {
	gin.SetMode(global.ServerSetting.RunMode)
	router := routers.NewRouter()
	server := &http.Server{ // 自定义http.Server
		Addr:           ":" + global.ServerSetting.HttpPort, // 设置监听的TCP 端口
		Handler:        router,                              // 处理路由程序
		ReadTimeout:    global.ServerSetting.ReadTimeout,    // 允许读取/写入的最大时间
		WriteTimeout:   global.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20, // 请求头的最大字节数
	}
	server.ListenAndServe()
}

/**
读取配置
*/
func setupSetting() error {
	setting, err := setting.NewSetting()
	if err != nil {
		return err
	}
	err = setting.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("App", &global.AppSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		return err
	}
	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second
	return nil
}
