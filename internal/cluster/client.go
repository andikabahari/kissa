package cluster

import (
	"flag"
	"path/filepath"
	"sync"

	"github.com/andikabahari/kissa/config"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

var Client *kubernetes.Clientset
var doOnce sync.Once

func InitClient() {
	doOnce.Do(initClient)
}

func initClient() {
	kubeconfig, err := kubeconfig()
	if err != nil {
		panic(err)
	}

	clientset, err := kubernetes.NewForConfig(kubeconfig)
	if err != nil {
		panic(err)
	}
	Client = clientset
}

func kubeconfig() (*rest.Config, error) {
	config := config.Get()
	if config.InClusterAccess {
		return rest.InClusterConfig()
	}

	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	return clientcmd.BuildConfigFromFlags("", *kubeconfig)
}
