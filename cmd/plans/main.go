package main

import (
	"net/http"

	"github.com/dosedetelemetria/projeto-otel-na-pratica/internal/app"
)

func main() {
	a := app.NewPlan()
	a.RegisterRoutes(http.DefaultServeMux)
	http.ListenAndServe(":8082", http.DefaultServeMux)
}
