package main

import (
	"github.com/HayKor/gocalc-api/pkg/server"
)

func main() {
	app := server.NewApp()
	app.Run()
}
