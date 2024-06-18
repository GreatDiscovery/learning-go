package main

import (
	"context"
	"github.com/sirupsen/logrus"
	"log"
	"os"
	"os/signal"
	"syscall"
	"testing"
)

func handleExitSignals(ctx context.Context, logger *logrus.Entry, cancel context.CancelFunc) {
	ch := make(chan os.Signal, 32)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)

	for {
		select {
		case s := <-ch:
			logger.WithField("signal", s).Debugf("Caught exit signal")
			cancel()
			return
		case <-ctx.Done():
			return
		}
	}
}

func TestSignal(t *testing.T) {
	// 创建一个channel用于接收操作系统发送的信号
	sigChan := make(chan os.Signal, 1)
	// 监听SIGINT和SIGTERM信号
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// 创建一个channel用于通知程序退出
	exitChan := make(chan struct{})

	// 启动一个goroutine来处理信号
	go func() {
		select {
		// 阻塞等待信号
		case sig := <-sigChan:
			// 收到信号后打印日志
			log.Printf("Received signal: %v", sig)
			// 向exitChan发送消息，通知程序退出
			close(exitChan)
		}
	}()

	// 在主goroutine中等待exitChan被关闭，然后退出程序
	<-exitChan
	log.Println("Shutting down gracefully...")
	// 在这里执行任何需要在退出时清理的操作
	log.Println("Exiting.")
}
