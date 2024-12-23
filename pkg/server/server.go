package server

import (
	"net/http"

	"github.com/HayKor/gocalc-api/internal/server/handlers"
)

type App struct {
	mux *http.ServeMux
}

func NewApp() *App {
	return &App{
		mux: http.NewServeMux(),
	}
}

func (a *App) Run() error {
	handlers.RegisterRoutes(a.mux)
	return http.ListenAndServe(":8080", a.mux)
}
