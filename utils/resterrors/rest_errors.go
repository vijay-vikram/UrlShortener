package resterrors

import "net/http"

//RestError ...
type RestError struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Title   string `json:"title"`
}

//NewBadRequestError ...
func NewBadRequestError(message string) *RestError {
	return &RestError{
		Message: message,
		Status:  http.StatusBadRequest,
		Title:   "bad_request",
	}
}

//NewInternalServerError ...
func NewInternalServerError(message string) *RestError {
	return &RestError{
		Message: message,
		Status:  http.StatusInternalServerError,
		Title:   "internal_server_error",
	}
}

//NewNotFoundError ...
func NewNotFoundError(message string) *RestError {
	return &RestError{
		Message: message,
		Status:  http.StatusNotFound,
		Title:   "not_found",
	}
}
