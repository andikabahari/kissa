package knative

import (
	"context"
	"fmt"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

type Knative interface {
	Client() *kubernetes.Clientset
	Namespace() string
	KnativeAPI
}

type knative struct {
	client    *kubernetes.Clientset
	namespace string
	crdPrefix string
}

func NewKnative(client *kubernetes.Clientset, namespace string) Knative {
	if namespace == "" {
		namespace = "default"
	}

	return &knative{
		client:    client,
		namespace: namespace,
		crdPrefix: fmt.Sprintf("/apis/serving.knative.dev/v1/namespaces/%s/", namespace),
	}
}

func (k *knative) Client() *kubernetes.Clientset {
	return k.client
}

func (k *knative) Namespace() string {
	return k.namespace
}

func (k *knative) List(resource string) rest.Result {
	crd := fmt.Sprintf("%s/%s", k.crdPrefix, resource)
	return k.client.RESTClient().Get().AbsPath(crd).Do(context.TODO())
}

func (k *knative) Get(resource, name string) rest.Result {
	crd := fmt.Sprintf("%s/%s/%s", k.crdPrefix, resource, name)
	return k.client.RESTClient().Get().AbsPath(crd).Do(context.TODO())
}

func (k *knative) Create(resource string, obj interface{}) rest.Result {
	crd := fmt.Sprintf("%s/%s", k.crdPrefix, resource)
	return k.client.RESTClient().Post().AbsPath(crd).Body(obj).Do(context.TODO())
}
