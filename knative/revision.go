package knative

import (
	"context"
	"fmt"

	"k8s.io/client-go/rest"
)

type revision struct {
	client     rest.Interface
	ns         string
	pathPrefix string
}

func newRevision(client rest.Interface, ns string) *revision {
	return &revision{
		client:     client,
		ns:         ns,
		pathPrefix: fmt.Sprintf("/apis/serving.knative.dev/v1/namespaces/%s", ns),
	}
}

func (s *revision) Get(ctx context.Context, name string) rest.Result {
	return s.client.
		Get().
		AbsPath(s.pathPrefix, "revisions", name).
		Do(ctx)
}

func (s *revision) List(ctx context.Context, labelSelector string) rest.Result {
	client := s.client.Get()

	if labelSelector != "" {
		client.Param("labelSelector", labelSelector)
	}

	return client.AbsPath(s.pathPrefix, "revisions").Do(ctx)
}

func (s *revision) Delete(ctx context.Context, name string) rest.Result {
	return s.client.
		Delete().
		AbsPath(s.pathPrefix, "revisions", name).
		Do(ctx)
}
