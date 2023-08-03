package api

import "github.com/labstack/echo/v4"

type responseWrapper struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func jsonResponse(ctx echo.Context, code int, message string, data interface{}) error {
	response := responseWrapper{message, data}
	return ctx.JSON(code, response)
}
