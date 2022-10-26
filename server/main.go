package main

import (
	"api/pkg/logging"
	"api/pkg/setting"
	"api/routers"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	log.Println("Hello,api 正在启动")
	setting.SetUp() //初始化配置
	logging.SetUp() //设置日志文件
	//models.SetUp()  //设置数据库

	router := routers.InitRouter() //初始化路由

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.ServerSetting.HttpPort),
		Handler:        router,
		ReadTimeout:    setting.ServerSetting.ReadTimeout,
		WriteTimeout:   setting.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	// 使用 http.Server - Shutdown() 优雅的关闭http服务
	go func() {
		if err := s.ListenAndServe(); err != nil {
			log.Printf("Listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	log.Println("Shutdown Server ...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown", err)
	}

	log.Println("程序服务关闭退出")
}
