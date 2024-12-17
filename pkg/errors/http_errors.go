package errors

import "net/http"

func InternalServerError() APIError {
	return New(http.StatusInternalServerError, "internal server error")
}

func InvalidInput() APIError {
	return New(http.StatusUnprocessableEntity, "invalid input")
}
