package api

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"text/template"

	"github.com/andikabahari/kissa/knative"
	"github.com/labstack/echo/v4"
)

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
