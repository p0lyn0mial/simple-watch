package main

import (
	"flag"
	"time"

	kubeinformers "k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/klog/v2"

	"github.com/openshift/library-go/pkg/config/helpers"
	configv1 "github.com/openshift/api/config/v1"
)


func main() {
	var kubeConfig string
	var stopCh chan struct{}

	klog.InitFlags(flag.CommandLine)
	flag.StringVar(&kubeConfig, "kubeconfig", "", "")
	flag.Parse()

	klog.Info("starting the app")

	config, err := helpers.GetKubeConfigOrInClusterConfig(kubeConfig, configv1.ClientConnectionOverrides{})
	if err != nil {
		panic(err.Error())
	}
	config.Timeout = 10 * time.Second

	client, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	klog.Info("starting configmap informer")
	kubeInformerFactory := kubeinformers.NewSharedInformerFactory(client, time.Second*30)
	cmInformer := kubeInformerFactory.Core().V1().ConfigMaps().Informer()
	cmInformer.Run(stopCh)
	klog.Info("done, exiting")
}

