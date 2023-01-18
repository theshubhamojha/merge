package dto

import (
	"encoding/json"
	"net/http"

	"github.com/megre/merrors"
)

type APIResponse struct {
	Data      interface{}           `json:"data,omitempty"`
	Message   string                `json:"message,omitempty"`
	ErrorCode merrors.ErrorCodeType `json:"error_code,omitempty"`
}

func SendAPIResponse(w http.ResponseWriter, apiResponse APIResponse, httpCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpCode)
	data, err := json.Marshal(apiResponse)
	if err != nil {
		w.Write([]byte("error preparing response data"))
		return
	}

	w.Write(data)
}
