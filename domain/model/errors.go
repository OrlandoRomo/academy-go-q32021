package model

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ErrMissingField struct {
	field string
}

func (e ErrMissingField) Error() string {
	return fmt.Sprintf("the field %s is required", e.field)
}

type ErrInvalidData struct {
	Field string
}

func (e ErrInvalidData) Error() string {
	return fmt.Sprintf("the field %s is invalid", e.Field)
}

type ErrNotFound struct {
	Term string
}

func (e ErrNotFound) Error() string {
	return fmt.Sprintf("he are not definitions for the term %s", e.Term)
}

type ErrNotFoundInCSV struct {
	Id string
}

func (e ErrNotFoundInCSV) Error() string {
	return fmt.Sprintf("there is not definition with id %s", e.Id)
}

type ErrMissingApiKey struct{}

func (e ErrMissingApiKey) Error() string {
	return "Urban Dictionary invalid api key"
}

// EncodeError encodes the error into a json format and writing the corresponding http status
func EncodeError(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	switch err.(type) {
	case ErrNotFound, ErrNotFoundInCSV:
		w.WriteHeader(http.StatusNotFound)
	case ErrMissingApiKey:
		w.WriteHeader(http.StatusForbidden)
	case ErrMissingField, ErrInvalidData:
		w.WriteHeader(http.StatusBadRequest)
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}
	_ = json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}
