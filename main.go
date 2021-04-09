package main

import (
	"flag"
	"fmt"
	configv1 "github.com/openshift/api/config/v1"
	"k8s.io/klog/v2"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/apiserver/pkg/server/dynamiccertificates"

	"github.com/openshift/library-go/pkg/config/helpers"
)


const (
	authenticationConfigMapNamespace = metav1.NamespaceSystem
	// authenticationConfigMapName is the name of ConfigMap in the kube-system namespace holding the root certificate
	// bundle to use to verify client certificates on incoming requests before trusting usernames in headers specified
	// by --requestheader-username-headers. This is created in the cluster by the kube-apiserver.
	// "WARNING: generally do not depend on authorization being already done for incoming requests.")
	authenticationConfigMapName = "extension-apiserver-authentication"
)


func main() {
	klog.InitFlags(flag.CommandLine)
	flag.Parse()
	klog.Info("starting the app")

	var kubeConfig string
	flag.StringVar(&kubeConfig, "kubeconfig", "", "")

	config, err := helpers.GetKubeConfigOrInClusterConfig(kubeConfig, configv1.ClientConnectionOverrides{})
	if err != nil {
		panic(err.Error())
	}
	config.Timeout = 10 * time.Second

	client, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	clientCAProvider, err := dynamiccertificates.NewDynamicCAFromConfigMapController("client-ca", authenticationConfigMapNamespace, authenticationConfigMapName, "client-ca-file", client)
	if err != nil {
		panic(fmt.Errorf("unable to load configmap based client CA file: %v", err))
	}

	var stopCh chan struct{}
	klog.Info("starting dynamic ca config map controller")
	clientCAProvider.Run(1, stopCh)
	klog.Info("done, exiting")
}

