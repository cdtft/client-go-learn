package main

import (
	"fmt"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/cache"
	"strings"
)

func UserIndexFunc(obj interface{}) ([]string, error) {
	pod := obj.(*v1.Pod)
	userString := pod.Annotations["user"]
	return strings.Split(userString, ","), nil
}

func main() {
	index := cache.NewIndexer(cache.MetaNamespaceKeyFunc, cache.Indexers{"byUser": UserIndexFunc})

	pod1 := &v1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name: "one",
			Annotations: map[string]string{
				"user": "cdtft,wangcheng",
			},
		},
	}
	pod2 := &v1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name: "tow",
			Annotations: map[string]string{
				"user": "cdtft,bert",
			},
		},
	}
	pod3 := &v1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name: "three",
			Annotations: map[string]string{
				"user": "cdtft,zs",
			},
		},
	}

	index.Add(pod1)
	index.Add(pod2)
	index.Add(pod3)

	cdtftPods, err := index.ByIndex("byUser", "cdtft")
	if err != nil {
		panic(err)
	}

	for _, cdtftPod := range cdtftPods {
		fmt.Printf("cdtft pod name is %s\n", cdtftPod.(*v1.Pod).Name)
	}
}
