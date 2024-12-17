package handlers

import (
	"net/http"

	"github.com/HayKor/gocalc-api/pkg/errors"
)

type APIFunc func(w http.ResponseWriter, r *http.Request) error

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
