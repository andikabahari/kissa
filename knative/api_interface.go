package knative

import "k8s.io/client-go/rest"

type KnativeAPI interface {
	Get(resource, labelSelector string) rest.Result
	Create(resource string, obj interface{}) rest.Result
	Update(resource string, obj interface{}) rest.Result
	Delete(resource string) rest.Result
}
