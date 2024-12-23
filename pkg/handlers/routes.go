package handlers

import (
	"encoding/json"
	"net/http"

	"log/slog"

	"github.com/HayKor/gocalc-api/pkg/calculator"
	"github.com/HayKor/gocalc-api/pkg/errors"
	"github.com/HayKor/gocalc-api/pkg/schemas"
)

func CalculateHandler(w http.ResponseWriter, r *http.Request) error {
	var (
		calcRequest schemas.CalcRequest
	)

	if r.Method != http.MethodPost {
		return errors.MethodNotAllowed()
	}

	if err := json.NewDecoder(r.Body).Decode(&calcRequest); err != nil {
		slog.Error(
			"error while decoding request",
			"error", err.Error(),
		)
		return errors.InternalServerError()
	}
	defer r.Body.Close()

	result, err := calculator.Calc(calcRequest.Expression)
	if err != nil {
		slog.Error(
			"error while calculating expression",
			"expression", calcRequest.Expression,
			"error", err.Error(),
		)
		return errors.InvalidInput()
	}

	calcResponse := schemas.CalcResponse{Result: result}
	writeJSON(w, http.StatusOK, calcResponse)

	return nil
}
