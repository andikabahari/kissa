package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"text/template"

	"github.com/andikabahari/kissa/constants"
	"github.com/andikabahari/kissa/dto"
	"github.com/andikabahari/kissa/knative"
	"github.com/labstack/echo/v4"
	"k8s.io/client-go/rest"
)

type ServiceHandler interface {
	List(c echo.Context) error
	Create(c echo.Context) error
	Update(c echo.Context) error
	Delete(c echo.Context) error
}

type serviceHandler struct {
	knative knative.Knative
}

func NewServiceHandler(kn knative.Knative) ServiceHandler {
	return &serviceHandler{
		knative: kn,
	}
}

func (h *serviceHandler) List(c echo.Context) error {
	resource := "services"

	serviceName := c.QueryParam("service_name")
	if serviceName != "" {
		resource += "/" + serviceName
	}

	result := h.knative.Get(resource)

	var code int
	result.StatusCode(&code)
	if code >= 400 {
		return dto.JSONResponse(c, code, result.Error().Error(), nil)
	}

	data, err := mapK8sResult(result)
	if err != nil {
		return err
	}

	return dto.JSONResponse(c, http.StatusOK, "success", data)
}

func (h *serviceHandler) Create(c echo.Context) error {
	request := dto.ServiceRequest{}

	if err := c.Bind(&request); err != nil {
		return err
	}

	buf, err := serviceBuf(request)
	if err != nil {
		return err
	}

	result := h.knative.Create("services", buf.Bytes())

	var code int
	result.StatusCode(&code)
	if code >= 400 {
		return dto.JSONResponse(c, code, result.Error().Error(), nil)
	}

	data, err := mapK8sResult(result)
	if err != nil {
		return err
	}

	return dto.JSONResponse(c, http.StatusOK, "success", data)
}

func (h *serviceHandler) Update(c echo.Context) error {
	request := dto.ServiceRequest{}

	if err := c.Bind(&request); err != nil {
		return err
	}

	serviceName := c.Param("service_name")
	request.Name = serviceName

	buf, err := serviceBuf(request)
	if err != nil {
		return err
	}

	result := h.knative.Update("services/"+serviceName, buf.Bytes())

	code := 0
	result.StatusCode(&code)
	if code >= 400 {
		return dto.JSONResponse(c, code, result.Error().Error(), nil)
	}

	data, err := mapK8sResult(result)
	if err != nil {
		return err
	}

	return dto.JSONResponse(c, http.StatusOK, "success", data)
}

func (h *serviceHandler) Delete(c echo.Context) error {
	serviceName := c.Param("service_name")

	result := h.knative.Delete("services/" + serviceName)

	code := 0
	result.StatusCode(&code)
	if code >= 400 {
		return dto.JSONResponse(c, code, result.Error().Error(), nil)
	}

	data, err := mapK8sResult(result)
	if err != nil {
		return err
	}

	return dto.JSONResponse(c, http.StatusOK, "success", data)
}

func mapK8sResult(result rest.Result) (map[string]interface{}, error) {
	raw, err := result.Raw()
	if err != nil {
		return nil, err
	}

	data := make(map[string]interface{})
	if err := json.Unmarshal(raw, &data); err != nil {
		return nil, err
	}

	return data, nil
}

func serviceBuf(request dto.ServiceRequest) (*bytes.Buffer, error) {
	tmpl, err := template.New("service").Parse(constants.KnativeServiceTemplate)
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
