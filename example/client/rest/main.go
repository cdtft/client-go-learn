package main

import (
	"fmt"
	corev1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	config, err := clientcmd.BuildConfigFromFlags("",
		"/Users/wangcheng/GolandProjects/kubeconfig/config")
	if err != nil {
		panic(err)
	}

	config.APIPath = "apis"
	// v1
	config.GroupVersion = &corev1.SchemeGroupVersion
	config.NegotiatedSerializer = scheme.Codecs

	restClient, err := rest.RESTClientFor(config)
	if err != nil {
		panic(err)
	}

	deployments := &corev1.DeploymentList{}

	err = restClient.Get().Namespace("default").Resource("deployments").
		VersionedParams(&metav1.ListOptions{Limit:500}, scheme.ParameterCodec).
		Do().
		Into(deployments)
	if err != nil {
		panic(err)
	}

	for _, dp := range deployments.Items {
		fmt.Println(dp.Name)
	}

}
