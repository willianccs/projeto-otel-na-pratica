package main

import (
	"net/http"

	"github.com/dosedetelemetria/projeto-otel-na-pratica/internal/app"
)

func main() {
	a := app.NewPayment()
	a.RegisterRoutes(http.DefaultServeMux)
	http.ListenAndServe(":8084", http.DefaultServeMux)
}
