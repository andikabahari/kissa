package handler

import (
	"net/http"

	"github.com/andikabahari/kissa/dto"
	"github.com/andikabahari/kissa/knative"
	"github.com/andikabahari/kissa/util"
	"github.com/labstack/echo/v4"
)

type RevisionHandler interface {
	List(c echo.Context) error
	Get(c echo.Context) error
	Delete(c echo.Context) error
}

type revisionHandler struct {
	knative knative.Knative
}

func NewRevisionHandler(kn knative.Knative) RevisionHandler {
	return &revisionHandler{
		knative: kn,
	}
}

func (h *revisionHandler) List(c echo.Context) error {
	labelSelector := ""
	serviceName := c.QueryParam("service_name")
	if serviceName != "" {
		labelSelector = "serving.knative.dev/service=" + serviceName
	}

	result := h.knative.Get("revisions", labelSelector)

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

func (h *revisionHandler) Get(c echo.Context) error {
	revisionName := c.Param("revision_name")
	result := h.knative.Get("revisions/"+revisionName, "")

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

func (h *revisionHandler) Delete(c echo.Context) error {
	revisionName := c.Param("revision_name")

	result := h.knative.Delete("revisions/" + revisionName)

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
