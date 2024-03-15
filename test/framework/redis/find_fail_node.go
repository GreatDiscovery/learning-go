package redis

import (
	"context"
	"fmt"
	"github.com/panjf2000/ants/v2"
	"github.com/redis/go-redis/v9"
	"math/rand"
	"strings"
	"sync"
	"time"
)

func shuffle(arr []string) {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(arr), func(i, j int) {
		arr[i], arr[j] = arr[j], arr[i]
	})
}

func FindClusterFailedIp(ip string) {
	if ip == "" {
		return
	}
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
	//fmt.Println(fmt.Sprintf("redis-cli -h %s cluster nodes", ip))
	nodes := client.ClusterNodes(context.TODO())
	if strings.Contains(nodes.String(), "fail") || strings.Contains(nodes.String(), "handshake") {
		fmt.Println("cluster has failed node ", ip)
	}
}

func main() {
	wg := sync.WaitGroup{}
	pool, _ := ants.NewPoolWithFunc(10, func(i interface{}) {
		FindClusterFailedIp(i.(string))
		wg.Done()
	})
	defer pool.Release()
	start := time.Now()
	allRedisIp := "10.214.175.239"
	arr := strings.Split(allRedisIp, ",")
	shuffle(arr)
	for _, redisIp := range arr {
		wg.Add(1)
		err := pool.Invoke(redisIp)
		if err != nil {
			fmt.Println("failed to submit task:", err)
		}
	}
	wg.Wait()
	fmt.Println("all tasks are finished, cost=", time.Now().Sub(start))
}
