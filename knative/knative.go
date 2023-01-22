package knative

import (
	"k8s.io/client-go/rest"
)

type Knative interface {
	Service() *service
	Revision() *revision
}

type knative struct {
	client rest.Interface
	ns     string
}

func New(client rest.Interface, ns string) *knative {
	return &knative{
		client: client,
		ns:     ns,
	}
}

func (kn *knative) Service() *service {
	return newService(kn.client, kn.ns)
}

func (kn *knative) Revision() *revision {
	return newRevision(kn.client, kn.ns)
}
