package errors

import "net/http"

type RestError struct {
	Message string
	Status  int
	Error   string
}

func RestBadRequestError(message string) *RestError {
	return &RestError{
		Message: message,
		Status:  http.StatusBadRequest,
		Error:   "Bad Request",
	}
}
