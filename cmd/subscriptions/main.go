package main

import (
	"net/http"

	"github.com/dosedetelemetria/projeto-otel-na-pratica/internal/app"
)

func main() {
	a := app.NewSubscription()
	a.RegisterRoutes(http.DefaultServeMux)
	http.ListenAndServe(":8083", http.DefaultServeMux)
}
