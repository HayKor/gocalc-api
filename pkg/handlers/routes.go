package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/HayKor/gocalc-api/pkg/calculator"
	"github.com/HayKor/gocalc-api/pkg/errors"
)

type CalcRequest struct {
	Expression string `json:"expression"`
}

type CalcResponse struct {
	Result float64 `json:"result"`
}

func Calculate(w http.ResponseWriter, r *http.Request) error {
	var (
		calcRequest CalcRequest
	)

	if err := json.NewDecoder(r.Body).Decode(&calcRequest); err != nil {
		return errors.InternalServerError()
	}

	result, err := calculator.Calc(calcRequest.Expression)
	if err != nil {
		return errors.InvalidInput()
	}

	calcResponse := CalcResponse{Result: result}
	writeJSON(w, http.StatusOK, calcResponse)

	return nil
}
