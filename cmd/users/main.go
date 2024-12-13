// Copyright Dose de Telemetria GmbH
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/dosedetelemetria/projeto-otel-na-pratica/internal/app"
	"github.com/dosedetelemetria/projeto-otel-na-pratica/internal/config"
)

func main() {
	configFlag := flag.String("config", "", "path to the config file")
	flag.Parse()

	c, err := config.LoadConfig(*configFlag)
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	a := app.NewUser(&c.Users)
	a.RegisterRoutes(http.DefaultServeMux)
	http.ListenAndServe(":8081", http.DefaultServeMux)
}
