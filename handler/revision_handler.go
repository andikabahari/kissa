package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) ListRevision(c echo.Context) error {
	labelSelector := ""
	serviceName := c.QueryParam("service_name")
	if serviceName != "" {
		labelSelector = "serving.knative.dev/service=" + serviceName
	}

	result := h.knative.Revision().List(c.Request().Context(), labelSelector)

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

func (h *Handler) GetRevision(c echo.Context) error {
	revisionName := c.Param("revision_name")
	result := h.knative.Revision().Get(c.Request().Context(), revisionName)

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

func (h *Handler) DeleteRevision(c echo.Context) error {
	revisionName := c.Param("revision_name")

	result := h.knative.Revision().Delete(c.Request().Context(), revisionName)

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
