package main

import (
	"net/http"

	"github.com/dosedetelemetria/projeto-otel-na-pratica/internal/app"
)

func main() {
	a := app.NewUser()
	a.RegisterRoutes(http.DefaultServeMux)
	http.ListenAndServe(":8081", http.DefaultServeMux)
}
