package handlers

import (
	"net/http"

	"github.com/HayKor/gocalc-api/internal/server/errors"
)

// Custom wrap around http.HandlerFunc
type APIFunc func(w http.ResponseWriter, r *http.Request) error

// Make `http.HandlerFunc` out of `APIFunc`
func Make(h APIFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := h(w, r); err != nil {
			switch err := err.(type) {
			case errors.APIError:
				writeJSON(w, err.StatusCode, err)
			default:
				errResp := map[string]any{
					"statusCode": http.StatusInternalServerError,
					"message":    "internal server error",
				}
				writeJSON(w, http.StatusInternalServerError, errResp)
			}
		}
	}
}
