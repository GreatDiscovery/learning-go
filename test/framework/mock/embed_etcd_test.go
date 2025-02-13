package mock

import (
	"go.etcd.io/etcd/server/v3/embed"
	"log"
	"testing"
	"time"
)

// 启动一个内嵌的etcd，方便进行单元测试
func TestNewEmbedEtcd(t *testing.T) {
	cfg := embed.NewConfig()
	cfg.Dir = "default.etcd"
	e, err := embed.StartEtcd(cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer e.Close()
	select {
	case <-e.Server.ReadyNotify():
		log.Printf("Server is ready!")
	case <-time.After(60 * time.Second):
		e.Server.Stop() // trigger a shutdown
		log.Printf("Server took too long to start!")
	}
	log.Fatal(<-e.Err())
}
