package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"aioc/internal/api"
	"aioc/internal/service"
	"aioc/pkg/setting"
)

// @title aioc
// @version 1.0
// @description  aioc
// @termsOfService

// @contact.name infrawaves
// @contact.url
// @contact.email liufanwh@126.com

// @license.name MIT
// @license.url

// @host   127.0.0.1:8000
// @BasePath /api/v1
func main() {
	if err := setting.LoadConfig(); err != nil {
		log.Fatalf("load config failed, err: %v", err)
	}
	if err := service.ServiceInit(); err != nil {
		log.Fatalf("service init failed, err: %v", err)
	}

	routersInit := api.InitRouter()
	readTimeout := setting.ServerSetting.ReadTimeout
	writeTimeout := setting.ServerSetting.WriteTimeout

	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)
	maxHeaderBytes := 1 << 20

	srv := &http.Server{
		Addr:           endPoint,
		Handler:        routersInit,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}

	go func() {
		// 开始监听
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// 设置优雅关闭的信号处理器
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")

	// 优雅关闭服务器，等待5秒超时
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exited")
}
