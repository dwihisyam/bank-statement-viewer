package response

import (
	"encoding/json"
	"net/http"
)

type ApiResponse struct {
	Status  string      `json:"status"` // "success" or "error"
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func JSON(w http.ResponseWriter, code int, resp ApiResponse) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(resp)
}

func NewError(msg string) ApiResponse {
	return ApiResponse{
		Status:  "error",
		Message: msg,
	}
}

func NewData(d interface{}) ApiResponse {
	return ApiResponse{
		Status: "success",
		Data:   d,
	}
}
