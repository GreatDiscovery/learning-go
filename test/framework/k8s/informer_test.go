package k8s

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/tools/cache"
	"log"
	"testing"
	"time"
)

//https://github.com/luozhiyun993/luozhiyun-SourceLearn/blob/master/%E6%B7%B1%E5%85%A5k8s/16.%E6%B7%B1%E5%85%A5k8s%EF%BC%9AInformer%E4%BD%BF%E7%94%A8%E5%8F%8A%E5%85%B6%E6%BA%90%E7%A0%81%E5%88%86%E6%9E%90.md

func TestInformer(t *testing.T) {
	clientSet, _ := InitClient()

	neverStopCh := make(chan struct{})
	defer close(neverStopCh)
	//表示每分钟进行一次resync，resync会周期性地执行List操作
	sharedInformers := informers.NewSharedInformerFactory(clientSet, time.Minute)

	informer := sharedInformers.Core().V1().Pods().Informer()

	_, err := informer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			mObj := obj.(v1.Object)
			log.Printf("New Pod Added to Store: %s", mObj.GetName())
		},
		UpdateFunc: func(oldObj, newObj interface{}) {
			oObj := oldObj.(v1.Object)
			nObj := newObj.(v1.Object)
			log.Printf("%s Pod Updated to %s", oObj.GetName(), nObj.GetName())
		},
		DeleteFunc: func(obj interface{}) {
			mObj := obj.(v1.Object)
			log.Printf("Pod Deleted from Store: %s", mObj.GetName())
		},
	})
	if err != nil {
		return
	}
	informer.Run(neverStopCh)
}
