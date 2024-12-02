package main

import (
	"github.com/dosedetelemetria/projeto-otel-na-pratica/internal/app/subscriptions"
)

func main() {
	app := subscriptions.NewApp()
	app.Start()
}
