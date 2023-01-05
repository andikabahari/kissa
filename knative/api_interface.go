package knative

import "k8s.io/client-go/rest"

type KnativeAPI interface {
	Get(resource string) rest.Result
	Create(resource string, obj interface{}) rest.Result
}
