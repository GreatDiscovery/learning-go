package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"net/http"
	_ "net/http/pprof"
	"time"
)

// 在创建大量客户端过程中，发现clusterClient占用内存非常多，所以写一个测试用例来观测clusterClient具体是哪里比较占用内存
// 最终定位到是客户端会缓存一些cluster stat导致内存不断上涨
func main() {
	{
		go func() {
			http.ListenAndServe("0.0.0.0:8899", nil)
		}()

		for i := 0; i < 100; i++ {
			client := redis.NewClusterClient(&redis.ClusterOptions{
				Addrs:         []string{"192.168.155.20:6379"},
				MaxRedirects:  1024,
				ReadOnly:      true,
				RouteRandomly: true,
			})
			statusCmd := client.Ping(context.Background())
			if statusCmd.Err() != nil {
				fmt.Println(statusCmd.Err())
			}
			time.Sleep(10 * time.Second)
		}
		time.Sleep(10 * time.Minute)
	}

}
