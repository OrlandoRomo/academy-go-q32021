package model

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/OrlandoRomo/academy-go-q32021/infrastructure/middleware"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestEncodeError(t *testing.T) {
	tests := []struct {
		name        string
		err         error
		code        int
		contentType string
	}{
		{"not found error type", ErrNotFound{}, http.StatusNotFound, "application/json; charset=utf-8"},
		{"not found in CSV error type", ErrNotFoundInCSV{}, http.StatusNotFound, "application/json; charset=utf-8"},
		{"missing api key error type", ErrMissingApiKey{}, http.StatusForbidden, "application/json; charset=utf-8"},
		{"missing field error type", ErrMissingField{}, http.StatusBadRequest, "application/json; charset=utf-8"},
		{"invalid data error type", ErrInvalidData{}, http.StatusBadRequest, "application/json; charset=utf-8"},
		{"random error type", errors.New("random unknown error"), http.StatusInternalServerError, "application/json; charset=utf-8"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			r := mux.NewRouter()
			r.Use(middleware.HeadersMiddleware)
			r.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {})
			req := httptest.NewRequest("GET", "/", nil)
			r.ServeHTTP(w, req)
			EncodeError(w, test.err)
			assert.Equal(t, test.code, w.Code)
			assert.Equal(t, test.contentType, w.Header().Get("Content-Type"))
		})
	}
}
