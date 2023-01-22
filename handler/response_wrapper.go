package handler

import "github.com/labstack/echo/v4"

type responseWrapper struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func jsonResponse(ctx echo.Context, code int, message string, data interface{}) error {
	response := responseWrapper{code, message, data}
	return ctx.JSON(code, response)
}
