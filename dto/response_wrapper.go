package dto

import "github.com/labstack/echo/v4"

type ResponseWrapper struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func JSONResponse(ctx echo.Context, code int, message string, data interface{}) error {
	response := ResponseWrapper{code, message, data}
	return ctx.JSON(code, response)
}
