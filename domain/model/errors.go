package model

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ErrInvalidData struct {
	Field string
}

func (e ErrInvalidData) Error() string {
	return fmt.Sprintf("the field `%s` is invalid", e.Field)
}

type ErrNotFound struct {
	Term string
}

func (e ErrNotFound) Error() string {
	return fmt.Sprintf("there are not definitions for the term %s", e.Term)
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

type ErrParsingDate struct {
	Date   string
	Format string
}

func (e ErrParsingDate) Error() string {
	return fmt.Sprintf("Could not parsed the date %s into the format %s", e.Date, e.Format)
}

// ErrInvalidDataType returned when we can not correctly incode struct
type ErrInvalidDataType struct {
	InvalidExpected string
}

func (e ErrInvalidDataType) Error() string {
	m := "invalid data type"
	if e.InvalidExpected != "" {
		return fmt.Sprintf("%s '%s'", m, e.InvalidExpected)
	}
	return m
}

// EncodeError encodes the error into a json format and writing the corresponding http status
func EncodeError(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	switch err.(type) {
	case ErrNotFound, ErrNotFoundInCSV:
		w.WriteHeader(http.StatusNotFound)
	case ErrMissingApiKey:
		w.WriteHeader(http.StatusForbidden)
	case ErrInvalidData, ErrInvalidDataType:
		w.WriteHeader(http.StatusBadRequest)
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}
	_ = json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}
