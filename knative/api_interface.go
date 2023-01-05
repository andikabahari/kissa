package knative

import "k8s.io/client-go/rest"

type KnativeAPI interface {
	List(resource string) rest.Result
	Get(resource, name string) rest.Result
	Create(resource string, obj interface{}) rest.Result
}
