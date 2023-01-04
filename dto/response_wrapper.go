package dto

import (
	"encoding/json"
	"net/http"
)

type ResponseWrapper struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func JSONResponse(w http.ResponseWriter, code int, message string, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	res := ResponseWrapper{
		Code:    code,
		Message: message,
		Data:    data,
	}
	json.NewEncoder(w).Encode(res)
}
