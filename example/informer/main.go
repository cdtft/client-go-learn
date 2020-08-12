package main

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
	"log"
	"time"
)

func main() {
	config, err := clientcmd.BuildConfigFromFlags("", "/Users/wangcheng/GolandProjects/kubeconfig/config")
	if err != nil {
		panic(err)
	}
	//informer通过clientset和APIServer进行交互
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	// informer是一个持久运行的goroutine，进程退出之前通知informer退出
	stopCh := make(chan struct{})
	defer close(stopCh)

	sharedInformers := informers.NewSharedInformerFactory(clientset, time.Minute)

	//得到Pod的informer
	informer := sharedInformers.Core().V1().Pods().Informer()
	//一般情况informer通过回调方法将将资源对象推送到WorkQueue
	//informer机制很容易的就可以监控我们关心的资源变化
	informer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			mObj := obj.(v1.Object)
			log.Printf("New pod added to store: %s\n", mObj.GetName())
		},
		UpdateFunc: func(oldObj, newObj interface{}) {
			mOldObj := oldObj.(v1.Object)
			mNewObj := newObj.(v1.Object)
			log.Printf("old pod:%s update to new pod %s\n", mOldObj.GetName(), mNewObj.GetName())
		},
		DeleteFunc: func(obj interface{}) {
			mObj := obj.(v1.Object)
			log.Printf("Pod deleted from stroe: %s", mObj.GetName())
		},
	})

	informer.Run(stopCh)
}
