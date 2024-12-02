package main

import "github.com/dosedetelemetria/projeto-otel-na-pratica/internal/app/plan"

func main() {
	app := plan.NewApp()
	app.Start()
}
