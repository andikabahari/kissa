package api

import (
	"strings"
	"time"

	"github.com/andikabahari/kissa/knative"
)

type serviceResponse struct {
	LastDeployed time.Time `json:"last_deployed,omitempty"`
	Name         string    `json:"name,omitempty"`
	Uid          string    `json:"uid,omitempty"`
	Url          string    `json:"url,omitempty"`
	Ready        bool      `json:"ready,omitempty"`
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

func newServicesResponse(obj knative.ServiceList) []serviceResponse {
	resp := make([]serviceResponse, 0)
	for _, item := range obj.Items {
		resp = append(resp, newServiceResponse(item))
	}
	return resp
}

func conditionReady(conditions []map[string]interface{}) bool {
	for _, condition := range conditions {
		if strings.EqualFold(condition["status"].(string), "true") {
			return true
		}
	}
	return false
}
