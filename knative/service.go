package knative

import (
	"context"
	"fmt"

	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/rest"
)

type service struct {
	client     rest.Interface
	ns         string
	pathPrefix string
}

func newService(client rest.Interface, ns string) *service {
	return &service{
		client:     client,
		ns:         ns,
		pathPrefix: fmt.Sprintf("/apis/serving.knative.dev/v1/namespaces/%s", ns),
	}
}

func (s *service) Get(ctx context.Context, name string) rest.Result {
	return s.client.
		Get().
		AbsPath(s.pathPrefix, "services", name).
		Do(ctx)
}

func (s *service) List(ctx context.Context) rest.Result {
	return s.client.
		Get().
		AbsPath(s.pathPrefix, "services").
		Do(ctx)
}

func (s *service) Create(ctx context.Context, obj interface{}) rest.Result {
	return s.client.
		Post().
		AbsPath(s.pathPrefix, "services").
		Body(obj).
		Do(ctx)
}

func (s *service) Update(ctx context.Context, name string, obj interface{}) rest.Result {
	return s.client.
		Patch(types.MergePatchType).
		AbsPath(s.pathPrefix, "services", name).
		Body(obj).
		Do(ctx)
}

func (s *service) Delete(ctx context.Context, name string) rest.Result {
	return s.client.
		Delete().
		AbsPath(s.pathPrefix, "services", name).
		Do(ctx)
}
