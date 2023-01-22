package main

import (
	"os"
	"strconv"

	"github.com/andikabahari/kissa/knative"
	"github.com/andikabahari/kissa/server"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	kcfg, err := clientcmd.BuildConfigFromFlags("", os.Getenv("KUBECONFIG_PATH"))
	if err != nil {
		panic(err)
	}

	kcl, err := kubernetes.NewForConfig(kcfg)
	if err != nil {
		panic(err)
	}

	ns, ok := os.LookupEnv("NAMESPACE")
	if !ok {
		ns = "default"
	}
	kn := knative.New(kcl.RESTClient(), ns)

	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		panic(err)
	}

	srv := server.New(kn)
	srv.Run(port)
}
