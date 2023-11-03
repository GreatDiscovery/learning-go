package k8s

import (
	"context"
	"fmt"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"testing"
)

// 测试资源预留，不知道接口调用后的实现逻辑
func TestResourceReserve(t *testing.T) {
	clientSet, _ := initClient()
	list, err := clientSet.ResourceV1alpha2().ResourceClaims("default").List(context.TODO(), v1.ListOptions{})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(list)
	}
}
