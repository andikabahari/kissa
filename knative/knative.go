package knative

import (
	"context"
	"encoding/json"
	"fmt"

	"k8s.io/client-go/kubernetes"
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

func mapByte(raw []byte) (map[string]interface{}, error) {
	data := make(map[string]interface{})
	if err := json.Unmarshal(raw, &data); err != nil {
		return nil, err
	}

	return data, nil
}

func (k *knative) List(resource string) ([]byte, error) {
	crd := fmt.Sprintf("%s/%s", k.crdPrefix, resource)
	return k.client.RESTClient().Get().AbsPath(crd).DoRaw(context.TODO())
}

func (k *knative) ListMap(resource string) (map[string]interface{}, error) {
	raw, err := k.List(resource)
	if err != nil {
		return nil, err
	}

	return mapByte(raw)
}

func (k *knative) Get(resource, name string) ([]byte, error) {
	crd := fmt.Sprintf("%s/%s/%s", k.crdPrefix, resource, name)
	return k.client.RESTClient().Get().AbsPath(crd).DoRaw(context.TODO())
}

func (k *knative) GetMap(resource, name string) (map[string]interface{}, error) {
	raw, err := k.Get(resource, name)
	if err != nil {
		return nil, err
	}

	return mapByte(raw)
}

func (k *knative) Create(resource string, obj interface{}) ([]byte, error) {
	crd := fmt.Sprintf("%s/%s", k.crdPrefix, resource)
	return k.client.RESTClient().Post().AbsPath(crd).Body(obj).DoRaw(context.TODO())
}

func (k *knative) CreateMap(resource string, obj interface{}) (map[string]interface{}, error) {
	raw, err := k.Create(resource, obj)
	if err != nil {
		return nil, err
	}

	return mapByte(raw)
}
