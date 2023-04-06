package handler

import (
	"bytes"
	"context"
	"net/http"
	"text/template"

	"github.com/andikabahari/kissa/knative"
	"github.com/labstack/echo/v4"
)

const serviceTemplate = `{
	"apiVersion": "serving.knative.dev/v1",
	"kind": "Service",
	"metadata": {
		"name": "{{.Name}}"
	},
	"spec": {
		"template": {
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

type serviceRequest struct {
	knative.ServiceObject
}

func (h *Handler) ListService(c echo.Context) error {
	result := h.knative.Service().List(context.Background())

	var code int
	result.StatusCode(&code)
	if code >= 400 {
		return jsonResponse(c, code, result.Error().Error(), nil)
	}

	data, err := mapK8sResult(result)
	if err != nil {
		return err
	}

	return jsonResponse(c, http.StatusOK, "success", data["items"])
}

func (h *Handler) GetService(c echo.Context) error {
	serviceName := c.Param("service_name")
	result := h.knative.Service().Get(context.Background(), serviceName)

	var code int
	result.StatusCode(&code)
	if code >= 400 {
		return jsonResponse(c, code, result.Error().Error(), nil)
	}

	data, err := mapK8sResult(result)
	if err != nil {
		return err
	}

	return jsonResponse(c, http.StatusOK, "success", data)
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

	data, err := mapK8sResult(result)
	if err != nil {
		return err
	}

	return jsonResponse(c, http.StatusOK, "success", data)
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

	result := h.knative.Service().Update(context.Background(), buf.Bytes())

	code := 0
	result.StatusCode(&code)
	if code >= 400 {
		return jsonResponse(c, code, result.Error().Error(), nil)
	}

	data, err := mapK8sResult(result)
	if err != nil {
		return err
	}

	return jsonResponse(c, http.StatusOK, "success", data)
}

func (h *Handler) DeleteService(c echo.Context) error {
	serviceName := c.Param("service_name")

	result := h.knative.Service().Delete(context.Background(), serviceName)

	code := 0
	result.StatusCode(&code)
	if code >= 400 {
		return jsonResponse(c, code, result.Error().Error(), nil)
	}

	data, err := mapK8sResult(result)
	if err != nil {
		return err
	}

	return jsonResponse(c, http.StatusOK, "success", data)
}

func serviceBuf(request serviceRequest) (*bytes.Buffer, error) {
	tmpl, err := template.New("service").Parse(serviceTemplate)
	if err != nil {
		return nil, err
	}

	obj := knative.ServiceObject{
		Name:          request.Name,
		Image:         request.Image,
		ContainerPort: request.ContainerPort,
		Env:           request.Env,
	}
	buf := new(bytes.Buffer)
	if err := tmpl.Execute(buf, obj); err != nil {
		return nil, err
	}

	return buf, nil
}
