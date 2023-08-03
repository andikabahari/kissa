package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"strings"
	"text/template"
	"time"

	"github.com/andikabahari/kissa/knative"
	"github.com/labstack/echo/v4"
)

type serviceRequest struct {
	Name          string `json:"name"`
	Image         string `json:"image"`
	ContainerPort int    `json:"container_port"`
	Env           []struct {
		Name  string `json:"name"`
		Value string `json:"value"`
	} `json:"env"`
	AutoscalingMetric string `json:"autoscaling_metric"`
	AutoscalingTarget int    `json:"autoscaling_target"`
	MaxScale          int    `json:"max_scale"`
	MinScale          int    `json:"min_scale"`
}

const serviceTemplate = `{
	"apiVersion": "serving.knative.dev/v1",
	"kind": "Service",
	"metadata": {
		"name": "{{.Name}}"
	},
	"spec": {
		"template": {
			"metadata": {
				"annotations": {
					"autoscaling.knative.dev/metric": "{{.AutoscalingMetric}}",
					"autoscaling.knative.dev/target": "{{.AutoscalingTarget}}",
					"autoscaling.knative.dev/min-scale": "{{.MinScale}}",
					"autoscaling.knative.dev/max-scale": "{{.MaxScale}}"
				}
			},
			"spec": {
				"containers": [
					{
						"image": "{{.Image}}",
						"ports": [
							{
								"containerPort": {{.ContainerPort}}
							}
						]
						{{if .Env}}
						,
						"env": [
							{{$env := .Env}}
							{{range $idx, $elem := .Env}}
							{{if $idx}},{{end}}
							{"name":"{{$elem.Name}}","value":"{{$elem.Value}}"}
							{{end}}
						]
						{{end}}
					}
				]
			}
		}
	}
}`

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

func (h *Handler) ListService(c echo.Context) error {
	result := h.knative.Service().List(context.Background())

	var code int
	result.StatusCode(&code)
	if code >= 400 {
		return jsonResponse(c, code, result.Error().Error(), nil)
	}

	raw, err := result.Raw()
	if err != nil {
		return err
	}

	obj := knative.ServiceList{}
	if err := json.Unmarshal(raw, &obj); err != nil {
		return err
	}

	resp := newServicesResponse(obj)

	return jsonResponse(c, http.StatusOK, "success", resp)
}

func (h *Handler) GetService(c echo.Context) error {
	serviceName := c.Param("service_name")
	result := h.knative.Service().Get(context.Background(), serviceName)

	var code int
	result.StatusCode(&code)
	if code >= 400 {
		return jsonResponse(c, code, result.Error().Error(), nil)
	}

	raw, err := result.Raw()
	if err != nil {
		return err
	}

	obj := knative.ServiceItem{}
	if err := json.Unmarshal(raw, &obj); err != nil {
		return err
	}

	resp := newServiceResponse(obj)

	return jsonResponse(c, http.StatusOK, "success", resp)
}

func (h *Handler) CreateService(c echo.Context) error {
	request := serviceRequest{}

	if err := c.Bind(&request); err != nil {
		return err
	}

	buf, err := serviceBuf(request)
	if err != nil {
		return err
	}

	result := h.knative.Service().Create(context.Background(), buf.Bytes())

	var code int
	result.StatusCode(&code)
	if code >= 400 {
		return jsonResponse(c, code, result.Error().Error(), nil)
	}

	raw, err := result.Raw()
	if err != nil {
		return err
	}

	obj := knative.ServiceItem{}
	if err := json.Unmarshal(raw, &obj); err != nil {
		return err
	}

	resp := newServiceResponse(obj)

	return jsonResponse(c, http.StatusOK, "success", resp)
}

func (h *Handler) UpdateService(c echo.Context) error {
	request := serviceRequest{}

	if err := c.Bind(&request); err != nil {
		return err
	}

	serviceName := c.Param("service_name")
	request.Name = serviceName

	buf, err := serviceBuf(request)
	if err != nil {
		return err
	}

	result := h.knative.Service().Update(context.Background(), serviceName, buf.Bytes())

	code := 0
	result.StatusCode(&code)
	if code >= 400 {
		return jsonResponse(c, code, result.Error().Error(), nil)
	}

	raw, err := result.Raw()
	if err != nil {
		return err
	}

	obj := knative.ServiceItem{}
	if err := json.Unmarshal(raw, &obj); err != nil {
		return err
	}

	resp := newServiceResponse(obj)

	return jsonResponse(c, http.StatusOK, "success", resp)
}

func (h *Handler) DeleteService(c echo.Context) error {
	serviceName := c.Param("service_name")

	result := h.knative.Service().Delete(context.Background(), serviceName)

	code := 0
	result.StatusCode(&code)
	if code >= 400 {
		return jsonResponse(c, code, result.Error().Error(), nil)
	}

	return jsonResponse(c, http.StatusOK, "success", nil)
}

func serviceBuf(request serviceRequest) (*bytes.Buffer, error) {
	tmpl, err := template.New("service").Parse(serviceTemplate)
	if err != nil {
		return nil, err
	}

	buf := new(bytes.Buffer)
	if err := tmpl.Execute(buf, request); err != nil {
		return nil, err
	}

	return buf, nil
}
