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

func (k *knative) Get(resource string) rest.Result {
	return k.client.
		RESTClient().
		Get().
		AbsPath(k.crdPrefix + resource).
		Do(context.TODO())
}

func (k *knative) Create(resource string, obj interface{}) rest.Result {
	return k.client.
		RESTClient().
		Post().
		AbsPath(k.crdPrefix + resource).
		Body(obj).Do(context.TODO())
}
