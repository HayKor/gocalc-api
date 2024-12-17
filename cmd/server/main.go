package main

import (
	"net/http"

	"github.com/HayKor/gocalc-api/pkg/handlers"
)

func main() {
	mux := http.NewServeMux()
	handlers.RegisterRoutes(mux)
	http.ListenAndServe(":8080", mux)
}
