package handler

import (
	"net/http"

	"github.com/andikabahari/kissa/dto"
	"github.com/andikabahari/kissa/knative"
	"github.com/andikabahari/kissa/util"
	"github.com/labstack/echo/v4"
)

type ServiceHandler interface {
	List(c echo.Context) error
	Get(c echo.Context) error
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
	result := h.knative.Get("services", "")

	var code int
	result.StatusCode(&code)
	if code >= 400 {
		return dto.JSONResponse(c, code, result.Error().Error(), nil)
	}

	data, err := util.MapK8sResult(result)
	if err != nil {
		return err
	}

	return dto.JSONResponse(c, http.StatusOK, "success", data["items"])
}

func (h *serviceHandler) Get(c echo.Context) error {
	serviceName := c.Param("service_name")
	result := h.knative.Get("services/"+serviceName, "")

	var code int
	result.StatusCode(&code)
	if code >= 400 {
		return dto.JSONResponse(c, code, result.Error().Error(), nil)
	}

	data, err := util.MapK8sResult(result)
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

	buf, err := util.ServiceBuf(request)
	if err != nil {
		return err
	}

	result := h.knative.Create("services", buf.Bytes())

	var code int
	result.StatusCode(&code)
	if code >= 400 {
		return dto.JSONResponse(c, code, result.Error().Error(), nil)
	}

	data, err := util.MapK8sResult(result)
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

	buf, err := util.ServiceBuf(request)
	if err != nil {
		return err
	}

	result := h.knative.Update("services/"+serviceName, buf.Bytes())

	code := 0
	result.StatusCode(&code)
	if code >= 400 {
		return dto.JSONResponse(c, code, result.Error().Error(), nil)
	}

	data, err := util.MapK8sResult(result)
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

	data, err := util.MapK8sResult(result)
	if err != nil {
		return err
	}

	return dto.JSONResponse(c, http.StatusOK, "success", data)
}
