package main

import (
	"net/http"

	"github.com/dosedetelemetria/projeto-otel-na-pratica/internal/app"
)

func main() {
	a, err := app.NewPayment()
	if err != nil {
		panic(err)
	}
	a.RegisterRoutes(http.DefaultServeMux)
	http.ListenAndServe(":8084", http.DefaultServeMux)
}
