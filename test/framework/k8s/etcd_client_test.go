package k8s

import (
	"context"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"testing"
	"time"
)

// 根据前缀批量拉取
func TestPaginatingQuery(t *testing.T) {
	// 创建 etcd 客户端
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379"}, // etcd 的地址
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		panic(err)
	}
	defer cli.Close()

	// 前缀
	prefix := "/example/prefix/"
	// 每批次限制
	limit := int64(512)
	// 开始键
	startKey := prefix
	// 结束键
	endKey := clientv3.GetPrefixRangeEnd(prefix)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	for {
		// 拉取带限制的 kv 数据
		resp, err := cli.Get(ctx, startKey, clientv3.WithRange(endKey), clientv3.WithLimit(limit))
		if err != nil {
			fmt.Printf("Error fetching data: %v\n", err)
			return
		}

		// 打印本批次的 kv 数据
		for _, kv := range resp.Kvs {
			fmt.Printf("Key: %s, Value: %s\n", kv.Key, kv.Value)
		}

		// 如果当前批次的数量小于限制，说明没有更多数据了
		if int64(len(resp.Kvs)) < limit {
			break
		}

		// 更新起始键，继续下一批次
		lastKey := resp.Kvs[len(resp.Kvs)-1].Key
		startKey = string(append(lastKey, 0)) // 设置为最后一个键之后的键
	}
}
