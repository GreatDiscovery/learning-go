package redis

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"
	"testing"
)

var ctx = context.Background()

func TestPipeline(t *testing.T) {
	// 创建Redis客户端
	client := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: []string{"localhost:6379"},
	})

	// 检查连接
	if err := client.Ping(ctx).Err(); err != nil {
		log.Fatal("无法连接到Redis服务器:", err)
	}

	// 创建管道
	pipe := client.Pipeline()

	// 添加管道命令
	pipe.Set(ctx, "key1", "value1", 0)
	pipe.Set(ctx, "key2", "value2", 0)
	k1 := pipe.Get(ctx, "key1")
	k2 := pipe.Get(ctx, "key2")

	// 执行管道命令并获取结果
	_, err := pipe.Exec(ctx)
	if err != nil {
		log.Fatal("管道执行失败:", err)
	}

	v1, _ := k1.Result()
	fmt.Println("k1=", v1)
	v2, _ := k2.Result()
	fmt.Println("k2=", v2)
}
