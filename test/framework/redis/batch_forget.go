package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"strings"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	allAliveIps := ""
	allInvalidNodeIds := ""
	for _, invalidNodeId := range strings.Split(allInvalidNodeIds, " ") {
		start := time.Now()
		for _, ip := range strings.Split(allAliveIps, " ") {
			wg.Add(1)
			go func(ip string) {
				defer wg.Done()

				if err := recover(); err != nil {
					fmt.Println("Worker recovered from:", err)
				}
				client := redis.NewClient(&redis.Options{
					Addr:        fmt.Sprintf("%s:6379", ip),
					DialTimeout: 200 * time.Millisecond,
				})
				defer func(client *redis.Client) {
					_ = client.Close()
				}(client)

				fmt.Println(fmt.Sprintf("redis-cli -h %s cluster forget %s", ip, invalidNodeId))
				result := client.ClusterForget(context.TODO(), invalidNodeId)
				fmt.Println("result=", result)
			}(ip)
		}
		wg.Wait()
		fmt.Println("total cost=", time.Now().Sub(start))
		time.Sleep(2 * time.Second)
	}
}
