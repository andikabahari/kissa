package api

import (
	"strings"
	"time"

	"github.com/andikabahari/kissa/knative"
)

type serviceResponse struct {
	LastDeployed time.Time `json:"last_deployed"`
	Name         string    `json:"name"`
	Uid          string    `json:"uid"`
	Url          string    `json:"url"`
	Ready        bool      `json:"ready"`
}

func newServiceResponse(obj knative.ServiceItem) serviceResponse {
	lastDeployed := obj.Metadata.CreationTimestamp
	if updateTimestamp := obj.Metadata.Annotations["client.knative.dev/updateTimestamp"]; updateTimestamp != nil {
		lastDeployed = updateTimestamp.(time.Time)
	}

	resp := serviceResponse{
		LastDeployed: lastDeployed,
		Name:         obj.Metadata.Name,
		Uid:          obj.Metadata.Uid,
		Url:          obj.Status.Url,
		Ready:        conditionReady(obj.Status.Conditions),
	}
	return resp
}

func newServiceResponses(obj knative.ServiceList) []serviceResponse {
	resp := make([]serviceResponse, len(obj.Items))
	for _, item := range obj.Items {
		resp = append(resp, newServiceResponse(item))
	}
	return resp
}

func conditionReady(conditions []map[string]interface{}) bool {
	for _, condition := range conditions {
		if strings.EqualFold(condition["status"].(string), "false") {
			return false
		}
	}
	return true
}
