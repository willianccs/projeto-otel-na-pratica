package main

import (
	"github.com/dosedetelemetria/projeto-otel-na-pratica/internal/app/users"
)

func main() {
	app := users.NewApp()
	app.Start()
}
