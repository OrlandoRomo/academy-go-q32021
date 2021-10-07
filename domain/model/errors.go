package model

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type HttpError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func NewHttpError(w http.ResponseWriter, code int, message string) {
	w.WriteHeader(code)
	httpErr := &HttpError{code, message}
	json.Marshal(httpErr)
	json.NewEncoder(w).Encode(httpErr)
}

func (h *HttpError) Error() string {
	return fmt.Sprintf("code=%d, message=%s", h.Code, h.Message)
}
