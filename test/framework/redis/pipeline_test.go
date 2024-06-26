package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"
	"testing"
)

var ctx = context.Background()

func TestGet(t *testing.T) {
	// 创建Redis客户端
	client := redis.NewClient(&redis.Options{
		//Addrs: []string{"192.168.143.120:12345"},
		Addr: "192.168.143.120:12345",
		// redis-go有bug，会导致某些corvus链接失败,https://github.com/redis/go-redis/issues/2616
		Protocol: 2,
	})

	// 检查连接
	//if err := client.Ping(ctx).Err(); err != nil {
	//	log.Fatal("无法连接到Redis服务器:", err)
	//}

	result := client.Get(ctx, "k1")
	fmt.Println("result=", result)

}

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
