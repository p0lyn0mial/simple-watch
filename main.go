package main

import (
	"context"
	"fmt"
	"k8s.io/apimachinery/pkg/api/meta"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"

    "github.com/openshift/library-go/pkg/config/helpers"
	configv1 "github.com/openshift/api/config/v1"
)

func main() {
	config, err := helpers.GetKubeConfigOrInClusterConfig("", configv1.ClientConnectionOverrides{})
	if err != nil {
		panic(err.Error())
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	for {
		watchCh, err := clientset.CoreV1().Secrets("test-01").Watch(context.TODO(), metav1.ListOptions{})
		if err != nil {
			panic(err.Error())
		}

		watchStart := time.Now()
		fmt.Println("starting watching Secrets in test-01 ns")
		for {
			event, ok := <-watchCh.ResultChan()
			if !ok {
				fmt.Println("watch closed")
				break
			}
			objMeta, err := meta.Accessor(event.Object)
			if err != nil {
				fmt.Println(fmt.Sprintf("got %v event, unable to get meta, err %v", event.Type, err))
				continue
			}
			fmt.Println(fmt.Sprintf("got %v event: object name %v", event.Type, objMeta.GetName()))
		}
		fmt.Println(fmt.Sprintf("watch ended, took %v", time.Now().Sub(watchStart)))
	}
}

