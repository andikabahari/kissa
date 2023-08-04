package api

import (
	"context"
	"encoding/json"
	"net/http"

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

	resp := newServiceResponses(obj)

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

	class := autoscalingClass(request.AutoscalingMetric)
	if class == "" {
		return jsonResponse(c, http.StatusBadRequest, "unsupported autoscaling metric", nil)
	}

	vars := serviceTemplateVars{
		serviceRequest:   request,
		AutoscalingClass: class,
	}
	buf, err := serviceBuf(vars)
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

	class := autoscalingClass(request.AutoscalingMetric)
	if class == "" {
		return jsonResponse(c, http.StatusBadRequest, "unsupported autoscaling metric", nil)
	}

	vars := serviceTemplateVars{
		serviceRequest:   request,
		AutoscalingClass: class,
	}
	buf, err := serviceBuf(vars)
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
