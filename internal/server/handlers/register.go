package handlers

import "net/http"

func RegisterRoutes(m *http.ServeMux) {
	m.HandleFunc("/api/v1/calculate", Make(CalculateHandler))
}
