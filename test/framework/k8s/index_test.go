package k8s

import (
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/cache"
	"testing"
)

func testIndexFunc(obj interface{}) ([]string, error) {
	pod := obj.(*v1.Pod)
	return []string{pod.Labels["foo"]}, nil
}

func TestGetIndexFuncValues(t *testing.T) {
	index := cache.NewIndexer(cache.MetaNamespaceKeyFunc, cache.Indexers{"testmodes": testIndexFunc})

	pod1 := &v1.Pod{ObjectMeta: metav1.ObjectMeta{Name: "one", Labels: map[string]string{"foo": "bar"}}}
	pod2 := &v1.Pod{ObjectMeta: metav1.ObjectMeta{Name: "two", Labels: map[string]string{"foo": "bar"}}}
	pod3 := &v1.Pod{ObjectMeta: metav1.ObjectMeta{Name: "tre", Labels: map[string]string{"foo": "biz"}}}

	index.Add(pod1)
	index.Add(pod2)
	index.Add(pod3)

	keys := index.ListIndexFuncValues("testmodes")
	if len(keys) != 2 {
		t.Errorf("Expected 2 keys but got %v", len(keys))
	}

	for _, key := range keys {
		if key != "bar" && key != "biz" {
			t.Errorf("Expected only 'bar' or 'biz' but got %s", key)
		}
	}
}
