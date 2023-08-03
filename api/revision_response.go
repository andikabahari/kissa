package api

import (
	"time"

	"github.com/andikabahari/kissa/knative"
)

type revisionResponse struct {
	CreatedAt       time.Time `json:"created_at"`
	Name            string    `json:"name"`
	Uid             string    `json:"uid"`
	Ready           bool      `json:"ready"`
	ActualReplicas  int       `json:"actual_replicas"`
	DesiredReplicas int       `json:"desired_replicas"`
}

func newRevisionResponse(obj knative.RevisionItem) revisionResponse {
	resp := revisionResponse{
		CreatedAt:       obj.Metadata.CreationTimestamp,
		Name:            obj.Metadata.Name,
		Uid:             obj.Metadata.Uid,
		Ready:           conditionReady(obj.Status.Conditions),
		ActualReplicas:  obj.Status.ActualReplicas,
		DesiredReplicas: obj.Status.DesiredReplicas,
	}
	return resp
}

func newRevisionResponses(obj knative.RevisionList) []revisionResponse {
	resp := make([]revisionResponse, len(obj.Items))
	for _, item := range obj.Items {
		resp = append(resp, newRevisionResponse(item))
	}
	return resp
}
